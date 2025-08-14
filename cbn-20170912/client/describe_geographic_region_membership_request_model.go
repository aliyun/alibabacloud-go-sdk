// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGeographicRegionMembershipRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGeographicRegionId(v string) *DescribeGeographicRegionMembershipRequest
	GetGeographicRegionId() *string
	SetOwnerAccount(v string) *DescribeGeographicRegionMembershipRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeGeographicRegionMembershipRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeGeographicRegionMembershipRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeGeographicRegionMembershipRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeGeographicRegionMembershipRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeGeographicRegionMembershipRequest
	GetResourceOwnerId() *int64
}

type DescribeGeographicRegionMembershipRequest struct {
	// The ID of the area that you want to query. Valid values:
	//
	// 	- **china**: the Chinese mainland
	//
	// 	- **asia-pacific**: Asia Pacific
	//
	// 	- **europe**: Europe
	//
	// 	- **north-america**: North America
	//
	// This parameter is required.
	//
	// example:
	//
	// china
	GeographicRegionId *string `json:"GeographicRegionId,omitempty" xml:"GeographicRegionId,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return per page. Default value: **10**. Valid values: **1*	- to **50**.
	//
	// example:
	//
	// 10
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeGeographicRegionMembershipRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGeographicRegionMembershipRequest) GoString() string {
	return s.String()
}

func (s *DescribeGeographicRegionMembershipRequest) GetGeographicRegionId() *string {
	return s.GeographicRegionId
}

func (s *DescribeGeographicRegionMembershipRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeGeographicRegionMembershipRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeGeographicRegionMembershipRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeGeographicRegionMembershipRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGeographicRegionMembershipRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeGeographicRegionMembershipRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeGeographicRegionMembershipRequest) SetGeographicRegionId(v string) *DescribeGeographicRegionMembershipRequest {
	s.GeographicRegionId = &v
	return s
}

func (s *DescribeGeographicRegionMembershipRequest) SetOwnerAccount(v string) *DescribeGeographicRegionMembershipRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeGeographicRegionMembershipRequest) SetOwnerId(v int64) *DescribeGeographicRegionMembershipRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeGeographicRegionMembershipRequest) SetPageNumber(v int32) *DescribeGeographicRegionMembershipRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGeographicRegionMembershipRequest) SetPageSize(v int32) *DescribeGeographicRegionMembershipRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGeographicRegionMembershipRequest) SetResourceOwnerAccount(v string) *DescribeGeographicRegionMembershipRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeGeographicRegionMembershipRequest) SetResourceOwnerId(v int64) *DescribeGeographicRegionMembershipRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeGeographicRegionMembershipRequest) Validate() error {
	return dara.Validate(s)
}
