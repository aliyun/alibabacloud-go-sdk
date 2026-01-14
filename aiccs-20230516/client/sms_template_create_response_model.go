// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSmsTemplateCreateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SmsTemplateCreateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SmsTemplateCreateResponse
	GetStatusCode() *int32
	SetBody(v *SmsTemplateCreateResponseBody) *SmsTemplateCreateResponse
	GetBody() *SmsTemplateCreateResponseBody
}

type SmsTemplateCreateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SmsTemplateCreateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SmsTemplateCreateResponse) String() string {
	return dara.Prettify(s)
}

func (s SmsTemplateCreateResponse) GoString() string {
	return s.String()
}

func (s *SmsTemplateCreateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SmsTemplateCreateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SmsTemplateCreateResponse) GetBody() *SmsTemplateCreateResponseBody {
	return s.Body
}

func (s *SmsTemplateCreateResponse) SetHeaders(v map[string]*string) *SmsTemplateCreateResponse {
	s.Headers = v
	return s
}

func (s *SmsTemplateCreateResponse) SetStatusCode(v int32) *SmsTemplateCreateResponse {
	s.StatusCode = &v
	return s
}

func (s *SmsTemplateCreateResponse) SetBody(v *SmsTemplateCreateResponseBody) *SmsTemplateCreateResponse {
	s.Body = v
	return s
}

func (s *SmsTemplateCreateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
