// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePackageDeductionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePackageDeductionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePackageDeductionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePackageDeductionsResponseBody) *DescribePackageDeductionsResponse
	GetBody() *DescribePackageDeductionsResponseBody
}

type DescribePackageDeductionsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePackageDeductionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePackageDeductionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePackageDeductionsResponse) GoString() string {
	return s.String()
}

func (s *DescribePackageDeductionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePackageDeductionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePackageDeductionsResponse) GetBody() *DescribePackageDeductionsResponseBody {
	return s.Body
}

func (s *DescribePackageDeductionsResponse) SetHeaders(v map[string]*string) *DescribePackageDeductionsResponse {
	s.Headers = v
	return s
}

func (s *DescribePackageDeductionsResponse) SetStatusCode(v int32) *DescribePackageDeductionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePackageDeductionsResponse) SetBody(v *DescribePackageDeductionsResponseBody) *DescribePackageDeductionsResponse {
	s.Body = v
	return s
}

func (s *DescribePackageDeductionsResponse) Validate() error {
	return dara.Validate(s)
}
