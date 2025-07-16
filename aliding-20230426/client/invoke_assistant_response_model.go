// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeAssistantResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InvokeAssistantResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InvokeAssistantResponse
	GetStatusCode() *int32
	SetBody(v *InvokeAssistantResponseBody) *InvokeAssistantResponse
	GetBody() *InvokeAssistantResponseBody
}

type InvokeAssistantResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InvokeAssistantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvokeAssistantResponse) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantResponse) GoString() string {
	return s.String()
}

func (s *InvokeAssistantResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InvokeAssistantResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InvokeAssistantResponse) GetBody() *InvokeAssistantResponseBody {
	return s.Body
}

func (s *InvokeAssistantResponse) SetHeaders(v map[string]*string) *InvokeAssistantResponse {
	s.Headers = v
	return s
}

func (s *InvokeAssistantResponse) SetStatusCode(v int32) *InvokeAssistantResponse {
	s.StatusCode = &v
	return s
}

func (s *InvokeAssistantResponse) SetBody(v *InvokeAssistantResponseBody) *InvokeAssistantResponse {
	s.Body = v
	return s
}

func (s *InvokeAssistantResponse) Validate() error {
	return dara.Validate(s)
}
