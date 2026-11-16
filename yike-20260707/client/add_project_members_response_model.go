// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddProjectMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddProjectMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddProjectMembersResponse
	GetStatusCode() *int32
	SetBody(v *AddProjectMembersResponseBody) *AddProjectMembersResponse
	GetBody() *AddProjectMembersResponseBody
}

type AddProjectMembersResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddProjectMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddProjectMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s AddProjectMembersResponse) GoString() string {
	return s.String()
}

func (s *AddProjectMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddProjectMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddProjectMembersResponse) GetBody() *AddProjectMembersResponseBody {
	return s.Body
}

func (s *AddProjectMembersResponse) SetHeaders(v map[string]*string) *AddProjectMembersResponse {
	s.Headers = v
	return s
}

func (s *AddProjectMembersResponse) SetStatusCode(v int32) *AddProjectMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *AddProjectMembersResponse) SetBody(v *AddProjectMembersResponseBody) *AddProjectMembersResponse {
	s.Body = v
	return s
}

func (s *AddProjectMembersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
