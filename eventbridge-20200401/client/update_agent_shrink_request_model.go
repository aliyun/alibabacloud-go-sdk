// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateAgentShrinkRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateAgentShrinkRequest
	GetDescription() *string
	SetMetadataShrink(v string) *UpdateAgentShrinkRequest
	GetMetadataShrink() *string
	SetName(v string) *UpdateAgentShrinkRequest
	GetName() *string
	SetPrompt(v string) *UpdateAgentShrinkRequest
	GetPrompt() *string
}

type UpdateAgentShrinkRequest struct {
	// example:
	//
	// TF-CreateRule-1652253755-aa33f762-7e99-4aee-bd27-d3370afa5625
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	MetadataShrink *string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// my-agent
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
}

func (s UpdateAgentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAgentShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateAgentShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateAgentShrinkRequest) GetMetadataShrink() *string {
	return s.MetadataShrink
}

func (s *UpdateAgentShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateAgentShrinkRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *UpdateAgentShrinkRequest) SetClientToken(v string) *UpdateAgentShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAgentShrinkRequest) SetDescription(v string) *UpdateAgentShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateAgentShrinkRequest) SetMetadataShrink(v string) *UpdateAgentShrinkRequest {
	s.MetadataShrink = &v
	return s
}

func (s *UpdateAgentShrinkRequest) SetName(v string) *UpdateAgentShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateAgentShrinkRequest) SetPrompt(v string) *UpdateAgentShrinkRequest {
	s.Prompt = &v
	return s
}

func (s *UpdateAgentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
