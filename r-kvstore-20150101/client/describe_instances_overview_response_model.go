// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancesOverviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstancesOverviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstancesOverviewResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstancesOverviewResponseBody) *DescribeInstancesOverviewResponse
	GetBody() *DescribeInstancesOverviewResponseBody
}

type DescribeInstancesOverviewResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstancesOverviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstancesOverviewResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesOverviewResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstancesOverviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstancesOverviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstancesOverviewResponse) GetBody() *DescribeInstancesOverviewResponseBody {
	return s.Body
}

func (s *DescribeInstancesOverviewResponse) SetHeaders(v map[string]*string) *DescribeInstancesOverviewResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstancesOverviewResponse) SetStatusCode(v int32) *DescribeInstancesOverviewResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstancesOverviewResponse) SetBody(v *DescribeInstancesOverviewResponseBody) *DescribeInstancesOverviewResponse {
	s.Body = v
	return s
}

func (s *DescribeInstancesOverviewResponse) Validate() error {
	return dara.Validate(s)
}
