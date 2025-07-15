// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAuditTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArticleId(v string) *QueryAuditTaskRequest
	GetArticleId() *string
	SetContentAuditTaskId(v string) *QueryAuditTaskRequest
	GetContentAuditTaskId() *string
	SetWorkspaceId(v string) *QueryAuditTaskRequest
	GetWorkspaceId() *string
}

type QueryAuditTaskRequest struct {
	// example:
	//
	// xxxx
	ArticleId *string `json:"ArticleId,omitempty" xml:"ArticleId,omitempty"`
	// example:
	//
	// xxx
	ContentAuditTaskId *string `json:"ContentAuditTaskId,omitempty" xml:"ContentAuditTaskId,omitempty"`
	// example:
	//
	// llm-xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryAuditTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAuditTaskRequest) GoString() string {
	return s.String()
}

func (s *QueryAuditTaskRequest) GetArticleId() *string {
	return s.ArticleId
}

func (s *QueryAuditTaskRequest) GetContentAuditTaskId() *string {
	return s.ContentAuditTaskId
}

func (s *QueryAuditTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryAuditTaskRequest) SetArticleId(v string) *QueryAuditTaskRequest {
	s.ArticleId = &v
	return s
}

func (s *QueryAuditTaskRequest) SetContentAuditTaskId(v string) *QueryAuditTaskRequest {
	s.ContentAuditTaskId = &v
	return s
}

func (s *QueryAuditTaskRequest) SetWorkspaceId(v string) *QueryAuditTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryAuditTaskRequest) Validate() error {
	return dara.Validate(s)
}
