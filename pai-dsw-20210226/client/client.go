// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type EcsSpec struct {
	// 实例类型
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// cpu数量
	Cpu *int64 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// gpu卡数
	Gpu *int64 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// 内存(GiB)
	MemoryInGiB *int64 `json:"MemoryInGiB,omitempty" xml:"MemoryInGiB,omitempty"`
	// 磁盘类型
	SystemDiskCategory *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	// 磁盘大小(GiB)
	SystemDiskSizeInGiB *int64 `json:"SystemDiskSizeInGiB,omitempty" xml:"SystemDiskSizeInGiB,omitempty"`
	// GPU卡类型
	GpuType *string `json:"GpuType,omitempty" xml:"GpuType,omitempty"`
}

func (s EcsSpec) String() string {
	return tea.Prettify(s)
}

func (s EcsSpec) GoString() string {
	return s.String()
}

func (s *EcsSpec) SetInstanceType(v string) *EcsSpec {
	s.InstanceType = &v
	return s
}

func (s *EcsSpec) SetCpu(v int64) *EcsSpec {
	s.Cpu = &v
	return s
}

func (s *EcsSpec) SetGpu(v int64) *EcsSpec {
	s.Gpu = &v
	return s
}

func (s *EcsSpec) SetMemoryInGiB(v int64) *EcsSpec {
	s.MemoryInGiB = &v
	return s
}

func (s *EcsSpec) SetSystemDiskCategory(v string) *EcsSpec {
	s.SystemDiskCategory = &v
	return s
}

func (s *EcsSpec) SetSystemDiskSizeInGiB(v int64) *EcsSpec {
	s.SystemDiskSizeInGiB = &v
	return s
}

func (s *EcsSpec) SetGpuType(v string) *EcsSpec {
	s.GpuType = &v
	return s
}

type InstanceSnapshot struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例快照ID
	InstanceSnapshotId *string `json:"InstanceSnapshotId,omitempty" xml:"InstanceSnapshotId,omitempty"`
	// 实例快照状态
	InstanceSnapshotStatus *string `json:"InstanceSnapshotStatus,omitempty" xml:"InstanceSnapshotStatus,omitempty"`
	// 实例快照名称
	InstanceSnapshotName *string `json:"InstanceSnapshotName,omitempty" xml:"InstanceSnapshotName,omitempty"`
	// 实例快照标签
	InstanceSnapshotTag *string `json:"InstanceSnapshotTag,omitempty" xml:"InstanceSnapshotTag,omitempty"`
	// 实例快照存储地址
	InstanceSnapshotRepoUrl *string `json:"InstanceSnapshotRepoUrl,omitempty" xml:"InstanceSnapshotRepoUrl,omitempty"`
	// 实例快照保存时间（GMT）
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 实例快照修改时间（GMT）
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// 实例快照描述
	InstanceSnapshotDescription *string `json:"InstanceSnapshotDescription,omitempty" xml:"InstanceSnapshotDescription,omitempty"`
}

func (s InstanceSnapshot) String() string {
	return tea.Prettify(s)
}

func (s InstanceSnapshot) GoString() string {
	return s.String()
}

func (s *InstanceSnapshot) SetInstanceId(v string) *InstanceSnapshot {
	s.InstanceId = &v
	return s
}

func (s *InstanceSnapshot) SetInstanceSnapshotId(v string) *InstanceSnapshot {
	s.InstanceSnapshotId = &v
	return s
}

func (s *InstanceSnapshot) SetInstanceSnapshotStatus(v string) *InstanceSnapshot {
	s.InstanceSnapshotStatus = &v
	return s
}

func (s *InstanceSnapshot) SetInstanceSnapshotName(v string) *InstanceSnapshot {
	s.InstanceSnapshotName = &v
	return s
}

func (s *InstanceSnapshot) SetInstanceSnapshotTag(v string) *InstanceSnapshot {
	s.InstanceSnapshotTag = &v
	return s
}

func (s *InstanceSnapshot) SetInstanceSnapshotRepoUrl(v string) *InstanceSnapshot {
	s.InstanceSnapshotRepoUrl = &v
	return s
}

func (s *InstanceSnapshot) SetGmtCreateTime(v string) *InstanceSnapshot {
	s.GmtCreateTime = &v
	return s
}

func (s *InstanceSnapshot) SetGmtModifiedTime(v string) *InstanceSnapshot {
	s.GmtModifiedTime = &v
	return s
}

func (s *InstanceSnapshot) SetInstanceSnapshotDescription(v string) *InstanceSnapshot {
	s.InstanceSnapshotDescription = &v
	return s
}

type Image struct {
	// 镜像ID
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 资源类型
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	// 算法框架
	Framework *string `json:"Framework,omitempty" xml:"Framework,omitempty"`
	// 算法框架版本
	FrameworkVersion *string `json:"FrameworkVersion,omitempty" xml:"FrameworkVersion,omitempty"`
	// 镜像操作系统分发版
	OS *string `json:"OS,omitempty" xml:"OS,omitempty"`
	// 分发版版本
	OSVersion *string `json:"OSVersion,omitempty" xml:"OSVersion,omitempty"`
	// Cuda版本
	CudaVersion *string `json:"CudaVersion,omitempty" xml:"CudaVersion,omitempty"`
	// 镜像类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s Image) String() string {
	return tea.Prettify(s)
}

func (s Image) GoString() string {
	return s.String()
}

func (s *Image) SetImageId(v string) *Image {
	s.ImageId = &v
	return s
}

func (s *Image) SetImageName(v string) *Image {
	s.ImageName = &v
	return s
}

func (s *Image) SetInstanceId(v string) *Image {
	s.InstanceId = &v
	return s
}

func (s *Image) SetAcceleratorType(v string) *Image {
	s.AcceleratorType = &v
	return s
}

func (s *Image) SetFramework(v string) *Image {
	s.Framework = &v
	return s
}

func (s *Image) SetFrameworkVersion(v string) *Image {
	s.FrameworkVersion = &v
	return s
}

