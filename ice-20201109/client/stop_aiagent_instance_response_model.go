// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAIAgentInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopAIAgentInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopAIAgentInstanceResponse
	GetStatusCode() *int32
	SetBody(v *StopAIAgentInstanceResponseBody) *StopAIAgentInstanceResponse
	GetBody() *StopAIAgentInstanceResponseBody
}

type StopAIAgentInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopAIAgentInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopAIAgentInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s StopAIAgentInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopAIAgentInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopAIAgentInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopAIAgentInstanceResponse) GetBody() *StopAIAgentInstanceResponseBody {
	return s.Body
}

func (s *StopAIAgentInstanceResponse) SetHeaders(v map[string]*string) *StopAIAgentInstanceResponse {
	s.Headers = v
	return s
}

func (s *StopAIAgentInstanceResponse) SetStatusCode(v int32) *StopAIAgentInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StopAIAgentInstanceResponse) SetBody(v *StopAIAgentInstanceResponseBody) *StopAIAgentInstanceResponse {
	s.Body = v
	return s
}

func (s *StopAIAgentInstanceResponse) Validate() error {
	return dara.Validate(s)
}
