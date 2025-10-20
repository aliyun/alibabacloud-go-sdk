// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserRolesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserRolesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserRolesResponse
	GetStatusCode() *int32
	SetBody(v *ListUserRolesResponseBody) *ListUserRolesResponse
	GetBody() *ListUserRolesResponseBody
}

type ListUserRolesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserRolesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserRolesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserRolesResponse) GoString() string {
	return s.String()
}

func (s *ListUserRolesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserRolesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserRolesResponse) GetBody() *ListUserRolesResponseBody {
	return s.Body
}

func (s *ListUserRolesResponse) SetHeaders(v map[string]*string) *ListUserRolesResponse {
	s.Headers = v
	return s
}

func (s *ListUserRolesResponse) SetStatusCode(v int32) *ListUserRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserRolesResponse) SetBody(v *ListUserRolesResponseBody) *ListUserRolesResponse {
	s.Body = v
	return s
}

func (s *ListUserRolesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
