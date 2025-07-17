// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedUsersForInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAuthorizedUsersForInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAuthorizedUsersForInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ListAuthorizedUsersForInstanceResponseBody) *ListAuthorizedUsersForInstanceResponse
	GetBody() *ListAuthorizedUsersForInstanceResponseBody
}

type ListAuthorizedUsersForInstanceResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAuthorizedUsersForInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAuthorizedUsersForInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedUsersForInstanceResponse) GoString() string {
	return s.String()
}

func (s *ListAuthorizedUsersForInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAuthorizedUsersForInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAuthorizedUsersForInstanceResponse) GetBody() *ListAuthorizedUsersForInstanceResponseBody {
	return s.Body
}

func (s *ListAuthorizedUsersForInstanceResponse) SetHeaders(v map[string]*string) *ListAuthorizedUsersForInstanceResponse {
	s.Headers = v
	return s
}

func (s *ListAuthorizedUsersForInstanceResponse) SetStatusCode(v int32) *ListAuthorizedUsersForInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAuthorizedUsersForInstanceResponse) SetBody(v *ListAuthorizedUsersForInstanceResponseBody) *ListAuthorizedUsersForInstanceResponse {
	s.Body = v
	return s
}

func (s *ListAuthorizedUsersForInstanceResponse) Validate() error {
	return dara.Validate(s)
}
