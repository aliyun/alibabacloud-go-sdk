// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCenBandwidthPackagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*DescribeCenBandwidthPackagesRequestFilter) *DescribeCenBandwidthPackagesRequest
	GetFilter() []*DescribeCenBandwidthPackagesRequestFilter
	SetIncludeReservationData(v bool) *DescribeCenBandwidthPackagesRequest
	GetIncludeReservationData() *bool
	SetIsOrKey(v bool) *DescribeCenBandwidthPackagesRequest
	GetIsOrKey() *bool
	SetOwnerAccount(v string) *DescribeCenBandwidthPackagesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeCenBandwidthPackagesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeCenBandwidthPackagesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCenBandwidthPackagesRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *DescribeCenBandwidthPackagesRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeCenBandwidthPackagesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeCenBandwidthPackagesRequest
	GetResourceOwnerId() *int64
	SetTag(v []*DescribeCenBandwidthPackagesRequestTag) *DescribeCenBandwidthPackagesRequest
	GetTag() []*DescribeCenBandwidthPackagesRequestTag
}

type DescribeCenBandwidthPackagesRequest struct {
	// The filter information.
	Filter []*DescribeCenBandwidthPackagesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// Specifies whether to include renewal data. Valid values:
	//
	// - **true**: Includes renewal data.
	//
	// - **false**: Does not include renewal data.
	//
	// example:
	//
	// true
	IncludeReservationData *bool `json:"IncludeReservationData,omitempty" xml:"IncludeReservationData,omitempty"`
	// The logical relationship between filter conditions. Valid values:
	//
	// - **false*	- (default): The filter conditions have an **AND*	- relationship. A bandwidth package must match all filter conditions to be returned.
	//
	// - **true**: The filter conditions have an **OR*	- relationship. A bandwidth package that matches any filter condition is returned.
	//
	// example:
	//
	// false
	IsOrKey      *bool   `json:"IsOrKey,omitempty" xml:"IsOrKey,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number of the list. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for a paged query. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfnwjeo4tv****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tag information list.
	//
	// You can specify up to 20 tags at a time.
	Tag []*DescribeCenBandwidthPackagesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeCenBandwidthPackagesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenBandwidthPackagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCenBandwidthPackagesRequest) GetFilter() []*DescribeCenBandwidthPackagesRequestFilter {
	return s.Filter
}

func (s *DescribeCenBandwidthPackagesRequest) GetIncludeReservationData() *bool {
	return s.IncludeReservationData
}

func (s *DescribeCenBandwidthPackagesRequest) GetIsOrKey() *bool {
	return s.IsOrKey
}

func (s *DescribeCenBandwidthPackagesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeCenBandwidthPackagesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCenBandwidthPackagesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCenBandwidthPackagesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCenBandwidthPackagesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeCenBandwidthPackagesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeCenBandwidthPackagesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCenBandwidthPackagesRequest) GetTag() []*DescribeCenBandwidthPackagesRequestTag {
	return s.Tag
}

func (s *DescribeCenBandwidthPackagesRequest) SetFilter(v []*DescribeCenBandwidthPackagesRequestFilter) *DescribeCenBandwidthPackagesRequest {
	s.Filter = v
	return s
}

func (s *DescribeCenBandwidthPackagesRequest) SetIncludeReservationData(v bool) *DescribeCenBandwidthPackagesRequest {
	s.IncludeReservationData = &v
	return s
}

func (s *DescribeCenBandwidthPackagesRequest) SetIsOrKey(v bool) *DescribeCenBandwidthPackagesRequest {
	s.IsOrKey = &v
	return s
}

func (s *DescribeCenBandwidthPackagesRequest) SetOwnerAccount(v string) *DescribeCenBandwidthPackagesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCenBandwidthPackagesRequest) SetOwnerId(v int64) *DescribeCenBandwidthPackagesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCenBandwidthPackagesRequest) SetPageNumber(v int32) *DescribeCenBandwidthPackagesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCenBandwidthPackagesRequest) SetPageSize(v int32) *DescribeCenBandwidthPackagesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCenBandwidthPackagesRequest) SetResourceGroupId(v string) *DescribeCenBandwidthPackagesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCenBandwidthPackagesRequest) SetResourceOwnerAccount(v string) *DescribeCenBandwidthPackagesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCenBandwidthPackagesRequest) SetResourceOwnerId(v int64) *DescribeCenBandwidthPackagesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCenBandwidthPackagesRequest) SetTag(v []*DescribeCenBandwidthPackagesRequestTag) *DescribeCenBandwidthPackagesRequest {
	s.Tag = v
	return s
}

func (s *DescribeCenBandwidthPackagesRequest) Validate() error {
	if s.Filter != nil {
		for _, item := range s.Filter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type DescribeCenBandwidthPackagesRequestFilter struct {
	// The filter condition.
	//
	// You can use filter conditions to filter the bandwidth package instances to query. The following filter conditions are supported:
	//
	// - **CenId**: The ID of the Cloud Enterprise Network (CEN) instance.
	//
	// - **Status**: The status of the bandwidth package instance. Valid values:
	//
	//     - **Idle**: Not associated.
	//
	//     - **InUse**: Associated.
	//
	// - **CenBandwidthPackageId**: The ID of the bandwidth package.
	//
	// - **Name**: The name of the bandwidth package.
	//
	// You can specify one or more filter conditions. The maximum value of **N*	- is **5**.
	//
	// example:
	//
	// CenId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The filter values based on the specified **Key**. You can specify multiple filter values for a single **Key**. The filter values have an **OR*	- relationship, which means that a bandwidth package matching any of the filter values is considered a match for the filter condition.
	//
	// example:
	//
	// Idle
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s DescribeCenBandwidthPackagesRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenBandwidthPackagesRequestFilter) GoString() string {
	return s.String()
}

func (s *DescribeCenBandwidthPackagesRequestFilter) GetKey() *string {
	return s.Key
}

func (s *DescribeCenBandwidthPackagesRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *DescribeCenBandwidthPackagesRequestFilter) SetKey(v string) *DescribeCenBandwidthPackagesRequestFilter {
	s.Key = &v
	return s
}

func (s *DescribeCenBandwidthPackagesRequestFilter) SetValue(v []*string) *DescribeCenBandwidthPackagesRequestFilter {
	s.Value = v
	return s
}

func (s *DescribeCenBandwidthPackagesRequestFilter) Validate() error {
	return dara.Validate(s)
}

type DescribeCenBandwidthPackagesRequestTag struct {
	// The tag key of the resource.
	//
	// Once specified, the tag key cannot be an empty string. The tag key can be up to 64 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// You can specify up to 20 tag keys at a time.
	//
	// example:
	//
	// TagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource.
	//
	// The tag value can be empty or a string of up to 128 characters. It cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// Each tag key corresponds to one tag value. You can specify up to 20 tag values at a time.
	//
	// example:
	//
	// TagValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeCenBandwidthPackagesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenBandwidthPackagesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeCenBandwidthPackagesRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeCenBandwidthPackagesRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeCenBandwidthPackagesRequestTag) SetKey(v string) *DescribeCenBandwidthPackagesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeCenBandwidthPackagesRequestTag) SetValue(v string) *DescribeCenBandwidthPackagesRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeCenBandwidthPackagesRequestTag) Validate() error {
	return dara.Validate(s)
}
