// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHealthCheckTemplateAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHealthCheckTemplateAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHealthCheckTemplateAttributeResponse
	GetStatusCode() *int32
	SetBody(v *GetHealthCheckTemplateAttributeResponseBody) *GetHealthCheckTemplateAttributeResponse
	GetBody() *GetHealthCheckTemplateAttributeResponseBody
}

type GetHealthCheckTemplateAttributeResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHealthCheckTemplateAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHealthCheckTemplateAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHealthCheckTemplateAttributeResponse) GoString() string {
	return s.String()
}

func (s *GetHealthCheckTemplateAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHealthCheckTemplateAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHealthCheckTemplateAttributeResponse) GetBody() *GetHealthCheckTemplateAttributeResponseBody {
	return s.Body
}

func (s *GetHealthCheckTemplateAttributeResponse) SetHeaders(v map[string]*string) *GetHealthCheckTemplateAttributeResponse {
	s.Headers = v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponse) SetStatusCode(v int32) *GetHealthCheckTemplateAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponse) SetBody(v *GetHealthCheckTemplateAttributeResponseBody) *GetHealthCheckTemplateAttributeResponse {
	s.Body = v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponse) Validate() error {
	return dara.Validate(s)
}
