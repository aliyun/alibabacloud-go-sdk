// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEgressProxyConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEgressProxyConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEgressProxyConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEgressProxyConfigurationResponseBody) *DeleteEgressProxyConfigurationResponse
	GetBody() *DeleteEgressProxyConfigurationResponseBody
}

type DeleteEgressProxyConfigurationResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEgressProxyConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEgressProxyConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEgressProxyConfigurationResponse) GoString() string {
	return s.String()
}

func (s *DeleteEgressProxyConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEgressProxyConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEgressProxyConfigurationResponse) GetBody() *DeleteEgressProxyConfigurationResponseBody {
	return s.Body
}

func (s *DeleteEgressProxyConfigurationResponse) SetHeaders(v map[string]*string) *DeleteEgressProxyConfigurationResponse {
	s.Headers = v
	return s
}

func (s *DeleteEgressProxyConfigurationResponse) SetStatusCode(v int32) *DeleteEgressProxyConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEgressProxyConfigurationResponse) SetBody(v *DeleteEgressProxyConfigurationResponseBody) *DeleteEgressProxyConfigurationResponse {
	s.Body = v
	return s
}

func (s *DeleteEgressProxyConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
