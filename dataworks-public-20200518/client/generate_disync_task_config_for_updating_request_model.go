// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateDISyncTaskConfigForUpdatingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GenerateDISyncTaskConfigForUpdatingRequest
	GetClientToken() *string
	SetProjectId(v int64) *GenerateDISyncTaskConfigForUpdatingRequest
	GetProjectId() *int64
	SetTaskId(v int64) *GenerateDISyncTaskConfigForUpdatingRequest
	GetTaskId() *int64
	SetTaskParam(v string) *GenerateDISyncTaskConfigForUpdatingRequest
	GetTaskParam() *string
	SetTaskType(v string) *GenerateDISyncTaskConfigForUpdatingRequest
	GetTaskType() *string
}

type GenerateDISyncTaskConfigForUpdatingRequest struct {
	// The client token that is used to ensure the idempotence of the request. This parameter is used to prevent repeated operations that are caused by multiple calls.
	//
	// example:
	//
	// ABFUOEUOTRTRJKE
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to obtain the workspace ID.
	//
	// You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The task ID.
	//
	// 	- If you set the TaskType parameter to DI_REALTIME, set the TaskId parameter to the value of the FileId parameter for the real-time synchronization task.
	//
	// 	- If you set the TaskType parameter to DI_SOLUTION, set the TaskId parameter to the value of the FileId parameter for the synchronization solution.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The script for updating the real-time synchronization task in Data Integration. DataWorks allows you to add or remove tables for a real-time synchronization task in Data Integration only in asynchronous mode. The following types of real-time synchronization tasks are supported:
	//
	// 	- Synchronization task that is used to synchronize data from MySQL to MaxCompute
	//
	// 	- Synchronization task that is used to synchronize data from MySQL data to Kafka
	//
	// 	- Synchronization task that is used to synchronize data from MySQL to Hologres
	//
	// The SelectedTables parameter is used to specify tables that you want to synchronize from multiple databases. The Tables parameter is used to specify tables that you want to synchronize from a single database.
	//
	// 	- If the script contains the SelectedTables parameter, the system synchronizes data from the tables that you specify in the SelectedTables parameter.
	//
	// 	- If the script contains the Tables parameter, the system synchronizes data from the tables that you specify in the Tables parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// {      "steps": [         {             "parameter": {                 "connection": [                     {                         "table": [                             "xyx"                         ]                     }                 ]             },             "name": "Reader",             "category": "reader"         }     ] }
	TaskParam *string `json:"TaskParam,omitempty" xml:"TaskParam,omitempty"`
	// The type of the object that you want to update in Data Integration in asynchronous mode. Valid values:
	//
	// 	- DI_REALTIME: real-time synchronization task
	//
	// 	- DI_SOLUTION: synchronization solution DataWorks allows you to create or update real-time synchronization tasks in Data Integration only in asynchronous mode.
	//
	// Valid values:
	//
	// 	- DI_OFFLINE
	//
	// 	- DI_REALTIME
	//
	// 	- DI_SOLUTION
	//
	// This parameter is required.
	//
	// example:
	//
	// DI_REALTIME
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GenerateDISyncTaskConfigForUpdatingRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateDISyncTaskConfigForUpdatingRequest) GoString() string {
	return s.String()
}

func (s *GenerateDISyncTaskConfigForUpdatingRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GenerateDISyncTaskConfigForUpdatingRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GenerateDISyncTaskConfigForUpdatingRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GenerateDISyncTaskConfigForUpdatingRequest) GetTaskParam() *string {
	return s.TaskParam
}

func (s *GenerateDISyncTaskConfigForUpdatingRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *GenerateDISyncTaskConfigForUpdatingRequest) SetClientToken(v string) *GenerateDISyncTaskConfigForUpdatingRequest {
	s.ClientToken = &v
	return s
}

func (s *GenerateDISyncTaskConfigForUpdatingRequest) SetProjectId(v int64) *GenerateDISyncTaskConfigForUpdatingRequest {
	s.ProjectId = &v
	return s
}

func (s *GenerateDISyncTaskConfigForUpdatingRequest) SetTaskId(v int64) *GenerateDISyncTaskConfigForUpdatingRequest {
	s.TaskId = &v
	return s
}

func (s *GenerateDISyncTaskConfigForUpdatingRequest) SetTaskParam(v string) *GenerateDISyncTaskConfigForUpdatingRequest {
	s.TaskParam = &v
	return s
}

func (s *GenerateDISyncTaskConfigForUpdatingRequest) SetTaskType(v string) *GenerateDISyncTaskConfigForUpdatingRequest {
	s.TaskType = &v
	return s
}

func (s *GenerateDISyncTaskConfigForUpdatingRequest) Validate() error {
	return dara.Validate(s)
}
