// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateAgentShrinkRequest
	GetDescription() *string
	SetMetadataShrink(v string) *CreateAgentShrinkRequest
	GetMetadataShrink() *string
	SetName(v string) *CreateAgentShrinkRequest
	GetName() *string
	SetPrompt(v string) *CreateAgentShrinkRequest
	GetPrompt() *string
}

type CreateAgentShrinkRequest struct {
	// The description of the event bus.
	//
	// example:
	//
	// 连接配置描述信息
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The associated metadata.
	MetadataShrink *string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The name of the agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-agent
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// You are an IoT data analytics assistant...
	//
	// example:
	//
	// 我想要她，你这样增加请求头获取用户IP CF-Connecting-IP%3B
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
}

func (s CreateAgentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAgentShrinkRequest) GetMetadataShrink() *string {
	return s.MetadataShrink
}

func (s *CreateAgentShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateAgentShrinkRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *CreateAgentShrinkRequest) SetDescription(v string) *CreateAgentShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateAgentShrinkRequest) SetMetadataShrink(v string) *CreateAgentShrinkRequest {
	s.MetadataShrink = &v
	return s
}

func (s *CreateAgentShrinkRequest) SetName(v string) *CreateAgentShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateAgentShrinkRequest) SetPrompt(v string) *CreateAgentShrinkRequest {
	s.Prompt = &v
	return s
}

func (s *CreateAgentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
