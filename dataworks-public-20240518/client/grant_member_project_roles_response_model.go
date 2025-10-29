// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantMemberProjectRolesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GrantMemberProjectRolesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GrantMemberProjectRolesResponse
	GetStatusCode() *int32
	SetBody(v *GrantMemberProjectRolesResponseBody) *GrantMemberProjectRolesResponse
	GetBody() *GrantMemberProjectRolesResponseBody
}

type GrantMemberProjectRolesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantMemberProjectRolesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantMemberProjectRolesResponse) String() string {
	return dara.Prettify(s)
}

func (s GrantMemberProjectRolesResponse) GoString() string {
	return s.String()
}

func (s *GrantMemberProjectRolesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GrantMemberProjectRolesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GrantMemberProjectRolesResponse) GetBody() *GrantMemberProjectRolesResponseBody {
	return s.Body
}

func (s *GrantMemberProjectRolesResponse) SetHeaders(v map[string]*string) *GrantMemberProjectRolesResponse {
	s.Headers = v
	return s
}

func (s *GrantMemberProjectRolesResponse) SetStatusCode(v int32) *GrantMemberProjectRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantMemberProjectRolesResponse) SetBody(v *GrantMemberProjectRolesResponseBody) *GrantMemberProjectRolesResponse {
	s.Body = v
	return s
}

func (s *GrantMemberProjectRolesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
