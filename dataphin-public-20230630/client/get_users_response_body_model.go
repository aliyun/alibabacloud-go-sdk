// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetUsersResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetUsersResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetUsersResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetUsersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetUsersResponseBody
	GetSuccess() *bool
	SetUserList(v []*GetUsersResponseBodyUserList) *GetUsersResponseBody
	GetUserList() []*GetUsersResponseBodyUserList
}

type GetUsersResponseBody struct {
	// The error code. A value of OK indicates that the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code returned by the backend.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The list of users.
	UserList []*GetUsersResponseBodyUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
}

func (s GetUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUsersResponseBody) GoString() string {
	return s.String()
}

func (s *GetUsersResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetUsersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetUsersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUsersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetUsersResponseBody) GetUserList() []*GetUsersResponseBodyUserList {
	return s.UserList
}

func (s *GetUsersResponseBody) SetCode(v string) *GetUsersResponseBody {
	s.Code = &v
	return s
}

func (s *GetUsersResponseBody) SetHttpStatusCode(v int32) *GetUsersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetUsersResponseBody) SetMessage(v string) *GetUsersResponseBody {
	s.Message = &v
	return s
}

func (s *GetUsersResponseBody) SetRequestId(v string) *GetUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUsersResponseBody) SetSuccess(v bool) *GetUsersResponseBody {
	s.Success = &v
	return s
}

func (s *GetUsersResponseBody) SetUserList(v []*GetUsersResponseBodyUserList) *GetUsersResponseBody {
	s.UserList = v
	return s
}

