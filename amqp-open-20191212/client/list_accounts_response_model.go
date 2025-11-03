// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccountsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAccountsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAccountsResponse
	GetStatusCode() *int32
	SetBody(v *ListAccountsResponseBody) *ListAccountsResponse
	GetBody() *ListAccountsResponseBody
}

type ListAccountsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAccountsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAccountsResponse) GoString() string {
	return s.String()
}

func (s *ListAccountsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAccountsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAccountsResponse) GetBody() *ListAccountsResponseBody {
	return s.Body
}

func (s *ListAccountsResponse) SetHeaders(v map[string]*string) *ListAccountsResponse {
	s.Headers = v
	return s
}

func (s *ListAccountsResponse) SetStatusCode(v int32) *ListAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccountsResponse) SetBody(v *ListAccountsResponseBody) *ListAccountsResponse {
	s.Body = v
	return s
}

func (s *ListAccountsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
