// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsNetSaleDistrictRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetDistrictCode(v string) *DescribeEnsNetSaleDistrictRequest
	GetNetDistrictCode() *string
	SetNetLevelCode(v string) *DescribeEnsNetSaleDistrictRequest
	GetNetLevelCode() *string
}

type DescribeEnsNetSaleDistrictRequest struct {
	// The district code.
	//
	// - If this parameter is left empty, only the nodes under the network level specified by NetLevelCode are queried.
	//
	// - If this parameter is specified, only the child nodes under the district specified by NetDistrictCode are queried.
	//
	// example:
	//
	// 100105
	NetDistrictCode *string `json:"NetDistrictCode,omitempty" xml:"NetDistrictCode,omitempty"`
	// The network level. Valid values:
	//
	// - **Big**: Major region.
	//
	// - **Middle**: Province.
	//
	// - **Small**: City.
	//
	// This parameter is required.
	//
	// example:
	//
	// Big
	NetLevelCode *string `json:"NetLevelCode,omitempty" xml:"NetLevelCode,omitempty"`
}

func (s DescribeEnsNetSaleDistrictRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsNetSaleDistrictRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetSaleDistrictRequest) GetNetDistrictCode() *string {
	return s.NetDistrictCode
}

func (s *DescribeEnsNetSaleDistrictRequest) GetNetLevelCode() *string {
	return s.NetLevelCode
}

func (s *DescribeEnsNetSaleDistrictRequest) SetNetDistrictCode(v string) *DescribeEnsNetSaleDistrictRequest {
	s.NetDistrictCode = &v
	return s
}

func (s *DescribeEnsNetSaleDistrictRequest) SetNetLevelCode(v string) *DescribeEnsNetSaleDistrictRequest {
	s.NetLevelCode = &v
	return s
}

func (s *DescribeEnsNetSaleDistrictRequest) Validate() error {
	return dara.Validate(s)
}
