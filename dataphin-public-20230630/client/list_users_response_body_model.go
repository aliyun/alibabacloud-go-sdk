// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListUsersResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListUsersResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListUsersResponseBody
	GetMessage() *string
	SetPageResult(v *ListUsersResponseBodyPageResult) *ListUsersResponseBody
	GetPageResult() *ListUsersResponseBodyPageResult
	SetRequestId(v string) *ListUsersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListUsersResponseBody
	GetSuccess() *bool
}

type ListUsersResponseBody struct {
	// The request error code. OK indicates that the request is successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code of the backend response.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request error message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The paged query result.
	PageResult *ListUsersResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 72F479F0-EE9F-5511-85C1-3AF3B7668E25
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListUsersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListUsersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListUsersResponseBody) GetPageResult() *ListUsersResponseBodyPageResult {
	return s.PageResult
}

func (s *ListUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUsersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListUsersResponseBody) SetCode(v string) *ListUsersResponseBody {
	s.Code = &v
	return s
}

func (s *ListUsersResponseBody) SetHttpStatusCode(v int32) *ListUsersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListUsersResponseBody) SetMessage(v string) *ListUsersResponseBody {
	s.Message = &v
	return s
}

func (s *ListUsersResponseBody) SetPageResult(v *ListUsersResponseBodyPageResult) *ListUsersResponseBody {
	s.PageResult = v
	return s
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersResponseBody) SetSuccess(v bool) *ListUsersResponseBody {
	s.Success = &v
	return s
}

func (s *ListUsersResponseBody) Validate() error {
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListUsersResponseBodyPageResult struct {
	// The total number of records.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The user list.
	UserList []*ListUsersResponseBodyPageResultUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyPageResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListUsersResponseBodyPageResult) GetUserList() []*ListUsersResponseBodyPageResultUserList {
	return s.UserList
}

func (s *ListUsersResponseBodyPageResult) SetTotalCount(v int32) *ListUsersResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListUsersResponseBodyPageResult) SetUserList(v []*ListUsersResponseBodyPageResultUserList) *ListUsersResponseBodyPageResult {
	s.UserList = v
	return s
}

func (s *ListUsersResponseBodyPageResult) Validate() error {
	if s.UserList != nil {
		for _, item := range s.UserList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUsersResponseBodyPageResultUserList struct {
	// The logon name of the original account.
	//
	// example:
	//
	// example-account
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Indicates whether the user is deleted.
	//
	// example:
	//
	// false
	Deleted *bool `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	// The department list.
	DeptList []*ListUsersResponseBodyPageResultUserListDeptList `json:"DeptList,omitempty" xml:"DeptList,omitempty" type:"Repeated"`
	// The DingTalk number.
	//
	// example:
	//
	// example-ding-number
	DingNumber *string `json:"DingNumber,omitempty" xml:"DingNumber,omitempty"`
	// The display name of the original account.
	//
	// example:
	//
	// Sample User
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The display name of the original account without the status suffix.
	//
	// example:
	//
	// Sample User
	DisplayNameWithoutStatus *string `json:"DisplayNameWithoutStatus,omitempty" xml:"DisplayNameWithoutStatus,omitempty"`
	// Indicates whether the IP whitelist is enabled.
	//
	// example:
	//
	// false
	EnableWhiteIp *bool `json:"EnableWhiteIp,omitempty" xml:"EnableWhiteIp,omitempty"`
	// The Lark bot.
	//
	// example:
	//
	// example-feishu-robot
	FeiShuRobot *string `json:"FeiShuRobot,omitempty" xml:"FeiShuRobot,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2026-01-02T00:00:00Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 10001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The email address.
	//
	// example:
	//
	// user@example.com
	Mail *string `json:"Mail,omitempty" xml:"Mail,omitempty"`
	// The mobile phone number.
	//
	// example:
	//
	// 13800000000
	MobilePhone *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	// The user name.
	//
	// example:
	//
	// example-user
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The nickname of the original account.
	//
	// example:
	//
	// Sample Nickname
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// The user ID of the tenant owner.
	//
	// example:
	//
	// 30001011
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The real name of the original account.
	//
	// example:
	//
	// John Doe
	RealName *string `json:"RealName,omitempty" xml:"RealName,omitempty"`
	// The original account ID.
	//
	// example:
	//
	// 10001
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The original account type.
	//
	// example:
	//
	// ALIYUN
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The WeCom bot.
	//
	// example:
	//
	// example-wechat-robot
	WeChatRobot *string `json:"WeChatRobot,omitempty" xml:"WeChatRobot,omitempty"`
	// The IP address whitelist.
	//
	// example:
	//
	// 192.0.2.1
	WhiteIp *string `json:"WhiteIp,omitempty" xml:"WhiteIp,omitempty"`
}

func (s ListUsersResponseBodyPageResultUserList) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBodyPageResultUserList) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyPageResultUserList) GetAccountName() *string {
	return s.AccountName
}

