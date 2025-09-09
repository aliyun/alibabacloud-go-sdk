// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDisplayName(v string) *ListUsersRequest
	GetDisplayName() *string
	SetInstanceId(v string) *ListUsersRequest
	GetInstanceId() *string
	SetMobile(v string) *ListUsersRequest
	GetMobile() *string
	SetPageNumber(v string) *ListUsersRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListUsersRequest
	GetPageSize() *string
	SetRegionId(v string) *ListUsersRequest
	GetRegionId() *string
	SetSource(v string) *ListUsersRequest
	GetSource() *string
	SetSourceUserId(v string) *ListUsersRequest
	GetSourceUserId() *string
	SetUserGroupId(v string) *ListUsersRequest
	GetUserGroupId() *string
	SetUserName(v string) *ListUsersRequest
	GetUserName() *string
	SetUserState(v string) *ListUsersRequest
	GetUserState() *string
}

type ListUsersRequest struct {
	// The display name of the user that you want to query. Only exact match is supported.
	//
	// example:
	//
	// abc
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the bastion host whose users you want to query.
	//
	// >  You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The mobile phone number of the user that you want to query. Only exact match is supported.
	//
	// example:
	//
	// 1359999****
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.\\
	//
	// Valid values: 1 to 100. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the bastion host whose users you want to query.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the user that you want to query. Valid values:
	//
	// 	- **Local**: a local user.
	//
	// 	- **Ram**: a Resource Access Management (RAM) user.
	//
	// 	- **AD**: an Active Directory (AD)-authenticated user.
	//
	// 	- **LDAP**: a Lightweight Directory Access Protocol (LDAP)-authenticated user.
	//
	// example:
	//
	// Local
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The unique ID of the user that you want to query. Only exact match is supported.
	//
	// >  This parameter uniquely identifies a RAM user of the bastion host. This parameter is valid if **Source*	- is set to **Ram**. You can call the [ListUsers](https://help.aliyun.com/document_detail/28684.html) operation in RAM to obtain the unique ID of the user from the **UserId*	- response parameter.
	//
	// example:
	//
	// 122748924538****
	SourceUserId *string `json:"SourceUserId,omitempty" xml:"SourceUserId,omitempty"`
	// The ID of the user group to which the user you want to query belongs.
	//
	// >  You can call the [ListUserGroups](https://help.aliyun.com/document_detail/204509.html) operation to query the user group ID.
	//
	// example:
	//
	// 1
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// The logon name of the user that you want to query. Only exact match is supported.
	//
	// example:
	//
	// abc
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The state of the user that you want to query. Valid values:
	//
	// 	- **Normal**: The user is in normal state.
	//
	// 	- **Frozen**: The user is locked.
	//
	// 	- **Expired**: The user has expired.
	//
	// example:
	//
	// Normal
	UserState *string `json:"UserState,omitempty" xml:"UserState,omitempty"`
}

func (s ListUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListUsersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListUsersRequest) GetMobile() *string {
	return s.Mobile
}

func (s *ListUsersRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListUsersRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListUsersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListUsersRequest) GetSource() *string {
	return s.Source
}

func (s *ListUsersRequest) GetSourceUserId() *string {
	return s.SourceUserId
}

func (s *ListUsersRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *ListUsersRequest) GetUserName() *string {
	return s.UserName
}

func (s *ListUsersRequest) GetUserState() *string {
	return s.UserState
}

func (s *ListUsersRequest) SetDisplayName(v string) *ListUsersRequest {
	s.DisplayName = &v
	return s
}

func (s *ListUsersRequest) SetInstanceId(v string) *ListUsersRequest {
	s.InstanceId = &v
	return s
}

func (s *ListUsersRequest) SetMobile(v string) *ListUsersRequest {
	s.Mobile = &v
	return s
}

func (s *ListUsersRequest) SetPageNumber(v string) *ListUsersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUsersRequest) SetPageSize(v string) *ListUsersRequest {
	s.PageSize = &v
	return s
}

func (s *ListUsersRequest) SetRegionId(v string) *ListUsersRequest {
	s.RegionId = &v
	return s
}

func (s *ListUsersRequest) SetSource(v string) *ListUsersRequest {
	s.Source = &v
	return s
}

func (s *ListUsersRequest) SetSourceUserId(v string) *ListUsersRequest {
	s.SourceUserId = &v
	return s
}

func (s *ListUsersRequest) SetUserGroupId(v string) *ListUsersRequest {
	s.UserGroupId = &v
	return s
}

func (s *ListUsersRequest) SetUserName(v string) *ListUsersRequest {
	s.UserName = &v
	return s
}

func (s *ListUsersRequest) SetUserState(v string) *ListUsersRequest {
	s.UserState = &v
	return s
}

func (s *ListUsersRequest) Validate() error {
	return dara.Validate(s)
}
