// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDownstreamTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDownstreamTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDownstreamTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListDownstreamTasksResponseBody) *ListDownstreamTasksResponse
	GetBody() *ListDownstreamTasksResponseBody
}

type ListDownstreamTasksResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDownstreamTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDownstreamTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDownstreamTasksResponse) GoString() string {
	return s.String()
}

func (s *ListDownstreamTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDownstreamTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDownstreamTasksResponse) GetBody() *ListDownstreamTasksResponseBody {
	return s.Body
}

func (s *ListDownstreamTasksResponse) SetHeaders(v map[string]*string) *ListDownstreamTasksResponse {
	s.Headers = v
	return s
}

func (s *ListDownstreamTasksResponse) SetStatusCode(v int32) *ListDownstreamTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDownstreamTasksResponse) SetBody(v *ListDownstreamTasksResponseBody) *ListDownstreamTasksResponse {
	s.Body = v
	return s
}

func (s *ListDownstreamTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
