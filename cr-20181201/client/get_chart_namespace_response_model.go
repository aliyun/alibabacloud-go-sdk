// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChartNamespaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetChartNamespaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetChartNamespaceResponse
	GetStatusCode() *int32
	SetBody(v *GetChartNamespaceResponseBody) *GetChartNamespaceResponse
	GetBody() *GetChartNamespaceResponseBody
}

type GetChartNamespaceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChartNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChartNamespaceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetChartNamespaceResponse) GoString() string {
	return s.String()
}

func (s *GetChartNamespaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetChartNamespaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetChartNamespaceResponse) GetBody() *GetChartNamespaceResponseBody {
	return s.Body
}

func (s *GetChartNamespaceResponse) SetHeaders(v map[string]*string) *GetChartNamespaceResponse {
	s.Headers = v
	return s
}

func (s *GetChartNamespaceResponse) SetStatusCode(v int32) *GetChartNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChartNamespaceResponse) SetBody(v *GetChartNamespaceResponseBody) *GetChartNamespaceResponse {
	s.Body = v
	return s
}

func (s *GetChartNamespaceResponse) Validate() error {
	return dara.Validate(s)
}
