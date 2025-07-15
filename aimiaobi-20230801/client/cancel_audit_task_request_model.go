// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAuditTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArticleId(v string) *CancelAuditTaskRequest
	GetArticleId() *string
	SetContentAuditTaskId(v string) *CancelAuditTaskRequest
	GetContentAuditTaskId() *string
	SetWorkspaceId(v string) *CancelAuditTaskRequest
	GetWorkspaceId() *string
}

type CancelAuditTaskRequest struct {
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
	// llm-xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CancelAuditTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelAuditTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelAuditTaskRequest) GetArticleId() *string {
	return s.ArticleId
}

func (s *CancelAuditTaskRequest) GetContentAuditTaskId() *string {
	return s.ContentAuditTaskId
}

func (s *CancelAuditTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CancelAuditTaskRequest) SetArticleId(v string) *CancelAuditTaskRequest {
	s.ArticleId = &v
	return s
}

func (s *CancelAuditTaskRequest) SetContentAuditTaskId(v string) *CancelAuditTaskRequest {
	s.ContentAuditTaskId = &v
	return s
}

func (s *CancelAuditTaskRequest) SetWorkspaceId(v string) *CancelAuditTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CancelAuditTaskRequest) Validate() error {
	return dara.Validate(s)
}
