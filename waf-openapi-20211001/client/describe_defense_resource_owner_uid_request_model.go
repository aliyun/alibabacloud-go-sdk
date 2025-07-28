// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseResourceOwnerUidRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeDefenseResourceOwnerUidRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeDefenseResourceOwnerUidRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeDefenseResourceOwnerUidRequest
	GetResourceManagerResourceGroupId() *string
	SetResourceNames(v string) *DescribeDefenseResourceOwnerUidRequest
	GetResourceNames() *string
}

type DescribeDefenseResourceOwnerUidRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn********60f
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a.com-waf,b.com-waf
	ResourceNames *string `json:"ResourceNames,omitempty" xml:"ResourceNames,omitempty"`
}

func (s DescribeDefenseResourceOwnerUidRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseResourceOwnerUidRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceOwnerUidRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDefenseResourceOwnerUidRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDefenseResourceOwnerUidRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeDefenseResourceOwnerUidRequest) GetResourceNames() *string {
	return s.ResourceNames
}

func (s *DescribeDefenseResourceOwnerUidRequest) SetInstanceId(v string) *DescribeDefenseResourceOwnerUidRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDefenseResourceOwnerUidRequest) SetRegionId(v string) *DescribeDefenseResourceOwnerUidRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDefenseResourceOwnerUidRequest) SetResourceManagerResourceGroupId(v string) *DescribeDefenseResourceOwnerUidRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeDefenseResourceOwnerUidRequest) SetResourceNames(v string) *DescribeDefenseResourceOwnerUidRequest {
	s.ResourceNames = &v
	return s
}

func (s *DescribeDefenseResourceOwnerUidRequest) Validate() error {
	return dara.Validate(s)
}
