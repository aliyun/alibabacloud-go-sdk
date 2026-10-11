// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEgressProxyConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEgressProxyConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEgressProxyConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *CreateEgressProxyConfigurationResponseBody) *CreateEgressProxyConfigurationResponse
	GetBody() *CreateEgressProxyConfigurationResponseBody
}

type CreateEgressProxyConfigurationResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEgressProxyConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEgressProxyConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyConfigurationResponse) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEgressProxyConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEgressProxyConfigurationResponse) GetBody() *CreateEgressProxyConfigurationResponseBody {
	return s.Body
}

func (s *CreateEgressProxyConfigurationResponse) SetHeaders(v map[string]*string) *CreateEgressProxyConfigurationResponse {
	s.Headers = v
	return s
}

func (s *CreateEgressProxyConfigurationResponse) SetStatusCode(v int32) *CreateEgressProxyConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEgressProxyConfigurationResponse) SetBody(v *CreateEgressProxyConfigurationResponseBody) *CreateEgressProxyConfigurationResponse {
	s.Body = v
	return s
}

func (s *CreateEgressProxyConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
