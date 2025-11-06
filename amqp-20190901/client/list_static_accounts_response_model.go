// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStaticAccountsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListStaticAccountsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListStaticAccountsResponse
	GetStatusCode() *int32
	SetBody(v *ListStaticAccountsResponseBody) *ListStaticAccountsResponse
	GetBody() *ListStaticAccountsResponseBody
}

type ListStaticAccountsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListStaticAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListStaticAccountsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListStaticAccountsResponse) GoString() string {
	return s.String()
}

func (s *ListStaticAccountsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListStaticAccountsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListStaticAccountsResponse) GetBody() *ListStaticAccountsResponseBody {
	return s.Body
}

func (s *ListStaticAccountsResponse) SetHeaders(v map[string]*string) *ListStaticAccountsResponse {
	s.Headers = v
	return s
}

func (s *ListStaticAccountsResponse) SetStatusCode(v int32) *ListStaticAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStaticAccountsResponse) SetBody(v *ListStaticAccountsResponseBody) *ListStaticAccountsResponse {
	s.Body = v
	return s
}

func (s *ListStaticAccountsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
