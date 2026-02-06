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
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The details of the cloud computer policies.
	DescribePolicyGroups []*DescribePolicyGroupsResponseBodyDescribePolicyGroups `json:"DescribePolicyGroups,omitempty" xml:"DescribePolicyGroups,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	AcademicProxy *string `json:"AcademicProxy,omitempty" xml:"AcademicProxy,omitempty"`
	// Indicates whether end users are granted the administrator permissions.
	//
	// >  This parameter is in invitational preview for specific users and not available to the public.
	//
	// example:
	//
	// deny
	AdminAccess *string `json:"AdminAccess,omitempty" xml:"AdminAccess,omitempty"`
	// Indicates whether the anti-screenshot feature is enabled.
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
	AppContentProtection *string `json:"AppContentProtection,omitempty" xml:"AppContentProtection,omitempty"`
	// The client IP address whitelist. End users can access cloud computers only from the IP addresses in the whitelist.
	AuthorizeAccessPolicyRules []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules `json:"AuthorizeAccessPolicyRules,omitempty" xml:"AuthorizeAccessPolicyRules,omitempty" type:"Repeated"`
	// The security group rules.
	AuthorizeSecurityPolicyRules []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules `json:"AuthorizeSecurityPolicyRules,omitempty" xml:"AuthorizeSecurityPolicyRules,omitempty" type:"Repeated"`
	// The automatic client connection recovery configurations.
	//
	// example:
	//
	// off
	AutoReconnect *string `json:"AutoReconnect,omitempty" xml:"AutoReconnect,omitempty"`
	// Indicates whether the webcam redirection feature is enabled.
	//
	// Valid values:
	//
	// 	- off
	//
	// 	- on (default)
	//
	// example:
	//
	// on
	CameraRedirect       *string `json:"CameraRedirect,omitempty" xml:"CameraRedirect,omitempty"`
	ClientControlMenu    *string `json:"ClientControlMenu,omitempty" xml:"ClientControlMenu,omitempty"`
	ClientCreateSnapshot *string `json:"ClientCreateSnapshot,omitempty" xml:"ClientCreateSnapshot,omitempty"`
	ClientHibernate      *string `json:"ClientHibernate,omitempty" xml:"ClientHibernate,omitempty"`
	ClientRestart        *string `json:"ClientRestart,omitempty" xml:"ClientRestart,omitempty"`
	ClientShutdown       *string `json:"ClientShutdown,omitempty" xml:"ClientShutdown,omitempty"`
	// The logon method control rules to limit the type of the Alibaba Cloud Workspace client used by end users to connect to cloud computers.
	ClientTypes []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsClientTypes `json:"ClientTypes,omitempty" xml:"ClientTypes,omitempty" type:"Repeated"`
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
	// 	- off: disables both one-way and two-way transfer. Files cannot be copied between local devices and cloud computers.
	//
	// example:
	//
	// readwrite
	Clipboard *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	// Indicates whether the Color Enhancement switch is turned on in design and 3D scenarios.
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
	ColorEnhancement  *string `json:"ColorEnhancement,omitempty" xml:"ColorEnhancement,omitempty"`
	CpdDriveClipboard *string `json:"CpdDriveClipboard,omitempty" xml:"CpdDriveClipboard,omitempty"`
	// The CPU underclocking duration. Valid values: 30 to 120. Unit: seconds.
	//
	// example:
	//
	// 30
	CpuDownGradeDuration *int32  `json:"CpuDownGradeDuration,omitempty" xml:"CpuDownGradeDuration,omitempty"`
	CpuOverload          *string `json:"CpuOverload,omitempty" xml:"CpuOverload,omitempty"`
	// The process whitelist that is not restricted by the CPU usage limit.
	CpuProcessors []*string `json:"CpuProcessors,omitempty" xml:"CpuProcessors,omitempty" type:"Repeated"`
	// Indicates whether the CPU spike protection switch is turned on.
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
	// The number of cloud computers bound with this policy.
	//
	// example:
	//
	// 1
	DesktopCount *int32 `json:"DesktopCount,omitempty" xml:"DesktopCount,omitempty"`
	// The number of shared cloud computers bound with this policy.
	//
	// example:
	//
	// 1
	DesktopGroupCount *int32  `json:"DesktopGroupCount,omitempty" xml:"DesktopGroupCount,omitempty"`
	DeviceConnectHint *string `json:"DeviceConnectHint,omitempty" xml:"DeviceConnectHint,omitempty"`
	// The device redirection rules.
	DeviceRedirects []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRedirects `json:"DeviceRedirects,omitempty" xml:"DeviceRedirects,omitempty" type:"Repeated"`
	// The custom peripheral rules.
	DeviceRules  []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsDeviceRules `json:"DeviceRules,omitempty" xml:"DeviceRules,omitempty" type:"Repeated"`
	DiskOverload *string                                                            `json:"DiskOverload,omitempty" xml:"DiskOverload,omitempty"`
	// The display mode.
	//
	// Valid values:
	//
	// 	- clientCustom: suitable for user-defined scenarios.
	//
	// 	- adminOffice: suitable for daily office scenarios.
	//
	// 	- adminDesign: suitable for 3D application scenarios.
	//
	// 	- adminCustom: administrator-customized scenarios
	//
	// example:
	//
	// adminCustom
	DisplayMode *string `json:"DisplayMode,omitempty" xml:"DisplayMode,omitempty"`
	// Specifies whether to enable the access control for domain names. Domain names support wildcards (\\*). Separate multiple domain names with commas (,).
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
	// The domain name resolution rules.
	DomainResolveRule []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsDomainResolveRule `json:"DomainResolveRule,omitempty" xml:"DomainResolveRule,omitempty" type:"Repeated"`
	// Indicates whether the switch for domain name resolution is turned on.
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
	DomainResolveRuleType *string `json:"DomainResolveRuleType,omitempty" xml:"DomainResolveRuleType,omitempty"`
	// The number of cloud computers that are associated with the policy. The number of cloud computers that are associated only with custom policies is returned.
	//
	// example:
	//
	// 1
	EdsCount *int32 `json:"EdsCount,omitempty" xml:"EdsCount,omitempty"`
	// Indicates whether the Contact Administrator for Help switch is turned on.
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
	EndUserApplyAdminCoordinate *string `json:"EndUserApplyAdminCoordinate,omitempty" xml:"EndUserApplyAdminCoordinate,omitempty"`
	// Indicates whether the User Stream Collaboration switch is turned on.
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
	EndUserGroupCoordinate *string `json:"EndUserGroupCoordinate,omitempty" xml:"EndUserGroupCoordinate,omitempty"`
	ExternalDrive          *string `json:"ExternalDrive,omitempty" xml:"ExternalDrive,omitempty"`
	FileMigrate            *string `json:"FileMigrate,omitempty" xml:"FileMigrate,omitempty"`
	// Transfers files.
	//
	// example:
	//
	// null
	FileTransfer              *string `json:"FileTransfer,omitempty" xml:"FileTransfer,omitempty"`
	FileTransferAddress       *string `json:"FileTransferAddress,omitempty" xml:"FileTransferAddress,omitempty"`
	FileTransferSpeed         *string `json:"FileTransferSpeed,omitempty" xml:"FileTransferSpeed,omitempty"`
	FileTransferSpeedLocation *string `json:"FileTransferSpeedLocation,omitempty" xml:"FileTransferSpeedLocation,omitempty"`
	// Indicates whether the Image Quality Control feature is enabled. If you have high requirements on the performance and user experience in scenarios such as professional design, we recommend that you enable this feature.
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
	HoverConfigMsg  *string `json:"HoverConfigMsg,omitempty" xml:"HoverConfigMsg,omitempty"`
	HoverHibernate  *string `json:"HoverHibernate,omitempty" xml:"HoverHibernate,omitempty"`
	HoverRestart    *string `json:"HoverRestart,omitempty" xml:"HoverRestart,omitempty"`
	HoverShutdown   *string `json:"HoverShutdown,omitempty" xml:"HoverShutdown,omitempty"`
	// Specifies whether to allow web client access.
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
	// 	- TCP (default): TCP.
	//
	// 	- BOTH: TCP and UDP.
	//
	// example:
	//
	// BOTH
	InternetCommunicationProtocol *string `json:"InternetCommunicationProtocol,omitempty" xml:"InternetCommunicationProtocol,omitempty"`
	InternetPrinter               *string `json:"InternetPrinter,omitempty" xml:"InternetPrinter,omitempty"`
	// The permissions on local disk mapping.
	//
	// Valid values:
	//
	// 	- read: read-only. Local disk mapping is available on cloud computers. However, you can only read (copy) local files but cannot modify the files.
	//
	// 	- readwrite: read and write. Local disk mapping is available on cloud computers. You can read (copy) and write (modify) local files.
	//
	// 	- off (default): none.
	//
	// example:
	//
	// readwrite
	LocalDrive *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	// The maximum retry period for reconnecting to cloud computers when the cloud computers are disconnected due to none-human reasons. Valid values: 30 to 7200. Unit: seconds.
	//
	// example:
	//
	// 120
	MaxReconnectTime *int32 `json:"MaxReconnectTime,omitempty" xml:"MaxReconnectTime,omitempty"`
	// The memory underclocking duration for a single process. Valid values: 30 to 120. Unit: seconds.
	//
	// example:
	//
	// 30
	MemoryDownGradeDuration *int32  `json:"MemoryDownGradeDuration,omitempty" xml:"MemoryDownGradeDuration,omitempty"`
	MemoryOverload          *string `json:"MemoryOverload,omitempty" xml:"MemoryOverload,omitempty"`
	// The whitelist of processes that are not restricted by the memory usage limit.
	MemoryProcessors []*string `json:"MemoryProcessors,omitempty" xml:"MemoryProcessors,omitempty" type:"Repeated"`
	// Indicates whether the memory spike protection switch is turned on.
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
	// The memory usage of a single process. Valid values: 30 to 60. Unit: %.
	//
	// example:
	//
	// 30
	MemorySingleRateLimit *int32 `json:"MemorySingleRateLimit,omitempty" xml:"MemorySingleRateLimit,omitempty"`
	// Specifies whether to display the restart button in the DesktopAssistant when the cloud computer is accessed from the Alibaba Cloud Workspace mobile clients (including the Android client and the iOS client).
	//
	// > Mobile clients of V7.4 and higher versions required.
	//
	// Valid values:
	//
	// - off: not provided.
	//
	// - on: provided.
	//
	// example:
	//
	// off
	MobileRestart *string `json:"MobileRestart,omitempty" xml:"MobileRestart,omitempty"`
	// Indicates whether the Windows security control is enabled for mobile clients.
	//
	// example:
	//
	// off
	MobileSafeMenu *string `json:"MobileSafeMenu,omitempty" xml:"MobileSafeMenu,omitempty"`
	// Specifies whether to display the shut down button in the DesktopAssistant when the cloud computer is accessed from the Alibaba Cloud Workspace mobile clients (including the Android client and the iOS client).
	//
	// > Mobile clients of V7.4 and higher versions required.
	//
	// Valid values:
	//
	// - off: not provided.
	//
	// - on: provided.
	//
	// example:
	//
	// off
	MobileShutdown *string `json:"MobileShutdown,omitempty" xml:"MobileShutdown,omitempty"`
	// Indicates whether the Cloud Computer Manager is enabled for mobile clients.
	//
	// example:
	//
	// off
	MobileWuyingKeeper *string `json:"MobileWuyingKeeper,omitempty" xml:"MobileWuyingKeeper,omitempty"`
	// Indicates whether the Xiaoying AI Assistant is enabled for mobile clients.
	//
	// example:
	//
	// off
	MobileWyAssistant *string `json:"MobileWyAssistant,omitempty" xml:"MobileWyAssistant,omitempty"`
	ModelLibrary      *string `json:"ModelLibrary,omitempty" xml:"ModelLibrary,omitempty"`
	MultiScreen       *string `json:"MultiScreen,omitempty" xml:"MultiScreen,omitempty"`
	// The name of the cloud computer policy.
	//
	// example:
	//
	// testPolicyGroupName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the network redirection feature is enabled.
	//
	// >  This parameter is in invitational preview for specific users and not available to the public.
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
	NetRedirect *string `json:"NetRedirect,omitempty" xml:"NetRedirect,omitempty"`
	// The network redirection rule.
	//
	// >  This parameter is in invitational preview for specific users and not available to the public.
	NetRedirectRule []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsNetRedirectRule `json:"NetRedirectRule,omitempty" xml:"NetRedirectRule,omitempty" type:"Repeated"`
	// The ID of the cloud computer policy.
	//
	// example:
	//
	// pg-gx2x1dhsmthe9****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The type of the cloud computer policy.
	//
	// Valid values:
	//
	// 	- SYSTEM
	//
	// 	- CUSTOM
	//
	// example:
	//
	// SYSTEM
	PolicyGroupType *string `json:"PolicyGroupType,omitempty" xml:"PolicyGroupType,omitempty"`
	// The status of the cloud computer policy.
	//
	// Valid values:
	//
	// 	- AVAILABLE
	//
	// 	- CREATING
	//
	// example:
	//
	// AVAILABLE
	PolicyStatus *string `json:"PolicyStatus,omitempty" xml:"PolicyStatus,omitempty"`
	PortProxy    *string `json:"PortProxy,omitempty" xml:"PortProxy,omitempty"`
	// The cloud computer preemption feature.
	//
	// >  To ensure user experience and data security, when a cloud computer is used by an end user, other end users cannot connect to the cloud computer. By default, this parameter is set to `off`, which cannot be modified.
	//
	// Valid values:
	//
	// 	- off: Preemption is not allowed.
	//
	// example:
	//
	// off
	PreemptLogin *string `json:"PreemptLogin,omitempty" xml:"PreemptLogin,omitempty"`
	// The usernames that can preempt to connect to the cloud computer.
	PreemptLoginUsers []*string `json:"PreemptLoginUsers,omitempty" xml:"PreemptLoginUsers,omitempty" type:"Repeated"`
	// Indicates whether the printer redirection feature is enabled.
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
	PrinterRedirection *string `json:"PrinterRedirection,omitempty" xml:"PrinterRedirection,omitempty"`
	// Indicates whether the Image Quality Enhancement switch is turned on for design and 3D scenarios.
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
	QualityEnhancement *string `json:"QualityEnhancement,omitempty" xml:"QualityEnhancement,omitempty"`
	// Indicates whether the custom screen recording feature is enabled.
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
	RecordContent *string `json:"RecordContent,omitempty" xml:"RecordContent,omitempty"`
	// The period when the custom screen recording can be retained before expiration. Default value: 30 days.
	//
	// example:
	//
	// 30
	RecordContentExpires *int64 `json:"RecordContentExpires,omitempty" xml:"RecordContentExpires,omitempty"`
	// The recording duration since a target event is detected by the screen recording audit policy. Unit: Minute. Valid values: 10-60.
	//
	// example:
	//
	// 10
	RecordEventDuration *int32 `json:"RecordEventDuration,omitempty" xml:"RecordEventDuration,omitempty"`
	// The screen recording file suffix.
	RecordEventFileExts []*string `json:"RecordEventFileExts,omitempty" xml:"RecordEventFileExts,omitempty" type:"Repeated"`
	// The array of absolute paths of the monitored files in the screen recording audit policy.
	RecordEventFilePaths []*string `json:"RecordEventFilePaths,omitempty" xml:"RecordEventFilePaths,omitempty" type:"Repeated"`
	// Indicates whether the screen recording event severity is enabled.
	RecordEventLevels []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsRecordEventLevels `json:"RecordEventLevels,omitempty" xml:"RecordEventLevels,omitempty" type:"Repeated"`
	// The array of absolute paths of the monitored registry entries in the screen recording audit policy.
	RecordEventRegisters []*string `json:"RecordEventRegisters,omitempty" xml:"RecordEventRegisters,omitempty" type:"Repeated"`
	// Indicates whether the screen recording feature is enabled.
	//
	// Valid values:
	//
	// 	- byaction_cmd_ft: enables the operation-triggered screen recording upon command execution and file transfer.
	//
	// 	- ALLTIME: enables the whole-process screen recording. That is, the recording starts when cloud computers are connected and ends when the cloud computers are disconnected.
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
	// Indicates whether audio files generated from cloud computers are recorded.
	//
	// Valid values:
	//
	// 	- off (default): records only video files.
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
	// 10
	RecordingDuration *int32 `json:"RecordingDuration,omitempty" xml:"RecordingDuration,omitempty"`
	// The time when the screen recording ended. The value is in the HH:MM:SS format. The value takes effect only when Recording is set to PERIOD.
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
	// 5
	RecordingFps *int64 `json:"RecordingFps,omitempty" xml:"RecordingFps,omitempty"`
	// The time when the screen recording was started. The value is in the HH:MM:SS format. The value takes effect only when Recording is set to PERIOD.
	//
	// example:
	//
	// 08:00:00
	RecordingStartTime *string `json:"RecordingStartTime,omitempty" xml:"RecordingStartTime,omitempty"`
	// Indicates whether the screen recording notification feature is enabled after end users log on to the Alibaba Cloud Workspace client.
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
	// Your desktop is being recorded
	RecordingUserNotifyMessage *string `json:"RecordingUserNotifyMessage,omitempty" xml:"RecordingUserNotifyMessage,omitempty"`
	// The permissions on keyboard and mouse control during remote assistance.
	//
	// Valid values:
	//
	// 	- optionalControl: By default, you are not granted the permissions. You can apply for the permissions.
	//
	// 	- fullControl: You are granted the full permissions.
	//
	// 	- disableControl: You are not granted the permissions.
	//
	// example:
	//
	// fullControl
	RemoteCoordinate *string `json:"RemoteCoordinate,omitempty" xml:"RemoteCoordinate,omitempty"`
	// Resets the cloud computer.
	//
	// example:
	//
	// null
	ResetDesktop     *string `json:"ResetDesktop,omitempty" xml:"ResetDesktop,omitempty"`
	ResolutionDpi    *int32  `json:"ResolutionDpi,omitempty" xml:"ResolutionDpi,omitempty"`
	ResolutionHeight *int32  `json:"ResolutionHeight,omitempty" xml:"ResolutionHeight,omitempty"`
	ResolutionModel  *string `json:"ResolutionModel,omitempty" xml:"ResolutionModel,omitempty"`
	ResolutionWidth  *int32  `json:"ResolutionWidth,omitempty" xml:"ResolutionWidth,omitempty"`
	// The number of resource groups bound with this policy.
	//
	// example:
	//
	// 1
	ResourceGroupCount *int32 `json:"ResourceGroupCount,omitempty" xml:"ResourceGroupCount,omitempty"`
	// The region of the cloud computer policy.
	//
	// > The value of a region-less policy is `center`.
	//
	// example:
	//
	// center
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	SafeMenu         *string `json:"SafeMenu,omitempty" xml:"SafeMenu,omitempty"`
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
	ScopeValue        []*string `json:"ScopeValue,omitempty" xml:"ScopeValue,omitempty" type:"Repeated"`
	ScreenDisplayMode *string   `json:"ScreenDisplayMode,omitempty" xml:"ScreenDisplayMode,omitempty"`
	// Indicates whether the Smooth Enhancement switch is turned on.
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
	SmoothEnhancement *string `json:"SmoothEnhancement,omitempty" xml:"SmoothEnhancement,omitempty"`
	// Specifies whether to provide the Metrics function in the DesktopAssistant. Valid values:
	//
	// - off: not provided.
	//
	// - on: provided.
	//
	// example:
	//
	// on
	StatusMonitor *string `json:"StatusMonitor,omitempty" xml:"StatusMonitor,omitempty"`
	// The streaming mode.
	//
	// Valid values:
	//
	// 	- intelligent: suitable for daily office scenarios (Intelligent Mode).
	//
	// 	- smooth: suitable for design and 3D application scenarios (Smooth Mode).
	//
	// example:
	//
	// smooth
	StreamingMode *string `json:"StreamingMode,omitempty" xml:"StreamingMode,omitempty"`
	// The destination frame rate. Valid values: 10 to 60. Unit: fps.
	//
	// example:
	//
	// 30
	TargetFps *int32 `json:"TargetFps,omitempty" xml:"TargetFps,omitempty"`
	// Indicates whether the USB redirection feature is enabled.
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
	UsbRedirect *string `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty"`
	// The USB redirection rule.
	UsbSupplyRedirectRule []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule `json:"UsbSupplyRedirectRule,omitempty" xml:"UsbSupplyRedirectRule,omitempty" type:"Repeated"`
	UseTime               *string                                                                      `json:"UseTime,omitempty" xml:"UseTime,omitempty"`
	// The average bitrate for video encoding. Valid values: 1000 to 50000.
	//
	// example:
	//
	// 1000
	VideoEncAvgKbps *int32 `json:"VideoEncAvgKbps,omitempty" xml:"VideoEncAvgKbps,omitempty"`
	// The maximum quantizer parameter (QP) of video files. A larger QP value indicates worse video quality. Valid values: 0 to 51.
	//
	// example:
	//
	// 20
	VideoEncMaxQP *int32 `json:"VideoEncMaxQP,omitempty" xml:"VideoEncMaxQP,omitempty"`
	// The minimum quantizer parameter (QP) of video files. A smaller QP value indicates higher video quality. Valid values: 0 to 51.
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
	// The video encoding feature.
	//
	// Valid values:
	//
	// 	- qualityFirst: The priority given to the image quality.
	//
	// 	- bandwidthFirst: The priority given to the bitrate.
	//
	// example:
	//
	// qualityFirst
	VideoEncPolicy *string `json:"VideoEncPolicy,omitempty" xml:"VideoEncPolicy,omitempty"`
	// Indicates whether the multimedia redirection feature is enabled.
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
	VideoRedirect *string `json:"VideoRedirect,omitempty" xml:"VideoRedirect,omitempty"`
	// The image display quality.
	//
	// Valid values:
	//
	// 	- high: high-definition (HD)
	//
	// 	- low: fluent
	//
	// 	- medium (default): adaptive
	//
	// 	- lossless: no quality loss
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
	// 	- off: The watermarking feature is disabled.
	//
	// 	- on: Visible watermarks are applied.
	//
	// example:
	//
	// on
	Watermark *string `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
	// Indicates whether the anti-screen photo feature is enabled for invisible watermarks.
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
	// If you set `WatermarkType` to `custom`, you must also specify `WatermarkCustomText`.
	//
	// example:
	//
	// test
	WatermarkCustomText *string `json:"WatermarkCustomText,omitempty" xml:"WatermarkCustomText,omitempty"`
	// The slope of the watermark. Valid values: -10 to -30.
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
	// Indicates whether the security priority feature is enabled for invisible watermarks.
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
	WatermarkShadow   *string `json:"WatermarkShadow,omitempty" xml:"WatermarkShadow,omitempty"`
	// The watermark transparency.
	//
	// Valid values:
	//
	// 	- LIGHT
	//
	// 	- DARK
	//
	// 	- MIDDLE
	//
	// example:
	//
	// LIGHT
	WatermarkTransparency *string `json:"WatermarkTransparency,omitempty" xml:"WatermarkTransparency,omitempty"`
	// The watermark transparency. A greater value indicates that the watermark is less transparent. Valid values: 10 to 100.
	//
	// example:
	//
	// 10
	WatermarkTransparencyValue *int32 `json:"WatermarkTransparencyValue,omitempty" xml:"WatermarkTransparencyValue,omitempty"`
	// The watermark content.
	//
	// Valid values:
	//
	// 	- EndUserId: the username.
	//
	// 	- Custom
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
	WuyingKeeper  *string `json:"WuyingKeeper,omitempty" xml:"WuyingKeeper,omitempty"`
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
	// The CIDR block that is allowed to access the client. The value is an IPv4 CIDR block.
	//
	// example:
	//
	// 47.100.XX.XX/16
	CidrIp *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	// The remarks on the CIDR block that is allowed to access the client.
	//
	// example:
	//
	// test
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
	// 	- tcp: Transmission Control Protocol (TCP)
	//
	// 	- udp: User Datagram Protocol (UDP)
	//
	// 	- all: all protocols
	//
	// 	- gre: Generic Routing Encapsulation (GRE)
	//
	// 	- icmp: Internet Control Message Protocol (ICMP) for IPv4
	//
	// example:
	//
	// tcp
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The authorization of the security group rule.
	//
	// Valid values:
	//
	// 	- drop: denies all access requests.
	//
	// 	- accept: accepts all requests.
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
	// Valid values:
	//
	// 	- outflow: outbound
	//
	// 	- inflow: inbound
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
	// Valid values:
	//
	// 	- html5: web client
	//
	// 	- android: Android client
	//
	// 	- windows: Windows client
	//
	// 	- ios: iOS client
	//
	// 	- macos: macOS client
	//
	// example:
	//
	// windows
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// Indicates whether end users are allowed to use a specific type of the client to connect to cloud computers.
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
	// The redirection type. Valid values:
	//
	// 	- usbRedirect
	//
	// 	- deviceRedirect
	//
	// 	- off: direction disabled.
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
	// 	- usbKey
	//
	// 	- other
	//
	// 	- graphicsTablet
	//
	// 	- printer
	//
	// 	- cardReader
	//
	// 	- scanner
	//
	// 	- storage
	//
	// 	- camera
	//
	// 	- adb
	//
	// 	- networkInterfaceCard: the NIC device.
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
	// 	- deviceRedirect
	//
	// 	- usbRedirect
	//
	// 	- off: redirection disabled.
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
	// The rule description.
	//
	// example:
	//
	// Test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination domain name.
	//
	// example:
	//
	// *.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Indicates whether the domain name resolution rule is allowed.
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
	// The rule content.
	//
	// example:
	//
	// *.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Indicates whether the rule is allowed.
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
	// The rule type.
	//
	// Valid values:
	//
	// 	- prc: process
	//
	// 	- domain: domain name
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
	// The event severity.
	EventLevel *string `json:"EventLevel,omitempty" xml:"EventLevel,omitempty"`
	// The event type.
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
	// Test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The device class. This parameter is required when `usbRuleType` is set to 1. For more information, see [Defined Class Codes](https://www.usb.org/defined-class-codes).
	//
	// example:
	//
	// 0Eh
	DeviceClass *string `json:"DeviceClass,omitempty" xml:"DeviceClass,omitempty"`
	// The subclass of the device. This parameter is required when `usbRuleType` is set to 1. For more information, see [Defined Class Codes](https://www.usb.org/defined-class-codes).
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
	// Indicates whether USB redirection is allowed.
	//
	// Valid values:
	//
	// 	- 1: allowed
	//
	// 	- 2: not allowed
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
	// 	- 2: by device vendor
	//
	// example:
	//
	// 1
	UsbRuleType *int64 `json:"UsbRuleType,omitempty" xml:"UsbRuleType,omitempty"`
	// The vendor ID (VID). For more information, see [Valid USB VIDs](https://www.usb.org/sites/default/files/vendor_ids032322.pdf_1.pdf).
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
