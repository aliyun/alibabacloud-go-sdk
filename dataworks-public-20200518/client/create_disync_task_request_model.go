// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDISyncTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateDISyncTaskRequest
	GetClientToken() *string
	SetProjectId(v int64) *CreateDISyncTaskRequest
	GetProjectId() *int64
	SetTaskContent(v string) *CreateDISyncTaskRequest
	GetTaskContent() *string
	SetTaskName(v string) *CreateDISyncTaskRequest
	GetTaskName() *string
	SetTaskParam(v string) *CreateDISyncTaskRequest
	GetTaskParam() *string
	SetTaskType(v string) *CreateDISyncTaskRequest
	GetTaskType() *string
}

type CreateDISyncTaskRequest struct {
	// The client token that is used to ensure the idempotence of the request. This parameter can be left empty.
	//
	// example:
	//
	// 0000-ABCD-EFG****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	TaskContent *string `json:"TaskContent,omitempty" xml:"TaskContent,omitempty"`
	// The name of the data synchronization task.
	//
	// example:
	//
	// new_di_task
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The configuration parameters of the data synchronization task. The following parameters are supported:
	//
	// 	- FileFolderPath: the storage path of the data synchronization task.
	//
	// 	- ResourceGroup: the identifier of the resource group for Data Integration that is used by the data synchronization task. You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/173913.html) operation to query the identifier of the resource group.
	//
	// 	- Cu: the specifications occupied by the data synchronization task in the serverless resource group. The value of this parameter must be a multiple of 0.5.
	//
	// example:
	//
	// {"FileFolderPath":"Business Flow/XXX/Data Integration","ResourceGroup":"S_res_group_XXX_XXXX"}
	TaskParam *string `json:"TaskParam,omitempty" xml:"TaskParam,omitempty"`
	// The type of the data synchronization task. Valid values: DI_OFFLINE, DI_REALTIME, and DI_SOLUTION.
	//
	// This parameter is required.
	//
	// example:
	//
	// DI_OFFLINE
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s CreateDISyncTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDISyncTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDISyncTaskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDISyncTaskRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateDISyncTaskRequest) GetTaskContent() *string {
	return s.TaskContent
}

func (s *CreateDISyncTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateDISyncTaskRequest) GetTaskParam() *string {
	return s.TaskParam
}

func (s *CreateDISyncTaskRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *CreateDISyncTaskRequest) SetClientToken(v string) *CreateDISyncTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDISyncTaskRequest) SetProjectId(v int64) *CreateDISyncTaskRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDISyncTaskRequest) SetTaskContent(v string) *CreateDISyncTaskRequest {
	s.TaskContent = &v
	return s
}

func (s *CreateDISyncTaskRequest) SetTaskName(v string) *CreateDISyncTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateDISyncTaskRequest) SetTaskParam(v string) *CreateDISyncTaskRequest {
	s.TaskParam = &v
	return s
}

func (s *CreateDISyncTaskRequest) SetTaskType(v string) *CreateDISyncTaskRequest {
	s.TaskType = &v
	return s
}

func (s *CreateDISyncTaskRequest) Validate() error {
	return dara.Validate(s)
}
