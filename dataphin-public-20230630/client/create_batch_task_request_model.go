// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBatchTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommand(v *CreateBatchTaskRequestCreateCommand) *CreateBatchTaskRequest
	GetCreateCommand() *CreateBatchTaskRequestCreateCommand
	SetOpTenantId(v int64) *CreateBatchTaskRequest
	GetOpTenantId() *int64
}

type CreateBatchTaskRequest struct {
	// This parameter is required.
	CreateCommand *CreateBatchTaskRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateBatchTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateBatchTaskRequest) GetCreateCommand() *CreateBatchTaskRequestCreateCommand {
	return s.CreateCommand
}

func (s *CreateBatchTaskRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateBatchTaskRequest) SetCreateCommand(v *CreateBatchTaskRequestCreateCommand) *CreateBatchTaskRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateBatchTaskRequest) SetOpTenantId(v int64) *CreateBatchTaskRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateBatchTaskRequest) Validate() error {
	return dara.Validate(s)
}

type CreateBatchTaskRequestCreateCommand struct {
	// example:
	//
	// mysql_catalog
	DataSourceCatalog *string `json:"DataSourceCatalog,omitempty" xml:"DataSourceCatalog,omitempty"`
	// example:
	//
	// 12131111
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// example:
	//
	// erp
	DataSourceSchema *string `json:"DataSourceSchema,omitempty" xml:"DataSourceSchema,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xx test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// /a/b
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// example:
	//
	// PYTHON3_7
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test111
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10121101
	ProjectId        *int64    `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	PythonModuleList []*string `json:"PythonModuleList,omitempty" xml:"PythonModuleList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	ScheduleType *int32 `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5
	TaskType *int32 `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s CreateBatchTaskRequestCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchTaskRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateBatchTaskRequestCreateCommand) GetDataSourceCatalog() *string {
	return s.DataSourceCatalog
}

func (s *CreateBatchTaskRequestCreateCommand) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *CreateBatchTaskRequestCreateCommand) GetDataSourceSchema() *string {
	return s.DataSourceSchema
}

func (s *CreateBatchTaskRequestCreateCommand) GetDescription() *string {
	return s.Description
}

func (s *CreateBatchTaskRequestCreateCommand) GetDirectory() *string {
	return s.Directory
}

func (s *CreateBatchTaskRequestCreateCommand) GetEngine() *string {
	return s.Engine
}

func (s *CreateBatchTaskRequestCreateCommand) GetName() *string {
	return s.Name
}

func (s *CreateBatchTaskRequestCreateCommand) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateBatchTaskRequestCreateCommand) GetPythonModuleList() []*string {
	return s.PythonModuleList
}

func (s *CreateBatchTaskRequestCreateCommand) GetScheduleType() *int32 {
	return s.ScheduleType
}

func (s *CreateBatchTaskRequestCreateCommand) GetTaskType() *int32 {
	return s.TaskType
}

func (s *CreateBatchTaskRequestCreateCommand) SetDataSourceCatalog(v string) *CreateBatchTaskRequestCreateCommand {
	s.DataSourceCatalog = &v
	return s
}

func (s *CreateBatchTaskRequestCreateCommand) SetDataSourceId(v string) *CreateBatchTaskRequestCreateCommand {
	s.DataSourceId = &v
	return s
}

func (s *CreateBatchTaskRequestCreateCommand) SetDataSourceSchema(v string) *CreateBatchTaskRequestCreateCommand {
	s.DataSourceSchema = &v
	return s
}

func (s *CreateBatchTaskRequestCreateCommand) SetDescription(v string) *CreateBatchTaskRequestCreateCommand {
	s.Description = &v
	return s
}

func (s *CreateBatchTaskRequestCreateCommand) SetDirectory(v string) *CreateBatchTaskRequestCreateCommand {
	s.Directory = &v
	return s
}

func (s *CreateBatchTaskRequestCreateCommand) SetEngine(v string) *CreateBatchTaskRequestCreateCommand {
	s.Engine = &v
	return s
}

func (s *CreateBatchTaskRequestCreateCommand) SetName(v string) *CreateBatchTaskRequestCreateCommand {
	s.Name = &v
	return s
}

func (s *CreateBatchTaskRequestCreateCommand) SetProjectId(v int64) *CreateBatchTaskRequestCreateCommand {
	s.ProjectId = &v
	return s
}

func (s *CreateBatchTaskRequestCreateCommand) SetPythonModuleList(v []*string) *CreateBatchTaskRequestCreateCommand {
	s.PythonModuleList = v
	return s
}

func (s *CreateBatchTaskRequestCreateCommand) SetScheduleType(v int32) *CreateBatchTaskRequestCreateCommand {
	s.ScheduleType = &v
	return s
}

func (s *CreateBatchTaskRequestCreateCommand) SetTaskType(v int32) *CreateBatchTaskRequestCreateCommand {
	s.TaskType = &v
	return s
}

func (s *CreateBatchTaskRequestCreateCommand) Validate() error {
	return dara.Validate(s)
}
