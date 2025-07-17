// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTextModerationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TextModerationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TextModerationResponse
	GetStatusCode() *int32
	SetBody(v *TextModerationResponseBody) *TextModerationResponse
	GetBody() *TextModerationResponseBody
}

type TextModerationResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TextModerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TextModerationResponse) String() string {
	return dara.Prettify(s)
}

func (s TextModerationResponse) GoString() string {
	return s.String()
}

func (s *TextModerationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TextModerationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TextModerationResponse) GetBody() *TextModerationResponseBody {
	return s.Body
}

func (s *TextModerationResponse) SetHeaders(v map[string]*string) *TextModerationResponse {
	s.Headers = v
	return s
}

func (s *TextModerationResponse) SetStatusCode(v int32) *TextModerationResponse {
	s.StatusCode = &v
	return s
}

func (s *TextModerationResponse) SetBody(v *TextModerationResponseBody) *TextModerationResponse {
	s.Body = v
	return s
}

func (s *TextModerationResponse) Validate() error {
	return dara.Validate(s)
}