func (s *Image) SetOS(v string) *Image {
	s.OS = &v
	return s
}

func (s *Image) SetOSVersion(v string) *Image {
	s.OSVersion = &v
	return s
}

func (s *Image) SetCudaVersion(v string) *Image {
	s.CudaVersion = &v
	return s
}

func (s *Image) SetType(v string) *Image {
	s.Type = &v
	return s
}

type InstanceShutdownTimer struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 定时关机时间
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	// 定时关机创建时间
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// 定时关机修改时间
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
}

func (s InstanceShutdownTimer) String() string {
	return tea.Prettify(s)
}

func (s InstanceShutdownTimer) GoString() string {
	return s.String()
}

func (s *InstanceShutdownTimer) SetInstanceId(v string) *InstanceShutdownTimer {
	s.InstanceId = &v
	return s
}

func (s *InstanceShutdownTimer) SetScheduleTime(v string) *InstanceShutdownTimer {
	s.ScheduleTime = &v
	return s
}

func (s *InstanceShutdownTimer) SetGmtModifiedTime(v string) *InstanceShutdownTimer {
	s.GmtModifiedTime = &v
	return s
}

func (s *InstanceShutdownTimer) SetGmtCreateTime(v string) *InstanceShutdownTimer {
	s.GmtCreateTime = &v
	return s
}

type Instance struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// ecs规格
	EcsSpec *string `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	// 实例状态
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// jupyter链接
	JupyterlabUrl *string `json:"JupyterlabUrl,omitempty" xml:"JupyterlabUrl,omitempty"`
	// webIde链接
	WebIdeUrl *string `json:"WebIdeUrl,omitempty" xml:"WebIdeUrl,omitempty"`
	// 命令行终端链接
	TerminalUrl *string `json:"TerminalUrl,omitempty" xml:"TerminalUrl,omitempty"`
	// 累计运行时间(分钟)
	AccumulativeRunningTimeInMinutes *int64 `json:"AccumulativeRunningTimeInMinutes,omitempty" xml:"AccumulativeRunningTimeInMinutes,omitempty"`
	// 镜像ID
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// 镜像链接
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// 创建时间(GMT)
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 修改时间(GMT)
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// 用户ID
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// nas文件系统ID
	NasFileSystemId *string `json:"NasFileSystemId,omitempty" xml:"NasFileSystemId,omitempty"`
	// 被打通VPC配置
	UserVpc *UserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty"`
	// 定时关机任务
	InstanceShutdownTimer *InstanceShutdownTimer `json:"InstanceShutdownTimer,omitempty" xml:"InstanceShutdownTimer,omitempty"`
	// 工作空间id
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// 工作空间名称
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s Instance) String() string {
	return tea.Prettify(s)
}

func (s Instance) GoString() string {
	return s.String()
}

func (s *Instance) SetInstanceId(v string) *Instance {
	s.InstanceId = &v
	return s
}

func (s *Instance) SetInstanceName(v string) *Instance {
	s.InstanceName = &v
	return s
}

func (s *Instance) SetEcsSpec(v string) *Instance {
	s.EcsSpec = &v
	return s
}

func (s *Instance) SetInstanceStatus(v string) *Instance {
	s.InstanceStatus = &v
	return s
}

func (s *Instance) SetJupyterlabUrl(v string) *Instance {
	s.JupyterlabUrl = &v
	return s
}

func (s *Instance) SetWebIdeUrl(v string) *Instance {
	s.WebIdeUrl = &v
	return s
}

func (s *Instance) SetTerminalUrl(v string) *Instance {
	s.TerminalUrl = &v
	return s
}

func (s *Instance) SetAccumulativeRunningTimeInMinutes(v int64) *Instance {
	s.AccumulativeRunningTimeInMinutes = &v
	return s
}

func (s *Instance) SetImageId(v string) *Instance {
	s.ImageId = &v
	return s
}

func (s *Instance) SetImageUrl(v string) *Instance {
	s.ImageUrl = &v
	return s
}

func (s *Instance) SetGmtCreateTime(v string) *Instance {
	s.GmtCreateTime = &v
	return s
}

func (s *Instance) SetGmtModifiedTime(v string) *Instance {
	s.GmtModifiedTime = &v
	return s
}

func (s *Instance) SetUserId(v string) *Instance {
	s.UserId = &v
	return s
}

func (s *Instance) SetNasFileSystemId(v string) *Instance {
	s.NasFileSystemId = &v
	return s
}

func (s *Instance) SetUserVpc(v *UserVpc) *Instance {
	s.UserVpc = v
	return s
}

func (s *Instance) SetInstanceShutdownTimer(v *InstanceShutdownTimer) *Instance {
	s.InstanceShutdownTimer = v
	return s
}

func (s *Instance) SetWorkspaceId(v string) *Instance {
	s.WorkspaceId = &v
	return s
}

func (s *Instance) SetWorkspaceName(v string) *Instance {
	s.WorkspaceName = &v
	return s
}

type UserVpc struct {
	// 虚拟网络ID
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// 虚拟交换机ID
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	// 角色标识码
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// 安全组ID
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s UserVpc) String() string {
	return tea.Prettify(s)
}

func (s UserVpc) GoString() string {
	return s.String()
}

func (s *UserVpc) SetVpcId(v string) *UserVpc {
	s.VpcId = &v
	return s
}

func (s *UserVpc) SetVswitchId(v string) *UserVpc {
	s.VswitchId = &v
	return s
}

func (s *UserVpc) SetRoleArn(v string) *UserVpc {
	s.RoleArn = &v
	return s
}

func (s *UserVpc) SetSecurityGroupId(v string) *UserVpc {
	s.SecurityGroupId = &v
	return s
}

type CreateInstanceSnapshotRequest struct {
	// 实例快照名称
	InstanceSnapshotName *string `json:"InstanceSnapshotName,omitempty" xml:"InstanceSnapshotName,omitempty"`
	// 实例快照存储地址（可选）
	InstanceSnapshotRepoUrl *string `json:"InstanceSnapshotRepoUrl,omitempty" xml:"InstanceSnapshotRepoUrl,omitempty"`
	// 实例快照描述
	InstanceSnapshotDescription *string `json:"InstanceSnapshotDescription,omitempty" xml:"InstanceSnapshotDescription,omitempty"`
}

