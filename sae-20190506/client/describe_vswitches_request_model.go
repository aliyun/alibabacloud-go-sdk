// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVSwitchesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVSwitchId(v string) *DescribeVSwitchesRequest
	GetVSwitchId() *string
	SetVSwitchName(v string) *DescribeVSwitchesRequest
	GetVSwitchName() *string
	SetVpcId(v string) *DescribeVSwitchesRequest
	GetVpcId() *string
}

type DescribeVSwitchesRequest struct {
	// example:
	//
	// vsw-bp165bn1eu4fyztdvbfvo,vsw-bp1iof8x4ypzrwfwg1h08
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// KKKK
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-wz97dyzl4m3dtuacb36ox
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeVSwitchesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchesRequest) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeVSwitchesRequest) GetVSwitchName() *string {
	return s.VSwitchName
}

func (s *DescribeVSwitchesRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVSwitchesRequest) SetVSwitchId(v string) *DescribeVSwitchesRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetVSwitchName(v string) *DescribeVSwitchesRequest {
	s.VSwitchName = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetVpcId(v string) *DescribeVSwitchesRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeVSwitchesRequest) Validate() error {
	return dara.Validate(s)
}
