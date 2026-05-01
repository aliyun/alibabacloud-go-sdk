// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentPlatform(v string) *CreateModelTemplateRequest
	GetAgentPlatform() *string
	SetAgentProvider(v string) *CreateModelTemplateRequest
	GetAgentProvider() *string
	SetBizType(v int32) *CreateModelTemplateRequest
	GetBizType() *int32
	SetDescription(v string) *CreateModelTemplateRequest
	GetDescription() *string
	SetName(v string) *CreateModelTemplateRequest
	GetName() *string
}

type CreateModelTemplateRequest struct {
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
	// 1
	BizType     *int32  `json:"BizType,omitempty" xml:"BizType,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// model-template-001
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateModelTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateModelTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateModelTemplateRequest) GetAgentPlatform() *string {
	return s.AgentPlatform
}

func (s *CreateModelTemplateRequest) GetAgentProvider() *string {
	return s.AgentProvider
}

func (s *CreateModelTemplateRequest) GetBizType() *int32 {
	return s.BizType
}

func (s *CreateModelTemplateRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateModelTemplateRequest) GetName() *string {
	return s.Name
}

func (s *CreateModelTemplateRequest) SetAgentPlatform(v string) *CreateModelTemplateRequest {
	s.AgentPlatform = &v
	return s
}

func (s *CreateModelTemplateRequest) SetAgentProvider(v string) *CreateModelTemplateRequest {
	s.AgentProvider = &v
	return s
}

func (s *CreateModelTemplateRequest) SetBizType(v int32) *CreateModelTemplateRequest {
	s.BizType = &v
	return s
}

func (s *CreateModelTemplateRequest) SetDescription(v string) *CreateModelTemplateRequest {
	s.Description = &v
	return s
}

func (s *CreateModelTemplateRequest) SetName(v string) *CreateModelTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateModelTemplateRequest) Validate() error {
	return dara.Validate(s)
}
