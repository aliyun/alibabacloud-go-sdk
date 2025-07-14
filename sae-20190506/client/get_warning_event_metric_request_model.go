// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWarningEventMetricRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppSource(v string) *GetWarningEventMetricRequest
	GetAppSource() *string
	SetCpuStrategy(v string) *GetWarningEventMetricRequest
	GetCpuStrategy() *string
	SetEndTime(v int64) *GetWarningEventMetricRequest
	GetEndTime() *int64
	SetLimit(v int64) *GetWarningEventMetricRequest
	GetLimit() *int64
	SetRegionId(v string) *GetWarningEventMetricRequest
	GetRegionId() *string
	SetStartTime(v int64) *GetWarningEventMetricRequest
	GetStartTime() *int64
}

type GetWarningEventMetricRequest struct {
	// The SAE application type. Valid values:
	//
	// 	- **micro_service**
	//
	// 	- **web**
	//
	// 	- **job**
	//
	// example:
	//
	// micro_service
	AppSource *string `json:"AppSource,omitempty" xml:"AppSource,omitempty"`
	// The CPU allocation policy. Valid values:
	//
	// 	- **request**: CPU cores are allocated only when a request is initiated.
	//
	// 	- **always**: Fixed CPU cores are always allocated.
	//
	// example:
	//
	// always
	CpuStrategy *string `json:"CpuStrategy,omitempty" xml:"CpuStrategy,omitempty"`
	// The end of the time range to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1675824035951
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of entries to return. Valid values: 0 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1675823135951
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetWarningEventMetricRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWarningEventMetricRequest) GoString() string {
	return s.String()
}

func (s *GetWarningEventMetricRequest) GetAppSource() *string {
	return s.AppSource
}

func (s *GetWarningEventMetricRequest) GetCpuStrategy() *string {
	return s.CpuStrategy
}

func (s *GetWarningEventMetricRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetWarningEventMetricRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *GetWarningEventMetricRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetWarningEventMetricRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetWarningEventMetricRequest) SetAppSource(v string) *GetWarningEventMetricRequest {
	s.AppSource = &v
	return s
}

func (s *GetWarningEventMetricRequest) SetCpuStrategy(v string) *GetWarningEventMetricRequest {
	s.CpuStrategy = &v
	return s
}

func (s *GetWarningEventMetricRequest) SetEndTime(v int64) *GetWarningEventMetricRequest {
	s.EndTime = &v
	return s
}

func (s *GetWarningEventMetricRequest) SetLimit(v int64) *GetWarningEventMetricRequest {
	s.Limit = &v
	return s
}

func (s *GetWarningEventMetricRequest) SetRegionId(v string) *GetWarningEventMetricRequest {
	s.RegionId = &v
	return s
}

func (s *GetWarningEventMetricRequest) SetStartTime(v int64) *GetWarningEventMetricRequest {
	s.StartTime = &v
	return s
}

func (s *GetWarningEventMetricRequest) Validate() error {
	return dara.Validate(s)
}
