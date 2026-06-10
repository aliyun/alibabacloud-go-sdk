// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCenterPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcademicProxy(v string) *CreateCenterPolicyRequest
	GetAcademicProxy() *string
	SetAdminAccess(v string) *CreateCenterPolicyRequest
	GetAdminAccess() *string
	SetAdminKeyboardOnFullScreen(v string) *CreateCenterPolicyRequest
	GetAdminKeyboardOnFullScreen() *string
	SetAdminKeyboardOnWindows(v string) *CreateCenterPolicyRequest
	GetAdminKeyboardOnWindows() *string
	SetAppContentProtection(v string) *CreateCenterPolicyRequest
	GetAppContentProtection() *string
	SetAuthorizeAccessPolicyRule(v []*CreateCenterPolicyRequestAuthorizeAccessPolicyRule) *CreateCenterPolicyRequest
	GetAuthorizeAccessPolicyRule() []*CreateCenterPolicyRequestAuthorizeAccessPolicyRule
	SetAuthorizeSecurityPolicyRule(v []*CreateCenterPolicyRequestAuthorizeSecurityPolicyRule) *CreateCenterPolicyRequest
	GetAuthorizeSecurityPolicyRule() []*CreateCenterPolicyRequestAuthorizeSecurityPolicyRule
	SetAutoReconnect(v string) *CreateCenterPolicyRequest
	GetAutoReconnect() *string
	SetBusinessChannel(v string) *CreateCenterPolicyRequest
	GetBusinessChannel() *string
	SetBusinessType(v int32) *CreateCenterPolicyRequest
	GetBusinessType() *int32
	SetCameraRedirect(v string) *CreateCenterPolicyRequest
	GetCameraRedirect() *string
	SetClientControlMenu(v string) *CreateCenterPolicyRequest
	GetClientControlMenu() *string
	SetClientCreateSnapshot(v string) *CreateCenterPolicyRequest
	GetClientCreateSnapshot() *string
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
	SetCpuOverload(v string) *CreateCenterPolicyRequest
	GetCpuOverload() *string
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
	SetDescription(v string) *CreateCenterPolicyRequest
	GetDescription() *string
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
	SetDiskOverload(v string) *CreateCenterPolicyRequest
	GetDiskOverload() *string
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
	SetExternalDrive(v string) *CreateCenterPolicyRequest
	GetExternalDrive() *string
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
	SetHoverConfigMsg(v string) *CreateCenterPolicyRequest
	GetHoverConfigMsg() *string
	SetHtml5FileTransfer(v string) *CreateCenterPolicyRequest
	GetHtml5FileTransfer() *string
	SetInternetCommunicationProtocol(v string) *CreateCenterPolicyRequest
	GetInternetCommunicationProtocol() *string
	SetInternetPrinter(v string) *CreateCenterPolicyRequest
	GetInternetPrinter() *string
	SetKeyboardControl(v string) *CreateCenterPolicyRequest
	GetKeyboardControl() *string
	SetLocalDrive(v string) *CreateCenterPolicyRequest
	GetLocalDrive() *string
	SetMaxReconnectTime(v int32) *CreateCenterPolicyRequest
	GetMaxReconnectTime() *int32
	SetMemoryDownGradeDuration(v int32) *CreateCenterPolicyRequest
	GetMemoryDownGradeDuration() *int32
	SetMemoryOverload(v string) *CreateCenterPolicyRequest
	GetMemoryOverload() *string
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
	SetModelLibrary(v string) *CreateCenterPolicyRequest
	GetModelLibrary() *string
	SetMultiScreen(v string) *CreateCenterPolicyRequest
	GetMultiScreen() *string
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
	SetPortProxy(v string) *CreateCenterPolicyRequest
	GetPortProxy() *string
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
	SetResolutionDpi(v int32) *CreateCenterPolicyRequest
	GetResolutionDpi() *int32
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
	SetWatermarkShadow(v string) *CreateCenterPolicyRequest
	GetWatermarkShadow() *string
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
	AcademicProxy *string `json:"AcademicProxy,omitempty" xml:"AcademicProxy,omitempty"`
	// Whether users have administrative permission after logging on to cloud desktops.
	//
	// > This feature is in invitational preview and is not publicly available.
	//
	// example:
	//
	// deny
	AdminAccess               *string `json:"AdminAccess,omitempty" xml:"AdminAccess,omitempty"`
	AdminKeyboardOnFullScreen *string `json:"AdminKeyboardOnFullScreen,omitempty" xml:"AdminKeyboardOnFullScreen,omitempty"`
	AdminKeyboardOnWindows    *string `json:"AdminKeyboardOnWindows,omitempty" xml:"AdminKeyboardOnWindows,omitempty"`
	// Enable screenshot prevention.
	//
	// example:
	//
	// off
	AppContentProtection *string `json:"AppContentProtection,omitempty" xml:"AppContentProtection,omitempty"`
	// Client IP address whitelist. After this parameter is configured, only IP addresses within the specified CIDR blocks can access cloud desktops.
	AuthorizeAccessPolicyRule []*CreateCenterPolicyRequestAuthorizeAccessPolicyRule `json:"AuthorizeAccessPolicyRule,omitempty" xml:"AuthorizeAccessPolicyRule,omitempty" type:"Repeated"`
	// List of security group control rules.
	AuthorizeSecurityPolicyRule []*CreateCenterPolicyRequestAuthorizeSecurityPolicyRule `json:"AuthorizeSecurityPolicyRule,omitempty" xml:"AuthorizeSecurityPolicyRule,omitempty" type:"Repeated"`
	// Client auto-reconnect switch
	//
	// example:
	//
	// off
	AutoReconnect   *string `json:"AutoReconnect,omitempty" xml:"AutoReconnect,omitempty"`
	BusinessChannel *string `json:"BusinessChannel,omitempty" xml:"BusinessChannel,omitempty"`
	// Business type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	BusinessType *int32 `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// Local camera redirection. This parameter takes effect only when no local camera redirection policy is configured in DeviceRedirects.
	//
	// example:
	//
	// off
	CameraRedirect       *string `json:"CameraRedirect,omitempty" xml:"CameraRedirect,omitempty"`
	ClientControlMenu    *string `json:"ClientControlMenu,omitempty" xml:"ClientControlMenu,omitempty"`
	ClientCreateSnapshot *string `json:"ClientCreateSnapshot,omitempty" xml:"ClientCreateSnapshot,omitempty"`
	// List of logon method control rules. These rules control which clients can access cloud desktops.
	ClientType []*CreateCenterPolicyRequestClientType `json:"ClientType,omitempty" xml:"ClientType,omitempty" type:"Repeated"`
	// Clipboard permission.
	//
	// example:
	//
	// off
	Clipboard *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	// Fine-grained clipboard control configuration
	ClipboardGraineds []*CreateCenterPolicyRequestClipboardGraineds `json:"ClipboardGraineds,omitempty" xml:"ClipboardGraineds,omitempty" type:"Repeated"`
	// Clipboard scope
	//
	// example:
	//
	// GLOBAL
	ClipboardScope *string `json:"ClipboardScope,omitempty" xml:"ClipboardScope,omitempty"`
	// Enable color enhancement for design and 3D application scenarios.
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
	// 30
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
	// CPU overall usage percentage. Valid values: 70 to 90.
	//
	// example:
	//
	// 70
	CpuRateLimit *int32 `json:"CpuRateLimit,omitempty" xml:"CpuRateLimit,omitempty"`
	// CPU overall sampling duration. Valid values: 10 to 60. Unit: seconds.
	//
	// example:
	//
	// 60
	CpuSampleDuration *int32 `json:"CpuSampleDuration,omitempty" xml:"CpuSampleDuration,omitempty"`
	// CPU single-core usage percentage. Valid values: 70 to 100.
	//
	// example:
	//
	// 70
	CpuSingleRateLimit *int32  `json:"CpuSingleRateLimit,omitempty" xml:"CpuSingleRateLimit,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Peripheral connection hint control.
	//
	// example:
	//
	// off
	DeviceConnectHint *string `json:"DeviceConnectHint,omitempty" xml:"DeviceConnectHint,omitempty"`
	// List of device redirection rules.
	DeviceRedirects []*CreateCenterPolicyRequestDeviceRedirects `json:"DeviceRedirects,omitempty" xml:"DeviceRedirects,omitempty" type:"Repeated"`
	// List of custom peripheral rules.
	DeviceRules []*CreateCenterPolicyRequestDeviceRules `json:"DeviceRules,omitempty" xml:"DeviceRules,omitempty" type:"Repeated"`
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
	// 30
	DisconnectKeepSessionTime *int32  `json:"DisconnectKeepSessionTime,omitempty" xml:"DisconnectKeepSessionTime,omitempty"`
	DiskOverload              *string `json:"DiskOverload,omitempty" xml:"DiskOverload,omitempty"`
	// Display mode.
	//
	// example:
	//
	// clientCustom
	DisplayMode *string `json:"DisplayMode,omitempty" xml:"DisplayMode,omitempty"`
	// Domain name resolution policy.
	DomainResolveRule []*CreateCenterPolicyRequestDomainResolveRule `json:"DomainResolveRule,omitempty" xml:"DomainResolveRule,omitempty" type:"Repeated"`
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
	// Users in the same office network share cloud desktops.
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
	FileMigrate               *string `json:"FileMigrate,omitempty" xml:"FileMigrate,omitempty"`
	FileTransferAddress       *string `json:"FileTransferAddress,omitempty" xml:"FileTransferAddress,omitempty"`
	FileTransferSpeed         *string `json:"FileTransferSpeed,omitempty" xml:"FileTransferSpeed,omitempty"`
	FileTransferSpeedLocation *string `json:"FileTransferSpeedLocation,omitempty" xml:"FileTransferSpeedLocation,omitempty"`
	// Enable image quality policy for graphics-intensive cloud desktops. Enable this policy if you require high performance and user experience, such as in professional design scenarios.
	//
	// example:
	//
	// off
	GpuAcceleration *string `json:"GpuAcceleration,omitempty" xml:"GpuAcceleration,omitempty"`
	HoverConfigMsg  *string `json:"HoverConfigMsg,omitempty" xml:"HoverConfigMsg,omitempty"`
	// Web client file transfer policy.
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
	InternetPrinter               *string `json:"InternetPrinter,omitempty" xml:"InternetPrinter,omitempty"`
	KeyboardControl               *string `json:"KeyboardControl,omitempty" xml:"KeyboardControl,omitempty"`
	// Local disk mapping permission.
	//
	// example:
	//
	// off
	LocalDrive *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	// Maximum reconnection retry time when a cloud desktop disconnects due to objective reasons. Valid values: 30 to 7200. Unit: seconds.
	//
	// example:
	//
	// 120
	MaxReconnectTime *int32 `json:"MaxReconnectTime,omitempty" xml:"MaxReconnectTime,omitempty"`
	// Memory downclocking duration per worker. Valid values: 30 to 120. Unit: seconds.
	//
	// example:
	//
	// 40
	MemoryDownGradeDuration *int32  `json:"MemoryDownGradeDuration,omitempty" xml:"MemoryDownGradeDuration,omitempty"`
	MemoryOverload          *string `json:"MemoryOverload,omitempty" xml:"MemoryOverload,omitempty"`
	// Process names.
	MemoryProcessors []*string `json:"MemoryProcessors,omitempty" xml:"MemoryProcessors,omitempty" type:"Repeated"`
	// Memory protection mode switch.
	//
	// example:
	//
	// off
	MemoryProtectedMode *string `json:"MemoryProtectedMode,omitempty" xml:"MemoryProtectedMode,omitempty"`
	// Memory overall usage percentage. Valid values: 70 to 90.
	//
	// example:
	//
	// 70
	MemoryRateLimit *int32 `json:"MemoryRateLimit,omitempty" xml:"MemoryRateLimit,omitempty"`
	// Memory overall sampling duration. Valid values: 30 to 60. Unit: seconds.
	//
	// example:
	//
	// 40
	MemorySampleDuration *int32 `json:"MemorySampleDuration,omitempty" xml:"MemorySampleDuration,omitempty"`
	// Memory usage percentage per worker. Valid values: 30 to 60.
	//
	// example:
	//
	// 40
	MemorySingleRateLimit *int32 `json:"MemorySingleRateLimit,omitempty" xml:"MemorySingleRateLimit,omitempty"`
	// Provide a restart button in the floating ball on the cloud desktop when connecting through mobile clients (Android and iOS clients).
	//
	// > This feature applies only to mobile clients V7.4 or later.
	//
	// example:
	//
	// off
	MobileRestart *string `json:"MobileRestart,omitempty" xml:"MobileRestart,omitempty"`
	// Mobile Windows security control switch
	//
	// example:
	//
	// off
	MobileSafeMenu *string `json:"MobileSafeMenu,omitempty" xml:"MobileSafeMenu,omitempty"`
	// Provide a shutdown button in the floating ball on the cloud desktop when connecting through mobile clients (Android and iOS clients).
	//
	// > This feature applies only to mobile clients V7.4 or later.
	//
	// example:
	//
	// off
	MobileShutdown *string `json:"MobileShutdown,omitempty" xml:"MobileShutdown,omitempty"`
	// Mobile Wuying Keeper switch
	//
	// example:
	//
	// off
	MobileWuyingKeeper *string `json:"MobileWuyingKeeper,omitempty" xml:"MobileWuyingKeeper,omitempty"`
	// Mobile Xiao Ying switch
	//
	// example:
	//
	// off
	MobileWyAssistant *string `json:"MobileWyAssistant,omitempty" xml:"MobileWyAssistant,omitempty"`
	ModelLibrary      *string `json:"ModelLibrary,omitempty" xml:"ModelLibrary,omitempty"`
	MultiScreen       *string `json:"MultiScreen,omitempty" xml:"MultiScreen,omitempty"`
	// Policy name.
	//
	// This parameter is required.
	//
	// example:
	//
	// testPolicyGroupName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Enable network redirection.
	//
	// > This feature is in invitational preview and is not publicly available.
	//
	// example:
	//
	// off
	NetRedirect *string `json:"NetRedirect,omitempty" xml:"NetRedirect,omitempty"`
	// Network redirection policy details.
	//
	// > This feature is in invitational preview and is not publicly available.
	NetRedirectRule []*CreateCenterPolicyRequestNetRedirectRule `json:"NetRedirectRule,omitempty" xml:"NetRedirectRule,omitempty" type:"Repeated"`
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
	NoOperationDisconnectTime *int32  `json:"NoOperationDisconnectTime,omitempty" xml:"NoOperationDisconnectTime,omitempty"`
	PortProxy                 *string `json:"PortProxy,omitempty" xml:"PortProxy,omitempty"`
	// Printer redirection policy. This parameter takes effect only when no printer redirection policy is configured in DeviceRedirects.
	//
	// example:
	//
	// off
	PrinterRedirect *string `json:"PrinterRedirect,omitempty" xml:"PrinterRedirect,omitempty"`
	// Enable image quality enhancement for design and 3D application scenarios.
	//
	// example:
	//
	// off
	QualityEnhancement *string `json:"QualityEnhancement,omitempty" xml:"QualityEnhancement,omitempty"`
	// Duration of screen recording after an event is detected in screen recording audit. Unit: minutes. Valid values: 10 to 60.
	//
	// example:
	//
	// 10
	RecordEventDuration *int32 `json:"RecordEventDuration,omitempty" xml:"RecordEventDuration,omitempty"`
	// File name extensions for screen recording events
	RecordEventFileExts []*string `json:"RecordEventFileExts,omitempty" xml:"RecordEventFileExts,omitempty" type:"Repeated"`
	// Absolute paths of files to monitor in screen recording audit.
	RecordEventFilePaths []*string `json:"RecordEventFilePaths,omitempty" xml:"RecordEventFilePaths,omitempty" type:"Repeated"`
	// Levels of screen recording events
	RecordEventLevels []*CreateCenterPolicyRequestRecordEventLevels `json:"RecordEventLevels,omitempty" xml:"RecordEventLevels,omitempty" type:"Repeated"`
	// Absolute paths of registry keys to monitor in screen recording audit.
	//
	// example:
	//
	// Computer\\HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Services\\USBSTOR
	RecordEventRegisters []*string `json:"RecordEventRegisters,omitempty" xml:"RecordEventRegisters,omitempty" type:"Repeated"`
	// List of screen recording events.
	RecordEvents []*string `json:"RecordEvents,omitempty" xml:"RecordEvents,omitempty" type:"Repeated"`
	// Enable screen recording.
	//
	// example:
	//
	// off
	Recording *string `json:"Recording,omitempty" xml:"Recording,omitempty"`
	// Options for recording cloud desktop audio.
	//
	// example:
	//
	// on
	RecordingAudio *string `json:"RecordingAudio,omitempty" xml:"RecordingAudio,omitempty"`
	// Duration for viewing recorded files. Unit: minutes. Recorded files are automatically split and uploaded to the bucket based on the specified duration. When a file reaches 300 MB, it is rolled over for updates. Valid values: 10 to 60.
	//
	// example:
	//
	// 10
	RecordingDuration *int32 `json:"RecordingDuration,omitempty" xml:"RecordingDuration,omitempty"`
	// Screen recording end time. Format: HH:MM:SS. This parameter takes effect only when `Recording` is set to `PERIOD`.
	//
	// example:
	//
	// 08:59:00
	RecordingEndTime *string `json:"RecordingEndTime,omitempty" xml:"RecordingEndTime,omitempty"`
	// Retention period for recorded files. Valid values: 1 to 180 days.
	//
	// example:
	//
	// 15
	RecordingExpires *int32 `json:"RecordingExpires,omitempty" xml:"RecordingExpires,omitempty"`
	// Screen recording frame rate. Unit: FPS.
	//
	// example:
	//
	// 2
	RecordingFps *string `json:"RecordingFps,omitempty" xml:"RecordingFps,omitempty"`
	// Screen recording start time. Format: HH:MM:SS. This parameter takes effect only when `Recording` is set to `PERIOD`.
	//
	// example:
	//
	// 08:00:00
	RecordingStartTime *string `json:"RecordingStartTime,omitempty" xml:"RecordingStartTime,omitempty"`
	// Notify end users that screen recording is enabled.
	//
	// example:
	//
	// off
	RecordingUserNotify *string `json:"RecordingUserNotify,omitempty" xml:"RecordingUserNotify,omitempty"`
	// Message to notify end users that screen recording is enabled.
	//
	// example:
	//
	// Screen recording is enabled.
	RecordingUserNotifyMessage *string `json:"RecordingUserNotifyMessage,omitempty" xml:"RecordingUserNotifyMessage,omitempty"`
	// Region ID. This feature is region-independent. Set this parameter to cn-shanghai.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Mouse and keyboard control permissions during remote assistance.
	//
	// example:
	//
	// fullControl
	RemoteCoordinate *string `json:"RemoteCoordinate,omitempty" xml:"RemoteCoordinate,omitempty"`
	// Reset cloud desktop.
	//
	// example:
	//
	// off
	ResetDesktop  *string `json:"ResetDesktop,omitempty" xml:"ResetDesktop,omitempty"`
	ResolutionDpi *int32  `json:"ResolutionDpi,omitempty" xml:"ResolutionDpi,omitempty"`
	// Resolution height. Unit: pixels. Valid values for cloud applications: 500 to 50000. Valid values for cloud desktops: 480 to 4096.
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
	// Resolution width. Unit: pixels. Valid values for cloud applications: 500 to 50000. Valid values for cloud desktops: 640 to 4096.
	//
	// example:
	//
	// 720
	ResolutionWidth *int32 `json:"ResolutionWidth,omitempty" xml:"ResolutionWidth,omitempty"`
	// Resource type.
	//
	// This parameter is required.
	//
	// example:
	//
	// desktop
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SafeMenu     *string `json:"SafeMenu,omitempty" xml:"SafeMenu,omitempty"`
	// Policy scope.
	//
	// example:
	//
	// GLOBAL
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// Specify this parameter when `Scope` is set to `IP`. This parameter takes effect only when `Scope` is set to `IP`.
	ScopeValue        []*string `json:"ScopeValue,omitempty" xml:"ScopeValue,omitempty" type:"Repeated"`
	ScreenDisplayMode *string   `json:"ScreenDisplayMode,omitempty" xml:"ScreenDisplayMode,omitempty"`
	// Maximum session bandwidth throttling rate. Unit: Kbps. Valid values: 2000 to 100000.
	//
	// example:
	//
	// 2000
	SessionMaxRateKbps *int32 `json:"SessionMaxRateKbps,omitempty" xml:"SessionMaxRateKbps,omitempty"`
	// Enable smoothness enhancement for daily office scenarios.
	//
	// example:
	//
	// off
	SmoothEnhancement *string `json:"SmoothEnhancement,omitempty" xml:"SmoothEnhancement,omitempty"`
	// Provide an entry point for status monitoring in the floating ball on the cloud desktop.
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
	UsbSupplyRedirectRule []*CreateCenterPolicyRequestUsbSupplyRedirectRule `json:"UsbSupplyRedirectRule,omitempty" xml:"UsbSupplyRedirectRule,omitempty" type:"Repeated"`
	UseTime               *string                                           `json:"UseTime,omitempty" xml:"UseTime,omitempty"`
	// Average video encoding bitrate. Unit: Kbps. Valid values: 1000 to 50000.
	//
	// example:
	//
	// 2000
	VideoEncAvgKbps *int32 `json:"VideoEncAvgKbps,omitempty" xml:"VideoEncAvgKbps,omitempty"`
	// Maximum QP for video encoding, representing the lowest image quality. Valid values: 0 to 51.
	//
	// example:
	//
	// 30
	VideoEncMaxQP *int32 `json:"VideoEncMaxQP,omitempty" xml:"VideoEncMaxQP,omitempty"`
	// Minimum QP for video encoding, representing the highest quality. Valid values: 0 to 51.
	//
	// example:
	//
	// 30
	VideoEncMinQP *int32 `json:"VideoEncMinQP,omitempty" xml:"VideoEncMinQP,omitempty"`
	// Peak video encoding bitrate. Unit: Kbps. Valid values: 1000 to 50000.
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
	// Image display quality policy.
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
	// If you set the `WatermarkType` parameter to `custom`, you must also specify custom text using the `WatermarkCustomText` parameter.
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
	// 3
	WatermarkRowAmount *int32 `json:"WatermarkRowAmount,omitempty" xml:"WatermarkRowAmount,omitempty"`
	// Blind watermark security priority rule.
	//
	// example:
	//
	// on
	WatermarkSecurity *string `json:"WatermarkSecurity,omitempty" xml:"WatermarkSecurity,omitempty"`
	WatermarkShadow   *string `json:"WatermarkShadow,omitempty" xml:"WatermarkShadow,omitempty"`
	// Watermark opacity. A higher value indicates lower transparency. Valid values: 10 to 100.
	//
	// example:
	//
	// 10
	WatermarkTransparencyValue *int32 `json:"WatermarkTransparencyValue,omitempty" xml:"WatermarkTransparencyValue,omitempty"`
	// Watermark type. You can select up to three types, separated by commas.
	//
	// > If you set this parameter to `custom`, you must also specify custom text using the `WatermarkCustomText` parameter.
	//
	// example:
	//
	// EndUserId,HostName,ClientTime
	WatermarkType *string `json:"WatermarkType,omitempty" xml:"WatermarkType,omitempty"`
	// Wuying Keeper switch.
	//
	// example:
	//
	// off
	WuyingKeeper *string `json:"WuyingKeeper,omitempty" xml:"WuyingKeeper,omitempty"`
	// Provide an entry point for Wuying AI Assistant in the floating ball on the cloud desktop when connecting through desktop clients (including Windows and macOS clients).
	//
	// > This feature applies only to desktop clients V7.7 or later.
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

func (s *CreateCenterPolicyRequest) GetAcademicProxy() *string {
	return s.AcademicProxy
}

func (s *CreateCenterPolicyRequest) GetAdminAccess() *string {
	return s.AdminAccess
}

func (s *CreateCenterPolicyRequest) GetAdminKeyboardOnFullScreen() *string {
	return s.AdminKeyboardOnFullScreen
}

func (s *CreateCenterPolicyRequest) GetAdminKeyboardOnWindows() *string {
	return s.AdminKeyboardOnWindows
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

func (s *CreateCenterPolicyRequest) GetBusinessChannel() *string {
	return s.BusinessChannel
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

func (s *CreateCenterPolicyRequest) GetClientCreateSnapshot() *string {
	return s.ClientCreateSnapshot
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

func (s *CreateCenterPolicyRequest) GetCpuOverload() *string {
	return s.CpuOverload
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

func (s *CreateCenterPolicyRequest) GetDescription() *string {
	return s.Description
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

func (s *CreateCenterPolicyRequest) GetDiskOverload() *string {
	return s.DiskOverload
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

func (s *CreateCenterPolicyRequest) GetExternalDrive() *string {
	return s.ExternalDrive
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

func (s *CreateCenterPolicyRequest) GetHoverConfigMsg() *string {
	return s.HoverConfigMsg
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

func (s *CreateCenterPolicyRequest) GetKeyboardControl() *string {
	return s.KeyboardControl
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

func (s *CreateCenterPolicyRequest) GetMemoryOverload() *string {
	return s.MemoryOverload
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

func (s *CreateCenterPolicyRequest) GetModelLibrary() *string {
	return s.ModelLibrary
}

func (s *CreateCenterPolicyRequest) GetMultiScreen() *string {
	return s.MultiScreen
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

func (s *CreateCenterPolicyRequest) GetPortProxy() *string {
	return s.PortProxy
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

func (s *CreateCenterPolicyRequest) GetResolutionDpi() *int32 {
	return s.ResolutionDpi
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

func (s *CreateCenterPolicyRequest) GetWatermarkShadow() *string {
	return s.WatermarkShadow
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

func (s *CreateCenterPolicyRequest) SetAcademicProxy(v string) *CreateCenterPolicyRequest {
	s.AcademicProxy = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetAdminAccess(v string) *CreateCenterPolicyRequest {
	s.AdminAccess = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetAdminKeyboardOnFullScreen(v string) *CreateCenterPolicyRequest {
	s.AdminKeyboardOnFullScreen = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetAdminKeyboardOnWindows(v string) *CreateCenterPolicyRequest {
	s.AdminKeyboardOnWindows = &v
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

func (s *CreateCenterPolicyRequest) SetBusinessChannel(v string) *CreateCenterPolicyRequest {
	s.BusinessChannel = &v
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

func (s *CreateCenterPolicyRequest) SetClientCreateSnapshot(v string) *CreateCenterPolicyRequest {
	s.ClientCreateSnapshot = &v
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

func (s *CreateCenterPolicyRequest) SetCpuOverload(v string) *CreateCenterPolicyRequest {
	s.CpuOverload = &v
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

func (s *CreateCenterPolicyRequest) SetDescription(v string) *CreateCenterPolicyRequest {
	s.Description = &v
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

func (s *CreateCenterPolicyRequest) SetDiskOverload(v string) *CreateCenterPolicyRequest {
	s.DiskOverload = &v
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

func (s *CreateCenterPolicyRequest) SetExternalDrive(v string) *CreateCenterPolicyRequest {
	s.ExternalDrive = &v
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

func (s *CreateCenterPolicyRequest) SetHoverConfigMsg(v string) *CreateCenterPolicyRequest {
	s.HoverConfigMsg = &v
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

func (s *CreateCenterPolicyRequest) SetKeyboardControl(v string) *CreateCenterPolicyRequest {
	s.KeyboardControl = &v
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

func (s *CreateCenterPolicyRequest) SetMemoryOverload(v string) *CreateCenterPolicyRequest {
	s.MemoryOverload = &v
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

func (s *CreateCenterPolicyRequest) SetModelLibrary(v string) *CreateCenterPolicyRequest {
	s.ModelLibrary = &v
	return s
}

func (s *CreateCenterPolicyRequest) SetMultiScreen(v string) *CreateCenterPolicyRequest {
	s.MultiScreen = &v
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

func (s *CreateCenterPolicyRequest) SetPortProxy(v string) *CreateCenterPolicyRequest {
	s.PortProxy = &v
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

func (s *CreateCenterPolicyRequest) SetResolutionDpi(v int32) *CreateCenterPolicyRequest {
	s.ResolutionDpi = &v
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

func (s *CreateCenterPolicyRequest) SetWatermarkShadow(v string) *CreateCenterPolicyRequest {
	s.WatermarkShadow = &v
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

type CreateCenterPolicyRequestAuthorizeAccessPolicyRule struct {
	// Client access IP address range. IPv4 CIDR block.
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
	// Object of the security group control rule. IPv4 CIDR block.
	//
	// example:
	//
	// 10.0.XX.XX/8
	CidrIp *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	// Description of the security group control rule.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Protocol type of the security group control rule.
	//
	// example:
	//
	// TCP
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// Authorization policy of the security group control rule.
	//
	// example:
	//
	// accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// Port range of the security group control rule. The port range depends on the protocol (IpProtocol):
	//
	// - TCP or UDP: Port range is 1 to 65535. Separate the start and end ports with a forward slash (/). Example: 1/200.
	//
	// - ICMP: Port is -1/-1.
	//
	// - GRE: Port is -1/-1.
	//
	// - IpProtocol is all: Port is -1/-1.
	//
	// For common ports used by typical applications, see [Common ports](https://help.aliyun.com/document_detail/40724.html).
	//
	// example:
	//
	// 22/22
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// Priority of the security group control rule. A smaller number indicates a higher priority.<br>
	//
	// Valid values: 1 to 60.<br>
	//
	// Default value: 1.<br><br>
	//
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// Direction of the security group control rule.
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
	// Logon method control. Specify the client type.
	//
	// example:
	//
	// windows
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// Logon method control. Specify whether to allow users to log on to cloud desktops using a specific type of client.
	//
	// > If you do not configure parameters related to `ClientType`, all types of clients are allowed to log on to cloud desktops by default.
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
	ClipboardSize        *int32  `json:"ClipboardSize,omitempty" xml:"ClipboardSize,omitempty"`
	ClipboardSizeUnit    *string `json:"ClipboardSizeUnit,omitempty" xml:"ClipboardSizeUnit,omitempty"`
	ClipboardType        *string `json:"ClipboardType,omitempty" xml:"ClipboardType,omitempty"`
	GrainedType          *string `json:"GrainedType,omitempty" xml:"GrainedType,omitempty"`
	InClipboardSize      *int32  `json:"InClipboardSize,omitempty" xml:"InClipboardSize,omitempty"`
	InClipboardSizeUnit  *string `json:"InClipboardSizeUnit,omitempty" xml:"InClipboardSizeUnit,omitempty"`
	OutClipboardSize     *int32  `json:"OutClipboardSize,omitempty" xml:"OutClipboardSize,omitempty"`
	OutClipboardSizeUnit *string `json:"OutClipboardSizeUnit,omitempty" xml:"OutClipboardSizeUnit,omitempty"`
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

func (s *CreateCenterPolicyRequestClipboardGraineds) GetClipboardSizeUnit() *string {
	return s.ClipboardSizeUnit
}

func (s *CreateCenterPolicyRequestClipboardGraineds) GetClipboardType() *string {
	return s.ClipboardType
}

func (s *CreateCenterPolicyRequestClipboardGraineds) GetGrainedType() *string {
	return s.GrainedType
}

func (s *CreateCenterPolicyRequestClipboardGraineds) GetInClipboardSize() *int32 {
	return s.InClipboardSize
}

func (s *CreateCenterPolicyRequestClipboardGraineds) GetInClipboardSizeUnit() *string {
	return s.InClipboardSizeUnit
}

func (s *CreateCenterPolicyRequestClipboardGraineds) GetOutClipboardSize() *int32 {
	return s.OutClipboardSize
}

func (s *CreateCenterPolicyRequestClipboardGraineds) GetOutClipboardSizeUnit() *string {
	return s.OutClipboardSizeUnit
}

func (s *CreateCenterPolicyRequestClipboardGraineds) SetClipboardSize(v int32) *CreateCenterPolicyRequestClipboardGraineds {
	s.ClipboardSize = &v
	return s
}

func (s *CreateCenterPolicyRequestClipboardGraineds) SetClipboardSizeUnit(v string) *CreateCenterPolicyRequestClipboardGraineds {
	s.ClipboardSizeUnit = &v
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

func (s *CreateCenterPolicyRequestClipboardGraineds) SetInClipboardSize(v int32) *CreateCenterPolicyRequestClipboardGraineds {
	s.InClipboardSize = &v
	return s
}

func (s *CreateCenterPolicyRequestClipboardGraineds) SetInClipboardSizeUnit(v string) *CreateCenterPolicyRequestClipboardGraineds {
	s.InClipboardSizeUnit = &v
	return s
}

func (s *CreateCenterPolicyRequestClipboardGraineds) SetOutClipboardSize(v int32) *CreateCenterPolicyRequestClipboardGraineds {
	s.OutClipboardSize = &v
	return s
}

func (s *CreateCenterPolicyRequestClipboardGraineds) SetOutClipboardSizeUnit(v string) *CreateCenterPolicyRequestClipboardGraineds {
	s.OutClipboardSizeUnit = &v
	return s
}

func (s *CreateCenterPolicyRequestClipboardGraineds) Validate() error {
	return dara.Validate(s)
}

type CreateCenterPolicyRequestDeviceRedirects struct {
	// Device type
	//
	// example:
	//
	// camera
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// Redirection type.
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
	// Policy description.
	//
	// example:
	//
	// 测试规则
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
	// allow
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// Rule type.
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
	// Rule description.
	//
	// example:
	//
	// 测试规则
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
	// 2
	UsbRuleType *string `json:"UsbRuleType,omitempty" xml:"UsbRuleType,omitempty"`
	// Vendor ID. For more information, see [Valid USB Vendor IDs (VIDs)](https://www.usb.org/sites/default/files/vendor_ids032322.pdf_1.pdf).
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
