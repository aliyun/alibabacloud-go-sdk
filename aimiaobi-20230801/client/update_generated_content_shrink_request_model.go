// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGeneratedContentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *UpdateGeneratedContentShrinkRequest
	GetAgentKey() *string
	SetContent(v string) *UpdateGeneratedContentShrinkRequest
	GetContent() *string
	SetContentText(v string) *UpdateGeneratedContentShrinkRequest
	GetContentText() *string
	SetId(v int64) *UpdateGeneratedContentShrinkRequest
	GetId() *int64
	SetKeywordsShrink(v string) *UpdateGeneratedContentShrinkRequest
	GetKeywordsShrink() *string
	SetPrompt(v string) *UpdateGeneratedContentShrinkRequest
	GetPrompt() *string
	SetTitle(v string) *UpdateGeneratedContentShrinkRequest
	GetTitle() *string
}

type UpdateGeneratedContentShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 正文
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 正文
	ContentText *string `json:"ContentText,omitempty" xml:"ContentText,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 36
	Id             *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	KeywordsShrink *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	// example:
	//
	// 创作xx文章
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// 评论类文章
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateGeneratedContentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGeneratedContentShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateGeneratedContentShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *UpdateGeneratedContentShrinkRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateGeneratedContentShrinkRequest) GetContentText() *string {
	return s.ContentText
}

func (s *UpdateGeneratedContentShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateGeneratedContentShrinkRequest) GetKeywordsShrink() *string {
	return s.KeywordsShrink
}

func (s *UpdateGeneratedContentShrinkRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *UpdateGeneratedContentShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateGeneratedContentShrinkRequest) SetAgentKey(v string) *UpdateGeneratedContentShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateGeneratedContentShrinkRequest) SetContent(v string) *UpdateGeneratedContentShrinkRequest {
	s.Content = &v
	return s
}

func (s *UpdateGeneratedContentShrinkRequest) SetContentText(v string) *UpdateGeneratedContentShrinkRequest {
	s.ContentText = &v
	return s
}

func (s *UpdateGeneratedContentShrinkRequest) SetId(v int64) *UpdateGeneratedContentShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateGeneratedContentShrinkRequest) SetKeywordsShrink(v string) *UpdateGeneratedContentShrinkRequest {
	s.KeywordsShrink = &v
	return s
}

func (s *UpdateGeneratedContentShrinkRequest) SetPrompt(v string) *UpdateGeneratedContentShrinkRequest {
	s.Prompt = &v
	return s
}

func (s *UpdateGeneratedContentShrinkRequest) SetTitle(v string) *UpdateGeneratedContentShrinkRequest {
	s.Title = &v
	return s
}

func (s *UpdateGeneratedContentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
