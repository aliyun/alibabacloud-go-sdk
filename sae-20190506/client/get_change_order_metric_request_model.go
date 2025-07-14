// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChangeOrderMetricRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetChangeOrderMetricRequest
	GetAppId() *string
	SetAppSource(v string) *GetChangeOrderMetricRequest
	GetAppSource() *string
	SetCoType(v string) *GetChangeOrderMetricRequest
	GetCoType() *string
	SetCpuStrategy(v string) *GetChangeOrderMetricRequest
	GetCpuStrategy() *string
	SetCreateTime(v string) *GetChangeOrderMetricRequest
	GetCreateTime() *string
	SetLimit(v int64) *GetChangeOrderMetricRequest
	GetLimit() *int64
	SetOrderBy(v string) *GetChangeOrderMetricRequest
	GetOrderBy() *string
	SetRegionId(v string) *GetChangeOrderMetricRequest
	GetRegionId() *string
}

type GetChangeOrderMetricRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
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
	CoType    *string `json:"CoType,omitempty" xml:"CoType,omitempty"`
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
	// The start time when the change order was created.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1661152748883
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The number of entries to return. Valid values: 0 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The field based on which you want to sort the returned entries.
	//
	// This parameter is required.
	//
	// example:
	//
	// errorPercent
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetChangeOrderMetricRequest) String() string {
	return dara.Prettify(s)
}

func (s GetChangeOrderMetricRequest) GoString() string {
	return s.String()
}

func (s *GetChangeOrderMetricRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetChangeOrderMetricRequest) GetAppSource() *string {
	return s.AppSource
}

func (s *GetChangeOrderMetricRequest) GetCoType() *string {
	return s.CoType
}

func (s *GetChangeOrderMetricRequest) GetCpuStrategy() *string {
	return s.CpuStrategy
}

func (s *GetChangeOrderMetricRequest) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetChangeOrderMetricRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *GetChangeOrderMetricRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *GetChangeOrderMetricRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetChangeOrderMetricRequest) SetAppId(v string) *GetChangeOrderMetricRequest {
	s.AppId = &v
	return s
}

func (s *GetChangeOrderMetricRequest) SetAppSource(v string) *GetChangeOrderMetricRequest {
	s.AppSource = &v
	return s
}

func (s *GetChangeOrderMetricRequest) SetCoType(v string) *GetChangeOrderMetricRequest {
	s.CoType = &v
	return s
}

func (s *GetChangeOrderMetricRequest) SetCpuStrategy(v string) *GetChangeOrderMetricRequest {
	s.CpuStrategy = &v
	return s
}

func (s *GetChangeOrderMetricRequest) SetCreateTime(v string) *GetChangeOrderMetricRequest {
	s.CreateTime = &v
	return s
}

func (s *GetChangeOrderMetricRequest) SetLimit(v int64) *GetChangeOrderMetricRequest {
	s.Limit = &v
	return s
}

func (s *GetChangeOrderMetricRequest) SetOrderBy(v string) *GetChangeOrderMetricRequest {
	s.OrderBy = &v
	return s
}

func (s *GetChangeOrderMetricRequest) SetRegionId(v string) *GetChangeOrderMetricRequest {
	s.RegionId = &v
	return s
}

func (s *GetChangeOrderMetricRequest) Validate() error {
	return dara.Validate(s)
}
