// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHealthCheckTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHealthCheckTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHealthCheckTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateHealthCheckTemplateResponseBody) *CreateHealthCheckTemplateResponse
	GetBody() *CreateHealthCheckTemplateResponseBody
}

type CreateHealthCheckTemplateResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHealthCheckTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHealthCheckTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHealthCheckTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateHealthCheckTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHealthCheckTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHealthCheckTemplateResponse) GetBody() *CreateHealthCheckTemplateResponseBody {
	return s.Body
}

func (s *CreateHealthCheckTemplateResponse) SetHeaders(v map[string]*string) *CreateHealthCheckTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateHealthCheckTemplateResponse) SetStatusCode(v int32) *CreateHealthCheckTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHealthCheckTemplateResponse) SetBody(v *CreateHealthCheckTemplateResponseBody) *CreateHealthCheckTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateHealthCheckTemplateResponse) Validate() error {
	return dara.Validate(s)
}
