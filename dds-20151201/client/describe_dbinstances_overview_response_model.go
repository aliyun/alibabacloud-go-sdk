// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancesOverviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstancesOverviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstancesOverviewResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstancesOverviewResponseBody) *DescribeDBInstancesOverviewResponse
	GetBody() *DescribeDBInstancesOverviewResponseBody
}

type DescribeDBInstancesOverviewResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstancesOverviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstancesOverviewResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesOverviewResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesOverviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstancesOverviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstancesOverviewResponse) GetBody() *DescribeDBInstancesOverviewResponseBody {
	return s.Body
}

func (s *DescribeDBInstancesOverviewResponse) SetHeaders(v map[string]*string) *DescribeDBInstancesOverviewResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstancesOverviewResponse) SetStatusCode(v int32) *DescribeDBInstancesOverviewResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstancesOverviewResponse) SetBody(v *DescribeDBInstancesOverviewResponseBody) *DescribeDBInstancesOverviewResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstancesOverviewResponse) Validate() error {
	return dara.Validate(s)
}
