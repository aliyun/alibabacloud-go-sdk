// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartMigrationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMigrationId(v int64) *StartMigrationRequest
	GetMigrationId() *int64
	SetProjectId(v int64) *StartMigrationRequest
	GetProjectId() *int64
}

type StartMigrationRequest struct {
	// The migration package ID. You can call the [CreateImportMigration](https://help.aliyun.com/document_detail/206094.html) operation to query the ID of the import package and call the [CreateExportMigration](https://help.aliyun.com/document_detail/349325.html) operation to query the ID of the export package.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	MigrationId *int64 `json:"MigrationId,omitempty" xml:"MigrationId,omitempty"`
	// The workspace ID. You can log on to the DataWorks console and go to the Workspace page to obtain the workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s StartMigrationRequest) String() string {
	return dara.Prettify(s)
}

func (s StartMigrationRequest) GoString() string {
	return s.String()
}

func (s *StartMigrationRequest) GetMigrationId() *int64 {
	return s.MigrationId
}

func (s *StartMigrationRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *StartMigrationRequest) SetMigrationId(v int64) *StartMigrationRequest {
	s.MigrationId = &v
	return s
}

func (s *StartMigrationRequest) SetProjectId(v int64) *StartMigrationRequest {
	s.ProjectId = &v
	return s
}

func (s *StartMigrationRequest) Validate() error {
	return dara.Validate(s)
}
