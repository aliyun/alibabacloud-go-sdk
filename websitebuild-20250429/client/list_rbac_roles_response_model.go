// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRbacRolesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRbacRolesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRbacRolesResponse
	GetStatusCode() *int32
	SetBody(v *ListRbacRolesResponseBody) *ListRbacRolesResponse
	GetBody() *ListRbacRolesResponseBody
}

type ListRbacRolesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRbacRolesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRbacRolesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRbacRolesResponse) GoString() string {
	return s.String()
}

func (s *ListRbacRolesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRbacRolesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRbacRolesResponse) GetBody() *ListRbacRolesResponseBody {
	return s.Body
}

func (s *ListRbacRolesResponse) SetHeaders(v map[string]*string) *ListRbacRolesResponse {
	s.Headers = v
	return s
}

func (s *ListRbacRolesResponse) SetStatusCode(v int32) *ListRbacRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRbacRolesResponse) SetBody(v *ListRbacRolesResponseBody) *ListRbacRolesResponse {
	s.Body = v
	return s
}

func (s *ListRbacRolesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
