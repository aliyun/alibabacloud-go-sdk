// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGeneratedContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *UpdateGeneratedContentRequest
	GetAgentKey() *string
	SetContent(v string) *UpdateGeneratedContentRequest
	GetContent() *string
	SetContentText(v string) *UpdateGeneratedContentRequest
	GetContentText() *string
	SetId(v int64) *UpdateGeneratedContentRequest
	GetId() *int64
	SetKeywords(v []*string) *UpdateGeneratedContentRequest
	GetKeywords() []*string
	SetPrompt(v string) *UpdateGeneratedContentRequest
	GetPrompt() *string
	SetTitle(v string) *UpdateGeneratedContentRequest
	GetTitle() *string
}

type UpdateGeneratedContentRequest struct {
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
	Id       *int64    `json:"Id,omitempty" xml:"Id,omitempty"`
	Keywords []*string `json:"Keywords,omitempty" xml:"Keywords,omitempty" type:"Repeated"`
	// example:
	//
	// 创作xx文章
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// 评论类文章
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateGeneratedContentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGeneratedContentRequest) GoString() string {
	return s.String()
}

func (s *UpdateGeneratedContentRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *UpdateGeneratedContentRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateGeneratedContentRequest) GetContentText() *string {
	return s.ContentText
}

func (s *UpdateGeneratedContentRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateGeneratedContentRequest) GetKeywords() []*string {
	return s.Keywords
}

func (s *UpdateGeneratedContentRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *UpdateGeneratedContentRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateGeneratedContentRequest) SetAgentKey(v string) *UpdateGeneratedContentRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateGeneratedContentRequest) SetContent(v string) *UpdateGeneratedContentRequest {
	s.Content = &v
	return s
}

func (s *UpdateGeneratedContentRequest) SetContentText(v string) *UpdateGeneratedContentRequest {
	s.ContentText = &v
	return s
}

func (s *UpdateGeneratedContentRequest) SetId(v int64) *UpdateGeneratedContentRequest {
	s.Id = &v
	return s
}

func (s *UpdateGeneratedContentRequest) SetKeywords(v []*string) *UpdateGeneratedContentRequest {
	s.Keywords = v
	return s
}

func (s *UpdateGeneratedContentRequest) SetPrompt(v string) *UpdateGeneratedContentRequest {
	s.Prompt = &v
	return s
}

func (s *UpdateGeneratedContentRequest) SetTitle(v string) *UpdateGeneratedContentRequest {
	s.Title = &v
	return s
}

func (s *UpdateGeneratedContentRequest) Validate() error {
	return dara.Validate(s)
}
