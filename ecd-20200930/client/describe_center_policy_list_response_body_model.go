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
	// The details of the cloud computer policy.
	DescribePolicyGroups []*DescribeCenterPolicyListResponseBodyDescribePolicyGroups `json:"DescribePolicyGroups,omitempty" xml:"DescribePolicyGroups,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of query results.
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
	if s.DescribePolicyGroups != nil {
		for _, item := range s.DescribePolicyGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCenterPolicyListResponseBodyDescribePolicyGroups struct {
	// Specifies whether the academic proxy feature is enabled. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	AcademicProxy *string `json:"AcademicProxy,omitempty" xml:"AcademicProxy,omitempty"`
	// Specifies whether the user has administrator permissions after connecting to the cloud computer.
	//
	// > This feature is in limited preview and is not publicly available.
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
	// Specifies whether to enable administrator keyboard control in Windows. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	AdminKeyboardOnWindows *string `json:"AdminKeyboardOnWindows,omitempty" xml:"AdminKeyboardOnWindows,omitempty"`
	// Specifies whether to enable administrator monitoring. Valid values:
	//
	// - 1: Enabled.
	//
	// - 0: Disabled.
	//
	// example:
	//
	// 0
	AdminMonitor *int32 `json:"AdminMonitor,omitempty" xml:"AdminMonitor,omitempty"`
	// Specifies whether to enable the screen capture prevention feature.
	//
	// example:
	//
	// off
	AppContentProtection *string `json:"AppContentProtection,omitempty" xml:"AppContentProtection,omitempty"`
	// The list of custom OSS buckets for mobile screenshots in screen capture prevention.
	AppScreenshotBuckets []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAppScreenshotBuckets `json:"AppScreenshotBuckets,omitempty" xml:"AppScreenshotBuckets,omitempty" type:"Repeated"`
	// The IP address whitelist for client access.
	AuthorizeAccessPolicyRules []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules `json:"AuthorizeAccessPolicyRules,omitempty" xml:"AuthorizeAccessPolicyRules,omitempty" type:"Repeated"`
	// The list of security group control rules.
	AuthorizeSecurityPolicyRules []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules `json:"AuthorizeSecurityPolicyRules,omitempty" xml:"AuthorizeSecurityPolicyRules,omitempty" type:"Repeated"`
	// Specifies whether to enable automatic reconnection for the client.
	//
	// example:
	//
	// off
	AutoReconnect *string `json:"AutoReconnect,omitempty" xml:"AutoReconnect,omitempty"`
	// Specifies whether to automatically save mobile screenshots. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	AutoSaveAppScreenshot *string `json:"AutoSaveAppScreenshot,omitempty" xml:"AutoSaveAppScreenshot,omitempty"`
	// Specifies whether to enable local camera redirection.
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
	// Specifies whether to enable the feature that allows the client to create custom snapshots. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	ClientCreateSnapshot *string `json:"ClientCreateSnapshot,omitempty" xml:"ClientCreateSnapshot,omitempty"`
	// Specifies whether to enable the hibernate option in the client menu. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	ClientHibernate *string `json:"ClientHibernate,omitempty" xml:"ClientHibernate,omitempty"`
	// Specifies whether to enable the restart option in the client menu. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	ClientRestart *string `json:"ClientRestart,omitempty" xml:"ClientRestart,omitempty"`
	// Specifies whether to enable the shutdown option in the client menu. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	ClientShutdown *string `json:"ClientShutdown,omitempty" xml:"ClientShutdown,omitempty"`
	// The logon method control rules.
	ClientTypes []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsClientTypes `json:"ClientTypes,omitempty" xml:"ClientTypes,omitempty" type:"Repeated"`
	// The clipboard permissions.
	//
	// example:
	//
	// off
	Clipboard *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	// Specifies whether color enhancement is enabled for design and 3D Common scenarios.
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
	// The CPU downclocking duration. Valid values: 30 to 120. Unit: seconds.
	//
	// example:
	//
	// 30
	CpuDownGradeDuration *int32 `json:"CpuDownGradeDuration,omitempty" xml:"CpuDownGradeDuration,omitempty"`
	// Specifies whether to enable CPU load protection. Valid values:
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
	// on
	CpuProtectedMode *string `json:"CpuProtectedMode,omitempty" xml:"CpuProtectedMode,omitempty"`
	// The overall CPU utilization percentage. Valid values: 70 to 90.
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
	// The single-core CPU utilization percentage. Valid values: 70 to 100.
	//
	// example:
	//
	// 70
	CpuSingleRateLimit *int32 `json:"CpuSingleRateLimit,omitempty" xml:"CpuSingleRateLimit,omitempty"`
	// The number of cloud computers associated with this policy.
	//
	// example:
	//
	// 1
	DesktopCount *int32 `json:"DesktopCount,omitempty" xml:"DesktopCount,omitempty"`
	// The number of cloud computer pools associated with this policy.
	//
	// example:
	//
	// 1
	DesktopGroupCount *int32 `json:"DesktopGroupCount,omitempty" xml:"DesktopGroupCount,omitempty"`
	// The peripheral connection prompt control.
	//
	// example:
	//
	// off
	DeviceConnectHint *string `json:"DeviceConnectHint,omitempty" xml:"DeviceConnectHint,omitempty"`
	// The device redirection rules.
	DeviceRedirects []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRedirects `json:"DeviceRedirects,omitempty" xml:"DeviceRedirects,omitempty" type:"Repeated"`
	// The custom peripheral rules.
	DeviceRules []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules `json:"DeviceRules,omitempty" xml:"DeviceRules,omitempty" type:"Repeated"`
	// Specifies whether to retain the session after disconnection.
	//
	// > This applies only to cloud application policies.
	//
	// example:
	//
	// persistent
	DisconnectKeepSession *string `json:"DisconnectKeepSession,omitempty" xml:"DisconnectKeepSession,omitempty"`
	// The session retention duration after disconnection. Unit: seconds.
	//
	// > This applies only to cloud application policies.
	//
	// example:
	//
	// 120
	DisconnectKeepSessionTime *int32 `json:"DisconnectKeepSessionTime,omitempty" xml:"DisconnectKeepSessionTime,omitempty"`
	// Specifies whether to enable disk load protection. Valid values:
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
	// adminCustom
	DisplayMode *string `json:"DisplayMode,omitempty" xml:"DisplayMode,omitempty"`
	// The effective field for domain name resolution.
	//
	// example:
	//
	// example.com
	DomainRegisterValue *string `json:"DomainRegisterValue,omitempty" xml:"DomainRegisterValue,omitempty"`
	// The details of the domain name resolution policy.
	DomainResolveRule []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDomainResolveRule `json:"DomainResolveRule,omitempty" xml:"DomainResolveRule,omitempty" type:"Repeated"`
	// Specifies whether to enable the domain name resolution policy.
	//
	// example:
	//
	// on
	DomainResolveRuleType *string `json:"DomainResolveRuleType,omitempty" xml:"DomainResolveRuleType,omitempty"`
	// Specifies whether to allow users to request administrator assistance.
	//
	// example:
	//
	// off
	EndUserApplyAdminCoordinate *string `json:"EndUserApplyAdminCoordinate,omitempty" xml:"EndUserApplyAdminCoordinate,omitempty"`
	// The number of associated end users.
	//
	// example:
	//
	// 3
	EndUserCount *string `json:"EndUserCount,omitempty" xml:"EndUserCount,omitempty"`
	// Specifies whether users in the same office network share cloud desktops.
	//
	// example:
	//
	// off
	EndUserGroupCoordinate *string `json:"EndUserGroupCoordinate,omitempty" xml:"EndUserGroupCoordinate,omitempty"`
	// Specifies whether to allow the use of external storage devices. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	ExternalDrive *string `json:"ExternalDrive,omitempty" xml:"ExternalDrive,omitempty"`
	// The file migration setting.
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
	// The maximum size of a file that can be transferred to the cloud desktop in a single operation. This parameter must be used together with the transfer-in unit parameter.
	//
	// example:
	//
	// 100
	FileTransferInSize *string `json:"FileTransferInSize,omitempty" xml:"FileTransferInSize,omitempty"`
	// The unit for the maximum size of a file that can be transferred to the cloud desktop in a single operation.
	//
	// example:
	//
	// MB
	FileTransferInUnit *string `json:"FileTransferInUnit,omitempty" xml:"FileTransferInUnit,omitempty"`
	// The maximum size of a file that can be transferred from the cloud desktop in a single operation. This parameter must be used together with the transfer-out unit parameter.
	//
	// example:
	//
	// 100
	FileTransferOutSize *string `json:"FileTransferOutSize,omitempty" xml:"FileTransferOutSize,omitempty"`
	// The unit for the maximum size of a file that can be transferred from the cloud desktop in a single operation.
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
	// Specifies whether to enable file synchronization. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	FolderSync *string `json:"FolderSync,omitempty" xml:"FolderSync,omitempty"`
	// Specifies whether to enable the image quality policy for cloud desktops with graphics-optimized specifications. Enable this policy when high performance and user experience are required, such as in professional design scenarios.
	//
	// example:
	//
	// off
	GpuAcceleration *string `json:"GpuAcceleration,omitempty" xml:"GpuAcceleration,omitempty"`
	// Specifies whether to enable message prompts for the floating ball configuration. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	HoverConfigMsg *string `json:"HoverConfigMsg,omitempty" xml:"HoverConfigMsg,omitempty"`
	// Specifies whether to enable the hibernate button on the floating ball. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	HoverHibernate *string `json:"HoverHibernate,omitempty" xml:"HoverHibernate,omitempty"`
	// Specifies whether to enable the restart button on the floating ball. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	HoverRestart *string `json:"HoverRestart,omitempty" xml:"HoverRestart,omitempty"`
	// Specifies whether to enable the shutdown button on the floating ball. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	HoverShutdown *string `json:"HoverShutdown,omitempty" xml:"HoverShutdown,omitempty"`
	// The access policy for the web client.
	//
	// example:
	//
	// off
	Html5Access *string `json:"Html5Access,omitempty" xml:"Html5Access,omitempty"`
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
	// tcp
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
	// Specifies whether to enable keyboard control on the floating ball. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	KeyboardControl *string `json:"KeyboardControl,omitempty" xml:"KeyboardControl,omitempty"`
	// Specifies whether to enable the KMS server. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	KmsServer *string `json:"KmsServer,omitempty" xml:"KmsServer,omitempty"`
	// The KMS server configuration in the address:port format.
	//
	// example:
	//
	// kms.example.com:1688
	KmsServerConfig *string `json:"KmsServerConfig,omitempty" xml:"KmsServerConfig,omitempty"`
	// The mapping permissions for local disks.
	//
	// example:
	//
	// readwrite
	LocalDrive *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	// The maximum connection retry time for a cloud desktop that is disconnected due to objective reasons. Valid values: 30 to 7200. Unit: seconds.
	//
	// example:
	//
	// 120
	MaxReconnectTime *int32 `json:"MaxReconnectTime,omitempty" xml:"MaxReconnectTime,omitempty"`
	// Specifies whether to enable the MCP service. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	McpService *string `json:"McpService,omitempty" xml:"McpService,omitempty"`
	// The memory downgrade duration for a single process. Valid values: 30 to 120. Unit: seconds.
	//
	// example:
	//
	// 30
	MemoryDownGradeDuration *int32 `json:"MemoryDownGradeDuration,omitempty" xml:"MemoryDownGradeDuration,omitempty"`
	// Specifies whether to enable memory load protection. Valid values:
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
	// Specifies whether to enable the memory protection mode.
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
	// 30
	MemorySampleDuration *int32 `json:"MemorySampleDuration,omitempty" xml:"MemorySampleDuration,omitempty"`
	// The memory usage percentage for a single process. Valid values: 30 to 60.
	//
	// example:
	//
	// 30
	MemorySingleRateLimit *int32 `json:"MemorySingleRateLimit,omitempty" xml:"MemorySingleRateLimit,omitempty"`
	// Specifies whether to provide a restart button on the floating ball when you connect to a cloud desktop from a mobile client (Android client<props="china"> and iOS client).
	//
	// > This parameter applies only to mobile clients of V7.4 or later.
	//
	// example:
	//
	// off
	MobileRestart *string `json:"MobileRestart,omitempty" xml:"MobileRestart,omitempty"`
	// Specifies whether to enable the Windows security control on the mobile client.
	//
	// example:
	//
	// off
	MobileSafeMenu *string `json:"MobileSafeMenu,omitempty" xml:"MobileSafeMenu,omitempty"`
	// Specifies whether to provide a shutdown button on the floating ball when you connect to a cloud desktop from a mobile client (Android client<props="china"> and iOS client).
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
	// Specifies whether to enable the Wuying Assistant on the mobile client.
	//
	// example:
	//
	// off
	MobileWyAssistant *string `json:"MobileWyAssistant,omitempty" xml:"MobileWyAssistant,omitempty"`
	// Specifies whether the model library feature is enabled. Valid values:
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
	// The network redirection setting.
	//
	// > This feature is in closed beta and is not publicly available.
	//
	// example:
	//
	// off
	NetRedirect *string `json:"NetRedirect,omitempty" xml:"NetRedirect,omitempty"`
	// The details of the network redirection policy.
	//
	// > This feature is in closed beta and is not publicly available.
	//
	// >
	NetRedirectRule []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsNetRedirectRule `json:"NetRedirectRule,omitempty" xml:"NetRedirectRule,omitempty" type:"Repeated"`
	// Specifies whether to enable the network printer feature. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	NetworkPrinter *string `json:"NetworkPrinter,omitempty" xml:"NetworkPrinter,omitempty"`
	// Specifies whether to disconnect the session when no operations are performed.
	//
	// > This parameter applies only to cloud application policies.
	//
	// example:
	//
	// off
	NoOperationDisconnect *string `json:"NoOperationDisconnect,omitempty" xml:"NoOperationDisconnect,omitempty"`
	// The idle timeout period before disconnection. Unit: seconds.
	//
	// > This parameter applies only to cloud application policies.
	//
	// example:
	//
	// 120
	NoOperationDisconnectTime *int32 `json:"NoOperationDisconnectTime,omitempty" xml:"NoOperationDisconnectTime,omitempty"`
	// The number of associated organizations.
	//
	// example:
	//
	// 2
	OrganizationCount *string `json:"OrganizationCount,omitempty" xml:"OrganizationCount,omitempty"`
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
	// The status of the cloud desktop policy.
	//
	// example:
	//
	// AVAILABLE
	PolicyStatus *string `json:"PolicyStatus,omitempty" xml:"PolicyStatus,omitempty"`
	// Specifies whether the port proxy feature is enabled. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	PortProxy *string `json:"PortProxy,omitempty" xml:"PortProxy,omitempty"`
	// The printer pop-up prompt setting. Valid values:
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
	// The printer redirection policy.
	//
	// example:
	//
	// off
	PrinterRedirection *string `json:"PrinterRedirection,omitempty" xml:"PrinterRedirection,omitempty"`
	// Specifies whether image quality enhancement is enabled for Common scenarios such as design and 3D.
	//
	// example:
	//
	// off
	QualityEnhancement *string `json:"QualityEnhancement,omitempty" xml:"QualityEnhancement,omitempty"`
	// Specifies whether to enable custom screen recording.
	//
	// example:
	//
	// off
	RecordContent *string `json:"RecordContent,omitempty" xml:"RecordContent,omitempty"`
	// The expiration time for custom screen recording files. Default value: 30 days.
	//
	// example:
	//
	// 30
	RecordContentExpires *int64 `json:"RecordContentExpires,omitempty" xml:"RecordContentExpires,omitempty"`
	// The screen recording duration after an event is detected in screen recording audit. Unit: minutes. Valid values: 10 to 60.
	//
	// example:
	//
	// 10
	RecordEventDuration *int32 `json:"RecordEventDuration,omitempty" xml:"RecordEventDuration,omitempty"`
	// The absolute paths for file monitoring in screen recording audit.
	RecordEventFilePaths []*string `json:"RecordEventFilePaths,omitempty" xml:"RecordEventFilePaths,omitempty" type:"Repeated"`
	// The absolute paths for registry monitoring in screen recording audit.
	RecordEventRegisters []*string `json:"RecordEventRegisters,omitempty" xml:"RecordEventRegisters,omitempty" type:"Repeated"`
	// Specifies whether to enable screen recording.
	//
	// example:
	//
	// off
	Recording *string `json:"Recording,omitempty" xml:"Recording,omitempty"`
	// The option for recording the audio of the cloud desktop.
	//
	// example:
	//
	// on
	RecordingAudio *string `json:"RecordingAudio,omitempty" xml:"RecordingAudio,omitempty"`
	// The duration of each screen recording file, in minutes. Recorded files are automatically split and uploaded to the storage space based on the specified duration. If a file reaches 300 MB, it is overwritten on a rolling basis.
	//
	// example:
	//
	// 10
	RecordingDuration *int32 `json:"RecordingDuration,omitempty" xml:"RecordingDuration,omitempty"`
	// The end time of screen recording. Format: HH:MM:SS. This parameter is valid only when the Recording parameter is set to period.
	//
	// example:
	//
	// 08:59:00
	RecordingEndTime *string `json:"RecordingEndTime,omitempty" xml:"RecordingEndTime,omitempty"`
	// The retention period of screen recording files. Valid values: 1 to 180 days.
	//
	// example:
	//
	// 15
	RecordingExpires *int64 `json:"RecordingExpires,omitempty" xml:"RecordingExpires,omitempty"`
	// The frame rate of screen recording. Unit: FPS.
	//
	// example:
	//
	// 5
	RecordingFps *int64 `json:"RecordingFps,omitempty" xml:"RecordingFps,omitempty"`
	// The start time of screen recording. Format: HH:MM:SS. This parameter is valid only when the Recording parameter is set to period.
	//
	// example:
	//
	// 08:00:00
	RecordingStartTime *string `json:"RecordingStartTime,omitempty" xml:"RecordingStartTime,omitempty"`
	// Specifies whether to notify the terminal user that screen recording is enabled.
	//
	// example:
	//
	// off
	RecordingUserNotify *string `json:"RecordingUserNotify,omitempty" xml:"RecordingUserNotify,omitempty"`
	// The notification message sent to the terminal user when screen recording is enabled.
	//
	// example:
	//
	// Screen recording is enabled
	RecordingUserNotifyMessage *string `json:"RecordingUserNotifyMessage,omitempty" xml:"RecordingUserNotifyMessage,omitempty"`
	// The permissions for keyboard and mouse control during remote assistance.
	//
	// example:
	//
	// fullControl
	RemoteCoordinate *string `json:"RemoteCoordinate,omitempty" xml:"RemoteCoordinate,omitempty"`
	// The permission to reset the cloud desktop.
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
	// The height of the resolution. Unit: pixels.
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
	// The width of the resolution. Unit: pixels.
	//
	// example:
	//
	// 1920
	ResolutionWidth *int32 `json:"ResolutionWidth,omitempty" xml:"ResolutionWidth,omitempty"`
	// The number of resource groups associated with the policy.
	//
	// example:
	//
	// 1
	ResourceGroupCount *int32 `json:"ResourceGroupCount,omitempty" xml:"ResourceGroupCount,omitempty"`
	// Specifies whether to enable the shortcut key for the security center. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
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
	// The effective CIDR block.
	ScopeValue []*string `json:"ScopeValue,omitempty" xml:"ScopeValue,omitempty" type:"Repeated"`
	// The screen display mode.
	//
	// example:
	//
	// auto
	ScreenDisplayMode *string `json:"ScreenDisplayMode,omitempty" xml:"ScreenDisplayMode,omitempty"`
	// Specifies whether to enable smoothness enhancement for daily office scenarios.
	//
	// example:
	//
	// off
	SmoothEnhancement *string `json:"SmoothEnhancement,omitempty" xml:"SmoothEnhancement,omitempty"`
	// Specifies whether to provide an entry for the status monitoring feature in the floating ball of the cloud desktop.
	//
	// example:
	//
	// on
	StatusMonitor *string `json:"StatusMonitor,omitempty" xml:"StatusMonitor,omitempty"`
	// The applicable scenario for the streaming mode.
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
	// Specifies whether to enable the triple-screen feature. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	ThreeScreen *string `json:"ThreeScreen,omitempty" xml:"ThreeScreen,omitempty"`
	// The USB redirection setting.
	//
	// example:
	//
	// off
	UsbRedirect *string `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty"`
	// The USB redirection rules.
	UsbSupplyRedirectRule []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule `json:"UsbSupplyRedirectRule,omitempty" xml:"UsbSupplyRedirectRule,omitempty" type:"Repeated"`
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
	// The average video encoding bitrate. Unit: Kbps. Valid values: 1000 to 50000.
	//
	// example:
	//
	// 1000
	VideoEncAvgKbps *int32 `json:"VideoEncAvgKbps,omitempty" xml:"VideoEncAvgKbps,omitempty"`
	// The maximum quantization parameter (QP) for video encoding, which indicates the lowest video quality. Valid values: 0 to 51.
	//
	// example:
	//
	// 20
	VideoEncMaxQP *int32 `json:"VideoEncMaxQP,omitempty" xml:"VideoEncMaxQP,omitempty"`
	// The minimum quantization parameter (QP) for video encoding, which indicates the highest video quality. Valid values: 0 to 51.
	//
	// example:
	//
	// 20
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
	// The multimedia redirection setting.
	//
	// example:
	//
	// off
	VideoRedirect *string `json:"VideoRedirect,omitempty" xml:"VideoRedirect,omitempty"`
	// The image display quality policy.
	//
	// example:
	//
	// medium
	VisualQuality *string `json:"VisualQuality,omitempty" xml:"VisualQuality,omitempty"`
	// The watermark setting.
	//
	// example:
	//
	// on
	Watermark *string `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
	// The blind watermark anti-camera capture feature.
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
	// If the WatermarkType parameter is set to custom, you must specify the custom text by using the WatermarkCustomText parameter.
	//
	// example:
	//
	// custom-watermark
	WatermarkCustomText *string `json:"WatermarkCustomText,omitempty" xml:"WatermarkCustomText,omitempty"`
	// The tilt angle of the watermark. Valid values: -10 to -30.
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
	// The blind watermark enhancement feature.
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
	// The security priority rule for blind watermarks.
	//
	// example:
	//
	// on
	WatermarkSecurity *string `json:"WatermarkSecurity,omitempty" xml:"WatermarkSecurity,omitempty"`
	// Specifies whether to enable the shadow effect for the watermark. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	WatermarkShadow *string `json:"WatermarkShadow,omitempty" xml:"WatermarkShadow,omitempty"`
	// The transparency of the watermark. A larger value indicates lower transparency. Valid values: 10 to 100.
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
	// Specifies whether to enable WUYING Keeper.
	//
	// example:
	//
	// off
	WuyingKeeper *string `json:"WuyingKeeper,omitempty" xml:"WuyingKeeper,omitempty"`
	// Specifies whether to provide an entry for the WUYING AI Assistant feature in the floating ball of the cloud desktop.
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

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetAcademicProxy() *string {
	return s.AcademicProxy
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetAdminAccess() *string {
	return s.AdminAccess
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetAdminKeyboardOnFullScreen() *string {
	return s.AdminKeyboardOnFullScreen
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetAdminKeyboardOnWindows() *string {
	return s.AdminKeyboardOnWindows
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetAdminMonitor() *int32 {
	return s.AdminMonitor
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetAppContentProtection() *string {
	return s.AppContentProtection
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetAppScreenshotBuckets() []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAppScreenshotBuckets {
	return s.AppScreenshotBuckets
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetAuthorizeAccessPolicyRules() []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules {
	return s.AuthorizeAccessPolicyRules
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetAuthorizeSecurityPolicyRules() []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	return s.AuthorizeSecurityPolicyRules
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetAutoReconnect() *string {
	return s.AutoReconnect
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetAutoSaveAppScreenshot() *string {
	return s.AutoSaveAppScreenshot
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetCameraRedirect() *string {
	return s.CameraRedirect
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetClientControlMenu() *string {
	return s.ClientControlMenu
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetClientCreateSnapshot() *string {
	return s.ClientCreateSnapshot
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetClientHibernate() *string {
	return s.ClientHibernate
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetClientRestart() *string {
	return s.ClientRestart
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetClientShutdown() *string {
	return s.ClientShutdown
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

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetCpuOverload() *string {
	return s.CpuOverload
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

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetDeviceConnectHint() *string {
	return s.DeviceConnectHint
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

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetDiskOverload() *string {
	return s.DiskOverload
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

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetEndUserCount() *string {
	return s.EndUserCount
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetEndUserGroupCoordinate() *string {
	return s.EndUserGroupCoordinate
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetExternalDrive() *string {
	return s.ExternalDrive
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetFileMigrate() *string {
	return s.FileMigrate
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetFileTransferAddress() *string {
	return s.FileTransferAddress
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetFileTransferInSize() *string {
	return s.FileTransferInSize
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetFileTransferInUnit() *string {
	return s.FileTransferInUnit
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetFileTransferOutSize() *string {
	return s.FileTransferOutSize
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetFileTransferOutUnit() *string {
	return s.FileTransferOutUnit
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetFileTransferSizeLimit() *string {
	return s.FileTransferSizeLimit
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetFileTransferSpeed() *string {
	return s.FileTransferSpeed
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetFileTransferSpeedLocation() *string {
	return s.FileTransferSpeedLocation
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetFolderSync() *string {
	return s.FolderSync
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetGpuAcceleration() *string {
	return s.GpuAcceleration
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetHoverConfigMsg() *string {
	return s.HoverConfigMsg
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetHoverHibernate() *string {
	return s.HoverHibernate
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetHoverRestart() *string {
	return s.HoverRestart
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetHoverShutdown() *string {
	return s.HoverShutdown
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

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetKeyboardControl() *string {
	return s.KeyboardControl
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetKmsServer() *string {
	return s.KmsServer
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetKmsServerConfig() *string {
	return s.KmsServerConfig
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetLocalDrive() *string {
	return s.LocalDrive
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetMaxReconnectTime() *int32 {
	return s.MaxReconnectTime
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetMcpService() *string {
	return s.McpService
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetMemoryDownGradeDuration() *int32 {
	return s.MemoryDownGradeDuration
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetMemoryOverload() *string {
	return s.MemoryOverload
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

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetMobileSafeMenu() *string {
	return s.MobileSafeMenu
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetMobileShutdown() *string {
	return s.MobileShutdown
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetMobileWuyingKeeper() *string {
	return s.MobileWuyingKeeper
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetMobileWyAssistant() *string {
	return s.MobileWyAssistant
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetModelLibrary() *string {
	return s.ModelLibrary
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetMultiScreen() *string {
	return s.MultiScreen
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

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetNetworkPrinter() *string {
	return s.NetworkPrinter
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetNoOperationDisconnect() *string {
	return s.NoOperationDisconnect
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetNoOperationDisconnectTime() *int32 {
	return s.NoOperationDisconnectTime
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetOrganizationCount() *string {
	return s.OrganizationCount
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

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetPortProxy() *string {
	return s.PortProxy
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetPrinterAlert() *string {
	return s.PrinterAlert
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetPrinterAlertContent() *string {
	return s.PrinterAlertContent
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetPrinterAlertTitle() *string {
	return s.PrinterAlertTitle
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

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetResetDesktop() *string {
	return s.ResetDesktop
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetResolutionDpi() *int32 {
	return s.ResolutionDpi
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

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetThreeScreen() *string {
	return s.ThreeScreen
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

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetWatermarkShadow() *string {
	return s.WatermarkShadow
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetWatermarkTransparencyValue() *int32 {
	return s.WatermarkTransparencyValue
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetWatermarkType() *string {
	return s.WatermarkType
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetWuyingKeeper() *string {
	return s.WuyingKeeper
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetWyAssistant() *string {
	return s.WyAssistant
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetAcademicProxy(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.AcademicProxy = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetAdminAccess(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.AdminAccess = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetAdminKeyboardOnFullScreen(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.AdminKeyboardOnFullScreen = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetAdminKeyboardOnWindows(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.AdminKeyboardOnWindows = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetAdminMonitor(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.AdminMonitor = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetAppContentProtection(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.AppContentProtection = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetAppScreenshotBuckets(v []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAppScreenshotBuckets) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.AppScreenshotBuckets = v
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

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetAutoReconnect(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.AutoReconnect = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetAutoSaveAppScreenshot(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.AutoSaveAppScreenshot = &v
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

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetClientCreateSnapshot(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.ClientCreateSnapshot = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetClientHibernate(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.ClientHibernate = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetClientRestart(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.ClientRestart = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetClientShutdown(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.ClientShutdown = &v
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

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetCpuOverload(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.CpuOverload = &v
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

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetDeviceConnectHint(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.DeviceConnectHint = &v
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

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetDiskOverload(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.DiskOverload = &v
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

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetEndUserCount(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.EndUserCount = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetEndUserGroupCoordinate(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.EndUserGroupCoordinate = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetExternalDrive(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.ExternalDrive = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetFileMigrate(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.FileMigrate = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetFileTransferAddress(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.FileTransferAddress = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetFileTransferInSize(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.FileTransferInSize = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetFileTransferInUnit(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.FileTransferInUnit = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetFileTransferOutSize(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.FileTransferOutSize = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetFileTransferOutUnit(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.FileTransferOutUnit = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetFileTransferSizeLimit(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.FileTransferSizeLimit = &v
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

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetFolderSync(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.FolderSync = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetGpuAcceleration(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.GpuAcceleration = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetHoverConfigMsg(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.HoverConfigMsg = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetHoverHibernate(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.HoverHibernate = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetHoverRestart(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.HoverRestart = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetHoverShutdown(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.HoverShutdown = &v
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

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetKeyboardControl(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.KeyboardControl = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetKmsServer(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.KmsServer = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetKmsServerConfig(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.KmsServerConfig = &v
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

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetMcpService(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.McpService = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetMemoryDownGradeDuration(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.MemoryDownGradeDuration = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetMemoryOverload(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.MemoryOverload = &v
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

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetMobileSafeMenu(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.MobileSafeMenu = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetMobileShutdown(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.MobileShutdown = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetMobileWuyingKeeper(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.MobileWuyingKeeper = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetMobileWyAssistant(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.MobileWyAssistant = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetModelLibrary(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.ModelLibrary = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetMultiScreen(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.MultiScreen = &v
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

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetNetworkPrinter(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.NetworkPrinter = &v
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

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetOrganizationCount(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.OrganizationCount = &v
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

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetPortProxy(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.PortProxy = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetPrinterAlert(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.PrinterAlert = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetPrinterAlertContent(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.PrinterAlertContent = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetPrinterAlertTitle(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.PrinterAlertTitle = &v
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

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetResetDesktop(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.ResetDesktop = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetResolutionDpi(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.ResolutionDpi = &v
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

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetThreeScreen(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.ThreeScreen = &v
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

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetWatermarkShadow(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.WatermarkShadow = &v
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

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetWuyingKeeper(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.WuyingKeeper = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetWyAssistant(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.WyAssistant = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) Validate() error {
	if s.AppScreenshotBuckets != nil {
		for _, item := range s.AppScreenshotBuckets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AuthorizeAccessPolicyRules != nil {
		for _, item := range s.AuthorizeAccessPolicyRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AuthorizeSecurityPolicyRules != nil {
		for _, item := range s.AuthorizeSecurityPolicyRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ClientTypes != nil {
		for _, item := range s.ClientTypes {
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

type DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAppScreenshotBuckets struct {
	// The name of the bucket.
	//
	// example:
	//
	// my-bucket-01
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The region where the user\\"s OSS bucket is located. It must be the same as the actual region of the bucket.
	//
	// example:
	//
	// cn-beijing
	BucketRegionId *string `json:"BucketRegionId,omitempty" xml:"BucketRegionId,omitempty"`
}

func (s DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAppScreenshotBuckets) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAppScreenshotBuckets) GoString() string {
	return s.String()
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAppScreenshotBuckets) GetBucketName() *string {
	return s.BucketName
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAppScreenshotBuckets) GetBucketRegionId() *string {
	return s.BucketRegionId
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAppScreenshotBuckets) SetBucketName(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAppScreenshotBuckets {
	s.BucketName = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAppScreenshotBuckets) SetBucketRegionId(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAppScreenshotBuckets {
	s.BucketRegionId = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAppScreenshotBuckets) Validate() error {
	return dara.Validate(s)
}

type DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules struct {
	// The CIDR block for client access. The value must be an IPv4 CIDR block.
	//
	// example:
	//
	// 47.100.XX.XX/16
	CidrIp *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	// The description of the CIDR block for client access.
	//
	// example:
	//
	// Corporate office CIDR block
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
	// The target of the security group control rule. The value must be an IPv4 CIDR block.
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
	// tcp
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The authorization policy of the security group control rule.
	//
	// example:
	//
	// accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The port range of the security group control rule.
	//
	// example:
	//
	// 22/22
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// The priority of the security group control rule. A smaller value indicates a higher priority.
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
	// Specifies whether to allow a specific type of client to log on to the cloud desktop.
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
	// The vendor ID. For more information, see [Valid USB Vendor IDs (VIDs)](https://www.usb.org/sites/default/files/vendor_ids032322.pdf_1.pdf).
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
	// The platform types to which the device rules apply.
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
	//
	// example:
	//
	// Test rule
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
	// The description of the rule.
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
	UsbRedirectType *int64 `json:"UsbRedirectType,omitempty" xml:"UsbRedirectType,omitempty"`
	// The type of the USB redirection rule.
	//
	// example:
	//
	// 1
	UsbRuleType *int64 `json:"UsbRuleType,omitempty" xml:"UsbRuleType,omitempty"`
	// The vendor ID. For more information, see [Valid USB Vendor IDs (VIDs)](https://www.usb.org/sites/default/files/vendor_ids032322.pdf_1.pdf).
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
