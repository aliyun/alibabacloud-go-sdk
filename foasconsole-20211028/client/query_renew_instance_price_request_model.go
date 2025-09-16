// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRenewInstancePriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDuration(v int32) *QueryRenewInstancePriceRequest
	GetDuration() *int32
	SetInstanceId(v string) *QueryRenewInstancePriceRequest
	GetInstanceId() *string
	SetPricingCycle(v string) *QueryRenewInstancePriceRequest
	GetPricingCycle() *string
	SetPromotionCode(v string) *QueryRenewInstancePriceRequest
	GetPromotionCode() *string
	SetRegion(v string) *QueryRenewInstancePriceRequest
	GetRegion() *string
	SetUsePromotionCode(v bool) *QueryRenewInstancePriceRequest
	GetUsePromotionCode() *bool
}

type QueryRenewInstancePriceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sc_flinkserverless_public_cn-7e22ae5sess
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Month
	PricingCycle  *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region           *string `json:"Region,omitempty" xml:"Region,omitempty"`
	UsePromotionCode *bool   `json:"UsePromotionCode,omitempty" xml:"UsePromotionCode,omitempty"`
}

func (s QueryRenewInstancePriceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRenewInstancePriceRequest) GoString() string {
	return s.String()
}

func (s *QueryRenewInstancePriceRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *QueryRenewInstancePriceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryRenewInstancePriceRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *QueryRenewInstancePriceRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *QueryRenewInstancePriceRequest) GetRegion() *string {
	return s.Region
}

func (s *QueryRenewInstancePriceRequest) GetUsePromotionCode() *bool {
	return s.UsePromotionCode
}

func (s *QueryRenewInstancePriceRequest) SetDuration(v int32) *QueryRenewInstancePriceRequest {
	s.Duration = &v
	return s
}

func (s *QueryRenewInstancePriceRequest) SetInstanceId(v string) *QueryRenewInstancePriceRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryRenewInstancePriceRequest) SetPricingCycle(v string) *QueryRenewInstancePriceRequest {
	s.PricingCycle = &v
	return s
}

func (s *QueryRenewInstancePriceRequest) SetPromotionCode(v string) *QueryRenewInstancePriceRequest {
	s.PromotionCode = &v
	return s
}

func (s *QueryRenewInstancePriceRequest) SetRegion(v string) *QueryRenewInstancePriceRequest {
	s.Region = &v
	return s
}

func (s *QueryRenewInstancePriceRequest) SetUsePromotionCode(v bool) *QueryRenewInstancePriceRequest {
	s.UsePromotionCode = &v
	return s
}

func (s *QueryRenewInstancePriceRequest) Validate() error {
	return dara.Validate(s)
}
