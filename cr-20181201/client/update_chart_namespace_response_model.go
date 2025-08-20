// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateChartNamespaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateChartNamespaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateChartNamespaceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateChartNamespaceResponseBody) *UpdateChartNamespaceResponse
	GetBody() *UpdateChartNamespaceResponseBody
}

type UpdateChartNamespaceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateChartNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateChartNamespaceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateChartNamespaceResponse) GoString() string {
	return s.String()
}

func (s *UpdateChartNamespaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateChartNamespaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateChartNamespaceResponse) GetBody() *UpdateChartNamespaceResponseBody {
	return s.Body
}

func (s *UpdateChartNamespaceResponse) SetHeaders(v map[string]*string) *UpdateChartNamespaceResponse {
	s.Headers = v
	return s
}

func (s *UpdateChartNamespaceResponse) SetStatusCode(v int32) *UpdateChartNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateChartNamespaceResponse) SetBody(v *UpdateChartNamespaceResponseBody) *UpdateChartNamespaceResponse {
	s.Body = v
	return s
}

func (s *UpdateChartNamespaceResponse) Validate() error {
	return dara.Validate(s)
}
