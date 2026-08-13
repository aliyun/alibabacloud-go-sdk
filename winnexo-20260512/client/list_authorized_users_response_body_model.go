// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthMode(v string) *ListAuthorizedUsersResponseBody
	GetAuthMode() *string
	SetCode(v string) *ListAuthorizedUsersResponseBody
	GetCode() *string
	SetItems(v []*ListAuthorizedUsersResponseBodyItems) *ListAuthorizedUsersResponseBody
	GetItems() []*ListAuthorizedUsersResponseBodyItems
	SetMessage(v string) *ListAuthorizedUsersResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAuthorizedUsersResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListAuthorizedUsersResponseBody
	GetTotal() *int64
}

type ListAuthorizedUsersResponseBody struct {
	// 授权模式：SPECIFIED_USERS / ALL_USERS
	//
	// example:
	//
	// string_value
	AuthMode *string `json:"authMode,omitempty" xml:"authMode,omitempty"`
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 已授权对象列表
	Items []*ListAuthorizedUsersResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 授权记录总数
	//
	// example:
	//
	// 1
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListAuthorizedUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListAuthorizedUsersResponseBody) GetAuthMode() *string {
	return s.AuthMode
}

func (s *ListAuthorizedUsersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAuthorizedUsersResponseBody) GetItems() []*ListAuthorizedUsersResponseBodyItems {
	return s.Items
}

func (s *ListAuthorizedUsersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAuthorizedUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAuthorizedUsersResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListAuthorizedUsersResponseBody) SetAuthMode(v string) *ListAuthorizedUsersResponseBody {
	s.AuthMode = &v
	return s
}

func (s *ListAuthorizedUsersResponseBody) SetCode(v string) *ListAuthorizedUsersResponseBody {
	s.Code = &v
	return s
}

func (s *ListAuthorizedUsersResponseBody) SetItems(v []*ListAuthorizedUsersResponseBodyItems) *ListAuthorizedUsersResponseBody {
	s.Items = v
	return s
}

func (s *ListAuthorizedUsersResponseBody) SetMessage(v string) *ListAuthorizedUsersResponseBody {
	s.Message = &v
	return s
}

func (s *ListAuthorizedUsersResponseBody) SetRequestId(v string) *ListAuthorizedUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAuthorizedUsersResponseBody) SetTotal(v int64) *ListAuthorizedUsersResponseBody {
	s.Total = &v
	return s
}

