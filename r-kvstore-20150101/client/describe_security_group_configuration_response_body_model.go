// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityGroupConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeSecurityGroupConfigurationResponseBodyItems) *DescribeSecurityGroupConfigurationResponseBody
	GetItems() *DescribeSecurityGroupConfigurationResponseBodyItems
	SetRequestId(v string) *DescribeSecurityGroupConfigurationResponseBody
	GetRequestId() *string
}

type DescribeSecurityGroupConfigurationResponseBody struct {
	// The list of security groups.
	Items *DescribeSecurityGroupConfigurationResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 981C0D6A-D9DD-466C-92DA-F29DF755****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSecurityGroupConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupConfigurationResponseBody) GetItems() *DescribeSecurityGroupConfigurationResponseBodyItems {
	return s.Items
}

func (s *DescribeSecurityGroupConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSecurityGroupConfigurationResponseBody) SetItems(v *DescribeSecurityGroupConfigurationResponseBodyItems) *DescribeSecurityGroupConfigurationResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSecurityGroupConfigurationResponseBody) SetRequestId(v string) *DescribeSecurityGroupConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityGroupConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSecurityGroupConfigurationResponseBodyItems struct {
	EcsSecurityGroupRelation []*DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation `json:"EcsSecurityGroupRelation,omitempty" xml:"EcsSecurityGroupRelation,omitempty" type:"Repeated"`
}

func (s DescribeSecurityGroupConfigurationResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupConfigurationResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupConfigurationResponseBodyItems) GetEcsSecurityGroupRelation() []*DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation {
	return s.EcsSecurityGroupRelation
}

func (s *DescribeSecurityGroupConfigurationResponseBodyItems) SetEcsSecurityGroupRelation(v []*DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation) *DescribeSecurityGroupConfigurationResponseBodyItems {
	s.EcsSecurityGroupRelation = v
	return s
}

func (s *DescribeSecurityGroupConfigurationResponseBodyItems) Validate() error {
	return dara.Validate(s)
}

type DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation struct {
	// The network type of the security group. Valid values:
	//
	// 	- **classic**: the classic network.
	//
	// 	- **vpc**: the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The ID of the region where the instance is deployed.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-bp14p9y07ns3gwq****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation) GetNetType() *string {
	return s.NetType
}

func (s *DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation) SetNetType(v string) *DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation {
	s.NetType = &v
	return s
}

func (s *DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation) SetRegionId(v string) *DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation {
	s.RegionId = &v
	return s
}

func (s *DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation) SetSecurityGroupId(v string) *DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation) Validate() error {
	return dara.Validate(s)
}
