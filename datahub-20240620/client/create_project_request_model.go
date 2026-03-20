// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *CreateProjectRequest
	GetComment() *string
	SetProjectName(v string) *CreateProjectRequest
	GetProjectName() *string
}

type CreateProjectRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test_comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s CreateProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateProjectRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateProjectRequest) SetComment(v string) *CreateProjectRequest {
	s.Comment = &v
	return s
}

func (s *CreateProjectRequest) SetProjectName(v string) *CreateProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateProjectRequest) Validate() error {
	return dara.Validate(s)
}
