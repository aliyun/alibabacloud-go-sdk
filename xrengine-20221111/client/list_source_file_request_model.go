// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSourceFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJwtToken(v string) *ListSourceFileRequest
	GetJwtToken() *string
	SetProjectId(v string) *ListSourceFileRequest
	GetProjectId() *string
}

type ListSourceFileRequest struct {
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	// This parameter is required.
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListSourceFileRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSourceFileRequest) GoString() string {
	return s.String()
}

func (s *ListSourceFileRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *ListSourceFileRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListSourceFileRequest) SetJwtToken(v string) *ListSourceFileRequest {
	s.JwtToken = &v
	return s
}

func (s *ListSourceFileRequest) SetProjectId(v string) *ListSourceFileRequest {
	s.ProjectId = &v
	return s
}

func (s *ListSourceFileRequest) Validate() error {
	return dara.Validate(s)
}
