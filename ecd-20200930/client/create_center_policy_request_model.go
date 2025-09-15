// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCenterPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdminAccess(v string) *CreateCenterPolicyRequest
	GetAdminAccess() *string
	SetAppContentProtection(v string) *CreateCenterPolicyRequest
	GetAppContentProtection() *string
	SetAuthorizeAccessPolicyRule(v []*CreateCenterPolicyRequestAuthorizeAccessPolicyRule) *CreateCenterPolicyRequest
	GetAuthorizeAccessPolicyRule() []*CreateCenterPolicyRequestAuthorizeAccessPolicyRule
	SetAuthorizeSecurityPolicyRule(v []*CreateCenterPolicyRequestAuthorizeSecurityPolicyRule) *CreateCenterPolicyRequest
	GetAuthorizeSecurityPolicyRule() []*CreateCenterPolicyRequestAuthorizeSecurityPolicyRule
	SetAutoReconnect(v string) *CreateCenterPolicyRequest
	GetAutoReconnect() *string
	SetBusinessType(v int32) *CreateCenterPolicyRequest
	GetBusinessType() *int32
	SetCameraRedirect(v string) *CreateCenterPolicyRequest
	GetCameraRedirect() *string
	SetClientControlMenu(v string) *CreateCenterPolicyRequest
	GetClientControlMenu() *string
	SetClientType(v []*CreateCenterPolicyRequestClientType) *CreateCenterPolicyRequest
	GetClientType() []*CreateCenterPolicyRequestClientType
	SetClipboard(v string) *CreateCenterPolicyRequest
	GetClipboard() *string
	SetClipboardGraineds(v []*CreateCenterPolicyRequestClipboardGraineds) *CreateCenterPolicyRequest
	GetClipboardGraineds() []*CreateCenterPolicyRequestClipboardGraineds
	SetClipboardScope(v string) *CreateCenterPolicyRequest
	GetClipboardScope() *string
	SetColorEnhancement(v string) *CreateCenterPolicyRequest
	GetColorEnhancement() *string
	SetCpdDriveClipboard(v string) *CreateCenterPolicyRequest
	GetCpdDriveClipboard() *string
	SetCpuDownGradeDuration(v int32) *CreateCenterPolicyRequest
	GetCpuDownGradeDuration() *int32
	SetCpuProcessors(v []*string) *CreateCenterPolicyRequest
	GetCpuProcessors() []*string
	SetCpuProtectedMode(v string) *CreateCenterPolicyRequest
	GetCpuProtectedMode() *string
	SetCpuRateLimit(v int32) *CreateCenterPolicyRequest
	GetCpuRateLimit() *int32
	SetCpuSampleDuration(v int32) *CreateCenterPolicyRequest
	GetCpuSampleDuration() *int32
	SetCpuSingleRateLimit(v int32) *CreateCenterPolicyRequest
	GetCpuSingleRateLimit() *int32
	SetDeviceConnectHint(v string) *CreateCenterPolicyRequest
	GetDeviceConnectHint() *string
	SetDeviceRedirects(v []*CreateCenterPolicyRequestDeviceRedirects) *CreateCenterPolicyRequest
	GetDeviceRedirects() []*CreateCenterPolicyRequestDeviceRedirects
	SetDeviceRules(v []*CreateCenterPolicyRequestDeviceRules) *CreateCenterPolicyRequest
	GetDeviceRules() []*CreateCenterPolicyRequestDeviceRules
	SetDisconnectKeepSession(v string) *CreateCenterPolicyRequest
	GetDisconnectKeepSession() *string
	SetDisconnectKeepSessionTime(v int32) *CreateCenterPolicyRequest
	GetDisconnectKeepSessionTime() *int32
	SetDisplayMode(v string) *CreateCenterPolicyRequest
	GetDisplayMode() *string
	SetDomainResolveRule(v []*CreateCenterPolicyRequestDomainResolveRule) *CreateCenterPolicyRequest
	GetDomainResolveRule() []*CreateCenterPolicyRequestDomainResolveRule
	SetDomainResolveRuleType(v string) *CreateCenterPolicyRequest
	GetDomainResolveRuleType() *string
	SetEnableSessionRateLimiting(v string) *CreateCenterPolicyRequest
	GetEnableSessionRateLimiting() *string
	SetEndUserApplyAdminCoordinate(v string) *CreateCenterPolicyRequest
	GetEndUserApplyAdminCoordinate() *string
	SetEndUserGroupCoordinate(v string) *CreateCenterPolicyRequest
	GetEndUserGroupCoordinate() *string
	SetFileMigrate(v string) *CreateCenterPolicyRequest
	GetFileMigrate() *string
	SetFileTransferAddress(v string) *CreateCenterPolicyRequest
	GetFileTransferAddress() *string
	SetFileTransferSpeed(v string) *CreateCenterPolicyRequest
	GetFileTransferSpeed() *string
	SetFileTransferSpeedLocation(v string) *CreateCenterPolicyRequest
	GetFileTransferSpeedLocation() *string
	SetGpuAcceleration(v string) *CreateCenterPolicyRequest
	GetGpuAcceleration() *string
	SetHtml5FileTransfer(v string) *CreateCenterPolicyRequest
	GetHtml5FileTransfer() *string
	SetInternetCommunicationProtocol(v string) *CreateCenterPolicyRequest
	GetInternetCommunicationProtocol() *string
	SetInternetPrinter(v string) *CreateCenterPolicyRequest
	GetInternetPrinter() *string
	SetLocalDrive(v string) *CreateCenterPolicyRequest
	GetLocalDrive() *string
	SetMaxReconnectTime(v int32) *CreateCenterPolicyRequest
	GetMaxReconnectTime() *int32
	SetMemoryDownGradeDuration(v int32) *CreateCenterPolicyRequest
	GetMemoryDownGradeDuration() *int32
	SetMemoryProcessors(v []*string) *CreateCenterPolicyRequest
	GetMemoryProcessors() []*string
	SetMemoryProtectedMode(v string) *CreateCenterPolicyRequest
	GetMemoryProtectedMode() *string
	SetMemoryRateLimit(v int32) *CreateCenterPolicyRequest
	GetMemoryRateLimit() *int32
	SetMemorySampleDuration(v int32) *CreateCenterPolicyRequest
	GetMemorySampleDuration() *int32
	SetMemorySingleRateLimit(v int32) *CreateCenterPolicyRequest
	GetMemorySingleRateLimit() *int32
	SetMobileRestart(v string) *CreateCenterPolicyRequest
	GetMobileRestart() *string
	SetMobileSafeMenu(v string) *CreateCenterPolicyRequest
	GetMobileSafeMenu() *string
	SetMobileShutdown(v string) *CreateCenterPolicyRequest
	GetMobileShutdown() *string
	SetMobileWuyingKeeper(v string) *CreateCenterPolicyRequest
	GetMobileWuyingKeeper() *string
	SetMobileWyAssistant(v string) *CreateCenterPolicyRequest
	GetMobileWyAssistant() *string
	SetName(v string) *CreateCenterPolicyRequest
	GetName() *string
	SetNetRedirect(v string) *CreateCenterPolicyRequest
	GetNetRedirect() *string
	SetNetRedirectRule(v []*CreateCenterPolicyRequestNetRedirectRule) *CreateCenterPolicyRequest
	GetNetRedirectRule() []*CreateCenterPolicyRequestNetRedirectRule
	SetNoOperationDisconnect(v string) *CreateCenterPolicyRequest
	GetNoOperationDisconnect() *string
	SetNoOperationDisconnectTime(v int32) *CreateCenterPolicyRequest
	GetNoOperationDisconnectTime() *int32
	SetPrinterRedirect(v string) *CreateCenterPolicyRequest
	GetPrinterRedirect() *string
	SetQualityEnhancement(v string) *CreateCenterPolicyRequest
	GetQualityEnhancement() *string
	SetRecordEventDuration(v int32) *CreateCenterPolicyRequest
	GetRecordEventDuration() *int32
	SetRecordEventFileExts(v []*string) *CreateCenterPolicyRequest
	GetRecordEventFileExts() []*string
	SetRecordEventFilePaths(v []*string) *CreateCenterPolicyRequest
	GetRecordEventFilePaths() []*string
	SetRecordEventLevels(v []*CreateCenterPolicyRequestRecordEventLevels) *CreateCenterPolicyRequest
	GetRecordEventLevels() []*CreateCenterPolicyRequestRecordEventLevels
	SetRecordEventRegisters(v []*string) *CreateCenterPolicyRequest
	GetRecordEventRegisters() []*string
	SetRecordEvents(v []*string) *CreateCenterPolicyRequest
	GetRecordEvents() []*string
	SetRecording(v string) *CreateCenterPolicyRequest
	GetRecording() *string
	SetRecordingAudio(v string) *CreateCenterPolicyRequest
	GetRecordingAudio() *string
	SetRecordingDuration(v int32) *CreateCenterPolicyRequest
	GetRecordingDuration() *int32
	SetRecordingEndTime(v string) *CreateCenterPolicyRequest
	GetRecordingEndTime() *string
	SetRecordingExpires(v int32) *CreateCenterPolicyRequest
	GetRecordingExpires() *int32
	SetRecordingFps(v string) *CreateCenterPolicyRequest
	GetRecordingFps() *string
	SetRecordingStartTime(v string) *CreateCenterPolicyRequest
	GetRecordingStartTime() *string
	SetRecordingUserNotify(v string) *CreateCenterPolicyRequest
	GetRecordingUserNotify() *string
	SetRecordingUserNotifyMessage(v string) *CreateCenterPolicyRequest
	GetRecordingUserNotifyMessage() *string
	SetRegionId(v string) *CreateCenterPolicyRequest
	GetRegionId() *string
	SetRemoteCoordinate(v string) *CreateCenterPolicyRequest
	GetRemoteCoordinate() *string
	SetResetDesktop(v string) *CreateCenterPolicyRequest
	GetResetDesktop() *string
	SetResolutionHeight(v int32) *CreateCenterPolicyRequest
	GetResolutionHeight() *int32
	SetResolutionModel(v string) *CreateCenterPolicyRequest
	GetResolutionModel() *string
	SetResolutionWidth(v int32) *CreateCenterPolicyRequest
	GetResolutionWidth() *int32
	SetResourceType(v string) *CreateCenterPolicyRequest
	GetResourceType() *string
	SetSafeMenu(v string) *CreateCenterPolicyRequest
	GetSafeMenu() *string
	SetScope(v string) *CreateCenterPolicyRequest
	GetScope() *string
	SetScopeValue(v []*string) *CreateCenterPolicyRequest
	GetScopeValue() []*string
	SetScreenDisplayMode(v string) *CreateCenterPolicyRequest
	GetScreenDisplayMode() *string
	SetSessionMaxRateKbps(v int32) *CreateCenterPolicyRequest
	GetSessionMaxRateKbps() *int32
	SetSmoothEnhancement(v string) *CreateCenterPolicyRequest
	GetSmoothEnhancement() *string
	SetStatusMonitor(v string) *CreateCenterPolicyRequest
	GetStatusMonitor() *string
	SetStreamingMode(v string) *CreateCenterPolicyRequest
	GetStreamingMode() *string
	SetTargetFps(v int32) *CreateCenterPolicyRequest
	GetTargetFps() *int32
	SetTaskbar(v string) *CreateCenterPolicyRequest
	GetTaskbar() *string
	SetUsbRedirect(v string) *CreateCenterPolicyRequest
	GetUsbRedirect() *string
	SetUsbSupplyRedirectRule(v []*CreateCenterPolicyRequestUsbSupplyRedirectRule) *CreateCenterPolicyRequest
	GetUsbSupplyRedirectRule() []*CreateCenterPolicyRequestUsbSupplyRedirectRule
	SetUseTime(v string) *CreateCenterPolicyRequest
	GetUseTime() *string
	SetVideoEncAvgKbps(v int32) *CreateCenterPolicyRequest
	GetVideoEncAvgKbps() *int32
	SetVideoEncMaxQP(v int32) *CreateCenterPolicyRequest
	GetVideoEncMaxQP() *int32
	SetVideoEncMinQP(v int32) *CreateCenterPolicyRequest
	GetVideoEncMinQP() *int32
	SetVideoEncPeakKbps(v int32) *CreateCenterPolicyRequest
	GetVideoEncPeakKbps() *int32
	SetVideoEncPolicy(v string) *CreateCenterPolicyRequest
	GetVideoEncPolicy() *string
	SetVideoRedirect(v string) *CreateCenterPolicyRequest
	GetVideoRedirect() *string
	SetVisualQuality(v string) *CreateCenterPolicyRequest
	GetVisualQuality() *string
	SetWatermark(v string) *CreateCenterPolicyRequest
	GetWatermark() *string
	SetWatermarkAntiCam(v string) *CreateCenterPolicyRequest
	GetWatermarkAntiCam() *string
	SetWatermarkColor(v int32) *CreateCenterPolicyRequest
	GetWatermarkColor() *int32
	SetWatermarkColumnAmount(v int32) *CreateCenterPolicyRequest
	GetWatermarkColumnAmount() *int32
	SetWatermarkCustomText(v string) *CreateCenterPolicyRequest
	GetWatermarkCustomText() *string
	SetWatermarkDegree(v float64) *CreateCenterPolicyRequest
	GetWatermarkDegree() *float64
	SetWatermarkFontSize(v int32) *CreateCenterPolicyRequest
	GetWatermarkFontSize() *int32
	SetWatermarkFontStyle(v string) *CreateCenterPolicyRequest
	GetWatermarkFontStyle() *string
	SetWatermarkPower(v string) *CreateCenterPolicyRequest
	GetWatermarkPower() *string
	SetWatermarkRowAmount(v int32) *CreateCenterPolicyRequest
	GetWatermarkRowAmount() *int32
	SetWatermarkSecurity(v string) *CreateCenterPolicyRequest
	GetWatermarkSecurity() *string
	SetWatermarkTransparencyValue(v int32) *CreateCenterPolicyRequest
	GetWatermarkTransparencyValue() *int32
	SetWatermarkType(v string) *CreateCenterPolicyRequest
	GetWatermarkType() *string
	SetWuyingKeeper(v string) *CreateCenterPolicyRequest
	GetWuyingKeeper() *string
	SetWyAssistant(v string) *CreateCenterPolicyRequest
	GetWyAssistant() *string
}

