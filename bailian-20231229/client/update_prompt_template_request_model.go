// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePromptTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *UpdatePromptTemplateRequest
	GetContent() *string
	SetName(v string) *UpdatePromptTemplateRequest
	GetName() *string
}

type UpdatePromptTemplateRequest struct {
	// The template content.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The template name.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdatePromptTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePromptTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdatePromptTemplateRequest) GetContent() *string {
	return s.Content
}

func (s *UpdatePromptTemplateRequest) GetName() *string {
	return s.Name
}

func (s *UpdatePromptTemplateRequest) SetContent(v string) *UpdatePromptTemplateRequest {
	s.Content = &v
	return s
}

func (s *UpdatePromptTemplateRequest) SetName(v string) *UpdatePromptTemplateRequest {
	s.Name = &v
	return s
}

func (s *UpdatePromptTemplateRequest) Validate() error {
	return dara.Validate(s)
}
