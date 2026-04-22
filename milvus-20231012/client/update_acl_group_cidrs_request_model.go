// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAclGroupCidrsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *UpdateAclGroupCidrsRequest
	GetGroupName() *string
	SetInstanceId(v string) *UpdateAclGroupCidrsRequest
	GetInstanceId() *string
	SetNewCidrs(v string) *UpdateAclGroupCidrsRequest
	GetNewCidrs() *string
}

type UpdateAclGroupCidrsRequest struct {
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// example:
	//
	// c-xxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// 140.205.11.0/24,140.205.11.2
	NewCidrs *string `json:"newCidrs,omitempty" xml:"newCidrs,omitempty"`
}

func (s UpdateAclGroupCidrsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAclGroupCidrsRequest) GoString() string {
	return s.String()
}

func (s *UpdateAclGroupCidrsRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateAclGroupCidrsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateAclGroupCidrsRequest) GetNewCidrs() *string {
	return s.NewCidrs
}

func (s *UpdateAclGroupCidrsRequest) SetGroupName(v string) *UpdateAclGroupCidrsRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateAclGroupCidrsRequest) SetInstanceId(v string) *UpdateAclGroupCidrsRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateAclGroupCidrsRequest) SetNewCidrs(v string) *UpdateAclGroupCidrsRequest {
	s.NewCidrs = &v
	return s
}

func (s *UpdateAclGroupCidrsRequest) Validate() error {
	return dara.Validate(s)
}
