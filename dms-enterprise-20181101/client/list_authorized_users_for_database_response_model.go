// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedUsersForDatabaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAuthorizedUsersForDatabaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAuthorizedUsersForDatabaseResponse
	GetStatusCode() *int32
	SetBody(v *ListAuthorizedUsersForDatabaseResponseBody) *ListAuthorizedUsersForDatabaseResponse
	GetBody() *ListAuthorizedUsersForDatabaseResponseBody
}

type ListAuthorizedUsersForDatabaseResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAuthorizedUsersForDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAuthorizedUsersForDatabaseResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedUsersForDatabaseResponse) GoString() string {
	return s.String()
}

func (s *ListAuthorizedUsersForDatabaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAuthorizedUsersForDatabaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAuthorizedUsersForDatabaseResponse) GetBody() *ListAuthorizedUsersForDatabaseResponseBody {
	return s.Body
}

func (s *ListAuthorizedUsersForDatabaseResponse) SetHeaders(v map[string]*string) *ListAuthorizedUsersForDatabaseResponse {
	s.Headers = v
	return s
}

func (s *ListAuthorizedUsersForDatabaseResponse) SetStatusCode(v int32) *ListAuthorizedUsersForDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAuthorizedUsersForDatabaseResponse) SetBody(v *ListAuthorizedUsersForDatabaseResponseBody) *ListAuthorizedUsersForDatabaseResponse {
	s.Body = v
	return s
}

func (s *ListAuthorizedUsersForDatabaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
