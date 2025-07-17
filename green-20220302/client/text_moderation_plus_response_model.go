// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTextModerationPlusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TextModerationPlusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TextModerationPlusResponse
	GetStatusCode() *int32
	SetBody(v *TextModerationPlusResponseBody) *TextModerationPlusResponse
	GetBody() *TextModerationPlusResponseBody
}

type TextModerationPlusResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TextModerationPlusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TextModerationPlusResponse) String() string {
	return dara.Prettify(s)
}

func (s TextModerationPlusResponse) GoString() string {
	return s.String()
}

func (s *TextModerationPlusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TextModerationPlusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TextModerationPlusResponse) GetBody() *TextModerationPlusResponseBody {
	return s.Body
}

func (s *TextModerationPlusResponse) SetHeaders(v map[string]*string) *TextModerationPlusResponse {
	s.Headers = v
	return s
}

func (s *TextModerationPlusResponse) SetStatusCode(v int32) *TextModerationPlusResponse {
	s.StatusCode = &v
	return s
}

func (s *TextModerationPlusResponse) SetBody(v *TextModerationPlusResponseBody) *TextModerationPlusResponse {
	s.Body = v
	return s
}

func (s *TextModerationPlusResponse) Validate() error {
	return dara.Validate(s)
}
