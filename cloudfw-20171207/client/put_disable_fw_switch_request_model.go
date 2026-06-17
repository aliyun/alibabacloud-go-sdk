// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutDisableFwSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpVersion(v string) *PutDisableFwSwitchRequest
	GetIpVersion() *string
	SetIpaddrList(v []*string) *PutDisableFwSwitchRequest
	GetIpaddrList() []*string
	SetLang(v string) *PutDisableFwSwitchRequest
	GetLang() *string
	SetMemberUid(v string) *PutDisableFwSwitchRequest
	GetMemberUid() *string
	SetRegionList(v []*string) *PutDisableFwSwitchRequest
	GetRegionList() []*string
	SetResourceTypeList(v []*string) *PutDisableFwSwitchRequest
	GetResourceTypeList() []*string
	SetSourceIp(v string) *PutDisableFwSwitchRequest
	GetSourceIp() *string
}

type PutDisableFwSwitchRequest struct {
	// The IP version.
	//
	// example:
	//
	// 4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The IP addresses.
	//
	// > You must specify a value for at least one of the following parameters: `IpaddrList`, `RegionList`, and `ResourceTypeList`.
	//
	// example:
	//
	// ["192.0.XX.XX","192.0.XX.XX"]
	IpaddrList []*string `json:"IpaddrList,omitempty" xml:"IpaddrList,omitempty" type:"Repeated"`
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The unique identifier of the member.
	//
	// example:
	//
	// 1234
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The regions.
	//
	// > You must specify a value for at least one of the following parameters: `IpaddrList`, `RegionList`, and `ResourceTypeList`.
	//
	// example:
	//
	// ["cn-hangzhou","cn-shanghai"]
	RegionList []*string `json:"RegionList,omitempty" xml:"RegionList,omitempty" type:"Repeated"`
	// The asset types. Valid values:
	//
	// - **BastionHostEgressIP**: The egress IP address of a bastion host.
	//
	// - **BastionHostIngressIP**: The ingress IP address of a bastion host.
	//
	// - **EcsEIP**: The Elastic IP Address (EIP) of an ECS instance.
	//
	// - **EcsPublicIP**: The public IP address of an ECS instance.
	//
	// - **EIP**: An Elastic IP Address (EIP).
	//
	// - **EniEIP**: The EIP of an elastic network interface (ENI).
	//
	// - **NatEIP**: The EIP of a NAT Gateway instance.
	//
	// - **SlbEIP**: The EIP of a Server Load Balancer (SLB) or Classic Load Balancer (CLB) instance.
	//
	// - **SlbPublicIP**: The public IP address of an SLB or CLB instance.
	//
	// - **NatPublicIP**: The public IP address of a NAT Gateway instance.
	//
	// - **HAVIP**: A High-availability Virtual IP (HAVIP).
	//
	// - **NlbEIP**: The EIP of a Network Load Balancer (NLB) instance.
	//
	// - **ApiGatewayEIP**: The public IP address of an API Gateway instance.
	//
	// - **AlbEIP**: The EIP of an Application Load Balancer (ALB) instance.
	//
	// - **AiGatewayEIP**: The public IP address of an AI Gateway instance.
	//
	// - **GaEIP**: The EIP of a Global Accelerator (GA) instance.
	//
	// - **SwasEIP**: The public IP address of a Simple Application Server instance.
	//
	// - **EcdEIP**: The public IP address of an Elastic Desktop Service (ECD) instance.
	//
	// - **BastionHostIP**: The IP address of a bastion host.
	//
	// > You must specify a value for at least one of the following parameters: `IpaddrList`, `RegionList`, and `ResourceTypeList`.
	//
	// example:
	//
	// ["EcsPublicIp","NatEip"]
	ResourceTypeList []*string `json:"ResourceTypeList,omitempty" xml:"ResourceTypeList,omitempty" type:"Repeated"`
	// Deprecated
	//
	// The source IP address of the request.
	//
	// example:
	//
	// 192.0.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s PutDisableFwSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s PutDisableFwSwitchRequest) GoString() string {
	return s.String()
}

func (s *PutDisableFwSwitchRequest) GetIpVersion() *string {
	return s.IpVersion
}

func (s *PutDisableFwSwitchRequest) GetIpaddrList() []*string {
	return s.IpaddrList
}

func (s *PutDisableFwSwitchRequest) GetLang() *string {
	return s.Lang
}

func (s *PutDisableFwSwitchRequest) GetMemberUid() *string {
	return s.MemberUid
}

func (s *PutDisableFwSwitchRequest) GetRegionList() []*string {
	return s.RegionList
}

func (s *PutDisableFwSwitchRequest) GetResourceTypeList() []*string {
	return s.ResourceTypeList
}

func (s *PutDisableFwSwitchRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *PutDisableFwSwitchRequest) SetIpVersion(v string) *PutDisableFwSwitchRequest {
	s.IpVersion = &v
	return s
}

func (s *PutDisableFwSwitchRequest) SetIpaddrList(v []*string) *PutDisableFwSwitchRequest {
	s.IpaddrList = v
	return s
}

func (s *PutDisableFwSwitchRequest) SetLang(v string) *PutDisableFwSwitchRequest {
	s.Lang = &v
	return s
}

func (s *PutDisableFwSwitchRequest) SetMemberUid(v string) *PutDisableFwSwitchRequest {
	s.MemberUid = &v
	return s
}

func (s *PutDisableFwSwitchRequest) SetRegionList(v []*string) *PutDisableFwSwitchRequest {
	s.RegionList = v
	return s
}

func (s *PutDisableFwSwitchRequest) SetResourceTypeList(v []*string) *PutDisableFwSwitchRequest {
	s.ResourceTypeList = v
	return s
}

func (s *PutDisableFwSwitchRequest) SetSourceIp(v string) *PutDisableFwSwitchRequest {
	s.SourceIp = &v
	return s
}

func (s *PutDisableFwSwitchRequest) Validate() error {
	return dara.Validate(s)
}
