// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhysicalConnectionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribePhysicalConnectionsRequest
	GetClientToken() *string
	SetFilter(v []*DescribePhysicalConnectionsRequestFilter) *DescribePhysicalConnectionsRequest
	GetFilter() []*DescribePhysicalConnectionsRequestFilter
	SetIncludeReservationData(v bool) *DescribePhysicalConnectionsRequest
	GetIncludeReservationData() *bool
	SetOwnerAccount(v string) *DescribePhysicalConnectionsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribePhysicalConnectionsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribePhysicalConnectionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePhysicalConnectionsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribePhysicalConnectionsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribePhysicalConnectionsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribePhysicalConnectionsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribePhysicalConnectionsRequest
	GetResourceOwnerId() *int64
	SetTags(v []*DescribePhysicalConnectionsRequestTags) *DescribePhysicalConnectionsRequest
	GetTags() []*DescribePhysicalConnectionsRequestTags
}

type DescribePhysicalConnectionsRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The filter keys.
	Filter []*DescribePhysicalConnectionsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// Specifies whether to return the data about pending orders. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	IncludeReservationData *bool   `json:"IncludeReservationData,omitempty" xml:"IncludeReservationData,omitempty"`
	OwnerAccount           *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **10**. Valid values: **1*	- to **50**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the Express Connect circuit.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Express Connect circuit belongs.
	//
	// example:
	//
	// rg-aek2yvwibxrmrkq
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tag list.
	Tags []*DescribePhysicalConnectionsRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribePhysicalConnectionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePhysicalConnectionsRequest) GoString() string {
	return s.String()
}

func (s *DescribePhysicalConnectionsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribePhysicalConnectionsRequest) GetFilter() []*DescribePhysicalConnectionsRequestFilter {
	return s.Filter
}

func (s *DescribePhysicalConnectionsRequest) GetIncludeReservationData() *bool {
	return s.IncludeReservationData
}

func (s *DescribePhysicalConnectionsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribePhysicalConnectionsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePhysicalConnectionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePhysicalConnectionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePhysicalConnectionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePhysicalConnectionsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribePhysicalConnectionsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribePhysicalConnectionsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribePhysicalConnectionsRequest) GetTags() []*DescribePhysicalConnectionsRequestTags {
	return s.Tags
}

func (s *DescribePhysicalConnectionsRequest) SetClientToken(v string) *DescribePhysicalConnectionsRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribePhysicalConnectionsRequest) SetFilter(v []*DescribePhysicalConnectionsRequestFilter) *DescribePhysicalConnectionsRequest {
	s.Filter = v
	return s
}

func (s *DescribePhysicalConnectionsRequest) SetIncludeReservationData(v bool) *DescribePhysicalConnectionsRequest {
	s.IncludeReservationData = &v
	return s
}

func (s *DescribePhysicalConnectionsRequest) SetOwnerAccount(v string) *DescribePhysicalConnectionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribePhysicalConnectionsRequest) SetOwnerId(v int64) *DescribePhysicalConnectionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePhysicalConnectionsRequest) SetPageNumber(v int32) *DescribePhysicalConnectionsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePhysicalConnectionsRequest) SetPageSize(v int32) *DescribePhysicalConnectionsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePhysicalConnectionsRequest) SetRegionId(v string) *DescribePhysicalConnectionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePhysicalConnectionsRequest) SetResourceGroupId(v string) *DescribePhysicalConnectionsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribePhysicalConnectionsRequest) SetResourceOwnerAccount(v string) *DescribePhysicalConnectionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePhysicalConnectionsRequest) SetResourceOwnerId(v int64) *DescribePhysicalConnectionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePhysicalConnectionsRequest) SetTags(v []*DescribePhysicalConnectionsRequestTags) *DescribePhysicalConnectionsRequest {
	s.Tags = v
	return s
}

func (s *DescribePhysicalConnectionsRequest) Validate() error {
	return dara.Validate(s)
}

