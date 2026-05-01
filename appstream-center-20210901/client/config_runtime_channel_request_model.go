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
	SetName(v string) *ConfigRuntimeChannelRequest
	GetName() *string
	SetRuntimeIds(v []*string) *ConfigRuntimeChannelRequest
	GetRuntimeIds() []*string
	SetRuntimeType(v string) *ConfigRuntimeChannelRequest
	GetRuntimeType() *string
}

type ConfigRuntimeChannelRequest struct {
	// example:
	//
	// ENTERPRISE
	AgentPlatform *string `json:"AgentPlatform,omitempty" xml:"AgentPlatform,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OpenClaw
	AgentProvider *string `json:"AgentProvider,omitempty" xml:"AgentProvider,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dingtalk-connector
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
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
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	RuntimeIds []*string `json:"RuntimeIds,omitempty" xml:"RuntimeIds,omitempty" type:"Repeated"`
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
