// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigWebIpSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigWebIpSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigWebIpSetResponse
	GetStatusCode() *int32
	SetBody(v *ConfigWebIpSetResponseBody) *ConfigWebIpSetResponse
	GetBody() *ConfigWebIpSetResponseBody
}

type ConfigWebIpSetResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigWebIpSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigWebIpSetResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigWebIpSetResponse) GoString() string {
	return s.String()
}

func (s *ConfigWebIpSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigWebIpSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigWebIpSetResponse) GetBody() *ConfigWebIpSetResponseBody {
	return s.Body
}

func (s *ConfigWebIpSetResponse) SetHeaders(v map[string]*string) *ConfigWebIpSetResponse {
	s.Headers = v
	return s
}

func (s *ConfigWebIpSetResponse) SetStatusCode(v int32) *ConfigWebIpSetResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigWebIpSetResponse) SetBody(v *ConfigWebIpSetResponseBody) *ConfigWebIpSetResponse {
	s.Body = v
	return s
}

func (s *ConfigWebIpSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
