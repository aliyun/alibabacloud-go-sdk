// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourcePackageProductResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeResourcePackageProductResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeResourcePackageProductResponse
	GetStatusCode() *int32
	SetBody(v *DescribeResourcePackageProductResponseBody) *DescribeResourcePackageProductResponse
	GetBody() *DescribeResourcePackageProductResponseBody
}

type DescribeResourcePackageProductResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResourcePackageProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResourcePackageProductResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcePackageProductResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackageProductResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeResourcePackageProductResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeResourcePackageProductResponse) GetBody() *DescribeResourcePackageProductResponseBody {
	return s.Body
}

func (s *DescribeResourcePackageProductResponse) SetHeaders(v map[string]*string) *DescribeResourcePackageProductResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourcePackageProductResponse) SetStatusCode(v int32) *DescribeResourcePackageProductResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourcePackageProductResponse) SetBody(v *DescribeResourcePackageProductResponseBody) *DescribeResourcePackageProductResponse {
	s.Body = v
	return s
}

func (s *DescribeResourcePackageProductResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
