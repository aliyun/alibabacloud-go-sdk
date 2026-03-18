// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v string) *UpdateProjectRequest
	GetProjectId() *string
	SetProjectName(v string) *UpdateProjectRequest
	GetProjectName() *string
}

type UpdateProjectRequest struct {
	// example:
	//
	// 123
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// example:
	//
	// MyProject
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
}

func (s UpdateProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *UpdateProjectRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *UpdateProjectRequest) SetProjectId(v string) *UpdateProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateProjectRequest) SetProjectName(v string) *UpdateProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateProjectRequest) Validate() error {
	return dara.Validate(s)
}
