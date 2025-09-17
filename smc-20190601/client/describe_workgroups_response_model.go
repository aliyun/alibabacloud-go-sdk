// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWorkgroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWorkgroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWorkgroupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWorkgroupsResponseBody) *DescribeWorkgroupsResponse
	GetBody() *DescribeWorkgroupsResponseBody
}

type DescribeWorkgroupsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWorkgroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWorkgroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWorkgroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeWorkgroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWorkgroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWorkgroupsResponse) GetBody() *DescribeWorkgroupsResponseBody {
	return s.Body
}

func (s *DescribeWorkgroupsResponse) SetHeaders(v map[string]*string) *DescribeWorkgroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeWorkgroupsResponse) SetStatusCode(v int32) *DescribeWorkgroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWorkgroupsResponse) SetBody(v *DescribeWorkgroupsResponseBody) *DescribeWorkgroupsResponse {
	s.Body = v
	return s
}

func (s *DescribeWorkgroupsResponse) Validate() error {
	return dara.Validate(s)
}
