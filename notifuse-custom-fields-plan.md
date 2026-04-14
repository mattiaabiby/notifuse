# Piano: Estendere i campi custom di Notifuse da 5 a 15

## Obiettivo
Aumentare i campi custom per contatto da 5 a 15 per ciascun tipo:
- `custom_string_1..15`
- `custom_number_1..15`
- `custom_datetime_1..15`

I `custom_json_1..5` rimangono invariati.

---

## Strategia di mantenimento (rebase sul main upstream)

Questo è un fork personale mantenuto in sync con `upstream/main` tramite **rebase**.
Il commit con queste modifiche viene sempre riportato in cima alla storia con:

```bash
git fetch upstream
git rebase upstream/main
```

### Numero della migration

Il file di migration **non ha un numero fisso**. Prima di ogni rebase,
verificare l'ultimo numero presente in `internal/migrations/` e assegnare
il numero successivo. Esempio:

```bash
# Trovare l'ultimo numero usato da upstream
ls internal/migrations/v*.go | grep -v test | sort -V | tail -1
# → v29.go  =>  usare v30.go

# Se dopo un rebase upstream ha aggiunto una v30, rinominare:
git mv internal/migrations/v30.go internal/migrations/v31.go
# e aggiornare GetMajorVersion() = 31.0 e il nome della struct dentro il file
```

---

## File da modificare (in ordine di esecuzione)

### 1. Migration DB — `internal/migrations/vN.go` (NUOVO FILE)

Il numero N va determinato come descritto sopra (attualmente N=30).
Creare il file seguendo esattamente il pattern di `v29.go`.

Struttura del file:
- Nome struct: `VNMigration` (es. `V30Migration`)
- `GetMajorVersion() = N.0`
- `HasSystemUpdate() = false`
- `HasWorkspaceUpdate() = true`
- `ShouldRestartServer() = false`
- `func init() { Register(&VNMigration{}) }`

`UpdateWorkspace` deve eseguire due operazioni:

**Operazione 1 — ALTER TABLE:**
```sql
ALTER TABLE contacts
  ADD COLUMN IF NOT EXISTS custom_string_6  TEXT,
  ADD COLUMN IF NOT EXISTS custom_string_7  TEXT,
  ADD COLUMN IF NOT EXISTS custom_string_8  TEXT,
  ADD COLUMN IF NOT EXISTS custom_string_9  TEXT,
  ADD COLUMN IF NOT EXISTS custom_string_10 TEXT,
  ADD COLUMN IF NOT EXISTS custom_string_11 TEXT,
  ADD COLUMN IF NOT EXISTS custom_string_12 TEXT,
  ADD COLUMN IF NOT EXISTS custom_string_13 TEXT,
  ADD COLUMN IF NOT EXISTS custom_string_14 TEXT,
  ADD COLUMN IF NOT EXISTS custom_string_15 TEXT,
  ADD COLUMN IF NOT EXISTS custom_number_6  DOUBLE PRECISION,
  ADD COLUMN IF NOT EXISTS custom_number_7  DOUBLE PRECISION,
  ADD COLUMN IF NOT EXISTS custom_number_8  DOUBLE PRECISION,
  ADD COLUMN IF NOT EXISTS custom_number_9  DOUBLE PRECISION,
  ADD COLUMN IF NOT EXISTS custom_number_10 DOUBLE PRECISION,
  ADD COLUMN IF NOT EXISTS custom_number_11 DOUBLE PRECISION,
  ADD COLUMN IF NOT EXISTS custom_number_12 DOUBLE PRECISION,
  ADD COLUMN IF NOT EXISTS custom_number_13 DOUBLE PRECISION,
  ADD COLUMN IF NOT EXISTS custom_number_14 DOUBLE PRECISION,
  ADD COLUMN IF NOT EXISTS custom_number_15 DOUBLE PRECISION,
  ADD COLUMN IF NOT EXISTS custom_datetime_6  TIMESTAMPTZ,
  ADD COLUMN IF NOT EXISTS custom_datetime_7  TIMESTAMPTZ,
  ADD COLUMN IF NOT EXISTS custom_datetime_8  TIMESTAMPTZ,
  ADD COLUMN IF NOT EXISTS custom_datetime_9  TIMESTAMPTZ,
  ADD COLUMN IF NOT EXISTS custom_datetime_10 TIMESTAMPTZ,
  ADD COLUMN IF NOT EXISTS custom_datetime_11 TIMESTAMPTZ,
  ADD COLUMN IF NOT EXISTS custom_datetime_12 TIMESTAMPTZ,
  ADD COLUMN IF NOT EXISTS custom_datetime_13 TIMESTAMPTZ,
  ADD COLUMN IF NOT EXISTS custom_datetime_14 TIMESTAMPTZ,
  ADD COLUMN IF NOT EXISTS custom_datetime_15 TIMESTAMPTZ;
```

**Operazione 2 — ricreare il trigger di change tracking:**

Il trigger confronta OLD e NEW per ogni campo e scrive le differenze nella
timeline del contatto. Prendere come riferimento il trigger in `v23.go`
(righe ~178-192) e aggiungere i blocchi per i campi 6..15:

```sql
IF OLD.custom_string_6 IS DISTINCT FROM NEW.custom_string_6 THEN
  changes_json := changes_json || jsonb_build_object('custom_string_6',
    jsonb_build_object('old', OLD.custom_string_6, 'new', NEW.custom_string_6));
END IF;
-- ripetere per _7.._15, poi per custom_number_6.._15, custom_datetime_6.._15
```

