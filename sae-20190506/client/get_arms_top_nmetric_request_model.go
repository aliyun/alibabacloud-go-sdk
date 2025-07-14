// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetArmsTopNMetricRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppSource(v string) *GetArmsTopNMetricRequest
	GetAppSource() *string
	SetCpuStrategy(v string) *GetArmsTopNMetricRequest
	GetCpuStrategy() *string
	SetEndTime(v int64) *GetArmsTopNMetricRequest
	GetEndTime() *int64
	SetLimit(v int64) *GetArmsTopNMetricRequest
	GetLimit() *int64
	SetOrderBy(v string) *GetArmsTopNMetricRequest
	GetOrderBy() *string
	SetRegionId(v string) *GetArmsTopNMetricRequest
	GetRegionId() *string
	SetStartTime(v int64) *GetArmsTopNMetricRequest
	GetStartTime() *int64
}

type GetArmsTopNMetricRequest struct {
	// The CPU allocation policy. Valid values:
	//
	// 	- **request**: CPU cores are allocated only when a request is initiated.
	//
	// 	- **always**: Fixed CPU cores are always allocated.
	//
	// example:
	//
	// micro_service
	AppSource *string `json:"AppSource,omitempty" xml:"AppSource,omitempty"`
	// The additional information that is returned. The following limits are imposed on the ID:
	//
	// 	- success: If the call is successful, **success*	- is returned.
	//
	// 	- An error code: If the call fails, an error code is returned.
	//
	// example:
	//
	// always
	CpuStrategy *string `json:"CpuStrategy,omitempty" xml:"CpuStrategy,omitempty"`
	// The SAE application type. Valid values:
	//
	// 	- **micro_service**
	//
	// 	- **web**
	//
	// 	- **job**
	//
	// This parameter is required.
	//
	// example:
	//
	// 1675824035951
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The number of entries to return. Valid values: 0 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// count
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The field based on which you want to sort the returned entries.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The end of the time range to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1675823135951
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetArmsTopNMetricRequest) String() string {
	return dara.Prettify(s)
}

func (s GetArmsTopNMetricRequest) GoString() string {
	return s.String()
}

func (s *GetArmsTopNMetricRequest) GetAppSource() *string {
	return s.AppSource
}

func (s *GetArmsTopNMetricRequest) GetCpuStrategy() *string {
	return s.CpuStrategy
}

func (s *GetArmsTopNMetricRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetArmsTopNMetricRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *GetArmsTopNMetricRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *GetArmsTopNMetricRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetArmsTopNMetricRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetArmsTopNMetricRequest) SetAppSource(v string) *GetArmsTopNMetricRequest {
	s.AppSource = &v
	return s
}

func (s *GetArmsTopNMetricRequest) SetCpuStrategy(v string) *GetArmsTopNMetricRequest {
	s.CpuStrategy = &v
	return s
}

func (s *GetArmsTopNMetricRequest) SetEndTime(v int64) *GetArmsTopNMetricRequest {
	s.EndTime = &v
	return s
}

func (s *GetArmsTopNMetricRequest) SetLimit(v int64) *GetArmsTopNMetricRequest {
	s.Limit = &v
	return s
}

func (s *GetArmsTopNMetricRequest) SetOrderBy(v string) *GetArmsTopNMetricRequest {
	s.OrderBy = &v
	return s
}

func (s *GetArmsTopNMetricRequest) SetRegionId(v string) *GetArmsTopNMetricRequest {
	s.RegionId = &v
	return s
}

func (s *GetArmsTopNMetricRequest) SetStartTime(v int64) *GetArmsTopNMetricRequest {
	s.StartTime = &v
	return s
}

func (s *GetArmsTopNMetricRequest) Validate() error {
	return dara.Validate(s)
}
