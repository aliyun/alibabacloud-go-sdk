// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCommonBandwidthPackageIpBandwidthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidthPackageId(v string) *CancelCommonBandwidthPackageIpBandwidthRequest
	GetBandwidthPackageId() *string
	SetEipId(v string) *CancelCommonBandwidthPackageIpBandwidthRequest
	GetEipId() *string
	SetIpInstanceId(v string) *CancelCommonBandwidthPackageIpBandwidthRequest
	GetIpInstanceId() *string
	SetIpType(v string) *CancelCommonBandwidthPackageIpBandwidthRequest
	GetIpType() *string
	SetOwnerAccount(v string) *CancelCommonBandwidthPackageIpBandwidthRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CancelCommonBandwidthPackageIpBandwidthRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CancelCommonBandwidthPackageIpBandwidthRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CancelCommonBandwidthPackageIpBandwidthRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CancelCommonBandwidthPackageIpBandwidthRequest
	GetResourceOwnerId() *int64
}

type CancelCommonBandwidthPackageIpBandwidthRequest struct {
	// The instance ID of the Internet Shared Bandwidth.
	//
	// This parameter is required.
	//
	// example:
	//
	// cbwp-bp13d0m4e2qv8xxxxxxxx
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The ID of the EIP that is added to the Internet Shared Bandwidth instance. This parameter applies only to EIPs. For IPv6 addresses, use IpInstanceId.
	//
	// example:
	//
	// eip-2zewysoansu0sxxxxxxxx
	EipId *string `json:"EipId,omitempty" xml:"EipId,omitempty"`
	// The ID of the public IP address instance for which you want to cancel the maximum available bandwidth setting in the Internet Shared Bandwidth instance. This parameter and `EipId` are mutually exclusive. We recommend that you use this parameter for new workloads. If IpType is set to **VpcIpv6**, specify the IPv6 instance ID.
	//
	// example:
	//
	// eip-2zewysoansu0svfbg****
	IpInstanceId *string `json:"IpInstanceId,omitempty" xml:"IpInstanceId,omitempty"`
	// The resource type of the public IP address for which you want to cancel the maximum available bandwidth setting in the Internet Shared Bandwidth instance. Valid values:
	//
	// - **EIP**: the public IP address of an EIP.
	//
	// - **VpcIpv6**: the IPv6 address of a VPC.
	//
	// If you do not specify this parameter, **EIP*	- is used.
	//
	// example:
	//
	// EIP
	IpType       *string `json:"IpType,omitempty" xml:"IpType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the Internet Shared Bandwidth instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query region IDs.
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

func (s CancelCommonBandwidthPackageIpBandwidthRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelCommonBandwidthPackageIpBandwidthRequest) GoString() string {
	return s.String()
}

func (s *CancelCommonBandwidthPackageIpBandwidthRequest) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *CancelCommonBandwidthPackageIpBandwidthRequest) GetEipId() *string {
	return s.EipId
}

func (s *CancelCommonBandwidthPackageIpBandwidthRequest) GetIpInstanceId() *string {
	return s.IpInstanceId
}

func (s *CancelCommonBandwidthPackageIpBandwidthRequest) GetIpType() *string {
	return s.IpType
}

func (s *CancelCommonBandwidthPackageIpBandwidthRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CancelCommonBandwidthPackageIpBandwidthRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CancelCommonBandwidthPackageIpBandwidthRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CancelCommonBandwidthPackageIpBandwidthRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CancelCommonBandwidthPackageIpBandwidthRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CancelCommonBandwidthPackageIpBandwidthRequest) SetBandwidthPackageId(v string) *CancelCommonBandwidthPackageIpBandwidthRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *CancelCommonBandwidthPackageIpBandwidthRequest) SetEipId(v string) *CancelCommonBandwidthPackageIpBandwidthRequest {
	s.EipId = &v
	return s
}

func (s *CancelCommonBandwidthPackageIpBandwidthRequest) SetIpInstanceId(v string) *CancelCommonBandwidthPackageIpBandwidthRequest {
	s.IpInstanceId = &v
	return s
}

func (s *CancelCommonBandwidthPackageIpBandwidthRequest) SetIpType(v string) *CancelCommonBandwidthPackageIpBandwidthRequest {
	s.IpType = &v
	return s
}

func (s *CancelCommonBandwidthPackageIpBandwidthRequest) SetOwnerAccount(v string) *CancelCommonBandwidthPackageIpBandwidthRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CancelCommonBandwidthPackageIpBandwidthRequest) SetOwnerId(v int64) *CancelCommonBandwidthPackageIpBandwidthRequest {
	s.OwnerId = &v
	return s
}

func (s *CancelCommonBandwidthPackageIpBandwidthRequest) SetRegionId(v string) *CancelCommonBandwidthPackageIpBandwidthRequest {
	s.RegionId = &v
	return s
}

func (s *CancelCommonBandwidthPackageIpBandwidthRequest) SetResourceOwnerAccount(v string) *CancelCommonBandwidthPackageIpBandwidthRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CancelCommonBandwidthPackageIpBandwidthRequest) SetResourceOwnerId(v int64) *CancelCommonBandwidthPackageIpBandwidthRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CancelCommonBandwidthPackageIpBandwidthRequest) Validate() error {
	return dara.Validate(s)
}
