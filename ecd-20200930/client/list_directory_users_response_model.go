// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDirectoryUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDirectoryUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDirectoryUsersResponse
	GetStatusCode() *int32
	SetBody(v *ListDirectoryUsersResponseBody) *ListDirectoryUsersResponse
	GetBody() *ListDirectoryUsersResponseBody
}

type ListDirectoryUsersResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDirectoryUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDirectoryUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDirectoryUsersResponse) GoString() string {
	return s.String()
}

func (s *ListDirectoryUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDirectoryUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDirectoryUsersResponse) GetBody() *ListDirectoryUsersResponseBody {
	return s.Body
}

func (s *ListDirectoryUsersResponse) SetHeaders(v map[string]*string) *ListDirectoryUsersResponse {
	s.Headers = v
	return s
}

func (s *ListDirectoryUsersResponse) SetStatusCode(v int32) *ListDirectoryUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDirectoryUsersResponse) SetBody(v *ListDirectoryUsersResponseBody) *ListDirectoryUsersResponse {
	s.Body = v
	return s
}

func (s *ListDirectoryUsersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
