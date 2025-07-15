// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRealtimeAgentStatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRealtimeAgentStatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRealtimeAgentStatesResponse
	GetStatusCode() *int32
	SetBody(v *ListRealtimeAgentStatesResponseBody) *ListRealtimeAgentStatesResponse
	GetBody() *ListRealtimeAgentStatesResponseBody
}

type ListRealtimeAgentStatesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRealtimeAgentStatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRealtimeAgentStatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRealtimeAgentStatesResponse) GoString() string {
	return s.String()
}

func (s *ListRealtimeAgentStatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRealtimeAgentStatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRealtimeAgentStatesResponse) GetBody() *ListRealtimeAgentStatesResponseBody {
	return s.Body
}

func (s *ListRealtimeAgentStatesResponse) SetHeaders(v map[string]*string) *ListRealtimeAgentStatesResponse {
	s.Headers = v
	return s
}

func (s *ListRealtimeAgentStatesResponse) SetStatusCode(v int32) *ListRealtimeAgentStatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRealtimeAgentStatesResponse) SetBody(v *ListRealtimeAgentStatesResponseBody) *ListRealtimeAgentStatesResponse {
	s.Body = v
	return s
}

func (s *ListRealtimeAgentStatesResponse) Validate() error {
	return dara.Validate(s)
}
