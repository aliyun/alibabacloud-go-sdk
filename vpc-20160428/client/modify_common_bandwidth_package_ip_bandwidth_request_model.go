// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCommonBandwidthPackageIpBandwidthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v string) *ModifyCommonBandwidthPackageIpBandwidthRequest
	GetBandwidth() *string
	SetBandwidthPackageId(v string) *ModifyCommonBandwidthPackageIpBandwidthRequest
	GetBandwidthPackageId() *string
	SetEipId(v string) *ModifyCommonBandwidthPackageIpBandwidthRequest
	GetEipId() *string
	SetIpInstanceId(v string) *ModifyCommonBandwidthPackageIpBandwidthRequest
	GetIpInstanceId() *string
	SetIpType(v string) *ModifyCommonBandwidthPackageIpBandwidthRequest
	GetIpType() *string
	SetOwnerAccount(v string) *ModifyCommonBandwidthPackageIpBandwidthRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyCommonBandwidthPackageIpBandwidthRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyCommonBandwidthPackageIpBandwidthRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyCommonBandwidthPackageIpBandwidthRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyCommonBandwidthPackageIpBandwidthRequest
	GetResourceOwnerId() *int64
}

type ModifyCommonBandwidthPackageIpBandwidthRequest struct {
	// The maximum bandwidth that can be allocated in the Internet Shared Bandwidth instance. Unit: Mbit/s.
	//
	// This parameter is required.
	//
	// example:
	//
	// 500
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The instance ID of the Internet Shared Bandwidth.
	//
	// This parameter is required.
	//
	// example:
	//
	// cbwp-2zep6hw5d6y8exscd****
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The ID of the EIP that is added to the Internet Shared Bandwidth instance. This parameter applies only to EIPs. For IPv6 addresses, use IpInstanceId.
	//
	// example:
	//
	// eip-2zewysoansu0svfbg****
	EipId *string `json:"EipId,omitempty" xml:"EipId,omitempty"`
	// The ID of the public IP address instance for which you want to set the maximum available bandwidth in the Internet Shared Bandwidth instance. This parameter and EipId are mutually exclusive. We recommend that you use this parameter for new workloads. If IpType is set to **VpcIpv6**, specify the ID of the IPv6 address.
	//
	// example:
	//
	// eip-2zewysoansu0svfbg****
	IpInstanceId *string `json:"IpInstanceId,omitempty" xml:"IpInstanceId,omitempty"`
	// The resource type of the public IP address for which you want to set the maximum available bandwidth in the Internet Shared Bandwidth instance. Valid values:
	//
	// - **EIP**: the public IP address of an EIP.
	//
	// - **VpcIpv6**: the IPv6 address of a VPC.
	//
	// If this parameter is not specified, the default value is **EIP**.
	//
	// example:
	//
	// EIP
	IpType       *string `json:"IpType,omitempty" xml:"IpType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the Internet Shared Bandwidth instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
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

func (s ModifyCommonBandwidthPackageIpBandwidthRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCommonBandwidthPackageIpBandwidthRequest) GoString() string {
	return s.String()
}

func (s *ModifyCommonBandwidthPackageIpBandwidthRequest) GetBandwidth() *string {
	return s.Bandwidth
}

func (s *ModifyCommonBandwidthPackageIpBandwidthRequest) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *ModifyCommonBandwidthPackageIpBandwidthRequest) GetEipId() *string {
	return s.EipId
}

func (s *ModifyCommonBandwidthPackageIpBandwidthRequest) GetIpInstanceId() *string {
	return s.IpInstanceId
}

func (s *ModifyCommonBandwidthPackageIpBandwidthRequest) GetIpType() *string {
	return s.IpType
}

func (s *ModifyCommonBandwidthPackageIpBandwidthRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyCommonBandwidthPackageIpBandwidthRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyCommonBandwidthPackageIpBandwidthRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyCommonBandwidthPackageIpBandwidthRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyCommonBandwidthPackageIpBandwidthRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyCommonBandwidthPackageIpBandwidthRequest) SetBandwidth(v string) *ModifyCommonBandwidthPackageIpBandwidthRequest {
	s.Bandwidth = &v
	return s
}

func (s *ModifyCommonBandwidthPackageIpBandwidthRequest) SetBandwidthPackageId(v string) *ModifyCommonBandwidthPackageIpBandwidthRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *ModifyCommonBandwidthPackageIpBandwidthRequest) SetEipId(v string) *ModifyCommonBandwidthPackageIpBandwidthRequest {
	s.EipId = &v
	return s
}

func (s *ModifyCommonBandwidthPackageIpBandwidthRequest) SetIpInstanceId(v string) *ModifyCommonBandwidthPackageIpBandwidthRequest {
	s.IpInstanceId = &v
	return s
}

func (s *ModifyCommonBandwidthPackageIpBandwidthRequest) SetIpType(v string) *ModifyCommonBandwidthPackageIpBandwidthRequest {
	s.IpType = &v
	return s
}

func (s *ModifyCommonBandwidthPackageIpBandwidthRequest) SetOwnerAccount(v string) *ModifyCommonBandwidthPackageIpBandwidthRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyCommonBandwidthPackageIpBandwidthRequest) SetOwnerId(v int64) *ModifyCommonBandwidthPackageIpBandwidthRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyCommonBandwidthPackageIpBandwidthRequest) SetRegionId(v string) *ModifyCommonBandwidthPackageIpBandwidthRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyCommonBandwidthPackageIpBandwidthRequest) SetResourceOwnerAccount(v string) *ModifyCommonBandwidthPackageIpBandwidthRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyCommonBandwidthPackageIpBandwidthRequest) SetResourceOwnerId(v int64) *ModifyCommonBandwidthPackageIpBandwidthRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyCommonBandwidthPackageIpBandwidthRequest) Validate() error {
	return dara.Validate(s)
}
