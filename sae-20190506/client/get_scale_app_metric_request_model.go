// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScaleAppMetricRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppSource(v string) *GetScaleAppMetricRequest
	GetAppSource() *string
	SetCpuStrategy(v string) *GetScaleAppMetricRequest
	GetCpuStrategy() *string
	SetLimit(v int64) *GetScaleAppMetricRequest
	GetLimit() *int64
	SetRegionId(v string) *GetScaleAppMetricRequest
	GetRegionId() *string
}

type GetScaleAppMetricRequest struct {
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

func (s GetScaleAppMetricRequest) String() string {
	return dara.Prettify(s)
}

func (s GetScaleAppMetricRequest) GoString() string {
	return s.String()
}

func (s *GetScaleAppMetricRequest) GetAppSource() *string {
	return s.AppSource
}

func (s *GetScaleAppMetricRequest) GetCpuStrategy() *string {
	return s.CpuStrategy
}

func (s *GetScaleAppMetricRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *GetScaleAppMetricRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetScaleAppMetricRequest) SetAppSource(v string) *GetScaleAppMetricRequest {
	s.AppSource = &v
	return s
}

func (s *GetScaleAppMetricRequest) SetCpuStrategy(v string) *GetScaleAppMetricRequest {
	s.CpuStrategy = &v
	return s
}

func (s *GetScaleAppMetricRequest) SetLimit(v int64) *GetScaleAppMetricRequest {
	s.Limit = &v
	return s
}

func (s *GetScaleAppMetricRequest) SetRegionId(v string) *GetScaleAppMetricRequest {
	s.RegionId = &v
	return s
}

func (s *GetScaleAppMetricRequest) Validate() error {
	return dara.Validate(s)
}
