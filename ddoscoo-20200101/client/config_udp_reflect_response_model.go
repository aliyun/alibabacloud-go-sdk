// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigUdpReflectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigUdpReflectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigUdpReflectResponse
	GetStatusCode() *int32
	SetBody(v *ConfigUdpReflectResponseBody) *ConfigUdpReflectResponse
	GetBody() *ConfigUdpReflectResponseBody
}

type ConfigUdpReflectResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigUdpReflectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigUdpReflectResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigUdpReflectResponse) GoString() string {
	return s.String()
}

func (s *ConfigUdpReflectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigUdpReflectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigUdpReflectResponse) GetBody() *ConfigUdpReflectResponseBody {
	return s.Body
}

func (s *ConfigUdpReflectResponse) SetHeaders(v map[string]*string) *ConfigUdpReflectResponse {
	s.Headers = v
	return s
}

func (s *ConfigUdpReflectResponse) SetStatusCode(v int32) *ConfigUdpReflectResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigUdpReflectResponse) SetBody(v *ConfigUdpReflectResponseBody) *ConfigUdpReflectResponse {
	s.Body = v
	return s
}

func (s *ConfigUdpReflectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
