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
	// testuser
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The instance ID of the bastion host for which you want to query the user list.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
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
	// The page number of the current page in a paging query. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of entries per page in a paging query.
	//
	// The maximum value of the PageSize parameter is 100. The default number of entries per page is 20. If PageSize is left empty, 20 entries are returned by default.
	//
	// > Do not leave PageSize empty.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the bastion host for which you want to query the user list.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The source of the user that you want to query. Valid values:
	//
	// - **Local**: local user
	//
	// - **Ram**: Resource Access Management (RAM) user
	//
	// - **AD**: AD user
	//
	// - **LDAP**: LDAP user
	//
	// example:
	//
	// Local
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The unique identity of the user that you want to query. Only exact match is supported.
	//
	// > This parameter is the unique identity of the Resource Access Management (RAM) user that corresponds to the bastion host user. This parameter takes effect when the source of the newly created user is a RAM user (that is, **Source*	- is set to **Ram**). You can invoke the [ListUsers](https://help.aliyun.com/document_detail/28684.html) operation of access control and obtain this parameter from the **UserId*	- field in the response.
	//
	// example:
	//
	// 122748924538****
	SourceUserId *string `json:"SourceUserId,omitempty" xml:"SourceUserId,omitempty"`
	// The ID of the user group that you want to query.
	//
	// > You can call the [ListUserGroups](https://help.aliyun.com/document_detail/204509.html) operation to obtain this parameter.
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
	// The status of the user that you want to query. Valid values:
	//
	// - **Normal**: normal
	//
	// - **Frozen**: locked
	//
	// - **Expired**: expired
	//
	// - **RemoteDeleted**: user source deleted
	//
	// - **Inactive**: inactive due to prolonged absence of logon
	//
	// - **PasswordExpired**: password expired
	//
	// - **RemoteDNChanged**: user DN updated
	//
	// - **RemoteFrozen**: frozen on the RAM side
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
