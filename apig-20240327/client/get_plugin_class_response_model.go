// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPluginClassResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPluginClassResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPluginClassResponse
	GetStatusCode() *int32
	SetBody(v *GetPluginClassResponseBody) *GetPluginClassResponse
	GetBody() *GetPluginClassResponseBody
}

type GetPluginClassResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPluginClassResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPluginClassResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPluginClassResponse) GoString() string {
	return s.String()
}

func (s *GetPluginClassResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPluginClassResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPluginClassResponse) GetBody() *GetPluginClassResponseBody {
	return s.Body
}

func (s *GetPluginClassResponse) SetHeaders(v map[string]*string) *GetPluginClassResponse {
	s.Headers = v
	return s
}

func (s *GetPluginClassResponse) SetStatusCode(v int32) *GetPluginClassResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPluginClassResponse) SetBody(v *GetPluginClassResponseBody) *GetPluginClassResponse {
	s.Body = v
	return s
}

func (s *GetPluginClassResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
