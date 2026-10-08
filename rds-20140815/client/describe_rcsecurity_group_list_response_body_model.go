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
	RCSecurityGroups []*DescribeRCSecurityGroupListResponseBodyRCSecurityGroups `json:"RCSecurityGroups,omitempty" xml:"RCSecurityGroups,omitempty" type:"Repeated"`
	RequestId        *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AvailableInstanceAmount *int32  `json:"AvailableInstanceAmount,omitempty" xml:"AvailableInstanceAmount,omitempty"`
	CreationTime            *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	InstanceCount     *int32  `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	SecurityGroupId   *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SecurityGroupType *string `json:"SecurityGroupType,omitempty" xml:"SecurityGroupType,omitempty"`
	VpcId             *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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
