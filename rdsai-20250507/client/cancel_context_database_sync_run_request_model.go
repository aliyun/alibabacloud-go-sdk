// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelContextDatabaseSyncRunRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CancelContextDatabaseSyncRunRequest
	GetClientToken() *string
	SetRunId(v string) *CancelContextDatabaseSyncRunRequest
	GetRunId() *string
	SetWorkspaceId(v string) *CancelContextDatabaseSyncRunRequest
	GetWorkspaceId() *string
}

type CancelContextDatabaseSyncRunRequest struct {
	// The idempotence token.
	//
	// example:
	//
	// sync-connection-20260923-001
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The run ID, a system-generated UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ws-as1llqmkol****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CancelContextDatabaseSyncRunRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelContextDatabaseSyncRunRequest) GoString() string {
	return s.String()
}

func (s *CancelContextDatabaseSyncRunRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CancelContextDatabaseSyncRunRequest) GetRunId() *string {
	return s.RunId
}

func (s *CancelContextDatabaseSyncRunRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CancelContextDatabaseSyncRunRequest) SetClientToken(v string) *CancelContextDatabaseSyncRunRequest {
	s.ClientToken = &v
	return s
}

func (s *CancelContextDatabaseSyncRunRequest) SetRunId(v string) *CancelContextDatabaseSyncRunRequest {
	s.RunId = &v
	return s
}

func (s *CancelContextDatabaseSyncRunRequest) SetWorkspaceId(v string) *CancelContextDatabaseSyncRunRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CancelContextDatabaseSyncRunRequest) Validate() error {
	return dara.Validate(s)
}
