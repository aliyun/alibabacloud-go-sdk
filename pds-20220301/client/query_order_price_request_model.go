// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOrderPriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryOrderPriceRequest
	GetCode() *string
	SetInstanceId(v string) *QueryOrderPriceRequest
	GetInstanceId() *string
	SetOrderType(v string) *QueryOrderPriceRequest
	GetOrderType() *string
	SetPackage(v string) *QueryOrderPriceRequest
	GetPackage() *string
	SetPeriod(v int64) *QueryOrderPriceRequest
	GetPeriod() *int64
	SetPeriodUnit(v string) *QueryOrderPriceRequest
	GetPeriodUnit() *string
	SetTotalSize(v int64) *QueryOrderPriceRequest
	GetTotalSize() *int64
	SetUserCount(v int64) *QueryOrderPriceRequest
	GetUserCount() *int64
}

type QueryOrderPriceRequest struct {
	// This parameter is required.
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id,omitempty"`
	// This parameter is required.
	OrderType *string `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// This parameter is required.
	Package *string `json:"package,omitempty" xml:"package,omitempty"`
	// This parameter is required.
	Period *int64 `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	PeriodUnit *string `json:"period_unit,omitempty" xml:"period_unit,omitempty"`
	// This parameter is required.
	TotalSize *int64 `json:"total_size,omitempty" xml:"total_size,omitempty"`
	// This parameter is required.
	UserCount *int64 `json:"user_count,omitempty" xml:"user_count,omitempty"`
}

func (s QueryOrderPriceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryOrderPriceRequest) GoString() string {
	return s.String()
}

func (s *QueryOrderPriceRequest) GetCode() *string {
	return s.Code
}

func (s *QueryOrderPriceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryOrderPriceRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *QueryOrderPriceRequest) GetPackage() *string {
	return s.Package
}

func (s *QueryOrderPriceRequest) GetPeriod() *int64 {
	return s.Period
}

func (s *QueryOrderPriceRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *QueryOrderPriceRequest) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *QueryOrderPriceRequest) GetUserCount() *int64 {
	return s.UserCount
}

func (s *QueryOrderPriceRequest) SetCode(v string) *QueryOrderPriceRequest {
	s.Code = &v
	return s
}

func (s *QueryOrderPriceRequest) SetInstanceId(v string) *QueryOrderPriceRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryOrderPriceRequest) SetOrderType(v string) *QueryOrderPriceRequest {
	s.OrderType = &v
	return s
}

func (s *QueryOrderPriceRequest) SetPackage(v string) *QueryOrderPriceRequest {
	s.Package = &v
	return s
}

func (s *QueryOrderPriceRequest) SetPeriod(v int64) *QueryOrderPriceRequest {
	s.Period = &v
	return s
}

func (s *QueryOrderPriceRequest) SetPeriodUnit(v string) *QueryOrderPriceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *QueryOrderPriceRequest) SetTotalSize(v int64) *QueryOrderPriceRequest {
	s.TotalSize = &v
	return s
}

func (s *QueryOrderPriceRequest) SetUserCount(v int64) *QueryOrderPriceRequest {
	s.UserCount = &v
	return s
}

func (s *QueryOrderPriceRequest) Validate() error {
	return dara.Validate(s)
}
