// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentlessTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAgentlessTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAgentlessTaskResponse
	GetStatusCode() *int32
	SetBody(v *ListAgentlessTaskResponseBody) *ListAgentlessTaskResponse
	GetBody() *ListAgentlessTaskResponseBody
}

type ListAgentlessTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgentlessTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgentlessTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgentlessTaskResponse) GoString() string {
	return s.String()
}

func (s *ListAgentlessTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAgentlessTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAgentlessTaskResponse) GetBody() *ListAgentlessTaskResponseBody {
	return s.Body
}

func (s *ListAgentlessTaskResponse) SetHeaders(v map[string]*string) *ListAgentlessTaskResponse {
	s.Headers = v
	return s
}

func (s *ListAgentlessTaskResponse) SetStatusCode(v int32) *ListAgentlessTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgentlessTaskResponse) SetBody(v *ListAgentlessTaskResponseBody) *ListAgentlessTaskResponse {
	s.Body = v
	return s
}

func (s *ListAgentlessTaskResponse) Validate() error {
	return dara.Validate(s)
}
