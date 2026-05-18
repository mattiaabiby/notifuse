# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

### Backend

```bash
# Build (CGO_ENABLED=1 is required for V8 JS runtime)
CGO_ENABLED=1 go build -o bin/server ./cmd/api
# or
make build

# Run (from source)
make run

# Dev with hot reload (Air)
make dev

# Run a single test
go test -race ./internal/service/... -run TestFunctionName

# Run ALL unit tests — use this, NOT make test-unit (verbose output crashes agents)
make test-agent

# Run tests by layer (safe to use -v individually)
make test-domain
make test-service
make test-repo
make test-http
make test-migrations
make test-database
make test-pkg

# Coverage report (HTML)
make coverage

# Integration tests (requires test DB via docker)
cd tests && docker compose -f compose.test.yaml up -d
make test-integration
```

### Frontend — Console (`console/`)

```bash
cd console
npm run dev             # Dev server
npm run build           # lingui compile + tsc + vite build
npm run lint            # ESLint (0 warnings allowed)
npm run lingui:extract  # Extract i18n strings to .po files
npm run lingui:compile  # Compile .po files (required before build)
npm test                # Vitest
```

### Frontend — Notification Center (`notification_center/`)

```bash
cd notification_center
npm run dev
npm run build
npm test
```

### Other

```bash
make keygen             # Generate secret keys
make openapi-bundle     # Bundle openapi/openapi.yaml → openapi.json
make docker-compose-up  # Start dev stack (Postgres + app)
```

## Architecture

**Entry point**: `cmd/api/main.go` → `internal/app/app.go` is the composition root (wires all repositories, services, and HTTP handlers).

**Clean Architecture layers** (all under `internal/`):
- `domain/` — entities, interfaces (Repository/Service), value objects, business rules. This is the source of truth for types.
- `service/` — business logic implementations; `service/broadcast/` is a sub-package for the broadcast orchestrator. Mock files live here alongside their service files (`mock_*.go`).
- `repository/` — PostgreSQL implementations using Squirrel query builder.
- `http/` — HTTP handlers (one file per domain, e.g. `contact_handler.go`). All handlers are tested in `*_handler_test.go`.
- `app/` — composition root only; no business logic here.
- `migrations/` — version-based schema migrations (`v1.go`, `v2.go`, …).

**Multi-tenant**: each workspace gets its own PostgreSQL database. `internal/app/app.go` manages per-workspace DB connections. System DB holds workspaces + users; workspace DBs hold everything else.

**API style**: RPC dot notation — `POST /api/contact.create`, `GET /api/contact.list`, etc.

**Module path**: `github.com/Notifuse/notifuse`

**Current version**: stored in `config/config.go` as `VERSION = "30.0"`. Major version increments trigger a new migration file.

## Mocks

- Domain interface mocks: `internal/domain/mocks/`
- Service-layer mocks: `internal/service/mock_*.go` (co-located)
- Package mocks: `pkg/mocks/`
- Generate with `//go:generate mockgen`

## Key Conventions

### Testing
- **Never test logging or tracing** — they are cross-cutting implementation details.
- Preferred assertions: `assert.NoError`, `require.NoError`, `mock.Anything`, `mock.AnythingOfType`.
- Database tests use `go-sqlmock`; service tests use GoMock.
- Keep test files in the same package as the code under test.

### Tracing (OpenCensus)
- Wrap tracing calls with codecov ignore comments so they don't affect coverage:
  ```go
  // codecov:ignore:start
  ctx, span := tracing.StartServiceSpan(ctx, "ServiceName", "MethodName")
  defer tracing.EndSpan(span, err)
  // codecov:ignore:end
  ```
- Pattern: `return tracing.TraceMethod(ctx, "Service", "Method", func(ctx context.Context) error { ... })`

### Error handling
- Wrap errors: `fmt.Errorf("context: %w", err)`
- Return errors directly; log at the handler layer.

### Imports
- UUID: `import uuid "github.com/google/uuid"`
- Logging: `import zerolog "github.com/rs/zerolog"` (but use the project logger pkg)
- Testify: `assert` and `require` aliases

## Database Migrations

When schema changes are needed:
1. Bump `VERSION` in `config/config.go` (major increment only for schema changes).
2. Create `internal/migrations/vN.go` implementing `MajorMigrationInterface`:
   - `GetMajorVersion() float64`
   - `HasSystemUpdate() bool` / `HasWorkspaceUpdate() bool`
   - `UpdateSystem(...)` / `UpdateWorkspace(...)`
3. Use `IF NOT EXISTS` / `ADD COLUMN IF NOT EXISTS` for idempotency.
4. Register with `func init() { Register(&VNMigration{}) }`.
5. Update `CHANGELOG.md`.

## i18n (Console only)

The console uses LinguiJS. All user-facing strings must be wrapped.

```tsx
import { useLingui } from '@lingui/react/macro'

const { t } = useLingui()
return <h1>{t`Create Broadcast`}</h1>
```

- Use template literals `` t`text` ``, not `t("text")`.
- For JSX children: `<Trans>text</Trans>` from `@lingui/react/macro`.
- After adding strings: `npm run lingui:extract` then commit the `.po` files.
- `npm run build` automatically runs `lingui compile` first.

## Domain Notes

### Contact Custom Fields

Contacts support **15 fields per type** across three types: `custom_string_1..15`, `custom_number_1..15`, `custom_datetime_1..15` (45 fields total). This was expanded from 5→15 in v30 (migration `internal/migrations/v30.go`).

If the limit needs to be raised again, every layer must be updated in lockstep:

| Layer | File(s) |
|---|---|
| Domain struct + parse + Merge | `internal/domain/contact.go` |
| DB columns, INSERT/UPDATE/BulkUpsert | `internal/repository/contact_postgres.go` |
| Segment filter whitelist | `internal/service/query_builder.go` |
| DB init schema | `internal/database/init.go` |
| Migration (ALTER TABLE + trigger rebuild) | `internal/migrations/vN.go` |
| TypeScript interface | `console/src/services/api/contacts.ts` |
| Settings UI | `console/src/components/settings/CustomFieldsConfiguration.tsx` |
| Segment filter schemas | `console/src/components/segment/table_schemas.ts` |
| OpenAPI schema | `openapi/components/schemas/contact.yaml` |

The v30 migration also rebuilds the `track_contact_changes()` trigger so new fields appear in the contact timeline.

## Plans

Save implementation plans to the `plans/` directory using kebab-case filenames:
`feature-name-plan.md`, `fix-description-plan.md`, `migration-vN-plan.md`.

## Claude Agent Rules

- No AI attribution in commits, changelogs, or code comments.
- No "Generated with Claude", "Co-Authored-By: Claude", or similar markers anywhere.
