// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectRolesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProjectRolesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProjectRolesResponse
	GetStatusCode() *int32
	SetBody(v *ListProjectRolesResponseBody) *ListProjectRolesResponse
	GetBody() *ListProjectRolesResponseBody
}

type ListProjectRolesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProjectRolesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProjectRolesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProjectRolesResponse) GoString() string {
	return s.String()
}

func (s *ListProjectRolesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProjectRolesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProjectRolesResponse) GetBody() *ListProjectRolesResponseBody {
	return s.Body
}

func (s *ListProjectRolesResponse) SetHeaders(v map[string]*string) *ListProjectRolesResponse {
	s.Headers = v
	return s
}

func (s *ListProjectRolesResponse) SetStatusCode(v int32) *ListProjectRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectRolesResponse) SetBody(v *ListProjectRolesResponseBody) *ListProjectRolesResponse {
	s.Body = v
	return s
}

func (s *ListProjectRolesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
