// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAIAgentInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartAIAgentInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartAIAgentInstanceResponse
	GetStatusCode() *int32
	SetBody(v *StartAIAgentInstanceResponseBody) *StartAIAgentInstanceResponse
	GetBody() *StartAIAgentInstanceResponseBody
}

type StartAIAgentInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartAIAgentInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartAIAgentInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s StartAIAgentInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartAIAgentInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartAIAgentInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartAIAgentInstanceResponse) GetBody() *StartAIAgentInstanceResponseBody {
	return s.Body
}

func (s *StartAIAgentInstanceResponse) SetHeaders(v map[string]*string) *StartAIAgentInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartAIAgentInstanceResponse) SetStatusCode(v int32) *StartAIAgentInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartAIAgentInstanceResponse) SetBody(v *StartAIAgentInstanceResponseBody) *StartAIAgentInstanceResponse {
	s.Body = v
	return s
}

func (s *StartAIAgentInstanceResponse) Validate() error {
	return dara.Validate(s)
}
