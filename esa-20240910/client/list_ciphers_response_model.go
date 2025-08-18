// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCiphersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCiphersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCiphersResponse
	GetStatusCode() *int32
	SetBody(v *ListCiphersResponseBody) *ListCiphersResponse
	GetBody() *ListCiphersResponseBody
}

type ListCiphersResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCiphersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCiphersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCiphersResponse) GoString() string {
	return s.String()
}

func (s *ListCiphersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCiphersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCiphersResponse) GetBody() *ListCiphersResponseBody {
	return s.Body
}

func (s *ListCiphersResponse) SetHeaders(v map[string]*string) *ListCiphersResponse {
	s.Headers = v
	return s
}

func (s *ListCiphersResponse) SetStatusCode(v int32) *ListCiphersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCiphersResponse) SetBody(v *ListCiphersResponseBody) *ListCiphersResponse {
	s.Body = v
	return s
}

func (s *ListCiphersResponse) Validate() error {
	return dara.Validate(s)
}
