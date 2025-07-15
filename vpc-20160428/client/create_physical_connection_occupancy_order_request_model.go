// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePhysicalConnectionOccupancyOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *CreatePhysicalConnectionOccupancyOrderRequest
	GetAutoPay() *bool
	SetClientToken(v string) *CreatePhysicalConnectionOccupancyOrderRequest
	GetClientToken() *string
	SetInstanceChargeType(v string) *CreatePhysicalConnectionOccupancyOrderRequest
	GetInstanceChargeType() *string
	SetOwnerAccount(v string) *CreatePhysicalConnectionOccupancyOrderRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreatePhysicalConnectionOccupancyOrderRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *CreatePhysicalConnectionOccupancyOrderRequest
	GetPeriod() *int32
	SetPhysicalConnectionId(v string) *CreatePhysicalConnectionOccupancyOrderRequest
	GetPhysicalConnectionId() *string
	SetPricingCycle(v string) *CreatePhysicalConnectionOccupancyOrderRequest
	GetPricingCycle() *string
	SetRegionId(v string) *CreatePhysicalConnectionOccupancyOrderRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreatePhysicalConnectionOccupancyOrderRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreatePhysicalConnectionOccupancyOrderRequest
	GetResourceOwnerId() *int64
}

type CreatePhysicalConnectionOccupancyOrderRequest struct {
	// Specifies whether to enable automatic payments. Valid values:
	//
	// 	- **true**: yes Make sure that you have a sufficient balance in your account. Otherwise, your order becomes invalid and is automatically canceled.
	//
	// 	- **false**: disables automatic payment. This is the default value.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests.
	//
	// example:
	//
	// CBCE910E-D396-4944
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The billing method. Set the value to
	//
	// **PrePaid**, which specifies the subscription billing method. If you choose this billing method, make sure that your Alibaba Cloud account supports balance payments or credit payments.
	//
	// example:
	//
	// PrePaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The subscription duration.
	//
	// 	- If **PricingCycle*	- is set to **Month**, set **Period*	- to a value from **1 to 9**.
	//
	// 	- If **PricingCycle*	- is set to **Year**, set **Period*	- to a value from **1 to 5**.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The ID of the Express Connect circuit.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-bp1hp0wr072f6****
	PhysicalConnectionId *string `json:"PhysicalConnectionId,omitempty" xml:"PhysicalConnectionId,omitempty"`
	// The billing cycle of the subscription. Valid values:
	//
	// 	- **Month*	- (default)
	//
	// 	- **Year**
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// The region ID of the Express Connect circuit.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreatePhysicalConnectionOccupancyOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePhysicalConnectionOccupancyOrderRequest) GoString() string {
	return s.String()
}

func (s *CreatePhysicalConnectionOccupancyOrderRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreatePhysicalConnectionOccupancyOrderRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreatePhysicalConnectionOccupancyOrderRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *CreatePhysicalConnectionOccupancyOrderRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreatePhysicalConnectionOccupancyOrderRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreatePhysicalConnectionOccupancyOrderRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreatePhysicalConnectionOccupancyOrderRequest) GetPhysicalConnectionId() *string {
	return s.PhysicalConnectionId
}

func (s *CreatePhysicalConnectionOccupancyOrderRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *CreatePhysicalConnectionOccupancyOrderRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePhysicalConnectionOccupancyOrderRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreatePhysicalConnectionOccupancyOrderRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreatePhysicalConnectionOccupancyOrderRequest) SetAutoPay(v bool) *CreatePhysicalConnectionOccupancyOrderRequest {
	s.AutoPay = &v
	return s
}

func (s *CreatePhysicalConnectionOccupancyOrderRequest) SetClientToken(v string) *CreatePhysicalConnectionOccupancyOrderRequest {
	s.ClientToken = &v
	return s
}

func (s *CreatePhysicalConnectionOccupancyOrderRequest) SetInstanceChargeType(v string) *CreatePhysicalConnectionOccupancyOrderRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *CreatePhysicalConnectionOccupancyOrderRequest) SetOwnerAccount(v string) *CreatePhysicalConnectionOccupancyOrderRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreatePhysicalConnectionOccupancyOrderRequest) SetOwnerId(v int64) *CreatePhysicalConnectionOccupancyOrderRequest {
	s.OwnerId = &v
	return s
}

func (s *CreatePhysicalConnectionOccupancyOrderRequest) SetPeriod(v int32) *CreatePhysicalConnectionOccupancyOrderRequest {
	s.Period = &v
	return s
}

func (s *CreatePhysicalConnectionOccupancyOrderRequest) SetPhysicalConnectionId(v string) *CreatePhysicalConnectionOccupancyOrderRequest {
	s.PhysicalConnectionId = &v
	return s
}

func (s *CreatePhysicalConnectionOccupancyOrderRequest) SetPricingCycle(v string) *CreatePhysicalConnectionOccupancyOrderRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreatePhysicalConnectionOccupancyOrderRequest) SetRegionId(v string) *CreatePhysicalConnectionOccupancyOrderRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePhysicalConnectionOccupancyOrderRequest) SetResourceOwnerAccount(v string) *CreatePhysicalConnectionOccupancyOrderRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreatePhysicalConnectionOccupancyOrderRequest) SetResourceOwnerId(v int64) *CreatePhysicalConnectionOccupancyOrderRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreatePhysicalConnectionOccupancyOrderRequest) Validate() error {
	return dara.Validate(s)
}
