// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceTrendingReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *GetInstanceTrendingReportRequest
	GetEndTime() *int64
	SetInstanceId(v string) *GetInstanceTrendingReportRequest
	GetInstanceId() *string
	SetStartTime(v int64) *GetInstanceTrendingReportRequest
	GetStartTime() *int64
	SetTimeInterval(v string) *GetInstanceTrendingReportRequest
	GetTimeInterval() *string
}

type GetInstanceTrendingReportRequest struct {
	// example:
	//
	// 1582103299434
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 82b2eaae-ce5c-45f8-8231-f15b6b27e55c
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1582267398628
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 1d
	TimeInterval *string `json:"TimeInterval,omitempty" xml:"TimeInterval,omitempty"`
}

func (s GetInstanceTrendingReportRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceTrendingReportRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceTrendingReportRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetInstanceTrendingReportRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceTrendingReportRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetInstanceTrendingReportRequest) GetTimeInterval() *string {
	return s.TimeInterval
}

func (s *GetInstanceTrendingReportRequest) SetEndTime(v int64) *GetInstanceTrendingReportRequest {
	s.EndTime = &v
	return s
}

func (s *GetInstanceTrendingReportRequest) SetInstanceId(v string) *GetInstanceTrendingReportRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceTrendingReportRequest) SetStartTime(v int64) *GetInstanceTrendingReportRequest {
	s.StartTime = &v
	return s
}

func (s *GetInstanceTrendingReportRequest) SetTimeInterval(v string) *GetInstanceTrendingReportRequest {
	s.TimeInterval = &v
	return s
}

func (s *GetInstanceTrendingReportRequest) Validate() error {
	return dara.Validate(s)
}
