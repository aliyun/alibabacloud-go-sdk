// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigRuntimeChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentPlatform(v string) *ConfigRuntimeChannelRequest
	GetAgentPlatform() *string
	SetAgentProvider(v string) *ConfigRuntimeChannelRequest
	GetAgentProvider() *string
	SetCode(v string) *ConfigRuntimeChannelRequest
	GetCode() *string
	SetConfig(v string) *ConfigRuntimeChannelRequest
	GetConfig() *string
	SetConfigMode(v string) *ConfigRuntimeChannelRequest
	GetConfigMode() *string
	SetName(v string) *ConfigRuntimeChannelRequest
	GetName() *string
	SetRuntimeIds(v []*string) *ConfigRuntimeChannelRequest
	GetRuntimeIds() []*string
	SetRuntimeType(v string) *ConfigRuntimeChannelRequest
	GetRuntimeType() *string
}

type ConfigRuntimeChannelRequest struct {
	// The Agent platform (such as ENTERPRISE or JVS).
	//
	// example:
	//
	// ENTERPRISE
	AgentPlatform *string `json:"AgentPlatform,omitempty" xml:"AgentPlatform,omitempty"`
	// The Agent provider.
	//
	// This parameter is required.
	//
	// example:
	//
	// OpenClaw
	AgentProvider *string `json:"AgentProvider,omitempty" xml:"AgentProvider,omitempty"`
	// The third-party channel code.
	//
	// This parameter is required.
	//
	// example:
	//
	// dingtalk-connector
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The channel configuration JSON string.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//     "appKey": "abc",
	//
	//     "appSecret": "efg"
	//
	// }
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The configuration mode.
	//
	// example:
	//
	// Simple
	ConfigMode *string `json:"ConfigMode,omitempty" xml:"ConfigMode,omitempty"`
	// The channel name.
	//
	// If you leave this parameter empty, the system automatically uses the value of Code as the name.
	//
	// example:
	//
	// 钉钉
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of runtime IDs.
	//
	// This parameter is required.
	RuntimeIds []*string `json:"RuntimeIds,omitempty" xml:"RuntimeIds,omitempty" type:"Repeated"`
	// The Agent runtime type.
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudDesktop
	RuntimeType *string `json:"RuntimeType,omitempty" xml:"RuntimeType,omitempty"`
}

func (s ConfigRuntimeChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigRuntimeChannelRequest) GoString() string {
	return s.String()
}

func (s *ConfigRuntimeChannelRequest) GetAgentPlatform() *string {
	return s.AgentPlatform
}

func (s *ConfigRuntimeChannelRequest) GetAgentProvider() *string {
	return s.AgentProvider
}

func (s *ConfigRuntimeChannelRequest) GetCode() *string {
	return s.Code
}

func (s *ConfigRuntimeChannelRequest) GetConfig() *string {
	return s.Config
}

func (s *ConfigRuntimeChannelRequest) GetConfigMode() *string {
	return s.ConfigMode
}

func (s *ConfigRuntimeChannelRequest) GetName() *string {
	return s.Name
}

func (s *ConfigRuntimeChannelRequest) GetRuntimeIds() []*string {
	return s.RuntimeIds
}

func (s *ConfigRuntimeChannelRequest) GetRuntimeType() *string {
	return s.RuntimeType
}

func (s *ConfigRuntimeChannelRequest) SetAgentPlatform(v string) *ConfigRuntimeChannelRequest {
	s.AgentPlatform = &v
	return s
}

func (s *ConfigRuntimeChannelRequest) SetAgentProvider(v string) *ConfigRuntimeChannelRequest {
	s.AgentProvider = &v
	return s
}

func (s *ConfigRuntimeChannelRequest) SetCode(v string) *ConfigRuntimeChannelRequest {
	s.Code = &v
	return s
}

func (s *ConfigRuntimeChannelRequest) SetConfig(v string) *ConfigRuntimeChannelRequest {
	s.Config = &v
	return s
}

func (s *ConfigRuntimeChannelRequest) SetConfigMode(v string) *ConfigRuntimeChannelRequest {
	s.ConfigMode = &v
	return s
}

func (s *ConfigRuntimeChannelRequest) SetName(v string) *ConfigRuntimeChannelRequest {
	s.Name = &v
	return s
}

func (s *ConfigRuntimeChannelRequest) SetRuntimeIds(v []*string) *ConfigRuntimeChannelRequest {
	s.RuntimeIds = v
	return s
}

func (s *ConfigRuntimeChannelRequest) SetRuntimeType(v string) *ConfigRuntimeChannelRequest {
	s.RuntimeType = &v
	return s
}

func (s *ConfigRuntimeChannelRequest) Validate() error {
	return dara.Validate(s)
}
