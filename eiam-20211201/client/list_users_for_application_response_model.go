// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersForApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUsersForApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUsersForApplicationResponse
	GetStatusCode() *int32
	SetBody(v *ListUsersForApplicationResponseBody) *ListUsersForApplicationResponse
	GetBody() *ListUsersForApplicationResponseBody
}

type ListUsersForApplicationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUsersForApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUsersForApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUsersForApplicationResponse) GoString() string {
	return s.String()
}

func (s *ListUsersForApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUsersForApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUsersForApplicationResponse) GetBody() *ListUsersForApplicationResponseBody {
	return s.Body
}

func (s *ListUsersForApplicationResponse) SetHeaders(v map[string]*string) *ListUsersForApplicationResponse {
	s.Headers = v
	return s
}

func (s *ListUsersForApplicationResponse) SetStatusCode(v int32) *ListUsersForApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUsersForApplicationResponse) SetBody(v *ListUsersForApplicationResponseBody) *ListUsersForApplicationResponse {
	s.Body = v
	return s
}

func (s *ListUsersForApplicationResponse) Validate() error {
	return dara.Validate(s)
}
