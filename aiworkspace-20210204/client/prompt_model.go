// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrompt interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *Prompt
	GetAccessibility() *string
	SetCreateTime(v string) *Prompt
	GetCreateTime() *string
	SetDescription(v string) *Prompt
	GetDescription() *string
	SetFrameworkContent(v string) *Prompt
	GetFrameworkContent() *string
	SetFrameworkType(v string) *Prompt
	GetFrameworkType() *string
	SetModifyTime(v string) *Prompt
	GetModifyTime() *string
	SetPromptId(v string) *Prompt
	GetPromptId() *string
	SetPromptName(v string) *Prompt
	GetPromptName() *string
}

type Prompt struct {
	Accessibility    *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	CreateTime       *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FrameworkContent *string `json:"FrameworkContent,omitempty" xml:"FrameworkContent,omitempty"`
	FrameworkType    *string `json:"FrameworkType,omitempty" xml:"FrameworkType,omitempty"`
	ModifyTime       *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	PromptId         *string `json:"PromptId,omitempty" xml:"PromptId,omitempty"`
	PromptName       *string `json:"PromptName,omitempty" xml:"PromptName,omitempty"`
}

func (s Prompt) String() string {
	return dara.Prettify(s)
}

func (s Prompt) GoString() string {
	return s.String()
}

func (s *Prompt) GetAccessibility() *string {
	return s.Accessibility
}

func (s *Prompt) GetCreateTime() *string {
	return s.CreateTime
}

func (s *Prompt) GetDescription() *string {
	return s.Description
}

func (s *Prompt) GetFrameworkContent() *string {
	return s.FrameworkContent
}

func (s *Prompt) GetFrameworkType() *string {
	return s.FrameworkType
}

func (s *Prompt) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *Prompt) GetPromptId() *string {
	return s.PromptId
}

func (s *Prompt) GetPromptName() *string {
	return s.PromptName
}

func (s *Prompt) SetAccessibility(v string) *Prompt {
	s.Accessibility = &v
	return s
}

func (s *Prompt) SetCreateTime(v string) *Prompt {
	s.CreateTime = &v
	return s
}

func (s *Prompt) SetDescription(v string) *Prompt {
	s.Description = &v
	return s
}

func (s *Prompt) SetFrameworkContent(v string) *Prompt {
	s.FrameworkContent = &v
	return s
}

func (s *Prompt) SetFrameworkType(v string) *Prompt {
	s.FrameworkType = &v
	return s
}

func (s *Prompt) SetModifyTime(v string) *Prompt {
	s.ModifyTime = &v
	return s
}

func (s *Prompt) SetPromptId(v string) *Prompt {
	s.PromptId = &v
	return s
}

func (s *Prompt) SetPromptName(v string) *Prompt {
	s.PromptName = &v
	return s
}

func (s *Prompt) Validate() error {
	return dara.Validate(s)
}
