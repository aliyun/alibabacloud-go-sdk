// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMultiPriceForLicenseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *DescribeMultiPriceForLicenseResponseBody
	GetData() *string
	SetRequestId(v string) *DescribeMultiPriceForLicenseResponseBody
	GetRequestId() *string
}

type DescribeMultiPriceForLicenseResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMultiPriceForLicenseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiPriceForLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMultiPriceForLicenseResponseBody) GetData() *string {
	return s.Data
}

func (s *DescribeMultiPriceForLicenseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMultiPriceForLicenseResponseBody) SetData(v string) *DescribeMultiPriceForLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeMultiPriceForLicenseResponseBody) SetRequestId(v string) *DescribeMultiPriceForLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMultiPriceForLicenseResponseBody) Validate() error {
	return dara.Validate(s)
}
