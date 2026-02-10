// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDigitalSmsTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDigitalSmsTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDigitalSmsTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateDigitalSmsTemplateResponseBody) *CreateDigitalSmsTemplateResponse
	GetBody() *CreateDigitalSmsTemplateResponseBody
}

type CreateDigitalSmsTemplateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDigitalSmsTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDigitalSmsTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDigitalSmsTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateDigitalSmsTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDigitalSmsTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDigitalSmsTemplateResponse) GetBody() *CreateDigitalSmsTemplateResponseBody {
	return s.Body
}

func (s *CreateDigitalSmsTemplateResponse) SetHeaders(v map[string]*string) *CreateDigitalSmsTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateDigitalSmsTemplateResponse) SetStatusCode(v int32) *CreateDigitalSmsTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDigitalSmsTemplateResponse) SetBody(v *CreateDigitalSmsTemplateResponseBody) *CreateDigitalSmsTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateDigitalSmsTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
