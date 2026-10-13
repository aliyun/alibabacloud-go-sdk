// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContextDatabaseSyncJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateContextDatabaseSyncJobShrinkRequest
	GetClientToken() *string
	SetEnabled(v bool) *UpdateContextDatabaseSyncJobShrinkRequest
	GetEnabled() *bool
	SetFiltersShrink(v string) *UpdateContextDatabaseSyncJobShrinkRequest
	GetFiltersShrink() *string
	SetJobId(v string) *UpdateContextDatabaseSyncJobShrinkRequest
	GetJobId() *string
	SetName(v string) *UpdateContextDatabaseSyncJobShrinkRequest
	GetName() *string
	SetTargetKnowledgeBaseId(v string) *UpdateContextDatabaseSyncJobShrinkRequest
	GetTargetKnowledgeBaseId() *string
	SetTargetKnowledgeBaseName(v string) *UpdateContextDatabaseSyncJobShrinkRequest
	GetTargetKnowledgeBaseName() *string
	SetWorkspaceId(v string) *UpdateContextDatabaseSyncJobShrinkRequest
	GetWorkspaceId() *string
}

type UpdateContextDatabaseSyncJobShrinkRequest struct {
	// example:
	//
	// sync-connection-20260923-001
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// {"include":["/研发文档/**"],"exclude":["/归档文档/**"],"excludedDocumentIds":[],"sizeLimit":2147483648,"typeAllowlist":["docx","pdf"]}
	FiltersShrink *string `json:"Filters,omitempty" xml:"Filters,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 研发团队文档导入
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// id-001
	TargetKnowledgeBaseId *string `json:"TargetKnowledgeBaseId,omitempty" xml:"TargetKnowledgeBaseId,omitempty"`
	// example:
	//
	// 示例名称
	TargetKnowledgeBaseName *string `json:"TargetKnowledgeBaseName,omitempty" xml:"TargetKnowledgeBaseName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ws-as1llqmkol****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateContextDatabaseSyncJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextDatabaseSyncJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateContextDatabaseSyncJobShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateContextDatabaseSyncJobShrinkRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateContextDatabaseSyncJobShrinkRequest) GetFiltersShrink() *string {
	return s.FiltersShrink
}

func (s *UpdateContextDatabaseSyncJobShrinkRequest) GetJobId() *string {
	return s.JobId
}

func (s *UpdateContextDatabaseSyncJobShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateContextDatabaseSyncJobShrinkRequest) GetTargetKnowledgeBaseId() *string {
	return s.TargetKnowledgeBaseId
}

func (s *UpdateContextDatabaseSyncJobShrinkRequest) GetTargetKnowledgeBaseName() *string {
	return s.TargetKnowledgeBaseName
}

func (s *UpdateContextDatabaseSyncJobShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateContextDatabaseSyncJobShrinkRequest) SetClientToken(v string) *UpdateContextDatabaseSyncJobShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobShrinkRequest) SetEnabled(v bool) *UpdateContextDatabaseSyncJobShrinkRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobShrinkRequest) SetFiltersShrink(v string) *UpdateContextDatabaseSyncJobShrinkRequest {
	s.FiltersShrink = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobShrinkRequest) SetJobId(v string) *UpdateContextDatabaseSyncJobShrinkRequest {
	s.JobId = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobShrinkRequest) SetName(v string) *UpdateContextDatabaseSyncJobShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobShrinkRequest) SetTargetKnowledgeBaseId(v string) *UpdateContextDatabaseSyncJobShrinkRequest {
	s.TargetKnowledgeBaseId = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobShrinkRequest) SetTargetKnowledgeBaseName(v string) *UpdateContextDatabaseSyncJobShrinkRequest {
	s.TargetKnowledgeBaseName = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobShrinkRequest) SetWorkspaceId(v string) *UpdateContextDatabaseSyncJobShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
