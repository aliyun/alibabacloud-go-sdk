// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContextDatabaseWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v string) *UpdateContextDatabaseWorkspaceResponseBody
	GetCreatedAt() *string
	SetRequestId(v string) *UpdateContextDatabaseWorkspaceResponseBody
	GetRequestId() *string
	SetStatus(v string) *UpdateContextDatabaseWorkspaceResponseBody
	GetStatus() *string
	SetType(v string) *UpdateContextDatabaseWorkspaceResponseBody
	GetType() *string
	SetWorkspaceId(v string) *UpdateContextDatabaseWorkspaceResponseBody
	GetWorkspaceId() *string
	SetWorkspaceName(v string) *UpdateContextDatabaseWorkspaceResponseBody
	GetWorkspaceName() *string
}

type UpdateContextDatabaseWorkspaceResponseBody struct {
	// The time when the workspace was created, in ISO-8601 format.
	//
	// example:
	//
	// 2026-05-28T17:59:55Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the workspace.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the workspace.
	//
	// example:
	//
	// personal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// ws-as1llqmkol****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The modified name of the workspace.
	//
	// example:
	//
	// my-workspace
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s UpdateContextDatabaseWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextDatabaseWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateContextDatabaseWorkspaceResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *UpdateContextDatabaseWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateContextDatabaseWorkspaceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UpdateContextDatabaseWorkspaceResponseBody) GetType() *string {
	return s.Type
}

func (s *UpdateContextDatabaseWorkspaceResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateContextDatabaseWorkspaceResponseBody) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *UpdateContextDatabaseWorkspaceResponseBody) SetCreatedAt(v string) *UpdateContextDatabaseWorkspaceResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *UpdateContextDatabaseWorkspaceResponseBody) SetRequestId(v string) *UpdateContextDatabaseWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateContextDatabaseWorkspaceResponseBody) SetStatus(v string) *UpdateContextDatabaseWorkspaceResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateContextDatabaseWorkspaceResponseBody) SetType(v string) *UpdateContextDatabaseWorkspaceResponseBody {
	s.Type = &v
	return s
}

func (s *UpdateContextDatabaseWorkspaceResponseBody) SetWorkspaceId(v string) *UpdateContextDatabaseWorkspaceResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateContextDatabaseWorkspaceResponseBody) SetWorkspaceName(v string) *UpdateContextDatabaseWorkspaceResponseBody {
	s.WorkspaceName = &v
	return s
}

func (s *UpdateContextDatabaseWorkspaceResponseBody) Validate() error {
	return dara.Validate(s)
}
