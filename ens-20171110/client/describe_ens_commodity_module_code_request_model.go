// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsCommodityModuleCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommodityCode(v string) *DescribeEnsCommodityModuleCodeRequest
	GetCommodityCode() *string
	SetModuleCode(v string) *DescribeEnsCommodityModuleCodeRequest
	GetModuleCode() *string
}

type DescribeEnsCommodityModuleCodeRequest struct {
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	ModuleCode    *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
}

func (s DescribeEnsCommodityModuleCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsCommodityModuleCodeRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnsCommodityModuleCodeRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeEnsCommodityModuleCodeRequest) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *DescribeEnsCommodityModuleCodeRequest) SetCommodityCode(v string) *DescribeEnsCommodityModuleCodeRequest {
	s.CommodityCode = &v
	return s
}

func (s *DescribeEnsCommodityModuleCodeRequest) SetModuleCode(v string) *DescribeEnsCommodityModuleCodeRequest {
	s.ModuleCode = &v
	return s
}

func (s *DescribeEnsCommodityModuleCodeRequest) Validate() error {
	return dara.Validate(s)
}
