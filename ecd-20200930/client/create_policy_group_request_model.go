// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicyGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdminAccess(v string) *CreatePolicyGroupRequest
	GetAdminAccess() *string
	SetAppContentProtection(v string) *CreatePolicyGroupRequest
	GetAppContentProtection() *string
	SetAuthorizeAccessPolicyRule(v []*CreatePolicyGroupRequestAuthorizeAccessPolicyRule) *CreatePolicyGroupRequest
	GetAuthorizeAccessPolicyRule() []*CreatePolicyGroupRequestAuthorizeAccessPolicyRule
	SetAuthorizeSecurityPolicyRule(v []*CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) *CreatePolicyGroupRequest
	GetAuthorizeSecurityPolicyRule() []*CreatePolicyGroupRequestAuthorizeSecurityPolicyRule
	SetCameraRedirect(v string) *CreatePolicyGroupRequest
	GetCameraRedirect() *string
	SetClientType(v []*CreatePolicyGroupRequestClientType) *CreatePolicyGroupRequest
	GetClientType() []*CreatePolicyGroupRequestClientType
	SetClipboard(v string) *CreatePolicyGroupRequest
	GetClipboard() *string
	SetDeviceRedirects(v []*CreatePolicyGroupRequestDeviceRedirects) *CreatePolicyGroupRequest
	GetDeviceRedirects() []*CreatePolicyGroupRequestDeviceRedirects
	SetDeviceRules(v []*CreatePolicyGroupRequestDeviceRules) *CreatePolicyGroupRequest
	GetDeviceRules() []*CreatePolicyGroupRequestDeviceRules
	SetDomainList(v string) *CreatePolicyGroupRequest
	GetDomainList() *string
	SetDomainResolveRule(v []*CreatePolicyGroupRequestDomainResolveRule) *CreatePolicyGroupRequest
	GetDomainResolveRule() []*CreatePolicyGroupRequestDomainResolveRule
	SetDomainResolveRuleType(v string) *CreatePolicyGroupRequest
	GetDomainResolveRuleType() *string
	SetEndUserApplyAdminCoordinate(v string) *CreatePolicyGroupRequest
	GetEndUserApplyAdminCoordinate() *string
	SetEndUserGroupCoordinate(v string) *CreatePolicyGroupRequest
	GetEndUserGroupCoordinate() *string
	SetGpuAcceleration(v string) *CreatePolicyGroupRequest
	GetGpuAcceleration() *string
	SetHtml5Access(v string) *CreatePolicyGroupRequest
	GetHtml5Access() *string
	SetHtml5FileTransfer(v string) *CreatePolicyGroupRequest
	GetHtml5FileTransfer() *string
	SetInternetCommunicationProtocol(v string) *CreatePolicyGroupRequest
	GetInternetCommunicationProtocol() *string
	SetLocalDrive(v string) *CreatePolicyGroupRequest
	GetLocalDrive() *string
	SetMaxReconnectTime(v int32) *CreatePolicyGroupRequest
	GetMaxReconnectTime() *int32
	SetName(v string) *CreatePolicyGroupRequest
	GetName() *string
	SetNetRedirect(v string) *CreatePolicyGroupRequest
	GetNetRedirect() *string
	SetPreemptLogin(v string) *CreatePolicyGroupRequest
	GetPreemptLogin() *string
	SetPreemptLoginUser(v []*string) *CreatePolicyGroupRequest
	GetPreemptLoginUser() []*string
	SetPrinterRedirection(v string) *CreatePolicyGroupRequest
	GetPrinterRedirection() *string
	SetRecordContent(v string) *CreatePolicyGroupRequest
	GetRecordContent() *string
	SetRecordContentExpires(v int64) *CreatePolicyGroupRequest
	GetRecordContentExpires() *int64
	SetRecording(v string) *CreatePolicyGroupRequest
	GetRecording() *string
	SetRecordingAudio(v string) *CreatePolicyGroupRequest
	GetRecordingAudio() *string
	SetRecordingDuration(v int32) *CreatePolicyGroupRequest
	GetRecordingDuration() *int32
	SetRecordingEndTime(v string) *CreatePolicyGroupRequest
	GetRecordingEndTime() *string
	SetRecordingExpires(v int64) *CreatePolicyGroupRequest
	GetRecordingExpires() *int64
	SetRecordingFps(v int64) *CreatePolicyGroupRequest
	GetRecordingFps() *int64
	SetRecordingStartTime(v string) *CreatePolicyGroupRequest
	GetRecordingStartTime() *string
	SetRecordingUserNotify(v string) *CreatePolicyGroupRequest
	GetRecordingUserNotify() *string
	SetRecordingUserNotifyMessage(v string) *CreatePolicyGroupRequest
	GetRecordingUserNotifyMessage() *string
	SetRegionId(v string) *CreatePolicyGroupRequest
	GetRegionId() *string
	SetRemoteCoordinate(v string) *CreatePolicyGroupRequest
	GetRemoteCoordinate() *string
	SetScope(v string) *CreatePolicyGroupRequest
	GetScope() *string
	SetScopeValue(v []*string) *CreatePolicyGroupRequest
	GetScopeValue() []*string
	SetUsbRedirect(v string) *CreatePolicyGroupRequest
	GetUsbRedirect() *string
	SetUsbSupplyRedirectRule(v []*CreatePolicyGroupRequestUsbSupplyRedirectRule) *CreatePolicyGroupRequest
	GetUsbSupplyRedirectRule() []*CreatePolicyGroupRequestUsbSupplyRedirectRule
	SetVideoRedirect(v string) *CreatePolicyGroupRequest
	GetVideoRedirect() *string
	SetVisualQuality(v string) *CreatePolicyGroupRequest
	GetVisualQuality() *string
	SetWatermark(v string) *CreatePolicyGroupRequest
	GetWatermark() *string
	SetWatermarkAntiCam(v string) *CreatePolicyGroupRequest
	GetWatermarkAntiCam() *string
	SetWatermarkColor(v int32) *CreatePolicyGroupRequest
	GetWatermarkColor() *int32
	SetWatermarkDegree(v float64) *CreatePolicyGroupRequest
	GetWatermarkDegree() *float64
	SetWatermarkFontSize(v int32) *CreatePolicyGroupRequest
	GetWatermarkFontSize() *int32
	SetWatermarkFontStyle(v string) *CreatePolicyGroupRequest
	GetWatermarkFontStyle() *string
	SetWatermarkPower(v string) *CreatePolicyGroupRequest
	GetWatermarkPower() *string
	SetWatermarkRowAmount(v int32) *CreatePolicyGroupRequest
	GetWatermarkRowAmount() *int32
	SetWatermarkSecurity(v string) *CreatePolicyGroupRequest
	GetWatermarkSecurity() *string
	SetWatermarkTransparency(v string) *CreatePolicyGroupRequest
	GetWatermarkTransparency() *string
	SetWatermarkTransparencyValue(v int32) *CreatePolicyGroupRequest
	GetWatermarkTransparencyValue() *int32
	SetWatermarkType(v string) *CreatePolicyGroupRequest
	GetWatermarkType() *string
	SetWyAssistant(v string) *CreatePolicyGroupRequest
	GetWyAssistant() *string
}

