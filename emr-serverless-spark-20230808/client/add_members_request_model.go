// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMemberArns(v []*string) *AddMembersRequest
	GetMemberArns() []*string
	SetWorkspaceId(v string) *AddMembersRequest
	GetWorkspaceId() *string
	SetRegionId(v string) *AddMembersRequest
	GetRegionId() *string
}

type AddMembersRequest struct {
	// This parameter is required.
	MemberArns []*string `json:"memberArns,omitempty" xml:"memberArns,omitempty" type:"Repeated"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// w-975bcfda9625****
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s AddMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s AddMembersRequest) GoString() string {
	return s.String()
}

func (s *AddMembersRequest) GetMemberArns() []*string {
	return s.MemberArns
}

func (s *AddMembersRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *AddMembersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddMembersRequest) SetMemberArns(v []*string) *AddMembersRequest {
	s.MemberArns = v
	return s
}

func (s *AddMembersRequest) SetWorkspaceId(v string) *AddMembersRequest {
	s.WorkspaceId = &v
	return s
}

func (s *AddMembersRequest) SetRegionId(v string) *AddMembersRequest {
	s.RegionId = &v
	return s
}

func (s *AddMembersRequest) Validate() error {
	return dara.Validate(s)
}
