// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentlessRelateMaliciousResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAgentlessRelateMaliciousResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAgentlessRelateMaliciousResponse
	GetStatusCode() *int32
	SetBody(v *ListAgentlessRelateMaliciousResponseBody) *ListAgentlessRelateMaliciousResponse
	GetBody() *ListAgentlessRelateMaliciousResponseBody
}

type ListAgentlessRelateMaliciousResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgentlessRelateMaliciousResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgentlessRelateMaliciousResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgentlessRelateMaliciousResponse) GoString() string {
	return s.String()
}

func (s *ListAgentlessRelateMaliciousResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAgentlessRelateMaliciousResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAgentlessRelateMaliciousResponse) GetBody() *ListAgentlessRelateMaliciousResponseBody {
	return s.Body
}

func (s *ListAgentlessRelateMaliciousResponse) SetHeaders(v map[string]*string) *ListAgentlessRelateMaliciousResponse {
	s.Headers = v
	return s
}

func (s *ListAgentlessRelateMaliciousResponse) SetStatusCode(v int32) *ListAgentlessRelateMaliciousResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgentlessRelateMaliciousResponse) SetBody(v *ListAgentlessRelateMaliciousResponseBody) *ListAgentlessRelateMaliciousResponse {
	s.Body = v
	return s
}

func (s *ListAgentlessRelateMaliciousResponse) Validate() error {
	return dara.Validate(s)
}
