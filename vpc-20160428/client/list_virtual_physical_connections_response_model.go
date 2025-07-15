// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirtualPhysicalConnectionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVirtualPhysicalConnectionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVirtualPhysicalConnectionsResponse
	GetStatusCode() *int32
	SetBody(v *ListVirtualPhysicalConnectionsResponseBody) *ListVirtualPhysicalConnectionsResponse
	GetBody() *ListVirtualPhysicalConnectionsResponseBody
}

type ListVirtualPhysicalConnectionsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVirtualPhysicalConnectionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVirtualPhysicalConnectionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVirtualPhysicalConnectionsResponse) GoString() string {
	return s.String()
}

func (s *ListVirtualPhysicalConnectionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVirtualPhysicalConnectionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVirtualPhysicalConnectionsResponse) GetBody() *ListVirtualPhysicalConnectionsResponseBody {
	return s.Body
}

func (s *ListVirtualPhysicalConnectionsResponse) SetHeaders(v map[string]*string) *ListVirtualPhysicalConnectionsResponse {
	s.Headers = v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponse) SetStatusCode(v int32) *ListVirtualPhysicalConnectionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponse) SetBody(v *ListVirtualPhysicalConnectionsResponseBody) *ListVirtualPhysicalConnectionsResponse {
	s.Body = v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponse) Validate() error {
	return dara.Validate(s)
}
