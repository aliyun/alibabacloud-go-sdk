// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCreditSeatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v bool) *CreateCreditSeatRequest
	GetAutoRenew() *bool
	SetClientToken(v string) *CreateCreditSeatRequest
	GetClientToken() *string
	SetPeriod(v int32) *CreateCreditSeatRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateCreditSeatRequest
	GetPeriodUnit() *string
	SetProductCode(v string) *CreateCreditSeatRequest
	GetProductCode() *string
	SetProductType(v string) *CreateCreditSeatRequest
	GetProductType() *string
	SetSubscriptionConfigs(v []*CreateCreditSeatRequestSubscriptionConfigs) *CreateCreditSeatRequest
	GetSubscriptionConfigs() []*CreateCreditSeatRequestSubscriptionConfigs
	SetSubscriptionType(v string) *CreateCreditSeatRequest
	GetSubscriptionType() *string
}

type CreateCreditSeatRequest struct {
	AutoRenew   *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	Period              *int32                                        `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit          *string                                       `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	ProductCode         *string                                       `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductType         *string                                       `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SubscriptionConfigs []*CreateCreditSeatRequestSubscriptionConfigs `json:"SubscriptionConfigs,omitempty" xml:"SubscriptionConfigs,omitempty" type:"Repeated"`
	SubscriptionType    *string                                       `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
}

func (s CreateCreditSeatRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCreditSeatRequest) GoString() string {
	return s.String()
}

func (s *CreateCreditSeatRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateCreditSeatRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCreditSeatRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateCreditSeatRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateCreditSeatRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *CreateCreditSeatRequest) GetProductType() *string {
	return s.ProductType
}

func (s *CreateCreditSeatRequest) GetSubscriptionConfigs() []*CreateCreditSeatRequestSubscriptionConfigs {
	return s.SubscriptionConfigs
}

func (s *CreateCreditSeatRequest) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *CreateCreditSeatRequest) SetAutoRenew(v bool) *CreateCreditSeatRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateCreditSeatRequest) SetClientToken(v string) *CreateCreditSeatRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCreditSeatRequest) SetPeriod(v int32) *CreateCreditSeatRequest {
	s.Period = &v
	return s
}

func (s *CreateCreditSeatRequest) SetPeriodUnit(v string) *CreateCreditSeatRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateCreditSeatRequest) SetProductCode(v string) *CreateCreditSeatRequest {
	s.ProductCode = &v
	return s
}

func (s *CreateCreditSeatRequest) SetProductType(v string) *CreateCreditSeatRequest {
	s.ProductType = &v
	return s
}

func (s *CreateCreditSeatRequest) SetSubscriptionConfigs(v []*CreateCreditSeatRequestSubscriptionConfigs) *CreateCreditSeatRequest {
	s.SubscriptionConfigs = v
	return s
}

func (s *CreateCreditSeatRequest) SetSubscriptionType(v string) *CreateCreditSeatRequest {
	s.SubscriptionType = &v
	return s
}

func (s *CreateCreditSeatRequest) Validate() error {
	if s.SubscriptionConfigs != nil {
		for _, item := range s.SubscriptionConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateCreditSeatRequestSubscriptionConfigs struct {
	Configs []*CreateCreditSeatRequestSubscriptionConfigsConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// This parameter is required.
	Seats *int64 `json:"Seats,omitempty" xml:"Seats,omitempty"`
}

func (s CreateCreditSeatRequestSubscriptionConfigs) String() string {
	return dara.Prettify(s)
}

func (s CreateCreditSeatRequestSubscriptionConfigs) GoString() string {
	return s.String()
}

func (s *CreateCreditSeatRequestSubscriptionConfigs) GetConfigs() []*CreateCreditSeatRequestSubscriptionConfigsConfigs {
	return s.Configs
}

func (s *CreateCreditSeatRequestSubscriptionConfigs) GetSeats() *int64 {
	return s.Seats
}

func (s *CreateCreditSeatRequestSubscriptionConfigs) SetConfigs(v []*CreateCreditSeatRequestSubscriptionConfigsConfigs) *CreateCreditSeatRequestSubscriptionConfigs {
	s.Configs = v
	return s
}

func (s *CreateCreditSeatRequestSubscriptionConfigs) SetSeats(v int64) *CreateCreditSeatRequestSubscriptionConfigs {
	s.Seats = &v
	return s
}

func (s *CreateCreditSeatRequestSubscriptionConfigs) Validate() error {
	if s.Configs != nil {
		for _, item := range s.Configs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateCreditSeatRequestSubscriptionConfigsConfigs struct {
	Code  *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateCreditSeatRequestSubscriptionConfigsConfigs) String() string {
	return dara.Prettify(s)
}

func (s CreateCreditSeatRequestSubscriptionConfigsConfigs) GoString() string {
	return s.String()
}

func (s *CreateCreditSeatRequestSubscriptionConfigsConfigs) GetCode() *string {
	return s.Code
}

func (s *CreateCreditSeatRequestSubscriptionConfigsConfigs) GetValue() *string {
	return s.Value
}

func (s *CreateCreditSeatRequestSubscriptionConfigsConfigs) SetCode(v string) *CreateCreditSeatRequestSubscriptionConfigsConfigs {
	s.Code = &v
	return s
}

func (s *CreateCreditSeatRequestSubscriptionConfigsConfigs) SetValue(v string) *CreateCreditSeatRequestSubscriptionConfigsConfigs {
	s.Value = &v
	return s
}

func (s *CreateCreditSeatRequestSubscriptionConfigsConfigs) Validate() error {
	return dara.Validate(s)
}
