// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeTasksRequest
	GetInstanceId() *string
	SetInstanceName(v string) *DescribeTasksRequest
	GetInstanceName() *string
	SetInvokeId(v string) *DescribeTasksRequest
	GetInvokeId() *string
	SetLevel(v int32) *DescribeTasksRequest
	GetLevel() *int32
	SetMaxResults(v int32) *DescribeTasksRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeTasksRequest
	GetNextToken() *string
	SetParam(v string) *DescribeTasksRequest
	GetParam() *string
	SetParentTaskId(v string) *DescribeTasksRequest
	GetParentTaskId() *string
	SetResourceIds(v []*string) *DescribeTasksRequest
	GetResourceIds() []*string
	SetTaskIds(v []*string) *DescribeTasksRequest
	GetTaskIds() []*string
	SetTaskStatus(v string) *DescribeTasksRequest
	GetTaskStatus() *string
	SetTaskStatuses(v []*string) *DescribeTasksRequest
	GetTaskStatuses() []*string
	SetTaskType(v string) *DescribeTasksRequest
	GetTaskType() *string
	SetTaskTypes(v []*string) *DescribeTasksRequest
	GetTaskTypes() []*string
}

type DescribeTasksRequest struct {
	// The ID of the cloud phone instance.
	//
	// example:
	//
	// acp-2zecay9ponatdc4m****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the cloud phone instance.
	//
	// example:
	//
	// defaultInstanceName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The ID of the command execution. You can set the value to the last returned request ID.
	//
	// example:
	//
	// B8ED2BA9-0C6A-5643-818F-B5D60A64****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The level of the task. A value of 1 specifies a batch task. A value of 2 specifies an instance-level task.
	//
	// example:
	//
	// 1
	Level *int32 `json:"Level,omitempty" xml:"Level,omitempty"`
	// The maximum number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. If the parameter is left empty, the data is queried from the first entry.
	//
	// example:
	//
	// FFbc8N4E1iOlcSxC+8boa0HHH2LKWbggYUinyrZWvtS1oTrMYCg1HuMLGuftj0****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The extension field.
	//
	// example:
	//
	// param
	Param *string `json:"Param,omitempty" xml:"Param,omitempty"`
	// The ID of the parent task.
	//
	// example:
	//
	// t-iaej5dkbnmivx****
	ParentTaskId *string `json:"ParentTaskId,omitempty" xml:"ParentTaskId,omitempty"`
	// The IDs of the resources.
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// The IDs of the tasks.
	TaskIds []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
	// The state of the task.
	//
	// Valid values:
	//
	// 	- PartFinished: The task is partially successful.
	//
	// 	- Finished: The task is completed.
	//
	// 	- Failed: The task failed.
	//
	// 	- Skipped: The task is skipped.
	//
	// 	- Processing: The task is running.
	//
	// 	- Waiting: The task is in queue.
	//
	// example:
	//
	// Processing
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The status of the tasks.
	TaskStatuses []*string `json:"TaskStatuses,omitempty" xml:"TaskStatuses,omitempty" type:"Repeated"`
	// The type of the task.
	//
	// Valid values:
	//
	// 	- BackupFile: backs up files.
	//
	// 	- StopInstance: stops cloud phone instances.
	//
	// 	- RebootInstance: restarts cloud phone instances.
	//
	// 	- StartApp: starts apps.
	//
	// 	- SendFile: uploads files.
	//
	// 	- RunCommand: sends remote command.
	//
	// 	- RestartApp: restarts apps.
	//
	// 	- ResetInstance: resets cloud phone instances.
	//
	// 	- RecoverFile: recovers files.
	//
	// 	- UninstallApp: uninstalls apps.
	//
	// 	- StopApp: stops apps.
	//
	// 	- Screenshot: takes screenshots.
	//
	// 	- InstallApp: installs apps.
	//
	// 	- FetchFile: downloads files.
	//
	// 	- UpdateGroupImage: replaces images.
	//
	// 	- StartInstance: starts instances.
	//
	// example:
	//
	// StartInstance
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The types of the tasks.
	TaskTypes []*string `json:"TaskTypes,omitempty" xml:"TaskTypes,omitempty" type:"Repeated"`
}

func (s DescribeTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeTasksRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeTasksRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeTasksRequest) GetInvokeId() *string {
	return s.InvokeId
}

func (s *DescribeTasksRequest) GetLevel() *int32 {
	return s.Level
}

func (s *DescribeTasksRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeTasksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeTasksRequest) GetParam() *string {
	return s.Param
}

func (s *DescribeTasksRequest) GetParentTaskId() *string {
	return s.ParentTaskId
}

func (s *DescribeTasksRequest) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *DescribeTasksRequest) GetTaskIds() []*string {
	return s.TaskIds
}

func (s *DescribeTasksRequest) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *DescribeTasksRequest) GetTaskStatuses() []*string {
	return s.TaskStatuses
}

func (s *DescribeTasksRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeTasksRequest) GetTaskTypes() []*string {
	return s.TaskTypes
}

func (s *DescribeTasksRequest) SetInstanceId(v string) *DescribeTasksRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTasksRequest) SetInstanceName(v string) *DescribeTasksRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeTasksRequest) SetInvokeId(v string) *DescribeTasksRequest {
	s.InvokeId = &v
	return s
}

func (s *DescribeTasksRequest) SetLevel(v int32) *DescribeTasksRequest {
	s.Level = &v
	return s
}

func (s *DescribeTasksRequest) SetMaxResults(v int32) *DescribeTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeTasksRequest) SetNextToken(v string) *DescribeTasksRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeTasksRequest) SetParam(v string) *DescribeTasksRequest {
	s.Param = &v
	return s
}

func (s *DescribeTasksRequest) SetParentTaskId(v string) *DescribeTasksRequest {
	s.ParentTaskId = &v
	return s
}

func (s *DescribeTasksRequest) SetResourceIds(v []*string) *DescribeTasksRequest {
	s.ResourceIds = v
	return s
}

func (s *DescribeTasksRequest) SetTaskIds(v []*string) *DescribeTasksRequest {
	s.TaskIds = v
	return s
}

func (s *DescribeTasksRequest) SetTaskStatus(v string) *DescribeTasksRequest {
	s.TaskStatus = &v
	return s
}

func (s *DescribeTasksRequest) SetTaskStatuses(v []*string) *DescribeTasksRequest {
	s.TaskStatuses = v
	return s
}

func (s *DescribeTasksRequest) SetTaskType(v string) *DescribeTasksRequest {
	s.TaskType = &v
	return s
}

func (s *DescribeTasksRequest) SetTaskTypes(v []*string) *DescribeTasksRequest {
	s.TaskTypes = v
	return s
}

func (s *DescribeTasksRequest) Validate() error {
	return dara.Validate(s)
}
