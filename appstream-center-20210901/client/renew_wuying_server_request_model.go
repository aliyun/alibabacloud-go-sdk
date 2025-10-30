// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewWuyingServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *RenewWuyingServerRequest
	GetAutoPay() *bool
	SetPeriod(v int32) *RenewWuyingServerRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *RenewWuyingServerRequest
	GetPeriodUnit() *string
	SetPromotionId(v string) *RenewWuyingServerRequest
	GetPromotionId() *string
	SetWuyingServerId(v string) *RenewWuyingServerRequest
	GetWuyingServerId() *string
}

type RenewWuyingServerRequest struct {
	// Automatic payment.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The renewal duration.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the renewal time.
	//
	// Valid values:
	//
	// 	- Month: month.
	//
	// 	- Year: year.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The discount ID.
	//
	// example:
	//
	// 17440009****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The ID of the workstation.
	//
	// example:
	//
	// ws-0bw2f11****dial
	WuyingServerId *string `json:"WuyingServerId,omitempty" xml:"WuyingServerId,omitempty"`
}

func (s RenewWuyingServerRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewWuyingServerRequest) GoString() string {
	return s.String()
}

func (s *RenewWuyingServerRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *RenewWuyingServerRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *RenewWuyingServerRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *RenewWuyingServerRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *RenewWuyingServerRequest) GetWuyingServerId() *string {
	return s.WuyingServerId
}

func (s *RenewWuyingServerRequest) SetAutoPay(v bool) *RenewWuyingServerRequest {
	s.AutoPay = &v
	return s
}

func (s *RenewWuyingServerRequest) SetPeriod(v int32) *RenewWuyingServerRequest {
	s.Period = &v
	return s
}

func (s *RenewWuyingServerRequest) SetPeriodUnit(v string) *RenewWuyingServerRequest {
	s.PeriodUnit = &v
	return s
}

func (s *RenewWuyingServerRequest) SetPromotionId(v string) *RenewWuyingServerRequest {
	s.PromotionId = &v
	return s
}

func (s *RenewWuyingServerRequest) SetWuyingServerId(v string) *RenewWuyingServerRequest {
	s.WuyingServerId = &v
	return s
}

func (s *RenewWuyingServerRequest) Validate() error {
	return dara.Validate(s)
}
