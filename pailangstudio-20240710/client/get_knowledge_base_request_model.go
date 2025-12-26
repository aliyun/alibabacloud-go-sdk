// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeBaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVersionName(v string) *GetKnowledgeBaseRequest
	GetVersionName() *string
	SetWorkspaceId(v string) *GetKnowledgeBaseRequest
	GetWorkspaceId() *string
}

type GetKnowledgeBaseRequest struct {
	// example:
	//
	// v1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 478**
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetKnowledgeBaseRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseRequest) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseRequest) GetVersionName() *string {
	return s.VersionName
}

func (s *GetKnowledgeBaseRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetKnowledgeBaseRequest) SetVersionName(v string) *GetKnowledgeBaseRequest {
	s.VersionName = &v
	return s
}

func (s *GetKnowledgeBaseRequest) SetWorkspaceId(v string) *GetKnowledgeBaseRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetKnowledgeBaseRequest) Validate() error {
	return dara.Validate(s)
}
