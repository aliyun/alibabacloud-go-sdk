// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConversionDataIntlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConversionDataIntlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConversionDataIntlResponse
	GetStatusCode() *int32
	SetBody(v *ConversionDataIntlResponseBody) *ConversionDataIntlResponse
	GetBody() *ConversionDataIntlResponseBody
}

type ConversionDataIntlResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConversionDataIntlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConversionDataIntlResponse) String() string {
	return dara.Prettify(s)
}

func (s ConversionDataIntlResponse) GoString() string {
	return s.String()
}

func (s *ConversionDataIntlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConversionDataIntlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConversionDataIntlResponse) GetBody() *ConversionDataIntlResponseBody {
	return s.Body
}

func (s *ConversionDataIntlResponse) SetHeaders(v map[string]*string) *ConversionDataIntlResponse {
	s.Headers = v
	return s
}

func (s *ConversionDataIntlResponse) SetStatusCode(v int32) *ConversionDataIntlResponse {
	s.StatusCode = &v
	return s
}

func (s *ConversionDataIntlResponse) SetBody(v *ConversionDataIntlResponseBody) *ConversionDataIntlResponse {
	s.Body = v
	return s
}

func (s *ConversionDataIntlResponse) Validate() error {
	return dara.Validate(s)
}
