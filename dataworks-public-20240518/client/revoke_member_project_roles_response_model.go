// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeMemberProjectRolesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeMemberProjectRolesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeMemberProjectRolesResponse
	GetStatusCode() *int32
	SetBody(v *RevokeMemberProjectRolesResponseBody) *RevokeMemberProjectRolesResponse
	GetBody() *RevokeMemberProjectRolesResponseBody
}

type RevokeMemberProjectRolesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeMemberProjectRolesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeMemberProjectRolesResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeMemberProjectRolesResponse) GoString() string {
	return s.String()
}

func (s *RevokeMemberProjectRolesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeMemberProjectRolesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeMemberProjectRolesResponse) GetBody() *RevokeMemberProjectRolesResponseBody {
	return s.Body
}

func (s *RevokeMemberProjectRolesResponse) SetHeaders(v map[string]*string) *RevokeMemberProjectRolesResponse {
	s.Headers = v
	return s
}

func (s *RevokeMemberProjectRolesResponse) SetStatusCode(v int32) *RevokeMemberProjectRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeMemberProjectRolesResponse) SetBody(v *RevokeMemberProjectRolesResponseBody) *RevokeMemberProjectRolesResponse {
	s.Body = v
	return s
}

func (s *RevokeMemberProjectRolesResponse) Validate() error {
	return dara.Validate(s)
}
