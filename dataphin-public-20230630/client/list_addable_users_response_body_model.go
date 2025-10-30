// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAddableUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAddableUsersResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListAddableUsersResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListAddableUsersResponseBody
	GetMessage() *string
	SetPageResult(v *ListAddableUsersResponseBodyPageResult) *ListAddableUsersResponseBody
	GetPageResult() *ListAddableUsersResponseBodyPageResult
	SetRequestId(v string) *ListAddableUsersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAddableUsersResponseBody
	GetSuccess() *bool
}

type ListAddableUsersResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message    *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	PageResult *ListAddableUsersResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAddableUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAddableUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListAddableUsersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAddableUsersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAddableUsersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAddableUsersResponseBody) GetPageResult() *ListAddableUsersResponseBodyPageResult {
	return s.PageResult
}

func (s *ListAddableUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAddableUsersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAddableUsersResponseBody) SetCode(v string) *ListAddableUsersResponseBody {
	s.Code = &v
	return s
}

func (s *ListAddableUsersResponseBody) SetHttpStatusCode(v int32) *ListAddableUsersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAddableUsersResponseBody) SetMessage(v string) *ListAddableUsersResponseBody {
	s.Message = &v
	return s
}

func (s *ListAddableUsersResponseBody) SetPageResult(v *ListAddableUsersResponseBodyPageResult) *ListAddableUsersResponseBody {
	s.PageResult = v
	return s
}

func (s *ListAddableUsersResponseBody) SetRequestId(v string) *ListAddableUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAddableUsersResponseBody) SetSuccess(v bool) *ListAddableUsersResponseBody {
	s.Success = &v
	return s
}

func (s *ListAddableUsersResponseBody) Validate() error {
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAddableUsersResponseBodyPageResult struct {
	// example:
	//
	// 66
	TotalCount *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	UserList   []*ListAddableUsersResponseBodyPageResultUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
}

func (s ListAddableUsersResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListAddableUsersResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListAddableUsersResponseBodyPageResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAddableUsersResponseBodyPageResult) GetUserList() []*ListAddableUsersResponseBodyPageResultUserList {
	return s.UserList
}

func (s *ListAddableUsersResponseBodyPageResult) SetTotalCount(v int32) *ListAddableUsersResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListAddableUsersResponseBodyPageResult) SetUserList(v []*ListAddableUsersResponseBodyPageResultUserList) *ListAddableUsersResponseBodyPageResult {
	s.UserList = v
	return s
}

func (s *ListAddableUsersResponseBodyPageResult) Validate() error {
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

type ListAddableUsersResponseBodyPageResultUserList struct {
	// example:
	//
	// 123@xx.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// 123@dingding
	DingNumber *string `json:"DingNumber,omitempty" xml:"DingNumber,omitempty"`
	// example:
	//
	// xx
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// xx
	DisplayNameWithoutStatus *string `json:"DisplayNameWithoutStatus,omitempty" xml:"DisplayNameWithoutStatus,omitempty"`
	// example:
	//
	// true
	EnableWhiteIp *string `json:"EnableWhiteIp,omitempty" xml:"EnableWhiteIp,omitempty"`
	// example:
	//
	// xx
	FeiShuRobot *string `json:"FeiShuRobot,omitempty" xml:"FeiShuRobot,omitempty"`
	// example:
	//
	// 1717343597000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1717343597000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 123231
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 123@xx.com
	Mail *string `json:"Mail,omitempty" xml:"Mail,omitempty"`
	// example:
	//
	// 13888888888
	MobilePhone *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	// example:
	//
	// xx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// xx
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// example:
	//
	// 231313
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// example:
	//
	// xx
	RealName *string `json:"RealName,omitempty" xml:"RealName,omitempty"`
	// example:
	//
	// 123@xx.com
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// example:
	//
	// aliyun
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// xx
	WeChatRobot *string `json:"WeChatRobot,omitempty" xml:"WeChatRobot,omitempty"`
	// example:
	//
	// *
	WhiteIp *string `json:"WhiteIp,omitempty" xml:"WhiteIp,omitempty"`
}

func (s ListAddableUsersResponseBodyPageResultUserList) String() string {
	return dara.Prettify(s)
}

func (s ListAddableUsersResponseBodyPageResultUserList) GoString() string {
	return s.String()
}

func (s *ListAddableUsersResponseBodyPageResultUserList) GetAccountName() *string {
	return s.AccountName
}

func (s *ListAddableUsersResponseBodyPageResultUserList) GetDingNumber() *string {
	return s.DingNumber
}

func (s *ListAddableUsersResponseBodyPageResultUserList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListAddableUsersResponseBodyPageResultUserList) GetDisplayNameWithoutStatus() *string {
	return s.DisplayNameWithoutStatus
}

func (s *ListAddableUsersResponseBodyPageResultUserList) GetEnableWhiteIp() *string {
	return s.EnableWhiteIp
}

func (s *ListAddableUsersResponseBodyPageResultUserList) GetFeiShuRobot() *string {
	return s.FeiShuRobot
}

func (s *ListAddableUsersResponseBodyPageResultUserList) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListAddableUsersResponseBodyPageResultUserList) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListAddableUsersResponseBodyPageResultUserList) GetId() *string {
	return s.Id
}

