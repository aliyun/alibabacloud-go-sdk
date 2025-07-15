// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCenterPolicyListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescribePolicyGroups(v []*DescribeCenterPolicyListResponseBodyDescribePolicyGroups) *DescribeCenterPolicyListResponseBody
	GetDescribePolicyGroups() []*DescribeCenterPolicyListResponseBodyDescribePolicyGroups
	SetRequestId(v string) *DescribeCenterPolicyListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeCenterPolicyListResponseBody
	GetTotalCount() *int32
}

type DescribeCenterPolicyListResponseBody struct {
	// The cloud computer policies.
	DescribePolicyGroups []*DescribeCenterPolicyListResponseBodyDescribePolicyGroups `json:"DescribePolicyGroups,omitempty" xml:"DescribePolicyGroups,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCenterPolicyListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenterPolicyListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCenterPolicyListResponseBody) GetDescribePolicyGroups() []*DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	return s.DescribePolicyGroups
}

func (s *DescribeCenterPolicyListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCenterPolicyListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCenterPolicyListResponseBody) SetDescribePolicyGroups(v []*DescribeCenterPolicyListResponseBodyDescribePolicyGroups) *DescribeCenterPolicyListResponseBody {
	s.DescribePolicyGroups = v
	return s
}

func (s *DescribeCenterPolicyListResponseBody) SetRequestId(v string) *DescribeCenterPolicyListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBody) SetTotalCount(v int32) *DescribeCenterPolicyListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCenterPolicyListResponseBodyDescribePolicyGroups struct {
	// Indicates whether the admin permissions are granted to end users.
	//
	// >  This parameter is in private preview and only available to specific users.
	//
	// example:
	//
	// deny
	AdminAccess *string `json:"AdminAccess,omitempty" xml:"AdminAccess,omitempty"`
	// Indicates whether anti-screenshot is enabled.
	//
	// example:
	//
	// off
	AppContentProtection *string `json:"AppContentProtection,omitempty" xml:"AppContentProtection,omitempty"`
	// The client IP address whitelists.
	AuthorizeAccessPolicyRules []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules `json:"AuthorizeAccessPolicyRules,omitempty" xml:"AuthorizeAccessPolicyRules,omitempty" type:"Repeated"`
	// The security group rules.
	AuthorizeSecurityPolicyRules []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules `json:"AuthorizeSecurityPolicyRules,omitempty" xml:"AuthorizeSecurityPolicyRules,omitempty" type:"Repeated"`
	// Indicates whether on-premises webcam redirection is enabled.
	//
	// example:
	//
	// on
	CameraRedirect    *string `json:"CameraRedirect,omitempty" xml:"CameraRedirect,omitempty"`
	ClientControlMenu *string `json:"ClientControlMenu,omitempty" xml:"ClientControlMenu,omitempty"`
	// The logon method control rules.
	ClientTypes []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsClientTypes `json:"ClientTypes,omitempty" xml:"ClientTypes,omitempty" type:"Repeated"`
	// The read/write permissions on the clipboard.
	//
	// example:
	//
	// off
	Clipboard *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	// Indicates whether color enhancement is enabled for design and 3D applications.
	//
	// example:
	//
	// off
	ColorEnhancement  *string `json:"ColorEnhancement,omitempty" xml:"ColorEnhancement,omitempty"`
	CpdDriveClipboard *string `json:"CpdDriveClipboard,omitempty" xml:"CpdDriveClipboard,omitempty"`
	// The CPU underclocking duration. Valid values: 30 to 120. Unit: seconds.
	//
	// example:
	//
	// 30
	CpuDownGradeDuration *int32 `json:"CpuDownGradeDuration,omitempty" xml:"CpuDownGradeDuration,omitempty"`
	// The CPU processors.
	CpuProcessors []*string `json:"CpuProcessors,omitempty" xml:"CpuProcessors,omitempty" type:"Repeated"`
	// The CPU spike protection policy.
	//
	// example:
	//
	// on
	CpuProtectedMode *string `json:"CpuProtectedMode,omitempty" xml:"CpuProtectedMode,omitempty"`
	// The overall CPU usage. Valid values: 70 to 90. Unit: percentage (%).
	//
	// example:
	//
	// 70
	CpuRateLimit *int32 `json:"CpuRateLimit,omitempty" xml:"CpuRateLimit,omitempty"`
	// The overall CPU sampling duration. Valid values: 10 to 60. Unit: seconds.
	//
	// example:
	//
	// 10
	CpuSampleDuration *int32 `json:"CpuSampleDuration,omitempty" xml:"CpuSampleDuration,omitempty"`
	// The single-CPU usage. Valid values: 70 to 100. Unit: %.
	//
	// example:
	//
	// 70
	CpuSingleRateLimit *int32 `json:"CpuSingleRateLimit,omitempty" xml:"CpuSingleRateLimit,omitempty"`
	// The number of cloud computers that are associated with the policy.
	//
	// example:
	//
	// 1
	DesktopCount *int32 `json:"DesktopCount,omitempty" xml:"DesktopCount,omitempty"`
	// The number of cloud computer shares that are associated with the policy.
	//
	// example:
	//
	// 1
	DesktopGroupCount *int32 `json:"DesktopGroupCount,omitempty" xml:"DesktopGroupCount,omitempty"`
	// The device redirection rules.
	DeviceRedirects []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRedirects `json:"DeviceRedirects,omitempty" xml:"DeviceRedirects,omitempty" type:"Repeated"`
	// The custom peripheral rules.
	DeviceRules []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules `json:"DeviceRules,omitempty" xml:"DeviceRules,omitempty" type:"Repeated"`
	// Indicates whether the session is retained after disconnection.
	//
	// >  This parameter applies only to cloud application policies.
	//
	// example:
	//
	// persistent
	DisconnectKeepSession *string `json:"DisconnectKeepSession,omitempty" xml:"DisconnectKeepSession,omitempty"`
	// The retention period of the session after disconnection. Unit: seconds.
	//
	// >  This parameter applies only to cloud application policies.
	//
	// example:
	//
	// 120
	DisconnectKeepSessionTime *int32 `json:"DisconnectKeepSessionTime,omitempty" xml:"DisconnectKeepSessionTime,omitempty"`
	// The display mode.
	//
	// example:
	//
	// adminCustom
	DisplayMode *string `json:"DisplayMode,omitempty" xml:"DisplayMode,omitempty"`
	// The field where the domain resolution policy is applied.
	//
	// example:
	//
	// xxxx
	DomainRegisterValue *string `json:"DomainRegisterValue,omitempty" xml:"DomainRegisterValue,omitempty"`
	// The domain resolution policies.
	DomainResolveRule []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDomainResolveRule `json:"DomainResolveRule,omitempty" xml:"DomainResolveRule,omitempty" type:"Repeated"`
	// Indicates whether domain name resolution is allowed.
	//
	// example:
	//
	// on
	DomainResolveRuleType *string `json:"DomainResolveRuleType,omitempty" xml:"DomainResolveRuleType,omitempty"`
	// Indicates whether end users are allowed to request administrator help.
	//
	// example:
	//
	// off
	EndUserApplyAdminCoordinate *string `json:"EndUserApplyAdminCoordinate,omitempty" xml:"EndUserApplyAdminCoordinate,omitempty"`
	// Indicates whether end users in the same office network can share cloud computers.
	//
	// example:
	//
	// off
	EndUserGroupCoordinate    *string `json:"EndUserGroupCoordinate,omitempty" xml:"EndUserGroupCoordinate,omitempty"`
	FileTransferAddress       *string `json:"FileTransferAddress,omitempty" xml:"FileTransferAddress,omitempty"`
	FileTransferSpeed         *string `json:"FileTransferSpeed,omitempty" xml:"FileTransferSpeed,omitempty"`
	FileTransferSpeedLocation *string `json:"FileTransferSpeedLocation,omitempty" xml:"FileTransferSpeedLocation,omitempty"`
	// Indicates whether image quality control is enabled. For optimal computer performance and user experience in professional design scenarios, we recommend enabling this feature.
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
	// The file transfer feature on the web client.
	//
	// example:
	//
	// off
	Html5FileTransfer *string `json:"Html5FileTransfer,omitempty" xml:"Html5FileTransfer,omitempty"`
	// The network communication protocol.
	//
	// example:
	//
	// tcp
	InternetCommunicationProtocol *string `json:"InternetCommunicationProtocol,omitempty" xml:"InternetCommunicationProtocol,omitempty"`
	InternetPrinter               *string `json:"InternetPrinter,omitempty" xml:"InternetPrinter,omitempty"`
	// The read/write permissions on the on-premises drive.
	//
	// example:
	//
	// readwrite
	LocalDrive *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	// The maximum duration to retry reconnecting to cloud computers after an unexpected disconnection (non-human causes). Valid values: 30 to 7200. Unit: seconds.
	//
	// example:
	//
	// 120
	MaxReconnectTime *int32 `json:"MaxReconnectTime,omitempty" xml:"MaxReconnectTime,omitempty"`
	// The memory underclocking duration per process. Valid values: 30 to 120. Unit: seconds.
	//
	// example:
	//
	// 30
	MemoryDownGradeDuration *int32 `json:"MemoryDownGradeDuration,omitempty" xml:"MemoryDownGradeDuration,omitempty"`
	// The memory processors.
	MemoryProcessors []*string `json:"MemoryProcessors,omitempty" xml:"MemoryProcessors,omitempty" type:"Repeated"`
	// The memory spike protection policy.
	//
	// example:
	//
	// off
	MemoryProtectedMode *string `json:"MemoryProtectedMode,omitempty" xml:"MemoryProtectedMode,omitempty"`
	// The overall memory usage. Valid values: 70 to 90. Unit: %.
	//
	// example:
	//
	// 70
	MemoryRateLimit *int32 `json:"MemoryRateLimit,omitempty" xml:"MemoryRateLimit,omitempty"`
	// The overall memory sampling duration. Valid values: 30 to 60. Unit: seconds.
	//
	// example:
	//
	// 30
	MemorySampleDuration *int32 `json:"MemorySampleDuration,omitempty" xml:"MemorySampleDuration,omitempty"`
	// The memory usage per process. Valid values: 30 to 60. Unit: %.
	//
	// example:
	//
	// 30
	MemorySingleRateLimit *int32 `json:"MemorySingleRateLimit,omitempty" xml:"MemorySingleRateLimit,omitempty"`
	// Indicates whether the Restart button is displayed in the DesktopAssistant menu when end users connect to cloud computers from Android clients.
	//
	// >  This feature applies to only mobile clients of version 7.4.0 or later.
	//
	// example:
	//
	// off
	MobileRestart *string `json:"MobileRestart,omitempty" xml:"MobileRestart,omitempty"`
	// Indicates whether the Stop button is displayed in the DesktopAssistant menu when end users connect to cloud computers from Android clients.
	//
	// >  This feature applies to only mobile clients of version 7.4.0 or later.
	//
	// example:
	//
	// off
	MobileShutdown *string `json:"MobileShutdown,omitempty" xml:"MobileShutdown,omitempty"`
	// The policy name.
	//
	// example:
	//
	// testPolicyGroupName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The network redirection policy.
	//
	// >  This parameter is in private preview and only available to specific users.
	//
	// example:
	//
	// off
	NetRedirect *string `json:"NetRedirect,omitempty" xml:"NetRedirect,omitempty"`
	// The network redirection policies.
	//
	// >  This parameter is in private preview and only available to specific users.
	NetRedirectRule []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsNetRedirectRule `json:"NetRedirectRule,omitempty" xml:"NetRedirectRule,omitempty" type:"Repeated"`
	// Indicates whether a disconnection is enforced upon inactivity.
	//
	// >  This parameter applies only to cloud application policies.
	//
	// example:
	//
	// off
	NoOperationDisconnect *string `json:"NoOperationDisconnect,omitempty" xml:"NoOperationDisconnect,omitempty"`
	// The duration of disconnection after inactivity. Unit: seconds.
	//
	// >  This parameter applies only to cloud application policies.
	//
	// example:
	//
	// 120
	NoOperationDisconnectTime *int32 `json:"NoOperationDisconnectTime,omitempty" xml:"NoOperationDisconnectTime,omitempty"`
	// The policy ID.
	//
	// example:
	//
	// pg-gx2x1dhsmthe9****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The policy type.
	//
	// example:
	//
	// SYSTEM
	PolicyGroupType *string `json:"PolicyGroupType,omitempty" xml:"PolicyGroupType,omitempty"`
	// The status of the cloud computer policy.
	//
	// example:
	//
	// AVAILABLE
	PolicyStatus *string `json:"PolicyStatus,omitempty" xml:"PolicyStatus,omitempty"`
	// The printer redirection policy.
	//
	// example:
	//
	// off
	PrinterRedirection *string `json:"PrinterRedirection,omitempty" xml:"PrinterRedirection,omitempty"`
	// Indicates whether image quality enhancement is enabled for design and 3D applications.
	//
	// example:
	//
	// off
	QualityEnhancement *string `json:"QualityEnhancement,omitempty" xml:"QualityEnhancement,omitempty"`
	// Indicates whether custom screen recording is enabled.
	//
	// example:
	//
	// off
	RecordContent *string `json:"RecordContent,omitempty" xml:"RecordContent,omitempty"`
	// The duration for which custom screen recordings are kept before they expire. Default value: 30 days.
	//
	// example:
	//
	// 30
	RecordContentExpires *int64 `json:"RecordContentExpires,omitempty" xml:"RecordContentExpires,omitempty"`
	// The duration of screen recording after the specified event is detected. Unit: minutes. Valid values: 10 to 60.
	//
	// example:
	//
	// 10
	RecordEventDuration *int32 `json:"RecordEventDuration,omitempty" xml:"RecordEventDuration,omitempty"`
	// The absolute paths to screen recording files.
	RecordEventFilePaths []*string `json:"RecordEventFilePaths,omitempty" xml:"RecordEventFilePaths,omitempty" type:"Repeated"`
	// The absolute paths to screen recording registries.
	RecordEventRegisters []*string `json:"RecordEventRegisters,omitempty" xml:"RecordEventRegisters,omitempty" type:"Repeated"`
	// Indicates whether screen recording is enabled.
	//
	// example:
	//
	// off
	Recording *string `json:"Recording,omitempty" xml:"Recording,omitempty"`
	// Indicates whether audio files generated on cloud computers are recorded.
	//
	// example:
	//
	// on
	RecordingAudio *string `json:"RecordingAudio,omitempty" xml:"RecordingAudio,omitempty"`
	// The length of the screen recording file. Unit: minutes. Screen recording files are split by the specified length and uploaded to OSS buckets. Once a file reaches 300 MB, the system prioritizes rolling updates for that file.
	//
	// example:
	//
	// 10
	RecordingDuration *int32 `json:"RecordingDuration,omitempty" xml:"RecordingDuration,omitempty"`
	// The end time of screen recording. The value is in the HH:MM:SS format. The value is meaningful only when you set Recording to period.
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
	// example:
	//
	// 5
	RecordingFps *int64 `json:"RecordingFps,omitempty" xml:"RecordingFps,omitempty"`
	// The start time of screen recording. The value is in the HH:MM:SS format. The value is meaningful only when you set Recording to period.
	//
	// example:
	//
	// 08:00:00
	RecordingStartTime *string `json:"RecordingStartTime,omitempty" xml:"RecordingStartTime,omitempty"`
	// Indicates whether to notify end users when screen recording is enabled.
	//
	// example:
	//
	// off
	RecordingUserNotify *string `json:"RecordingUserNotify,omitempty" xml:"RecordingUserNotify,omitempty"`
	// The notification sent to end users when screen recording is enabled.
	RecordingUserNotifyMessage *string `json:"RecordingUserNotifyMessage,omitempty" xml:"RecordingUserNotifyMessage,omitempty"`
	// The keyboard and mouse control permissions during remote assistance.
	//
	// example:
	//
	// fullControl
	RemoteCoordinate *string `json:"RemoteCoordinate,omitempty" xml:"RemoteCoordinate,omitempty"`
	// The height of the resolution. Unit: pixel.
	//
	// example:
	//
	// 1280
	ResolutionHeight *int32 `json:"ResolutionHeight,omitempty" xml:"ResolutionHeight,omitempty"`
	// The resolution type.
	//
	// example:
	//
	// adaptive
	ResolutionModel *string `json:"ResolutionModel,omitempty" xml:"ResolutionModel,omitempty"`
	// The width of the resolution. Unit: pixel.
	//
	// example:
	//
	// 1920
	ResolutionWidth *int32 `json:"ResolutionWidth,omitempty" xml:"ResolutionWidth,omitempty"`
	// The number of resource groups that are associated with the policy.
	//
	// example:
	//
	// 1
	ResourceGroupCount *int32  `json:"ResourceGroupCount,omitempty" xml:"ResourceGroupCount,omitempty"`
	SafeMenu           *string `json:"SafeMenu,omitempty" xml:"SafeMenu,omitempty"`
	// The effective scope of the policy.
	//
	// example:
	//
	// GLOBAL
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The effective scopes specified by CIDR blocks.
	ScopeValue        []*string `json:"ScopeValue,omitempty" xml:"ScopeValue,omitempty" type:"Repeated"`
	ScreenDisplayMode *string   `json:"ScreenDisplayMode,omitempty" xml:"ScreenDisplayMode,omitempty"`
	// Indicates whether smoothness enhancement is enabled for daily office use.
	//
	// example:
	//
	// off
	SmoothEnhancement *string `json:"SmoothEnhancement,omitempty" xml:"SmoothEnhancement,omitempty"`
	// Indicates whether the metric status entry is displayed in the DesktopAssistant menu.
	//
	// example:
	//
	// on
	StatusMonitor *string `json:"StatusMonitor,omitempty" xml:"StatusMonitor,omitempty"`
	// The streaming mode.
	//
	// example:
	//
	// smooth
	StreamingMode *string `json:"StreamingMode,omitempty" xml:"StreamingMode,omitempty"`
	// The target frame rate. Valid values: 10 to 60. Unit: fps.
	//
	// example:
	//
	// 30
	TargetFps *int32 `json:"TargetFps,omitempty" xml:"TargetFps,omitempty"`
	// Indicates whether the application taskbar is displayed.
	//
	// >  This parameter applies only to cloud application policies.
	//
	// example:
	//
	// off
	Taskbar *string `json:"Taskbar,omitempty" xml:"Taskbar,omitempty"`
	// The USB redirection policy.
	//
	// example:
	//
	// off
	UsbRedirect *string `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty"`
	// The USB redirection rules.
	UsbSupplyRedirectRule []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule `json:"UsbSupplyRedirectRule,omitempty" xml:"UsbSupplyRedirectRule,omitempty" type:"Repeated"`
	UseTime               *string                                                                          `json:"UseTime,omitempty" xml:"UseTime,omitempty"`
	// The average bitrate for video encoding. Unit: Kbit/s. Valid values: 1000 to 50000.
	//
	// example:
	//
	// 1000
	VideoEncAvgKbps *int32 `json:"VideoEncAvgKbps,omitempty" xml:"VideoEncAvgKbps,omitempty"`
	// The maximum QP for video files. Higher QP values result in lower video quality. Valid values: 0 to 51.
	//
	// example:
	//
	// 20
	VideoEncMaxQP *int32 `json:"VideoEncMaxQP,omitempty" xml:"VideoEncMaxQP,omitempty"`
	// The minimum quantizer parameter (QP) for video files. A lower QP means better video quality. Valid values: 0 to 51.
	//
	// example:
	//
	// 20
	VideoEncMinQP *int32 `json:"VideoEncMinQP,omitempty" xml:"VideoEncMinQP,omitempty"`
	// The peak bitrate for video encoding. Unit: Kbit/s. Valid values: 1000 to 50000.
	//
	// example:
	//
	// 2000
	VideoEncPeakKbps *int32 `json:"VideoEncPeakKbps,omitempty" xml:"VideoEncPeakKbps,omitempty"`
	// The video encoding policy.
	//
	// example:
	//
	// qualityFirst
	VideoEncPolicy *string `json:"VideoEncPolicy,omitempty" xml:"VideoEncPolicy,omitempty"`
	// Indicates whether multimedia redirection is enabled.
	//
	// example:
	//
	// off
	VideoRedirect *string `json:"VideoRedirect,omitempty" xml:"VideoRedirect,omitempty"`
	// The image quality policy.
	//
	// example:
	//
	// medium
	VisualQuality *string `json:"VisualQuality,omitempty" xml:"VisualQuality,omitempty"`
	// The watermark policy.
	//
	// example:
	//
	// on
	Watermark *string `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
	// Indicates whether anti-screen capture is enabled for invisible watermarks.
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
	// If you set `WatermarkType` to `custom`, you must also specify `WatermarkCustomText`.
	//
	// example:
	//
	// custom-watermark
	WatermarkCustomText *string `json:"WatermarkCustomText,omitempty" xml:"WatermarkCustomText,omitempty"`
	// The watermark rotation. Valid values: -10 to -30.
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
	// The enhancement level for invisible watermarks.
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
	// Indicates whether security priority is enabled for invisible watermarks.
	//
	// example:
	//
	// on
	WatermarkSecurity *string `json:"WatermarkSecurity,omitempty" xml:"WatermarkSecurity,omitempty"`
	// The watermark transparency. A higher value means the watermark is less transparent. Valid values: 10 to 100.
	//
	// example:
	//
	// 10
	WatermarkTransparencyValue *int32 `json:"WatermarkTransparencyValue,omitempty" xml:"WatermarkTransparencyValue,omitempty"`
	// The watermark type.
	//
	// example:
	//
	// EndUserId
	WatermarkType *string `json:"WatermarkType,omitempty" xml:"WatermarkType,omitempty"`
	// Indicates whether the Xiaoying AI Assistant entry is displayed in the DesktopAssistant menu.
	//
	// example:
	//
	// on
	WyAssistant *string `json:"WyAssistant,omitempty" xml:"WyAssistant,omitempty"`
}

func (s DescribeCenterPolicyListResponseBodyDescribePolicyGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GoString() string {
	return s.String()
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetAdminAccess() *string {
	return s.AdminAccess
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetAppContentProtection() *string {
	return s.AppContentProtection
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetAuthorizeAccessPolicyRules() []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules {
	return s.AuthorizeAccessPolicyRules
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetAuthorizeSecurityPolicyRules() []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	return s.AuthorizeSecurityPolicyRules
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetCameraRedirect() *string {
	return s.CameraRedirect
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetClientControlMenu() *string {
	return s.ClientControlMenu
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetClientTypes() []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsClientTypes {
	return s.ClientTypes
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetClipboard() *string {
	return s.Clipboard
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetColorEnhancement() *string {
	return s.ColorEnhancement
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetCpdDriveClipboard() *string {
	return s.CpdDriveClipboard
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetCpuDownGradeDuration() *int32 {
	return s.CpuDownGradeDuration
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetCpuProcessors() []*string {
	return s.CpuProcessors
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetCpuProtectedMode() *string {
	return s.CpuProtectedMode
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetCpuRateLimit() *int32 {
	return s.CpuRateLimit
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetCpuSampleDuration() *int32 {
	return s.CpuSampleDuration
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetCpuSingleRateLimit() *int32 {
	return s.CpuSingleRateLimit
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetDesktopCount() *int32 {
	return s.DesktopCount
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetDesktopGroupCount() *int32 {
	return s.DesktopGroupCount
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetDeviceRedirects() []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRedirects {
	return s.DeviceRedirects
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetDeviceRules() []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules {
	return s.DeviceRules
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetDisconnectKeepSession() *string {
	return s.DisconnectKeepSession
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetDisconnectKeepSessionTime() *int32 {
	return s.DisconnectKeepSessionTime
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetDisplayMode() *string {
	return s.DisplayMode
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetDomainRegisterValue() *string {
	return s.DomainRegisterValue
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetDomainResolveRule() []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDomainResolveRule {
	return s.DomainResolveRule
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetDomainResolveRuleType() *string {
	return s.DomainResolveRuleType
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetEndUserApplyAdminCoordinate() *string {
	return s.EndUserApplyAdminCoordinate
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetEndUserGroupCoordinate() *string {
	return s.EndUserGroupCoordinate
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetFileTransferAddress() *string {
	return s.FileTransferAddress
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetFileTransferSpeed() *string {
	return s.FileTransferSpeed
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetFileTransferSpeedLocation() *string {
	return s.FileTransferSpeedLocation
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetGpuAcceleration() *string {
	return s.GpuAcceleration
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetHtml5Access() *string {
	return s.Html5Access
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetHtml5FileTransfer() *string {
	return s.Html5FileTransfer
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetInternetCommunicationProtocol() *string {
	return s.InternetCommunicationProtocol
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetInternetPrinter() *string {
	return s.InternetPrinter
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetLocalDrive() *string {
	return s.LocalDrive
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetMaxReconnectTime() *int32 {
	return s.MaxReconnectTime
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetMemoryDownGradeDuration() *int32 {
	return s.MemoryDownGradeDuration
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetMemoryProcessors() []*string {
	return s.MemoryProcessors
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetMemoryProtectedMode() *string {
	return s.MemoryProtectedMode
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetMemoryRateLimit() *int32 {
	return s.MemoryRateLimit
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetMemorySampleDuration() *int32 {
	return s.MemorySampleDuration
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetMemorySingleRateLimit() *int32 {
	return s.MemorySingleRateLimit
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetMobileRestart() *string {
	return s.MobileRestart
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetMobileShutdown() *string {
	return s.MobileShutdown
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetName() *string {
	return s.Name
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetNetRedirect() *string {
	return s.NetRedirect
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetNetRedirectRule() []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsNetRedirectRule {
	return s.NetRedirectRule
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetNoOperationDisconnect() *string {
	return s.NoOperationDisconnect
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetNoOperationDisconnectTime() *int32 {
	return s.NoOperationDisconnectTime
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetPolicyGroupType() *string {
	return s.PolicyGroupType
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetPolicyStatus() *string {
	return s.PolicyStatus
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetPrinterRedirection() *string {
	return s.PrinterRedirection
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetQualityEnhancement() *string {
	return s.QualityEnhancement
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetRecordContent() *string {
	return s.RecordContent
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetRecordContentExpires() *int64 {
	return s.RecordContentExpires
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetRecordEventDuration() *int32 {
	return s.RecordEventDuration
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetRecordEventFilePaths() []*string {
	return s.RecordEventFilePaths
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetRecordEventRegisters() []*string {
	return s.RecordEventRegisters
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetRecording() *string {
	return s.Recording
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetRecordingAudio() *string {
	return s.RecordingAudio
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetRecordingDuration() *int32 {
	return s.RecordingDuration
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetRecordingEndTime() *string {
	return s.RecordingEndTime
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetRecordingExpires() *int64 {
	return s.RecordingExpires
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetRecordingFps() *int64 {
	return s.RecordingFps
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetRecordingStartTime() *string {
	return s.RecordingStartTime
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetRecordingUserNotify() *string {
	return s.RecordingUserNotify
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetRecordingUserNotifyMessage() *string {
	return s.RecordingUserNotifyMessage
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetRemoteCoordinate() *string {
	return s.RemoteCoordinate
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetResolutionHeight() *int32 {
	return s.ResolutionHeight
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetResolutionModel() *string {
	return s.ResolutionModel
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetResolutionWidth() *int32 {
	return s.ResolutionWidth
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetResourceGroupCount() *int32 {
	return s.ResourceGroupCount
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetSafeMenu() *string {
	return s.SafeMenu
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetScope() *string {
	return s.Scope
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetScopeValue() []*string {
	return s.ScopeValue
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetScreenDisplayMode() *string {
	return s.ScreenDisplayMode
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetSmoothEnhancement() *string {
	return s.SmoothEnhancement
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetStatusMonitor() *string {
	return s.StatusMonitor
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetStreamingMode() *string {
	return s.StreamingMode
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetTargetFps() *int32 {
	return s.TargetFps
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetTaskbar() *string {
	return s.Taskbar
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetUsbRedirect() *string {
	return s.UsbRedirect
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetUsbSupplyRedirectRule() []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule {
	return s.UsbSupplyRedirectRule
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetUseTime() *string {
	return s.UseTime
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetVideoEncAvgKbps() *int32 {
	return s.VideoEncAvgKbps
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetVideoEncMaxQP() *int32 {
	return s.VideoEncMaxQP
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetVideoEncMinQP() *int32 {
	return s.VideoEncMinQP
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetVideoEncPeakKbps() *int32 {
	return s.VideoEncPeakKbps
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetVideoEncPolicy() *string {
	return s.VideoEncPolicy
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetVideoRedirect() *string {
	return s.VideoRedirect
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetVisualQuality() *string {
	return s.VisualQuality
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetWatermark() *string {
	return s.Watermark
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetWatermarkAntiCam() *string {
	return s.WatermarkAntiCam
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetWatermarkColor() *int32 {
	return s.WatermarkColor
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetWatermarkCustomText() *string {
	return s.WatermarkCustomText
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetWatermarkDegree() *float64 {
	return s.WatermarkDegree
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetWatermarkFontSize() *int32 {
	return s.WatermarkFontSize
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetWatermarkFontStyle() *string {
	return s.WatermarkFontStyle
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetWatermarkPower() *string {
	return s.WatermarkPower
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetWatermarkRowAmount() *int32 {
	return s.WatermarkRowAmount
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetWatermarkSecurity() *string {
	return s.WatermarkSecurity
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetWatermarkTransparencyValue() *int32 {
	return s.WatermarkTransparencyValue
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetWatermarkType() *string {
	return s.WatermarkType
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetWyAssistant() *string {
	return s.WyAssistant
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetAdminAccess(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.AdminAccess = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetAppContentProtection(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.AppContentProtection = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetAuthorizeAccessPolicyRules(v []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.AuthorizeAccessPolicyRules = v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetAuthorizeSecurityPolicyRules(v []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.AuthorizeSecurityPolicyRules = v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetCameraRedirect(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.CameraRedirect = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetClientControlMenu(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.ClientControlMenu = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetClientTypes(v []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsClientTypes) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.ClientTypes = v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetClipboard(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.Clipboard = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetColorEnhancement(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.ColorEnhancement = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetCpdDriveClipboard(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.CpdDriveClipboard = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetCpuDownGradeDuration(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.CpuDownGradeDuration = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetCpuProcessors(v []*string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.CpuProcessors = v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetCpuProtectedMode(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.CpuProtectedMode = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetCpuRateLimit(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.CpuRateLimit = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetCpuSampleDuration(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.CpuSampleDuration = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetCpuSingleRateLimit(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.CpuSingleRateLimit = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetDesktopCount(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.DesktopCount = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetDesktopGroupCount(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.DesktopGroupCount = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetDeviceRedirects(v []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRedirects) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.DeviceRedirects = v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetDeviceRules(v []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.DeviceRules = v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetDisconnectKeepSession(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.DisconnectKeepSession = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetDisconnectKeepSessionTime(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.DisconnectKeepSessionTime = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetDisplayMode(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.DisplayMode = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetDomainRegisterValue(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.DomainRegisterValue = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetDomainResolveRule(v []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDomainResolveRule) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.DomainResolveRule = v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetDomainResolveRuleType(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.DomainResolveRuleType = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetEndUserApplyAdminCoordinate(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.EndUserApplyAdminCoordinate = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetEndUserGroupCoordinate(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.EndUserGroupCoordinate = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetFileTransferAddress(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.FileTransferAddress = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetFileTransferSpeed(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.FileTransferSpeed = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetFileTransferSpeedLocation(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.FileTransferSpeedLocation = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetGpuAcceleration(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.GpuAcceleration = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetHtml5Access(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.Html5Access = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetHtml5FileTransfer(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.Html5FileTransfer = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetInternetCommunicationProtocol(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.InternetCommunicationProtocol = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetInternetPrinter(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.InternetPrinter = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetLocalDrive(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.LocalDrive = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetMaxReconnectTime(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.MaxReconnectTime = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetMemoryDownGradeDuration(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.MemoryDownGradeDuration = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetMemoryProcessors(v []*string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.MemoryProcessors = v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetMemoryProtectedMode(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.MemoryProtectedMode = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetMemoryRateLimit(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.MemoryRateLimit = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetMemorySampleDuration(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.MemorySampleDuration = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetMemorySingleRateLimit(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.MemorySingleRateLimit = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetMobileRestart(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.MobileRestart = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetMobileShutdown(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.MobileShutdown = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetName(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.Name = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetNetRedirect(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.NetRedirect = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetNetRedirectRule(v []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsNetRedirectRule) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.NetRedirectRule = v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetNoOperationDisconnect(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.NoOperationDisconnect = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetNoOperationDisconnectTime(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.NoOperationDisconnectTime = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetPolicyGroupId(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.PolicyGroupId = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetPolicyGroupType(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.PolicyGroupType = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetPolicyStatus(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.PolicyStatus = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetPrinterRedirection(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.PrinterRedirection = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetQualityEnhancement(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.QualityEnhancement = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetRecordContent(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.RecordContent = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetRecordContentExpires(v int64) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.RecordContentExpires = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetRecordEventDuration(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.RecordEventDuration = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetRecordEventFilePaths(v []*string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.RecordEventFilePaths = v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetRecordEventRegisters(v []*string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.RecordEventRegisters = v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetRecording(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.Recording = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetRecordingAudio(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.RecordingAudio = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetRecordingDuration(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.RecordingDuration = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetRecordingEndTime(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.RecordingEndTime = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetRecordingExpires(v int64) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.RecordingExpires = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetRecordingFps(v int64) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.RecordingFps = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetRecordingStartTime(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.RecordingStartTime = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetRecordingUserNotify(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.RecordingUserNotify = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetRecordingUserNotifyMessage(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.RecordingUserNotifyMessage = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetRemoteCoordinate(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.RemoteCoordinate = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetResolutionHeight(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.ResolutionHeight = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetResolutionModel(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.ResolutionModel = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetResolutionWidth(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.ResolutionWidth = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetResourceGroupCount(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.ResourceGroupCount = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetSafeMenu(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.SafeMenu = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetScope(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.Scope = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetScopeValue(v []*string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.ScopeValue = v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetScreenDisplayMode(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.ScreenDisplayMode = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetSmoothEnhancement(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.SmoothEnhancement = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetStatusMonitor(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.StatusMonitor = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetStreamingMode(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.StreamingMode = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetTargetFps(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.TargetFps = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetTaskbar(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.Taskbar = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetUsbRedirect(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.UsbRedirect = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetUsbSupplyRedirectRule(v []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.UsbSupplyRedirectRule = v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetUseTime(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.UseTime = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetVideoEncAvgKbps(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.VideoEncAvgKbps = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetVideoEncMaxQP(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.VideoEncMaxQP = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetVideoEncMinQP(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.VideoEncMinQP = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetVideoEncPeakKbps(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.VideoEncPeakKbps = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetVideoEncPolicy(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.VideoEncPolicy = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetVideoRedirect(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.VideoRedirect = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetVisualQuality(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.VisualQuality = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetWatermark(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.Watermark = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetWatermarkAntiCam(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.WatermarkAntiCam = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetWatermarkColor(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.WatermarkColor = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetWatermarkCustomText(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.WatermarkCustomText = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetWatermarkDegree(v float64) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.WatermarkDegree = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetWatermarkFontSize(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.WatermarkFontSize = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetWatermarkFontStyle(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.WatermarkFontStyle = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetWatermarkPower(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.WatermarkPower = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetWatermarkRowAmount(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.WatermarkRowAmount = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetWatermarkSecurity(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.WatermarkSecurity = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetWatermarkTransparencyValue(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.WatermarkTransparencyValue = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetWatermarkType(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.WatermarkType = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetWyAssistant(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.WyAssistant = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) Validate() error {
	return dara.Validate(s)
}

type DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules struct {
	// The client CIDR block from which end users can connect to cloud computers. The value is an IPv4 CIDR block.
	//
	// example:
	//
	// 47.100.XX.XX/16
	CidrIp *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	// The remarks on the client CIDR block.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules) GoString() string {
	return s.String()
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules) GetCidrIp() *string {
	return s.CidrIp
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules) GetDescription() *string {
	return s.Description
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules) SetCidrIp(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules {
	s.CidrIp = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules) SetDescription(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules {
	s.Description = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules) Validate() error {
	return dara.Validate(s)
}

type DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules struct {
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
	// The port range of the security group rule.
	//
	// example:
	//
	// 22/22
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// The priority of the security group rule. A smaller value indicates a higher priority.
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

func (s DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) GoString() string {
	return s.String()
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) GetCidrIp() *string {
	return s.CidrIp
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) GetDescription() *string {
	return s.Description
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) GetPolicy() *string {
	return s.Policy
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) GetPortRange() *string {
	return s.PortRange
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) GetPriority() *string {
	return s.Priority
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) GetType() *string {
	return s.Type
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetCidrIp(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.CidrIp = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetDescription(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.Description = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetIpProtocol(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.IpProtocol = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetPolicy(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.Policy = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetPortRange(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.PortRange = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetPriority(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.Priority = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetType(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.Type = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) Validate() error {
	return dara.Validate(s)
}

type DescribeCenterPolicyListResponseBodyDescribePolicyGroupsClientTypes struct {
	// The client type.
	//
	// example:
	//
	// windows
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// Indicates whether a specific client type can connect to cloud computers.
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCenterPolicyListResponseBodyDescribePolicyGroupsClientTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenterPolicyListResponseBodyDescribePolicyGroupsClientTypes) GoString() string {
	return s.String()
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsClientTypes) GetClientType() *string {
	return s.ClientType
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsClientTypes) GetStatus() *string {
	return s.Status
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsClientTypes) SetClientType(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsClientTypes {
	s.ClientType = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsClientTypes) SetStatus(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsClientTypes {
	s.Status = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsClientTypes) Validate() error {
	return dara.Validate(s)
}

type DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRedirects struct {
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
	// usbRedirect
	RedirectType *string `json:"RedirectType,omitempty" xml:"RedirectType,omitempty"`
}

func (s DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRedirects) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRedirects) GoString() string {
	return s.String()
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRedirects) GetDeviceType() *string {
	return s.DeviceType
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRedirects) GetRedirectType() *string {
	return s.RedirectType
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRedirects) SetDeviceType(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRedirects {
	s.DeviceType = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRedirects) SetRedirectType(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRedirects {
	s.RedirectType = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRedirects) Validate() error {
	return dara.Validate(s)
}

type DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules struct {
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
	// example:
	//
	// usbRedirect
	RedirectType *string `json:"RedirectType,omitempty" xml:"RedirectType,omitempty"`
}

func (s DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules) GoString() string {
	return s.String()
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules) GetDeviceName() *string {
	return s.DeviceName
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules) GetDevicePid() *string {
	return s.DevicePid
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules) GetDeviceType() *string {
	return s.DeviceType
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules) GetDeviceVid() *string {
	return s.DeviceVid
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules) GetOptCommand() *string {
	return s.OptCommand
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules) GetPlatforms() *string {
	return s.Platforms
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules) GetRedirectType() *string {
	return s.RedirectType
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules) SetDeviceName(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules {
	s.DeviceName = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules) SetDevicePid(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules {
	s.DevicePid = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules) SetDeviceType(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules {
	s.DeviceType = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules) SetDeviceVid(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules {
	s.DeviceVid = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules) SetOptCommand(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules {
	s.OptCommand = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules) SetPlatforms(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules {
	s.Platforms = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules) SetRedirectType(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules {
	s.RedirectType = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules) Validate() error {
	return dara.Validate(s)
}

type DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDomainResolveRule struct {
	// The policy description.
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

func (s DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDomainResolveRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDomainResolveRule) GoString() string {
	return s.String()
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDomainResolveRule) GetDescription() *string {
	return s.Description
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDomainResolveRule) GetDomain() *string {
	return s.Domain
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDomainResolveRule) GetPolicy() *string {
	return s.Policy
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDomainResolveRule) SetDescription(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDomainResolveRule {
	s.Description = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDomainResolveRule) SetDomain(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDomainResolveRule {
	s.Domain = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDomainResolveRule) SetPolicy(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDomainResolveRule {
	s.Policy = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDomainResolveRule) Validate() error {
	return dara.Validate(s)
}

type DescribeCenterPolicyListResponseBodyDescribePolicyGroupsNetRedirectRule struct {
	// The domain name.
	//
	// example:
	//
	// *.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The redirection policy.
	//
	// example:
	//
	// allow
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The rule type.
	//
	// example:
	//
	// domain
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s DescribeCenterPolicyListResponseBodyDescribePolicyGroupsNetRedirectRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenterPolicyListResponseBodyDescribePolicyGroupsNetRedirectRule) GoString() string {
	return s.String()
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsNetRedirectRule) GetDomain() *string {
	return s.Domain
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsNetRedirectRule) GetPolicy() *string {
	return s.Policy
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsNetRedirectRule) GetRuleType() *string {
	return s.RuleType
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsNetRedirectRule) SetDomain(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsNetRedirectRule {
	s.Domain = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsNetRedirectRule) SetPolicy(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsNetRedirectRule {
	s.Policy = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsNetRedirectRule) SetRuleType(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsNetRedirectRule {
	s.RuleType = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsNetRedirectRule) Validate() error {
	return dara.Validate(s)
}

type DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule struct {
	// The rule description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The product ID (PID).
	//
	// example:
	//
	// 08**
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// Indicates whether USB redirection is allowed.
	//
	// example:
	//
	// 1
	UsbRedirectType *int64 `json:"UsbRedirectType,omitempty" xml:"UsbRedirectType,omitempty"`
	// The type of the USB redirection rule.
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

func (s DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) GoString() string {
	return s.String()
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) GetDescription() *string {
	return s.Description
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) GetProductId() *string {
	return s.ProductId
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) GetUsbRedirectType() *int64 {
	return s.UsbRedirectType
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) GetUsbRuleType() *int64 {
	return s.UsbRuleType
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) GetVendorId() *string {
	return s.VendorId
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) SetDescription(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule {
	s.Description = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) SetProductId(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule {
	s.ProductId = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) SetUsbRedirectType(v int64) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule {
	s.UsbRedirectType = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) SetUsbRuleType(v int64) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule {
	s.UsbRuleType = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) SetVendorId(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule {
	s.VendorId = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) Validate() error {
	return dara.Validate(s)
}
