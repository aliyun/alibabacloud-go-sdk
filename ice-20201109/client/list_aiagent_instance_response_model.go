// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIAgentInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAIAgentInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAIAgentInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ListAIAgentInstanceResponseBody) *ListAIAgentInstanceResponse
	GetBody() *ListAIAgentInstanceResponseBody
}

type ListAIAgentInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAIAgentInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAIAgentInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAIAgentInstanceResponse) GoString() string {
	return s.String()
}

func (s *ListAIAgentInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAIAgentInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAIAgentInstanceResponse) GetBody() *ListAIAgentInstanceResponseBody {
	return s.Body
}

func (s *ListAIAgentInstanceResponse) SetHeaders(v map[string]*string) *ListAIAgentInstanceResponse {
	s.Headers = v
	return s
}

func (s *ListAIAgentInstanceResponse) SetStatusCode(v int32) *ListAIAgentInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAIAgentInstanceResponse) SetBody(v *ListAIAgentInstanceResponseBody) *ListAIAgentInstanceResponse {
	s.Body = v
	return s
}

func (s *ListAIAgentInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
