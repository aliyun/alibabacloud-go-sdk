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
	SetId(v string) *InvokeAssistantResponse
	GetId() *string
	SetEvent(v string) *InvokeAssistantResponse
	GetEvent() *string
	SetBody(v *InvokeAssistantResponseBody) *InvokeAssistantResponse
	GetBody() *InvokeAssistantResponseBody
}

type InvokeAssistantResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                      `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                      `json:"event,omitempty" xml:"event,omitempty"`
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

func (s *InvokeAssistantResponse) GetId() *string {
	return s.Id
}

func (s *InvokeAssistantResponse) GetEvent() *string {
	return s.Event
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

func (s *InvokeAssistantResponse) SetId(v string) *InvokeAssistantResponse {
	s.Id = &v
	return s
}

func (s *InvokeAssistantResponse) SetEvent(v string) *InvokeAssistantResponse {
	s.Event = &v
	return s
}

func (s *InvokeAssistantResponse) SetBody(v *InvokeAssistantResponseBody) *InvokeAssistantResponse {
	s.Body = v
	return s
}

func (s *InvokeAssistantResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
