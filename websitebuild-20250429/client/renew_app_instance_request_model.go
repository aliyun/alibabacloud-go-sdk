// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewAppInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *RenewAppInstanceRequest
	GetBizId() *string
	SetClientToken(v string) *RenewAppInstanceRequest
	GetClientToken() *string
	SetDuration(v int32) *RenewAppInstanceRequest
	GetDuration() *int32
	SetExtend(v string) *RenewAppInstanceRequest
	GetExtend() *string
	SetPaymentType(v string) *RenewAppInstanceRequest
	GetPaymentType() *string
	SetPricingCycle(v string) *RenewAppInstanceRequest
	GetPricingCycle() *string
}

type RenewAppInstanceRequest struct {
	// example:
	//
	// WD20250718165839000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// fdfjafdfadfenjeora
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// 1
	Duration *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Extend   *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// example:
	//
	// PayAsYouGo
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
}

func (s RenewAppInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewAppInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewAppInstanceRequest) GetBizId() *string {
	return s.BizId
}

func (s *RenewAppInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RenewAppInstanceRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *RenewAppInstanceRequest) GetExtend() *string {
	return s.Extend
}

func (s *RenewAppInstanceRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *RenewAppInstanceRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *RenewAppInstanceRequest) SetBizId(v string) *RenewAppInstanceRequest {
	s.BizId = &v
	return s
}

func (s *RenewAppInstanceRequest) SetClientToken(v string) *RenewAppInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RenewAppInstanceRequest) SetDuration(v int32) *RenewAppInstanceRequest {
	s.Duration = &v
	return s
}

func (s *RenewAppInstanceRequest) SetExtend(v string) *RenewAppInstanceRequest {
	s.Extend = &v
	return s
}

func (s *RenewAppInstanceRequest) SetPaymentType(v string) *RenewAppInstanceRequest {
	s.PaymentType = &v
	return s
}

func (s *RenewAppInstanceRequest) SetPricingCycle(v string) *RenewAppInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *RenewAppInstanceRequest) Validate() error {
	return dara.Validate(s)
}
