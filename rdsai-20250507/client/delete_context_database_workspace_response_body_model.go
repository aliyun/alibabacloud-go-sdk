// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContextDatabaseWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v string) *DeleteContextDatabaseWorkspaceResponseBody
	GetCreatedAt() *string
	SetRequestId(v string) *DeleteContextDatabaseWorkspaceResponseBody
	GetRequestId() *string
	SetStatus(v string) *DeleteContextDatabaseWorkspaceResponseBody
	GetStatus() *string
	SetType(v string) *DeleteContextDatabaseWorkspaceResponseBody
	GetType() *string
	SetWorkspaceId(v string) *DeleteContextDatabaseWorkspaceResponseBody
	GetWorkspaceId() *string
	SetWorkspaceName(v string) *DeleteContextDatabaseWorkspaceResponseBody
	GetWorkspaceName() *string
}

type DeleteContextDatabaseWorkspaceResponseBody struct {
	// example:
	//
	// 2026-05-28T17:59:55Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Deleted
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// personal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// ws-as1llqmkol****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// example:
	//
	// my-workspace
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s DeleteContextDatabaseWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteContextDatabaseWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContextDatabaseWorkspaceResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *DeleteContextDatabaseWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteContextDatabaseWorkspaceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DeleteContextDatabaseWorkspaceResponseBody) GetType() *string {
	return s.Type
}

func (s *DeleteContextDatabaseWorkspaceResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteContextDatabaseWorkspaceResponseBody) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *DeleteContextDatabaseWorkspaceResponseBody) SetCreatedAt(v string) *DeleteContextDatabaseWorkspaceResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *DeleteContextDatabaseWorkspaceResponseBody) SetRequestId(v string) *DeleteContextDatabaseWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteContextDatabaseWorkspaceResponseBody) SetStatus(v string) *DeleteContextDatabaseWorkspaceResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteContextDatabaseWorkspaceResponseBody) SetType(v string) *DeleteContextDatabaseWorkspaceResponseBody {
	s.Type = &v
	return s
}

func (s *DeleteContextDatabaseWorkspaceResponseBody) SetWorkspaceId(v string) *DeleteContextDatabaseWorkspaceResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteContextDatabaseWorkspaceResponseBody) SetWorkspaceName(v string) *DeleteContextDatabaseWorkspaceResponseBody {
	s.WorkspaceName = &v
	return s
}

func (s *DeleteContextDatabaseWorkspaceResponseBody) Validate() error {
	return dara.Validate(s)
}
