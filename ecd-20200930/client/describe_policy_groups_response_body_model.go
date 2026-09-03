// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolicyGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribePolicyGroupsResponseBody
	GetCount() *int32
	SetDescribePolicyGroups(v []*DescribePolicyGroupsResponseBodyDescribePolicyGroups) *DescribePolicyGroupsResponseBody
	GetDescribePolicyGroups() []*DescribePolicyGroupsResponseBodyDescribePolicyGroups
	SetNextToken(v string) *DescribePolicyGroupsResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *DescribePolicyGroupsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePolicyGroupsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribePolicyGroupsResponseBody
	GetRequestId() *string
}

type DescribePolicyGroupsResponseBody struct {
	// The total number of entries.
	//
	// example:
	//
	// 40
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The details of the cloud computer policies.
	DescribePolicyGroups []*DescribePolicyGroupsResponseBodyDescribePolicyGroups `json:"DescribePolicyGroups,omitempty" xml:"DescribePolicyGroups,omitempty" type:"Repeated"`
	// The token for the next query. If NextToken is empty, no more results exist.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number of the current page for a paged query.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of entries per page for a paged query.
	//
	// Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePolicyGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolicyGroupsResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribePolicyGroupsResponseBody) GetDescribePolicyGroups() []*DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	return s.DescribePolicyGroups
}

func (s *DescribePolicyGroupsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribePolicyGroupsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePolicyGroupsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePolicyGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePolicyGroupsResponseBody) SetCount(v int32) *DescribePolicyGroupsResponseBody {
	s.Count = &v
	return s
}

func (s *DescribePolicyGroupsResponseBody) SetDescribePolicyGroups(v []*DescribePolicyGroupsResponseBodyDescribePolicyGroups) *DescribePolicyGroupsResponseBody {
	s.DescribePolicyGroups = v
	return s
}

