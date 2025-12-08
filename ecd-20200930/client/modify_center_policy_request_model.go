// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCenterPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcademicProxy(v string) *ModifyCenterPolicyRequest
	GetAcademicProxy() *string
	SetAdminAccess(v string) *ModifyCenterPolicyRequest
	GetAdminAccess() *string
	SetAppContentProtection(v string) *ModifyCenterPolicyRequest
	GetAppContentProtection() *string
	SetAuthorizeAccessPolicyRule(v []*ModifyCenterPolicyRequestAuthorizeAccessPolicyRule) *ModifyCenterPolicyRequest
	GetAuthorizeAccessPolicyRule() []*ModifyCenterPolicyRequestAuthorizeAccessPolicyRule
	SetAuthorizeSecurityPolicyRule(v []*ModifyCenterPolicyRequestAuthorizeSecurityPolicyRule) *ModifyCenterPolicyRequest
	GetAuthorizeSecurityPolicyRule() []*ModifyCenterPolicyRequestAuthorizeSecurityPolicyRule
	SetAutoReconnect(v string) *ModifyCenterPolicyRequest
	GetAutoReconnect() *string
	SetBusinessChannel(v string) *ModifyCenterPolicyRequest
	GetBusinessChannel() *string
	SetBusinessType(v int32) *ModifyCenterPolicyRequest
	GetBusinessType() *int32
	SetCameraRedirect(v string) *ModifyCenterPolicyRequest
	GetCameraRedirect() *string
	SetClientControlMenu(v string) *ModifyCenterPolicyRequest
	GetClientControlMenu() *string
	SetClientCreateSnapshot(v string) *ModifyCenterPolicyRequest
	GetClientCreateSnapshot() *string
	SetClientType(v []*ModifyCenterPolicyRequestClientType) *ModifyCenterPolicyRequest
	GetClientType() []*ModifyCenterPolicyRequestClientType
	SetClipboard(v string) *ModifyCenterPolicyRequest
	GetClipboard() *string
	SetClipboardGraineds(v []*ModifyCenterPolicyRequestClipboardGraineds) *ModifyCenterPolicyRequest
	GetClipboardGraineds() []*ModifyCenterPolicyRequestClipboardGraineds
	SetClipboardScope(v string) *ModifyCenterPolicyRequest
	GetClipboardScope() *string
	SetColorEnhancement(v string) *ModifyCenterPolicyRequest
	GetColorEnhancement() *string
	SetCpdDriveClipboard(v string) *ModifyCenterPolicyRequest
	GetCpdDriveClipboard() *string
	SetCpuDownGradeDuration(v int32) *ModifyCenterPolicyRequest
	GetCpuDownGradeDuration() *int32
	SetCpuOverload(v string) *ModifyCenterPolicyRequest
	GetCpuOverload() *string
	SetCpuProcessors(v []*string) *ModifyCenterPolicyRequest
	GetCpuProcessors() []*string
	SetCpuProtectedMode(v string) *ModifyCenterPolicyRequest
	GetCpuProtectedMode() *string
	SetCpuRateLimit(v int32) *ModifyCenterPolicyRequest
	GetCpuRateLimit() *int32
	SetCpuSampleDuration(v int32) *ModifyCenterPolicyRequest
	GetCpuSampleDuration() *int32
	SetCpuSingleRateLimit(v int32) *ModifyCenterPolicyRequest
	GetCpuSingleRateLimit() *int32
	SetDeviceConnectHint(v string) *ModifyCenterPolicyRequest
	GetDeviceConnectHint() *string
	SetDeviceRedirects(v []*ModifyCenterPolicyRequestDeviceRedirects) *ModifyCenterPolicyRequest
	GetDeviceRedirects() []*ModifyCenterPolicyRequestDeviceRedirects
	SetDeviceRules(v []*ModifyCenterPolicyRequestDeviceRules) *ModifyCenterPolicyRequest
	GetDeviceRules() []*ModifyCenterPolicyRequestDeviceRules
	SetDisconnectKeepSession(v string) *ModifyCenterPolicyRequest
	GetDisconnectKeepSession() *string
	SetDisconnectKeepSessionTime(v int32) *ModifyCenterPolicyRequest
	GetDisconnectKeepSessionTime() *int32
	SetDiskOverload(v string) *ModifyCenterPolicyRequest
	GetDiskOverload() *string
	SetDisplayMode(v string) *ModifyCenterPolicyRequest
	GetDisplayMode() *string
	SetDomainResolveRule(v []*ModifyCenterPolicyRequestDomainResolveRule) *ModifyCenterPolicyRequest
	GetDomainResolveRule() []*ModifyCenterPolicyRequestDomainResolveRule
	SetDomainResolveRuleType(v string) *ModifyCenterPolicyRequest
	GetDomainResolveRuleType() *string
	SetEnableSessionRateLimiting(v string) *ModifyCenterPolicyRequest
	GetEnableSessionRateLimiting() *string
	SetEndUserApplyAdminCoordinate(v string) *ModifyCenterPolicyRequest
	GetEndUserApplyAdminCoordinate() *string
	SetEndUserGroupCoordinate(v string) *ModifyCenterPolicyRequest
	GetEndUserGroupCoordinate() *string
	SetExternalDrive(v string) *ModifyCenterPolicyRequest
	GetExternalDrive() *string
	SetFileMigrate(v string) *ModifyCenterPolicyRequest
	GetFileMigrate() *string
	SetFileTransferAddress(v string) *ModifyCenterPolicyRequest
	GetFileTransferAddress() *string
	SetFileTransferSpeed(v string) *ModifyCenterPolicyRequest
	GetFileTransferSpeed() *string
	SetFileTransferSpeedLocation(v string) *ModifyCenterPolicyRequest
	GetFileTransferSpeedLocation() *string
	SetGpuAcceleration(v string) *ModifyCenterPolicyRequest
	GetGpuAcceleration() *string
	SetHtml5FileTransfer(v string) *ModifyCenterPolicyRequest
	GetHtml5FileTransfer() *string
	SetInternetCommunicationProtocol(v string) *ModifyCenterPolicyRequest
	GetInternetCommunicationProtocol() *string
	SetInternetPrinter(v string) *ModifyCenterPolicyRequest
	GetInternetPrinter() *string
	SetLocalDrive(v string) *ModifyCenterPolicyRequest
	GetLocalDrive() *string
	SetMaxReconnectTime(v int32) *ModifyCenterPolicyRequest
	GetMaxReconnectTime() *int32
	SetMemoryDownGradeDuration(v int32) *ModifyCenterPolicyRequest
	GetMemoryDownGradeDuration() *int32
	SetMemoryOverload(v string) *ModifyCenterPolicyRequest
	GetMemoryOverload() *string
	SetMemoryProcessors(v []*string) *ModifyCenterPolicyRequest
	GetMemoryProcessors() []*string
	SetMemoryProtectedMode(v string) *ModifyCenterPolicyRequest
	GetMemoryProtectedMode() *string
	SetMemoryRateLimit(v int32) *ModifyCenterPolicyRequest
	GetMemoryRateLimit() *int32
	SetMemorySampleDuration(v int32) *ModifyCenterPolicyRequest
	GetMemorySampleDuration() *int32
	SetMemorySingleRateLimit(v int32) *ModifyCenterPolicyRequest
	GetMemorySingleRateLimit() *int32
	SetMobileRestart(v string) *ModifyCenterPolicyRequest
	GetMobileRestart() *string
	SetMobileSafeMenu(v string) *ModifyCenterPolicyRequest
	GetMobileSafeMenu() *string
	SetMobileShutdown(v string) *ModifyCenterPolicyRequest
	GetMobileShutdown() *string
	SetMobileWuyingKeeper(v string) *ModifyCenterPolicyRequest
	GetMobileWuyingKeeper() *string
	SetMobileWyAssistant(v string) *ModifyCenterPolicyRequest
	GetMobileWyAssistant() *string
	SetModelLibrary(v string) *ModifyCenterPolicyRequest
	GetModelLibrary() *string
	SetName(v string) *ModifyCenterPolicyRequest
	GetName() *string
	SetNetRedirect(v string) *ModifyCenterPolicyRequest
	GetNetRedirect() *string
	SetNetRedirectRule(v []*ModifyCenterPolicyRequestNetRedirectRule) *ModifyCenterPolicyRequest
	GetNetRedirectRule() []*ModifyCenterPolicyRequestNetRedirectRule
	SetNoOperationDisconnect(v string) *ModifyCenterPolicyRequest
	GetNoOperationDisconnect() *string
	SetNoOperationDisconnectTime(v int32) *ModifyCenterPolicyRequest
	GetNoOperationDisconnectTime() *int32
	SetPolicyGroupId(v string) *ModifyCenterPolicyRequest
	GetPolicyGroupId() *string
	SetPortProxy(v string) *ModifyCenterPolicyRequest
	GetPortProxy() *string
	SetPrinterRedirect(v string) *ModifyCenterPolicyRequest
	GetPrinterRedirect() *string
	SetQualityEnhancement(v string) *ModifyCenterPolicyRequest
	GetQualityEnhancement() *string
	SetRecordEventDuration(v int32) *ModifyCenterPolicyRequest
	GetRecordEventDuration() *int32
	SetRecordEventFileExts(v []*string) *ModifyCenterPolicyRequest
	GetRecordEventFileExts() []*string
	SetRecordEventFilePaths(v []*string) *ModifyCenterPolicyRequest
	GetRecordEventFilePaths() []*string
	SetRecordEventLevels(v []*ModifyCenterPolicyRequestRecordEventLevels) *ModifyCenterPolicyRequest
	GetRecordEventLevels() []*ModifyCenterPolicyRequestRecordEventLevels
	SetRecordEventRegisters(v []*string) *ModifyCenterPolicyRequest
	GetRecordEventRegisters() []*string
	SetRecordEvents(v []*string) *ModifyCenterPolicyRequest
	GetRecordEvents() []*string
	SetRecording(v string) *ModifyCenterPolicyRequest
	GetRecording() *string
	SetRecordingAudio(v string) *ModifyCenterPolicyRequest
	GetRecordingAudio() *string
	SetRecordingDuration(v int32) *ModifyCenterPolicyRequest
	GetRecordingDuration() *int32
	SetRecordingEndTime(v string) *ModifyCenterPolicyRequest
	GetRecordingEndTime() *string
	SetRecordingExpires(v int32) *ModifyCenterPolicyRequest
	GetRecordingExpires() *int32
	SetRecordingFps(v string) *ModifyCenterPolicyRequest
	GetRecordingFps() *string
	SetRecordingStartTime(v string) *ModifyCenterPolicyRequest
	GetRecordingStartTime() *string
	SetRecordingUserNotify(v string) *ModifyCenterPolicyRequest
	GetRecordingUserNotify() *string
	SetRecordingUserNotifyMessage(v string) *ModifyCenterPolicyRequest
	GetRecordingUserNotifyMessage() *string
	SetRegionId(v string) *ModifyCenterPolicyRequest
	GetRegionId() *string
	SetRemoteCoordinate(v string) *ModifyCenterPolicyRequest
	GetRemoteCoordinate() *string
	SetResetDesktop(v string) *ModifyCenterPolicyRequest
	GetResetDesktop() *string
	SetResolutionHeight(v int32) *ModifyCenterPolicyRequest
	GetResolutionHeight() *int32
	SetResolutionModel(v string) *ModifyCenterPolicyRequest
	GetResolutionModel() *string
	SetResolutionWidth(v int32) *ModifyCenterPolicyRequest
	GetResolutionWidth() *int32
	SetResourceType(v string) *ModifyCenterPolicyRequest
	GetResourceType() *string
	SetRevokeAccessPolicyRule(v []*ModifyCenterPolicyRequestRevokeAccessPolicyRule) *ModifyCenterPolicyRequest
	GetRevokeAccessPolicyRule() []*ModifyCenterPolicyRequestRevokeAccessPolicyRule
	SetRevokeSecurityPolicyRule(v []*ModifyCenterPolicyRequestRevokeSecurityPolicyRule) *ModifyCenterPolicyRequest
	GetRevokeSecurityPolicyRule() []*ModifyCenterPolicyRequestRevokeSecurityPolicyRule
	SetSafeMenu(v string) *ModifyCenterPolicyRequest
	GetSafeMenu() *string
	SetScope(v string) *ModifyCenterPolicyRequest
	GetScope() *string
	SetScopeValue(v []*string) *ModifyCenterPolicyRequest
	GetScopeValue() []*string
	SetScreenDisplayMode(v string) *ModifyCenterPolicyRequest
	GetScreenDisplayMode() *string
	SetSessionMaxRateKbps(v int32) *ModifyCenterPolicyRequest
	GetSessionMaxRateKbps() *int32
	SetSmoothEnhancement(v string) *ModifyCenterPolicyRequest
	GetSmoothEnhancement() *string
	SetStatusMonitor(v string) *ModifyCenterPolicyRequest
	GetStatusMonitor() *string
	SetStreamingMode(v string) *ModifyCenterPolicyRequest
	GetStreamingMode() *string
	SetTargetFps(v int32) *ModifyCenterPolicyRequest
	GetTargetFps() *int32
	SetTaskbar(v string) *ModifyCenterPolicyRequest
	GetTaskbar() *string
	SetUsbRedirect(v string) *ModifyCenterPolicyRequest
	GetUsbRedirect() *string
	SetUsbSupplyRedirectRule(v []*ModifyCenterPolicyRequestUsbSupplyRedirectRule) *ModifyCenterPolicyRequest
	GetUsbSupplyRedirectRule() []*ModifyCenterPolicyRequestUsbSupplyRedirectRule
	SetUseTime(v string) *ModifyCenterPolicyRequest
	GetUseTime() *string
	SetVideoEncAvgKbps(v int32) *ModifyCenterPolicyRequest
	GetVideoEncAvgKbps() *int32
	SetVideoEncMaxQP(v int32) *ModifyCenterPolicyRequest
	GetVideoEncMaxQP() *int32
	SetVideoEncMinQP(v int32) *ModifyCenterPolicyRequest
	GetVideoEncMinQP() *int32
	SetVideoEncPeakKbps(v int32) *ModifyCenterPolicyRequest
	GetVideoEncPeakKbps() *int32
	SetVideoEncPolicy(v string) *ModifyCenterPolicyRequest
	GetVideoEncPolicy() *string
	SetVideoRedirect(v string) *ModifyCenterPolicyRequest
	GetVideoRedirect() *string
	SetVisualQuality(v string) *ModifyCenterPolicyRequest
	GetVisualQuality() *string
	SetWatermark(v string) *ModifyCenterPolicyRequest
	GetWatermark() *string
	SetWatermarkAntiCam(v string) *ModifyCenterPolicyRequest
	GetWatermarkAntiCam() *string
	SetWatermarkColor(v int32) *ModifyCenterPolicyRequest
	GetWatermarkColor() *int32
	SetWatermarkColumnAmount(v int32) *ModifyCenterPolicyRequest
	GetWatermarkColumnAmount() *int32
	SetWatermarkCustomText(v string) *ModifyCenterPolicyRequest
	GetWatermarkCustomText() *string
	SetWatermarkDegree(v float64) *ModifyCenterPolicyRequest
	GetWatermarkDegree() *float64
	SetWatermarkFontSize(v int32) *ModifyCenterPolicyRequest
	GetWatermarkFontSize() *int32
	SetWatermarkFontStyle(v string) *ModifyCenterPolicyRequest
	GetWatermarkFontStyle() *string
	SetWatermarkPower(v string) *ModifyCenterPolicyRequest
	GetWatermarkPower() *string
	SetWatermarkRowAmount(v int32) *ModifyCenterPolicyRequest
	GetWatermarkRowAmount() *int32
	SetWatermarkSecurity(v string) *ModifyCenterPolicyRequest
	GetWatermarkSecurity() *string
	SetWatermarkShadow(v string) *ModifyCenterPolicyRequest
	GetWatermarkShadow() *string
	SetWatermarkTransparencyValue(v int32) *ModifyCenterPolicyRequest
	GetWatermarkTransparencyValue() *int32
	SetWatermarkType(v string) *ModifyCenterPolicyRequest
	GetWatermarkType() *string
	SetWuyingKeeper(v string) *ModifyCenterPolicyRequest
	GetWuyingKeeper() *string
	SetWyAssistant(v string) *ModifyCenterPolicyRequest
	GetWyAssistant() *string
}

