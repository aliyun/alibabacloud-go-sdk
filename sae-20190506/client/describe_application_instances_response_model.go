// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApplicationInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApplicationInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApplicationInstancesResponseBody) *DescribeApplicationInstancesResponse
	GetBody() *DescribeApplicationInstancesResponseBody
}

type DescribeApplicationInstancesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApplicationInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApplicationInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApplicationInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApplicationInstancesResponse) GetBody() *DescribeApplicationInstancesResponseBody {
	return s.Body
}

func (s *DescribeApplicationInstancesResponse) SetHeaders(v map[string]*string) *DescribeApplicationInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationInstancesResponse) SetStatusCode(v int32) *DescribeApplicationInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationInstancesResponse) SetBody(v *DescribeApplicationInstancesResponseBody) *DescribeApplicationInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeApplicationInstancesResponse) Validate() error {
	return dara.Validate(s)
}
