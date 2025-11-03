// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServerRelatedGlobalAccelerationInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeServerRelatedGlobalAccelerationInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeServerRelatedGlobalAccelerationInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeServerRelatedGlobalAccelerationInstancesResponseBody) *DescribeServerRelatedGlobalAccelerationInstancesResponse
	GetBody() *DescribeServerRelatedGlobalAccelerationInstancesResponseBody
}

type DescribeServerRelatedGlobalAccelerationInstancesResponse struct {
	Headers    map[string]*string                                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServerRelatedGlobalAccelerationInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServerRelatedGlobalAccelerationInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeServerRelatedGlobalAccelerationInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeServerRelatedGlobalAccelerationInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeServerRelatedGlobalAccelerationInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeServerRelatedGlobalAccelerationInstancesResponse) GetBody() *DescribeServerRelatedGlobalAccelerationInstancesResponseBody {
	return s.Body
}

func (s *DescribeServerRelatedGlobalAccelerationInstancesResponse) SetHeaders(v map[string]*string) *DescribeServerRelatedGlobalAccelerationInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeServerRelatedGlobalAccelerationInstancesResponse) SetStatusCode(v int32) *DescribeServerRelatedGlobalAccelerationInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServerRelatedGlobalAccelerationInstancesResponse) SetBody(v *DescribeServerRelatedGlobalAccelerationInstancesResponseBody) *DescribeServerRelatedGlobalAccelerationInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeServerRelatedGlobalAccelerationInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