type ModifyCenterPolicyRequest struct {
	AcademicProxy *string `json:"AcademicProxy,omitempty" xml:"AcademicProxy,omitempty"`
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
	// on
	AppContentProtection *string `json:"AppContentProtection,omitempty" xml:"AppContentProtection,omitempty"`
	// The client IP address whitelists that you want to add.
	AuthorizeAccessPolicyRule []*ModifyCenterPolicyRequestAuthorizeAccessPolicyRule `json:"AuthorizeAccessPolicyRule,omitempty" xml:"AuthorizeAccessPolicyRule,omitempty" type:"Repeated"`
	// The security group rules.
	AuthorizeSecurityPolicyRule []*ModifyCenterPolicyRequestAuthorizeSecurityPolicyRule `json:"AuthorizeSecurityPolicyRule,omitempty" xml:"AuthorizeSecurityPolicyRule,omitempty" type:"Repeated"`
	// example:
	//
	// off
	AutoReconnect   *string `json:"AutoReconnect,omitempty" xml:"AutoReconnect,omitempty"`
	BusinessChannel *string `json:"BusinessChannel,omitempty" xml:"BusinessChannel,omitempty"`
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
	// on
	CameraRedirect       *string `json:"CameraRedirect,omitempty" xml:"CameraRedirect,omitempty"`
	ClientControlMenu    *string `json:"ClientControlMenu,omitempty" xml:"ClientControlMenu,omitempty"`
	ClientCreateSnapshot *string `json:"ClientCreateSnapshot,omitempty" xml:"ClientCreateSnapshot,omitempty"`
	// The types of Alibaba Cloud Workspace clients that end users can use to connect to cloud computers.
	ClientType []*ModifyCenterPolicyRequestClientType `json:"ClientType,omitempty" xml:"ClientType,omitempty" type:"Repeated"`
	// The read/write permissions on the clipboard.
	//
	// Valid values:
	//
	// 	- read: specifies one-way transfer. You can copy files only from on-premises devices to cloud computers.
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
	ClipboardGraineds []*ModifyCenterPolicyRequestClipboardGraineds `json:"ClipboardGraineds,omitempty" xml:"ClipboardGraineds,omitempty" type:"Repeated"`
	ClipboardScope    *string                                       `json:"ClipboardScope,omitempty" xml:"ClipboardScope,omitempty"`
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
	// 50
	CpuDownGradeDuration *int32  `json:"CpuDownGradeDuration,omitempty" xml:"CpuDownGradeDuration,omitempty"`
	CpuOverload          *string `json:"CpuOverload,omitempty" xml:"CpuOverload,omitempty"`
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
	// 30
	CpuSampleDuration *int32 `json:"CpuSampleDuration,omitempty" xml:"CpuSampleDuration,omitempty"`
	// The single-CPU usage. Valid values: 70 to 100. Unit: %.
	//
	// example:
	//
	// 80
	CpuSingleRateLimit *int32 `json:"CpuSingleRateLimit,omitempty" xml:"CpuSingleRateLimit,omitempty"`
	// Specifies whether to display the peripheral connection prompt.
	//
	// Valid values:
	//
	// 	- off: doesn\\"t display the peripheral connection prompt.
	//
	// 	- on: displays the peripheral connection prompt.
	//
	// example:
	//
	// off
	DeviceConnectHint *string `json:"DeviceConnectHint,omitempty" xml:"DeviceConnectHint,omitempty"`
	// The device redirection rules.
	DeviceRedirects []*ModifyCenterPolicyRequestDeviceRedirects `json:"DeviceRedirects,omitempty" xml:"DeviceRedirects,omitempty" type:"Repeated"`
	// The custom peripheral rules.
	DeviceRules []*ModifyCenterPolicyRequestDeviceRules `json:"DeviceRules,omitempty" xml:"DeviceRules,omitempty" type:"Repeated"`
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
	// 120
	DisconnectKeepSessionTime *int32  `json:"DisconnectKeepSessionTime,omitempty" xml:"DisconnectKeepSessionTime,omitempty"`
	DiskOverload              *string `json:"DiskOverload,omitempty" xml:"DiskOverload,omitempty"`
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
	DomainResolveRule []*ModifyCenterPolicyRequestDomainResolveRule `json:"DomainResolveRule,omitempty" xml:"DomainResolveRule,omitempty" type:"Repeated"`
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
	// Specifies whether to enforce a bandwidth limit for sessions.
	//
	// Valid values:
	//
	// 	- off: doesn\\"t enforce a bandwidth limit for sessions.
	//
	// 	- on: enforces a bandwidth limit for sessions.
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
	// Specifies whether to allow end users from the same office network to share cloud computers.
	//
	// Valid values:
	//
	// 	- off: doesn\\"t allow end users from the same office network to share cloud computers.
	//
	// 	- on: allows end users from the same office network to share cloud computers.
	//
	// example:
	//
	// off
	EndUserGroupCoordinate *string `json:"EndUserGroupCoordinate,omitempty" xml:"EndUserGroupCoordinate,omitempty"`
	ExternalDrive          *string `json:"ExternalDrive,omitempty" xml:"ExternalDrive,omitempty"`
	// Specifies whether to enable file transfer.
	//
	// Valid values:
	//
	// 	- off: enables file transfer.
	//
	// 	- on: disables file transfer.
	//
	// example:
	//
	// off
	FileMigrate               *string `json:"FileMigrate,omitempty" xml:"FileMigrate,omitempty"`
	FileTransferAddress       *string `json:"FileTransferAddress,omitempty" xml:"FileTransferAddress,omitempty"`
	FileTransferSpeed         *string `json:"FileTransferSpeed,omitempty" xml:"FileTransferSpeed,omitempty"`
	FileTransferSpeedLocation *string `json:"FileTransferSpeedLocation,omitempty" xml:"FileTransferSpeedLocation,omitempty"`
	// Specifies whether to enable Image Quality Control. This feature is highly recommended for professional design scenarios where performance and user experience are critical.
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
	// The file transfer policy on the web client.
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
	// The network communication protocol.
	//
	// Valid values:
	//
	// 	- tcp: TCP is used when UDP/AST is restricted.
	//
	// 	- rtc: AST is used for high-frequency audio and video streaming.
	//
	// 	- auto: UTO enables automatic switch between AST and UDP modes based on desktop content.
	//
	// 	- both: UDP is ideal for office and HD graphic design use.
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
	MemoryDownGradeDuration *int32  `json:"MemoryDownGradeDuration,omitempty" xml:"MemoryDownGradeDuration,omitempty"`
	MemoryOverload          *string `json:"MemoryOverload,omitempty" xml:"MemoryOverload,omitempty"`
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
	// 	- off: doesn\\"t display the Restart button in the DesktopAssistant menu when end users connect to cloud computers from Android clients.
	//
	// 	- on: displays the Restart button in the DesktopAssistant menu when end users connect to cloud computers from Android clients.
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
	// 	- off: doesn\\"t display the Stop button in the DesktopAssistant menu when end users connect to cloud computers from Android clients.
	//
	// 	- on: displays the Stop button in the DesktopAssistant menu when end users connect to cloud computers from Android clients.
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
	ModelLibrary      *string `json:"ModelLibrary,omitempty" xml:"ModelLibrary,omitempty"`
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
	// on
	NetRedirect *string `json:"NetRedirect,omitempty" xml:"NetRedirect,omitempty"`
	// The network redirection rules.
	//
	// >  This parameter is in private preview and only available to specific users.
	NetRedirectRule []*ModifyCenterPolicyRequestNetRedirectRule `json:"NetRedirectRule,omitempty" xml:"NetRedirectRule,omitempty" type:"Repeated"`
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
	// The cloud computer policy ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pg-53iyi2aar0nd6****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	PortProxy     *string `json:"PortProxy,omitempty" xml:"PortProxy,omitempty"`
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
	RecordEventLevels    []*ModifyCenterPolicyRequestRecordEventLevels `json:"RecordEventLevels,omitempty" xml:"RecordEventLevels,omitempty" type:"Repeated"`
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
	// Specifies whether to record audio files generated by cloud computers.
	//
	// Valid values:
	//
	// 	- off: doesn\\"t record audio files generated by cloud computers.
	//
	// 	- on: records audio files generated by cloud computers.
	//
	// example:
	//
	// on
	RecordingAudio *string `json:"RecordingAudio,omitempty" xml:"RecordingAudio,omitempty"`
	// The length of the screen recording file (in minutes). Screen recordings are split based on the specified duration and uploaded to Object Storage Service (OSS) buckets. If a file reaches 300 MB, the system prioritizes rolling updates for that file. Valid values: 10 to 60.
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
	// 5
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
	// The region ID. Set the value to `cn-shanghai`.
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
	// Valid values:
	//
	// 	- off: disables the reset setting.
	//
	// 	- on: enables the reset setting.
	//
	// example:
	//
	// off
	ResetDesktop *string `json:"ResetDesktop,omitempty" xml:"ResetDesktop,omitempty"`
	// The height of the resolution. Unit: pixel. Valid values for cloud applications: 500 to 50000. Valid values for cloud computers: 480 to 4096.
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
	// The width of the resolution. Unit: pixel. Valid values for cloud applications: 500 to 50000. Valid values for cloud computers: 480 to 4096.
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
	// The client IP address whitelists that you want to delete.
	RevokeAccessPolicyRule []*ModifyCenterPolicyRequestRevokeAccessPolicyRule `json:"RevokeAccessPolicyRule,omitempty" xml:"RevokeAccessPolicyRule,omitempty" type:"Repeated"`
	// The security group rules that you want to delete.
	RevokeSecurityPolicyRule []*ModifyCenterPolicyRequestRevokeSecurityPolicyRule `json:"RevokeSecurityPolicyRule,omitempty" xml:"RevokeSecurityPolicyRule,omitempty" type:"Repeated"`
	SafeMenu                 *string                                              `json:"SafeMenu,omitempty" xml:"SafeMenu,omitempty"`
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
	// Specifies whether to enable the USB redirection feature.
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
	UsbRedirect *string `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty"`
	// The USB redirection rules.
	UsbSupplyRedirectRule []*ModifyCenterPolicyRequestUsbSupplyRedirectRule `json:"UsbSupplyRedirectRule,omitempty" xml:"UsbSupplyRedirectRule,omitempty" type:"Repeated"`
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
	// The peak bitrate allowed for video encoding. Unit: Kbit/s. Valid values: 1000 to 50000.
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
	// 	- off: doesn\\"t enable anti-screen capture for invisible watermarks.
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
	// 5
	WatermarkRowAmount *int32 `json:"WatermarkRowAmount,omitempty" xml:"WatermarkRowAmount,omitempty"`
	// Specifies whether to enable security priority for invisible watermarks.
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
	WatermarkSecurity *string `json:"WatermarkSecurity,omitempty" xml:"WatermarkSecurity,omitempty"`
	WatermarkShadow   *string `json:"WatermarkShadow,omitempty" xml:"WatermarkShadow,omitempty"`
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
	// EndUserId
	WatermarkType *string `json:"WatermarkType,omitempty" xml:"WatermarkType,omitempty"`
	// Specifies whether to enable Cloud Computer Manager.
	//
	// Valid values:
	//
	// 	- off: disables Cloud Computer Manager.
	//
	// 	- on: enables Cloud Computer Manager.
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

func (s ModifyCenterPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCenterPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyCenterPolicyRequest) GetAcademicProxy() *string {
	return s.AcademicProxy
}

func (s *ModifyCenterPolicyRequest) GetAdminAccess() *string {
	return s.AdminAccess
}

func (s *ModifyCenterPolicyRequest) GetAppContentProtection() *string {
	return s.AppContentProtection
}

func (s *ModifyCenterPolicyRequest) GetAuthorizeAccessPolicyRule() []*ModifyCenterPolicyRequestAuthorizeAccessPolicyRule {
	return s.AuthorizeAccessPolicyRule
}

func (s *ModifyCenterPolicyRequest) GetAuthorizeSecurityPolicyRule() []*ModifyCenterPolicyRequestAuthorizeSecurityPolicyRule {
	return s.AuthorizeSecurityPolicyRule
}

func (s *ModifyCenterPolicyRequest) GetAutoReconnect() *string {
	return s.AutoReconnect
}

func (s *ModifyCenterPolicyRequest) GetBusinessChannel() *string {
	return s.BusinessChannel
}

func (s *ModifyCenterPolicyRequest) GetBusinessType() *int32 {
	return s.BusinessType
}

func (s *ModifyCenterPolicyRequest) GetCameraRedirect() *string {
	return s.CameraRedirect
}

func (s *ModifyCenterPolicyRequest) GetClientControlMenu() *string {
	return s.ClientControlMenu
}

func (s *ModifyCenterPolicyRequest) GetClientCreateSnapshot() *string {
	return s.ClientCreateSnapshot
}

func (s *ModifyCenterPolicyRequest) GetClientType() []*ModifyCenterPolicyRequestClientType {
	return s.ClientType
}

func (s *ModifyCenterPolicyRequest) GetClipboard() *string {
	return s.Clipboard
}

func (s *ModifyCenterPolicyRequest) GetClipboardGraineds() []*ModifyCenterPolicyRequestClipboardGraineds {
	return s.ClipboardGraineds
}

func (s *ModifyCenterPolicyRequest) GetClipboardScope() *string {
	return s.ClipboardScope
}

func (s *ModifyCenterPolicyRequest) GetColorEnhancement() *string {
	return s.ColorEnhancement
}

func (s *ModifyCenterPolicyRequest) GetCpdDriveClipboard() *string {
	return s.CpdDriveClipboard
}

func (s *ModifyCenterPolicyRequest) GetCpuDownGradeDuration() *int32 {
	return s.CpuDownGradeDuration
}

func (s *ModifyCenterPolicyRequest) GetCpuOverload() *string {
	return s.CpuOverload
}

func (s *ModifyCenterPolicyRequest) GetCpuProcessors() []*string {
	return s.CpuProcessors
}

func (s *ModifyCenterPolicyRequest) GetCpuProtectedMode() *string {
	return s.CpuProtectedMode
}

func (s *ModifyCenterPolicyRequest) GetCpuRateLimit() *int32 {
	return s.CpuRateLimit
}

func (s *ModifyCenterPolicyRequest) GetCpuSampleDuration() *int32 {
	return s.CpuSampleDuration
}

func (s *ModifyCenterPolicyRequest) GetCpuSingleRateLimit() *int32 {
	return s.CpuSingleRateLimit
}

func (s *ModifyCenterPolicyRequest) GetDeviceConnectHint() *string {
	return s.DeviceConnectHint
}

func (s *ModifyCenterPolicyRequest) GetDeviceRedirects() []*ModifyCenterPolicyRequestDeviceRedirects {
	return s.DeviceRedirects
}

func (s *ModifyCenterPolicyRequest) GetDeviceRules() []*ModifyCenterPolicyRequestDeviceRules {
	return s.DeviceRules
}

func (s *ModifyCenterPolicyRequest) GetDisconnectKeepSession() *string {
	return s.DisconnectKeepSession
}

func (s *ModifyCenterPolicyRequest) GetDisconnectKeepSessionTime() *int32 {
	return s.DisconnectKeepSessionTime
}

func (s *ModifyCenterPolicyRequest) GetDiskOverload() *string {
	return s.DiskOverload
}

func (s *ModifyCenterPolicyRequest) GetDisplayMode() *string {
	return s.DisplayMode
}

func (s *ModifyCenterPolicyRequest) GetDomainResolveRule() []*ModifyCenterPolicyRequestDomainResolveRule {
	return s.DomainResolveRule
}

func (s *ModifyCenterPolicyRequest) GetDomainResolveRuleType() *string {
	return s.DomainResolveRuleType
}

func (s *ModifyCenterPolicyRequest) GetEnableSessionRateLimiting() *string {
	return s.EnableSessionRateLimiting
}

func (s *ModifyCenterPolicyRequest) GetEndUserApplyAdminCoordinate() *string {
	return s.EndUserApplyAdminCoordinate
}

func (s *ModifyCenterPolicyRequest) GetEndUserGroupCoordinate() *string {
	return s.EndUserGroupCoordinate
}

func (s *ModifyCenterPolicyRequest) GetExternalDrive() *string {
	return s.ExternalDrive
}

func (s *ModifyCenterPolicyRequest) GetFileMigrate() *string {
	return s.FileMigrate
}

func (s *ModifyCenterPolicyRequest) GetFileTransferAddress() *string {
	return s.FileTransferAddress
}

func (s *ModifyCenterPolicyRequest) GetFileTransferSpeed() *string {
	return s.FileTransferSpeed
}

func (s *ModifyCenterPolicyRequest) GetFileTransferSpeedLocation() *string {
	return s.FileTransferSpeedLocation
}

func (s *ModifyCenterPolicyRequest) GetGpuAcceleration() *string {
	return s.GpuAcceleration
}

func (s *ModifyCenterPolicyRequest) GetHtml5FileTransfer() *string {
	return s.Html5FileTransfer
}

func (s *ModifyCenterPolicyRequest) GetInternetCommunicationProtocol() *string {
	return s.InternetCommunicationProtocol
}

func (s *ModifyCenterPolicyRequest) GetInternetPrinter() *string {
	return s.InternetPrinter
}

func (s *ModifyCenterPolicyRequest) GetLocalDrive() *string {
	return s.LocalDrive
}

func (s *ModifyCenterPolicyRequest) GetMaxReconnectTime() *int32 {
	return s.MaxReconnectTime
}

func (s *ModifyCenterPolicyRequest) GetMemoryDownGradeDuration() *int32 {
	return s.MemoryDownGradeDuration
}

func (s *ModifyCenterPolicyRequest) GetMemoryOverload() *string {
	return s.MemoryOverload
}

func (s *ModifyCenterPolicyRequest) GetMemoryProcessors() []*string {
	return s.MemoryProcessors
}

func (s *ModifyCenterPolicyRequest) GetMemoryProtectedMode() *string {
	return s.MemoryProtectedMode
}

func (s *ModifyCenterPolicyRequest) GetMemoryRateLimit() *int32 {
	return s.MemoryRateLimit
}

func (s *ModifyCenterPolicyRequest) GetMemorySampleDuration() *int32 {
	return s.MemorySampleDuration
}

func (s *ModifyCenterPolicyRequest) GetMemorySingleRateLimit() *int32 {
	return s.MemorySingleRateLimit
}

func (s *ModifyCenterPolicyRequest) GetMobileRestart() *string {
	return s.MobileRestart
}

func (s *ModifyCenterPolicyRequest) GetMobileSafeMenu() *string {
	return s.MobileSafeMenu
}

func (s *ModifyCenterPolicyRequest) GetMobileShutdown() *string {
	return s.MobileShutdown
}

func (s *ModifyCenterPolicyRequest) GetMobileWuyingKeeper() *string {
	return s.MobileWuyingKeeper
}

func (s *ModifyCenterPolicyRequest) GetMobileWyAssistant() *string {
	return s.MobileWyAssistant
}

func (s *ModifyCenterPolicyRequest) GetModelLibrary() *string {
	return s.ModelLibrary
}

func (s *ModifyCenterPolicyRequest) GetName() *string {
	return s.Name
}

func (s *ModifyCenterPolicyRequest) GetNetRedirect() *string {
	return s.NetRedirect
}

func (s *ModifyCenterPolicyRequest) GetNetRedirectRule() []*ModifyCenterPolicyRequestNetRedirectRule {
	return s.NetRedirectRule
}

func (s *ModifyCenterPolicyRequest) GetNoOperationDisconnect() *string {
	return s.NoOperationDisconnect
}

func (s *ModifyCenterPolicyRequest) GetNoOperationDisconnectTime() *int32 {
	return s.NoOperationDisconnectTime
}

func (s *ModifyCenterPolicyRequest) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *ModifyCenterPolicyRequest) GetPortProxy() *string {
	return s.PortProxy
}

