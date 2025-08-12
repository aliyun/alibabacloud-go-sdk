// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridMonitorTaskListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *DescribeHybridMonitorTaskListRequest
	GetGroupId() *string
	SetIncludeAliyunTask(v bool) *DescribeHybridMonitorTaskListRequest
	GetIncludeAliyunTask() *bool
	SetKeyword(v string) *DescribeHybridMonitorTaskListRequest
	GetKeyword() *string
	SetNamespace(v string) *DescribeHybridMonitorTaskListRequest
	GetNamespace() *string
	SetPageNumber(v int32) *DescribeHybridMonitorTaskListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHybridMonitorTaskListRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeHybridMonitorTaskListRequest
	GetRegionId() *string
	SetTargetUserId(v int64) *DescribeHybridMonitorTaskListRequest
	GetTargetUserId() *int64
	SetTaskId(v string) *DescribeHybridMonitorTaskListRequest
	GetTaskId() *string
	SetTaskType(v string) *DescribeHybridMonitorTaskListRequest
	GetTaskType() *string
}

type DescribeHybridMonitorTaskListRequest struct {
	// The ID of the application group.
	//
	// For information about how to obtain the ID of an application group, see [DescribeMonitorGroups](https://help.aliyun.com/document_detail/115032.html).
	//
	// example:
	//
	// 3607****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Specifies whether the returned result includes metric import tasks for Alibaba Cloud services. Valid values:
	//
	// 	- true (default): The returned result includes metric import tasks for Alibaba Cloud services.
	//
	// 	- false: The returned result excludes metric import tasks for Alibaba Cloud services.
	//
	// example:
	//
	// true
	IncludeAliyunTask *bool `json:"IncludeAliyunTask,omitempty" xml:"IncludeAliyunTask,omitempty"`
	// The keyword that is used for the search.
	//
	// example:
	//
	// task
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The name of the namespace.
	//
	// For information about how to obtain the name of a namespace, see [DescribeHybridMonitorNamespaceList](https://help.aliyun.com/document_detail/428880.html).
	//
	// example:
	//
	// aliyun
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The page number.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Pages start from page 1. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the member account.
	//
	// > This parameter is required only if you use a management account to call this operation to delete the metric import tasks that belong to a member in a resource directory. In this case, the `TaskType` parameter is set to `aliyun_fc`.
	//
	// example:
	//
	// 120886317861****
	TargetUserId *int64 `json:"TargetUserId,omitempty" xml:"TargetUserId,omitempty"`
	// The ID of the metric import task.
	//
	// example:
	//
	// 36****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The type of the metric import task. Valid values:
	//
	// 	- aliyun_fc: metric import tasks for Alibaba Cloud services
	//
	// 	- aliyun_sls: metrics for logs imported from Log Service
	//
	// example:
	//
	// aliyun_fc
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeHybridMonitorTaskListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridMonitorTaskListRequest) GoString() string {
	return s.String()
}

func (s *DescribeHybridMonitorTaskListRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeHybridMonitorTaskListRequest) GetIncludeAliyunTask() *bool {
	return s.IncludeAliyunTask
}

func (s *DescribeHybridMonitorTaskListRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeHybridMonitorTaskListRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeHybridMonitorTaskListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHybridMonitorTaskListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHybridMonitorTaskListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHybridMonitorTaskListRequest) GetTargetUserId() *int64 {
	return s.TargetUserId
}

func (s *DescribeHybridMonitorTaskListRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeHybridMonitorTaskListRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeHybridMonitorTaskListRequest) SetGroupId(v string) *DescribeHybridMonitorTaskListRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeHybridMonitorTaskListRequest) SetIncludeAliyunTask(v bool) *DescribeHybridMonitorTaskListRequest {
	s.IncludeAliyunTask = &v
	return s
}

func (s *DescribeHybridMonitorTaskListRequest) SetKeyword(v string) *DescribeHybridMonitorTaskListRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeHybridMonitorTaskListRequest) SetNamespace(v string) *DescribeHybridMonitorTaskListRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeHybridMonitorTaskListRequest) SetPageNumber(v int32) *DescribeHybridMonitorTaskListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHybridMonitorTaskListRequest) SetPageSize(v int32) *DescribeHybridMonitorTaskListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHybridMonitorTaskListRequest) SetRegionId(v string) *DescribeHybridMonitorTaskListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHybridMonitorTaskListRequest) SetTargetUserId(v int64) *DescribeHybridMonitorTaskListRequest {
	s.TargetUserId = &v
	return s
}

func (s *DescribeHybridMonitorTaskListRequest) SetTaskId(v string) *DescribeHybridMonitorTaskListRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeHybridMonitorTaskListRequest) SetTaskType(v string) *DescribeHybridMonitorTaskListRequest {
	s.TaskType = &v
	return s
}

func (s *DescribeHybridMonitorTaskListRequest) Validate() error {
	return dara.Validate(s)
}
