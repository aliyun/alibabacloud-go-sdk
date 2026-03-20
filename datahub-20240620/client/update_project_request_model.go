// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *UpdateProjectRequest
	GetComment() *string
	SetProjectName(v string) *UpdateProjectRequest
	GetProjectName() *string
}

type UpdateProjectRequest struct {
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s UpdateProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectRequest) GetComment() *string {
	return s.Comment
}

func (s *UpdateProjectRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *UpdateProjectRequest) SetComment(v string) *UpdateProjectRequest {
	s.Comment = &v
	return s
}

func (s *UpdateProjectRequest) SetProjectName(v string) *UpdateProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateProjectRequest) Validate() error {
	return dara.Validate(s)
}
