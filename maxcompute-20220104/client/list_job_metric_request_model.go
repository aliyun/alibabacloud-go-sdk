// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobMetricRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroup(v string) *ListJobMetricRequest
	GetGroup() *string
	SetMetric(v string) *ListJobMetricRequest
	GetMetric() *string
	SetPeriod(v int64) *ListJobMetricRequest
	GetPeriod() *int64
	SetProject(v []*string) *ListJobMetricRequest
	GetProject() []*string
	SetQuota(v []*string) *ListJobMetricRequest
	GetQuota() []*string
	SetType(v string) *ListJobMetricRequest
	GetType() *string
	SetEndTime(v int64) *ListJobMetricRequest
	GetEndTime() *int64
	SetStartTime(v int64) *ListJobMetricRequest
	GetStartTime() *int64
}

type ListJobMetricRequest struct {
	// Grouping basis.
	//
	// > Available values: project, quota, type, status. Meanings:
	//
	// >- project: Group and aggregate by project;
	//
	// >- quota: Group and aggregate by quota;
	//
	// >- type: Group and aggregate by job type;
	//
	// >- status: Group and aggregate by job status.
	//
	// example:
	//
	// project
	Group *string `json:"group,omitempty" xml:"group,omitempty"`
	// The name of observation metric.
	//
	// example:
	//
	// num
	Metric *string `json:"metric,omitempty" xml:"metric,omitempty"`
	// The monitoring statistical period.Unit:Second(s).
	//
	// example:
	//
	// 3600
	Period *int64 `json:"period,omitempty" xml:"period,omitempty"`
	// The name of MaxCompute projects (for filtering).
	Project []*string `json:"project,omitempty" xml:"project,omitempty" type:"Repeated"`
	// The nickname of computing Quota nickname used by the job (for filtering).
	Quota []*string `json:"quota,omitempty" xml:"quota,omitempty" type:"Repeated"`
	// The type of observation metric.
	//
	// example:
	//
	// total
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The end time for the observation interval.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1735536322
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The start time for the observation interval.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1735534322
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ListJobMetricRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobMetricRequest) GoString() string {
	return s.String()
}

func (s *ListJobMetricRequest) GetGroup() *string {
	return s.Group
}

func (s *ListJobMetricRequest) GetMetric() *string {
	return s.Metric
}

func (s *ListJobMetricRequest) GetPeriod() *int64 {
	return s.Period
}

func (s *ListJobMetricRequest) GetProject() []*string {
	return s.Project
}

func (s *ListJobMetricRequest) GetQuota() []*string {
	return s.Quota
}

func (s *ListJobMetricRequest) GetType() *string {
	return s.Type
}

func (s *ListJobMetricRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListJobMetricRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListJobMetricRequest) SetGroup(v string) *ListJobMetricRequest {
	s.Group = &v
	return s
}

func (s *ListJobMetricRequest) SetMetric(v string) *ListJobMetricRequest {
	s.Metric = &v
	return s
}

func (s *ListJobMetricRequest) SetPeriod(v int64) *ListJobMetricRequest {
	s.Period = &v
	return s
}

func (s *ListJobMetricRequest) SetProject(v []*string) *ListJobMetricRequest {
	s.Project = v
	return s
}

func (s *ListJobMetricRequest) SetQuota(v []*string) *ListJobMetricRequest {
	s.Quota = v
	return s
}

func (s *ListJobMetricRequest) SetType(v string) *ListJobMetricRequest {
	s.Type = &v
	return s
}

func (s *ListJobMetricRequest) SetEndTime(v int64) *ListJobMetricRequest {
	s.EndTime = &v
	return s
}

func (s *ListJobMetricRequest) SetStartTime(v int64) *ListJobMetricRequest {
	s.StartTime = &v
	return s
}

func (s *ListJobMetricRequest) Validate() error {
	return dara.Validate(s)
}
