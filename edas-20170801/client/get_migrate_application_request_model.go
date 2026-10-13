// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMigrateApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMigrationId(v string) *GetMigrateApplicationRequest
	GetMigrationId() *string
	SetMigrationType(v int64) *GetMigrateApplicationRequest
	GetMigrationType() *int64
	SetRegionId(v string) *GetMigrateApplicationRequest
	GetRegionId() *string
}

type GetMigrateApplicationRequest struct {
	// The migration ID.
	//
	// example:
	//
	// afe9995d-31fd-40ff-a669-c125d76d8267
	MigrationId *string `json:"migrationId,omitempty" xml:"migrationId,omitempty"`
	// The operation type of the migration task. Valid values: 1: import. 2: export. 3: clone. If this parameter is not specified, all types are returned.
	//
	// example:
	//
	// 1
	MigrationType *int64 `json:"migrationType,omitempty" xml:"migrationType,omitempty"`
	// regionId
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GetMigrateApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMigrateApplicationRequest) GoString() string {
	return s.String()
}

func (s *GetMigrateApplicationRequest) GetMigrationId() *string {
	return s.MigrationId
}

func (s *GetMigrateApplicationRequest) GetMigrationType() *int64 {
	return s.MigrationType
}

func (s *GetMigrateApplicationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetMigrateApplicationRequest) SetMigrationId(v string) *GetMigrateApplicationRequest {
	s.MigrationId = &v
	return s
}

func (s *GetMigrateApplicationRequest) SetMigrationType(v int64) *GetMigrateApplicationRequest {
	s.MigrationType = &v
	return s
}

func (s *GetMigrateApplicationRequest) SetRegionId(v string) *GetMigrateApplicationRequest {
	s.RegionId = &v
	return s
}

func (s *GetMigrateApplicationRequest) Validate() error {
	return dara.Validate(s)
}
