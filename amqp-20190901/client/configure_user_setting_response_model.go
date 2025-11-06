// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureUserSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigureUserSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigureUserSettingResponse
	GetStatusCode() *int32
	SetBody(v *ConfigureUserSettingResponseBody) *ConfigureUserSettingResponse
	GetBody() *ConfigureUserSettingResponseBody
}

type ConfigureUserSettingResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigureUserSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigureUserSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigureUserSettingResponse) GoString() string {
	return s.String()
}

func (s *ConfigureUserSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigureUserSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigureUserSettingResponse) GetBody() *ConfigureUserSettingResponseBody {
	return s.Body
}

func (s *ConfigureUserSettingResponse) SetHeaders(v map[string]*string) *ConfigureUserSettingResponse {
	s.Headers = v
	return s
}

func (s *ConfigureUserSettingResponse) SetStatusCode(v int32) *ConfigureUserSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigureUserSettingResponse) SetBody(v *ConfigureUserSettingResponseBody) *ConfigureUserSettingResponse {
	s.Body = v
	return s
}

func (s *ConfigureUserSettingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