type DescribePhysicalConnectionsRequestFilter struct {
	// The key of the filter. Valid values:
	//
	// 	- **PhysicalConnectionId**: the ID of the Express Connect circuit.
	//
	// 	- **AccessPointId**: the ID of the access point.
	//
	// 	- **Type**: the type of resource to which the Express Connect circuit is connected. You can set Type only to **VPC**.
	//
	// 	- **LineOperator**: the connectivity provider of the Express Connect circuit. Valid values:
	//
	//     	- **CT**: China Telecom.
	//
	//     	- **CU**: China Unicom.
	//
	//     	- **CM**: China Mobile.
	//
	//     	- **CO**: other connectivity providers in the Chinese mainland.
	//
	//     	- **Equinix**: Equinix.
	//
	//     	- **Other**: other connectivity providers outside the Chinese mainland.
	//
	// 	- **Spec**: the specification of the Express Connect circuit. Valid values:
	//
	//     	- **1G and below**
	//
	//     	- **10G**
	//
	//     	- **40G**
	//
	//     	- **100G**
	//
	// >  By default, you cannot set the value to **40G*	- or **100G**. To use these values, you must first contact your account manager.
	//
	// 	- **Status**: the status of the Express Connect circuit. Valid values:
	//
	//     	- **Initial**: The application is under review.
	//
	//     	- **Approved**: The application is approved.
	//
	//     	- **Allocating**: The system is allocating resources.
	//
	//     	- **Allocated**: The Express Connect circuit is under construction.
	//
	//     	- **Confirmed**: The Express Connect circuit is pending for user confirmation.
	//
	//     	- **Enabled**: The Express Connect circuit is enabled.
	//
	//     	- **Rejected**: The application is rejected.
	//
	//     	- **Canceled**: The application is canceled.
	//
	//     	- **Allocation Failed**: The system failed to allocate resources.
	//
	//     	- **Terminating**: The Express Connect circuit is being disabled.
	//
	//     	- **Terminated**: The Express Connect circuit is disabled.
	//
	// 	- **Name**: the name of the Express Connect circuit.
	//
	// 	- **ProductType**: the type of the Express Connect circuit. Valid values:
	//
	//     	- **VirtualPhysicalConnection**: shared Express Connect circuit
	//
	//     	- **PhysicalConnection**: dedicated Express Connect circuit.
	//
	// You can specify at most five filter conditions in each request. The logical relation among the filter conditions is **AND**. Therefore, an Express Connect circuit is returned only when all specified filter conditions are matched.
	//
	// example:
	//
	// Name
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The filter values.
	//
	// example:
	//
	// 1
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s DescribePhysicalConnectionsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribePhysicalConnectionsRequestFilter) GoString() string {
	return s.String()
}

func (s *DescribePhysicalConnectionsRequestFilter) GetKey() *string {
	return s.Key
}

func (s *DescribePhysicalConnectionsRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *DescribePhysicalConnectionsRequestFilter) SetKey(v string) *DescribePhysicalConnectionsRequestFilter {
	s.Key = &v
	return s
}

func (s *DescribePhysicalConnectionsRequestFilter) SetValue(v []*string) *DescribePhysicalConnectionsRequestFilter {
	s.Value = v
	return s
}

func (s *DescribePhysicalConnectionsRequestFilter) Validate() error {
	return dara.Validate(s)
}

type DescribePhysicalConnectionsRequestTags struct {
	// The key of tag N to add to the resource. You can specify at most 20 tag keys. The tag key cannot be an empty string.
	//
	// It can be up to 64 characters in length and can contain digits, periods (.), underscores (_), and hyphens (-). It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the resource. You can specify at most 20 tag values. The tag value can be an empty string.
	//
	// It can be up to 128 characters in length and can contain digits, periods (.), underscores (_), and hyphens (-). It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribePhysicalConnectionsRequestTags) String() string {
	return dara.Prettify(s)
}

func (s DescribePhysicalConnectionsRequestTags) GoString() string {
	return s.String()
}

func (s *DescribePhysicalConnectionsRequestTags) GetKey() *string {
	return s.Key
}

func (s *DescribePhysicalConnectionsRequestTags) GetValue() *string {
	return s.Value
}

func (s *DescribePhysicalConnectionsRequestTags) SetKey(v string) *DescribePhysicalConnectionsRequestTags {
	s.Key = &v
	return s
}

func (s *DescribePhysicalConnectionsRequestTags) SetValue(v string) *DescribePhysicalConnectionsRequestTags {
	s.Value = &v
	return s
}

func (s *DescribePhysicalConnectionsRequestTags) Validate() error {
	return dara.Validate(s)
}
