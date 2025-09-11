// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspaceUserRolesByUserIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v string) *ListWorkspaceUserRolesByUserIdRequest
	GetUserId() *string
}

type ListWorkspaceUserRolesByUserIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// afas-********asfasg
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListWorkspaceUserRolesByUserIdRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspaceUserRolesByUserIdRequest) GoString() string {
	return s.String()
}

func (s *ListWorkspaceUserRolesByUserIdRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListWorkspaceUserRolesByUserIdRequest) SetUserId(v string) *ListWorkspaceUserRolesByUserIdRequest {
	s.UserId = &v
	return s
}

func (s *ListWorkspaceUserRolesByUserIdRequest) Validate() error {
	return dara.Validate(s)
}
