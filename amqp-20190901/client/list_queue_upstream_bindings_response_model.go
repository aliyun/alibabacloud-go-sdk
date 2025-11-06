// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQueueUpstreamBindingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListQueueUpstreamBindingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListQueueUpstreamBindingsResponse
	GetStatusCode() *int32
	SetBody(v *ListQueueUpstreamBindingsResponseBody) *ListQueueUpstreamBindingsResponse
	GetBody() *ListQueueUpstreamBindingsResponseBody
}

type ListQueueUpstreamBindingsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQueueUpstreamBindingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQueueUpstreamBindingsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListQueueUpstreamBindingsResponse) GoString() string {
	return s.String()
}

func (s *ListQueueUpstreamBindingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListQueueUpstreamBindingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListQueueUpstreamBindingsResponse) GetBody() *ListQueueUpstreamBindingsResponseBody {
	return s.Body
}

func (s *ListQueueUpstreamBindingsResponse) SetHeaders(v map[string]*string) *ListQueueUpstreamBindingsResponse {
	s.Headers = v
	return s
}

func (s *ListQueueUpstreamBindingsResponse) SetStatusCode(v int32) *ListQueueUpstreamBindingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQueueUpstreamBindingsResponse) SetBody(v *ListQueueUpstreamBindingsResponseBody) *ListQueueUpstreamBindingsResponse {
	s.Body = v
	return s
}

func (s *ListQueueUpstreamBindingsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
