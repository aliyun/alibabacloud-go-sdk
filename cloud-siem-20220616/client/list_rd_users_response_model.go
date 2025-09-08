// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRdUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRdUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRdUsersResponse
	GetStatusCode() *int32
	SetBody(v *ListRdUsersResponseBody) *ListRdUsersResponse
	GetBody() *ListRdUsersResponseBody
}

type ListRdUsersResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRdUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRdUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRdUsersResponse) GoString() string {
	return s.String()
}

func (s *ListRdUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRdUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRdUsersResponse) GetBody() *ListRdUsersResponseBody {
	return s.Body
}

func (s *ListRdUsersResponse) SetHeaders(v map[string]*string) *ListRdUsersResponse {
	s.Headers = v
	return s
}

func (s *ListRdUsersResponse) SetStatusCode(v int32) *ListRdUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRdUsersResponse) SetBody(v *ListRdUsersResponseBody) *ListRdUsersResponse {
	s.Body = v
	return s
}

func (s *ListRdUsersResponse) Validate() error {
	return dara.Validate(s)
}
