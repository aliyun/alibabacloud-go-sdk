// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMultiPriceForLicenseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMultiPriceForLicenseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMultiPriceForLicenseResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMultiPriceForLicenseResponseBody) *DescribeMultiPriceForLicenseResponse
	GetBody() *DescribeMultiPriceForLicenseResponseBody
}

type DescribeMultiPriceForLicenseResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMultiPriceForLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMultiPriceForLicenseResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiPriceForLicenseResponse) GoString() string {
	return s.String()
}

func (s *DescribeMultiPriceForLicenseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMultiPriceForLicenseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMultiPriceForLicenseResponse) GetBody() *DescribeMultiPriceForLicenseResponseBody {
	return s.Body
}

func (s *DescribeMultiPriceForLicenseResponse) SetHeaders(v map[string]*string) *DescribeMultiPriceForLicenseResponse {
	s.Headers = v
	return s
}

func (s *DescribeMultiPriceForLicenseResponse) SetStatusCode(v int32) *DescribeMultiPriceForLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMultiPriceForLicenseResponse) SetBody(v *DescribeMultiPriceForLicenseResponseBody) *DescribeMultiPriceForLicenseResponse {
	s.Body = v
	return s
}

func (s *DescribeMultiPriceForLicenseResponse) Validate() error {
	return dara.Validate(s)
}
