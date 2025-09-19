// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpcListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpcListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpcListResponseBody) *DescribeVpcListResponse
	GetBody() *DescribeVpcListResponseBody
}

type DescribeVpcListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpcListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpcListResponse) GetBody() *DescribeVpcListResponseBody {
	return s.Body
}

func (s *DescribeVpcListResponse) SetHeaders(v map[string]*string) *DescribeVpcListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcListResponse) SetStatusCode(v int32) *DescribeVpcListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcListResponse) SetBody(v *DescribeVpcListResponseBody) *DescribeVpcListResponse {
	s.Body = v
	return s
}

func (s *DescribeVpcListResponse) Validate() error {
	return dara.Validate(s)
}
