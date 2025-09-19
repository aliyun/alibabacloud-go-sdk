// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupedInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGroupedInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGroupedInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGroupedInstancesResponseBody) *DescribeGroupedInstancesResponse
	GetBody() *DescribeGroupedInstancesResponseBody
}

type DescribeGroupedInstancesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGroupedInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGroupedInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupedInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeGroupedInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGroupedInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGroupedInstancesResponse) GetBody() *DescribeGroupedInstancesResponseBody {
	return s.Body
}

func (s *DescribeGroupedInstancesResponse) SetHeaders(v map[string]*string) *DescribeGroupedInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeGroupedInstancesResponse) SetStatusCode(v int32) *DescribeGroupedInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGroupedInstancesResponse) SetBody(v *DescribeGroupedInstancesResponseBody) *DescribeGroupedInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeGroupedInstancesResponse) Validate() error {
	return dara.Validate(s)
}
