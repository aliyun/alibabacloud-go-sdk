// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublishedAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPublishedAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPublishedAgentResponse
	GetStatusCode() *int32
	SetBody(v *ListPublishedAgentResponseBody) *ListPublishedAgentResponse
	GetBody() *ListPublishedAgentResponseBody
}

type ListPublishedAgentResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPublishedAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPublishedAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPublishedAgentResponse) GoString() string {
	return s.String()
}

func (s *ListPublishedAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPublishedAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPublishedAgentResponse) GetBody() *ListPublishedAgentResponseBody {
	return s.Body
}

func (s *ListPublishedAgentResponse) SetHeaders(v map[string]*string) *ListPublishedAgentResponse {
	s.Headers = v
	return s
}

func (s *ListPublishedAgentResponse) SetStatusCode(v int32) *ListPublishedAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPublishedAgentResponse) SetBody(v *ListPublishedAgentResponseBody) *ListPublishedAgentResponse {
	s.Body = v
	return s
}

func (s *ListPublishedAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
