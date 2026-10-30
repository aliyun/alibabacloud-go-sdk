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
	// The region code.
	//
	// If this parameter is left empty, only the edge zones under the specified NetLevelCode are queried.
	//
	// If this parameter is specified, only the child edge zones under the specified NetDistrictCode are queried.
	//
	// example:
	//
	// 100106
	NetDistrictCode *string `json:"NetDistrictCode,omitempty" xml:"NetDistrictCode,omitempty"`
	// The region code.
	NetDistrictCodeNode *bool `json:"NetDistrictCodeNode,omitempty" xml:"NetDistrictCodeNode,omitempty"`
	// The region level. Valid values:
	//
	// - **Big**: Area.
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
