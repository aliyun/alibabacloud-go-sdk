// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRoleUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRoleUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRoleUsersResponse
	GetStatusCode() *int32
	SetBody(v *ListRoleUsersResponseBody) *ListRoleUsersResponse
	GetBody() *ListRoleUsersResponseBody
}

type ListRoleUsersResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRoleUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRoleUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRoleUsersResponse) GoString() string {
	return s.String()
}

func (s *ListRoleUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRoleUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRoleUsersResponse) GetBody() *ListRoleUsersResponseBody {
	return s.Body
}

func (s *ListRoleUsersResponse) SetHeaders(v map[string]*string) *ListRoleUsersResponse {
	s.Headers = v
	return s
}

func (s *ListRoleUsersResponse) SetStatusCode(v int32) *ListRoleUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRoleUsersResponse) SetBody(v *ListRoleUsersResponseBody) *ListRoleUsersResponse {
	s.Body = v
	return s
}

func (s *ListRoleUsersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
