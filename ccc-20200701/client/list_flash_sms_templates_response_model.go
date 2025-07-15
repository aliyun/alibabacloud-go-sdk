// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlashSmsTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFlashSmsTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFlashSmsTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListFlashSmsTemplatesResponseBody) *ListFlashSmsTemplatesResponse
	GetBody() *ListFlashSmsTemplatesResponseBody
}

type ListFlashSmsTemplatesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFlashSmsTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFlashSmsTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFlashSmsTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListFlashSmsTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFlashSmsTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFlashSmsTemplatesResponse) GetBody() *ListFlashSmsTemplatesResponseBody {
	return s.Body
}

func (s *ListFlashSmsTemplatesResponse) SetHeaders(v map[string]*string) *ListFlashSmsTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListFlashSmsTemplatesResponse) SetStatusCode(v int32) *ListFlashSmsTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFlashSmsTemplatesResponse) SetBody(v *ListFlashSmsTemplatesResponseBody) *ListFlashSmsTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListFlashSmsTemplatesResponse) Validate() error {
	return dara.Validate(s)
}
