// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigEndpointProbeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigEndpointProbeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigEndpointProbeResponse
	GetStatusCode() *int32
	SetBody(v *ConfigEndpointProbeResponseBody) *ConfigEndpointProbeResponse
	GetBody() *ConfigEndpointProbeResponseBody
}

type ConfigEndpointProbeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigEndpointProbeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigEndpointProbeResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigEndpointProbeResponse) GoString() string {
	return s.String()
}

func (s *ConfigEndpointProbeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigEndpointProbeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigEndpointProbeResponse) GetBody() *ConfigEndpointProbeResponseBody {
	return s.Body
}

func (s *ConfigEndpointProbeResponse) SetHeaders(v map[string]*string) *ConfigEndpointProbeResponse {
	s.Headers = v
	return s
}

func (s *ConfigEndpointProbeResponse) SetStatusCode(v int32) *ConfigEndpointProbeResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigEndpointProbeResponse) SetBody(v *ConfigEndpointProbeResponseBody) *ConfigEndpointProbeResponse {
	s.Body = v
	return s
}

func (s *ConfigEndpointProbeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
