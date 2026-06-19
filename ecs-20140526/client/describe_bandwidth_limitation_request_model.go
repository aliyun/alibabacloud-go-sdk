// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBandwidthLimitationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceChargeType(v string) *DescribeBandwidthLimitationRequest
	GetInstanceChargeType() *string
	SetInstanceType(v string) *DescribeBandwidthLimitationRequest
	GetInstanceType() *string
	SetOperationType(v string) *DescribeBandwidthLimitationRequest
	GetOperationType() *string
	SetOwnerAccount(v string) *DescribeBandwidthLimitationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeBandwidthLimitationRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeBandwidthLimitationRequest
	GetRegionId() *string
	SetResourceId(v string) *DescribeBandwidthLimitationRequest
	GetResourceId() *string
	SetResourceOwnerAccount(v string) *DescribeBandwidthLimitationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeBandwidthLimitationRequest
	GetResourceOwnerId() *int64
	SetSpotStrategy(v string) *DescribeBandwidthLimitationRequest
	GetSpotStrategy() *string
}

type DescribeBandwidthLimitationRequest struct {
	// The billing method of the instance. For more information, see [Billing overview](https://help.aliyun.com/document_detail/25398.html). Valid values:
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
	// The instance type. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs.g5.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The operation type for querying public bandwidth limits. Valid values:
	//
	// - Upgrade: upgrades the public bandwidth.
	//
	// - Downgrade: downgrades the public bandwidth.
	//
	// - Create: creates an ECS instance.
	//
	// Default value: Create.
	//
	// example:
	//
	// Upgrade
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the target region. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource ID.
	//
	// > When you set the OperationType parameter to Upgrade or Downgrade, the ResourceId parameter is required.
	//
	// example:
	//
	// i-bp67acfmxazb4ph***
	ResourceId           *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The bidding strategy for pay-as-you-go instances. Valid values:
	//
	// - NoSpot: a regular pay-as-you-go instance.
	//
	// - SpotWithPriceLimit: a preemptible instance with a user-defined maximum hourly price.
	//
	// - SpotAsPriceGo: a preemptible instance priced at the market price at the time of purchase.
	//
	// Default value: NoSpot.
	//
	// > The SpotStrategy parameter takes effect only when the InstanceChargeType parameter is set to PostPaid.
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
}

func (s DescribeBandwidthLimitationRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBandwidthLimitationRequest) GoString() string {
	return s.String()
}

func (s *DescribeBandwidthLimitationRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeBandwidthLimitationRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeBandwidthLimitationRequest) GetOperationType() *string {
	return s.OperationType
}

func (s *DescribeBandwidthLimitationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeBandwidthLimitationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeBandwidthLimitationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeBandwidthLimitationRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeBandwidthLimitationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeBandwidthLimitationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeBandwidthLimitationRequest) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *DescribeBandwidthLimitationRequest) SetInstanceChargeType(v string) *DescribeBandwidthLimitationRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeBandwidthLimitationRequest) SetInstanceType(v string) *DescribeBandwidthLimitationRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeBandwidthLimitationRequest) SetOperationType(v string) *DescribeBandwidthLimitationRequest {
	s.OperationType = &v
	return s
}

func (s *DescribeBandwidthLimitationRequest) SetOwnerAccount(v string) *DescribeBandwidthLimitationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeBandwidthLimitationRequest) SetOwnerId(v int64) *DescribeBandwidthLimitationRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBandwidthLimitationRequest) SetRegionId(v string) *DescribeBandwidthLimitationRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBandwidthLimitationRequest) SetResourceId(v string) *DescribeBandwidthLimitationRequest {
	s.ResourceId = &v
	return s
}

func (s *DescribeBandwidthLimitationRequest) SetResourceOwnerAccount(v string) *DescribeBandwidthLimitationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeBandwidthLimitationRequest) SetResourceOwnerId(v int64) *DescribeBandwidthLimitationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeBandwidthLimitationRequest) SetSpotStrategy(v string) *DescribeBandwidthLimitationRequest {
	s.SpotStrategy = &v
	return s
}

func (s *DescribeBandwidthLimitationRequest) Validate() error {
	return dara.Validate(s)
}
