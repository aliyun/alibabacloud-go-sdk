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
	// The application ID.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The type of the Serverless App Engine (SAE) application.
	//
	// - **micro_service**
	//
	// - **web**
	//
	// - **job**
	//
	// example:
	//
	// micro_service
	AppSource *string `json:"AppSource,omitempty" xml:"AppSource,omitempty"`
	// The change order type.
	//
	// example:
	//
	// CoCreateApp
	CoType *string `json:"CoType,omitempty" xml:"CoType,omitempty"`
	// The CPU allocation strategy. Valid values:
	//
	// - **request**: Allocates CPU resources only when requests are received.
	//
	// - **always**: Always allocates a fixed amount of CPU resources.
	//
	// example:
	//
	// always
	CpuStrategy *string `json:"CpuStrategy,omitempty" xml:"CpuStrategy,omitempty"`
	// The start time of change order creation.
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
	// The field used to sort the query results.
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
