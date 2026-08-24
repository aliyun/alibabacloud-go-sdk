// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseWorkspacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListContextDatabaseWorkspacesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListContextDatabaseWorkspacesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListContextDatabaseWorkspacesResponseBody
	GetRequestId() *string
	SetWorkspaces(v []*ListContextDatabaseWorkspacesResponseBodyWorkspaces) *ListContextDatabaseWorkspacesResponseBody
	GetWorkspaces() []*ListContextDatabaseWorkspacesResponseBodyWorkspaces
}

type ListContextDatabaseWorkspacesResponseBody struct {
	// This field is empty.
	//
	// example:
	//
	// (null)
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// This field is empty.
	//
	// example:
	//
	// (null)
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of workspaces.
	Workspaces []*ListContextDatabaseWorkspacesResponseBodyWorkspaces `json:"Workspaces,omitempty" xml:"Workspaces,omitempty" type:"Repeated"`
}

func (s ListContextDatabaseWorkspacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseWorkspacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseWorkspacesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListContextDatabaseWorkspacesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListContextDatabaseWorkspacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListContextDatabaseWorkspacesResponseBody) GetWorkspaces() []*ListContextDatabaseWorkspacesResponseBodyWorkspaces {
	return s.Workspaces
}

func (s *ListContextDatabaseWorkspacesResponseBody) SetMaxResults(v int32) *ListContextDatabaseWorkspacesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListContextDatabaseWorkspacesResponseBody) SetNextToken(v string) *ListContextDatabaseWorkspacesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListContextDatabaseWorkspacesResponseBody) SetRequestId(v string) *ListContextDatabaseWorkspacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListContextDatabaseWorkspacesResponseBody) SetWorkspaces(v []*ListContextDatabaseWorkspacesResponseBodyWorkspaces) *ListContextDatabaseWorkspacesResponseBody {
	s.Workspaces = v
	return s
}

func (s *ListContextDatabaseWorkspacesResponseBody) Validate() error {
	if s.Workspaces != nil {
		for _, item := range s.Workspaces {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListContextDatabaseWorkspacesResponseBodyWorkspaces struct {
	// The time when the workspace was created, in ISO-8601 format.
	//
	// example:
	//
	// 2026-05-28T17:59:55Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The workspace status.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The workspace type.
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
	// The workspace name.
	//
	// example:
	//
	// my-workspace
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s ListContextDatabaseWorkspacesResponseBodyWorkspaces) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseWorkspacesResponseBodyWorkspaces) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseWorkspacesResponseBodyWorkspaces) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListContextDatabaseWorkspacesResponseBodyWorkspaces) GetStatus() *string {
	return s.Status
}

func (s *ListContextDatabaseWorkspacesResponseBodyWorkspaces) GetType() *string {
	return s.Type
}

func (s *ListContextDatabaseWorkspacesResponseBodyWorkspaces) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListContextDatabaseWorkspacesResponseBodyWorkspaces) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *ListContextDatabaseWorkspacesResponseBodyWorkspaces) SetCreatedAt(v string) *ListContextDatabaseWorkspacesResponseBodyWorkspaces {
	s.CreatedAt = &v
	return s
}

func (s *ListContextDatabaseWorkspacesResponseBodyWorkspaces) SetStatus(v string) *ListContextDatabaseWorkspacesResponseBodyWorkspaces {
	s.Status = &v
	return s
}

func (s *ListContextDatabaseWorkspacesResponseBodyWorkspaces) SetType(v string) *ListContextDatabaseWorkspacesResponseBodyWorkspaces {
	s.Type = &v
	return s
}

func (s *ListContextDatabaseWorkspacesResponseBodyWorkspaces) SetWorkspaceId(v string) *ListContextDatabaseWorkspacesResponseBodyWorkspaces {
	s.WorkspaceId = &v
	return s
}

func (s *ListContextDatabaseWorkspacesResponseBodyWorkspaces) SetWorkspaceName(v string) *ListContextDatabaseWorkspacesResponseBodyWorkspaces {
	s.WorkspaceName = &v
	return s
}

func (s *ListContextDatabaseWorkspacesResponseBodyWorkspaces) Validate() error {
	return dara.Validate(s)
}
