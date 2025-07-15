// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivilegesOfUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPrivilegesOfUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPrivilegesOfUserResponse
	GetStatusCode() *int32
	SetBody(v *ListPrivilegesOfUserResponseBody) *ListPrivilegesOfUserResponse
	GetBody() *ListPrivilegesOfUserResponseBody
}

type ListPrivilegesOfUserResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPrivilegesOfUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPrivilegesOfUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPrivilegesOfUserResponse) GoString() string {
	return s.String()
}

func (s *ListPrivilegesOfUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPrivilegesOfUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPrivilegesOfUserResponse) GetBody() *ListPrivilegesOfUserResponseBody {
	return s.Body
}

func (s *ListPrivilegesOfUserResponse) SetHeaders(v map[string]*string) *ListPrivilegesOfUserResponse {
	s.Headers = v
	return s
}

func (s *ListPrivilegesOfUserResponse) SetStatusCode(v int32) *ListPrivilegesOfUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrivilegesOfUserResponse) SetBody(v *ListPrivilegesOfUserResponseBody) *ListPrivilegesOfUserResponse {
	s.Body = v
	return s
}

func (s *ListPrivilegesOfUserResponse) Validate() error {
	return dara.Validate(s)
}
