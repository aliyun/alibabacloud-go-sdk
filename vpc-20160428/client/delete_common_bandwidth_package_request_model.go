// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCommonBandwidthPackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidthPackageId(v string) *DeleteCommonBandwidthPackageRequest
	GetBandwidthPackageId() *string
	SetForce(v string) *DeleteCommonBandwidthPackageRequest
	GetForce() *string
	SetOwnerAccount(v string) *DeleteCommonBandwidthPackageRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteCommonBandwidthPackageRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteCommonBandwidthPackageRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteCommonBandwidthPackageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteCommonBandwidthPackageRequest
	GetResourceOwnerId() *int64
}

type DeleteCommonBandwidthPackageRequest struct {
	// The ID of the Internet Shared Bandwidth instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cbwp-2ze2ic1xd2qeqk145pn4u
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// Specifies whether to forcefully delete the Internet Shared Bandwidth instance. Valid values:
	//
	// 	- **false*	- (default): deletes the Internet Shared Bandwidth instance only when no EIPs are associated with the Internet Shared Bandwidth instance.
	//
	// 	- **true**: disassociates all EIPs from the Internet Shared Bandwidth instance and deletes the Internet Shared Bandwidth instance.
	//
	// example:
	//
	// false
	Force        *string `json:"Force,omitempty" xml:"Force,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the Internet Shared Bandwidth instance is created.
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

func (s DeleteCommonBandwidthPackageRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCommonBandwidthPackageRequest) GoString() string {
	return s.String()
}

func (s *DeleteCommonBandwidthPackageRequest) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *DeleteCommonBandwidthPackageRequest) GetForce() *string {
	return s.Force
}

func (s *DeleteCommonBandwidthPackageRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteCommonBandwidthPackageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteCommonBandwidthPackageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteCommonBandwidthPackageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteCommonBandwidthPackageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteCommonBandwidthPackageRequest) SetBandwidthPackageId(v string) *DeleteCommonBandwidthPackageRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *DeleteCommonBandwidthPackageRequest) SetForce(v string) *DeleteCommonBandwidthPackageRequest {
	s.Force = &v
	return s
}

func (s *DeleteCommonBandwidthPackageRequest) SetOwnerAccount(v string) *DeleteCommonBandwidthPackageRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteCommonBandwidthPackageRequest) SetOwnerId(v int64) *DeleteCommonBandwidthPackageRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCommonBandwidthPackageRequest) SetRegionId(v string) *DeleteCommonBandwidthPackageRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCommonBandwidthPackageRequest) SetResourceOwnerAccount(v string) *DeleteCommonBandwidthPackageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteCommonBandwidthPackageRequest) SetResourceOwnerId(v int64) *DeleteCommonBandwidthPackageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteCommonBandwidthPackageRequest) Validate() error {
	return dara.Validate(s)
}
