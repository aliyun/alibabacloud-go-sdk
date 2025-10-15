// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationAccountsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApplicationAccountsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApplicationAccountsResponse
	GetStatusCode() *int32
	SetBody(v *ListApplicationAccountsResponseBody) *ListApplicationAccountsResponse
	GetBody() *ListApplicationAccountsResponseBody
}

type ListApplicationAccountsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationAccountsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationAccountsResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationAccountsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApplicationAccountsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApplicationAccountsResponse) GetBody() *ListApplicationAccountsResponseBody {
	return s.Body
}

func (s *ListApplicationAccountsResponse) SetHeaders(v map[string]*string) *ListApplicationAccountsResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationAccountsResponse) SetStatusCode(v int32) *ListApplicationAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationAccountsResponse) SetBody(v *ListApplicationAccountsResponseBody) *ListApplicationAccountsResponse {
	s.Body = v
	return s
}

func (s *ListApplicationAccountsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
