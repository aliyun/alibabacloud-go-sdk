// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePublicIpAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpVersion(v string) *DescribePublicIpAddressRequest
	GetIpVersion() *string
	SetOwnerAccount(v string) *DescribePublicIpAddressRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribePublicIpAddressRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribePublicIpAddressRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePublicIpAddressRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribePublicIpAddressRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribePublicIpAddressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribePublicIpAddressRequest
	GetResourceOwnerId() *int64
}

type DescribePublicIpAddressRequest struct {
	// The IP version. Valid values:
	//
	// 	- **IPv4*	- (default)
	//
	// 	- **IPv6**
	//
	// example:
	//
	// ipv4
	IpVersion    *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Valid values: **1*	- to **100**. Default value: **100**.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region that you want to query. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribePublicIpAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePublicIpAddressRequest) GoString() string {
	return s.String()
}

func (s *DescribePublicIpAddressRequest) GetIpVersion() *string {
	return s.IpVersion
}

func (s *DescribePublicIpAddressRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribePublicIpAddressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePublicIpAddressRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePublicIpAddressRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePublicIpAddressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePublicIpAddressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribePublicIpAddressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribePublicIpAddressRequest) SetIpVersion(v string) *DescribePublicIpAddressRequest {
	s.IpVersion = &v
	return s
}

func (s *DescribePublicIpAddressRequest) SetOwnerAccount(v string) *DescribePublicIpAddressRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribePublicIpAddressRequest) SetOwnerId(v int64) *DescribePublicIpAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePublicIpAddressRequest) SetPageNumber(v int32) *DescribePublicIpAddressRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePublicIpAddressRequest) SetPageSize(v int32) *DescribePublicIpAddressRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePublicIpAddressRequest) SetRegionId(v string) *DescribePublicIpAddressRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePublicIpAddressRequest) SetResourceOwnerAccount(v string) *DescribePublicIpAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePublicIpAddressRequest) SetResourceOwnerId(v int64) *DescribePublicIpAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePublicIpAddressRequest) Validate() error {
	return dara.Validate(s)
}
