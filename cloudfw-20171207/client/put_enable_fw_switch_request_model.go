// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutEnableFwSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpaddrList(v []*string) *PutEnableFwSwitchRequest
	GetIpaddrList() []*string
	SetLang(v string) *PutEnableFwSwitchRequest
	GetLang() *string
	SetRegionList(v []*string) *PutEnableFwSwitchRequest
	GetRegionList() []*string
	SetResourceTypeList(v []*string) *PutEnableFwSwitchRequest
	GetResourceTypeList() []*string
	SetSourceIp(v string) *PutEnableFwSwitchRequest
	GetSourceIp() *string
}

type PutEnableFwSwitchRequest struct {
	// The IP addresses.
	//
	// > You must specify at least one of the IpaddrList, RegionList, and ResourceTypeList parameters.
	//
	// example:
	//
	// ["192.0.X.X","192.0.X.X"]
	IpaddrList []*string `json:"IpaddrList,omitempty" xml:"IpaddrList,omitempty" type:"Repeated"`
	// The language of the content within the response.
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The regions.
	//
	// > You must specify at least one of the IpaddrList, RegionList, and ResourceTypeList parameters.
	//
	// example:
	//
	// ["cn-hangzhou","cn-shanghai"]
	RegionList []*string `json:"RegionList,omitempty" xml:"RegionList,omitempty" type:"Repeated"`
	// The types of the assets.
	//
	// Valid values:
	//
	// 	- BastionHostIP: the egress IP address of a bastion host
	//
	// 	- BastionHostIngressIP: the ingress IP address of a bastion host
	//
	// 	- EcsEIP: the elastic IP address (EIP) of an Elastic Compute Service (ECS) instance
	//
	// 	- EcsPublicIP: the public IP address of an ECS instance
	//
	// 	- EIP: the EIP
	//
	// 	- EniEIP: the EIP of an elastic network interface (ENI)
	//
	// 	- NatEIP: the EIP of a NAT gateway
	//
	// 	- SlbEIP: the EIP of a Server Load Balancer (SLB) instance
	//
	// 	- SlbPublicIP: the public IP address of an SLB instance
	//
	// 	- NatPublicIP: the public IP address of a NAT gateway
	//
	// 	- HAVIP: the high-availability virtual IP address (HAVIP)
	//
	// > You must specify at least one of the IpaddrList, RegionList, and ResourceTypeList parameters.
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
	// 192.0.X.X
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s PutEnableFwSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s PutEnableFwSwitchRequest) GoString() string {
	return s.String()
}

func (s *PutEnableFwSwitchRequest) GetIpaddrList() []*string {
	return s.IpaddrList
}

func (s *PutEnableFwSwitchRequest) GetLang() *string {
	return s.Lang
}

func (s *PutEnableFwSwitchRequest) GetRegionList() []*string {
	return s.RegionList
}

func (s *PutEnableFwSwitchRequest) GetResourceTypeList() []*string {
	return s.ResourceTypeList
}

func (s *PutEnableFwSwitchRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *PutEnableFwSwitchRequest) SetIpaddrList(v []*string) *PutEnableFwSwitchRequest {
	s.IpaddrList = v
	return s
}

func (s *PutEnableFwSwitchRequest) SetLang(v string) *PutEnableFwSwitchRequest {
	s.Lang = &v
	return s
}

func (s *PutEnableFwSwitchRequest) SetRegionList(v []*string) *PutEnableFwSwitchRequest {
	s.RegionList = v
	return s
}

func (s *PutEnableFwSwitchRequest) SetResourceTypeList(v []*string) *PutEnableFwSwitchRequest {
	s.ResourceTypeList = v
	return s
}

func (s *PutEnableFwSwitchRequest) SetSourceIp(v string) *PutEnableFwSwitchRequest {
	s.SourceIp = &v
	return s
}

func (s *PutEnableFwSwitchRequest) Validate() error {
	return dara.Validate(s)
}
