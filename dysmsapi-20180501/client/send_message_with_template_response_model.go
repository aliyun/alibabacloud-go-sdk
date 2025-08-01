// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendMessageWithTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendMessageWithTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendMessageWithTemplateResponse
	GetStatusCode() *int32
	SetBody(v *SendMessageWithTemplateResponseBody) *SendMessageWithTemplateResponse
	GetBody() *SendMessageWithTemplateResponseBody
}

type SendMessageWithTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendMessageWithTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendMessageWithTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s SendMessageWithTemplateResponse) GoString() string {
	return s.String()
}

func (s *SendMessageWithTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendMessageWithTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendMessageWithTemplateResponse) GetBody() *SendMessageWithTemplateResponseBody {
	return s.Body
}

func (s *SendMessageWithTemplateResponse) SetHeaders(v map[string]*string) *SendMessageWithTemplateResponse {
	s.Headers = v
	return s
}

func (s *SendMessageWithTemplateResponse) SetStatusCode(v int32) *SendMessageWithTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *SendMessageWithTemplateResponse) SetBody(v *SendMessageWithTemplateResponseBody) *SendMessageWithTemplateResponse {
	s.Body = v
	return s
}

func (s *SendMessageWithTemplateResponse) Validate() error {
	return dara.Validate(s)
}
