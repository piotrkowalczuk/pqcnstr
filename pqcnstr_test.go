package pqcnstr

import "testing"

func TestAny(t *testing.T) {
	success := map[Constraint]struct {
		schema, table, kind string
		columns             []string
	}{
		"schema.table_c1_c2_kind": {
			schema:  "schema",
			table:   "table",
			kind:    "kind",
			columns: []string{"c1", "c2"},
		},
		"public.user_pkey": {
			table: "user",
			kind:  "pkey",
		},
	}

	for expected, given := range success {
		cnstr := any(given.kind, given.schema, given.table, given.columns...)

		if cnstr != expected {
			t.Errorf("incorrect constraint, expected %s, got %s", expected, cnstr)
		}
	}
}

func TestConstraint_Type(t *testing.T) {
	success := map[string]string{
		"idx":   "schema.user_c1_c2_idx",
		"pkey":  "schema.user_pkey",
		"key":   "schema.user_id_key",
		"excl":  "schema.user_first_name_last_name_excl",
		"fkey":  "schema.user_details_id_fkey",
		"check": "schema.user_email_check",
	}

	for expected, given := range success {
		c := Constraint(given)

		if c.Type() != expected {
			t.Errorf("expected to be type of %s, got %s", expected, given)
		}
	}
}

func TestConstraint_IsIndex(t *testing.T) {
	success := []string{
		"schema.user_c1_c2_idx",
	}

	for _, given := range success {
		idx := Constraint(given)

		if !idx.IsIndex() {
			t.Errorf("expected to be recognize as index: %s", given)
		}
	}
}

func TestConstraint_IsForeignKey(t *testing.T) {
	success := []string{
		"schema.user_c1_c2_fkey",
	}

	for _, given := range success {
		fkey := Constraint(given)

		if !fkey.IsForeignKey() {
			t.Errorf("expected to be recognize as foreign key: %s", given)
		}
	}
}

func TestConstraint_IsCheck(t *testing.T) {
	success := []string{
		"schema.user_c1_c2_check",
	}

	for _, given := range success {
		check := Constraint(given)

		if !check.IsCheck() {
			t.Errorf("expected to be recognize as check: %s", given)
		}
	}
}

func TestConstraint_IsUnique(t *testing.T) {
	success := []string{
		"schema.user_c1_c2_key",
	}

	for _, given := range success {
		key := Constraint(given)

		if !key.IsUnique() {
			t.Errorf("expected to be recognize as check: %s", given)
		}
	}
}

func TestConstraint_IsExclusion(t *testing.T) {
	success := []string{
		"schema.user_c1_c2_excl",
	}

	for _, given := range success {
		excl := Constraint(given)

		if !excl.IsExclusion() {
			t.Errorf("expected to be recognize as exclusion: %s", given)
		}
	}
}
func TestConstraint_IsPrimaryKey(t *testing.T) {
	success := []string{
		"schema.user_c1_c2_pkey",
	}

	for _, given := range success {
		pkey := Constraint(given)

		if !pkey.IsPrimaryKey() {
			t.Errorf("expected to be recognize as primary key: %s", given)
		}
	}
}
