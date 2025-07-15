// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadVpnConnectionConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DownloadVpnConnectionConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DownloadVpnConnectionConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DownloadVpnConnectionConfigRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DownloadVpnConnectionConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DownloadVpnConnectionConfigRequest
	GetResourceOwnerId() *int64
	SetVpnConnectionId(v string) *DownloadVpnConnectionConfigRequest
	GetVpnConnectionId() *string
}

type DownloadVpnConnectionConfigRequest struct {
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
	// cn-shanghai
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

func (s DownloadVpnConnectionConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DownloadVpnConnectionConfigRequest) GoString() string {
	return s.String()
}

func (s *DownloadVpnConnectionConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DownloadVpnConnectionConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DownloadVpnConnectionConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DownloadVpnConnectionConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DownloadVpnConnectionConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DownloadVpnConnectionConfigRequest) GetVpnConnectionId() *string {
	return s.VpnConnectionId
}

func (s *DownloadVpnConnectionConfigRequest) SetOwnerAccount(v string) *DownloadVpnConnectionConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DownloadVpnConnectionConfigRequest) SetOwnerId(v int64) *DownloadVpnConnectionConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DownloadVpnConnectionConfigRequest) SetRegionId(v string) *DownloadVpnConnectionConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DownloadVpnConnectionConfigRequest) SetResourceOwnerAccount(v string) *DownloadVpnConnectionConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DownloadVpnConnectionConfigRequest) SetResourceOwnerId(v int64) *DownloadVpnConnectionConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DownloadVpnConnectionConfigRequest) SetVpnConnectionId(v string) *DownloadVpnConnectionConfigRequest {
	s.VpnConnectionId = &v
	return s
}

func (s *DownloadVpnConnectionConfigRequest) Validate() error {
	return dara.Validate(s)
}