type CreateCenterPolicyRequest struct {
	// Specifies whether to grant the admin permissions to end users.
	//
	// >  This parameter is in private preview and only available to specific users.
	//
	// Valid values:
	//
	// 	- allow: forcibly grants admin permissions.
	//
	// 	- deny: forcibly rejects granting admin permissions.
	//
	// 	- inherited: inherits the admin permissions from the user dimension.
	//
	// example:
	//
	// deny
	AdminAccess *string `json:"AdminAccess,omitempty" xml:"AdminAccess,omitempty"`
	// The anti-screenshot policy.
	//
	// Valid values:
	//
	// 	- off (default): disables anti-screenshot.
	//
	// 	- on: enables anti-screenshot.
	//
	// example:
	//
	// off
	AppContentProtection *string `json:"AppContentProtection,omitempty" xml:"AppContentProtection,omitempty"`
	// The client IP address whitelists that you want to add. Once an IP address whitelist is configured, end users can only access cloud computers from the IP addresses listed in it.
	AuthorizeAccessPolicyRule []*CreateCenterPolicyRequestAuthorizeAccessPolicyRule `json:"AuthorizeAccessPolicyRule,omitempty" xml:"AuthorizeAccessPolicyRule,omitempty" type:"Repeated"`
	// The security group rule.
	AuthorizeSecurityPolicyRule []*CreateCenterPolicyRequestAuthorizeSecurityPolicyRule `json:"AuthorizeSecurityPolicyRule,omitempty" xml:"AuthorizeSecurityPolicyRule,omitempty" type:"Repeated"`
	// example:
	//
	// off
	AutoReconnect *string `json:"AutoReconnect,omitempty" xml:"AutoReconnect,omitempty"`
	// The business type.
	//
	// Valid values:
	//
	// 	- 1: public cloud
	//
	// 	- 8: commercial edition.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	BusinessType *int32 `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The on-premises camera redirection policy. This parameter only applies if DeviceRedirects does not include an on-premises camera redirection policy.
	//
	// Valid values:
	//
	// 	- deviceRedirect: enables device redirection.
	//
	// 	- off: disables device redirection.
	//
	// example:
	//
	// off
	CameraRedirect    *string `json:"CameraRedirect,omitempty" xml:"CameraRedirect,omitempty"`
	ClientControlMenu *string `json:"ClientControlMenu,omitempty" xml:"ClientControlMenu,omitempty"`
	// The types of Alibaba Cloud Workspace clients that end users can use to connect to cloud computers.
	ClientType []*CreateCenterPolicyRequestClientType `json:"ClientType,omitempty" xml:"ClientType,omitempty" type:"Repeated"`
	// The read/write permissions on the clipboard.
	//
	// Valid values:
	//
	// 	- read: specifies one-way transfer You can copy files only from on-premises devices to cloud computers.
	//
	// 	- readwrite: specifies two-way transfer. You can copy files between on-premises devices and cloud computers.
	//
	// 	- write: specifies one-way transfer. You can only copy files from cloud computers to on-premises devices.
	//
	// 	- off (default): disables all transfers, both one-way and two-way. Files cannot be copied directly between on-premises devices and cloud computers.
	//
	// example:
	//
	// off
	Clipboard         *string                                       `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	ClipboardGraineds []*CreateCenterPolicyRequestClipboardGraineds `json:"ClipboardGraineds,omitempty" xml:"ClipboardGraineds,omitempty" type:"Repeated"`
	// example:
	//
	// GLOBAL
	ClipboardScope *string `json:"ClipboardScope,omitempty" xml:"ClipboardScope,omitempty"`
	// Specifies whether to enable color enhancement for design and 3D applications.
	//
	// Valid values:
	//
	// 	- off: doesn\\"t enable color enhancement for design and 3D applications.
	//
	// 	- on: enables color enhancement for design and 3D applications.
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
	// Valid values:
	//
	// 	- off: disables CPU spike protection.
	//
	// 	- on: enables CPU spike protection.
	//
	// example:
	//
	// off
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
	// 60
	CpuSampleDuration *int32 `json:"CpuSampleDuration,omitempty" xml:"CpuSampleDuration,omitempty"`
	// The single-CPU usage. Valid values: 70 to 100. Unit: %.
	//
	// example:
	//
	// 70
	CpuSingleRateLimit *int32 `json:"CpuSingleRateLimit,omitempty" xml:"CpuSingleRateLimit,omitempty"`
	// Specifies whether to display the peripheral connection prompt.
	//
	// example:
	//
	// off
	DeviceConnectHint *string `json:"DeviceConnectHint,omitempty" xml:"DeviceConnectHint,omitempty"`
	// The device redirection rules.
	DeviceRedirects []*CreateCenterPolicyRequestDeviceRedirects `json:"DeviceRedirects,omitempty" xml:"DeviceRedirects,omitempty" type:"Repeated"`
	// The custom peripheral rules.
	DeviceRules []*CreateCenterPolicyRequestDeviceRules `json:"DeviceRules,omitempty" xml:"DeviceRules,omitempty" type:"Repeated"`
	// Specifies whether to retain the session upon disconnection.
	//
	// >  This parameter applies only to cloud application policies.
	//
	// Valid values:
	//
	// 	- customTime: retains the session for a specified time period.
	//
	// 	- persistent: retains the session permanently.
	//
	// example:
	//
	// customTime
	DisconnectKeepSession *string `json:"DisconnectKeepSession,omitempty" xml:"DisconnectKeepSession,omitempty"`
	// The retention period of the session after disconnection. Valid values: 30 to 7200. Unit: seconds.
	//
	// >  This parameter applies only to cloud application policies.
	//
	// example:
	//
	// 30
	DisconnectKeepSessionTime *int32 `json:"DisconnectKeepSessionTime,omitempty" xml:"DisconnectKeepSessionTime,omitempty"`
	// The display mode.
	//
	// Valid values:
	//
	// 	- clientCustom: suitable for user-defined scenarios.
	//
	// 	- adminOffice: suitable for daily office scenarios.
	//
	// 	- adminDesign: suitable for design and 3D application scenarios.
	//
	// 	- adminCustom: suitable for admin-customized scenarios.
	//
	// example:
	//
	// clientCustom
	DisplayMode *string `json:"DisplayMode,omitempty" xml:"DisplayMode,omitempty"`
	// The domain resolution policies.
	DomainResolveRule []*CreateCenterPolicyRequestDomainResolveRule `json:"DomainResolveRule,omitempty" xml:"DomainResolveRule,omitempty" type:"Repeated"`
	// Specifies whether to enforce the domain resolution policy.
	//
	// Valid values:
	//
	// 	- off: disables the domain resolution policy.
	//
	// 	- on: enables the domain resolution policy.
	//
	// example:
	//
	// off
	DomainResolveRuleType *string `json:"DomainResolveRuleType,omitempty" xml:"DomainResolveRuleType,omitempty"`
	// Specifies whether to enforce the peak bandwidth limit for sessions.
	//
	// Valid values:
	//
	// 	- off: doesn\\"t enforce the peak bandwidth limit for sessions.
	//
	// 	- on: enforces the peak bandwidth limit for sessions.
	//
	// example:
	//
	// off
	EnableSessionRateLimiting *string `json:"EnableSessionRateLimiting,omitempty" xml:"EnableSessionRateLimiting,omitempty"`
	// Specifies whether to enable end users to request administrator help.
	//
	// Valid values:
	//
	// 	- off: disables end users to request administrator help.
	//
	// 	- on: enables end users to request administrator help.
	//
	// example:
	//
	// off
	EndUserApplyAdminCoordinate *string `json:"EndUserApplyAdminCoordinate,omitempty" xml:"EndUserApplyAdminCoordinate,omitempty"`
	// Specifies whether to allow end users in the same office network to share cloud computers.
	//
	// Valid values:
	//
	// 	- off: doesn\\"t allow end users in the same office network to share cloud computers.
	//
	// 	- on: allows end users in the same office network to share cloud computers.
	//
	// example:
	//
	// off
	EndUserGroupCoordinate *string `json:"EndUserGroupCoordinate,omitempty" xml:"EndUserGroupCoordinate,omitempty"`
	// Specifies whether to enable file transfer.
	//
	// example:
	//
	// off
	FileMigrate               *string `json:"FileMigrate,omitempty" xml:"FileMigrate,omitempty"`
	FileTransferAddress       *string `json:"FileTransferAddress,omitempty" xml:"FileTransferAddress,omitempty"`
	FileTransferSpeed         *string `json:"FileTransferSpeed,omitempty" xml:"FileTransferSpeed,omitempty"`
	FileTransferSpeedLocation *string `json:"FileTransferSpeedLocation,omitempty" xml:"FileTransferSpeedLocation,omitempty"`
	// Specifies whether to enable image quality control. This feature is highly recommended for professional design scenarios where computer performance and user experience are critical.
	//
	// Valid values:
	//
	// 	- off: doesn\\"t enable image quality control.
	//
	// 	- on: enables image quality control.
	//
	// example:
	//
	// off
	GpuAcceleration *string `json:"GpuAcceleration,omitempty" xml:"GpuAcceleration,omitempty"`
	// The file transfer feature on the web client.
	//
	// Valid values:
	//
	// 	- all: File upload and download are supported.
	//
	// 	- download: Only file download is supported.
	//
	// 	- upload: Only file upload is supported.
	//
	// 	- off (default): File upload and download are not supported.
	//
	// example:
	//
	// off
	Html5FileTransfer *string `json:"Html5FileTransfer,omitempty" xml:"Html5FileTransfer,omitempty"`
	// The protocol for network communication.
	//
	// Valid values:
	//
	// 	- tcp: TCP is used when UDP/AST is restricted.
	//
	// 	- rtc: AST is used for high-frequency audio and video streaming.
	//
	// 	- auto: UTO is used to enable automatic switch between AST and UDP modes based on desktop content.
	//
	// 	- both: UDP is used for office and HD graphic design use.
	//
	// example:
	//
	// both
	InternetCommunicationProtocol *string `json:"InternetCommunicationProtocol,omitempty" xml:"InternetCommunicationProtocol,omitempty"`
	InternetPrinter               *string `json:"InternetPrinter,omitempty" xml:"InternetPrinter,omitempty"`
	// The read/write permissions on the on-premises drive.
	//
	// Valid values:
	//
	// 	- read: read-only. Cloud computers support on-premises disk mapping, but only for reading (copying) files—not modifying them.
	//
	// 	- readwrite: read and write. Cloud computers support on-premises disk mapping, allowing you to read (copy) and write (modify) on-premises files.
	//
	// 	- off (default): none. Cloud computers don\\"t support on-premises disk mapping.
	//
	// example:
	//
	// off
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
	// 40
	MemoryDownGradeDuration *int32 `json:"MemoryDownGradeDuration,omitempty" xml:"MemoryDownGradeDuration,omitempty"`
	// The memory processors.
	MemoryProcessors []*string `json:"MemoryProcessors,omitempty" xml:"MemoryProcessors,omitempty" type:"Repeated"`
	// The memory spike protection policy.
	//
	// Valid values:
	//
	// 	- off: disables memory spike protection.
	//
	// 	- on: enables memory spike protection.
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
	// 40
	MemorySampleDuration *int32 `json:"MemorySampleDuration,omitempty" xml:"MemorySampleDuration,omitempty"`
	// The memory usage per process. Valid values: 30 to 60. Unit: %.
	//
	// example:
	//
	// 40
	MemorySingleRateLimit *int32 `json:"MemorySingleRateLimit,omitempty" xml:"MemorySingleRateLimit,omitempty"`
	// Specifies whether to display the Restart button in the DesktopAssistant menu when end users connect to cloud computers from Android clients.
	//
	// >  This feature applies to only mobile clients of version 7.4.0 or later.
	//
	// Valid values:
	//
	// 	- off: doesn\\"t display the Restart button in the DesktopAssistant menu.
	//
	// 	- on: displays the Restart button in the DesktopAssistant menu.
	//
	// example:
	//
	// off
	MobileRestart *string `json:"MobileRestart,omitempty" xml:"MobileRestart,omitempty"`
	// example:
	//
	// off
	MobileSafeMenu *string `json:"MobileSafeMenu,omitempty" xml:"MobileSafeMenu,omitempty"`
	// Specifies whether to display the Stop button in the DesktopAssistant menu when end users connect to cloud computers from Android clients.
	//
	// >  This feature applies to only mobile clients of version 7.4.0 or later.
	//
	// Valid values:
	//
	// 	- off: doesn\\"t display the Stop button in the DesktopAssistant menu.
	//
	// 	- on: displays the Stop button in the DesktopAssistant menu.
	//
	// example:
	//
	// off
	MobileShutdown *string `json:"MobileShutdown,omitempty" xml:"MobileShutdown,omitempty"`
	// example:
	//
	// off
	MobileWuyingKeeper *string `json:"MobileWuyingKeeper,omitempty" xml:"MobileWuyingKeeper,omitempty"`
	// example:
	//
	// off
	MobileWyAssistant *string `json:"MobileWyAssistant,omitempty" xml:"MobileWyAssistant,omitempty"`
	// The policy name.
	//
	// This parameter is required.
	//
	// example:
	//
	// testPolicyGroupName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The network redirection policy.
	//
	// >  This parameter is in private preview and only available to specific users.
	//
	// Valid values:
	//
	// 	- all: enables network redirection globally.
	//
	// 	- off (default): disables network redirection.
	//
	// 	- on: enables the whitelist mode.
	//
	// example:
	//
	// off
	NetRedirect *string `json:"NetRedirect,omitempty" xml:"NetRedirect,omitempty"`
	// The network redirection policy.
	//
	// >  This parameter is in private preview and only available to specific users.
	NetRedirectRule []*CreateCenterPolicyRequestNetRedirectRule `json:"NetRedirectRule,omitempty" xml:"NetRedirectRule,omitempty" type:"Repeated"`
	// Specifies whether to enforce a disconnection upon inactivity.
	//
	// >  This parameter applies only to cloud application policies.
	//
	// Valid values:
	//
	// 	- off: doesn\\"t enforce a disconnection upon inactivity.
	//
	// 	- on: enforces a disconnection upon inactivity.
	//
	// example:
	//
	// off
	NoOperationDisconnect *string `json:"NoOperationDisconnect,omitempty" xml:"NoOperationDisconnect,omitempty"`
	// The duration of disconnection after inactivity. Valid values: 120 to 7200. Unit: seconds.
	//
	// >  This parameter applies only to cloud application policies.
	//
	// example:
	//
	// 120
	NoOperationDisconnectTime *int32 `json:"NoOperationDisconnectTime,omitempty" xml:"NoOperationDisconnectTime,omitempty"`
	// The printer redirection policy. This parameter only applies if DeviceRedirects does not include a printer redirection policy.
	//
	// Valid values:
	//
	// 	- deviceRedirect (default):enables device redirection.
	//
	// 	- usbRedirect: enables USB redirection.
	//
	// 	- off: disables any type of redirection.
	//
	// example:
	//
	// off
	PrinterRedirect *string `json:"PrinterRedirect,omitempty" xml:"PrinterRedirect,omitempty"`
	// Specifies whether to enable image quality enhancement for design and 3D applications.
	//
	// Valid values:
	//
	// 	- off: doesn\\"t enable image quality enhancement for design and 3D applications.
	//
	// 	- on: enables image quality enhancement for design and 3D applications.
	//
	// example:
	//
	// off
	QualityEnhancement *string `json:"QualityEnhancement,omitempty" xml:"QualityEnhancement,omitempty"`
	// The duration of screen recording after the specified event is detected. Unit: minutes. Valid values: 10 to 60.
	//
	// example:
	//
	// 10
	RecordEventDuration *int32    `json:"RecordEventDuration,omitempty" xml:"RecordEventDuration,omitempty"`
	RecordEventFileExts []*string `json:"RecordEventFileExts,omitempty" xml:"RecordEventFileExts,omitempty" type:"Repeated"`
	// The absolute paths to screen recording files.
	RecordEventFilePaths []*string                                     `json:"RecordEventFilePaths,omitempty" xml:"RecordEventFilePaths,omitempty" type:"Repeated"`
	RecordEventLevels    []*CreateCenterPolicyRequestRecordEventLevels `json:"RecordEventLevels,omitempty" xml:"RecordEventLevels,omitempty" type:"Repeated"`
	// The absolute paths to screen recording registries.
	RecordEventRegisters []*string `json:"RecordEventRegisters,omitempty" xml:"RecordEventRegisters,omitempty" type:"Repeated"`
	// The events that trigger screen recording.
	RecordEvents []*string `json:"RecordEvents,omitempty" xml:"RecordEvents,omitempty" type:"Repeated"`
	// The screen recording policy.
	//
	// Valid values:
	//
	// 	- period: Screen recording occurs at set intervals.
	//
	// 	- session: Screen recording is limited to sessions only.
	//
	// 	- off: Screen recording is disabled.
	//
	// 	- alltime: Screen recording is always enabled.
	//
	// example:
	//
	// off
	Recording *string `json:"Recording,omitempty" xml:"Recording,omitempty"`
	// Specifies whether to record audio files generated on cloud computers.
	//
	// Valid values:
	//
	// 	- off: doesn\\"t record audio files generated on cloud computers.
	//
	// 	- on: records audio files generated on cloud computers.
	//
	// example:
	//
	// on
	RecordingAudio *string `json:"RecordingAudio,omitempty" xml:"RecordingAudio,omitempty"`
	// The frame rate of screen recording. Screen recordings are split based on the specified duration and uploaded to Object Storage Service (OSS) buckets. If a file reaches 300 MB, the system prioritizes rolling updates for that file. Valid values: 10 to 60
	//
	// example:
	//
	// 10
	RecordingDuration *int32 `json:"RecordingDuration,omitempty" xml:"RecordingDuration,omitempty"`
	// The screen recording\\"s end time in HH:MM:SS format. The value is meaningful only if `Recording` is set to `PERIOD`.
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
	RecordingExpires *int32 `json:"RecordingExpires,omitempty" xml:"RecordingExpires,omitempty"`
	// The frame rate of screen recording. Unit: fps.
	//
	// example:
	//
	// 2
	RecordingFps *string `json:"RecordingFps,omitempty" xml:"RecordingFps,omitempty"`
	// The screen recording\\"s start time in HH:MM:SS format. The value is meaningful only if `Recording` is set to `PERIOD`.
	//
	// example:
	//
	// 08:00:00
	RecordingStartTime *string `json:"RecordingStartTime,omitempty" xml:"RecordingStartTime,omitempty"`
	// Specifies whether to notify end users when screen recording is enabled.
	//
	// Valid values:
	//
	// 	- off: doesn\\"t notify end users when screen recording is enabled.
	//
	// 	- on: notifies end users when screen recording is enabled.
	//
	// example:
	//
	// off
	RecordingUserNotify *string `json:"RecordingUserNotify,omitempty" xml:"RecordingUserNotify,omitempty"`
	// The notification sent to end users when screen recording is enabled.
	RecordingUserNotifyMessage *string `json:"RecordingUserNotifyMessage,omitempty" xml:"RecordingUserNotifyMessage,omitempty"`
	// The region ID. Set the value to cn-shanghai.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The keyboard and mouse control permissions during remote assistance.
	//
	// Valid values:
	//
	// 	- optionalControl: By default, keyboard and mouse control is disabled during remote assistance. You can request permissions as needed.
	//
	// 	- fullControl: Keyboard and mouse control is enabled during remote assistance.
	//
	// 	- disableControl: Keyboard and mouse control is disabled during remote assistance.
	//
	// example:
	//
	// fullControl
	RemoteCoordinate *string `json:"RemoteCoordinate,omitempty" xml:"RemoteCoordinate,omitempty"`
	// The computer reset setting.
	//
	// example:
	//
	// off
	ResetDesktop *string `json:"ResetDesktop,omitempty" xml:"ResetDesktop,omitempty"`
	// The resolution height. Unit: pixel. Valid values for cloud applications: 500 to 50000. Valid values for cloud computers: 480 to 4096.
	//
	// example:
	//
	// 1280
	ResolutionHeight *int32 `json:"ResolutionHeight,omitempty" xml:"ResolutionHeight,omitempty"`
	// The resolution type.
	//
	// Valid values:
	//
	// 	- adaptive: adaptive resolution.
	//
	// 	- customer: fixed resolution.
	//
	// example:
	//
	// adaptive
	ResolutionModel *string `json:"ResolutionModel,omitempty" xml:"ResolutionModel,omitempty"`
	// The resolution width. Unit: pixel. Valid values for cloud applications: 500 to 50000. Valid values for cloud computers: 480 to 4096.
	//
	// example:
	//
	// 720
	ResolutionWidth *int32 `json:"ResolutionWidth,omitempty" xml:"ResolutionWidth,omitempty"`
	// The resource type.
	//
	// Valid values:
	//
	// 	- app: cloud applications.
	//
	// 	- desktop: cloud computers.
	//
	// This parameter is required.
	//
	// example:
	//
	// desktop
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SafeMenu     *string `json:"SafeMenu,omitempty" xml:"SafeMenu,omitempty"`
	// The effective scope of the policy.
	//
	// Valid values:
	//
	// 	- IP: The policy applies to specific IP addresses.
	//
	// 	- GLOBAL: The policy applies globally.
	//
	// example:
	//
	// GLOBAL
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The effective scopes. This parameter is required when `Scope` is set to `IP`. If `Scope` is set to `IP`, this parameter doesn\\"t take effect.
	ScopeValue        []*string `json:"ScopeValue,omitempty" xml:"ScopeValue,omitempty" type:"Repeated"`
	ScreenDisplayMode *string   `json:"ScreenDisplayMode,omitempty" xml:"ScreenDisplayMode,omitempty"`
	// The bandwidth peak allowed for sessions. Unit: Kbit/s. Valid values: 2000 to 100000.
	//
	// example:
	//
	// 2000
	SessionMaxRateKbps *int32 `json:"SessionMaxRateKbps,omitempty" xml:"SessionMaxRateKbps,omitempty"`
	// Specifies whether to enable smoothness enhancement for daily office use.
	//
	// Valid values:
	//
	// 	- off: doesn\\"t enable smoothness enhancement for daily office use.
	//
	// 	- on: enables smoothness enhancement for daily office use.
	//
	// example:
	//
	// off
	SmoothEnhancement *string `json:"SmoothEnhancement,omitempty" xml:"SmoothEnhancement,omitempty"`
	// Specifies whether to display the metric status entry in the DesktopAssistant menu.
	//
	// Valid values:
	//
	// 	- off: doesn\\"t display the metric status entry in the DesktopAssistant menu.
	//
	// 	- on: displays the metric status entry in the DesktopAssistant menu.
	//
	// example:
	//
	// off
	StatusMonitor *string `json:"StatusMonitor,omitempty" xml:"StatusMonitor,omitempty"`
	// The streaming mode.
	//
	// Valid values:
	//
	// 	- intelligent
	//
	// 	- smooth
	//
	// example:
	//
	// smooth
	StreamingMode *string `json:"StreamingMode,omitempty" xml:"StreamingMode,omitempty"`
	// The target frame rate. Valid values: 10 to 60.
	//
	// example:
	//
	// 30
	TargetFps *int32 `json:"TargetFps,omitempty" xml:"TargetFps,omitempty"`
	// Specifies whether to display the application taskbar.
	//
	// >  This parameter applies only to cloud application policies.
	//
	// Valid values:
	//
	// 	- off: doesn\\"t display the application taskbar.
	//
	// 	- on: displays the application taskbar.
	//
	// example:
	//
	// off
	Taskbar *string `json:"Taskbar,omitempty" xml:"Taskbar,omitempty"`
	// Specifies whether to enable USB redirection.
	//
	// Valid values:
	//
	// 	- off (default): doesn\\"t enable USB redirection.
	//
	// 	- on: enables USB redirection.
	//
	// example:
	//
	// off
	UsbRedirect *string `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty"`
	// The USB redirection rules.
	UsbSupplyRedirectRule []*CreateCenterPolicyRequestUsbSupplyRedirectRule `json:"UsbSupplyRedirectRule,omitempty" xml:"UsbSupplyRedirectRule,omitempty" type:"Repeated"`
	UseTime               *string                                           `json:"UseTime,omitempty" xml:"UseTime,omitempty"`
	// The average bitrate for video encoding. Unit: Kbit/s. Valid values: 1000 to 50000.
	//
	// example:
	//
	// 2000
	VideoEncAvgKbps *int32 `json:"VideoEncAvgKbps,omitempty" xml:"VideoEncAvgKbps,omitempty"`
	// The maximum QP for video files. Higher QP values result in lower video quality. Valid values: 0 to 51.
	//
	// example:
	//
	// 30
	VideoEncMaxQP *int32 `json:"VideoEncMaxQP,omitempty" xml:"VideoEncMaxQP,omitempty"`
	// The minimum quantizer parameter (QP) for video files. A lower QP means better video quality. Valid values: 0 to 51.
	//
	// example:
	//
	// 30
	VideoEncMinQP *int32 `json:"VideoEncMinQP,omitempty" xml:"VideoEncMinQP,omitempty"`
	// The peak bitrate for video encoding. Unit: Kbit/s. Valid values: 1000 to 50000.
	//
	// example:
	//
	// 2000
	VideoEncPeakKbps *int32 `json:"VideoEncPeakKbps,omitempty" xml:"VideoEncPeakKbps,omitempty"`
	// The video encoding policy.
	//
	// Valid values:
	//
	// 	- qualityFirst: prioritizes image quality.
	//
	// 	- bandwidthFirst: prioritizes bandwidth.
	//
	// example:
	//
	// qualityFirst
	VideoEncPolicy *string `json:"VideoEncPolicy,omitempty" xml:"VideoEncPolicy,omitempty"`
	// The multimedia redirection policy.
	//
	// Valid values:
	//
	// 	- off: disables multimedia redirection.
	//
	// 	- on: enables multimedia redirection.
	//
	// example:
	//
	// on
	VideoRedirect *string `json:"VideoRedirect,omitempty" xml:"VideoRedirect,omitempty"`
	// The image display quality.
	//
	// Valid values:
	//
	// 	- high: high-definition (HD).
	//
	// 	- low: smoothness.
	//
	// 	- lossless: no quality loss.
	//
	// 	- medium (default): scenario-specific adaptation.
	//
	// example:
	//
	// low
	VisualQuality *string `json:"VisualQuality,omitempty" xml:"VisualQuality,omitempty"`
	// The watermark policy.
	//
	// Valid values:
	//
	// 	- blind: displays invisible watermarks.
	//
	// 	- off (default): displays no watermark.
	//
	// 	- on: displays visible watermarks.
	//
	// example:
	//
	// off
	Watermark *string `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
	// Specifies whether to enable anti-screen capture for invisible watermarks.
	//
	// Valid values:
	//
	// 	- off: disables anti-screen capture for invisible watermarks.
	//
	// 	- on: enables anti-screen capture for invisible watermarks.
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
	// The number of watermark columns. Valid values: 3 to 10.
	//
	// example:
	//
	// 3
	WatermarkColumnAmount *int32 `json:"WatermarkColumnAmount,omitempty" xml:"WatermarkColumnAmount,omitempty"`
	// If you set `WatermarkType` to `custom`, you must also specify `WatermarkCustomText`.
	//
	// example:
	//
	// test
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
	// The enhancement level for invisible watermarks.
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
	// The number of watermark rows. Valid values: 3 to 10.
	//
	// example:
	//
	// 3
	WatermarkRowAmount *int32 `json:"WatermarkRowAmount,omitempty" xml:"WatermarkRowAmount,omitempty"`
	// Specifies whether to enable security priority for invisible watermarks.
	//
	// Valid values:
	//
	// 	- off: disables security priority for invisible watermarks.
	//
	// 	- on: enables security priority for invisible watermarks.
	//
	// example:
	//
	// on
	WatermarkSecurity *string `json:"WatermarkSecurity,omitempty" xml:"WatermarkSecurity,omitempty"`
	// The watermark opacity. A higher value makes the watermark more opaque. Valid values: 10 to 100.
	//
	// example:
	//
	// 10
	WatermarkTransparencyValue *int32 `json:"WatermarkTransparencyValue,omitempty" xml:"WatermarkTransparencyValue,omitempty"`
	// The watermark type. You can specify up to three types. Separate multiple values with commas (,).
	//
	// >  If you provide `custom` as the value for this parameter, you must configure `WatermarkCustomText` to specify custom text.
	//
	// Valid values:
	//
	// 	- EndUserId: the username.
	//
	// 	- Custom: the custom text.
	//
	// 	- DesktopIp: the IP address of the cloud computer.
	//
	// 	- ClientIp: the IP address of the client.
	//
	// 	- HostName: the rightmost 15 digits of the cloud computer ID.
	//
	// 	- ClientTime: the current time displayed on the cloud computer.
	//
	// example:
	//
	// EndUserId,HostName,ClientTime
	WatermarkType *string `json:"WatermarkType,omitempty" xml:"WatermarkType,omitempty"`
	// Specifies whether to enable Cloud Computer Manager.
	//
	// example:
	//
	// off
	WuyingKeeper *string `json:"WuyingKeeper,omitempty" xml:"WuyingKeeper,omitempty"`
	// Specifies whether to display the Xiaoying AI Assistant option in the DesktopAssistant menu when end users connect to cloud computers via desktop clients (Windows and macOS).
	//
	// >  This feature applies to only desktop clients of version 7.7.0 or later.
	//
	// Valid values:
	//
	// 	- off: doesn\\"t display the Xiaoying AI Assistant option in the DesktopAssistant menu.
	//
	// 	- on: displays the Xiaoying AI Assistant option in the DesktopAssistant menu.
	//
	// example:
	//
	// on
	WyAssistant *string `json:"WyAssistant,omitempty" xml:"WyAssistant,omitempty"`
}

func (s CreateCenterPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCenterPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateCenterPolicyRequest) GetAdminAccess() *string {
	return s.AdminAccess
}

func (s *CreateCenterPolicyRequest) GetAppContentProtection() *string {
	return s.AppContentProtection
}

func (s *CreateCenterPolicyRequest) GetAuthorizeAccessPolicyRule() []*CreateCenterPolicyRequestAuthorizeAccessPolicyRule {
	return s.AuthorizeAccessPolicyRule
}

func (s *CreateCenterPolicyRequest) GetAuthorizeSecurityPolicyRule() []*CreateCenterPolicyRequestAuthorizeSecurityPolicyRule {
	return s.AuthorizeSecurityPolicyRule
}

func (s *CreateCenterPolicyRequest) GetAutoReconnect() *string {
	return s.AutoReconnect
}

func (s *CreateCenterPolicyRequest) GetBusinessType() *int32 {
	return s.BusinessType
}

func (s *CreateCenterPolicyRequest) GetCameraRedirect() *string {
	return s.CameraRedirect
}

func (s *CreateCenterPolicyRequest) GetClientControlMenu() *string {
	return s.ClientControlMenu
}

func (s *CreateCenterPolicyRequest) GetClientType() []*CreateCenterPolicyRequestClientType {
	return s.ClientType
}

func (s *CreateCenterPolicyRequest) GetClipboard() *string {
	return s.Clipboard
}

func (s *CreateCenterPolicyRequest) GetClipboardGraineds() []*CreateCenterPolicyRequestClipboardGraineds {
	return s.ClipboardGraineds
}

func (s *CreateCenterPolicyRequest) GetClipboardScope() *string {
	return s.ClipboardScope
}

func (s *CreateCenterPolicyRequest) GetColorEnhancement() *string {
	return s.ColorEnhancement
}

func (s *CreateCenterPolicyRequest) GetCpdDriveClipboard() *string {
	return s.CpdDriveClipboard
}

func (s *CreateCenterPolicyRequest) GetCpuDownGradeDuration() *int32 {
	return s.CpuDownGradeDuration
}

func (s *CreateCenterPolicyRequest) GetCpuProcessors() []*string {
	return s.CpuProcessors
}

func (s *CreateCenterPolicyRequest) GetCpuProtectedMode() *string {
	return s.CpuProtectedMode
}

func (s *CreateCenterPolicyRequest) GetCpuRateLimit() *int32 {
	return s.CpuRateLimit
}

func (s *CreateCenterPolicyRequest) GetCpuSampleDuration() *int32 {
	return s.CpuSampleDuration
}

func (s *CreateCenterPolicyRequest) GetCpuSingleRateLimit() *int32 {
	return s.CpuSingleRateLimit
}

func (s *CreateCenterPolicyRequest) GetDeviceConnectHint() *string {
	return s.DeviceConnectHint
}

func (s *CreateCenterPolicyRequest) GetDeviceRedirects() []*CreateCenterPolicyRequestDeviceRedirects {
	return s.DeviceRedirects
}

func (s *CreateCenterPolicyRequest) GetDeviceRules() []*CreateCenterPolicyRequestDeviceRules {
	return s.DeviceRules
}

func (s *CreateCenterPolicyRequest) GetDisconnectKeepSession() *string {
	return s.DisconnectKeepSession
}

func (s *CreateCenterPolicyRequest) GetDisconnectKeepSessionTime() *int32 {
	return s.DisconnectKeepSessionTime
}

func (s *CreateCenterPolicyRequest) GetDisplayMode() *string {
	return s.DisplayMode
}

func (s *CreateCenterPolicyRequest) GetDomainResolveRule() []*CreateCenterPolicyRequestDomainResolveRule {
	return s.DomainResolveRule
}

func (s *CreateCenterPolicyRequest) GetDomainResolveRuleType() *string {
	return s.DomainResolveRuleType
}

func (s *CreateCenterPolicyRequest) GetEnableSessionRateLimiting() *string {
	return s.EnableSessionRateLimiting
}

func (s *CreateCenterPolicyRequest) GetEndUserApplyAdminCoordinate() *string {
	return s.EndUserApplyAdminCoordinate
}

func (s *CreateCenterPolicyRequest) GetEndUserGroupCoordinate() *string {
	return s.EndUserGroupCoordinate
}

func (s *CreateCenterPolicyRequest) GetFileMigrate() *string {
	return s.FileMigrate
}

func (s *CreateCenterPolicyRequest) GetFileTransferAddress() *string {
	return s.FileTransferAddress
}

func (s *CreateCenterPolicyRequest) GetFileTransferSpeed() *string {
	return s.FileTransferSpeed
}

func (s *CreateCenterPolicyRequest) GetFileTransferSpeedLocation() *string {
	return s.FileTransferSpeedLocation
}

func (s *CreateCenterPolicyRequest) GetGpuAcceleration() *string {
	return s.GpuAcceleration
}

func (s *CreateCenterPolicyRequest) GetHtml5FileTransfer() *string {
	return s.Html5FileTransfer
}

func (s *CreateCenterPolicyRequest) GetInternetCommunicationProtocol() *string {
	return s.InternetCommunicationProtocol
}

func (s *CreateCenterPolicyRequest) GetInternetPrinter() *string {
	return s.InternetPrinter
}

func (s *CreateCenterPolicyRequest) GetLocalDrive() *string {
	return s.LocalDrive
}

func (s *CreateCenterPolicyRequest) GetMaxReconnectTime() *int32 {
	return s.MaxReconnectTime
}

func (s *CreateCenterPolicyRequest) GetMemoryDownGradeDuration() *int32 {
	return s.MemoryDownGradeDuration
}

func (s *CreateCenterPolicyRequest) GetMemoryProcessors() []*string {
	return s.MemoryProcessors
}

func (s *CreateCenterPolicyRequest) GetMemoryProtectedMode() *string {
	return s.MemoryProtectedMode
}

func (s *CreateCenterPolicyRequest) GetMemoryRateLimit() *int32 {
	return s.MemoryRateLimit
}

func (s *CreateCenterPolicyRequest) GetMemorySampleDuration() *int32 {
	return s.MemorySampleDuration
}

func (s *CreateCenterPolicyRequest) GetMemorySingleRateLimit() *int32 {
	return s.MemorySingleRateLimit
}

func (s *CreateCenterPolicyRequest) GetMobileRestart() *string {
	return s.MobileRestart
}

func (s *CreateCenterPolicyRequest) GetMobileSafeMenu() *string {
	return s.MobileSafeMenu
}

func (s *CreateCenterPolicyRequest) GetMobileShutdown() *string {
	return s.MobileShutdown
}

func (s *CreateCenterPolicyRequest) GetMobileWuyingKeeper() *string {
	return s.MobileWuyingKeeper
}

func (s *CreateCenterPolicyRequest) GetMobileWyAssistant() *string {
	return s.MobileWyAssistant
}

func (s *CreateCenterPolicyRequest) GetName() *string {
	return s.Name
}

func (s *CreateCenterPolicyRequest) GetNetRedirect() *string {
	return s.NetRedirect
}

func (s *CreateCenterPolicyRequest) GetNetRedirectRule() []*CreateCenterPolicyRequestNetRedirectRule {
	return s.NetRedirectRule
}

func (s *CreateCenterPolicyRequest) GetNoOperationDisconnect() *string {
	return s.NoOperationDisconnect
}

func (s *CreateCenterPolicyRequest) GetNoOperationDisconnectTime() *int32 {
	return s.NoOperationDisconnectTime
}

func (s *CreateCenterPolicyRequest) GetPrinterRedirect() *string {
	return s.PrinterRedirect
}

func (s *CreateCenterPolicyRequest) GetQualityEnhancement() *string {
	return s.QualityEnhancement
}

func (s *CreateCenterPolicyRequest) GetRecordEventDuration() *int32 {
	return s.RecordEventDuration
}

func (s *CreateCenterPolicyRequest) GetRecordEventFileExts() []*string {
	return s.RecordEventFileExts
}

func (s *CreateCenterPolicyRequest) GetRecordEventFilePaths() []*string {
	return s.RecordEventFilePaths
}

func (s *CreateCenterPolicyRequest) GetRecordEventLevels() []*CreateCenterPolicyRequestRecordEventLevels {
	return s.RecordEventLevels
}

func (s *CreateCenterPolicyRequest) GetRecordEventRegisters() []*string {
	return s.RecordEventRegisters
}

func (s *CreateCenterPolicyRequest) GetRecordEvents() []*string {
	return s.RecordEvents
}

func (s *CreateCenterPolicyRequest) GetRecording() *string {
	return s.Recording
}

func (s *CreateCenterPolicyRequest) GetRecordingAudio() *string {
	return s.RecordingAudio
}

func (s *CreateCenterPolicyRequest) GetRecordingDuration() *int32 {
	return s.RecordingDuration
}

func (s *CreateCenterPolicyRequest) GetRecordingEndTime() *string {
	return s.RecordingEndTime
}

func (s *CreateCenterPolicyRequest) GetRecordingExpires() *int32 {
	return s.RecordingExpires
}

func (s *CreateCenterPolicyRequest) GetRecordingFps() *string {
	return s.RecordingFps
}

func (s *CreateCenterPolicyRequest) GetRecordingStartTime() *string {
	return s.RecordingStartTime
}

func (s *CreateCenterPolicyRequest) GetRecordingUserNotify() *string {
	return s.RecordingUserNotify
}

func (s *CreateCenterPolicyRequest) GetRecordingUserNotifyMessage() *string {
	return s.RecordingUserNotifyMessage
}

func (s *CreateCenterPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCenterPolicyRequest) GetRemoteCoordinate() *string {
	return s.RemoteCoordinate
}

func (s *CreateCenterPolicyRequest) GetResetDesktop() *string {
	return s.ResetDesktop
}

func (s *CreateCenterPolicyRequest) GetResolutionHeight() *int32 {
	return s.ResolutionHeight
}

func (s *CreateCenterPolicyRequest) GetResolutionModel() *string {
	return s.ResolutionModel
}

func (s *CreateCenterPolicyRequest) GetResolutionWidth() *int32 {
	return s.ResolutionWidth
}

func (s *CreateCenterPolicyRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateCenterPolicyRequest) GetSafeMenu() *string {
	return s.SafeMenu
}

func (s *CreateCenterPolicyRequest) GetScope() *string {
	return s.Scope
}

func (s *CreateCenterPolicyRequest) GetScopeValue() []*string {
	return s.ScopeValue
}

func (s *CreateCenterPolicyRequest) GetScreenDisplayMode() *string {
	return s.ScreenDisplayMode
}

func (s *CreateCenterPolicyRequest) GetSessionMaxRateKbps() *int32 {
	return s.SessionMaxRateKbps
}

func (s *CreateCenterPolicyRequest) GetSmoothEnhancement() *string {
	return s.SmoothEnhancement
}

func (s *CreateCenterPolicyRequest) GetStatusMonitor() *string {
	return s.StatusMonitor
}

func (s *CreateCenterPolicyRequest) GetStreamingMode() *string {
	return s.StreamingMode
}

func (s *CreateCenterPolicyRequest) GetTargetFps() *int32 {
	return s.TargetFps
}

func (s *CreateCenterPolicyRequest) GetTaskbar() *string {
	return s.Taskbar
}

func (s *CreateCenterPolicyRequest) GetUsbRedirect() *string {
	return s.UsbRedirect
}

func (s *CreateCenterPolicyRequest) GetUsbSupplyRedirectRule() []*CreateCenterPolicyRequestUsbSupplyRedirectRule {
	return s.UsbSupplyRedirectRule
}

func (s *CreateCenterPolicyRequest) GetUseTime() *string {
	return s.UseTime
}

func (s *CreateCenterPolicyRequest) GetVideoEncAvgKbps() *int32 {
	return s.VideoEncAvgKbps
}

func (s *CreateCenterPolicyRequest) GetVideoEncMaxQP() *int32 {
	return s.VideoEncMaxQP
}

func (s *CreateCenterPolicyRequest) GetVideoEncMinQP() *int32 {
	return s.VideoEncMinQP
}

func (s *CreateCenterPolicyRequest) GetVideoEncPeakKbps() *int32 {
	return s.VideoEncPeakKbps
}

func (s *CreateCenterPolicyRequest) GetVideoEncPolicy() *string {
	return s.VideoEncPolicy
}

func (s *CreateCenterPolicyRequest) GetVideoRedirect() *string {
	return s.VideoRedirect
}

func (s *CreateCenterPolicyRequest) GetVisualQuality() *string {
	return s.VisualQuality
}

func (s *CreateCenterPolicyRequest) GetWatermark() *string {
	return s.Watermark
}

func (s *CreateCenterPolicyRequest) GetWatermarkAntiCam() *string {
	return s.WatermarkAntiCam
}

func (s *CreateCenterPolicyRequest) GetWatermarkColor() *int32 {
	return s.WatermarkColor
}

func (s *CreateCenterPolicyRequest) GetWatermarkColumnAmount() *int32 {
	return s.WatermarkColumnAmount
}

func (s *CreateCenterPolicyRequest) GetWatermarkCustomText() *string {
	return s.WatermarkCustomText
}

func (s *CreateCenterPolicyRequest) GetWatermarkDegree() *float64 {
	return s.WatermarkDegree
}

func (s *CreateCenterPolicyRequest) GetWatermarkFontSize() *int32 {
	return s.WatermarkFontSize
}

func (s *CreateCenterPolicyRequest) GetWatermarkFontStyle() *string {
	return s.WatermarkFontStyle
}

func (s *CreateCenterPolicyRequest) GetWatermarkPower() *string {
	return s.WatermarkPower
}

func (s *CreateCenterPolicyRequest) GetWatermarkRowAmount() *int32 {
	return s.WatermarkRowAmount
}

func (s *CreateCenterPolicyRequest) GetWatermarkSecurity() *string {
	return s.WatermarkSecurity
}

func (s *CreateCenterPolicyRequest) GetWatermarkTransparencyValue() *int32 {
	return s.WatermarkTransparencyValue
}

func (s *CreateCenterPolicyRequest) GetWatermarkType() *string {
	return s.WatermarkType
}

func (s *CreateCenterPolicyRequest) GetWuyingKeeper() *string {
	return s.WuyingKeeper
}

func (s *CreateCenterPolicyRequest) GetWyAssistant() *string {
	return s.WyAssistant
}

func (s *CreateCenterPolicyRequest) SetAdminAccess(v string) *CreateCenterPolicyRequest {
	s.AdminAccess = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetAppContentProtection(v string) *CreateCenterPolicyRequest {
	s.AppContentProtection = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetAuthorizeAccessPolicyRule(v []*CreateCenterPolicyRequestAuthorizeAccessPolicyRule) *CreateCenterPolicyRequest {
	s.AuthorizeAccessPolicyRule = v
	return s
}

func (s *CreateCenterPolicyRequest) SetAuthorizeSecurityPolicyRule(v []*CreateCenterPolicyRequestAuthorizeSecurityPolicyRule) *CreateCenterPolicyRequest {
	s.AuthorizeSecurityPolicyRule = v
	return s
}

func (s *CreateCenterPolicyRequest) SetAutoReconnect(v string) *CreateCenterPolicyRequest {
	s.AutoReconnect = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetBusinessType(v int32) *CreateCenterPolicyRequest {
	s.BusinessType = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetCameraRedirect(v string) *CreateCenterPolicyRequest {
	s.CameraRedirect = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetClientControlMenu(v string) *CreateCenterPolicyRequest {
	s.ClientControlMenu = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetClientType(v []*CreateCenterPolicyRequestClientType) *CreateCenterPolicyRequest {
	s.ClientType = v
	return s
}

func (s *CreateCenterPolicyRequest) SetClipboard(v string) *CreateCenterPolicyRequest {
	s.Clipboard = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetClipboardGraineds(v []*CreateCenterPolicyRequestClipboardGraineds) *CreateCenterPolicyRequest {
	s.ClipboardGraineds = v
	return s
}

func (s *CreateCenterPolicyRequest) SetClipboardScope(v string) *CreateCenterPolicyRequest {
	s.ClipboardScope = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetColorEnhancement(v string) *CreateCenterPolicyRequest {
	s.ColorEnhancement = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetCpdDriveClipboard(v string) *CreateCenterPolicyRequest {
	s.CpdDriveClipboard = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetCpuDownGradeDuration(v int32) *CreateCenterPolicyRequest {
	s.CpuDownGradeDuration = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetCpuProcessors(v []*string) *CreateCenterPolicyRequest {
	s.CpuProcessors = v
	return s
}

func (s *CreateCenterPolicyRequest) SetCpuProtectedMode(v string) *CreateCenterPolicyRequest {
	s.CpuProtectedMode = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetCpuRateLimit(v int32) *CreateCenterPolicyRequest {
	s.CpuRateLimit = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetCpuSampleDuration(v int32) *CreateCenterPolicyRequest {
	s.CpuSampleDuration = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetCpuSingleRateLimit(v int32) *CreateCenterPolicyRequest {
	s.CpuSingleRateLimit = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetDeviceConnectHint(v string) *CreateCenterPolicyRequest {
	s.DeviceConnectHint = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetDeviceRedirects(v []*CreateCenterPolicyRequestDeviceRedirects) *CreateCenterPolicyRequest {
	s.DeviceRedirects = v
	return s
}

func (s *CreateCenterPolicyRequest) SetDeviceRules(v []*CreateCenterPolicyRequestDeviceRules) *CreateCenterPolicyRequest {
	s.DeviceRules = v
	return s
}

func (s *CreateCenterPolicyRequest) SetDisconnectKeepSession(v string) *CreateCenterPolicyRequest {
	s.DisconnectKeepSession = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetDisconnectKeepSessionTime(v int32) *CreateCenterPolicyRequest {
	s.DisconnectKeepSessionTime = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetDisplayMode(v string) *CreateCenterPolicyRequest {
	s.DisplayMode = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetDomainResolveRule(v []*CreateCenterPolicyRequestDomainResolveRule) *CreateCenterPolicyRequest {
	s.DomainResolveRule = v
	return s
}

func (s *CreateCenterPolicyRequest) SetDomainResolveRuleType(v string) *CreateCenterPolicyRequest {
	s.DomainResolveRuleType = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetEnableSessionRateLimiting(v string) *CreateCenterPolicyRequest {
	s.EnableSessionRateLimiting = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetEndUserApplyAdminCoordinate(v string) *CreateCenterPolicyRequest {
	s.EndUserApplyAdminCoordinate = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetEndUserGroupCoordinate(v string) *CreateCenterPolicyRequest {
	s.EndUserGroupCoordinate = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetFileMigrate(v string) *CreateCenterPolicyRequest {
	s.FileMigrate = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetFileTransferAddress(v string) *CreateCenterPolicyRequest {
	s.FileTransferAddress = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetFileTransferSpeed(v string) *CreateCenterPolicyRequest {
	s.FileTransferSpeed = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetFileTransferSpeedLocation(v string) *CreateCenterPolicyRequest {
	s.FileTransferSpeedLocation = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetGpuAcceleration(v string) *CreateCenterPolicyRequest {
	s.GpuAcceleration = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetHtml5FileTransfer(v string) *CreateCenterPolicyRequest {
	s.Html5FileTransfer = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetInternetCommunicationProtocol(v string) *CreateCenterPolicyRequest {
	s.InternetCommunicationProtocol = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetInternetPrinter(v string) *CreateCenterPolicyRequest {
	s.InternetPrinter = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetLocalDrive(v string) *CreateCenterPolicyRequest {
	s.LocalDrive = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetMaxReconnectTime(v int32) *CreateCenterPolicyRequest {
	s.MaxReconnectTime = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetMemoryDownGradeDuration(v int32) *CreateCenterPolicyRequest {
	s.MemoryDownGradeDuration = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetMemoryProcessors(v []*string) *CreateCenterPolicyRequest {
	s.MemoryProcessors = v
	return s
}

func (s *CreateCenterPolicyRequest) SetMemoryProtectedMode(v string) *CreateCenterPolicyRequest {
	s.MemoryProtectedMode = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetMemoryRateLimit(v int32) *CreateCenterPolicyRequest {
	s.MemoryRateLimit = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetMemorySampleDuration(v int32) *CreateCenterPolicyRequest {
	s.MemorySampleDuration = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetMemorySingleRateLimit(v int32) *CreateCenterPolicyRequest {
	s.MemorySingleRateLimit = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetMobileRestart(v string) *CreateCenterPolicyRequest {
	s.MobileRestart = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetMobileSafeMenu(v string) *CreateCenterPolicyRequest {
	s.MobileSafeMenu = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetMobileShutdown(v string) *CreateCenterPolicyRequest {
	s.MobileShutdown = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetMobileWuyingKeeper(v string) *CreateCenterPolicyRequest {
	s.MobileWuyingKeeper = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetMobileWyAssistant(v string) *CreateCenterPolicyRequest {
	s.MobileWyAssistant = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetName(v string) *CreateCenterPolicyRequest {
	s.Name = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetNetRedirect(v string) *CreateCenterPolicyRequest {
	s.NetRedirect = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetNetRedirectRule(v []*CreateCenterPolicyRequestNetRedirectRule) *CreateCenterPolicyRequest {
	s.NetRedirectRule = v
	return s
}

func (s *CreateCenterPolicyRequest) SetNoOperationDisconnect(v string) *CreateCenterPolicyRequest {
	s.NoOperationDisconnect = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetNoOperationDisconnectTime(v int32) *CreateCenterPolicyRequest {
	s.NoOperationDisconnectTime = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetPrinterRedirect(v string) *CreateCenterPolicyRequest {
	s.PrinterRedirect = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetQualityEnhancement(v string) *CreateCenterPolicyRequest {
	s.QualityEnhancement = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetRecordEventDuration(v int32) *CreateCenterPolicyRequest {
	s.RecordEventDuration = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetRecordEventFileExts(v []*string) *CreateCenterPolicyRequest {
	s.RecordEventFileExts = v
	return s
}

func (s *CreateCenterPolicyRequest) SetRecordEventFilePaths(v []*string) *CreateCenterPolicyRequest {
	s.RecordEventFilePaths = v
	return s
}

func (s *CreateCenterPolicyRequest) SetRecordEventLevels(v []*CreateCenterPolicyRequestRecordEventLevels) *CreateCenterPolicyRequest {
	s.RecordEventLevels = v
	return s
}

func (s *CreateCenterPolicyRequest) SetRecordEventRegisters(v []*string) *CreateCenterPolicyRequest {
	s.RecordEventRegisters = v
	return s
}

func (s *CreateCenterPolicyRequest) SetRecordEvents(v []*string) *CreateCenterPolicyRequest {
	s.RecordEvents = v
	return s
}

func (s *CreateCenterPolicyRequest) SetRecording(v string) *CreateCenterPolicyRequest {
	s.Recording = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetRecordingAudio(v string) *CreateCenterPolicyRequest {
	s.RecordingAudio = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetRecordingDuration(v int32) *CreateCenterPolicyRequest {
	s.RecordingDuration = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetRecordingEndTime(v string) *CreateCenterPolicyRequest {
	s.RecordingEndTime = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetRecordingExpires(v int32) *CreateCenterPolicyRequest {
	s.RecordingExpires = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetRecordingFps(v string) *CreateCenterPolicyRequest {
	s.RecordingFps = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetRecordingStartTime(v string) *CreateCenterPolicyRequest {
	s.RecordingStartTime = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetRecordingUserNotify(v string) *CreateCenterPolicyRequest {
	s.RecordingUserNotify = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetRecordingUserNotifyMessage(v string) *CreateCenterPolicyRequest {
	s.RecordingUserNotifyMessage = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetRegionId(v string) *CreateCenterPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetRemoteCoordinate(v string) *CreateCenterPolicyRequest {
	s.RemoteCoordinate = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetResetDesktop(v string) *CreateCenterPolicyRequest {
	s.ResetDesktop = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetResolutionHeight(v int32) *CreateCenterPolicyRequest {
	s.ResolutionHeight = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetResolutionModel(v string) *CreateCenterPolicyRequest {
	s.ResolutionModel = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetResolutionWidth(v int32) *CreateCenterPolicyRequest {
	s.ResolutionWidth = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetResourceType(v string) *CreateCenterPolicyRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetSafeMenu(v string) *CreateCenterPolicyRequest {
	s.SafeMenu = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetScope(v string) *CreateCenterPolicyRequest {
	s.Scope = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetScopeValue(v []*string) *CreateCenterPolicyRequest {
	s.ScopeValue = v
	return s
}

func (s *CreateCenterPolicyRequest) SetScreenDisplayMode(v string) *CreateCenterPolicyRequest {
	s.ScreenDisplayMode = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetSessionMaxRateKbps(v int32) *CreateCenterPolicyRequest {
	s.SessionMaxRateKbps = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetSmoothEnhancement(v string) *CreateCenterPolicyRequest {
	s.SmoothEnhancement = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetStatusMonitor(v string) *CreateCenterPolicyRequest {
	s.StatusMonitor = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetStreamingMode(v string) *CreateCenterPolicyRequest {
	s.StreamingMode = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetTargetFps(v int32) *CreateCenterPolicyRequest {
	s.TargetFps = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetTaskbar(v string) *CreateCenterPolicyRequest {
	s.Taskbar = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetUsbRedirect(v string) *CreateCenterPolicyRequest {
	s.UsbRedirect = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetUsbSupplyRedirectRule(v []*CreateCenterPolicyRequestUsbSupplyRedirectRule) *CreateCenterPolicyRequest {
	s.UsbSupplyRedirectRule = v
	return s
}

func (s *CreateCenterPolicyRequest) SetUseTime(v string) *CreateCenterPolicyRequest {
	s.UseTime = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetVideoEncAvgKbps(v int32) *CreateCenterPolicyRequest {
	s.VideoEncAvgKbps = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetVideoEncMaxQP(v int32) *CreateCenterPolicyRequest {
	s.VideoEncMaxQP = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetVideoEncMinQP(v int32) *CreateCenterPolicyRequest {
	s.VideoEncMinQP = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetVideoEncPeakKbps(v int32) *CreateCenterPolicyRequest {
	s.VideoEncPeakKbps = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetVideoEncPolicy(v string) *CreateCenterPolicyRequest {
	s.VideoEncPolicy = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetVideoRedirect(v string) *CreateCenterPolicyRequest {
	s.VideoRedirect = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetVisualQuality(v string) *CreateCenterPolicyRequest {
	s.VisualQuality = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetWatermark(v string) *CreateCenterPolicyRequest {
	s.Watermark = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetWatermarkAntiCam(v string) *CreateCenterPolicyRequest {
	s.WatermarkAntiCam = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetWatermarkColor(v int32) *CreateCenterPolicyRequest {
	s.WatermarkColor = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetWatermarkColumnAmount(v int32) *CreateCenterPolicyRequest {
	s.WatermarkColumnAmount = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetWatermarkCustomText(v string) *CreateCenterPolicyRequest {
	s.WatermarkCustomText = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetWatermarkDegree(v float64) *CreateCenterPolicyRequest {
	s.WatermarkDegree = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetWatermarkFontSize(v int32) *CreateCenterPolicyRequest {
	s.WatermarkFontSize = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetWatermarkFontStyle(v string) *CreateCenterPolicyRequest {
	s.WatermarkFontStyle = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetWatermarkPower(v string) *CreateCenterPolicyRequest {
	s.WatermarkPower = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetWatermarkRowAmount(v int32) *CreateCenterPolicyRequest {
	s.WatermarkRowAmount = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetWatermarkSecurity(v string) *CreateCenterPolicyRequest {
	s.WatermarkSecurity = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetWatermarkTransparencyValue(v int32) *CreateCenterPolicyRequest {
	s.WatermarkTransparencyValue = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetWatermarkType(v string) *CreateCenterPolicyRequest {
	s.WatermarkType = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetWuyingKeeper(v string) *CreateCenterPolicyRequest {
	s.WuyingKeeper = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetWyAssistant(v string) *CreateCenterPolicyRequest {
	s.WyAssistant = &v
	return s
}

func (s *CreateCenterPolicyRequest) Validate() error {
	return dara.Validate(s)
}

type CreateCenterPolicyRequestAuthorizeAccessPolicyRule struct {
	// The client CIDR block from which end users can connect to cloud computers. Specify an IPv4 CIDR block.
	//
	// example:
	//
	// 47.100.XX.XX/16
	CidrIp *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	// The description of the client IP address whitelist.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s CreateCenterPolicyRequestAuthorizeAccessPolicyRule) String() string {
	return dara.Prettify(s)
}

func (s CreateCenterPolicyRequestAuthorizeAccessPolicyRule) GoString() string {
	return s.String()
}

func (s *CreateCenterPolicyRequestAuthorizeAccessPolicyRule) GetCidrIp() *string {
	return s.CidrIp
}

func (s *CreateCenterPolicyRequestAuthorizeAccessPolicyRule) GetDescription() *string {
	return s.Description
}

func (s *CreateCenterPolicyRequestAuthorizeAccessPolicyRule) SetCidrIp(v string) *CreateCenterPolicyRequestAuthorizeAccessPolicyRule {
	s.CidrIp = &v
	return s
}

func (s *CreateCenterPolicyRequestAuthorizeAccessPolicyRule) SetDescription(v string) *CreateCenterPolicyRequestAuthorizeAccessPolicyRule {
	s.Description = &v
	return s
}

func (s *CreateCenterPolicyRequestAuthorizeAccessPolicyRule) Validate() error {
	return dara.Validate(s)
}

type CreateCenterPolicyRequestAuthorizeSecurityPolicyRule struct {
	// The object of the security group rule. Specify an IPv4 CIDR block.
	//
	// example:
	//
	// 10.0.XX.XX/8
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
	// 	- ALL: any type of protocol.
	//
	// 	- GRE: the Generic Routing Encapsulation (GRE) protocol.
	//
	// 	- ICMP: the Internet Control Message Protocol (ICMP) for (IPv4).
	//
	// example:
	//
	// TCP
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The authorization policy of the security group rule.
	//
	// Valid values:
	//
	// 	- drop: denies all access requests. If no \\"\\"access denied\\"\\" messages are returned, the requests either timed out or failed.
	//
	// 	- accept (default): accepts all requests.
	//
	// example:
	//
	// accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The port range of the security group rule. The value range of this parameter varies based on the value of IpProtocol.
	//
	// 	- If IpProtocol is set to TCP or UDP, the port range is 1 to 65535. Separate the start port number and the end port number with a forward slash (/). Example: 1/200.
	//
	// 	- If IpProtocol is set to ICMP, set the value to -1/-1.
	//
	// 	- If IpProtocol is set to GRE, set the value to -1/-1.
	//
	// 	- If IpProtocol is set to ALL, set the value to -1/-1.
	//
	// For more information about the common ports, see [Common ports](https://help.aliyun.com/document_detail/40724.html).
	//
	// example:
	//
	// 22/22
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// The priority of the security group rule. A smaller value specifies a higher priority.\\
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

func (s CreateCenterPolicyRequestAuthorizeSecurityPolicyRule) String() string {
	return dara.Prettify(s)
}

func (s CreateCenterPolicyRequestAuthorizeSecurityPolicyRule) GoString() string {
	return s.String()
}

func (s *CreateCenterPolicyRequestAuthorizeSecurityPolicyRule) GetCidrIp() *string {
	return s.CidrIp
}

func (s *CreateCenterPolicyRequestAuthorizeSecurityPolicyRule) GetDescription() *string {
	return s.Description
}

func (s *CreateCenterPolicyRequestAuthorizeSecurityPolicyRule) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *CreateCenterPolicyRequestAuthorizeSecurityPolicyRule) GetPolicy() *string {
	return s.Policy
}

func (s *CreateCenterPolicyRequestAuthorizeSecurityPolicyRule) GetPortRange() *string {
	return s.PortRange
}

func (s *CreateCenterPolicyRequestAuthorizeSecurityPolicyRule) GetPriority() *string {
	return s.Priority
}

func (s *CreateCenterPolicyRequestAuthorizeSecurityPolicyRule) GetType() *string {
	return s.Type
}

func (s *CreateCenterPolicyRequestAuthorizeSecurityPolicyRule) SetCidrIp(v string) *CreateCenterPolicyRequestAuthorizeSecurityPolicyRule {
	s.CidrIp = &v
	return s
}

func (s *CreateCenterPolicyRequestAuthorizeSecurityPolicyRule) SetDescription(v string) *CreateCenterPolicyRequestAuthorizeSecurityPolicyRule {
	s.Description = &v
	return s
}

func (s *CreateCenterPolicyRequestAuthorizeSecurityPolicyRule) SetIpProtocol(v string) *CreateCenterPolicyRequestAuthorizeSecurityPolicyRule {
	s.IpProtocol = &v
	return s
}

func (s *CreateCenterPolicyRequestAuthorizeSecurityPolicyRule) SetPolicy(v string) *CreateCenterPolicyRequestAuthorizeSecurityPolicyRule {
	s.Policy = &v
	return s
}

func (s *CreateCenterPolicyRequestAuthorizeSecurityPolicyRule) SetPortRange(v string) *CreateCenterPolicyRequestAuthorizeSecurityPolicyRule {
	s.PortRange = &v
	return s
}

func (s *CreateCenterPolicyRequestAuthorizeSecurityPolicyRule) SetPriority(v string) *CreateCenterPolicyRequestAuthorizeSecurityPolicyRule {
	s.Priority = &v
	return s
}

func (s *CreateCenterPolicyRequestAuthorizeSecurityPolicyRule) SetType(v string) *CreateCenterPolicyRequestAuthorizeSecurityPolicyRule {
	s.Type = &v
	return s
}

func (s *CreateCenterPolicyRequestAuthorizeSecurityPolicyRule) Validate() error {
	return dara.Validate(s)
}

type CreateCenterPolicyRequestClientType struct {
	// The type of the Alibaba Cloud Workspace client that end users can use to connect to cloud computers.
	//
	// Valid values:
	//
	// 	- html5: the web client.
	//
	// 	- android: the Android client.
	//
	// 	- ios: the iOS client.
	//
	// 	- windows: the Windows client.
	//
	// 	- macos: the macOS client.
	//
	// example:
	//
	// windows
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// Specifies whether end users can use the specified type of Alibaba Cloud Workspace client to connect to cloud computers.
	//
	// >  If you don\\"t specify `ClientType`, any client can be used to connect to cloud computers.
	//
	// Valid values:
	//
	// 	- off: End users cannot use the specified type of Alibaba Cloud Workspace client to connect to cloud computers.
	//
	// 	- on: End users can use the specified type of Alibaba Cloud Workspace client to connect to cloud computers.
	//
	// example:
	//
	// off
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateCenterPolicyRequestClientType) String() string {
	return dara.Prettify(s)
}

func (s CreateCenterPolicyRequestClientType) GoString() string {
	return s.String()
}

func (s *CreateCenterPolicyRequestClientType) GetClientType() *string {
	return s.ClientType
}

func (s *CreateCenterPolicyRequestClientType) GetStatus() *string {
	return s.Status
}

func (s *CreateCenterPolicyRequestClientType) SetClientType(v string) *CreateCenterPolicyRequestClientType {
	s.ClientType = &v
	return s
}

func (s *CreateCenterPolicyRequestClientType) SetStatus(v string) *CreateCenterPolicyRequestClientType {
	s.Status = &v
	return s
}

func (s *CreateCenterPolicyRequestClientType) Validate() error {
	return dara.Validate(s)
}

type CreateCenterPolicyRequestClipboardGraineds struct {
	ClipboardSize *int32  `json:"ClipboardSize,omitempty" xml:"ClipboardSize,omitempty"`
	ClipboardType *string `json:"ClipboardType,omitempty" xml:"ClipboardType,omitempty"`
	GrainedType   *string `json:"GrainedType,omitempty" xml:"GrainedType,omitempty"`
}

func (s CreateCenterPolicyRequestClipboardGraineds) String() string {
	return dara.Prettify(s)
}

func (s CreateCenterPolicyRequestClipboardGraineds) GoString() string {
	return s.String()
}

func (s *CreateCenterPolicyRequestClipboardGraineds) GetClipboardSize() *int32 {
	return s.ClipboardSize
}

func (s *CreateCenterPolicyRequestClipboardGraineds) GetClipboardType() *string {
	return s.ClipboardType
}

func (s *CreateCenterPolicyRequestClipboardGraineds) GetGrainedType() *string {
	return s.GrainedType
}

func (s *CreateCenterPolicyRequestClipboardGraineds) SetClipboardSize(v int32) *CreateCenterPolicyRequestClipboardGraineds {
	s.ClipboardSize = &v
	return s
}

func (s *CreateCenterPolicyRequestClipboardGraineds) SetClipboardType(v string) *CreateCenterPolicyRequestClipboardGraineds {
	s.ClipboardType = &v
	return s
}

func (s *CreateCenterPolicyRequestClipboardGraineds) SetGrainedType(v string) *CreateCenterPolicyRequestClipboardGraineds {
	s.GrainedType = &v
	return s
}

func (s *CreateCenterPolicyRequestClipboardGraineds) Validate() error {
	return dara.Validate(s)
}

type CreateCenterPolicyRequestDeviceRedirects struct {
	// The peripheral type.
	//
	// Valid values:
	//
	// 	- printer
	//
	// 	- scanner
	//
	// 	- serialport
	//
	// 	- camera
	//
	// 	- adb
	//
	// example:
	//
	// camera
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// The redirection type.
	//
	// Valid values:
	//
	// 	- deviceRedirect: enables device redirection.
	//
	// 	- usbRedirect: enables USB redirection.
	//
	// 	- off: disables any type of redirection.
	//
	// example:
	//
	// usbRedirect
	RedirectType *string `json:"RedirectType,omitempty" xml:"RedirectType,omitempty"`
}

func (s CreateCenterPolicyRequestDeviceRedirects) String() string {
	return dara.Prettify(s)
}

func (s CreateCenterPolicyRequestDeviceRedirects) GoString() string {
	return s.String()
}

func (s *CreateCenterPolicyRequestDeviceRedirects) GetDeviceType() *string {
	return s.DeviceType
}

func (s *CreateCenterPolicyRequestDeviceRedirects) GetRedirectType() *string {
	return s.RedirectType
}

func (s *CreateCenterPolicyRequestDeviceRedirects) SetDeviceType(v string) *CreateCenterPolicyRequestDeviceRedirects {
	s.DeviceType = &v
	return s
}

func (s *CreateCenterPolicyRequestDeviceRedirects) SetRedirectType(v string) *CreateCenterPolicyRequestDeviceRedirects {
	s.RedirectType = &v
	return s
}

func (s *CreateCenterPolicyRequestDeviceRedirects) Validate() error {
	return dara.Validate(s)
}

type CreateCenterPolicyRequestDeviceRules struct {
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
	// 	- cardReader: card readers.
	//
	// 	- printer: printers.
	//
	// 	- scanner: scanners.
	//
	// 	- storage: storage devices.
	//
	// 	- camera: cameras.
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
	// 	- deviceRedirect: device redirection.
	//
	// 	- usbRedirect: USB redirection.
	//
	// 	- off: redirection disabled.
	//
	// example:
	//
	// usbRedirect
	RedirectType *string `json:"RedirectType,omitempty" xml:"RedirectType,omitempty"`
}

func (s CreateCenterPolicyRequestDeviceRules) String() string {
	return dara.Prettify(s)
}

func (s CreateCenterPolicyRequestDeviceRules) GoString() string {
	return s.String()
}

func (s *CreateCenterPolicyRequestDeviceRules) GetDeviceName() *string {
	return s.DeviceName
}

func (s *CreateCenterPolicyRequestDeviceRules) GetDevicePid() *string {
	return s.DevicePid
}

func (s *CreateCenterPolicyRequestDeviceRules) GetDeviceType() *string {
	return s.DeviceType
}

func (s *CreateCenterPolicyRequestDeviceRules) GetDeviceVid() *string {
	return s.DeviceVid
}

func (s *CreateCenterPolicyRequestDeviceRules) GetOptCommand() *string {
	return s.OptCommand
}

func (s *CreateCenterPolicyRequestDeviceRules) GetPlatforms() *string {
	return s.Platforms
}

func (s *CreateCenterPolicyRequestDeviceRules) GetRedirectType() *string {
	return s.RedirectType
}

func (s *CreateCenterPolicyRequestDeviceRules) SetDeviceName(v string) *CreateCenterPolicyRequestDeviceRules {
	s.DeviceName = &v
	return s
}

func (s *CreateCenterPolicyRequestDeviceRules) SetDevicePid(v string) *CreateCenterPolicyRequestDeviceRules {
	s.DevicePid = &v
	return s
}

func (s *CreateCenterPolicyRequestDeviceRules) SetDeviceType(v string) *CreateCenterPolicyRequestDeviceRules {
	s.DeviceType = &v
	return s
}

func (s *CreateCenterPolicyRequestDeviceRules) SetDeviceVid(v string) *CreateCenterPolicyRequestDeviceRules {
	s.DeviceVid = &v
	return s
}

func (s *CreateCenterPolicyRequestDeviceRules) SetOptCommand(v string) *CreateCenterPolicyRequestDeviceRules {
	s.OptCommand = &v
	return s
}

func (s *CreateCenterPolicyRequestDeviceRules) SetPlatforms(v string) *CreateCenterPolicyRequestDeviceRules {
	s.Platforms = &v
	return s
}

func (s *CreateCenterPolicyRequestDeviceRules) SetRedirectType(v string) *CreateCenterPolicyRequestDeviceRules {
	s.RedirectType = &v
	return s
}

func (s *CreateCenterPolicyRequestDeviceRules) Validate() error {
	return dara.Validate(s)
}

type CreateCenterPolicyRequestDomainResolveRule struct {
	// The policy description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The domain name.
	//
	// example:
	//
	// *.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Specifies whether to allow the domain resolution policy to take effect.
	//
	// Valid values:
	//
	// 	- allow
	//
	// 	- block
	//
	// example:
	//
	// allow
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
}

func (s CreateCenterPolicyRequestDomainResolveRule) String() string {
	return dara.Prettify(s)
}

func (s CreateCenterPolicyRequestDomainResolveRule) GoString() string {
	return s.String()
}

func (s *CreateCenterPolicyRequestDomainResolveRule) GetDescription() *string {
	return s.Description
}

func (s *CreateCenterPolicyRequestDomainResolveRule) GetDomain() *string {
	return s.Domain
}

func (s *CreateCenterPolicyRequestDomainResolveRule) GetPolicy() *string {
	return s.Policy
}

func (s *CreateCenterPolicyRequestDomainResolveRule) SetDescription(v string) *CreateCenterPolicyRequestDomainResolveRule {
	s.Description = &v
	return s
}

func (s *CreateCenterPolicyRequestDomainResolveRule) SetDomain(v string) *CreateCenterPolicyRequestDomainResolveRule {
	s.Domain = &v
	return s
}

func (s *CreateCenterPolicyRequestDomainResolveRule) SetPolicy(v string) *CreateCenterPolicyRequestDomainResolveRule {
	s.Policy = &v
	return s
}

func (s *CreateCenterPolicyRequestDomainResolveRule) Validate() error {
	return dara.Validate(s)
}

type CreateCenterPolicyRequestNetRedirectRule struct {
	// The domain name.
	//
	// example:
	//
	// *.taobao.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The redirection policy.
	//
	// example:
	//
	// allow
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The rule type.
	//
	// Valid values:
	//
	// 	- prc: process.
	//
	// 	- domain: domain name.
	//
	// example:
	//
	// domain
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s CreateCenterPolicyRequestNetRedirectRule) String() string {
	return dara.Prettify(s)
}

func (s CreateCenterPolicyRequestNetRedirectRule) GoString() string {
	return s.String()
}

func (s *CreateCenterPolicyRequestNetRedirectRule) GetDomain() *string {
	return s.Domain
}

func (s *CreateCenterPolicyRequestNetRedirectRule) GetPolicy() *string {
	return s.Policy
}

func (s *CreateCenterPolicyRequestNetRedirectRule) GetRuleType() *string {
	return s.RuleType
}

func (s *CreateCenterPolicyRequestNetRedirectRule) SetDomain(v string) *CreateCenterPolicyRequestNetRedirectRule {
	s.Domain = &v
	return s
}

func (s *CreateCenterPolicyRequestNetRedirectRule) SetPolicy(v string) *CreateCenterPolicyRequestNetRedirectRule {
	s.Policy = &v
	return s
}

func (s *CreateCenterPolicyRequestNetRedirectRule) SetRuleType(v string) *CreateCenterPolicyRequestNetRedirectRule {
	s.RuleType = &v
	return s
}

func (s *CreateCenterPolicyRequestNetRedirectRule) Validate() error {
	return dara.Validate(s)
}

type CreateCenterPolicyRequestRecordEventLevels struct {
	// example:
	//
	// HIGH
	EventLevel *string `json:"EventLevel,omitempty" xml:"EventLevel,omitempty"`
	// example:
	//
	// StartApplication
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
}

func (s CreateCenterPolicyRequestRecordEventLevels) String() string {
	return dara.Prettify(s)
}

func (s CreateCenterPolicyRequestRecordEventLevels) GoString() string {
	return s.String()
}

func (s *CreateCenterPolicyRequestRecordEventLevels) GetEventLevel() *string {
	return s.EventLevel
}

func (s *CreateCenterPolicyRequestRecordEventLevels) GetEventType() *string {
	return s.EventType
}

func (s *CreateCenterPolicyRequestRecordEventLevels) SetEventLevel(v string) *CreateCenterPolicyRequestRecordEventLevels {
	s.EventLevel = &v
	return s
}

func (s *CreateCenterPolicyRequestRecordEventLevels) SetEventType(v string) *CreateCenterPolicyRequestRecordEventLevels {
	s.EventType = &v
	return s
}

func (s *CreateCenterPolicyRequestRecordEventLevels) Validate() error {
	return dara.Validate(s)
}

type CreateCenterPolicyRequestUsbSupplyRedirectRule struct {
	// The rule description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The product ID (PID).
	//
	// example:
	//
	// 08**
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// Specifies whether to allow USB redirection.
	//
	// Valid values:
	//
	// 	- 1: allows USB redirection.
	//
	// 	- 2: forbids USB redirection.
	//
	// example:
	//
	// 1
	UsbRedirectType *string `json:"UsbRedirectType,omitempty" xml:"UsbRedirectType,omitempty"`
	// The type of the USB redirection rule.
	//
	// Valid values:
	//
	// 	- 2: enables USB redirection based on products.
	//
	// example:
	//
	// 2
	UsbRuleType *string `json:"UsbRuleType,omitempty" xml:"UsbRuleType,omitempty"`
	// The vendor ID (VID). For more information, see [Valid USB Vendor IDs (VIDs)](https://www.usb.org/sites/default/files/vendor_ids032322.pdf_1.pdf).
	//
	// example:
	//
	// 04**
	VendorId *string `json:"VendorId,omitempty" xml:"VendorId,omitempty"`
}

func (s CreateCenterPolicyRequestUsbSupplyRedirectRule) String() string {
	return dara.Prettify(s)
}

func (s CreateCenterPolicyRequestUsbSupplyRedirectRule) GoString() string {
	return s.String()
}

func (s *CreateCenterPolicyRequestUsbSupplyRedirectRule) GetDescription() *string {
	return s.Description
}

func (s *CreateCenterPolicyRequestUsbSupplyRedirectRule) GetProductId() *string {
	return s.ProductId
}

func (s *CreateCenterPolicyRequestUsbSupplyRedirectRule) GetUsbRedirectType() *string {
	return s.UsbRedirectType
}

func (s *CreateCenterPolicyRequestUsbSupplyRedirectRule) GetUsbRuleType() *string {
	return s.UsbRuleType
}

func (s *CreateCenterPolicyRequestUsbSupplyRedirectRule) GetVendorId() *string {
	return s.VendorId
}

func (s *CreateCenterPolicyRequestUsbSupplyRedirectRule) SetDescription(v string) *CreateCenterPolicyRequestUsbSupplyRedirectRule {
	s.Description = &v
	return s
}

func (s *CreateCenterPolicyRequestUsbSupplyRedirectRule) SetProductId(v string) *CreateCenterPolicyRequestUsbSupplyRedirectRule {
	s.ProductId = &v
	return s
}

func (s *CreateCenterPolicyRequestUsbSupplyRedirectRule) SetUsbRedirectType(v string) *CreateCenterPolicyRequestUsbSupplyRedirectRule {
	s.UsbRedirectType = &v
	return s
}

func (s *CreateCenterPolicyRequestUsbSupplyRedirectRule) SetUsbRuleType(v string) *CreateCenterPolicyRequestUsbSupplyRedirectRule {
	s.UsbRuleType = &v
	return s
}

func (s *CreateCenterPolicyRequestUsbSupplyRedirectRule) SetVendorId(v string) *CreateCenterPolicyRequestUsbSupplyRedirectRule {
	s.VendorId = &v
	return s
}

func (s *CreateCenterPolicyRequestUsbSupplyRedirectRule) Validate() error {
	return dara.Validate(s)
}
