// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigL7UsKeepaliveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigL7UsKeepaliveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigL7UsKeepaliveResponse
	GetStatusCode() *int32
	SetBody(v *ConfigL7UsKeepaliveResponseBody) *ConfigL7UsKeepaliveResponse
	GetBody() *ConfigL7UsKeepaliveResponseBody
}

type ConfigL7UsKeepaliveResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigL7UsKeepaliveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigL7UsKeepaliveResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigL7UsKeepaliveResponse) GoString() string {
	return s.String()
}

func (s *ConfigL7UsKeepaliveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigL7UsKeepaliveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigL7UsKeepaliveResponse) GetBody() *ConfigL7UsKeepaliveResponseBody {
	return s.Body
}

func (s *ConfigL7UsKeepaliveResponse) SetHeaders(v map[string]*string) *ConfigL7UsKeepaliveResponse {
	s.Headers = v
	return s
}

func (s *ConfigL7UsKeepaliveResponse) SetStatusCode(v int32) *ConfigL7UsKeepaliveResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigL7UsKeepaliveResponse) SetBody(v *ConfigL7UsKeepaliveResponseBody) *ConfigL7UsKeepaliveResponse {
	s.Body = v
	return s
}

func (s *ConfigL7UsKeepaliveResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
