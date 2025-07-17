// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedDatabasesForUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAuthorizedDatabasesForUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAuthorizedDatabasesForUserResponse
	GetStatusCode() *int32
	SetBody(v *ListAuthorizedDatabasesForUserResponseBody) *ListAuthorizedDatabasesForUserResponse
	GetBody() *ListAuthorizedDatabasesForUserResponseBody
}

type ListAuthorizedDatabasesForUserResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAuthorizedDatabasesForUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAuthorizedDatabasesForUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedDatabasesForUserResponse) GoString() string {
	return s.String()
}

func (s *ListAuthorizedDatabasesForUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAuthorizedDatabasesForUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAuthorizedDatabasesForUserResponse) GetBody() *ListAuthorizedDatabasesForUserResponseBody {
	return s.Body
}

func (s *ListAuthorizedDatabasesForUserResponse) SetHeaders(v map[string]*string) *ListAuthorizedDatabasesForUserResponse {
	s.Headers = v
	return s
}

func (s *ListAuthorizedDatabasesForUserResponse) SetStatusCode(v int32) *ListAuthorizedDatabasesForUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAuthorizedDatabasesForUserResponse) SetBody(v *ListAuthorizedDatabasesForUserResponseBody) *ListAuthorizedDatabasesForUserResponse {
	s.Body = v
	return s
}

func (s *ListAuthorizedDatabasesForUserResponse) Validate() error {
	return dara.Validate(s)
}
