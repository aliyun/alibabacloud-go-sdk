// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigRuntimeModelTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigRuntimeModelTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigRuntimeModelTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ConfigRuntimeModelTemplateResponseBody) *ConfigRuntimeModelTemplateResponse
	GetBody() *ConfigRuntimeModelTemplateResponseBody
}

type ConfigRuntimeModelTemplateResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigRuntimeModelTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigRuntimeModelTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigRuntimeModelTemplateResponse) GoString() string {
	return s.String()
}

func (s *ConfigRuntimeModelTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigRuntimeModelTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigRuntimeModelTemplateResponse) GetBody() *ConfigRuntimeModelTemplateResponseBody {
	return s.Body
}

func (s *ConfigRuntimeModelTemplateResponse) SetHeaders(v map[string]*string) *ConfigRuntimeModelTemplateResponse {
	s.Headers = v
	return s
}

func (s *ConfigRuntimeModelTemplateResponse) SetStatusCode(v int32) *ConfigRuntimeModelTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigRuntimeModelTemplateResponse) SetBody(v *ConfigRuntimeModelTemplateResponseBody) *ConfigRuntimeModelTemplateResponse {
	s.Body = v
	return s
}

func (s *ConfigRuntimeModelTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
