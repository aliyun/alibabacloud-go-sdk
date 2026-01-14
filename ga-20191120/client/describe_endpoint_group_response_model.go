// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEndpointGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEndpointGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEndpointGroupResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEndpointGroupResponseBody) *DescribeEndpointGroupResponse
	GetBody() *DescribeEndpointGroupResponseBody
}

type DescribeEndpointGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEndpointGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEndpointGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEndpointGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeEndpointGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEndpointGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEndpointGroupResponse) GetBody() *DescribeEndpointGroupResponseBody {
	return s.Body
}

func (s *DescribeEndpointGroupResponse) SetHeaders(v map[string]*string) *DescribeEndpointGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeEndpointGroupResponse) SetStatusCode(v int32) *DescribeEndpointGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEndpointGroupResponse) SetBody(v *DescribeEndpointGroupResponseBody) *DescribeEndpointGroupResponse {
	s.Body = v
	return s
}

func (s *DescribeEndpointGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
