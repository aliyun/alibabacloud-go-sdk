// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachGadInstanceMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGadInstanceName(v string) *DetachGadInstanceMemberRequest
	GetGadInstanceName() *string
	SetMemberInstanceName(v string) *DetachGadInstanceMemberRequest
	GetMemberInstanceName() *string
	SetRegionId(v string) *DetachGadInstanceMemberRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DetachGadInstanceMemberRequest
	GetResourceGroupId() *string
}

type DetachGadInstanceMemberRequest struct {
	// The ID of the global active database cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// gad-rm-bp1npi2j8********
	GadInstanceName *string `json:"GadInstanceName,omitempty" xml:"GadInstanceName,omitempty"`
	// The ID of the instance that serves as the unit node you want to remove. You can call the DescribeGadInstances query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp1npi2j8********
	MemberInstanceName *string `json:"MemberInstanceName,omitempty" xml:"MemberInstanceName,omitempty"`
	// The region ID of the central node. You can call the DescribeGadInstances operation to query the region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. You can call the DescribeDBInstanceAttribute operation to query the resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DetachGadInstanceMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachGadInstanceMemberRequest) GoString() string {
	return s.String()
}

func (s *DetachGadInstanceMemberRequest) GetGadInstanceName() *string {
	return s.GadInstanceName
}

func (s *DetachGadInstanceMemberRequest) GetMemberInstanceName() *string {
	return s.MemberInstanceName
}

func (s *DetachGadInstanceMemberRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachGadInstanceMemberRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DetachGadInstanceMemberRequest) SetGadInstanceName(v string) *DetachGadInstanceMemberRequest {
	s.GadInstanceName = &v
	return s
}

func (s *DetachGadInstanceMemberRequest) SetMemberInstanceName(v string) *DetachGadInstanceMemberRequest {
	s.MemberInstanceName = &v
	return s
}

func (s *DetachGadInstanceMemberRequest) SetRegionId(v string) *DetachGadInstanceMemberRequest {
	s.RegionId = &v
	return s
}

func (s *DetachGadInstanceMemberRequest) SetResourceGroupId(v string) *DetachGadInstanceMemberRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DetachGadInstanceMemberRequest) Validate() error {
	return dara.Validate(s)
}
