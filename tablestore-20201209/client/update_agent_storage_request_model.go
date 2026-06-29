// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentStorageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentStorageDescription(v string) *UpdateAgentStorageRequest
	GetAgentStorageDescription() *string
	SetAgentStorageName(v string) *UpdateAgentStorageRequest
	GetAgentStorageName() *string
	SetAliasName(v string) *UpdateAgentStorageRequest
	GetAliasName() *string
	SetNetwork(v string) *UpdateAgentStorageRequest
	GetNetwork() *string
	SetNetworkSourceACL(v []*string) *UpdateAgentStorageRequest
	GetNetworkSourceACL() []*string
	SetNetworkTypeACL(v []*string) *UpdateAgentStorageRequest
	GetNetworkTypeACL() []*string
}

type UpdateAgentStorageRequest struct {
	// agent storage description
	//
	// example:
	//
	// description for agent storage
	AgentStorageDescription *string `json:"AgentStorageDescription,omitempty" xml:"AgentStorageDescription,omitempty"`
	// agent storage name
	//
	// This parameter is required.
	//
	// example:
	//
	// agent-test
	AgentStorageName *string `json:"AgentStorageName,omitempty" xml:"AgentStorageName,omitempty"`
	// The alias of the agent storage.
	//
	// example:
	//
	// agent-test
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// (Deprecated) The network type of the agent storage. Valid values: NORMAL and VPC_CONSOLE. Default value: NORMAL.
	//
	// example:
	//
	// VPC
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The list of network sources allowed for the agent storage. All sources are allowed by default. Valid values:
	//
	// - TRUST_PROXY: console.
	NetworkSourceACL []*string `json:"NetworkSourceACL,omitempty" xml:"NetworkSourceACL,omitempty" type:"Repeated"`
	// The list of network types allowed for the agent storage. All types are allowed by default. Valid values:
	//
	// - CLASSIC: classic network.
	//
	// - INTERNET: public network.
	//
	// - VPC: VPC network.
	NetworkTypeACL []*string `json:"NetworkTypeACL,omitempty" xml:"NetworkTypeACL,omitempty" type:"Repeated"`
}

func (s UpdateAgentStorageRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentStorageRequest) GoString() string {
	return s.String()
}

func (s *UpdateAgentStorageRequest) GetAgentStorageDescription() *string {
	return s.AgentStorageDescription
}

func (s *UpdateAgentStorageRequest) GetAgentStorageName() *string {
	return s.AgentStorageName
}

func (s *UpdateAgentStorageRequest) GetAliasName() *string {
	return s.AliasName
}

func (s *UpdateAgentStorageRequest) GetNetwork() *string {
	return s.Network
}

func (s *UpdateAgentStorageRequest) GetNetworkSourceACL() []*string {
	return s.NetworkSourceACL
}

func (s *UpdateAgentStorageRequest) GetNetworkTypeACL() []*string {
	return s.NetworkTypeACL
}

func (s *UpdateAgentStorageRequest) SetAgentStorageDescription(v string) *UpdateAgentStorageRequest {
	s.AgentStorageDescription = &v
	return s
}

func (s *UpdateAgentStorageRequest) SetAgentStorageName(v string) *UpdateAgentStorageRequest {
	s.AgentStorageName = &v
	return s
}

func (s *UpdateAgentStorageRequest) SetAliasName(v string) *UpdateAgentStorageRequest {
	s.AliasName = &v
	return s
}

func (s *UpdateAgentStorageRequest) SetNetwork(v string) *UpdateAgentStorageRequest {
	s.Network = &v
	return s
}

func (s *UpdateAgentStorageRequest) SetNetworkSourceACL(v []*string) *UpdateAgentStorageRequest {
	s.NetworkSourceACL = v
	return s
}

func (s *UpdateAgentStorageRequest) SetNetworkTypeACL(v []*string) *UpdateAgentStorageRequest {
	s.NetworkTypeACL = v
	return s
}

func (s *UpdateAgentStorageRequest) Validate() error {
	return dara.Validate(s)
}