func (s CreateInstanceSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceSnapshotRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceSnapshotRequest) SetInstanceSnapshotName(v string) *CreateInstanceSnapshotRequest {
	s.InstanceSnapshotName = &v
	return s
}

func (s *CreateInstanceSnapshotRequest) SetInstanceSnapshotRepoUrl(v string) *CreateInstanceSnapshotRequest {
	s.InstanceSnapshotRepoUrl = &v
	return s
}

func (s *CreateInstanceSnapshotRequest) SetInstanceSnapshotDescription(v string) *CreateInstanceSnapshotRequest {
	s.InstanceSnapshotDescription = &v
	return s
}

type CreateInstanceSnapshotResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例快照ID
	InstanceSnapshotId *string `json:"InstanceSnapshotId,omitempty" xml:"InstanceSnapshotId,omitempty"`
}

func (s CreateInstanceSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceSnapshotResponseBody) SetRequestId(v string) *CreateInstanceSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceSnapshotResponseBody) SetInstanceId(v string) *CreateInstanceSnapshotResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceSnapshotResponseBody) SetInstanceSnapshotId(v string) *CreateInstanceSnapshotResponseBody {
	s.InstanceSnapshotId = &v
	return s
}

type CreateInstanceSnapshotResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateInstanceSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateInstanceSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceSnapshotResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceSnapshotResponse) SetHeaders(v map[string]*string) *CreateInstanceSnapshotResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceSnapshotResponse) SetBody(v *CreateInstanceSnapshotResponseBody) *CreateInstanceSnapshotResponse {
	s.Body = v
	return s
}

type DeleteInstanceShutdownTimerResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteInstanceShutdownTimerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceShutdownTimerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceShutdownTimerResponseBody) SetRequestId(v string) *DeleteInstanceShutdownTimerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstanceShutdownTimerResponseBody) SetInstanceId(v string) *DeleteInstanceShutdownTimerResponseBody {
	s.InstanceId = &v
	return s
}

type DeleteInstanceShutdownTimerResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteInstanceShutdownTimerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteInstanceShutdownTimerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceShutdownTimerResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceShutdownTimerResponse) SetHeaders(v map[string]*string) *DeleteInstanceShutdownTimerResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceShutdownTimerResponse) SetBody(v *DeleteInstanceShutdownTimerResponseBody) *DeleteInstanceShutdownTimerResponse {
	s.Body = v
	return s
}

type GetInstanceShutdownTimerResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 任务创建时间(GMT)
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 任务修改时间(GMT)
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// 定时关机时间(GMT)
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
}

func (s GetInstanceShutdownTimerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceShutdownTimerResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceShutdownTimerResponseBody) SetRequestId(v string) *GetInstanceShutdownTimerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceShutdownTimerResponseBody) SetInstanceId(v string) *GetInstanceShutdownTimerResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceShutdownTimerResponseBody) SetGmtCreateTime(v string) *GetInstanceShutdownTimerResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetInstanceShutdownTimerResponseBody) SetGmtModifiedTime(v string) *GetInstanceShutdownTimerResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetInstanceShutdownTimerResponseBody) SetScheduleTime(v string) *GetInstanceShutdownTimerResponseBody {
	s.ScheduleTime = &v
	return s
}

type GetInstanceShutdownTimerResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetInstanceShutdownTimerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInstanceShutdownTimerResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceShutdownTimerResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceShutdownTimerResponse) SetHeaders(v map[string]*string) *GetInstanceShutdownTimerResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceShutdownTimerResponse) SetBody(v *GetInstanceShutdownTimerResponseBody) *GetInstanceShutdownTimerResponse {
	s.Body = v
	return s
}

type StopInstanceResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s StopInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopInstanceResponseBody) SetRequestId(v string) *StopInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopInstanceResponseBody) SetInstanceId(v string) *StopInstanceResponseBody {
	s.InstanceId = &v
	return s
}

type StopInstanceResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StopInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopInstanceResponse) SetHeaders(v map[string]*string) *StopInstanceResponse {
	s.Headers = v
	return s
}

func (s *StopInstanceResponse) SetBody(v *StopInstanceResponseBody) *StopInstanceResponse {
	s.Body = v
	return s
}

type UpdateInstanceRequest struct {
	// 修改后实例名称
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s UpdateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequest) SetInstanceName(v string) *UpdateInstanceRequest {
	s.InstanceName = &v
	return s
}

type UpdateInstanceResponseBody struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponseBody) SetInstanceId(v string) *UpdateInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetRequestId(v string) *UpdateInstanceResponseBody {
	s.RequestId = &v
	return s
}

type UpdateInstanceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponse) SetHeaders(v map[string]*string) *UpdateInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceResponse) SetBody(v *UpdateInstanceResponseBody) *UpdateInstanceResponse {
	s.Body = v
	return s
}

type ListEcsSpecsRequest struct {
	// 每页返回的实例数
	AcceleratorTypeEquals *string `json:"AcceleratorTypeEquals,omitempty" xml:"AcceleratorTypeEquals,omitempty"`
}

func (s ListEcsSpecsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEcsSpecsRequest) GoString() string {
	return s.String()
}

func (s *ListEcsSpecsRequest) SetAcceleratorTypeEquals(v string) *ListEcsSpecsRequest {
	s.AcceleratorTypeEquals = &v
	return s
}

type ListEcsSpecsResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 请求ecs规格列表
	EcsSpecs []*EcsSpec `json:"EcsSpecs,omitempty" xml:"EcsSpecs,omitempty" type:"Repeated"`
	// 符合要求的ecs规格数量
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEcsSpecsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEcsSpecsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEcsSpecsResponseBody) SetRequestId(v string) *ListEcsSpecsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEcsSpecsResponseBody) SetEcsSpecs(v []*EcsSpec) *ListEcsSpecsResponseBody {
	s.EcsSpecs = v
	return s
}

