// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDbInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDbInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDbInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDbInstancesResponseBody) *DescribeDbInstancesResponse
	GetBody() *DescribeDbInstancesResponseBody
}

type DescribeDbInstancesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDbInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDbInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDbInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDbInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDbInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDbInstancesResponse) GetBody() *DescribeDbInstancesResponseBody {
	return s.Body
}

func (s *DescribeDbInstancesResponse) SetHeaders(v map[string]*string) *DescribeDbInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDbInstancesResponse) SetStatusCode(v int32) *DescribeDbInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDbInstancesResponse) SetBody(v *DescribeDbInstancesResponseBody) *DescribeDbInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeDbInstancesResponse) Validate() error {
	return dara.Validate(s)
}
