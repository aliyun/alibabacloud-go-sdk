// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCreditSeatsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AddCreditSeatsRequest
	GetClientToken() *string
	SetConfigs(v []*AddCreditSeatsRequestConfigs) *AddCreditSeatsRequest
	GetConfigs() []*AddCreditSeatsRequestConfigs
	SetProductCode(v string) *AddCreditSeatsRequest
	GetProductCode() *string
	SetProductType(v string) *AddCreditSeatsRequest
	GetProductType() *string
	SetSeats(v int64) *AddCreditSeatsRequest
	GetSeats() *int64
	SetSubscriptionType(v string) *AddCreditSeatsRequest
	GetSubscriptionType() *string
}

type AddCreditSeatsRequest struct {
	ClientToken *string                         `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Configs     []*AddCreditSeatsRequestConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	ProductCode *string                         `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductType *string                         `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// This parameter is required.
	Seats            *int64  `json:"Seats,omitempty" xml:"Seats,omitempty"`
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
}

func (s AddCreditSeatsRequest) String() string {
	return dara.Prettify(s)
}

func (s AddCreditSeatsRequest) GoString() string {
	return s.String()
}

func (s *AddCreditSeatsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddCreditSeatsRequest) GetConfigs() []*AddCreditSeatsRequestConfigs {
	return s.Configs
}

func (s *AddCreditSeatsRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *AddCreditSeatsRequest) GetProductType() *string {
	return s.ProductType
}

func (s *AddCreditSeatsRequest) GetSeats() *int64 {
	return s.Seats
}

func (s *AddCreditSeatsRequest) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *AddCreditSeatsRequest) SetClientToken(v string) *AddCreditSeatsRequest {
	s.ClientToken = &v
	return s
}

func (s *AddCreditSeatsRequest) SetConfigs(v []*AddCreditSeatsRequestConfigs) *AddCreditSeatsRequest {
	s.Configs = v
	return s
}

func (s *AddCreditSeatsRequest) SetProductCode(v string) *AddCreditSeatsRequest {
	s.ProductCode = &v
	return s
}

func (s *AddCreditSeatsRequest) SetProductType(v string) *AddCreditSeatsRequest {
	s.ProductType = &v
	return s
}

func (s *AddCreditSeatsRequest) SetSeats(v int64) *AddCreditSeatsRequest {
	s.Seats = &v
	return s
}

func (s *AddCreditSeatsRequest) SetSubscriptionType(v string) *AddCreditSeatsRequest {
	s.SubscriptionType = &v
	return s
}

func (s *AddCreditSeatsRequest) Validate() error {
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

type AddCreditSeatsRequestConfigs struct {
	Code  *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AddCreditSeatsRequestConfigs) String() string {
	return dara.Prettify(s)
}

func (s AddCreditSeatsRequestConfigs) GoString() string {
	return s.String()
}

func (s *AddCreditSeatsRequestConfigs) GetCode() *string {
	return s.Code
}

func (s *AddCreditSeatsRequestConfigs) GetValue() *string {
	return s.Value
}

func (s *AddCreditSeatsRequestConfigs) SetCode(v string) *AddCreditSeatsRequestConfigs {
	s.Code = &v
	return s
}

func (s *AddCreditSeatsRequestConfigs) SetValue(v string) *AddCreditSeatsRequestConfigs {
	s.Value = &v
	return s
}

func (s *AddCreditSeatsRequestConfigs) Validate() error {
	return dara.Validate(s)
}
