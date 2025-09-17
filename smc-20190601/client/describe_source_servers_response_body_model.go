// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSourceServersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeSourceServersResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSourceServersResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeSourceServersResponseBody
	GetRequestId() *string
	SetSourceServers(v *DescribeSourceServersResponseBodySourceServers) *DescribeSourceServersResponseBody
	GetSourceServers() *DescribeSourceServersResponseBodySourceServers
	SetTotalCount(v int32) *DescribeSourceServersResponseBody
	GetTotalCount() *int32
}

type DescribeSourceServersResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 410E6073-66D0-45D3-AB3E-4DC3F5E4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the migration source.
	SourceServers *DescribeSourceServersResponseBodySourceServers `json:"SourceServers,omitempty" xml:"SourceServers,omitempty" type:"Struct"`
	// The total number of migration sources returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSourceServersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSourceServersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSourceServersResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSourceServersResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSourceServersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSourceServersResponseBody) GetSourceServers() *DescribeSourceServersResponseBodySourceServers {
	return s.SourceServers
}

func (s *DescribeSourceServersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSourceServersResponseBody) SetPageNumber(v int32) *DescribeSourceServersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSourceServersResponseBody) SetPageSize(v int32) *DescribeSourceServersResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSourceServersResponseBody) SetRequestId(v string) *DescribeSourceServersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSourceServersResponseBody) SetSourceServers(v *DescribeSourceServersResponseBodySourceServers) *DescribeSourceServersResponseBody {
	s.SourceServers = v
	return s
}

func (s *DescribeSourceServersResponseBody) SetTotalCount(v int32) *DescribeSourceServersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSourceServersResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSourceServersResponseBodySourceServers struct {
	SourceServer []*DescribeSourceServersResponseBodySourceServersSourceServer `json:"SourceServer,omitempty" xml:"SourceServer,omitempty" type:"Repeated"`
}

func (s DescribeSourceServersResponseBodySourceServers) String() string {
	return dara.Prettify(s)
}

func (s DescribeSourceServersResponseBodySourceServers) GoString() string {
	return s.String()
}

func (s *DescribeSourceServersResponseBodySourceServers) GetSourceServer() []*DescribeSourceServersResponseBodySourceServersSourceServer {
	return s.SourceServer
}

func (s *DescribeSourceServersResponseBodySourceServers) SetSourceServer(v []*DescribeSourceServersResponseBodySourceServersSourceServer) *DescribeSourceServersResponseBodySourceServers {
	s.SourceServer = v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServers) Validate() error {
	return dara.Validate(s)
}

