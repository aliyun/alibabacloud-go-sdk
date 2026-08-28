// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentIMChannelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAgentIMChannelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAgentIMChannelsResponse
	GetStatusCode() *int32
	SetBody(v *ListAgentIMChannelsResponseBody) *ListAgentIMChannelsResponse
	GetBody() *ListAgentIMChannelsResponseBody
}

type ListAgentIMChannelsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgentIMChannelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgentIMChannelsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgentIMChannelsResponse) GoString() string {
	return s.String()
}

func (s *ListAgentIMChannelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAgentIMChannelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAgentIMChannelsResponse) GetBody() *ListAgentIMChannelsResponseBody {
	return s.Body
}

func (s *ListAgentIMChannelsResponse) SetHeaders(v map[string]*string) *ListAgentIMChannelsResponse {
	s.Headers = v
	return s
}

func (s *ListAgentIMChannelsResponse) SetStatusCode(v int32) *ListAgentIMChannelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgentIMChannelsResponse) SetBody(v *ListAgentIMChannelsResponseBody) *ListAgentIMChannelsResponse {
	s.Body = v
	return s
}

func (s *ListAgentIMChannelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
