// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJwtToken(v string) *DeleteProjectRequest
	GetJwtToken() *string
	SetProjectId(v string) *DeleteProjectRequest
	GetProjectId() *string
}

type DeleteProjectRequest struct {
	JwtToken  *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteProjectRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *DeleteProjectRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *DeleteProjectRequest) SetJwtToken(v string) *DeleteProjectRequest {
	s.JwtToken = &v
	return s
}

func (s *DeleteProjectRequest) SetProjectId(v string) *DeleteProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteProjectRequest) Validate() error {
	return dara.Validate(s)
}
