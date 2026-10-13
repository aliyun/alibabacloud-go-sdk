// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContextDatabaseSyncSourceSelectionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateContextDatabaseSyncSourceSelectionShrinkRequest
	GetClientToken() *string
	SetSelectedNodeKeysShrink(v string) *UpdateContextDatabaseSyncSourceSelectionShrinkRequest
	GetSelectedNodeKeysShrink() *string
	SetSelectionPolicy(v string) *UpdateContextDatabaseSyncSourceSelectionShrinkRequest
	GetSelectionPolicy() *string
	SetSourceId(v string) *UpdateContextDatabaseSyncSourceSelectionShrinkRequest
	GetSourceId() *string
	SetWorkspaceId(v string) *UpdateContextDatabaseSyncSourceSelectionShrinkRequest
	GetWorkspaceId() *string
}

type UpdateContextDatabaseSyncSourceSelectionShrinkRequest struct {
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
	SelectedNodeKeysShrink *string `json:"SelectedNodeKeys,omitempty" xml:"SelectedNodeKeys,omitempty"`
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

func (s UpdateContextDatabaseSyncSourceSelectionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextDatabaseSyncSourceSelectionShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateContextDatabaseSyncSourceSelectionShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateContextDatabaseSyncSourceSelectionShrinkRequest) GetSelectedNodeKeysShrink() *string {
	return s.SelectedNodeKeysShrink
}

func (s *UpdateContextDatabaseSyncSourceSelectionShrinkRequest) GetSelectionPolicy() *string {
	return s.SelectionPolicy
}

func (s *UpdateContextDatabaseSyncSourceSelectionShrinkRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *UpdateContextDatabaseSyncSourceSelectionShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateContextDatabaseSyncSourceSelectionShrinkRequest) SetClientToken(v string) *UpdateContextDatabaseSyncSourceSelectionShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateContextDatabaseSyncSourceSelectionShrinkRequest) SetSelectedNodeKeysShrink(v string) *UpdateContextDatabaseSyncSourceSelectionShrinkRequest {
	s.SelectedNodeKeysShrink = &v
	return s
}

func (s *UpdateContextDatabaseSyncSourceSelectionShrinkRequest) SetSelectionPolicy(v string) *UpdateContextDatabaseSyncSourceSelectionShrinkRequest {
	s.SelectionPolicy = &v
	return s
}

func (s *UpdateContextDatabaseSyncSourceSelectionShrinkRequest) SetSourceId(v string) *UpdateContextDatabaseSyncSourceSelectionShrinkRequest {
	s.SourceId = &v
	return s
}

func (s *UpdateContextDatabaseSyncSourceSelectionShrinkRequest) SetWorkspaceId(v string) *UpdateContextDatabaseSyncSourceSelectionShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateContextDatabaseSyncSourceSelectionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
