// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassociateHaVipRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UnassociateHaVipRequest
	GetClientToken() *string
	SetForce(v string) *UnassociateHaVipRequest
	GetForce() *string
	SetHaVipId(v string) *UnassociateHaVipRequest
	GetHaVipId() *string
	SetInstanceId(v string) *UnassociateHaVipRequest
	GetInstanceId() *string
	SetInstanceType(v string) *UnassociateHaVipRequest
	GetInstanceType() *string
	SetOwnerAccount(v string) *UnassociateHaVipRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UnassociateHaVipRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UnassociateHaVipRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UnassociateHaVipRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UnassociateHaVipRequest
	GetResourceOwnerId() *int64
}

type UnassociateHaVipRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The `token` can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe63****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to forcefully disassociate the HAVIP from the ECS instance or ENI. Valid values:
	//
	// 	- **True**
	//
	// 	- **False*	- (default)
	//
	// >  If you set the value to **False**, you cannot disassociate the HAVIP from the primary instance.
	//
	// example:
	//
	// True
	Force *string `json:"Force,omitempty" xml:"Force,omitempty"`
	// The ID of the HAVIP that you want to disassociate.
	//
	// This parameter is required.
	//
	// example:
	//
	// havip-2zeo05qre24nhrqpy****
	HaVipId *string `json:"HaVipId,omitempty" xml:"HaVipId,omitempty"`
	// The ID of the ECS instance or ENI from which you want to disassociate the HAVIP.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-faf344422ffsfad****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the instance from which you want to disassociate the HAVIP. Valid values:
	//
	// 	- **EcsInstance**: an ECS instance
	//
	// 	- **NetworkInterface**: an ENI
	//
	// >  If you want to disassociate the HAVIP from an ENI, this parameter is required.
	//
	// example:
	//
	// EcsInstance
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the HAVIP.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UnassociateHaVipRequest) String() string {
	return dara.Prettify(s)
}

func (s UnassociateHaVipRequest) GoString() string {
	return s.String()
}

func (s *UnassociateHaVipRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UnassociateHaVipRequest) GetForce() *string {
	return s.Force
}

func (s *UnassociateHaVipRequest) GetHaVipId() *string {
	return s.HaVipId
}

func (s *UnassociateHaVipRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UnassociateHaVipRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *UnassociateHaVipRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UnassociateHaVipRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnassociateHaVipRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UnassociateHaVipRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UnassociateHaVipRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UnassociateHaVipRequest) SetClientToken(v string) *UnassociateHaVipRequest {
	s.ClientToken = &v
	return s
}

func (s *UnassociateHaVipRequest) SetForce(v string) *UnassociateHaVipRequest {
	s.Force = &v
	return s
}

func (s *UnassociateHaVipRequest) SetHaVipId(v string) *UnassociateHaVipRequest {
	s.HaVipId = &v
	return s
}

func (s *UnassociateHaVipRequest) SetInstanceId(v string) *UnassociateHaVipRequest {
	s.InstanceId = &v
	return s
}

func (s *UnassociateHaVipRequest) SetInstanceType(v string) *UnassociateHaVipRequest {
	s.InstanceType = &v
	return s
}

func (s *UnassociateHaVipRequest) SetOwnerAccount(v string) *UnassociateHaVipRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UnassociateHaVipRequest) SetOwnerId(v int64) *UnassociateHaVipRequest {
	s.OwnerId = &v
	return s
}

func (s *UnassociateHaVipRequest) SetRegionId(v string) *UnassociateHaVipRequest {
	s.RegionId = &v
	return s
}

func (s *UnassociateHaVipRequest) SetResourceOwnerAccount(v string) *UnassociateHaVipRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnassociateHaVipRequest) SetResourceOwnerId(v int64) *UnassociateHaVipRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnassociateHaVipRequest) Validate() error {
	return dara.Validate(s)
}
