// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDistributionProductsLinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCodes(v []*string) *DescribeDistributionProductsLinkRequest
	GetCodes() []*string
}

type DescribeDistributionProductsLinkRequest struct {
	// This parameter is required.
	Codes []*string `json:"Codes,omitempty" xml:"Codes,omitempty" type:"Repeated"`
}

func (s DescribeDistributionProductsLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDistributionProductsLinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeDistributionProductsLinkRequest) GetCodes() []*string {
	return s.Codes
}

func (s *DescribeDistributionProductsLinkRequest) SetCodes(v []*string) *DescribeDistributionProductsLinkRequest {
	s.Codes = v
	return s
}

func (s *DescribeDistributionProductsLinkRequest) Validate() error {
	return dara.Validate(s)
}
