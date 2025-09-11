// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSmsTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddSmsTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddSmsTemplateResponse
	GetStatusCode() *int32
	SetBody(v *AddSmsTemplateResponseBody) *AddSmsTemplateResponse
	GetBody() *AddSmsTemplateResponseBody
}

type AddSmsTemplateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddSmsTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddSmsTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s AddSmsTemplateResponse) GoString() string {
	return s.String()
}

func (s *AddSmsTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddSmsTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddSmsTemplateResponse) GetBody() *AddSmsTemplateResponseBody {
	return s.Body
}

func (s *AddSmsTemplateResponse) SetHeaders(v map[string]*string) *AddSmsTemplateResponse {
	s.Headers = v
	return s
}

func (s *AddSmsTemplateResponse) SetStatusCode(v int32) *AddSmsTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *AddSmsTemplateResponse) SetBody(v *AddSmsTemplateResponseBody) *AddSmsTemplateResponse {
	s.Body = v
	return s
}

func (s *AddSmsTemplateResponse) Validate() error {
	return dara.Validate(s)
}
