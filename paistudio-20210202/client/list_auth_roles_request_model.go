// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthRolesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsGenerateToken(v string) *ListAuthRolesRequest
	GetIsGenerateToken() *string
	SetWorkspaceId(v string) *ListAuthRolesRequest
	GetWorkspaceId() *string
}

type ListAuthRolesRequest struct {
	// example:
	//
	// true
	IsGenerateToken *string `json:"IsGenerateToken,omitempty" xml:"IsGenerateToken,omitempty"`
	// example:
	//
	// 12345
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListAuthRolesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAuthRolesRequest) GoString() string {
	return s.String()
}

func (s *ListAuthRolesRequest) GetIsGenerateToken() *string {
	return s.IsGenerateToken
}

func (s *ListAuthRolesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListAuthRolesRequest) SetIsGenerateToken(v string) *ListAuthRolesRequest {
	s.IsGenerateToken = &v
	return s
}

func (s *ListAuthRolesRequest) SetWorkspaceId(v string) *ListAuthRolesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListAuthRolesRequest) Validate() error {
	return dara.Validate(s)
}
