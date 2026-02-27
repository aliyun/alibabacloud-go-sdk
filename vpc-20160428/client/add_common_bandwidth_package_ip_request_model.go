// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCommonBandwidthPackageIpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidthPackageId(v string) *AddCommonBandwidthPackageIpRequest
	GetBandwidthPackageId() *string
	SetClientToken(v string) *AddCommonBandwidthPackageIpRequest
	GetClientToken() *string
	SetIpInstanceId(v string) *AddCommonBandwidthPackageIpRequest
	GetIpInstanceId() *string
	SetIpType(v string) *AddCommonBandwidthPackageIpRequest
	GetIpType() *string
	SetOwnerAccount(v string) *AddCommonBandwidthPackageIpRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AddCommonBandwidthPackageIpRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AddCommonBandwidthPackageIpRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AddCommonBandwidthPackageIpRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddCommonBandwidthPackageIpRequest
	GetResourceOwnerId() *int64
}

type AddCommonBandwidthPackageIpRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cbwp-2ze2ic1xd2qeqasdf****
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eip-2zeerraiwb7uqwed****
	IpInstanceId *string `json:"IpInstanceId,omitempty" xml:"IpInstanceId,omitempty"`
	// example:
	//
	// EIP
	IpType       *string `json:"IpType,omitempty" xml:"IpType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AddCommonBandwidthPackageIpRequest) String() string {
	return dara.Prettify(s)
}

func (s AddCommonBandwidthPackageIpRequest) GoString() string {
	return s.String()
}

func (s *AddCommonBandwidthPackageIpRequest) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *AddCommonBandwidthPackageIpRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddCommonBandwidthPackageIpRequest) GetIpInstanceId() *string {
	return s.IpInstanceId
}

func (s *AddCommonBandwidthPackageIpRequest) GetIpType() *string {
	return s.IpType
}

func (s *AddCommonBandwidthPackageIpRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddCommonBandwidthPackageIpRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddCommonBandwidthPackageIpRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddCommonBandwidthPackageIpRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddCommonBandwidthPackageIpRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddCommonBandwidthPackageIpRequest) SetBandwidthPackageId(v string) *AddCommonBandwidthPackageIpRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *AddCommonBandwidthPackageIpRequest) SetClientToken(v string) *AddCommonBandwidthPackageIpRequest {
	s.ClientToken = &v
	return s
}

func (s *AddCommonBandwidthPackageIpRequest) SetIpInstanceId(v string) *AddCommonBandwidthPackageIpRequest {
	s.IpInstanceId = &v
	return s
}

func (s *AddCommonBandwidthPackageIpRequest) SetIpType(v string) *AddCommonBandwidthPackageIpRequest {
	s.IpType = &v
	return s
}

func (s *AddCommonBandwidthPackageIpRequest) SetOwnerAccount(v string) *AddCommonBandwidthPackageIpRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddCommonBandwidthPackageIpRequest) SetOwnerId(v int64) *AddCommonBandwidthPackageIpRequest {
	s.OwnerId = &v
	return s
}

func (s *AddCommonBandwidthPackageIpRequest) SetRegionId(v string) *AddCommonBandwidthPackageIpRequest {
	s.RegionId = &v
	return s
}

func (s *AddCommonBandwidthPackageIpRequest) SetResourceOwnerAccount(v string) *AddCommonBandwidthPackageIpRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddCommonBandwidthPackageIpRequest) SetResourceOwnerId(v int64) *AddCommonBandwidthPackageIpRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddCommonBandwidthPackageIpRequest) Validate() error {
	return dara.Validate(s)
}
