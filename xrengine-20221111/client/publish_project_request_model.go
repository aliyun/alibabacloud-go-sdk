// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJwtToken(v string) *PublishProjectRequest
	GetJwtToken() *string
	SetProjectId(v string) *PublishProjectRequest
	GetProjectId() *string
}

type PublishProjectRequest struct {
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	// This parameter is required.
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s PublishProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishProjectRequest) GoString() string {
	return s.String()
}

func (s *PublishProjectRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *PublishProjectRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *PublishProjectRequest) SetJwtToken(v string) *PublishProjectRequest {
	s.JwtToken = &v
	return s
}

func (s *PublishProjectRequest) SetProjectId(v string) *PublishProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *PublishProjectRequest) Validate() error {
	return dara.Validate(s)
}
