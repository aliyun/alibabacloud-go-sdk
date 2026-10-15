// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTextTranslateAsyncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TextTranslateAsyncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TextTranslateAsyncResponse
	GetStatusCode() *int32
	SetBody(v *TextTranslateAsyncResponseBody) *TextTranslateAsyncResponse
	GetBody() *TextTranslateAsyncResponseBody
}

type TextTranslateAsyncResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TextTranslateAsyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TextTranslateAsyncResponse) String() string {
	return dara.Prettify(s)
}

func (s TextTranslateAsyncResponse) GoString() string {
	return s.String()
}

func (s *TextTranslateAsyncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TextTranslateAsyncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TextTranslateAsyncResponse) GetBody() *TextTranslateAsyncResponseBody {
	return s.Body
}

func (s *TextTranslateAsyncResponse) SetHeaders(v map[string]*string) *TextTranslateAsyncResponse {
	s.Headers = v
	return s
}

func (s *TextTranslateAsyncResponse) SetStatusCode(v int32) *TextTranslateAsyncResponse {
	s.StatusCode = &v
	return s
}

func (s *TextTranslateAsyncResponse) SetBody(v *TextTranslateAsyncResponseBody) *TextTranslateAsyncResponse {
	s.Body = v
	return s
}

func (s *TextTranslateAsyncResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
