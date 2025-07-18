// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstancesResponseBody) *DescribeDBInstancesResponse
	GetBody() *DescribeDBInstancesResponseBody
}

type DescribeDBInstancesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstancesResponse) GetBody() *DescribeDBInstancesResponseBody {
	return s.Body
}

func (s *DescribeDBInstancesResponse) SetHeaders(v map[string]*string) *DescribeDBInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstancesResponse) SetStatusCode(v int32) *DescribeDBInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstancesResponse) SetBody(v *DescribeDBInstancesResponseBody) *DescribeDBInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstancesResponse) Validate() error {
	return dara.Validate(s)
}
