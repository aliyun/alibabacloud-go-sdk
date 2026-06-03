// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSmsTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSmsTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSmsTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ListSmsTemplateResponseBody) *ListSmsTemplateResponse
	GetBody() *ListSmsTemplateResponseBody
}

type ListSmsTemplateResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSmsTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSmsTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSmsTemplateResponse) GoString() string {
	return s.String()
}

func (s *ListSmsTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSmsTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSmsTemplateResponse) GetBody() *ListSmsTemplateResponseBody {
	return s.Body
}

func (s *ListSmsTemplateResponse) SetHeaders(v map[string]*string) *ListSmsTemplateResponse {
	s.Headers = v
	return s
}

func (s *ListSmsTemplateResponse) SetStatusCode(v int32) *ListSmsTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSmsTemplateResponse) SetBody(v *ListSmsTemplateResponseBody) *ListSmsTemplateResponse {
	s.Body = v
	return s
}

func (s *ListSmsTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
