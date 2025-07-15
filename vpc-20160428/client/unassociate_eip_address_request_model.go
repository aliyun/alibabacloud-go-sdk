// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassociateEipAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllocationId(v string) *UnassociateEipAddressRequest
	GetAllocationId() *string
	SetClientToken(v string) *UnassociateEipAddressRequest
	GetClientToken() *string
	SetForce(v bool) *UnassociateEipAddressRequest
	GetForce() *bool
	SetInstanceId(v string) *UnassociateEipAddressRequest
	GetInstanceId() *string
	SetInstanceType(v string) *UnassociateEipAddressRequest
	GetInstanceType() *string
	SetOwnerAccount(v string) *UnassociateEipAddressRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UnassociateEipAddressRequest
	GetOwnerId() *int64
	SetPrivateIpAddress(v string) *UnassociateEipAddressRequest
	GetPrivateIpAddress() *string
	SetRegionId(v string) *UnassociateEipAddressRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UnassociateEipAddressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UnassociateEipAddressRequest
	GetResourceOwnerId() *int64
}

type UnassociateEipAddressRequest struct {
	// The ID of the EIP that you want to disassociate.
	//
	// This parameter is required.
	//
	// example:
	//
	// eip-2zeerraiwb7uj6i0d****
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to disassociate the EIP from a NAT gateway if a DNAT or SNAT entry is added to the NAT gateway. Valid values:
	//
	// 	- **false*	- (default)
	//
	// 	- **true**
	//
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// The ID of the instance from which you want to disassociate the EIP.
	//
	// example:
	//
	// i-hp3akk9irtd69jad****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of instance from which you want to disassociate the EIP. Valid values:
	//
	// 	- **EcsInstance*	- (default): an Elastic Compute Service (ECS) instance in a virtual private cloud (VPC)
	//
	// 	- **SlbInstance**: a Server Load Balancer (SLB) instance in a VPC
	//
	// 	- **NetworkInterface**: a secondary elastic network interface (ENI) in a VPC
	//
	// 	- **Nat**: a NAT gateway
	//
	// 	- **HaVip**: a high-availability virtual IP address (HAVIP)
	//
	// example:
	//
	// EcsInstance
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The private IP address of the ECS instance or the secondary ENI from which you want to disassociate the EIP.
	//
	// example:
	//
	// 192.XX.XX.2
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The ID of the region to which the EIP belongs. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UnassociateEipAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s UnassociateEipAddressRequest) GoString() string {
	return s.String()
}

func (s *UnassociateEipAddressRequest) GetAllocationId() *string {
	return s.AllocationId
}

func (s *UnassociateEipAddressRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UnassociateEipAddressRequest) GetForce() *bool {
	return s.Force
}

func (s *UnassociateEipAddressRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UnassociateEipAddressRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *UnassociateEipAddressRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UnassociateEipAddressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnassociateEipAddressRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *UnassociateEipAddressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UnassociateEipAddressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UnassociateEipAddressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UnassociateEipAddressRequest) SetAllocationId(v string) *UnassociateEipAddressRequest {
	s.AllocationId = &v
	return s
}

func (s *UnassociateEipAddressRequest) SetClientToken(v string) *UnassociateEipAddressRequest {
	s.ClientToken = &v
	return s
}

func (s *UnassociateEipAddressRequest) SetForce(v bool) *UnassociateEipAddressRequest {
	s.Force = &v
	return s
}

func (s *UnassociateEipAddressRequest) SetInstanceId(v string) *UnassociateEipAddressRequest {
	s.InstanceId = &v
	return s
}

func (s *UnassociateEipAddressRequest) SetInstanceType(v string) *UnassociateEipAddressRequest {
	s.InstanceType = &v
	return s
}

func (s *UnassociateEipAddressRequest) SetOwnerAccount(v string) *UnassociateEipAddressRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UnassociateEipAddressRequest) SetOwnerId(v int64) *UnassociateEipAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *UnassociateEipAddressRequest) SetPrivateIpAddress(v string) *UnassociateEipAddressRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *UnassociateEipAddressRequest) SetRegionId(v string) *UnassociateEipAddressRequest {
	s.RegionId = &v
	return s
}

func (s *UnassociateEipAddressRequest) SetResourceOwnerAccount(v string) *UnassociateEipAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnassociateEipAddressRequest) SetResourceOwnerId(v int64) *UnassociateEipAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnassociateEipAddressRequest) Validate() error {
	return dara.Validate(s)
}