func (s *ListEcsSpecsResponseBody) SetTotalCount(v int64) *ListEcsSpecsResponseBody {
	s.TotalCount = &v
	return s
}

type ListEcsSpecsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListEcsSpecsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEcsSpecsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEcsSpecsResponse) GoString() string {
	return s.String()
}

func (s *ListEcsSpecsResponse) SetHeaders(v map[string]*string) *ListEcsSpecsResponse {
	s.Headers = v
	return s
}

func (s *ListEcsSpecsResponse) SetBody(v *ListEcsSpecsResponseBody) *ListEcsSpecsResponse {
	s.Body = v
	return s
}

type GetInstancesStatisticsRequest struct {
	// 工作空间id列表
	WorkspaceIds *string `json:"WorkspaceIds,omitempty" xml:"WorkspaceIds,omitempty"`
}

func (s GetInstancesStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetInstancesStatisticsRequest) SetWorkspaceIds(v string) *GetInstancesStatisticsRequest {
	s.WorkspaceIds = &v
	return s
}

type GetInstancesStatisticsResponseBody struct {
	// Id of the request
	RequestID *string `json:"RequestID,omitempty" xml:"RequestID,omitempty"`
	// 统计数据
	Statistics map[string]interface{} `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
}

func (s GetInstancesStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstancesStatisticsResponseBody) SetRequestID(v string) *GetInstancesStatisticsResponseBody {
	s.RequestID = &v
	return s
}

func (s *GetInstancesStatisticsResponseBody) SetStatistics(v map[string]interface{}) *GetInstancesStatisticsResponseBody {
	s.Statistics = v
	return s
}

type GetInstancesStatisticsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetInstancesStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInstancesStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetInstancesStatisticsResponse) SetHeaders(v map[string]*string) *GetInstancesStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetInstancesStatisticsResponse) SetBody(v *GetInstancesStatisticsResponseBody) *GetInstancesStatisticsResponse {
	s.Body = v
	return s
}

type DeleteInstanceSnapshotResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例快照ID
	InstanceSnapshotId *string `json:"InstanceSnapshotId,omitempty" xml:"InstanceSnapshotId,omitempty"`
}

func (s DeleteInstanceSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceSnapshotResponseBody) SetRequestId(v string) *DeleteInstanceSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstanceSnapshotResponseBody) SetInstanceId(v string) *DeleteInstanceSnapshotResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DeleteInstanceSnapshotResponseBody) SetInstanceSnapshotId(v string) *DeleteInstanceSnapshotResponseBody {
	s.InstanceSnapshotId = &v
	return s
}

type DeleteInstanceSnapshotResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteInstanceSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteInstanceSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceSnapshotResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceSnapshotResponse) SetHeaders(v map[string]*string) *DeleteInstanceSnapshotResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceSnapshotResponse) SetBody(v *DeleteInstanceSnapshotResponseBody) *DeleteInstanceSnapshotResponse {
	s.Body = v
	return s
}

type UpdateInstanceSnapshotRequest struct {
	// 实例快照名称
	InstanceSnapshotName *string `json:"InstanceSnapshotName,omitempty" xml:"InstanceSnapshotName,omitempty"`
}

func (s UpdateInstanceSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceSnapshotRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceSnapshotRequest) SetInstanceSnapshotName(v string) *UpdateInstanceSnapshotRequest {
	s.InstanceSnapshotName = &v
	return s
}

type UpdateInstanceSnapshotResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例镜像ID
	InstanceSnapshotId *string `json:"InstanceSnapshotId,omitempty" xml:"InstanceSnapshotId,omitempty"`
}

func (s UpdateInstanceSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceSnapshotResponseBody) SetRequestId(v string) *UpdateInstanceSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceSnapshotResponseBody) SetInstanceId(v string) *UpdateInstanceSnapshotResponseBody {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceSnapshotResponseBody) SetInstanceSnapshotId(v string) *UpdateInstanceSnapshotResponseBody {
	s.InstanceSnapshotId = &v
	return s
}

type UpdateInstanceSnapshotResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateInstanceSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateInstanceSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceSnapshotResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceSnapshotResponse) SetHeaders(v map[string]*string) *UpdateInstanceSnapshotResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceSnapshotResponse) SetBody(v *UpdateInstanceSnapshotResponseBody) *UpdateInstanceSnapshotResponse {
	s.Body = v
	return s
}

type CreateInstanceRequest struct {
	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// 实例规格
	EcsSpec *string `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	// 镜像id
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// nas文件系统id
	NasFileSystemId *string `json:"NasFileSystemId,omitempty" xml:"NasFileSystemId,omitempty"`
	// 打通的vpc网络配置
	UserVpc *UserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty"`
	// 镜像地址
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) SetInstanceName(v string) *CreateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceRequest) SetEcsSpec(v string) *CreateInstanceRequest {
	s.EcsSpec = &v
	return s
}

func (s *CreateInstanceRequest) SetImageId(v string) *CreateInstanceRequest {
	s.ImageId = &v
	return s
}

func (s *CreateInstanceRequest) SetNasFileSystemId(v string) *CreateInstanceRequest {
	s.NasFileSystemId = &v
	return s
}

func (s *CreateInstanceRequest) SetUserVpc(v *UserVpc) *CreateInstanceRequest {
	s.UserVpc = v
	return s
}

func (s *CreateInstanceRequest) SetImageUrl(v string) *CreateInstanceRequest {
	s.ImageUrl = &v
	return s
}

type CreateInstanceResponseBody struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) SetInstanceId(v string) *CreateInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

type CreateInstanceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponse) SetHeaders(v map[string]*string) *CreateInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceResponse) SetBody(v *CreateInstanceResponseBody) *CreateInstanceResponse {
	s.Body = v
	return s
}

type GetAuthorizationResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 是否已经给DSW服务账号授权
	Authorized *bool `json:"Authorized,omitempty" xml:"Authorized,omitempty"`
	// 授权失败错误代码
	AuthorizationFailedCode *string `json:"AuthorizationFailedCode,omitempty" xml:"AuthorizationFailedCode,omitempty"`
	// 授权失败错误消息
	AuthorizationFailedMessage *string `json:"AuthorizationFailedMessage,omitempty" xml:"AuthorizationFailedMessage,omitempty"`
}

func (s GetAuthorizationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuthorizationResponseBody) SetRequestId(v string) *GetAuthorizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAuthorizationResponseBody) SetAuthorized(v bool) *GetAuthorizationResponseBody {
	s.Authorized = &v
	return s
}

func (s *GetAuthorizationResponseBody) SetAuthorizationFailedCode(v string) *GetAuthorizationResponseBody {
	s.AuthorizationFailedCode = &v
	return s
}

func (s *GetAuthorizationResponseBody) SetAuthorizationFailedMessage(v string) *GetAuthorizationResponseBody {
	s.AuthorizationFailedMessage = &v
	return s
}

type GetAuthorizationResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAuthorizationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *GetAuthorizationResponse) SetHeaders(v map[string]*string) *GetAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *GetAuthorizationResponse) SetBody(v *GetAuthorizationResponseBody) *GetAuthorizationResponse {
	s.Body = v
	return s
}

type ListImagesRequest struct {
	// 资源类型
	AcceleratorTypeEquals *string `json:"AcceleratorTypeEquals,omitempty" xml:"AcceleratorTypeEquals,omitempty"`
	// 镜像类型
	ImageTypeEquals *string `json:"ImageTypeEquals,omitempty" xml:"ImageTypeEquals,omitempty"`
	// 容器名称关键字
	ImageNameContains *string `json:"ImageNameContains,omitempty" xml:"ImageNameContains,omitempty"`
}

func (s ListImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListImagesRequest) GoString() string {
	return s.String()
}

func (s *ListImagesRequest) SetAcceleratorTypeEquals(v string) *ListImagesRequest {
	s.AcceleratorTypeEquals = &v
	return s
}

func (s *ListImagesRequest) SetImageTypeEquals(v string) *ListImagesRequest {
	s.ImageTypeEquals = &v
	return s
}

func (s *ListImagesRequest) SetImageNameContains(v string) *ListImagesRequest {
	s.ImageNameContains = &v
	return s
}

type ListImagesResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 镜像列表
	Images []*Image `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
}

func (s ListImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBody) SetRequestId(v string) *ListImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListImagesResponseBody) SetImages(v []*Image) *ListImagesResponseBody {
	s.Images = v
	return s
}

type ListImagesResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponse) GoString() string {
	return s.String()
}

func (s *ListImagesResponse) SetHeaders(v map[string]*string) *ListImagesResponse {
	s.Headers = v
	return s
}

func (s *ListImagesResponse) SetBody(v *ListImagesResponseBody) *ListImagesResponse {
	s.Body = v
	return s
}

type FoobarResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s FoobarResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FoobarResponseBody) GoString() string {
	return s.String()
}

func (s *FoobarResponseBody) SetRequestId(v string) *FoobarResponseBody {
	s.RequestId = &v
	return s
}

type FoobarResponse struct {
	Headers map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FoobarResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FoobarResponse) String() string {
	return tea.Prettify(s)
}

func (s FoobarResponse) GoString() string {
	return s.String()
}

func (s *FoobarResponse) SetHeaders(v map[string]*string) *FoobarResponse {
	s.Headers = v
	return s
}

func (s *FoobarResponse) SetBody(v *FoobarResponseBody) *FoobarResponse {
	s.Body = v
	return s
}

type ListInstanceSnapshotsResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 镜像快照列表
	InstanceSnapshots []*InstanceSnapshot `json:"InstanceSnapshots,omitempty" xml:"InstanceSnapshots,omitempty" type:"Repeated"`
}

func (s ListInstanceSnapshotsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceSnapshotsResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceSnapshotsResponseBody) SetRequestId(v string) *ListInstanceSnapshotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceSnapshotsResponseBody) SetInstanceSnapshots(v []*InstanceSnapshot) *ListInstanceSnapshotsResponseBody {
	s.InstanceSnapshots = v
	return s
}

type ListInstanceSnapshotsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInstanceSnapshotsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInstanceSnapshotsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceSnapshotsResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceSnapshotsResponse) SetHeaders(v map[string]*string) *ListInstanceSnapshotsResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceSnapshotsResponse) SetBody(v *ListInstanceSnapshotsResponseBody) *ListInstanceSnapshotsResponse {
	s.Body = v
	return s
}

type CreateInstanceShutdownTimerRequest struct {
	// 定时关机时间(GMT)
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
}

func (s CreateInstanceShutdownTimerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceShutdownTimerRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceShutdownTimerRequest) SetScheduleTime(v string) *CreateInstanceShutdownTimerRequest {
	s.ScheduleTime = &v
	return s
}

type CreateInstanceShutdownTimerResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateInstanceShutdownTimerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceShutdownTimerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceShutdownTimerResponseBody) SetRequestId(v string) *CreateInstanceShutdownTimerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceShutdownTimerResponseBody) SetInstanceId(v string) *CreateInstanceShutdownTimerResponseBody {
	s.InstanceId = &v
	return s
}

type CreateInstanceShutdownTimerResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateInstanceShutdownTimerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateInstanceShutdownTimerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceShutdownTimerResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceShutdownTimerResponse) SetHeaders(v map[string]*string) *CreateInstanceShutdownTimerResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceShutdownTimerResponse) SetBody(v *CreateInstanceShutdownTimerResponseBody) *CreateInstanceShutdownTimerResponse {
	s.Body = v
	return s
}

type GetInstanceResponseBody struct {
	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// ecs规格
	EcsSpec *string `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	// 实例状态
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// jupyter链接
	JupyterlabUrl *string `json:"JupyterlabUrl,omitempty" xml:"JupyterlabUrl,omitempty"`
	// web ide链接
	WebIdeUrl *string `json:"WebIdeUrl,omitempty" xml:"WebIdeUrl,omitempty"`
	// 命令行终端链接
	TerminalUrl *string `json:"TerminalUrl,omitempty" xml:"TerminalUrl,omitempty"`
	// 累计运行时间(分钟)
	AccumulativeRunningTimeInMinutes *int64 `json:"AccumulativeRunningTimeInMinutes,omitempty" xml:"AccumulativeRunningTimeInMinutes,omitempty"`
	// 镜像ID
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// 镜像链接
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// 实例创建时间(GMT)
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 实例修改时间(GMT)
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// 用户ID
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// nas文件系统ID
	NasFileSystemId *string `json:"NasFileSystemId,omitempty" xml:"NasFileSystemId,omitempty"`
	// 被打通VPC配置
	UserVpc *UserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty"`
	// 定时关机任务
	InstanceShutdownTimer *InstanceShutdownTimer `json:"InstanceShutdownTimer,omitempty" xml:"InstanceShutdownTimer,omitempty"`
}

