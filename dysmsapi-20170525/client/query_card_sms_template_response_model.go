// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCardSmsTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCardSmsTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCardSmsTemplateResponse
	GetStatusCode() *int32
	SetBody(v *QueryCardSmsTemplateResponseBody) *QueryCardSmsTemplateResponse
	GetBody() *QueryCardSmsTemplateResponseBody
}

type QueryCardSmsTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCardSmsTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCardSmsTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCardSmsTemplateResponse) GoString() string {
	return s.String()
}

func (s *QueryCardSmsTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCardSmsTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCardSmsTemplateResponse) GetBody() *QueryCardSmsTemplateResponseBody {
	return s.Body
}

func (s *QueryCardSmsTemplateResponse) SetHeaders(v map[string]*string) *QueryCardSmsTemplateResponse {
	s.Headers = v
	return s
}

func (s *QueryCardSmsTemplateResponse) SetStatusCode(v int32) *QueryCardSmsTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCardSmsTemplateResponse) SetBody(v *QueryCardSmsTemplateResponseBody) *QueryCardSmsTemplateResponse {
	s.Body = v
	return s
}

func (s *QueryCardSmsTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
