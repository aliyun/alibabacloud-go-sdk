// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLoadBalancerAddressTypeConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressType(v string) *UpdateLoadBalancerAddressTypeConfigRequest
	GetAddressType() *string
	SetClientToken(v string) *UpdateLoadBalancerAddressTypeConfigRequest
	GetClientToken() *string
	SetDryRun(v string) *UpdateLoadBalancerAddressTypeConfigRequest
	GetDryRun() *string
	SetLoadBalancerId(v string) *UpdateLoadBalancerAddressTypeConfigRequest
	GetLoadBalancerId() *string
	SetZoneMappings(v []*UpdateLoadBalancerAddressTypeConfigRequestZoneMappings) *UpdateLoadBalancerAddressTypeConfigRequest
	GetZoneMappings() []*UpdateLoadBalancerAddressTypeConfigRequestZoneMappings
}

type UpdateLoadBalancerAddressTypeConfigRequest struct {
	// The new network type. Valid values:
	//
	// 	- **Internet**: The ALB instance uses a public IP address. The domain name of the ALB instance is resolved to the public IP address. Therefore, the ALB instance can be accessed over the Internet.
	//
	// 	- **Intranet**: The ALB instance uses a private IP address. The domain name of the ALB instance is resolved to the private IP address. In this case, the ALB instance can be accessed over the virtual private cloud (VPC) where the ALB instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// Internet
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ALB instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-bp1o94dp5i6ea****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The zones and the vSwitches in the zones. You can specify a maximum of 10 zones. If the selected region supports two or more zones, select at least two zones to ensure the high availability of your service.
	ZoneMappings []*UpdateLoadBalancerAddressTypeConfigRequestZoneMappings `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty" type:"Repeated"`
}

func (s UpdateLoadBalancerAddressTypeConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLoadBalancerAddressTypeConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerAddressTypeConfigRequest) GetAddressType() *string {
	return s.AddressType
}

func (s *UpdateLoadBalancerAddressTypeConfigRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateLoadBalancerAddressTypeConfigRequest) GetDryRun() *string {
	return s.DryRun
}

func (s *UpdateLoadBalancerAddressTypeConfigRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *UpdateLoadBalancerAddressTypeConfigRequest) GetZoneMappings() []*UpdateLoadBalancerAddressTypeConfigRequestZoneMappings {
	return s.ZoneMappings
}

func (s *UpdateLoadBalancerAddressTypeConfigRequest) SetAddressType(v string) *UpdateLoadBalancerAddressTypeConfigRequest {
	s.AddressType = &v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigRequest) SetClientToken(v string) *UpdateLoadBalancerAddressTypeConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigRequest) SetDryRun(v string) *UpdateLoadBalancerAddressTypeConfigRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigRequest) SetLoadBalancerId(v string) *UpdateLoadBalancerAddressTypeConfigRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigRequest) SetZoneMappings(v []*UpdateLoadBalancerAddressTypeConfigRequestZoneMappings) *UpdateLoadBalancerAddressTypeConfigRequest {
	s.ZoneMappings = v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateLoadBalancerAddressTypeConfigRequestZoneMappings struct {
	// The ID of the elastic IP address (EIP). You can specify a maximum of 10 zones.
	//
	// >  This parameter is required if you want to change the network type from internal-facing to Internet-facing.
	//
	// example:
	//
	// eip-bp1aedxso6u80u0qf****
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// The type of the EIP. Valid values:
	//
	// 	- Common (default): indicates an EIP
	//
	// 	- Anycast: indicates an Anycast EIP
	//
	// >  For more information about the regions in which ALB supports Anycast EIPs, see [Limits](https://help.aliyun.com/document_detail/460727.html).
	//
	// example:
	//
	// Common
	EipType *string `json:"EipType,omitempty" xml:"EipType,omitempty"`
	// The vSwitch in the zone. You can specify only one vSwitch (subnet) in each zone of an ALB instance. You can specify a maximum of 10 zones. If the selected region supports two or more zones, select at least two zones to ensure the high availability of your service.
	//
	// example:
	//
	// vsw-bp10ttov87felojcn****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID of the ALB instance. You can specify a maximum of 10 zones. If the selected region supports two or more zones, select at least two zones to ensure the high availability of your service.
	//
	// You can call the [DescribeZones](https://help.aliyun.com/document_detail/189196.html) operation to query the information about the zone.
	//
	// example:
	//
	// cn-hangzhou-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s UpdateLoadBalancerAddressTypeConfigRequestZoneMappings) String() string {
	return dara.Prettify(s)
}

func (s UpdateLoadBalancerAddressTypeConfigRequestZoneMappings) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerAddressTypeConfigRequestZoneMappings) GetAllocationId() *string {
	return s.AllocationId
}

func (s *UpdateLoadBalancerAddressTypeConfigRequestZoneMappings) GetEipType() *string {
	return s.EipType
}

func (s *UpdateLoadBalancerAddressTypeConfigRequestZoneMappings) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *UpdateLoadBalancerAddressTypeConfigRequestZoneMappings) GetZoneId() *string {
	return s.ZoneId
}

func (s *UpdateLoadBalancerAddressTypeConfigRequestZoneMappings) SetAllocationId(v string) *UpdateLoadBalancerAddressTypeConfigRequestZoneMappings {
	s.AllocationId = &v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigRequestZoneMappings) SetEipType(v string) *UpdateLoadBalancerAddressTypeConfigRequestZoneMappings {
	s.EipType = &v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigRequestZoneMappings) SetVSwitchId(v string) *UpdateLoadBalancerAddressTypeConfigRequestZoneMappings {
	s.VSwitchId = &v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigRequestZoneMappings) SetZoneId(v string) *UpdateLoadBalancerAddressTypeConfigRequestZoneMappings {
	s.ZoneId = &v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigRequestZoneMappings) Validate() error {
	return dara.Validate(s)
}
