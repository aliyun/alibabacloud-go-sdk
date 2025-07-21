// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatFlowTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListChatFlowTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListChatFlowTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ListChatFlowTemplateResponseBody) *ListChatFlowTemplateResponse
	GetBody() *ListChatFlowTemplateResponseBody
}

type ListChatFlowTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListChatFlowTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListChatFlowTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ListChatFlowTemplateResponse) GoString() string {
	return s.String()
}

func (s *ListChatFlowTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListChatFlowTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListChatFlowTemplateResponse) GetBody() *ListChatFlowTemplateResponseBody {
	return s.Body
}

func (s *ListChatFlowTemplateResponse) SetHeaders(v map[string]*string) *ListChatFlowTemplateResponse {
	s.Headers = v
	return s
}

func (s *ListChatFlowTemplateResponse) SetStatusCode(v int32) *ListChatFlowTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ListChatFlowTemplateResponse) SetBody(v *ListChatFlowTemplateResponseBody) *ListChatFlowTemplateResponse {
	s.Body = v
	return s
}

func (s *ListChatFlowTemplateResponse) Validate() error {
	return dara.Validate(s)
}
