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
	SetAdminKeyboardOnFullScreen(v string) *ModifyCenterPolicyRequest
	GetAdminKeyboardOnFullScreen() *string
	SetAdminKeyboardOnWindows(v string) *ModifyCenterPolicyRequest
	GetAdminKeyboardOnWindows() *string
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
	SetDescription(v string) *ModifyCenterPolicyRequest
	GetDescription() *string
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
	SetFileTransferInSize(v string) *ModifyCenterPolicyRequest
	GetFileTransferInSize() *string
	SetFileTransferInUnit(v string) *ModifyCenterPolicyRequest
	GetFileTransferInUnit() *string
	SetFileTransferOutSize(v string) *ModifyCenterPolicyRequest
	GetFileTransferOutSize() *string
	SetFileTransferOutUnit(v string) *ModifyCenterPolicyRequest
	GetFileTransferOutUnit() *string
	SetFileTransferSizeLimit(v string) *ModifyCenterPolicyRequest
	GetFileTransferSizeLimit() *string
	SetFileTransferSpeed(v string) *ModifyCenterPolicyRequest
	GetFileTransferSpeed() *string
	SetFileTransferSpeedLocation(v string) *ModifyCenterPolicyRequest
	GetFileTransferSpeedLocation() *string
	SetGpuAcceleration(v string) *ModifyCenterPolicyRequest
	GetGpuAcceleration() *string
	SetHoverConfigMsg(v string) *ModifyCenterPolicyRequest
	GetHoverConfigMsg() *string
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
	SetMultiScreen(v string) *ModifyCenterPolicyRequest
	GetMultiScreen() *string
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
	SetPrinterAlert(v string) *ModifyCenterPolicyRequest
	GetPrinterAlert() *string
	SetPrinterAlertContent(v string) *ModifyCenterPolicyRequest
	GetPrinterAlertContent() *string
	SetPrinterAlertTitle(v string) *ModifyCenterPolicyRequest
	GetPrinterAlertTitle() *string
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
	SetResolutionDpi(v int32) *ModifyCenterPolicyRequest
	GetResolutionDpi() *int32
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
	SetThreeScreen(v string) *ModifyCenterPolicyRequest
	GetThreeScreen() *string
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
	// Specifies whether to enable the academic proxy feature. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	AcademicProxy *string `json:"AcademicProxy,omitempty" xml:"AcademicProxy,omitempty"`
	// Specifies whether the user has administrator permissions after logging on to the cloud desktop.
	//
	// > This feature is in invitational preview and is not publicly available.
	//
	// example:
	//
	// deny
	AdminAccess *string `json:"AdminAccess,omitempty" xml:"AdminAccess,omitempty"`
	// Specifies whether to enable administrator keyboard control in full-screen mode. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	AdminKeyboardOnFullScreen *string `json:"AdminKeyboardOnFullScreen,omitempty" xml:"AdminKeyboardOnFullScreen,omitempty"`
	// Specifies whether to enable administrator keyboard control within the Windows system. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	AdminKeyboardOnWindows *string `json:"AdminKeyboardOnWindows,omitempty" xml:"AdminKeyboardOnWindows,omitempty"`
	// Specifies whether to enable the screenshot prevention feature.
	//
	// example:
	//
	// on
	AppContentProtection *string `json:"AppContentProtection,omitempty" xml:"AppContentProtection,omitempty"`
	// The list of client IP whitelist entries to add.
	AuthorizeAccessPolicyRule []*ModifyCenterPolicyRequestAuthorizeAccessPolicyRule `json:"AuthorizeAccessPolicyRule,omitempty" xml:"AuthorizeAccessPolicyRule,omitempty" type:"Repeated"`
	// The list of security group control rules to add.
	AuthorizeSecurityPolicyRule []*ModifyCenterPolicyRequestAuthorizeSecurityPolicyRule `json:"AuthorizeSecurityPolicyRule,omitempty" xml:"AuthorizeSecurityPolicyRule,omitempty" type:"Repeated"`
	// Specifies whether to automatically reconnect after disconnection.
	//
	// example:
	//
	// off
	AutoReconnect *string `json:"AutoReconnect,omitempty" xml:"AutoReconnect,omitempty"`
	// The business channel. Valid values:
	//
	// - Enterprise: Enterprise Edition.
	//
	// - Business: Business Edition.
	//
	// example:
	//
	// Enterprise
	BusinessChannel *string `json:"BusinessChannel,omitempty" xml:"BusinessChannel,omitempty"`
	// The business type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	BusinessType *int32 `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The local camera redirection policy. This parameter takes effect only when no local camera redirection policy is specified in DeviceRedirects.
	//
	// example:
	//
	// on
	CameraRedirect *string `json:"CameraRedirect,omitempty" xml:"CameraRedirect,omitempty"`
	// Specifies whether to display the client control menu. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	ClientControlMenu *string `json:"ClientControlMenu,omitempty" xml:"ClientControlMenu,omitempty"`
	// Specifies whether to enable the custom snapshot creation feature on the client. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	ClientCreateSnapshot *string `json:"ClientCreateSnapshot,omitempty" xml:"ClientCreateSnapshot,omitempty"`
	// The list of logon method control rules. Specifies which client types can be used to access cloud computers.
	ClientType []*ModifyCenterPolicyRequestClientType `json:"ClientType,omitempty" xml:"ClientType,omitempty" type:"Repeated"`
	// The clipboard permission.
	//
	// example:
	//
	// off
	Clipboard *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	// The fine-grained clipboard control configurations.
	ClipboardGraineds []*ModifyCenterPolicyRequestClipboardGraineds `json:"ClipboardGraineds,omitempty" xml:"ClipboardGraineds,omitempty" type:"Repeated"`
	// The effective scope of the clipboard.
	//
	// example:
	//
	// GLOBAL
	ClipboardScope *string `json:"ClipboardScope,omitempty" xml:"ClipboardScope,omitempty"`
	// Indicates whether color enhancement is enabled for the design and 3D application common scenarios.
	//
	// example:
	//
	// off
	ColorEnhancement *string `json:"ColorEnhancement,omitempty" xml:"ColorEnhancement,omitempty"`
	// Specifies whether to enable the local drive clipboard feature. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	CpdDriveClipboard *string `json:"CpdDriveClipboard,omitempty" xml:"CpdDriveClipboard,omitempty"`
	// The CPU throttling duration. Valid values: 30 to 120. Unit: seconds.
	//
	// example:
	//
	// 50
	CpuDownGradeDuration *int32 `json:"CpuDownGradeDuration,omitempty" xml:"CpuDownGradeDuration,omitempty"`
	// Specifies whether to enable CPU overload protection. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	CpuOverload *string `json:"CpuOverload,omitempty" xml:"CpuOverload,omitempty"`
	// The process name.
	CpuProcessors []*string `json:"CpuProcessors,omitempty" xml:"CpuProcessors,omitempty" type:"Repeated"`
	// Specifies whether to enable CPU protection mode.
	//
	// example:
	//
	// off
	CpuProtectedMode *string `json:"CpuProtectedMode,omitempty" xml:"CpuProtectedMode,omitempty"`
	// The overall CPU usage percentage. Valid values: 70 to 90.
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
	// The single-core CPU usage percentage. Valid values: 70 to 100.
	//
	// example:
	//
	// 80
	CpuSingleRateLimit *int32 `json:"CpuSingleRateLimit,omitempty" xml:"CpuSingleRateLimit,omitempty"`
	// The description of the NAS file system.
	//
	// example:
	//
	// newDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The peripheral connection notification control.
	//
	// example:
	//
	// off
	DeviceConnectHint *string `json:"DeviceConnectHint,omitempty" xml:"DeviceConnectHint,omitempty"`
	// The list of device redirection rules.
	DeviceRedirects []*ModifyCenterPolicyRequestDeviceRedirects `json:"DeviceRedirects,omitempty" xml:"DeviceRedirects,omitempty" type:"Repeated"`
	// The list of custom peripheral rules.
	DeviceRules []*ModifyCenterPolicyRequestDeviceRules `json:"DeviceRules,omitempty" xml:"DeviceRules,omitempty" type:"Repeated"`
	// The session retention policy after disconnection.
	//
	// > This parameter applies only to cloud application policies.
	//
	// example:
	//
	// customTime
	DisconnectKeepSession *string `json:"DisconnectKeepSession,omitempty" xml:"DisconnectKeepSession,omitempty"`
	// The session retention period after disconnection. Valid values: 30 to 7200. Unit: seconds.
	//
	// > This parameter applies only to cloud application policies.
	//
	// example:
	//
	// 120
	DisconnectKeepSessionTime *int32 `json:"DisconnectKeepSessionTime,omitempty" xml:"DisconnectKeepSessionTime,omitempty"`
	// Specifies whether to enable disk overload protection. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	DiskOverload *string `json:"DiskOverload,omitempty" xml:"DiskOverload,omitempty"`
	// The display mode.
	//
	// example:
	//
	// clientCustom
	DisplayMode *string `json:"DisplayMode,omitempty" xml:"DisplayMode,omitempty"`
	// The domain name resolution policy.
	DomainResolveRule []*ModifyCenterPolicyRequestDomainResolveRule `json:"DomainResolveRule,omitempty" xml:"DomainResolveRule,omitempty" type:"Repeated"`
	// The domain name resolution policy type.
	//
	// example:
	//
	// off
	DomainResolveRuleType *string `json:"DomainResolveRuleType,omitempty" xml:"DomainResolveRuleType,omitempty"`
	// Specifies whether to enable session bandwidth throttling.
	//
	// example:
	//
	// off
	EnableSessionRateLimiting *string `json:"EnableSessionRateLimiting,omitempty" xml:"EnableSessionRateLimiting,omitempty"`
	// Specifies whether users can request assistance from administrators.
	//
	// example:
	//
	// off
	EndUserApplyAdminCoordinate *string `json:"EndUserApplyAdminCoordinate,omitempty" xml:"EndUserApplyAdminCoordinate,omitempty"`
	// Specifies whether users in the same office network can share cloud desktops.
	//
	// example:
	//
	// off
	EndUserGroupCoordinate *string `json:"EndUserGroupCoordinate,omitempty" xml:"EndUserGroupCoordinate,omitempty"`
	// Specifies whether to enable external storage devices. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	ExternalDrive *string `json:"ExternalDrive,omitempty" xml:"ExternalDrive,omitempty"`
	// Specifies whether to enable file migration.
	//
	// example:
	//
	// off
	FileMigrate *string `json:"FileMigrate,omitempty" xml:"FileMigrate,omitempty"`
	// The service address for the file transfer feature.
	//
	// example:
	//
	// filetransfer.example.com
	FileTransferAddress *string `json:"FileTransferAddress,omitempty" xml:"FileTransferAddress,omitempty"`
	// The maximum file size for a single transfer to the cloud desktop. This parameter must be used together with the transfer-in unit parameter.
	//
	// example:
	//
	// 100
	FileTransferInSize *string `json:"FileTransferInSize,omitempty" xml:"FileTransferInSize,omitempty"`
	// The unit of the maximum file size for a single transfer to the cloud desktop.
	//
	// example:
	//
	// MB
	FileTransferInUnit *string `json:"FileTransferInUnit,omitempty" xml:"FileTransferInUnit,omitempty"`
	// The maximum file size for a single transfer from the cloud desktop. This parameter must be used together with the transfer-out unit parameter.
	//
	// example:
	//
	// 100
	FileTransferOutSize *string `json:"FileTransferOutSize,omitempty" xml:"FileTransferOutSize,omitempty"`
	// The unit of the maximum file size for a single transfer from the cloud desktop.
	//
	// example:
	//
	// MB
	FileTransferOutUnit *string `json:"FileTransferOutUnit,omitempty" xml:"FileTransferOutUnit,omitempty"`
	// Specifies whether to enable the file transfer size limit. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	FileTransferSizeLimit *string `json:"FileTransferSizeLimit,omitempty" xml:"FileTransferSizeLimit,omitempty"`
	// The file transfer speed level.
	//
	// example:
	//
	// default
	FileTransferSpeed *string `json:"FileTransferSpeed,omitempty" xml:"FileTransferSpeed,omitempty"`
	// The location where the file transfer speed configured on the client takes effect.
	//
	// example:
	//
	// client
	FileTransferSpeedLocation *string `json:"FileTransferSpeedLocation,omitempty" xml:"FileTransferSpeedLocation,omitempty"`
	// Specifies whether to enable the image quality policy for GPU-accelerated cloud desktops. Enable this policy when high cloud desktop performance and user experience are required, such as in professional design scenarios.
	//
	// example:
	//
	// off
	GpuAcceleration *string `json:"GpuAcceleration,omitempty" xml:"GpuAcceleration,omitempty"`
	// Specifies whether to enable the floating ball configuration message prompt. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	HoverConfigMsg *string `json:"HoverConfigMsg,omitempty" xml:"HoverConfigMsg,omitempty"`
	// The file transfer policy for the web client.
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
	// Specifies whether to enable the network printer feature. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	InternetPrinter *string `json:"InternetPrinter,omitempty" xml:"InternetPrinter,omitempty"`
	// The local disk mapping permission.
	//
	// example:
	//
	// off
	LocalDrive *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	// The maximum connection retry time when a cloud computer is disconnected due to objective reasons. Valid values: 30 to 7200. Unit: seconds.
	//
	// example:
	//
	// 120
	MaxReconnectTime *int32 `json:"MaxReconnectTime,omitempty" xml:"MaxReconnectTime,omitempty"`
	// The memory throttling duration of a single process. Valid values: 30 to 120. Unit: seconds.
	//
	// example:
	//
	// 40
	MemoryDownGradeDuration *int32 `json:"MemoryDownGradeDuration,omitempty" xml:"MemoryDownGradeDuration,omitempty"`
	// Specifies whether to enable memory overload protection. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	MemoryOverload *string `json:"MemoryOverload,omitempty" xml:"MemoryOverload,omitempty"`
	// The process name.
	MemoryProcessors []*string `json:"MemoryProcessors,omitempty" xml:"MemoryProcessors,omitempty" type:"Repeated"`
	// Specifies whether to enable memory protection mode.
	//
	// example:
	//
	// off
	MemoryProtectedMode *string `json:"MemoryProtectedMode,omitempty" xml:"MemoryProtectedMode,omitempty"`
	// The overall memory usage percentage. Valid values: 70 to 90.
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
	// The memory usage percentage of a single process. Valid values: 30 to 60.
	//
	// example:
	//
	// 40
	MemorySingleRateLimit *int32 `json:"MemorySingleRateLimit,omitempty" xml:"MemorySingleRateLimit,omitempty"`
	// Specifies whether to provide the restart button in the cloud desktop floating ball when connecting to the cloud desktop from a mobile client (Android client<props="china"> and iOS client).
	//
	// > This parameter applies only to mobile clients of V7.4 or later.
	//
	// example:
	//
	// off
	MobileRestart *string `json:"MobileRestart,omitempty" xml:"MobileRestart,omitempty"`
	// Specifies whether to enable the security button for Windows on the mobile client.
	//
	// example:
	//
	// off
	MobileSafeMenu *string `json:"MobileSafeMenu,omitempty" xml:"MobileSafeMenu,omitempty"`
	// Specifies whether to provide the shutdown button in the cloud desktop floating ball when connecting to the cloud desktop from a mobile client (Android client<props="china"> and iOS client).
	//
	// > This parameter applies only to mobile clients of V7.4 or later.
	//
	// example:
	//
	// off
	MobileShutdown *string `json:"MobileShutdown,omitempty" xml:"MobileShutdown,omitempty"`
	// Specifies whether to enable WUYING Keeper on the mobile client.
	//
	// example:
	//
	// off
	MobileWuyingKeeper *string `json:"MobileWuyingKeeper,omitempty" xml:"MobileWuyingKeeper,omitempty"`
	// Specifies whether to enable WUYING Assistant on the mobile client.
	//
	// example:
	//
	// off
	MobileWyAssistant *string `json:"MobileWyAssistant,omitempty" xml:"MobileWyAssistant,omitempty"`
	// Specifies whether to enable the model library feature. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	ModelLibrary *string `json:"ModelLibrary,omitempty" xml:"ModelLibrary,omitempty"`
	// Specifies whether to enable the multi-screen display feature. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	MultiScreen *string `json:"MultiScreen,omitempty" xml:"MultiScreen,omitempty"`
	// The policy name.
	//
	// example:
	//
	// testPolicyGroupName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether to enable network redirection.
	//
	// > This feature is in invitational preview and is not publicly available.
	//
	// example:
	//
	// on
	NetRedirect *string `json:"NetRedirect,omitempty" xml:"NetRedirect,omitempty"`
	// The details of the network redirect policy.
	//
	// > This feature is in invitational preview and is not publicly available.
	//
	// >
	NetRedirectRule []*ModifyCenterPolicyRequestNetRedirectRule `json:"NetRedirectRule,omitempty" xml:"NetRedirectRule,omitempty" type:"Repeated"`
	// Specifies whether to disconnect the session when no operation is performed.
	//
	// > This parameter applies only to cloud application policies.
	//
	// example:
	//
	// off
	NoOperationDisconnect *string `json:"NoOperationDisconnect,omitempty" xml:"NoOperationDisconnect,omitempty"`
	// The idle disconnection period. Valid values: 120 to 7200. Unit: seconds.
	//
	// > This parameter applies only to cloud application policies.
	//
	// example:
	//
	// 120
	NoOperationDisconnectTime *int32 `json:"NoOperationDisconnectTime,omitempty" xml:"NoOperationDisconnectTime,omitempty"`
	// The ID of the cloud desktop policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// pg-53iyi2aar0nd6****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// Specifies whether to enable the port proxy feature. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	PortProxy *string `json:"PortProxy,omitempty" xml:"PortProxy,omitempty"`
	// The printer pop-up prompt. Valid values:
	//
	// - default: Default value.
	//
	// - off: Disabled.
	//
	// - custom: Custom.
	//
	// example:
	//
	// off
	PrinterAlert *string `json:"PrinterAlert,omitempty" xml:"PrinterAlert,omitempty"`
	// The content of the printer pop-up prompt.
	//
	// example:
	//
	// Print Content.
	PrinterAlertContent *string `json:"PrinterAlertContent,omitempty" xml:"PrinterAlertContent,omitempty"`
	// The title of the printer pop-up prompt.
	//
	// example:
	//
	// Print Title
	PrinterAlertTitle *string `json:"PrinterAlertTitle,omitempty" xml:"PrinterAlertTitle,omitempty"`
	// The printer redirection policy. This parameter takes effect only when no printer redirection policy is specified in DeviceRedirects.
	//
	// example:
	//
	// off
	PrinterRedirect *string `json:"PrinterRedirect,omitempty" xml:"PrinterRedirect,omitempty"`
	// Indicates whether image quality enhancement is enabled for the design and 3D application common scenarios.
	//
	// example:
	//
	// off
	QualityEnhancement *string `json:"QualityEnhancement,omitempty" xml:"QualityEnhancement,omitempty"`
	// The duration of screen recording after an event is detected in screen recording audits. Unit: minutes. Valid values: 10 to 60.
	//
	// example:
	//
	// 10
	RecordEventDuration *int32 `json:"RecordEventDuration,omitempty" xml:"RecordEventDuration,omitempty"`
	// The file extensions of recording events.
	RecordEventFileExts []*string `json:"RecordEventFileExts,omitempty" xml:"RecordEventFileExts,omitempty" type:"Repeated"`
	// The absolute paths for file monitoring in screen recording audits.
	RecordEventFilePaths []*string `json:"RecordEventFilePaths,omitempty" xml:"RecordEventFilePaths,omitempty" type:"Repeated"`
	// The levels of recording events.
	RecordEventLevels []*ModifyCenterPolicyRequestRecordEventLevels `json:"RecordEventLevels,omitempty" xml:"RecordEventLevels,omitempty" type:"Repeated"`
	// The absolute paths for registry monitoring in screen recording audits.
	RecordEventRegisters []*string `json:"RecordEventRegisters,omitempty" xml:"RecordEventRegisters,omitempty" type:"Repeated"`
	// The list of screen recording events.
	RecordEvents []*string `json:"RecordEvents,omitempty" xml:"RecordEvents,omitempty" type:"Repeated"`
	// Specifies whether to enable screen recording.
	//
	// example:
	//
	// off
	Recording *string `json:"Recording,omitempty" xml:"Recording,omitempty"`
	// The audio recording option for cloud desktops.
	//
	// example:
	//
	// on
	RecordingAudio *string `json:"RecordingAudio,omitempty" xml:"RecordingAudio,omitempty"`
	// The duration of each screen recording file, in minutes. Recording files are automatically split and uploaded to the storage space based on the specified duration. Files are rolled over when they reach 300 MB. Valid values: 10 to 60.
	//
	// example:
	//
	// 10
	RecordingDuration *int32 `json:"RecordingDuration,omitempty" xml:"RecordingDuration,omitempty"`
	// The end time of screen recording. Format: HH:MM:SS. This response value is meaningful only when `Recording` is set to `PERIOD`.
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
	RecordingExpires *int32 `json:"RecordingExpires,omitempty" xml:"RecordingExpires,omitempty"`
	// The screen recording frame rate. Unit: FPS (frames per second).
	//
	// example:
	//
	// 5
	RecordingFps *string `json:"RecordingFps,omitempty" xml:"RecordingFps,omitempty"`
	// The start time of screen recording. Format: HH:MM:SS. This response value is meaningful only when `Recording` is set to `PERIOD`.
	//
	// example:
	//
	// 08:00:00
	RecordingStartTime *string `json:"RecordingStartTime,omitempty" xml:"RecordingStartTime,omitempty"`
	// Specifies whether to notify end users that screen recording is enabled.
	//
	// example:
	//
	// off
	RecordingUserNotify *string `json:"RecordingUserNotify,omitempty" xml:"RecordingUserNotify,omitempty"`
	// The notification message displayed to end users when screen recording is enabled.
	//
	// example:
	//
	// Screen recording is enabled
	RecordingUserNotifyMessage *string `json:"RecordingUserNotifyMessage,omitempty" xml:"RecordingUserNotifyMessage,omitempty"`
	// The region ID. This feature is region-independent. Set this parameter to `cn-shanghai`.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The keyboard and mouse control permission for remote assistance.
	//
	// example:
	//
	// fullControl
	RemoteCoordinate *string `json:"RemoteCoordinate,omitempty" xml:"RemoteCoordinate,omitempty"`
	// Specifies whether to allow resetting the cloud desktop.
	//
	// example:
	//
	// off
	ResetDesktop *string `json:"ResetDesktop,omitempty" xml:"ResetDesktop,omitempty"`
	// The DPI value of the screen resolution.
	//
	// example:
	//
	// 96
	ResolutionDpi *int32 `json:"ResolutionDpi,omitempty" xml:"ResolutionDpi,omitempty"`
	// The height of the resolution. Unit: pixels. Valid values for cloud applications: 500 to 50000. Valid values for cloud desktops: 480 to 4096.
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
	// The width of the resolution. Unit: pixels. Valid values for cloud applications: 500 to 50000. Valid values for cloud desktops: 480 to 4096.
	//
	// example:
	//
	// 720
	ResolutionWidth *int32 `json:"ResolutionWidth,omitempty" xml:"ResolutionWidth,omitempty"`
	// The resource type.
	//
	// This parameter is required.
	//
	// example:
	//
	// desktop
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The list of client IP whitelist entries to delete.
	RevokeAccessPolicyRule []*ModifyCenterPolicyRequestRevokeAccessPolicyRule `json:"RevokeAccessPolicyRule,omitempty" xml:"RevokeAccessPolicyRule,omitempty" type:"Repeated"`
	// The list of security group rules to delete.
	RevokeSecurityPolicyRule []*ModifyCenterPolicyRequestRevokeSecurityPolicyRule `json:"RevokeSecurityPolicyRule,omitempty" xml:"RevokeSecurityPolicyRule,omitempty" type:"Repeated"`
	// Specifies whether to enable the Security Center shortcut key. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	SafeMenu *string `json:"SafeMenu,omitempty" xml:"SafeMenu,omitempty"`
	// The scope in which the policy takes effect.
	//
	// example:
	//
	// GLOBAL
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The value to specify when `Scope` is set to `IP`. This parameter takes effect only when `Scope` is set to `IP`.
	ScopeValue []*string `json:"ScopeValue,omitempty" xml:"ScopeValue,omitempty" type:"Repeated"`
	// The screen display mode.
	//
	// example:
	//
	// auto
	ScreenDisplayMode *string `json:"ScreenDisplayMode,omitempty" xml:"ScreenDisplayMode,omitempty"`
	// The maximum value of session bandwidth throttling. Unit: Kbps. Valid values: 2000 to 100000.
	//
	// example:
	//
	// 2000
	SessionMaxRateKbps *int32 `json:"SessionMaxRateKbps,omitempty" xml:"SessionMaxRateKbps,omitempty"`
	// Specifies whether to enable smoothness enhancement for the daily office scenario.
	//
	// example:
	//
	// off
	SmoothEnhancement *string `json:"SmoothEnhancement,omitempty" xml:"SmoothEnhancement,omitempty"`
	// Specifies whether to provide the status monitoring entry in the cloud desktop floating ball.
	//
	// example:
	//
	// off
	StatusMonitor *string `json:"StatusMonitor,omitempty" xml:"StatusMonitor,omitempty"`
	// The streaming mode adaptation scenario.
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
	// The application taskbar.
	//
	// > This parameter applies only to cloud application policies.
	//
	// example:
	//
	// off
	Taskbar *string `json:"Taskbar,omitempty" xml:"Taskbar,omitempty"`
	// Specifies whether to enable the three-screen feature. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	ThreeScreen *string `json:"ThreeScreen,omitempty" xml:"ThreeScreen,omitempty"`
	// The USB redirection policy.
	//
	// example:
	//
	// off
	UsbRedirect *string `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty"`
	// The USB redirection rule.
	UsbSupplyRedirectRule []*ModifyCenterPolicyRequestUsbSupplyRedirectRule `json:"UsbSupplyRedirectRule,omitempty" xml:"UsbSupplyRedirectRule,omitempty" type:"Repeated"`
	// Specifies whether to display the usage duration in the floating ball. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	UseTime *string `json:"UseTime,omitempty" xml:"UseTime,omitempty"`
	// The average bitrate for video encoding. Unit: Kbps. Valid values: 1000 to 50000.
	//
	// example:
	//
	// 2000
	VideoEncAvgKbps *int32 `json:"VideoEncAvgKbps,omitempty" xml:"VideoEncAvgKbps,omitempty"`
	// The maximum QP for video encoding, which represents the lowest quality. Valid values: 0 to 51.
	//
	// example:
	//
	// 30
	VideoEncMaxQP *int32 `json:"VideoEncMaxQP,omitempty" xml:"VideoEncMaxQP,omitempty"`
	// The minimum QP for video encoding, which represents the highest quality. Valid values: 0 to 51.
	//
	// example:
	//
	// 30
	VideoEncMinQP *int32 `json:"VideoEncMinQP,omitempty" xml:"VideoEncMinQP,omitempty"`
	// The peak video encoding bitrate. Unit: Kbps. Valid values: 1000 to 50000.
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
	// The multimedia redirection policy.
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
	// The invisible watermark anti-photography feature.
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
	// The number of watermark columns. Valid values: 3 to 10.
	//
	// example:
	//
	// 3
	WatermarkColumnAmount *int32 `json:"WatermarkColumnAmount,omitempty" xml:"WatermarkColumnAmount,omitempty"`
	// If the `WatermarkType` parameter is set to `custom`, you must also specify the custom text content by using the `WatermarkCustomText` parameter.
	//
	// example:
	//
	// Internal Document
	WatermarkCustomText *string `json:"WatermarkCustomText,omitempty" xml:"WatermarkCustomText,omitempty"`
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
	// The invisible watermark enhancement feature.
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
	// The security priority rule for invisible watermarks.
	//
	// example:
	//
	// off
	WatermarkSecurity *string `json:"WatermarkSecurity,omitempty" xml:"WatermarkSecurity,omitempty"`
	// Specifies whether to enable the watermark shadow effect. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	WatermarkShadow *string `json:"WatermarkShadow,omitempty" xml:"WatermarkShadow,omitempty"`
	// The watermark opacity. A larger value indicates lower transparency. Valid values: 10 to 100.
	//
	// example:
	//
	// 10
	WatermarkTransparencyValue *int32 `json:"WatermarkTransparencyValue,omitempty" xml:"WatermarkTransparencyValue,omitempty"`
	// The watermark type. You can select up to three types, separated by commas (,).
	//
	// > If you set this parameter to `custom`, you must also specify the custom text content by using the `WatermarkCustomText` parameter.
	//
	// example:
	//
	// EndUserId
	WatermarkType *string `json:"WatermarkType,omitempty" xml:"WatermarkType,omitempty"`
	// Specifies whether to enable WUYING Keeper.
	//
	// example:
	//
	// off
	WuyingKeeper *string `json:"WuyingKeeper,omitempty" xml:"WuyingKeeper,omitempty"`
	// Specifies whether to provide the WUYING AI Assistant entry in the floating ball when connecting to a cloud computer through a desktop client (including Windows client and macOS client).
	//
	// > This feature applies only to desktop clients of V7.7 or later.
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

