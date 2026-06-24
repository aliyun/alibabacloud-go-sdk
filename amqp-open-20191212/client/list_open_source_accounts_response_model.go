// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOpenSourceAccountsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOpenSourceAccountsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOpenSourceAccountsResponse
	GetStatusCode() *int32
	SetBody(v *ListOpenSourceAccountsResponseBody) *ListOpenSourceAccountsResponse
	GetBody() *ListOpenSourceAccountsResponseBody
}

type ListOpenSourceAccountsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOpenSourceAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOpenSourceAccountsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOpenSourceAccountsResponse) GoString() string {
	return s.String()
}

func (s *ListOpenSourceAccountsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOpenSourceAccountsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOpenSourceAccountsResponse) GetBody() *ListOpenSourceAccountsResponseBody {
	return s.Body
}

func (s *ListOpenSourceAccountsResponse) SetHeaders(v map[string]*string) *ListOpenSourceAccountsResponse {
	s.Headers = v
	return s
}

func (s *ListOpenSourceAccountsResponse) SetStatusCode(v int32) *ListOpenSourceAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOpenSourceAccountsResponse) SetBody(v *ListOpenSourceAccountsResponseBody) *ListOpenSourceAccountsResponse {
	s.Body = v
	return s
}

func (s *ListOpenSourceAccountsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
