// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAclResourceNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclResourcePatternType(v string) *DescribeAclResourceNameRequest
	GetAclResourcePatternType() *string
	SetAclResourceType(v string) *DescribeAclResourceNameRequest
	GetAclResourceType() *string
	SetInstanceId(v string) *DescribeAclResourceNameRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeAclResourceNameRequest
	GetRegionId() *string
}

type DescribeAclResourceNameRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// LITERAL
	AclResourcePatternType *string `json:"AclResourcePatternType,omitempty" xml:"AclResourcePatternType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Topic
	AclResourceType *string `json:"AclResourceType,omitempty" xml:"AclResourceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-mp919o4v****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAclResourceNameRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclResourceNameRequest) GoString() string {
	return s.String()
}

func (s *DescribeAclResourceNameRequest) GetAclResourcePatternType() *string {
	return s.AclResourcePatternType
}

func (s *DescribeAclResourceNameRequest) GetAclResourceType() *string {
	return s.AclResourceType
}

func (s *DescribeAclResourceNameRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAclResourceNameRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAclResourceNameRequest) SetAclResourcePatternType(v string) *DescribeAclResourceNameRequest {
	s.AclResourcePatternType = &v
	return s
}

func (s *DescribeAclResourceNameRequest) SetAclResourceType(v string) *DescribeAclResourceNameRequest {
	s.AclResourceType = &v
	return s
}

func (s *DescribeAclResourceNameRequest) SetInstanceId(v string) *DescribeAclResourceNameRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAclResourceNameRequest) SetRegionId(v string) *DescribeAclResourceNameRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAclResourceNameRequest) Validate() error {
	return dara.Validate(s)
}
