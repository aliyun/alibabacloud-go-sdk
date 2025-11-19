// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePackageOrdersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePackageOrdersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePackageOrdersResponse
	GetStatusCode() *int32
	SetBody(v *DescribePackageOrdersResponseBody) *DescribePackageOrdersResponse
	GetBody() *DescribePackageOrdersResponseBody
}

type DescribePackageOrdersResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePackageOrdersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePackageOrdersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePackageOrdersResponse) GoString() string {
	return s.String()
}

func (s *DescribePackageOrdersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePackageOrdersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePackageOrdersResponse) GetBody() *DescribePackageOrdersResponseBody {
	return s.Body
}

func (s *DescribePackageOrdersResponse) SetHeaders(v map[string]*string) *DescribePackageOrdersResponse {
	s.Headers = v
	return s
}

func (s *DescribePackageOrdersResponse) SetStatusCode(v int32) *DescribePackageOrdersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePackageOrdersResponse) SetBody(v *DescribePackageOrdersResponseBody) *DescribePackageOrdersResponse {
	s.Body = v
	return s
}

func (s *DescribePackageOrdersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
