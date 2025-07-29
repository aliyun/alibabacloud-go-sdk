// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOverviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *GetOverviewRequest
	GetEndTime() *int64
	SetGroupId(v string) *GetOverviewRequest
	GetGroupId() *string
	SetMetricType(v int32) *GetOverviewRequest
	GetMetricType() *int32
	SetNamespace(v string) *GetOverviewRequest
	GetNamespace() *string
	SetNamespaceSource(v string) *GetOverviewRequest
	GetNamespaceSource() *string
	SetOperate(v string) *GetOverviewRequest
	GetOperate() *string
	SetRegionId(v string) *GetOverviewRequest
	GetRegionId() *string
	SetStartTime(v int64) *GetOverviewRequest
	GetStartTime() *int64
}

type GetOverviewRequest struct {
	// example:
	//
	// 1684166400
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	MetricType *int32 `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// query
	Operate *string `json:"Operate,omitempty" xml:"Operate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1684166400
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetOverviewRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOverviewRequest) GoString() string {
	return s.String()
}

func (s *GetOverviewRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetOverviewRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *GetOverviewRequest) GetMetricType() *int32 {
	return s.MetricType
}

func (s *GetOverviewRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetOverviewRequest) GetNamespaceSource() *string {
	return s.NamespaceSource
}

func (s *GetOverviewRequest) GetOperate() *string {
	return s.Operate
}

func (s *GetOverviewRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetOverviewRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetOverviewRequest) SetEndTime(v int64) *GetOverviewRequest {
	s.EndTime = &v
	return s
}

func (s *GetOverviewRequest) SetGroupId(v string) *GetOverviewRequest {
	s.GroupId = &v
	return s
}

func (s *GetOverviewRequest) SetMetricType(v int32) *GetOverviewRequest {
	s.MetricType = &v
	return s
}

func (s *GetOverviewRequest) SetNamespace(v string) *GetOverviewRequest {
	s.Namespace = &v
	return s
}

func (s *GetOverviewRequest) SetNamespaceSource(v string) *GetOverviewRequest {
	s.NamespaceSource = &v
	return s
}

func (s *GetOverviewRequest) SetOperate(v string) *GetOverviewRequest {
	s.Operate = &v
	return s
}

func (s *GetOverviewRequest) SetRegionId(v string) *GetOverviewRequest {
	s.RegionId = &v
	return s
}

func (s *GetOverviewRequest) SetStartTime(v int64) *GetOverviewRequest {
	s.StartTime = &v
	return s
}

func (s *GetOverviewRequest) Validate() error {
	return dara.Validate(s)
}
