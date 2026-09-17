// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPromptAgentSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PromptAgentSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PromptAgentSessionResponse
	GetStatusCode() *int32
	SetId(v string) *PromptAgentSessionResponse
	GetId() *string
	SetEvent(v string) *PromptAgentSessionResponse
	GetEvent() *string
	SetBody(v *PromptAgentSessionResponseBody) *PromptAgentSessionResponse
	GetBody() *PromptAgentSessionResponseBody
}

type PromptAgentSessionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                         `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                         `json:"event,omitempty" xml:"event,omitempty"`
	Body       *PromptAgentSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PromptAgentSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s PromptAgentSessionResponse) GoString() string {
	return s.String()
}

func (s *PromptAgentSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PromptAgentSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PromptAgentSessionResponse) GetId() *string {
	return s.Id
}

func (s *PromptAgentSessionResponse) GetEvent() *string {
	return s.Event
}

func (s *PromptAgentSessionResponse) GetBody() *PromptAgentSessionResponseBody {
	return s.Body
}

func (s *PromptAgentSessionResponse) SetHeaders(v map[string]*string) *PromptAgentSessionResponse {
	s.Headers = v
	return s
}

func (s *PromptAgentSessionResponse) SetStatusCode(v int32) *PromptAgentSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *PromptAgentSessionResponse) SetId(v string) *PromptAgentSessionResponse {
	s.Id = &v
	return s
}

func (s *PromptAgentSessionResponse) SetEvent(v string) *PromptAgentSessionResponse {
	s.Event = &v
	return s
}

func (s *PromptAgentSessionResponse) SetBody(v *PromptAgentSessionResponseBody) *PromptAgentSessionResponse {
	s.Body = v
	return s
}

func (s *PromptAgentSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
