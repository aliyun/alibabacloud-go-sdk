// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYaoChiAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetYaoChiAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetYaoChiAgentResponse
	GetStatusCode() *int32
	SetId(v string) *GetYaoChiAgentResponse
	GetId() *string
	SetEvent(v string) *GetYaoChiAgentResponse
	GetEvent() *string
	SetBody(v *GetYaoChiAgentResponseBody) *GetYaoChiAgentResponse
	GetBody() *GetYaoChiAgentResponseBody
}

type GetYaoChiAgentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                     `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                     `json:"event,omitempty" xml:"event,omitempty"`
	Body       *GetYaoChiAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetYaoChiAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetYaoChiAgentResponse) GoString() string {
	return s.String()
}

func (s *GetYaoChiAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetYaoChiAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetYaoChiAgentResponse) GetId() *string {
	return s.Id
}

func (s *GetYaoChiAgentResponse) GetEvent() *string {
	return s.Event
}

func (s *GetYaoChiAgentResponse) GetBody() *GetYaoChiAgentResponseBody {
	return s.Body
}

func (s *GetYaoChiAgentResponse) SetHeaders(v map[string]*string) *GetYaoChiAgentResponse {
	s.Headers = v
	return s
}

func (s *GetYaoChiAgentResponse) SetStatusCode(v int32) *GetYaoChiAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetYaoChiAgentResponse) SetId(v string) *GetYaoChiAgentResponse {
	s.Id = &v
	return s
}

func (s *GetYaoChiAgentResponse) SetEvent(v string) *GetYaoChiAgentResponse {
	s.Event = &v
	return s
}

func (s *GetYaoChiAgentResponse) SetBody(v *GetYaoChiAgentResponseBody) *GetYaoChiAgentResponse {
	s.Body = v
	return s
}

func (s *GetYaoChiAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
