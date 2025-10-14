// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkInterfaceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateNetworkInterfaceShrinkRequest
	GetDescription() *string
	SetName(v string) *CreateNetworkInterfaceShrinkRequest
	GetName() *string
	SetSecurityGroupIdsShrink(v string) *CreateNetworkInterfaceShrinkRequest
	GetSecurityGroupIdsShrink() *string
	SetVSwitchId(v string) *CreateNetworkInterfaceShrinkRequest
	GetVSwitchId() *string
}

type CreateNetworkInterfaceShrinkRequest struct {
	// Description of the ENI.
	//
	// example:
	//
	// example
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// ENI name.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Security group ID.
	//
	// This parameter is required.
	SecurityGroupIdsShrink *string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty"`
	// vSwitch ID.
	//
	// example:
	//
	// vsw-5****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateNetworkInterfaceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkInterfaceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateNetworkInterfaceShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateNetworkInterfaceShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateNetworkInterfaceShrinkRequest) GetSecurityGroupIdsShrink() *string {
	return s.SecurityGroupIdsShrink
}

func (s *CreateNetworkInterfaceShrinkRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateNetworkInterfaceShrinkRequest) SetDescription(v string) *CreateNetworkInterfaceShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateNetworkInterfaceShrinkRequest) SetName(v string) *CreateNetworkInterfaceShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateNetworkInterfaceShrinkRequest) SetSecurityGroupIdsShrink(v string) *CreateNetworkInterfaceShrinkRequest {
	s.SecurityGroupIdsShrink = &v
	return s
}

func (s *CreateNetworkInterfaceShrinkRequest) SetVSwitchId(v string) *CreateNetworkInterfaceShrinkRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateNetworkInterfaceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
