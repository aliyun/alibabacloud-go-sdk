// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAgentVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAgentVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListAgentVersionsResponseBody) *ListAgentVersionsResponse
	GetBody() *ListAgentVersionsResponseBody
}

type ListAgentVersionsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgentVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgentVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgentVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListAgentVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAgentVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAgentVersionsResponse) GetBody() *ListAgentVersionsResponseBody {
	return s.Body
}

func (s *ListAgentVersionsResponse) SetHeaders(v map[string]*string) *ListAgentVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListAgentVersionsResponse) SetStatusCode(v int32) *ListAgentVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgentVersionsResponse) SetBody(v *ListAgentVersionsResponseBody) *ListAgentVersionsResponse {
	s.Body = v
	return s
}

func (s *ListAgentVersionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