func (s *ListUsersResponseBodyPageResultUserList) GetDeleted() *bool {
	return s.Deleted
}

func (s *ListUsersResponseBodyPageResultUserList) GetDeptList() []*ListUsersResponseBodyPageResultUserListDeptList {
	return s.DeptList
}

func (s *ListUsersResponseBodyPageResultUserList) GetDingNumber() *string {
	return s.DingNumber
}

func (s *ListUsersResponseBodyPageResultUserList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListUsersResponseBodyPageResultUserList) GetDisplayNameWithoutStatus() *string {
	return s.DisplayNameWithoutStatus
}

func (s *ListUsersResponseBodyPageResultUserList) GetEnableWhiteIp() *bool {
	return s.EnableWhiteIp
}

func (s *ListUsersResponseBodyPageResultUserList) GetFeiShuRobot() *string {
	return s.FeiShuRobot
}

func (s *ListUsersResponseBodyPageResultUserList) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListUsersResponseBodyPageResultUserList) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListUsersResponseBodyPageResultUserList) GetId() *string {
	return s.Id
}

func (s *ListUsersResponseBodyPageResultUserList) GetMail() *string {
	return s.Mail
}

func (s *ListUsersResponseBodyPageResultUserList) GetMobilePhone() *string {
	return s.MobilePhone
}

func (s *ListUsersResponseBodyPageResultUserList) GetName() *string {
	return s.Name
}

func (s *ListUsersResponseBodyPageResultUserList) GetNickName() *string {
	return s.NickName
}

func (s *ListUsersResponseBodyPageResultUserList) GetParentId() *string {
	return s.ParentId
}

func (s *ListUsersResponseBodyPageResultUserList) GetRealName() *string {
	return s.RealName
}

func (s *ListUsersResponseBodyPageResultUserList) GetSourceId() *string {
	return s.SourceId
}

func (s *ListUsersResponseBodyPageResultUserList) GetSourceType() *string {
	return s.SourceType
}

func (s *ListUsersResponseBodyPageResultUserList) GetWeChatRobot() *string {
	return s.WeChatRobot
}

func (s *ListUsersResponseBodyPageResultUserList) GetWhiteIp() *string {
	return s.WhiteIp
}

func (s *ListUsersResponseBodyPageResultUserList) SetAccountName(v string) *ListUsersResponseBodyPageResultUserList {
	s.AccountName = &v
	return s
}

func (s *ListUsersResponseBodyPageResultUserList) SetDeleted(v bool) *ListUsersResponseBodyPageResultUserList {
	s.Deleted = &v
	return s
}

func (s *ListUsersResponseBodyPageResultUserList) SetDeptList(v []*ListUsersResponseBodyPageResultUserListDeptList) *ListUsersResponseBodyPageResultUserList {
	s.DeptList = v
	return s
}

func (s *ListUsersResponseBodyPageResultUserList) SetDingNumber(v string) *ListUsersResponseBodyPageResultUserList {
	s.DingNumber = &v
	return s
}

func (s *ListUsersResponseBodyPageResultUserList) SetDisplayName(v string) *ListUsersResponseBodyPageResultUserList {
	s.DisplayName = &v
	return s
}

func (s *ListUsersResponseBodyPageResultUserList) SetDisplayNameWithoutStatus(v string) *ListUsersResponseBodyPageResultUserList {
	s.DisplayNameWithoutStatus = &v
	return s
}

func (s *ListUsersResponseBodyPageResultUserList) SetEnableWhiteIp(v bool) *ListUsersResponseBodyPageResultUserList {
	s.EnableWhiteIp = &v
	return s
}

func (s *ListUsersResponseBodyPageResultUserList) SetFeiShuRobot(v string) *ListUsersResponseBodyPageResultUserList {
	s.FeiShuRobot = &v
	return s
}

func (s *ListUsersResponseBodyPageResultUserList) SetGmtCreateTime(v string) *ListUsersResponseBodyPageResultUserList {
	s.GmtCreateTime = &v
	return s
}

func (s *ListUsersResponseBodyPageResultUserList) SetGmtModifiedTime(v string) *ListUsersResponseBodyPageResultUserList {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListUsersResponseBodyPageResultUserList) SetId(v string) *ListUsersResponseBodyPageResultUserList {
	s.Id = &v
	return s
}

