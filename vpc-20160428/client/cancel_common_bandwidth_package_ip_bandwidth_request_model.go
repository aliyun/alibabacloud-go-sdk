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
	// The ID of the Internet Shared Bandwidth instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cbwp-bp13d0m4e2qv8xxxxxxxx
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The ID of the EIP that is associated with the Internet Shared Bandwidth instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// eip-2zewysoansu0sxxxxxxxx
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