type DescribeSourceServersResponseBodySourceServersSourceServer struct {
	// The version number of the SMC client.
	//
	// example:
	//
	// 1.5.2.3
	AgentVersion *string `json:"AgentVersion,omitempty" xml:"AgentVersion,omitempty"`
	// The system architecture of the migration source.
	//
	// example:
	//
	// x86_64
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// The time when the migration source was created.
	//
	// example:
	//
	// 2019-06-27T02:58:09Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The data disks on the migration source.
	DataDisks *DescribeSourceServersResponseBodySourceServersSourceServerDataDisks `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Struct"`
	// The description of the migration source.
	//
	// example:
	//
	// Server Source Imported By GotoAliyun.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The information about the disk.
	//
	// if can be null:
	// false
	Disks *DescribeSourceServersResponseBodySourceServersSourceServerDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Struct"`
	// The error code of the migration source.
	//
	// example:
	//
	// SourceServer.Offline
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The interval at which heartbeats are sent from the SMC client. Unit: seconds.
	//
	// example:
	//
	// 30
	HeartbeatRate *int32 `json:"HeartbeatRate,omitempty" xml:"HeartbeatRate,omitempty"`
	// The ID of the last migration job.
	//
	// example:
	//
	// j-bp19vlwm0tyigbmj****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The kernel level of the migration source.
	//
	// example:
	//
	// 1
	KernelLevel *int32 `json:"KernelLevel,omitempty" xml:"KernelLevel,omitempty"`
	// The name of the migration source.
	//
	// example:
	//
	// SourceServerName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The operating system of the migration source.
	//
	// example:
	//
	// OpenSUSE
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The replication driver used for migration. Default value: SMT.
	//
	// example:
	//
	// SMT
	ReplicationDriver *string `json:"ReplicationDriver,omitempty" xml:"ReplicationDriver,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmw3ty5y7****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the migration source.
	//
	// example:
	//
	// s-bp1e2fsl57knvuug****
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The state of the migration source.
	//
	// example:
	//
	// InUse
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The status information of the migration source. This parameter is returned if the migration source is in the Unavailable state. The value of this parameter consists of key-value pairs in the JSON format. Sample keys:
	//
	//     error_code: The error code.
	//
	//     error_msg: the error message.
	//
	// example:
	//
	// {"error_code": "S1", "error_msg": "Rsync not found. Please install rsync."}
	StatusInfo *string `json:"StatusInfo,omitempty" xml:"StatusInfo,omitempty"`
	// The information about the system disk partition.
	SystemDiskParts *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskParts `json:"SystemDiskParts,omitempty" xml:"SystemDiskParts,omitempty" type:"Struct"`
	// The system disk size of the migration source. Unit: GiB
	//
	// example:
	//
	// 40
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// The system information of the migration source. The parameter must be specified as key-value pairs in the JSON format. The key-value pairs are extensible and have fixed keys. Maximum value: 1 KB. Example:
	//
	//     agent_mode: the running mode.
	//
	//     agent_type: the type of the run.
	//
	//     client_type: the type of the client.
	//
	//     hostname : the hostname.
	//
	//     ipv4:IPv4 address
	//
	//     ipv6: IPv6 address
	//
	//     cores: the number of CPU cores.
	//
	//     cpu_usage: the CPU utilization.
	//
	//     memory: the memory size.
	//
	//     memory_usage: the memory usage.
	//
	// example:
	//
	// {\\"agent_mode\\":\\"daemon\\",\\"agent_type\\":\\"aliyun\\",\\"client_type\\":\\"\\",\\"cores\\":\\"2\\",\\"cpu_usage\\":\\"0.00\\",\\"hostname\\":\\"ixxxxxxxxxx\\",\\"ipv4\\":\\"10.0.0.1\\",\\"memory\\":\\"8.00\\",\\"memory_usage\\":\\"3.61\\"}
	SystemInfo *string `json:"SystemInfo,omitempty" xml:"SystemInfo,omitempty"`
	// The tag details.
	Tags *DescribeSourceServersResponseBodySourceServersSourceServerTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The workgroup ID.
	//
	// example:
	//
	// w-bp1ja22kdqphehlj****
	WorkgroupId *string `json:"WorkgroupId,omitempty" xml:"WorkgroupId,omitempty"`
}

func (s DescribeSourceServersResponseBodySourceServersSourceServer) String() string {
	return dara.Prettify(s)
}

func (s DescribeSourceServersResponseBodySourceServersSourceServer) GoString() string {
	return s.String()
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) GetArchitecture() *string {
	return s.Architecture
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) GetDataDisks() *DescribeSourceServersResponseBodySourceServersSourceServerDataDisks {
	return s.DataDisks
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) GetDescription() *string {
	return s.Description
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) GetDisks() *DescribeSourceServersResponseBodySourceServersSourceServerDisks {
	return s.Disks
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) GetHeartbeatRate() *int32 {
	return s.HeartbeatRate
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) GetJobId() *string {
	return s.JobId
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) GetKernelLevel() *int32 {
	return s.KernelLevel
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) GetName() *string {
	return s.Name
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) GetPlatform() *string {
	return s.Platform
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) GetReplicationDriver() *string {
	return s.ReplicationDriver
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) GetSourceId() *string {
	return s.SourceId
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) GetState() *string {
	return s.State
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) GetStatusInfo() *string {
	return s.StatusInfo
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) GetSystemDiskParts() *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskParts {
	return s.SystemDiskParts
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) GetSystemInfo() *string {
	return s.SystemInfo
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) GetTags() *DescribeSourceServersResponseBodySourceServersSourceServerTags {
	return s.Tags
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) GetWorkgroupId() *string {
	return s.WorkgroupId
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetAgentVersion(v string) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.AgentVersion = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetArchitecture(v string) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.Architecture = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetCreationTime(v string) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.CreationTime = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetDataDisks(v *DescribeSourceServersResponseBodySourceServersSourceServerDataDisks) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.DataDisks = v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetDescription(v string) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.Description = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetDisks(v *DescribeSourceServersResponseBodySourceServersSourceServerDisks) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.Disks = v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetErrorCode(v string) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.ErrorCode = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetHeartbeatRate(v int32) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.HeartbeatRate = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetJobId(v string) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.JobId = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetKernelLevel(v int32) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.KernelLevel = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetName(v string) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.Name = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetPlatform(v string) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.Platform = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetReplicationDriver(v string) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.ReplicationDriver = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetResourceGroupId(v string) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetSourceId(v string) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.SourceId = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetState(v string) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.State = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetStatusInfo(v string) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.StatusInfo = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetSystemDiskParts(v *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskParts) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.SystemDiskParts = v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetSystemDiskSize(v int32) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetSystemInfo(v string) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.SystemInfo = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetTags(v *DescribeSourceServersResponseBodySourceServersSourceServerTags) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.Tags = v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetWorkgroupId(v string) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.WorkgroupId = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) Validate() error {
	return dara.Validate(s)
}

type DescribeSourceServersResponseBodySourceServersSourceServerDataDisks struct {
	DataDisk []*DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerDataDisks) String() string {
	return dara.Prettify(s)
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerDataDisks) GoString() string {
	return s.String()
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisks) GetDataDisk() []*DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDisk {
	return s.DataDisk
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisks) SetDataDisk(v []*DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDisk) *DescribeSourceServersResponseBodySourceServersSourceServerDataDisks {
	s.DataDisk = v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisks) Validate() error {
	return dara.Validate(s)
}

type DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDisk struct {
	// The index number of the data disk.
	//
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The information about the data disk partition.
	Parts *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskParts `json:"Parts,omitempty" xml:"Parts,omitempty" type:"Struct"`
	// The path of data disk N.
	//
	// example:
	//
	// /home/data
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The size of data disk N. Unit: GiB.
	//
	// example:
	//
	// 20
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDisk) String() string {
	return dara.Prettify(s)
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDisk) GoString() string {
	return s.String()
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDisk) GetIndex() *int32 {
	return s.Index
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDisk) GetParts() *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskParts {
	return s.Parts
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDisk) GetPath() *string {
	return s.Path
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDisk) GetSize() *int32 {
	return s.Size
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDisk) SetIndex(v int32) *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDisk {
	s.Index = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDisk) SetParts(v *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskParts) *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDisk {
	s.Parts = v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDisk) SetPath(v string) *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDisk {
	s.Path = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDisk) SetSize(v int32) *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDisk {
	s.Size = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDisk) Validate() error {
	return dara.Validate(s)
}

type DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskParts struct {
	Part []*DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart `json:"Part,omitempty" xml:"Part,omitempty" type:"Repeated"`
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskParts) String() string {
	return dara.Prettify(s)
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskParts) GoString() string {
	return s.String()
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskParts) GetPart() []*DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart {
	return s.Part
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskParts) SetPart(v []*DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart) *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskParts {
	s.Part = v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskParts) Validate() error {
	return dara.Validate(s)
}

type DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart struct {
	// Indicates whether block replication is enabled for the data disk partition.
	//
	// example:
	//
	// false
	CanBlock *bool `json:"CanBlock,omitempty" xml:"CanBlock,omitempty"`
	// The device ID of the data disk partition.
	//
	// example:
	//
	// 1_0
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// Indicates whether the data disk partition must be selected.
	//
	// example:
	//
	// false
	Need *bool `json:"Need,omitempty" xml:"Need,omitempty"`
	// The path of the data disk partition.
	//
	// example:
	//
	// /home/data
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The size of the data disk partition. Unit: bytes.
	//
	// example:
	//
	// 21474836480
	SizeBytes *int64 `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart) String() string {
	return dara.Prettify(s)
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart) GoString() string {
	return s.String()
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart) GetCanBlock() *bool {
	return s.CanBlock
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart) GetDevice() *string {
	return s.Device
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart) GetNeed() *bool {
	return s.Need
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart) GetPath() *string {
	return s.Path
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart) GetSizeBytes() *int64 {
	return s.SizeBytes
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart) SetCanBlock(v bool) *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart {
	s.CanBlock = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart) SetDevice(v string) *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart {
	s.Device = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart) SetNeed(v bool) *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart {
	s.Need = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart) SetPath(v string) *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart {
	s.Path = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart) SetSizeBytes(v int64) *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart {
	s.SizeBytes = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart) Validate() error {
	return dara.Validate(s)
}

