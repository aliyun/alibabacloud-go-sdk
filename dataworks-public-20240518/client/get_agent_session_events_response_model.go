// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentSessionEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentSessionEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentSessionEventsResponse
	GetStatusCode() *int32
	SetId(v string) *GetAgentSessionEventsResponse
	GetId() *string
	SetEvent(v string) *GetAgentSessionEventsResponse
	GetEvent() *string
	SetBody(v *GetAgentSessionEventsResponseBody) *GetAgentSessionEventsResponse
	GetBody() *GetAgentSessionEventsResponseBody
}

type GetAgentSessionEventsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                            `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                            `json:"event,omitempty" xml:"event,omitempty"`
	Body       *GetAgentSessionEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentSessionEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSessionEventsResponse) GoString() string {
	return s.String()
}

func (s *GetAgentSessionEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentSessionEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentSessionEventsResponse) GetId() *string {
	return s.Id
}

func (s *GetAgentSessionEventsResponse) GetEvent() *string {
	return s.Event
}

func (s *GetAgentSessionEventsResponse) GetBody() *GetAgentSessionEventsResponseBody {
	return s.Body
}

func (s *GetAgentSessionEventsResponse) SetHeaders(v map[string]*string) *GetAgentSessionEventsResponse {
	s.Headers = v
	return s
}

func (s *GetAgentSessionEventsResponse) SetStatusCode(v int32) *GetAgentSessionEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentSessionEventsResponse) SetId(v string) *GetAgentSessionEventsResponse {
	s.Id = &v
	return s
}

func (s *GetAgentSessionEventsResponse) SetEvent(v string) *GetAgentSessionEventsResponse {
	s.Event = &v
	return s
}

func (s *GetAgentSessionEventsResponse) SetBody(v *GetAgentSessionEventsResponseBody) *GetAgentSessionEventsResponse {
	s.Body = v
	return s
}

func (s *GetAgentSessionEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
