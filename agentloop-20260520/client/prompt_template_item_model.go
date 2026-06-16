// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPromptTemplateItem interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *PromptTemplateItem
	GetContent() *string
	SetRole(v string) *PromptTemplateItem
	GetRole() *string
}

type PromptTemplateItem struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	Role    *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s PromptTemplateItem) String() string {
	return dara.Prettify(s)
}

func (s PromptTemplateItem) GoString() string {
	return s.String()
}

func (s *PromptTemplateItem) GetContent() *string {
	return s.Content
}

func (s *PromptTemplateItem) GetRole() *string {
	return s.Role
}

func (s *PromptTemplateItem) SetContent(v string) *PromptTemplateItem {
	s.Content = &v
	return s
}

func (s *PromptTemplateItem) SetRole(v string) *PromptTemplateItem {
	s.Role = &v
	return s
}

func (s *PromptTemplateItem) Validate() error {
	return dara.Validate(s)
}