func (s *ModifyCenterPolicyRequest) GetPrinterRedirect() *string {
	return s.PrinterRedirect
}

func (s *ModifyCenterPolicyRequest) GetQualityEnhancement() *string {
	return s.QualityEnhancement
}

func (s *ModifyCenterPolicyRequest) GetRecordEventDuration() *int32 {
	return s.RecordEventDuration
}

func (s *ModifyCenterPolicyRequest) GetRecordEventFileExts() []*string {
	return s.RecordEventFileExts
}

func (s *ModifyCenterPolicyRequest) GetRecordEventFilePaths() []*string {
	return s.RecordEventFilePaths
}

func (s *ModifyCenterPolicyRequest) GetRecordEventLevels() []*ModifyCenterPolicyRequestRecordEventLevels {
	return s.RecordEventLevels
}

func (s *ModifyCenterPolicyRequest) GetRecordEventRegisters() []*string {
	return s.RecordEventRegisters
}

func (s *ModifyCenterPolicyRequest) GetRecordEvents() []*string {
	return s.RecordEvents
}

func (s *ModifyCenterPolicyRequest) GetRecording() *string {
	return s.Recording
}

func (s *ModifyCenterPolicyRequest) GetRecordingAudio() *string {
	return s.RecordingAudio
}

func (s *ModifyCenterPolicyRequest) GetRecordingDuration() *int32 {
	return s.RecordingDuration
}

