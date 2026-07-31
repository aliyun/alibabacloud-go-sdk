// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceVpcAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyInstanceVpcAttributeRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *ModifyInstanceVpcAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyInstanceVpcAttributeRequest
	GetOwnerId() *int64
	SetPrivateIpAddress(v string) *ModifyInstanceVpcAttributeRequest
	GetPrivateIpAddress() *string
	SetResourceOwnerAccount(v string) *ModifyInstanceVpcAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyInstanceVpcAttributeRequest
	GetResourceOwnerId() *int64
	SetSecurityGroupId(v []*string) *ModifyInstanceVpcAttributeRequest
	GetSecurityGroupId() []*string
	SetVSwitchId(v string) *ModifyInstanceVpcAttributeRequest
	GetVSwitchId() *string
	SetVpcId(v string) *ModifyInstanceVpcAttributeRequest
	GetVpcId() *string
}

type ModifyInstanceVpcAttributeRequest struct {
	// The instance ID.
	//
	// > When you call this operation, the ECS instance must be in the **Stopped*	- state. For other restrictions on the instance, carefully read the **operation description*	- section.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp1iudwa5b1tqag1****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The new private IP address.
	//
	// > The `PrivateIpAddress` parameter depends on `VSwitchId`. The specified IP address must be within the CIDR block of the vSwitch.
	//
	// Default value: If this parameter is not specified, a private IP address is randomly assigned from the CIDR block of the vSwitch.
	//
	// example:
	//
	// ``172.17.**.**``
	PrivateIpAddress     *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The IDs of the security groups to which the instance is added after the VPC is changed. This parameter is required only when the `VpcId` parameter is specified.
	//
	// - The security groups must belong to the destination VPC.
	//
	// - You can specify one or more security groups. The number of security groups is subject to the limits on the number of security groups to which an instance can belong. For more information, see [Limits](~~25412#SecurityGroupQuota1~~).
	//
	// - All security groups in the list must be of the same type.
	//
	// - Switching between security group types is supported. When you switch an ECS instance between security group types, make sure that you understand the differences in security group rule configurations between the two types to avoid impacts on instance networking. For more information, see [Security group overview](https://help.aliyun.com/document_detail/25387.html).
	//
	// example:
	//
	// sg-o6w9l8bc8dgmkw87****
	SecurityGroupId []*string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Repeated"`
	// The vSwitch ID.
	//
	// - If the specified ID is the current vSwitch of the instance, the vSwitch remains unchanged.
	//
	// - If the specified ID is a new vSwitch and the `VpcId` parameter is empty, the new and old vSwitches must belong to the same zone and the same VPC.
	//
	// - If the `VpcId` parameter is not empty, the vSwitch specified by this parameter must belong to the specified VPC and must be in the same zone as the original vSwitch.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1s5fnvk4gn3tw12****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the destination VPC.
	//
	// example:
	//
	// vpc-bp1vwnn14rqpyiczj****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ModifyInstanceVpcAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceVpcAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceVpcAttributeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceVpcAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyInstanceVpcAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyInstanceVpcAttributeRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *ModifyInstanceVpcAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyInstanceVpcAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyInstanceVpcAttributeRequest) GetSecurityGroupId() []*string {
	return s.SecurityGroupId
}

func (s *ModifyInstanceVpcAttributeRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ModifyInstanceVpcAttributeRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ModifyInstanceVpcAttributeRequest) SetInstanceId(v string) *ModifyInstanceVpcAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceVpcAttributeRequest) SetOwnerAccount(v string) *ModifyInstanceVpcAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceVpcAttributeRequest) SetOwnerId(v int64) *ModifyInstanceVpcAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceVpcAttributeRequest) SetPrivateIpAddress(v string) *ModifyInstanceVpcAttributeRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *ModifyInstanceVpcAttributeRequest) SetResourceOwnerAccount(v string) *ModifyInstanceVpcAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceVpcAttributeRequest) SetResourceOwnerId(v int64) *ModifyInstanceVpcAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceVpcAttributeRequest) SetSecurityGroupId(v []*string) *ModifyInstanceVpcAttributeRequest {
	s.SecurityGroupId = v
	return s
}

func (s *ModifyInstanceVpcAttributeRequest) SetVSwitchId(v string) *ModifyInstanceVpcAttributeRequest {
	s.VSwitchId = &v
	return s
}

func (s *ModifyInstanceVpcAttributeRequest) SetVpcId(v string) *ModifyInstanceVpcAttributeRequest {
	s.VpcId = &v
	return s
}

func (s *ModifyInstanceVpcAttributeRequest) Validate() error {
	return dara.Validate(s)
}
