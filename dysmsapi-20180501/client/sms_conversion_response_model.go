// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSmsConversionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SmsConversionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SmsConversionResponse
	GetStatusCode() *int32
	SetBody(v *SmsConversionResponseBody) *SmsConversionResponse
	GetBody() *SmsConversionResponseBody
}

type SmsConversionResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SmsConversionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SmsConversionResponse) String() string {
	return dara.Prettify(s)
}

func (s SmsConversionResponse) GoString() string {
	return s.String()
}

func (s *SmsConversionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SmsConversionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SmsConversionResponse) GetBody() *SmsConversionResponseBody {
	return s.Body
}

func (s *SmsConversionResponse) SetHeaders(v map[string]*string) *SmsConversionResponse {
	s.Headers = v
	return s
}

func (s *SmsConversionResponse) SetStatusCode(v int32) *SmsConversionResponse {
	s.StatusCode = &v
	return s
}

func (s *SmsConversionResponse) SetBody(v *SmsConversionResponseBody) *SmsConversionResponse {
	s.Body = v
	return s
}

func (s *SmsConversionResponse) Validate() error {
	return dara.Validate(s)
}
