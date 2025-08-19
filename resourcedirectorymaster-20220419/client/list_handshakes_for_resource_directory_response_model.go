// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHandshakesForResourceDirectoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHandshakesForResourceDirectoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHandshakesForResourceDirectoryResponse
	GetStatusCode() *int32
	SetBody(v *ListHandshakesForResourceDirectoryResponseBody) *ListHandshakesForResourceDirectoryResponse
	GetBody() *ListHandshakesForResourceDirectoryResponseBody
}

type ListHandshakesForResourceDirectoryResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHandshakesForResourceDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHandshakesForResourceDirectoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHandshakesForResourceDirectoryResponse) GoString() string {
	return s.String()
}

func (s *ListHandshakesForResourceDirectoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHandshakesForResourceDirectoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHandshakesForResourceDirectoryResponse) GetBody() *ListHandshakesForResourceDirectoryResponseBody {
	return s.Body
}

func (s *ListHandshakesForResourceDirectoryResponse) SetHeaders(v map[string]*string) *ListHandshakesForResourceDirectoryResponse {
	s.Headers = v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponse) SetStatusCode(v int32) *ListHandshakesForResourceDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponse) SetBody(v *ListHandshakesForResourceDirectoryResponseBody) *ListHandshakesForResourceDirectoryResponse {
	s.Body = v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponse) Validate() error {
	return dara.Validate(s)
}
