// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentRuntimesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAgentRuntimesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAgentRuntimesResponse
	GetStatusCode() *int32
	SetBody(v *ListAgentRuntimesResult) *ListAgentRuntimesResponse
	GetBody() *ListAgentRuntimesResult
}

type ListAgentRuntimesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgentRuntimesResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgentRuntimesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgentRuntimesResponse) GoString() string {
	return s.String()
}

func (s *ListAgentRuntimesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAgentRuntimesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAgentRuntimesResponse) GetBody() *ListAgentRuntimesResult {
	return s.Body
}

func (s *ListAgentRuntimesResponse) SetHeaders(v map[string]*string) *ListAgentRuntimesResponse {
	s.Headers = v
	return s
}

func (s *ListAgentRuntimesResponse) SetStatusCode(v int32) *ListAgentRuntimesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgentRuntimesResponse) SetBody(v *ListAgentRuntimesResult) *ListAgentRuntimesResponse {
	s.Body = v
	return s
}

func (s *ListAgentRuntimesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
