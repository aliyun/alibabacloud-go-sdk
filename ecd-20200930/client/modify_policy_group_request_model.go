// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPolicyGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdminAccess(v string) *ModifyPolicyGroupRequest
	GetAdminAccess() *string
	SetAppContentProtection(v string) *ModifyPolicyGroupRequest
	GetAppContentProtection() *string
	SetAuthorizeAccessPolicyRule(v []*ModifyPolicyGroupRequestAuthorizeAccessPolicyRule) *ModifyPolicyGroupRequest
	GetAuthorizeAccessPolicyRule() []*ModifyPolicyGroupRequestAuthorizeAccessPolicyRule
	SetAuthorizeSecurityPolicyRule(v []*ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) *ModifyPolicyGroupRequest
	GetAuthorizeSecurityPolicyRule() []*ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule
	SetCameraRedirect(v string) *ModifyPolicyGroupRequest
	GetCameraRedirect() *string
	SetClientType(v []*ModifyPolicyGroupRequestClientType) *ModifyPolicyGroupRequest
	GetClientType() []*ModifyPolicyGroupRequestClientType
	SetClipboard(v string) *ModifyPolicyGroupRequest
	GetClipboard() *string
	SetDeviceRedirects(v []*ModifyPolicyGroupRequestDeviceRedirects) *ModifyPolicyGroupRequest
	GetDeviceRedirects() []*ModifyPolicyGroupRequestDeviceRedirects
	SetDeviceRules(v []*ModifyPolicyGroupRequestDeviceRules) *ModifyPolicyGroupRequest
	GetDeviceRules() []*ModifyPolicyGroupRequestDeviceRules
	SetDomainList(v string) *ModifyPolicyGroupRequest
	GetDomainList() *string
	SetDomainResolveRule(v []*ModifyPolicyGroupRequestDomainResolveRule) *ModifyPolicyGroupRequest
	GetDomainResolveRule() []*ModifyPolicyGroupRequestDomainResolveRule
	SetDomainResolveRuleType(v string) *ModifyPolicyGroupRequest
	GetDomainResolveRuleType() *string
	SetEndUserApplyAdminCoordinate(v string) *ModifyPolicyGroupRequest
	GetEndUserApplyAdminCoordinate() *string
	SetEndUserGroupCoordinate(v string) *ModifyPolicyGroupRequest
	GetEndUserGroupCoordinate() *string
	SetGpuAcceleration(v string) *ModifyPolicyGroupRequest
	GetGpuAcceleration() *string
	SetHtml5Access(v string) *ModifyPolicyGroupRequest
	GetHtml5Access() *string
	SetHtml5FileTransfer(v string) *ModifyPolicyGroupRequest
	GetHtml5FileTransfer() *string
	SetInternetCommunicationProtocol(v string) *ModifyPolicyGroupRequest
	GetInternetCommunicationProtocol() *string
	SetLocalDrive(v string) *ModifyPolicyGroupRequest
	GetLocalDrive() *string
	SetMaxReconnectTime(v int32) *ModifyPolicyGroupRequest
	GetMaxReconnectTime() *int32
	SetName(v string) *ModifyPolicyGroupRequest
	GetName() *string
	SetNetRedirect(v string) *ModifyPolicyGroupRequest
	GetNetRedirect() *string
	SetPolicyGroupId(v string) *ModifyPolicyGroupRequest
	GetPolicyGroupId() *string
	SetPreemptLogin(v string) *ModifyPolicyGroupRequest
	GetPreemptLogin() *string
	SetPreemptLoginUser(v []*string) *ModifyPolicyGroupRequest
	GetPreemptLoginUser() []*string
	SetPrinterRedirection(v string) *ModifyPolicyGroupRequest
	GetPrinterRedirection() *string
	SetRecordContent(v string) *ModifyPolicyGroupRequest
	GetRecordContent() *string
	SetRecordContentExpires(v int64) *ModifyPolicyGroupRequest
	GetRecordContentExpires() *int64
	SetRecording(v string) *ModifyPolicyGroupRequest
	GetRecording() *string
	SetRecordingAudio(v string) *ModifyPolicyGroupRequest
	GetRecordingAudio() *string
	SetRecordingDuration(v int32) *ModifyPolicyGroupRequest
	GetRecordingDuration() *int32
	SetRecordingEndTime(v string) *ModifyPolicyGroupRequest
	GetRecordingEndTime() *string
	SetRecordingExpires(v int64) *ModifyPolicyGroupRequest
	GetRecordingExpires() *int64
	SetRecordingFps(v int64) *ModifyPolicyGroupRequest
	GetRecordingFps() *int64
	SetRecordingStartTime(v string) *ModifyPolicyGroupRequest
	GetRecordingStartTime() *string
	SetRecordingUserNotify(v string) *ModifyPolicyGroupRequest
	GetRecordingUserNotify() *string
	SetRecordingUserNotifyMessage(v string) *ModifyPolicyGroupRequest
	GetRecordingUserNotifyMessage() *string
	SetRegionId(v string) *ModifyPolicyGroupRequest
	GetRegionId() *string
	SetRemoteCoordinate(v string) *ModifyPolicyGroupRequest
	GetRemoteCoordinate() *string
	SetRevokeAccessPolicyRule(v []*ModifyPolicyGroupRequestRevokeAccessPolicyRule) *ModifyPolicyGroupRequest
	GetRevokeAccessPolicyRule() []*ModifyPolicyGroupRequestRevokeAccessPolicyRule
	SetRevokeSecurityPolicyRule(v []*ModifyPolicyGroupRequestRevokeSecurityPolicyRule) *ModifyPolicyGroupRequest
	GetRevokeSecurityPolicyRule() []*ModifyPolicyGroupRequestRevokeSecurityPolicyRule
	SetScope(v string) *ModifyPolicyGroupRequest
	GetScope() *string
	SetScopeValue(v []*string) *ModifyPolicyGroupRequest
	GetScopeValue() []*string
	SetUsbRedirect(v string) *ModifyPolicyGroupRequest
	GetUsbRedirect() *string
	SetUsbSupplyRedirectRule(v []*ModifyPolicyGroupRequestUsbSupplyRedirectRule) *ModifyPolicyGroupRequest
	GetUsbSupplyRedirectRule() []*ModifyPolicyGroupRequestUsbSupplyRedirectRule
	SetVideoRedirect(v string) *ModifyPolicyGroupRequest
	GetVideoRedirect() *string
	SetVisualQuality(v string) *ModifyPolicyGroupRequest
	GetVisualQuality() *string
	SetWatermark(v string) *ModifyPolicyGroupRequest
	GetWatermark() *string
	SetWatermarkAntiCam(v string) *ModifyPolicyGroupRequest
	GetWatermarkAntiCam() *string
	SetWatermarkColor(v int32) *ModifyPolicyGroupRequest
	GetWatermarkColor() *int32
	SetWatermarkDegree(v float64) *ModifyPolicyGroupRequest
	GetWatermarkDegree() *float64
	SetWatermarkFontSize(v int32) *ModifyPolicyGroupRequest
	GetWatermarkFontSize() *int32
	SetWatermarkFontStyle(v string) *ModifyPolicyGroupRequest
	GetWatermarkFontStyle() *string
	SetWatermarkPower(v string) *ModifyPolicyGroupRequest
	GetWatermarkPower() *string
	SetWatermarkRowAmount(v int32) *ModifyPolicyGroupRequest
	GetWatermarkRowAmount() *int32
	SetWatermarkSecurity(v string) *ModifyPolicyGroupRequest
	GetWatermarkSecurity() *string
	SetWatermarkTransparency(v string) *ModifyPolicyGroupRequest
	GetWatermarkTransparency() *string
	SetWatermarkTransparencyValue(v int32) *ModifyPolicyGroupRequest
	GetWatermarkTransparencyValue() *int32
	SetWatermarkType(v string) *ModifyPolicyGroupRequest
	GetWatermarkType() *string
	SetWyAssistant(v string) *ModifyPolicyGroupRequest
	GetWyAssistant() *string
}

