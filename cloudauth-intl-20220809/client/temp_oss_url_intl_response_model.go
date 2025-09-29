// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTempOssUrlIntlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TempOssUrlIntlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TempOssUrlIntlResponse
	GetStatusCode() *int32
	SetBody(v *TempOssUrlIntlResponseBody) *TempOssUrlIntlResponse
	GetBody() *TempOssUrlIntlResponseBody
}

type TempOssUrlIntlResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TempOssUrlIntlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TempOssUrlIntlResponse) String() string {
	return dara.Prettify(s)
}

func (s TempOssUrlIntlResponse) GoString() string {
	return s.String()
}

func (s *TempOssUrlIntlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TempOssUrlIntlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TempOssUrlIntlResponse) GetBody() *TempOssUrlIntlResponseBody {
	return s.Body
}

func (s *TempOssUrlIntlResponse) SetHeaders(v map[string]*string) *TempOssUrlIntlResponse {
	s.Headers = v
	return s
}

func (s *TempOssUrlIntlResponse) SetStatusCode(v int32) *TempOssUrlIntlResponse {
	s.StatusCode = &v
	return s
}

func (s *TempOssUrlIntlResponse) SetBody(v *TempOssUrlIntlResponseBody) *TempOssUrlIntlResponse {
	s.Body = v
	return s
}

func (s *TempOssUrlIntlResponse) Validate() error {
	return dara.Validate(s)
}
