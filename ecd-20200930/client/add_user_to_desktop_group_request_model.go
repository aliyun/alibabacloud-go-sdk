// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserToDesktopGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AddUserToDesktopGroupRequest
	GetClientToken() *string
	SetDesktopGroupId(v string) *AddUserToDesktopGroupRequest
	GetDesktopGroupId() *string
	SetDesktopGroupIds(v []*string) *AddUserToDesktopGroupRequest
	GetDesktopGroupIds() []*string
	SetEndUserIds(v []*string) *AddUserToDesktopGroupRequest
	GetEndUserIds() []*string
	SetOrgId(v string) *AddUserToDesktopGroupRequest
	GetOrgId() *string
	SetRegionId(v string) *AddUserToDesktopGroupRequest
	GetRegionId() *string
	SetSimpleUserGroupId(v string) *AddUserToDesktopGroupRequest
	GetSimpleUserGroupId() *string
	SetUserGroupName(v string) *AddUserToDesktopGroupRequest
	GetUserGroupName() *string
	SetUserOuPath(v string) *AddUserToDesktopGroupRequest
	GetUserOuPath() *string
}

type AddUserToDesktopGroupRequest struct {
	// A client token that ensures the idempotence of a request. Generate a unique value for this parameter from your client. The token can contain only ASCII characters and must be no more than 64 characters long. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the shared cloud desktop to which to add an authorized user.
	//
	// example:
	//
	// dg-2i8qxpv6t1a03****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// A list of shared cloud desktop group IDs.
	DesktopGroupIds []*string `json:"DesktopGroupIds,omitempty" xml:"DesktopGroupIds,omitempty" type:"Repeated"`
	// A list of users to authorize.
	EndUserIds []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	OrgId      *string   `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// The region ID. Call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to obtain a list of regions where WUYING Workspace is available.
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

func (s AddUserToDesktopGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AddUserToDesktopGroupRequest) GoString() string {
	return s.String()
}

func (s *AddUserToDesktopGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddUserToDesktopGroupRequest) GetDesktopGroupId() *string {
	return s.DesktopGroupId
}

func (s *AddUserToDesktopGroupRequest) GetDesktopGroupIds() []*string {
	return s.DesktopGroupIds
}

func (s *AddUserToDesktopGroupRequest) GetEndUserIds() []*string {
	return s.EndUserIds
}

func (s *AddUserToDesktopGroupRequest) GetOrgId() *string {
	return s.OrgId
}

func (s *AddUserToDesktopGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddUserToDesktopGroupRequest) GetSimpleUserGroupId() *string {
	return s.SimpleUserGroupId
}

func (s *AddUserToDesktopGroupRequest) GetUserGroupName() *string {
	return s.UserGroupName
}

func (s *AddUserToDesktopGroupRequest) GetUserOuPath() *string {
	return s.UserOuPath
}

func (s *AddUserToDesktopGroupRequest) SetClientToken(v string) *AddUserToDesktopGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *AddUserToDesktopGroupRequest) SetDesktopGroupId(v string) *AddUserToDesktopGroupRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *AddUserToDesktopGroupRequest) SetDesktopGroupIds(v []*string) *AddUserToDesktopGroupRequest {
	s.DesktopGroupIds = v
	return s
}

func (s *AddUserToDesktopGroupRequest) SetEndUserIds(v []*string) *AddUserToDesktopGroupRequest {
	s.EndUserIds = v
	return s
}

func (s *AddUserToDesktopGroupRequest) SetOrgId(v string) *AddUserToDesktopGroupRequest {
	s.OrgId = &v
	return s
}

func (s *AddUserToDesktopGroupRequest) SetRegionId(v string) *AddUserToDesktopGroupRequest {
	s.RegionId = &v
	return s
}

func (s *AddUserToDesktopGroupRequest) SetSimpleUserGroupId(v string) *AddUserToDesktopGroupRequest {
	s.SimpleUserGroupId = &v
	return s
}

func (s *AddUserToDesktopGroupRequest) SetUserGroupName(v string) *AddUserToDesktopGroupRequest {
	s.UserGroupName = &v
	return s
}

func (s *AddUserToDesktopGroupRequest) SetUserOuPath(v string) *AddUserToDesktopGroupRequest {
	s.UserOuPath = &v
	return s
}

func (s *AddUserToDesktopGroupRequest) Validate() error {
	return dara.Validate(s)
}
