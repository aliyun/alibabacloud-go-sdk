// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAsyncTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceGroupId(v string) *DeleteAsyncTaskRequest
	GetResourceGroupId() *string
	SetTaskId(v int32) *DeleteAsyncTaskRequest
	GetTaskId() *int32
}

type DeleteAsyncTaskRequest struct {
	// The ID of the resource group to which the instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the task that you want to delete.
	//
	// >  You can call the [DescribeAsyncTasks](~~DescribeAsyncTasks~~) operation to query the IDs of all asynchronous export tasks.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteAsyncTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAsyncTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteAsyncTaskRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteAsyncTaskRequest) GetTaskId() *int32 {
	return s.TaskId
}

func (s *DeleteAsyncTaskRequest) SetResourceGroupId(v string) *DeleteAsyncTaskRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteAsyncTaskRequest) SetTaskId(v int32) *DeleteAsyncTaskRequest {
	s.TaskId = &v
	return s
}

func (s *DeleteAsyncTaskRequest) Validate() error {
	return dara.Validate(s)
}