type DescribeSourceServersResponseBodySourceServersSourceServerDisks struct {
	// The list of data disk information.
	Data *DescribeSourceServersResponseBodySourceServersSourceServerDisksData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The information about the system disk.
	//
	// if can be null:
	// false
	System *DescribeSourceServersResponseBodySourceServersSourceServerDisksSystem `json:"System,omitempty" xml:"System,omitempty" type:"Struct"`
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerDisks) String() string {
	return dara.Prettify(s)
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerDisks) GoString() string {
	return s.String()
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisks) GetData() *DescribeSourceServersResponseBodySourceServersSourceServerDisksData {
	return s.Data
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisks) GetSystem() *DescribeSourceServersResponseBodySourceServersSourceServerDisksSystem {
	return s.System
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisks) SetData(v *DescribeSourceServersResponseBodySourceServersSourceServerDisksData) *DescribeSourceServersResponseBodySourceServersSourceServerDisks {
	s.Data = v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisks) SetSystem(v *DescribeSourceServersResponseBodySourceServersSourceServerDisksSystem) *DescribeSourceServersResponseBodySourceServersSourceServerDisks {
	s.System = v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisks) Validate() error {
	return dara.Validate(s)
}

type DescribeSourceServersResponseBodySourceServersSourceServerDisksData struct {
	Data []*DescribeSourceServersResponseBodySourceServersSourceServerDisksDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerDisksData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerDisksData) GoString() string {
	return s.String()
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksData) GetData() []*DescribeSourceServersResponseBodySourceServersSourceServerDisksDataData {
	return s.Data
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksData) SetData(v []*DescribeSourceServersResponseBodySourceServersSourceServerDisksDataData) *DescribeSourceServersResponseBodySourceServersSourceServerDisksData {
	s.Data = v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksData) Validate() error {
	return dara.Validate(s)
}

