// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSmsTemplatePageListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SmsTemplatePageListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SmsTemplatePageListResponse
	GetStatusCode() *int32
	SetBody(v *SmsTemplatePageListResponseBody) *SmsTemplatePageListResponse
	GetBody() *SmsTemplatePageListResponseBody
}

type SmsTemplatePageListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SmsTemplatePageListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SmsTemplatePageListResponse) String() string {
	return dara.Prettify(s)
}

func (s SmsTemplatePageListResponse) GoString() string {
	return s.String()
}

func (s *SmsTemplatePageListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SmsTemplatePageListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SmsTemplatePageListResponse) GetBody() *SmsTemplatePageListResponseBody {
	return s.Body
}

func (s *SmsTemplatePageListResponse) SetHeaders(v map[string]*string) *SmsTemplatePageListResponse {
	s.Headers = v
	return s
}

func (s *SmsTemplatePageListResponse) SetStatusCode(v int32) *SmsTemplatePageListResponse {
	s.StatusCode = &v
	return s
}

func (s *SmsTemplatePageListResponse) SetBody(v *SmsTemplatePageListResponseBody) *SmsTemplatePageListResponse {
	s.Body = v
	return s
}

func (s *SmsTemplatePageListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
