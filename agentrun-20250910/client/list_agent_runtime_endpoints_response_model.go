// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentRuntimeEndpointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAgentRuntimeEndpointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAgentRuntimeEndpointsResponse
	GetStatusCode() *int32
	SetBody(v *ListAgentRuntimeEndpointsResult) *ListAgentRuntimeEndpointsResponse
	GetBody() *ListAgentRuntimeEndpointsResult
}

type ListAgentRuntimeEndpointsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgentRuntimeEndpointsResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgentRuntimeEndpointsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgentRuntimeEndpointsResponse) GoString() string {
	return s.String()
}

func (s *ListAgentRuntimeEndpointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAgentRuntimeEndpointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAgentRuntimeEndpointsResponse) GetBody() *ListAgentRuntimeEndpointsResult {
	return s.Body
}

func (s *ListAgentRuntimeEndpointsResponse) SetHeaders(v map[string]*string) *ListAgentRuntimeEndpointsResponse {
	s.Headers = v
	return s
}

func (s *ListAgentRuntimeEndpointsResponse) SetStatusCode(v int32) *ListAgentRuntimeEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgentRuntimeEndpointsResponse) SetBody(v *ListAgentRuntimeEndpointsResult) *ListAgentRuntimeEndpointsResponse {
	s.Body = v
	return s
}

func (s *ListAgentRuntimeEndpointsResponse) Validate() error {
	return dara.Validate(s)
}
