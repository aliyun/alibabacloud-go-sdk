// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatabaseAccountsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDatabaseAccountsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDatabaseAccountsResponse
	GetStatusCode() *int32
	SetBody(v *ListDatabaseAccountsResponseBody) *ListDatabaseAccountsResponse
	GetBody() *ListDatabaseAccountsResponseBody
}

type ListDatabaseAccountsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDatabaseAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDatabaseAccountsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDatabaseAccountsResponse) GoString() string {
	return s.String()
}

func (s *ListDatabaseAccountsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDatabaseAccountsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDatabaseAccountsResponse) GetBody() *ListDatabaseAccountsResponseBody {
	return s.Body
}

func (s *ListDatabaseAccountsResponse) SetHeaders(v map[string]*string) *ListDatabaseAccountsResponse {
	s.Headers = v
	return s
}

func (s *ListDatabaseAccountsResponse) SetStatusCode(v int32) *ListDatabaseAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDatabaseAccountsResponse) SetBody(v *ListDatabaseAccountsResponseBody) *ListDatabaseAccountsResponse {
	s.Body = v
	return s
}

func (s *ListDatabaseAccountsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
