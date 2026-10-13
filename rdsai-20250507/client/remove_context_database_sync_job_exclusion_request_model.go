// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveContextDatabaseSyncJobExclusionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RemoveContextDatabaseSyncJobExclusionRequest
	GetClientToken() *string
	SetDocumentId(v string) *RemoveContextDatabaseSyncJobExclusionRequest
	GetDocumentId() *string
	SetJobId(v string) *RemoveContextDatabaseSyncJobExclusionRequest
	GetJobId() *string
	SetWorkspaceId(v string) *RemoveContextDatabaseSyncJobExclusionRequest
	GetWorkspaceId() *string
}

type RemoveContextDatabaseSyncJobExclusionRequest struct {
	// The idempotency token.
	//
	// example:
	//
	// sync-connection-20260923-001
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The document ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// doc-as1llqmkol****
	DocumentId *string `json:"DocumentId,omitempty" xml:"DocumentId,omitempty"`
	// The job ID, which is a system-generated UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ws-as1llqmkol****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RemoveContextDatabaseSyncJobExclusionRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveContextDatabaseSyncJobExclusionRequest) GoString() string {
	return s.String()
}

func (s *RemoveContextDatabaseSyncJobExclusionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RemoveContextDatabaseSyncJobExclusionRequest) GetDocumentId() *string {
	return s.DocumentId
}

func (s *RemoveContextDatabaseSyncJobExclusionRequest) GetJobId() *string {
	return s.JobId
}

func (s *RemoveContextDatabaseSyncJobExclusionRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RemoveContextDatabaseSyncJobExclusionRequest) SetClientToken(v string) *RemoveContextDatabaseSyncJobExclusionRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveContextDatabaseSyncJobExclusionRequest) SetDocumentId(v string) *RemoveContextDatabaseSyncJobExclusionRequest {
	s.DocumentId = &v
	return s
}

func (s *RemoveContextDatabaseSyncJobExclusionRequest) SetJobId(v string) *RemoveContextDatabaseSyncJobExclusionRequest {
	s.JobId = &v
	return s
}

func (s *RemoveContextDatabaseSyncJobExclusionRequest) SetWorkspaceId(v string) *RemoveContextDatabaseSyncJobExclusionRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RemoveContextDatabaseSyncJobExclusionRequest) Validate() error {
	return dara.Validate(s)
}
