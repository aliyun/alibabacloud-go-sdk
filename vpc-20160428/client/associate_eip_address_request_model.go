// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateEipAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllocationId(v string) *AssociateEipAddressRequest
	GetAllocationId() *string
	SetClientToken(v string) *AssociateEipAddressRequest
	GetClientToken() *string
	SetInstanceId(v string) *AssociateEipAddressRequest
	GetInstanceId() *string
	SetInstanceRegionId(v string) *AssociateEipAddressRequest
	GetInstanceRegionId() *string
	SetInstanceType(v string) *AssociateEipAddressRequest
	GetInstanceType() *string
	SetMode(v string) *AssociateEipAddressRequest
	GetMode() *string
	SetOwnerAccount(v string) *AssociateEipAddressRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AssociateEipAddressRequest
	GetOwnerId() *int64
	SetPrivateIpAddress(v string) *AssociateEipAddressRequest
	GetPrivateIpAddress() *string
	SetRegionId(v string) *AssociateEipAddressRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AssociateEipAddressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AssociateEipAddressRequest
	GetResourceOwnerId() *int64
	SetVpcId(v string) *AssociateEipAddressRequest
	GetVpcId() *string
}

type AssociateEipAddressRequest struct {
	// The ID of the EIP to be associated with the cloud resource instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// eip-2zeerraiwb7ujsxdc****
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe63****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID of the instance to be associated with the EIP.
	//
	// You can enter the instance ID of a NAT gateway, a Classic Load Balancer (CLB) instance, an Elastic Compute Service (ECS) instance, a secondary elastic network interface controller (NIC) instance, a high-availability virtual IP address instance, or an IP address.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-2zebb08phyczzawe****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the instance to be associated with the EIP.
	//
	// > This parameter is required only after the EIP is added to a shared-bandwidth Global Accelerator (GA) instance.
	//
	// example:
	//
	// cn-hangzhou
	InstanceRegionId *string `json:"InstanceRegionId,omitempty" xml:"InstanceRegionId,omitempty"`
	// The type of the instance to be associated with the EIP. Valid values:
	//
	// - **Nat**: NAT gateway.
	//
	// - **SlbInstance**: Classic Load Balancer (CLB).
	//
	// - **EcsInstance*	- (default): Elastic Compute Service (ECS).
	//
	// - **NetworkInterface**: secondary elastic network interface controller (NIC).
	//
	// - **HaVip**: high-availability virtual IP address.
	//
	// - **IpAddress**: IP address.
	//
	// > The default instance type is **EcsInstance**. If the instance type is not **EcsInstance**, this parameter is required.
	//
	// example:
	//
	// EcsInstance
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The association mode. Valid values:
	//
	// - **NAT*	- (default): NAT mode (standard mode).
	//
	// - **MULTI_BINDED**: multi-EIP-to-ENI mode.
	//
	// - **BINDED**: EIP-to-ENI mode.
	//
	//
	// > This parameter is required only when **InstanceType*	- is set to **NetworkInterface**.
	//
	// example:
	//
	// NAT
	Mode         *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// An IP address in the vSwitch CIDR block.
	//
	// If you do not specify this parameter, the system automatically assigns a private IP address based on the VPC ID and vSwitch ID.
	//
	// > If **InstanceType*	- is set to **NetworkInterface**, this parameter is required. Enter the private IP address to be associated.
	//
	// example:
	//
	// 192.168.XX.XX
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The region ID of the EIP to be associated with the cloud resource instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the VPC that has the IPv4 gateway feature enabled and is in the same region as the EIP.
	//
	// When the EIP is associated with an IP address, the system can use the route configuration of the VPC to enable public network access for the associated IP address.
	//
	// > This parameter is required when **InstanceType*	- is set to **IpAddress**.
	//
	// example:
	//
	// vpc-257gqcdfvx6n****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s AssociateEipAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateEipAddressRequest) GoString() string {
	return s.String()
}

func (s *AssociateEipAddressRequest) GetAllocationId() *string {
	return s.AllocationId
}

func (s *AssociateEipAddressRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AssociateEipAddressRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AssociateEipAddressRequest) GetInstanceRegionId() *string {
	return s.InstanceRegionId
}

func (s *AssociateEipAddressRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *AssociateEipAddressRequest) GetMode() *string {
	return s.Mode
}

func (s *AssociateEipAddressRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AssociateEipAddressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AssociateEipAddressRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *AssociateEipAddressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AssociateEipAddressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AssociateEipAddressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AssociateEipAddressRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *AssociateEipAddressRequest) SetAllocationId(v string) *AssociateEipAddressRequest {
	s.AllocationId = &v
	return s
}

func (s *AssociateEipAddressRequest) SetClientToken(v string) *AssociateEipAddressRequest {
	s.ClientToken = &v
	return s
}

func (s *AssociateEipAddressRequest) SetInstanceId(v string) *AssociateEipAddressRequest {
	s.InstanceId = &v
	return s
}

func (s *AssociateEipAddressRequest) SetInstanceRegionId(v string) *AssociateEipAddressRequest {
	s.InstanceRegionId = &v
	return s
}

func (s *AssociateEipAddressRequest) SetInstanceType(v string) *AssociateEipAddressRequest {
	s.InstanceType = &v
	return s
}

func (s *AssociateEipAddressRequest) SetMode(v string) *AssociateEipAddressRequest {
	s.Mode = &v
	return s
}

func (s *AssociateEipAddressRequest) SetOwnerAccount(v string) *AssociateEipAddressRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AssociateEipAddressRequest) SetOwnerId(v int64) *AssociateEipAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *AssociateEipAddressRequest) SetPrivateIpAddress(v string) *AssociateEipAddressRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *AssociateEipAddressRequest) SetRegionId(v string) *AssociateEipAddressRequest {
	s.RegionId = &v
	return s
}

func (s *AssociateEipAddressRequest) SetResourceOwnerAccount(v string) *AssociateEipAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AssociateEipAddressRequest) SetResourceOwnerId(v int64) *AssociateEipAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AssociateEipAddressRequest) SetVpcId(v string) *AssociateEipAddressRequest {
	s.VpcId = &v
	return s
}

func (s *AssociateEipAddressRequest) Validate() error {
	return dara.Validate(s)
}