func (s *ModifyCenterPolicyRequest) GetAdminKeyboardOnFullScreen() *string {
	return s.AdminKeyboardOnFullScreen
}

func (s *ModifyCenterPolicyRequest) GetAdminKeyboardOnWindows() *string {
	return s.AdminKeyboardOnWindows
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

func (s *ModifyCenterPolicyRequest) GetDescription() *string {
	return s.Description
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

func (s *ModifyCenterPolicyRequest) GetFileTransferInSize() *string {
	return s.FileTransferInSize
}

func (s *ModifyCenterPolicyRequest) GetFileTransferInUnit() *string {
	return s.FileTransferInUnit
}

func (s *ModifyCenterPolicyRequest) GetFileTransferOutSize() *string {
	return s.FileTransferOutSize
}

func (s *ModifyCenterPolicyRequest) GetFileTransferOutUnit() *string {
	return s.FileTransferOutUnit
}

func (s *ModifyCenterPolicyRequest) GetFileTransferSizeLimit() *string {
	return s.FileTransferSizeLimit
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

func (s *ModifyCenterPolicyRequest) GetHoverConfigMsg() *string {
	return s.HoverConfigMsg
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

func (s *ModifyCenterPolicyRequest) GetMultiScreen() *string {
	return s.MultiScreen
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

func (s *ModifyCenterPolicyRequest) GetPrinterAlert() *string {
	return s.PrinterAlert
}

func (s *ModifyCenterPolicyRequest) GetPrinterAlertContent() *string {
	return s.PrinterAlertContent
}

func (s *ModifyCenterPolicyRequest) GetPrinterAlertTitle() *string {
	return s.PrinterAlertTitle
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

func (s *ModifyCenterPolicyRequest) GetResolutionDpi() *int32 {
	return s.ResolutionDpi
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

func (s *ModifyCenterPolicyRequest) GetThreeScreen() *string {
	return s.ThreeScreen
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

func (s *ModifyCenterPolicyRequest) SetAdminKeyboardOnFullScreen(v string) *ModifyCenterPolicyRequest {
	s.AdminKeyboardOnFullScreen = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetAdminKeyboardOnWindows(v string) *ModifyCenterPolicyRequest {
	s.AdminKeyboardOnWindows = &v
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

func (s *ModifyCenterPolicyRequest) SetDescription(v string) *ModifyCenterPolicyRequest {
	s.Description = &v
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

func (s *ModifyCenterPolicyRequest) SetFileTransferInSize(v string) *ModifyCenterPolicyRequest {
	s.FileTransferInSize = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetFileTransferInUnit(v string) *ModifyCenterPolicyRequest {
	s.FileTransferInUnit = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetFileTransferOutSize(v string) *ModifyCenterPolicyRequest {
	s.FileTransferOutSize = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetFileTransferOutUnit(v string) *ModifyCenterPolicyRequest {
	s.FileTransferOutUnit = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetFileTransferSizeLimit(v string) *ModifyCenterPolicyRequest {
	s.FileTransferSizeLimit = &v
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

func (s *ModifyCenterPolicyRequest) SetHoverConfigMsg(v string) *ModifyCenterPolicyRequest {
	s.HoverConfigMsg = &v
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

func (s *ModifyCenterPolicyRequest) SetMultiScreen(v string) *ModifyCenterPolicyRequest {
	s.MultiScreen = &v
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

func (s *ModifyCenterPolicyRequest) SetPrinterAlert(v string) *ModifyCenterPolicyRequest {
	s.PrinterAlert = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetPrinterAlertContent(v string) *ModifyCenterPolicyRequest {
	s.PrinterAlertContent = &v
	return s
}

func (s *ModifyCenterPolicyRequest) SetPrinterAlertTitle(v string) *ModifyCenterPolicyRequest {
	s.PrinterAlertTitle = &v
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

func (s *ModifyCenterPolicyRequest) SetResolutionDpi(v int32) *ModifyCenterPolicyRequest {
	s.ResolutionDpi = &v
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

func (s *ModifyCenterPolicyRequest) SetThreeScreen(v string) *ModifyCenterPolicyRequest {
	s.ThreeScreen = &v
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
	// The client IP address range. The value is an IPv4 CIDR block.
	//
	// example:
	//
	// 47.100.XX.XX/16
	CidrIp *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	// The description of the client IP whitelist entry.
	//
	// example:
	//
	// Corporate office network segment
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
	// The object of the security group control rule. An IPv4 CIDR block in CIDR notation.
	//
	// example:
	//
	// 47.100.XX.XX/16
	CidrIp *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	// The description of the security group control rule.
	//
	// example:
	//
	// Allow access to the internal R&D environment
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The protocol type of the security group control rule.
	//
	// example:
	//
	// TCP
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The authorization policy of the security group control rule.
	//
	// example:
	//
	// accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The port range of the security group control rule. The port range is determined by the value of the protocol (IpProtocol):
	//
	// - TCP or UDP: Valid values: 1 to 65535. Separate the start port and end port with a forward slash (/). Example: 1/200.
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
	// The priority of the security group control rule. A smaller value indicates a higher priority. Valid values: 1 to 60. Default value: 1.
	//
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The direction of the security group control rule.
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
	// The client type for logon method control.
	//
	// example:
	//
	// android
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// Specifies whether to allow a specific type of client to log on to cloud computers.
	//
	// > If you do not set the `ClientType` parameters, all client types are allowed to log on to cloud computers by default.
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
	// The maximum size of a single clipboard transfer. Use this parameter together with the size unit parameter.
	//
	// example:
	//
	// 10
	ClipboardSize *int32 `json:"ClipboardSize,omitempty" xml:"ClipboardSize,omitempty"`
	// The unit of the maximum size of a single clipboard transfer.
	//
	// example:
	//
	// MB
	ClipboardSizeUnit *string `json:"ClipboardSizeUnit,omitempty" xml:"ClipboardSizeUnit,omitempty"`
	// The fine-grained clipboard control type. Valid values:
	//
	// - off: Clipboard usage is disabled.
	//
	// - read: Read-only.
	//
	// - write: Write-only.
	//
	// - readwrite: Read and write.
	//
	// example:
	//
	// readwrite
	ClipboardType *string `json:"ClipboardType,omitempty" xml:"ClipboardType,omitempty"`
	// The content type for fine-grained clipboard control. Valid values:
	//
	// - text: Text.
	//
	// - richtext: Rich text.
	//
	// - file: File.
	//
	// - picture: Image.
	//
	// example:
	//
	// text
	GrainedType *string `json:"GrainedType,omitempty" xml:"GrainedType,omitempty"`
	// The maximum size of a single clipboard transfer to the cloud desktop. Use this parameter together with the inbound size unit parameter.
	//
	// example:
	//
	// 10
	InClipboardSize *int32 `json:"InClipboardSize,omitempty" xml:"InClipboardSize,omitempty"`
	// The unit of the maximum size of a single clipboard transfer to the cloud desktop.
	//
	// example:
	//
	// MB
	InClipboardSizeUnit *string `json:"InClipboardSizeUnit,omitempty" xml:"InClipboardSizeUnit,omitempty"`
	// The maximum size of a single clipboard transfer from the cloud desktop. Use this parameter together with the outbound size unit parameter.
	//
	// example:
	//
	// 10
	OutClipboardSize *int32 `json:"OutClipboardSize,omitempty" xml:"OutClipboardSize,omitempty"`
	// The unit of the maximum size of a single clipboard transfer from the cloud desktop.
	//
	// example:
	//
	// MB
	OutClipboardSizeUnit *string `json:"OutClipboardSizeUnit,omitempty" xml:"OutClipboardSizeUnit,omitempty"`
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

func (s *ModifyCenterPolicyRequestClipboardGraineds) GetClipboardSizeUnit() *string {
	return s.ClipboardSizeUnit
}

func (s *ModifyCenterPolicyRequestClipboardGraineds) GetClipboardType() *string {
	return s.ClipboardType
}

func (s *ModifyCenterPolicyRequestClipboardGraineds) GetGrainedType() *string {
	return s.GrainedType
}

func (s *ModifyCenterPolicyRequestClipboardGraineds) GetInClipboardSize() *int32 {
	return s.InClipboardSize
}

func (s *ModifyCenterPolicyRequestClipboardGraineds) GetInClipboardSizeUnit() *string {
	return s.InClipboardSizeUnit
}

func (s *ModifyCenterPolicyRequestClipboardGraineds) GetOutClipboardSize() *int32 {
	return s.OutClipboardSize
}

func (s *ModifyCenterPolicyRequestClipboardGraineds) GetOutClipboardSizeUnit() *string {
	return s.OutClipboardSizeUnit
}

func (s *ModifyCenterPolicyRequestClipboardGraineds) SetClipboardSize(v int32) *ModifyCenterPolicyRequestClipboardGraineds {
	s.ClipboardSize = &v
	return s
}

func (s *ModifyCenterPolicyRequestClipboardGraineds) SetClipboardSizeUnit(v string) *ModifyCenterPolicyRequestClipboardGraineds {
	s.ClipboardSizeUnit = &v
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

func (s *ModifyCenterPolicyRequestClipboardGraineds) SetInClipboardSize(v int32) *ModifyCenterPolicyRequestClipboardGraineds {
	s.InClipboardSize = &v
	return s
}

func (s *ModifyCenterPolicyRequestClipboardGraineds) SetInClipboardSizeUnit(v string) *ModifyCenterPolicyRequestClipboardGraineds {
	s.InClipboardSizeUnit = &v
	return s
}

func (s *ModifyCenterPolicyRequestClipboardGraineds) SetOutClipboardSize(v int32) *ModifyCenterPolicyRequestClipboardGraineds {
	s.OutClipboardSize = &v
	return s
}

func (s *ModifyCenterPolicyRequestClipboardGraineds) SetOutClipboardSizeUnit(v string) *ModifyCenterPolicyRequestClipboardGraineds {
	s.OutClipboardSizeUnit = &v
	return s
}

func (s *ModifyCenterPolicyRequestClipboardGraineds) Validate() error {
	return dara.Validate(s)
}

type ModifyCenterPolicyRequestDeviceRedirects struct {
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
	// The product ID.
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
	// The vendor ID. See [Valid USB Vendor IDs (VIDs)](https://www.usb.org/sites/default/files/vendor_ids032322.pdf_1.pdf).
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
	// *.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The redirect policy.
	//
	// example:
	//
	// Allow
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The rule type.
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
	// The event level.
	//
	// example:
	//
	// HIGH
	EventLevel *string `json:"EventLevel,omitempty" xml:"EventLevel,omitempty"`
	// The event type.
	//
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
	// Corporate office network segment
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
	// TCP
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The authorization policy of the security group rule to delete.
	//
	// example:
	//
	// accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The port range of the security group rule to delete. The port range is determined by the value of IpProtocol:
	//
	// - TCP or UDP: The port range is 1 to 65535. Separate the start port and end port with a forward slash (/). Example: 1/200.
	//
	// - ICMP: -1/-1.
	//
	// - GRE: -1/-1.
	//
	// - If IpProtocol is set to all: -1/-1.
	//
	// For more information about common ports of typical applications, see [Common ports](https://help.aliyun.com/document_detail/40724.html).
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
	//
	// example:
	//
	// Test rule
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
	UsbRedirectType *string `json:"UsbRedirectType,omitempty" xml:"UsbRedirectType,omitempty"`
	// The USB redirection rule type.
	//
	// example:
	//
	// 1
	UsbRuleType *string `json:"UsbRuleType,omitempty" xml:"UsbRuleType,omitempty"`
	// The vendor ID. See [Valid USB Vendor IDs (VIDs)](https://www.usb.org/sites/default/files/vendor_ids032322.pdf_1.pdf).
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
