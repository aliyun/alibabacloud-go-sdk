// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigWebCCTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigWebCCTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigWebCCTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ConfigWebCCTemplateResponseBody) *ConfigWebCCTemplateResponse
	GetBody() *ConfigWebCCTemplateResponseBody
}

type ConfigWebCCTemplateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigWebCCTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigWebCCTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigWebCCTemplateResponse) GoString() string {
	return s.String()
}

func (s *ConfigWebCCTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigWebCCTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigWebCCTemplateResponse) GetBody() *ConfigWebCCTemplateResponseBody {
	return s.Body
}

func (s *ConfigWebCCTemplateResponse) SetHeaders(v map[string]*string) *ConfigWebCCTemplateResponse {
	s.Headers = v
	return s
}

func (s *ConfigWebCCTemplateResponse) SetStatusCode(v int32) *ConfigWebCCTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigWebCCTemplateResponse) SetBody(v *ConfigWebCCTemplateResponseBody) *ConfigWebCCTemplateResponse {
	s.Body = v
	return s
}

func (s *ConfigWebCCTemplateResponse) Validate() error {
	return dara.Validate(s)
}
