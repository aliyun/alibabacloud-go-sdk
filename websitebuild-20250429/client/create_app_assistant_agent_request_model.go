// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppAssistantAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *CreateAppAssistantAgentRequest
	GetAgentName() *string
	SetBizId(v string) *CreateAppAssistantAgentRequest
	GetBizId() *string
	SetPlatformType(v string) *CreateAppAssistantAgentRequest
	GetPlatformType() *string
}

type CreateAppAssistantAgentRequest struct {
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// example:
	//
	// WS20250731233102000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// LINUX64
	PlatformType *string `json:"PlatformType,omitempty" xml:"PlatformType,omitempty"`
}

func (s CreateAppAssistantAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppAssistantAgentRequest) GoString() string {
	return s.String()
}

func (s *CreateAppAssistantAgentRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *CreateAppAssistantAgentRequest) GetBizId() *string {
	return s.BizId
}

func (s *CreateAppAssistantAgentRequest) GetPlatformType() *string {
	return s.PlatformType
}

func (s *CreateAppAssistantAgentRequest) SetAgentName(v string) *CreateAppAssistantAgentRequest {
	s.AgentName = &v
	return s
}

func (s *CreateAppAssistantAgentRequest) SetBizId(v string) *CreateAppAssistantAgentRequest {
	s.BizId = &v
	return s
}

func (s *CreateAppAssistantAgentRequest) SetPlatformType(v string) *CreateAppAssistantAgentRequest {
	s.PlatformType = &v
	return s
}

func (s *CreateAppAssistantAgentRequest) Validate() error {
	return dara.Validate(s)
}
