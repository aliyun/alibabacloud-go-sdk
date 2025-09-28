// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsNetDistrictRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetDistrictCode(v string) *DescribeEnsNetDistrictRequest
	GetNetDistrictCode() *string
	SetNetDistrictCodeNode(v bool) *DescribeEnsNetDistrictRequest
	GetNetDistrictCodeNode() *bool
	SetNetLevelCode(v string) *DescribeEnsNetDistrictRequest
	GetNetLevelCode() *string
}

type DescribeEnsNetDistrictRequest struct {
	// The code of the region.
	//
	// If you do not specify this parameter, only nodes in the regions of the level that is specified by the NetLevelCode parameter are queried.
	//
	// If you specify this parameter, only nodes in the regions of the level that is specified by this parameter are queried.
	//
	// example:
	//
	// 100106
	NetDistrictCode     *string `json:"NetDistrictCode,omitempty" xml:"NetDistrictCode,omitempty"`
	NetDistrictCodeNode *bool   `json:"NetDistrictCodeNode,omitempty" xml:"NetDistrictCodeNode,omitempty"`
	// The level of the region.
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

func (s DescribeEnsNetDistrictRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsNetDistrictRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetDistrictRequest) GetNetDistrictCode() *string {
	return s.NetDistrictCode
}

func (s *DescribeEnsNetDistrictRequest) GetNetDistrictCodeNode() *bool {
	return s.NetDistrictCodeNode
}

func (s *DescribeEnsNetDistrictRequest) GetNetLevelCode() *string {
	return s.NetLevelCode
}

func (s *DescribeEnsNetDistrictRequest) SetNetDistrictCode(v string) *DescribeEnsNetDistrictRequest {
	s.NetDistrictCode = &v
	return s
}

func (s *DescribeEnsNetDistrictRequest) SetNetDistrictCodeNode(v bool) *DescribeEnsNetDistrictRequest {
	s.NetDistrictCodeNode = &v
	return s
}

func (s *DescribeEnsNetDistrictRequest) SetNetLevelCode(v string) *DescribeEnsNetDistrictRequest {
	s.NetLevelCode = &v
	return s
}

func (s *DescribeEnsNetDistrictRequest) Validate() error {
	return dara.Validate(s)
}
