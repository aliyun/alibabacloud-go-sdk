// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrafficControlTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAll(v bool) *ListTrafficControlTasksRequest
	GetAll() *bool
	SetControlTargetFilter(v string) *ListTrafficControlTasksRequest
	GetControlTargetFilter() *string
	SetEnvironment(v string) *ListTrafficControlTasksRequest
	GetEnvironment() *string
	SetInstanceId(v string) *ListTrafficControlTasksRequest
	GetInstanceId() *string
	SetName(v string) *ListTrafficControlTasksRequest
	GetName() *string
	SetOrder(v string) *ListTrafficControlTasksRequest
	GetOrder() *string
	SetPageNumber(v string) *ListTrafficControlTasksRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListTrafficControlTasksRequest
	GetPageSize() *string
	SetSceneId(v string) *ListTrafficControlTasksRequest
	GetSceneId() *string
	SetSortBy(v string) *ListTrafficControlTasksRequest
	GetSortBy() *string
	SetStatus(v string) *ListTrafficControlTasksRequest
	GetStatus() *string
	SetTrafficControlTaskId(v string) *ListTrafficControlTasksRequest
	GetTrafficControlTaskId() *string
	SetVersion(v string) *ListTrafficControlTasksRequest
	GetVersion() *string
}

type ListTrafficControlTasksRequest struct {
	// Specifies whether to return all results without pagination.
	//
	// example:
	//
	// false
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The filter for traffic control targets. Valid values:
	//
	// - `All`: Returns all traffic control targets.
	//
	// - `Valid`: Returns only active traffic control targets. A traffic control target is active if the current time is within the specified start and end time.
	//
	// - `None`: Does not return any traffic control targets.
	//
	// example:
	//
	// All
	ControlTargetFilter *string `json:"ControlTargetFilter,omitempty" xml:"ControlTargetFilter,omitempty"`
	// The environment. Valid values:
	//
	// - `Daily`: the daily environment
	//
	// - `Pre`: the pre-production environment
	//
	// - `Prod`: the production environment
	//
	// example:
	//
	// Pre
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// The instance ID. You can obtain this ID by calling the [ListInstances](https://icms.alibaba-inc.com/content/learn/pai?l=1\\&m=16768\\&n=4300782) operation.
	//
	// example:
	//
	// pairec-cn-********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the traffic control task.
	//
	// example:
	//
	// task_1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The sort order. Valid values:
	//
	// - `ASC`: ascending order
	//
	// - `DESC`: descending order
	//
	// example:
	//
	// ASC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the scene. You can obtain this ID by calling the [ListScenes](https://help.aliyun.com/document_detail/2402581.html) operation.
	//
	// example:
	//
	// 1
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The field by which to sort the results. Valid values:
	//
	// - `GmtCreateTime`: Sorts the results by creation time.
	//
	// - `GmtModifiedTime`: Sorts the results by modification time.
	//
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The status of the task. Valid values:
	//
	// - `NotRunning`: The task is not running.
	//
	// - `Ready`: The task is ready to run.
	//
	// - `Running`: The task is running.
	//
	// - `Finished`: The task is finished.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the traffic control task.
	//
	// example:
	//
	// 1
	TrafficControlTaskId *string `json:"TrafficControlTaskId,omitempty" xml:"TrafficControlTaskId,omitempty"`
	// The version. Valid values:
	//
	// - `Latest`: The latest version. This is the default value.
	//
	// - `Released`: the released version
	//
	// example:
	//
	// Latest
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListTrafficControlTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTrafficControlTasksRequest) GoString() string {
	return s.String()
}

func (s *ListTrafficControlTasksRequest) GetAll() *bool {
	return s.All
}

func (s *ListTrafficControlTasksRequest) GetControlTargetFilter() *string {
	return s.ControlTargetFilter
}

func (s *ListTrafficControlTasksRequest) GetEnvironment() *string {
	return s.Environment
}

func (s *ListTrafficControlTasksRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListTrafficControlTasksRequest) GetName() *string {
	return s.Name
}

func (s *ListTrafficControlTasksRequest) GetOrder() *string {
	return s.Order
}

func (s *ListTrafficControlTasksRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListTrafficControlTasksRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListTrafficControlTasksRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *ListTrafficControlTasksRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListTrafficControlTasksRequest) GetStatus() *string {
	return s.Status
}

func (s *ListTrafficControlTasksRequest) GetTrafficControlTaskId() *string {
	return s.TrafficControlTaskId
}

func (s *ListTrafficControlTasksRequest) GetVersion() *string {
	return s.Version
}

func (s *ListTrafficControlTasksRequest) SetAll(v bool) *ListTrafficControlTasksRequest {
	s.All = &v
	return s
}

func (s *ListTrafficControlTasksRequest) SetControlTargetFilter(v string) *ListTrafficControlTasksRequest {
	s.ControlTargetFilter = &v
	return s
}

func (s *ListTrafficControlTasksRequest) SetEnvironment(v string) *ListTrafficControlTasksRequest {
	s.Environment = &v
	return s
}

func (s *ListTrafficControlTasksRequest) SetInstanceId(v string) *ListTrafficControlTasksRequest {
	s.InstanceId = &v
	return s
}

func (s *ListTrafficControlTasksRequest) SetName(v string) *ListTrafficControlTasksRequest {
	s.Name = &v
	return s
}

func (s *ListTrafficControlTasksRequest) SetOrder(v string) *ListTrafficControlTasksRequest {
	s.Order = &v
	return s
}

func (s *ListTrafficControlTasksRequest) SetPageNumber(v string) *ListTrafficControlTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTrafficControlTasksRequest) SetPageSize(v string) *ListTrafficControlTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListTrafficControlTasksRequest) SetSceneId(v string) *ListTrafficControlTasksRequest {
	s.SceneId = &v
	return s
}

func (s *ListTrafficControlTasksRequest) SetSortBy(v string) *ListTrafficControlTasksRequest {
	s.SortBy = &v
	return s
}

func (s *ListTrafficControlTasksRequest) SetStatus(v string) *ListTrafficControlTasksRequest {
	s.Status = &v
	return s
}

func (s *ListTrafficControlTasksRequest) SetTrafficControlTaskId(v string) *ListTrafficControlTasksRequest {
	s.TrafficControlTaskId = &v
	return s
}

func (s *ListTrafficControlTasksRequest) SetVersion(v string) *ListTrafficControlTasksRequest {
	s.Version = &v
	return s
}

func (s *ListTrafficControlTasksRequest) Validate() error {
	return dara.Validate(s)
}
