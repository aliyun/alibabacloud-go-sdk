// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliasName(v string) *UpdateInstanceRequest
	GetAliasName() *string
	SetInstanceDescription(v string) *UpdateInstanceRequest
	GetInstanceDescription() *string
	SetInstanceName(v string) *UpdateInstanceRequest
	GetInstanceName() *string
	SetNetwork(v string) *UpdateInstanceRequest
	GetNetwork() *string
	SetNetworkSourceACL(v []*string) *UpdateInstanceRequest
	GetNetworkSourceACL() []*string
	SetNetworkTypeACL(v []*string) *UpdateInstanceRequest
	GetNetworkTypeACL() []*string
}

type UpdateInstanceRequest struct {
	// The alias of the instance.
	//
	// example:
	//
	// instance-test
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// the test instance
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// The name of the instance whose information you want to update.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance-test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// (Deprecated) The network type of the instance. Valid values: NORMAL and VPC_CONSOLE. Default value: NORMAL.
	//
	// example:
	//
	// VPC
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The new sources of the network from which access is allowed. By default, all sources of networks are allowed. Valid value:
	//
	// TRUST_PROXY: the console
	NetworkSourceACL []*string `json:"NetworkSourceACL,omitempty" xml:"NetworkSourceACL,omitempty" type:"Repeated"`
	// The new types of the network from which access is allowed. By default, all types of networks are allowed. Valid values:
	//
	// 	- INTERNET: the Internet
	//
	// 	- VPC: VPCs
	//
	// 	- CLASSIC: the classic network
	NetworkTypeACL []*string `json:"NetworkTypeACL,omitempty" xml:"NetworkTypeACL,omitempty" type:"Repeated"`
}

func (s UpdateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequest) GetAliasName() *string {
	return s.AliasName
}

func (s *UpdateInstanceRequest) GetInstanceDescription() *string {
	return s.InstanceDescription
}

func (s *UpdateInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *UpdateInstanceRequest) GetNetwork() *string {
	return s.Network
}

func (s *UpdateInstanceRequest) GetNetworkSourceACL() []*string {
	return s.NetworkSourceACL
}

func (s *UpdateInstanceRequest) GetNetworkTypeACL() []*string {
	return s.NetworkTypeACL
}

func (s *UpdateInstanceRequest) SetAliasName(v string) *UpdateInstanceRequest {
	s.AliasName = &v
	return s
}

func (s *UpdateInstanceRequest) SetInstanceDescription(v string) *UpdateInstanceRequest {
	s.InstanceDescription = &v
	return s
}

func (s *UpdateInstanceRequest) SetInstanceName(v string) *UpdateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *UpdateInstanceRequest) SetNetwork(v string) *UpdateInstanceRequest {
	s.Network = &v
	return s
}

func (s *UpdateInstanceRequest) SetNetworkSourceACL(v []*string) *UpdateInstanceRequest {
	s.NetworkSourceACL = v
	return s
}

func (s *UpdateInstanceRequest) SetNetworkTypeACL(v []*string) *UpdateInstanceRequest {
	s.NetworkTypeACL = v
	return s
}

func (s *UpdateInstanceRequest) Validate() error {
	return dara.Validate(s)
}
