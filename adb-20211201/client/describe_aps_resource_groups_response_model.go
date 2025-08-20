// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApsResourceGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApsResourceGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApsResourceGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApsResourceGroupsResponseBody) *DescribeApsResourceGroupsResponse
	GetBody() *DescribeApsResourceGroupsResponseBody
}

type DescribeApsResourceGroupsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApsResourceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApsResourceGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsResourceGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeApsResourceGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApsResourceGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApsResourceGroupsResponse) GetBody() *DescribeApsResourceGroupsResponseBody {
	return s.Body
}

func (s *DescribeApsResourceGroupsResponse) SetHeaders(v map[string]*string) *DescribeApsResourceGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeApsResourceGroupsResponse) SetStatusCode(v int32) *DescribeApsResourceGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApsResourceGroupsResponse) SetBody(v *DescribeApsResourceGroupsResponseBody) *DescribeApsResourceGroupsResponse {
	s.Body = v
	return s
}

func (s *DescribeApsResourceGroupsResponse) Validate() error {
	return dara.Validate(s)
}
