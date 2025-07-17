// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSchemasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSchemasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSchemasResponse
	GetStatusCode() *int32
	SetBody(v *ListSchemasResponseBody) *ListSchemasResponse
	GetBody() *ListSchemasResponseBody
}

type ListSchemasResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSchemasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSchemasResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSchemasResponse) GoString() string {
	return s.String()
}

func (s *ListSchemasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSchemasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSchemasResponse) GetBody() *ListSchemasResponseBody {
	return s.Body
}

func (s *ListSchemasResponse) SetHeaders(v map[string]*string) *ListSchemasResponse {
	s.Headers = v
	return s
}

func (s *ListSchemasResponse) SetStatusCode(v int32) *ListSchemasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSchemasResponse) SetBody(v *ListSchemasResponseBody) *ListSchemasResponse {
	s.Body = v
	return s
}

func (s *ListSchemasResponse) Validate() error {
	return dara.Validate(s)
}
