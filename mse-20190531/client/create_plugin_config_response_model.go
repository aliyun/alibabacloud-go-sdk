// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePluginConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePluginConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePluginConfigResponse
	GetStatusCode() *int32
	SetBody(v *CreatePluginConfigResponseBody) *CreatePluginConfigResponse
	GetBody() *CreatePluginConfigResponseBody
}

type CreatePluginConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePluginConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePluginConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePluginConfigResponse) GoString() string {
	return s.String()
}

func (s *CreatePluginConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePluginConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePluginConfigResponse) GetBody() *CreatePluginConfigResponseBody {
	return s.Body
}

func (s *CreatePluginConfigResponse) SetHeaders(v map[string]*string) *CreatePluginConfigResponse {
	s.Headers = v
	return s
}

func (s *CreatePluginConfigResponse) SetStatusCode(v int32) *CreatePluginConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePluginConfigResponse) SetBody(v *CreatePluginConfigResponseBody) *CreatePluginConfigResponse {
	s.Body = v
	return s
}

func (s *CreatePluginConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
