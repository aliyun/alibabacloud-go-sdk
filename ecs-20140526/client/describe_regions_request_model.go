// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DescribeRegionsRequest
	GetAcceptLanguage() *string
	SetInstanceChargeType(v string) *DescribeRegionsRequest
	GetInstanceChargeType() *string
	SetOwnerAccount(v string) *DescribeRegionsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeRegionsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeRegionsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeRegionsRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *DescribeRegionsRequest
	GetResourceType() *string
}

type DescribeRegionsRequest struct {
	// The natural language that is used to filter responses. For more information, see [RFC 7231](https://tools.ietf.org/html/rfc7231). Valid values:
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
	// The billing method of the instance. For more information, see [Billing overview](https://help.aliyun.com/document_detail/25398.html). Valid values:
	//
	// - PrePaid: subscription. If you set this parameter to PrePaid, confirm that your account supports balance payment or credit payment. Otherwise, the InvalidPayMethod error is returned.
	//
	// - PostPaid: pay-as-you-go.
	//
	// - SpotWithPriceLimit: spot instance with a maximum price limit.
	//
	// - SpotAsPriceGo: spot instance priced at the market price with the pay-as-you-go price as the upper limit.
	//
	// Default value: PostPaid.
	//
	// example:
	//
	// PrePaid
	InstanceChargeType   *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The resource type. Valid values:
	//
	// - instance: ECS instance.
	//
	// - disk: cloud disk.
	//
	// - reservedinstance: reserved instance.
	//
	// - scu: storage capacity unit (SCU).
	//
	// Default value: instance.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DescribeRegionsRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeRegionsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeRegionsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRegionsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeRegionsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeRegionsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeRegionsRequest) SetInstanceChargeType(v string) *DescribeRegionsRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeRegionsRequest) SetOwnerAccount(v string) *DescribeRegionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRegionsRequest) SetOwnerId(v int64) *DescribeRegionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceOwnerAccount(v string) *DescribeRegionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceOwnerId(v int64) *DescribeRegionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceType(v string) *DescribeRegionsRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeRegionsRequest) Validate() error {
	return dara.Validate(s)
}
