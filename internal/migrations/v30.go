package migrations

import (
	"context"
	"fmt"

	"github.com/Notifuse/notifuse/config"
	"github.com/Notifuse/notifuse/internal/domain"
)

// V30Migration extends custom fields from 5 to 15 per type.
//
// This migration:
// - Workspace: adds custom_string_6..15, custom_number_6..15, custom_datetime_6..15 columns
// - Workspace: recreates track_contact_changes() trigger to include the new fields
type V30Migration struct{}

func (m *V30Migration) GetMajorVersion() float64 {
	return 30.0
}

func (m *V30Migration) HasSystemUpdate() bool {
	return false
}

func (m *V30Migration) HasWorkspaceUpdate() bool {
	return true
}

func (m *V30Migration) ShouldRestartServer() bool {
	return false
}

func (m *V30Migration) UpdateSystem(ctx context.Context, cfg *config.Config, db DBExecutor) error {
	return nil
}

func (m *V30Migration) UpdateWorkspace(ctx context.Context, cfg *config.Config, workspace *domain.Workspace, db DBExecutor) error {
	// PART 1: Add new custom field columns
	_, err := db.ExecContext(ctx, `
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
		  ADD COLUMN IF NOT EXISTS custom_datetime_15 TIMESTAMPTZ
	`)
	if err != nil {
		return fmt.Errorf("failed to add custom field columns: %w", err)
	}

	// PART 2: Recreate track_contact_changes() to include all custom fields 1..15
	_, err = db.ExecContext(ctx, `
		CREATE OR REPLACE FUNCTION track_contact_changes()
		RETURNS TRIGGER AS $$
		DECLARE
			changes_json JSONB := '{}'::jsonb;
			op VARCHAR(20);
		BEGIN
			IF TG_OP = 'INSERT' THEN
				op := 'insert';
				changes_json := NULL;
			ELSIF TG_OP = 'UPDATE' THEN
				op := 'update';
				IF OLD.external_id IS DISTINCT FROM NEW.external_id THEN changes_json := changes_json || jsonb_build_object('external_id', jsonb_build_object('old', OLD.external_id, 'new', NEW.external_id)); END IF;
				IF OLD.timezone IS DISTINCT FROM NEW.timezone THEN changes_json := changes_json || jsonb_build_object('timezone', jsonb_build_object('old', OLD.timezone, 'new', NEW.timezone)); END IF;
				IF OLD.language IS DISTINCT FROM NEW.language THEN changes_json := changes_json || jsonb_build_object('language', jsonb_build_object('old', OLD.language, 'new', NEW.language)); END IF;
				IF OLD.first_name IS DISTINCT FROM NEW.first_name THEN changes_json := changes_json || jsonb_build_object('first_name', jsonb_build_object('old', OLD.first_name, 'new', NEW.first_name)); END IF;
				IF OLD.last_name IS DISTINCT FROM NEW.last_name THEN changes_json := changes_json || jsonb_build_object('last_name', jsonb_build_object('old', OLD.last_name, 'new', NEW.last_name)); END IF;
				IF OLD.full_name IS DISTINCT FROM NEW.full_name THEN changes_json := changes_json || jsonb_build_object('full_name', jsonb_build_object('old', OLD.full_name, 'new', NEW.full_name)); END IF;
				IF OLD.phone IS DISTINCT FROM NEW.phone THEN changes_json := changes_json || jsonb_build_object('phone', jsonb_build_object('old', OLD.phone, 'new', NEW.phone)); END IF;
				IF OLD.address_line_1 IS DISTINCT FROM NEW.address_line_1 THEN changes_json := changes_json || jsonb_build_object('address_line_1', jsonb_build_object('old', OLD.address_line_1, 'new', NEW.address_line_1)); END IF;
				IF OLD.address_line_2 IS DISTINCT FROM NEW.address_line_2 THEN changes_json := changes_json || jsonb_build_object('address_line_2', jsonb_build_object('old', OLD.address_line_2, 'new', NEW.address_line_2)); END IF;
				IF OLD.country IS DISTINCT FROM NEW.country THEN changes_json := changes_json || jsonb_build_object('country', jsonb_build_object('old', OLD.country, 'new', NEW.country)); END IF;
				IF OLD.postcode IS DISTINCT FROM NEW.postcode THEN changes_json := changes_json || jsonb_build_object('postcode', jsonb_build_object('old', OLD.postcode, 'new', NEW.postcode)); END IF;
				IF OLD.state IS DISTINCT FROM NEW.state THEN changes_json := changes_json || jsonb_build_object('state', jsonb_build_object('old', OLD.state, 'new', NEW.state)); END IF;
				IF OLD.job_title IS DISTINCT FROM NEW.job_title THEN changes_json := changes_json || jsonb_build_object('job_title', jsonb_build_object('old', OLD.job_title, 'new', NEW.job_title)); END IF;
				IF OLD.custom_string_1 IS DISTINCT FROM NEW.custom_string_1 THEN changes_json := changes_json || jsonb_build_object('custom_string_1', jsonb_build_object('old', OLD.custom_string_1, 'new', NEW.custom_string_1)); END IF;
				IF OLD.custom_string_2 IS DISTINCT FROM NEW.custom_string_2 THEN changes_json := changes_json || jsonb_build_object('custom_string_2', jsonb_build_object('old', OLD.custom_string_2, 'new', NEW.custom_string_2)); END IF;
				IF OLD.custom_string_3 IS DISTINCT FROM NEW.custom_string_3 THEN changes_json := changes_json || jsonb_build_object('custom_string_3', jsonb_build_object('old', OLD.custom_string_3, 'new', NEW.custom_string_3)); END IF;
				IF OLD.custom_string_4 IS DISTINCT FROM NEW.custom_string_4 THEN changes_json := changes_json || jsonb_build_object('custom_string_4', jsonb_build_object('old', OLD.custom_string_4, 'new', NEW.custom_string_4)); END IF;
				IF OLD.custom_string_5 IS DISTINCT FROM NEW.custom_string_5 THEN changes_json := changes_json || jsonb_build_object('custom_string_5', jsonb_build_object('old', OLD.custom_string_5, 'new', NEW.custom_string_5)); END IF;
				IF OLD.custom_string_6 IS DISTINCT FROM NEW.custom_string_6 THEN changes_json := changes_json || jsonb_build_object('custom_string_6', jsonb_build_object('old', OLD.custom_string_6, 'new', NEW.custom_string_6)); END IF;
				IF OLD.custom_string_7 IS DISTINCT FROM NEW.custom_string_7 THEN changes_json := changes_json || jsonb_build_object('custom_string_7', jsonb_build_object('old', OLD.custom_string_7, 'new', NEW.custom_string_7)); END IF;
				IF OLD.custom_string_8 IS DISTINCT FROM NEW.custom_string_8 THEN changes_json := changes_json || jsonb_build_object('custom_string_8', jsonb_build_object('old', OLD.custom_string_8, 'new', NEW.custom_string_8)); END IF;
				IF OLD.custom_string_9 IS DISTINCT FROM NEW.custom_string_9 THEN changes_json := changes_json || jsonb_build_object('custom_string_9', jsonb_build_object('old', OLD.custom_string_9, 'new', NEW.custom_string_9)); END IF;
				IF OLD.custom_string_10 IS DISTINCT FROM NEW.custom_string_10 THEN changes_json := changes_json || jsonb_build_object('custom_string_10', jsonb_build_object('old', OLD.custom_string_10, 'new', NEW.custom_string_10)); END IF;
				IF OLD.custom_string_11 IS DISTINCT FROM NEW.custom_string_11 THEN changes_json := changes_json || jsonb_build_object('custom_string_11', jsonb_build_object('old', OLD.custom_string_11, 'new', NEW.custom_string_11)); END IF;
				IF OLD.custom_string_12 IS DISTINCT FROM NEW.custom_string_12 THEN changes_json := changes_json || jsonb_build_object('custom_string_12', jsonb_build_object('old', OLD.custom_string_12, 'new', NEW.custom_string_12)); END IF;
				IF OLD.custom_string_13 IS DISTINCT FROM NEW.custom_string_13 THEN changes_json := changes_json || jsonb_build_object('custom_string_13', jsonb_build_object('old', OLD.custom_string_13, 'new', NEW.custom_string_13)); END IF;
				IF OLD.custom_string_14 IS DISTINCT FROM NEW.custom_string_14 THEN changes_json := changes_json || jsonb_build_object('custom_string_14', jsonb_build_object('old', OLD.custom_string_14, 'new', NEW.custom_string_14)); END IF;
				IF OLD.custom_string_15 IS DISTINCT FROM NEW.custom_string_15 THEN changes_json := changes_json || jsonb_build_object('custom_string_15', jsonb_build_object('old', OLD.custom_string_15, 'new', NEW.custom_string_15)); END IF;
				IF OLD.custom_number_1 IS DISTINCT FROM NEW.custom_number_1 THEN changes_json := changes_json || jsonb_build_object('custom_number_1', jsonb_build_object('old', OLD.custom_number_1, 'new', NEW.custom_number_1)); END IF;
				IF OLD.custom_number_2 IS DISTINCT FROM NEW.custom_number_2 THEN changes_json := changes_json || jsonb_build_object('custom_number_2', jsonb_build_object('old', OLD.custom_number_2, 'new', NEW.custom_number_2)); END IF;
				IF OLD.custom_number_3 IS DISTINCT FROM NEW.custom_number_3 THEN changes_json := changes_json || jsonb_build_object('custom_number_3', jsonb_build_object('old', OLD.custom_number_3, 'new', NEW.custom_number_3)); END IF;
				IF OLD.custom_number_4 IS DISTINCT FROM NEW.custom_number_4 THEN changes_json := changes_json || jsonb_build_object('custom_number_4', jsonb_build_object('old', OLD.custom_number_4, 'new', NEW.custom_number_4)); END IF;
				IF OLD.custom_number_5 IS DISTINCT FROM NEW.custom_number_5 THEN changes_json := changes_json || jsonb_build_object('custom_number_5', jsonb_build_object('old', OLD.custom_number_5, 'new', NEW.custom_number_5)); END IF;
				IF OLD.custom_number_6 IS DISTINCT FROM NEW.custom_number_6 THEN changes_json := changes_json || jsonb_build_object('custom_number_6', jsonb_build_object('old', OLD.custom_number_6, 'new', NEW.custom_number_6)); END IF;
				IF OLD.custom_number_7 IS DISTINCT FROM NEW.custom_number_7 THEN changes_json := changes_json || jsonb_build_object('custom_number_7', jsonb_build_object('old', OLD.custom_number_7, 'new', NEW.custom_number_7)); END IF;
				IF OLD.custom_number_8 IS DISTINCT FROM NEW.custom_number_8 THEN changes_json := changes_json || jsonb_build_object('custom_number_8', jsonb_build_object('old', OLD.custom_number_8, 'new', NEW.custom_number_8)); END IF;
				IF OLD.custom_number_9 IS DISTINCT FROM NEW.custom_number_9 THEN changes_json := changes_json || jsonb_build_object('custom_number_9', jsonb_build_object('old', OLD.custom_number_9, 'new', NEW.custom_number_9)); END IF;
				IF OLD.custom_number_10 IS DISTINCT FROM NEW.custom_number_10 THEN changes_json := changes_json || jsonb_build_object('custom_number_10', jsonb_build_object('old', OLD.custom_number_10, 'new', NEW.custom_number_10)); END IF;
				IF OLD.custom_number_11 IS DISTINCT FROM NEW.custom_number_11 THEN changes_json := changes_json || jsonb_build_object('custom_number_11', jsonb_build_object('old', OLD.custom_number_11, 'new', NEW.custom_number_11)); END IF;
				IF OLD.custom_number_12 IS DISTINCT FROM NEW.custom_number_12 THEN changes_json := changes_json || jsonb_build_object('custom_number_12', jsonb_build_object('old', OLD.custom_number_12, 'new', NEW.custom_number_12)); END IF;
				IF OLD.custom_number_13 IS DISTINCT FROM NEW.custom_number_13 THEN changes_json := changes_json || jsonb_build_object('custom_number_13', jsonb_build_object('old', OLD.custom_number_13, 'new', NEW.custom_number_13)); END IF;
				IF OLD.custom_number_14 IS DISTINCT FROM NEW.custom_number_14 THEN changes_json := changes_json || jsonb_build_object('custom_number_14', jsonb_build_object('old', OLD.custom_number_14, 'new', NEW.custom_number_14)); END IF;
				IF OLD.custom_number_15 IS DISTINCT FROM NEW.custom_number_15 THEN changes_json := changes_json || jsonb_build_object('custom_number_15', jsonb_build_object('old', OLD.custom_number_15, 'new', NEW.custom_number_15)); END IF;
				IF OLD.custom_datetime_1 IS DISTINCT FROM NEW.custom_datetime_1 THEN changes_json := changes_json || jsonb_build_object('custom_datetime_1', jsonb_build_object('old', OLD.custom_datetime_1, 'new', NEW.custom_datetime_1)); END IF;
				IF OLD.custom_datetime_2 IS DISTINCT FROM NEW.custom_datetime_2 THEN changes_json := changes_json || jsonb_build_object('custom_datetime_2', jsonb_build_object('old', OLD.custom_datetime_2, 'new', NEW.custom_datetime_2)); END IF;
				IF OLD.custom_datetime_3 IS DISTINCT FROM NEW.custom_datetime_3 THEN changes_json := changes_json || jsonb_build_object('custom_datetime_3', jsonb_build_object('old', OLD.custom_datetime_3, 'new', NEW.custom_datetime_3)); END IF;
				IF OLD.custom_datetime_4 IS DISTINCT FROM NEW.custom_datetime_4 THEN changes_json := changes_json || jsonb_build_object('custom_datetime_4', jsonb_build_object('old', OLD.custom_datetime_4, 'new', NEW.custom_datetime_4)); END IF;
				IF OLD.custom_datetime_5 IS DISTINCT FROM NEW.custom_datetime_5 THEN changes_json := changes_json || jsonb_build_object('custom_datetime_5', jsonb_build_object('old', OLD.custom_datetime_5, 'new', NEW.custom_datetime_5)); END IF;
				IF OLD.custom_datetime_6 IS DISTINCT FROM NEW.custom_datetime_6 THEN changes_json := changes_json || jsonb_build_object('custom_datetime_6', jsonb_build_object('old', OLD.custom_datetime_6, 'new', NEW.custom_datetime_6)); END IF;
				IF OLD.custom_datetime_7 IS DISTINCT FROM NEW.custom_datetime_7 THEN changes_json := changes_json || jsonb_build_object('custom_datetime_7', jsonb_build_object('old', OLD.custom_datetime_7, 'new', NEW.custom_datetime_7)); END IF;
				IF OLD.custom_datetime_8 IS DISTINCT FROM NEW.custom_datetime_8 THEN changes_json := changes_json || jsonb_build_object('custom_datetime_8', jsonb_build_object('old', OLD.custom_datetime_8, 'new', NEW.custom_datetime_8)); END IF;
				IF OLD.custom_datetime_9 IS DISTINCT FROM NEW.custom_datetime_9 THEN changes_json := changes_json || jsonb_build_object('custom_datetime_9', jsonb_build_object('old', OLD.custom_datetime_9, 'new', NEW.custom_datetime_9)); END IF;
				IF OLD.custom_datetime_10 IS DISTINCT FROM NEW.custom_datetime_10 THEN changes_json := changes_json || jsonb_build_object('custom_datetime_10', jsonb_build_object('old', OLD.custom_datetime_10, 'new', NEW.custom_datetime_10)); END IF;
				IF OLD.custom_datetime_11 IS DISTINCT FROM NEW.custom_datetime_11 THEN changes_json := changes_json || jsonb_build_object('custom_datetime_11', jsonb_build_object('old', OLD.custom_datetime_11, 'new', NEW.custom_datetime_11)); END IF;
				IF OLD.custom_datetime_12 IS DISTINCT FROM NEW.custom_datetime_12 THEN changes_json := changes_json || jsonb_build_object('custom_datetime_12', jsonb_build_object('old', OLD.custom_datetime_12, 'new', NEW.custom_datetime_12)); END IF;
				IF OLD.custom_datetime_13 IS DISTINCT FROM NEW.custom_datetime_13 THEN changes_json := changes_json || jsonb_build_object('custom_datetime_13', jsonb_build_object('old', OLD.custom_datetime_13, 'new', NEW.custom_datetime_13)); END IF;
				IF OLD.custom_datetime_14 IS DISTINCT FROM NEW.custom_datetime_14 THEN changes_json := changes_json || jsonb_build_object('custom_datetime_14', jsonb_build_object('old', OLD.custom_datetime_14, 'new', NEW.custom_datetime_14)); END IF;
				IF OLD.custom_datetime_15 IS DISTINCT FROM NEW.custom_datetime_15 THEN changes_json := changes_json || jsonb_build_object('custom_datetime_15', jsonb_build_object('old', OLD.custom_datetime_15, 'new', NEW.custom_datetime_15)); END IF;
				IF OLD.custom_json_1 IS DISTINCT FROM NEW.custom_json_1 THEN changes_json := changes_json || jsonb_build_object('custom_json_1', jsonb_build_object('old', OLD.custom_json_1, 'new', NEW.custom_json_1)); END IF;
				IF OLD.custom_json_2 IS DISTINCT FROM NEW.custom_json_2 THEN changes_json := changes_json || jsonb_build_object('custom_json_2', jsonb_build_object('old', OLD.custom_json_2, 'new', NEW.custom_json_2)); END IF;
				IF OLD.custom_json_3 IS DISTINCT FROM NEW.custom_json_3 THEN changes_json := changes_json || jsonb_build_object('custom_json_3', jsonb_build_object('old', OLD.custom_json_3, 'new', NEW.custom_json_3)); END IF;
				IF OLD.custom_json_4 IS DISTINCT FROM NEW.custom_json_4 THEN changes_json := changes_json || jsonb_build_object('custom_json_4', jsonb_build_object('old', OLD.custom_json_4, 'new', NEW.custom_json_4)); END IF;
				IF OLD.custom_json_5 IS DISTINCT FROM NEW.custom_json_5 THEN changes_json := changes_json || jsonb_build_object('custom_json_5', jsonb_build_object('old', OLD.custom_json_5, 'new', NEW.custom_json_5)); END IF;
				IF changes_json = '{}'::jsonb THEN RETURN NEW; END IF;
			END IF;
			IF TG_OP = 'INSERT' THEN
				INSERT INTO contact_timeline (email, operation, entity_type, kind, changes, created_at)
				VALUES (NEW.email, op, 'contact', 'contact.created', changes_json, NEW.created_at);
			ELSE
				INSERT INTO contact_timeline (email, operation, entity_type, kind, changes, created_at)
				VALUES (NEW.email, op, 'contact', 'contact.updated', changes_json, NEW.updated_at);
			END IF;
			RETURN NEW;
		END;
		$$ LANGUAGE plpgsql;
	`)
	if err != nil {
		return fmt.Errorf("failed to update track_contact_changes function: %w", err)
	}

	return nil
}

func init() {
	Register(&V30Migration{})
}
