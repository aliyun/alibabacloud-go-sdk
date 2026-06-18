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
	// Prompt template content.
	//
	// This parameter is required.
	//
	// example:
	//
	// 请写一篇小红书种草笔记，增加丰富的emoji元素，结尾作总结，并加上相关标签。主题为：${theme}
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// Template name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 小红书文案助手
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
