// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessPointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*DescribeAccessPointsRequestFilter) *DescribeAccessPointsRequest
	GetFilter() []*DescribeAccessPointsRequestFilter
	SetOwnerId(v int64) *DescribeAccessPointsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeAccessPointsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAccessPointsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeAccessPointsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeAccessPointsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAccessPointsRequest
	GetResourceOwnerId() *int64
	SetType(v string) *DescribeAccessPointsRequest
	GetType() *string
}

type DescribeAccessPointsRequest struct {
	// The filters to apply to the query results.
	Filter  []*DescribeAccessPointsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	OwnerId *int64                               `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the access points are located. Call the `DescribeRegions` operation to query the latest list of regions.
	//
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the access point. Set the value to `ecs`.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAccessPointsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessPointsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccessPointsRequest) GetFilter() []*DescribeAccessPointsRequestFilter {
	return s.Filter
}

func (s *DescribeAccessPointsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAccessPointsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAccessPointsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAccessPointsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAccessPointsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAccessPointsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAccessPointsRequest) GetType() *string {
	return s.Type
}

func (s *DescribeAccessPointsRequest) SetFilter(v []*DescribeAccessPointsRequestFilter) *DescribeAccessPointsRequest {
	s.Filter = v
	return s
}

func (s *DescribeAccessPointsRequest) SetOwnerId(v int64) *DescribeAccessPointsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAccessPointsRequest) SetPageNumber(v int32) *DescribeAccessPointsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccessPointsRequest) SetPageSize(v int32) *DescribeAccessPointsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAccessPointsRequest) SetRegionId(v string) *DescribeAccessPointsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAccessPointsRequest) SetResourceOwnerAccount(v string) *DescribeAccessPointsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAccessPointsRequest) SetResourceOwnerId(v int64) *DescribeAccessPointsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAccessPointsRequest) SetType(v string) *DescribeAccessPointsRequest {
	s.Type = &v
	return s
}

func (s *DescribeAccessPointsRequest) Validate() error {
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

type DescribeAccessPointsRequestFilter struct {
	// The filter key. Valid values:
	//
	// - `AccessPointId`: Filter by access point ID.
	//
	// - `AccessPointName`: Filter by access point name.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The filter values.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s DescribeAccessPointsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessPointsRequestFilter) GoString() string {
	return s.String()
}

func (s *DescribeAccessPointsRequestFilter) GetKey() *string {
	return s.Key
}

func (s *DescribeAccessPointsRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *DescribeAccessPointsRequestFilter) SetKey(v string) *DescribeAccessPointsRequestFilter {
	s.Key = &v
	return s
}

func (s *DescribeAccessPointsRequestFilter) SetValue(v []*string) *DescribeAccessPointsRequestFilter {
	s.Value = v
	return s
}

func (s *DescribeAccessPointsRequestFilter) Validate() error {
	return dara.Validate(s)
}
