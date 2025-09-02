// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExportMigrationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateExportMigrationRequest
	GetDescription() *string
	SetExportMode(v string) *CreateExportMigrationRequest
	GetExportMode() *string
	SetExportObjectStatus(v string) *CreateExportMigrationRequest
	GetExportObjectStatus() *string
	SetIncrementalSince(v int64) *CreateExportMigrationRequest
	GetIncrementalSince() *int64
	SetName(v string) *CreateExportMigrationRequest
	GetName() *string
	SetProjectId(v int64) *CreateExportMigrationRequest
	GetProjectId() *int64
}

type CreateExportMigrationRequest struct {
	// The description of the export task.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The export mode of the export task. Valid values:
	//
	// 	- FULL: The export task is used to export all data objects.
	//
	// 	- INCREMENTAL: The export task is used to export data objects that were modified since the specified point in time. If you set this parameter to INCREMENTAL, you must configure the IncrementalSince parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// FULL
	ExportMode *string `json:"ExportMode,omitempty" xml:"ExportMode,omitempty"`
	// The status of the data objects that you want to export in the export task. The system exports data objects in the state that is specified by this parameter. Valid values:
	//
	// 	- SAVED
	//
	// 	- SUBMITTED
	//
	// 	- DEPLOYED
	//
	// if can be null:
	// true
	//
	// example:
	//
	// SAVED
	ExportObjectStatus *string `json:"ExportObjectStatus,omitempty" xml:"ExportObjectStatus,omitempty"`
	// The start time of the incremental export task.
	//
	// The IncrementalSince parameter takes effect only when the ExportMode parameter is set to INCREMENTAL.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 1589904000000
	IncrementalSince *int64 `json:"IncrementalSince,omitempty" xml:"IncrementalSince,omitempty"`
	// The name of the export task.
	//
	// The name of each export task must be unique. You must ensure that no duplicate export task exists in the current workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_export_01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to obtain the workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreateExportMigrationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateExportMigrationRequest) GoString() string {
	return s.String()
}

func (s *CreateExportMigrationRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateExportMigrationRequest) GetExportMode() *string {
	return s.ExportMode
}

func (s *CreateExportMigrationRequest) GetExportObjectStatus() *string {
	return s.ExportObjectStatus
}

func (s *CreateExportMigrationRequest) GetIncrementalSince() *int64 {
	return s.IncrementalSince
}

func (s *CreateExportMigrationRequest) GetName() *string {
	return s.Name
}

func (s *CreateExportMigrationRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateExportMigrationRequest) SetDescription(v string) *CreateExportMigrationRequest {
	s.Description = &v
	return s
}

func (s *CreateExportMigrationRequest) SetExportMode(v string) *CreateExportMigrationRequest {
	s.ExportMode = &v
	return s
}

func (s *CreateExportMigrationRequest) SetExportObjectStatus(v string) *CreateExportMigrationRequest {
	s.ExportObjectStatus = &v
	return s
}

func (s *CreateExportMigrationRequest) SetIncrementalSince(v int64) *CreateExportMigrationRequest {
	s.IncrementalSince = &v
	return s
}

func (s *CreateExportMigrationRequest) SetName(v string) *CreateExportMigrationRequest {
	s.Name = &v
	return s
}

func (s *CreateExportMigrationRequest) SetProjectId(v int64) *CreateExportMigrationRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateExportMigrationRequest) Validate() error {
	return dara.Validate(s)
}
