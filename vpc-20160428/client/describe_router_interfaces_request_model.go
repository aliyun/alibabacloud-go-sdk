// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRouterInterfacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*DescribeRouterInterfacesRequestFilter) *DescribeRouterInterfacesRequest
	GetFilter() []*DescribeRouterInterfacesRequestFilter
	SetIncludeReservationData(v bool) *DescribeRouterInterfacesRequest
	GetIncludeReservationData() *bool
	SetOwnerId(v int64) *DescribeRouterInterfacesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeRouterInterfacesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRouterInterfacesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeRouterInterfacesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeRouterInterfacesRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeRouterInterfacesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeRouterInterfacesRequest
	GetResourceOwnerId() *int64
	SetTags(v []*DescribeRouterInterfacesRequestTags) *DescribeRouterInterfacesRequest
	GetTags() []*DescribeRouterInterfacesRequestTags
}

type DescribeRouterInterfacesRequest struct {
	// The filter conditions.
	Filter []*DescribeRouterInterfacesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// Specifies whether to include renewal data. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false*	- (default): No.
	//
	// example:
	//
	// false
	IncludeReservationData *bool  `json:"IncludeReservationData,omitempty" xml:"IncludeReservationData,omitempty"`
	OwnerId                *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// The region ID of the router interface.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// For more information about resource groups, see [What is a resource group?](https://help.aliyun.com/document_detail/2381067.html).
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags of the resource.
	Tags []*DescribeRouterInterfacesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeRouterInterfacesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouterInterfacesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRouterInterfacesRequest) GetFilter() []*DescribeRouterInterfacesRequestFilter {
	return s.Filter
}

func (s *DescribeRouterInterfacesRequest) GetIncludeReservationData() *bool {
	return s.IncludeReservationData
}

func (s *DescribeRouterInterfacesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRouterInterfacesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRouterInterfacesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRouterInterfacesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRouterInterfacesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeRouterInterfacesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeRouterInterfacesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeRouterInterfacesRequest) GetTags() []*DescribeRouterInterfacesRequestTags {
	return s.Tags
}

func (s *DescribeRouterInterfacesRequest) SetFilter(v []*DescribeRouterInterfacesRequestFilter) *DescribeRouterInterfacesRequest {
	s.Filter = v
	return s
}

func (s *DescribeRouterInterfacesRequest) SetIncludeReservationData(v bool) *DescribeRouterInterfacesRequest {
	s.IncludeReservationData = &v
	return s
}

func (s *DescribeRouterInterfacesRequest) SetOwnerId(v int64) *DescribeRouterInterfacesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRouterInterfacesRequest) SetPageNumber(v int32) *DescribeRouterInterfacesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRouterInterfacesRequest) SetPageSize(v int32) *DescribeRouterInterfacesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRouterInterfacesRequest) SetRegionId(v string) *DescribeRouterInterfacesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRouterInterfacesRequest) SetResourceGroupId(v string) *DescribeRouterInterfacesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeRouterInterfacesRequest) SetResourceOwnerAccount(v string) *DescribeRouterInterfacesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRouterInterfacesRequest) SetResourceOwnerId(v int64) *DescribeRouterInterfacesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRouterInterfacesRequest) SetTags(v []*DescribeRouterInterfacesRequestTags) *DescribeRouterInterfacesRequest {
	s.Tags = v
	return s
}

func (s *DescribeRouterInterfacesRequest) Validate() error {
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

type DescribeRouterInterfacesRequestFilter struct {
	// The filter key. You can specify up to five filter keys. The following filter keys are supported:
	//
	// - **RouterInterfaceId**: The ID of the router interface.
	//
	// - **RouterId**: The ID of the router.
	//
	// - **RouterType**: The type of the router. Valid values: **VRouter*	- and **VBR**.
	//
	// - **RouterInterfaceOwnerId**: The ID of the Alibaba Cloud account to which the router interface belongs.
	//
	// - **OppositeInterfaceId**: The ID of the peer router interface.
	//
	// - **OppositeRouterType**: The type of the peer router. Valid values: **VRouter*	- and **VBR**.
	//
	// - **OppositeRouterId**: The ID of the peer router.
	//
	// - **OppositeInterfaceOwnerId**: The ID of the Alibaba Cloud account to which the peer router interface belongs.
	//
	// - **Status**: The status of the router interface.
	//
	// - **Name**: The name of the router interface.
	//
	// > Multiple values for a filter key are evaluated by using the OR operator. A resource matches the filter condition if it matches one of the values. Multiple filter keys are evaluated by using the AND operator. A resource is returned in the query result only if it matches all the specified filter conditions.
	//
	// example:
	//
	// Filter.1.Status
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The filter value specified for the filter key. You can specify multiple filter values for a filter key. Multiple filter values are evaluated by using the OR operator. A resource matches the filter condition if it matches one of the filter values.
	//
	// example:
	//
	// Filter.1.Active 1
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s DescribeRouterInterfacesRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouterInterfacesRequestFilter) GoString() string {
	return s.String()
}

func (s *DescribeRouterInterfacesRequestFilter) GetKey() *string {
	return s.Key
}

func (s *DescribeRouterInterfacesRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *DescribeRouterInterfacesRequestFilter) SetKey(v string) *DescribeRouterInterfacesRequestFilter {
	s.Key = &v
	return s
}

func (s *DescribeRouterInterfacesRequestFilter) SetValue(v []*string) *DescribeRouterInterfacesRequestFilter {
	s.Value = v
	return s
}

func (s *DescribeRouterInterfacesRequestFilter) Validate() error {
	return dara.Validate(s)
}

type DescribeRouterInterfacesRequestTags struct {
	// The tag key of the resource. You must specify at least one tag key and can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// A tag key can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// A tag value can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeRouterInterfacesRequestTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouterInterfacesRequestTags) GoString() string {
	return s.String()
}

func (s *DescribeRouterInterfacesRequestTags) GetKey() *string {
	return s.Key
}

func (s *DescribeRouterInterfacesRequestTags) GetValue() *string {
	return s.Value
}

func (s *DescribeRouterInterfacesRequestTags) SetKey(v string) *DescribeRouterInterfacesRequestTags {
	s.Key = &v
	return s
}

func (s *DescribeRouterInterfacesRequestTags) SetValue(v string) *DescribeRouterInterfacesRequestTags {
	s.Value = &v
	return s
}

func (s *DescribeRouterInterfacesRequestTags) Validate() error {
	return dara.Validate(s)
}
