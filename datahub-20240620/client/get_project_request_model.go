// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectName(v string) *GetProjectRequest
	GetProjectName() *string
}

type GetProjectRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test_project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s GetProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProjectRequest) GoString() string {
	return s.String()
}

func (s *GetProjectRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetProjectRequest) SetProjectName(v string) *GetProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *GetProjectRequest) Validate() error {
	return dara.Validate(s)
}
