// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIndexTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIndexTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIndexTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListIndexTemplatesResponseBody) *ListIndexTemplatesResponse
	GetBody() *ListIndexTemplatesResponseBody
}

type ListIndexTemplatesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIndexTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIndexTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIndexTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListIndexTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIndexTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIndexTemplatesResponse) GetBody() *ListIndexTemplatesResponseBody {
	return s.Body
}

func (s *ListIndexTemplatesResponse) SetHeaders(v map[string]*string) *ListIndexTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListIndexTemplatesResponse) SetStatusCode(v int32) *ListIndexTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIndexTemplatesResponse) SetBody(v *ListIndexTemplatesResponseBody) *ListIndexTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListIndexTemplatesResponse) Validate() error {
	return dara.Validate(s)
}
