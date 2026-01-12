// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryBuildBreakpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJwtToken(v string) *QueryBuildBreakpointRequest
	GetJwtToken() *string
	SetProjectId(v string) *QueryBuildBreakpointRequest
	GetProjectId() *string
}

type QueryBuildBreakpointRequest struct {
	JwtToken  *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s QueryBuildBreakpointRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryBuildBreakpointRequest) GoString() string {
	return s.String()
}

func (s *QueryBuildBreakpointRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *QueryBuildBreakpointRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *QueryBuildBreakpointRequest) SetJwtToken(v string) *QueryBuildBreakpointRequest {
	s.JwtToken = &v
	return s
}

func (s *QueryBuildBreakpointRequest) SetProjectId(v string) *QueryBuildBreakpointRequest {
	s.ProjectId = &v
	return s
}

func (s *QueryBuildBreakpointRequest) Validate() error {
	return dara.Validate(s)
}
