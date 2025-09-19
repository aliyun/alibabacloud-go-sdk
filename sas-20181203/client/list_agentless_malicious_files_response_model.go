// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentlessMaliciousFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAgentlessMaliciousFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAgentlessMaliciousFilesResponse
	GetStatusCode() *int32
	SetBody(v *ListAgentlessMaliciousFilesResponseBody) *ListAgentlessMaliciousFilesResponse
	GetBody() *ListAgentlessMaliciousFilesResponseBody
}

type ListAgentlessMaliciousFilesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgentlessMaliciousFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgentlessMaliciousFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgentlessMaliciousFilesResponse) GoString() string {
	return s.String()
}

func (s *ListAgentlessMaliciousFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAgentlessMaliciousFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAgentlessMaliciousFilesResponse) GetBody() *ListAgentlessMaliciousFilesResponseBody {
	return s.Body
}

func (s *ListAgentlessMaliciousFilesResponse) SetHeaders(v map[string]*string) *ListAgentlessMaliciousFilesResponse {
	s.Headers = v
	return s
}

func (s *ListAgentlessMaliciousFilesResponse) SetStatusCode(v int32) *ListAgentlessMaliciousFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgentlessMaliciousFilesResponse) SetBody(v *ListAgentlessMaliciousFilesResponseBody) *ListAgentlessMaliciousFilesResponse {
	s.Body = v
	return s
}

func (s *ListAgentlessMaliciousFilesResponse) Validate() error {
	return dara.Validate(s)
}
