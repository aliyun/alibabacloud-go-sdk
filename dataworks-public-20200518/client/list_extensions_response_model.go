// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExtensionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListExtensionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListExtensionsResponse
	GetStatusCode() *int32
	SetBody(v *ListExtensionsResponseBody) *ListExtensionsResponse
	GetBody() *ListExtensionsResponseBody
}

type ListExtensionsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExtensionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExtensionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListExtensionsResponse) GoString() string {
	return s.String()
}

func (s *ListExtensionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListExtensionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListExtensionsResponse) GetBody() *ListExtensionsResponseBody {
	return s.Body
}

func (s *ListExtensionsResponse) SetHeaders(v map[string]*string) *ListExtensionsResponse {
	s.Headers = v
	return s
}

func (s *ListExtensionsResponse) SetStatusCode(v int32) *ListExtensionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExtensionsResponse) SetBody(v *ListExtensionsResponseBody) *ListExtensionsResponse {
	s.Body = v
	return s
}

func (s *ListExtensionsResponse) Validate() error {
	return dara.Validate(s)
}
