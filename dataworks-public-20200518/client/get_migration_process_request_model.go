// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMigrationProcessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMigrationId(v int64) *GetMigrationProcessRequest
	GetMigrationId() *int64
	SetProjectId(v int64) *GetMigrationProcessRequest
	GetProjectId() *int64
}

type GetMigrationProcessRequest struct {
	// The migration package ID. You can call the CreateImportMigration operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	MigrationId *int64 `json:"MigrationId,omitempty" xml:"MigrationId,omitempty"`
	// The workspace ID. You can log on to the DataWorks console and go to the Workspace page to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetMigrationProcessRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMigrationProcessRequest) GoString() string {
	return s.String()
}

func (s *GetMigrationProcessRequest) GetMigrationId() *int64 {
	return s.MigrationId
}

func (s *GetMigrationProcessRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetMigrationProcessRequest) SetMigrationId(v int64) *GetMigrationProcessRequest {
	s.MigrationId = &v
	return s
}

func (s *GetMigrationProcessRequest) SetProjectId(v int64) *GetMigrationProcessRequest {
	s.ProjectId = &v
	return s
}

func (s *GetMigrationProcessRequest) Validate() error {
	return dara.Validate(s)
}
