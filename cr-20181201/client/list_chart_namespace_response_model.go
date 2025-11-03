// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChartNamespaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListChartNamespaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListChartNamespaceResponse
	GetStatusCode() *int32
	SetBody(v *ListChartNamespaceResponseBody) *ListChartNamespaceResponse
	GetBody() *ListChartNamespaceResponseBody
}

type ListChartNamespaceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListChartNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListChartNamespaceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListChartNamespaceResponse) GoString() string {
	return s.String()
}

func (s *ListChartNamespaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListChartNamespaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListChartNamespaceResponse) GetBody() *ListChartNamespaceResponseBody {
	return s.Body
}

func (s *ListChartNamespaceResponse) SetHeaders(v map[string]*string) *ListChartNamespaceResponse {
	s.Headers = v
	return s
}

func (s *ListChartNamespaceResponse) SetStatusCode(v int32) *ListChartNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListChartNamespaceResponse) SetBody(v *ListChartNamespaceResponseBody) *ListChartNamespaceResponse {
	s.Body = v
	return s
}

func (s *ListChartNamespaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
