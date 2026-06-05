// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEdgeWorkersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListEdgeWorkersShrinkRequest
	GetEndTime() *string
	SetHiveIdsShrink(v string) *ListEdgeWorkersShrinkRequest
	GetHiveIdsShrink() *string
	SetInstanceIdsShrink(v string) *ListEdgeWorkersShrinkRequest
	GetInstanceIdsShrink() *string
	SetPageNumber(v int32) *ListEdgeWorkersShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListEdgeWorkersShrinkRequest
	GetPageSize() *int32
	SetPlanIdsShrink(v string) *ListEdgeWorkersShrinkRequest
	GetPlanIdsShrink() *string
	SetSpec(v string) *ListEdgeWorkersShrinkRequest
	GetSpec() *string
	SetStartTime(v string) *ListEdgeWorkersShrinkRequest
	GetStartTime() *string
	SetStatusesShrink(v string) *ListEdgeWorkersShrinkRequest
	GetStatusesShrink() *string
}

type ListEdgeWorkersShrinkRequest struct {
	// example:
	//
	// 2025-05-14T15:20:37+08:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// ["hive-4fbf3928d40e43948b98acdb4fb5aaed"]
	HiveIdsShrink *string `json:"HiveIds,omitempty" xml:"HiveIds,omitempty"`
	// example:
	//
	// ew-xxxxxx
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
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
	PlanIdsShrink *string `json:"PlanIds,omitempty" xml:"PlanIds,omitempty"`
	// example:
	//
	// ew.gn8t6xlarge-rb.x1p
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// example:
	//
	// 2026-05-25T06:35:26+08:00
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StatusesShrink *string `json:"Statuses,omitempty" xml:"Statuses,omitempty"`
}

func (s ListEdgeWorkersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeWorkersShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListEdgeWorkersShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListEdgeWorkersShrinkRequest) GetHiveIdsShrink() *string {
	return s.HiveIdsShrink
}

func (s *ListEdgeWorkersShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *ListEdgeWorkersShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListEdgeWorkersShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEdgeWorkersShrinkRequest) GetPlanIdsShrink() *string {
	return s.PlanIdsShrink
}

func (s *ListEdgeWorkersShrinkRequest) GetSpec() *string {
	return s.Spec
}

func (s *ListEdgeWorkersShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListEdgeWorkersShrinkRequest) GetStatusesShrink() *string {
	return s.StatusesShrink
}

func (s *ListEdgeWorkersShrinkRequest) SetEndTime(v string) *ListEdgeWorkersShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *ListEdgeWorkersShrinkRequest) SetHiveIdsShrink(v string) *ListEdgeWorkersShrinkRequest {
	s.HiveIdsShrink = &v
	return s
}

func (s *ListEdgeWorkersShrinkRequest) SetInstanceIdsShrink(v string) *ListEdgeWorkersShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *ListEdgeWorkersShrinkRequest) SetPageNumber(v int32) *ListEdgeWorkersShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEdgeWorkersShrinkRequest) SetPageSize(v int32) *ListEdgeWorkersShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListEdgeWorkersShrinkRequest) SetPlanIdsShrink(v string) *ListEdgeWorkersShrinkRequest {
	s.PlanIdsShrink = &v
	return s
}

func (s *ListEdgeWorkersShrinkRequest) SetSpec(v string) *ListEdgeWorkersShrinkRequest {
	s.Spec = &v
	return s
}

func (s *ListEdgeWorkersShrinkRequest) SetStartTime(v string) *ListEdgeWorkersShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *ListEdgeWorkersShrinkRequest) SetStatusesShrink(v string) *ListEdgeWorkersShrinkRequest {
	s.StatusesShrink = &v
	return s
}

func (s *ListEdgeWorkersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
