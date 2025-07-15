// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpnConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeVpnConnectionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeVpnConnectionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeVpnConnectionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeVpnConnectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeVpnConnectionRequest
	GetResourceOwnerId() *int64
	SetVpnConnectionId(v string) *DescribeVpnConnectionRequest
	GetVpnConnectionId() *string
}

type DescribeVpnConnectionRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the IPsec-VPN connection is created.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the IPsec-VPN connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// vco-bp1bbi27hojx80nck****
	VpnConnectionId *string `json:"VpnConnectionId,omitempty" xml:"VpnConnectionId,omitempty"`
}

func (s DescribeVpnConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnConnectionRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpnConnectionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeVpnConnectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVpnConnectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVpnConnectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeVpnConnectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeVpnConnectionRequest) GetVpnConnectionId() *string {
	return s.VpnConnectionId
}

func (s *DescribeVpnConnectionRequest) SetOwnerAccount(v string) *DescribeVpnConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeVpnConnectionRequest) SetOwnerId(v int64) *DescribeVpnConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVpnConnectionRequest) SetRegionId(v string) *DescribeVpnConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVpnConnectionRequest) SetResourceOwnerAccount(v string) *DescribeVpnConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVpnConnectionRequest) SetResourceOwnerId(v int64) *DescribeVpnConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVpnConnectionRequest) SetVpnConnectionId(v string) *DescribeVpnConnectionRequest {
	s.VpnConnectionId = &v
	return s
}

func (s *DescribeVpnConnectionRequest) Validate() error {
	return dara.Validate(s)
}
