// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMultiPriceForLicenseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParamStr(v string) *DescribeMultiPriceForLicenseRequest
	GetParamStr() *string
}

type DescribeMultiPriceForLicenseRequest struct {
	ParamStr *string `json:"ParamStr,omitempty" xml:"ParamStr,omitempty"`
}

func (s DescribeMultiPriceForLicenseRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiPriceForLicenseRequest) GoString() string {
	return s.String()
}

func (s *DescribeMultiPriceForLicenseRequest) GetParamStr() *string {
	return s.ParamStr
}

func (s *DescribeMultiPriceForLicenseRequest) SetParamStr(v string) *DescribeMultiPriceForLicenseRequest {
	s.ParamStr = &v
	return s
}

func (s *DescribeMultiPriceForLicenseRequest) Validate() error {
	return dara.Validate(s)
}
