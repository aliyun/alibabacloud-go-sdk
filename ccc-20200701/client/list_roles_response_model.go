// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRolesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRolesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRolesResponse
	GetStatusCode() *int32
	SetBody(v *ListRolesResponseBody) *ListRolesResponse
	GetBody() *ListRolesResponseBody
}

type ListRolesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRolesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRolesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRolesResponse) GoString() string {
	return s.String()
}

func (s *ListRolesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRolesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRolesResponse) GetBody() *ListRolesResponseBody {
	return s.Body
}

func (s *ListRolesResponse) SetHeaders(v map[string]*string) *ListRolesResponse {
	s.Headers = v
	return s
}

func (s *ListRolesResponse) SetStatusCode(v int32) *ListRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRolesResponse) SetBody(v *ListRolesResponseBody) *ListRolesResponse {
	s.Body = v
	return s
}

func (s *ListRolesResponse) Validate() error {
	return dara.Validate(s)
}
