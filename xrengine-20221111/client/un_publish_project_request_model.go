// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnPublishProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJwtToken(v string) *UnPublishProjectRequest
	GetJwtToken() *string
	SetProjectId(v string) *UnPublishProjectRequest
	GetProjectId() *string
}

type UnPublishProjectRequest struct {
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	// This parameter is required.
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s UnPublishProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s UnPublishProjectRequest) GoString() string {
	return s.String()
}

func (s *UnPublishProjectRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *UnPublishProjectRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *UnPublishProjectRequest) SetJwtToken(v string) *UnPublishProjectRequest {
	s.JwtToken = &v
	return s
}

func (s *UnPublishProjectRequest) SetProjectId(v string) *UnPublishProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *UnPublishProjectRequest) Validate() error {
	return dara.Validate(s)
}
