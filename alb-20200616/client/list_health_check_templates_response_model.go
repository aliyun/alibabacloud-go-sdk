// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHealthCheckTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHealthCheckTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHealthCheckTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListHealthCheckTemplatesResponseBody) *ListHealthCheckTemplatesResponse
	GetBody() *ListHealthCheckTemplatesResponseBody
}

type ListHealthCheckTemplatesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHealthCheckTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHealthCheckTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHealthCheckTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListHealthCheckTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHealthCheckTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHealthCheckTemplatesResponse) GetBody() *ListHealthCheckTemplatesResponseBody {
	return s.Body
}

func (s *ListHealthCheckTemplatesResponse) SetHeaders(v map[string]*string) *ListHealthCheckTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListHealthCheckTemplatesResponse) SetStatusCode(v int32) *ListHealthCheckTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHealthCheckTemplatesResponse) SetBody(v *ListHealthCheckTemplatesResponseBody) *ListHealthCheckTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListHealthCheckTemplatesResponse) Validate() error {
	return dara.Validate(s)
}
