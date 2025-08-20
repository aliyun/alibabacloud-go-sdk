// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChartNamespaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteChartNamespaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteChartNamespaceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteChartNamespaceResponseBody) *DeleteChartNamespaceResponse
	GetBody() *DeleteChartNamespaceResponseBody
}

type DeleteChartNamespaceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteChartNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteChartNamespaceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteChartNamespaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteChartNamespaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteChartNamespaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteChartNamespaceResponse) GetBody() *DeleteChartNamespaceResponseBody {
	return s.Body
}

func (s *DeleteChartNamespaceResponse) SetHeaders(v map[string]*string) *DeleteChartNamespaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteChartNamespaceResponse) SetStatusCode(v int32) *DeleteChartNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteChartNamespaceResponse) SetBody(v *DeleteChartNamespaceResponseBody) *DeleteChartNamespaceResponse {
	s.Body = v
	return s
}

func (s *DeleteChartNamespaceResponse) Validate() error {
	return dara.Validate(s)
}
