// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectRolesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *ListProjectRolesRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListProjectRolesRequest
	GetOpUserId() *string
	SetProjectType(v string) *ListProjectRolesRequest
	GetProjectType() *string
}

type ListProjectRolesRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The user ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The project type. Valid values: BASIC (basic mode project), DEV (development environment project), PROD (production environment project), and TAG (tag platform project).
	//
	// This parameter is required.
	//
	// example:
	//
	// DEV
	ProjectType *string `json:"ProjectType,omitempty" xml:"ProjectType,omitempty"`
}

func (s ListProjectRolesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectRolesRequest) GoString() string {
	return s.String()
}

func (s *ListProjectRolesRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListProjectRolesRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListProjectRolesRequest) GetProjectType() *string {
	return s.ProjectType
}

func (s *ListProjectRolesRequest) SetOpTenantId(v int64) *ListProjectRolesRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListProjectRolesRequest) SetOpUserId(v string) *ListProjectRolesRequest {
	s.OpUserId = &v
	return s
}

func (s *ListProjectRolesRequest) SetProjectType(v string) *ListProjectRolesRequest {
	s.ProjectType = &v
	return s
}

func (s *ListProjectRolesRequest) Validate() error {
	return dara.Validate(s)
}
