// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryFailedContextDatabaseSyncRunRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RetryFailedContextDatabaseSyncRunRequest
	GetClientToken() *string
	SetRunId(v string) *RetryFailedContextDatabaseSyncRunRequest
	GetRunId() *string
	SetWorkspaceId(v string) *RetryFailedContextDatabaseSyncRunRequest
	GetWorkspaceId() *string
}

type RetryFailedContextDatabaseSyncRunRequest struct {
	// example:
	//
	// sync-connection-20260923-001
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ws-as1llqmkol****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RetryFailedContextDatabaseSyncRunRequest) String() string {
	return dara.Prettify(s)
}

func (s RetryFailedContextDatabaseSyncRunRequest) GoString() string {
	return s.String()
}

func (s *RetryFailedContextDatabaseSyncRunRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RetryFailedContextDatabaseSyncRunRequest) GetRunId() *string {
	return s.RunId
}

func (s *RetryFailedContextDatabaseSyncRunRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RetryFailedContextDatabaseSyncRunRequest) SetClientToken(v string) *RetryFailedContextDatabaseSyncRunRequest {
	s.ClientToken = &v
	return s
}

func (s *RetryFailedContextDatabaseSyncRunRequest) SetRunId(v string) *RetryFailedContextDatabaseSyncRunRequest {
	s.RunId = &v
	return s
}

func (s *RetryFailedContextDatabaseSyncRunRequest) SetWorkspaceId(v string) *RetryFailedContextDatabaseSyncRunRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RetryFailedContextDatabaseSyncRunRequest) Validate() error {
	return dara.Validate(s)
}
