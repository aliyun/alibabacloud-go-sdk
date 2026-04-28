// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *CreateOrderRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *CreateOrderRequest
	GetAutoRenew() *bool
	SetCode(v string) *CreateOrderRequest
	GetCode() *string
	SetInstanceId(v string) *CreateOrderRequest
	GetInstanceId() *string
	SetOrderType(v string) *CreateOrderRequest
	GetOrderType() *string
	SetPackage(v string) *CreateOrderRequest
	GetPackage() *string
	SetPeriod(v int64) *CreateOrderRequest
	GetPeriod() *int64
	SetPeriodUnit(v string) *CreateOrderRequest
	GetPeriodUnit() *string
	SetTotalSize(v int64) *CreateOrderRequest
	GetTotalSize() *int64
	SetUserCount(v int64) *CreateOrderRequest
	GetUserCount() *int64
}

type CreateOrderRequest struct {
	AutoPay   *bool `json:"auto_pay,omitempty" xml:"auto_pay,omitempty"`
	AutoRenew *bool `json:"auto_renew,omitempty" xml:"auto_renew,omitempty"`
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

func (s CreateOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateOrderRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateOrderRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateOrderRequest) GetCode() *string {
	return s.Code
}

func (s *CreateOrderRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateOrderRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *CreateOrderRequest) GetPackage() *string {
	return s.Package
}

func (s *CreateOrderRequest) GetPeriod() *int64 {
	return s.Period
}

func (s *CreateOrderRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateOrderRequest) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *CreateOrderRequest) GetUserCount() *int64 {
	return s.UserCount
}

func (s *CreateOrderRequest) SetAutoPay(v bool) *CreateOrderRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateOrderRequest) SetAutoRenew(v bool) *CreateOrderRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateOrderRequest) SetCode(v string) *CreateOrderRequest {
	s.Code = &v
	return s
}

func (s *CreateOrderRequest) SetInstanceId(v string) *CreateOrderRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateOrderRequest) SetOrderType(v string) *CreateOrderRequest {
	s.OrderType = &v
	return s
}

func (s *CreateOrderRequest) SetPackage(v string) *CreateOrderRequest {
	s.Package = &v
	return s
}

func (s *CreateOrderRequest) SetPeriod(v int64) *CreateOrderRequest {
	s.Period = &v
	return s
}

func (s *CreateOrderRequest) SetPeriodUnit(v string) *CreateOrderRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateOrderRequest) SetTotalSize(v int64) *CreateOrderRequest {
	s.TotalSize = &v
	return s
}

func (s *CreateOrderRequest) SetUserCount(v int64) *CreateOrderRequest {
	s.UserCount = &v
	return s
}

func (s *CreateOrderRequest) Validate() error {
	return dara.Validate(s)
}
