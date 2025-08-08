// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProvidedSharesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProvidedSharesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProvidedSharesResponse
	GetStatusCode() *int32
	SetBody(v *ListProvidedSharesResponseBody) *ListProvidedSharesResponse
	GetBody() *ListProvidedSharesResponseBody
}

type ListProvidedSharesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProvidedSharesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProvidedSharesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProvidedSharesResponse) GoString() string {
	return s.String()
}

func (s *ListProvidedSharesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProvidedSharesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProvidedSharesResponse) GetBody() *ListProvidedSharesResponseBody {
	return s.Body
}

func (s *ListProvidedSharesResponse) SetHeaders(v map[string]*string) *ListProvidedSharesResponse {
	s.Headers = v
	return s
}

func (s *ListProvidedSharesResponse) SetStatusCode(v int32) *ListProvidedSharesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProvidedSharesResponse) SetBody(v *ListProvidedSharesResponseBody) *ListProvidedSharesResponse {
	s.Body = v
	return s
}

func (s *ListProvidedSharesResponse) Validate() error {
	return dara.Validate(s)
}
