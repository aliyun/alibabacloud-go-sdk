// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCommonBandwidthPackageSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v string) *ModifyCommonBandwidthPackageSpecRequest
	GetBandwidth() *string
	SetBandwidthPackageId(v string) *ModifyCommonBandwidthPackageSpecRequest
	GetBandwidthPackageId() *string
	SetOwnerAccount(v string) *ModifyCommonBandwidthPackageSpecRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyCommonBandwidthPackageSpecRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyCommonBandwidthPackageSpecRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyCommonBandwidthPackageSpecRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyCommonBandwidthPackageSpecRequest
	GetResourceOwnerId() *int64
}

type ModifyCommonBandwidthPackageSpecRequest struct {
	// The maximum bandwidth of the Internet Shared Bandwidth instance. Unit: Mbit/s.
	//
	// Valid values: **1*	- to **1000**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The ID of the Internet Shared Bandwidth instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cbwp-2ze2ic1xd2qeqk145****
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the Internet Shared Bandwidth instance.
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
}

func (s ModifyCommonBandwidthPackageSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCommonBandwidthPackageSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyCommonBandwidthPackageSpecRequest) GetBandwidth() *string {
	return s.Bandwidth
}

func (s *ModifyCommonBandwidthPackageSpecRequest) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *ModifyCommonBandwidthPackageSpecRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyCommonBandwidthPackageSpecRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyCommonBandwidthPackageSpecRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyCommonBandwidthPackageSpecRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyCommonBandwidthPackageSpecRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyCommonBandwidthPackageSpecRequest) SetBandwidth(v string) *ModifyCommonBandwidthPackageSpecRequest {
	s.Bandwidth = &v
	return s
}

func (s *ModifyCommonBandwidthPackageSpecRequest) SetBandwidthPackageId(v string) *ModifyCommonBandwidthPackageSpecRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *ModifyCommonBandwidthPackageSpecRequest) SetOwnerAccount(v string) *ModifyCommonBandwidthPackageSpecRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyCommonBandwidthPackageSpecRequest) SetOwnerId(v int64) *ModifyCommonBandwidthPackageSpecRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyCommonBandwidthPackageSpecRequest) SetRegionId(v string) *ModifyCommonBandwidthPackageSpecRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyCommonBandwidthPackageSpecRequest) SetResourceOwnerAccount(v string) *ModifyCommonBandwidthPackageSpecRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyCommonBandwidthPackageSpecRequest) SetResourceOwnerId(v int64) *ModifyCommonBandwidthPackageSpecRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyCommonBandwidthPackageSpecRequest) Validate() error {
	return dara.Validate(s)
}