type DescribeSourceServersResponseBodySourceServersSourceServerDisksDataData struct {
	// The start offset of the first partition of the data disk. Unit: bytes.
	//
	// example:
	//
	// 1024
	Offset *int64 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The information about the data disk partition.
	Parts *DescribeSourceServersResponseBodySourceServersSourceServerDisksDataDataParts `json:"Parts,omitempty" xml:"Parts,omitempty" type:"Struct"`
	// The data disk size of the migration source. Unit: GiB.
	//
	// example:
	//
	// 80
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerDisksDataData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerDisksDataData) GoString() string {
	return s.String()
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksDataData) GetOffset() *int64 {
	return s.Offset
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksDataData) GetParts() *DescribeSourceServersResponseBodySourceServersSourceServerDisksDataDataParts {
	return s.Parts
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksDataData) GetSize() *int32 {
	return s.Size
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksDataData) SetOffset(v int64) *DescribeSourceServersResponseBodySourceServersSourceServerDisksDataData {
	s.Offset = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksDataData) SetParts(v *DescribeSourceServersResponseBodySourceServersSourceServerDisksDataDataParts) *DescribeSourceServersResponseBodySourceServersSourceServerDisksDataData {
	s.Parts = v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksDataData) SetSize(v int32) *DescribeSourceServersResponseBodySourceServersSourceServerDisksDataData {
	s.Size = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksDataData) Validate() error {
	return dara.Validate(s)
}

