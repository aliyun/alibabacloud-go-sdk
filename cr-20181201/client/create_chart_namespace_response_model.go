// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChartNamespaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateChartNamespaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateChartNamespaceResponse
	GetStatusCode() *int32
	SetBody(v *CreateChartNamespaceResponseBody) *CreateChartNamespaceResponse
	GetBody() *CreateChartNamespaceResponseBody
}

type CreateChartNamespaceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateChartNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateChartNamespaceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateChartNamespaceResponse) GoString() string {
	return s.String()
}

func (s *CreateChartNamespaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateChartNamespaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateChartNamespaceResponse) GetBody() *CreateChartNamespaceResponseBody {
	return s.Body
}

func (s *CreateChartNamespaceResponse) SetHeaders(v map[string]*string) *CreateChartNamespaceResponse {
	s.Headers = v
	return s
}

func (s *CreateChartNamespaceResponse) SetStatusCode(v int32) *CreateChartNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateChartNamespaceResponse) SetBody(v *CreateChartNamespaceResponseBody) *CreateChartNamespaceResponse {
	s.Body = v
	return s
}

func (s *CreateChartNamespaceResponse) Validate() error {
	return dara.Validate(s)
}
