// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProjectUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProjectUsersResponse
	GetStatusCode() *int32
	SetBody(v *ListProjectUsersResponseBody) *ListProjectUsersResponse
	GetBody() *ListProjectUsersResponseBody
}

type ListProjectUsersResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProjectUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProjectUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProjectUsersResponse) GoString() string {
	return s.String()
}

func (s *ListProjectUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProjectUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProjectUsersResponse) GetBody() *ListProjectUsersResponseBody {
	return s.Body
}

func (s *ListProjectUsersResponse) SetHeaders(v map[string]*string) *ListProjectUsersResponse {
	s.Headers = v
	return s
}

func (s *ListProjectUsersResponse) SetStatusCode(v int32) *ListProjectUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectUsersResponse) SetBody(v *ListProjectUsersResponseBody) *ListProjectUsersResponse {
	s.Body = v
	return s
}

func (s *ListProjectUsersResponse) Validate() error {
	return dara.Validate(s)
}
