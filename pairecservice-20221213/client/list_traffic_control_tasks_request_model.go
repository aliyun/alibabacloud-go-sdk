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
	All                  *bool   `json:"All,omitempty" xml:"All,omitempty"`
	ControlTargetFilter  *string `json:"ControlTargetFilter,omitempty" xml:"ControlTargetFilter,omitempty"`
	Environment          *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Order                *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber           *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SceneId              *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	SortBy               *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TrafficControlTaskId *string `json:"TrafficControlTaskId,omitempty" xml:"TrafficControlTaskId,omitempty"`
	Version              *string `json:"Version,omitempty" xml:"Version,omitempty"`
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
