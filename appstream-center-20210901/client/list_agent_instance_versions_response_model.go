// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentInstanceVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAgentInstanceVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAgentInstanceVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListAgentInstanceVersionsResponseBody) *ListAgentInstanceVersionsResponse
	GetBody() *ListAgentInstanceVersionsResponseBody
}

type ListAgentInstanceVersionsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgentInstanceVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgentInstanceVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgentInstanceVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListAgentInstanceVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAgentInstanceVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAgentInstanceVersionsResponse) GetBody() *ListAgentInstanceVersionsResponseBody {
	return s.Body
}

func (s *ListAgentInstanceVersionsResponse) SetHeaders(v map[string]*string) *ListAgentInstanceVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListAgentInstanceVersionsResponse) SetStatusCode(v int32) *ListAgentInstanceVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgentInstanceVersionsResponse) SetBody(v *ListAgentInstanceVersionsResponseBody) *ListAgentInstanceVersionsResponse {
	s.Body = v
	return s
}

func (s *ListAgentInstanceVersionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