func (s *ListAddableUsersResponseBodyPageResultUserList) GetMail() *string {
	return s.Mail
}

func (s *ListAddableUsersResponseBodyPageResultUserList) GetMobilePhone() *string {
	return s.MobilePhone
}

func (s *ListAddableUsersResponseBodyPageResultUserList) GetName() *string {
	return s.Name
}

func (s *ListAddableUsersResponseBodyPageResultUserList) GetNickName() *string {
	return s.NickName
}

func (s *ListAddableUsersResponseBodyPageResultUserList) GetParentId() *string {
	return s.ParentId
}

func (s *ListAddableUsersResponseBodyPageResultUserList) GetRealName() *string {
	return s.RealName
}

func (s *ListAddableUsersResponseBodyPageResultUserList) GetSourceId() *string {
	return s.SourceId
}

func (s *ListAddableUsersResponseBodyPageResultUserList) GetSourceType() *string {
	return s.SourceType
}

func (s *ListAddableUsersResponseBodyPageResultUserList) GetWeChatRobot() *string {
	return s.WeChatRobot
}

func (s *ListAddableUsersResponseBodyPageResultUserList) GetWhiteIp() *string {
	return s.WhiteIp
}

func (s *ListAddableUsersResponseBodyPageResultUserList) SetAccountName(v string) *ListAddableUsersResponseBodyPageResultUserList {
	s.AccountName = &v
	return s
}

func (s *ListAddableUsersResponseBodyPageResultUserList) SetDingNumber(v string) *ListAddableUsersResponseBodyPageResultUserList {
	s.DingNumber = &v
	return s
}

func (s *ListAddableUsersResponseBodyPageResultUserList) SetDisplayName(v string) *ListAddableUsersResponseBodyPageResultUserList {
	s.DisplayName = &v
	return s
}

func (s *ListAddableUsersResponseBodyPageResultUserList) SetDisplayNameWithoutStatus(v string) *ListAddableUsersResponseBodyPageResultUserList {
	s.DisplayNameWithoutStatus = &v
	return s
}

func (s *ListAddableUsersResponseBodyPageResultUserList) SetEnableWhiteIp(v string) *ListAddableUsersResponseBodyPageResultUserList {
	s.EnableWhiteIp = &v
	return s
}

func (s *ListAddableUsersResponseBodyPageResultUserList) SetFeiShuRobot(v string) *ListAddableUsersResponseBodyPageResultUserList {
	s.FeiShuRobot = &v
	return s
}

func (s *ListAddableUsersResponseBodyPageResultUserList) SetGmtCreate(v int64) *ListAddableUsersResponseBodyPageResultUserList {
	s.GmtCreate = &v
	return s
}

func (s *ListAddableUsersResponseBodyPageResultUserList) SetGmtModified(v int64) *ListAddableUsersResponseBodyPageResultUserList {
	s.GmtModified = &v
	return s
}

func (s *ListAddableUsersResponseBodyPageResultUserList) SetId(v string) *ListAddableUsersResponseBodyPageResultUserList {
	s.Id = &v
	return s
}

func (s *ListAddableUsersResponseBodyPageResultUserList) SetMail(v string) *ListAddableUsersResponseBodyPageResultUserList {
	s.Mail = &v
	return s
}

func (s *ListAddableUsersResponseBodyPageResultUserList) SetMobilePhone(v string) *ListAddableUsersResponseBodyPageResultUserList {
	s.MobilePhone = &v
	return s
}

func (s *ListAddableUsersResponseBodyPageResultUserList) SetName(v string) *ListAddableUsersResponseBodyPageResultUserList {
	s.Name = &v
	return s
}

func (s *ListAddableUsersResponseBodyPageResultUserList) SetNickName(v string) *ListAddableUsersResponseBodyPageResultUserList {
	s.NickName = &v
	return s
}

func (s *ListAddableUsersResponseBodyPageResultUserList) SetParentId(v string) *ListAddableUsersResponseBodyPageResultUserList {
	s.ParentId = &v
	return s
}

func (s *ListAddableUsersResponseBodyPageResultUserList) SetRealName(v string) *ListAddableUsersResponseBodyPageResultUserList {
	s.RealName = &v
	return s
}

func (s *ListAddableUsersResponseBodyPageResultUserList) SetSourceId(v string) *ListAddableUsersResponseBodyPageResultUserList {
	s.SourceId = &v
	return s
}

func (s *ListAddableUsersResponseBodyPageResultUserList) SetSourceType(v string) *ListAddableUsersResponseBodyPageResultUserList {
	s.SourceType = &v
	return s
}

func (s *ListAddableUsersResponseBodyPageResultUserList) SetWeChatRobot(v string) *ListAddableUsersResponseBodyPageResultUserList {
	s.WeChatRobot = &v
	return s
}

func (s *ListAddableUsersResponseBodyPageResultUserList) SetWhiteIp(v string) *ListAddableUsersResponseBodyPageResultUserList {
	s.WhiteIp = &v
	return s
}

func (s *ListAddableUsersResponseBodyPageResultUserList) Validate() error {
	return dara.Validate(s)
}
