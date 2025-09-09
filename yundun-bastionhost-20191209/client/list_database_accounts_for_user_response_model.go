// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatabaseAccountsForUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDatabaseAccountsForUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDatabaseAccountsForUserResponse
	GetStatusCode() *int32
	SetBody(v *ListDatabaseAccountsForUserResponseBody) *ListDatabaseAccountsForUserResponse
	GetBody() *ListDatabaseAccountsForUserResponseBody
}

type ListDatabaseAccountsForUserResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDatabaseAccountsForUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDatabaseAccountsForUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDatabaseAccountsForUserResponse) GoString() string {
	return s.String()
}

func (s *ListDatabaseAccountsForUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDatabaseAccountsForUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDatabaseAccountsForUserResponse) GetBody() *ListDatabaseAccountsForUserResponseBody {
	return s.Body
}

func (s *ListDatabaseAccountsForUserResponse) SetHeaders(v map[string]*string) *ListDatabaseAccountsForUserResponse {
	s.Headers = v
	return s
}

func (s *ListDatabaseAccountsForUserResponse) SetStatusCode(v int32) *ListDatabaseAccountsForUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDatabaseAccountsForUserResponse) SetBody(v *ListDatabaseAccountsForUserResponseBody) *ListDatabaseAccountsForUserResponse {
	s.Body = v
	return s
}

func (s *ListDatabaseAccountsForUserResponse) Validate() error {
	return dara.Validate(s)
}
