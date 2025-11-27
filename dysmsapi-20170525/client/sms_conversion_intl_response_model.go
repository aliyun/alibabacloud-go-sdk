// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSmsConversionIntlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SmsConversionIntlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SmsConversionIntlResponse
	GetStatusCode() *int32
	SetBody(v *SmsConversionIntlResponseBody) *SmsConversionIntlResponse
	GetBody() *SmsConversionIntlResponseBody
}

type SmsConversionIntlResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SmsConversionIntlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SmsConversionIntlResponse) String() string {
	return dara.Prettify(s)
}

func (s SmsConversionIntlResponse) GoString() string {
	return s.String()
}

func (s *SmsConversionIntlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SmsConversionIntlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SmsConversionIntlResponse) GetBody() *SmsConversionIntlResponseBody {
	return s.Body
}

func (s *SmsConversionIntlResponse) SetHeaders(v map[string]*string) *SmsConversionIntlResponse {
	s.Headers = v
	return s
}

func (s *SmsConversionIntlResponse) SetStatusCode(v int32) *SmsConversionIntlResponse {
	s.StatusCode = &v
	return s
}

func (s *SmsConversionIntlResponse) SetBody(v *SmsConversionIntlResponseBody) *SmsConversionIntlResponse {
	s.Body = v
	return s
}

func (s *SmsConversionIntlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
