// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMultiModalAgentSSEResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MultiModalAgentSSEResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MultiModalAgentSSEResponse
	GetStatusCode() *int32
	SetId(v string) *MultiModalAgentSSEResponse
	GetId() *string
	SetEvent(v string) *MultiModalAgentSSEResponse
	GetEvent() *string
	SetBody(v *MultiModalAgentSSEResponseBody) *MultiModalAgentSSEResponse
	GetBody() *MultiModalAgentSSEResponseBody
}

type MultiModalAgentSSEResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                         `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                         `json:"event,omitempty" xml:"event,omitempty"`
	Body       *MultiModalAgentSSEResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MultiModalAgentSSEResponse) String() string {
	return dara.Prettify(s)
}

func (s MultiModalAgentSSEResponse) GoString() string {
	return s.String()
}

func (s *MultiModalAgentSSEResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MultiModalAgentSSEResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MultiModalAgentSSEResponse) GetId() *string {
	return s.Id
}

func (s *MultiModalAgentSSEResponse) GetEvent() *string {
	return s.Event
}

func (s *MultiModalAgentSSEResponse) GetBody() *MultiModalAgentSSEResponseBody {
	return s.Body
}

func (s *MultiModalAgentSSEResponse) SetHeaders(v map[string]*string) *MultiModalAgentSSEResponse {
	s.Headers = v
	return s
}

func (s *MultiModalAgentSSEResponse) SetStatusCode(v int32) *MultiModalAgentSSEResponse {
	s.StatusCode = &v
	return s
}

func (s *MultiModalAgentSSEResponse) SetId(v string) *MultiModalAgentSSEResponse {
	s.Id = &v
	return s
}

func (s *MultiModalAgentSSEResponse) SetEvent(v string) *MultiModalAgentSSEResponse {
	s.Event = &v
	return s
}

func (s *MultiModalAgentSSEResponse) SetBody(v *MultiModalAgentSSEResponseBody) *MultiModalAgentSSEResponse {
	s.Body = v
	return s
}

func (s *MultiModalAgentSSEResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
