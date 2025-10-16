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
	// Specifies whether end users have the administrator permissions.
	//
	// >  This parameter is in invitational preview for specific users and not available to the public.
	//
	// example:
	//
	// deny
	AdminAccess *string `json:"AdminAccess,omitempty" xml:"AdminAccess,omitempty"`
	// Specifies whether to enable the anti-screenshot feature.
	//
	// Valid values:
	//
	// 	- off: Anti-screenshot is disabled. This value is the default value.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- on: Anti-screenshot is enabled.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// off
	AppContentProtection *string `json:"AppContentProtection,omitempty" xml:"AppContentProtection,omitempty"`
	// The client IP address whitelist. After you configure the whitelist, end users can access cloud computers only from the IP addresses in the whitelist.
	AuthorizeAccessPolicyRule []*CreatePolicyGroupRequestAuthorizeAccessPolicyRule `json:"AuthorizeAccessPolicyRule,omitempty" xml:"AuthorizeAccessPolicyRule,omitempty" type:"Repeated"`
	// The security group rules.
	AuthorizeSecurityPolicyRule []*CreatePolicyGroupRequestAuthorizeSecurityPolicyRule `json:"AuthorizeSecurityPolicyRule,omitempty" xml:"AuthorizeSecurityPolicyRule,omitempty" type:"Repeated"`
	// Specifies whether to enable the webcam redirection feature.
	//
	// Valid values:
	//
	// 	- off: Webcam redirection is disabled.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- on: Webcam redirection is enabled. This value is the default value.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// on
	CameraRedirect *string `json:"CameraRedirect,omitempty" xml:"CameraRedirect,omitempty"`
	// The logon method control rules to limit the type of the Alibaba Cloud Workspace client used by end users to connect to cloud computers.
	ClientType []*CreatePolicyGroupRequestClientType `json:"ClientType,omitempty" xml:"ClientType,omitempty" type:"Repeated"`
	// The permissions on the clipboard.
	//
	// Valid values:
	//
	// 	- read: specifies one-way transfer. You can copy files only from local devices to cloud computers.
	//
	// 	- readwrite: specifies two-way transfer. You can copy files between local devices and cloud computers.
	//
	// 	- write: specifies one-way transfer. You can only copy files from cloud computers to local devices.
	//
	// 	- off (default): disables both one-way and two-way transfer. Files cannot be copied between local devices and cloud computers.
	//
	// example:
	//
	// off
	Clipboard *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	// The device redirection rules.
	DeviceRedirects []*CreatePolicyGroupRequestDeviceRedirects `json:"DeviceRedirects,omitempty" xml:"DeviceRedirects,omitempty" type:"Repeated"`
	// The custom peripheral rules.
	DeviceRules []*CreatePolicyGroupRequestDeviceRules `json:"DeviceRules,omitempty" xml:"DeviceRules,omitempty" type:"Repeated"`
	// Specifies whether the access control for domain names is enabled. Domain names support wildcards (\\*). Separate multiple domain names with commas (,).
	//
	// Valid values:
	//
	// 	- off
	//
	// 	- on
	//
	// example:
	//
	// off
	DomainList *string `json:"DomainList,omitempty" xml:"DomainList,omitempty"`
	// The details of the domain name resolution rule.
	DomainResolveRule []*CreatePolicyGroupRequestDomainResolveRule `json:"DomainResolveRule,omitempty" xml:"DomainResolveRule,omitempty" type:"Repeated"`
	// The type of the domain name resolution policy.
	//
	// Valid values:
	//
	// 	- OFF
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- ON
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// OFF
	DomainResolveRuleType *string `json:"DomainResolveRuleType,omitempty" xml:"DomainResolveRuleType,omitempty"`
	// Specifies whether to turn on the Contact Administrator for Help switch.
	//
	// Valid values:
	//
	// 	- OFF
	//
	// 	- ON
	//
	// example:
	//
	// ON
	EndUserApplyAdminCoordinate *string `json:"EndUserApplyAdminCoordinate,omitempty" xml:"EndUserApplyAdminCoordinate,omitempty"`
	// Specifies whether to turn on the User Stream Collaboration switch.
	//
	// Valid values:
	//
	// 	- OFF
	//
	// 	- ON
	//
	// example:
	//
	// ON
	EndUserGroupCoordinate *string `json:"EndUserGroupCoordinate,omitempty" xml:"EndUserGroupCoordinate,omitempty"`
	// Specifies whether to enable the Image Quality Control feature. If you have high requirements on the performance and user experience in scenarios such as professional design, we recommend that you enable this feature.
	//
	// Valid values:
	//
	// 	- off
	//
	// 	- on
	//
	// example:
	//
	// off
	GpuAcceleration *string `json:"GpuAcceleration,omitempty" xml:"GpuAcceleration,omitempty"`
	// Specifies whether to allow web client access.
	//
	// >  We recommend that you use the ClientType-related parameters to control the Alibaba Cloud Workspace client type for cloud computer logon.``
	//
	// Valid values:
	//
	// 	- off (default)
	//
	// 	- on
	//
	// example:
	//
	// off
	Html5Access *string `json:"Html5Access,omitempty" xml:"Html5Access,omitempty"`
	// The file transfer feature on the web client.
	//
	// Valid values:
	//
	// 	- all: Files can be uploaded and downloaded between local computers and the web client.
	//
	// 	- download: Files on the web client can be downloaded to local computers.
	//
	// 	- upload: Files on local computers can be uploaded to the web client.
	//
	// 	- off (default): Files cannot be transferred between the web client and local computers.
	//
	// example:
	//
	// off
	Html5FileTransfer *string `json:"Html5FileTransfer,omitempty" xml:"Html5FileTransfer,omitempty"`
	// The protocol for network communication.
	//
	// Valid values:
	//
	// 	- TCP (default): TCP
	//
	// 	- BOTH: TCP and UDP
	//
	// example:
	//
	// both
	InternetCommunicationProtocol *string `json:"InternetCommunicationProtocol,omitempty" xml:"InternetCommunicationProtocol,omitempty"`
	// The permissions on local disk mapping.
	//
	// Valid values:
	//
	// 	- read: read-only. Local disk mapping is available on cloud computers. However, you can only read (copy) local files but cannot modify the files.
	//
	// 	- readwrite: read and write. Local disk mapping is available on cloud computers. You can read (copy) and write (modify) local files.
	//
	// 	- off (default): disabled. Local disk mapping is unavailable on cloud computers.
	//
	// example:
	//
	// off
	LocalDrive *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	// The maximum retry period for reconnecting to cloud computers when the cloud computers are disconnected due to none-human reasons. Valid values: 30 to 7200. Unit: seconds.
	//
	// example:
	//
	// 120
	MaxReconnectTime *int32 `json:"MaxReconnectTime,omitempty" xml:"MaxReconnectTime,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// testPolicyGroupName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether to enable the network redirection feature.
	//
	// > This feature is in invitational preview and is not available to the public.
	//
	// Valid values:
	//
	// 	- off (default): The network redirection feature is disabled.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- on: The network redirection feature is enabled.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// off
	NetRedirect *string `json:"NetRedirect,omitempty" xml:"NetRedirect,omitempty"`
	// The cloud computer preemption feature.
	//
	// >  To ensure user experience and data security, when a cloud computer is used by an end user, other end users cannot connect to the cloud computer. By default, this parameter is set to `off`, which cannot be modified.
	//
	// Valid values:
	//
	// 	- off (default): Multiple end users cannot connect to the same cloud computer at the same time.
	//
	// example:
	//
	// off
	PreemptLogin *string `json:"PreemptLogin,omitempty" xml:"PreemptLogin,omitempty"`
	// The usernames that are allowed to connect to the cloud computer in use. You can specify up to five usernames.
	//
	// >  To ensure user experience and data security, other end users cannot connect to the cloud computer that is used by an end user.
	//
	// example:
	//
	// Alice
	PreemptLoginUser []*string `json:"PreemptLoginUser,omitempty" xml:"PreemptLoginUser,omitempty" type:"Repeated"`
	// The policy for printer redirection.
	//
	// Valid values:
	//
	// 	- off: Printer redirection is disabled.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- on: Printer redirection is enabled.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// on
	PrinterRedirection *string `json:"PrinterRedirection,omitempty" xml:"PrinterRedirection,omitempty"`
	// Specifies whether to enable the custom screen recording feature.
	//
	// Valid values:
	//
	// 	- off: Custom screen recording is disabled. This value is the default value.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- on: Custom screen recording is enabled.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// OFF
	RecordContent *string `json:"RecordContent,omitempty" xml:"RecordContent,omitempty"`
	// The duration in which the custom screen recording is valid. Default value: 30. Unit: days.
	//
	// example:
	//
	// 30
	RecordContentExpires *int64 `json:"RecordContentExpires,omitempty" xml:"RecordContentExpires,omitempty"`
	// Specifies whether to enable the screen recording feature.
	//
	// Valid values:
	//
	// 	- byaction_cmd_ft: enables the operation-triggered screen recording upon command execution and file transfer.
	//
	// 	- ALLTIME: enables the whole-process screen recording. That is, the recording starts when cloud computers are connected and ends when the cloud computers are disconnected.
	//
	// 	- session: enables the screen recording for session lifecycle listening.
	//
	// 	- PERIOD: enables the interval-based screen recording. You must specify an interval between the start time and end time of this type of recording.
	//
	// 	- byaction_commands: enables the operation-triggered screen recording upon command execution.
	//
	// 	- OFF: disables the screen recording feature.
	//
	// 	- byaction_file_transfer: enables the operation-triggered screen recording upon file transfer.
	//
	// example:
	//
	// OFF
	Recording *string `json:"Recording,omitempty" xml:"Recording,omitempty"`
	// Specifies whether to record audio files generated from cloud computers.
	//
	// Valid values:
	//
	// 	- off: records only video files.
	//
	// 	- on: records video and audio files.
	//
	// example:
	//
	// on
	RecordingAudio *string `json:"RecordingAudio,omitempty" xml:"RecordingAudio,omitempty"`
	// The file length of the screen recording. Unit: minutes. Screen recording files are split based on the specified file length and uploaded to Object Storage Service (OSS) buckets. When a screen recording file reaches 300 MB in size, the system preferentially performs rolling update for the file.
	//
	// Valid values:
	//
	// 	- 10
	//
	// 	- 20
	//
	// 	- 30
	//
	// 	- 60
	//
	// example:
	//
	// 15
	RecordingDuration *int32 `json:"RecordingDuration,omitempty" xml:"RecordingDuration,omitempty"`
	// The time when the screen recording ends. The value is in the HH:MM:SS format. The value is meaningful only when you set the `Recording` parameter to `PERIOD`.
	//
	// example:
	//
	// 08:59:00
	RecordingEndTime *string `json:"RecordingEndTime,omitempty" xml:"RecordingEndTime,omitempty"`
	// The retention period of the screen recording file. Valid values: 1 to 180. Unit: days.
	//
	// example:
	//
	// 15
	RecordingExpires *int64 `json:"RecordingExpires,omitempty" xml:"RecordingExpires,omitempty"`
	// The frame rate of screen recording. Unit: fps.
	//
	// Valid values:
	//
	// 	- 2
	//
	// 	- 5
	//
	// 	- 10
	//
	// 	- 15
	//
	// example:
	//
	// 2
	RecordingFps *int64 `json:"RecordingFps,omitempty" xml:"RecordingFps,omitempty"`
	// The time when the screen recording starts. The value is in the HH:MM:SS format. The value is meaningful only when you set the `Recording` parameter to `PERIOD`.
	//
	// example:
	//
	// 08:00:00
	RecordingStartTime *string `json:"RecordingStartTime,omitempty" xml:"RecordingStartTime,omitempty"`
	// Specifies whether to enable the screen recording notification feature after end users log on to the Alibaba Cloud Workspace client.
	//
	// Valid values:
	//
	// 	- off
	//
	// 	- on
	//
	// example:
	//
	// off
	RecordingUserNotify *string `json:"RecordingUserNotify,omitempty" xml:"RecordingUserNotify,omitempty"`
	// The notification content of screen recording. By default, this parameter is left empty.
	//
	// example:
	//
	// Your desktop is being recorded.
	RecordingUserNotifyMessage *string `json:"RecordingUserNotifyMessage,omitempty" xml:"RecordingUserNotifyMessage,omitempty"`
	// The region ID. You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the regions supported by Elastic Desktop Service (EDS).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The permission to control the keyboard and the mouse during remote assistance.
	//
	// Valid values:
	//
	// 	- optionalControl: By default, this feature is disabled. You can enable it by applying permissions.
	//
	// 	- fullControl: The permission is granted.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- disableControl: The permission is revoked.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// fullControl
	RemoteCoordinate *string `json:"RemoteCoordinate,omitempty" xml:"RemoteCoordinate,omitempty"`
	// The effective scope of the policy.
	//
	// Valid values:
	//
	// 	- IP: The policy takes effect based on the IP address.
	//
	// 	- GLOBAL: The policy takes effect globally.
	//
	// example:
	//
	// GLOBAL
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// This parameter is required when the `Scope` parameter is set to `IP`.````
	ScopeValue []*string `json:"ScopeValue,omitempty" xml:"ScopeValue,omitempty" type:"Repeated"`
	// Specifies whether to enable USB redirection.
	//
	// Valid values:
	//
	// 	- off: USB redirection is disabled. This value is the default value.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- on: USB redirection is enabled.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// off
	UsbRedirect *string `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty"`
	// The USB redirection rules.
	UsbSupplyRedirectRule []*CreatePolicyGroupRequestUsbSupplyRedirectRule `json:"UsbSupplyRedirectRule,omitempty" xml:"UsbSupplyRedirectRule,omitempty" type:"Repeated"`
	// Specifies whether to enable the multimedia redirection switch.
	//
	// Valid values:
	//
	// 	- off
	//
	// 	- on
	//
	// example:
	//
	// on
	VideoRedirect *string `json:"VideoRedirect,omitempty" xml:"VideoRedirect,omitempty"`
	// The policy for image display quality.
	//
	// Valid values:
	//
	// 	- high
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- low
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- lossless
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- medium: adaptive. This value is the default value.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// medium
	VisualQuality *string `json:"VisualQuality,omitempty" xml:"VisualQuality,omitempty"`
	// The watermarking feature.
	//
	// Valid values:
	//
	// 	- blind: Invisible watermarks are applied.
	//
	// 	- off (default): The watermarking feature is disabled.
	//
	// 	- on: Visible watermarks are applied.
	//
	// example:
	//
	// off
	Watermark *string `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
	// Specifies whether to enable the anti-screen photo feature for invisible watermarks.
	//
	// Valid values:
	//
	// 	- off
	//
	// 	- on
	//
	// example:
	//
	// off
	WatermarkAntiCam *string `json:"WatermarkAntiCam,omitempty" xml:"WatermarkAntiCam,omitempty"`
	// The font color in red, green, and blue (RGB) of the watermark. Valid values: 0 to 16777215.
	//
	// example:
	//
	// 0
	WatermarkColor *int32 `json:"WatermarkColor,omitempty" xml:"WatermarkColor,omitempty"`
	// The watermark rotation. Valid values: -10 to -30.
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
	// Valid values:
	//
	// 	- plain
	//
	// 	- bold
	//
	// example:
	//
	// plain
	WatermarkFontStyle *string `json:"WatermarkFontStyle,omitempty" xml:"WatermarkFontStyle,omitempty"`
	// The watermark enhancement feature.
	//
	// Valid values:
	//
	// 	- high
	//
	// 	- low
	//
	// 	- medium
	//
	// example:
	//
	// medium
	WatermarkPower *string `json:"WatermarkPower,omitempty" xml:"WatermarkPower,omitempty"`
	// The number of watermark rows.
	//
	// >  This parameter is not available for public use.
	//
	// example:
	//
	// 5
	WatermarkRowAmount *int32 `json:"WatermarkRowAmount,omitempty" xml:"WatermarkRowAmount,omitempty"`
	// Specifies whether to enable the security priority feature for invisible watermarks.
	//
	// Valid values:
	//
	// 	- off
	//
	// 	- on
	//
	// example:
	//
	// on
	WatermarkSecurity *string `json:"WatermarkSecurity,omitempty" xml:"WatermarkSecurity,omitempty"`
	// The transparency of the watermark.
	//
	// Valid values:
	//
	// 	- LIGHT
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- DARK
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- MIDDLE
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// LIGHT
	WatermarkTransparency *string `json:"WatermarkTransparency,omitempty" xml:"WatermarkTransparency,omitempty"`
	// The watermark opacity. A larger value indicates more opaque watermarks. Valid values: 10 to 100.
	//
	// example:
	//
	// 10
	WatermarkTransparencyValue *int32 `json:"WatermarkTransparencyValue,omitempty" xml:"WatermarkTransparencyValue,omitempty"`
	// The watermark content. You can select up to three items as the watermark content. Separate multiple items with commas (,).
	//
	// >  If you set this parameter to `Custom`, specify `WatermarkCustomText`
	//
	// Valid values:
	//
	// 	- EndUserId: the username.
	//
	// 	- Custom: the custom text.
	//
	// 	- DesktopIp: the IP address of the cloud computer.
	//
	// 	- ClientIp: the IP address of the Alibaba Cloud Workspace client.
	//
	// 	- HostName: the rightmost 15 digits of the cloud computer ID.
	//
	// 	- ClientTime: the current time displayed on the cloud computer.
	//
	// example:
	//
	// EndUserId
	WatermarkType *string `json:"WatermarkType,omitempty" xml:"WatermarkType,omitempty"`
	// Specifies whether to provide the AI Assistant function in the DesktopAssistant when the cloud computer is accessed from the Alibaba Cloud Workspace desktop clients (including the Windows client and the macOS client).
	//
	// > Desktop clients of V7.7 and higher versions required.
	//
	// Valid values:
	//
	// - off: the AI Aisstant function is not provided.
	//
	// - on: the AI Aisstant function is provided.
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
	// The client CIDR block from which end users can connect to cloud computers. The value is an IPv4 CIDR block.
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
	// The object to which the security group rule applies. The value is an IPv4 CIDR block.
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
	// Valid values:
	//
	// 	- TCP: the Transmission Control Protocol (TCP) protocol.
	//
	// 	- UDP: the User Datagram Protocol (UDP) protocol.
	//
	// 	- ALL: all protocols.
	//
	// 	- GRE: the Generic Routing Encapsulation (GRE) protocol.
	//
	// 	- ICMP: the Internet Control Message Protocol (ICMP) for IPv4.
	//
	// example:
	//
	// tcp
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The authorization of the security group rule.
	//
	// Valid values:
	//
	// 	- drop: denies all access requests. If no messages of access denied are returned, the requests timed out or failed.
	//
	// 	- accept (default): accepts all requests.
	//
	// example:
	//
	// accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The port range of the security group rule. The value range of this parameter varies based on the value of the IpProtocol parameter.
	//
	// 	- If the IpProtocol parameter is set to TCP or UDP, the port range is 1 to 65535. Separate the start port number and the end port number with a forward slash (/). Example: 1/200.
	//
	// 	- If the IpProtocol parameter is set to ICMP, set the value to -1/-1.
	//
	// 	- If the IpProtocol parameter is set to GRE, set the value to -1/-1.
	//
	// 	- If the IpProtocol parameter is set to ALL, set the value to -1/-1.
	//
	// For more information about the common ports applied in EDS, see [Common ports](https://help.aliyun.com/document_detail/40724.html).
	//
	// example:
	//
	// 22/22
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// The priority of the security group rule. A smaller value indicates a higher priority.\\
	//
	// Valid values: 1 to 60.\\
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The direction of the security group rule.
	//
	// Valid values:
	//
	// 	- outflow: outbound.
	//
	// 	- inflow: inbound.
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
	// The type of the Alibaba Cloud Workspace client.
	//
	// >  If you do not specify the `ClientType` parameter, all types of the client are allowed by default.
	//
	// Valid values:
	//
	// 	- html5: web client
	//
	// 	- android: Android client
	//
	// 	- ios: iOS client
	//
	// 	- windows: Windows client
	//
	// 	- macos: macOS client
	//
	// example:
	//
	// windows
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// Specifies whether to allow end users to use a specific type of the client to connect to cloud computers.
	//
	// >  If you do not specify the `ClientType` parameter, all types of the client are allowed by default.
	//
	// Valid values:
	//
	// 	- OFF
	//
	// 	- ON
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
	// Valid values:
	//
	// 	- printer
	//
	// 	- scanner
	//
	// 	- camera
	//
	// 	- adb: the Android Debug Bridge (ADB) device.
	//
	// example:
	//
	// camera
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// The redirection type.
	//
	// Valid values:
	//
	// 	- deviceRedirect: device redirection
	//
	// 	- usbRedirect: USB redirection
	//
	// 	- off: redirection disabled
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
	// Valid values:
	//
	// 	- usbKey: UKeys.
	//
	// 	- other: other peripheral devices.
	//
	// 	- graphicsTablet: graphics tablets.
	//
	// 	- printer: printers.
	//
	// 	- cardReader: card readers.
	//
	// 	- scanner: scanners.
	//
	// 	- storage: storage devices.
	//
	// 	- camera: web cameras.
	//
	// 	- adb: Android Debug Bridge (ADB) devices.
	//
	// 	- networkInterfaceCard: NIC devices.
	//
	// example:
	//
	// storage
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// The vendor ID (VID). For more information, see [Valid USB VIDs](https://www.usb.org/sites/default/files/vendor_ids032322.pdf_1.pdf).
	//
	// example:
	//
	// 0x0781
	DeviceVid *string `json:"DeviceVid,omitempty" xml:"DeviceVid,omitempty"`
	// The link optimization command.
	//
	// example:
	//
	// 2:0
	OptCommand *string `json:"OptCommand,omitempty" xml:"OptCommand,omitempty"`
	Platforms  *string `json:"Platforms,omitempty" xml:"Platforms,omitempty"`
	// The redirection type.
	//
	// Valid values:
	//
	// 	- deviceRedirect: device redirection
	//
	// 	- usbRedirect: USB redirection
	//
	// 	- off: redirection disabled
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
	// The description of domain name resolution rule.
	//
	// example:
	//
	// system policy
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The domain name.
	//
	// example:
	//
	// *.baidu.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Specifies whether to allow the domain name resolution rule.
	//
	// Valid values:
	//
	// 	- allow: allows the rule.
	//
	// 	- block: denies the rule.
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
	// The description of the rule.
	//
	// example:
	//
	// Test rule
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The class of the device. If you set the `usbRuleType` parameter to 1, you must specify this parameter. For more information, see [Defined Class Codes](https://www.usb.org/defined-class-codes).
	//
	// example:
	//
	// 0Eh
	DeviceClass *string `json:"DeviceClass,omitempty" xml:"DeviceClass,omitempty"`
	// The subclass of the device. If you set the `usbRuleType` parameter to 1, you must specify this parameter. For more information, see [Defined Class Codes](https://www.usb.org/defined-class-codes).
	//
	// example:
	//
	// xxh
	DeviceSubclass *string `json:"DeviceSubclass,omitempty" xml:"DeviceSubclass,omitempty"`
	// The ID of the service.
	//
	// example:
	//
	// 08**
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The type of USB redirection.
	//
	// Valid values:
	//
	// 	- 1: allows USB redirection
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- 2: forbids USB redirection
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// 1
	UsbRedirectType *int64 `json:"UsbRedirectType,omitempty" xml:"UsbRedirectType,omitempty"`
	// The type of the USB redirection rule.
	//
	// Valid values:
	//
	// 	- 1: by device class
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- 2: by device vendor
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// 1
	UsbRuleType *int64 `json:"UsbRuleType,omitempty" xml:"UsbRuleType,omitempty"`
	// The ID of the vendor. For more information, see [Valid USB Vendor IDs (VIDs)](https://www.usb.org/sites/default/files/vendor_ids032322.pdf_1.pdf).
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
