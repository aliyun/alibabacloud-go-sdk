// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAuthorizedUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAuthorizedUsersResponse
	GetStatusCode() *int32
	SetBody(v *ListAuthorizedUsersResponseBody) *ListAuthorizedUsersResponse
	GetBody() *ListAuthorizedUsersResponseBody
}

type ListAuthorizedUsersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAuthorizedUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAuthorizedUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedUsersResponse) GoString() string {
	return s.String()
}

func (s *ListAuthorizedUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAuthorizedUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAuthorizedUsersResponse) GetBody() *ListAuthorizedUsersResponseBody {
	return s.Body
}

func (s *ListAuthorizedUsersResponse) SetHeaders(v map[string]*string) *ListAuthorizedUsersResponse {
	s.Headers = v
	return s
}

func (s *ListAuthorizedUsersResponse) SetStatusCode(v int32) *ListAuthorizedUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAuthorizedUsersResponse) SetBody(v *ListAuthorizedUsersResponseBody) *ListAuthorizedUsersResponse {
	s.Body = v
	return s
}

func (s *ListAuthorizedUsersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
