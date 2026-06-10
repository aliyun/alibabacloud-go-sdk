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
	// Specifies whether users have administrator permissions after logging on to cloud computers.
	//
	// > This feature is in invitational preview and not available to the public.
	//
	// example:
	//
	// deny
	AdminAccess               *string `json:"AdminAccess,omitempty" xml:"AdminAccess,omitempty"`
	AdminKeyboardOnFullScreen *string `json:"AdminKeyboardOnFullScreen,omitempty" xml:"AdminKeyboardOnFullScreen,omitempty"`
	AdminKeyboardOnWindows    *string `json:"AdminKeyboardOnWindows,omitempty" xml:"AdminKeyboardOnWindows,omitempty"`
	// Specifies whether to enable anti-screenshot protection.
	//
	// example:
	//
	// on
	AppContentProtection *string `json:"AppContentProtection,omitempty" xml:"AppContentProtection,omitempty"`
	// List of new client IP address whitelists.
	AuthorizeAccessPolicyRule []*ModifyCenterPolicyRequestAuthorizeAccessPolicyRule `json:"AuthorizeAccessPolicyRule,omitempty" xml:"AuthorizeAccessPolicyRule,omitempty" type:"Repeated"`
	// List of new security group control rules.
	AuthorizeSecurityPolicyRule []*ModifyCenterPolicyRequestAuthorizeSecurityPolicyRule `json:"AuthorizeSecurityPolicyRule,omitempty" xml:"AuthorizeSecurityPolicyRule,omitempty" type:"Repeated"`
	// Automatically reconnect after disconnection
	//
	// example:
	//
	// off
	AutoReconnect   *string `json:"AutoReconnect,omitempty" xml:"AutoReconnect,omitempty"`
	BusinessChannel *string `json:"BusinessChannel,omitempty" xml:"BusinessChannel,omitempty"`
	// The business type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	BusinessType *int32 `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// Local camera redirection. This parameter takes effect only if DeviceRedirects does not include a local camera redirection policy.
	//
	// example:
	//
	// on
	CameraRedirect *string `json:"CameraRedirect,omitempty" xml:"CameraRedirect,omitempty"`
	// Event level for screen recording
	ClientControlMenu    *string `json:"ClientControlMenu,omitempty" xml:"ClientControlMenu,omitempty"`
	ClientCreateSnapshot *string `json:"ClientCreateSnapshot,omitempty" xml:"ClientCreateSnapshot,omitempty"`
	// List of client login control rules. Controls which clients can access cloud computers.
	ClientType []*ModifyCenterPolicyRequestClientType `json:"ClientType,omitempty" xml:"ClientType,omitempty" type:"Repeated"`
	// The clipboard permissions.
	//
	// example:
	//
	// off
	Clipboard         *string                                       `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	ClipboardGraineds []*ModifyCenterPolicyRequestClipboardGraineds `json:"ClipboardGraineds,omitempty" xml:"ClipboardGraineds,omitempty" type:"Repeated"`
	ClipboardScope    *string                                       `json:"ClipboardScope,omitempty" xml:"ClipboardScope,omitempty"`
	// Specifies whether to enable color enhancement for design and 3D applications.
	//
	// example:
	//
	// off
	ColorEnhancement  *string `json:"ColorEnhancement,omitempty" xml:"ColorEnhancement,omitempty"`
	CpdDriveClipboard *string `json:"CpdDriveClipboard,omitempty" xml:"CpdDriveClipboard,omitempty"`
	// CPU downclocking duration. Valid values: 30 to 120. Unit: seconds.
	//
	// example:
	//
	// 50
	CpuDownGradeDuration *int32  `json:"CpuDownGradeDuration,omitempty" xml:"CpuDownGradeDuration,omitempty"`
	CpuOverload          *string `json:"CpuOverload,omitempty" xml:"CpuOverload,omitempty"`
	// The name of the process.
	CpuProcessors []*string `json:"CpuProcessors,omitempty" xml:"CpuProcessors,omitempty" type:"Repeated"`
	// CPU protection mode switch.
	//
	// example:
	//
	// off
	CpuProtectedMode *string `json:"CpuProtectedMode,omitempty" xml:"CpuProtectedMode,omitempty"`
	// Overall CPU usage percentage. Valid values: 70 to 90.
	//
	// example:
	//
	// 70
	CpuRateLimit *int32 `json:"CpuRateLimit,omitempty" xml:"CpuRateLimit,omitempty"`
	// Overall CPU sampling duration. Valid values: 10 to 60. Unit: seconds.
	//
	// example:
	//
	// 30
	CpuSampleDuration *int32 `json:"CpuSampleDuration,omitempty" xml:"CpuSampleDuration,omitempty"`
	// Single-core CPU usage percentage. Valid values: 70 to 100.
	//
	// example:
	//
	// 80
	CpuSingleRateLimit *int32  `json:"CpuSingleRateLimit,omitempty" xml:"CpuSingleRateLimit,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Peripheral connection prompt control.
	//
	// example:
	//
	// off
	DeviceConnectHint *string `json:"DeviceConnectHint,omitempty" xml:"DeviceConnectHint,omitempty"`
	// Device redirection rules.
	DeviceRedirects []*ModifyCenterPolicyRequestDeviceRedirects `json:"DeviceRedirects,omitempty" xml:"DeviceRedirects,omitempty" type:"Repeated"`
	// Custom peripheral rules.
	DeviceRules []*ModifyCenterPolicyRequestDeviceRules `json:"DeviceRules,omitempty" xml:"DeviceRules,omitempty" type:"Repeated"`
	// Session retention after disconnection.
	//
	// > This parameter applies only to cloud application policies.
	//
	// example:
	//
	// customTime
	DisconnectKeepSession *string `json:"DisconnectKeepSession,omitempty" xml:"DisconnectKeepSession,omitempty"`
	// Session retention duration after disconnection. Valid values: 30 to 7200. Unit: seconds.
	//
	// > This parameter applies only to cloud application policies.
	//
	// example:
	//
	// 120
	DisconnectKeepSessionTime *int32  `json:"DisconnectKeepSessionTime,omitempty" xml:"DisconnectKeepSessionTime,omitempty"`
	DiskOverload              *string `json:"DiskOverload,omitempty" xml:"DiskOverload,omitempty"`
	// Display mode.
	//
	// example:
	//
	// clientCustom
	DisplayMode *string `json:"DisplayMode,omitempty" xml:"DisplayMode,omitempty"`
	// Domain name resolution policies.
	DomainResolveRule []*ModifyCenterPolicyRequestDomainResolveRule `json:"DomainResolveRule,omitempty" xml:"DomainResolveRule,omitempty" type:"Repeated"`
	// Domain name resolution policy type.
	//
	// example:
	//
	// off
	DomainResolveRuleType *string `json:"DomainResolveRuleType,omitempty" xml:"DomainResolveRuleType,omitempty"`
	// Session bandwidth throttling.
	//
	// example:
	//
	// off
	EnableSessionRateLimiting *string `json:"EnableSessionRateLimiting,omitempty" xml:"EnableSessionRateLimiting,omitempty"`
	// User requests administrator assistance.
	//
	// example:
	//
	// off
	EndUserApplyAdminCoordinate *string `json:"EndUserApplyAdminCoordinate,omitempty" xml:"EndUserApplyAdminCoordinate,omitempty"`
	// Users on the same office network share cloud computers.
	//
	// example:
	//
	// off
	EndUserGroupCoordinate *string `json:"EndUserGroupCoordinate,omitempty" xml:"EndUserGroupCoordinate,omitempty"`
	ExternalDrive          *string `json:"ExternalDrive,omitempty" xml:"ExternalDrive,omitempty"`
	// File migration.
	//
	// example:
	//
	// off
	FileMigrate         *string `json:"FileMigrate,omitempty" xml:"FileMigrate,omitempty"`
	FileTransferAddress *string `json:"FileTransferAddress,omitempty" xml:"FileTransferAddress,omitempty"`
	FileTransferSpeed   *string `json:"FileTransferSpeed,omitempty" xml:"FileTransferSpeed,omitempty"`
	// Screen recording event suffix
	FileTransferSpeedLocation *string `json:"FileTransferSpeedLocation,omitempty" xml:"FileTransferSpeedLocation,omitempty"`
	// Specifies whether to enable the image quality policy for graphics-intensive cloud computers. Enable this policy for scenarios such as professional design where high performance and user experience are required.
	//
	// example:
	//
	// off
	GpuAcceleration *string `json:"GpuAcceleration,omitempty" xml:"GpuAcceleration,omitempty"`
	HoverConfigMsg  *string `json:"HoverConfigMsg,omitempty" xml:"HoverConfigMsg,omitempty"`
	// The file transfer policy for web clients.
	//
	// example:
	//
	// off
	Html5FileTransfer *string `json:"Html5FileTransfer,omitempty" xml:"Html5FileTransfer,omitempty"`
	// Network communication protocol.
	//
	// example:
	//
	// both
	InternetCommunicationProtocol *string `json:"InternetCommunicationProtocol,omitempty" xml:"InternetCommunicationProtocol,omitempty"`
	// Wuying Keeper toggle for mobile devices
	//
	// example:
	//
	// off
	InternetPrinter *string `json:"InternetPrinter,omitempty" xml:"InternetPrinter,omitempty"`
	// The local disk mapping permissions.
	//
	// example:
	//
	// off
	LocalDrive *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	// Maximum reconnection retry time after an unexpected disconnection. Valid values: 30 to 7200. Unit: seconds.
	//
	// example:
	//
	// 120
	MaxReconnectTime *int32 `json:"MaxReconnectTime,omitempty" xml:"MaxReconnectTime,omitempty"`
	// Memory downclocking duration per process. Valid values: 30 to 120. Unit: seconds.
	//
	// example:
	//
	// 40
	MemoryDownGradeDuration *int32  `json:"MemoryDownGradeDuration,omitempty" xml:"MemoryDownGradeDuration,omitempty"`
	MemoryOverload          *string `json:"MemoryOverload,omitempty" xml:"MemoryOverload,omitempty"`
	// The name of the process.
	MemoryProcessors []*string `json:"MemoryProcessors,omitempty" xml:"MemoryProcessors,omitempty" type:"Repeated"`
	// Memory protection mode switch.
	//
	// example:
	//
	// off
	MemoryProtectedMode *string `json:"MemoryProtectedMode,omitempty" xml:"MemoryProtectedMode,omitempty"`
	// Overall memory usage percentage. Valid values: 70 to 90.
	//
	// example:
	//
	// 70
	MemoryRateLimit *int32 `json:"MemoryRateLimit,omitempty" xml:"MemoryRateLimit,omitempty"`
	// Overall memory sampling duration. Valid values: 30 to 60. Unit: seconds.
	//
	// example:
	//
	// 40
	MemorySampleDuration *int32 `json:"MemorySampleDuration,omitempty" xml:"MemorySampleDuration,omitempty"`
	// Memory usage per process percentage. Valid values: 30 to 60.
	//
	// example:
	//
	// 40
	MemorySingleRateLimit *int32 `json:"MemorySingleRateLimit,omitempty" xml:"MemorySingleRateLimit,omitempty"`
	// Specifies whether to provide the Restart button in the cloud computer floating ball when connecting via mobile clients (Android and iOS clients).
	//
	// > This feature applies only to mobile clients of version 7.4 or later.
	//
	// example:
	//
	// off
	MobileRestart *string `json:"MobileRestart,omitempty" xml:"MobileRestart,omitempty"`
	// The security button toggle for Windows systems on mobile devices
	//
	// example:
	//
	// off
	MobileSafeMenu *string `json:"MobileSafeMenu,omitempty" xml:"MobileSafeMenu,omitempty"`
	// Specifies whether to provide the Shutdown button in the cloud computer floating ball when connecting via mobile clients (Android and iOS clients).
	//
	// > This feature applies only to mobile clients of version 7.4 or later.
	//
	// example:
	//
	// off
	MobileShutdown *string `json:"MobileShutdown,omitempty" xml:"MobileShutdown,omitempty"`
	// Wuying Keeper toggle for mobile devices
	//
	// example:
	//
	// off
	MobileWuyingKeeper *string `json:"MobileWuyingKeeper,omitempty" xml:"MobileWuyingKeeper,omitempty"`
	// Mobile Wy Assistant Toggle
	//
	// example:
	//
	// off
	MobileWyAssistant *string `json:"MobileWyAssistant,omitempty" xml:"MobileWyAssistant,omitempty"`
	ModelLibrary      *string `json:"ModelLibrary,omitempty" xml:"ModelLibrary,omitempty"`
	MultiScreen       *string `json:"MultiScreen,omitempty" xml:"MultiScreen,omitempty"`
	// The policy name.
	//
	// example:
	//
	// testPolicyGroupName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether to enable network redirection.
	//
	// > This feature is in invitational preview and not available to the public.
	//
	// example:
	//
	// on
	NetRedirect *string `json:"NetRedirect,omitempty" xml:"NetRedirect,omitempty"`
	// Network redirection policy details.
	//
	// > This feature is in invitational preview and not available to the public.
	NetRedirectRule []*ModifyCenterPolicyRequestNetRedirectRule `json:"NetRedirectRule,omitempty" xml:"NetRedirectRule,omitempty" type:"Repeated"`
	// Disconnect on inactivity.
	//
	// > This parameter applies only to cloud application policies.
	//
	// example:
	//
	// off
	NoOperationDisconnect *string `json:"NoOperationDisconnect,omitempty" xml:"NoOperationDisconnect,omitempty"`
	// Inactivity disconnect duration. Valid values: 120 to 7200. Unit: seconds.
	//
	// > This parameter applies only to cloud application policies.
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
	// The printer redirection policy. This parameter takes effect only if DeviceRedirects does not include a printer redirection policy.
	//
	// example:
	//
	// off
	PrinterRedirect *string `json:"PrinterRedirect,omitempty" xml:"PrinterRedirect,omitempty"`
	// Specifies whether to enable image quality enhancement for design and 3D applications.
	//
	// example:
	//
	// off
	QualityEnhancement *string `json:"QualityEnhancement,omitempty" xml:"QualityEnhancement,omitempty"`
	// Screen recording duration after an event is detected in audit. Unit: minutes. Valid values: 10 to 60.
	//
	// example:
	//
	// 10
	RecordEventDuration *int32 `json:"RecordEventDuration,omitempty" xml:"RecordEventDuration,omitempty"`
	// File extensions for screen recording events
	RecordEventFileExts []*string `json:"RecordEventFileExts,omitempty" xml:"RecordEventFileExts,omitempty" type:"Repeated"`
	// Absolute paths for file monitoring in screen recording audit.
	RecordEventFilePaths []*string `json:"RecordEventFilePaths,omitempty" xml:"RecordEventFilePaths,omitempty" type:"Repeated"`
	// Levels of screen recording events
	RecordEventLevels []*ModifyCenterPolicyRequestRecordEventLevels `json:"RecordEventLevels,omitempty" xml:"RecordEventLevels,omitempty" type:"Repeated"`
	// Absolute paths for registry monitoring in screen recording audit.
	//
	// example:
	//
	// Computer\\HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Services\\USBSTOR
	RecordEventRegisters []*string `json:"RecordEventRegisters,omitempty" xml:"RecordEventRegisters,omitempty" type:"Repeated"`
	// List of screen recording events.
	RecordEvents []*string `json:"RecordEvents,omitempty" xml:"RecordEvents,omitempty" type:"Repeated"`
	// Specifies whether to enable screen recording.
	//
	// example:
	//
	// off
	Recording *string `json:"Recording,omitempty" xml:"Recording,omitempty"`
	// Cloud computer audio recording option.
	//
	// example:
	//
	// on
	RecordingAudio *string `json:"RecordingAudio,omitempty" xml:"RecordingAudio,omitempty"`
	// Screen recording file viewing duration in minutes. Recording files are automatically split based on the specified duration and uploaded to the storage bucket. When a file reaches 300 MB, it is prioritized for rolling updates. Valid values: 10 to 60.
	//
	// example:
	//
	// 10
	RecordingDuration *int32 `json:"RecordingDuration,omitempty" xml:"RecordingDuration,omitempty"`
	// Screen recording end time in HH:MM:SS format. This parameter is meaningful only when `Recording` is set to `PERIOD`.
	//
	// example:
	//
	// 08:59:00
	RecordingEndTime *string `json:"RecordingEndTime,omitempty" xml:"RecordingEndTime,omitempty"`
	// Retention period of screen recording files. Valid values: 1 to 180 days.
	//
	// example:
	//
	// 15
	RecordingExpires *int32 `json:"RecordingExpires,omitempty" xml:"RecordingExpires,omitempty"`
	// Screen recording frame rate. Unit: FPS.
	//
	// example:
	//
	// 5
	RecordingFps *string `json:"RecordingFps,omitempty" xml:"RecordingFps,omitempty"`
	// Screen recording start time in HH:MM:SS format. This parameter is meaningful only when `Recording` is set to `PERIOD`.
	//
	// example:
	//
	// 08:00:00
	RecordingStartTime *string `json:"RecordingStartTime,omitempty" xml:"RecordingStartTime,omitempty"`
	// Notify end users when screen recording is enabled.
	//
	// example:
	//
	// off
	RecordingUserNotify *string `json:"RecordingUserNotify,omitempty" xml:"RecordingUserNotify,omitempty"`
	// Message to notify end users when screen recording is enabled.
	//
	// example:
	//
	// Screen recording is enabled.
	RecordingUserNotifyMessage *string `json:"RecordingUserNotifyMessage,omitempty" xml:"RecordingUserNotifyMessage,omitempty"`
	// The region ID. Set the value to `cn-shanghai`.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Keyboard and mouse control permissions for remote assistance.
	//
	// example:
	//
	// fullControl
	RemoteCoordinate *string `json:"RemoteCoordinate,omitempty" xml:"RemoteCoordinate,omitempty"`
	// Reset cloud computer.
	//
	// example:
	//
	// off
	ResetDesktop  *string `json:"ResetDesktop,omitempty" xml:"ResetDesktop,omitempty"`
	ResolutionDpi *int32  `json:"ResolutionDpi,omitempty" xml:"ResolutionDpi,omitempty"`
	// Resolution height. Unit: pixels. Valid values for cloud applications: 500 to 50000. Valid values for cloud computers: 480 to 4096.
	//
	// example:
	//
	// 1280
	ResolutionHeight *int32 `json:"ResolutionHeight,omitempty" xml:"ResolutionHeight,omitempty"`
	// Resolution type.
	//
	// example:
	//
	// adaptive
	ResolutionModel *string `json:"ResolutionModel,omitempty" xml:"ResolutionModel,omitempty"`
	// Resolution width. Unit: pixels. Valid values for cloud applications: 500 to 50000. Valid values for cloud computers: 480 to 4096.
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
	// The list of client IP address whitelists to delete.
	RevokeAccessPolicyRule []*ModifyCenterPolicyRequestRevokeAccessPolicyRule `json:"RevokeAccessPolicyRule,omitempty" xml:"RevokeAccessPolicyRule,omitempty" type:"Repeated"`
	// List of security group control rules to delete.
	RevokeSecurityPolicyRule []*ModifyCenterPolicyRequestRevokeSecurityPolicyRule `json:"RevokeSecurityPolicyRule,omitempty" xml:"RevokeSecurityPolicyRule,omitempty" type:"Repeated"`
	// Security button toggle for Windows systems on mobile devices
	//
	// example:
	//
	// off
	SafeMenu *string `json:"SafeMenu,omitempty" xml:"SafeMenu,omitempty"`
	// The effective scope of the policy.
	//
	// example:
	//
	// GLOBAL
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// Specify when `Scope` is set to `IP`. Takes effect only when `Scope` is `IP`.
	ScopeValue []*string `json:"ScopeValue,omitempty" xml:"ScopeValue,omitempty" type:"Repeated"`
	// Xiao Ying toggle for mobile devices
	//
	// example:
	//
	// off
	ScreenDisplayMode *string `json:"ScreenDisplayMode,omitempty" xml:"ScreenDisplayMode,omitempty"`
	// Maximum session bandwidth throttling value. Unit: Kbps. Valid values: 2000 to 100000.
	//
	// example:
	//
	// 2000
	SessionMaxRateKbps *int32 `json:"SessionMaxRateKbps,omitempty" xml:"SessionMaxRateKbps,omitempty"`
	// Specifies whether to enable smoothness enhancement for daily office scenarios.
	//
	// example:
	//
	// off
	SmoothEnhancement *string `json:"SmoothEnhancement,omitempty" xml:"SmoothEnhancement,omitempty"`
	// Specifies whether to provide the status monitoring entry in the cloud computer floating ball.
	//
	// example:
	//
	// off
	StatusMonitor *string `json:"StatusMonitor,omitempty" xml:"StatusMonitor,omitempty"`
	// Streaming mode adaptation scenario.
	//
	// example:
	//
	// smooth
	StreamingMode *string `json:"StreamingMode,omitempty" xml:"StreamingMode,omitempty"`
	// Target frame rate. Valid values: 10 to 60.
	//
	// example:
	//
	// 30
	TargetFps *int32 `json:"TargetFps,omitempty" xml:"TargetFps,omitempty"`
	// Application taskbar.
	//
	// > This parameter applies only to cloud application policies.
	//
	// example:
	//
	// off
	Taskbar *string `json:"Taskbar,omitempty" xml:"Taskbar,omitempty"`
	// USB redirection.
	//
	// example:
	//
	// off
	UsbRedirect *string `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty"`
	// USB redirection rules.
	UsbSupplyRedirectRule []*ModifyCenterPolicyRequestUsbSupplyRedirectRule `json:"UsbSupplyRedirectRule,omitempty" xml:"UsbSupplyRedirectRule,omitempty" type:"Repeated"`
	UseTime               *string                                           `json:"UseTime,omitempty" xml:"UseTime,omitempty"`
	// Average bitrate for video encoding. Unit: Kbps. Valid values: 1000 to 50000.
	//
	// example:
	//
	// 2000
	VideoEncAvgKbps *int32 `json:"VideoEncAvgKbps,omitempty" xml:"VideoEncAvgKbps,omitempty"`
	// Maximum QP for video encoding, representing lowest quality. Valid values: 0 to 51.
	//
	// example:
	//
	// 30
	VideoEncMaxQP *int32 `json:"VideoEncMaxQP,omitempty" xml:"VideoEncMaxQP,omitempty"`
	// Minimum QP for video encoding, representing highest quality. Valid values: 0 to 51.
	//
	// example:
	//
	// 30
	VideoEncMinQP *int32 `json:"VideoEncMinQP,omitempty" xml:"VideoEncMinQP,omitempty"`
	// Peak bitrate for video encoding. Unit: Kbps. Valid values: 1000 to 50000.
	//
	// example:
	//
	// 2000
	VideoEncPeakKbps *int32 `json:"VideoEncPeakKbps,omitempty" xml:"VideoEncPeakKbps,omitempty"`
	// Video encoding policy.
	//
	// example:
	//
	// qualityFirst
	VideoEncPolicy *string `json:"VideoEncPolicy,omitempty" xml:"VideoEncPolicy,omitempty"`
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
	// Watermark.
	//
	// example:
	//
	// off
	Watermark *string `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
	// Blind watermark anti-photo feature.
	//
	// example:
	//
	// off
	WatermarkAntiCam *string `json:"WatermarkAntiCam,omitempty" xml:"WatermarkAntiCam,omitempty"`
	// Watermark font color. Valid values: 0 to 16777215.
	//
	// example:
	//
	// 0
	WatermarkColor *int32 `json:"WatermarkColor,omitempty" xml:"WatermarkColor,omitempty"`
	// Number of watermark columns. Valid values: 3 to 10.
	//
	// example:
	//
	// 3
	WatermarkColumnAmount *int32 `json:"WatermarkColumnAmount,omitempty" xml:"WatermarkColumnAmount,omitempty"`
	// If you set `WatermarkType` to `custom`, you must also specify custom text using the `WatermarkCustomText` parameter.
	//
	// example:
	//
	// test
	WatermarkCustomText *string `json:"WatermarkCustomText,omitempty" xml:"WatermarkCustomText,omitempty"`
	// Watermark tilt angle. Valid values: -10 to -30.
	//
	// example:
	//
	// -10
	WatermarkDegree *float64 `json:"WatermarkDegree,omitempty" xml:"WatermarkDegree,omitempty"`
	// Watermark font size. Valid values: 10 to 20.
	//
	// example:
	//
	// 10
	WatermarkFontSize *int32 `json:"WatermarkFontSize,omitempty" xml:"WatermarkFontSize,omitempty"`
	// Watermark font style.
	//
	// example:
	//
	// plain
	WatermarkFontStyle *string `json:"WatermarkFontStyle,omitempty" xml:"WatermarkFontStyle,omitempty"`
	// Blind watermark enhancement feature.
	//
	// example:
	//
	// medium
	WatermarkPower *string `json:"WatermarkPower,omitempty" xml:"WatermarkPower,omitempty"`
	// Number of watermark rows. Valid values: 3 to 10.
	//
	// example:
	//
	// 5
	WatermarkRowAmount *int32 `json:"WatermarkRowAmount,omitempty" xml:"WatermarkRowAmount,omitempty"`
	// Blind watermark security priority rule.
	//
	// example:
	//
	// off
	WatermarkSecurity *string `json:"WatermarkSecurity,omitempty" xml:"WatermarkSecurity,omitempty"`
	WatermarkShadow   *string `json:"WatermarkShadow,omitempty" xml:"WatermarkShadow,omitempty"`
	// Watermark opacity. A higher value means lower transparency. Valid values: 10 to 100.
	//
	// example:
	//
	// 10
	WatermarkTransparencyValue *int32 `json:"WatermarkTransparencyValue,omitempty" xml:"WatermarkTransparencyValue,omitempty"`
	// The watermark type. You can select up to three types, separated by commas.
	//
	// > If you set this parameter to `custom`, you must also specify custom text using the `WatermarkCustomText` parameter.
	//
	// example:
	//
	// EndUserId
	WatermarkType *string `json:"WatermarkType,omitempty" xml:"WatermarkType,omitempty"`
	// Wuying Keeper switch.
	//
	// example:
	//
	// off
	WuyingKeeper *string `json:"WuyingKeeper,omitempty" xml:"WuyingKeeper,omitempty"`
	// Specifies whether to provide the Wuying AI Assistant entry in the cloud computer floating ball when connecting via desktop clients (including Windows and macOS clients).
	//
	// > This feature applies only to desktop clients of version 7.7 or later.
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
	// Client access IP address range. Specify an IPv4 CIDr block.
	//
	// example:
	//
	// 47.100.XX.XX/16
	CidrIp *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	// Description of the client IP address whitelist.
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
	// Target of the security group control rule. Specify an IPv4 CIDR block.
	//
	// example:
	//
	// 47.100.XX.XX/16
	CidrIp *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	// Description of the security group control rule.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Protocol type for the security group control rule.
	//
	// example:
	//
	// TCP
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// Authorization policy for the security group control rule.
	//
	// example:
	//
	// accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// Port range for the security group control rule. The port range depends on the protocol (IpProtocol):
	//
	// - For TCP or UDP protocols: Port range is 1 to 65535. Separate the start and end ports with a forward slash (/). Example: 1/200.
	//
	// - For ICMP protocol: Set to -1/-1.
	//
	// - For GRE protocol: Set to -1/-1.
	//
	// - When IpProtocol is set to all: Set to -1/-1.
	//
	// For more information about common ports for typical applications, see [Common ports](https://help.aliyun.com/document_detail/40724.html).
	//
	// example:
	//
	// 22/22
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// Priority of the security group control rule. A smaller number indicates a higher priority. Valid values: 1 to 60. Default value: 1.
	//
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// Rule direction for the security group control rule.
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
	// Client login control. Specifies the client type.
	//
	// example:
	//
	// android
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// Client login control. Specifies whether to allow using a specific client type to log on to cloud computers.
	//
	// > If you don\\"t configure `ClientType` parameters, all client types are allowed by default.
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
	ClipboardSize        *int32  `json:"ClipboardSize,omitempty" xml:"ClipboardSize,omitempty"`
	ClipboardSizeUnit    *string `json:"ClipboardSizeUnit,omitempty" xml:"ClipboardSizeUnit,omitempty"`
	ClipboardType        *string `json:"ClipboardType,omitempty" xml:"ClipboardType,omitempty"`
	GrainedType          *string `json:"GrainedType,omitempty" xml:"GrainedType,omitempty"`
	InClipboardSize      *int32  `json:"InClipboardSize,omitempty" xml:"InClipboardSize,omitempty"`
	InClipboardSizeUnit  *string `json:"InClipboardSizeUnit,omitempty" xml:"InClipboardSizeUnit,omitempty"`
	OutClipboardSize     *int32  `json:"OutClipboardSize,omitempty" xml:"OutClipboardSize,omitempty"`
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
	// Peripheral type.
	//
	// example:
	//
	// camera
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// Redirection type.
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
	// Device name.
	//
	// example:
	//
	// sandisk
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// Product ID.
	//
	// example:
	//
	// 0x55b1
	DevicePid *string `json:"DevicePid,omitempty" xml:"DevicePid,omitempty"`
	// Peripheral type.
	//
	// example:
	//
	// storage
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// Vendor ID. For more information, see [Valid USB Vendor IDs (VIDs)](https://www.usb.org/sites/default/files/vendor_ids032322.pdf_1.pdf).
	//
	// example:
	//
	// 0x0781
	DeviceVid *string `json:"DeviceVid,omitempty" xml:"DeviceVid,omitempty"`
	// Link optimization command.
	//
	// example:
	//
	// 2:0
	OptCommand *string `json:"OptCommand,omitempty" xml:"OptCommand,omitempty"`
	Platforms  *string `json:"Platforms,omitempty" xml:"Platforms,omitempty"`
	// Redirection type.
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
	// Policy description.
	//
	// example:
	//
	// Policy description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Domain name.
	//
	// example:
	//
	// *.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Resolution policy.
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
	// Domain name.
	//
	// example:
	//
	// *.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Redirection policy.
	//
	// example:
	//
	// Allow
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// Rule type.
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
	// Event level
	//
	// example:
	//
	// HIGH
	EventLevel *string `json:"EventLevel,omitempty" xml:"EventLevel,omitempty"`
	// Event type
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
	// Client access IP address range to delete. Specify an IPv4 CIDR block.
	//
	// example:
	//
	// 47.100.XX.XX/16
	CidrIp *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	// Description of the client IP address whitelist to delete.
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
	// Target of the security group control rule to delete. Specify an IPv4 CIDR block.
	//
	// example:
	//
	// 47.100.XX.XX/16
	CidrIp *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	// Description of the security group control rule to delete.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Protocol type for the security group control rule to delete.
	//
	// example:
	//
	// TCP
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// Authorization policy for the security group control rule to delete.
	//
	// example:
	//
	// accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// Port range for the security group control rule to delete. The port range depends on the protocol (IpProtocol):
	//
	// - For TCP or UDP protocols: Port range is 1 to 65535. Separate the start and end ports with a forward slash (/). Example: 1/200.
	//
	// - For ICMP protocol: Set to -1/-1.
	//
	// - For GRE protocol: Set to -1/-1.
	//
	// - When IpProtocol is set to all: Set to -1/-1.
	//
	// For more information about common ports for typical applications, see [Common ports](https://help.aliyun.com/document_detail/40724.html).
	//
	// example:
	//
	// 22/22
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// Priority of the security group control rule to delete. A smaller number indicates a higher priority. Valid values: 1 to 60. Default value: 1.
	//
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// Rule direction for the security group control rule to delete.
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
	// Rule description.
	//
	// example:
	//
	// Test rule
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Product ID.
	//
	// example:
	//
	// 08**
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// USB redirection type.
	//
	// example:
	//
	// 1
	UsbRedirectType *string `json:"UsbRedirectType,omitempty" xml:"UsbRedirectType,omitempty"`
	// USB redirection rule type.
	//
	// example:
	//
	// 1
	UsbRuleType *string `json:"UsbRuleType,omitempty" xml:"UsbRuleType,omitempty"`
	// Vendor ID. For more information, see [Valid USB Vendor IDs (VIDs)](https://www.usb.org/sites/default/files/vendor_ids032322.pdf_1.pdf).
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
