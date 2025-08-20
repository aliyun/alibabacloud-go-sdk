// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAPSADBInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAPSADBInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAPSADBInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAPSADBInstancesResponseBody) *DescribeAPSADBInstancesResponse
	GetBody() *DescribeAPSADBInstancesResponseBody
}

type DescribeAPSADBInstancesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAPSADBInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAPSADBInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAPSADBInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAPSADBInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAPSADBInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAPSADBInstancesResponse) GetBody() *DescribeAPSADBInstancesResponseBody {
	return s.Body
}

func (s *DescribeAPSADBInstancesResponse) SetHeaders(v map[string]*string) *DescribeAPSADBInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAPSADBInstancesResponse) SetStatusCode(v int32) *DescribeAPSADBInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAPSADBInstancesResponse) SetBody(v *DescribeAPSADBInstancesResponseBody) *DescribeAPSADBInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeAPSADBInstancesResponse) Validate() error {
	return dara.Validate(s)
}
