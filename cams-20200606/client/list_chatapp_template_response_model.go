// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatappTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListChatappTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListChatappTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ListChatappTemplateResponseBody) *ListChatappTemplateResponse
	GetBody() *ListChatappTemplateResponseBody
}

type ListChatappTemplateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListChatappTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListChatappTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ListChatappTemplateResponse) GoString() string {
	return s.String()
}

func (s *ListChatappTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListChatappTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListChatappTemplateResponse) GetBody() *ListChatappTemplateResponseBody {
	return s.Body
}

func (s *ListChatappTemplateResponse) SetHeaders(v map[string]*string) *ListChatappTemplateResponse {
	s.Headers = v
	return s
}

func (s *ListChatappTemplateResponse) SetStatusCode(v int32) *ListChatappTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ListChatappTemplateResponse) SetBody(v *ListChatappTemplateResponseBody) *ListChatappTemplateResponse {
	s.Body = v
	return s
}

func (s *ListChatappTemplateResponse) Validate() error {
	return dara.Validate(s)
}
