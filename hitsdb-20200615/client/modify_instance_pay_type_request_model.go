// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstancePayTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDuration(v int32) *ModifyInstancePayTypeRequest
	GetDuration() *int32
	SetInstanceId(v string) *ModifyInstancePayTypeRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *ModifyInstancePayTypeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyInstancePayTypeRequest
	GetOwnerId() *int64
	SetPayType(v string) *ModifyInstancePayTypeRequest
	GetPayType() *string
	SetPricingCycle(v string) *ModifyInstancePayTypeRequest
	GetPricingCycle() *string
	SetResourceOwnerAccount(v string) *ModifyInstancePayTypeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyInstancePayTypeRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ModifyInstancePayTypeRequest
	GetSecurityToken() *string
}

type ModifyInstancePayTypeRequest struct {
	// The subscription duration of the instance. The parameter is required if the instance is an subscription instance.
	//
	// 	- If PricingCycle is set to Month, set this parameter to an integer that ranges from 1 to 9.
	//
	// 	- If PricingCycle is set to Year, set this parameter to an integer that ranges from 1 to 3.
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ld-bp1z3506imz2f****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- **PREPAY**: subscription.
	//
	// 	- **POSTPAY**: pay-as-you-go.
	//
	// This parameter is required.
	//
	// example:
	//
	// POSTPAY
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The unit of the subscription duration for the instance. Valid values:
	//
	// 	- Month
	//
	// 	- Year
	//
	// example:
	//
	// Month
	PricingCycle         *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyInstancePayTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstancePayTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstancePayTypeRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *ModifyInstancePayTypeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstancePayTypeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyInstancePayTypeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyInstancePayTypeRequest) GetPayType() *string {
	return s.PayType
}

func (s *ModifyInstancePayTypeRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *ModifyInstancePayTypeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyInstancePayTypeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyInstancePayTypeRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyInstancePayTypeRequest) SetDuration(v int32) *ModifyInstancePayTypeRequest {
	s.Duration = &v
	return s
}

func (s *ModifyInstancePayTypeRequest) SetInstanceId(v string) *ModifyInstancePayTypeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstancePayTypeRequest) SetOwnerAccount(v string) *ModifyInstancePayTypeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstancePayTypeRequest) SetOwnerId(v int64) *ModifyInstancePayTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstancePayTypeRequest) SetPayType(v string) *ModifyInstancePayTypeRequest {
	s.PayType = &v
	return s
}

func (s *ModifyInstancePayTypeRequest) SetPricingCycle(v string) *ModifyInstancePayTypeRequest {
	s.PricingCycle = &v
	return s
}

func (s *ModifyInstancePayTypeRequest) SetResourceOwnerAccount(v string) *ModifyInstancePayTypeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstancePayTypeRequest) SetResourceOwnerId(v int64) *ModifyInstancePayTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstancePayTypeRequest) SetSecurityToken(v string) *ModifyInstancePayTypeRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyInstancePayTypeRequest) Validate() error {
	return dara.Validate(s)
}
