// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEdgeWorkersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListEdgeWorkersRequest
	GetEndTime() *string
	SetHiveIds(v []*string) *ListEdgeWorkersRequest
	GetHiveIds() []*string
	SetInstanceIds(v []*string) *ListEdgeWorkersRequest
	GetInstanceIds() []*string
	SetPageNumber(v int32) *ListEdgeWorkersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListEdgeWorkersRequest
	GetPageSize() *int32
	SetPlanIds(v []*string) *ListEdgeWorkersRequest
	GetPlanIds() []*string
	SetSpec(v string) *ListEdgeWorkersRequest
	GetSpec() *string
	SetStartTime(v string) *ListEdgeWorkersRequest
	GetStartTime() *string
	SetStatuses(v []*string) *ListEdgeWorkersRequest
	GetStatuses() []*string
}

type ListEdgeWorkersRequest struct {
	// example:
	//
	// 2025-05-14T15:20:37+08:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// ["hive-4fbf3928d40e43948b98acdb4fb5aaed"]
	HiveIds []*string `json:"HiveIds,omitempty" xml:"HiveIds,omitempty" type:"Repeated"`
	// example:
	//
	// ew-xxxxxx
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// ["pk-4fbf3928d40e43948b98acdb4fb5aaed"]
	PlanIds []*string `json:"PlanIds,omitempty" xml:"PlanIds,omitempty" type:"Repeated"`
	// example:
	//
	// ew.gn8t6xlarge-rb.x1p
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// example:
	//
	// 2026-05-25T06:35:26+08:00
	StartTime *string   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Statuses  []*string `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
}

func (s ListEdgeWorkersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeWorkersRequest) GoString() string {
	return s.String()
}

func (s *ListEdgeWorkersRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListEdgeWorkersRequest) GetHiveIds() []*string {
	return s.HiveIds
}

func (s *ListEdgeWorkersRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ListEdgeWorkersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListEdgeWorkersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEdgeWorkersRequest) GetPlanIds() []*string {
	return s.PlanIds
}

func (s *ListEdgeWorkersRequest) GetSpec() *string {
	return s.Spec
}

func (s *ListEdgeWorkersRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListEdgeWorkersRequest) GetStatuses() []*string {
	return s.Statuses
}

func (s *ListEdgeWorkersRequest) SetEndTime(v string) *ListEdgeWorkersRequest {
	s.EndTime = &v
	return s
}

func (s *ListEdgeWorkersRequest) SetHiveIds(v []*string) *ListEdgeWorkersRequest {
	s.HiveIds = v
	return s
}

func (s *ListEdgeWorkersRequest) SetInstanceIds(v []*string) *ListEdgeWorkersRequest {
	s.InstanceIds = v
	return s
}

func (s *ListEdgeWorkersRequest) SetPageNumber(v int32) *ListEdgeWorkersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEdgeWorkersRequest) SetPageSize(v int32) *ListEdgeWorkersRequest {
	s.PageSize = &v
	return s
}

func (s *ListEdgeWorkersRequest) SetPlanIds(v []*string) *ListEdgeWorkersRequest {
	s.PlanIds = v
	return s
}

func (s *ListEdgeWorkersRequest) SetSpec(v string) *ListEdgeWorkersRequest {
	s.Spec = &v
	return s
}

func (s *ListEdgeWorkersRequest) SetStartTime(v string) *ListEdgeWorkersRequest {
	s.StartTime = &v
	return s
}

func (s *ListEdgeWorkersRequest) SetStatuses(v []*string) *ListEdgeWorkersRequest {
	s.Statuses = v
	return s
}

func (s *ListEdgeWorkersRequest) Validate() error {
	return dara.Validate(s)
}
