// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcListLiteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpcListLiteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpcListLiteResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpcListLiteResponseBody) *DescribeVpcListLiteResponse
	GetBody() *DescribeVpcListLiteResponseBody
}

type DescribeVpcListLiteResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcListLiteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcListLiteResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcListLiteResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcListLiteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpcListLiteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpcListLiteResponse) GetBody() *DescribeVpcListLiteResponseBody {
	return s.Body
}

func (s *DescribeVpcListLiteResponse) SetHeaders(v map[string]*string) *DescribeVpcListLiteResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcListLiteResponse) SetStatusCode(v int32) *DescribeVpcListLiteResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcListLiteResponse) SetBody(v *DescribeVpcListLiteResponseBody) *DescribeVpcListLiteResponse {
	s.Body = v
	return s
}

func (s *DescribeVpcListLiteResponse) Validate() error {
	return dara.Validate(s)
}
