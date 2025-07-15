// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIPv6TranslatorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *CreateIPv6TranslatorRequest
	GetAutoPay() *bool
	SetBandwidth(v int32) *CreateIPv6TranslatorRequest
	GetBandwidth() *int32
	SetClientToken(v string) *CreateIPv6TranslatorRequest
	GetClientToken() *string
	SetDuration(v int32) *CreateIPv6TranslatorRequest
	GetDuration() *int32
	SetName(v string) *CreateIPv6TranslatorRequest
	GetName() *string
	SetOwnerAccount(v string) *CreateIPv6TranslatorRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateIPv6TranslatorRequest
	GetOwnerId() *int64
	SetPayType(v string) *CreateIPv6TranslatorRequest
	GetPayType() *string
	SetPricingCycle(v string) *CreateIPv6TranslatorRequest
	GetPricingCycle() *string
	SetRegionId(v string) *CreateIPv6TranslatorRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateIPv6TranslatorRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateIPv6TranslatorRequest
	GetResourceOwnerId() *int64
	SetSpec(v string) *CreateIPv6TranslatorRequest
	GetSpec() *string
}

type CreateIPv6TranslatorRequest struct {
	// Specifies whether to enable automatic payment. Valid values: **true and false**.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The bandwidth of the IPv6 Translation Service instance. Unit: Mbit/s. Valid values: **1*	- to **200**. If you do not specify the bandwidth for the mapping entry, the bandwidth is shared with the mapping entry.
	//
	// > If you do not specify this parameter, the default bandwidth is 10 Mbit/s.
	//
	// example:
	//
	// 10
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// sha111
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The subscription duration.
	//
	// 	- If the billing cycle is **Month**, valid values are **1*	- to **9**.
	//
	// 	- If the billing cycle is **Year**, set the value to **3**.
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The name of the IPv6 Translation Service instance. The default name is the instance ID. It must be 2 to 100 characters in length and must start with a letter. It can contain letters, digits, periods (.), underscores (_), and hyphens (-). It cannot start with `http://` or `https://`.
	//
	// example:
	//
	// ipv6_1
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing method of the IPv6 Translation Service instance. Valid values:
	//
	// 	- **PREPAY**: subscription
	//
	// 	- **POSTPAY**: pay-as-you-go
	//
	// example:
	//
	// PREPAY
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
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
	// The region of the IPv6 Translation Service instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cm-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The specification of the IPv6 Translation Service instance. Set the value to **small**.
	//
	// example:
	//
	// small
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s CreateIPv6TranslatorRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIPv6TranslatorRequest) GoString() string {
	return s.String()
}

func (s *CreateIPv6TranslatorRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateIPv6TranslatorRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *CreateIPv6TranslatorRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateIPv6TranslatorRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *CreateIPv6TranslatorRequest) GetName() *string {
	return s.Name
}

func (s *CreateIPv6TranslatorRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateIPv6TranslatorRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateIPv6TranslatorRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateIPv6TranslatorRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *CreateIPv6TranslatorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateIPv6TranslatorRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateIPv6TranslatorRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateIPv6TranslatorRequest) GetSpec() *string {
	return s.Spec
}

func (s *CreateIPv6TranslatorRequest) SetAutoPay(v bool) *CreateIPv6TranslatorRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateIPv6TranslatorRequest) SetBandwidth(v int32) *CreateIPv6TranslatorRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateIPv6TranslatorRequest) SetClientToken(v string) *CreateIPv6TranslatorRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateIPv6TranslatorRequest) SetDuration(v int32) *CreateIPv6TranslatorRequest {
	s.Duration = &v
	return s
}

func (s *CreateIPv6TranslatorRequest) SetName(v string) *CreateIPv6TranslatorRequest {
	s.Name = &v
	return s
}

func (s *CreateIPv6TranslatorRequest) SetOwnerAccount(v string) *CreateIPv6TranslatorRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateIPv6TranslatorRequest) SetOwnerId(v int64) *CreateIPv6TranslatorRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateIPv6TranslatorRequest) SetPayType(v string) *CreateIPv6TranslatorRequest {
	s.PayType = &v
	return s
}

func (s *CreateIPv6TranslatorRequest) SetPricingCycle(v string) *CreateIPv6TranslatorRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateIPv6TranslatorRequest) SetRegionId(v string) *CreateIPv6TranslatorRequest {
	s.RegionId = &v
	return s
}

func (s *CreateIPv6TranslatorRequest) SetResourceOwnerAccount(v string) *CreateIPv6TranslatorRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateIPv6TranslatorRequest) SetResourceOwnerId(v int64) *CreateIPv6TranslatorRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateIPv6TranslatorRequest) SetSpec(v string) *CreateIPv6TranslatorRequest {
	s.Spec = &v
	return s
}

func (s *CreateIPv6TranslatorRequest) Validate() error {
	return dara.Validate(s)
}
