// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserFromDesktopGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopGroupId(v string) *RemoveUserFromDesktopGroupRequest
	GetDesktopGroupId() *string
	SetDesktopGroupIds(v []*string) *RemoveUserFromDesktopGroupRequest
	GetDesktopGroupIds() []*string
	SetEndUserIds(v []*string) *RemoveUserFromDesktopGroupRequest
	GetEndUserIds() []*string
	SetOrgId(v string) *RemoveUserFromDesktopGroupRequest
	GetOrgId() *string
	SetRegionId(v string) *RemoveUserFromDesktopGroupRequest
	GetRegionId() *string
	SetSimpleUserGroupId(v string) *RemoveUserFromDesktopGroupRequest
	GetSimpleUserGroupId() *string
	SetUserGroupName(v string) *RemoveUserFromDesktopGroupRequest
	GetUserGroupName() *string
	SetUserOuPath(v string) *RemoveUserFromDesktopGroupRequest
	GetUserOuPath() *string
}

type RemoveUserFromDesktopGroupRequest struct {
	// The ID of the shared cloud desktop from which you revoke the user’s permission.
	//
	// example:
	//
	// dg-2i8qxpv6t1a03****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// A list of shared desktop group IDs.
	DesktopGroupIds []*string `json:"DesktopGroupIds,omitempty" xml:"DesktopGroupIds,omitempty" type:"Repeated"`
	// The list of authorized users to remove.
	EndUserIds []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	OrgId      *string   `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// The region ID. Call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to get a list of regions where WUYING Workspace is available.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SimpleUserGroupId *string `json:"SimpleUserGroupId,omitempty" xml:"SimpleUserGroupId,omitempty"`
	UserGroupName     *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
	UserOuPath        *string `json:"UserOuPath,omitempty" xml:"UserOuPath,omitempty"`
}

func (s RemoveUserFromDesktopGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserFromDesktopGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveUserFromDesktopGroupRequest) GetDesktopGroupId() *string {
	return s.DesktopGroupId
}

func (s *RemoveUserFromDesktopGroupRequest) GetDesktopGroupIds() []*string {
	return s.DesktopGroupIds
}

func (s *RemoveUserFromDesktopGroupRequest) GetEndUserIds() []*string {
	return s.EndUserIds
}

func (s *RemoveUserFromDesktopGroupRequest) GetOrgId() *string {
	return s.OrgId
}

func (s *RemoveUserFromDesktopGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RemoveUserFromDesktopGroupRequest) GetSimpleUserGroupId() *string {
	return s.SimpleUserGroupId
}

func (s *RemoveUserFromDesktopGroupRequest) GetUserGroupName() *string {
	return s.UserGroupName
}

func (s *RemoveUserFromDesktopGroupRequest) GetUserOuPath() *string {
	return s.UserOuPath
}

func (s *RemoveUserFromDesktopGroupRequest) SetDesktopGroupId(v string) *RemoveUserFromDesktopGroupRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *RemoveUserFromDesktopGroupRequest) SetDesktopGroupIds(v []*string) *RemoveUserFromDesktopGroupRequest {
	s.DesktopGroupIds = v
	return s
}

func (s *RemoveUserFromDesktopGroupRequest) SetEndUserIds(v []*string) *RemoveUserFromDesktopGroupRequest {
	s.EndUserIds = v
	return s
}

func (s *RemoveUserFromDesktopGroupRequest) SetOrgId(v string) *RemoveUserFromDesktopGroupRequest {
	s.OrgId = &v
	return s
}

func (s *RemoveUserFromDesktopGroupRequest) SetRegionId(v string) *RemoveUserFromDesktopGroupRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveUserFromDesktopGroupRequest) SetSimpleUserGroupId(v string) *RemoveUserFromDesktopGroupRequest {
	s.SimpleUserGroupId = &v
	return s
}

func (s *RemoveUserFromDesktopGroupRequest) SetUserGroupName(v string) *RemoveUserFromDesktopGroupRequest {
	s.UserGroupName = &v
	return s
}

func (s *RemoveUserFromDesktopGroupRequest) SetUserOuPath(v string) *RemoveUserFromDesktopGroupRequest {
	s.UserOuPath = &v
	return s
}

func (s *RemoveUserFromDesktopGroupRequest) Validate() error {
	return dara.Validate(s)
}
