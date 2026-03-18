// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillingInstanceIds(v string) *RenewInstanceRequest
	GetBillingInstanceIds() *string
	SetDuration(v int32) *RenewInstanceRequest
	GetDuration() *int32
	SetInstanceId(v string) *RenewInstanceRequest
	GetInstanceId() *string
	SetPricingCycle(v string) *RenewInstanceRequest
	GetPricingCycle() *string
	SetPromotionOptionNo(v string) *RenewInstanceRequest
	GetPromotionOptionNo() *string
}

type RenewInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ng-dd8933281e46****
	BillingInstanceIds *string `json:"BillingInstanceIds,omitempty" xml:"BillingInstanceIds,omitempty"`
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
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Year
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// example:
	//
	// youhuiquan_12378dfj6
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
}

func (s RenewInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewInstanceRequest) GetBillingInstanceIds() *string {
	return s.BillingInstanceIds
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

func (s *RenewInstanceRequest) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *RenewInstanceRequest) SetBillingInstanceIds(v string) *RenewInstanceRequest {
	s.BillingInstanceIds = &v
	return s
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

func (s *RenewInstanceRequest) SetPromotionOptionNo(v string) *RenewInstanceRequest {
	s.PromotionOptionNo = &v
	return s
}

func (s *RenewInstanceRequest) Validate() error {
	return dara.Validate(s)
}
