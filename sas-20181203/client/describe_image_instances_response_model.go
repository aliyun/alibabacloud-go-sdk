// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageInstancesResponseBody) *DescribeImageInstancesResponse
	GetBody() *DescribeImageInstancesResponseBody
}

type DescribeImageInstancesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageInstancesResponse) GetBody() *DescribeImageInstancesResponseBody {
	return s.Body
}

func (s *DescribeImageInstancesResponse) SetHeaders(v map[string]*string) *DescribeImageInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageInstancesResponse) SetStatusCode(v int32) *DescribeImageInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageInstancesResponse) SetBody(v *DescribeImageInstancesResponseBody) *DescribeImageInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeImageInstancesResponse) Validate() error {
	return dara.Validate(s)
}
