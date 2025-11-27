// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancesByExpireTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpirePeriod(v int32) *DescribeDBInstancesByExpireTimeRequest
	GetExpirePeriod() *int32
	SetExpired(v bool) *DescribeDBInstancesByExpireTimeRequest
	GetExpired() *bool
	SetOwnerAccount(v string) *DescribeDBInstancesByExpireTimeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBInstancesByExpireTimeRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeDBInstancesByExpireTimeRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDBInstancesByExpireTimeRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDBInstancesByExpireTimeRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDBInstancesByExpireTimeRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeDBInstancesByExpireTimeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBInstancesByExpireTimeRequest
	GetResourceOwnerId() *int64
	SetTags(v string) *DescribeDBInstancesByExpireTimeRequest
	GetTags() *string
	SetProxyId(v string) *DescribeDBInstancesByExpireTimeRequest
	GetProxyId() *string
}

type DescribeDBInstancesByExpireTimeRequest struct {
	// The number of remaining days for which the instances are available. Valid values: **0 to 180**.
	//
	// example:
	//
	// 180
	ExpirePeriod *int32 `json:"ExpirePeriod,omitempty" xml:"ExpirePeriod,omitempty"`
	// Specifies whether to query instances that have expired. Valid values:
	//
	// 	- **True**: queries instances that have expired.
	//
	// 	- **False**: does not query instances that have expired.
	//
	// example:
	//
	// True
	Expired      *bool   `json:"Expired,omitempty" xml:"Expired,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Valid values: any **non-zero*	- positive integer.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: **1 to 100**.
	//
	// Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. You can call the DescribeDBInstanceAttribute operation to obtain the resource group ID.
	//
	// example:
	//
	// rg-acfmy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tag that is added to the instance. Each tag is a key-value pair that consists of two parts: TagKey and TagValue. You can specify a maximum of five tags in the following format for each request: `{"key1":"value1","key2":"value2"...}`.
	//
	// example:
	//
	// {"key1":"value1"}
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// A deprecated parameter. You do not need to configure this parameter.
	//
	// example:
	//
	// None
	ProxyId *string `json:"proxyId,omitempty" xml:"proxyId,omitempty"`
}

func (s DescribeDBInstancesByExpireTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesByExpireTimeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesByExpireTimeRequest) GetExpirePeriod() *int32 {
	return s.ExpirePeriod
}

func (s *DescribeDBInstancesByExpireTimeRequest) GetExpired() *bool {
	return s.Expired
}

func (s *DescribeDBInstancesByExpireTimeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBInstancesByExpireTimeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBInstancesByExpireTimeRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBInstancesByExpireTimeRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDBInstancesByExpireTimeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstancesByExpireTimeRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBInstancesByExpireTimeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBInstancesByExpireTimeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInstancesByExpireTimeRequest) GetTags() *string {
	return s.Tags
}

func (s *DescribeDBInstancesByExpireTimeRequest) GetProxyId() *string {
	return s.ProxyId
}

func (s *DescribeDBInstancesByExpireTimeRequest) SetExpirePeriod(v int32) *DescribeDBInstancesByExpireTimeRequest {
	s.ExpirePeriod = &v
	return s
}

func (s *DescribeDBInstancesByExpireTimeRequest) SetExpired(v bool) *DescribeDBInstancesByExpireTimeRequest {
	s.Expired = &v
	return s
}

func (s *DescribeDBInstancesByExpireTimeRequest) SetOwnerAccount(v string) *DescribeDBInstancesByExpireTimeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBInstancesByExpireTimeRequest) SetOwnerId(v int64) *DescribeDBInstancesByExpireTimeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstancesByExpireTimeRequest) SetPageNumber(v int32) *DescribeDBInstancesByExpireTimeRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstancesByExpireTimeRequest) SetPageSize(v int32) *DescribeDBInstancesByExpireTimeRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBInstancesByExpireTimeRequest) SetRegionId(v string) *DescribeDBInstancesByExpireTimeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstancesByExpireTimeRequest) SetResourceGroupId(v string) *DescribeDBInstancesByExpireTimeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstancesByExpireTimeRequest) SetResourceOwnerAccount(v string) *DescribeDBInstancesByExpireTimeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInstancesByExpireTimeRequest) SetResourceOwnerId(v int64) *DescribeDBInstancesByExpireTimeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstancesByExpireTimeRequest) SetTags(v string) *DescribeDBInstancesByExpireTimeRequest {
	s.Tags = &v
	return s
}

func (s *DescribeDBInstancesByExpireTimeRequest) SetProxyId(v string) *DescribeDBInstancesByExpireTimeRequest {
	s.ProxyId = &v
	return s
}

func (s *DescribeDBInstancesByExpireTimeRequest) Validate() error {
	return dara.Validate(s)
}
