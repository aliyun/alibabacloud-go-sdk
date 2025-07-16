// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDistributeTableListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDistributeTableListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDistributeTableListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDistributeTableListResponseBody) *DescribeDistributeTableListResponse
	GetBody() *DescribeDistributeTableListResponseBody
}

type DescribeDistributeTableListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDistributeTableListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDistributeTableListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDistributeTableListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDistributeTableListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDistributeTableListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDistributeTableListResponse) GetBody() *DescribeDistributeTableListResponseBody {
	return s.Body
}

func (s *DescribeDistributeTableListResponse) SetHeaders(v map[string]*string) *DescribeDistributeTableListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDistributeTableListResponse) SetStatusCode(v int32) *DescribeDistributeTableListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDistributeTableListResponse) SetBody(v *DescribeDistributeTableListResponseBody) *DescribeDistributeTableListResponse {
	s.Body = v
	return s
}

func (s *DescribeDistributeTableListResponse) Validate() error {
	return dara.Validate(s)
}
