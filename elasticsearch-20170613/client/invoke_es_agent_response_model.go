// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeEsAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InvokeEsAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InvokeEsAgentResponse
	GetStatusCode() *int32
	SetId(v string) *InvokeEsAgentResponse
	GetId() *string
	SetEvent(v string) *InvokeEsAgentResponse
	GetEvent() *string
	SetBody(v *InvokeEsAgentResponseBody) *InvokeEsAgentResponse
	GetBody() *InvokeEsAgentResponseBody
}

type InvokeEsAgentResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                    `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                    `json:"event,omitempty" xml:"event,omitempty"`
	Body       *InvokeEsAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvokeEsAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s InvokeEsAgentResponse) GoString() string {
	return s.String()
}

func (s *InvokeEsAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InvokeEsAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InvokeEsAgentResponse) GetId() *string {
	return s.Id
}

func (s *InvokeEsAgentResponse) GetEvent() *string {
	return s.Event
}

func (s *InvokeEsAgentResponse) GetBody() *InvokeEsAgentResponseBody {
	return s.Body
}

func (s *InvokeEsAgentResponse) SetHeaders(v map[string]*string) *InvokeEsAgentResponse {
	s.Headers = v
	return s
}

func (s *InvokeEsAgentResponse) SetStatusCode(v int32) *InvokeEsAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *InvokeEsAgentResponse) SetId(v string) *InvokeEsAgentResponse {
	s.Id = &v
	return s
}

func (s *InvokeEsAgentResponse) SetEvent(v string) *InvokeEsAgentResponse {
	s.Event = &v
	return s
}

func (s *InvokeEsAgentResponse) SetBody(v *InvokeEsAgentResponseBody) *InvokeEsAgentResponse {
	s.Body = v
	return s
}

func (s *InvokeEsAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
