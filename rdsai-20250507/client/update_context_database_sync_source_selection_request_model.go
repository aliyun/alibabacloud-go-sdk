// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContextDatabaseSyncSourceSelectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateContextDatabaseSyncSourceSelectionRequest
	GetClientToken() *string
	SetSelectedNodeKeys(v []*string) *UpdateContextDatabaseSyncSourceSelectionRequest
	GetSelectedNodeKeys() []*string
	SetSelectionPolicy(v string) *UpdateContextDatabaseSyncSourceSelectionRequest
	GetSelectionPolicy() *string
	SetSourceId(v string) *UpdateContextDatabaseSyncSourceSelectionRequest
	GetSourceId() *string
	SetWorkspaceId(v string) *UpdateContextDatabaseSyncSourceSelectionRequest
	GetWorkspaceId() *string
}

type UpdateContextDatabaseSyncSourceSelectionRequest struct {
	// The idempotency token.
	//
	// example:
	//
	// sync-connection-20260923-001
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The list of selected node identifiers.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["doc:1234567890"]
	SelectedNodeKeys []*string `json:"SelectedNodeKeys,omitempty" xml:"SelectedNodeKeys,omitempty" type:"Repeated"`
	// The selection policy.
	//
	// example:
	//
	// subtree
	SelectionPolicy *string `json:"SelectionPolicy,omitempty" xml:"SelectionPolicy,omitempty"`
	// The data source ID, which is a system-generated UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ws-as1llqmkol****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateContextDatabaseSyncSourceSelectionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextDatabaseSyncSourceSelectionRequest) GoString() string {
	return s.String()
}

func (s *UpdateContextDatabaseSyncSourceSelectionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateContextDatabaseSyncSourceSelectionRequest) GetSelectedNodeKeys() []*string {
	return s.SelectedNodeKeys
}

func (s *UpdateContextDatabaseSyncSourceSelectionRequest) GetSelectionPolicy() *string {
	return s.SelectionPolicy
}

func (s *UpdateContextDatabaseSyncSourceSelectionRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *UpdateContextDatabaseSyncSourceSelectionRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateContextDatabaseSyncSourceSelectionRequest) SetClientToken(v string) *UpdateContextDatabaseSyncSourceSelectionRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateContextDatabaseSyncSourceSelectionRequest) SetSelectedNodeKeys(v []*string) *UpdateContextDatabaseSyncSourceSelectionRequest {
	s.SelectedNodeKeys = v
	return s
}

func (s *UpdateContextDatabaseSyncSourceSelectionRequest) SetSelectionPolicy(v string) *UpdateContextDatabaseSyncSourceSelectionRequest {
	s.SelectionPolicy = &v
	return s
}

func (s *UpdateContextDatabaseSyncSourceSelectionRequest) SetSourceId(v string) *UpdateContextDatabaseSyncSourceSelectionRequest {
	s.SourceId = &v
	return s
}

func (s *UpdateContextDatabaseSyncSourceSelectionRequest) SetWorkspaceId(v string) *UpdateContextDatabaseSyncSourceSelectionRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateContextDatabaseSyncSourceSelectionRequest) Validate() error {
	return dara.Validate(s)
}
