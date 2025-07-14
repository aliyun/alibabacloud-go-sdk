// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAvailabilityMetricRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppSource(v string) *GetAvailabilityMetricRequest
	GetAppSource() *string
	SetCpuStrategy(v string) *GetAvailabilityMetricRequest
	GetCpuStrategy() *string
	SetLimit(v int64) *GetAvailabilityMetricRequest
	GetLimit() *int64
	SetRegionId(v string) *GetAvailabilityMetricRequest
	GetRegionId() *string
}

type GetAvailabilityMetricRequest struct {
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
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetAvailabilityMetricRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAvailabilityMetricRequest) GoString() string {
	return s.String()
}

func (s *GetAvailabilityMetricRequest) GetAppSource() *string {
	return s.AppSource
}

func (s *GetAvailabilityMetricRequest) GetCpuStrategy() *string {
	return s.CpuStrategy
}

func (s *GetAvailabilityMetricRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *GetAvailabilityMetricRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAvailabilityMetricRequest) SetAppSource(v string) *GetAvailabilityMetricRequest {
	s.AppSource = &v
	return s
}

func (s *GetAvailabilityMetricRequest) SetCpuStrategy(v string) *GetAvailabilityMetricRequest {
	s.CpuStrategy = &v
	return s
}

func (s *GetAvailabilityMetricRequest) SetLimit(v int64) *GetAvailabilityMetricRequest {
	s.Limit = &v
	return s
}

func (s *GetAvailabilityMetricRequest) SetRegionId(v string) *GetAvailabilityMetricRequest {
	s.RegionId = &v
	return s
}

func (s *GetAvailabilityMetricRequest) Validate() error {
	return dara.Validate(s)
}
