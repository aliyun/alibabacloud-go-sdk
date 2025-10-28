// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHealthCheckUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateHealthCheckUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateHealthCheckUrlResponse
	GetStatusCode() *int32
	SetBody(v *UpdateHealthCheckUrlResponseBody) *UpdateHealthCheckUrlResponse
	GetBody() *UpdateHealthCheckUrlResponseBody
}

type UpdateHealthCheckUrlResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHealthCheckUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHealthCheckUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateHealthCheckUrlResponse) GoString() string {
	return s.String()
}

func (s *UpdateHealthCheckUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateHealthCheckUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateHealthCheckUrlResponse) GetBody() *UpdateHealthCheckUrlResponseBody {
	return s.Body
}

func (s *UpdateHealthCheckUrlResponse) SetHeaders(v map[string]*string) *UpdateHealthCheckUrlResponse {
	s.Headers = v
	return s
}

func (s *UpdateHealthCheckUrlResponse) SetStatusCode(v int32) *UpdateHealthCheckUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHealthCheckUrlResponse) SetBody(v *UpdateHealthCheckUrlResponseBody) *UpdateHealthCheckUrlResponse {
	s.Body = v
	return s
}

func (s *UpdateHealthCheckUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
