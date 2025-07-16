// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceDLinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationCIDRs(v string) *UpdateResourceDLinkRequest
	GetDestinationCIDRs() *string
	SetSecurityGroupId(v string) *UpdateResourceDLinkRequest
	GetSecurityGroupId() *string
	SetVSwitchId(v string) *UpdateResourceDLinkRequest
	GetVSwitchId() *string
	SetVSwitchIdList(v []*string) *UpdateResourceDLinkRequest
	GetVSwitchIdList() []*string
}

type UpdateResourceDLinkRequest struct {
	// The CIDR blocks of the clients that you want to connect to. After this parameter is specified, the CIDR blocks are added to the back-to-origin route of the server. Either this parameter or the VSwitchIdList parameter can be used to determine CIDR blocks.
	//
	// example:
	//
	// 72.16.0.0/16
	DestinationCIDRs *string `json:"DestinationCIDRs,omitempty" xml:"DestinationCIDRs,omitempty"`
	// The ID of the security group to which the Elastic Compute Service (ECS) instance belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// sg-bp149cedsfx2rfspd2d
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The ID of the peer primary vSwitch. After this parameter is specified, an elastic network interface (ENI) is created in the VSwitch.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-uf66uio7md****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The vSwitches of the clients that you want to connect to. After this parameter is specified, the CIDR blocks of these vSwitches are added to the back-to-origin route of the server.
	VSwitchIdList []*string `json:"VSwitchIdList,omitempty" xml:"VSwitchIdList,omitempty" type:"Repeated"`
}

func (s UpdateResourceDLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceDLinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceDLinkRequest) GetDestinationCIDRs() *string {
	return s.DestinationCIDRs
}

func (s *UpdateResourceDLinkRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *UpdateResourceDLinkRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *UpdateResourceDLinkRequest) GetVSwitchIdList() []*string {
	return s.VSwitchIdList
}

func (s *UpdateResourceDLinkRequest) SetDestinationCIDRs(v string) *UpdateResourceDLinkRequest {
	s.DestinationCIDRs = &v
	return s
}

func (s *UpdateResourceDLinkRequest) SetSecurityGroupId(v string) *UpdateResourceDLinkRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *UpdateResourceDLinkRequest) SetVSwitchId(v string) *UpdateResourceDLinkRequest {
	s.VSwitchId = &v
	return s
}

func (s *UpdateResourceDLinkRequest) SetVSwitchIdList(v []*string) *UpdateResourceDLinkRequest {
	s.VSwitchIdList = v
	return s
}

func (s *UpdateResourceDLinkRequest) Validate() error {
	return dara.Validate(s)
}
