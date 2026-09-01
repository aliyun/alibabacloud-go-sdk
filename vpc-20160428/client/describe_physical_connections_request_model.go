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
	// Generate a parameter value from your client to ensure uniqueness across different requests. ClientToken supports only ASCII characters.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The list of filter conditions.
	Filter []*DescribePhysicalConnectionsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// Specifies whether to return order data that has not taken effect. Valid values:
	//
	// 	- **true**: Returns order data that has not taken effect.
	//
	// 	- **false*	- (default): Does not return order data that has not taken effect.
	//
	// example:
	//
	// false
	IncludeReservationData *bool   `json:"IncludeReservationData,omitempty" xml:"IncludeReservationData,omitempty"`
	OwnerAccount           *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number of the list. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page in a paged query. Default value: **10**. Valid values: **1*	- to **50**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the Express Connect circuit.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
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
	// rg-aek2yvwibxr****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The list of tags.
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

type DescribePhysicalConnectionsRequestFilter struct {
	// The filter condition. Valid values:
	//
	// - **PhysicalConnectionId**: the ID of the Express Connect circuit.
	//
	// - **AccessPointId**: the ID of the access point.
	//
	// - **Type**: the type of the Express Connect circuit. This filter condition supports only the filter value **VPC**.
	//
	// - **LineOperator**: the carrier of the Express Connect circuit. This filter condition supports the following filter values:
	//
	//     - **CT**: China Telecom.
	//
	//     - **CU**: China Unicom.
	//
	//     - **CM**: China shift.
	//
	//     - **CO**: Other carriers in the Chinese mainland.
	//
	//     - **Equinix**: Equinix.
	//
	//     - **Other**: Other carriers outside the Chinese mainland.
	//
	// - **Spec**: the specification of the Express Connect circuit. This filter condition supports the following filter values:
	//
	//     - **1G and below**.
	//
	//     - **10G**.
	//
	//     - **40G**.
	//
	//     - **100G**.
	//
	// >  The **40G*	- and **100G*	- specifications are not available by default. Only users who have submitted an application to their account manager and received approval can use these filter values.
	//
	// - **Status**: the status of the Express Connect circuit. This filter condition supports the following filter values:
	//
	//     - **Initial**: pending application.
	//
	//     - **Approved**: approved.
	//
	//     - **Allocating**: allocating resources.
	//
	//     - **Allocated**: under construction.
	//
	//     - **Confirmed**: pending user confirmation.
	//
	//     - **Enabled**: enabled.
	//
	//     - **Rejected**: application rejected.
	//
	//     - **Canceled**: canceled.
	//
	//     - **Allocation Failed**: resource allocation failed.
	//
	//     - **Terminating**: stopping.
	//
	//     - **Terminated**: stopped.
	//
	// - **Name**: the name of the Express Connect circuit.
	//
	// - **ProductType**: the type of the Express Connect circuit. Valid values:
	//
	//     - **VirtualPhysicalConnection**: shared Express Connect circuit.
	//
	//     - **PhysicalConnection**: dedicated Express Connect circuit.
	//
	//
	//
	//
	// You can specify up to 5 filter conditions at a time. The filter conditions have an **AND*	- relationship. Results are returned only when all filter conditions are met.
	//
	// example:
	//
	// Name
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The list of filter values.
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
	// The tag key of the resource. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and can contain digits, periods (.), underscores (_), and hyphens (-). It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and can contain digits, periods (.), underscores (_), and hyphens (-). It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
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
