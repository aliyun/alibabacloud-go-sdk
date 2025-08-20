// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserEniVswitchOptionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbClusterId(v string) *ModifyUserEniVswitchOptionsRequest
	GetDbClusterId() *string
	SetOwnerAccount(v string) *ModifyUserEniVswitchOptionsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyUserEniVswitchOptionsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyUserEniVswitchOptionsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyUserEniVswitchOptionsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyUserEniVswitchOptionsRequest
	GetResourceOwnerId() *int64
	SetVSwitchOptions(v string) *ModifyUserEniVswitchOptionsRequest
	GetVSwitchOptions() *string
}

type ModifyUserEniVswitchOptionsRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DbClusterId  *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/454314.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The vSwitches that you want to use. The vSwitches must reside in the same virtual private cloud (VPC) and zone as ENIs. You can specify up to three vSwitches. Separate multiple vSwitches with commas (,).
	//
	// >
	//
	// 	- The vSwitches that you specify overwrite the existing vSwitches that are connected to ENIs.
	//
	// 	- You can call the [DescribeDBClusterAttribute](https://help.aliyun.com/document_detail/612399.html) operation to query the network information about ENIs in a cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-rj9ixufmywqq98z******,vsw-rj95ij6wcz656v7******
	VSwitchOptions *string `json:"VSwitchOptions,omitempty" xml:"VSwitchOptions,omitempty"`
}

func (s ModifyUserEniVswitchOptionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserEniVswitchOptionsRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserEniVswitchOptionsRequest) GetDbClusterId() *string {
	return s.DbClusterId
}

func (s *ModifyUserEniVswitchOptionsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyUserEniVswitchOptionsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyUserEniVswitchOptionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyUserEniVswitchOptionsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyUserEniVswitchOptionsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyUserEniVswitchOptionsRequest) GetVSwitchOptions() *string {
	return s.VSwitchOptions
}

func (s *ModifyUserEniVswitchOptionsRequest) SetDbClusterId(v string) *ModifyUserEniVswitchOptionsRequest {
	s.DbClusterId = &v
	return s
}

func (s *ModifyUserEniVswitchOptionsRequest) SetOwnerAccount(v string) *ModifyUserEniVswitchOptionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyUserEniVswitchOptionsRequest) SetOwnerId(v int64) *ModifyUserEniVswitchOptionsRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyUserEniVswitchOptionsRequest) SetRegionId(v string) *ModifyUserEniVswitchOptionsRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyUserEniVswitchOptionsRequest) SetResourceOwnerAccount(v string) *ModifyUserEniVswitchOptionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyUserEniVswitchOptionsRequest) SetResourceOwnerId(v int64) *ModifyUserEniVswitchOptionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyUserEniVswitchOptionsRequest) SetVSwitchOptions(v string) *ModifyUserEniVswitchOptionsRequest {
	s.VSwitchOptions = &v
	return s
}

func (s *ModifyUserEniVswitchOptionsRequest) Validate() error {
	return dara.Validate(s)
}