func (s *ListAuthorizedUsersResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAuthorizedUsersResponseBodyItems struct {
	// 授权截止时间戳（毫秒）
	//
	// example:
	//
	// 1
	ExpireDate *int64 `json:"expireDate,omitempty" xml:"expireDate,omitempty"`
	// 创建时间
	//
	// example:
	//
	// string_value
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 最后修改时间
	//
	// example:
	//
	// string_value
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// 授权人用户 ID
	//
	// example:
	//
	// 1
	GrantedBy *int64 `json:"grantedBy,omitempty" xml:"grantedBy,omitempty"`
	// 被授权对象 ID
	//
	// example:
	//
	// exampleGranteeId
	GranteeId *string `json:"granteeId,omitempty" xml:"granteeId,omitempty"`
	// 被授权对象类型：USER / USER_GROUP
	//
	// example:
	//
	// string_value
	GranteeType *string `json:"granteeType,omitempty" xml:"granteeType,omitempty"`
	// 授权记录 ID
	//
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 用户组成员数
	//
	// example:
	//
	// 1
	MemberCount *int64 `json:"memberCount,omitempty" xml:"memberCount,omitempty"`
	// 已授权的权限列表
	//
	// example:
	//
	// string_value
	Permissions []*string `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
	// 用户组 ID（granteeType=USER_GROUP 时有值）
	//
	// example:
	//
	// exampleUserGroupId
	UserGroupId *string `json:"userGroupId,omitempty" xml:"userGroupId,omitempty"`
	// 用户组名
	//
	// example:
	//
	// string_value
	UserGroupName *string `json:"userGroupName,omitempty" xml:"userGroupName,omitempty"`
	// 用户 ID（granteeType=USER 时有值）
	//
	// example:
	//
	// 1
	UserId *int64 `json:"userId,omitempty" xml:"userId,omitempty"`
	// 用户名
	//
	// example:
	//
	// string_value
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s ListAuthorizedUsersResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedUsersResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListAuthorizedUsersResponseBodyItems) GetExpireDate() *int64 {
	return s.ExpireDate
}

func (s *ListAuthorizedUsersResponseBodyItems) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListAuthorizedUsersResponseBodyItems) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListAuthorizedUsersResponseBodyItems) GetGrantedBy() *int64 {
	return s.GrantedBy
}

func (s *ListAuthorizedUsersResponseBodyItems) GetGranteeId() *string {
	return s.GranteeId
}

func (s *ListAuthorizedUsersResponseBodyItems) GetGranteeType() *string {
	return s.GranteeType
}

func (s *ListAuthorizedUsersResponseBodyItems) GetId() *int64 {
	return s.Id
}

func (s *ListAuthorizedUsersResponseBodyItems) GetMemberCount() *int64 {
	return s.MemberCount
}

func (s *ListAuthorizedUsersResponseBodyItems) GetPermissions() []*string {
	return s.Permissions
}

func (s *ListAuthorizedUsersResponseBodyItems) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *ListAuthorizedUsersResponseBodyItems) GetUserGroupName() *string {
	return s.UserGroupName
}

func (s *ListAuthorizedUsersResponseBodyItems) GetUserId() *int64 {
	return s.UserId
}

func (s *ListAuthorizedUsersResponseBodyItems) GetUserName() *string {
	return s.UserName
}

func (s *ListAuthorizedUsersResponseBodyItems) SetExpireDate(v int64) *ListAuthorizedUsersResponseBodyItems {
	s.ExpireDate = &v
	return s
}

func (s *ListAuthorizedUsersResponseBodyItems) SetGmtCreate(v string) *ListAuthorizedUsersResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *ListAuthorizedUsersResponseBodyItems) SetGmtModified(v string) *ListAuthorizedUsersResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *ListAuthorizedUsersResponseBodyItems) SetGrantedBy(v int64) *ListAuthorizedUsersResponseBodyItems {
	s.GrantedBy = &v
	return s
}

func (s *ListAuthorizedUsersResponseBodyItems) SetGranteeId(v string) *ListAuthorizedUsersResponseBodyItems {
	s.GranteeId = &v
	return s
}

func (s *ListAuthorizedUsersResponseBodyItems) SetGranteeType(v string) *ListAuthorizedUsersResponseBodyItems {
	s.GranteeType = &v
	return s
}

func (s *ListAuthorizedUsersResponseBodyItems) SetId(v int64) *ListAuthorizedUsersResponseBodyItems {
	s.Id = &v
	return s
}

func (s *ListAuthorizedUsersResponseBodyItems) SetMemberCount(v int64) *ListAuthorizedUsersResponseBodyItems {
	s.MemberCount = &v
	return s
}

func (s *ListAuthorizedUsersResponseBodyItems) SetPermissions(v []*string) *ListAuthorizedUsersResponseBodyItems {
	s.Permissions = v
	return s
}

func (s *ListAuthorizedUsersResponseBodyItems) SetUserGroupId(v string) *ListAuthorizedUsersResponseBodyItems {
	s.UserGroupId = &v
	return s
}

func (s *ListAuthorizedUsersResponseBodyItems) SetUserGroupName(v string) *ListAuthorizedUsersResponseBodyItems {
	s.UserGroupName = &v
	return s
}

func (s *ListAuthorizedUsersResponseBodyItems) SetUserId(v int64) *ListAuthorizedUsersResponseBodyItems {
	s.UserId = &v
	return s
}

func (s *ListAuthorizedUsersResponseBodyItems) SetUserName(v string) *ListAuthorizedUsersResponseBodyItems {
	s.UserName = &v
	return s
}

func (s *ListAuthorizedUsersResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
