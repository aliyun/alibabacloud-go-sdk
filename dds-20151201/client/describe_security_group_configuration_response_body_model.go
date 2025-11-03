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
	// Details about the ECS security groups.
	Items *DescribeSecurityGroupConfigurationResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 3C4A2494-85C4-45C5-93CF-548DB3375193
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
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSecurityGroupConfigurationResponseBodyItems struct {
	RdsEcsSecurityGroupRel []*DescribeSecurityGroupConfigurationResponseBodyItemsRdsEcsSecurityGroupRel `json:"RdsEcsSecurityGroupRel,omitempty" xml:"RdsEcsSecurityGroupRel,omitempty" type:"Repeated"`
}

func (s DescribeSecurityGroupConfigurationResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupConfigurationResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupConfigurationResponseBodyItems) GetRdsEcsSecurityGroupRel() []*DescribeSecurityGroupConfigurationResponseBodyItemsRdsEcsSecurityGroupRel {
	return s.RdsEcsSecurityGroupRel
}

func (s *DescribeSecurityGroupConfigurationResponseBodyItems) SetRdsEcsSecurityGroupRel(v []*DescribeSecurityGroupConfigurationResponseBodyItemsRdsEcsSecurityGroupRel) *DescribeSecurityGroupConfigurationResponseBodyItems {
	s.RdsEcsSecurityGroupRel = v
	return s
}

func (s *DescribeSecurityGroupConfigurationResponseBodyItems) Validate() error {
	if s.RdsEcsSecurityGroupRel != nil {
		for _, item := range s.RdsEcsSecurityGroupRel {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSecurityGroupConfigurationResponseBodyItemsRdsEcsSecurityGroupRel struct {
	// The network type of the ECS security group. Valid values:
	//
	// 	- **vpc**
	//
	// 	- **classic**
	//
	// example:
	//
	// vpc
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The region ID of the ECS security group.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the ECS security group.
	//
	// example:
	//
	// sg-bpxxxxxxxx
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s DescribeSecurityGroupConfigurationResponseBodyItemsRdsEcsSecurityGroupRel) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupConfigurationResponseBodyItemsRdsEcsSecurityGroupRel) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupConfigurationResponseBodyItemsRdsEcsSecurityGroupRel) GetNetType() *string {
	return s.NetType
}

func (s *DescribeSecurityGroupConfigurationResponseBodyItemsRdsEcsSecurityGroupRel) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSecurityGroupConfigurationResponseBodyItemsRdsEcsSecurityGroupRel) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeSecurityGroupConfigurationResponseBodyItemsRdsEcsSecurityGroupRel) SetNetType(v string) *DescribeSecurityGroupConfigurationResponseBodyItemsRdsEcsSecurityGroupRel {
	s.NetType = &v
	return s
}

func (s *DescribeSecurityGroupConfigurationResponseBodyItemsRdsEcsSecurityGroupRel) SetRegionId(v string) *DescribeSecurityGroupConfigurationResponseBodyItemsRdsEcsSecurityGroupRel {
	s.RegionId = &v
	return s
}

func (s *DescribeSecurityGroupConfigurationResponseBodyItemsRdsEcsSecurityGroupRel) SetSecurityGroupId(v string) *DescribeSecurityGroupConfigurationResponseBodyItemsRdsEcsSecurityGroupRel {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeSecurityGroupConfigurationResponseBodyItemsRdsEcsSecurityGroupRel) Validate() error {
	return dara.Validate(s)
}
