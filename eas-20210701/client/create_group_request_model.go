// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabels(v map[string]*string) *CreateGroupRequest
	GetLabels() map[string]*string
	SetName(v string) *CreateGroupRequest
	GetName() *string
	SetNetwork(v *CreateGroupRequestNetwork) *CreateGroupRequest
	GetNetwork() *CreateGroupRequestNetwork
	SetWorkSpaceId(v string) *CreateGroupRequest
	GetWorkSpaceId() *string
}

type CreateGroupRequest struct {
	// The user-defined labels.
	Labels map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The name of the group.
	//
	// example:
	//
	// foo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The virtual private cloud (VPC) configuration.
	Network *CreateGroupRequestNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	// The workspace ID.
	//
	// example:
	//
	// 12xx34
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" xml:"WorkSpaceId,omitempty"`
}

func (s CreateGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupRequest) GetLabels() map[string]*string {
	return s.Labels
}

func (s *CreateGroupRequest) GetName() *string {
	return s.Name
}

func (s *CreateGroupRequest) GetNetwork() *CreateGroupRequestNetwork {
	return s.Network
}

func (s *CreateGroupRequest) GetWorkSpaceId() *string {
	return s.WorkSpaceId
}

func (s *CreateGroupRequest) SetLabels(v map[string]*string) *CreateGroupRequest {
	s.Labels = v
	return s
}

func (s *CreateGroupRequest) SetName(v string) *CreateGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateGroupRequest) SetNetwork(v *CreateGroupRequestNetwork) *CreateGroupRequest {
	s.Network = v
	return s
}

func (s *CreateGroupRequest) SetWorkSpaceId(v string) *CreateGroupRequest {
	s.WorkSpaceId = &v
	return s
}

func (s *CreateGroupRequest) Validate() error {
	if s.Network != nil {
		if err := s.Network.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateGroupRequestNetwork struct {
	// The dedicated gateway ID.
	//
	// example:
	//
	// gw-248xxxxxxvlkhtbrda
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-584xxxxxx7h08llvoww5tv5gl
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp1xxxxxxwmssgq28gye8
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-bp1xxxxxx0qrykjr9b
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateGroupRequestNetwork) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupRequestNetwork) GoString() string {
	return s.String()
}

func (s *CreateGroupRequestNetwork) GetGatewayId() *string {
	return s.GatewayId
}

func (s *CreateGroupRequestNetwork) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateGroupRequestNetwork) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateGroupRequestNetwork) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateGroupRequestNetwork) SetGatewayId(v string) *CreateGroupRequestNetwork {
	s.GatewayId = &v
	return s
}

func (s *CreateGroupRequestNetwork) SetSecurityGroupId(v string) *CreateGroupRequestNetwork {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateGroupRequestNetwork) SetVSwitchId(v string) *CreateGroupRequestNetwork {
	s.VSwitchId = &v
	return s
}

func (s *CreateGroupRequestNetwork) SetVpcId(v string) *CreateGroupRequestNetwork {
	s.VpcId = &v
	return s
}

func (s *CreateGroupRequestNetwork) Validate() error {
	return dara.Validate(s)
}
