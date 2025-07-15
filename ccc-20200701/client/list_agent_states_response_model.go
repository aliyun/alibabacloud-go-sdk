// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentStatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAgentStatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAgentStatesResponse
	GetStatusCode() *int32
	SetBody(v *ListAgentStatesResponseBody) *ListAgentStatesResponse
	GetBody() *ListAgentStatesResponseBody
}

type ListAgentStatesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgentStatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgentStatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgentStatesResponse) GoString() string {
	return s.String()
}

func (s *ListAgentStatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAgentStatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAgentStatesResponse) GetBody() *ListAgentStatesResponseBody {
	return s.Body
}

func (s *ListAgentStatesResponse) SetHeaders(v map[string]*string) *ListAgentStatesResponse {
	s.Headers = v
	return s
}

func (s *ListAgentStatesResponse) SetStatusCode(v int32) *ListAgentStatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgentStatesResponse) SetBody(v *ListAgentStatesResponseBody) *ListAgentStatesResponse {
	s.Body = v
	return s
}

func (s *ListAgentStatesResponse) Validate() error {
	return dara.Validate(s)
}