func (s *ModifyCenterPolicyRequest) GetRecordingEndTime() *string {
	return s.RecordingEndTime
}

func (s *ModifyCenterPolicyRequest) GetRecordingExpires() *int32 {
	return s.RecordingExpires
}

func (s *ModifyCenterPolicyRequest) GetRecordingFps() *string {
	return s.RecordingFps
}

func (s *ModifyCenterPolicyRequest) GetRecordingStartTime() *string {
	return s.RecordingStartTime
}

func (s *ModifyCenterPolicyRequest) GetRecordingUserNotify() *string {
	return s.RecordingUserNotify
}

func (s *ModifyCenterPolicyRequest) GetRecordingUserNotifyMessage() *string {
	return s.RecordingUserNotifyMessage
}

func (s *ModifyCenterPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyCenterPolicyRequest) GetRemoteCoordinate() *string {
	return s.RemoteCoordinate
}

func (s *ModifyCenterPolicyRequest) GetResetDesktop() *string {
	return s.ResetDesktop
}

func (s *ModifyCenterPolicyRequest) GetResolutionHeight() *int32 {
	return s.ResolutionHeight
}

func (s *ModifyCenterPolicyRequest) GetResolutionModel() *string {
	return s.ResolutionModel
}

func (s *ModifyCenterPolicyRequest) GetResolutionWidth() *int32 {
	return s.ResolutionWidth
}

func (s *ModifyCenterPolicyRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ModifyCenterPolicyRequest) GetRevokeAccessPolicyRule() []*ModifyCenterPolicyRequestRevokeAccessPolicyRule {
	return s.RevokeAccessPolicyRule
}

func (s *ModifyCenterPolicyRequest) GetRevokeSecurityPolicyRule() []*ModifyCenterPolicyRequestRevokeSecurityPolicyRule {
	return s.RevokeSecurityPolicyRule
}

func (s *ModifyCenterPolicyRequest) GetSafeMenu() *string {
	return s.SafeMenu
}

func (s *ModifyCenterPolicyRequest) GetScope() *string {
	return s.Scope
}

func (s *ModifyCenterPolicyRequest) GetScopeValue() []*string {
	return s.ScopeValue
}

func (s *ModifyCenterPolicyRequest) GetScreenDisplayMode() *string {
	return s.ScreenDisplayMode
}

func (s *ModifyCenterPolicyRequest) GetSessionMaxRateKbps() *int32 {
	return s.SessionMaxRateKbps
}

func (s *ModifyCenterPolicyRequest) GetSmoothEnhancement() *string {
	return s.SmoothEnhancement
}

func (s *ModifyCenterPolicyRequest) GetStatusMonitor() *string {
	return s.StatusMonitor
}

func (s *ModifyCenterPolicyRequest) GetStreamingMode() *string {
	return s.StreamingMode
}

func (s *ModifyCenterPolicyRequest) GetTargetFps() *int32 {
	return s.TargetFps
}

func (s *ModifyCenterPolicyRequest) GetTaskbar() *string {
	return s.Taskbar
}

func (s *ModifyCenterPolicyRequest) GetUsbRedirect() *string {
	return s.UsbRedirect
}

func (s *ModifyCenterPolicyRequest) GetUsbSupplyRedirectRule() []*ModifyCenterPolicyRequestUsbSupplyRedirectRule {
	return s.UsbSupplyRedirectRule
}

func (s *ModifyCenterPolicyRequest) GetUseTime() *string {
	return s.UseTime
}

func (s *ModifyCenterPolicyRequest) GetVideoEncAvgKbps() *int32 {
	return s.VideoEncAvgKbps
}

func (s *ModifyCenterPolicyRequest) GetVideoEncMaxQP() *int32 {
	return s.VideoEncMaxQP
}

func (s *ModifyCenterPolicyRequest) GetVideoEncMinQP() *int32 {
	return s.VideoEncMinQP
}

func (s *ModifyCenterPolicyRequest) GetVideoEncPeakKbps() *int32 {
	return s.VideoEncPeakKbps
}

