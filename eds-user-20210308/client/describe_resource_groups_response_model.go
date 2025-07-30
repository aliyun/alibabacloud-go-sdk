// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeResourceGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeResourceGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeResourceGroupsResponseBody) *DescribeResourceGroupsResponse
	GetBody() *DescribeResourceGroupsResponseBody
}

type DescribeResourceGroupsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResourceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResourceGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeResourceGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeResourceGroupsResponse) GetBody() *DescribeResourceGroupsResponseBody {
	return s.Body
}

func (s *DescribeResourceGroupsResponse) SetHeaders(v map[string]*string) *DescribeResourceGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceGroupsResponse) SetStatusCode(v int32) *DescribeResourceGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourceGroupsResponse) SetBody(v *DescribeResourceGroupsResponseBody) *DescribeResourceGroupsResponse {
	s.Body = v
	return s
}

func (s *DescribeResourceGroupsResponse) Validate() error {
	return dara.Validate(s)
}
