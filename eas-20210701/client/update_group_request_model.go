// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabels(v map[string]*string) *UpdateGroupRequest
	GetLabels() map[string]*string
	SetNetwork(v *UpdateGroupRequestNetwork) *UpdateGroupRequest
	GetNetwork() *UpdateGroupRequestNetwork
	SetTrafficMode(v string) *UpdateGroupRequest
	GetTrafficMode() *string
}

type UpdateGroupRequest struct {
	// The user-defined labels.
	Labels map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The VPC configuration.
	Network *UpdateGroupRequestNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	// The traffic mode. Valid values:
	//
	// - auto: automatically assigns weights based on the proportion of instances.
	//
	// - customized: distributes traffic based on custom fixed weights.
	//
	// example:
	//
	// auto
	TrafficMode *string `json:"TrafficMode,omitempty" xml:"TrafficMode,omitempty"`
}

func (s UpdateGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupRequest) GetLabels() map[string]*string {
	return s.Labels
}

func (s *UpdateGroupRequest) GetNetwork() *UpdateGroupRequestNetwork {
	return s.Network
}

func (s *UpdateGroupRequest) GetTrafficMode() *string {
	return s.TrafficMode
}

func (s *UpdateGroupRequest) SetLabels(v map[string]*string) *UpdateGroupRequest {
	s.Labels = v
	return s
}

func (s *UpdateGroupRequest) SetNetwork(v *UpdateGroupRequestNetwork) *UpdateGroupRequest {
	s.Network = v
	return s
}

func (s *UpdateGroupRequest) SetTrafficMode(v string) *UpdateGroupRequest {
	s.TrafficMode = &v
	return s
}

func (s *UpdateGroupRequest) Validate() error {
	if s.Network != nil {
		if err := s.Network.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateGroupRequestNetwork struct {
	// The dedicated gateway ID.
	//
	// example:
	//
	// gw-rcgxxxxxxzytgq9zrj
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-uf6xxxxxxc3lysxabg72i
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp1xxxxxxjeqwbo3z2pux
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-bp1xxxxxx17e0qrykjr9b
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s UpdateGroupRequestNetwork) String() string {
	return dara.Prettify(s)
}

func (s UpdateGroupRequestNetwork) GoString() string {
	return s.String()
}

func (s *UpdateGroupRequestNetwork) GetGatewayId() *string {
	return s.GatewayId
}

func (s *UpdateGroupRequestNetwork) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *UpdateGroupRequestNetwork) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *UpdateGroupRequestNetwork) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateGroupRequestNetwork) SetGatewayId(v string) *UpdateGroupRequestNetwork {
	s.GatewayId = &v
	return s
}

func (s *UpdateGroupRequestNetwork) SetSecurityGroupId(v string) *UpdateGroupRequestNetwork {
	s.SecurityGroupId = &v
	return s
}

func (s *UpdateGroupRequestNetwork) SetVSwitchId(v string) *UpdateGroupRequestNetwork {
	s.VSwitchId = &v
	return s
}

func (s *UpdateGroupRequestNetwork) SetVpcId(v string) *UpdateGroupRequestNetwork {
	s.VpcId = &v
	return s
}

func (s *UpdateGroupRequestNetwork) Validate() error {
	return dara.Validate(s)
}
