// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLoadAgentSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LoadAgentSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LoadAgentSessionResponse
	GetStatusCode() *int32
	SetId(v string) *LoadAgentSessionResponse
	GetId() *string
	SetEvent(v string) *LoadAgentSessionResponse
	GetEvent() *string
	SetBody(v *LoadAgentSessionResponseBody) *LoadAgentSessionResponse
	GetBody() *LoadAgentSessionResponseBody
}

type LoadAgentSessionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                       `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                       `json:"event,omitempty" xml:"event,omitempty"`
	Body       *LoadAgentSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LoadAgentSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s LoadAgentSessionResponse) GoString() string {
	return s.String()
}

func (s *LoadAgentSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LoadAgentSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LoadAgentSessionResponse) GetId() *string {
	return s.Id
}

func (s *LoadAgentSessionResponse) GetEvent() *string {
	return s.Event
}

func (s *LoadAgentSessionResponse) GetBody() *LoadAgentSessionResponseBody {
	return s.Body
}

func (s *LoadAgentSessionResponse) SetHeaders(v map[string]*string) *LoadAgentSessionResponse {
	s.Headers = v
	return s
}

func (s *LoadAgentSessionResponse) SetStatusCode(v int32) *LoadAgentSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *LoadAgentSessionResponse) SetId(v string) *LoadAgentSessionResponse {
	s.Id = &v
	return s
}

func (s *LoadAgentSessionResponse) SetEvent(v string) *LoadAgentSessionResponse {
	s.Event = &v
	return s
}

func (s *LoadAgentSessionResponse) SetBody(v *LoadAgentSessionResponseBody) *LoadAgentSessionResponse {
	s.Body = v
	return s
}

func (s *LoadAgentSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
