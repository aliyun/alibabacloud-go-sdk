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
	SetOwnerId(v int64) *DescribeVirtualBorderRoutersRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeVirtualBorderRoutersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVirtualBorderRoutersRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeVirtualBorderRoutersRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeVirtualBorderRoutersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeVirtualBorderRoutersRequest
	GetResourceOwnerId() *int64
}

type DescribeVirtualBorderRoutersRequest struct {
	// The filter conditions.
	Filter  []*DescribeVirtualBorderRoutersRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	OwnerId *int64                                       `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages are numbered starting from 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 50. Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the Virtual Border Router (VBR) is located. You can call the `DescribeRegions` operation to obtain the most recent list of regions.
	//
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
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

func (s *DescribeVirtualBorderRoutersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeVirtualBorderRoutersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeVirtualBorderRoutersRequest) SetFilter(v []*DescribeVirtualBorderRoutersRequestFilter) *DescribeVirtualBorderRoutersRequest {
	s.Filter = v
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

func (s *DescribeVirtualBorderRoutersRequest) SetResourceOwnerAccount(v string) *DescribeVirtualBorderRoutersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVirtualBorderRoutersRequest) SetResourceOwnerId(v int64) *DescribeVirtualBorderRoutersRequest {
	s.ResourceOwnerId = &v
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
	return nil
}

type DescribeVirtualBorderRoutersRequestFilter struct {
	// The filter key. Set the value to `VbrId`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The filter value. The value must be an array of VBR IDs.
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
