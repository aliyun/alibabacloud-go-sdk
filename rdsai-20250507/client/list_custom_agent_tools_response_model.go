// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomAgentToolsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCustomAgentToolsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCustomAgentToolsResponse
	GetStatusCode() *int32
	SetBody(v *ListCustomAgentToolsResponseBody) *ListCustomAgentToolsResponse
	GetBody() *ListCustomAgentToolsResponseBody
}

type ListCustomAgentToolsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomAgentToolsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomAgentToolsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAgentToolsResponse) GoString() string {
	return s.String()
}

func (s *ListCustomAgentToolsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCustomAgentToolsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCustomAgentToolsResponse) GetBody() *ListCustomAgentToolsResponseBody {
	return s.Body
}

func (s *ListCustomAgentToolsResponse) SetHeaders(v map[string]*string) *ListCustomAgentToolsResponse {
	s.Headers = v
	return s
}

func (s *ListCustomAgentToolsResponse) SetStatusCode(v int32) *ListCustomAgentToolsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomAgentToolsResponse) SetBody(v *ListCustomAgentToolsResponseBody) *ListCustomAgentToolsResponse {
	s.Body = v
	return s
}

func (s *ListCustomAgentToolsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
