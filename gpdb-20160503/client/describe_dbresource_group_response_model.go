// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBResourceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBResourceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBResourceGroupResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBResourceGroupResponseBody) *DescribeDBResourceGroupResponse
	GetBody() *DescribeDBResourceGroupResponseBody
}

type DescribeDBResourceGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBResourceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBResourceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBResourceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBResourceGroupResponse) GetBody() *DescribeDBResourceGroupResponseBody {
	return s.Body
}

func (s *DescribeDBResourceGroupResponse) SetHeaders(v map[string]*string) *DescribeDBResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBResourceGroupResponse) SetStatusCode(v int32) *DescribeDBResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBResourceGroupResponse) SetBody(v *DescribeDBResourceGroupResponseBody) *DescribeDBResourceGroupResponse {
	s.Body = v
	return s
}

func (s *DescribeDBResourceGroupResponse) Validate() error {
	return dara.Validate(s)
}
