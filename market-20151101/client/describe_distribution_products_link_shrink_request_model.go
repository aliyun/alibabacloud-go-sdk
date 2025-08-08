// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDistributionProductsLinkShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCodesShrink(v string) *DescribeDistributionProductsLinkShrinkRequest
	GetCodesShrink() *string
}

type DescribeDistributionProductsLinkShrinkRequest struct {
	// This parameter is required.
	CodesShrink *string `json:"Codes,omitempty" xml:"Codes,omitempty"`
}

func (s DescribeDistributionProductsLinkShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDistributionProductsLinkShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeDistributionProductsLinkShrinkRequest) GetCodesShrink() *string {
	return s.CodesShrink
}

func (s *DescribeDistributionProductsLinkShrinkRequest) SetCodesShrink(v string) *DescribeDistributionProductsLinkShrinkRequest {
	s.CodesShrink = &v
	return s
}

func (s *DescribeDistributionProductsLinkShrinkRequest) Validate() error {
	return dara.Validate(s)
}
