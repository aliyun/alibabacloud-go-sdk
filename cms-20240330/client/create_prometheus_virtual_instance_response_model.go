// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrometheusVirtualInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePrometheusVirtualInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePrometheusVirtualInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreatePrometheusVirtualInstanceResponseBody) *CreatePrometheusVirtualInstanceResponse
	GetBody() *CreatePrometheusVirtualInstanceResponseBody
}

type CreatePrometheusVirtualInstanceResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePrometheusVirtualInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePrometheusVirtualInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePrometheusVirtualInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreatePrometheusVirtualInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePrometheusVirtualInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePrometheusVirtualInstanceResponse) GetBody() *CreatePrometheusVirtualInstanceResponseBody {
	return s.Body
}

func (s *CreatePrometheusVirtualInstanceResponse) SetHeaders(v map[string]*string) *CreatePrometheusVirtualInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreatePrometheusVirtualInstanceResponse) SetStatusCode(v int32) *CreatePrometheusVirtualInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePrometheusVirtualInstanceResponse) SetBody(v *CreatePrometheusVirtualInstanceResponseBody) *CreatePrometheusVirtualInstanceResponse {
	s.Body = v
	return s
}

func (s *CreatePrometheusVirtualInstanceResponse) Validate() error {
	return dara.Validate(s)
}
