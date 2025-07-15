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
	// The maximum bandwidth for the EIP. This value cannot be larger than the maximum bandwidth of the Internet Shared Bandwidth instance. Unit: Mbit/s.
	//
	// This parameter is required.
	//
	// example:
	//
	// 500
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The ID of the Internet Shared Bandwidth instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cbwp-2zep6hw5d6y8exscd****
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The ID of the EIP that is associated with the Internet Shared Bandwidth instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// eip-2zewysoansu0svfbg****
	EipId        *string `json:"EipId,omitempty" xml:"EipId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the Internet Shared Bandwidth instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
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
