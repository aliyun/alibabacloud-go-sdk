// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProductInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeProductInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeProductInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeProductInstancesResponseBody) *DescribeProductInstancesResponse
	GetBody() *DescribeProductInstancesResponseBody
}

type DescribeProductInstancesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeProductInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeProductInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeProductInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeProductInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeProductInstancesResponse) GetBody() *DescribeProductInstancesResponseBody {
	return s.Body
}

func (s *DescribeProductInstancesResponse) SetHeaders(v map[string]*string) *DescribeProductInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeProductInstancesResponse) SetStatusCode(v int32) *DescribeProductInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProductInstancesResponse) SetBody(v *DescribeProductInstancesResponseBody) *DescribeProductInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeProductInstancesResponse) Validate() error {
	return dara.Validate(s)
}
