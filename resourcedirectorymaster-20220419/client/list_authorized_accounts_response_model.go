// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedAccountsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAuthorizedAccountsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAuthorizedAccountsResponse
	GetStatusCode() *int32
	SetBody(v *ListAuthorizedAccountsResponseBody) *ListAuthorizedAccountsResponse
	GetBody() *ListAuthorizedAccountsResponseBody
}

type ListAuthorizedAccountsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAuthorizedAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAuthorizedAccountsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedAccountsResponse) GoString() string {
	return s.String()
}

func (s *ListAuthorizedAccountsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAuthorizedAccountsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAuthorizedAccountsResponse) GetBody() *ListAuthorizedAccountsResponseBody {
	return s.Body
}

func (s *ListAuthorizedAccountsResponse) SetHeaders(v map[string]*string) *ListAuthorizedAccountsResponse {
	s.Headers = v
	return s
}

func (s *ListAuthorizedAccountsResponse) SetStatusCode(v int32) *ListAuthorizedAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAuthorizedAccountsResponse) SetBody(v *ListAuthorizedAccountsResponseBody) *ListAuthorizedAccountsResponse {
	s.Body = v
	return s
}

func (s *ListAuthorizedAccountsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
