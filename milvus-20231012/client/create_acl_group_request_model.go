// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAclGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCidrs(v string) *CreateAclGroupRequest
	GetCidrs() *string
	SetGroupName(v string) *CreateAclGroupRequest
	GetGroupName() *string
	SetInstanceId(v string) *CreateAclGroupRequest
	GetInstanceId() *string
	SetRegionId(v string) *CreateAclGroupRequest
	GetRegionId() *string
}

type CreateAclGroupRequest struct {
	// example:
	//
	// 140.205.11.0/24,140.205.11.2
	Cidrs *string `json:"cidrs,omitempty" xml:"cidrs,omitempty"`
	// example:
	//
	// test
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// example:
	//
	// c-xxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s CreateAclGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAclGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateAclGroupRequest) GetCidrs() *string {
	return s.Cidrs
}

func (s *CreateAclGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateAclGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateAclGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAclGroupRequest) SetCidrs(v string) *CreateAclGroupRequest {
	s.Cidrs = &v
	return s
}

func (s *CreateAclGroupRequest) SetGroupName(v string) *CreateAclGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateAclGroupRequest) SetInstanceId(v string) *CreateAclGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateAclGroupRequest) SetRegionId(v string) *CreateAclGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAclGroupRequest) Validate() error {
	return dara.Validate(s)
}
