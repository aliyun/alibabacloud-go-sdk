// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectName(v string) *CreateProjectRequest
	GetProjectName() *string
	SetProjectType(v string) *CreateProjectRequest
	GetProjectType() *string
}

type CreateProjectRequest struct {
	// example:
	//
	// MyProject
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	// example:
	//
	// online_oral_evaluation_post_paid_call_count
	ProjectType *string `json:"projectType,omitempty" xml:"projectType,omitempty"`
}

func (s CreateProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateProjectRequest) GetProjectType() *string {
	return s.ProjectType
}

func (s *CreateProjectRequest) SetProjectName(v string) *CreateProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateProjectRequest) SetProjectType(v string) *CreateProjectRequest {
	s.ProjectType = &v
	return s
}

func (s *CreateProjectRequest) Validate() error {
	return dara.Validate(s)
}
