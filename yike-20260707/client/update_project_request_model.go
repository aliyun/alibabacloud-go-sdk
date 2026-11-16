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
	SetTitle(v string) *UpdateProjectRequest
	GetTitle() *string
}

type UpdateProjectRequest struct {
	// The project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pd_061716****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The new name of the project.
	//
	// example:
	//
	// Product Demo
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
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

func (s *UpdateProjectRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateProjectRequest) SetProjectId(v string) *UpdateProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateProjectRequest) SetTitle(v string) *UpdateProjectRequest {
	s.Title = &v
	return s
}

func (s *UpdateProjectRequest) Validate() error {
	return dara.Validate(s)
}
