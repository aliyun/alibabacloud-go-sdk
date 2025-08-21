// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRenderingProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v string) *DeleteRenderingProjectRequest
	GetProjectId() *string
}

type DeleteRenderingProjectRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// project-422bc38dfgh5eb44149f135ef76304f63b
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteRenderingProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRenderingProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteRenderingProjectRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *DeleteRenderingProjectRequest) SetProjectId(v string) *DeleteRenderingProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteRenderingProjectRequest) Validate() error {
	return dara.Validate(s)
}
