// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClientUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClientUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClientUsersResponse
	GetStatusCode() *int32
	SetBody(v *ListClientUsersResponseBody) *ListClientUsersResponse
	GetBody() *ListClientUsersResponseBody
}

type ListClientUsersResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClientUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClientUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClientUsersResponse) GoString() string {
	return s.String()
}

func (s *ListClientUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClientUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClientUsersResponse) GetBody() *ListClientUsersResponseBody {
	return s.Body
}

func (s *ListClientUsersResponse) SetHeaders(v map[string]*string) *ListClientUsersResponse {
	s.Headers = v
	return s
}

func (s *ListClientUsersResponse) SetStatusCode(v int32) *ListClientUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClientUsersResponse) SetBody(v *ListClientUsersResponseBody) *ListClientUsersResponse {
	s.Body = v
	return s
}

func (s *ListClientUsersResponse) Validate() error {
	return dara.Validate(s)
}
