// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTerminalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTerminalResponseBody
	GetCode() *string
	SetData(v []*ListTerminalResponseBodyData) *ListTerminalResponseBody
	GetData() []*ListTerminalResponseBodyData
	SetHttpStatusCode(v int32) *ListTerminalResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListTerminalResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListTerminalResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTerminalResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTerminalResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListTerminalResponseBody
	GetTotalCount() *int32
}

type ListTerminalResponseBody struct {
	Code           *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*ListTerminalResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                          `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken      *string                         `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId      *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int32                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTerminalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTerminalResponseBody) GoString() string {
	return s.String()
}

func (s *ListTerminalResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTerminalResponseBody) GetData() []*ListTerminalResponseBodyData {
	return s.Data
}

func (s *ListTerminalResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListTerminalResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTerminalResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTerminalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTerminalResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTerminalResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTerminalResponseBody) SetCode(v string) *ListTerminalResponseBody {
	s.Code = &v
	return s
}

func (s *ListTerminalResponseBody) SetData(v []*ListTerminalResponseBodyData) *ListTerminalResponseBody {
	s.Data = v
	return s
}

func (s *ListTerminalResponseBody) SetHttpStatusCode(v int32) *ListTerminalResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTerminalResponseBody) SetMessage(v string) *ListTerminalResponseBody {
	s.Message = &v
	return s
}

func (s *ListTerminalResponseBody) SetNextToken(v string) *ListTerminalResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTerminalResponseBody) SetRequestId(v string) *ListTerminalResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTerminalResponseBody) SetSuccess(v bool) *ListTerminalResponseBody {
	s.Success = &v
	return s
}

func (s *ListTerminalResponseBody) SetTotalCount(v int32) *ListTerminalResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTerminalResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTerminalResponseBodyData struct {
	Alias         *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	BindUserCount *int32  `json:"BindUserCount,omitempty" xml:"BindUserCount,omitempty"`
	BindUserId    *string `json:"BindUserId,omitempty" xml:"BindUserId,omitempty"`
	BuildId       *string `json:"BuildId,omitempty" xml:"BuildId,omitempty"`
	ClientType    *int32  `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	DesktopId     *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	InManage      *bool   `json:"InManage,omitempty" xml:"InManage,omitempty"`
	Ipv4          *string `json:"Ipv4,omitempty" xml:"Ipv4,omitempty"`
	// Deprecated
	LastLoginUser   *string `json:"LastLoginUser,omitempty" xml:"LastLoginUser,omitempty"`
	LocationInfo    *string `json:"LocationInfo,omitempty" xml:"LocationInfo,omitempty"`
	LockSettings    *bool   `json:"LockSettings,omitempty" xml:"LockSettings,omitempty"`
	LoginUser       *string `json:"LoginUser,omitempty" xml:"LoginUser,omitempty"`
	Model           *string `json:"Model,omitempty" xml:"Model,omitempty"`
	OnlineStatus    *bool   `json:"OnlineStatus,omitempty" xml:"OnlineStatus,omitempty"`
	SerialNumber    *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	TerminalGroupId *string `json:"TerminalGroupId,omitempty" xml:"TerminalGroupId,omitempty"`
	Uuid            *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListTerminalResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTerminalResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTerminalResponseBodyData) GetAlias() *string {
	return s.Alias
}

func (s *ListTerminalResponseBodyData) GetBindUserCount() *int32 {
	return s.BindUserCount
}

func (s *ListTerminalResponseBodyData) GetBindUserId() *string {
	return s.BindUserId
}

func (s *ListTerminalResponseBodyData) GetBuildId() *string {
	return s.BuildId
}

func (s *ListTerminalResponseBodyData) GetClientType() *int32 {
	return s.ClientType
}

func (s *ListTerminalResponseBodyData) GetDesktopId() *string {
	return s.DesktopId
}

func (s *ListTerminalResponseBodyData) GetInManage() *bool {
	return s.InManage
}

func (s *ListTerminalResponseBodyData) GetIpv4() *string {
	return s.Ipv4
}

func (s *ListTerminalResponseBodyData) GetLastLoginUser() *string {
	return s.LastLoginUser
}

func (s *ListTerminalResponseBodyData) GetLocationInfo() *string {
	return s.LocationInfo
}

func (s *ListTerminalResponseBodyData) GetLockSettings() *bool {
	return s.LockSettings
}

func (s *ListTerminalResponseBodyData) GetLoginUser() *string {
	return s.LoginUser
}

func (s *ListTerminalResponseBodyData) GetModel() *string {
	return s.Model
}

func (s *ListTerminalResponseBodyData) GetOnlineStatus() *bool {
	return s.OnlineStatus
}

func (s *ListTerminalResponseBodyData) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *ListTerminalResponseBodyData) GetTerminalGroupId() *string {
	return s.TerminalGroupId
}

func (s *ListTerminalResponseBodyData) GetUuid() *string {
	return s.Uuid
}

func (s *ListTerminalResponseBodyData) SetAlias(v string) *ListTerminalResponseBodyData {
	s.Alias = &v
	return s
}

func (s *ListTerminalResponseBodyData) SetBindUserCount(v int32) *ListTerminalResponseBodyData {
	s.BindUserCount = &v
	return s
}

func (s *ListTerminalResponseBodyData) SetBindUserId(v string) *ListTerminalResponseBodyData {
	s.BindUserId = &v
	return s
}

func (s *ListTerminalResponseBodyData) SetBuildId(v string) *ListTerminalResponseBodyData {
	s.BuildId = &v
	return s
}

func (s *ListTerminalResponseBodyData) SetClientType(v int32) *ListTerminalResponseBodyData {
	s.ClientType = &v
	return s
}

func (s *ListTerminalResponseBodyData) SetDesktopId(v string) *ListTerminalResponseBodyData {
	s.DesktopId = &v
	return s
}

func (s *ListTerminalResponseBodyData) SetInManage(v bool) *ListTerminalResponseBodyData {
	s.InManage = &v
	return s
}

func (s *ListTerminalResponseBodyData) SetIpv4(v string) *ListTerminalResponseBodyData {
	s.Ipv4 = &v
	return s
}

func (s *ListTerminalResponseBodyData) SetLastLoginUser(v string) *ListTerminalResponseBodyData {
	s.LastLoginUser = &v
	return s
}

func (s *ListTerminalResponseBodyData) SetLocationInfo(v string) *ListTerminalResponseBodyData {
	s.LocationInfo = &v
	return s
}

func (s *ListTerminalResponseBodyData) SetLockSettings(v bool) *ListTerminalResponseBodyData {
	s.LockSettings = &v
	return s
}

func (s *ListTerminalResponseBodyData) SetLoginUser(v string) *ListTerminalResponseBodyData {
	s.LoginUser = &v
	return s
}

func (s *ListTerminalResponseBodyData) SetModel(v string) *ListTerminalResponseBodyData {
	s.Model = &v
	return s
}

func (s *ListTerminalResponseBodyData) SetOnlineStatus(v bool) *ListTerminalResponseBodyData {
	s.OnlineStatus = &v
	return s
}

func (s *ListTerminalResponseBodyData) SetSerialNumber(v string) *ListTerminalResponseBodyData {
	s.SerialNumber = &v
	return s
}

func (s *ListTerminalResponseBodyData) SetTerminalGroupId(v string) *ListTerminalResponseBodyData {
	s.TerminalGroupId = &v
	return s
}

func (s *ListTerminalResponseBodyData) SetUuid(v string) *ListTerminalResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *ListTerminalResponseBodyData) Validate() error {
	return dara.Validate(s)
}
