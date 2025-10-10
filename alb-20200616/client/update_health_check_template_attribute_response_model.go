// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHealthCheckTemplateAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateHealthCheckTemplateAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateHealthCheckTemplateAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateHealthCheckTemplateAttributeResponseBody) *UpdateHealthCheckTemplateAttributeResponse
	GetBody() *UpdateHealthCheckTemplateAttributeResponseBody
}

type UpdateHealthCheckTemplateAttributeResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHealthCheckTemplateAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHealthCheckTemplateAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateHealthCheckTemplateAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateHealthCheckTemplateAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateHealthCheckTemplateAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateHealthCheckTemplateAttributeResponse) GetBody() *UpdateHealthCheckTemplateAttributeResponseBody {
	return s.Body
}

func (s *UpdateHealthCheckTemplateAttributeResponse) SetHeaders(v map[string]*string) *UpdateHealthCheckTemplateAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeResponse) SetStatusCode(v int32) *UpdateHealthCheckTemplateAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeResponse) SetBody(v *UpdateHealthCheckTemplateAttributeResponseBody) *UpdateHealthCheckTemplateAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeResponse) Validate() error {
	return dara.Validate(s)
}
