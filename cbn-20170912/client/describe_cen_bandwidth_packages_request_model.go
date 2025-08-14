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
	// The filter configurations.
	Filter []*DescribeCenBandwidthPackagesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// Specifies whether to include renewal data. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IncludeReservationData *bool `json:"IncludeReservationData,omitempty" xml:"IncludeReservationData,omitempty"`
	// The logical operator between the filter conditions. Valid values:
	//
	// 	- **false*	- (default): **AND*	- Bandwidth plans that meet all filter conditions are returned.
	//
	// 	- **true**: **OR*	- Bandwidth plans that meet one of the filter conditions are returned.
	//
	// example:
	//
	// false
	IsOrKey      *bool   `json:"IsOrKey,omitempty" xml:"IsOrKey,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfnwjeo4tv****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The information about the tags.
	//
	// You can specify at most 20 tags in each call.
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
	return dara.Validate(s)
}

type DescribeCenBandwidthPackagesRequestFilter struct {
	// The filter conditions. You can use filter conditions to filter the bandwidth plans that you want to query. The following filter conditions are supported:
	//
	// 	- **CenId**: CEN instance ID
	//
	// 	- **Status**: bandwidth plan status. Valid values:
	//
	//     	- **Idle**: not associated with a CEN instance.
	//
	//     	- **InUse**: associated with a CEN instance.
	//
	// 	- **CenBandwidthPackageId**: bandwidth plan ID
	//
	// 	- **Name**: bandwidth plan name You can specify one or more filter conditions. The maximum value of **N*	- is **5**.
	//
	// example:
	//
	// CenId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Specify a filter value based on the **Key*	- parameter. You can specify multiple filter values for each **Key**. The logical operator between filter values is **OR**. If one filter value is matched, the filter condition is matched.
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
	// The tag keys.
	//
	// The tag keys cannot be an empty string. The tag keys can be up to 64 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// You can specify at most 20 tag keys.
	//
	// example:
	//
	// TagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag values.
	//
	// The tag values can be 0 to 128 characters in length, and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// The tag value of each tag key must be unique. You can specify at most 20 tag values in each call.
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