func (s *ListUsersResponseBodyPageResultUserList) SetMail(v string) *ListUsersResponseBodyPageResultUserList {
	s.Mail = &v
	return s
}

func (s *ListUsersResponseBodyPageResultUserList) SetMobilePhone(v string) *ListUsersResponseBodyPageResultUserList {
	s.MobilePhone = &v
	return s
}

func (s *ListUsersResponseBodyPageResultUserList) SetName(v string) *ListUsersResponseBodyPageResultUserList {
	s.Name = &v
	return s
}

func (s *ListUsersResponseBodyPageResultUserList) SetNickName(v string) *ListUsersResponseBodyPageResultUserList {
	s.NickName = &v
	return s
}

func (s *ListUsersResponseBodyPageResultUserList) SetParentId(v string) *ListUsersResponseBodyPageResultUserList {
	s.ParentId = &v
	return s
}

func (s *ListUsersResponseBodyPageResultUserList) SetRealName(v string) *ListUsersResponseBodyPageResultUserList {
	s.RealName = &v
	return s
}

func (s *ListUsersResponseBodyPageResultUserList) SetSourceId(v string) *ListUsersResponseBodyPageResultUserList {
	s.SourceId = &v
	return s
}

func (s *ListUsersResponseBodyPageResultUserList) SetSourceType(v string) *ListUsersResponseBodyPageResultUserList {
	s.SourceType = &v
	return s
}

func (s *ListUsersResponseBodyPageResultUserList) SetWeChatRobot(v string) *ListUsersResponseBodyPageResultUserList {
	s.WeChatRobot = &v
	return s
}

func (s *ListUsersResponseBodyPageResultUserList) SetWhiteIp(v string) *ListUsersResponseBodyPageResultUserList {
	s.WhiteIp = &v
	return s
}

func (s *ListUsersResponseBodyPageResultUserList) Validate() error {
	if s.DeptList != nil {
		for _, item := range s.DeptList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUsersResponseBodyPageResultUserListDeptList struct {
	// The sub-department list.
	Children []map[string]interface{} `json:"Children,omitempty" xml:"Children,omitempty" type:"Repeated"`
	// The full path of the department.
	//
	// example:
	//
	// Sample Company/R&D Department
	FullName *string `json:"FullName,omitempty" xml:"FullName,omitempty"`
	// The department ID.
	//
	// example:
	//
	// dept-001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The department name.
	//
	// example:
	//
	// R&D Department
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parent department ID.
	//
	// example:
	//
	// dept-root
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The parent department list.
	Parents []map[string]interface{} `json:"Parents,omitempty" xml:"Parents,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBodyPageResultUserListDeptList) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBodyPageResultUserListDeptList) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyPageResultUserListDeptList) GetChildren() []map[string]interface{} {
	return s.Children
}

func (s *ListUsersResponseBodyPageResultUserListDeptList) GetFullName() *string {
	return s.FullName
}

func (s *ListUsersResponseBodyPageResultUserListDeptList) GetId() *string {
	return s.Id
}

func (s *ListUsersResponseBodyPageResultUserListDeptList) GetName() *string {
	return s.Name
}

func (s *ListUsersResponseBodyPageResultUserListDeptList) GetParentId() *string {
	return s.ParentId
}

func (s *ListUsersResponseBodyPageResultUserListDeptList) GetParents() []map[string]interface{} {
	return s.Parents
}

func (s *ListUsersResponseBodyPageResultUserListDeptList) SetChildren(v []map[string]interface{}) *ListUsersResponseBodyPageResultUserListDeptList {
	s.Children = v
	return s
}

func (s *ListUsersResponseBodyPageResultUserListDeptList) SetFullName(v string) *ListUsersResponseBodyPageResultUserListDeptList {
	s.FullName = &v
	return s
}

func (s *ListUsersResponseBodyPageResultUserListDeptList) SetId(v string) *ListUsersResponseBodyPageResultUserListDeptList {
	s.Id = &v
	return s
}

func (s *ListUsersResponseBodyPageResultUserListDeptList) SetName(v string) *ListUsersResponseBodyPageResultUserListDeptList {
	s.Name = &v
	return s
}

func (s *ListUsersResponseBodyPageResultUserListDeptList) SetParentId(v string) *ListUsersResponseBodyPageResultUserListDeptList {
	s.ParentId = &v
	return s
}

func (s *ListUsersResponseBodyPageResultUserListDeptList) SetParents(v []map[string]interface{}) *ListUsersResponseBodyPageResultUserListDeptList {
	s.Parents = v
	return s
}

func (s *ListUsersResponseBodyPageResultUserListDeptList) Validate() error {
	return dara.Validate(s)
}
