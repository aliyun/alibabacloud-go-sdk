// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationDatabaseAccountsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOperationDatabaseAccountsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOperationDatabaseAccountsResponse
	GetStatusCode() *int32
	SetBody(v *ListOperationDatabaseAccountsResponseBody) *ListOperationDatabaseAccountsResponse
	GetBody() *ListOperationDatabaseAccountsResponseBody
}

type ListOperationDatabaseAccountsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOperationDatabaseAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOperationDatabaseAccountsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOperationDatabaseAccountsResponse) GoString() string {
	return s.String()
}

func (s *ListOperationDatabaseAccountsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOperationDatabaseAccountsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOperationDatabaseAccountsResponse) GetBody() *ListOperationDatabaseAccountsResponseBody {
	return s.Body
}

func (s *ListOperationDatabaseAccountsResponse) SetHeaders(v map[string]*string) *ListOperationDatabaseAccountsResponse {
	s.Headers = v
	return s
}

func (s *ListOperationDatabaseAccountsResponse) SetStatusCode(v int32) *ListOperationDatabaseAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOperationDatabaseAccountsResponse) SetBody(v *ListOperationDatabaseAccountsResponseBody) *ListOperationDatabaseAccountsResponse {
	s.Body = v
	return s
}

func (s *ListOperationDatabaseAccountsResponse) Validate() error {
	return dara.Validate(s)
}
