// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePluginClassResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePluginClassResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePluginClassResponse
	GetStatusCode() *int32
	SetBody(v *DeletePluginClassResponseBody) *DeletePluginClassResponse
	GetBody() *DeletePluginClassResponseBody
}

type DeletePluginClassResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePluginClassResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePluginClassResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePluginClassResponse) GoString() string {
	return s.String()
}

func (s *DeletePluginClassResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePluginClassResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePluginClassResponse) GetBody() *DeletePluginClassResponseBody {
	return s.Body
}

func (s *DeletePluginClassResponse) SetHeaders(v map[string]*string) *DeletePluginClassResponse {
	s.Headers = v
	return s
}

func (s *DeletePluginClassResponse) SetStatusCode(v int32) *DeletePluginClassResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePluginClassResponse) SetBody(v *DeletePluginClassResponseBody) *DeletePluginClassResponse {
	s.Body = v
	return s
}

func (s *DeletePluginClassResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
