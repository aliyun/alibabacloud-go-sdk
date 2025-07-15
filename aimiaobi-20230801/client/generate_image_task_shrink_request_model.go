// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateImageTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *GenerateImageTaskShrinkRequest
	GetAgentKey() *string
	SetArticleTaskId(v string) *GenerateImageTaskShrinkRequest
	GetArticleTaskId() *string
	SetParagraphListShrink(v string) *GenerateImageTaskShrinkRequest
	GetParagraphListShrink() *string
	SetSize(v string) *GenerateImageTaskShrinkRequest
	GetSize() *string
	SetStyle(v string) *GenerateImageTaskShrinkRequest
	GetStyle() *string
}

type GenerateImageTaskShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// e1be065b-adc3-435e-bd01-1c18c5ed75d3
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e1be065b-adc3-435e-bd01-1c18c5ed75d3
	ArticleTaskId *string `json:"ArticleTaskId,omitempty" xml:"ArticleTaskId,omitempty"`
	// This parameter is required.
	ParagraphListShrink *string `json:"ParagraphList,omitempty" xml:"ParagraphList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1024*1024
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// <auto>
	Style *string `json:"Style,omitempty" xml:"Style,omitempty"`
}

func (s GenerateImageTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateImageTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *GenerateImageTaskShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *GenerateImageTaskShrinkRequest) GetArticleTaskId() *string {
	return s.ArticleTaskId
}

func (s *GenerateImageTaskShrinkRequest) GetParagraphListShrink() *string {
	return s.ParagraphListShrink
}

func (s *GenerateImageTaskShrinkRequest) GetSize() *string {
	return s.Size
}

func (s *GenerateImageTaskShrinkRequest) GetStyle() *string {
	return s.Style
}

func (s *GenerateImageTaskShrinkRequest) SetAgentKey(v string) *GenerateImageTaskShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *GenerateImageTaskShrinkRequest) SetArticleTaskId(v string) *GenerateImageTaskShrinkRequest {
	s.ArticleTaskId = &v
	return s
}

func (s *GenerateImageTaskShrinkRequest) SetParagraphListShrink(v string) *GenerateImageTaskShrinkRequest {
	s.ParagraphListShrink = &v
	return s
}

func (s *GenerateImageTaskShrinkRequest) SetSize(v string) *GenerateImageTaskShrinkRequest {
	s.Size = &v
	return s
}

func (s *GenerateImageTaskShrinkRequest) SetStyle(v string) *GenerateImageTaskShrinkRequest {
	s.Style = &v
	return s
}

func (s *GenerateImageTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
