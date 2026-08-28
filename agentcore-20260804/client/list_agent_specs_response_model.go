// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentSpecsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAgentSpecsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAgentSpecsResponse
	GetStatusCode() *int32
	SetBody(v *ListAgentSpecsResponseBody) *ListAgentSpecsResponse
	GetBody() *ListAgentSpecsResponseBody
}

type ListAgentSpecsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgentSpecsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgentSpecsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSpecsResponse) GoString() string {
	return s.String()
}

func (s *ListAgentSpecsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAgentSpecsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAgentSpecsResponse) GetBody() *ListAgentSpecsResponseBody {
	return s.Body
}

func (s *ListAgentSpecsResponse) SetHeaders(v map[string]*string) *ListAgentSpecsResponse {
	s.Headers = v
	return s
}

func (s *ListAgentSpecsResponse) SetStatusCode(v int32) *ListAgentSpecsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgentSpecsResponse) SetBody(v *ListAgentSpecsResponseBody) *ListAgentSpecsResponse {
	s.Body = v
	return s
}

func (s *ListAgentSpecsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
