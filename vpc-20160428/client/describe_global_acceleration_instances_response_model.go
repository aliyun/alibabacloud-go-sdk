// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGlobalAccelerationInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGlobalAccelerationInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGlobalAccelerationInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGlobalAccelerationInstancesResponseBody) *DescribeGlobalAccelerationInstancesResponse
	GetBody() *DescribeGlobalAccelerationInstancesResponseBody
}

type DescribeGlobalAccelerationInstancesResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGlobalAccelerationInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGlobalAccelerationInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalAccelerationInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeGlobalAccelerationInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGlobalAccelerationInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGlobalAccelerationInstancesResponse) GetBody() *DescribeGlobalAccelerationInstancesResponseBody {
	return s.Body
}

func (s *DescribeGlobalAccelerationInstancesResponse) SetHeaders(v map[string]*string) *DescribeGlobalAccelerationInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeGlobalAccelerationInstancesResponse) SetStatusCode(v int32) *DescribeGlobalAccelerationInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGlobalAccelerationInstancesResponse) SetBody(v *DescribeGlobalAccelerationInstancesResponseBody) *DescribeGlobalAccelerationInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeGlobalAccelerationInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
