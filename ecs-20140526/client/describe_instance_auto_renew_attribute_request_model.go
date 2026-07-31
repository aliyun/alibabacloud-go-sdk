// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceAutoRenewAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeInstanceAutoRenewAttributeRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DescribeInstanceAutoRenewAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeInstanceAutoRenewAttributeRequest
	GetOwnerId() *int64
	SetPageNumber(v string) *DescribeInstanceAutoRenewAttributeRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeInstanceAutoRenewAttributeRequest
	GetPageSize() *string
	SetRegionId(v string) *DescribeInstanceAutoRenewAttributeRequest
	GetRegionId() *string
	SetRenewalStatus(v string) *DescribeInstanceAutoRenewAttributeRequest
	GetRenewalStatus() *string
	SetResourceOwnerAccount(v string) *DescribeInstanceAutoRenewAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeInstanceAutoRenewAttributeRequest
	GetResourceOwnerId() *int64
}

type DescribeInstanceAutoRenewAttributeRequest struct {
	// The instance ID. You can specify up to 100 subscription instance IDs at a time. Separate multiple instance IDs with commas (,).
	//
	// > You must specify at least one of the InstanceId and RenewalStatus parameters.
	//
	// example:
	//
	// i-bp18x3z4hc7bixhx****,i-bp1g6zv0ce8oghu7****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number of the returned page.
	//
	// Minimum value: 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Unit: entries.
	//
	// Maximum value: 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the instance. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent list of Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The auto-renewal status of ECS instance. Valid values:
	//
	// - AutoRenewal: Auto-renewal is enabled.
	//
	// - Normal: Auto-renewal is disabled.
	//
	// - NotRenewal: ECS instance will not be renewed. The system no longer sends expiration reminders but sends only a non-renewal reminder three days before ECS instance expires. You can call [ModifyInstanceAutoRenewAttribute](https://help.aliyun.com/document_detail/52843.html) to change the renewal status of an instance from NotRenewal to `Normal`, and then manually renew ECS instance or enable auto-renewal for ECS instance.
	//
	// example:
	//
	// AutoRenewal
	RenewalStatus        *string `json:"RenewalStatus,omitempty" xml:"RenewalStatus,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeInstanceAutoRenewAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAutoRenewAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAutoRenewAttributeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceAutoRenewAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeInstanceAutoRenewAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeInstanceAutoRenewAttributeRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeInstanceAutoRenewAttributeRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeInstanceAutoRenewAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceAutoRenewAttributeRequest) GetRenewalStatus() *string {
	return s.RenewalStatus
}

func (s *DescribeInstanceAutoRenewAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeInstanceAutoRenewAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeInstanceAutoRenewAttributeRequest) SetInstanceId(v string) *DescribeInstanceAutoRenewAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeRequest) SetOwnerAccount(v string) *DescribeInstanceAutoRenewAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeRequest) SetOwnerId(v int64) *DescribeInstanceAutoRenewAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeRequest) SetPageNumber(v string) *DescribeInstanceAutoRenewAttributeRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeRequest) SetPageSize(v string) *DescribeInstanceAutoRenewAttributeRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeRequest) SetRegionId(v string) *DescribeInstanceAutoRenewAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeRequest) SetRenewalStatus(v string) *DescribeInstanceAutoRenewAttributeRequest {
	s.RenewalStatus = &v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeRequest) SetResourceOwnerAccount(v string) *DescribeInstanceAutoRenewAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeRequest) SetResourceOwnerId(v int64) *DescribeInstanceAutoRenewAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeRequest) Validate() error {
	return dara.Validate(s)
}
