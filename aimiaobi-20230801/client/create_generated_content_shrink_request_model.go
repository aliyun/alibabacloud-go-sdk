// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGeneratedContentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *CreateGeneratedContentShrinkRequest
	GetAgentKey() *string
	SetContent(v string) *CreateGeneratedContentShrinkRequest
	GetContent() *string
	SetContentDomain(v string) *CreateGeneratedContentShrinkRequest
	GetContentDomain() *string
	SetContentText(v string) *CreateGeneratedContentShrinkRequest
	GetContentText() *string
	SetKeywordsShrink(v string) *CreateGeneratedContentShrinkRequest
	GetKeywordsShrink() *string
	SetPrompt(v string) *CreateGeneratedContentShrinkRequest
	GetPrompt() *string
	SetTaskId(v string) *CreateGeneratedContentShrinkRequest
	GetTaskId() *string
	SetTitle(v string) *CreateGeneratedContentShrinkRequest
	GetTitle() *string
	SetUuid(v string) *CreateGeneratedContentShrinkRequest
	GetUuid() *string
}

type CreateGeneratedContentShrinkRequest struct {
	// The unique identifier of the workspace. For more information, see [AgentKey](https://help.aliyun.com/document_detail/2587494.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The body of the content, in rich text format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 杭州亚运会
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The domain for content generation.
	//
	// example:
	//
	// government
	ContentDomain *string `json:"ContentDomain,omitempty" xml:"ContentDomain,omitempty"`
	// The body of the content, in plain text format.
	//
	// example:
	//
	// 杭州亚运会
	ContentText *string `json:"ContentText,omitempty" xml:"ContentText,omitempty"`
	// The keywords.
	KeywordsShrink *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	// The last generated prompt.
	//
	// example:
	//
	// 创作xxx文章
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// The unique identifier of the task.
	//
	// > The system automatically generates a task ID. You do not need to specify this parameter. If you specify the same task ID for multiple tasks, they are grouped into a single conversation.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The title.
	//
	// This parameter is required.
	//
	// example:
	//
	// 杭州亚运会
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The traceability UUID.
	//
	// example:
	//
	// xxxx
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s CreateGeneratedContentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGeneratedContentShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateGeneratedContentShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *CreateGeneratedContentShrinkRequest) GetContent() *string {
	return s.Content
}

func (s *CreateGeneratedContentShrinkRequest) GetContentDomain() *string {
	return s.ContentDomain
}

func (s *CreateGeneratedContentShrinkRequest) GetContentText() *string {
	return s.ContentText
}

func (s *CreateGeneratedContentShrinkRequest) GetKeywordsShrink() *string {
	return s.KeywordsShrink
}

func (s *CreateGeneratedContentShrinkRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *CreateGeneratedContentShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateGeneratedContentShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateGeneratedContentShrinkRequest) GetUuid() *string {
	return s.Uuid
}

func (s *CreateGeneratedContentShrinkRequest) SetAgentKey(v string) *CreateGeneratedContentShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateGeneratedContentShrinkRequest) SetContent(v string) *CreateGeneratedContentShrinkRequest {
	s.Content = &v
	return s
}

func (s *CreateGeneratedContentShrinkRequest) SetContentDomain(v string) *CreateGeneratedContentShrinkRequest {
	s.ContentDomain = &v
	return s
}

func (s *CreateGeneratedContentShrinkRequest) SetContentText(v string) *CreateGeneratedContentShrinkRequest {
	s.ContentText = &v
	return s
}

func (s *CreateGeneratedContentShrinkRequest) SetKeywordsShrink(v string) *CreateGeneratedContentShrinkRequest {
	s.KeywordsShrink = &v
	return s
}

func (s *CreateGeneratedContentShrinkRequest) SetPrompt(v string) *CreateGeneratedContentShrinkRequest {
	s.Prompt = &v
	return s
}

func (s *CreateGeneratedContentShrinkRequest) SetTaskId(v string) *CreateGeneratedContentShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *CreateGeneratedContentShrinkRequest) SetTitle(v string) *CreateGeneratedContentShrinkRequest {
	s.Title = &v
	return s
}

func (s *CreateGeneratedContentShrinkRequest) SetUuid(v string) *CreateGeneratedContentShrinkRequest {
	s.Uuid = &v
	return s
}

func (s *CreateGeneratedContentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
