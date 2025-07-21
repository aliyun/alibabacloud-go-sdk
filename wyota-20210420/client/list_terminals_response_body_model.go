// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTerminalsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTerminalsResponseBody
	GetCode() *string
	SetData(v []*ListTerminalsResponseBodyData) *ListTerminalsResponseBody
	GetData() []*ListTerminalsResponseBodyData
	SetHttpStatusCode(v int32) *ListTerminalsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListTerminalsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListTerminalsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTerminalsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTerminalsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListTerminalsResponseBody
	GetTotalCount() *int32
}

type ListTerminalsResponseBody struct {
	// example:
	//
	// TERMINAL_NOT_FOUND
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListTerminalsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// terminal not found
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// AAAAAdEdsXbwG2ZlbWCzN4wTTg6wQvfp7u1BJl4bxCAby41POSaYAlCvfULQpkAnb0ff****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// C5DCE54A-B266-522E-A6ED-468AF45F5AAA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTerminalsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTerminalsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTerminalsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTerminalsResponseBody) GetData() []*ListTerminalsResponseBodyData {
	return s.Data
}

func (s *ListTerminalsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListTerminalsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTerminalsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTerminalsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTerminalsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTerminalsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTerminalsResponseBody) SetCode(v string) *ListTerminalsResponseBody {
	s.Code = &v
	return s
}

func (s *ListTerminalsResponseBody) SetData(v []*ListTerminalsResponseBodyData) *ListTerminalsResponseBody {
	s.Data = v
	return s
}

func (s *ListTerminalsResponseBody) SetHttpStatusCode(v int32) *ListTerminalsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTerminalsResponseBody) SetMessage(v string) *ListTerminalsResponseBody {
	s.Message = &v
	return s
}

func (s *ListTerminalsResponseBody) SetNextToken(v string) *ListTerminalsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTerminalsResponseBody) SetRequestId(v string) *ListTerminalsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTerminalsResponseBody) SetSuccess(v bool) *ListTerminalsResponseBody {
	s.Success = &v
	return s
}

