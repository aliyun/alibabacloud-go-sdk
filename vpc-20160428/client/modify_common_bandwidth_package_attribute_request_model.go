// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCommonBandwidthPackageAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidthPackageId(v string) *ModifyCommonBandwidthPackageAttributeRequest
	GetBandwidthPackageId() *string
	SetDescription(v string) *ModifyCommonBandwidthPackageAttributeRequest
	GetDescription() *string
	SetName(v string) *ModifyCommonBandwidthPackageAttributeRequest
	GetName() *string
	SetOwnerAccount(v string) *ModifyCommonBandwidthPackageAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyCommonBandwidthPackageAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyCommonBandwidthPackageAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyCommonBandwidthPackageAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyCommonBandwidthPackageAttributeRequest
	GetResourceOwnerId() *int64
}

type ModifyCommonBandwidthPackageAttributeRequest struct {
	// The ID of the EIP bandwidth plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// cbwp-2ze2ic1xd2qeqk145****
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The description of the EIP bandwidth plan. The description must be 2 to 256 characters in length. It must start with a letter but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the EIP bandwidth plan. The name must be 2 to 128 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter. It cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test123
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the EIP bandwidth plan is created.
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

func (s ModifyCommonBandwidthPackageAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCommonBandwidthPackageAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyCommonBandwidthPackageAttributeRequest) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *ModifyCommonBandwidthPackageAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyCommonBandwidthPackageAttributeRequest) GetName() *string {
	return s.Name
}

func (s *ModifyCommonBandwidthPackageAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyCommonBandwidthPackageAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyCommonBandwidthPackageAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyCommonBandwidthPackageAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyCommonBandwidthPackageAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyCommonBandwidthPackageAttributeRequest) SetBandwidthPackageId(v string) *ModifyCommonBandwidthPackageAttributeRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *ModifyCommonBandwidthPackageAttributeRequest) SetDescription(v string) *ModifyCommonBandwidthPackageAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyCommonBandwidthPackageAttributeRequest) SetName(v string) *ModifyCommonBandwidthPackageAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifyCommonBandwidthPackageAttributeRequest) SetOwnerAccount(v string) *ModifyCommonBandwidthPackageAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyCommonBandwidthPackageAttributeRequest) SetOwnerId(v int64) *ModifyCommonBandwidthPackageAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyCommonBandwidthPackageAttributeRequest) SetRegionId(v string) *ModifyCommonBandwidthPackageAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyCommonBandwidthPackageAttributeRequest) SetResourceOwnerAccount(v string) *ModifyCommonBandwidthPackageAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyCommonBandwidthPackageAttributeRequest) SetResourceOwnerId(v int64) *ModifyCommonBandwidthPackageAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyCommonBandwidthPackageAttributeRequest) Validate() error {
	return dara.Validate(s)
}
