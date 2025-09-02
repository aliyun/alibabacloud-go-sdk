// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectRolesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v int64) *ListProjectRolesRequest
	GetProjectId() *int64
}

type ListProjectRolesRequest struct {
	// The ID of the DataWorks workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 27
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListProjectRolesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectRolesRequest) GoString() string {
	return s.String()
}

func (s *ListProjectRolesRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListProjectRolesRequest) SetProjectId(v int64) *ListProjectRolesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListProjectRolesRequest) Validate() error {
	return dara.Validate(s)
}
