// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentCodeRuntimesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAgentCodeRuntimesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAgentCodeRuntimesResponse
	GetStatusCode() *int32
	SetBody(v *ListAgentCodeRuntimesResponseBody) *ListAgentCodeRuntimesResponse
	GetBody() *ListAgentCodeRuntimesResponseBody
}

type ListAgentCodeRuntimesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgentCodeRuntimesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgentCodeRuntimesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgentCodeRuntimesResponse) GoString() string {
	return s.String()
}

func (s *ListAgentCodeRuntimesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAgentCodeRuntimesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAgentCodeRuntimesResponse) GetBody() *ListAgentCodeRuntimesResponseBody {
	return s.Body
}

func (s *ListAgentCodeRuntimesResponse) SetHeaders(v map[string]*string) *ListAgentCodeRuntimesResponse {
	s.Headers = v
	return s
}

func (s *ListAgentCodeRuntimesResponse) SetStatusCode(v int32) *ListAgentCodeRuntimesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgentCodeRuntimesResponse) SetBody(v *ListAgentCodeRuntimesResponseBody) *ListAgentCodeRuntimesResponse {
	s.Body = v
	return s
}

func (s *ListAgentCodeRuntimesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
