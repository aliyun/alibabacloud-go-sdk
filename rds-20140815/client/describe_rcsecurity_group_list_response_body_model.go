// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCSecurityGroupListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRCSecurityGroups(v []*DescribeRCSecurityGroupListResponseBodyRCSecurityGroups) *DescribeRCSecurityGroupListResponseBody
	GetRCSecurityGroups() []*DescribeRCSecurityGroupListResponseBodyRCSecurityGroups
	SetRequestId(v string) *DescribeRCSecurityGroupListResponseBody
	GetRequestId() *string
}

type DescribeRCSecurityGroupListResponseBody struct {
	// The basic information about the security groups.
	RCSecurityGroups []*DescribeRCSecurityGroupListResponseBodyRCSecurityGroups `json:"RCSecurityGroups,omitempty" xml:"RCSecurityGroups,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 7A41C147-C8D0-4DAE-A1A2-17EBCD60DFA1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRCSecurityGroupListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCSecurityGroupListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRCSecurityGroupListResponseBody) GetRCSecurityGroups() []*DescribeRCSecurityGroupListResponseBodyRCSecurityGroups {
	return s.RCSecurityGroups
}

func (s *DescribeRCSecurityGroupListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRCSecurityGroupListResponseBody) SetRCSecurityGroups(v []*DescribeRCSecurityGroupListResponseBodyRCSecurityGroups) *DescribeRCSecurityGroupListResponseBody {
	s.RCSecurityGroups = v
	return s
}

func (s *DescribeRCSecurityGroupListResponseBody) SetRequestId(v string) *DescribeRCSecurityGroupListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRCSecurityGroupListResponseBody) Validate() error {
	if s.RCSecurityGroups != nil {
		for _, item := range s.RCSecurityGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRCSecurityGroupListResponseBodyRCSecurityGroups struct {
	// The number of instances that can be added to the security group.
	//
	// example:
	//
	// 48
	AvailableInstanceAmount *int32 `json:"AvailableInstanceAmount,omitempty" xml:"AvailableInstanceAmount,omitempty"`
	// The time when the security group was created. The time follows the ISO 8601 standard and is in the `yyyy-MM-ddThh:mmZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2025-05-31T03:12:29Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the security group.
	//
	// example:
	//
	// TestDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The number of instances that are added to the security group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	InstanceCount *int32 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-2ze27hs990o2hn9****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The type of the security group. Valid values:
	//
	// 	- **normal**: a normal security group.
	//
	// 	- **enterprise**: an advanced security group.
	//
	// example:
	//
	// normal
	SecurityGroupType *string `json:"SecurityGroupType,omitempty" xml:"SecurityGroupType,omitempty"`
	// The ID of the VPC to which the security group belongs.
	//
	// example:
	//
	// vpc-bp1opxu1zkhn****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeRCSecurityGroupListResponseBodyRCSecurityGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCSecurityGroupListResponseBodyRCSecurityGroups) GoString() string {
	return s.String()
}

func (s *DescribeRCSecurityGroupListResponseBodyRCSecurityGroups) GetAvailableInstanceAmount() *int32 {
	return s.AvailableInstanceAmount
}

func (s *DescribeRCSecurityGroupListResponseBodyRCSecurityGroups) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeRCSecurityGroupListResponseBodyRCSecurityGroups) GetDescription() *string {
	return s.Description
}

func (s *DescribeRCSecurityGroupListResponseBodyRCSecurityGroups) GetInstanceCount() *int32 {
	return s.InstanceCount
}

func (s *DescribeRCSecurityGroupListResponseBodyRCSecurityGroups) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeRCSecurityGroupListResponseBodyRCSecurityGroups) GetSecurityGroupType() *string {
	return s.SecurityGroupType
}

func (s *DescribeRCSecurityGroupListResponseBodyRCSecurityGroups) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeRCSecurityGroupListResponseBodyRCSecurityGroups) SetAvailableInstanceAmount(v int32) *DescribeRCSecurityGroupListResponseBodyRCSecurityGroups {
	s.AvailableInstanceAmount = &v
	return s
}

func (s *DescribeRCSecurityGroupListResponseBodyRCSecurityGroups) SetCreationTime(v string) *DescribeRCSecurityGroupListResponseBodyRCSecurityGroups {
	s.CreationTime = &v
	return s
}

func (s *DescribeRCSecurityGroupListResponseBodyRCSecurityGroups) SetDescription(v string) *DescribeRCSecurityGroupListResponseBodyRCSecurityGroups {
	s.Description = &v
	return s
}

func (s *DescribeRCSecurityGroupListResponseBodyRCSecurityGroups) SetInstanceCount(v int32) *DescribeRCSecurityGroupListResponseBodyRCSecurityGroups {
	s.InstanceCount = &v
	return s
}

func (s *DescribeRCSecurityGroupListResponseBodyRCSecurityGroups) SetSecurityGroupId(v string) *DescribeRCSecurityGroupListResponseBodyRCSecurityGroups {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeRCSecurityGroupListResponseBodyRCSecurityGroups) SetSecurityGroupType(v string) *DescribeRCSecurityGroupListResponseBodyRCSecurityGroups {
	s.SecurityGroupType = &v
	return s
}

func (s *DescribeRCSecurityGroupListResponseBodyRCSecurityGroups) SetVpcId(v string) *DescribeRCSecurityGroupListResponseBodyRCSecurityGroups {
	s.VpcId = &v
	return s
}

func (s *DescribeRCSecurityGroupListResponseBodyRCSecurityGroups) Validate() error {
	return dara.Validate(s)
}
