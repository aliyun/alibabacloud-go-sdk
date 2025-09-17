// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSellerInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSellerInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSellerInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSellerInstancesResponseBody) *DescribeSellerInstancesResponse
	GetBody() *DescribeSellerInstancesResponseBody
}

type DescribeSellerInstancesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSellerInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSellerInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSellerInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSellerInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSellerInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSellerInstancesResponse) GetBody() *DescribeSellerInstancesResponseBody {
	return s.Body
}

func (s *DescribeSellerInstancesResponse) SetHeaders(v map[string]*string) *DescribeSellerInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSellerInstancesResponse) SetStatusCode(v int32) *DescribeSellerInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSellerInstancesResponse) SetBody(v *DescribeSellerInstancesResponseBody) *DescribeSellerInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeSellerInstancesResponse) Validate() error {
	return dara.Validate(s)
}