func (s *DescribePolicyGroupsResponseBody) SetNextToken(v string) *DescribePolicyGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribePolicyGroupsResponseBody) SetPageNumber(v int32) *DescribePolicyGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePolicyGroupsResponseBody) SetPageSize(v int32) *DescribePolicyGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePolicyGroupsResponseBody) SetRequestId(v string) *DescribePolicyGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePolicyGroupsResponseBody) Validate() error {
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

type DescribePolicyGroupsResponseBodyDescribePolicyGroups struct {
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
	// Indicates whether the user has administrator permissions after connecting to the cloud computer.
	//
	// > This feature is in invitational preview and is not publicly available.
	//
	// example:
	//
	// deny
	AdminAccess *string `json:"AdminAccess,omitempty" xml:"AdminAccess,omitempty"`
	// Specifies whether the administrator keyboard control in full-screen mode is enabled. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	AdminKeyboardOnFullScreen *string `json:"AdminKeyboardOnFullScreen,omitempty" xml:"AdminKeyboardOnFullScreen,omitempty"`
	// Specifies whether the administrator keyboard control within the Windows system is enabled. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	AdminKeyboardOnWindows *string `json:"AdminKeyboardOnWindows,omitempty" xml:"AdminKeyboardOnWindows,omitempty"`
	// Specifies whether the screenshot prevention feature is enabled.
	//
	// example:
	//
	// off
	AppContentProtection *string `json:"AppContentProtection,omitempty" xml:"AppContentProtection,omitempty"`
	// The client IP whitelist. Only IP addresses within the whitelisted CIDR blocks can access cloud desktops.
	AuthorizeAccessPolicyRules []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules `json:"AuthorizeAccessPolicyRules,omitempty" xml:"AuthorizeAccessPolicyRules,omitempty" type:"Repeated"`
	// The list of security group rules.
	AuthorizeSecurityPolicyRules []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules `json:"AuthorizeSecurityPolicyRules,omitempty" xml:"AuthorizeSecurityPolicyRules,omitempty" type:"Repeated"`
	// The client auto-reconnect configuration.
	//
	// example:
	//
	// off
	AutoReconnect *string `json:"AutoReconnect,omitempty" xml:"AutoReconnect,omitempty"`
	// Specifies whether local camera redirection is enabled.
	//
	// example:
	//
	// on
	CameraRedirect *string `json:"CameraRedirect,omitempty" xml:"CameraRedirect,omitempty"`
	// The client control menu display switch. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	ClientControlMenu *string `json:"ClientControlMenu,omitempty" xml:"ClientControlMenu,omitempty"`
	// Specifies whether the client custom snapshot creation feature is enabled. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	ClientCreateSnapshot *string `json:"ClientCreateSnapshot,omitempty" xml:"ClientCreateSnapshot,omitempty"`
	// Specifies whether the hibernate option in the client menu is enabled. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	ClientHibernate *string `json:"ClientHibernate,omitempty" xml:"ClientHibernate,omitempty"`
	// Specifies whether the restart option in the client menu is enabled. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	ClientRestart *string `json:"ClientRestart,omitempty" xml:"ClientRestart,omitempty"`
	// Specifies whether the shutdown option in the client menu is enabled. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	ClientShutdown *string `json:"ClientShutdown,omitempty" xml:"ClientShutdown,omitempty"`
	// The logon method control list. Specifies which client types are allowed to access cloud desktops.
	ClientTypes []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsClientTypes `json:"ClientTypes,omitempty" xml:"ClientTypes,omitempty" type:"Repeated"`
	// The clipboard permission.
	//
	// example:
	//
	// off
	Clipboard *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	// Indicates whether color enhancement is enabled for common scenarios of design and 3D applications.
	//
	// example:
	//
	// off
	ColorEnhancement *string `json:"ColorEnhancement,omitempty" xml:"ColorEnhancement,omitempty"`
	// Specifies whether the local drive clipboard feature is enabled. Valid values:
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
	// 30
	CpuDownGradeDuration *int32 `json:"CpuDownGradeDuration,omitempty" xml:"CpuDownGradeDuration,omitempty"`
	// Specifies whether CPU overload protection is enabled. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	CpuOverload *string `json:"CpuOverload,omitempty" xml:"CpuOverload,omitempty"`
	// The whitelist of processes that are not subject to CPU usage limits.
	CpuProcessors []*string `json:"CpuProcessors,omitempty" xml:"CpuProcessors,omitempty" type:"Repeated"`
	// Specifies whether to enable CPU protection mode.
	//
	// example:
	//
	// on
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
	// 10
	CpuSampleDuration *int32 `json:"CpuSampleDuration,omitempty" xml:"CpuSampleDuration,omitempty"`
	// The single-core CPU usage percentage. Valid values: 70 to 100.
	//
	// example:
	//
	// 70
	CpuSingleRateLimit *int32 `json:"CpuSingleRateLimit,omitempty" xml:"CpuSingleRateLimit,omitempty"`
	// The description of the NAS file system.
	//
	// example:
	//
	// newDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The number of cloud computers associated with the policy.
	//
	// example:
	//
	// 1
	DesktopCount *int32 `json:"DesktopCount,omitempty" xml:"DesktopCount,omitempty"`
	// The number of cloud computer pools associated with the policy.
	//
	// example:
	//
	// 1
	DesktopGroupCount *int32 `json:"DesktopGroupCount,omitempty" xml:"DesktopGroupCount,omitempty"`
	// The peripheral connection hint control.
	//
	// example:
	//
	// off
	DeviceConnectHint *string `json:"DeviceConnectHint,omitempty" xml:"DeviceConnectHint,omitempty"`
	// The list of device redirection rules.
	DeviceRedirects []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRedirects `json:"DeviceRedirects,omitempty" xml:"DeviceRedirects,omitempty" type:"Repeated"`
	// The list of custom peripheral rules.
	DeviceRules []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRules `json:"DeviceRules,omitempty" xml:"DeviceRules,omitempty" type:"Repeated"`
	// Specifies whether disk overload protection is enabled. Valid values:
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
	// The access domain name permission control. Domain names support wildcards (\\*). Separate multiple domain names with commas (,).
	//
	// example:
	//
	// off
	DomainList *string `json:"DomainList,omitempty" xml:"DomainList,omitempty"`
	// The domain name resolution policy list.
	DomainResolveRule []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsDomainResolveRule `json:"DomainResolveRule,omitempty" xml:"DomainResolveRule,omitempty" type:"Repeated"`
	// The switch for the domain name resolution policy.
	//
	// example:
	//
	// on
	DomainResolveRuleType *string `json:"DomainResolveRuleType,omitempty" xml:"DomainResolveRuleType,omitempty"`
	// The total number of cloud computers and cloud computer pools associated with this policy. This value is returned only for custom policies.
	//
	// example:
	//
	// 2
	EdsCount *int32 `json:"EdsCount,omitempty" xml:"EdsCount,omitempty"`
	// Specifies whether to enable the feature that allows users to request administrator assistance.
	//
	// example:
	//
	// on
	EndUserApplyAdminCoordinate *string `json:"EndUserApplyAdminCoordinate,omitempty" xml:"EndUserApplyAdminCoordinate,omitempty"`
	// The number of associated end users.
	//
	// example:
	//
	// 3
	EndUserCount *string `json:"EndUserCount,omitempty" xml:"EndUserCount,omitempty"`
	// Specifies whether to enable stream collaboration between users.
	//
	// example:
	//
	// on
	EndUserGroupCoordinate *string `json:"EndUserGroupCoordinate,omitempty" xml:"EndUserGroupCoordinate,omitempty"`
	// Specifies whether the use of external storage devices is enabled. Valid values:
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
	// The file transfer setting.
	//
	// example:
	//
	// off
	FileTransfer *string `json:"FileTransfer,omitempty" xml:"FileTransfer,omitempty"`
	// The service address for the file transfer feature.
	//
	// example:
	//
	// filetransfer.example.com
	FileTransferAddress *string `json:"FileTransferAddress,omitempty" xml:"FileTransferAddress,omitempty"`
	// The file size limit for a single file transfer to the cloud desktop. Use this parameter together with the inbound unit parameter.
	//
	// example:
	//
	// 100
	FileTransferInSize *int32 `json:"FileTransferInSize,omitempty" xml:"FileTransferInSize,omitempty"`
	// The unit for the file size limit of a single file transfer to the cloud desktop.
	//
	// example:
	//
	// MB
	FileTransferInUnit *string `json:"FileTransferInUnit,omitempty" xml:"FileTransferInUnit,omitempty"`
	// The file size limit for a single file transfer from the cloud desktop. Use this parameter together with the outbound unit parameter.
	//
	// example:
	//
	// 100
	FileTransferOutSize *int32 `json:"FileTransferOutSize,omitempty" xml:"FileTransferOutSize,omitempty"`
	// The unit for the file size limit of a single file transfer from the cloud desktop.
	//
	// example:
	//
	// MB
	FileTransferOutUnit *string `json:"FileTransferOutUnit,omitempty" xml:"FileTransferOutUnit,omitempty"`
	// Specifies whether the file transfer size limit is enabled. Valid values:
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
	// Specifies whether the image quality policy is enabled for GPU-accelerated cloud desktops. Enable this policy when high performance and user experience are required, such as in professional design scenarios.
	//
	// example:
	//
	// off
	GpuAcceleration *string `json:"GpuAcceleration,omitempty" xml:"GpuAcceleration,omitempty"`
	// Specifies whether the floating ball configuration message prompt is enabled. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	HoverConfigMsg *string `json:"HoverConfigMsg,omitempty" xml:"HoverConfigMsg,omitempty"`
	// Specifies whether the hibernate button on the floating ball is enabled. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	HoverHibernate *string `json:"HoverHibernate,omitempty" xml:"HoverHibernate,omitempty"`
	// Specifies whether the restart button on the floating ball is enabled. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	HoverRestart *string `json:"HoverRestart,omitempty" xml:"HoverRestart,omitempty"`
	// Specifies whether the shutdown button on the floating ball is enabled. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	HoverShutdown *string `json:"HoverShutdown,omitempty" xml:"HoverShutdown,omitempty"`
	// The web client access policy.
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
	// BOTH
	InternetCommunicationProtocol *string `json:"InternetCommunicationProtocol,omitempty" xml:"InternetCommunicationProtocol,omitempty"`
	// The network printer feature switch. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	InternetPrinter *string `json:"InternetPrinter,omitempty" xml:"InternetPrinter,omitempty"`
	// Specifies whether the keyboard control on the floating ball is enabled. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	KeyboardControl *string `json:"KeyboardControl,omitempty" xml:"KeyboardControl,omitempty"`
	// The local drive mapping permission.
	//
	// example:
	//
	// readwrite
	LocalDrive *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	// The maximum reconnection retry time when the cloud computer is disconnected due to external reasons. Valid values: 30 to 7200. Unit: seconds.
	//
	// example:
	//
	// 120
	MaxReconnectTime *int32 `json:"MaxReconnectTime,omitempty" xml:"MaxReconnectTime,omitempty"`
	// The memory throttling duration of a single process. Valid values: 30 to 120. Unit: seconds.
	//
	// example:
	//
	// 30
	MemoryDownGradeDuration *int32 `json:"MemoryDownGradeDuration,omitempty" xml:"MemoryDownGradeDuration,omitempty"`
	// Specifies whether memory overload protection is enabled. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	MemoryOverload *string `json:"MemoryOverload,omitempty" xml:"MemoryOverload,omitempty"`
	// The whitelist of processes that are not subject to memory usage limits.
	MemoryProcessors []*string `json:"MemoryProcessors,omitempty" xml:"MemoryProcessors,omitempty" type:"Repeated"`
	// Specifies whether to enable memory protection mode.
	//
	// example:
	//
	// on
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
	// The memory usage percentage of a single process. Valid values: 30 to 60.
	//
	// example:
	//
	// 30
	MemorySingleRateLimit *int32 `json:"MemorySingleRateLimit,omitempty" xml:"MemorySingleRateLimit,omitempty"`
	// Specifies whether the restart button is provided in the cloud computer floating ball when connecting to a cloud computer from a mobile client (including Android and iOS clients).
	//
	// > This applies only to mobile clients of V7.4 or later.
	//
	// example:
	//
	// off
	MobileRestart *string `json:"MobileRestart,omitempty" xml:"MobileRestart,omitempty"`
	// Specifies whether the Windows security control feature is enabled on mobile clients.
	//
	// example:
	//
	// off
	MobileSafeMenu *string `json:"MobileSafeMenu,omitempty" xml:"MobileSafeMenu,omitempty"`
	// Specifies whether the shutdown button is provided in the cloud computer floating ball when connecting to a cloud computer from a mobile client (including Android and iOS clients).
	//
	// > This applies only to mobile clients of V7.4 or later.
	//
	// example:
	//
	// off
	MobileShutdown *string `json:"MobileShutdown,omitempty" xml:"MobileShutdown,omitempty"`
	// Specifies whether the WUYING Keeper feature is enabled on mobile clients.
	//
	// example:
	//
	// off
	MobileWuyingKeeper *string `json:"MobileWuyingKeeper,omitempty" xml:"MobileWuyingKeeper,omitempty"`
	// Specifies whether the WUYING Assistant feature is enabled on mobile clients.
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
	// Specifies whether the multi-screen display feature is enabled. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	MultiScreen *string `json:"MultiScreen,omitempty" xml:"MultiScreen,omitempty"`
	// The Policy Name of the cloud computer policy.
	//
	// example:
	//
	// testPolicyGroupName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The network redirection setting.
	//
	// > This feature is in invitational preview and is not publicly available.
	//
	// example:
	//
	// off
	NetRedirect *string `json:"NetRedirect,omitempty" xml:"NetRedirect,omitempty"`
	// The network redirection policy list.
	//
	// > This feature is in invitational preview and is not publicly available.
	NetRedirectRule []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsNetRedirectRule `json:"NetRedirectRule,omitempty" xml:"NetRedirectRule,omitempty" type:"Repeated"`
	// The network printer feature switch. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	NetworkPrinter *string `json:"NetworkPrinter,omitempty" xml:"NetworkPrinter,omitempty"`
	// The number of associated organizations.
	//
	// example:
	//
	// 2
	OrganizationCount *string `json:"OrganizationCount,omitempty" xml:"OrganizationCount,omitempty"`
	// The cloud computer policy ID.
	//
	// example:
	//
	// pg-gx2x1dhsmthe9****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The type of the cloud computer policy.
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
	// The preemption policy for the cloud computer.
	//
	// > To ensure the user experience and data security of end users who are using cloud computers, preemption among multiple users is not allowed. This means the configuration is set to `off` by default and cannot be modified.
	//
	// example:
	//
	// off
	PreemptLogin *string `json:"PreemptLogin,omitempty" xml:"PreemptLogin,omitempty"`
	// The usernames of users who can preempt cloud desktops.
	PreemptLoginUsers []*string `json:"PreemptLoginUsers,omitempty" xml:"PreemptLoginUsers,omitempty" type:"Repeated"`
	// The printer pop-up alert setting. Valid values:
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
	// The content of the printer pop-up alert.
	//
	// example:
	//
	// Print Content
	PrinterAlertContent *string `json:"PrinterAlertContent,omitempty" xml:"PrinterAlertContent,omitempty"`
	// The title of the printer pop-up alert.
	//
	// example:
	//
	// Print Title
	PrinterAlertTitle *string `json:"PrinterAlertTitle,omitempty" xml:"PrinterAlertTitle,omitempty"`
	// The printer redirection policy.
	//
	// example:
	//
	// on
	PrinterRedirection *string `json:"PrinterRedirection,omitempty" xml:"PrinterRedirection,omitempty"`
	// Specifies whether image quality enhancement is enabled for design and 3D common scenarios.
	//
	// example:
	//
	// off
	QualityEnhancement *string `json:"QualityEnhancement,omitempty" xml:"QualityEnhancement,omitempty"`
	// Specifies whether custom screen recording is enabled.
	//
	// example:
	//
	// off
	RecordContent *string `json:"RecordContent,omitempty" xml:"RecordContent,omitempty"`
	// The expiration time of custom screen recording files. Default value: 30 days.
	//
	// example:
	//
	// 30
	RecordContentExpires *int64 `json:"RecordContentExpires,omitempty" xml:"RecordContentExpires,omitempty"`
	// The recording duration after an event is detected in screen recording audit. Unit: minutes. Valid values: 10 to 60.
	//
	// example:
	//
	// 10
	RecordEventDuration *int32 `json:"RecordEventDuration,omitempty" xml:"RecordEventDuration,omitempty"`
	// The file extensions for screen recording events.
	RecordEventFileExts []*string `json:"RecordEventFileExts,omitempty" xml:"RecordEventFileExts,omitempty" type:"Repeated"`
	// The list of absolute paths for file monitoring in screen recording audit.
	RecordEventFilePaths []*string `json:"RecordEventFilePaths,omitempty" xml:"RecordEventFilePaths,omitempty" type:"Repeated"`
	// The screen recording event level settings.
	RecordEventLevels []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsRecordEventLevels `json:"RecordEventLevels,omitempty" xml:"RecordEventLevels,omitempty" type:"Repeated"`
	// The list of absolute paths for registry monitoring in screen recording audit.
	RecordEventRegisters []*string `json:"RecordEventRegisters,omitempty" xml:"RecordEventRegisters,omitempty" type:"Repeated"`
	// Specifies whether screen recording is enabled.
	//
	// example:
	//
	// OFF
	Recording *string `json:"Recording,omitempty" xml:"Recording,omitempty"`
	// The option for recording cloud computer audio.
	//
	// example:
	//
	// on
	RecordingAudio *string `json:"RecordingAudio,omitempty" xml:"RecordingAudio,omitempty"`
	// The duration of a screen recording file, in minutes. Recording files are automatically split and uploaded to the storage space based on the duration you specify. When a file reaches 300 MB, it is rolled over first.
	//
	// example:
	//
	// 10
	RecordingDuration *int32 `json:"RecordingDuration,omitempty" xml:"RecordingDuration,omitempty"`
	// The screen recording end time in the format of HH:MM:SS. This parameter is meaningful only when Recording is set to PERIOD.
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
	// The screen recording frame rate. Unit: FPS (frames per second).
	//
	// example:
	//
	// 5
	RecordingFps *int64 `json:"RecordingFps,omitempty" xml:"RecordingFps,omitempty"`
	// The screen recording start time in the format of HH:MM:SS. This parameter is meaningful only when Recording is set to PERIOD.
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
	// Your cloud computer is being recorded
	RecordingUserNotifyMessage *string `json:"RecordingUserNotifyMessage,omitempty" xml:"RecordingUserNotifyMessage,omitempty"`
	// The keyboard and mouse control permission for remote assistance.
	//
	// example:
	//
	// fullControl
	RemoteCoordinate *string `json:"RemoteCoordinate,omitempty" xml:"RemoteCoordinate,omitempty"`
	// The setting for resetting the cloud computer.
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
	// The height of the resolution. Unit: pixels. Valid values for cloud applications: 500 to 50000. Valid values for cloud computers: 480 to 4096.
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
	// The width of the resolution. Unit: pixels. Valid values for cloud applications: 500 to 50000. Valid values for cloud computers: 640 to 4096.
	//
	// example:
	//
	// 720
	ResolutionWidth *int32 `json:"ResolutionWidth,omitempty" xml:"ResolutionWidth,omitempty"`
	// The number of resource groups associated with the policy.
	//
	// example:
	//
	// 1
	ResourceGroupCount *int32 `json:"ResourceGroupCount,omitempty" xml:"ResourceGroupCount,omitempty"`
	// The region to which the cloud computer policy belongs.
	//
	// > If the policy is a region-independent policy, this value is `center`.
	//
	// example:
	//
	// center
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// The security center shortcut key switch. Valid values:
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
	// This parameter is required when `Scope` is set to `IP`. This parameter takes effect only when `Scope` is set to `IP`.
	ScopeValue []*string `json:"ScopeValue,omitempty" xml:"ScopeValue,omitempty" type:"Repeated"`
	// The screen display mode.
	//
	// example:
	//
	// auto
	ScreenDisplayMode *string `json:"ScreenDisplayMode,omitempty" xml:"ScreenDisplayMode,omitempty"`
	// Specifies whether smoothness enhancement is enabled for daily office scenarios.
	//
	// example:
	//
	// off
	SmoothEnhancement *string `json:"SmoothEnhancement,omitempty" xml:"SmoothEnhancement,omitempty"`
	// Specifies whether the status monitoring entry is provided in the cloud computer floating ball.
	//
	// example:
	//
	// on
	StatusMonitor *string `json:"StatusMonitor,omitempty" xml:"StatusMonitor,omitempty"`
	// The streaming mode for scenario adaptation.
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
	// Specifies whether the three-screen feature is enabled. Valid values:
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
	// on
	UsbRedirect *string `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty"`
	// The USB redirection rules.
	UsbSupplyRedirectRule []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule `json:"UsbSupplyRedirectRule,omitempty" xml:"UsbSupplyRedirectRule,omitempty" type:"Repeated"`
	// Specifies whether the usage duration display on the floating ball is enabled. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	UseTime *string `json:"UseTime,omitempty" xml:"UseTime,omitempty"`
	// The average bitrate for video encoding. Valid values: 1000 to 50000.
	//
	// example:
	//
	// 1000
	VideoEncAvgKbps *int32 `json:"VideoEncAvgKbps,omitempty" xml:"VideoEncAvgKbps,omitempty"`
	// The maximum QP for video encoding, which represents the lowest image quality. Valid values: 0 to 51.
	//
	// example:
	//
	// 20
	VideoEncMaxQP *int32 `json:"VideoEncMaxQP,omitempty" xml:"VideoEncMaxQP,omitempty"`
	// The minimum QP for video encoding, which represents the highest quality. Valid values: 0 to 51.
	//
	// example:
	//
	// 20
	VideoEncMinQP *int32 `json:"VideoEncMinQP,omitempty" xml:"VideoEncMinQP,omitempty"`
	// The peak bitrate for video encoding. Valid values: 1000 to 50000.
	//
	// example:
	//
	// 1000
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
	// The watermark policy.
	//
	// example:
	//
	// on
	Watermark *string `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
	// The anti-camera capture feature for invisible watermarks.
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
	// If the `WatermarkType` parameter is set to `custom`, you must also specify the custom text content by using the `WatermarkCustomText` parameter.
	//
	// example:
	//
	// custom-watermark
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
	// The enhancement feature for invisible watermarks.
	//
	// example:
	//
	// medium
	WatermarkPower *string `json:"WatermarkPower,omitempty" xml:"WatermarkPower,omitempty"`
	// The number of watermark rows.
	//
	// > This parameter is not yet available for use.
	//
	// example:
	//
	// 5
	WatermarkRowAmount *int32 `json:"WatermarkRowAmount,omitempty" xml:"WatermarkRowAmount,omitempty"`
	// The security-first rule for invisible watermarks.
	//
	// example:
	//
	// on
	WatermarkSecurity *string `json:"WatermarkSecurity,omitempty" xml:"WatermarkSecurity,omitempty"`
	// Specifies whether the watermark shadow effect is enabled. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	WatermarkShadow *string `json:"WatermarkShadow,omitempty" xml:"WatermarkShadow,omitempty"`
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
	// The watermark type.
	//
	// example:
	//
	// EndUserId
	WatermarkType *string `json:"WatermarkType,omitempty" xml:"WatermarkType,omitempty"`
	// The WUYING Keeper switch.
	//
	// example:
	//
	// off
	WuyingKeeper *string `json:"WuyingKeeper,omitempty" xml:"WuyingKeeper,omitempty"`
	// Specifies whether the WUYING AI Assistant entry is provided in the cloud computer floating ball.
	//
	// example:
	//
	// on
	WyAssistant *string `json:"WyAssistant,omitempty" xml:"WyAssistant,omitempty"`
}

func (s DescribePolicyGroupsResponseBodyDescribePolicyGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyGroupsResponseBodyDescribePolicyGroups) GoString() string {
	return s.String()
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetAcademicProxy() *string {
	return s.AcademicProxy
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetAdminAccess() *string {
	return s.AdminAccess
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetAdminKeyboardOnFullScreen() *string {
	return s.AdminKeyboardOnFullScreen
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetAdminKeyboardOnWindows() *string {
	return s.AdminKeyboardOnWindows
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetAppContentProtection() *string {
	return s.AppContentProtection
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetAuthorizeAccessPolicyRules() []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules {
	return s.AuthorizeAccessPolicyRules
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetAuthorizeSecurityPolicyRules() []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	return s.AuthorizeSecurityPolicyRules
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetAutoReconnect() *string {
	return s.AutoReconnect
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetCameraRedirect() *string {
	return s.CameraRedirect
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetClientControlMenu() *string {
	return s.ClientControlMenu
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetClientCreateSnapshot() *string {
	return s.ClientCreateSnapshot
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetClientHibernate() *string {
	return s.ClientHibernate
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetClientRestart() *string {
	return s.ClientRestart
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetClientShutdown() *string {
	return s.ClientShutdown
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetClientTypes() []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsClientTypes {
	return s.ClientTypes
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetClipboard() *string {
	return s.Clipboard
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetColorEnhancement() *string {
	return s.ColorEnhancement
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetCpdDriveClipboard() *string {
	return s.CpdDriveClipboard
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetCpuDownGradeDuration() *int32 {
	return s.CpuDownGradeDuration
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetCpuOverload() *string {
	return s.CpuOverload
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetCpuProcessors() []*string {
	return s.CpuProcessors
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetCpuProtectedMode() *string {
	return s.CpuProtectedMode
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetCpuRateLimit() *int32 {
	return s.CpuRateLimit
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetCpuSampleDuration() *int32 {
	return s.CpuSampleDuration
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetCpuSingleRateLimit() *int32 {
	return s.CpuSingleRateLimit
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetDescription() *string {
	return s.Description
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetDesktopCount() *int32 {
	return s.DesktopCount
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetDesktopGroupCount() *int32 {
	return s.DesktopGroupCount
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetDeviceConnectHint() *string {
	return s.DeviceConnectHint
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetDeviceRedirects() []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRedirects {
	return s.DeviceRedirects
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetDeviceRules() []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRules {
	return s.DeviceRules
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetDiskOverload() *string {
	return s.DiskOverload
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetDisplayMode() *string {
	return s.DisplayMode
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetDomainList() *string {
	return s.DomainList
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetDomainResolveRule() []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsDomainResolveRule {
	return s.DomainResolveRule
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetDomainResolveRuleType() *string {
	return s.DomainResolveRuleType
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetEdsCount() *int32 {
	return s.EdsCount
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetEndUserApplyAdminCoordinate() *string {
	return s.EndUserApplyAdminCoordinate
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetEndUserCount() *string {
	return s.EndUserCount
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetEndUserGroupCoordinate() *string {
	return s.EndUserGroupCoordinate
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetExternalDrive() *string {
	return s.ExternalDrive
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetFileMigrate() *string {
	return s.FileMigrate
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetFileTransfer() *string {
	return s.FileTransfer
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetFileTransferAddress() *string {
	return s.FileTransferAddress
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetFileTransferInSize() *int32 {
	return s.FileTransferInSize
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetFileTransferInUnit() *string {
	return s.FileTransferInUnit
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetFileTransferOutSize() *int32 {
	return s.FileTransferOutSize
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetFileTransferOutUnit() *string {
	return s.FileTransferOutUnit
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetFileTransferSizeLimit() *string {
	return s.FileTransferSizeLimit
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetFileTransferSpeed() *string {
	return s.FileTransferSpeed
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetFileTransferSpeedLocation() *string {
	return s.FileTransferSpeedLocation
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetGpuAcceleration() *string {
	return s.GpuAcceleration
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetHoverConfigMsg() *string {
	return s.HoverConfigMsg
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetHoverHibernate() *string {
	return s.HoverHibernate
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetHoverRestart() *string {
	return s.HoverRestart
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetHoverShutdown() *string {
	return s.HoverShutdown
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetHtml5Access() *string {
	return s.Html5Access
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetHtml5FileTransfer() *string {
	return s.Html5FileTransfer
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetInternetCommunicationProtocol() *string {
	return s.InternetCommunicationProtocol
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetInternetPrinter() *string {
	return s.InternetPrinter
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetKeyboardControl() *string {
	return s.KeyboardControl
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetLocalDrive() *string {
	return s.LocalDrive
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetMaxReconnectTime() *int32 {
	return s.MaxReconnectTime
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetMemoryDownGradeDuration() *int32 {
	return s.MemoryDownGradeDuration
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetMemoryOverload() *string {
	return s.MemoryOverload
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetMemoryProcessors() []*string {
	return s.MemoryProcessors
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetMemoryProtectedMode() *string {
	return s.MemoryProtectedMode
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetMemoryRateLimit() *int32 {
	return s.MemoryRateLimit
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetMemorySampleDuration() *int32 {
	return s.MemorySampleDuration
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetMemorySingleRateLimit() *int32 {
	return s.MemorySingleRateLimit
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetMobileRestart() *string {
	return s.MobileRestart
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetMobileSafeMenu() *string {
	return s.MobileSafeMenu
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetMobileShutdown() *string {
	return s.MobileShutdown
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetMobileWuyingKeeper() *string {
	return s.MobileWuyingKeeper
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetMobileWyAssistant() *string {
	return s.MobileWyAssistant
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetModelLibrary() *string {
	return s.ModelLibrary
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetMultiScreen() *string {
	return s.MultiScreen
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetName() *string {
	return s.Name
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetNetRedirect() *string {
	return s.NetRedirect
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetNetRedirectRule() []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsNetRedirectRule {
	return s.NetRedirectRule
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetNetworkPrinter() *string {
	return s.NetworkPrinter
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetOrganizationCount() *string {
	return s.OrganizationCount
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetPolicyGroupType() *string {
	return s.PolicyGroupType
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetPolicyStatus() *string {
	return s.PolicyStatus
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetPortProxy() *string {
	return s.PortProxy
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetPreemptLogin() *string {
	return s.PreemptLogin
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetPreemptLoginUsers() []*string {
	return s.PreemptLoginUsers
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetPrinterAlert() *string {
	return s.PrinterAlert
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetPrinterAlertContent() *string {
	return s.PrinterAlertContent
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetPrinterAlertTitle() *string {
	return s.PrinterAlertTitle
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetPrinterRedirection() *string {
	return s.PrinterRedirection
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetQualityEnhancement() *string {
	return s.QualityEnhancement
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetRecordContent() *string {
	return s.RecordContent
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetRecordContentExpires() *int64 {
	return s.RecordContentExpires
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetRecordEventDuration() *int32 {
	return s.RecordEventDuration
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetRecordEventFileExts() []*string {
	return s.RecordEventFileExts
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetRecordEventFilePaths() []*string {
	return s.RecordEventFilePaths
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetRecordEventLevels() []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsRecordEventLevels {
	return s.RecordEventLevels
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetRecordEventRegisters() []*string {
	return s.RecordEventRegisters
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetRecording() *string {
	return s.Recording
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetRecordingAudio() *string {
	return s.RecordingAudio
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetRecordingDuration() *int32 {
	return s.RecordingDuration
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetRecordingEndTime() *string {
	return s.RecordingEndTime
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetRecordingExpires() *int64 {
	return s.RecordingExpires
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetRecordingFps() *int64 {
	return s.RecordingFps
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetRecordingStartTime() *string {
	return s.RecordingStartTime
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetRecordingUserNotify() *string {
	return s.RecordingUserNotify
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetRecordingUserNotifyMessage() *string {
	return s.RecordingUserNotifyMessage
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetRemoteCoordinate() *string {
	return s.RemoteCoordinate
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetResetDesktop() *string {
	return s.ResetDesktop
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetResolutionDpi() *int32 {
	return s.ResolutionDpi
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetResolutionHeight() *int32 {
	return s.ResolutionHeight
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetResolutionModel() *string {
	return s.ResolutionModel
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetResolutionWidth() *int32 {
	return s.ResolutionWidth
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetResourceGroupCount() *int32 {
	return s.ResourceGroupCount
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetSafeMenu() *string {
	return s.SafeMenu
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetScope() *string {
	return s.Scope
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetScopeValue() []*string {
	return s.ScopeValue
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetScreenDisplayMode() *string {
	return s.ScreenDisplayMode
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetSmoothEnhancement() *string {
	return s.SmoothEnhancement
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetStatusMonitor() *string {
	return s.StatusMonitor
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetStreamingMode() *string {
	return s.StreamingMode
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetTargetFps() *int32 {
	return s.TargetFps
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetThreeScreen() *string {
	return s.ThreeScreen
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetUsbRedirect() *string {
	return s.UsbRedirect
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetUsbSupplyRedirectRule() []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule {
	return s.UsbSupplyRedirectRule
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetUseTime() *string {
	return s.UseTime
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetVideoEncAvgKbps() *int32 {
	return s.VideoEncAvgKbps
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetVideoEncMaxQP() *int32 {
	return s.VideoEncMaxQP
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetVideoEncMinQP() *int32 {
	return s.VideoEncMinQP
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetVideoEncPeakKbps() *int32 {
	return s.VideoEncPeakKbps
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetVideoEncPolicy() *string {
	return s.VideoEncPolicy
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetVideoRedirect() *string {
	return s.VideoRedirect
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetVisualQuality() *string {
	return s.VisualQuality
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetWatermark() *string {
	return s.Watermark
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetWatermarkAntiCam() *string {
	return s.WatermarkAntiCam
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetWatermarkColor() *int32 {
	return s.WatermarkColor
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetWatermarkCustomText() *string {
	return s.WatermarkCustomText
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetWatermarkDegree() *float64 {
	return s.WatermarkDegree
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetWatermarkFontSize() *int32 {
	return s.WatermarkFontSize
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetWatermarkFontStyle() *string {
	return s.WatermarkFontStyle
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetWatermarkPower() *string {
	return s.WatermarkPower
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetWatermarkRowAmount() *int32 {
	return s.WatermarkRowAmount
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetWatermarkSecurity() *string {
	return s.WatermarkSecurity
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetWatermarkShadow() *string {
	return s.WatermarkShadow
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetWatermarkTransparency() *string {
	return s.WatermarkTransparency
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetWatermarkTransparencyValue() *int32 {
	return s.WatermarkTransparencyValue
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetWatermarkType() *string {
	return s.WatermarkType
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetWuyingKeeper() *string {
	return s.WuyingKeeper
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) GetWyAssistant() *string {
	return s.WyAssistant
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetAcademicProxy(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.AcademicProxy = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetAdminAccess(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.AdminAccess = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetAdminKeyboardOnFullScreen(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.AdminKeyboardOnFullScreen = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetAdminKeyboardOnWindows(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.AdminKeyboardOnWindows = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetAppContentProtection(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.AppContentProtection = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetAuthorizeAccessPolicyRules(v []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.AuthorizeAccessPolicyRules = v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetAuthorizeSecurityPolicyRules(v []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.AuthorizeSecurityPolicyRules = v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetAutoReconnect(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.AutoReconnect = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetCameraRedirect(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.CameraRedirect = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetClientControlMenu(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.ClientControlMenu = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetClientCreateSnapshot(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.ClientCreateSnapshot = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetClientHibernate(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.ClientHibernate = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetClientRestart(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.ClientRestart = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetClientShutdown(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.ClientShutdown = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetClientTypes(v []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsClientTypes) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.ClientTypes = v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetClipboard(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.Clipboard = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetColorEnhancement(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.ColorEnhancement = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetCpdDriveClipboard(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.CpdDriveClipboard = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetCpuDownGradeDuration(v int32) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.CpuDownGradeDuration = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetCpuOverload(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.CpuOverload = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetCpuProcessors(v []*string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.CpuProcessors = v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetCpuProtectedMode(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.CpuProtectedMode = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetCpuRateLimit(v int32) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.CpuRateLimit = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetCpuSampleDuration(v int32) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.CpuSampleDuration = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetCpuSingleRateLimit(v int32) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.CpuSingleRateLimit = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetDescription(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.Description = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetDesktopCount(v int32) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.DesktopCount = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetDesktopGroupCount(v int32) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.DesktopGroupCount = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetDeviceConnectHint(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.DeviceConnectHint = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetDeviceRedirects(v []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRedirects) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.DeviceRedirects = v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetDeviceRules(v []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRules) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.DeviceRules = v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetDiskOverload(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.DiskOverload = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetDisplayMode(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.DisplayMode = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetDomainList(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.DomainList = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetDomainResolveRule(v []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsDomainResolveRule) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.DomainResolveRule = v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetDomainResolveRuleType(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.DomainResolveRuleType = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetEdsCount(v int32) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.EdsCount = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetEndUserApplyAdminCoordinate(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.EndUserApplyAdminCoordinate = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetEndUserCount(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.EndUserCount = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetEndUserGroupCoordinate(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.EndUserGroupCoordinate = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetExternalDrive(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.ExternalDrive = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetFileMigrate(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.FileMigrate = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetFileTransfer(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.FileTransfer = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetFileTransferAddress(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.FileTransferAddress = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetFileTransferInSize(v int32) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.FileTransferInSize = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetFileTransferInUnit(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.FileTransferInUnit = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetFileTransferOutSize(v int32) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.FileTransferOutSize = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetFileTransferOutUnit(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.FileTransferOutUnit = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetFileTransferSizeLimit(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.FileTransferSizeLimit = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetFileTransferSpeed(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.FileTransferSpeed = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetFileTransferSpeedLocation(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.FileTransferSpeedLocation = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetGpuAcceleration(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.GpuAcceleration = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetHoverConfigMsg(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.HoverConfigMsg = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetHoverHibernate(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.HoverHibernate = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetHoverRestart(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.HoverRestart = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetHoverShutdown(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.HoverShutdown = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetHtml5Access(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.Html5Access = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetHtml5FileTransfer(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.Html5FileTransfer = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetInternetCommunicationProtocol(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.InternetCommunicationProtocol = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetInternetPrinter(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.InternetPrinter = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetKeyboardControl(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.KeyboardControl = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetLocalDrive(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.LocalDrive = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetMaxReconnectTime(v int32) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.MaxReconnectTime = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetMemoryDownGradeDuration(v int32) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.MemoryDownGradeDuration = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetMemoryOverload(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.MemoryOverload = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetMemoryProcessors(v []*string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.MemoryProcessors = v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetMemoryProtectedMode(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.MemoryProtectedMode = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetMemoryRateLimit(v int32) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.MemoryRateLimit = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetMemorySampleDuration(v int32) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.MemorySampleDuration = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetMemorySingleRateLimit(v int32) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.MemorySingleRateLimit = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetMobileRestart(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.MobileRestart = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetMobileSafeMenu(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.MobileSafeMenu = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetMobileShutdown(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.MobileShutdown = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetMobileWuyingKeeper(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.MobileWuyingKeeper = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetMobileWyAssistant(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.MobileWyAssistant = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetModelLibrary(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.ModelLibrary = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetMultiScreen(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.MultiScreen = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetName(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.Name = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetNetRedirect(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.NetRedirect = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetNetRedirectRule(v []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsNetRedirectRule) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.NetRedirectRule = v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetNetworkPrinter(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.NetworkPrinter = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetOrganizationCount(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.OrganizationCount = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetPolicyGroupId(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.PolicyGroupId = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetPolicyGroupType(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.PolicyGroupType = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetPolicyStatus(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.PolicyStatus = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetPortProxy(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.PortProxy = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetPreemptLogin(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.PreemptLogin = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetPreemptLoginUsers(v []*string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.PreemptLoginUsers = v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetPrinterAlert(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.PrinterAlert = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetPrinterAlertContent(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.PrinterAlertContent = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetPrinterAlertTitle(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.PrinterAlertTitle = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetPrinterRedirection(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.PrinterRedirection = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetQualityEnhancement(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.QualityEnhancement = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetRecordContent(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.RecordContent = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetRecordContentExpires(v int64) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.RecordContentExpires = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetRecordEventDuration(v int32) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.RecordEventDuration = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetRecordEventFileExts(v []*string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.RecordEventFileExts = v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetRecordEventFilePaths(v []*string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.RecordEventFilePaths = v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetRecordEventLevels(v []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsRecordEventLevels) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.RecordEventLevels = v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetRecordEventRegisters(v []*string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.RecordEventRegisters = v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetRecording(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.Recording = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetRecordingAudio(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.RecordingAudio = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetRecordingDuration(v int32) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.RecordingDuration = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetRecordingEndTime(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.RecordingEndTime = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetRecordingExpires(v int64) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.RecordingExpires = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetRecordingFps(v int64) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.RecordingFps = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetRecordingStartTime(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.RecordingStartTime = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetRecordingUserNotify(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.RecordingUserNotify = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetRecordingUserNotifyMessage(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.RecordingUserNotifyMessage = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetRemoteCoordinate(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.RemoteCoordinate = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetResetDesktop(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.ResetDesktop = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetResolutionDpi(v int32) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.ResolutionDpi = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetResolutionHeight(v int32) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.ResolutionHeight = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetResolutionModel(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.ResolutionModel = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetResolutionWidth(v int32) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.ResolutionWidth = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetResourceGroupCount(v int32) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.ResourceGroupCount = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetResourceRegionId(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.ResourceRegionId = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetSafeMenu(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.SafeMenu = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetScope(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.Scope = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetScopeValue(v []*string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.ScopeValue = v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetScreenDisplayMode(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.ScreenDisplayMode = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetSmoothEnhancement(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.SmoothEnhancement = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetStatusMonitor(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.StatusMonitor = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetStreamingMode(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.StreamingMode = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetTargetFps(v int32) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.TargetFps = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetThreeScreen(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.ThreeScreen = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetUsbRedirect(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.UsbRedirect = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetUsbSupplyRedirectRule(v []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.UsbSupplyRedirectRule = v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetUseTime(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.UseTime = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetVideoEncAvgKbps(v int32) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.VideoEncAvgKbps = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetVideoEncMaxQP(v int32) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.VideoEncMaxQP = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetVideoEncMinQP(v int32) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.VideoEncMinQP = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetVideoEncPeakKbps(v int32) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.VideoEncPeakKbps = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetVideoEncPolicy(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.VideoEncPolicy = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetVideoRedirect(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.VideoRedirect = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetVisualQuality(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.VisualQuality = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetWatermark(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.Watermark = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetWatermarkAntiCam(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.WatermarkAntiCam = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetWatermarkColor(v int32) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.WatermarkColor = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetWatermarkCustomText(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.WatermarkCustomText = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetWatermarkDegree(v float64) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.WatermarkDegree = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetWatermarkFontSize(v int32) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.WatermarkFontSize = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetWatermarkFontStyle(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.WatermarkFontStyle = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetWatermarkPower(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.WatermarkPower = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetWatermarkRowAmount(v int32) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.WatermarkRowAmount = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetWatermarkSecurity(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.WatermarkSecurity = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetWatermarkShadow(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.WatermarkShadow = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetWatermarkTransparency(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.WatermarkTransparency = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetWatermarkTransparencyValue(v int32) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.WatermarkTransparencyValue = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetWatermarkType(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.WatermarkType = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetWuyingKeeper(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.WuyingKeeper = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetWyAssistant(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.WyAssistant = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) Validate() error {
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

type DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules struct {
	// The client access IP CIDR block. The value is an IPv4 CIDR block.
	//
	// example:
	//
	// 47.100.XX.XX/16
	CidrIp *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	// The description of the client access IP CIDR block.
	//
	// example:
	//
	// Corporate office network
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules) GoString() string {
	return s.String()
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules) GetCidrIp() *string {
	return s.CidrIp
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules) GetDescription() *string {
	return s.Description
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules) SetCidrIp(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules {
	s.CidrIp = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules) SetDescription(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules {
	s.Description = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules) Validate() error {
	return dara.Validate(s)
}

type DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules struct {
	// The target of the security group rule. The value is an IPv4 CIDR block.
	//
	// example:
	//
	// 47.100.XX.XX/16
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

func (s DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) GoString() string {
	return s.String()
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) GetCidrIp() *string {
	return s.CidrIp
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) GetDescription() *string {
	return s.Description
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) GetPolicy() *string {
	return s.Policy
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) GetPortRange() *string {
	return s.PortRange
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) GetPriority() *string {
	return s.Priority
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) GetType() *string {
	return s.Type
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetCidrIp(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.CidrIp = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetDescription(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.Description = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetIpProtocol(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.IpProtocol = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetPolicy(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.Policy = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetPortRange(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.PortRange = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetPriority(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.Priority = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetType(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.Type = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) Validate() error {
	return dara.Validate(s)
}

type DescribePolicyGroupsResponseBodyDescribePolicyGroupsClientTypes struct {
	// The client type.
	//
	// example:
	//
	// windows
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// Specifies whether a specific type of client is allowed to connect to cloud desktops.
	//
	// example:
	//
	// ON
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribePolicyGroupsResponseBodyDescribePolicyGroupsClientTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyGroupsResponseBodyDescribePolicyGroupsClientTypes) GoString() string {
	return s.String()
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsClientTypes) GetClientType() *string {
	return s.ClientType
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsClientTypes) GetStatus() *string {
	return s.Status
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsClientTypes) SetClientType(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsClientTypes {
	s.ClientType = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsClientTypes) SetStatus(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsClientTypes {
	s.Status = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsClientTypes) Validate() error {
	return dara.Validate(s)
}

type DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRedirects struct {
	// The peripheral type.
	//
	// example:
	//
	// camera
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// The redirection type. Valid values:
	//
	// - usbRedirect: USB redirection.
	//
	// - deviceRedirect: device redirection.
	//
	// - off: disabled.
	//
	// example:
	//
	// usbRedirect
	RedirectType *string `json:"RedirectType,omitempty" xml:"RedirectType,omitempty"`
}

func (s DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRedirects) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRedirects) GoString() string {
	return s.String()
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRedirects) GetDeviceType() *string {
	return s.DeviceType
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRedirects) GetRedirectType() *string {
	return s.RedirectType
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRedirects) SetDeviceType(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRedirects {
	s.DeviceType = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRedirects) SetRedirectType(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRedirects {
	s.RedirectType = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRedirects) Validate() error {
	return dara.Validate(s)
}

type DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRules struct {
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

func (s DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRules) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRules) GoString() string {
	return s.String()
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRules) GetDeviceName() *string {
	return s.DeviceName
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRules) GetDevicePid() *string {
	return s.DevicePid
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRules) GetDeviceType() *string {
	return s.DeviceType
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRules) GetDeviceVid() *string {
	return s.DeviceVid
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRules) GetOptCommand() *string {
	return s.OptCommand
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRules) GetPlatforms() *string {
	return s.Platforms
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRules) GetRedirectType() *string {
	return s.RedirectType
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRules) SetDeviceName(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRules {
	s.DeviceName = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRules) SetDevicePid(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRules {
	s.DevicePid = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRules) SetDeviceType(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRules {
	s.DeviceType = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRules) SetDeviceVid(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRules {
	s.DeviceVid = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRules) SetOptCommand(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRules {
	s.OptCommand = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRules) SetPlatforms(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRules {
	s.Platforms = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRules) SetRedirectType(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRules {
	s.RedirectType = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRules) Validate() error {
	return dara.Validate(s)
}

type DescribePolicyGroupsResponseBodyDescribePolicyGroupsDomainResolveRule struct {
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

func (s DescribePolicyGroupsResponseBodyDescribePolicyGroupsDomainResolveRule) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyGroupsResponseBodyDescribePolicyGroupsDomainResolveRule) GoString() string {
	return s.String()
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsDomainResolveRule) GetDescription() *string {
	return s.Description
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsDomainResolveRule) GetDomain() *string {
	return s.Domain
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsDomainResolveRule) GetPolicy() *string {
	return s.Policy
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsDomainResolveRule) SetDescription(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsDomainResolveRule {
	s.Description = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsDomainResolveRule) SetDomain(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsDomainResolveRule {
	s.Domain = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsDomainResolveRule) SetPolicy(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsDomainResolveRule {
	s.Policy = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsDomainResolveRule) Validate() error {
	return dara.Validate(s)
}

type DescribePolicyGroupsResponseBodyDescribePolicyGroupsNetRedirectRule struct {
	// The policy content.
	//
	// example:
	//
	// *.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The policy type.
	//
	// example:
	//
	// allow
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The policy type.
	//
	// example:
	//
	// domain
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s DescribePolicyGroupsResponseBodyDescribePolicyGroupsNetRedirectRule) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyGroupsResponseBodyDescribePolicyGroupsNetRedirectRule) GoString() string {
	return s.String()
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsNetRedirectRule) GetDomain() *string {
	return s.Domain
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsNetRedirectRule) GetPolicy() *string {
	return s.Policy
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsNetRedirectRule) GetRuleType() *string {
	return s.RuleType
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsNetRedirectRule) SetDomain(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsNetRedirectRule {
	s.Domain = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsNetRedirectRule) SetPolicy(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsNetRedirectRule {
	s.Policy = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsNetRedirectRule) SetRuleType(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsNetRedirectRule {
	s.RuleType = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsNetRedirectRule) Validate() error {
	return dara.Validate(s)
}

type DescribePolicyGroupsResponseBodyDescribePolicyGroupsRecordEventLevels struct {
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

func (s DescribePolicyGroupsResponseBodyDescribePolicyGroupsRecordEventLevels) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyGroupsResponseBodyDescribePolicyGroupsRecordEventLevels) GoString() string {
	return s.String()
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsRecordEventLevels) GetEventLevel() *string {
	return s.EventLevel
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsRecordEventLevels) GetEventType() *string {
	return s.EventType
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsRecordEventLevels) SetEventLevel(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsRecordEventLevels {
	s.EventLevel = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsRecordEventLevels) SetEventType(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsRecordEventLevels {
	s.EventType = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsRecordEventLevels) Validate() error {
	return dara.Validate(s)
}

type DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule struct {
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

func (s DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) GoString() string {
	return s.String()
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) GetDescription() *string {
	return s.Description
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) GetDeviceClass() *string {
	return s.DeviceClass
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) GetDeviceSubclass() *string {
	return s.DeviceSubclass
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) GetProductId() *string {
	return s.ProductId
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) GetUsbRedirectType() *int64 {
	return s.UsbRedirectType
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) GetUsbRuleType() *int64 {
	return s.UsbRuleType
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) GetVendorId() *string {
	return s.VendorId
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) SetDescription(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule {
	s.Description = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) SetDeviceClass(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule {
	s.DeviceClass = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) SetDeviceSubclass(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule {
	s.DeviceSubclass = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) SetProductId(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule {
	s.ProductId = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) SetUsbRedirectType(v int64) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule {
	s.UsbRedirectType = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) SetUsbRuleType(v int64) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule {
	s.UsbRuleType = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) SetVendorId(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule {
	s.VendorId = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) Validate() error {
	return dara.Validate(s)
}
