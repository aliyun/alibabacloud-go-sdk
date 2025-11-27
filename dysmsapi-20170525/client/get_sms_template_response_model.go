// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSmsTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSmsTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSmsTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetSmsTemplateResponseBody) *GetSmsTemplateResponse
	GetBody() *GetSmsTemplateResponseBody
}

type GetSmsTemplateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSmsTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSmsTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSmsTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetSmsTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSmsTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSmsTemplateResponse) GetBody() *GetSmsTemplateResponseBody {
	return s.Body
}

func (s *GetSmsTemplateResponse) SetHeaders(v map[string]*string) *GetSmsTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetSmsTemplateResponse) SetStatusCode(v int32) *GetSmsTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSmsTemplateResponse) SetBody(v *GetSmsTemplateResponseBody) *GetSmsTemplateResponse {
	s.Body = v
	return s
}

func (s *GetSmsTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
