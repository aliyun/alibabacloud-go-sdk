// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAIAgentInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAIAgentInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAIAgentInstanceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAIAgentInstanceResponseBody) *UpdateAIAgentInstanceResponse
	GetBody() *UpdateAIAgentInstanceResponseBody
}

type UpdateAIAgentInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAIAgentInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAIAgentInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAIAgentInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateAIAgentInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAIAgentInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAIAgentInstanceResponse) GetBody() *UpdateAIAgentInstanceResponseBody {
	return s.Body
}

func (s *UpdateAIAgentInstanceResponse) SetHeaders(v map[string]*string) *UpdateAIAgentInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateAIAgentInstanceResponse) SetStatusCode(v int32) *UpdateAIAgentInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAIAgentInstanceResponse) SetBody(v *UpdateAIAgentInstanceResponseBody) *UpdateAIAgentInstanceResponse {
	s.Body = v
	return s
}

func (s *UpdateAIAgentInstanceResponse) Validate() error {
	return dara.Validate(s)
}
