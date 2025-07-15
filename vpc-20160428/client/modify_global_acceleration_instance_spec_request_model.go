// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGlobalAccelerationInstanceSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v string) *ModifyGlobalAccelerationInstanceSpecRequest
	GetBandwidth() *string
	SetGlobalAccelerationInstanceId(v string) *ModifyGlobalAccelerationInstanceSpecRequest
	GetGlobalAccelerationInstanceId() *string
	SetOwnerAccount(v string) *ModifyGlobalAccelerationInstanceSpecRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyGlobalAccelerationInstanceSpecRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyGlobalAccelerationInstanceSpecRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyGlobalAccelerationInstanceSpecRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyGlobalAccelerationInstanceSpecRequest
	GetResourceOwnerId() *int64
}

type ModifyGlobalAccelerationInstanceSpecRequest struct {
	// The maximum bandwidth of the GA instance. Unit: Mbit/s. Set the value to **10**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The ID of the GA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-32s33s****
	GlobalAccelerationInstanceId *string `json:"GlobalAccelerationInstanceId,omitempty" xml:"GlobalAccelerationInstanceId,omitempty"`
	OwnerAccount                 *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the GA instance.
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

func (s ModifyGlobalAccelerationInstanceSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyGlobalAccelerationInstanceSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyGlobalAccelerationInstanceSpecRequest) GetBandwidth() *string {
	return s.Bandwidth
}

func (s *ModifyGlobalAccelerationInstanceSpecRequest) GetGlobalAccelerationInstanceId() *string {
	return s.GlobalAccelerationInstanceId
}

func (s *ModifyGlobalAccelerationInstanceSpecRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyGlobalAccelerationInstanceSpecRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyGlobalAccelerationInstanceSpecRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyGlobalAccelerationInstanceSpecRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyGlobalAccelerationInstanceSpecRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyGlobalAccelerationInstanceSpecRequest) SetBandwidth(v string) *ModifyGlobalAccelerationInstanceSpecRequest {
	s.Bandwidth = &v
	return s
}

func (s *ModifyGlobalAccelerationInstanceSpecRequest) SetGlobalAccelerationInstanceId(v string) *ModifyGlobalAccelerationInstanceSpecRequest {
	s.GlobalAccelerationInstanceId = &v
	return s
}

func (s *ModifyGlobalAccelerationInstanceSpecRequest) SetOwnerAccount(v string) *ModifyGlobalAccelerationInstanceSpecRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyGlobalAccelerationInstanceSpecRequest) SetOwnerId(v int64) *ModifyGlobalAccelerationInstanceSpecRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyGlobalAccelerationInstanceSpecRequest) SetRegionId(v string) *ModifyGlobalAccelerationInstanceSpecRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyGlobalAccelerationInstanceSpecRequest) SetResourceOwnerAccount(v string) *ModifyGlobalAccelerationInstanceSpecRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyGlobalAccelerationInstanceSpecRequest) SetResourceOwnerId(v int64) *ModifyGlobalAccelerationInstanceSpecRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyGlobalAccelerationInstanceSpecRequest) Validate() error {
	return dara.Validate(s)
}
