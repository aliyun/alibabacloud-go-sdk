// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVirtualBorderRoutersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*DescribeVirtualBorderRoutersRequestFilter) *DescribeVirtualBorderRoutersRequest
	GetFilter() []*DescribeVirtualBorderRoutersRequestFilter
	SetIncludeCrossAccountVbr(v bool) *DescribeVirtualBorderRoutersRequest
	GetIncludeCrossAccountVbr() *bool
	SetOwnerId(v int64) *DescribeVirtualBorderRoutersRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeVirtualBorderRoutersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVirtualBorderRoutersRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeVirtualBorderRoutersRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeVirtualBorderRoutersRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeVirtualBorderRoutersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeVirtualBorderRoutersRequest
	GetResourceOwnerId() *int64
	SetTags(v []*DescribeVirtualBorderRoutersRequestTags) *DescribeVirtualBorderRoutersRequest
	GetTags() []*DescribeVirtualBorderRoutersRequestTags
}

type DescribeVirtualBorderRoutersRequest struct {
	// The filter information.
	Filter []*DescribeVirtualBorderRoutersRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// Specifies whether to include cross-account Virtual Border Routers.
	//
	// - **true**: Included.
	//
	// - **false*	- (default): Not included.
	//
	// example:
	//
	// false
	IncludeCrossAccountVbr *bool  `json:"IncludeCrossAccountVbr,omitempty" xml:"IncludeCrossAccountVbr,omitempty"`
	OwnerId                *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number of the list. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page in a paged query. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the VBR. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// For more information about resource groups, see [What is a resource group?](https://help.aliyun.com/document_detail/94475.html).
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags of the resource.
	Tags []*DescribeVirtualBorderRoutersRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeVirtualBorderRoutersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVirtualBorderRoutersRequest) GoString() string {
	return s.String()
}

func (s *DescribeVirtualBorderRoutersRequest) GetFilter() []*DescribeVirtualBorderRoutersRequestFilter {
	return s.Filter
}

func (s *DescribeVirtualBorderRoutersRequest) GetIncludeCrossAccountVbr() *bool {
	return s.IncludeCrossAccountVbr
}

func (s *DescribeVirtualBorderRoutersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVirtualBorderRoutersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVirtualBorderRoutersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVirtualBorderRoutersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVirtualBorderRoutersRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeVirtualBorderRoutersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeVirtualBorderRoutersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeVirtualBorderRoutersRequest) GetTags() []*DescribeVirtualBorderRoutersRequestTags {
	return s.Tags
}

func (s *DescribeVirtualBorderRoutersRequest) SetFilter(v []*DescribeVirtualBorderRoutersRequestFilter) *DescribeVirtualBorderRoutersRequest {
	s.Filter = v
	return s
}

func (s *DescribeVirtualBorderRoutersRequest) SetIncludeCrossAccountVbr(v bool) *DescribeVirtualBorderRoutersRequest {
	s.IncludeCrossAccountVbr = &v
	return s
}

func (s *DescribeVirtualBorderRoutersRequest) SetOwnerId(v int64) *DescribeVirtualBorderRoutersRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersRequest) SetPageNumber(v int32) *DescribeVirtualBorderRoutersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVirtualBorderRoutersRequest) SetPageSize(v int32) *DescribeVirtualBorderRoutersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVirtualBorderRoutersRequest) SetRegionId(v string) *DescribeVirtualBorderRoutersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersRequest) SetResourceGroupId(v string) *DescribeVirtualBorderRoutersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersRequest) SetResourceOwnerAccount(v string) *DescribeVirtualBorderRoutersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVirtualBorderRoutersRequest) SetResourceOwnerId(v int64) *DescribeVirtualBorderRoutersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersRequest) SetTags(v []*DescribeVirtualBorderRoutersRequestTags) *DescribeVirtualBorderRoutersRequest {
	s.Tags = v
	return s
}

func (s *DescribeVirtualBorderRoutersRequest) Validate() error {
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

type DescribeVirtualBorderRoutersRequestFilter struct {
	// The filter condition. You can specify up to 5 filter conditions. The following filter conditions are supported:
	//
	// 	- **PhysicalConnectionId**: instance ID of the Express Connect circuit instance.
	//
	// 	- **VbrId**: instance ID of the Virtual Border Router instance.
	//
	// 	- **Status**: the status of the Virtual Border Router.
	//
	// 	- **Name**: the name of the Virtual Border Router.
	//
	// 	- **AccessPointId**: instance ID of the access point.
	//
	// 	- **eccId**: instance ID of the Express Cloud Connect instance.
	//
	// 	- **type**: the type of the Express Connect circuit.
	//
	// example:
	//
	// Status
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The filter value based on the specified Key. You can specify multiple filter values for a Key. The relationship between filter values is OR, which means that a match is found if any of the filter values is met.
	//
	// example:
	//
	// Active
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s DescribeVirtualBorderRoutersRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeVirtualBorderRoutersRequestFilter) GoString() string {
	return s.String()
}

func (s *DescribeVirtualBorderRoutersRequestFilter) GetKey() *string {
	return s.Key
}

func (s *DescribeVirtualBorderRoutersRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *DescribeVirtualBorderRoutersRequestFilter) SetKey(v string) *DescribeVirtualBorderRoutersRequestFilter {
	s.Key = &v
	return s
}

func (s *DescribeVirtualBorderRoutersRequestFilter) SetValue(v []*string) *DescribeVirtualBorderRoutersRequestFilter {
	s.Value = v
	return s
}

func (s *DescribeVirtualBorderRoutersRequestFilter) Validate() error {
	return dara.Validate(s)
}

type DescribeVirtualBorderRoutersRequestTags struct {
	// The tag key of the resource. You must specify at least 1 tag key and can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// A tag key can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVirtualBorderRoutersRequestTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeVirtualBorderRoutersRequestTags) GoString() string {
	return s.String()
}

func (s *DescribeVirtualBorderRoutersRequestTags) GetKey() *string {
	return s.Key
}

func (s *DescribeVirtualBorderRoutersRequestTags) GetValue() *string {
	return s.Value
}

func (s *DescribeVirtualBorderRoutersRequestTags) SetKey(v string) *DescribeVirtualBorderRoutersRequestTags {
	s.Key = &v
	return s
}

func (s *DescribeVirtualBorderRoutersRequestTags) SetValue(v string) *DescribeVirtualBorderRoutersRequestTags {
	s.Value = &v
	return s
}

func (s *DescribeVirtualBorderRoutersRequestTags) Validate() error {
	return dara.Validate(s)
}