func (s *GetUsersResponseBody) Validate() error {
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

type GetUsersResponseBodyUserList struct {
	// The account name.
	//
	// example:
	//
	// 123@xx.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The DingTalk number.
	//
	// example:
	//
	// 123@dingding
	DingNumber *string `json:"DingNumber,omitempty" xml:"DingNumber,omitempty"`
	// The display name of the user.
	//
	// example:
	//
	// xx测试
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The display name of the user without status.
	//
	// example:
	//
	// xx测试
	DisplayNameWithoutStatus *string `json:"DisplayNameWithoutStatus,omitempty" xml:"DisplayNameWithoutStatus,omitempty"`
	// Indicates whether the IP whitelist is enabled.
	//
	// example:
	//
	// true
	EnableWhiteIp *string `json:"EnableWhiteIp,omitempty" xml:"EnableWhiteIp,omitempty"`
	// The Lark chatbot.
	//
	// example:
	//
	// xx
	FeiShuRobot *string `json:"FeiShuRobot,omitempty" xml:"FeiShuRobot,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1717343597000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The last modification time.
	//
	// example:
	//
	// 1717343597000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 1233121
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The email address.
	//
	// example:
	//
	// 123@xx.com
	Mail *string `json:"Mail,omitempty" xml:"Mail,omitempty"`
	// The phone number.
	//
	// example:
	//
	// 1388888888
	MobilePhone *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	// The username.
	//
	// example:
	//
	// xx测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The nickname of the user.
	//
	// example:
	//
	// xx测试
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// The ID of the primary account to which the user belongs.
	//
	// example:
	//
	// 231231
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The real name of the user.
	//
	// example:
	//
	// xx测试
	RealName *string `json:"RealName,omitempty" xml:"RealName,omitempty"`
	// The source ID of the user.
	//
	// example:
	//
	// 123@xx.com
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The source type of the user.
	//
	// example:
	//
	// ALIYUN
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The WeCom chatbot.
	//
	// example:
	//
	// xx
	WeChatRobot *string `json:"WeChatRobot,omitempty" xml:"WeChatRobot,omitempty"`
	// The IP whitelist.
	//
	// example:
	//
	// *
	WhiteIp *string `json:"WhiteIp,omitempty" xml:"WhiteIp,omitempty"`
}

func (s GetUsersResponseBodyUserList) String() string {
	return dara.Prettify(s)
}

func (s GetUsersResponseBodyUserList) GoString() string {
	return s.String()
}

func (s *GetUsersResponseBodyUserList) GetAccountName() *string {
	return s.AccountName
}

func (s *GetUsersResponseBodyUserList) GetDingNumber() *string {
	return s.DingNumber
}

func (s *GetUsersResponseBodyUserList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetUsersResponseBodyUserList) GetDisplayNameWithoutStatus() *string {
	return s.DisplayNameWithoutStatus
}

func (s *GetUsersResponseBodyUserList) GetEnableWhiteIp() *string {
	return s.EnableWhiteIp
}

func (s *GetUsersResponseBodyUserList) GetFeiShuRobot() *string {
	return s.FeiShuRobot
}

func (s *GetUsersResponseBodyUserList) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *GetUsersResponseBodyUserList) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *GetUsersResponseBodyUserList) GetId() *string {
	return s.Id
}

func (s *GetUsersResponseBodyUserList) GetMail() *string {
	return s.Mail
}

func (s *GetUsersResponseBodyUserList) GetMobilePhone() *string {
	return s.MobilePhone
}

func (s *GetUsersResponseBodyUserList) GetName() *string {
	return s.Name
}

func (s *GetUsersResponseBodyUserList) GetNickName() *string {
	return s.NickName
}

func (s *GetUsersResponseBodyUserList) GetParentId() *string {
	return s.ParentId
}

func (s *GetUsersResponseBodyUserList) GetRealName() *string {
	return s.RealName
}

func (s *GetUsersResponseBodyUserList) GetSourceId() *string {
	return s.SourceId
}

func (s *GetUsersResponseBodyUserList) GetSourceType() *string {
	return s.SourceType
}

func (s *GetUsersResponseBodyUserList) GetWeChatRobot() *string {
	return s.WeChatRobot
}

func (s *GetUsersResponseBodyUserList) GetWhiteIp() *string {
	return s.WhiteIp
}

func (s *GetUsersResponseBodyUserList) SetAccountName(v string) *GetUsersResponseBodyUserList {
	s.AccountName = &v
	return s
}

func (s *GetUsersResponseBodyUserList) SetDingNumber(v string) *GetUsersResponseBodyUserList {
	s.DingNumber = &v
	return s
}

func (s *GetUsersResponseBodyUserList) SetDisplayName(v string) *GetUsersResponseBodyUserList {
	s.DisplayName = &v
	return s
}

func (s *GetUsersResponseBodyUserList) SetDisplayNameWithoutStatus(v string) *GetUsersResponseBodyUserList {
	s.DisplayNameWithoutStatus = &v
	return s
}

func (s *GetUsersResponseBodyUserList) SetEnableWhiteIp(v string) *GetUsersResponseBodyUserList {
	s.EnableWhiteIp = &v
	return s
}

func (s *GetUsersResponseBodyUserList) SetFeiShuRobot(v string) *GetUsersResponseBodyUserList {
	s.FeiShuRobot = &v
	return s
}

func (s *GetUsersResponseBodyUserList) SetGmtCreate(v int64) *GetUsersResponseBodyUserList {
	s.GmtCreate = &v
	return s
}

func (s *GetUsersResponseBodyUserList) SetGmtModified(v int64) *GetUsersResponseBodyUserList {
	s.GmtModified = &v
	return s
}

func (s *GetUsersResponseBodyUserList) SetId(v string) *GetUsersResponseBodyUserList {
	s.Id = &v
	return s
}

func (s *GetUsersResponseBodyUserList) SetMail(v string) *GetUsersResponseBodyUserList {
	s.Mail = &v
	return s
}

func (s *GetUsersResponseBodyUserList) SetMobilePhone(v string) *GetUsersResponseBodyUserList {
	s.MobilePhone = &v
	return s
}

func (s *GetUsersResponseBodyUserList) SetName(v string) *GetUsersResponseBodyUserList {
	s.Name = &v
	return s
}

func (s *GetUsersResponseBodyUserList) SetNickName(v string) *GetUsersResponseBodyUserList {
	s.NickName = &v
	return s
}

func (s *GetUsersResponseBodyUserList) SetParentId(v string) *GetUsersResponseBodyUserList {
	s.ParentId = &v
	return s
}

func (s *GetUsersResponseBodyUserList) SetRealName(v string) *GetUsersResponseBodyUserList {
	s.RealName = &v
	return s
}

func (s *GetUsersResponseBodyUserList) SetSourceId(v string) *GetUsersResponseBodyUserList {
	s.SourceId = &v
	return s
}

func (s *GetUsersResponseBodyUserList) SetSourceType(v string) *GetUsersResponseBodyUserList {
	s.SourceType = &v
	return s
}

func (s *GetUsersResponseBodyUserList) SetWeChatRobot(v string) *GetUsersResponseBodyUserList {
	s.WeChatRobot = &v
	return s
}

func (s *GetUsersResponseBodyUserList) SetWhiteIp(v string) *GetUsersResponseBodyUserList {
	s.WhiteIp = &v
	return s
}

func (s *GetUsersResponseBodyUserList) Validate() error {
	return dara.Validate(s)
}
