// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHealthCheckTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHealthCheckTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHealthCheckTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHealthCheckTemplatesResponseBody) *DeleteHealthCheckTemplatesResponse
	GetBody() *DeleteHealthCheckTemplatesResponseBody
}

type DeleteHealthCheckTemplatesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHealthCheckTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHealthCheckTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHealthCheckTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DeleteHealthCheckTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHealthCheckTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHealthCheckTemplatesResponse) GetBody() *DeleteHealthCheckTemplatesResponseBody {
	return s.Body
}

func (s *DeleteHealthCheckTemplatesResponse) SetHeaders(v map[string]*string) *DeleteHealthCheckTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DeleteHealthCheckTemplatesResponse) SetStatusCode(v int32) *DeleteHealthCheckTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHealthCheckTemplatesResponse) SetBody(v *DeleteHealthCheckTemplatesResponseBody) *DeleteHealthCheckTemplatesResponse {
	s.Body = v
	return s
}

func (s *DeleteHealthCheckTemplatesResponse) Validate() error {
	return dara.Validate(s)
}
