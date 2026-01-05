// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPluginClassesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPluginClassesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPluginClassesResponse
	GetStatusCode() *int32
	SetBody(v *ListPluginClassesResponseBody) *ListPluginClassesResponse
	GetBody() *ListPluginClassesResponseBody
}

type ListPluginClassesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPluginClassesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPluginClassesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPluginClassesResponse) GoString() string {
	return s.String()
}

func (s *ListPluginClassesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPluginClassesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPluginClassesResponse) GetBody() *ListPluginClassesResponseBody {
	return s.Body
}

func (s *ListPluginClassesResponse) SetHeaders(v map[string]*string) *ListPluginClassesResponse {
	s.Headers = v
	return s
}

func (s *ListPluginClassesResponse) SetStatusCode(v int32) *ListPluginClassesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPluginClassesResponse) SetBody(v *ListPluginClassesResponseBody) *ListPluginClassesResponse {
	s.Body = v
	return s
}

func (s *ListPluginClassesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
