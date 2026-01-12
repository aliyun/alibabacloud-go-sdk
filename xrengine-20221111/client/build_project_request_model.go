// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBuildProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJwtToken(v string) *BuildProjectRequest
	GetJwtToken() *string
	SetProjectId(v string) *BuildProjectRequest
	GetProjectId() *string
	SetVideoName(v string) *BuildProjectRequest
	GetVideoName() *string
}

type BuildProjectRequest struct {
	JwtToken  *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	VideoName *string `json:"VideoName,omitempty" xml:"VideoName,omitempty"`
}

func (s BuildProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s BuildProjectRequest) GoString() string {
	return s.String()
}

func (s *BuildProjectRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *BuildProjectRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *BuildProjectRequest) GetVideoName() *string {
	return s.VideoName
}

func (s *BuildProjectRequest) SetJwtToken(v string) *BuildProjectRequest {
	s.JwtToken = &v
	return s
}

func (s *BuildProjectRequest) SetProjectId(v string) *BuildProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *BuildProjectRequest) SetVideoName(v string) *BuildProjectRequest {
	s.VideoName = &v
	return s
}

func (s *BuildProjectRequest) Validate() error {
	return dara.Validate(s)
}
