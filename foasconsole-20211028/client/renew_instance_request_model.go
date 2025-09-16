// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDuration(v int32) *RenewInstanceRequest
	GetDuration() *int32
	SetInstanceId(v string) *RenewInstanceRequest
	GetInstanceId() *string
	SetPricingCycle(v string) *RenewInstanceRequest
	GetPricingCycle() *string
	SetPromotionCode(v string) *RenewInstanceRequest
	GetPromotionCode() *string
	SetRegion(v string) *RenewInstanceRequest
	GetRegion() *string
	SetUsePromotionCode(v bool) *RenewInstanceRequest
	GetUsePromotionCode() *bool
}

type RenewInstanceRequest struct {
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

func (s RenewInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewInstanceRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *RenewInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RenewInstanceRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *RenewInstanceRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *RenewInstanceRequest) GetRegion() *string {
	return s.Region
}

func (s *RenewInstanceRequest) GetUsePromotionCode() *bool {
	return s.UsePromotionCode
}

func (s *RenewInstanceRequest) SetDuration(v int32) *RenewInstanceRequest {
	s.Duration = &v
	return s
}

func (s *RenewInstanceRequest) SetInstanceId(v string) *RenewInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RenewInstanceRequest) SetPricingCycle(v string) *RenewInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *RenewInstanceRequest) SetPromotionCode(v string) *RenewInstanceRequest {
	s.PromotionCode = &v
	return s
}

func (s *RenewInstanceRequest) SetRegion(v string) *RenewInstanceRequest {
	s.Region = &v
	return s
}

func (s *RenewInstanceRequest) SetUsePromotionCode(v bool) *RenewInstanceRequest {
	s.UsePromotionCode = &v
	return s
}

func (s *RenewInstanceRequest) Validate() error {
	return dara.Validate(s)
}
