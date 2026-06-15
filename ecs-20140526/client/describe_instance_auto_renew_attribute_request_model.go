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
	// The instance IDs. You can specify up to 100 subscription instances at a time. Separate multiple instance IDs with commas.
	//
	// > You must specify either `InstanceId` or `RenewalStatus`.
	//
	// example:
	//
	// i-bp18x3z4hc7bixhx****,i-bp1g6zv0ce8oghu7****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return.
	//
	// Minimum value: 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Maximum value: 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to view the latest list of Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The auto-renewal status of the instance. Valid values:
	//
	// - AutoRenewal: Auto-renewal is enabled.
	//
	// - Normal: Auto-renewal is disabled.
	//
	// - NotRenewal: The instance will not be renewed. The system does not send expiration reminders but sends a non-renewal reminder three days before the expiration date. To renew an ECS instance with this status, you must first call [ModifyInstanceAutoRenewAttribute](https://help.aliyun.com/document_detail/52843.html) to change its status to `Normal`. You can then manually renew the instance or enable auto-renewal.
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
