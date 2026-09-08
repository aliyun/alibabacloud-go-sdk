// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCenBandwidthPackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *CreateCenBandwidthPackageRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *CreateCenBandwidthPackageRequest
	GetAutoRenew() *bool
	SetAutoRenewDuration(v int32) *CreateCenBandwidthPackageRequest
	GetAutoRenewDuration() *int32
	SetBandwidth(v int32) *CreateCenBandwidthPackageRequest
	GetBandwidth() *int32
	SetBandwidthPackageChargeType(v string) *CreateCenBandwidthPackageRequest
	GetBandwidthPackageChargeType() *string
	SetClientToken(v string) *CreateCenBandwidthPackageRequest
	GetClientToken() *string
	SetDescription(v string) *CreateCenBandwidthPackageRequest
	GetDescription() *string
	SetGeographicRegionAId(v string) *CreateCenBandwidthPackageRequest
	GetGeographicRegionAId() *string
	SetGeographicRegionBId(v string) *CreateCenBandwidthPackageRequest
	GetGeographicRegionBId() *string
	SetName(v string) *CreateCenBandwidthPackageRequest
	GetName() *string
	SetOwnerAccount(v string) *CreateCenBandwidthPackageRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateCenBandwidthPackageRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *CreateCenBandwidthPackageRequest
	GetPeriod() *int32
	SetPricingCycle(v string) *CreateCenBandwidthPackageRequest
	GetPricingCycle() *string
	SetResourceOwnerAccount(v string) *CreateCenBandwidthPackageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateCenBandwidthPackageRequest
	GetResourceOwnerId() *int64
	SetTag(v []*CreateCenBandwidthPackageRequestTag) *CreateCenBandwidthPackageRequest
	GetTag() []*CreateCenBandwidthPackageRequestTag
}

type CreateCenBandwidthPackageRequest struct {
	// Specifies whether to enable automatic payment for the bill of the bandwidth plan instance. Valid values:
	//
	// - **true**: enables automatic payment.
	//
	// - **false*	- (default): disables automatic payment.
	//
	// If you set this parameter to false, go to the Order Center in the console to complete the payment after you invoke this operation. Otherwise, the instance cannot be created.
	//
	// example:
	//
	// false
	AutoPay           *bool  `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoRenew         *bool  `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoRenewDuration *int32 `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	// The maximum bandwidth value of the bandwidth plan. Unit: Mbit/s. Valid values: **2*	- to **10000**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The billing method of the bandwidth plan. Valid values: **PREPAY**, which indicates the subscription billing method.
	//
	// example:
	//
	// PREPAY
	BandwidthPackageChargeType *string `json:"BandwidthPackageChargeType,omitempty" xml:"BandwidthPackageChargeType,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that the value is unique among different requests. The token can be up to 64 ASCII characters in length.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the bandwidth plan.
	//
	// The description can be empty or 1 to 256 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// namedesc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The area to which the network instance belongs. Valid values:
	//
	// - **China**: the Chinese mainland.
	//
	// - **North-America**: North America.
	//
	// - **Asia-Pacific**: Asia Pacific.
	//
	// - **Europe**: Europe.
	//
	// This parameter is required.
	//
	// example:
	//
	// China
	GeographicRegionAId *string `json:"GeographicRegionAId,omitempty" xml:"GeographicRegionAId,omitempty"`
	// The area to which the other network instance belongs. Valid values:
	//
	// - **China**: the Chinese mainland.
	//
	// - **North-America**: North America.
	//
	// - **Asia-Pacific**: Asia Pacific.
	//
	// - **Europe**: Europe.
	//
	// This parameter is required.
	//
	// example:
	//
	// China
	GeographicRegionBId *string `json:"GeographicRegionBId,omitempty" xml:"GeographicRegionBId,omitempty"`
	// The name of the bandwidth plan.
	//
	// The name can be empty or 1 to 128 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// nametest
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The subscription duration of the bandwidth plan. Default value: 1.
	//
	// - If **PricingCycle*	- is set to **Month**, valid values for **Period*	- are **1*	- to **3*	- and **6**.
	//
	// - If **PricingCycle*	- is set to **Year**, valid values for **Period*	- are **1*	- to **3**.
	//
	// > This parameter is required when **BandwidthPackageChargeType*	- is set to **PREPAY**.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The billing cycle of the bandwidth plan. Valid values:
	//
	// - **Month*	- (default): billed on a monthly basis.
	//
	// - **Year**: billed on a yearly basis.
	//
	// example:
	//
	// Month
	PricingCycle         *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tag information.
	//
	// You can specify up to 20 tags at a time.
	Tag []*CreateCenBandwidthPackageRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateCenBandwidthPackageRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCenBandwidthPackageRequest) GoString() string {
	return s.String()
}

func (s *CreateCenBandwidthPackageRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateCenBandwidthPackageRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateCenBandwidthPackageRequest) GetAutoRenewDuration() *int32 {
	return s.AutoRenewDuration
}

func (s *CreateCenBandwidthPackageRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *CreateCenBandwidthPackageRequest) GetBandwidthPackageChargeType() *string {
	return s.BandwidthPackageChargeType
}

func (s *CreateCenBandwidthPackageRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCenBandwidthPackageRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateCenBandwidthPackageRequest) GetGeographicRegionAId() *string {
	return s.GeographicRegionAId
}

func (s *CreateCenBandwidthPackageRequest) GetGeographicRegionBId() *string {
	return s.GeographicRegionBId
}

func (s *CreateCenBandwidthPackageRequest) GetName() *string {
	return s.Name
}

func (s *CreateCenBandwidthPackageRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateCenBandwidthPackageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateCenBandwidthPackageRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateCenBandwidthPackageRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *CreateCenBandwidthPackageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateCenBandwidthPackageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateCenBandwidthPackageRequest) GetTag() []*CreateCenBandwidthPackageRequestTag {
	return s.Tag
}

func (s *CreateCenBandwidthPackageRequest) SetAutoPay(v bool) *CreateCenBandwidthPackageRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateCenBandwidthPackageRequest) SetAutoRenew(v bool) *CreateCenBandwidthPackageRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateCenBandwidthPackageRequest) SetAutoRenewDuration(v int32) *CreateCenBandwidthPackageRequest {
	s.AutoRenewDuration = &v
	return s
}

func (s *CreateCenBandwidthPackageRequest) SetBandwidth(v int32) *CreateCenBandwidthPackageRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateCenBandwidthPackageRequest) SetBandwidthPackageChargeType(v string) *CreateCenBandwidthPackageRequest {
	s.BandwidthPackageChargeType = &v
	return s
}

func (s *CreateCenBandwidthPackageRequest) SetClientToken(v string) *CreateCenBandwidthPackageRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCenBandwidthPackageRequest) SetDescription(v string) *CreateCenBandwidthPackageRequest {
	s.Description = &v
	return s
}

func (s *CreateCenBandwidthPackageRequest) SetGeographicRegionAId(v string) *CreateCenBandwidthPackageRequest {
	s.GeographicRegionAId = &v
	return s
}

func (s *CreateCenBandwidthPackageRequest) SetGeographicRegionBId(v string) *CreateCenBandwidthPackageRequest {
	s.GeographicRegionBId = &v
	return s
}

func (s *CreateCenBandwidthPackageRequest) SetName(v string) *CreateCenBandwidthPackageRequest {
	s.Name = &v
	return s
}

func (s *CreateCenBandwidthPackageRequest) SetOwnerAccount(v string) *CreateCenBandwidthPackageRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateCenBandwidthPackageRequest) SetOwnerId(v int64) *CreateCenBandwidthPackageRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCenBandwidthPackageRequest) SetPeriod(v int32) *CreateCenBandwidthPackageRequest {
	s.Period = &v
	return s
}

func (s *CreateCenBandwidthPackageRequest) SetPricingCycle(v string) *CreateCenBandwidthPackageRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateCenBandwidthPackageRequest) SetResourceOwnerAccount(v string) *CreateCenBandwidthPackageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateCenBandwidthPackageRequest) SetResourceOwnerId(v int64) *CreateCenBandwidthPackageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateCenBandwidthPackageRequest) SetTag(v []*CreateCenBandwidthPackageRequestTag) *CreateCenBandwidthPackageRequest {
	s.Tag = v
	return s
}

func (s *CreateCenBandwidthPackageRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateCenBandwidthPackageRequestTag struct {
	// The tag key of the resource.
	//
	// The tag key cannot be an empty string. The tag key can be up to 64 characters in length and cannot start with `aliyun` or `acs:`. The tag key cannot contain `http://` or `https://`.
	//
	// You can specify up to 20 tag keys at a time.
	//
	// example:
	//
	// tagtest
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource.
	//
	// The tag value cannot be empty. The tag value can be up to 128 characters in length and cannot start with aliyun or acs:. The tag value cannot contain http:// or https://.
	//
	// Each tag key has a unique tag value. You can specify up to 20 tag values at a time.
	//
	// example:
	//
	// tagtest
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateCenBandwidthPackageRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateCenBandwidthPackageRequestTag) GoString() string {
	return s.String()
}

func (s *CreateCenBandwidthPackageRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateCenBandwidthPackageRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateCenBandwidthPackageRequestTag) SetKey(v string) *CreateCenBandwidthPackageRequestTag {
	s.Key = &v
	return s
}

func (s *CreateCenBandwidthPackageRequestTag) SetValue(v string) *CreateCenBandwidthPackageRequestTag {
	s.Value = &v
	return s
}

func (s *CreateCenBandwidthPackageRequestTag) Validate() error {
	return dara.Validate(s)
}
