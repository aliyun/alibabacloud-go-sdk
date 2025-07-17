// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePromptTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *CreatePromptTemplateRequest
	GetContent() *string
	SetName(v string) *CreatePromptTemplateRequest
	GetName() *string
}

type CreatePromptTemplateRequest struct {
	// This parameter is required.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreatePromptTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePromptTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreatePromptTemplateRequest) GetContent() *string {
	return s.Content
}

func (s *CreatePromptTemplateRequest) GetName() *string {
	return s.Name
}

func (s *CreatePromptTemplateRequest) SetContent(v string) *CreatePromptTemplateRequest {
	s.Content = &v
	return s
}

func (s *CreatePromptTemplateRequest) SetName(v string) *CreatePromptTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreatePromptTemplateRequest) Validate() error {
	return dara.Validate(s)
}