func (s *ModifyCenterPolicyRequest) GetVideoEncPolicy() *string {
	return s.VideoEncPolicy
}

func (s *ModifyCenterPolicyRequest) GetVideoRedirect() *string {
	return s.VideoRedirect
}

func (s *ModifyCenterPolicyRequest) GetVisualQuality() *string {
	return s.VisualQuality
}

func (s *ModifyCenterPolicyRequest) GetWatermark() *string {
	return s.Watermark
}

func (s *ModifyCenterPolicyRequest) GetWatermarkAntiCam() *string {
	return s.WatermarkAntiCam
}

func (s *ModifyCenterPolicyRequest) GetWatermarkColor() *int32 {
	return s.WatermarkColor
}

func (s *ModifyCenterPolicyRequest) GetWatermarkColumnAmount() *int32 {
	return s.WatermarkColumnAmount
}

func (s *ModifyCenterPolicyRequest) GetWatermarkCustomText() *string {
	return s.WatermarkCustomText
}

func (s *ModifyCenterPolicyRequest) GetWatermarkDegree() *float64 {
	return s.WatermarkDegree
}

func (s *ModifyCenterPolicyRequest) GetWatermarkFontSize() *int32 {
	return s.WatermarkFontSize
}

func (s *ModifyCenterPolicyRequest) GetWatermarkFontStyle() *string {
	return s.WatermarkFontStyle
}

func (s *ModifyCenterPolicyRequest) GetWatermarkPower() *string {
	return s.WatermarkPower
}

func (s *ModifyCenterPolicyRequest) GetWatermarkRowAmount() *int32 {
	return s.WatermarkRowAmount
}

func (s *ModifyCenterPolicyRequest) GetWatermarkSecurity() *string {
	return s.WatermarkSecurity
}

func (s *ModifyCenterPolicyRequest) GetWatermarkShadow() *string {
	return s.WatermarkShadow
}

func (s *ModifyCenterPolicyRequest) GetWatermarkTransparencyValue() *int32 {
	return s.WatermarkTransparencyValue
}

func (s *ModifyCenterPolicyRequest) GetWatermarkType() *string {
	return s.WatermarkType
}

func (s *ModifyCenterPolicyRequest) GetWuyingKeeper() *string {
	return s.WuyingKeeper
}

func (s *ModifyCenterPolicyRequest) GetWyAssistant() *string {
	return s.WyAssistant
}