func (s GetInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBody) SetInstanceName(v string) *GetInstanceResponseBody {
	s.InstanceName = &v
	return s
}

func (s *GetInstanceResponseBody) SetInstanceId(v string) *GetInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceResponseBody) SetRequestId(v string) *GetInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceResponseBody) SetEcsSpec(v string) *GetInstanceResponseBody {
	s.EcsSpec = &v
	return s
}

func (s *GetInstanceResponseBody) SetInstanceStatus(v string) *GetInstanceResponseBody {
	s.InstanceStatus = &v
	return s
}

func (s *GetInstanceResponseBody) SetJupyterlabUrl(v string) *GetInstanceResponseBody {
	s.JupyterlabUrl = &v
	return s
}

func (s *GetInstanceResponseBody) SetWebIdeUrl(v string) *GetInstanceResponseBody {
	s.WebIdeUrl = &v
	return s
}

func (s *GetInstanceResponseBody) SetTerminalUrl(v string) *GetInstanceResponseBody {
	s.TerminalUrl = &v
	return s
}

func (s *GetInstanceResponseBody) SetAccumulativeRunningTimeInMinutes(v int64) *GetInstanceResponseBody {
	s.AccumulativeRunningTimeInMinutes = &v
	return s
}

func (s *GetInstanceResponseBody) SetImageId(v string) *GetInstanceResponseBody {
	s.ImageId = &v
	return s
}

func (s *GetInstanceResponseBody) SetImageUrl(v string) *GetInstanceResponseBody {
	s.ImageUrl = &v
	return s
}

func (s *GetInstanceResponseBody) SetGmtCreateTime(v string) *GetInstanceResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetInstanceResponseBody) SetGmtModifiedTime(v string) *GetInstanceResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetInstanceResponseBody) SetUserId(v string) *GetInstanceResponseBody {
	s.UserId = &v
	return s
}

func (s *GetInstanceResponseBody) SetNasFileSystemId(v string) *GetInstanceResponseBody {
	s.NasFileSystemId = &v
	return s
}

func (s *GetInstanceResponseBody) SetUserVpc(v *UserVpc) *GetInstanceResponseBody {
	s.UserVpc = v
	return s
}

func (s *GetInstanceResponseBody) SetInstanceShutdownTimer(v *InstanceShutdownTimer) *GetInstanceResponseBody {
	s.InstanceShutdownTimer = v
	return s
}

type GetInstanceResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceResponse) SetHeaders(v map[string]*string) *GetInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceResponse) SetBody(v *GetInstanceResponseBody) *GetInstanceResponse {
	s.Body = v
	return s
}

type ListInstancesRequest struct {
	// 实例名称关键字
	InstanceNameContains *string `json:"InstanceNameContains,omitempty" xml:"InstanceNameContains,omitempty"`
	// 实例状态
	InstanceStatusEquals *string `json:"InstanceStatusEquals,omitempty" xml:"InstanceStatusEquals,omitempty"`
	// 当前页
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页返回的实例数
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 升序降序
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// 排序字段
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) SetInstanceNameContains(v string) *ListInstancesRequest {
	s.InstanceNameContains = &v
	return s
}

func (s *ListInstancesRequest) SetInstanceStatusEquals(v string) *ListInstancesRequest {
	s.InstanceStatusEquals = &v
	return s
}

func (s *ListInstancesRequest) SetPageNumber(v string) *ListInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesRequest) SetPageSize(v string) *ListInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstancesRequest) SetSortOrder(v string) *ListInstancesRequest {
	s.SortOrder = &v
	return s
}

func (s *ListInstancesRequest) SetSortBy(v string) *ListInstancesRequest {
	s.SortBy = &v
	return s
}

type ListInstancesResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 实例列表
	Instances []*Instance `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// 当前页
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页返回的实例数
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 符合条件的实例数
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetInstances(v []*Instance) *ListInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *ListInstancesResponseBody) SetPageNumber(v int64) *ListInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesResponseBody) SetPageSize(v int64) *ListInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListInstancesResponseBody) SetTotalCount(v int64) *ListInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type ListInstancesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListInstancesResponse) SetHeaders(v map[string]*string) *ListInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListInstancesResponse) SetBody(v *ListInstancesResponseBody) *ListInstancesResponse {
	s.Body = v
	return s
}

type GetInstanceSnapshotResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例快照ID
	InstanceSnapshotId *string `json:"InstanceSnapshotId,omitempty" xml:"InstanceSnapshotId,omitempty"`
	// 实例快照状态
	InstanceSnapshotStatus *string `json:"InstanceSnapshotStatus,omitempty" xml:"InstanceSnapshotStatus,omitempty"`
	// 实例快照名称
	InstanceSnapshotName *string `json:"InstanceSnapshotName,omitempty" xml:"InstanceSnapshotName,omitempty"`
	// 实例快照标签
	InstanceSnapshotTag *string `json:"InstanceSnapshotTag,omitempty" xml:"InstanceSnapshotTag,omitempty"`
	// 实例快照存储地址
	InstanceSnapshotRepoUrl *string `json:"InstanceSnapshotRepoUrl,omitempty" xml:"InstanceSnapshotRepoUrl,omitempty"`
	// 实例快照保存时间（GMT）
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 实例快照修改时间（GMT）
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// 实例快照描述
	InstanceSnapshotDescription *string `json:"InstanceSnapshotDescription,omitempty" xml:"InstanceSnapshotDescription,omitempty"`
}

func (s GetInstanceSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceSnapshotResponseBody) SetRequestId(v string) *GetInstanceSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetInstanceId(v string) *GetInstanceSnapshotResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetInstanceSnapshotId(v string) *GetInstanceSnapshotResponseBody {
	s.InstanceSnapshotId = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetInstanceSnapshotStatus(v string) *GetInstanceSnapshotResponseBody {
	s.InstanceSnapshotStatus = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetInstanceSnapshotName(v string) *GetInstanceSnapshotResponseBody {
	s.InstanceSnapshotName = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetInstanceSnapshotTag(v string) *GetInstanceSnapshotResponseBody {
	s.InstanceSnapshotTag = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetInstanceSnapshotRepoUrl(v string) *GetInstanceSnapshotResponseBody {
	s.InstanceSnapshotRepoUrl = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetGmtCreateTime(v string) *GetInstanceSnapshotResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetGmtModifiedTime(v string) *GetInstanceSnapshotResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetInstanceSnapshotDescription(v string) *GetInstanceSnapshotResponseBody {
	s.InstanceSnapshotDescription = &v
	return s
}

type GetInstanceSnapshotResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetInstanceSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInstanceSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceSnapshotResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceSnapshotResponse) SetHeaders(v map[string]*string) *GetInstanceSnapshotResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceSnapshotResponse) SetBody(v *GetInstanceSnapshotResponseBody) *GetInstanceSnapshotResponse {
	s.Body = v
	return s
}

type StartInstanceResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s StartInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartInstanceResponseBody) SetRequestId(v string) *StartInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartInstanceResponseBody) SetInstanceId(v string) *StartInstanceResponseBody {
	s.InstanceId = &v
	return s
}

type StartInstanceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartInstanceResponse) SetHeaders(v map[string]*string) *StartInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartInstanceResponse) SetBody(v *StartInstanceResponseBody) *StartInstanceResponse {
	s.Body = v
	return s
}

type DeleteInstanceResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponseBody) SetRequestId(v string) *DeleteInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetInstanceId(v string) *DeleteInstanceResponseBody {
	s.InstanceId = &v
	return s
}

type DeleteInstanceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponse) SetHeaders(v map[string]*string) *DeleteInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceResponse) SetBody(v *DeleteInstanceResponseBody) *DeleteInstanceResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("pai-dsw"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateInstanceSnapshot(InstanceId *string, request *CreateInstanceSnapshotRequest) (_result *CreateInstanceSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateInstanceSnapshotResponse{}
	_body, _err := client.CreateInstanceSnapshotWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateInstanceSnapshotWithOptions(InstanceId *string, request *CreateInstanceSnapshotRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateInstanceSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceSnapshotName)) {
		body["InstanceSnapshotName"] = request.InstanceSnapshotName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceSnapshotRepoUrl)) {
		body["InstanceSnapshotRepoUrl"] = request.InstanceSnapshotRepoUrl
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceSnapshotDescription)) {
		body["InstanceSnapshotDescription"] = request.InstanceSnapshotDescription
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateInstanceSnapshotResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateInstanceSnapshot"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v1/instances/"+tea.StringValue(InstanceId)+"/snapshots"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteInstanceShutdownTimer(InstanceId *string) (_result *DeleteInstanceShutdownTimerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteInstanceShutdownTimerResponse{}
	_body, _err := client.DeleteInstanceShutdownTimerWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteInstanceShutdownTimerWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteInstanceShutdownTimerResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DeleteInstanceShutdownTimerResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteInstanceShutdownTimer"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/api/v1/instances/"+tea.StringValue(InstanceId)+"/shutdownTimer"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstanceShutdownTimer(InstanceId *string) (_result *GetInstanceShutdownTimerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceShutdownTimerResponse{}
	_body, _err := client.GetInstanceShutdownTimerWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstanceShutdownTimerWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetInstanceShutdownTimerResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetInstanceShutdownTimerResponse{}
	_body, _err := client.DoROARequest(tea.String("GetInstanceShutdownTimer"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/instances/"+tea.StringValue(InstanceId)+"/shutdownTimer"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopInstance(InstanceId *string) (_result *StopInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopInstanceResponse{}
	_body, _err := client.StopInstanceWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopInstanceWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopInstanceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &StopInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("StopInstance"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/api/v1/instances/"+tea.StringValue(InstanceId)+"/stop"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateInstance(InstanceId *string, request *UpdateInstanceRequest) (_result *UpdateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateInstanceResponse{}
	_body, _err := client.UpdateInstanceWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateInstanceWithOptions(InstanceId *string, request *UpdateInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &UpdateInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateInstance"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/api/v1/instances/"+tea.StringValue(InstanceId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEcsSpecs(request *ListEcsSpecsRequest) (_result *ListEcsSpecsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEcsSpecsResponse{}
	_body, _err := client.ListEcsSpecsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEcsSpecsWithOptions(request *ListEcsSpecsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListEcsSpecsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorTypeEquals)) {
		query["AcceleratorTypeEquals"] = request.AcceleratorTypeEquals
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListEcsSpecsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListEcsSpecs"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/ecsSpecs"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstancesStatistics(request *GetInstancesStatisticsRequest) (_result *GetInstancesStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstancesStatisticsResponse{}
	_body, _err := client.GetInstancesStatisticsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstancesStatisticsWithOptions(request *GetInstancesStatisticsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetInstancesStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceIds)) {
		query["WorkspaceIds"] = request.WorkspaceIds
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetInstancesStatisticsResponse{}
	_body, _err := client.DoROARequest(tea.String("GetInstancesStatistics"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/statistics/instances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteInstanceSnapshot(InstanceId *string, InstanceSnapshotId *string) (_result *DeleteInstanceSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteInstanceSnapshotResponse{}
	_body, _err := client.DeleteInstanceSnapshotWithOptions(InstanceId, InstanceSnapshotId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteInstanceSnapshotWithOptions(InstanceId *string, InstanceSnapshotId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteInstanceSnapshotResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DeleteInstanceSnapshotResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteInstanceSnapshot"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/api/v1/instances/"+tea.StringValue(InstanceId)+"/snapshots/"+tea.StringValue(InstanceSnapshotId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateInstanceSnapshot(InstanceId *string, InstanceSnapshotId *string, request *UpdateInstanceSnapshotRequest) (_result *UpdateInstanceSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateInstanceSnapshotResponse{}
	_body, _err := client.UpdateInstanceSnapshotWithOptions(InstanceId, InstanceSnapshotId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateInstanceSnapshotWithOptions(InstanceId *string, InstanceSnapshotId *string, request *UpdateInstanceSnapshotRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateInstanceSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceSnapshotName)) {
		body["InstanceSnapshotName"] = request.InstanceSnapshotName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &UpdateInstanceSnapshotResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateInstanceSnapshot"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/api/v1/instances/"+tea.StringValue(InstanceId)+"/snapshots/"+tea.StringValue(InstanceSnapshotId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateInstance(request *CreateInstanceRequest) (_result *CreateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateInstanceResponse{}
	_body, _err := client.CreateInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.EcsSpec)) {
		body["EcsSpec"] = request.EcsSpec
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		body["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.NasFileSystemId)) {
		body["NasFileSystemId"] = request.NasFileSystemId
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.UserVpc))) {
		body["UserVpc"] = request.UserVpc
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["ImageUrl"] = request.ImageUrl
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateInstance"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v1/instances/"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAuthorization() (_result *GetAuthorizationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAuthorizationResponse{}
	_body, _err := client.GetAuthorizationWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAuthorizationWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAuthorizationResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetAuthorizationResponse{}
	_body, _err := client.DoROARequest(tea.String("GetAuthorization"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/authorization"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListImages(request *ListImagesRequest) (_result *ListImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListImagesResponse{}
	_body, _err := client.ListImagesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListImagesWithOptions(request *ListImagesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorTypeEquals)) {
		query["AcceleratorTypeEquals"] = request.AcceleratorTypeEquals
	}

	if !tea.BoolValue(util.IsUnset(request.ImageTypeEquals)) {
		query["ImageTypeEquals"] = request.ImageTypeEquals
	}

	if !tea.BoolValue(util.IsUnset(request.ImageNameContains)) {
		query["ImageNameContains"] = request.ImageNameContains
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListImagesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListImages"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/images"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Foobar() (_result *FoobarResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &FoobarResponse{}
	_body, _err := client.FoobarWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FoobarWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *FoobarResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &FoobarResponse{}
	_body, _err := client.DoROARequest(tea.String("Foobar"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/foobar"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstanceSnapshots(InstanceId *string) (_result *ListInstanceSnapshotsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstanceSnapshotsResponse{}
	_body, _err := client.ListInstanceSnapshotsWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstanceSnapshotsWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstanceSnapshotsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ListInstanceSnapshotsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListInstanceSnapshots"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/instances/"+tea.StringValue(InstanceId)+"/snapshots"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateInstanceShutdownTimer(InstanceId *string, request *CreateInstanceShutdownTimerRequest) (_result *CreateInstanceShutdownTimerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateInstanceShutdownTimerResponse{}
	_body, _err := client.CreateInstanceShutdownTimerWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateInstanceShutdownTimerWithOptions(InstanceId *string, request *CreateInstanceShutdownTimerRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateInstanceShutdownTimerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ScheduleTime)) {
		body["ScheduleTime"] = request.ScheduleTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateInstanceShutdownTimerResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateInstanceShutdownTimer"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v1/instances/"+tea.StringValue(InstanceId)+"/shutdownTimer"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstance(InstanceId *string) (_result *GetInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceResponse{}
	_body, _err := client.GetInstanceWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstanceWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("GetInstance"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/instances/"+tea.StringValue(InstanceId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstances(request *ListInstancesRequest) (_result *ListInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstancesResponse{}
	_body, _err := client.ListInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstancesWithOptions(request *ListInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceNameContains)) {
		query["InstanceNameContains"] = request.InstanceNameContains
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceStatusEquals)) {
		query["InstanceStatusEquals"] = request.InstanceStatusEquals
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrder)) {
		query["SortOrder"] = request.SortOrder
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListInstancesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListInstances"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/instances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstanceSnapshot(InstanceId *string, InstanceSnapshotId *string) (_result *GetInstanceSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceSnapshotResponse{}
	_body, _err := client.GetInstanceSnapshotWithOptions(InstanceId, InstanceSnapshotId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstanceSnapshotWithOptions(InstanceId *string, InstanceSnapshotId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetInstanceSnapshotResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetInstanceSnapshotResponse{}
	_body, _err := client.DoROARequest(tea.String("GetInstanceSnapshot"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/instances/"+tea.StringValue(InstanceId)+"/snapshots/"+tea.StringValue(InstanceSnapshotId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartInstance(InstanceId *string) (_result *StartInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartInstanceResponse{}
	_body, _err := client.StartInstanceWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartInstanceWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartInstanceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &StartInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("StartInstance"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/api/v1/instances/"+tea.StringValue(InstanceId)+"/start"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteInstance(InstanceId *string) (_result *DeleteInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DeleteInstanceWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteInstanceWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteInstance"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/api/v1/instances/"+tea.StringValue(InstanceId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
