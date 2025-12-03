// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressIPVersion(v string) *DescribeAvailableResourceRequest
	GetAddressIPVersion() *string
	SetAddressType(v string) *DescribeAvailableResourceRequest
	GetAddressType() *string
	SetOwnerAccount(v string) *DescribeAvailableResourceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAvailableResourceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeAvailableResourceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeAvailableResourceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAvailableResourceRequest
	GetResourceOwnerId() *int64
}

type DescribeAvailableResourceRequest struct {
	// The type of the IP address.
	//
	// Valid values: **ipv4 and ipv6**.
	//
	// example:
	//
	// ipv4
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	// The network type.
	//
	// Valid values: **vpc, classic-internet, and classic-intranet**.
	//
	// vpc: an internal Classic Load Balancer (CLB) instance that is deployed in a virtual private cloud (VPC).
	//
	// classic_internet: a public-facing CLB instance.
	//
	// classic_intranet: an internal CLB instance that is deployed in a classic network.
	//
	// example:
	//
	// vpc
	AddressType  *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
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

func (s DescribeAvailableResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceRequest) GetAddressIPVersion() *string {
	return s.AddressIPVersion
}

func (s *DescribeAvailableResourceRequest) GetAddressType() *string {
	return s.AddressType
}

func (s *DescribeAvailableResourceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAvailableResourceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAvailableResourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAvailableResourceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAvailableResourceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAvailableResourceRequest) SetAddressIPVersion(v string) *DescribeAvailableResourceRequest {
	s.AddressIPVersion = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetAddressType(v string) *DescribeAvailableResourceRequest {
	s.AddressType = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetOwnerAccount(v string) *DescribeAvailableResourceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetOwnerId(v int64) *DescribeAvailableResourceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetRegionId(v string) *DescribeAvailableResourceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetResourceOwnerAccount(v string) *DescribeAvailableResourceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetResourceOwnerId(v int64) *DescribeAvailableResourceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) Validate() error {
	return dara.Validate(s)
}
