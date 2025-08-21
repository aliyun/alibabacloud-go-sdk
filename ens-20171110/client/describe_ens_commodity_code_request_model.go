// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsCommodityCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommodityCode(v string) *DescribeEnsCommodityCodeRequest
	GetCommodityCode() *string
}

type DescribeEnsCommodityCodeRequest struct {
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
}

func (s DescribeEnsCommodityCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsCommodityCodeRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnsCommodityCodeRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeEnsCommodityCodeRequest) SetCommodityCode(v string) *DescribeEnsCommodityCodeRequest {
	s.CommodityCode = &v
	return s
}

func (s *DescribeEnsCommodityCodeRequest) Validate() error {
	return dara.Validate(s)
}
