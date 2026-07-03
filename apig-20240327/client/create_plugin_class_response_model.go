// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePluginClassResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePluginClassResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePluginClassResponse
	GetStatusCode() *int32
	SetBody(v *CreatePluginClassResponseBody) *CreatePluginClassResponse
	GetBody() *CreatePluginClassResponseBody
}

type CreatePluginClassResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePluginClassResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePluginClassResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePluginClassResponse) GoString() string {
	return s.String()
}

func (s *CreatePluginClassResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePluginClassResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePluginClassResponse) GetBody() *CreatePluginClassResponseBody {
	return s.Body
}

func (s *CreatePluginClassResponse) SetHeaders(v map[string]*string) *CreatePluginClassResponse {
	s.Headers = v
	return s
}

func (s *CreatePluginClassResponse) SetStatusCode(v int32) *CreatePluginClassResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePluginClassResponse) SetBody(v *CreatePluginClassResponseBody) *CreatePluginClassResponse {
	s.Body = v
	return s
}

func (s *CreatePluginClassResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
