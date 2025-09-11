// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSmsTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSmsTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSmsTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateSmsTemplateResponseBody) *CreateSmsTemplateResponse
	GetBody() *CreateSmsTemplateResponseBody
}

type CreateSmsTemplateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSmsTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSmsTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSmsTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateSmsTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSmsTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSmsTemplateResponse) GetBody() *CreateSmsTemplateResponseBody {
	return s.Body
}

func (s *CreateSmsTemplateResponse) SetHeaders(v map[string]*string) *CreateSmsTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateSmsTemplateResponse) SetStatusCode(v int32) *CreateSmsTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSmsTemplateResponse) SetBody(v *CreateSmsTemplateResponseBody) *CreateSmsTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateSmsTemplateResponse) Validate() error {
	return dara.Validate(s)
}
