// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMigrationSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMigrationId(v int64) *GetMigrationSummaryRequest
	GetMigrationId() *int64
	SetProjectId(v int64) *GetMigrationSummaryRequest
	GetProjectId() *int64
}

type GetMigrationSummaryRequest struct {
	// The migration task ID.
	//
	// You can call the [CreateImportMigration](https://help.aliyun.com/document_detail/2780280.html) operation to obtain the ID of the import task and call the [CreateExportMigration](https://help.aliyun.com/document_detail/2780281.html) operation to obtain the ID of the export task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	MigrationId *int64 `json:"MigrationId,omitempty" xml:"MigrationId,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to obtain the workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetMigrationSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMigrationSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetMigrationSummaryRequest) GetMigrationId() *int64 {
	return s.MigrationId
}

func (s *GetMigrationSummaryRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetMigrationSummaryRequest) SetMigrationId(v int64) *GetMigrationSummaryRequest {
	s.MigrationId = &v
	return s
}

func (s *GetMigrationSummaryRequest) SetProjectId(v int64) *GetMigrationSummaryRequest {
	s.ProjectId = &v
	return s
}

func (s *GetMigrationSummaryRequest) Validate() error {
	return dara.Validate(s)
}