func (s *ModifyCenterPolicyRequest) SetAcademicProxy(v string) *ModifyCenterPolicyRequest {
	s.AcademicProxy = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetAdminAccess(v string) *ModifyCenterPolicyRequest {
	s.AdminAccess = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetAppContentProtection(v string) *ModifyCenterPolicyRequest {
	s.AppContentProtection = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetAuthorizeAccessPolicyRule(v []*ModifyCenterPolicyRequestAuthorizeAccessPolicyRule) *ModifyCenterPolicyRequest {
	s.AuthorizeAccessPolicyRule = v
	return s
}

func (s *ModifyCenterPolicyRequest) SetAuthorizeSecurityPolicyRule(v []*ModifyCenterPolicyRequestAuthorizeSecurityPolicyRule) *ModifyCenterPolicyRequest {
	s.AuthorizeSecurityPolicyRule = v
	return s
}

func (s *ModifyCenterPolicyRequest) SetAutoReconnect(v string) *ModifyCenterPolicyRequest {
	s.AutoReconnect = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetBusinessChannel(v string) *ModifyCenterPolicyRequest {
	s.BusinessChannel = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetBusinessType(v int32) *ModifyCenterPolicyRequest {
	s.BusinessType = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetCameraRedirect(v string) *ModifyCenterPolicyRequest {
	s.CameraRedirect = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetClientControlMenu(v string) *ModifyCenterPolicyRequest {
	s.ClientControlMenu = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetClientCreateSnapshot(v string) *ModifyCenterPolicyRequest {
	s.ClientCreateSnapshot = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetClientType(v []*ModifyCenterPolicyRequestClientType) *ModifyCenterPolicyRequest {
	s.ClientType = v
	return s
}

func (s *ModifyCenterPolicyRequest) SetClipboard(v string) *ModifyCenterPolicyRequest {
	s.Clipboard = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetClipboardGraineds(v []*ModifyCenterPolicyRequestClipboardGraineds) *ModifyCenterPolicyRequest {
	s.ClipboardGraineds = v
	return s
}

func (s *ModifyCenterPolicyRequest) SetClipboardScope(v string) *ModifyCenterPolicyRequest {
	s.ClipboardScope = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetColorEnhancement(v string) *ModifyCenterPolicyRequest {
	s.ColorEnhancement = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetCpdDriveClipboard(v string) *ModifyCenterPolicyRequest {
	s.CpdDriveClipboard = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetCpuDownGradeDuration(v int32) *ModifyCenterPolicyRequest {
	s.CpuDownGradeDuration = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetCpuOverload(v string) *ModifyCenterPolicyRequest {
	s.CpuOverload = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetCpuProcessors(v []*string) *ModifyCenterPolicyRequest {
	s.CpuProcessors = v
	return s
}

func (s *ModifyCenterPolicyRequest) SetCpuProtectedMode(v string) *ModifyCenterPolicyRequest {
	s.CpuProtectedMode = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetCpuRateLimit(v int32) *ModifyCenterPolicyRequest {
	s.CpuRateLimit = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetCpuSampleDuration(v int32) *ModifyCenterPolicyRequest {
	s.CpuSampleDuration = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetCpuSingleRateLimit(v int32) *ModifyCenterPolicyRequest {
	s.CpuSingleRateLimit = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetDeviceConnectHint(v string) *ModifyCenterPolicyRequest {
	s.DeviceConnectHint = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetDeviceRedirects(v []*ModifyCenterPolicyRequestDeviceRedirects) *ModifyCenterPolicyRequest {
	s.DeviceRedirects = v
	return s
}

func (s *ModifyCenterPolicyRequest) SetDeviceRules(v []*ModifyCenterPolicyRequestDeviceRules) *ModifyCenterPolicyRequest {
	s.DeviceRules = v
	return s
}

func (s *ModifyCenterPolicyRequest) SetDisconnectKeepSession(v string) *ModifyCenterPolicyRequest {
	s.DisconnectKeepSession = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetDisconnectKeepSessionTime(v int32) *ModifyCenterPolicyRequest {
	s.DisconnectKeepSessionTime = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetDiskOverload(v string) *ModifyCenterPolicyRequest {
	s.DiskOverload = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetDisplayMode(v string) *ModifyCenterPolicyRequest {
	s.DisplayMode = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetDomainResolveRule(v []*ModifyCenterPolicyRequestDomainResolveRule) *ModifyCenterPolicyRequest {
	s.DomainResolveRule = v
	return s
}

func (s *ModifyCenterPolicyRequest) SetDomainResolveRuleType(v string) *ModifyCenterPolicyRequest {
	s.DomainResolveRuleType = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetEnableSessionRateLimiting(v string) *ModifyCenterPolicyRequest {
	s.EnableSessionRateLimiting = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetEndUserApplyAdminCoordinate(v string) *ModifyCenterPolicyRequest {
	s.EndUserApplyAdminCoordinate = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetEndUserGroupCoordinate(v string) *ModifyCenterPolicyRequest {
	s.EndUserGroupCoordinate = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetExternalDrive(v string) *ModifyCenterPolicyRequest {
	s.ExternalDrive = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetFileMigrate(v string) *ModifyCenterPolicyRequest {
	s.FileMigrate = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetFileTransferAddress(v string) *ModifyCenterPolicyRequest {
	s.FileTransferAddress = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetFileTransferSpeed(v string) *ModifyCenterPolicyRequest {
	s.FileTransferSpeed = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetFileTransferSpeedLocation(v string) *ModifyCenterPolicyRequest {
	s.FileTransferSpeedLocation = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetGpuAcceleration(v string) *ModifyCenterPolicyRequest {
	s.GpuAcceleration = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetHtml5FileTransfer(v string) *ModifyCenterPolicyRequest {
	s.Html5FileTransfer = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetInternetCommunicationProtocol(v string) *ModifyCenterPolicyRequest {
	s.InternetCommunicationProtocol = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetInternetPrinter(v string) *ModifyCenterPolicyRequest {
	s.InternetPrinter = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetLocalDrive(v string) *ModifyCenterPolicyRequest {
	s.LocalDrive = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetMaxReconnectTime(v int32) *ModifyCenterPolicyRequest {
	s.MaxReconnectTime = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetMemoryDownGradeDuration(v int32) *ModifyCenterPolicyRequest {
	s.MemoryDownGradeDuration = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetMemoryOverload(v string) *ModifyCenterPolicyRequest {
	s.MemoryOverload = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetMemoryProcessors(v []*string) *ModifyCenterPolicyRequest {
	s.MemoryProcessors = v
	return s
}

func (s *ModifyCenterPolicyRequest) SetMemoryProtectedMode(v string) *ModifyCenterPolicyRequest {
	s.MemoryProtectedMode = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetMemoryRateLimit(v int32) *ModifyCenterPolicyRequest {
	s.MemoryRateLimit = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetMemorySampleDuration(v int32) *ModifyCenterPolicyRequest {
	s.MemorySampleDuration = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetMemorySingleRateLimit(v int32) *ModifyCenterPolicyRequest {
	s.MemorySingleRateLimit = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetMobileRestart(v string) *ModifyCenterPolicyRequest {
	s.MobileRestart = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetMobileSafeMenu(v string) *ModifyCenterPolicyRequest {
	s.MobileSafeMenu = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetMobileShutdown(v string) *ModifyCenterPolicyRequest {
	s.MobileShutdown = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetMobileWuyingKeeper(v string) *ModifyCenterPolicyRequest {
	s.MobileWuyingKeeper = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetMobileWyAssistant(v string) *ModifyCenterPolicyRequest {
	s.MobileWyAssistant = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetModelLibrary(v string) *ModifyCenterPolicyRequest {
	s.ModelLibrary = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetName(v string) *ModifyCenterPolicyRequest {
	s.Name = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetNetRedirect(v string) *ModifyCenterPolicyRequest {
	s.NetRedirect = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetNetRedirectRule(v []*ModifyCenterPolicyRequestNetRedirectRule) *ModifyCenterPolicyRequest {
	s.NetRedirectRule = v
	return s
}

func (s *ModifyCenterPolicyRequest) SetNoOperationDisconnect(v string) *ModifyCenterPolicyRequest {
	s.NoOperationDisconnect = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetNoOperationDisconnectTime(v int32) *ModifyCenterPolicyRequest {
	s.NoOperationDisconnectTime = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetPolicyGroupId(v string) *ModifyCenterPolicyRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetPortProxy(v string) *ModifyCenterPolicyRequest {
	s.PortProxy = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetPrinterRedirect(v string) *ModifyCenterPolicyRequest {
	s.PrinterRedirect = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetQualityEnhancement(v string) *ModifyCenterPolicyRequest {
	s.QualityEnhancement = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetRecordEventDuration(v int32) *ModifyCenterPolicyRequest {
	s.RecordEventDuration = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetRecordEventFileExts(v []*string) *ModifyCenterPolicyRequest {
	s.RecordEventFileExts = v
	return s
}

func (s *ModifyCenterPolicyRequest) SetRecordEventFilePaths(v []*string) *ModifyCenterPolicyRequest {
	s.RecordEventFilePaths = v
	return s
}

func (s *ModifyCenterPolicyRequest) SetRecordEventLevels(v []*ModifyCenterPolicyRequestRecordEventLevels) *ModifyCenterPolicyRequest {
	s.RecordEventLevels = v
	return s
}

func (s *ModifyCenterPolicyRequest) SetRecordEventRegisters(v []*string) *ModifyCenterPolicyRequest {
	s.RecordEventRegisters = v
	return s
}

func (s *ModifyCenterPolicyRequest) SetRecordEvents(v []*string) *ModifyCenterPolicyRequest {
	s.RecordEvents = v
	return s
}

func (s *ModifyCenterPolicyRequest) SetRecording(v string) *ModifyCenterPolicyRequest {
	s.Recording = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetRecordingAudio(v string) *ModifyCenterPolicyRequest {
	s.RecordingAudio = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetRecordingDuration(v int32) *ModifyCenterPolicyRequest {
	s.RecordingDuration = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetRecordingEndTime(v string) *ModifyCenterPolicyRequest {
	s.RecordingEndTime = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetRecordingExpires(v int32) *ModifyCenterPolicyRequest {
	s.RecordingExpires = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetRecordingFps(v string) *ModifyCenterPolicyRequest {
	s.RecordingFps = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetRecordingStartTime(v string) *ModifyCenterPolicyRequest {
	s.RecordingStartTime = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetRecordingUserNotify(v string) *ModifyCenterPolicyRequest {
	s.RecordingUserNotify = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetRecordingUserNotifyMessage(v string) *ModifyCenterPolicyRequest {
	s.RecordingUserNotifyMessage = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetRegionId(v string) *ModifyCenterPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetRemoteCoordinate(v string) *ModifyCenterPolicyRequest {
	s.RemoteCoordinate = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetResetDesktop(v string) *ModifyCenterPolicyRequest {
	s.ResetDesktop = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetResolutionHeight(v int32) *ModifyCenterPolicyRequest {
	s.ResolutionHeight = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetResolutionModel(v string) *ModifyCenterPolicyRequest {
	s.ResolutionModel = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetResolutionWidth(v int32) *ModifyCenterPolicyRequest {
	s.ResolutionWidth = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetResourceType(v string) *ModifyCenterPolicyRequest {
	s.ResourceType = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetRevokeAccessPolicyRule(v []*ModifyCenterPolicyRequestRevokeAccessPolicyRule) *ModifyCenterPolicyRequest {
	s.RevokeAccessPolicyRule = v
	return s
}

func (s *ModifyCenterPolicyRequest) SetRevokeSecurityPolicyRule(v []*ModifyCenterPolicyRequestRevokeSecurityPolicyRule) *ModifyCenterPolicyRequest {
	s.RevokeSecurityPolicyRule = v
	return s
}

func (s *ModifyCenterPolicyRequest) SetSafeMenu(v string) *ModifyCenterPolicyRequest {
	s.SafeMenu = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetScope(v string) *ModifyCenterPolicyRequest {
	s.Scope = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetScopeValue(v []*string) *ModifyCenterPolicyRequest {
	s.ScopeValue = v
	return s
}

func (s *ModifyCenterPolicyRequest) SetScreenDisplayMode(v string) *ModifyCenterPolicyRequest {
	s.ScreenDisplayMode = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetSessionMaxRateKbps(v int32) *ModifyCenterPolicyRequest {
	s.SessionMaxRateKbps = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetSmoothEnhancement(v string) *ModifyCenterPolicyRequest {
	s.SmoothEnhancement = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetStatusMonitor(v string) *ModifyCenterPolicyRequest {
	s.StatusMonitor = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetStreamingMode(v string) *ModifyCenterPolicyRequest {
	s.StreamingMode = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetTargetFps(v int32) *ModifyCenterPolicyRequest {
	s.TargetFps = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetTaskbar(v string) *ModifyCenterPolicyRequest {
	s.Taskbar = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetUsbRedirect(v string) *ModifyCenterPolicyRequest {
	s.UsbRedirect = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetUsbSupplyRedirectRule(v []*ModifyCenterPolicyRequestUsbSupplyRedirectRule) *ModifyCenterPolicyRequest {
	s.UsbSupplyRedirectRule = v
	return s
}

func (s *ModifyCenterPolicyRequest) SetUseTime(v string) *ModifyCenterPolicyRequest {
	s.UseTime = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetVideoEncAvgKbps(v int32) *ModifyCenterPolicyRequest {
	s.VideoEncAvgKbps = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetVideoEncMaxQP(v int32) *ModifyCenterPolicyRequest {
	s.VideoEncMaxQP = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetVideoEncMinQP(v int32) *ModifyCenterPolicyRequest {
	s.VideoEncMinQP = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetVideoEncPeakKbps(v int32) *ModifyCenterPolicyRequest {
	s.VideoEncPeakKbps = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetVideoEncPolicy(v string) *ModifyCenterPolicyRequest {
	s.VideoEncPolicy = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetVideoRedirect(v string) *ModifyCenterPolicyRequest {
	s.VideoRedirect = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetVisualQuality(v string) *ModifyCenterPolicyRequest {
	s.VisualQuality = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetWatermark(v string) *ModifyCenterPolicyRequest {
	s.Watermark = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetWatermarkAntiCam(v string) *ModifyCenterPolicyRequest {
	s.WatermarkAntiCam = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetWatermarkColor(v int32) *ModifyCenterPolicyRequest {
	s.WatermarkColor = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetWatermarkColumnAmount(v int32) *ModifyCenterPolicyRequest {
	s.WatermarkColumnAmount = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetWatermarkCustomText(v string) *ModifyCenterPolicyRequest {
	s.WatermarkCustomText = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetWatermarkDegree(v float64) *ModifyCenterPolicyRequest {
	s.WatermarkDegree = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetWatermarkFontSize(v int32) *ModifyCenterPolicyRequest {
	s.WatermarkFontSize = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetWatermarkFontStyle(v string) *ModifyCenterPolicyRequest {
	s.WatermarkFontStyle = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetWatermarkPower(v string) *ModifyCenterPolicyRequest {
	s.WatermarkPower = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetWatermarkRowAmount(v int32) *ModifyCenterPolicyRequest {
	s.WatermarkRowAmount = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetWatermarkSecurity(v string) *ModifyCenterPolicyRequest {
	s.WatermarkSecurity = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetWatermarkShadow(v string) *ModifyCenterPolicyRequest {
	s.WatermarkShadow = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetWatermarkTransparencyValue(v int32) *ModifyCenterPolicyRequest {
	s.WatermarkTransparencyValue = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetWatermarkType(v string) *ModifyCenterPolicyRequest {
	s.WatermarkType = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetWuyingKeeper(v string) *ModifyCenterPolicyRequest {
	s.WuyingKeeper = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetWyAssistant(v string) *ModifyCenterPolicyRequest {
	s.WyAssistant = &v
	return s
}

func (s *ModifyCenterPolicyRequest) Validate() error {
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
	if s.ClipboardGraineds != nil {
		for _, item := range s.ClipboardGraineds {
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
	if s.NetRedirectRule != nil {
		for _, item := range s.NetRedirectRule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RecordEventLevels != nil {
		for _, item := range s.RecordEventLevels {
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

type ModifyCenterPolicyRequestAuthorizeAccessPolicyRule struct {
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

func (s ModifyCenterPolicyRequestAuthorizeAccessPolicyRule) String() string {
	return dara.Prettify(s)
}

func (s ModifyCenterPolicyRequestAuthorizeAccessPolicyRule) GoString() string {
	return s.String()
}

func (s *ModifyCenterPolicyRequestAuthorizeAccessPolicyRule) GetCidrIp() *string {
	return s.CidrIp
}

func (s *ModifyCenterPolicyRequestAuthorizeAccessPolicyRule) GetDescription() *string {
	return s.Description
}

func (s *ModifyCenterPolicyRequestAuthorizeAccessPolicyRule) SetCidrIp(v string) *ModifyCenterPolicyRequestAuthorizeAccessPolicyRule {
	s.CidrIp = &v
	return s
}

func (s *ModifyCenterPolicyRequestAuthorizeAccessPolicyRule) SetDescription(v string) *ModifyCenterPolicyRequestAuthorizeAccessPolicyRule {
	s.Description = &v
	return s
}

func (s *ModifyCenterPolicyRequestAuthorizeAccessPolicyRule) Validate() error {
	return dara.Validate(s)
}

type ModifyCenterPolicyRequestAuthorizeSecurityPolicyRule struct {
	// The object of the security group rule. Specify an IPv4 CIDR block.
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
	// The priority of the security group rule. A smaller value specifies a higher priority. Valid values: 1 to 60. Default value: 1.
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

func (s ModifyCenterPolicyRequestAuthorizeSecurityPolicyRule) String() string {
	return dara.Prettify(s)
}

func (s ModifyCenterPolicyRequestAuthorizeSecurityPolicyRule) GoString() string {
	return s.String()
}

func (s *ModifyCenterPolicyRequestAuthorizeSecurityPolicyRule) GetCidrIp() *string {
	return s.CidrIp
}

func (s *ModifyCenterPolicyRequestAuthorizeSecurityPolicyRule) GetDescription() *string {
	return s.Description
}

func (s *ModifyCenterPolicyRequestAuthorizeSecurityPolicyRule) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *ModifyCenterPolicyRequestAuthorizeSecurityPolicyRule) GetPolicy() *string {
	return s.Policy
}

func (s *ModifyCenterPolicyRequestAuthorizeSecurityPolicyRule) GetPortRange() *string {
	return s.PortRange
}

func (s *ModifyCenterPolicyRequestAuthorizeSecurityPolicyRule) GetPriority() *string {
	return s.Priority
}

func (s *ModifyCenterPolicyRequestAuthorizeSecurityPolicyRule) GetType() *string {
	return s.Type
}

func (s *ModifyCenterPolicyRequestAuthorizeSecurityPolicyRule) SetCidrIp(v string) *ModifyCenterPolicyRequestAuthorizeSecurityPolicyRule {
	s.CidrIp = &v
	return s
}

func (s *ModifyCenterPolicyRequestAuthorizeSecurityPolicyRule) SetDescription(v string) *ModifyCenterPolicyRequestAuthorizeSecurityPolicyRule {
	s.Description = &v
	return s
}

func (s *ModifyCenterPolicyRequestAuthorizeSecurityPolicyRule) SetIpProtocol(v string) *ModifyCenterPolicyRequestAuthorizeSecurityPolicyRule {
	s.IpProtocol = &v
	return s
}

func (s *ModifyCenterPolicyRequestAuthorizeSecurityPolicyRule) SetPolicy(v string) *ModifyCenterPolicyRequestAuthorizeSecurityPolicyRule {
	s.Policy = &v
	return s
}

func (s *ModifyCenterPolicyRequestAuthorizeSecurityPolicyRule) SetPortRange(v string) *ModifyCenterPolicyRequestAuthorizeSecurityPolicyRule {
	s.PortRange = &v
	return s
}

func (s *ModifyCenterPolicyRequestAuthorizeSecurityPolicyRule) SetPriority(v string) *ModifyCenterPolicyRequestAuthorizeSecurityPolicyRule {
	s.Priority = &v
	return s
}

func (s *ModifyCenterPolicyRequestAuthorizeSecurityPolicyRule) SetType(v string) *ModifyCenterPolicyRequestAuthorizeSecurityPolicyRule {
	s.Type = &v
	return s
}

func (s *ModifyCenterPolicyRequestAuthorizeSecurityPolicyRule) Validate() error {
	return dara.Validate(s)
}

type ModifyCenterPolicyRequestClientType struct {
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
	// android
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

func (s ModifyCenterPolicyRequestClientType) String() string {
	return dara.Prettify(s)
}

func (s ModifyCenterPolicyRequestClientType) GoString() string {
	return s.String()
}

func (s *ModifyCenterPolicyRequestClientType) GetClientType() *string {
	return s.ClientType
}

func (s *ModifyCenterPolicyRequestClientType) GetStatus() *string {
	return s.Status
}

func (s *ModifyCenterPolicyRequestClientType) SetClientType(v string) *ModifyCenterPolicyRequestClientType {
	s.ClientType = &v
	return s
}

func (s *ModifyCenterPolicyRequestClientType) SetStatus(v string) *ModifyCenterPolicyRequestClientType {
	s.Status = &v
	return s
}

func (s *ModifyCenterPolicyRequestClientType) Validate() error {
	return dara.Validate(s)
}

type ModifyCenterPolicyRequestClipboardGraineds struct {
	ClipboardSize *int32  `json:"ClipboardSize,omitempty" xml:"ClipboardSize,omitempty"`
	ClipboardType *string `json:"ClipboardType,omitempty" xml:"ClipboardType,omitempty"`
	GrainedType   *string `json:"GrainedType,omitempty" xml:"GrainedType,omitempty"`
}

func (s ModifyCenterPolicyRequestClipboardGraineds) String() string {
	return dara.Prettify(s)
}

func (s ModifyCenterPolicyRequestClipboardGraineds) GoString() string {
	return s.String()
}

func (s *ModifyCenterPolicyRequestClipboardGraineds) GetClipboardSize() *int32 {
	return s.ClipboardSize
}

func (s *ModifyCenterPolicyRequestClipboardGraineds) GetClipboardType() *string {
	return s.ClipboardType
}

func (s *ModifyCenterPolicyRequestClipboardGraineds) GetGrainedType() *string {
	return s.GrainedType
}

func (s *ModifyCenterPolicyRequestClipboardGraineds) SetClipboardSize(v int32) *ModifyCenterPolicyRequestClipboardGraineds {
	s.ClipboardSize = &v
	return s
}

func (s *ModifyCenterPolicyRequestClipboardGraineds) SetClipboardType(v string) *ModifyCenterPolicyRequestClipboardGraineds {
	s.ClipboardType = &v
	return s
}

func (s *ModifyCenterPolicyRequestClipboardGraineds) SetGrainedType(v string) *ModifyCenterPolicyRequestClipboardGraineds {
	s.GrainedType = &v
	return s
}

func (s *ModifyCenterPolicyRequestClipboardGraineds) Validate() error {
	return dara.Validate(s)
}

type ModifyCenterPolicyRequestDeviceRedirects struct {
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
	// 	- deviceRedirect: device redirection
	//
	// 	- usbRedirect: USB redirection.
	//
	// 	- off: any type of redirection.
	//
	// example:
	//
	// deviceRedirect
	RedirectType *string `json:"RedirectType,omitempty" xml:"RedirectType,omitempty"`
}

func (s ModifyCenterPolicyRequestDeviceRedirects) String() string {
	return dara.Prettify(s)
}

func (s ModifyCenterPolicyRequestDeviceRedirects) GoString() string {
	return s.String()
}

func (s *ModifyCenterPolicyRequestDeviceRedirects) GetDeviceType() *string {
	return s.DeviceType
}

func (s *ModifyCenterPolicyRequestDeviceRedirects) GetRedirectType() *string {
	return s.RedirectType
}

func (s *ModifyCenterPolicyRequestDeviceRedirects) SetDeviceType(v string) *ModifyCenterPolicyRequestDeviceRedirects {
	s.DeviceType = &v
	return s
}

func (s *ModifyCenterPolicyRequestDeviceRedirects) SetRedirectType(v string) *ModifyCenterPolicyRequestDeviceRedirects {
	s.RedirectType = &v
	return s
}

func (s *ModifyCenterPolicyRequestDeviceRedirects) Validate() error {
	return dara.Validate(s)
}

type ModifyCenterPolicyRequestDeviceRules struct {
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
	// 	- camera: web cameras.
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

func (s ModifyCenterPolicyRequestDeviceRules) String() string {
	return dara.Prettify(s)
}

func (s ModifyCenterPolicyRequestDeviceRules) GoString() string {
	return s.String()
}

func (s *ModifyCenterPolicyRequestDeviceRules) GetDeviceName() *string {
	return s.DeviceName
}

func (s *ModifyCenterPolicyRequestDeviceRules) GetDevicePid() *string {
	return s.DevicePid
}

func (s *ModifyCenterPolicyRequestDeviceRules) GetDeviceType() *string {
	return s.DeviceType
}

func (s *ModifyCenterPolicyRequestDeviceRules) GetDeviceVid() *string {
	return s.DeviceVid
}

func (s *ModifyCenterPolicyRequestDeviceRules) GetOptCommand() *string {
	return s.OptCommand
}

func (s *ModifyCenterPolicyRequestDeviceRules) GetPlatforms() *string {
	return s.Platforms
}

func (s *ModifyCenterPolicyRequestDeviceRules) GetRedirectType() *string {
	return s.RedirectType
}

func (s *ModifyCenterPolicyRequestDeviceRules) SetDeviceName(v string) *ModifyCenterPolicyRequestDeviceRules {
	s.DeviceName = &v
	return s
}

func (s *ModifyCenterPolicyRequestDeviceRules) SetDevicePid(v string) *ModifyCenterPolicyRequestDeviceRules {
	s.DevicePid = &v
	return s
}

func (s *ModifyCenterPolicyRequestDeviceRules) SetDeviceType(v string) *ModifyCenterPolicyRequestDeviceRules {
	s.DeviceType = &v
	return s
}

func (s *ModifyCenterPolicyRequestDeviceRules) SetDeviceVid(v string) *ModifyCenterPolicyRequestDeviceRules {
	s.DeviceVid = &v
	return s
}

func (s *ModifyCenterPolicyRequestDeviceRules) SetOptCommand(v string) *ModifyCenterPolicyRequestDeviceRules {
	s.OptCommand = &v
	return s
}

func (s *ModifyCenterPolicyRequestDeviceRules) SetPlatforms(v string) *ModifyCenterPolicyRequestDeviceRules {
	s.Platforms = &v
	return s
}

func (s *ModifyCenterPolicyRequestDeviceRules) SetRedirectType(v string) *ModifyCenterPolicyRequestDeviceRules {
	s.RedirectType = &v
	return s
}

func (s *ModifyCenterPolicyRequestDeviceRules) Validate() error {
	return dara.Validate(s)
}

type ModifyCenterPolicyRequestDomainResolveRule struct {
	// The policy description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The domain name.
	//
	// example:
	//
	// *.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Specifies whether to allow the domain name resolution rule.
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

func (s ModifyCenterPolicyRequestDomainResolveRule) String() string {
	return dara.Prettify(s)
}

func (s ModifyCenterPolicyRequestDomainResolveRule) GoString() string {
	return s.String()
}

func (s *ModifyCenterPolicyRequestDomainResolveRule) GetDescription() *string {
	return s.Description
}

func (s *ModifyCenterPolicyRequestDomainResolveRule) GetDomain() *string {
	return s.Domain
}

func (s *ModifyCenterPolicyRequestDomainResolveRule) GetPolicy() *string {
	return s.Policy
}

func (s *ModifyCenterPolicyRequestDomainResolveRule) SetDescription(v string) *ModifyCenterPolicyRequestDomainResolveRule {
	s.Description = &v
	return s
}

func (s *ModifyCenterPolicyRequestDomainResolveRule) SetDomain(v string) *ModifyCenterPolicyRequestDomainResolveRule {
	s.Domain = &v
	return s
}

func (s *ModifyCenterPolicyRequestDomainResolveRule) SetPolicy(v string) *ModifyCenterPolicyRequestDomainResolveRule {
	s.Policy = &v
	return s
}

func (s *ModifyCenterPolicyRequestDomainResolveRule) Validate() error {
	return dara.Validate(s)
}

type ModifyCenterPolicyRequestNetRedirectRule struct {
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
	// Allow
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

func (s ModifyCenterPolicyRequestNetRedirectRule) String() string {
	return dara.Prettify(s)
}

func (s ModifyCenterPolicyRequestNetRedirectRule) GoString() string {
	return s.String()
}

func (s *ModifyCenterPolicyRequestNetRedirectRule) GetDomain() *string {
	return s.Domain
}

func (s *ModifyCenterPolicyRequestNetRedirectRule) GetPolicy() *string {
	return s.Policy
}

func (s *ModifyCenterPolicyRequestNetRedirectRule) GetRuleType() *string {
	return s.RuleType
}

func (s *ModifyCenterPolicyRequestNetRedirectRule) SetDomain(v string) *ModifyCenterPolicyRequestNetRedirectRule {
	s.Domain = &v
	return s
}

func (s *ModifyCenterPolicyRequestNetRedirectRule) SetPolicy(v string) *ModifyCenterPolicyRequestNetRedirectRule {
	s.Policy = &v
	return s
}

func (s *ModifyCenterPolicyRequestNetRedirectRule) SetRuleType(v string) *ModifyCenterPolicyRequestNetRedirectRule {
	s.RuleType = &v
	return s
}

func (s *ModifyCenterPolicyRequestNetRedirectRule) Validate() error {
	return dara.Validate(s)
}

type ModifyCenterPolicyRequestRecordEventLevels struct {
	// example:
	//
	// HIGH
	EventLevel *string `json:"EventLevel,omitempty" xml:"EventLevel,omitempty"`
	// example:
	//
	// StartApplication
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
}

func (s ModifyCenterPolicyRequestRecordEventLevels) String() string {
	return dara.Prettify(s)
}

func (s ModifyCenterPolicyRequestRecordEventLevels) GoString() string {
	return s.String()
}

func (s *ModifyCenterPolicyRequestRecordEventLevels) GetEventLevel() *string {
	return s.EventLevel
}

func (s *ModifyCenterPolicyRequestRecordEventLevels) GetEventType() *string {
	return s.EventType
}

func (s *ModifyCenterPolicyRequestRecordEventLevels) SetEventLevel(v string) *ModifyCenterPolicyRequestRecordEventLevels {
	s.EventLevel = &v
	return s
}

func (s *ModifyCenterPolicyRequestRecordEventLevels) SetEventType(v string) *ModifyCenterPolicyRequestRecordEventLevels {
	s.EventType = &v
	return s
}

func (s *ModifyCenterPolicyRequestRecordEventLevels) Validate() error {
	return dara.Validate(s)
}

type ModifyCenterPolicyRequestRevokeAccessPolicyRule struct {
	// The client CIDR block that you want to delete. Specify an IPv4 CIDR block.
	//
	// example:
	//
	// 47.100.XX.XX/16
	CidrIp *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	// The description of the client IP address whitelist that you want to delete.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ModifyCenterPolicyRequestRevokeAccessPolicyRule) String() string {
	return dara.Prettify(s)
}

func (s ModifyCenterPolicyRequestRevokeAccessPolicyRule) GoString() string {
	return s.String()
}

func (s *ModifyCenterPolicyRequestRevokeAccessPolicyRule) GetCidrIp() *string {
	return s.CidrIp
}

func (s *ModifyCenterPolicyRequestRevokeAccessPolicyRule) GetDescription() *string {
	return s.Description
}

func (s *ModifyCenterPolicyRequestRevokeAccessPolicyRule) SetCidrIp(v string) *ModifyCenterPolicyRequestRevokeAccessPolicyRule {
	s.CidrIp = &v
	return s
}

func (s *ModifyCenterPolicyRequestRevokeAccessPolicyRule) SetDescription(v string) *ModifyCenterPolicyRequestRevokeAccessPolicyRule {
	s.Description = &v
	return s
}

func (s *ModifyCenterPolicyRequestRevokeAccessPolicyRule) Validate() error {
	return dara.Validate(s)
}

type ModifyCenterPolicyRequestRevokeSecurityPolicyRule struct {
	// The object of the security group rule that you want to delete. Specify an IPv4 CIDR block.
	//
	// example:
	//
	// 47.100.XX.XX/16
	CidrIp *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	// The description of the security group rule that you want to delete.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The protocol type of the security group rule that you want to delete.
	//
	// Valid values:
	//
	// 	- TCP: the TCP protocol.
	//
	// 	- UDP: the UDP protocol.
	//
	// 	- ALL: any type of protocol.
	//
	// 	- GRE: the GRE protocol.
	//
	// 	- ICMP: the ICMP for IPv4.
	//
	// example:
	//
	// TCP
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The authorization of the security group rule that you want to delete.
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
	// The port range of the security group rule that you want to delete. The value range of this parameter varies based on the value of IpProtocol.
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
	// The priority of the security group rule that you want to delete. A smaller value specifies a higher priority. Valid values: 1 to 60. Default value: 1.
	//
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The direction of the security group rule that you want to delete.
	//
	// Valid values:
	//
	// 	- outflow: outbound.
	//
	// 	- inflow: inbound.
	//
	// example:
	//
	// outflow
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyCenterPolicyRequestRevokeSecurityPolicyRule) String() string {
	return dara.Prettify(s)
}

func (s ModifyCenterPolicyRequestRevokeSecurityPolicyRule) GoString() string {
	return s.String()
}

func (s *ModifyCenterPolicyRequestRevokeSecurityPolicyRule) GetCidrIp() *string {
	return s.CidrIp
}

func (s *ModifyCenterPolicyRequestRevokeSecurityPolicyRule) GetDescription() *string {
	return s.Description
}

func (s *ModifyCenterPolicyRequestRevokeSecurityPolicyRule) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *ModifyCenterPolicyRequestRevokeSecurityPolicyRule) GetPolicy() *string {
	return s.Policy
}

func (s *ModifyCenterPolicyRequestRevokeSecurityPolicyRule) GetPortRange() *string {
	return s.PortRange
}

func (s *ModifyCenterPolicyRequestRevokeSecurityPolicyRule) GetPriority() *string {
	return s.Priority
}

func (s *ModifyCenterPolicyRequestRevokeSecurityPolicyRule) GetType() *string {
	return s.Type
}

func (s *ModifyCenterPolicyRequestRevokeSecurityPolicyRule) SetCidrIp(v string) *ModifyCenterPolicyRequestRevokeSecurityPolicyRule {
	s.CidrIp = &v
	return s
}

func (s *ModifyCenterPolicyRequestRevokeSecurityPolicyRule) SetDescription(v string) *ModifyCenterPolicyRequestRevokeSecurityPolicyRule {
	s.Description = &v
	return s
}

func (s *ModifyCenterPolicyRequestRevokeSecurityPolicyRule) SetIpProtocol(v string) *ModifyCenterPolicyRequestRevokeSecurityPolicyRule {
	s.IpProtocol = &v
	return s
}

func (s *ModifyCenterPolicyRequestRevokeSecurityPolicyRule) SetPolicy(v string) *ModifyCenterPolicyRequestRevokeSecurityPolicyRule {
	s.Policy = &v
	return s
}

func (s *ModifyCenterPolicyRequestRevokeSecurityPolicyRule) SetPortRange(v string) *ModifyCenterPolicyRequestRevokeSecurityPolicyRule {
	s.PortRange = &v
	return s
}

func (s *ModifyCenterPolicyRequestRevokeSecurityPolicyRule) SetPriority(v string) *ModifyCenterPolicyRequestRevokeSecurityPolicyRule {
	s.Priority = &v
	return s
}

func (s *ModifyCenterPolicyRequestRevokeSecurityPolicyRule) SetType(v string) *ModifyCenterPolicyRequestRevokeSecurityPolicyRule {
	s.Type = &v
	return s
}

func (s *ModifyCenterPolicyRequestRevokeSecurityPolicyRule) Validate() error {
	return dara.Validate(s)
}

type ModifyCenterPolicyRequestUsbSupplyRedirectRule struct {
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
	// 	- 1: enables USB redirection based on device manufacturers.
	//
	// example:
	//
	// 1
	UsbRuleType *string `json:"UsbRuleType,omitempty" xml:"UsbRuleType,omitempty"`
	// The vendor ID (VID). For more information, see [Valid USB Vendor IDs (VIDs)](https://www.usb.org/sites/default/files/vendor_ids032322.pdf_1.pdf).
	//
	// example:
	//
	// 04**
	VendorId *string `json:"VendorId,omitempty" xml:"VendorId,omitempty"`
}

func (s ModifyCenterPolicyRequestUsbSupplyRedirectRule) String() string {
	return dara.Prettify(s)
}

func (s ModifyCenterPolicyRequestUsbSupplyRedirectRule) GoString() string {
	return s.String()
}

func (s *ModifyCenterPolicyRequestUsbSupplyRedirectRule) GetDescription() *string {
	return s.Description
}

func (s *ModifyCenterPolicyRequestUsbSupplyRedirectRule) GetProductId() *string {
	return s.ProductId
}

func (s *ModifyCenterPolicyRequestUsbSupplyRedirectRule) GetUsbRedirectType() *string {
	return s.UsbRedirectType
}

func (s *ModifyCenterPolicyRequestUsbSupplyRedirectRule) GetUsbRuleType() *string {
	return s.UsbRuleType
}

func (s *ModifyCenterPolicyRequestUsbSupplyRedirectRule) GetVendorId() *string {
	return s.VendorId
}

func (s *ModifyCenterPolicyRequestUsbSupplyRedirectRule) SetDescription(v string) *ModifyCenterPolicyRequestUsbSupplyRedirectRule {
	s.Description = &v
	return s
}

func (s *ModifyCenterPolicyRequestUsbSupplyRedirectRule) SetProductId(v string) *ModifyCenterPolicyRequestUsbSupplyRedirectRule {
	s.ProductId = &v
	return s
}

func (s *ModifyCenterPolicyRequestUsbSupplyRedirectRule) SetUsbRedirectType(v string) *ModifyCenterPolicyRequestUsbSupplyRedirectRule {
	s.UsbRedirectType = &v
	return s
}

func (s *ModifyCenterPolicyRequestUsbSupplyRedirectRule) SetUsbRuleType(v string) *ModifyCenterPolicyRequestUsbSupplyRedirectRule {
	s.UsbRuleType = &v
	return s
}

func (s *ModifyCenterPolicyRequestUsbSupplyRedirectRule) SetVendorId(v string) *ModifyCenterPolicyRequestUsbSupplyRedirectRule {
	s.VendorId = &v
	return s
}

func (s *ModifyCenterPolicyRequestUsbSupplyRedirectRule) Validate() error {
	return dara.Validate(s)
}
