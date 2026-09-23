// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewCreditSeatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v bool) *RenewCreditSeatRequest
	GetAutoRenew() *bool
	SetClientToken(v string) *RenewCreditSeatRequest
	GetClientToken() *string
	SetPeriod(v int32) *RenewCreditSeatRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *RenewCreditSeatRequest
	GetPeriodUnit() *string
	SetProductCode(v string) *RenewCreditSeatRequest
	GetProductCode() *string
	SetProductType(v string) *RenewCreditSeatRequest
	GetProductType() *string
	SetSubscriptionType(v string) *RenewCreditSeatRequest
	GetSubscriptionType() *string
}

type RenewCreditSeatRequest struct {
	AutoRenew        *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Period           *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit       *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	ProductCode      *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductType      *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
}

func (s RenewCreditSeatRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewCreditSeatRequest) GoString() string {
	return s.String()
}

func (s *RenewCreditSeatRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *RenewCreditSeatRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RenewCreditSeatRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *RenewCreditSeatRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *RenewCreditSeatRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *RenewCreditSeatRequest) GetProductType() *string {
	return s.ProductType
}

func (s *RenewCreditSeatRequest) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *RenewCreditSeatRequest) SetAutoRenew(v bool) *RenewCreditSeatRequest {
	s.AutoRenew = &v
	return s
}

func (s *RenewCreditSeatRequest) SetClientToken(v string) *RenewCreditSeatRequest {
	s.ClientToken = &v
	return s
}

func (s *RenewCreditSeatRequest) SetPeriod(v int32) *RenewCreditSeatRequest {
	s.Period = &v
	return s
}

func (s *RenewCreditSeatRequest) SetPeriodUnit(v string) *RenewCreditSeatRequest {
	s.PeriodUnit = &v
	return s
}

func (s *RenewCreditSeatRequest) SetProductCode(v string) *RenewCreditSeatRequest {
	s.ProductCode = &v
	return s
}

func (s *RenewCreditSeatRequest) SetProductType(v string) *RenewCreditSeatRequest {
	s.ProductType = &v
	return s
}

func (s *RenewCreditSeatRequest) SetSubscriptionType(v string) *RenewCreditSeatRequest {
	s.SubscriptionType = &v
	return s
}

func (s *RenewCreditSeatRequest) Validate() error {
	return dara.Validate(s)
}
