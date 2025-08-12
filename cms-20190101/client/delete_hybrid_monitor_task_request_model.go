// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHybridMonitorTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespace(v string) *DeleteHybridMonitorTaskRequest
	GetNamespace() *string
	SetRegionId(v string) *DeleteHybridMonitorTaskRequest
	GetRegionId() *string
	SetTargetUserId(v string) *DeleteHybridMonitorTaskRequest
	GetTargetUserId() *string
	SetTaskId(v string) *DeleteHybridMonitorTaskRequest
	GetTaskId() *string
}

type DeleteHybridMonitorTaskRequest struct {
	// The name of the namespace.
	//
	// The name can contain uppercase letters, lowercase letters, digits, and hyphens (-).
	//
	// > This parameter is required only if you call this operation to delete metric import tasks for Alibaba Cloud services. In this case, the `TaskType` parameter is set to `aliyun_fc`.
	//
	// example:
	//
	// aliyun
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the member account.
	//
	// > This parameter is required only if you use a management account to call this operation to query metric import tasks that belong to a member in a resource directory. In this case, the `TaskType` parameter is set to `aliyun_fc`.
	//
	// example:
	//
	// 120886317861****
	TargetUserId *string `json:"TargetUserId,omitempty" xml:"TargetUserId,omitempty"`
	// The ID of the metric import task.
	//
	// For information about how to obtain the ID of a metric import task, see [DescribeHybridMonitorTaskList](https://help.aliyun.com/document_detail/428624.html).
	//
	// > This parameter is required only if you call this operation to delete metrics for the logs that are imported from Log Service. In this case, the `TaskType` parameter is set to `aliyun_sls`.
	//
	// example:
	//
	// 36****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteHybridMonitorTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHybridMonitorTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteHybridMonitorTaskRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteHybridMonitorTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteHybridMonitorTaskRequest) GetTargetUserId() *string {
	return s.TargetUserId
}

func (s *DeleteHybridMonitorTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DeleteHybridMonitorTaskRequest) SetNamespace(v string) *DeleteHybridMonitorTaskRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteHybridMonitorTaskRequest) SetRegionId(v string) *DeleteHybridMonitorTaskRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteHybridMonitorTaskRequest) SetTargetUserId(v string) *DeleteHybridMonitorTaskRequest {
	s.TargetUserId = &v
	return s
}

func (s *DeleteHybridMonitorTaskRequest) SetTaskId(v string) *DeleteHybridMonitorTaskRequest {
	s.TaskId = &v
	return s
}

func (s *DeleteHybridMonitorTaskRequest) Validate() error {
	return dara.Validate(s)
}
