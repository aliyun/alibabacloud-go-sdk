// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAtiAgentRegisterInfoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentDescription(v string) *CreateAtiAgentRegisterInfoShrinkRequest
	GetAgentDescription() *string
	SetAgentDisplayName(v string) *CreateAtiAgentRegisterInfoShrinkRequest
	GetAgentDisplayName() *string
	SetAgentHost(v string) *CreateAtiAgentRegisterInfoShrinkRequest
	GetAgentHost() *string
	SetAgentVersion(v string) *CreateAtiAgentRegisterInfoShrinkRequest
	GetAgentVersion() *string
	SetClientToken(v string) *CreateAtiAgentRegisterInfoShrinkRequest
	GetClientToken() *string
	SetEndpointsShrink(v string) *CreateAtiAgentRegisterInfoShrinkRequest
	GetEndpointsShrink() *string
	SetRegistrantId(v string) *CreateAtiAgentRegisterInfoShrinkRequest
	GetRegistrantId() *string
}

type CreateAtiAgentRegisterInfoShrinkRequest struct {
	// The description of the agent capabilities.
	//
	// example:
	//
	// 支付服务
	AgentDescription *string `json:"AgentDescription,omitempty" xml:"AgentDescription,omitempty"`
	// The display name of the agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// 测试Agent
	AgentDisplayName *string `json:"AgentDisplayName,omitempty" xml:"AgentDisplayName,omitempty"`
	// The endpoint domain name through which the agent provides services.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	AgentHost *string `json:"AgentHost,omitempty" xml:"AgentHost,omitempty"`
	// The version of the agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0.1
	AgentVersion *string `json:"AgentVersion,omitempty" xml:"AgentVersion,omitempty"`
	// Provides idempotency. Within 3 minutes, the same value takes effect only once.
	//
	// example:
	//
	// eyJhbGciOiJIUzI1NiIsInR5cC.....
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The endpoint information of the agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{\\"EndpointValue\\":\\"http://www.baidu.com\\",\\"EndpointType\\":\\"http\\"}]
	EndpointsShrink *string `json:"Endpoints,omitempty" xml:"Endpoints,omitempty"`
	// The ID of the verified registrant. Obtain this ID by invoking the identity verification API operation or from the ATS console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2072277378616354816
	RegistrantId *string `json:"RegistrantId,omitempty" xml:"RegistrantId,omitempty"`
}

func (s CreateAtiAgentRegisterInfoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAtiAgentRegisterInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAtiAgentRegisterInfoShrinkRequest) GetAgentDescription() *string {
	return s.AgentDescription
}

func (s *CreateAtiAgentRegisterInfoShrinkRequest) GetAgentDisplayName() *string {
	return s.AgentDisplayName
}

func (s *CreateAtiAgentRegisterInfoShrinkRequest) GetAgentHost() *string {
	return s.AgentHost
}

func (s *CreateAtiAgentRegisterInfoShrinkRequest) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *CreateAtiAgentRegisterInfoShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAtiAgentRegisterInfoShrinkRequest) GetEndpointsShrink() *string {
	return s.EndpointsShrink
}

func (s *CreateAtiAgentRegisterInfoShrinkRequest) GetRegistrantId() *string {
	return s.RegistrantId
}

func (s *CreateAtiAgentRegisterInfoShrinkRequest) SetAgentDescription(v string) *CreateAtiAgentRegisterInfoShrinkRequest {
	s.AgentDescription = &v
	return s
}

func (s *CreateAtiAgentRegisterInfoShrinkRequest) SetAgentDisplayName(v string) *CreateAtiAgentRegisterInfoShrinkRequest {
	s.AgentDisplayName = &v
	return s
}

func (s *CreateAtiAgentRegisterInfoShrinkRequest) SetAgentHost(v string) *CreateAtiAgentRegisterInfoShrinkRequest {
	s.AgentHost = &v
	return s
}

func (s *CreateAtiAgentRegisterInfoShrinkRequest) SetAgentVersion(v string) *CreateAtiAgentRegisterInfoShrinkRequest {
	s.AgentVersion = &v
	return s
}

func (s *CreateAtiAgentRegisterInfoShrinkRequest) SetClientToken(v string) *CreateAtiAgentRegisterInfoShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAtiAgentRegisterInfoShrinkRequest) SetEndpointsShrink(v string) *CreateAtiAgentRegisterInfoShrinkRequest {
	s.EndpointsShrink = &v
	return s
}

func (s *CreateAtiAgentRegisterInfoShrinkRequest) SetRegistrantId(v string) *CreateAtiAgentRegisterInfoShrinkRequest {
	s.RegistrantId = &v
	return s
}

func (s *CreateAtiAgentRegisterInfoShrinkRequest) Validate() error {
	return dara.Validate(s)
}