type DescribeSourceServersResponseBodySourceServersSourceServerDisksDataDataParts struct {
	Part []*DescribeSourceServersResponseBodySourceServersSourceServerDisksDataDataPartsPart `json:"Part,omitempty" xml:"Part,omitempty" type:"Repeated"`
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerDisksDataDataParts) String() string {
	return dara.Prettify(s)
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerDisksDataDataParts) GoString() string {
	return s.String()
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksDataDataParts) GetPart() []*DescribeSourceServersResponseBodySourceServersSourceServerDisksDataDataPartsPart {
	return s.Part
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksDataDataParts) SetPart(v []*DescribeSourceServersResponseBodySourceServersSourceServerDisksDataDataPartsPart) *DescribeSourceServersResponseBodySourceServersSourceServerDisksDataDataParts {
	s.Part = v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksDataDataParts) Validate() error {
	return dara.Validate(s)
}

type DescribeSourceServersResponseBodySourceServersSourceServerDisksDataDataPartsPart struct {
	// Whether block replication is enabled for the data disk partition. Valid values:
	//
	// 	- true: Block replication is enabled for the data disk partition.
	//
	// 	- false: Block replication is disabled for the data disk partition.
	//
	// example:
	//
	// false
	CanBlock *bool `json:"CanBlock,omitempty" xml:"CanBlock,omitempty"`
	// The path of the data disk partition.
	//
	// example:
	//
	// /home/data
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The size of the data disk partition. Unit: bytes.
	//
	// example:
	//
	// 21474836480
	SizeBytes *int64 `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
	// The type of the data disk partition. Valid values:
	//
	// 	- Normal: normal partition.
	//
	// 	- System: system partition.
	//
	// 	- Boot: boot partition.
	//
	// example:
	//
	// Normal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerDisksDataDataPartsPart) String() string {
	return dara.Prettify(s)
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerDisksDataDataPartsPart) GoString() string {
	return s.String()
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksDataDataPartsPart) GetCanBlock() *bool {
	return s.CanBlock
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksDataDataPartsPart) GetPath() *string {
	return s.Path
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksDataDataPartsPart) GetSizeBytes() *int64 {
	return s.SizeBytes
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksDataDataPartsPart) GetType() *string {
	return s.Type
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksDataDataPartsPart) SetCanBlock(v bool) *DescribeSourceServersResponseBodySourceServersSourceServerDisksDataDataPartsPart {
	s.CanBlock = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksDataDataPartsPart) SetPath(v string) *DescribeSourceServersResponseBodySourceServersSourceServerDisksDataDataPartsPart {
	s.Path = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksDataDataPartsPart) SetSizeBytes(v int64) *DescribeSourceServersResponseBodySourceServersSourceServerDisksDataDataPartsPart {
	s.SizeBytes = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksDataDataPartsPart) SetType(v string) *DescribeSourceServersResponseBodySourceServersSourceServerDisksDataDataPartsPart {
	s.Type = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksDataDataPartsPart) Validate() error {
	return dara.Validate(s)
}

type DescribeSourceServersResponseBodySourceServersSourceServerDisksSystem struct {
	// The start offset of the first partition of the system disk. Unit: bytes.
	//
	// example:
	//
	// 1024
	Offset *int64 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The information about the system disk partition.
	Parts *DescribeSourceServersResponseBodySourceServersSourceServerDisksSystemParts `json:"Parts,omitempty" xml:"Parts,omitempty" type:"Struct"`
	// The size of the source system disk. Unit: GiB. Valid values: 20 to 32768.
	//
	// >  The parameter value must be greater than the actual used space of the data disk on the source server. For example, if the size of the source disk is 500 GiB but the actual used space is 100 GiB, you must set this parameter to a value greater than 100 GiB.
	//
	// example:
	//
	// 100
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerDisksSystem) String() string {
	return dara.Prettify(s)
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerDisksSystem) GoString() string {
	return s.String()
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksSystem) GetOffset() *int64 {
	return s.Offset
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksSystem) GetParts() *DescribeSourceServersResponseBodySourceServersSourceServerDisksSystemParts {
	return s.Parts
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksSystem) GetSize() *int32 {
	return s.Size
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksSystem) SetOffset(v int64) *DescribeSourceServersResponseBodySourceServersSourceServerDisksSystem {
	s.Offset = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksSystem) SetParts(v *DescribeSourceServersResponseBodySourceServersSourceServerDisksSystemParts) *DescribeSourceServersResponseBodySourceServersSourceServerDisksSystem {
	s.Parts = v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksSystem) SetSize(v int32) *DescribeSourceServersResponseBodySourceServersSourceServerDisksSystem {
	s.Size = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksSystem) Validate() error {
	return dara.Validate(s)
}

type DescribeSourceServersResponseBodySourceServersSourceServerDisksSystemParts struct {
	Part []*DescribeSourceServersResponseBodySourceServersSourceServerDisksSystemPartsPart `json:"Part,omitempty" xml:"Part,omitempty" type:"Repeated"`
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerDisksSystemParts) String() string {
	return dara.Prettify(s)
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerDisksSystemParts) GoString() string {
	return s.String()
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksSystemParts) GetPart() []*DescribeSourceServersResponseBodySourceServersSourceServerDisksSystemPartsPart {
	return s.Part
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksSystemParts) SetPart(v []*DescribeSourceServersResponseBodySourceServersSourceServerDisksSystemPartsPart) *DescribeSourceServersResponseBodySourceServersSourceServerDisksSystemParts {
	s.Part = v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksSystemParts) Validate() error {
	return dara.Validate(s)
}

type DescribeSourceServersResponseBodySourceServersSourceServerDisksSystemPartsPart struct {
	// Indicates whether block replication is enabled for the system disk partition. Valid values:
	//
	// 	- true: Block replication is enabled for the system disk partition.
	//
	// 	- false: Block replication is disabled for the system disk partition.
	//
	// example:
	//
	// false
	CanBlock *bool `json:"CanBlock,omitempty" xml:"CanBlock,omitempty"`
	// The path of the system disk partition.
	//
	// example:
	//
	// /home/data
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The size of the system disk partition. Unit: bytes.
	//
	// example:
	//
	// 21474836480
	SizeBytes *int64 `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
	// The type of the system disk partition. Valid values:
	//
	// 	- Normal: normal partition.
	//
	// 	- System: system partition.
	//
	// 	- Boot: boot partition.
	//
	// example:
	//
	// Normal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerDisksSystemPartsPart) String() string {
	return dara.Prettify(s)
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerDisksSystemPartsPart) GoString() string {
	return s.String()
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksSystemPartsPart) GetCanBlock() *bool {
	return s.CanBlock
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksSystemPartsPart) GetPath() *string {
	return s.Path
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksSystemPartsPart) GetSizeBytes() *int64 {
	return s.SizeBytes
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksSystemPartsPart) GetType() *string {
	return s.Type
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksSystemPartsPart) SetCanBlock(v bool) *DescribeSourceServersResponseBodySourceServersSourceServerDisksSystemPartsPart {
	s.CanBlock = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksSystemPartsPart) SetPath(v string) *DescribeSourceServersResponseBodySourceServersSourceServerDisksSystemPartsPart {
	s.Path = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksSystemPartsPart) SetSizeBytes(v int64) *DescribeSourceServersResponseBodySourceServersSourceServerDisksSystemPartsPart {
	s.SizeBytes = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksSystemPartsPart) SetType(v string) *DescribeSourceServersResponseBodySourceServersSourceServerDisksSystemPartsPart {
	s.Type = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDisksSystemPartsPart) Validate() error {
	return dara.Validate(s)
}

type DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskParts struct {
	SystemDiskPart []*DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart `json:"SystemDiskPart,omitempty" xml:"SystemDiskPart,omitempty" type:"Repeated"`
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskParts) String() string {
	return dara.Prettify(s)
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskParts) GoString() string {
	return s.String()
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskParts) GetSystemDiskPart() []*DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart {
	return s.SystemDiskPart
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskParts) SetSystemDiskPart(v []*DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart) *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskParts {
	s.SystemDiskPart = v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskParts) Validate() error {
	return dara.Validate(s)
}

type DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart struct {
	// Indicates whether block replication is enabled for the system disk partition.
	//
	// example:
	//
	// true
	CanBlock *bool `json:"CanBlock,omitempty" xml:"CanBlock,omitempty"`
	// The device ID of the system disk partition.
	//
	// example:
	//
	// 0_0
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// Indicates whether the system disk partition must be selected.
	//
	// example:
	//
	// true
	Need *bool `json:"Need,omitempty" xml:"Need,omitempty"`
	// The path of the system disk partition.
	//
	// example:
	//
	// /boot
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The size of the system disk partition. Unit: bytes.
	//
	// example:
	//
	// 254803968
	SizeBytes *int64 `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart) String() string {
	return dara.Prettify(s)
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart) GoString() string {
	return s.String()
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart) GetCanBlock() *bool {
	return s.CanBlock
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart) GetDevice() *string {
	return s.Device
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart) GetNeed() *bool {
	return s.Need
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart) GetPath() *string {
	return s.Path
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart) GetSizeBytes() *int64 {
	return s.SizeBytes
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart) SetCanBlock(v bool) *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart {
	s.CanBlock = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart) SetDevice(v string) *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart {
	s.Device = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart) SetNeed(v bool) *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart {
	s.Need = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart) SetPath(v string) *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart {
	s.Path = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart) SetSizeBytes(v int64) *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart {
	s.SizeBytes = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart) Validate() error {
	return dara.Validate(s)
}

type DescribeSourceServersResponseBodySourceServersSourceServerTags struct {
	Tag []*DescribeSourceServersResponseBodySourceServersSourceServerTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerTags) GoString() string {
	return s.String()
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerTags) GetTag() []*DescribeSourceServersResponseBodySourceServersSourceServerTagsTag {
	return s.Tag
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerTags) SetTag(v []*DescribeSourceServersResponseBodySourceServersSourceServerTagsTag) *DescribeSourceServersResponseBodySourceServersSourceServerTags {
	s.Tag = v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerTags) Validate() error {
	return dara.Validate(s)
}

type DescribeSourceServersResponseBodySourceServersSourceServerTagsTag struct {
	// The key of tag N that is attached to the SMC resource. Valid values of N: 1 to 20.
	//
	// You cannot specify an empty string as a tag key. The tag key can be up to 64 characters in length and cannot contain http:// or https://. The tag key cannot start with acs: or aliyun.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N that is attached to the SMC resource. Valid values of N: 1 to 20.
	//
	// The tag key can be an empty string. The tag value can be up to 64 characters in length and cannot contain http:// or https://.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerTagsTag) SetKey(v string) *DescribeSourceServersResponseBodySourceServersSourceServerTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerTagsTag) SetValue(v string) *DescribeSourceServersResponseBodySourceServersSourceServerTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerTagsTag) Validate() error {
	return dara.Validate(s)
}
