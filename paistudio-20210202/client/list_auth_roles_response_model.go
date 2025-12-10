// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthRolesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAuthRolesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAuthRolesResponse
	GetStatusCode() *int32
	SetBody(v *ListAuthRolesResponseBody) *ListAuthRolesResponse
	GetBody() *ListAuthRolesResponseBody
}

type ListAuthRolesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAuthRolesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAuthRolesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAuthRolesResponse) GoString() string {
	return s.String()
}

func (s *ListAuthRolesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAuthRolesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAuthRolesResponse) GetBody() *ListAuthRolesResponseBody {
	return s.Body
}

func (s *ListAuthRolesResponse) SetHeaders(v map[string]*string) *ListAuthRolesResponse {
	s.Headers = v
	return s
}

func (s *ListAuthRolesResponse) SetStatusCode(v int32) *ListAuthRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAuthRolesResponse) SetBody(v *ListAuthRolesResponseBody) *ListAuthRolesResponse {
	s.Body = v
	return s
}

func (s *ListAuthRolesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
