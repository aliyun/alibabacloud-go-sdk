// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DescribeZonesRequest
	GetAcceptLanguage() *string
	SetInstanceChargeType(v string) *DescribeZonesRequest
	GetInstanceChargeType() *string
	SetOwnerAccount(v string) *DescribeZonesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeZonesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeZonesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeZonesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeZonesRequest
	GetResourceOwnerId() *int64
	SetSpotStrategy(v string) *DescribeZonesRequest
	GetSpotStrategy() *string
	SetVerbose(v bool) *DescribeZonesRequest
	GetVerbose() *bool
}

type DescribeZonesRequest struct {
	// The natural language that is used to filter response results. For more information, see [RFC 7231](https://tools.ietf.org/html/rfc7231). Valid values:
	//
	//
	//
	// - zh-CN: simplified Chinese.
	//
	// - zh-TW: traditional Chinese.
	//
	// - en-US: English.
	//
	// - ja: Japanese.
	//
	// - fr: French.
	//
	// - de: German.
	//
	// - ko: Korean.
	//
	// Default value: zh-CN.
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The billing method supported by resources in the active zone. For more information, see [Billing overview](https://help.aliyun.com/document_detail/25398.html). Valid values:
	//
	// - PrePaid: subscription.
	//
	// - PostPaid: pay-as-you-go.
	//
	// Default value: PostPaid.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the zones are located. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The bidding policy for pay-as-you-go instances. You can specify this parameter when `InstanceChargeType=PostPaid`. For more information, see [Spot instances](https://help.aliyun.com/document_detail/52088.html). Valid values:
	//
	//
	//
	// - NoSpot: regular pay-as-you-go instance.
	//
	// - SpotWithPriceLimit: spot instance with a maximum price limit.
	//
	// - SpotAsPriceGo: spot instance priced at the market price with the pay-as-you-go price as the upper limit.
	//
	// Default value: NoSpot.
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// Specifies whether to display detailed information.
	//
	// - true: Display detailed information.
	//
	// - false: Do not display detailed information.
	//
	// Default value: true.
	//
	// example:
	//
	// false
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
}

func (s DescribeZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeZonesRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DescribeZonesRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeZonesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeZonesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeZonesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeZonesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeZonesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeZonesRequest) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *DescribeZonesRequest) GetVerbose() *bool {
	return s.Verbose
}

func (s *DescribeZonesRequest) SetAcceptLanguage(v string) *DescribeZonesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeZonesRequest) SetInstanceChargeType(v string) *DescribeZonesRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeZonesRequest) SetOwnerAccount(v string) *DescribeZonesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeZonesRequest) SetOwnerId(v int64) *DescribeZonesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeZonesRequest) SetRegionId(v string) *DescribeZonesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeZonesRequest) SetResourceOwnerAccount(v string) *DescribeZonesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeZonesRequest) SetResourceOwnerId(v int64) *DescribeZonesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeZonesRequest) SetSpotStrategy(v string) *DescribeZonesRequest {
	s.SpotStrategy = &v
	return s
}

func (s *DescribeZonesRequest) SetVerbose(v bool) *DescribeZonesRequest {
	s.Verbose = &v
	return s
}

func (s *DescribeZonesRequest) Validate() error {
	return dara.Validate(s)
}
