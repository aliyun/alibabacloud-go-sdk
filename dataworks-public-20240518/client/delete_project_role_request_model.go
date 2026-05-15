// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProjectRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteProjectRoleRequest
	GetCode() *string
	SetProjectId(v int64) *DeleteProjectRoleRequest
	GetProjectId() *int64
}

type DeleteProjectRoleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// base_role_xxx
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteProjectRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteProjectRoleRequest) GoString() string {
	return s.String()
}

func (s *DeleteProjectRoleRequest) GetCode() *string {
	return s.Code
}

func (s *DeleteProjectRoleRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteProjectRoleRequest) SetCode(v string) *DeleteProjectRoleRequest {
	s.Code = &v
	return s
}

func (s *DeleteProjectRoleRequest) SetProjectId(v int64) *DeleteProjectRoleRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteProjectRoleRequest) Validate() error {
	return dara.Validate(s)
}
