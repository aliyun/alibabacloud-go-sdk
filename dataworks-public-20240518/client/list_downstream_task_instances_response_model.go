// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDownstreamTaskInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDownstreamTaskInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDownstreamTaskInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListDownstreamTaskInstancesResponseBody) *ListDownstreamTaskInstancesResponse
	GetBody() *ListDownstreamTaskInstancesResponseBody
}

type ListDownstreamTaskInstancesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDownstreamTaskInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDownstreamTaskInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDownstreamTaskInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListDownstreamTaskInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDownstreamTaskInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDownstreamTaskInstancesResponse) GetBody() *ListDownstreamTaskInstancesResponseBody {
	return s.Body
}

func (s *ListDownstreamTaskInstancesResponse) SetHeaders(v map[string]*string) *ListDownstreamTaskInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListDownstreamTaskInstancesResponse) SetStatusCode(v int32) *ListDownstreamTaskInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDownstreamTaskInstancesResponse) SetBody(v *ListDownstreamTaskInstancesResponseBody) *ListDownstreamTaskInstancesResponse {
	s.Body = v
	return s
}

func (s *ListDownstreamTaskInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
