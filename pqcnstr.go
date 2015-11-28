package pqcnstr

import (
	"strings"

	"github.com/lib/pq"
)

const (
	kindUnknown    = "unknown"
	kindPrimaryKey = "pkey"
	kindCheck      = "check"
	kindUnique     = "key"
	kindIndex      = "idx"
	kindForeignKey = "fkey"
	kindExclusion  = "excl"
)

// Constraint is helper wrapper for constraint that gives few extra options.
type Constraint string

func any(kind, schema, table string, columns ...string) Constraint {
	if schema == "" {
		schema = "public"
	}
	if len(columns) == 0 {
		return Constraint(schema + "." + table + "_" + kind)
	}
	return Constraint(schema + "." + table + "_" + strings.Join(columns, "_") + "_" + kind)
}

// Unique constraint ensure that the data contained in a column or a group of columns is unique with respect to all the rows in the table.
func Unique(schema, table string, columns ...string) Constraint {
	return any(kindUnique, schema, table, columns...)
}

// PrimaryKey constraint is simply a combination of a unique constraint and a not-null constraint.
func PrimaryKey(schema, table string) Constraint {
	return any(kindPrimaryKey, schema, table)
}

// Exclusion constraint ensure that if any two rows are compared on the specified columns
// or expressions using the specified operators,
// at least one of these operator comparisons will return false or null.
func Exclusion(schema, table string, columns ...string) Constraint {
	return any(kindExclusion, schema, table, columns...)
}

// ForeignKey constraint specifies that the values in a column (or a group of columns)
// must match the values appearing in some row of another table.
// We say this maintains the referential integrity between two related tables.
func ForeignKey(schema, table string, columns ...string) Constraint {
	return any(kindForeignKey, schema, table, columns...)
}

// Index ...
func Index(schema, table string, columns ...string) Constraint {
	return any(kindIndex, schema, table, columns...)
}

// String implements Stringer interface.
func (c Constraint) String() string {
	return string(c)
}

// Type returns name of a type. It can be one of unknown, pkey, check, fkey, idx, key or excl.
func (c Constraint) Type() string {
	if c == "" {
		return kindUnknown
	}

	parts := strings.Split(c.String(), "_")

	return parts[len(parts)-1]
}

// IsForeignKey returns true if constraint has suffix "_fkey".
func (c Constraint) IsForeignKey() bool {
	return IsForeignKey(c.String())
}

// IsUnique returns true if constraint has suffix "_key".
func (c Constraint) IsUnique() bool {
	return IsUnique(c.String())
}

// IsPrimaryKey returns true if constraint has suffix "_pkey".
func (c Constraint) IsPrimaryKey() bool {
	return IsPrimaryKey(c.String())
}

// IsCheck returns true if constraint has suffix "_check".
func (c Constraint) IsCheck() bool {
	return IsCheck(c.String())
}

// IsExclusion returns true if constraint has suffix "_excl".
func (c Constraint) IsExclusion() bool {
	return IsExclusion(c.String())
}

// IsIndex returns true if constraint has suffix "_idx".
func (c Constraint) IsIndex() bool {
	return IsIndex(c.String())
}

// IsEmpty returns true if constraint is an empty string.
func (c Constraint) IsEmpty() bool {
	return c == ""
}

// IsForeignKey returns true if string has suffix "_fkey".
func IsForeignKey(c string) bool {
	return strings.HasSuffix(c, kindForeignKey)
}

// IsUnique returns true if string has suffix "_key".
func IsUnique(c string) bool {
	return strings.HasSuffix(c, kindUnique)
}

// IsPrimaryKey returns true if string has suffix "_pkey".
func IsPrimaryKey(c string) bool {
	return strings.HasSuffix(c, kindPrimaryKey)
}

// IsCheck returns true if string has suffix "_check".
func IsCheck(c string) bool {
	return strings.HasSuffix(c, kindCheck)
}

// IsExclusion returns true if string has suffix "_excl".
func IsExclusion(c string) bool {
	return strings.HasSuffix(c, kindExclusion)
}

// IsIndex returns true if string has suffix "_idx".
func IsIndex(c string) bool {
	return strings.HasSuffix(c, kindIndex)
}

// FromError retrieves Constraint from error if possible. Returns false if its not.
func FromError(err error) (cnstr Constraint) {
	var ok bool
	var e *pq.Error

	e, ok = err.(*pq.Error)
	if !ok {
		return
	}

	switch e.Code.Name() {
	case "unique_violation":
		cnstr = Constraint(e.Constraint)
	}

	return
}
