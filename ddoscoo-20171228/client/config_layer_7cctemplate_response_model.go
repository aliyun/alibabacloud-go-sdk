// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigLayer7CCTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigLayer7CCTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigLayer7CCTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ConfigLayer7CCTemplateResponseBody) *ConfigLayer7CCTemplateResponse
	GetBody() *ConfigLayer7CCTemplateResponseBody
}

type ConfigLayer7CCTemplateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigLayer7CCTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigLayer7CCTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigLayer7CCTemplateResponse) GoString() string {
	return s.String()
}

func (s *ConfigLayer7CCTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigLayer7CCTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigLayer7CCTemplateResponse) GetBody() *ConfigLayer7CCTemplateResponseBody {
	return s.Body
}

func (s *ConfigLayer7CCTemplateResponse) SetHeaders(v map[string]*string) *ConfigLayer7CCTemplateResponse {
	s.Headers = v
	return s
}

func (s *ConfigLayer7CCTemplateResponse) SetStatusCode(v int32) *ConfigLayer7CCTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigLayer7CCTemplateResponse) SetBody(v *ConfigLayer7CCTemplateResponseBody) *ConfigLayer7CCTemplateResponse {
	s.Body = v
	return s
}

func (s *ConfigLayer7CCTemplateResponse) Validate() error {
	return dara.Validate(s)
}
