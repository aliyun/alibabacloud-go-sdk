// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigLayer4RuleBakModeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigLayer4RuleBakModeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigLayer4RuleBakModeResponse
	GetStatusCode() *int32
	SetBody(v *ConfigLayer4RuleBakModeResponseBody) *ConfigLayer4RuleBakModeResponse
	GetBody() *ConfigLayer4RuleBakModeResponseBody
}

type ConfigLayer4RuleBakModeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigLayer4RuleBakModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigLayer4RuleBakModeResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigLayer4RuleBakModeResponse) GoString() string {
	return s.String()
}

func (s *ConfigLayer4RuleBakModeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigLayer4RuleBakModeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigLayer4RuleBakModeResponse) GetBody() *ConfigLayer4RuleBakModeResponseBody {
	return s.Body
}

func (s *ConfigLayer4RuleBakModeResponse) SetHeaders(v map[string]*string) *ConfigLayer4RuleBakModeResponse {
	s.Headers = v
	return s
}

func (s *ConfigLayer4RuleBakModeResponse) SetStatusCode(v int32) *ConfigLayer4RuleBakModeResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigLayer4RuleBakModeResponse) SetBody(v *ConfigLayer4RuleBakModeResponseBody) *ConfigLayer4RuleBakModeResponse {
	s.Body = v
	return s
}

func (s *ConfigLayer4RuleBakModeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
