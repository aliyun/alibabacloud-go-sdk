// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InvokeAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InvokeAgentResponse
	GetStatusCode() *int32
	SetId(v string) *InvokeAgentResponse
	GetId() *string
	SetEvent(v string) *InvokeAgentResponse
	GetEvent() *string
	SetBody(v *InvokeAgentResponseBody) *InvokeAgentResponse
	GetBody() *InvokeAgentResponseBody
}

type InvokeAgentResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                  `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                  `json:"event,omitempty" xml:"event,omitempty"`
	Body       *InvokeAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvokeAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s InvokeAgentResponse) GoString() string {
	return s.String()
}

func (s *InvokeAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InvokeAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InvokeAgentResponse) GetId() *string {
	return s.Id
}

func (s *InvokeAgentResponse) GetEvent() *string {
	return s.Event
}

func (s *InvokeAgentResponse) GetBody() *InvokeAgentResponseBody {
	return s.Body
}

func (s *InvokeAgentResponse) SetHeaders(v map[string]*string) *InvokeAgentResponse {
	s.Headers = v
	return s
}

func (s *InvokeAgentResponse) SetStatusCode(v int32) *InvokeAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *InvokeAgentResponse) SetId(v string) *InvokeAgentResponse {
	s.Id = &v
	return s
}

func (s *InvokeAgentResponse) SetEvent(v string) *InvokeAgentResponse {
	s.Event = &v
	return s
}

func (s *InvokeAgentResponse) SetBody(v *InvokeAgentResponseBody) *InvokeAgentResponse {
	s.Body = v
	return s
}

func (s *InvokeAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
