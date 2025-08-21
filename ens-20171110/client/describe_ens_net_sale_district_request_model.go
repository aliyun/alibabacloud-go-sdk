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
	// The region code.
	//
	// 	- If you do not specify this parameter, only nodes under the area level that is specified by NetLevelCode are queried.
	//
	// 	- If you specify this parameter, only child nodes in the area that is specified by NetDistrictCode are queried.
	//
	// example:
	//
	// 100105
	NetDistrictCode *string `json:"NetDistrictCode,omitempty" xml:"NetDistrictCode,omitempty"`
	// The network level. Valid values:
	//
	// 	- **Big**: area
	//
	// 	- **Middle**: province
	//
	// 	- **Small**: city
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
