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
	// The name of the project. For more information, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
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
