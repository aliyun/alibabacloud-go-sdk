// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVirtualBorderRoutersForPhysicalConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*DescribeVirtualBorderRoutersForPhysicalConnectionRequestFilter) *DescribeVirtualBorderRoutersForPhysicalConnectionRequest
	GetFilter() []*DescribeVirtualBorderRoutersForPhysicalConnectionRequestFilter
	SetOwnerId(v int64) *DescribeVirtualBorderRoutersForPhysicalConnectionRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeVirtualBorderRoutersForPhysicalConnectionRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVirtualBorderRoutersForPhysicalConnectionRequest
	GetPageSize() *int32
	SetPhysicalConnectionId(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionRequest
	GetPhysicalConnectionId() *string
	SetRegionId(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeVirtualBorderRoutersForPhysicalConnectionRequest
	GetResourceOwnerId() *int64
}

type DescribeVirtualBorderRoutersForPhysicalConnectionRequest struct {
	// The filter keys.
	Filter  []*DescribeVirtualBorderRoutersForPhysicalConnectionRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	OwnerId *int64                                                            `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the Express Connect circuit.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-119mfj****
	PhysicalConnectionId *string `json:"PhysicalConnectionId,omitempty" xml:"PhysicalConnectionId,omitempty"`
	// The region in which the Express Connect circuit is deployed. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to obtain the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeVirtualBorderRoutersForPhysicalConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVirtualBorderRoutersForPhysicalConnectionRequest) GoString() string {
	return s.String()
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequest) GetFilter() []*DescribeVirtualBorderRoutersForPhysicalConnectionRequestFilter {
	return s.Filter
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequest) GetPhysicalConnectionId() *string {
	return s.PhysicalConnectionId
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequest) SetFilter(v []*DescribeVirtualBorderRoutersForPhysicalConnectionRequestFilter) *DescribeVirtualBorderRoutersForPhysicalConnectionRequest {
	s.Filter = v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequest) SetOwnerId(v int64) *DescribeVirtualBorderRoutersForPhysicalConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequest) SetPageNumber(v int32) *DescribeVirtualBorderRoutersForPhysicalConnectionRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequest) SetPageSize(v int32) *DescribeVirtualBorderRoutersForPhysicalConnectionRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequest) SetPhysicalConnectionId(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionRequest {
	s.PhysicalConnectionId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequest) SetRegionId(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequest) SetResourceOwnerAccount(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequest) SetResourceOwnerId(v int64) *DescribeVirtualBorderRoutersForPhysicalConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequest) Validate() error {
	if s.Filter != nil {
		for _, item := range s.Filter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVirtualBorderRoutersForPhysicalConnectionRequestFilter struct {
	// The filter conditions. You can specify at most five filter conditions. The following filter conditions are supported:
	//
	// 	- **PhysicalConnectionId**: filter VBRs by the Express Connect circuit ID.
	//
	// 	- **VbrId**: filter VBRs by ID.
	//
	// 	- **Status**: filter VBRs by status.
	//
	// 	- **Name**: filter VBRs by name.
	//
	// 	- **AccessPointId**: filter VBRs by access point ID.
	//
	// 	- **eccId**: filter VBRs by ID of Express Cloud Connect (ECC) instance.
	//
	// 	- **type**: filter VBRs by type.
	//
	// example:
	//
	// Status
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The filter value for the key. You can specify multiple filter values for one key. The logical operator among filter values is OR. If one filter value is matched, the filter condition is matched.
	//
	// example:
	//
	// Active
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s DescribeVirtualBorderRoutersForPhysicalConnectionRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeVirtualBorderRoutersForPhysicalConnectionRequestFilter) GoString() string {
	return s.String()
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequestFilter) GetKey() *string {
	return s.Key
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequestFilter) SetKey(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionRequestFilter {
	s.Key = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequestFilter) SetValue(v []*string) *DescribeVirtualBorderRoutersForPhysicalConnectionRequestFilter {
	s.Value = v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequestFilter) Validate() error {
	return dara.Validate(s)
}
