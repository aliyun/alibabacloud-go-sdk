// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentStateLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAgentStateLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAgentStateLogsResponse
	GetStatusCode() *int32
	SetBody(v *ListAgentStateLogsResponseBody) *ListAgentStateLogsResponse
	GetBody() *ListAgentStateLogsResponseBody
}

type ListAgentStateLogsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgentStateLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgentStateLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgentStateLogsResponse) GoString() string {
	return s.String()
}

func (s *ListAgentStateLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAgentStateLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAgentStateLogsResponse) GetBody() *ListAgentStateLogsResponseBody {
	return s.Body
}

func (s *ListAgentStateLogsResponse) SetHeaders(v map[string]*string) *ListAgentStateLogsResponse {
	s.Headers = v
	return s
}

func (s *ListAgentStateLogsResponse) SetStatusCode(v int32) *ListAgentStateLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgentStateLogsResponse) SetBody(v *ListAgentStateLogsResponseBody) *ListAgentStateLogsResponse {
	s.Body = v
	return s
}

func (s *ListAgentStateLogsResponse) Validate() error {
	return dara.Validate(s)
}
