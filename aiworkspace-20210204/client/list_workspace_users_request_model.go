// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspaceUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v string) *ListWorkspaceUsersRequest
	GetUserId() *string
	SetUserName(v string) *ListWorkspaceUsersRequest
	GetUserName() *string
}

type ListWorkspaceUsersRequest struct {
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The display names of users who can be added to the workspace as members.
	//
	// example:
	//
	// doctest****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListWorkspaceUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspaceUsersRequest) GoString() string {
	return s.String()
}

func (s *ListWorkspaceUsersRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListWorkspaceUsersRequest) GetUserName() *string {
	return s.UserName
}

func (s *ListWorkspaceUsersRequest) SetUserId(v string) *ListWorkspaceUsersRequest {
	s.UserId = &v
	return s
}

func (s *ListWorkspaceUsersRequest) SetUserName(v string) *ListWorkspaceUsersRequest {
	s.UserName = &v
	return s
}

func (s *ListWorkspaceUsersRequest) Validate() error {
	return dara.Validate(s)
}
