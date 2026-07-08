// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAuditTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArticleId(v string) *SubmitAuditTaskRequest
	GetArticleId() *string
	SetContent(v string) *SubmitAuditTaskRequest
	GetContent() *string
	SetHtmlContent(v string) *SubmitAuditTaskRequest
	GetHtmlContent() *string
	SetTitle(v string) *SubmitAuditTaskRequest
	GetTitle() *string
	SetWorkspaceId(v string) *SubmitAuditTaskRequest
	GetWorkspaceId() *string
}

type SubmitAuditTaskRequest struct {
	// The ID of the article to be audited.
	//
	// example:
	//
	// xxxx
	ArticleId *string `json:"ArticleId,omitempty" xml:"ArticleId,omitempty"`
	// The content to be audited.
	//
	// example:
	//
	// 待审核的内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The content to be audited, in HTML format.
	//
	// example:
	//
	// 待审核的内容（HTML格式）
	HtmlContent *string `json:"HtmlContent,omitempty" xml:"HtmlContent,omitempty"`
	// The title of the article to be audited.
	//
	// example:
	//
	// 审核时的文章标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// [Workspace ID](https://help.aliyun.com/document_detail/2782167.html)
	//
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SubmitAuditTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitAuditTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitAuditTaskRequest) GetArticleId() *string {
	return s.ArticleId
}

func (s *SubmitAuditTaskRequest) GetContent() *string {
	return s.Content
}

func (s *SubmitAuditTaskRequest) GetHtmlContent() *string {
	return s.HtmlContent
}

func (s *SubmitAuditTaskRequest) GetTitle() *string {
	return s.Title
}

func (s *SubmitAuditTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SubmitAuditTaskRequest) SetArticleId(v string) *SubmitAuditTaskRequest {
	s.ArticleId = &v
	return s
}

func (s *SubmitAuditTaskRequest) SetContent(v string) *SubmitAuditTaskRequest {
	s.Content = &v
	return s
}

func (s *SubmitAuditTaskRequest) SetHtmlContent(v string) *SubmitAuditTaskRequest {
	s.HtmlContent = &v
	return s
}

func (s *SubmitAuditTaskRequest) SetTitle(v string) *SubmitAuditTaskRequest {
	s.Title = &v
	return s
}

func (s *SubmitAuditTaskRequest) SetWorkspaceId(v string) *SubmitAuditTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SubmitAuditTaskRequest) Validate() error {
	return dara.Validate(s)
}
