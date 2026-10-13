// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContextDatabaseSyncJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateContextDatabaseSyncJobRequest
	GetClientToken() *string
	SetEnabled(v bool) *UpdateContextDatabaseSyncJobRequest
	GetEnabled() *bool
	SetFilters(v map[string]interface{}) *UpdateContextDatabaseSyncJobRequest
	GetFilters() map[string]interface{}
	SetJobId(v string) *UpdateContextDatabaseSyncJobRequest
	GetJobId() *string
	SetName(v string) *UpdateContextDatabaseSyncJobRequest
	GetName() *string
	SetTargetKnowledgeBaseId(v string) *UpdateContextDatabaseSyncJobRequest
	GetTargetKnowledgeBaseId() *string
	SetTargetKnowledgeBaseName(v string) *UpdateContextDatabaseSyncJobRequest
	GetTargetKnowledgeBaseName() *string
	SetWorkspaceId(v string) *UpdateContextDatabaseSyncJobRequest
	GetWorkspaceId() *string
}

type UpdateContextDatabaseSyncJobRequest struct {
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
	Filters map[string]interface{} `json:"Filters,omitempty" xml:"Filters,omitempty"`
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

func (s UpdateContextDatabaseSyncJobRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextDatabaseSyncJobRequest) GoString() string {
	return s.String()
}

func (s *UpdateContextDatabaseSyncJobRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateContextDatabaseSyncJobRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateContextDatabaseSyncJobRequest) GetFilters() map[string]interface{} {
	return s.Filters
}

func (s *UpdateContextDatabaseSyncJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *UpdateContextDatabaseSyncJobRequest) GetName() *string {
	return s.Name
}

func (s *UpdateContextDatabaseSyncJobRequest) GetTargetKnowledgeBaseId() *string {
	return s.TargetKnowledgeBaseId
}

func (s *UpdateContextDatabaseSyncJobRequest) GetTargetKnowledgeBaseName() *string {
	return s.TargetKnowledgeBaseName
}

func (s *UpdateContextDatabaseSyncJobRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateContextDatabaseSyncJobRequest) SetClientToken(v string) *UpdateContextDatabaseSyncJobRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobRequest) SetEnabled(v bool) *UpdateContextDatabaseSyncJobRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobRequest) SetFilters(v map[string]interface{}) *UpdateContextDatabaseSyncJobRequest {
	s.Filters = v
	return s
}

func (s *UpdateContextDatabaseSyncJobRequest) SetJobId(v string) *UpdateContextDatabaseSyncJobRequest {
	s.JobId = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobRequest) SetName(v string) *UpdateContextDatabaseSyncJobRequest {
	s.Name = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobRequest) SetTargetKnowledgeBaseId(v string) *UpdateContextDatabaseSyncJobRequest {
	s.TargetKnowledgeBaseId = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobRequest) SetTargetKnowledgeBaseName(v string) *UpdateContextDatabaseSyncJobRequest {
	s.TargetKnowledgeBaseName = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobRequest) SetWorkspaceId(v string) *UpdateContextDatabaseSyncJobRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobRequest) Validate() error {
	return dara.Validate(s)
}