type CreatePolicyGroupRequest struct {
	// Specifies whether a user has administrative permissions after logging on to the cloud computer.
	//
	// > This feature is in invitational preview and is not available to the public.
	//
	// example:
	//
	// deny
	AdminAccess *string `json:"AdminAccess,omitempty" xml:"AdminAccess,omitempty"`
	// Specifies whether to enable the anti-screenshot feature.
	//
	// example:
	//
	// off
	AppContentProtection *string `json:"AppContentProtection,omitempty" xml:"AppContentProtection,omitempty"`
	// The client IP address whitelist. After you configure this parameter, only IP addresses in the whitelist can access the cloud computer.
	AuthorizeAccessPolicyRule []*CreatePolicyGroupRequestAuthorizeAccessPolicyRule `json:"AuthorizeAccessPolicyRule,omitempty" xml:"AuthorizeAccessPolicyRule,omitempty" type:"Repeated"`
	// The list of security group rules.
	AuthorizeSecurityPolicyRule []*CreatePolicyGroupRequestAuthorizeSecurityPolicyRule `json:"AuthorizeSecurityPolicyRule,omitempty" xml:"AuthorizeSecurityPolicyRule,omitempty" type:"Repeated"`
	// Specifies whether to enable local camera redirection.
	//
	// example:
	//
	// on
	CameraRedirect *string `json:"CameraRedirect,omitempty" xml:"CameraRedirect,omitempty"`
	// The list of logon method control rules. These rules control which clients can be used to access the cloud computer.
	ClientType []*CreatePolicyGroupRequestClientType `json:"ClientType,omitempty" xml:"ClientType,omitempty" type:"Repeated"`
	// The clipboard permission.
	//
	// example:
	//
	// off
	Clipboard *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	// The list of device redirection rules.
	DeviceRedirects []*CreatePolicyGroupRequestDeviceRedirects `json:"DeviceRedirects,omitempty" xml:"DeviceRedirects,omitempty" type:"Repeated"`
	// The list of custom peripheral rules.
	DeviceRules []*CreatePolicyGroupRequestDeviceRules `json:"DeviceRules,omitempty" xml:"DeviceRules,omitempty" type:"Repeated"`
	// The policy for controlling access to domain names. You can use a wildcard character (\\*). Separate multiple domain names with commas (,).
	//
	// example:
	//
	// off
	DomainList *string `json:"DomainList,omitempty" xml:"DomainList,omitempty"`
	// The details of the domain name resolution policy.
	DomainResolveRule []*CreatePolicyGroupRequestDomainResolveRule `json:"DomainResolveRule,omitempty" xml:"DomainResolveRule,omitempty" type:"Repeated"`
	// The type of the domain name resolution policy.
	//
	// example:
	//
	// OFF
	DomainResolveRuleType *string `json:"DomainResolveRuleType,omitempty" xml:"DomainResolveRuleType,omitempty"`
	// Specifies whether to allow end users to request assistance from administrators.
	//
	// example:
	//
	// ON
	EndUserApplyAdminCoordinate *string `json:"EndUserApplyAdminCoordinate,omitempty" xml:"EndUserApplyAdminCoordinate,omitempty"`
	// Specifies whether to enable stream collaboration between users.
	//
	// example:
	//
	// ON
	EndUserGroupCoordinate *string `json:"EndUserGroupCoordinate,omitempty" xml:"EndUserGroupCoordinate,omitempty"`
	// Specifies whether to enable the image quality policy for graphics cloud computers. Enable this policy for scenarios that require high performance and user experience, such as professional design.
	//
	// example:
	//
	// off
	GpuAcceleration *string `json:"GpuAcceleration,omitempty" xml:"GpuAcceleration,omitempty"`
	// The policy for access from web clients.
	//
	// > Use the `ClientType` parameters to manage logon methods.
	//
	// example:
	//
	// off
	Html5Access *string `json:"Html5Access,omitempty" xml:"Html5Access,omitempty"`
	// The file transfer policy for web clients.
	//
	// example:
	//
	// off
	Html5FileTransfer *string `json:"Html5FileTransfer,omitempty" xml:"Html5FileTransfer,omitempty"`
	// The network communication protocol.
	//
	// example:
	//
	// both
	InternetCommunicationProtocol *string `json:"InternetCommunicationProtocol,omitempty" xml:"InternetCommunicationProtocol,omitempty"`
	// The local disk mapping permission.
	//
	// example:
	//
	// off
	LocalDrive *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	// The maximum amount of time to retry the connection if the cloud computer is disconnected due to an unexpected event. Valid values: 30 to 7200. Unit: seconds.
	//
	// example:
	//
	// 120
	MaxReconnectTime *int32 `json:"MaxReconnectTime,omitempty" xml:"MaxReconnectTime,omitempty"`
	// The policy name.
	//
	// example:
	//
	// testPolicyGroupName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether to enable network redirection.
	//
	// > This feature is in invitational preview and is not available to the public.
	//
	// example:
	//
	// off
	NetRedirect *string `json:"NetRedirect,omitempty" xml:"NetRedirect,omitempty"`
	// The preemption policy.
	//
	// > To ensure the user experience and data security of the end users who are using cloud computers, mutual preemption among multiple users is not allowed. This parameter is set to `off` by default and cannot be changed.
	//
	// example:
	//
	// off
	PreemptLogin *string `json:"PreemptLogin,omitempty" xml:"PreemptLogin,omitempty"`
	// The usernames of the users that are allowed to preempt the cloud computer. You can specify up to five usernames.
	//
	// > To ensure the user experience and data security of the end users who are using cloud computers, mutual preemption among multiple users is not allowed.
	//
	// example:
	//
	// Alice
	PreemptLoginUser []*string `json:"PreemptLoginUser,omitempty" xml:"PreemptLoginUser,omitempty" type:"Repeated"`
	// The printer redirection policy.
	//
	// example:
	//
	// on
	PrinterRedirection *string `json:"PrinterRedirection,omitempty" xml:"PrinterRedirection,omitempty"`
	// Specifies whether to enable custom screen recording.
	//
	// example:
	//
	// OFF
	RecordContent *string `json:"RecordContent,omitempty" xml:"RecordContent,omitempty"`
	// The expiration time of custom recording files. The default value is 30. Unit: days.
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
	// The option to record audio from the cloud computer.
	//
	// example:
	//
	// on
	RecordingAudio *string `json:"RecordingAudio,omitempty" xml:"RecordingAudio,omitempty"`
	// The duration for viewing the recording file. Unit: minutes. The recording file is automatically split based on the specified duration and uploaded to a bucket. If a file reaches 300 MB, it is rolled over first.
	//
	// example:
	//
	// 10
	RecordingDuration *int32 `json:"RecordingDuration,omitempty" xml:"RecordingDuration,omitempty"`
	// The time when screen recording ends. The value is in the HH:MM:SS format. This parameter is valid only when \\`Recording\\` is set to \\`PERIOD\\`.
	//
	// example:
	//
	// 08:59:00
	RecordingEndTime *string `json:"RecordingEndTime,omitempty" xml:"RecordingEndTime,omitempty"`
	// The retention period of the recording file. Valid values: 1 to 180. Unit: days.
	//
	// example:
	//
	// 15
	RecordingExpires *int64 `json:"RecordingExpires,omitempty" xml:"RecordingExpires,omitempty"`
	// The frame rate for screen recording. Unit: frames per second (fps).
	//
	// example:
	//
	// 2
	RecordingFps *int64 `json:"RecordingFps,omitempty" xml:"RecordingFps,omitempty"`
	// The time when screen recording starts. The value is in the HH:MM:SS format. This parameter is valid only when \\`Recording\\` is set to \\`PERIOD\\`.
	//
	// example:
	//
	// 08:00:00
	RecordingStartTime *string `json:"RecordingStartTime,omitempty" xml:"RecordingStartTime,omitempty"`
	// The feature that sends notifications to the client when screen recording is in progress.
	//
	// example:
	//
	// off
	RecordingUserNotify *string `json:"RecordingUserNotify,omitempty" xml:"RecordingUserNotify,omitempty"`
	// The content of the notification that is sent to the client when screen recording is in progress. You do not need to specify this parameter.
	//
	// example:
	//
	// Your cloud desktop is being recorded.
	RecordingUserNotifyMessage *string `json:"RecordingUserNotifyMessage,omitempty" xml:"RecordingUserNotifyMessage,omitempty"`
	// The region ID. Call the [DescribeRegions](~~DescribeRegions~~) operation to obtain the list of regions that support WUYING Workspace.
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
	// The scope of the policy.
	//
	// example:
	//
	// GLOBAL
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// This parameter is required when `Scope` is set to `IP`. It takes effect only when `Scope` is set to `IP`.
	ScopeValue []*string `json:"ScopeValue,omitempty" xml:"ScopeValue,omitempty" type:"Repeated"`
	// USB redirection.
	//
	// example:
	//
	// off
	UsbRedirect *string `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty"`
	// The USB redirection rules.
	UsbSupplyRedirectRule []*CreatePolicyGroupRequestUsbSupplyRedirectRule `json:"UsbSupplyRedirectRule,omitempty" xml:"UsbSupplyRedirectRule,omitempty" type:"Repeated"`
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
	// medium
	VisualQuality *string `json:"VisualQuality,omitempty" xml:"VisualQuality,omitempty"`
	// The watermark feature.
	//
	// example:
	//
	// off
	Watermark *string `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
	// The anti-screen-recording feature for invisible watermarks.
	//
	// example:
	//
	// off
	WatermarkAntiCam *string `json:"WatermarkAntiCam,omitempty" xml:"WatermarkAntiCam,omitempty"`
	// The font color of the watermark. Valid values: 0 to 16777215.
	//
	// example:
	//
	// 0
	WatermarkColor *int32 `json:"WatermarkColor,omitempty" xml:"WatermarkColor,omitempty"`
	// The rotation angle of the watermark. Valid values: -10 to -30.
	//
	// example:
	//
	// -10
	WatermarkDegree *float64 `json:"WatermarkDegree,omitempty" xml:"WatermarkDegree,omitempty"`
	// The font size of the watermark. Valid values: 10 to 20.
	//
	// example:
	//
	// 10
	WatermarkFontSize *int32 `json:"WatermarkFontSize,omitempty" xml:"WatermarkFontSize,omitempty"`
	// The font style of the watermark.
	//
	// example:
	//
	// plain
	WatermarkFontStyle *string `json:"WatermarkFontStyle,omitempty" xml:"WatermarkFontStyle,omitempty"`
	// The enhanced feature for invisible watermarks.
	//
	// example:
	//
	// medium
	WatermarkPower *string `json:"WatermarkPower,omitempty" xml:"WatermarkPower,omitempty"`
	// The number of watermark rows.
	//
	// > This parameter is not yet available.
	//
	// example:
	//
	// 5
	WatermarkRowAmount *int32 `json:"WatermarkRowAmount,omitempty" xml:"WatermarkRowAmount,omitempty"`
	// The security priority rule for invisible watermarks.
	//
	// example:
	//
	// on
	WatermarkSecurity *string `json:"WatermarkSecurity,omitempty" xml:"WatermarkSecurity,omitempty"`
	// The transparency of the watermark.
	//
	// example:
	//
	// LIGHT
	WatermarkTransparency *string `json:"WatermarkTransparency,omitempty" xml:"WatermarkTransparency,omitempty"`
	// The opacity of the watermark. A larger value indicates lower transparency. Valid values: 10 to 100.
	//
	// example:
	//
	// 10
	WatermarkTransparencyValue *int32 `json:"WatermarkTransparencyValue,omitempty" xml:"WatermarkTransparencyValue,omitempty"`
	// The type of watermark. You can specify up to three types. Separate multiple types with commas (,).
	//
	// > If you set this parameter to `custom`, you must also specify the `WatermarkCustomText` parameter.
	//
	// example:
	//
	// EndUserId
	WatermarkType *string `json:"WatermarkType,omitempty" xml:"WatermarkType,omitempty"`
	// When you connect to a cloud computer from a desktop client (including a Windows client and a macOS client), specifies whether to display the entry for the WUYING AI assistant in the floating ball on the cloud computer.
	//
	// > This feature is available only for desktop clients of V7.7 or later.
	//
	// example:
	//
	// on
	WyAssistant *string `json:"WyAssistant,omitempty" xml:"WyAssistant,omitempty"`
}

func (s CreatePolicyGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyGroupRequest) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupRequest) GetAdminAccess() *string {
	return s.AdminAccess
}

func (s *CreatePolicyGroupRequest) GetAppContentProtection() *string {
	return s.AppContentProtection
}

func (s *CreatePolicyGroupRequest) GetAuthorizeAccessPolicyRule() []*CreatePolicyGroupRequestAuthorizeAccessPolicyRule {
	return s.AuthorizeAccessPolicyRule
}

func (s *CreatePolicyGroupRequest) GetAuthorizeSecurityPolicyRule() []*CreatePolicyGroupRequestAuthorizeSecurityPolicyRule {
	return s.AuthorizeSecurityPolicyRule
}

func (s *CreatePolicyGroupRequest) GetCameraRedirect() *string {
	return s.CameraRedirect
}

func (s *CreatePolicyGroupRequest) GetClientType() []*CreatePolicyGroupRequestClientType {
	return s.ClientType
}

func (s *CreatePolicyGroupRequest) GetClipboard() *string {
	return s.Clipboard
}

func (s *CreatePolicyGroupRequest) GetDeviceRedirects() []*CreatePolicyGroupRequestDeviceRedirects {
	return s.DeviceRedirects
}

func (s *CreatePolicyGroupRequest) GetDeviceRules() []*CreatePolicyGroupRequestDeviceRules {
	return s.DeviceRules
}

func (s *CreatePolicyGroupRequest) GetDomainList() *string {
	return s.DomainList
}

func (s *CreatePolicyGroupRequest) GetDomainResolveRule() []*CreatePolicyGroupRequestDomainResolveRule {
	return s.DomainResolveRule
}

func (s *CreatePolicyGroupRequest) GetDomainResolveRuleType() *string {
	return s.DomainResolveRuleType
}

func (s *CreatePolicyGroupRequest) GetEndUserApplyAdminCoordinate() *string {
	return s.EndUserApplyAdminCoordinate
}

func (s *CreatePolicyGroupRequest) GetEndUserGroupCoordinate() *string {
	return s.EndUserGroupCoordinate
}

func (s *CreatePolicyGroupRequest) GetGpuAcceleration() *string {
	return s.GpuAcceleration
}

func (s *CreatePolicyGroupRequest) GetHtml5Access() *string {
	return s.Html5Access
}

func (s *CreatePolicyGroupRequest) GetHtml5FileTransfer() *string {
	return s.Html5FileTransfer
}

func (s *CreatePolicyGroupRequest) GetInternetCommunicationProtocol() *string {
	return s.InternetCommunicationProtocol
}

func (s *CreatePolicyGroupRequest) GetLocalDrive() *string {
	return s.LocalDrive
}

func (s *CreatePolicyGroupRequest) GetMaxReconnectTime() *int32 {
	return s.MaxReconnectTime
}

func (s *CreatePolicyGroupRequest) GetName() *string {
	return s.Name
}

func (s *CreatePolicyGroupRequest) GetNetRedirect() *string {
	return s.NetRedirect
}

func (s *CreatePolicyGroupRequest) GetPreemptLogin() *string {
	return s.PreemptLogin
}

func (s *CreatePolicyGroupRequest) GetPreemptLoginUser() []*string {
	return s.PreemptLoginUser
}

func (s *CreatePolicyGroupRequest) GetPrinterRedirection() *string {
	return s.PrinterRedirection
}

func (s *CreatePolicyGroupRequest) GetRecordContent() *string {
	return s.RecordContent
}

func (s *CreatePolicyGroupRequest) GetRecordContentExpires() *int64 {
	return s.RecordContentExpires
}

func (s *CreatePolicyGroupRequest) GetRecording() *string {
	return s.Recording
}

func (s *CreatePolicyGroupRequest) GetRecordingAudio() *string {
	return s.RecordingAudio
}

func (s *CreatePolicyGroupRequest) GetRecordingDuration() *int32 {
	return s.RecordingDuration
}

func (s *CreatePolicyGroupRequest) GetRecordingEndTime() *string {
	return s.RecordingEndTime
}

func (s *CreatePolicyGroupRequest) GetRecordingExpires() *int64 {
	return s.RecordingExpires
}

func (s *CreatePolicyGroupRequest) GetRecordingFps() *int64 {
	return s.RecordingFps
}

func (s *CreatePolicyGroupRequest) GetRecordingStartTime() *string {
	return s.RecordingStartTime
}

func (s *CreatePolicyGroupRequest) GetRecordingUserNotify() *string {
	return s.RecordingUserNotify
}

func (s *CreatePolicyGroupRequest) GetRecordingUserNotifyMessage() *string {
	return s.RecordingUserNotifyMessage
}

func (s *CreatePolicyGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePolicyGroupRequest) GetRemoteCoordinate() *string {
	return s.RemoteCoordinate
}

func (s *CreatePolicyGroupRequest) GetScope() *string {
	return s.Scope
}

func (s *CreatePolicyGroupRequest) GetScopeValue() []*string {
	return s.ScopeValue
}

func (s *CreatePolicyGroupRequest) GetUsbRedirect() *string {
	return s.UsbRedirect
}

func (s *CreatePolicyGroupRequest) GetUsbSupplyRedirectRule() []*CreatePolicyGroupRequestUsbSupplyRedirectRule {
	return s.UsbSupplyRedirectRule
}

func (s *CreatePolicyGroupRequest) GetVideoRedirect() *string {
	return s.VideoRedirect
}

func (s *CreatePolicyGroupRequest) GetVisualQuality() *string {
	return s.VisualQuality
}

func (s *CreatePolicyGroupRequest) GetWatermark() *string {
	return s.Watermark
}

func (s *CreatePolicyGroupRequest) GetWatermarkAntiCam() *string {
	return s.WatermarkAntiCam
}

func (s *CreatePolicyGroupRequest) GetWatermarkColor() *int32 {
	return s.WatermarkColor
}

func (s *CreatePolicyGroupRequest) GetWatermarkDegree() *float64 {
	return s.WatermarkDegree
}

func (s *CreatePolicyGroupRequest) GetWatermarkFontSize() *int32 {
	return s.WatermarkFontSize
}

func (s *CreatePolicyGroupRequest) GetWatermarkFontStyle() *string {
	return s.WatermarkFontStyle
}

func (s *CreatePolicyGroupRequest) GetWatermarkPower() *string {
	return s.WatermarkPower
}

func (s *CreatePolicyGroupRequest) GetWatermarkRowAmount() *int32 {
	return s.WatermarkRowAmount
}

func (s *CreatePolicyGroupRequest) GetWatermarkSecurity() *string {
	return s.WatermarkSecurity
}

func (s *CreatePolicyGroupRequest) GetWatermarkTransparency() *string {
	return s.WatermarkTransparency
}

func (s *CreatePolicyGroupRequest) GetWatermarkTransparencyValue() *int32 {
	return s.WatermarkTransparencyValue
}

func (s *CreatePolicyGroupRequest) GetWatermarkType() *string {
	return s.WatermarkType
}

func (s *CreatePolicyGroupRequest) GetWyAssistant() *string {
	return s.WyAssistant
}

func (s *CreatePolicyGroupRequest) SetAdminAccess(v string) *CreatePolicyGroupRequest {
	s.AdminAccess = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetAppContentProtection(v string) *CreatePolicyGroupRequest {
	s.AppContentProtection = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetAuthorizeAccessPolicyRule(v []*CreatePolicyGroupRequestAuthorizeAccessPolicyRule) *CreatePolicyGroupRequest {
	s.AuthorizeAccessPolicyRule = v
	return s
}

func (s *CreatePolicyGroupRequest) SetAuthorizeSecurityPolicyRule(v []*CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) *CreatePolicyGroupRequest {
	s.AuthorizeSecurityPolicyRule = v
	return s
}

func (s *CreatePolicyGroupRequest) SetCameraRedirect(v string) *CreatePolicyGroupRequest {
	s.CameraRedirect = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetClientType(v []*CreatePolicyGroupRequestClientType) *CreatePolicyGroupRequest {
	s.ClientType = v
	return s
}

func (s *CreatePolicyGroupRequest) SetClipboard(v string) *CreatePolicyGroupRequest {
	s.Clipboard = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetDeviceRedirects(v []*CreatePolicyGroupRequestDeviceRedirects) *CreatePolicyGroupRequest {
	s.DeviceRedirects = v
	return s
}

func (s *CreatePolicyGroupRequest) SetDeviceRules(v []*CreatePolicyGroupRequestDeviceRules) *CreatePolicyGroupRequest {
	s.DeviceRules = v
	return s
}

func (s *CreatePolicyGroupRequest) SetDomainList(v string) *CreatePolicyGroupRequest {
	s.DomainList = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetDomainResolveRule(v []*CreatePolicyGroupRequestDomainResolveRule) *CreatePolicyGroupRequest {
	s.DomainResolveRule = v
	return s
}

func (s *CreatePolicyGroupRequest) SetDomainResolveRuleType(v string) *CreatePolicyGroupRequest {
	s.DomainResolveRuleType = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetEndUserApplyAdminCoordinate(v string) *CreatePolicyGroupRequest {
	s.EndUserApplyAdminCoordinate = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetEndUserGroupCoordinate(v string) *CreatePolicyGroupRequest {
	s.EndUserGroupCoordinate = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetGpuAcceleration(v string) *CreatePolicyGroupRequest {
	s.GpuAcceleration = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetHtml5Access(v string) *CreatePolicyGroupRequest {
	s.Html5Access = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetHtml5FileTransfer(v string) *CreatePolicyGroupRequest {
	s.Html5FileTransfer = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetInternetCommunicationProtocol(v string) *CreatePolicyGroupRequest {
	s.InternetCommunicationProtocol = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetLocalDrive(v string) *CreatePolicyGroupRequest {
	s.LocalDrive = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetMaxReconnectTime(v int32) *CreatePolicyGroupRequest {
	s.MaxReconnectTime = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetName(v string) *CreatePolicyGroupRequest {
	s.Name = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetNetRedirect(v string) *CreatePolicyGroupRequest {
	s.NetRedirect = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetPreemptLogin(v string) *CreatePolicyGroupRequest {
	s.PreemptLogin = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetPreemptLoginUser(v []*string) *CreatePolicyGroupRequest {
	s.PreemptLoginUser = v
	return s
}

func (s *CreatePolicyGroupRequest) SetPrinterRedirection(v string) *CreatePolicyGroupRequest {
	s.PrinterRedirection = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetRecordContent(v string) *CreatePolicyGroupRequest {
	s.RecordContent = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetRecordContentExpires(v int64) *CreatePolicyGroupRequest {
	s.RecordContentExpires = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetRecording(v string) *CreatePolicyGroupRequest {
	s.Recording = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetRecordingAudio(v string) *CreatePolicyGroupRequest {
	s.RecordingAudio = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetRecordingDuration(v int32) *CreatePolicyGroupRequest {
	s.RecordingDuration = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetRecordingEndTime(v string) *CreatePolicyGroupRequest {
	s.RecordingEndTime = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetRecordingExpires(v int64) *CreatePolicyGroupRequest {
	s.RecordingExpires = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetRecordingFps(v int64) *CreatePolicyGroupRequest {
	s.RecordingFps = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetRecordingStartTime(v string) *CreatePolicyGroupRequest {
	s.RecordingStartTime = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetRecordingUserNotify(v string) *CreatePolicyGroupRequest {
	s.RecordingUserNotify = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetRecordingUserNotifyMessage(v string) *CreatePolicyGroupRequest {
	s.RecordingUserNotifyMessage = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetRegionId(v string) *CreatePolicyGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetRemoteCoordinate(v string) *CreatePolicyGroupRequest {
	s.RemoteCoordinate = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetScope(v string) *CreatePolicyGroupRequest {
	s.Scope = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetScopeValue(v []*string) *CreatePolicyGroupRequest {
	s.ScopeValue = v
	return s
}

func (s *CreatePolicyGroupRequest) SetUsbRedirect(v string) *CreatePolicyGroupRequest {
	s.UsbRedirect = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetUsbSupplyRedirectRule(v []*CreatePolicyGroupRequestUsbSupplyRedirectRule) *CreatePolicyGroupRequest {
	s.UsbSupplyRedirectRule = v
	return s
}

func (s *CreatePolicyGroupRequest) SetVideoRedirect(v string) *CreatePolicyGroupRequest {
	s.VideoRedirect = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetVisualQuality(v string) *CreatePolicyGroupRequest {
	s.VisualQuality = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetWatermark(v string) *CreatePolicyGroupRequest {
	s.Watermark = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetWatermarkAntiCam(v string) *CreatePolicyGroupRequest {
	s.WatermarkAntiCam = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetWatermarkColor(v int32) *CreatePolicyGroupRequest {
	s.WatermarkColor = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetWatermarkDegree(v float64) *CreatePolicyGroupRequest {
	s.WatermarkDegree = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetWatermarkFontSize(v int32) *CreatePolicyGroupRequest {
	s.WatermarkFontSize = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetWatermarkFontStyle(v string) *CreatePolicyGroupRequest {
	s.WatermarkFontStyle = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetWatermarkPower(v string) *CreatePolicyGroupRequest {
	s.WatermarkPower = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetWatermarkRowAmount(v int32) *CreatePolicyGroupRequest {
	s.WatermarkRowAmount = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetWatermarkSecurity(v string) *CreatePolicyGroupRequest {
	s.WatermarkSecurity = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetWatermarkTransparency(v string) *CreatePolicyGroupRequest {
	s.WatermarkTransparency = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetWatermarkTransparencyValue(v int32) *CreatePolicyGroupRequest {
	s.WatermarkTransparencyValue = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetWatermarkType(v string) *CreatePolicyGroupRequest {
	s.WatermarkType = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetWyAssistant(v string) *CreatePolicyGroupRequest {
	s.WyAssistant = &v
	return s
}

func (s *CreatePolicyGroupRequest) Validate() error {
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

type CreatePolicyGroupRequestAuthorizeAccessPolicyRule struct {
	// The client IP address CIDR block. The value is an IPv4 CIDR block.
	//
	// example:
	//
	// 47.100.XX.XX/16
	CidrIp *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	// The description of the client IP address whitelist.
	//
	// example:
	//
	// North China Branch
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s CreatePolicyGroupRequestAuthorizeAccessPolicyRule) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyGroupRequestAuthorizeAccessPolicyRule) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupRequestAuthorizeAccessPolicyRule) GetCidrIp() *string {
	return s.CidrIp
}

func (s *CreatePolicyGroupRequestAuthorizeAccessPolicyRule) GetDescription() *string {
	return s.Description
}

func (s *CreatePolicyGroupRequestAuthorizeAccessPolicyRule) SetCidrIp(v string) *CreatePolicyGroupRequestAuthorizeAccessPolicyRule {
	s.CidrIp = &v
	return s
}

func (s *CreatePolicyGroupRequestAuthorizeAccessPolicyRule) SetDescription(v string) *CreatePolicyGroupRequestAuthorizeAccessPolicyRule {
	s.Description = &v
	return s
}

func (s *CreatePolicyGroupRequestAuthorizeAccessPolicyRule) Validate() error {
	return dara.Validate(s)
}

type CreatePolicyGroupRequestAuthorizeSecurityPolicyRule struct {
	// The object of the security group rule. The value is an IPv4 CIDR block.
	//
	// example:
	//
	// 47.100.XX.XX/16
	CidrIp *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	// The description of the security group rule.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The protocol type of the security group rule.
	//
	// example:
	//
	// tcp
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The authorization policy of the security group rule.
	//
	// example:
	//
	// accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The port range of the security group rule. The value of this parameter depends on the value of the \\`IpProtocol\\` parameter.
	//
	// - If \\`IpProtocol\\` is set to TCP or UDP, the port range is 1 to 65535. Use a forward slash (/) to separate the start port and the end port. For example: 1/200.
	//
	// - If \\`IpProtocol\\` is set to ICMP, the port range is -1/-1.
	//
	// - If \\`IpProtocol\\` is set to GRE, the port range is -1/-1.
	//
	// - If \\`IpProtocol\\` is set to all, the port range is -1/-1.
	//
	// For more information about common ports, see [Common ports](https://help.aliyun.com/document_detail/40724.html).
	//
	// example:
	//
	// 22/22
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// The priority of the security group rule. A smaller value indicates a higher priority.<br>
	//
	// Valid values: 1 to 60.<br>
	//
	// Default value: 1.<br><br>
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

func (s CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) GetCidrIp() *string {
	return s.CidrIp
}

func (s *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) GetDescription() *string {
	return s.Description
}

func (s *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) GetPolicy() *string {
	return s.Policy
}

func (s *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) GetPortRange() *string {
	return s.PortRange
}

func (s *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) GetPriority() *string {
	return s.Priority
}

func (s *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) GetType() *string {
	return s.Type
}

func (s *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) SetCidrIp(v string) *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.CidrIp = &v
	return s
}

func (s *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) SetDescription(v string) *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.Description = &v
	return s
}

func (s *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) SetIpProtocol(v string) *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.IpProtocol = &v
	return s
}

func (s *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) SetPolicy(v string) *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.Policy = &v
	return s
}

func (s *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) SetPortRange(v string) *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.PortRange = &v
	return s
}

func (s *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) SetPriority(v string) *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.Priority = &v
	return s
}

func (s *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) SetType(v string) *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.Type = &v
	return s
}

func (s *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) Validate() error {
	return dara.Validate(s)
}

type CreatePolicyGroupRequestClientType struct {
	// Logon method control. Specifies the client type.
	//
	// > If you do not configure the `ClientType` parameters, all types of clients are allowed to log on to the cloud computer by default.
	//
	// example:
	//
	// windows
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// Logon method control. Specifies whether to allow a specific type of client to log on to the cloud computer.
	//
	// > If you do not configure the `ClientType` parameters, all types of clients are allowed to log on to the cloud computer by default.
	//
	// example:
	//
	// ON
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreatePolicyGroupRequestClientType) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyGroupRequestClientType) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupRequestClientType) GetClientType() *string {
	return s.ClientType
}

func (s *CreatePolicyGroupRequestClientType) GetStatus() *string {
	return s.Status
}

func (s *CreatePolicyGroupRequestClientType) SetClientType(v string) *CreatePolicyGroupRequestClientType {
	s.ClientType = &v
	return s
}

func (s *CreatePolicyGroupRequestClientType) SetStatus(v string) *CreatePolicyGroupRequestClientType {
	s.Status = &v
	return s
}

func (s *CreatePolicyGroupRequestClientType) Validate() error {
	return dara.Validate(s)
}

type CreatePolicyGroupRequestDeviceRedirects struct {
	// The peripheral type.
	//
	// example:
	//
	// camera
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// The redirection type.
	//
	// example:
	//
	// deviceRedirect
	RedirectType *string `json:"RedirectType,omitempty" xml:"RedirectType,omitempty"`
}

func (s CreatePolicyGroupRequestDeviceRedirects) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyGroupRequestDeviceRedirects) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupRequestDeviceRedirects) GetDeviceType() *string {
	return s.DeviceType
}

func (s *CreatePolicyGroupRequestDeviceRedirects) GetRedirectType() *string {
	return s.RedirectType
}

func (s *CreatePolicyGroupRequestDeviceRedirects) SetDeviceType(v string) *CreatePolicyGroupRequestDeviceRedirects {
	s.DeviceType = &v
	return s
}

func (s *CreatePolicyGroupRequestDeviceRedirects) SetRedirectType(v string) *CreatePolicyGroupRequestDeviceRedirects {
	s.RedirectType = &v
	return s
}

func (s *CreatePolicyGroupRequestDeviceRedirects) Validate() error {
	return dara.Validate(s)
}

type CreatePolicyGroupRequestDeviceRules struct {
	// The device name.
	//
	// example:
	//
	// sandisk
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The product ID (PID).
	//
	// example:
	//
	// 0x55b1
	DevicePid *string `json:"DevicePid,omitempty" xml:"DevicePid,omitempty"`
	// The peripheral type.
	//
	// example:
	//
	// storage
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// The vendor ID (VID). For more information, see [Valid USB Vendor IDs (VIDs)](https://www.usb.org/sites/default/files/vendor_ids032322.pdf_1.pdf).
	//
	// example:
	//
	// 0x0781
	DeviceVid *string `json:"DeviceVid,omitempty" xml:"DeviceVid,omitempty"`
	// The link optimization instruction.
	//
	// example:
	//
	// 2:0
	OptCommand *string `json:"OptCommand,omitempty" xml:"OptCommand,omitempty"`
	Platforms  *string `json:"Platforms,omitempty" xml:"Platforms,omitempty"`
	// The redirection type.
	//
	// example:
	//
	// usbRedirect
	RedirectType *string `json:"RedirectType,omitempty" xml:"RedirectType,omitempty"`
}

func (s CreatePolicyGroupRequestDeviceRules) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyGroupRequestDeviceRules) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupRequestDeviceRules) GetDeviceName() *string {
	return s.DeviceName
}

func (s *CreatePolicyGroupRequestDeviceRules) GetDevicePid() *string {
	return s.DevicePid
}

func (s *CreatePolicyGroupRequestDeviceRules) GetDeviceType() *string {
	return s.DeviceType
}

func (s *CreatePolicyGroupRequestDeviceRules) GetDeviceVid() *string {
	return s.DeviceVid
}

func (s *CreatePolicyGroupRequestDeviceRules) GetOptCommand() *string {
	return s.OptCommand
}

func (s *CreatePolicyGroupRequestDeviceRules) GetPlatforms() *string {
	return s.Platforms
}

func (s *CreatePolicyGroupRequestDeviceRules) GetRedirectType() *string {
	return s.RedirectType
}

func (s *CreatePolicyGroupRequestDeviceRules) SetDeviceName(v string) *CreatePolicyGroupRequestDeviceRules {
	s.DeviceName = &v
	return s
}

func (s *CreatePolicyGroupRequestDeviceRules) SetDevicePid(v string) *CreatePolicyGroupRequestDeviceRules {
	s.DevicePid = &v
	return s
}

func (s *CreatePolicyGroupRequestDeviceRules) SetDeviceType(v string) *CreatePolicyGroupRequestDeviceRules {
	s.DeviceType = &v
	return s
}

func (s *CreatePolicyGroupRequestDeviceRules) SetDeviceVid(v string) *CreatePolicyGroupRequestDeviceRules {
	s.DeviceVid = &v
	return s
}

func (s *CreatePolicyGroupRequestDeviceRules) SetOptCommand(v string) *CreatePolicyGroupRequestDeviceRules {
	s.OptCommand = &v
	return s
}

func (s *CreatePolicyGroupRequestDeviceRules) SetPlatforms(v string) *CreatePolicyGroupRequestDeviceRules {
	s.Platforms = &v
	return s
}

func (s *CreatePolicyGroupRequestDeviceRules) SetRedirectType(v string) *CreatePolicyGroupRequestDeviceRules {
	s.RedirectType = &v
	return s
}

func (s *CreatePolicyGroupRequestDeviceRules) Validate() error {
	return dara.Validate(s)
}

type CreatePolicyGroupRequestDomainResolveRule struct {
	// The policy description.
	//
	// example:
	//
	// 测试规则
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

func (s CreatePolicyGroupRequestDomainResolveRule) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyGroupRequestDomainResolveRule) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupRequestDomainResolveRule) GetDescription() *string {
	return s.Description
}

func (s *CreatePolicyGroupRequestDomainResolveRule) GetDomain() *string {
	return s.Domain
}

func (s *CreatePolicyGroupRequestDomainResolveRule) GetPolicy() *string {
	return s.Policy
}

func (s *CreatePolicyGroupRequestDomainResolveRule) SetDescription(v string) *CreatePolicyGroupRequestDomainResolveRule {
	s.Description = &v
	return s
}

func (s *CreatePolicyGroupRequestDomainResolveRule) SetDomain(v string) *CreatePolicyGroupRequestDomainResolveRule {
	s.Domain = &v
	return s
}

func (s *CreatePolicyGroupRequestDomainResolveRule) SetPolicy(v string) *CreatePolicyGroupRequestDomainResolveRule {
	s.Policy = &v
	return s
}

func (s *CreatePolicyGroupRequestDomainResolveRule) Validate() error {
	return dara.Validate(s)
}

type CreatePolicyGroupRequestUsbSupplyRedirectRule struct {
	// The rule description.
	//
	// example:
	//
	// 测试规则
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The device class. This parameter is required when `usbRuleType` is set to 1. For more information, see [Defined Class Codes](https://www.usb.org/defined-class-codes).
	//
	// example:
	//
	// 0Eh
	DeviceClass *string `json:"DeviceClass,omitempty" xml:"DeviceClass,omitempty"`
	// The device subclass. This parameter is required when `usbRuleType` is set to 1. For more information, see [Defined Class Codes](https://www.usb.org/defined-class-codes).
	//
	// example:
	//
	// xxh
	DeviceSubclass *string `json:"DeviceSubclass,omitempty" xml:"DeviceSubclass,omitempty"`
	// The product ID (PID).
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
	// The vendor ID (VID). For more information, see [Valid USB Vendor IDs (VIDs)](https://www.usb.org/sites/default/files/vendor_ids032322.pdf_1.pdf).
	//
	// example:
	//
	// 04**
	VendorId *string `json:"VendorId,omitempty" xml:"VendorId,omitempty"`
}

func (s CreatePolicyGroupRequestUsbSupplyRedirectRule) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyGroupRequestUsbSupplyRedirectRule) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupRequestUsbSupplyRedirectRule) GetDescription() *string {
	return s.Description
}

func (s *CreatePolicyGroupRequestUsbSupplyRedirectRule) GetDeviceClass() *string {
	return s.DeviceClass
}

func (s *CreatePolicyGroupRequestUsbSupplyRedirectRule) GetDeviceSubclass() *string {
	return s.DeviceSubclass
}

func (s *CreatePolicyGroupRequestUsbSupplyRedirectRule) GetProductId() *string {
	return s.ProductId
}

func (s *CreatePolicyGroupRequestUsbSupplyRedirectRule) GetUsbRedirectType() *int64 {
	return s.UsbRedirectType
}

func (s *CreatePolicyGroupRequestUsbSupplyRedirectRule) GetUsbRuleType() *int64 {
	return s.UsbRuleType
}

func (s *CreatePolicyGroupRequestUsbSupplyRedirectRule) GetVendorId() *string {
	return s.VendorId
}

func (s *CreatePolicyGroupRequestUsbSupplyRedirectRule) SetDescription(v string) *CreatePolicyGroupRequestUsbSupplyRedirectRule {
	s.Description = &v
	return s
}

func (s *CreatePolicyGroupRequestUsbSupplyRedirectRule) SetDeviceClass(v string) *CreatePolicyGroupRequestUsbSupplyRedirectRule {
	s.DeviceClass = &v
	return s
}

func (s *CreatePolicyGroupRequestUsbSupplyRedirectRule) SetDeviceSubclass(v string) *CreatePolicyGroupRequestUsbSupplyRedirectRule {
	s.DeviceSubclass = &v
	return s
}

func (s *CreatePolicyGroupRequestUsbSupplyRedirectRule) SetProductId(v string) *CreatePolicyGroupRequestUsbSupplyRedirectRule {
	s.ProductId = &v
	return s
}

func (s *CreatePolicyGroupRequestUsbSupplyRedirectRule) SetUsbRedirectType(v int64) *CreatePolicyGroupRequestUsbSupplyRedirectRule {
	s.UsbRedirectType = &v
	return s
}

func (s *CreatePolicyGroupRequestUsbSupplyRedirectRule) SetUsbRuleType(v int64) *CreatePolicyGroupRequestUsbSupplyRedirectRule {
	s.UsbRuleType = &v
	return s
}

func (s *CreatePolicyGroupRequestUsbSupplyRedirectRule) SetVendorId(v string) *CreatePolicyGroupRequestUsbSupplyRedirectRule {
	s.VendorId = &v
	return s
}

func (s *CreatePolicyGroupRequestUsbSupplyRedirectRule) Validate() error {
	return dara.Validate(s)
}