type ModifyPolicyGroupRequest struct {
	// Specifies whether the user has administrator permissions after logging on to the cloud desktop.
	//
	// example:
	//
	// deny
	AdminAccess *string `json:"AdminAccess,omitempty" xml:"AdminAccess,omitempty"`
	// Specifies whether to enable the screenshot protection feature.
	//
	// example:
	//
	// on
	AppContentProtection *string `json:"AppContentProtection,omitempty" xml:"AppContentProtection,omitempty"`
	// The list of client IP whitelist entries.
	AuthorizeAccessPolicyRule []*ModifyPolicyGroupRequestAuthorizeAccessPolicyRule `json:"AuthorizeAccessPolicyRule,omitempty" xml:"AuthorizeAccessPolicyRule,omitempty" type:"Repeated"`
	// The list of security group rules.
	AuthorizeSecurityPolicyRule []*ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule `json:"AuthorizeSecurityPolicyRule,omitempty" xml:"AuthorizeSecurityPolicyRule,omitempty" type:"Repeated"`
	// Specifies whether to enable local camera redirection.
	//
	// example:
	//
	// on
	CameraRedirect *string `json:"CameraRedirect,omitempty" xml:"CameraRedirect,omitempty"`
	// The list of logon method control rules.
	ClientType []*ModifyPolicyGroupRequestClientType `json:"ClientType,omitempty" xml:"ClientType,omitempty" type:"Repeated"`
	// The clipboard permission.
	//
	// example:
	//
	// off
	Clipboard *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	// The list of device redirection rules.
	DeviceRedirects []*ModifyPolicyGroupRequestDeviceRedirects `json:"DeviceRedirects,omitempty" xml:"DeviceRedirects,omitempty" type:"Repeated"`
	// The list of custom peripheral rules.
	//
	// if can be null:
	// false
	DeviceRules []*ModifyPolicyGroupRequestDeviceRules `json:"DeviceRules,omitempty" xml:"DeviceRules,omitempty" type:"Repeated"`
	// The domain names for access control. Domain names support the wildcard character (*). Separate multiple domain names with commas (,).
	//
	// example:
	//
	// on
	DomainList *string `json:"DomainList,omitempty" xml:"DomainList,omitempty"`
	// The domain name resolution policy details.
	DomainResolveRule []*ModifyPolicyGroupRequestDomainResolveRule `json:"DomainResolveRule,omitempty" xml:"DomainResolveRule,omitempty" type:"Repeated"`
	// The domain name resolution policy type.
	//
	// example:
	//
	// OFF
	DomainResolveRuleType *string `json:"DomainResolveRuleType,omitempty" xml:"DomainResolveRuleType,omitempty"`
	// Specifies whether to enable the feature that allows users to request administrator assistance.
	//
	// example:
	//
	// on
	EndUserApplyAdminCoordinate *string `json:"EndUserApplyAdminCoordinate,omitempty" xml:"EndUserApplyAdminCoordinate,omitempty"`
	// Specifies whether to enable stream collaboration between users.
	//
	// example:
	//
	// on
	EndUserGroupCoordinate *string `json:"EndUserGroupCoordinate,omitempty" xml:"EndUserGroupCoordinate,omitempty"`
	// Specifies whether to enable the image quality policy for graphics-type cloud desktops. Enable this policy when high performance and user experience are required, such as in professional design scenarios.
	//
	// example:
	//
	// off
	GpuAcceleration *string `json:"GpuAcceleration,omitempty" xml:"GpuAcceleration,omitempty"`
	// The web client access policy.
	//
	// example:
	//
	// off
	Html5Access *string `json:"Html5Access,omitempty" xml:"Html5Access,omitempty"`
	// The web client file transfer policy.
	//
	// example:
	//
	// off
	Html5FileTransfer *string `json:"Html5FileTransfer,omitempty" xml:"Html5FileTransfer,omitempty"`
	// The network communication protocol.
	//
	// example:
	//
	// BOTH
	InternetCommunicationProtocol *string `json:"InternetCommunicationProtocol,omitempty" xml:"InternetCommunicationProtocol,omitempty"`
	// The local disk mapping permission.
	//
	// example:
	//
	// off
	LocalDrive *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	// The maximum reconnection retry time when the cloud desktop is disconnected due to objective reasons. Valid values: 30 to 7200. Unit: seconds.
	//
	// example:
	//
	// 120
	MaxReconnectTime *int32 `json:"MaxReconnectTime,omitempty" xml:"MaxReconnectTime,omitempty"`
	// The Policy Name of the cloud computer policy.
	//
	// example:
	//
	// testPolicyGroupName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Network redirection.
	//
	// example:
	//
	// on
	NetRedirect *string `json:"NetRedirect,omitempty" xml:"NetRedirect,omitempty"`
	// The cloud computer policy ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pg-gx2x1dhsmthe9****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The preemption policy for the cloud computer.
	//
	// > To ensure the user experience and data security of end users who are currently using the cloud computer, preemption between multiple users is not allowed. This configuration is set to `off` by default and cannot be modified.
	//
	// example:
	//
	// off
	PreemptLogin *string `json:"PreemptLogin,omitempty" xml:"PreemptLogin,omitempty"`
	// The usernames of users who can preempt the cloud computer. You can specify up to five usernames.
	PreemptLoginUser []*string `json:"PreemptLoginUser,omitempty" xml:"PreemptLoginUser,omitempty" type:"Repeated"`
	// The printer redirection policy.
	//
	// example:
	//
	// off
	PrinterRedirection *string `json:"PrinterRedirection,omitempty" xml:"PrinterRedirection,omitempty"`
	// Specifies whether to enable custom screen recording.
	//
	// example:
	//
	// OFF
	RecordContent *string `json:"RecordContent,omitempty" xml:"RecordContent,omitempty"`
	// The expiration time of custom screen recording files. Default value: 30 days.
	//
	// example:
	//
	// 30
	RecordContentExpires *int64 `json:"RecordContentExpires,omitempty" xml:"RecordContentExpires,omitempty"`
	// Specifies whether to enable screen recording.
	//
	// example:
	//
	// OFF
	Recording *string `json:"Recording,omitempty" xml:"Recording,omitempty"`
	// The audio recording option for cloud desktops.
	//
	// example:
	//
	// on
	RecordingAudio *string `json:"RecordingAudio,omitempty" xml:"RecordingAudio,omitempty"`
	// The duration of screen recording files, in minutes. Recording files are automatically split and uploaded to the storage space based on the duration you specify. When a file reaches 300 MB, it is rolled over first.
	//
	// example:
	//
	// 10
	RecordingDuration *int32 `json:"RecordingDuration,omitempty" xml:"RecordingDuration,omitempty"`
	// The end time of screen recording. Format: HH:MM:SS. This parameter takes effect only when `Recording` is set to `PERIOD`.
	//
	// example:
	//
	// 08:59:00
	RecordingEndTime *string `json:"RecordingEndTime,omitempty" xml:"RecordingEndTime,omitempty"`
	// The retention period of screen recording files. Valid values: 1 to 180. Unit: days.
	//
	// example:
	//
	// 15
	RecordingExpires *int64 `json:"RecordingExpires,omitempty" xml:"RecordingExpires,omitempty"`
	// The frame rate of screen recording. Unit: FPS (frames per second).
	//
	// example:
	//
	// 5
	RecordingFps *int64 `json:"RecordingFps,omitempty" xml:"RecordingFps,omitempty"`
	// The start time of screen recording. Format: HH:MM:SS. This parameter takes effect only when `Recording` is set to `PERIOD`.
	//
	// example:
	//
	// 08:00:00
	RecordingStartTime *string `json:"RecordingStartTime,omitempty" xml:"RecordingStartTime,omitempty"`
	// The client notification feature for screen recording.
	//
	// example:
	//
	// off
	RecordingUserNotify *string `json:"RecordingUserNotify,omitempty" xml:"RecordingUserNotify,omitempty"`
	// The notification content for the screen recording client. Leave this parameter empty by default.
	//
	// example:
	//
	// Your cloud desktop is being recorded
	RecordingUserNotifyMessage *string `json:"RecordingUserNotifyMessage,omitempty" xml:"RecordingUserNotifyMessage,omitempty"`
	// The region ID. You can call [DescribeRegions](~~DescribeRegions~~) to query the regions supported by WUYING Workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The keyboard and mouse control permissions for remote assistance.
	//
	// example:
	//
	// fullControl
	RemoteCoordinate *string `json:"RemoteCoordinate,omitempty" xml:"RemoteCoordinate,omitempty"`
	// The list of client IP whitelist entries to delete.
	RevokeAccessPolicyRule []*ModifyPolicyGroupRequestRevokeAccessPolicyRule `json:"RevokeAccessPolicyRule,omitempty" xml:"RevokeAccessPolicyRule,omitempty" type:"Repeated"`
	// The list of security group rules to delete.
	RevokeSecurityPolicyRule []*ModifyPolicyGroupRequestRevokeSecurityPolicyRule `json:"RevokeSecurityPolicyRule,omitempty" xml:"RevokeSecurityPolicyRule,omitempty" type:"Repeated"`
	// The effective scope of the policy.
	//
	// example:
	//
	// GLOBAL
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// This parameter is required when `Scope` is set to `IP`. This parameter takes effect only when `Scope` is set to `IP`.
	ScopeValue []*string `json:"ScopeValue,omitempty" xml:"ScopeValue,omitempty" type:"Repeated"`
	// The USB redirection policy.
	//
	// example:
	//
	// off
	UsbRedirect *string `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty"`
	// The USB redirection rules.
	UsbSupplyRedirectRule []*ModifyPolicyGroupRequestUsbSupplyRedirectRule `json:"UsbSupplyRedirectRule,omitempty" xml:"UsbSupplyRedirectRule,omitempty" type:"Repeated"`
	// Multimedia redirection.
	//
	// example:
	//
	// on
	VideoRedirect *string `json:"VideoRedirect,omitempty" xml:"VideoRedirect,omitempty"`
	// The image display quality policy.
	//
	// example:
	//
	// low
	VisualQuality *string `json:"VisualQuality,omitempty" xml:"VisualQuality,omitempty"`
	// The watermark policy.
	//
	// example:
	//
	// off
	Watermark *string `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
	// The anti-photography feature for invisible watermarks.
	//
	// example:
	//
	// off
	WatermarkAntiCam *string `json:"WatermarkAntiCam,omitempty" xml:"WatermarkAntiCam,omitempty"`
	// The watermark font color. Valid values: 0 to 16777215.
	//
	// example:
	//
	// 0
	WatermarkColor *int32 `json:"WatermarkColor,omitempty" xml:"WatermarkColor,omitempty"`
	// The watermark tilt angle. Valid values: -10 to -30.
	//
	// example:
	//
	// -10
	WatermarkDegree *float64 `json:"WatermarkDegree,omitempty" xml:"WatermarkDegree,omitempty"`
	// The watermark font size. Valid values: 10 to 20.
	//
	// example:
	//
	// 10
	WatermarkFontSize *int32 `json:"WatermarkFontSize,omitempty" xml:"WatermarkFontSize,omitempty"`
	// The watermark font style.
	//
	// example:
	//
	// plain
	WatermarkFontStyle *string `json:"WatermarkFontStyle,omitempty" xml:"WatermarkFontStyle,omitempty"`
	// The enhancement feature for invisible watermarks.
	//
	// example:
	//
	// medium
	WatermarkPower *string `json:"WatermarkPower,omitempty" xml:"WatermarkPower,omitempty"`
	// The number of watermark rows.
	//
	// example:
	//
	// 3
	WatermarkRowAmount *int32 `json:"WatermarkRowAmount,omitempty" xml:"WatermarkRowAmount,omitempty"`
	// The security priority rule for invisible watermarks.
	//
	// example:
	//
	// off
	WatermarkSecurity *string `json:"WatermarkSecurity,omitempty" xml:"WatermarkSecurity,omitempty"`
	// The transparency level of the watermark.
	//
	// example:
	//
	// LIGHT
	WatermarkTransparency *string `json:"WatermarkTransparency,omitempty" xml:"WatermarkTransparency,omitempty"`
	// The watermark transparency. A larger value indicates lower transparency. Valid values: 10 to 100.
	//
	// example:
	//
	// 10
	WatermarkTransparencyValue *int32 `json:"WatermarkTransparencyValue,omitempty" xml:"WatermarkTransparencyValue,omitempty"`
	// The watermark type. You can select up to three types, separated by commas (,).
	//
	// example:
	//
	// EndUserId
	WatermarkType *string `json:"WatermarkType,omitempty" xml:"WatermarkType,omitempty"`
	// Specifies whether to provide the WUYING AI Assistant entry in the floating ball when connecting to a cloud computer through a desktop client (including Windows client and macOS client).
	//
	// > Applicable only to desktop clients of V7.7 or later.
	//
	// example:
	//
	// on
	WyAssistant *string `json:"WyAssistant,omitempty" xml:"WyAssistant,omitempty"`
}

func (s ModifyPolicyGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupRequest) GetAdminAccess() *string {
	return s.AdminAccess
}

func (s *ModifyPolicyGroupRequest) GetAppContentProtection() *string {
	return s.AppContentProtection
}

func (s *ModifyPolicyGroupRequest) GetAuthorizeAccessPolicyRule() []*ModifyPolicyGroupRequestAuthorizeAccessPolicyRule {
	return s.AuthorizeAccessPolicyRule
}

func (s *ModifyPolicyGroupRequest) GetAuthorizeSecurityPolicyRule() []*ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule {
	return s.AuthorizeSecurityPolicyRule
}

func (s *ModifyPolicyGroupRequest) GetCameraRedirect() *string {
	return s.CameraRedirect
}

func (s *ModifyPolicyGroupRequest) GetClientType() []*ModifyPolicyGroupRequestClientType {
	return s.ClientType
}

func (s *ModifyPolicyGroupRequest) GetClipboard() *string {
	return s.Clipboard
}

func (s *ModifyPolicyGroupRequest) GetDeviceRedirects() []*ModifyPolicyGroupRequestDeviceRedirects {
	return s.DeviceRedirects
}

func (s *ModifyPolicyGroupRequest) GetDeviceRules() []*ModifyPolicyGroupRequestDeviceRules {
	return s.DeviceRules
}

func (s *ModifyPolicyGroupRequest) GetDomainList() *string {
	return s.DomainList
}

func (s *ModifyPolicyGroupRequest) GetDomainResolveRule() []*ModifyPolicyGroupRequestDomainResolveRule {
	return s.DomainResolveRule
}

func (s *ModifyPolicyGroupRequest) GetDomainResolveRuleType() *string {
	return s.DomainResolveRuleType
}

func (s *ModifyPolicyGroupRequest) GetEndUserApplyAdminCoordinate() *string {
	return s.EndUserApplyAdminCoordinate
}

func (s *ModifyPolicyGroupRequest) GetEndUserGroupCoordinate() *string {
	return s.EndUserGroupCoordinate
}

func (s *ModifyPolicyGroupRequest) GetGpuAcceleration() *string {
	return s.GpuAcceleration
}

func (s *ModifyPolicyGroupRequest) GetHtml5Access() *string {
	return s.Html5Access
}

func (s *ModifyPolicyGroupRequest) GetHtml5FileTransfer() *string {
	return s.Html5FileTransfer
}

func (s *ModifyPolicyGroupRequest) GetInternetCommunicationProtocol() *string {
	return s.InternetCommunicationProtocol
}

func (s *ModifyPolicyGroupRequest) GetLocalDrive() *string {
	return s.LocalDrive
}

func (s *ModifyPolicyGroupRequest) GetMaxReconnectTime() *int32 {
	return s.MaxReconnectTime
}

func (s *ModifyPolicyGroupRequest) GetName() *string {
	return s.Name
}

func (s *ModifyPolicyGroupRequest) GetNetRedirect() *string {
	return s.NetRedirect
}

func (s *ModifyPolicyGroupRequest) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *ModifyPolicyGroupRequest) GetPreemptLogin() *string {
	return s.PreemptLogin
}

func (s *ModifyPolicyGroupRequest) GetPreemptLoginUser() []*string {
	return s.PreemptLoginUser
}

func (s *ModifyPolicyGroupRequest) GetPrinterRedirection() *string {
	return s.PrinterRedirection
}

func (s *ModifyPolicyGroupRequest) GetRecordContent() *string {
	return s.RecordContent
}

func (s *ModifyPolicyGroupRequest) GetRecordContentExpires() *int64 {
	return s.RecordContentExpires
}

func (s *ModifyPolicyGroupRequest) GetRecording() *string {
	return s.Recording
}

func (s *ModifyPolicyGroupRequest) GetRecordingAudio() *string {
	return s.RecordingAudio
}

func (s *ModifyPolicyGroupRequest) GetRecordingDuration() *int32 {
	return s.RecordingDuration
}

func (s *ModifyPolicyGroupRequest) GetRecordingEndTime() *string {
	return s.RecordingEndTime
}

func (s *ModifyPolicyGroupRequest) GetRecordingExpires() *int64 {
	return s.RecordingExpires
}

func (s *ModifyPolicyGroupRequest) GetRecordingFps() *int64 {
	return s.RecordingFps
}

func (s *ModifyPolicyGroupRequest) GetRecordingStartTime() *string {
	return s.RecordingStartTime
}

func (s *ModifyPolicyGroupRequest) GetRecordingUserNotify() *string {
	return s.RecordingUserNotify
}

func (s *ModifyPolicyGroupRequest) GetRecordingUserNotifyMessage() *string {
	return s.RecordingUserNotifyMessage
}

func (s *ModifyPolicyGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyPolicyGroupRequest) GetRemoteCoordinate() *string {
	return s.RemoteCoordinate
}

func (s *ModifyPolicyGroupRequest) GetRevokeAccessPolicyRule() []*ModifyPolicyGroupRequestRevokeAccessPolicyRule {
	return s.RevokeAccessPolicyRule
}

func (s *ModifyPolicyGroupRequest) GetRevokeSecurityPolicyRule() []*ModifyPolicyGroupRequestRevokeSecurityPolicyRule {
	return s.RevokeSecurityPolicyRule
}

func (s *ModifyPolicyGroupRequest) GetScope() *string {
	return s.Scope
}

func (s *ModifyPolicyGroupRequest) GetScopeValue() []*string {
	return s.ScopeValue
}

func (s *ModifyPolicyGroupRequest) GetUsbRedirect() *string {
	return s.UsbRedirect
}

func (s *ModifyPolicyGroupRequest) GetUsbSupplyRedirectRule() []*ModifyPolicyGroupRequestUsbSupplyRedirectRule {
	return s.UsbSupplyRedirectRule
}

func (s *ModifyPolicyGroupRequest) GetVideoRedirect() *string {
	return s.VideoRedirect
}

func (s *ModifyPolicyGroupRequest) GetVisualQuality() *string {
	return s.VisualQuality
}

func (s *ModifyPolicyGroupRequest) GetWatermark() *string {
	return s.Watermark
}

func (s *ModifyPolicyGroupRequest) GetWatermarkAntiCam() *string {
	return s.WatermarkAntiCam
}

func (s *ModifyPolicyGroupRequest) GetWatermarkColor() *int32 {
	return s.WatermarkColor
}

func (s *ModifyPolicyGroupRequest) GetWatermarkDegree() *float64 {
	return s.WatermarkDegree
}

func (s *ModifyPolicyGroupRequest) GetWatermarkFontSize() *int32 {
	return s.WatermarkFontSize
}

func (s *ModifyPolicyGroupRequest) GetWatermarkFontStyle() *string {
	return s.WatermarkFontStyle
}

func (s *ModifyPolicyGroupRequest) GetWatermarkPower() *string {
	return s.WatermarkPower
}

func (s *ModifyPolicyGroupRequest) GetWatermarkRowAmount() *int32 {
	return s.WatermarkRowAmount
}

func (s *ModifyPolicyGroupRequest) GetWatermarkSecurity() *string {
	return s.WatermarkSecurity
}

func (s *ModifyPolicyGroupRequest) GetWatermarkTransparency() *string {
	return s.WatermarkTransparency
}

func (s *ModifyPolicyGroupRequest) GetWatermarkTransparencyValue() *int32 {
	return s.WatermarkTransparencyValue
}

func (s *ModifyPolicyGroupRequest) GetWatermarkType() *string {
	return s.WatermarkType
}

func (s *ModifyPolicyGroupRequest) GetWyAssistant() *string {
	return s.WyAssistant
}

func (s *ModifyPolicyGroupRequest) SetAdminAccess(v string) *ModifyPolicyGroupRequest {
	s.AdminAccess = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetAppContentProtection(v string) *ModifyPolicyGroupRequest {
	s.AppContentProtection = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetAuthorizeAccessPolicyRule(v []*ModifyPolicyGroupRequestAuthorizeAccessPolicyRule) *ModifyPolicyGroupRequest {
	s.AuthorizeAccessPolicyRule = v
	return s
}

func (s *ModifyPolicyGroupRequest) SetAuthorizeSecurityPolicyRule(v []*ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) *ModifyPolicyGroupRequest {
	s.AuthorizeSecurityPolicyRule = v
	return s
}

func (s *ModifyPolicyGroupRequest) SetCameraRedirect(v string) *ModifyPolicyGroupRequest {
	s.CameraRedirect = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetClientType(v []*ModifyPolicyGroupRequestClientType) *ModifyPolicyGroupRequest {
	s.ClientType = v
	return s
}

func (s *ModifyPolicyGroupRequest) SetClipboard(v string) *ModifyPolicyGroupRequest {
	s.Clipboard = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetDeviceRedirects(v []*ModifyPolicyGroupRequestDeviceRedirects) *ModifyPolicyGroupRequest {
	s.DeviceRedirects = v
	return s
}

func (s *ModifyPolicyGroupRequest) SetDeviceRules(v []*ModifyPolicyGroupRequestDeviceRules) *ModifyPolicyGroupRequest {
	s.DeviceRules = v
	return s
}

func (s *ModifyPolicyGroupRequest) SetDomainList(v string) *ModifyPolicyGroupRequest {
	s.DomainList = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetDomainResolveRule(v []*ModifyPolicyGroupRequestDomainResolveRule) *ModifyPolicyGroupRequest {
	s.DomainResolveRule = v
	return s
}

func (s *ModifyPolicyGroupRequest) SetDomainResolveRuleType(v string) *ModifyPolicyGroupRequest {
	s.DomainResolveRuleType = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetEndUserApplyAdminCoordinate(v string) *ModifyPolicyGroupRequest {
	s.EndUserApplyAdminCoordinate = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetEndUserGroupCoordinate(v string) *ModifyPolicyGroupRequest {
	s.EndUserGroupCoordinate = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetGpuAcceleration(v string) *ModifyPolicyGroupRequest {
	s.GpuAcceleration = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetHtml5Access(v string) *ModifyPolicyGroupRequest {
	s.Html5Access = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetHtml5FileTransfer(v string) *ModifyPolicyGroupRequest {
	s.Html5FileTransfer = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetInternetCommunicationProtocol(v string) *ModifyPolicyGroupRequest {
	s.InternetCommunicationProtocol = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetLocalDrive(v string) *ModifyPolicyGroupRequest {
	s.LocalDrive = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetMaxReconnectTime(v int32) *ModifyPolicyGroupRequest {
	s.MaxReconnectTime = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetName(v string) *ModifyPolicyGroupRequest {
	s.Name = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetNetRedirect(v string) *ModifyPolicyGroupRequest {
	s.NetRedirect = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetPolicyGroupId(v string) *ModifyPolicyGroupRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetPreemptLogin(v string) *ModifyPolicyGroupRequest {
	s.PreemptLogin = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetPreemptLoginUser(v []*string) *ModifyPolicyGroupRequest {
	s.PreemptLoginUser = v
	return s
}

func (s *ModifyPolicyGroupRequest) SetPrinterRedirection(v string) *ModifyPolicyGroupRequest {
	s.PrinterRedirection = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetRecordContent(v string) *ModifyPolicyGroupRequest {
	s.RecordContent = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetRecordContentExpires(v int64) *ModifyPolicyGroupRequest {
	s.RecordContentExpires = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetRecording(v string) *ModifyPolicyGroupRequest {
	s.Recording = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetRecordingAudio(v string) *ModifyPolicyGroupRequest {
	s.RecordingAudio = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetRecordingDuration(v int32) *ModifyPolicyGroupRequest {
	s.RecordingDuration = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetRecordingEndTime(v string) *ModifyPolicyGroupRequest {
	s.RecordingEndTime = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetRecordingExpires(v int64) *ModifyPolicyGroupRequest {
	s.RecordingExpires = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetRecordingFps(v int64) *ModifyPolicyGroupRequest {
	s.RecordingFps = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetRecordingStartTime(v string) *ModifyPolicyGroupRequest {
	s.RecordingStartTime = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetRecordingUserNotify(v string) *ModifyPolicyGroupRequest {
	s.RecordingUserNotify = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetRecordingUserNotifyMessage(v string) *ModifyPolicyGroupRequest {
	s.RecordingUserNotifyMessage = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetRegionId(v string) *ModifyPolicyGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetRemoteCoordinate(v string) *ModifyPolicyGroupRequest {
	s.RemoteCoordinate = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetRevokeAccessPolicyRule(v []*ModifyPolicyGroupRequestRevokeAccessPolicyRule) *ModifyPolicyGroupRequest {
	s.RevokeAccessPolicyRule = v
	return s
}

func (s *ModifyPolicyGroupRequest) SetRevokeSecurityPolicyRule(v []*ModifyPolicyGroupRequestRevokeSecurityPolicyRule) *ModifyPolicyGroupRequest {
	s.RevokeSecurityPolicyRule = v
	return s
}

func (s *ModifyPolicyGroupRequest) SetScope(v string) *ModifyPolicyGroupRequest {
	s.Scope = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetScopeValue(v []*string) *ModifyPolicyGroupRequest {
	s.ScopeValue = v
	return s
}

func (s *ModifyPolicyGroupRequest) SetUsbRedirect(v string) *ModifyPolicyGroupRequest {
	s.UsbRedirect = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetUsbSupplyRedirectRule(v []*ModifyPolicyGroupRequestUsbSupplyRedirectRule) *ModifyPolicyGroupRequest {
	s.UsbSupplyRedirectRule = v
	return s
}

func (s *ModifyPolicyGroupRequest) SetVideoRedirect(v string) *ModifyPolicyGroupRequest {
	s.VideoRedirect = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetVisualQuality(v string) *ModifyPolicyGroupRequest {
	s.VisualQuality = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetWatermark(v string) *ModifyPolicyGroupRequest {
	s.Watermark = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetWatermarkAntiCam(v string) *ModifyPolicyGroupRequest {
	s.WatermarkAntiCam = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetWatermarkColor(v int32) *ModifyPolicyGroupRequest {
	s.WatermarkColor = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetWatermarkDegree(v float64) *ModifyPolicyGroupRequest {
	s.WatermarkDegree = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetWatermarkFontSize(v int32) *ModifyPolicyGroupRequest {
	s.WatermarkFontSize = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetWatermarkFontStyle(v string) *ModifyPolicyGroupRequest {
	s.WatermarkFontStyle = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetWatermarkPower(v string) *ModifyPolicyGroupRequest {
	s.WatermarkPower = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetWatermarkRowAmount(v int32) *ModifyPolicyGroupRequest {
	s.WatermarkRowAmount = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetWatermarkSecurity(v string) *ModifyPolicyGroupRequest {
	s.WatermarkSecurity = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetWatermarkTransparency(v string) *ModifyPolicyGroupRequest {
	s.WatermarkTransparency = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetWatermarkTransparencyValue(v int32) *ModifyPolicyGroupRequest {
	s.WatermarkTransparencyValue = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetWatermarkType(v string) *ModifyPolicyGroupRequest {
	s.WatermarkType = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetWyAssistant(v string) *ModifyPolicyGroupRequest {
	s.WyAssistant = &v
	return s
}

func (s *ModifyPolicyGroupRequest) Validate() error {
	if s.AuthorizeAccessPolicyRule != nil {
		for _, item := range s.AuthorizeAccessPolicyRule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AuthorizeSecurityPolicyRule != nil {
		for _, item := range s.AuthorizeSecurityPolicyRule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ClientType != nil {
		for _, item := range s.ClientType {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DeviceRedirects != nil {
		for _, item := range s.DeviceRedirects {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DeviceRules != nil {
		for _, item := range s.DeviceRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DomainResolveRule != nil {
		for _, item := range s.DomainResolveRule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RevokeAccessPolicyRule != nil {
		for _, item := range s.RevokeAccessPolicyRule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RevokeSecurityPolicyRule != nil {
		for _, item := range s.RevokeSecurityPolicyRule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UsbSupplyRedirectRule != nil {
		for _, item := range s.UsbSupplyRedirectRule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyPolicyGroupRequestAuthorizeAccessPolicyRule struct {
	// The client IP address range. The value is an IPv4 CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// 47.100.XX.XX/16
	CidrIp *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	// The description of the client IP whitelist entry.
	//
	// example:
	//
	// Office network segment
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ModifyPolicyGroupRequestAuthorizeAccessPolicyRule) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyGroupRequestAuthorizeAccessPolicyRule) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupRequestAuthorizeAccessPolicyRule) GetCidrIp() *string {
	return s.CidrIp
}

func (s *ModifyPolicyGroupRequestAuthorizeAccessPolicyRule) GetDescription() *string {
	return s.Description
}

func (s *ModifyPolicyGroupRequestAuthorizeAccessPolicyRule) SetCidrIp(v string) *ModifyPolicyGroupRequestAuthorizeAccessPolicyRule {
	s.CidrIp = &v
	return s
}

func (s *ModifyPolicyGroupRequestAuthorizeAccessPolicyRule) SetDescription(v string) *ModifyPolicyGroupRequestAuthorizeAccessPolicyRule {
	s.Description = &v
	return s
}

func (s *ModifyPolicyGroupRequestAuthorizeAccessPolicyRule) Validate() error {
	return dara.Validate(s)
}

type ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule struct {
	// The target of the security group rule. The value is an IPv4 CIDR block.
	//
	// example:
	//
	// 10.0.XX.XX/8
	CidrIp *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	// The description of the security group rule.
	//
	// example:
	//
	// Allow access to the internal R&D environment
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The protocol type of the security group rule.
	//
	// example:
	//
	// TCP
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The authorization policy of the security group rule.
	//
	// example:
	//
	// accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The port range of the security group rule. The valid values of this parameter are determined by the value of IpProtocol:
	//
	// - TCP or UDP: Valid values are 1 to 65535. Separate the start port and end port with a forward slash (/). Example: 1/200.
	//
	// - ICMP: -1/-1.
	//
	// - GRE: -1/-1.
	//
	// - If IpProtocol is set to all: -1/-1.
	//
	// For common ports of typical applications, see [Common ports](https://help.aliyun.com/document_detail/40724.html).
	//
	// example:
	//
	// 22/22
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// The priority of the security group rule. A smaller number indicates a higher priority. Valid values: 1 to 60. Default value: 1.
	//
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The direction of the security group rule.
	//
	// example:
	//
	// inflow
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) GetCidrIp() *string {
	return s.CidrIp
}

func (s *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) GetDescription() *string {
	return s.Description
}

func (s *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) GetPolicy() *string {
	return s.Policy
}

func (s *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) GetPortRange() *string {
	return s.PortRange
}

func (s *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) GetPriority() *string {
	return s.Priority
}

func (s *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) GetType() *string {
	return s.Type
}

func (s *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) SetCidrIp(v string) *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.CidrIp = &v
	return s
}

func (s *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) SetDescription(v string) *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.Description = &v
	return s
}

func (s *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) SetIpProtocol(v string) *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.IpProtocol = &v
	return s
}

func (s *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) SetPolicy(v string) *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.Policy = &v
	return s
}

func (s *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) SetPortRange(v string) *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.PortRange = &v
	return s
}

func (s *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) SetPriority(v string) *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.Priority = &v
	return s
}

func (s *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) SetType(v string) *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.Type = &v
	return s
}

func (s *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) Validate() error {
	return dara.Validate(s)
}

type ModifyPolicyGroupRequestClientType struct {
	// Specifies the client type for logon method control.
	//
	// > If you do not set `ClientType` parameters, all types of clients are allowed to log on to cloud desktops by default.
	//
	// example:
	//
	// windows
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// Specifies whether to allow logon to cloud desktops from a specific client type.
	//
	// example:
	//
	// ON
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyPolicyGroupRequestClientType) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyGroupRequestClientType) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupRequestClientType) GetClientType() *string {
	return s.ClientType
}

func (s *ModifyPolicyGroupRequestClientType) GetStatus() *string {
	return s.Status
}

func (s *ModifyPolicyGroupRequestClientType) SetClientType(v string) *ModifyPolicyGroupRequestClientType {
	s.ClientType = &v
	return s
}

func (s *ModifyPolicyGroupRequestClientType) SetStatus(v string) *ModifyPolicyGroupRequestClientType {
	s.Status = &v
	return s
}

func (s *ModifyPolicyGroupRequestClientType) Validate() error {
	return dara.Validate(s)
}

type ModifyPolicyGroupRequestDeviceRedirects struct {
	// The peripheral device type.
	//
	// example:
	//
	// camera
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// The redirection type.
	//
	// example:
	//
	// usbRedirect
	RedirectType *string `json:"RedirectType,omitempty" xml:"RedirectType,omitempty"`
}

func (s ModifyPolicyGroupRequestDeviceRedirects) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyGroupRequestDeviceRedirects) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupRequestDeviceRedirects) GetDeviceType() *string {
	return s.DeviceType
}

func (s *ModifyPolicyGroupRequestDeviceRedirects) GetRedirectType() *string {
	return s.RedirectType
}

func (s *ModifyPolicyGroupRequestDeviceRedirects) SetDeviceType(v string) *ModifyPolicyGroupRequestDeviceRedirects {
	s.DeviceType = &v
	return s
}

func (s *ModifyPolicyGroupRequestDeviceRedirects) SetRedirectType(v string) *ModifyPolicyGroupRequestDeviceRedirects {
	s.RedirectType = &v
	return s
}

func (s *ModifyPolicyGroupRequestDeviceRedirects) Validate() error {
	return dara.Validate(s)
}

type ModifyPolicyGroupRequestDeviceRules struct {
	// The device name.
	//
	// example:
	//
	// sandisk
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The product ID.
	//
	// example:
	//
	// 0x0781
	DevicePid *string `json:"DevicePid,omitempty" xml:"DevicePid,omitempty"`
	// The peripheral device type.
	//
	// example:
	//
	// storage
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// The vendor ID. See [Valid USB Vendor IDs (VIDs)](https://www.usb.org/sites/default/files/vendor_ids032322.pdf_1.pdf).
	//
	// example:
	//
	// 0x55b1
	DeviceVid *string `json:"DeviceVid,omitempty" xml:"DeviceVid,omitempty"`
	// The link optimization instruction.
	//
	// example:
	//
	// 2:0
	OptCommand *string `json:"OptCommand,omitempty" xml:"OptCommand,omitempty"`
	// The platform types to which the device rule applies.
	//
	// example:
	//
	// Windows
	Platforms *string `json:"Platforms,omitempty" xml:"Platforms,omitempty"`
	// The redirection type.
	//
	// example:
	//
	// usbRedirect
	RedirectType *string `json:"RedirectType,omitempty" xml:"RedirectType,omitempty"`
}

func (s ModifyPolicyGroupRequestDeviceRules) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyGroupRequestDeviceRules) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupRequestDeviceRules) GetDeviceName() *string {
	return s.DeviceName
}

func (s *ModifyPolicyGroupRequestDeviceRules) GetDevicePid() *string {
	return s.DevicePid
}

func (s *ModifyPolicyGroupRequestDeviceRules) GetDeviceType() *string {
	return s.DeviceType
}

func (s *ModifyPolicyGroupRequestDeviceRules) GetDeviceVid() *string {
	return s.DeviceVid
}

func (s *ModifyPolicyGroupRequestDeviceRules) GetOptCommand() *string {
	return s.OptCommand
}

func (s *ModifyPolicyGroupRequestDeviceRules) GetPlatforms() *string {
	return s.Platforms
}

func (s *ModifyPolicyGroupRequestDeviceRules) GetRedirectType() *string {
	return s.RedirectType
}

func (s *ModifyPolicyGroupRequestDeviceRules) SetDeviceName(v string) *ModifyPolicyGroupRequestDeviceRules {
	s.DeviceName = &v
	return s
}

func (s *ModifyPolicyGroupRequestDeviceRules) SetDevicePid(v string) *ModifyPolicyGroupRequestDeviceRules {
	s.DevicePid = &v
	return s
}

func (s *ModifyPolicyGroupRequestDeviceRules) SetDeviceType(v string) *ModifyPolicyGroupRequestDeviceRules {
	s.DeviceType = &v
	return s
}

func (s *ModifyPolicyGroupRequestDeviceRules) SetDeviceVid(v string) *ModifyPolicyGroupRequestDeviceRules {
	s.DeviceVid = &v
	return s
}

func (s *ModifyPolicyGroupRequestDeviceRules) SetOptCommand(v string) *ModifyPolicyGroupRequestDeviceRules {
	s.OptCommand = &v
	return s
}

func (s *ModifyPolicyGroupRequestDeviceRules) SetPlatforms(v string) *ModifyPolicyGroupRequestDeviceRules {
	s.Platforms = &v
	return s
}

func (s *ModifyPolicyGroupRequestDeviceRules) SetRedirectType(v string) *ModifyPolicyGroupRequestDeviceRules {
	s.RedirectType = &v
	return s
}

func (s *ModifyPolicyGroupRequestDeviceRules) Validate() error {
	return dara.Validate(s)
}

type ModifyPolicyGroupRequestDomainResolveRule struct {
	// The policy description.
	//
	// example:
	//
	// Policy description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The domain name.
	//
	// example:
	//
	// *.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The resolution policy.
	//
	// example:
	//
	// allow
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
}

func (s ModifyPolicyGroupRequestDomainResolveRule) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyGroupRequestDomainResolveRule) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupRequestDomainResolveRule) GetDescription() *string {
	return s.Description
}

func (s *ModifyPolicyGroupRequestDomainResolveRule) GetDomain() *string {
	return s.Domain
}

func (s *ModifyPolicyGroupRequestDomainResolveRule) GetPolicy() *string {
	return s.Policy
}

func (s *ModifyPolicyGroupRequestDomainResolveRule) SetDescription(v string) *ModifyPolicyGroupRequestDomainResolveRule {
	s.Description = &v
	return s
}

func (s *ModifyPolicyGroupRequestDomainResolveRule) SetDomain(v string) *ModifyPolicyGroupRequestDomainResolveRule {
	s.Domain = &v
	return s
}

func (s *ModifyPolicyGroupRequestDomainResolveRule) SetPolicy(v string) *ModifyPolicyGroupRequestDomainResolveRule {
	s.Policy = &v
	return s
}

func (s *ModifyPolicyGroupRequestDomainResolveRule) Validate() error {
	return dara.Validate(s)
}

type ModifyPolicyGroupRequestRevokeAccessPolicyRule struct {
	// The client IP address range to delete. The value is an IPv4 CIDR block.
	//
	// example:
	//
	// 47.100.XX.XX/16
	CidrIp *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	// The description of the client IP whitelist entry to delete.
	//
	// example:
	//
	// Office network segment
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ModifyPolicyGroupRequestRevokeAccessPolicyRule) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyGroupRequestRevokeAccessPolicyRule) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupRequestRevokeAccessPolicyRule) GetCidrIp() *string {
	return s.CidrIp
}

func (s *ModifyPolicyGroupRequestRevokeAccessPolicyRule) GetDescription() *string {
	return s.Description
}

func (s *ModifyPolicyGroupRequestRevokeAccessPolicyRule) SetCidrIp(v string) *ModifyPolicyGroupRequestRevokeAccessPolicyRule {
	s.CidrIp = &v
	return s
}

func (s *ModifyPolicyGroupRequestRevokeAccessPolicyRule) SetDescription(v string) *ModifyPolicyGroupRequestRevokeAccessPolicyRule {
	s.Description = &v
	return s
}

func (s *ModifyPolicyGroupRequestRevokeAccessPolicyRule) Validate() error {
	return dara.Validate(s)
}

type ModifyPolicyGroupRequestRevokeSecurityPolicyRule struct {
	// The target of the security group rule to delete. The value is an IPv4 CIDR block.
	//
	// example:
	//
	// 47.100.XX.XX/16
	CidrIp *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	// The description of the security group rule to delete.
	//
	// example:
	//
	// Allow access to the internal R&D environment
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The protocol type of the security group rule to delete.
	//
	// example:
	//
	// tcp
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The authorization policy of the security group rule to delete.
	//
	// example:
	//
	// accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The port range of the security group rule to be deleted. The valid values of this parameter are determined by the value of IpProtocol:
	//
	// - TCP or UDP: Valid values are 1 to 65535. Separate the start port and the end port with a forward slash (/). Example: 1/200.
	//
	// - ICMP: -1/-1.
	//
	// - GRE: -1/-1.
	//
	// - If IpProtocol is set to all: -1/-1.
	//
	// For the common ports of typical applications, see [Common ports](https://help.aliyun.com/document_detail/40724.html).
	//
	// example:
	//
	// 22/22
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// The priority of the security group rule to delete. A smaller value indicates a higher priority. Valid values: 1 to 60. Default value: 1.
	//
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The direction of the security group rule to delete.
	//
	// example:
	//
	// outflow
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyPolicyGroupRequestRevokeSecurityPolicyRule) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyGroupRequestRevokeSecurityPolicyRule) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupRequestRevokeSecurityPolicyRule) GetCidrIp() *string {
	return s.CidrIp
}

func (s *ModifyPolicyGroupRequestRevokeSecurityPolicyRule) GetDescription() *string {
	return s.Description
}

func (s *ModifyPolicyGroupRequestRevokeSecurityPolicyRule) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *ModifyPolicyGroupRequestRevokeSecurityPolicyRule) GetPolicy() *string {
	return s.Policy
}

func (s *ModifyPolicyGroupRequestRevokeSecurityPolicyRule) GetPortRange() *string {
	return s.PortRange
}

func (s *ModifyPolicyGroupRequestRevokeSecurityPolicyRule) GetPriority() *string {
	return s.Priority
}

func (s *ModifyPolicyGroupRequestRevokeSecurityPolicyRule) GetType() *string {
	return s.Type
}

func (s *ModifyPolicyGroupRequestRevokeSecurityPolicyRule) SetCidrIp(v string) *ModifyPolicyGroupRequestRevokeSecurityPolicyRule {
	s.CidrIp = &v
	return s
}

func (s *ModifyPolicyGroupRequestRevokeSecurityPolicyRule) SetDescription(v string) *ModifyPolicyGroupRequestRevokeSecurityPolicyRule {
	s.Description = &v
	return s
}

func (s *ModifyPolicyGroupRequestRevokeSecurityPolicyRule) SetIpProtocol(v string) *ModifyPolicyGroupRequestRevokeSecurityPolicyRule {
	s.IpProtocol = &v
	return s
}

func (s *ModifyPolicyGroupRequestRevokeSecurityPolicyRule) SetPolicy(v string) *ModifyPolicyGroupRequestRevokeSecurityPolicyRule {
	s.Policy = &v
	return s
}

func (s *ModifyPolicyGroupRequestRevokeSecurityPolicyRule) SetPortRange(v string) *ModifyPolicyGroupRequestRevokeSecurityPolicyRule {
	s.PortRange = &v
	return s
}

func (s *ModifyPolicyGroupRequestRevokeSecurityPolicyRule) SetPriority(v string) *ModifyPolicyGroupRequestRevokeSecurityPolicyRule {
	s.Priority = &v
	return s
}

func (s *ModifyPolicyGroupRequestRevokeSecurityPolicyRule) SetType(v string) *ModifyPolicyGroupRequestRevokeSecurityPolicyRule {
	s.Type = &v
	return s
}

func (s *ModifyPolicyGroupRequestRevokeSecurityPolicyRule) Validate() error {
	return dara.Validate(s)
}

type ModifyPolicyGroupRequestUsbSupplyRedirectRule struct {
	// The rule description.
	//
	// example:
	//
	// Test rule
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The device class. This parameter is required when `usbRuleType` is set to 1. See [Defined Class Codes](https://www.usb.org/defined-class-codes).
	//
	// example:
	//
	// 0Eh
	DeviceClass *string `json:"DeviceClass,omitempty" xml:"DeviceClass,omitempty"`
	// The device subclass. This parameter is required when `usbRuleType` is set to 1. See [Defined Class Codes](https://www.usb.org/defined-class-codes).
	//
	// example:
	//
	// xxh
	DeviceSubclass *string `json:"DeviceSubclass,omitempty" xml:"DeviceSubclass,omitempty"`
	// The product ID.
	//
	// example:
	//
	// 08**
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The USB redirection type.
	//
	// example:
	//
	// 1
	UsbRedirectType *int64 `json:"UsbRedirectType,omitempty" xml:"UsbRedirectType,omitempty"`
	// The USB redirection rule type.
	//
	// example:
	//
	// 1
	UsbRuleType *int64 `json:"UsbRuleType,omitempty" xml:"UsbRuleType,omitempty"`
	// The vendor ID. See [Valid USB Vendor IDs (VIDs)](https://www.usb.org/sites/default/files/vendor_ids032322.pdf_1.pdf).
	//
	// example:
	//
	// 04**
	VendorId *string `json:"VendorId,omitempty" xml:"VendorId,omitempty"`
}

func (s ModifyPolicyGroupRequestUsbSupplyRedirectRule) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyGroupRequestUsbSupplyRedirectRule) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupRequestUsbSupplyRedirectRule) GetDescription() *string {
	return s.Description
}

func (s *ModifyPolicyGroupRequestUsbSupplyRedirectRule) GetDeviceClass() *string {
	return s.DeviceClass
}

func (s *ModifyPolicyGroupRequestUsbSupplyRedirectRule) GetDeviceSubclass() *string {
	return s.DeviceSubclass
}

func (s *ModifyPolicyGroupRequestUsbSupplyRedirectRule) GetProductId() *string {
	return s.ProductId
}

func (s *ModifyPolicyGroupRequestUsbSupplyRedirectRule) GetUsbRedirectType() *int64 {
	return s.UsbRedirectType
}

func (s *ModifyPolicyGroupRequestUsbSupplyRedirectRule) GetUsbRuleType() *int64 {
	return s.UsbRuleType
}

func (s *ModifyPolicyGroupRequestUsbSupplyRedirectRule) GetVendorId() *string {
	return s.VendorId
}

func (s *ModifyPolicyGroupRequestUsbSupplyRedirectRule) SetDescription(v string) *ModifyPolicyGroupRequestUsbSupplyRedirectRule {
	s.Description = &v
	return s
}

func (s *ModifyPolicyGroupRequestUsbSupplyRedirectRule) SetDeviceClass(v string) *ModifyPolicyGroupRequestUsbSupplyRedirectRule {
	s.DeviceClass = &v
	return s
}

func (s *ModifyPolicyGroupRequestUsbSupplyRedirectRule) SetDeviceSubclass(v string) *ModifyPolicyGroupRequestUsbSupplyRedirectRule {
	s.DeviceSubclass = &v
	return s
}

func (s *ModifyPolicyGroupRequestUsbSupplyRedirectRule) SetProductId(v string) *ModifyPolicyGroupRequestUsbSupplyRedirectRule {
	s.ProductId = &v
	return s
}

func (s *ModifyPolicyGroupRequestUsbSupplyRedirectRule) SetUsbRedirectType(v int64) *ModifyPolicyGroupRequestUsbSupplyRedirectRule {
	s.UsbRedirectType = &v
	return s
}

func (s *ModifyPolicyGroupRequestUsbSupplyRedirectRule) SetUsbRuleType(v int64) *ModifyPolicyGroupRequestUsbSupplyRedirectRule {
	s.UsbRuleType = &v
	return s
}

func (s *ModifyPolicyGroupRequestUsbSupplyRedirectRule) SetVendorId(v string) *ModifyPolicyGroupRequestUsbSupplyRedirectRule {
	s.VendorId = &v
	return s
}

func (s *ModifyPolicyGroupRequestUsbSupplyRedirectRule) Validate() error {
	return dara.Validate(s)
}
