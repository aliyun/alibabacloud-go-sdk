// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUpstreamTaskInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUpstreamTaskInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUpstreamTaskInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListUpstreamTaskInstancesResponseBody) *ListUpstreamTaskInstancesResponse
	GetBody() *ListUpstreamTaskInstancesResponseBody
}

type ListUpstreamTaskInstancesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUpstreamTaskInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUpstreamTaskInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUpstreamTaskInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListUpstreamTaskInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUpstreamTaskInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUpstreamTaskInstancesResponse) GetBody() *ListUpstreamTaskInstancesResponseBody {
	return s.Body
}

func (s *ListUpstreamTaskInstancesResponse) SetHeaders(v map[string]*string) *ListUpstreamTaskInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListUpstreamTaskInstancesResponse) SetStatusCode(v int32) *ListUpstreamTaskInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUpstreamTaskInstancesResponse) SetBody(v *ListUpstreamTaskInstancesResponseBody) *ListUpstreamTaskInstancesResponse {
	s.Body = v
	return s
}

func (s *ListUpstreamTaskInstancesResponse) Validate() error {
	return dara.Validate(s)
}
