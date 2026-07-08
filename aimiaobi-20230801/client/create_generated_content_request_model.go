// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGeneratedContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *CreateGeneratedContentRequest
	GetAgentKey() *string
	SetContent(v string) *CreateGeneratedContentRequest
	GetContent() *string
	SetContentDomain(v string) *CreateGeneratedContentRequest
	GetContentDomain() *string
	SetContentText(v string) *CreateGeneratedContentRequest
	GetContentText() *string
	SetKeywords(v []*string) *CreateGeneratedContentRequest
	GetKeywords() []*string
	SetPrompt(v string) *CreateGeneratedContentRequest
	GetPrompt() *string
	SetTaskId(v string) *CreateGeneratedContentRequest
	GetTaskId() *string
	SetTitle(v string) *CreateGeneratedContentRequest
	GetTitle() *string
	SetUuid(v string) *CreateGeneratedContentRequest
	GetUuid() *string
}

type CreateGeneratedContentRequest struct {
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
	Keywords []*string `json:"Keywords,omitempty" xml:"Keywords,omitempty" type:"Repeated"`
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

func (s CreateGeneratedContentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGeneratedContentRequest) GoString() string {
	return s.String()
}

func (s *CreateGeneratedContentRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *CreateGeneratedContentRequest) GetContent() *string {
	return s.Content
}

func (s *CreateGeneratedContentRequest) GetContentDomain() *string {
	return s.ContentDomain
}

func (s *CreateGeneratedContentRequest) GetContentText() *string {
	return s.ContentText
}

func (s *CreateGeneratedContentRequest) GetKeywords() []*string {
	return s.Keywords
}

func (s *CreateGeneratedContentRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *CreateGeneratedContentRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateGeneratedContentRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateGeneratedContentRequest) GetUuid() *string {
	return s.Uuid
}

func (s *CreateGeneratedContentRequest) SetAgentKey(v string) *CreateGeneratedContentRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateGeneratedContentRequest) SetContent(v string) *CreateGeneratedContentRequest {
	s.Content = &v
	return s
}

func (s *CreateGeneratedContentRequest) SetContentDomain(v string) *CreateGeneratedContentRequest {
	s.ContentDomain = &v
	return s
}

func (s *CreateGeneratedContentRequest) SetContentText(v string) *CreateGeneratedContentRequest {
	s.ContentText = &v
	return s
}

func (s *CreateGeneratedContentRequest) SetKeywords(v []*string) *CreateGeneratedContentRequest {
	s.Keywords = v
	return s
}

func (s *CreateGeneratedContentRequest) SetPrompt(v string) *CreateGeneratedContentRequest {
	s.Prompt = &v
	return s
}

func (s *CreateGeneratedContentRequest) SetTaskId(v string) *CreateGeneratedContentRequest {
	s.TaskId = &v
	return s
}

func (s *CreateGeneratedContentRequest) SetTitle(v string) *CreateGeneratedContentRequest {
	s.Title = &v
	return s
}

func (s *CreateGeneratedContentRequest) SetUuid(v string) *CreateGeneratedContentRequest {
	s.Uuid = &v
	return s
}

func (s *CreateGeneratedContentRequest) Validate() error {
	return dara.Validate(s)
}