Il trigger va DROP + CREATE (come fanno le migration precedenti) per includere
tutti i campi vecchi e nuovi. Senza questa parte, le modifiche ai campi 6..15
non appaiono nella timeline del contatto.

---

### 2. Domain struct — `internal/domain/contact.go`

Aggiungere i campi 6..15 subito dopo i rispettivi _5:

```go
// dopo CustomString5:
CustomString6  *NullableString `json:"custom_string_6,omitempty"  valid:"optional"`
// ... fino a CustomString15

// dopo CustomNumber5:
CustomNumber6  *NullableFloat64 `json:"custom_number_6,omitempty"  valid:"optional"`
// ... fino a CustomNumber15

// dopo CustomDatetime5:
CustomDatetime6  *NullableTime `json:"custom_datetime_6,omitempty"  valid:"optional"`
// ... fino a CustomDatetime15
```

Nella funzione di parsing (cerca `parseNullableString(jsonResult, "custom_string_1"`),
aggiungere le chiamate per i campi 6..15:

```go
if err := parseNullableString(jsonResult, "custom_string_6", &contact.CustomString6); err != nil { ... }
// ... fino a _15
if err := parseNullableFloat(jsonResult, "custom_number_6", &contact.CustomNumber6); err != nil { ... }
// ... fino a _15
if err := parseNullableTime(jsonResult, "custom_datetime_6", &contact.CustomDatetime6); err != nil { ... }
// ... fino a _15
```

---

### 3. Repository — `internal/repository/contact_postgres.go`

Cercare `custom_string_5` per trovare tutti i punti. Ce ne sono **3**:

**A) Lista colonne SELECT** (intorno alla riga 1202):
Aggiungere `custom_string_6..15` dopo `custom_string_5`. Stessa cosa per number e datetime.

**B) Lista colonne INSERT/UPSERT** (intorno alle righe 822-825):
Aggiungere `"custom_string_6"..` agli slice. Stessa cosa per number e datetime.

**C) Map dei valori per UPDATE** (intorno alle righe 1082-1101):
```go
"custom_string_6":  customString6SQL,
// ... fino a _15
```
Le variabili `customString6SQL` ecc. vanno dichiarate seguendo il pattern
di `customString1SQL`..`customString5SQL` già presenti sopra.

---

### 4. Query Builder — `internal/service/query_builder.go`

In `initializeContactFields()`, aggiungere dopo ogni `_5` esistente:

```go
// dopo custom_string_5:
"custom_string_6", "custom_string_7", "custom_string_8", "custom_string_9",
"custom_string_10", "custom_string_11", "custom_string_12", "custom_string_13",
"custom_string_14", "custom_string_15",

// dopo custom_number_5:
"custom_number_6", ..., "custom_number_15",

// dopo custom_datetime_5:
"custom_datetime_6", ..., "custom_datetime_15",
```

---

### 5. Frontend TypeScript — `console/src/services/api/contacts.ts`

Nell'interfaccia `Contact`, aggiungere dopo `custom_string_5`:
```typescript
custom_string_6?: string
// ... fino a custom_string_15
// stessa cosa per custom_number_6..15 (number) e custom_datetime_6..15 (string)
```

---

### 6. Frontend Settings — `console/src/components/settings/CustomFieldsConfiguration.tsx`

Tre righe, cambiare solo il numero `5` → `15`:
```typescript
String:   Array.from({ length: 15 }, (_, i) => `custom_string_${i + 1}`),
Number:   Array.from({ length: 15 }, (_, i) => `custom_number_${i + 1}`),
Datetime: Array.from({ length: 15 }, (_, i) => `custom_datetime_${i + 1}`),
```

---

### 7. Frontend Segmenti — `console/src/components/segment/table_schemas.ts`

Aggiungere le entry per i campi 6..15 seguendo esattamente il pattern
delle entry 1..5. Cercare `custom_string_5` per trovare il punto di
inserimento e replicare il blocco per ogni nuovo campo.

---

### 8. OpenAPI schema — `openapi/components/schemas/contact.yaml`

Aggiungere le proprietà `custom_string_6..15`, `custom_number_6..15`,
`custom_datetime_6..15` seguendo il formato delle proprietà 1..5.

---

## File che NON vanno modificati

- `internal/migrations/v7.go`, `v18.go`, `v19.go`, `v23.go` — migration storiche
- `internal/domain/tree.go` — non dipende dal numero di campi
- `console/src/components/contacts/fieldTypes.ts` — usa `startsWith()`, già dinamico
- `console/src/hooks/useCustomFieldLabel.ts` — genera label dinamicamente

---

## Gestione conflitti durante i rebase futuri

I file più soggetti a conflitti sono:

- **`contact_postgres.go`** — file lungo, toccato spesso. I conflitti saranno
  quasi sempre nella sezione custom field: tenere le modifiche di upstream
  e ri-aggiungere le righe 6..15 dopo il `_5` corrispondente.

- **`table_schemas.ts`** — frontend che cambia frequentemente. Stesso approccio.

- **`contact.go`** — meno frequente, ma possibile se upstream aggiunge campi.

Regola generale: **le tue righe vanno sempre immediatamente dopo il `_5` di
upstream**, mai mescolate nel mezzo. Questo rende i conflitti meccanici e
facili da risolvere.
