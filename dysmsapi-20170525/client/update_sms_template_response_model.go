// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmsTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSmsTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSmsTemplateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSmsTemplateResponseBody) *UpdateSmsTemplateResponse
	GetBody() *UpdateSmsTemplateResponseBody
}

type UpdateSmsTemplateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSmsTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSmsTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmsTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateSmsTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSmsTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSmsTemplateResponse) GetBody() *UpdateSmsTemplateResponseBody {
	return s.Body
}

func (s *UpdateSmsTemplateResponse) SetHeaders(v map[string]*string) *UpdateSmsTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateSmsTemplateResponse) SetStatusCode(v int32) *UpdateSmsTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSmsTemplateResponse) SetBody(v *UpdateSmsTemplateResponseBody) *UpdateSmsTemplateResponse {
	s.Body = v
	return s
}

func (s *UpdateSmsTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
