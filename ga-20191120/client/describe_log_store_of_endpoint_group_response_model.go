// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogStoreOfEndpointGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLogStoreOfEndpointGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLogStoreOfEndpointGroupResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLogStoreOfEndpointGroupResponseBody) *DescribeLogStoreOfEndpointGroupResponse
	GetBody() *DescribeLogStoreOfEndpointGroupResponseBody
}

type DescribeLogStoreOfEndpointGroupResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLogStoreOfEndpointGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLogStoreOfEndpointGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogStoreOfEndpointGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogStoreOfEndpointGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLogStoreOfEndpointGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLogStoreOfEndpointGroupResponse) GetBody() *DescribeLogStoreOfEndpointGroupResponseBody {
	return s.Body
}

func (s *DescribeLogStoreOfEndpointGroupResponse) SetHeaders(v map[string]*string) *DescribeLogStoreOfEndpointGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogStoreOfEndpointGroupResponse) SetStatusCode(v int32) *DescribeLogStoreOfEndpointGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLogStoreOfEndpointGroupResponse) SetBody(v *DescribeLogStoreOfEndpointGroupResponseBody) *DescribeLogStoreOfEndpointGroupResponse {
	s.Body = v
	return s
}

func (s *DescribeLogStoreOfEndpointGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
