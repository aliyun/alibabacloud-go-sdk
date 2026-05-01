// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveRuntimeChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentPlatform(v string) *RemoveRuntimeChannelRequest
	GetAgentPlatform() *string
	SetAgentProvider(v string) *RemoveRuntimeChannelRequest
	GetAgentProvider() *string
	SetCode(v string) *RemoveRuntimeChannelRequest
	GetCode() *string
	SetRuntimeIds(v []*string) *RemoveRuntimeChannelRequest
	GetRuntimeIds() []*string
	SetRuntimeType(v string) *RemoveRuntimeChannelRequest
	GetRuntimeType() *string
}

type RemoveRuntimeChannelRequest struct {
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
	RuntimeIds []*string `json:"RuntimeIds,omitempty" xml:"RuntimeIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// CloudDesktop
	RuntimeType *string `json:"RuntimeType,omitempty" xml:"RuntimeType,omitempty"`
}

func (s RemoveRuntimeChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveRuntimeChannelRequest) GoString() string {
	return s.String()
}

func (s *RemoveRuntimeChannelRequest) GetAgentPlatform() *string {
	return s.AgentPlatform
}

func (s *RemoveRuntimeChannelRequest) GetAgentProvider() *string {
	return s.AgentProvider
}

func (s *RemoveRuntimeChannelRequest) GetCode() *string {
	return s.Code
}

func (s *RemoveRuntimeChannelRequest) GetRuntimeIds() []*string {
	return s.RuntimeIds
}

func (s *RemoveRuntimeChannelRequest) GetRuntimeType() *string {
	return s.RuntimeType
}

func (s *RemoveRuntimeChannelRequest) SetAgentPlatform(v string) *RemoveRuntimeChannelRequest {
	s.AgentPlatform = &v
	return s
}

func (s *RemoveRuntimeChannelRequest) SetAgentProvider(v string) *RemoveRuntimeChannelRequest {
	s.AgentProvider = &v
	return s
}

func (s *RemoveRuntimeChannelRequest) SetCode(v string) *RemoveRuntimeChannelRequest {
	s.Code = &v
	return s
}

func (s *RemoveRuntimeChannelRequest) SetRuntimeIds(v []*string) *RemoveRuntimeChannelRequest {
	s.RuntimeIds = v
	return s
}

func (s *RemoveRuntimeChannelRequest) SetRuntimeType(v string) *RemoveRuntimeChannelRequest {
	s.RuntimeType = &v
	return s
}

func (s *RemoveRuntimeChannelRequest) Validate() error {
	return dara.Validate(s)
}
