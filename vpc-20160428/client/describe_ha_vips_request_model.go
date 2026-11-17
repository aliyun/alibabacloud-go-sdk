// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHaVipsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*DescribeHaVipsRequestFilter) *DescribeHaVipsRequest
	GetFilter() []*DescribeHaVipsRequestFilter
	SetOwnerAccount(v string) *DescribeHaVipsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeHaVipsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeHaVipsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHaVipsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeHaVipsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeHaVipsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeHaVipsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeHaVipsRequest
	GetResourceOwnerId() *int64
	SetTags(v []*DescribeHaVipsRequestTags) *DescribeHaVipsRequest
	GetTags() []*DescribeHaVipsRequestTags
}

type DescribeHaVipsRequest struct {
	// The details of the filter conditions.
	Filter       []*DescribeHaVipsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	OwnerAccount *string                        `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64                         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page in a paging query. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the HaVip.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to obtain the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the HaVip belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4ph****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The list of tags.
	Tags []*DescribeHaVipsRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeHaVipsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHaVipsRequest) GoString() string {
	return s.String()
}

func (s *DescribeHaVipsRequest) GetFilter() []*DescribeHaVipsRequestFilter {
	return s.Filter
}

func (s *DescribeHaVipsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeHaVipsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeHaVipsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHaVipsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHaVipsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHaVipsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeHaVipsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeHaVipsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeHaVipsRequest) GetTags() []*DescribeHaVipsRequestTags {
	return s.Tags
}

func (s *DescribeHaVipsRequest) SetFilter(v []*DescribeHaVipsRequestFilter) *DescribeHaVipsRequest {
	s.Filter = v
	return s
}

func (s *DescribeHaVipsRequest) SetOwnerAccount(v string) *DescribeHaVipsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeHaVipsRequest) SetOwnerId(v int64) *DescribeHaVipsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeHaVipsRequest) SetPageNumber(v int32) *DescribeHaVipsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHaVipsRequest) SetPageSize(v int32) *DescribeHaVipsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHaVipsRequest) SetRegionId(v string) *DescribeHaVipsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHaVipsRequest) SetResourceGroupId(v string) *DescribeHaVipsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeHaVipsRequest) SetResourceOwnerAccount(v string) *DescribeHaVipsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeHaVipsRequest) SetResourceOwnerId(v int64) *DescribeHaVipsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeHaVipsRequest) SetTags(v []*DescribeHaVipsRequestTags) *DescribeHaVipsRequest {
	s.Tags = v
	return s
}

func (s *DescribeHaVipsRequest) Validate() error {
	if s.Filter != nil {
		for _, item := range s.Filter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeHaVipsRequestFilter struct {
	// The filter condition. You can specify up to 5 filter conditions. Valid values of **N**: **1*	- to **5**.
	//
	// Supported filter conditions:
	//
	// - **VpcId**: The ID of the virtual private cloud (VPC).
	//
	// - **VSwitchId**: The vSwitch ID.
	//
	// - **Status**: The status of the HaVip.
	//
	// - **HaVipId**: The HaVip ID.
	//
	// - **HaVipAddress**: The IP address of the HaVip.
	//
	// You can specify multiple values for each filter condition. The values are evaluated by using the OR operator. A resource is considered to match the filter condition of the parameter if it matches at least one of the values.
	//
	// Multiple filter conditions are evaluated by using the AND operator. A resource is returned only if it matches the filter conditions of all parameters.
	//
	// example:
	//
	// HaVipId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the specified filter condition. Valid values of **N**: **1*	- to **5**.
	//
	// example:
	//
	// havip-bp19o63nequs01i8d****
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s DescribeHaVipsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeHaVipsRequestFilter) GoString() string {
	return s.String()
}

func (s *DescribeHaVipsRequestFilter) GetKey() *string {
	return s.Key
}

func (s *DescribeHaVipsRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *DescribeHaVipsRequestFilter) SetKey(v string) *DescribeHaVipsRequestFilter {
	s.Key = &v
	return s
}

func (s *DescribeHaVipsRequestFilter) SetValue(v []*string) *DescribeHaVipsRequestFilter {
	s.Value = v
	return s
}

func (s *DescribeHaVipsRequestFilter) Validate() error {
	return dara.Validate(s)
}

type DescribeHaVipsRequestTags struct {
	// The tag key of the resource. You can specify up to 20 tag keys. If you specify this parameter, you cannot enter an empty string.
	//
	// A tag key can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify up to 20 tag values. If you specify this parameter, you can enter an empty string.
	//
	// A tag value can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHaVipsRequestTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeHaVipsRequestTags) GoString() string {
	return s.String()
}

func (s *DescribeHaVipsRequestTags) GetKey() *string {
	return s.Key
}

func (s *DescribeHaVipsRequestTags) GetValue() *string {
	return s.Value
}

func (s *DescribeHaVipsRequestTags) SetKey(v string) *DescribeHaVipsRequestTags {
	s.Key = &v
	return s
}

func (s *DescribeHaVipsRequestTags) SetValue(v string) *DescribeHaVipsRequestTags {
	s.Value = &v
	return s
}

func (s *DescribeHaVipsRequestTags) Validate() error {
	return dara.Validate(s)
}
