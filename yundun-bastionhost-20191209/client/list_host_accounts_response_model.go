// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostAccountsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHostAccountsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHostAccountsResponse
	GetStatusCode() *int32
	SetBody(v *ListHostAccountsResponseBody) *ListHostAccountsResponse
	GetBody() *ListHostAccountsResponseBody
}

type ListHostAccountsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHostAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHostAccountsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHostAccountsResponse) GoString() string {
	return s.String()
}

func (s *ListHostAccountsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHostAccountsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHostAccountsResponse) GetBody() *ListHostAccountsResponseBody {
	return s.Body
}

func (s *ListHostAccountsResponse) SetHeaders(v map[string]*string) *ListHostAccountsResponse {
	s.Headers = v
	return s
}

func (s *ListHostAccountsResponse) SetStatusCode(v int32) *ListHostAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHostAccountsResponse) SetBody(v *ListHostAccountsResponseBody) *ListHostAccountsResponse {
	s.Body = v
	return s
}

func (s *ListHostAccountsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
