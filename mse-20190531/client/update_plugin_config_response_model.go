// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePluginConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePluginConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePluginConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePluginConfigResponseBody) *UpdatePluginConfigResponse
	GetBody() *UpdatePluginConfigResponseBody
}

type UpdatePluginConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePluginConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePluginConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePluginConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdatePluginConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePluginConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePluginConfigResponse) GetBody() *UpdatePluginConfigResponseBody {
	return s.Body
}

func (s *UpdatePluginConfigResponse) SetHeaders(v map[string]*string) *UpdatePluginConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdatePluginConfigResponse) SetStatusCode(v int32) *UpdatePluginConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePluginConfigResponse) SetBody(v *UpdatePluginConfigResponseBody) *UpdatePluginConfigResponse {
	s.Body = v
	return s
}

func (s *UpdatePluginConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