func (s *ListTerminalsResponseBody) SetTotalCount(v int32) *ListTerminalsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTerminalsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTerminalsResponseBodyData struct {
	// example:
	//
	// DemoDevice
	Alias         *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	BindUserCount *int32  `json:"BindUserCount,omitempty" xml:"BindUserCount,omitempty"`
	// example:
	//
	// 7.0.2-RS-20240805.044924
	BuildId *string `json:"BuildId,omitempty" xml:"BuildId,omitempty"`
	// example:
	//
	// 1
	ClientType *int32 `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// example:
	//
	// ecd-drqmaogzbmbdf****
	CurrentConnectDesktop *string `json:"CurrentConnectDesktop,omitempty" xml:"CurrentConnectDesktop,omitempty"`
	// example:
	//
	// alice
	CurrentLoginUser *string `json:"CurrentLoginUser,omitempty" xml:"CurrentLoginUser,omitempty"`
	// example:
	//
	// 192.168.XX.XX
	Ipv4          *string `json:"Ipv4,omitempty" xml:"Ipv4,omitempty"`
	LastLoginUser *string `json:"LastLoginUser,omitempty" xml:"LastLoginUser,omitempty"`
	LocationInfo  *string `json:"LocationInfo,omitempty" xml:"LocationInfo,omitempty"`
	ManageTime    *string `json:"ManageTime,omitempty" xml:"ManageTime,omitempty"`
	// example:
	//
	// US01
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// true
	Online *bool `json:"Online,omitempty" xml:"Online,omitempty"`
	// example:
	//
	// alice
	PasswordFreeLoginUser *string `json:"PasswordFreeLoginUser,omitempty" xml:"PasswordFreeLoginUser,omitempty"`
	PublicIpv4            *string `json:"PublicIpv4,omitempty" xml:"PublicIpv4,omitempty"`
	// example:
	//
	// ODN49YQCPQYC****
	SerialNumber                 *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	SetPasswordFreeLoginUserTime *string `json:"SetPasswordFreeLoginUserTime,omitempty" xml:"SetPasswordFreeLoginUserTime,omitempty"`
	// example:
	//
	// tg-default
	TerminalGroupId *string `json:"TerminalGroupId,omitempty" xml:"TerminalGroupId,omitempty"`
	// example:
	//
	// 04873D3898B51A7DF2455C1E1DC9****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListTerminalsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTerminalsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTerminalsResponseBodyData) GetAlias() *string {
	return s.Alias
}

func (s *ListTerminalsResponseBodyData) GetBindUserCount() *int32 {
	return s.BindUserCount
}

func (s *ListTerminalsResponseBodyData) GetBuildId() *string {
	return s.BuildId
}

func (s *ListTerminalsResponseBodyData) GetClientType() *int32 {
	return s.ClientType
}

func (s *ListTerminalsResponseBodyData) GetCurrentConnectDesktop() *string {
	return s.CurrentConnectDesktop
}

func (s *ListTerminalsResponseBodyData) GetCurrentLoginUser() *string {
	return s.CurrentLoginUser
}

func (s *ListTerminalsResponseBodyData) GetIpv4() *string {
	return s.Ipv4
}

func (s *ListTerminalsResponseBodyData) GetLastLoginUser() *string {
	return s.LastLoginUser
}

func (s *ListTerminalsResponseBodyData) GetLocationInfo() *string {
	return s.LocationInfo
}

func (s *ListTerminalsResponseBodyData) GetManageTime() *string {
	return s.ManageTime
}

func (s *ListTerminalsResponseBodyData) GetModel() *string {
	return s.Model
}

func (s *ListTerminalsResponseBodyData) GetOnline() *bool {
	return s.Online
}

func (s *ListTerminalsResponseBodyData) GetPasswordFreeLoginUser() *string {
	return s.PasswordFreeLoginUser
}

func (s *ListTerminalsResponseBodyData) GetPublicIpv4() *string {
	return s.PublicIpv4
}

func (s *ListTerminalsResponseBodyData) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *ListTerminalsResponseBodyData) GetSetPasswordFreeLoginUserTime() *string {
	return s.SetPasswordFreeLoginUserTime
}

func (s *ListTerminalsResponseBodyData) GetTerminalGroupId() *string {
	return s.TerminalGroupId
}

func (s *ListTerminalsResponseBodyData) GetUuid() *string {
	return s.Uuid
}

func (s *ListTerminalsResponseBodyData) SetAlias(v string) *ListTerminalsResponseBodyData {
	s.Alias = &v
	return s
}

func (s *ListTerminalsResponseBodyData) SetBindUserCount(v int32) *ListTerminalsResponseBodyData {
	s.BindUserCount = &v
	return s
}

func (s *ListTerminalsResponseBodyData) SetBuildId(v string) *ListTerminalsResponseBodyData {
	s.BuildId = &v
	return s
}

func (s *ListTerminalsResponseBodyData) SetClientType(v int32) *ListTerminalsResponseBodyData {
	s.ClientType = &v
	return s
}

func (s *ListTerminalsResponseBodyData) SetCurrentConnectDesktop(v string) *ListTerminalsResponseBodyData {
	s.CurrentConnectDesktop = &v
	return s
}

func (s *ListTerminalsResponseBodyData) SetCurrentLoginUser(v string) *ListTerminalsResponseBodyData {
	s.CurrentLoginUser = &v
	return s
}

func (s *ListTerminalsResponseBodyData) SetIpv4(v string) *ListTerminalsResponseBodyData {
	s.Ipv4 = &v
	return s
}

func (s *ListTerminalsResponseBodyData) SetLastLoginUser(v string) *ListTerminalsResponseBodyData {
	s.LastLoginUser = &v
	return s
}

func (s *ListTerminalsResponseBodyData) SetLocationInfo(v string) *ListTerminalsResponseBodyData {
	s.LocationInfo = &v
	return s
}

func (s *ListTerminalsResponseBodyData) SetManageTime(v string) *ListTerminalsResponseBodyData {
	s.ManageTime = &v
	return s
}

func (s *ListTerminalsResponseBodyData) SetModel(v string) *ListTerminalsResponseBodyData {
	s.Model = &v
	return s
}

func (s *ListTerminalsResponseBodyData) SetOnline(v bool) *ListTerminalsResponseBodyData {
	s.Online = &v
	return s
}

func (s *ListTerminalsResponseBodyData) SetPasswordFreeLoginUser(v string) *ListTerminalsResponseBodyData {
	s.PasswordFreeLoginUser = &v
	return s
}

func (s *ListTerminalsResponseBodyData) SetPublicIpv4(v string) *ListTerminalsResponseBodyData {
	s.PublicIpv4 = &v
	return s
}

func (s *ListTerminalsResponseBodyData) SetSerialNumber(v string) *ListTerminalsResponseBodyData {
	s.SerialNumber = &v
	return s
}

func (s *ListTerminalsResponseBodyData) SetSetPasswordFreeLoginUserTime(v string) *ListTerminalsResponseBodyData {
	s.SetPasswordFreeLoginUserTime = &v
	return s
}

func (s *ListTerminalsResponseBodyData) SetTerminalGroupId(v string) *ListTerminalsResponseBodyData {
	s.TerminalGroupId = &v
	return s
}

func (s *ListTerminalsResponseBodyData) SetUuid(v string) *ListTerminalsResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *ListTerminalsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
