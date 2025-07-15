// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRamUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRamUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRamUsersResponse
	GetStatusCode() *int32
	SetBody(v *ListRamUsersResponseBody) *ListRamUsersResponse
	GetBody() *ListRamUsersResponseBody
}

type ListRamUsersResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRamUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRamUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRamUsersResponse) GoString() string {
	return s.String()
}

func (s *ListRamUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRamUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRamUsersResponse) GetBody() *ListRamUsersResponseBody {
	return s.Body
}

func (s *ListRamUsersResponse) SetHeaders(v map[string]*string) *ListRamUsersResponse {
	s.Headers = v
	return s
}

func (s *ListRamUsersResponse) SetStatusCode(v int32) *ListRamUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRamUsersResponse) SetBody(v *ListRamUsersResponseBody) *ListRamUsersResponse {
	s.Body = v
	return s
}

func (s *ListRamUsersResponse) Validate() error {
	return dara.Validate(s)
}
