// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentRuntimeVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAgentRuntimeVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAgentRuntimeVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListAgentRuntimeVersionsResult) *ListAgentRuntimeVersionsResponse
	GetBody() *ListAgentRuntimeVersionsResult
}

type ListAgentRuntimeVersionsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgentRuntimeVersionsResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgentRuntimeVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgentRuntimeVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListAgentRuntimeVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAgentRuntimeVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAgentRuntimeVersionsResponse) GetBody() *ListAgentRuntimeVersionsResult {
	return s.Body
}

func (s *ListAgentRuntimeVersionsResponse) SetHeaders(v map[string]*string) *ListAgentRuntimeVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListAgentRuntimeVersionsResponse) SetStatusCode(v int32) *ListAgentRuntimeVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgentRuntimeVersionsResponse) SetBody(v *ListAgentRuntimeVersionsResult) *ListAgentRuntimeVersionsResponse {
	s.Body = v
	return s
}

func (s *ListAgentRuntimeVersionsResponse) Validate() error {
	return dara.Validate(s)
}
