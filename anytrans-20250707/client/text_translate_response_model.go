// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTextTranslateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TextTranslateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TextTranslateResponse
	GetStatusCode() *int32
	SetBody(v *TextTranslateResponseBody) *TextTranslateResponse
	GetBody() *TextTranslateResponseBody
}

type TextTranslateResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TextTranslateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TextTranslateResponse) String() string {
	return dara.Prettify(s)
}

func (s TextTranslateResponse) GoString() string {
	return s.String()
}

func (s *TextTranslateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TextTranslateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TextTranslateResponse) GetBody() *TextTranslateResponseBody {
	return s.Body
}

func (s *TextTranslateResponse) SetHeaders(v map[string]*string) *TextTranslateResponse {
	s.Headers = v
	return s
}

func (s *TextTranslateResponse) SetStatusCode(v int32) *TextTranslateResponse {
	s.StatusCode = &v
	return s
}

func (s *TextTranslateResponse) SetBody(v *TextTranslateResponseBody) *TextTranslateResponse {
	s.Body = v
	return s
}

func (s *TextTranslateResponse) Validate() error {
	return dara.Validate(s)
}
