// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatappTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateChatappTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateChatappTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateChatappTemplateResponseBody) *CreateChatappTemplateResponse
	GetBody() *CreateChatappTemplateResponseBody
}

type CreateChatappTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateChatappTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateChatappTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateChatappTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateChatappTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateChatappTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateChatappTemplateResponse) GetBody() *CreateChatappTemplateResponseBody {
	return s.Body
}

func (s *CreateChatappTemplateResponse) SetHeaders(v map[string]*string) *CreateChatappTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateChatappTemplateResponse) SetStatusCode(v int32) *CreateChatappTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateChatappTemplateResponse) SetBody(v *CreateChatappTemplateResponseBody) *CreateChatappTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateChatappTemplateResponse) Validate() error {
	return dara.Validate(s)
}
