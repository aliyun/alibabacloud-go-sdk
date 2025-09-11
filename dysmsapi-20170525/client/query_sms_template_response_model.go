// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySmsTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySmsTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySmsTemplateResponse
	GetStatusCode() *int32
	SetBody(v *QuerySmsTemplateResponseBody) *QuerySmsTemplateResponse
	GetBody() *QuerySmsTemplateResponseBody
}

type QuerySmsTemplateResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySmsTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySmsTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySmsTemplateResponse) GoString() string {
	return s.String()
}

func (s *QuerySmsTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySmsTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySmsTemplateResponse) GetBody() *QuerySmsTemplateResponseBody {
	return s.Body
}

func (s *QuerySmsTemplateResponse) SetHeaders(v map[string]*string) *QuerySmsTemplateResponse {
	s.Headers = v
	return s
}

func (s *QuerySmsTemplateResponse) SetStatusCode(v int32) *QuerySmsTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySmsTemplateResponse) SetBody(v *QuerySmsTemplateResponseBody) *QuerySmsTemplateResponse {
	s.Body = v
	return s
}

func (s *QuerySmsTemplateResponse) Validate() error {
	return dara.Validate(s)
}
