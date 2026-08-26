// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrationStatusFilter interface {
	dara.Model
	String() string
	GoString() string
	SetIsMigrated(v bool) *MigrationStatusFilter
	GetIsMigrated() *bool
}

type MigrationStatusFilter struct {
	// Specifies whether to filter by migration rule. Valid values:
	//
	// - true: Only migrated rules (migration_status is not 0 or NULL).
	//
	// - false: Only native rules (migration_status = 0).
	IsMigrated *bool `json:"isMigrated,omitempty" xml:"isMigrated,omitempty"`
}

func (s MigrationStatusFilter) String() string {
	return dara.Prettify(s)
}

func (s MigrationStatusFilter) GoString() string {
	return s.String()
}

func (s *MigrationStatusFilter) GetIsMigrated() *bool {
	return s.IsMigrated
}

func (s *MigrationStatusFilter) SetIsMigrated(v bool) *MigrationStatusFilter {
	s.IsMigrated = &v
	return s
}

func (s *MigrationStatusFilter) Validate() error {
	return dara.Validate(s)
}
