// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeAIAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InvokeAIAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InvokeAIAgentResponse
	GetStatusCode() *int32
	SetId(v string) *InvokeAIAgentResponse
	GetId() *string
	SetEvent(v string) *InvokeAIAgentResponse
	GetEvent() *string
	SetBody(v *InvokeAIAgentResponseBody) *InvokeAIAgentResponse
	GetBody() *InvokeAIAgentResponseBody
}

type InvokeAIAgentResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                    `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                    `json:"event,omitempty" xml:"event,omitempty"`
	Body       *InvokeAIAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvokeAIAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s InvokeAIAgentResponse) GoString() string {
	return s.String()
}

func (s *InvokeAIAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InvokeAIAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InvokeAIAgentResponse) GetId() *string {
	return s.Id
}

func (s *InvokeAIAgentResponse) GetEvent() *string {
	return s.Event
}

func (s *InvokeAIAgentResponse) GetBody() *InvokeAIAgentResponseBody {
	return s.Body
}

func (s *InvokeAIAgentResponse) SetHeaders(v map[string]*string) *InvokeAIAgentResponse {
	s.Headers = v
	return s
}

func (s *InvokeAIAgentResponse) SetStatusCode(v int32) *InvokeAIAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *InvokeAIAgentResponse) SetId(v string) *InvokeAIAgentResponse {
	s.Id = &v
	return s
}

func (s *InvokeAIAgentResponse) SetEvent(v string) *InvokeAIAgentResponse {
	s.Event = &v
	return s
}

func (s *InvokeAIAgentResponse) SetBody(v *InvokeAIAgentResponseBody) *InvokeAIAgentResponse {
	s.Body = v
	return s
}

func (s *InvokeAIAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
