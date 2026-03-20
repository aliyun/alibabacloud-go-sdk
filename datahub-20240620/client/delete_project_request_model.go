// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectName(v string) *DeleteProjectRequest
	GetProjectName() *string
}

type DeleteProjectRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// jr_techout
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s DeleteProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteProjectRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DeleteProjectRequest) SetProjectName(v string) *DeleteProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *DeleteProjectRequest) Validate() error {
	return dara.Validate(s)
}
