// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContextDatabaseWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMemberName(v string) *CreateContextDatabaseWorkspaceRequest
	GetMemberName() *string
	SetWorkspaceName(v string) *CreateContextDatabaseWorkspaceRequest
	GetWorkspaceName() *string
}

type CreateContextDatabaseWorkspaceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// my-member
	MemberName *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-workspace
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s CreateContextDatabaseWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateContextDatabaseWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *CreateContextDatabaseWorkspaceRequest) GetMemberName() *string {
	return s.MemberName
}

func (s *CreateContextDatabaseWorkspaceRequest) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *CreateContextDatabaseWorkspaceRequest) SetMemberName(v string) *CreateContextDatabaseWorkspaceRequest {
	s.MemberName = &v
	return s
}

func (s *CreateContextDatabaseWorkspaceRequest) SetWorkspaceName(v string) *CreateContextDatabaseWorkspaceRequest {
	s.WorkspaceName = &v
	return s
}

func (s *CreateContextDatabaseWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}
