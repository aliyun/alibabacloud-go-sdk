// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceDLinkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuxVSwitchList(v []*string) *DescribeResourceDLinkResponseBody
	GetAuxVSwitchList() []*string
	SetDestinationCIDRs(v string) *DescribeResourceDLinkResponseBody
	GetDestinationCIDRs() *string
	SetRequestId(v string) *DescribeResourceDLinkResponseBody
	GetRequestId() *string
	SetSecurityGroupId(v string) *DescribeResourceDLinkResponseBody
	GetSecurityGroupId() *string
	SetVSwitchId(v string) *DescribeResourceDLinkResponseBody
	GetVSwitchId() *string
	SetVpcId(v string) *DescribeResourceDLinkResponseBody
	GetVpcId() *string
}

type DescribeResourceDLinkResponseBody struct {
	// The IDs of the secondary vSwitches that are directly connected.
	AuxVSwitchList []*string `json:"AuxVSwitchList,omitempty" xml:"AuxVSwitchList,omitempty" type:"Repeated"`
	// The CIDR blocks of the clients that you want to connect to. After this parameter is specified, the CIDR blocks are added to the back-to-origin route of the server. Either this parameter or the VSwitchIdList parameter can be used to determine CIDR blocks.
	//
	// example:
	//
	// 72.16.0.0/16
	DestinationCIDRs *string `json:"DestinationCIDRs,omitempty" xml:"DestinationCIDRs,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the security group that is directly connected.
	//
	// example:
	//
	// sg-bp1j1z7297hcink9d****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The ID of the primary vSwitch that is directly connected.
	//
	// example:
	//
	// vsw-8vbqn2at0kljjxxxx****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC that is directly connected.
	//
	// example:
	//
	// vpc-uf66uio7md****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeResourceDLinkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceDLinkResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceDLinkResponseBody) GetAuxVSwitchList() []*string {
	return s.AuxVSwitchList
}

func (s *DescribeResourceDLinkResponseBody) GetDestinationCIDRs() *string {
	return s.DestinationCIDRs
}

func (s *DescribeResourceDLinkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeResourceDLinkResponseBody) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeResourceDLinkResponseBody) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeResourceDLinkResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeResourceDLinkResponseBody) SetAuxVSwitchList(v []*string) *DescribeResourceDLinkResponseBody {
	s.AuxVSwitchList = v
	return s
}

func (s *DescribeResourceDLinkResponseBody) SetDestinationCIDRs(v string) *DescribeResourceDLinkResponseBody {
	s.DestinationCIDRs = &v
	return s
}

func (s *DescribeResourceDLinkResponseBody) SetRequestId(v string) *DescribeResourceDLinkResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceDLinkResponseBody) SetSecurityGroupId(v string) *DescribeResourceDLinkResponseBody {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeResourceDLinkResponseBody) SetVSwitchId(v string) *DescribeResourceDLinkResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribeResourceDLinkResponseBody) SetVpcId(v string) *DescribeResourceDLinkResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeResourceDLinkResponseBody) Validate() error {
	return dara.Validate(s)
}
