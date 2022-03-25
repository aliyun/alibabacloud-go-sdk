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

type Cluster struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 集群类型
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// 集群配置
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// 集群名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// nas文件id
	Nas *string `json:"Nas,omitempty" xml:"Nas,omitempty"`
	// 集群owner
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// 个人nas挂载到容器路径
	PrivateContainerPath *string `json:"PrivateContainerPath,omitempty" xml:"PrivateContainerPath,omitempty"`
	// 个人nas挂载路径
	PrivateNasPath *string `json:"PrivateNasPath,omitempty" xml:"PrivateNasPath,omitempty"`
	// 公共nas挂载到容器的路径
	PublicContainerPath *string `json:"PublicContainerPath,omitempty" xml:"PublicContainerPath,omitempty"`
	// 公共nas挂载路径
	PublicNasPath *string `json:"PublicNasPath,omitempty" xml:"PublicNasPath,omitempty"`
	// 集群状态
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// 集群vSwitch
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// 集群vpc
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s Cluster) String() string {
	return tea.Prettify(s)
}

func (s Cluster) GoString() string {
	return s.String()
}

func (s *Cluster) SetClusterId(v string) *Cluster {
	s.ClusterId = &v
	return s
}

func (s *Cluster) SetClusterType(v string) *Cluster {
	s.ClusterType = &v
	return s
}

func (s *Cluster) SetConfig(v string) *Cluster {
	s.Config = &v
	return s
}

func (s *Cluster) SetName(v string) *Cluster {
	s.Name = &v
	return s
}

func (s *Cluster) SetNas(v string) *Cluster {
	s.Nas = &v
	return s
}

func (s *Cluster) SetOwner(v string) *Cluster {
	s.Owner = &v
	return s
}

func (s *Cluster) SetPrivateContainerPath(v string) *Cluster {
	s.PrivateContainerPath = &v
	return s
}

func (s *Cluster) SetPrivateNasPath(v string) *Cluster {
	s.PrivateNasPath = &v
	return s
}

func (s *Cluster) SetPublicContainerPath(v string) *Cluster {
	s.PublicContainerPath = &v
	return s
}

func (s *Cluster) SetPublicNasPath(v string) *Cluster {
	s.PublicNasPath = &v
	return s
}

func (s *Cluster) SetState(v string) *Cluster {
	s.State = &v
	return s
}

func (s *Cluster) SetVSwitchId(v string) *Cluster {
	s.VSwitchId = &v
	return s
}

func (s *Cluster) SetVpcId(v string) *Cluster {
	s.VpcId = &v
	return s
}

type Config struct {
	// 配置名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 配置数值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s Config) String() string {
	return tea.Prettify(s)
}

func (s Config) GoString() string {
	return s.String()
}

func (s *Config) SetName(v string) *Config {
	s.Name = &v
	return s
}

func (s *Config) SetValue(v string) *Config {
	s.Value = &v
	return s
}

type Dataset struct {
	// 数据集id
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// 默认挂载路径
	DefaultMountPath *string `json:"DefaultMountPath,omitempty" xml:"DefaultMountPath,omitempty"`
	// 文件系统Id
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// 挂载路径
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// 数据集名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 文件系统被挂载路径
	NasPath *string `json:"NasPath,omitempty" xml:"NasPath,omitempty"`
}

func (s Dataset) String() string {
	return tea.Prettify(s)
}

func (s Dataset) GoString() string {
	return s.String()
}

func (s *Dataset) SetDatasetId(v string) *Dataset {
	s.DatasetId = &v
	return s
}

func (s *Dataset) SetDefaultMountPath(v string) *Dataset {
	s.DefaultMountPath = &v
	return s
}

func (s *Dataset) SetFileSystemId(v string) *Dataset {
	s.FileSystemId = &v
	return s
}

func (s *Dataset) SetMountPath(v string) *Dataset {
	s.MountPath = &v
	return s
}

func (s *Dataset) SetName(v string) *Dataset {
	s.Name = &v
	return s
}

func (s *Dataset) SetNasPath(v string) *Dataset {
	s.NasPath = &v
	return s
}

type EcsSpec struct {
	// cpu数量
	Cpu *int64 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// gpu卡数
	Gpu *int64 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// GPU卡类型
	GpuType *string `json:"GpuType,omitempty" xml:"GpuType,omitempty"`
	// 实例类型
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// 内存(GiB)
	MemoryInGiB *int64 `json:"MemoryInGiB,omitempty" xml:"MemoryInGiB,omitempty"`
	// 磁盘类型
	SystemDiskCategory *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	// 磁盘大小(GiB)
	SystemDiskSizeInGiB *int64 `json:"SystemDiskSizeInGiB,omitempty" xml:"SystemDiskSizeInGiB,omitempty"`
}

func (s EcsSpec) String() string {
	return tea.Prettify(s)
}

func (s EcsSpec) GoString() string {
	return s.String()
}

func (s *EcsSpec) SetCpu(v int64) *EcsSpec {
	s.Cpu = &v
	return s
}

func (s *EcsSpec) SetGpu(v int64) *EcsSpec {
	s.Gpu = &v
	return s
}

func (s *EcsSpec) SetGpuType(v string) *EcsSpec {
	s.GpuType = &v
	return s
}

func (s *EcsSpec) SetInstanceType(v string) *EcsSpec {
	s.InstanceType = &v
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

type Image struct {
	// 资源类型
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	// 镜像作者
	Author *string `json:"Author,omitempty" xml:"Author,omitempty"`
	// Cuda版本
	CudaVersion *string `json:"CudaVersion,omitempty" xml:"CudaVersion,omitempty"`
	// 镜像描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Ecs镜像key
	EcsImageKey *string `json:"EcsImageKey,omitempty" xml:"EcsImageKey,omitempty"`
	// 算法框架
	Framework *string `json:"Framework,omitempty" xml:"Framework,omitempty"`
	// 算法框架版本
	FrameworkVersion *string `json:"FrameworkVersion,omitempty" xml:"FrameworkVersion,omitempty"`
	// 镜像父镜像
	FromImageId *string `json:"FromImageId,omitempty" xml:"FromImageId,omitempty"`
	// 镜像名称
	FromImageName *string `json:"FromImageName,omitempty" xml:"FromImageName,omitempty"`
	// 创建时间
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 修改时间
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// 镜像ID
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// 镜像url
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 镜像命名空间
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// 镜像操作系统分发版
	OS *string `json:"OS,omitempty" xml:"OS,omitempty"`
	// 分发版版本
	OSVersion *string `json:"OSVersion,omitempty" xml:"OSVersion,omitempty"`
	// python版本
	PythonVersion *string `json:"PythonVersion,omitempty" xml:"PythonVersion,omitempty"`
	// 地区
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// 镜像仓库
	Repository *string `json:"Repository,omitempty" xml:"Repository,omitempty"`
	// 跳转的镜像站点页面
	RepositoryPage *string `json:"RepositoryPage,omitempty" xml:"RepositoryPage,omitempty"`
	// 资源类型
	ResourceType *int64 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// 镜像的根镜像
	RootImageId *string `json:"RootImageId,omitempty" xml:"RootImageId,omitempty"`
	// 镜像是否被其他实例共享
	Shared *bool `json:"Shared,omitempty" xml:"Shared,omitempty"`
	// 镜像短url
	ShortImageUrl *string `json:"ShortImageUrl,omitempty" xml:"ShortImageUrl,omitempty"`
	// 镜像仓库短名称
	ShortRepository *string `json:"ShortRepository,omitempty" xml:"ShortRepository,omitempty"`
	// 镜像状态
	Stage *string `json:"Stage,omitempty" xml:"Stage,omitempty"`
	// 镜像状态代码
	StageCode *int64 `json:"StageCode,omitempty" xml:"StageCode,omitempty"`
	// 保存镜像建议的名称
	SuggestedName *string `json:"SuggestedName,omitempty" xml:"SuggestedName,omitempty"`
	// Tag
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// 镜像类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 工作空间镜像id
	WorkspaceImageId *string `json:"WorkspaceImageId,omitempty" xml:"WorkspaceImageId,omitempty"`
}

func (s Image) String() string {
	return tea.Prettify(s)
}

func (s Image) GoString() string {
	return s.String()
}

func (s *Image) SetAcceleratorType(v string) *Image {
	s.AcceleratorType = &v
	return s
}

func (s *Image) SetAuthor(v string) *Image {
	s.Author = &v
	return s
}

func (s *Image) SetCudaVersion(v string) *Image {
	s.CudaVersion = &v
	return s
}

func (s *Image) SetDescription(v string) *Image {
	s.Description = &v
	return s
}

func (s *Image) SetEcsImageKey(v string) *Image {
	s.EcsImageKey = &v
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

func (s *Image) SetFromImageId(v string) *Image {
	s.FromImageId = &v
	return s
}

func (s *Image) SetFromImageName(v string) *Image {
	s.FromImageName = &v
	return s
}

func (s *Image) SetGmtCreateTime(v string) *Image {
	s.GmtCreateTime = &v
	return s
}

func (s *Image) SetGmtModifiedTime(v string) *Image {
	s.GmtModifiedTime = &v
	return s
}

func (s *Image) SetImageId(v string) *Image {
	s.ImageId = &v
	return s
}

func (s *Image) SetImageName(v string) *Image {
	s.ImageName = &v
	return s
}

func (s *Image) SetImageUrl(v string) *Image {
	s.ImageUrl = &v
	return s
}

func (s *Image) SetInstanceId(v string) *Image {
	s.InstanceId = &v
	return s
}

func (s *Image) SetNamespace(v string) *Image {
	s.Namespace = &v
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

func (s *Image) SetPythonVersion(v string) *Image {
	s.PythonVersion = &v
	return s
}

func (s *Image) SetRegion(v string) *Image {
	s.Region = &v
	return s
}

func (s *Image) SetRepository(v string) *Image {
	s.Repository = &v
	return s
}

func (s *Image) SetRepositoryPage(v string) *Image {
	s.RepositoryPage = &v
	return s
}

func (s *Image) SetResourceType(v int64) *Image {
	s.ResourceType = &v
	return s
}

func (s *Image) SetRootImageId(v string) *Image {
	s.RootImageId = &v
	return s
}

func (s *Image) SetShared(v bool) *Image {
	s.Shared = &v
	return s
}

func (s *Image) SetShortImageUrl(v string) *Image {
	s.ShortImageUrl = &v
	return s
}

func (s *Image) SetShortRepository(v string) *Image {
	s.ShortRepository = &v
	return s
}

func (s *Image) SetStage(v string) *Image {
	s.Stage = &v
	return s
}

func (s *Image) SetStageCode(v int64) *Image {
	s.StageCode = &v
	return s
}

func (s *Image) SetSuggestedName(v string) *Image {
	s.SuggestedName = &v
	return s
}

func (s *Image) SetTag(v string) *Image {
	s.Tag = &v
	return s
}

func (s *Image) SetType(v string) *Image {
	s.Type = &v
	return s
}

func (s *Image) SetWorkspaceImageId(v string) *Image {
	s.WorkspaceImageId = &v
	return s
}

type ImageNamespace struct {
	// 命名空间名称
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// 命名空间状态
	NamespaceStatus *string `json:"NamespaceStatus,omitempty" xml:"NamespaceStatus,omitempty"`
}

func (s ImageNamespace) String() string {
	return tea.Prettify(s)
}

func (s ImageNamespace) GoString() string {
	return s.String()
}

func (s *ImageNamespace) SetNamespace(v string) *ImageNamespace {
	s.Namespace = &v
	return s
}

func (s *ImageNamespace) SetNamespaceStatus(v string) *ImageNamespace {
	s.NamespaceStatus = &v
	return s
}

type ImageRepository struct {
	// 创建时间
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 修改时间
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// 地区Id
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// 仓库命名空间
	RepoNamespace *string `json:"RepoNamespace,omitempty" xml:"RepoNamespace,omitempty"`
	// 仓库状态
	RepoStatus *string `json:"RepoStatus,omitempty" xml:"RepoStatus,omitempty"`
	// 仓库地址
	Repository *string `json:"Repository,omitempty" xml:"Repository,omitempty"`
}

func (s ImageRepository) String() string {
	return tea.Prettify(s)
}

func (s ImageRepository) GoString() string {
	return s.String()
}

func (s *ImageRepository) SetGmtCreateTime(v string) *ImageRepository {
	s.GmtCreateTime = &v
	return s
}

func (s *ImageRepository) SetGmtModifiedTime(v string) *ImageRepository {
	s.GmtModifiedTime = &v
	return s
}

func (s *ImageRepository) SetRegionId(v string) *ImageRepository {
	s.RegionId = &v
	return s
}

func (s *ImageRepository) SetRepoName(v string) *ImageRepository {
	s.RepoName = &v
	return s
}

func (s *ImageRepository) SetRepoNamespace(v string) *ImageRepository {
	s.RepoNamespace = &v
	return s
}

func (s *ImageRepository) SetRepoStatus(v string) *ImageRepository {
	s.RepoStatus = &v
	return s
}

func (s *ImageRepository) SetRepository(v string) *ImageRepository {
	s.Repository = &v
	return s
}

type Instance struct {
	// 运行时间，毫秒数
	AccumulativeRunningTimeInMillis *int64 `json:"AccumulativeRunningTimeInMillis,omitempty" xml:"AccumulativeRunningTimeInMillis,omitempty"`
	// 累计运行时间(分钟)
	AccumulativeRunningTimeInMinutes *int64 `json:"AccumulativeRunningTimeInMinutes,omitempty" xml:"AccumulativeRunningTimeInMinutes,omitempty"`
	// 创建者
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// 数据集列表
	DatasetList []*Dataset `json:"DatasetList,omitempty" xml:"DatasetList,omitempty" type:"Repeated"`
	// ecs规格
	EcsSpec *string `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	// 创建时间(GMT)
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 修改时间(GMT)
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// 镜像ID
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// 镜像类型
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// 镜像链接
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// 定时关机任务
	InstanceShutdownTimer *InstanceShutdownTimer `json:"InstanceShutdownTimer,omitempty" xml:"InstanceShutdownTimer,omitempty"`
	// 实例状态
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// dsw实例链接
	InstanceUrl *string `json:"InstanceUrl,omitempty" xml:"InstanceUrl,omitempty"`
	// 实例版本
	InstanceVersion *string `json:"InstanceVersion,omitempty" xml:"InstanceVersion,omitempty"`
	// 是否他人可见
	IsPublic *int64 `json:"IsPublic,omitempty" xml:"IsPublic,omitempty"`
	// jupyter链接
	JupyterlabUrl *string `json:"JupyterlabUrl,omitempty" xml:"JupyterlabUrl,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// nas文件系统ID
	NasFileSystemId *string `json:"NasFileSystemId,omitempty" xml:"NasFileSystemId,omitempty"`
	// 付费类型代码
	PayType *int64 `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// 付费类型名称
	PayTypeName *string `json:"PayTypeName,omitempty" xml:"PayTypeName,omitempty"`
	// 资源类型名称
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// 资源类型代码
	ResourceType *int64 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// 是否支持定时关机
	ShutdownEnabled *bool `json:"ShutdownEnabled,omitempty" xml:"ShutdownEnabled,omitempty"`
	// 命令行终端链接
	TerminalUrl *string `json:"TerminalUrl,omitempty" xml:"TerminalUrl,omitempty"`
	// 用户ID
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 保存用户镜像列表
	UserImageList []*Image `json:"UserImageList,omitempty" xml:"UserImageList,omitempty" type:"Repeated"`
	// 被打通VPC配置
	UserVpc *UserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty"`
	// webIde链接
	WebIdeUrl *string `json:"WebIdeUrl,omitempty" xml:"WebIdeUrl,omitempty"`
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

func (s *Instance) SetAccumulativeRunningTimeInMillis(v int64) *Instance {
	s.AccumulativeRunningTimeInMillis = &v
	return s
}

func (s *Instance) SetAccumulativeRunningTimeInMinutes(v int64) *Instance {
	s.AccumulativeRunningTimeInMinutes = &v
	return s
}

func (s *Instance) SetCreateUser(v string) *Instance {
	s.CreateUser = &v
	return s
}

func (s *Instance) SetDatasetList(v []*Dataset) *Instance {
	s.DatasetList = v
	return s
}

func (s *Instance) SetEcsSpec(v string) *Instance {
	s.EcsSpec = &v
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

func (s *Instance) SetImageId(v string) *Instance {
	s.ImageId = &v
	return s
}

func (s *Instance) SetImageName(v string) *Instance {
	s.ImageName = &v
	return s
}

func (s *Instance) SetImageType(v string) *Instance {
	s.ImageType = &v
	return s
}

func (s *Instance) SetImageUrl(v string) *Instance {
	s.ImageUrl = &v
	return s
}

func (s *Instance) SetInstanceId(v string) *Instance {
	s.InstanceId = &v
	return s
}

func (s *Instance) SetInstanceName(v string) *Instance {
	s.InstanceName = &v
	return s
}

func (s *Instance) SetInstanceShutdownTimer(v *InstanceShutdownTimer) *Instance {
	s.InstanceShutdownTimer = v
	return s
}

func (s *Instance) SetInstanceStatus(v string) *Instance {
	s.InstanceStatus = &v
	return s
}

func (s *Instance) SetInstanceUrl(v string) *Instance {
	s.InstanceUrl = &v
	return s
}

func (s *Instance) SetInstanceVersion(v string) *Instance {
	s.InstanceVersion = &v
	return s
}

func (s *Instance) SetIsPublic(v int64) *Instance {
	s.IsPublic = &v
	return s
}

func (s *Instance) SetJupyterlabUrl(v string) *Instance {
	s.JupyterlabUrl = &v
	return s
}

func (s *Instance) SetMessage(v string) *Instance {
	s.Message = &v
	return s
}

func (s *Instance) SetNasFileSystemId(v string) *Instance {
	s.NasFileSystemId = &v
	return s
}

func (s *Instance) SetPayType(v int64) *Instance {
	s.PayType = &v
	return s
}

func (s *Instance) SetPayTypeName(v string) *Instance {
	s.PayTypeName = &v
	return s
}

func (s *Instance) SetResource(v string) *Instance {
	s.Resource = &v
	return s
}

func (s *Instance) SetResourceType(v int64) *Instance {
	s.ResourceType = &v
	return s
}

func (s *Instance) SetShutdownEnabled(v bool) *Instance {
	s.ShutdownEnabled = &v
	return s
}

func (s *Instance) SetTerminalUrl(v string) *Instance {
	s.TerminalUrl = &v
	return s
}

func (s *Instance) SetUserId(v string) *Instance {
	s.UserId = &v
	return s
}

func (s *Instance) SetUserImageList(v []*Image) *Instance {
	s.UserImageList = v
	return s
}

func (s *Instance) SetUserVpc(v *UserVpc) *Instance {
	s.UserVpc = v
	return s
}

func (s *Instance) SetWebIdeUrl(v string) *Instance {
	s.WebIdeUrl = &v
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

type InstanceShutdownTimer struct {
	// 定时关机修改时间
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 定时关机创建时间
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 定时关机时间
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	// 多少毫秒后定时关机（如果设定可以覆盖ScheduleTime）
	TtlInMillis *int64 `json:"TtlInMillis,omitempty" xml:"TtlInMillis,omitempty"`
}

func (s InstanceShutdownTimer) String() string {
	return tea.Prettify(s)
}

func (s InstanceShutdownTimer) GoString() string {
	return s.String()
}

func (s *InstanceShutdownTimer) SetGmtCreateTime(v string) *InstanceShutdownTimer {
	s.GmtCreateTime = &v
	return s
}

func (s *InstanceShutdownTimer) SetGmtModifiedTime(v string) *InstanceShutdownTimer {
	s.GmtModifiedTime = &v
	return s
}

func (s *InstanceShutdownTimer) SetInstanceId(v string) *InstanceShutdownTimer {
	s.InstanceId = &v
	return s
}

func (s *InstanceShutdownTimer) SetScheduleTime(v string) *InstanceShutdownTimer {
	s.ScheduleTime = &v
	return s
}

func (s *InstanceShutdownTimer) SetTtlInMillis(v int64) *InstanceShutdownTimer {
	s.TtlInMillis = &v
	return s
}

type InstanceSnapshot struct {
	// 实例快照保存时间（GMT）
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 实例快照修改时间（GMT）
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例快照描述
	InstanceSnapshotDescription *string `json:"InstanceSnapshotDescription,omitempty" xml:"InstanceSnapshotDescription,omitempty"`
	// 实例快照ID
	InstanceSnapshotId *string `json:"InstanceSnapshotId,omitempty" xml:"InstanceSnapshotId,omitempty"`
	// 实例快照名称
	InstanceSnapshotName *string `json:"InstanceSnapshotName,omitempty" xml:"InstanceSnapshotName,omitempty"`
	// 实例快照存储地址
	InstanceSnapshotRepoUrl *string `json:"InstanceSnapshotRepoUrl,omitempty" xml:"InstanceSnapshotRepoUrl,omitempty"`
	// 实例快照状态
	InstanceSnapshotStatus *string `json:"InstanceSnapshotStatus,omitempty" xml:"InstanceSnapshotStatus,omitempty"`
	// 实例快照标签
	InstanceSnapshotTag *string `json:"InstanceSnapshotTag,omitempty" xml:"InstanceSnapshotTag,omitempty"`
}

func (s InstanceSnapshot) String() string {
	return tea.Prettify(s)
}

func (s InstanceSnapshot) GoString() string {
	return s.String()
}

func (s *InstanceSnapshot) SetGmtCreateTime(v string) *InstanceSnapshot {
	s.GmtCreateTime = &v
	return s
}

func (s *InstanceSnapshot) SetGmtModifiedTime(v string) *InstanceSnapshot {
	s.GmtModifiedTime = &v
	return s
}

func (s *InstanceSnapshot) SetInstanceId(v string) *InstanceSnapshot {
	s.InstanceId = &v
	return s
}

func (s *InstanceSnapshot) SetInstanceSnapshotDescription(v string) *InstanceSnapshot {
	s.InstanceSnapshotDescription = &v
	return s
}

func (s *InstanceSnapshot) SetInstanceSnapshotId(v string) *InstanceSnapshot {
	s.InstanceSnapshotId = &v
	return s
}

func (s *InstanceSnapshot) SetInstanceSnapshotName(v string) *InstanceSnapshot {
	s.InstanceSnapshotName = &v
	return s
}

func (s *InstanceSnapshot) SetInstanceSnapshotRepoUrl(v string) *InstanceSnapshot {
	s.InstanceSnapshotRepoUrl = &v
	return s
}

func (s *InstanceSnapshot) SetInstanceSnapshotStatus(v string) *InstanceSnapshot {
	s.InstanceSnapshotStatus = &v
	return s
}

func (s *InstanceSnapshot) SetInstanceSnapshotTag(v string) *InstanceSnapshot {
	s.InstanceSnapshotTag = &v
	return s
}

type InstanceStatus struct {
	// 累计运行时间（分钟）
	AccumulativeRunningTimeInMinutes *int64 `json:"AccumulativeRunningTimeInMinutes,omitempty" xml:"AccumulativeRunningTimeInMinutes,omitempty"`
	// 实例ID
	InstanceId            *string                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceShutdownTimer *InstanceShutdownTimer `json:"InstanceShutdownTimer,omitempty" xml:"InstanceShutdownTimer,omitempty"`
	// 实例状态
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// 实例消息
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// 是否允许使用定时关机
	ShutdownEnabled *bool `json:"ShutdownEnabled,omitempty" xml:"ShutdownEnabled,omitempty"`
	// 实例类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s InstanceStatus) String() string {
	return tea.Prettify(s)
}

func (s InstanceStatus) GoString() string {
	return s.String()
}

func (s *InstanceStatus) SetAccumulativeRunningTimeInMinutes(v int64) *InstanceStatus {
	s.AccumulativeRunningTimeInMinutes = &v
	return s
}

func (s *InstanceStatus) SetInstanceId(v string) *InstanceStatus {
	s.InstanceId = &v
	return s
}

func (s *InstanceStatus) SetInstanceShutdownTimer(v *InstanceShutdownTimer) *InstanceStatus {
	s.InstanceShutdownTimer = v
	return s
}

func (s *InstanceStatus) SetInstanceStatus(v string) *InstanceStatus {
	s.InstanceStatus = &v
	return s
}

func (s *InstanceStatus) SetMsg(v string) *InstanceStatus {
	s.Msg = &v
	return s
}

func (s *InstanceStatus) SetShutdownEnabled(v bool) *InstanceStatus {
	s.ShutdownEnabled = &v
	return s
}

func (s *InstanceStatus) SetType(v string) *InstanceStatus {
	s.Type = &v
	return s
}

type InstanceType struct {
	// CPU核数
	CpuCoreCount *int64 `json:"CpuCoreCount,omitempty" xml:"CpuCoreCount,omitempty"`
	// 折扣
	Discount *float32 `json:"Discount,omitempty" xml:"Discount,omitempty"`
	// 内部价
	DomesticPrice *float32 `json:"DomesticPrice,omitempty" xml:"DomesticPrice,omitempty"`
	// GPU卡数
	GPUAmount *int64 `json:"GPUAmount,omitempty" xml:"GPUAmount,omitempty"`
	// GPU规格
	GPUSpec *string `json:"GPUSpec,omitempty" xml:"GPUSpec,omitempty"`
	// 实例接收带宽
	InstanceBandwidthRx *int64 `json:"InstanceBandwidthRx,omitempty" xml:"InstanceBandwidthRx,omitempty"`
	// 实例发送带宽
	InstanceBandwidthTx *int64 `json:"InstanceBandwidthTx,omitempty" xml:"InstanceBandwidthTx,omitempty"`
	// 实例每秒发包数量
	InstancePpsRx *int64 `json:"InstancePpsRx,omitempty" xml:"InstancePpsRx,omitempty"`
	// 实例每秒收包数量
	InstancePpsTx *int64 `json:"InstancePpsTx,omitempty" xml:"InstancePpsTx,omitempty"`
	// 实例规格族
	InstanceTypeFamily *string `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
	// 实例类型Id
	InstanceTypeId *string `json:"InstanceTypeId,omitempty" xml:"InstanceTypeId,omitempty"`
	// 是否国际站
	International *bool `json:"International,omitempty" xml:"International,omitempty"`
	// 本地磁盘容量
	LocalStorageCapacity *int64 `json:"LocalStorageCapacity,omitempty" xml:"LocalStorageCapacity,omitempty"`
	// 内存容量
	MemorySize *float32 `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	// 价格
	Price *float32 `json:"Price,omitempty" xml:"Price,omitempty"`
	// 价格（人民币）
	PriceCNY *float32 `json:"PriceCNY,omitempty" xml:"PriceCNY,omitempty"`
	// 价格（美元）
	PriceUSD *float32 `json:"PriceUSD,omitempty" xml:"PriceUSD,omitempty"`
	// 资源类型
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// 磁盘存储类型
	SystemDiskCategory *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	// 磁盘容量
	SystemDiskSize *int64 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s InstanceType) String() string {
	return tea.Prettify(s)
}

func (s InstanceType) GoString() string {
	return s.String()
}

func (s *InstanceType) SetCpuCoreCount(v int64) *InstanceType {
	s.CpuCoreCount = &v
	return s
}

func (s *InstanceType) SetDiscount(v float32) *InstanceType {
	s.Discount = &v
	return s
}

func (s *InstanceType) SetDomesticPrice(v float32) *InstanceType {
	s.DomesticPrice = &v
	return s
}

func (s *InstanceType) SetGPUAmount(v int64) *InstanceType {
	s.GPUAmount = &v
	return s
}

func (s *InstanceType) SetGPUSpec(v string) *InstanceType {
	s.GPUSpec = &v
	return s
}

func (s *InstanceType) SetInstanceBandwidthRx(v int64) *InstanceType {
	s.InstanceBandwidthRx = &v
	return s
}

func (s *InstanceType) SetInstanceBandwidthTx(v int64) *InstanceType {
	s.InstanceBandwidthTx = &v
	return s
}

func (s *InstanceType) SetInstancePpsRx(v int64) *InstanceType {
	s.InstancePpsRx = &v
	return s
}

func (s *InstanceType) SetInstancePpsTx(v int64) *InstanceType {
	s.InstancePpsTx = &v
	return s
}

func (s *InstanceType) SetInstanceTypeFamily(v string) *InstanceType {
	s.InstanceTypeFamily = &v
	return s
}

func (s *InstanceType) SetInstanceTypeId(v string) *InstanceType {
	s.InstanceTypeId = &v
	return s
}

func (s *InstanceType) SetInternational(v bool) *InstanceType {
	s.International = &v
	return s
}

func (s *InstanceType) SetLocalStorageCapacity(v int64) *InstanceType {
	s.LocalStorageCapacity = &v
	return s
}

func (s *InstanceType) SetMemorySize(v float32) *InstanceType {
	s.MemorySize = &v
	return s
}

func (s *InstanceType) SetPrice(v float32) *InstanceType {
	s.Price = &v
	return s
}

func (s *InstanceType) SetPriceCNY(v float32) *InstanceType {
	s.PriceCNY = &v
	return s
}

func (s *InstanceType) SetPriceUSD(v float32) *InstanceType {
	s.PriceUSD = &v
	return s
}

func (s *InstanceType) SetResourceType(v string) *InstanceType {
	s.ResourceType = &v
	return s
}

func (s *InstanceType) SetSystemDiskCategory(v string) *InstanceType {
	s.SystemDiskCategory = &v
	return s
}

func (s *InstanceType) SetSystemDiskSize(v int64) *InstanceType {
	s.SystemDiskSize = &v
	return s
}

type Nas struct {
	// Nas盘描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Nas文件系统Id
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// Nas盘状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s Nas) String() string {
	return tea.Prettify(s)
}

func (s Nas) GoString() string {
	return s.String()
}

func (s *Nas) SetDescription(v string) *Nas {
	s.Description = &v
	return s
}

func (s *Nas) SetFileSystemId(v string) *Nas {
	s.FileSystemId = &v
	return s
}

func (s *Nas) SetStatus(v string) *Nas {
	s.Status = &v
	return s
}

type Region struct {
	// 城市
	RegionCity *string `json:"RegionCity,omitempty" xml:"RegionCity,omitempty"`
	// id
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 名称
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// 州省
	RegionState *string `json:"RegionState,omitempty" xml:"RegionState,omitempty"`
	// 服务地址
	ServiceUrl *string `json:"ServiceUrl,omitempty" xml:"ServiceUrl,omitempty"`
}

func (s Region) String() string {
	return tea.Prettify(s)
}

func (s Region) GoString() string {
	return s.String()
}

func (s *Region) SetRegionCity(v string) *Region {
	s.RegionCity = &v
	return s
}

func (s *Region) SetRegionId(v string) *Region {
	s.RegionId = &v
	return s
}

func (s *Region) SetRegionName(v string) *Region {
	s.RegionName = &v
	return s
}

func (s *Region) SetRegionState(v string) *Region {
	s.RegionState = &v
	return s
}

func (s *Region) SetServiceUrl(v string) *Region {
	s.ServiceUrl = &v
	return s
}

type ResourceInfo struct {
	// 显卡类型
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 支付类型
	PayType *int64 `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// 资源类型
	ResourceType *int64 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ResourceInfo) String() string {
	return tea.Prettify(s)
}

func (s ResourceInfo) GoString() string {
	return s.String()
}

func (s *ResourceInfo) SetName(v string) *ResourceInfo {
	s.Name = &v
	return s
}

func (s *ResourceInfo) SetPayType(v int64) *ResourceInfo {
	s.PayType = &v
	return s
}

func (s *ResourceInfo) SetResourceType(v int64) *ResourceInfo {
	s.ResourceType = &v
	return s
}

type SecurityGroup struct {
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 安全组id
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// 名称
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
	// vpc id
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s SecurityGroup) String() string {
	return tea.Prettify(s)
}

func (s SecurityGroup) GoString() string {
	return s.String()
}

func (s *SecurityGroup) SetCreateTime(v string) *SecurityGroup {
	s.CreateTime = &v
	return s
}

func (s *SecurityGroup) SetDescription(v string) *SecurityGroup {
	s.Description = &v
	return s
}

func (s *SecurityGroup) SetSecurityGroupId(v string) *SecurityGroup {
	s.SecurityGroupId = &v
	return s
}

func (s *SecurityGroup) SetSecurityGroupName(v string) *SecurityGroup {
	s.SecurityGroupName = &v
	return s
}

func (s *SecurityGroup) SetVpcId(v string) *SecurityGroup {
	s.VpcId = &v
	return s
}

type Status struct {
	// 累计运行时间（分钟）
	AccumulativeRunningTimeInMinutes *int64 `json:"AccumulativeRunningTimeInMinutes,omitempty" xml:"AccumulativeRunningTimeInMinutes,omitempty"`
	// 实例ID
	InstanceId            *string                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceShutdownTimer *InstanceShutdownTimer `json:"InstanceShutdownTimer,omitempty" xml:"InstanceShutdownTimer,omitempty"`
	// 实例状态
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// 实例消息
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// 实例类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s Status) String() string {
	return tea.Prettify(s)
}

func (s Status) GoString() string {
	return s.String()
}

func (s *Status) SetAccumulativeRunningTimeInMinutes(v int64) *Status {
	s.AccumulativeRunningTimeInMinutes = &v
	return s
}

func (s *Status) SetInstanceId(v string) *Status {
	s.InstanceId = &v
	return s
}

func (s *Status) SetInstanceShutdownTimer(v *InstanceShutdownTimer) *Status {
	s.InstanceShutdownTimer = v
	return s
}

func (s *Status) SetInstanceStatus(v string) *Status {
	s.InstanceStatus = &v
	return s
}

func (s *Status) SetMsg(v string) *Status {
	s.Msg = &v
	return s
}

func (s *Status) SetType(v string) *Status {
	s.Type = &v
	return s
}

type UserVpc struct {
	// 角色标识码
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// 安全组ID
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// 虚拟网络ID
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// 虚拟交换机ID
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s UserVpc) String() string {
	return tea.Prettify(s)
}

func (s UserVpc) GoString() string {
	return s.String()
}

func (s *UserVpc) SetRoleArn(v string) *UserVpc {
	s.RoleArn = &v
	return s
}

func (s *UserVpc) SetSecurityGroupId(v string) *UserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *UserVpc) SetVpcId(v string) *UserVpc {
	s.VpcId = &v
	return s
}

func (s *UserVpc) SetVswitchId(v string) *UserVpc {
	s.VswitchId = &v
	return s
}

type VSwitch struct {
	// 可用ip数量
	AvailableIpAddressCount *int64 `json:"AvailableIpAddressCount,omitempty" xml:"AvailableIpAddressCount,omitempty"`
	// 子网
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 是否默认
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// 状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// VSwitch Id
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// 名称
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
	// vpc id
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// 可用区
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s VSwitch) String() string {
	return tea.Prettify(s)
}

func (s VSwitch) GoString() string {
	return s.String()
}

func (s *VSwitch) SetAvailableIpAddressCount(v int64) *VSwitch {
	s.AvailableIpAddressCount = &v
	return s
}

func (s *VSwitch) SetCidrBlock(v string) *VSwitch {
	s.CidrBlock = &v
	return s
}

func (s *VSwitch) SetCreateTime(v string) *VSwitch {
	s.CreateTime = &v
	return s
}

func (s *VSwitch) SetDescription(v string) *VSwitch {
	s.Description = &v
	return s
}

func (s *VSwitch) SetIsDefault(v bool) *VSwitch {
	s.IsDefault = &v
	return s
}

func (s *VSwitch) SetStatus(v string) *VSwitch {
	s.Status = &v
	return s
}

func (s *VSwitch) SetVSwitchId(v string) *VSwitch {
	s.VSwitchId = &v
	return s
}

func (s *VSwitch) SetVSwitchName(v string) *VSwitch {
	s.VSwitchName = &v
	return s
}

func (s *VSwitch) SetVpcId(v string) *VSwitch {
	s.VpcId = &v
	return s
}

func (s *VSwitch) SetZoneId(v string) *VSwitch {
	s.ZoneId = &v
	return s
}

type Vpc struct {
	// vpc子网
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// vpc描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 是否默认
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// vpc状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 路由id
	VRouterId *string `json:"VRouterId,omitempty" xml:"VRouterId,omitempty"`
	// vpc id
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// vpc名称
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s Vpc) String() string {
	return tea.Prettify(s)
}

func (s Vpc) GoString() string {
	return s.String()
}

func (s *Vpc) SetCidrBlock(v string) *Vpc {
	s.CidrBlock = &v
	return s
}

func (s *Vpc) SetCreateTime(v string) *Vpc {
	s.CreateTime = &v
	return s
}

func (s *Vpc) SetDescription(v string) *Vpc {
	s.Description = &v
	return s
}

func (s *Vpc) SetIsDefault(v bool) *Vpc {
	s.IsDefault = &v
	return s
}

func (s *Vpc) SetStatus(v string) *Vpc {
	s.Status = &v
	return s
}

func (s *Vpc) SetVRouterId(v string) *Vpc {
	s.VRouterId = &v
	return s
}

func (s *Vpc) SetVpcId(v string) *Vpc {
	s.VpcId = &v
	return s
}

func (s *Vpc) SetVpcName(v string) *Vpc {
	s.VpcName = &v
	return s
}

type CreateInstanceRequest struct {
	DatasetList []*Dataset `json:"DatasetList,omitempty" xml:"DatasetList,omitempty" type:"Repeated"`
	// 实例规格
	EcsSpec *string `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	// 环境参数
	Environments map[string]interface{} `json:"Environments,omitempty" xml:"Environments,omitempty"`
	// 镜像id
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// 镜像地址
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	IsPublic     *int64  `json:"IsPublic,omitempty" xml:"IsPublic,omitempty"`
	// nas文件系统id
	NasFileSystemId *string `json:"NasFileSystemId,omitempty" xml:"NasFileSystemId,omitempty"`
	// 实例的真实用户名称
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// 打通的vpc网络配置
	UserVpc     *UserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty"`
	WorkspaceId *string  `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) SetDatasetList(v []*Dataset) *CreateInstanceRequest {
	s.DatasetList = v
	return s
}

func (s *CreateInstanceRequest) SetEcsSpec(v string) *CreateInstanceRequest {
	s.EcsSpec = &v
	return s
}

func (s *CreateInstanceRequest) SetEnvironments(v map[string]interface{}) *CreateInstanceRequest {
	s.Environments = v
	return s
}

func (s *CreateInstanceRequest) SetImageId(v string) *CreateInstanceRequest {
	s.ImageId = &v
	return s
}

func (s *CreateInstanceRequest) SetImageUrl(v string) *CreateInstanceRequest {
	s.ImageUrl = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceName(v string) *CreateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceRequest) SetIsPublic(v int64) *CreateInstanceRequest {
	s.IsPublic = &v
	return s
}

func (s *CreateInstanceRequest) SetNasFileSystemId(v string) *CreateInstanceRequest {
	s.NasFileSystemId = &v
	return s
}

func (s *CreateInstanceRequest) SetUserName(v string) *CreateInstanceRequest {
	s.UserName = &v
	return s
}

func (s *CreateInstanceRequest) SetUserVpc(v *UserVpc) *CreateInstanceRequest {
	s.UserVpc = v
	return s
}

func (s *CreateInstanceRequest) SetWorkspaceId(v string) *CreateInstanceRequest {
	s.WorkspaceId = &v
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

type CreateInstanceShutdownTimerRequest struct {
	// 定时关机时间（GMT）
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	// 多少毫秒后定时关机（如果设定可以覆盖ScheduleTime）
	TtlInMillis *int64 `json:"TtlInMillis,omitempty" xml:"TtlInMillis,omitempty"`
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

func (s *CreateInstanceShutdownTimerRequest) SetTtlInMillis(v int64) *CreateInstanceShutdownTimerRequest {
	s.TtlInMillis = &v
	return s
}

type CreateInstanceShutdownTimerResponseBody struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInstanceShutdownTimerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceShutdownTimerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceShutdownTimerResponseBody) SetInstanceId(v string) *CreateInstanceShutdownTimerResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceShutdownTimerResponseBody) SetRequestId(v string) *CreateInstanceShutdownTimerResponseBody {
	s.RequestId = &v
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

type CreateInstanceShutdownTimerV2Request struct {
	// 定时关机设定时间
	DueTime *string `json:"DueTime,omitempty" xml:"DueTime,omitempty"`
	// 距离定时关机时间段
	RemainingTimeInMs *int64 `json:"RemainingTimeInMs,omitempty" xml:"RemainingTimeInMs,omitempty"`
}

func (s CreateInstanceShutdownTimerV2Request) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceShutdownTimerV2Request) GoString() string {
	return s.String()
}

func (s *CreateInstanceShutdownTimerV2Request) SetDueTime(v string) *CreateInstanceShutdownTimerV2Request {
	s.DueTime = &v
	return s
}

func (s *CreateInstanceShutdownTimerV2Request) SetRemainingTimeInMs(v int64) *CreateInstanceShutdownTimerV2Request {
	s.RemainingTimeInMs = &v
	return s
}

type CreateInstanceShutdownTimerV2ResponseBody struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInstanceShutdownTimerV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceShutdownTimerV2ResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceShutdownTimerV2ResponseBody) SetInstanceId(v string) *CreateInstanceShutdownTimerV2ResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceShutdownTimerV2ResponseBody) SetRequestId(v string) *CreateInstanceShutdownTimerV2ResponseBody {
	s.RequestId = &v
	return s
}

type CreateInstanceShutdownTimerV2Response struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateInstanceShutdownTimerV2ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateInstanceShutdownTimerV2Response) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceShutdownTimerV2Response) GoString() string {
	return s.String()
}

func (s *CreateInstanceShutdownTimerV2Response) SetHeaders(v map[string]*string) *CreateInstanceShutdownTimerV2Response {
	s.Headers = v
	return s
}

func (s *CreateInstanceShutdownTimerV2Response) SetBody(v *CreateInstanceShutdownTimerV2ResponseBody) *CreateInstanceShutdownTimerV2Response {
	s.Body = v
	return s
}

type CreateInstanceSnapshotRequest struct {
	// 实例快照描述
	InstanceSnapshotDescription *string `json:"InstanceSnapshotDescription,omitempty" xml:"InstanceSnapshotDescription,omitempty"`
	// 实例快照名称
	InstanceSnapshotName *string `json:"InstanceSnapshotName,omitempty" xml:"InstanceSnapshotName,omitempty"`
	// 实例快照存储地址（可选）
	InstanceSnapshotRepoUrl *string `json:"InstanceSnapshotRepoUrl,omitempty" xml:"InstanceSnapshotRepoUrl,omitempty"`
}

func (s CreateInstanceSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceSnapshotRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceSnapshotRequest) SetInstanceSnapshotDescription(v string) *CreateInstanceSnapshotRequest {
	s.InstanceSnapshotDescription = &v
	return s
}

func (s *CreateInstanceSnapshotRequest) SetInstanceSnapshotName(v string) *CreateInstanceSnapshotRequest {
	s.InstanceSnapshotName = &v
	return s
}

func (s *CreateInstanceSnapshotRequest) SetInstanceSnapshotRepoUrl(v string) *CreateInstanceSnapshotRequest {
	s.InstanceSnapshotRepoUrl = &v
	return s
}

type CreateInstanceSnapshotResponseBody struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例快照ID
	InstanceSnapshotId *string `json:"InstanceSnapshotId,omitempty" xml:"InstanceSnapshotId,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInstanceSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceSnapshotResponseBody) SetInstanceId(v string) *CreateInstanceSnapshotResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceSnapshotResponseBody) SetInstanceSnapshotId(v string) *CreateInstanceSnapshotResponseBody {
	s.InstanceSnapshotId = &v
	return s
}

func (s *CreateInstanceSnapshotResponseBody) SetRequestId(v string) *CreateInstanceSnapshotResponseBody {
	s.RequestId = &v
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

type CreateInstanceSnapshotV2Request struct {
	// 镜像地址
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// 实例快照描述
	SnapshotDescription *string `json:"SnapshotDescription,omitempty" xml:"SnapshotDescription,omitempty"`
	// 实例快照名称
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
}

func (s CreateInstanceSnapshotV2Request) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceSnapshotV2Request) GoString() string {
	return s.String()
}

func (s *CreateInstanceSnapshotV2Request) SetImageUrl(v string) *CreateInstanceSnapshotV2Request {
	s.ImageUrl = &v
	return s
}

func (s *CreateInstanceSnapshotV2Request) SetSnapshotDescription(v string) *CreateInstanceSnapshotV2Request {
	s.SnapshotDescription = &v
	return s
}

func (s *CreateInstanceSnapshotV2Request) SetSnapshotName(v string) *CreateInstanceSnapshotV2Request {
	s.SnapshotName = &v
	return s
}

type CreateInstanceSnapshotV2ResponseBody struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 实例快照Id
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s CreateInstanceSnapshotV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceSnapshotV2ResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceSnapshotV2ResponseBody) SetInstanceId(v string) *CreateInstanceSnapshotV2ResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceSnapshotV2ResponseBody) SetRequestId(v string) *CreateInstanceSnapshotV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceSnapshotV2ResponseBody) SetSnapshotId(v string) *CreateInstanceSnapshotV2ResponseBody {
	s.SnapshotId = &v
	return s
}

type CreateInstanceSnapshotV2Response struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateInstanceSnapshotV2ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateInstanceSnapshotV2Response) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceSnapshotV2Response) GoString() string {
	return s.String()
}

func (s *CreateInstanceSnapshotV2Response) SetHeaders(v map[string]*string) *CreateInstanceSnapshotV2Response {
	s.Headers = v
	return s
}

func (s *CreateInstanceSnapshotV2Response) SetBody(v *CreateInstanceSnapshotV2ResponseBody) *CreateInstanceSnapshotV2Response {
	s.Body = v
	return s
}

type CreateInstanceV2Request struct {
	// 工作空间内是否他人可见
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// 数据集集合
	Datasets []*CreateInstanceV2RequestDatasets `json:"Datasets,omitempty" xml:"Datasets,omitempty" type:"Repeated"`
	// 实例对应的Ecs规格
	EcsSpec *string `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	// 环境变量
	EnvironmentVariables map[string]*string `json:"EnvironmentVariables,omitempty" xml:"EnvironmentVariables,omitempty"`
	// 镜像Id
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// 镜像地址
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// user vpc配置
	UserVpc *CreateInstanceV2RequestUserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	// 工作空间Id
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateInstanceV2Request) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceV2Request) GoString() string {
	return s.String()
}

func (s *CreateInstanceV2Request) SetAccessibility(v string) *CreateInstanceV2Request {
	s.Accessibility = &v
	return s
}

func (s *CreateInstanceV2Request) SetDatasets(v []*CreateInstanceV2RequestDatasets) *CreateInstanceV2Request {
	s.Datasets = v
	return s
}

func (s *CreateInstanceV2Request) SetEcsSpec(v string) *CreateInstanceV2Request {
	s.EcsSpec = &v
	return s
}

func (s *CreateInstanceV2Request) SetEnvironmentVariables(v map[string]*string) *CreateInstanceV2Request {
	s.EnvironmentVariables = v
	return s
}

func (s *CreateInstanceV2Request) SetImageId(v string) *CreateInstanceV2Request {
	s.ImageId = &v
	return s
}

func (s *CreateInstanceV2Request) SetImageUrl(v string) *CreateInstanceV2Request {
	s.ImageUrl = &v
	return s
}

func (s *CreateInstanceV2Request) SetInstanceName(v string) *CreateInstanceV2Request {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceV2Request) SetUserVpc(v *CreateInstanceV2RequestUserVpc) *CreateInstanceV2Request {
	s.UserVpc = v
	return s
}

func (s *CreateInstanceV2Request) SetWorkspaceId(v string) *CreateInstanceV2Request {
	s.WorkspaceId = &v
	return s
}

type CreateInstanceV2RequestDatasets struct {
	// 数据集Id
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// 容器内挂载路径
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
}

func (s CreateInstanceV2RequestDatasets) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceV2RequestDatasets) GoString() string {
	return s.String()
}

func (s *CreateInstanceV2RequestDatasets) SetDatasetId(v string) *CreateInstanceV2RequestDatasets {
	s.DatasetId = &v
	return s
}

func (s *CreateInstanceV2RequestDatasets) SetMountPath(v string) *CreateInstanceV2RequestDatasets {
	s.MountPath = &v
	return s
}

type CreateInstanceV2RequestUserVpc struct {
	// Security Group Id
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// VSwitch Id
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// Vpc Id
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateInstanceV2RequestUserVpc) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceV2RequestUserVpc) GoString() string {
	return s.String()
}

func (s *CreateInstanceV2RequestUserVpc) SetSecurityGroupId(v string) *CreateInstanceV2RequestUserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateInstanceV2RequestUserVpc) SetVSwitchId(v string) *CreateInstanceV2RequestUserVpc {
	s.VSwitchId = &v
	return s
}

func (s *CreateInstanceV2RequestUserVpc) SetVpcId(v string) *CreateInstanceV2RequestUserVpc {
	s.VpcId = &v
	return s
}

type CreateInstanceV2ResponseBody struct {
	// 工作空间Id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInstanceV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceV2ResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceV2ResponseBody) SetInstanceId(v string) *CreateInstanceV2ResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceV2ResponseBody) SetRequestId(v string) *CreateInstanceV2ResponseBody {
	s.RequestId = &v
	return s
}

type CreateInstanceV2Response struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateInstanceV2ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateInstanceV2Response) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceV2Response) GoString() string {
	return s.String()
}

func (s *CreateInstanceV2Response) SetHeaders(v map[string]*string) *CreateInstanceV2Response {
	s.Headers = v
	return s
}

func (s *CreateInstanceV2Response) SetBody(v *CreateInstanceV2ResponseBody) *CreateInstanceV2Response {
	s.Body = v
	return s
}

type DeleteInstanceResponseBody struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponseBody) SetInstanceId(v string) *DeleteInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetRequestId(v string) *DeleteInstanceResponseBody {
	s.RequestId = &v
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

type DeleteInstanceShutdownTimerResponseBody struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInstanceShutdownTimerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceShutdownTimerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceShutdownTimerResponseBody) SetInstanceId(v string) *DeleteInstanceShutdownTimerResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DeleteInstanceShutdownTimerResponseBody) SetRequestId(v string) *DeleteInstanceShutdownTimerResponseBody {
	s.RequestId = &v
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

type DeleteInstanceShutdownTimerV2ResponseBody struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInstanceShutdownTimerV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceShutdownTimerV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceShutdownTimerV2ResponseBody) SetInstanceId(v string) *DeleteInstanceShutdownTimerV2ResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DeleteInstanceShutdownTimerV2ResponseBody) SetRequestId(v string) *DeleteInstanceShutdownTimerV2ResponseBody {
	s.RequestId = &v
	return s
}

type DeleteInstanceShutdownTimerV2Response struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteInstanceShutdownTimerV2ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteInstanceShutdownTimerV2Response) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceShutdownTimerV2Response) GoString() string {
	return s.String()
}

func (s *DeleteInstanceShutdownTimerV2Response) SetHeaders(v map[string]*string) *DeleteInstanceShutdownTimerV2Response {
	s.Headers = v
	return s
}

func (s *DeleteInstanceShutdownTimerV2Response) SetBody(v *DeleteInstanceShutdownTimerV2ResponseBody) *DeleteInstanceShutdownTimerV2Response {
	s.Body = v
	return s
}

type DeleteInstanceSnapshotResponseBody struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例快照ID
	InstanceSnapshotId *string `json:"InstanceSnapshotId,omitempty" xml:"InstanceSnapshotId,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInstanceSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceSnapshotResponseBody) SetInstanceId(v string) *DeleteInstanceSnapshotResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DeleteInstanceSnapshotResponseBody) SetInstanceSnapshotId(v string) *DeleteInstanceSnapshotResponseBody {
	s.InstanceSnapshotId = &v
	return s
}

func (s *DeleteInstanceSnapshotResponseBody) SetRequestId(v string) *DeleteInstanceSnapshotResponseBody {
	s.RequestId = &v
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

type DeleteInstanceSnapshotV2ResponseBody struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 实例快照Id
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s DeleteInstanceSnapshotV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceSnapshotV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceSnapshotV2ResponseBody) SetInstanceId(v string) *DeleteInstanceSnapshotV2ResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DeleteInstanceSnapshotV2ResponseBody) SetRequestId(v string) *DeleteInstanceSnapshotV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstanceSnapshotV2ResponseBody) SetSnapshotId(v string) *DeleteInstanceSnapshotV2ResponseBody {
	s.SnapshotId = &v
	return s
}

type DeleteInstanceSnapshotV2Response struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteInstanceSnapshotV2ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteInstanceSnapshotV2Response) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceSnapshotV2Response) GoString() string {
	return s.String()
}

func (s *DeleteInstanceSnapshotV2Response) SetHeaders(v map[string]*string) *DeleteInstanceSnapshotV2Response {
	s.Headers = v
	return s
}

func (s *DeleteInstanceSnapshotV2Response) SetBody(v *DeleteInstanceSnapshotV2ResponseBody) *DeleteInstanceSnapshotV2Response {
	s.Body = v
	return s
}

type DeleteInstanceV2ResponseBody struct {
	// 工作空间Id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInstanceV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceV2ResponseBody) SetInstanceId(v string) *DeleteInstanceV2ResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DeleteInstanceV2ResponseBody) SetRequestId(v string) *DeleteInstanceV2ResponseBody {
	s.RequestId = &v
	return s
}

type DeleteInstanceV2Response struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteInstanceV2ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteInstanceV2Response) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceV2Response) GoString() string {
	return s.String()
}

func (s *DeleteInstanceV2Response) SetHeaders(v map[string]*string) *DeleteInstanceV2Response {
	s.Headers = v
	return s
}

func (s *DeleteInstanceV2Response) SetBody(v *DeleteInstanceV2ResponseBody) *DeleteInstanceV2Response {
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

type Foobar1Response struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    interface{}        `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s Foobar1Response) String() string {
	return tea.Prettify(s)
}

func (s Foobar1Response) GoString() string {
	return s.String()
}

func (s *Foobar1Response) SetHeaders(v map[string]*string) *Foobar1Response {
	s.Headers = v
	return s
}

func (s *Foobar1Response) SetBody(v interface{}) *Foobar1Response {
	s.Body = v
	return s
}

type GetAuthorizationResponseBody struct {
	// 授权失败错误代码
	AuthorizationFailedCode *string `json:"AuthorizationFailedCode,omitempty" xml:"AuthorizationFailedCode,omitempty"`
	// 授权失败错误消息
	AuthorizationFailedMessage *string `json:"AuthorizationFailedMessage,omitempty" xml:"AuthorizationFailedMessage,omitempty"`
	// 是否已经给DSW服务账号授权
	Authorized *bool `json:"Authorized,omitempty" xml:"Authorized,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAuthorizationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuthorizationResponseBody) SetAuthorizationFailedCode(v string) *GetAuthorizationResponseBody {
	s.AuthorizationFailedCode = &v
	return s
}

func (s *GetAuthorizationResponseBody) SetAuthorizationFailedMessage(v string) *GetAuthorizationResponseBody {
	s.AuthorizationFailedMessage = &v
	return s
}

func (s *GetAuthorizationResponseBody) SetAuthorized(v bool) *GetAuthorizationResponseBody {
	s.Authorized = &v
	return s
}

func (s *GetAuthorizationResponseBody) SetRequestId(v string) *GetAuthorizationResponseBody {
	s.RequestId = &v
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

type GetDashboardStatisticsRequest struct {
	// 工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetDashboardStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDashboardStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetDashboardStatisticsRequest) SetWorkspaceId(v string) *GetDashboardStatisticsRequest {
	s.WorkspaceId = &v
	return s
}

type GetDashboardStatisticsResponseBody struct {
	// 实例数
	InstanceTotal *int64 `json:"InstanceTotal,omitempty" xml:"InstanceTotal,omitempty"`
	// 运行实例数
	InstsanceRunningTotal *int64 `json:"InstsanceRunningTotal,omitempty" xml:"InstsanceRunningTotal,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDashboardStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDashboardStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetDashboardStatisticsResponseBody) SetInstanceTotal(v int64) *GetDashboardStatisticsResponseBody {
	s.InstanceTotal = &v
	return s
}

func (s *GetDashboardStatisticsResponseBody) SetInstsanceRunningTotal(v int64) *GetDashboardStatisticsResponseBody {
	s.InstsanceRunningTotal = &v
	return s
}

func (s *GetDashboardStatisticsResponseBody) SetRequestId(v string) *GetDashboardStatisticsResponseBody {
	s.RequestId = &v
	return s
}

type GetDashboardStatisticsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDashboardStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDashboardStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDashboardStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetDashboardStatisticsResponse) SetHeaders(v map[string]*string) *GetDashboardStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetDashboardStatisticsResponse) SetBody(v *GetDashboardStatisticsResponseBody) *GetDashboardStatisticsResponse {
	s.Body = v
	return s
}

type GetInstanceResponseBody struct {
	// 累计运行时间(分钟)
	AccumulativeRunningTimeInMinutes *int64 `json:"AccumulativeRunningTimeInMinutes,omitempty" xml:"AccumulativeRunningTimeInMinutes,omitempty"`
	// ecs规格
	EcsSpec *string `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	// 实例创建时间(GMT)
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 实例修改时间(GMT)
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// 镜像ID
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// 镜像链接
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// 定时关机任务
	InstanceShutdownTimer *InstanceShutdownTimer `json:"InstanceShutdownTimer,omitempty" xml:"InstanceShutdownTimer,omitempty"`
	// 实例状态
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// jupyter链接
	JupyterlabUrl *string `json:"JupyterlabUrl,omitempty" xml:"JupyterlabUrl,omitempty"`
	// nas文件系统ID
	NasFileSystemId *string `json:"NasFileSystemId,omitempty" xml:"NasFileSystemId,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 命令行终端链接
	TerminalUrl *string `json:"TerminalUrl,omitempty" xml:"TerminalUrl,omitempty"`
	// 用户ID
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 被打通VPC配置
	UserVpc *UserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty"`
	// web ide链接
	WebIdeUrl *string `json:"WebIdeUrl,omitempty" xml:"WebIdeUrl,omitempty"`
}

func (s GetInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBody) SetAccumulativeRunningTimeInMinutes(v int64) *GetInstanceResponseBody {
	s.AccumulativeRunningTimeInMinutes = &v
	return s
}

func (s *GetInstanceResponseBody) SetEcsSpec(v string) *GetInstanceResponseBody {
	s.EcsSpec = &v
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

func (s *GetInstanceResponseBody) SetImageId(v string) *GetInstanceResponseBody {
	s.ImageId = &v
	return s
}

func (s *GetInstanceResponseBody) SetImageUrl(v string) *GetInstanceResponseBody {
	s.ImageUrl = &v
	return s
}

func (s *GetInstanceResponseBody) SetInstanceId(v string) *GetInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceResponseBody) SetInstanceName(v string) *GetInstanceResponseBody {
	s.InstanceName = &v
	return s
}

func (s *GetInstanceResponseBody) SetInstanceShutdownTimer(v *InstanceShutdownTimer) *GetInstanceResponseBody {
	s.InstanceShutdownTimer = v
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

func (s *GetInstanceResponseBody) SetNasFileSystemId(v string) *GetInstanceResponseBody {
	s.NasFileSystemId = &v
	return s
}

func (s *GetInstanceResponseBody) SetRequestId(v string) *GetInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceResponseBody) SetTerminalUrl(v string) *GetInstanceResponseBody {
	s.TerminalUrl = &v
	return s
}

func (s *GetInstanceResponseBody) SetUserId(v string) *GetInstanceResponseBody {
	s.UserId = &v
	return s
}

func (s *GetInstanceResponseBody) SetUserVpc(v *UserVpc) *GetInstanceResponseBody {
	s.UserVpc = v
	return s
}

func (s *GetInstanceResponseBody) SetWebIdeUrl(v string) *GetInstanceResponseBody {
	s.WebIdeUrl = &v
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

type GetInstanceShutdownTimerResponseBody struct {
	// 任务创建时间(GMT)
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 任务修改时间(GMT)
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 定时关机时间(GMT)
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
}

func (s GetInstanceShutdownTimerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceShutdownTimerResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceShutdownTimerResponseBody) SetGmtCreateTime(v string) *GetInstanceShutdownTimerResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetInstanceShutdownTimerResponseBody) SetGmtModifiedTime(v string) *GetInstanceShutdownTimerResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetInstanceShutdownTimerResponseBody) SetInstanceId(v string) *GetInstanceShutdownTimerResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceShutdownTimerResponseBody) SetRequestId(v string) *GetInstanceShutdownTimerResponseBody {
	s.RequestId = &v
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

type GetInstanceShutdownTimerV2ResponseBody struct {
	// 设定关机时间
	DueTime *string `json:"DueTime,omitempty" xml:"DueTime,omitempty"`
	// 创建时间
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 修改时间
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 剩余关机时间（ms）
	RemainingTimeInMs *int64 `json:"RemainingTimeInMs,omitempty" xml:"RemainingTimeInMs,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInstanceShutdownTimerV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceShutdownTimerV2ResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceShutdownTimerV2ResponseBody) SetDueTime(v string) *GetInstanceShutdownTimerV2ResponseBody {
	s.DueTime = &v
	return s
}

func (s *GetInstanceShutdownTimerV2ResponseBody) SetGmtCreateTime(v string) *GetInstanceShutdownTimerV2ResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetInstanceShutdownTimerV2ResponseBody) SetGmtModifiedTime(v string) *GetInstanceShutdownTimerV2ResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetInstanceShutdownTimerV2ResponseBody) SetInstanceId(v string) *GetInstanceShutdownTimerV2ResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceShutdownTimerV2ResponseBody) SetRemainingTimeInMs(v int64) *GetInstanceShutdownTimerV2ResponseBody {
	s.RemainingTimeInMs = &v
	return s
}

func (s *GetInstanceShutdownTimerV2ResponseBody) SetRequestId(v string) *GetInstanceShutdownTimerV2ResponseBody {
	s.RequestId = &v
	return s
}

type GetInstanceShutdownTimerV2Response struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetInstanceShutdownTimerV2ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInstanceShutdownTimerV2Response) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceShutdownTimerV2Response) GoString() string {
	return s.String()
}

func (s *GetInstanceShutdownTimerV2Response) SetHeaders(v map[string]*string) *GetInstanceShutdownTimerV2Response {
	s.Headers = v
	return s
}

func (s *GetInstanceShutdownTimerV2Response) SetBody(v *GetInstanceShutdownTimerV2ResponseBody) *GetInstanceShutdownTimerV2Response {
	s.Body = v
	return s
}

type GetInstanceSnapshotResponseBody struct {
	// 实例快照保存时间（GMT）
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 实例快照修改时间（GMT）
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例快照描述
	InstanceSnapshotDescription *string `json:"InstanceSnapshotDescription,omitempty" xml:"InstanceSnapshotDescription,omitempty"`
	// 实例快照ID
	InstanceSnapshotId *string `json:"InstanceSnapshotId,omitempty" xml:"InstanceSnapshotId,omitempty"`
	// 实例快照名称
	InstanceSnapshotName *string `json:"InstanceSnapshotName,omitempty" xml:"InstanceSnapshotName,omitempty"`
	// 实例快照存储地址
	InstanceSnapshotRepoUrl *string `json:"InstanceSnapshotRepoUrl,omitempty" xml:"InstanceSnapshotRepoUrl,omitempty"`
	// 实例快照状态
	InstanceSnapshotStatus *string `json:"InstanceSnapshotStatus,omitempty" xml:"InstanceSnapshotStatus,omitempty"`
	// 实例快照标签
	InstanceSnapshotTag *string `json:"InstanceSnapshotTag,omitempty" xml:"InstanceSnapshotTag,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInstanceSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceSnapshotResponseBody) SetGmtCreateTime(v string) *GetInstanceSnapshotResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetGmtModifiedTime(v string) *GetInstanceSnapshotResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetInstanceId(v string) *GetInstanceSnapshotResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetInstanceSnapshotDescription(v string) *GetInstanceSnapshotResponseBody {
	s.InstanceSnapshotDescription = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetInstanceSnapshotId(v string) *GetInstanceSnapshotResponseBody {
	s.InstanceSnapshotId = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetInstanceSnapshotName(v string) *GetInstanceSnapshotResponseBody {
	s.InstanceSnapshotName = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetInstanceSnapshotRepoUrl(v string) *GetInstanceSnapshotResponseBody {
	s.InstanceSnapshotRepoUrl = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetInstanceSnapshotStatus(v string) *GetInstanceSnapshotResponseBody {
	s.InstanceSnapshotStatus = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetInstanceSnapshotTag(v string) *GetInstanceSnapshotResponseBody {
	s.InstanceSnapshotTag = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetRequestId(v string) *GetInstanceSnapshotResponseBody {
	s.RequestId = &v
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

type GetInstanceSnapshotV2ResponseBody struct {
	// 实例快照创建时间
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 实例快照修改时间
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// 实例快照的镜像Id
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// 实例快照的镜像地址
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例快照错误代码
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// 实例快照错误消息
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 实例快照Id
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// 实例快照名称
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	// 实例快照状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetInstanceSnapshotV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceSnapshotV2ResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceSnapshotV2ResponseBody) SetGmtCreateTime(v string) *GetInstanceSnapshotV2ResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetInstanceSnapshotV2ResponseBody) SetGmtModifiedTime(v string) *GetInstanceSnapshotV2ResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetInstanceSnapshotV2ResponseBody) SetImageId(v string) *GetInstanceSnapshotV2ResponseBody {
	s.ImageId = &v
	return s
}

func (s *GetInstanceSnapshotV2ResponseBody) SetImageUrl(v string) *GetInstanceSnapshotV2ResponseBody {
	s.ImageUrl = &v
	return s
}

func (s *GetInstanceSnapshotV2ResponseBody) SetInstanceId(v string) *GetInstanceSnapshotV2ResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceSnapshotV2ResponseBody) SetReasonCode(v string) *GetInstanceSnapshotV2ResponseBody {
	s.ReasonCode = &v
	return s
}

func (s *GetInstanceSnapshotV2ResponseBody) SetReasonMessage(v string) *GetInstanceSnapshotV2ResponseBody {
	s.ReasonMessage = &v
	return s
}

func (s *GetInstanceSnapshotV2ResponseBody) SetRequestId(v string) *GetInstanceSnapshotV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceSnapshotV2ResponseBody) SetSnapshotId(v string) *GetInstanceSnapshotV2ResponseBody {
	s.SnapshotId = &v
	return s
}

func (s *GetInstanceSnapshotV2ResponseBody) SetSnapshotName(v string) *GetInstanceSnapshotV2ResponseBody {
	s.SnapshotName = &v
	return s
}

func (s *GetInstanceSnapshotV2ResponseBody) SetStatus(v string) *GetInstanceSnapshotV2ResponseBody {
	s.Status = &v
	return s
}

type GetInstanceSnapshotV2Response struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetInstanceSnapshotV2ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInstanceSnapshotV2Response) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceSnapshotV2Response) GoString() string {
	return s.String()
}

func (s *GetInstanceSnapshotV2Response) SetHeaders(v map[string]*string) *GetInstanceSnapshotV2Response {
	s.Headers = v
	return s
}

func (s *GetInstanceSnapshotV2Response) SetBody(v *GetInstanceSnapshotV2ResponseBody) *GetInstanceSnapshotV2Response {
	s.Body = v
	return s
}

type GetInstanceV2ResponseBody struct {
	// 实例计算类型
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	// 工作空间内是否他人可见
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// 累计运行时间（ms）
	AccumulatedRunningTimeInMs *int64 `json:"AccumulatedRunningTimeInMs,omitempty" xml:"AccumulatedRunningTimeInMs,omitempty"`
	// 数据集集合
	Datasets []*GetInstanceV2ResponseBodyDatasets `json:"Datasets,omitempty" xml:"Datasets,omitempty" type:"Repeated"`
	// 实例对应的Ecs规格
	EcsSpec *string `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	// 环境变量
	EnvironmentVariables map[string]*string `json:"EnvironmentVariables,omitempty" xml:"EnvironmentVariables,omitempty"`
	// 实例创建时间
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 实例修改时间
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// 镜像Id
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// 镜像地址
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// 定时关机任务
	InstanceShutdownTimer *GetInstanceV2ResponseBodyInstanceShutdownTimer `json:"InstanceShutdownTimer,omitempty" xml:"InstanceShutdownTimer,omitempty" type:"Struct"`
	// 实例Url
	InstanceUrl *string `json:"InstanceUrl,omitempty" xml:"InstanceUrl,omitempty"`
	// Jupyterlab Url
	JupyterlabUrl *string `json:"JupyterlabUrl,omitempty" xml:"JupyterlabUrl,omitempty"`
	// 最新保存的用户镜像
	LatestSnapshot *GetInstanceV2ResponseBodyLatestSnapshot `json:"LatestSnapshot,omitempty" xml:"LatestSnapshot,omitempty" type:"Struct"`
	// 支付类型
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// 实例错误代码
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// 实例错误原因
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 实例状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// terminal url
	TerminalUrl *string `json:"TerminalUrl,omitempty" xml:"TerminalUrl,omitempty"`
	// 用户Id
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// user vpc配置
	UserVpc *GetInstanceV2ResponseBodyUserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	// web ide url
	WebIDEUrl *string `json:"WebIDEUrl,omitempty" xml:"WebIDEUrl,omitempty"`
	// 工作空间Id
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// 工作空间名称
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s GetInstanceV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceV2ResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceV2ResponseBody) SetAcceleratorType(v string) *GetInstanceV2ResponseBody {
	s.AcceleratorType = &v
	return s
}

func (s *GetInstanceV2ResponseBody) SetAccessibility(v string) *GetInstanceV2ResponseBody {
	s.Accessibility = &v
	return s
}

func (s *GetInstanceV2ResponseBody) SetAccumulatedRunningTimeInMs(v int64) *GetInstanceV2ResponseBody {
	s.AccumulatedRunningTimeInMs = &v
	return s
}

func (s *GetInstanceV2ResponseBody) SetDatasets(v []*GetInstanceV2ResponseBodyDatasets) *GetInstanceV2ResponseBody {
	s.Datasets = v
	return s
}

func (s *GetInstanceV2ResponseBody) SetEcsSpec(v string) *GetInstanceV2ResponseBody {
	s.EcsSpec = &v
	return s
}

func (s *GetInstanceV2ResponseBody) SetEnvironmentVariables(v map[string]*string) *GetInstanceV2ResponseBody {
	s.EnvironmentVariables = v
	return s
}

func (s *GetInstanceV2ResponseBody) SetGmtCreateTime(v string) *GetInstanceV2ResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetInstanceV2ResponseBody) SetGmtModifiedTime(v string) *GetInstanceV2ResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetInstanceV2ResponseBody) SetImageId(v string) *GetInstanceV2ResponseBody {
	s.ImageId = &v
	return s
}

func (s *GetInstanceV2ResponseBody) SetImageName(v string) *GetInstanceV2ResponseBody {
	s.ImageName = &v
	return s
}

func (s *GetInstanceV2ResponseBody) SetImageUrl(v string) *GetInstanceV2ResponseBody {
	s.ImageUrl = &v
	return s
}

func (s *GetInstanceV2ResponseBody) SetInstanceId(v string) *GetInstanceV2ResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceV2ResponseBody) SetInstanceName(v string) *GetInstanceV2ResponseBody {
	s.InstanceName = &v
	return s
}

func (s *GetInstanceV2ResponseBody) SetInstanceShutdownTimer(v *GetInstanceV2ResponseBodyInstanceShutdownTimer) *GetInstanceV2ResponseBody {
	s.InstanceShutdownTimer = v
	return s
}

func (s *GetInstanceV2ResponseBody) SetInstanceUrl(v string) *GetInstanceV2ResponseBody {
	s.InstanceUrl = &v
	return s
}

func (s *GetInstanceV2ResponseBody) SetJupyterlabUrl(v string) *GetInstanceV2ResponseBody {
	s.JupyterlabUrl = &v
	return s
}

func (s *GetInstanceV2ResponseBody) SetLatestSnapshot(v *GetInstanceV2ResponseBodyLatestSnapshot) *GetInstanceV2ResponseBody {
	s.LatestSnapshot = v
	return s
}

func (s *GetInstanceV2ResponseBody) SetPaymentType(v string) *GetInstanceV2ResponseBody {
	s.PaymentType = &v
	return s
}

func (s *GetInstanceV2ResponseBody) SetReasonCode(v string) *GetInstanceV2ResponseBody {
	s.ReasonCode = &v
	return s
}

func (s *GetInstanceV2ResponseBody) SetReasonMessage(v string) *GetInstanceV2ResponseBody {
	s.ReasonMessage = &v
	return s
}

func (s *GetInstanceV2ResponseBody) SetRequestId(v string) *GetInstanceV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceV2ResponseBody) SetStatus(v string) *GetInstanceV2ResponseBody {
	s.Status = &v
	return s
}

func (s *GetInstanceV2ResponseBody) SetTerminalUrl(v string) *GetInstanceV2ResponseBody {
	s.TerminalUrl = &v
	return s
}

func (s *GetInstanceV2ResponseBody) SetUserId(v string) *GetInstanceV2ResponseBody {
	s.UserId = &v
	return s
}

func (s *GetInstanceV2ResponseBody) SetUserVpc(v *GetInstanceV2ResponseBodyUserVpc) *GetInstanceV2ResponseBody {
	s.UserVpc = v
	return s
}

func (s *GetInstanceV2ResponseBody) SetWebIDEUrl(v string) *GetInstanceV2ResponseBody {
	s.WebIDEUrl = &v
	return s
}

func (s *GetInstanceV2ResponseBody) SetWorkspaceId(v string) *GetInstanceV2ResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetInstanceV2ResponseBody) SetWorkspaceName(v string) *GetInstanceV2ResponseBody {
	s.WorkspaceName = &v
	return s
}

type GetInstanceV2ResponseBodyDatasets struct {
	// 数据集Id
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// 容器内挂载路径
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
}

func (s GetInstanceV2ResponseBodyDatasets) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceV2ResponseBodyDatasets) GoString() string {
	return s.String()
}

func (s *GetInstanceV2ResponseBodyDatasets) SetDatasetId(v string) *GetInstanceV2ResponseBodyDatasets {
	s.DatasetId = &v
	return s
}

func (s *GetInstanceV2ResponseBodyDatasets) SetMountPath(v string) *GetInstanceV2ResponseBodyDatasets {
	s.MountPath = &v
	return s
}

type GetInstanceV2ResponseBodyInstanceShutdownTimer struct {
	// 设定关机时间
	DueTime *string `json:"DueTime,omitempty" xml:"DueTime,omitempty"`
	// 创建时间
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 修改时间
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 剩余关机时间（ms）
	RemainingTimeInMs *int64 `json:"RemainingTimeInMs,omitempty" xml:"RemainingTimeInMs,omitempty"`
}

func (s GetInstanceV2ResponseBodyInstanceShutdownTimer) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceV2ResponseBodyInstanceShutdownTimer) GoString() string {
	return s.String()
}

func (s *GetInstanceV2ResponseBodyInstanceShutdownTimer) SetDueTime(v string) *GetInstanceV2ResponseBodyInstanceShutdownTimer {
	s.DueTime = &v
	return s
}

func (s *GetInstanceV2ResponseBodyInstanceShutdownTimer) SetGmtCreateTime(v string) *GetInstanceV2ResponseBodyInstanceShutdownTimer {
	s.GmtCreateTime = &v
	return s
}

func (s *GetInstanceV2ResponseBodyInstanceShutdownTimer) SetGmtModifiedTime(v string) *GetInstanceV2ResponseBodyInstanceShutdownTimer {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetInstanceV2ResponseBodyInstanceShutdownTimer) SetInstanceId(v string) *GetInstanceV2ResponseBodyInstanceShutdownTimer {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceV2ResponseBodyInstanceShutdownTimer) SetRemainingTimeInMs(v int64) *GetInstanceV2ResponseBodyInstanceShutdownTimer {
	s.RemainingTimeInMs = &v
	return s
}

type GetInstanceV2ResponseBodyLatestSnapshot struct {
	// 快照创建时间
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 快照修改时间
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// 镜像Id
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// 镜像Url
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// 镜像仓库Url
	RepositoryUrl *string `json:"RepositoryUrl,omitempty" xml:"RepositoryUrl,omitempty"`
}

func (s GetInstanceV2ResponseBodyLatestSnapshot) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceV2ResponseBodyLatestSnapshot) GoString() string {
	return s.String()
}

func (s *GetInstanceV2ResponseBodyLatestSnapshot) SetGmtCreateTime(v string) *GetInstanceV2ResponseBodyLatestSnapshot {
	s.GmtCreateTime = &v
	return s
}

func (s *GetInstanceV2ResponseBodyLatestSnapshot) SetGmtModifiedTime(v string) *GetInstanceV2ResponseBodyLatestSnapshot {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetInstanceV2ResponseBodyLatestSnapshot) SetImageId(v string) *GetInstanceV2ResponseBodyLatestSnapshot {
	s.ImageId = &v
	return s
}

func (s *GetInstanceV2ResponseBodyLatestSnapshot) SetImageName(v string) *GetInstanceV2ResponseBodyLatestSnapshot {
	s.ImageName = &v
	return s
}

func (s *GetInstanceV2ResponseBodyLatestSnapshot) SetImageUrl(v string) *GetInstanceV2ResponseBodyLatestSnapshot {
	s.ImageUrl = &v
	return s
}

func (s *GetInstanceV2ResponseBodyLatestSnapshot) SetRepositoryUrl(v string) *GetInstanceV2ResponseBodyLatestSnapshot {
	s.RepositoryUrl = &v
	return s
}

type GetInstanceV2ResponseBodyUserVpc struct {
	// Security Group Id
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// VSwitch Id
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// Vpc Id
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetInstanceV2ResponseBodyUserVpc) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceV2ResponseBodyUserVpc) GoString() string {
	return s.String()
}

func (s *GetInstanceV2ResponseBodyUserVpc) SetSecurityGroupId(v string) *GetInstanceV2ResponseBodyUserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *GetInstanceV2ResponseBodyUserVpc) SetVSwitchId(v string) *GetInstanceV2ResponseBodyUserVpc {
	s.VSwitchId = &v
	return s
}

func (s *GetInstanceV2ResponseBodyUserVpc) SetVpcId(v string) *GetInstanceV2ResponseBodyUserVpc {
	s.VpcId = &v
	return s
}

type GetInstanceV2Response struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetInstanceV2ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInstanceV2Response) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceV2Response) GoString() string {
	return s.String()
}

func (s *GetInstanceV2Response) SetHeaders(v map[string]*string) *GetInstanceV2Response {
	s.Headers = v
	return s
}

func (s *GetInstanceV2Response) SetBody(v *GetInstanceV2ResponseBody) *GetInstanceV2Response {
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 统计数据
	Statistics map[string]interface{} `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
}

func (s GetInstancesStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstancesStatisticsResponseBody) SetRequestId(v string) *GetInstancesStatisticsResponseBody {
	s.RequestId = &v
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

type GetUserConfigV2ResponseBody struct {
	// 用户账号金额是否充足
	AccountSufficient *bool `json:"AccountSufficient,omitempty" xml:"AccountSufficient,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetUserConfigV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserConfigV2ResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserConfigV2ResponseBody) SetAccountSufficient(v bool) *GetUserConfigV2ResponseBody {
	s.AccountSufficient = &v
	return s
}

func (s *GetUserConfigV2ResponseBody) SetRequestId(v string) *GetUserConfigV2ResponseBody {
	s.RequestId = &v
	return s
}

type GetUserConfigV2Response struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUserConfigV2ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserConfigV2Response) String() string {
	return tea.Prettify(s)
}

func (s GetUserConfigV2Response) GoString() string {
	return s.String()
}

func (s *GetUserConfigV2Response) SetHeaders(v map[string]*string) *GetUserConfigV2Response {
	s.Headers = v
	return s
}

func (s *GetUserConfigV2Response) SetBody(v *GetUserConfigV2ResponseBody) *GetUserConfigV2Response {
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
	// 请求ecs规格列表
	EcsSpecs []*EcsSpec `json:"EcsSpecs,omitempty" xml:"EcsSpecs,omitempty" type:"Repeated"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 符合要求的ecs规格数量
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEcsSpecsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEcsSpecsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEcsSpecsResponseBody) SetEcsSpecs(v []*EcsSpec) *ListEcsSpecsResponseBody {
	s.EcsSpecs = v
	return s
}

func (s *ListEcsSpecsResponseBody) SetRequestId(v string) *ListEcsSpecsResponseBody {
	s.RequestId = &v
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

type ListEcsSpecsV2Request struct {
	// 加速类型
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	// 排序顺序
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// 页数
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页大小
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 排序字段
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListEcsSpecsV2Request) String() string {
	return tea.Prettify(s)
}

func (s ListEcsSpecsV2Request) GoString() string {
	return s.String()
}

func (s *ListEcsSpecsV2Request) SetAcceleratorType(v string) *ListEcsSpecsV2Request {
	s.AcceleratorType = &v
	return s
}

func (s *ListEcsSpecsV2Request) SetOrder(v string) *ListEcsSpecsV2Request {
	s.Order = &v
	return s
}

func (s *ListEcsSpecsV2Request) SetPageNumber(v int64) *ListEcsSpecsV2Request {
	s.PageNumber = &v
	return s
}

func (s *ListEcsSpecsV2Request) SetPageSize(v int64) *ListEcsSpecsV2Request {
	s.PageSize = &v
	return s
}

func (s *ListEcsSpecsV2Request) SetSortBy(v string) *ListEcsSpecsV2Request {
	s.SortBy = &v
	return s
}

type ListEcsSpecsV2ResponseBody struct {
	// 本分页中请求的实例列表
	EcsSpecs []*ListEcsSpecsV2ResponseBodyEcsSpecs `json:"EcsSpecs,omitempty" xml:"EcsSpecs,omitempty" type:"Repeated"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 实例总数
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEcsSpecsV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEcsSpecsV2ResponseBody) GoString() string {
	return s.String()
}

func (s *ListEcsSpecsV2ResponseBody) SetEcsSpecs(v []*ListEcsSpecsV2ResponseBodyEcsSpecs) *ListEcsSpecsV2ResponseBody {
	s.EcsSpecs = v
	return s
}

func (s *ListEcsSpecsV2ResponseBody) SetRequestId(v string) *ListEcsSpecsV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEcsSpecsV2ResponseBody) SetTotalCount(v int64) *ListEcsSpecsV2ResponseBody {
	s.TotalCount = &v
	return s
}

type ListEcsSpecsV2ResponseBodyEcsSpecs struct {
	// 资源类型
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	// CPU核数
	CPU *int64 `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// 货币单位
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// GPU卡数
	GPU *int64 `json:"GPU,omitempty" xml:"GPU,omitempty"`
	// 显卡类型
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	// 实例接收带宽
	InstanceBandwidthRx *int64 `json:"InstanceBandwidthRx,omitempty" xml:"InstanceBandwidthRx,omitempty"`
	// 实例规格
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// 内存大小(GB)
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// 价格
	Price *float64 `json:"Price,omitempty" xml:"Price,omitempty"`
	// 系统盘大小(GB)
	SystemDiskCapacity *int64 `json:"SystemDiskCapacity,omitempty" xml:"SystemDiskCapacity,omitempty"`
}

func (s ListEcsSpecsV2ResponseBodyEcsSpecs) String() string {
	return tea.Prettify(s)
}

func (s ListEcsSpecsV2ResponseBodyEcsSpecs) GoString() string {
	return s.String()
}

func (s *ListEcsSpecsV2ResponseBodyEcsSpecs) SetAcceleratorType(v string) *ListEcsSpecsV2ResponseBodyEcsSpecs {
	s.AcceleratorType = &v
	return s
}

func (s *ListEcsSpecsV2ResponseBodyEcsSpecs) SetCPU(v int64) *ListEcsSpecsV2ResponseBodyEcsSpecs {
	s.CPU = &v
	return s
}

func (s *ListEcsSpecsV2ResponseBodyEcsSpecs) SetCurrency(v string) *ListEcsSpecsV2ResponseBodyEcsSpecs {
	s.Currency = &v
	return s
}

func (s *ListEcsSpecsV2ResponseBodyEcsSpecs) SetGPU(v int64) *ListEcsSpecsV2ResponseBodyEcsSpecs {
	s.GPU = &v
	return s
}

func (s *ListEcsSpecsV2ResponseBodyEcsSpecs) SetGPUType(v string) *ListEcsSpecsV2ResponseBodyEcsSpecs {
	s.GPUType = &v
	return s
}

func (s *ListEcsSpecsV2ResponseBodyEcsSpecs) SetInstanceBandwidthRx(v int64) *ListEcsSpecsV2ResponseBodyEcsSpecs {
	s.InstanceBandwidthRx = &v
	return s
}

func (s *ListEcsSpecsV2ResponseBodyEcsSpecs) SetInstanceType(v string) *ListEcsSpecsV2ResponseBodyEcsSpecs {
	s.InstanceType = &v
	return s
}

func (s *ListEcsSpecsV2ResponseBodyEcsSpecs) SetMemory(v float32) *ListEcsSpecsV2ResponseBodyEcsSpecs {
	s.Memory = &v
	return s
}

func (s *ListEcsSpecsV2ResponseBodyEcsSpecs) SetPrice(v float64) *ListEcsSpecsV2ResponseBodyEcsSpecs {
	s.Price = &v
	return s
}

func (s *ListEcsSpecsV2ResponseBodyEcsSpecs) SetSystemDiskCapacity(v int64) *ListEcsSpecsV2ResponseBodyEcsSpecs {
	s.SystemDiskCapacity = &v
	return s
}

type ListEcsSpecsV2Response struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListEcsSpecsV2ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEcsSpecsV2Response) String() string {
	return tea.Prettify(s)
}

func (s ListEcsSpecsV2Response) GoString() string {
	return s.String()
}

func (s *ListEcsSpecsV2Response) SetHeaders(v map[string]*string) *ListEcsSpecsV2Response {
	s.Headers = v
	return s
}

func (s *ListEcsSpecsV2Response) SetBody(v *ListEcsSpecsV2ResponseBody) *ListEcsSpecsV2Response {
	s.Body = v
	return s
}

type ListImagesRequest struct {
	// 资源类型
	AcceleratorTypeEquals *string `json:"AcceleratorTypeEquals,omitempty" xml:"AcceleratorTypeEquals,omitempty"`
	// 容器名称关键字
	ImageNameContains *string `json:"ImageNameContains,omitempty" xml:"ImageNameContains,omitempty"`
	// 镜像类型
	ImageTypeEquals *string `json:"ImageTypeEquals,omitempty" xml:"ImageTypeEquals,omitempty"`
	// 产品
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// 工作空间Id
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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

func (s *ListImagesRequest) SetImageNameContains(v string) *ListImagesRequest {
	s.ImageNameContains = &v
	return s
}

func (s *ListImagesRequest) SetImageTypeEquals(v string) *ListImagesRequest {
	s.ImageTypeEquals = &v
	return s
}

func (s *ListImagesRequest) SetProduct(v string) *ListImagesRequest {
	s.Product = &v
	return s
}

func (s *ListImagesRequest) SetWorkspaceId(v string) *ListImagesRequest {
	s.WorkspaceId = &v
	return s
}

type ListImagesResponseBody struct {
	// 镜像列表
	Images []*Image `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBody) SetImages(v []*Image) *ListImagesResponseBody {
	s.Images = v
	return s
}

func (s *ListImagesResponseBody) SetRequestId(v string) *ListImagesResponseBody {
	s.RequestId = &v
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

type ListInstanceSnapshotV2Request struct {
	Order      *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortBy     *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListInstanceSnapshotV2Request) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceSnapshotV2Request) GoString() string {
	return s.String()
}

func (s *ListInstanceSnapshotV2Request) SetOrder(v string) *ListInstanceSnapshotV2Request {
	s.Order = &v
	return s
}

func (s *ListInstanceSnapshotV2Request) SetPageNumber(v int64) *ListInstanceSnapshotV2Request {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceSnapshotV2Request) SetPageSize(v int64) *ListInstanceSnapshotV2Request {
	s.PageSize = &v
	return s
}

func (s *ListInstanceSnapshotV2Request) SetSortBy(v string) *ListInstanceSnapshotV2Request {
	s.SortBy = &v
	return s
}

type ListInstanceSnapshotV2ResponseBody struct {
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 本分页中请求的实例镜像列表
	Snapshots []*ListInstanceSnapshotV2ResponseBodySnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Repeated"`
	// 实例总数
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstanceSnapshotV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceSnapshotV2ResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceSnapshotV2ResponseBody) SetRequestId(v string) *ListInstanceSnapshotV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceSnapshotV2ResponseBody) SetSnapshots(v []*ListInstanceSnapshotV2ResponseBodySnapshots) *ListInstanceSnapshotV2ResponseBody {
	s.Snapshots = v
	return s
}

func (s *ListInstanceSnapshotV2ResponseBody) SetTotalCount(v int64) *ListInstanceSnapshotV2ResponseBody {
	s.TotalCount = &v
	return s
}

type ListInstanceSnapshotV2ResponseBodySnapshots struct {
	// 实例快照创建时间
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 实例快照修改时间
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// 实例快照的镜像Id
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// 实例快照的镜像地址
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例快照错误代码
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// 实例快照错误消息
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// 实例快照Id
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// 实例快照名称
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	// 实例快照状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListInstanceSnapshotV2ResponseBodySnapshots) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceSnapshotV2ResponseBodySnapshots) GoString() string {
	return s.String()
}

func (s *ListInstanceSnapshotV2ResponseBodySnapshots) SetGmtCreateTime(v string) *ListInstanceSnapshotV2ResponseBodySnapshots {
	s.GmtCreateTime = &v
	return s
}

func (s *ListInstanceSnapshotV2ResponseBodySnapshots) SetGmtModifiedTime(v string) *ListInstanceSnapshotV2ResponseBodySnapshots {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListInstanceSnapshotV2ResponseBodySnapshots) SetImageId(v string) *ListInstanceSnapshotV2ResponseBodySnapshots {
	s.ImageId = &v
	return s
}

func (s *ListInstanceSnapshotV2ResponseBodySnapshots) SetImageUrl(v string) *ListInstanceSnapshotV2ResponseBodySnapshots {
	s.ImageUrl = &v
	return s
}

func (s *ListInstanceSnapshotV2ResponseBodySnapshots) SetInstanceId(v string) *ListInstanceSnapshotV2ResponseBodySnapshots {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceSnapshotV2ResponseBodySnapshots) SetReasonCode(v string) *ListInstanceSnapshotV2ResponseBodySnapshots {
	s.ReasonCode = &v
	return s
}

func (s *ListInstanceSnapshotV2ResponseBodySnapshots) SetReasonMessage(v string) *ListInstanceSnapshotV2ResponseBodySnapshots {
	s.ReasonMessage = &v
	return s
}

func (s *ListInstanceSnapshotV2ResponseBodySnapshots) SetSnapshotId(v string) *ListInstanceSnapshotV2ResponseBodySnapshots {
	s.SnapshotId = &v
	return s
}

func (s *ListInstanceSnapshotV2ResponseBodySnapshots) SetSnapshotName(v string) *ListInstanceSnapshotV2ResponseBodySnapshots {
	s.SnapshotName = &v
	return s
}

func (s *ListInstanceSnapshotV2ResponseBodySnapshots) SetStatus(v string) *ListInstanceSnapshotV2ResponseBodySnapshots {
	s.Status = &v
	return s
}

type ListInstanceSnapshotV2Response struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInstanceSnapshotV2ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInstanceSnapshotV2Response) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceSnapshotV2Response) GoString() string {
	return s.String()
}

func (s *ListInstanceSnapshotV2Response) SetHeaders(v map[string]*string) *ListInstanceSnapshotV2Response {
	s.Headers = v
	return s
}

func (s *ListInstanceSnapshotV2Response) SetBody(v *ListInstanceSnapshotV2ResponseBody) *ListInstanceSnapshotV2Response {
	s.Body = v
	return s
}

type ListInstanceSnapshotsResponseBody struct {
	// 镜像快照列表
	InstanceSnapshots []*InstanceSnapshot `json:"InstanceSnapshots,omitempty" xml:"InstanceSnapshots,omitempty" type:"Repeated"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInstanceSnapshotsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceSnapshotsResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceSnapshotsResponseBody) SetInstanceSnapshots(v []*InstanceSnapshot) *ListInstanceSnapshotsResponseBody {
	s.InstanceSnapshots = v
	return s
}

func (s *ListInstanceSnapshotsResponseBody) SetRequestId(v string) *ListInstanceSnapshotsResponseBody {
	s.RequestId = &v
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

type ListInstanceStatisticsV2Request struct {
	// 工作空间列表
	WorkspaceIds *string `json:"WorkspaceIds,omitempty" xml:"WorkspaceIds,omitempty"`
}

func (s ListInstanceStatisticsV2Request) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceStatisticsV2Request) GoString() string {
	return s.String()
}

func (s *ListInstanceStatisticsV2Request) SetWorkspaceIds(v string) *ListInstanceStatisticsV2Request {
	s.WorkspaceIds = &v
	return s
}

type ListInstanceStatisticsV2ResponseBody struct {
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 统计信息
	Statistics map[string]map[string]interface{} `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
}

func (s ListInstanceStatisticsV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceStatisticsV2ResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceStatisticsV2ResponseBody) SetRequestId(v string) *ListInstanceStatisticsV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceStatisticsV2ResponseBody) SetStatistics(v map[string]map[string]interface{}) *ListInstanceStatisticsV2ResponseBody {
	s.Statistics = v
	return s
}

type ListInstanceStatisticsV2Response struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInstanceStatisticsV2ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInstanceStatisticsV2Response) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceStatisticsV2Response) GoString() string {
	return s.String()
}

func (s *ListInstanceStatisticsV2Response) SetHeaders(v map[string]*string) *ListInstanceStatisticsV2Response {
	s.Headers = v
	return s
}

func (s *ListInstanceStatisticsV2Response) SetBody(v *ListInstanceStatisticsV2ResponseBody) *ListInstanceStatisticsV2Response {
	s.Body = v
	return s
}

type ListInstancesRequest struct {
	// 是否在工作空间内查询
	InWorkspace *bool `json:"InWorkspace,omitempty" xml:"InWorkspace,omitempty"`
	// 实例名称关键字
	InstanceNameContains *string `json:"InstanceNameContains,omitempty" xml:"InstanceNameContains,omitempty"`
	// 实例状态
	InstanceStatusEquals *string `json:"InstanceStatusEquals,omitempty" xml:"InstanceStatusEquals,omitempty"`
	// 当前页
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页返回的实例数
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 排序字段
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// 升序降序
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// 工作空间Id
	WorkspaceIdEquals *string `json:"WorkspaceIdEquals,omitempty" xml:"WorkspaceIdEquals,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) SetInWorkspace(v bool) *ListInstancesRequest {
	s.InWorkspace = &v
	return s
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

func (s *ListInstancesRequest) SetSortBy(v string) *ListInstancesRequest {
	s.SortBy = &v
	return s
}

func (s *ListInstancesRequest) SetSortOrder(v string) *ListInstancesRequest {
	s.SortOrder = &v
	return s
}

func (s *ListInstancesRequest) SetWorkspaceIdEquals(v string) *ListInstancesRequest {
	s.WorkspaceIdEquals = &v
	return s
}

type ListInstancesResponseBody struct {
	// 实例列表
	Instances []*Instance `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// 当前页
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页返回的实例数
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 符合条件的实例数
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
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

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
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

type ListInstancesStatusRequest struct {
	// 实例Id列表
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s ListInstancesStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesStatusRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesStatusRequest) SetInstanceIds(v string) *ListInstancesStatusRequest {
	s.InstanceIds = &v
	return s
}

type ListInstancesStatusResponseBody struct {
	// Id of the request
	RequestId *string           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Statuses  []*InstanceStatus `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
}

func (s ListInstancesStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesStatusResponseBody) SetRequestId(v string) *ListInstancesStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesStatusResponseBody) SetStatuses(v []*InstanceStatus) *ListInstancesStatusResponseBody {
	s.Statuses = v
	return s
}

type ListInstancesStatusResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInstancesStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInstancesStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesStatusResponse) GoString() string {
	return s.String()
}

func (s *ListInstancesStatusResponse) SetHeaders(v map[string]*string) *ListInstancesStatusResponse {
	s.Headers = v
	return s
}

func (s *ListInstancesStatusResponse) SetBody(v *ListInstancesStatusResponseBody) *ListInstancesStatusResponse {
	s.Body = v
	return s
}

type ListInstancesV2Request struct {
	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// 排列顺序
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// 页码
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页数量大小
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 排序字段
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// 实例状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 工作空间Id
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListInstancesV2Request) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesV2Request) GoString() string {
	return s.String()
}

func (s *ListInstancesV2Request) SetInstanceName(v string) *ListInstancesV2Request {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesV2Request) SetOrder(v string) *ListInstancesV2Request {
	s.Order = &v
	return s
}

func (s *ListInstancesV2Request) SetPageNumber(v int64) *ListInstancesV2Request {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesV2Request) SetPageSize(v int64) *ListInstancesV2Request {
	s.PageSize = &v
	return s
}

func (s *ListInstancesV2Request) SetSortBy(v string) *ListInstancesV2Request {
	s.SortBy = &v
	return s
}

func (s *ListInstancesV2Request) SetStatus(v string) *ListInstancesV2Request {
	s.Status = &v
	return s
}

func (s *ListInstancesV2Request) SetWorkspaceId(v string) *ListInstancesV2Request {
	s.WorkspaceId = &v
	return s
}

type ListInstancesV2ResponseBody struct {
	// 本分页中请求的实例列表
	Instances []*ListInstancesV2ResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 实例总数
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstancesV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesV2ResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesV2ResponseBody) SetInstances(v []*ListInstancesV2ResponseBodyInstances) *ListInstancesV2ResponseBody {
	s.Instances = v
	return s
}

func (s *ListInstancesV2ResponseBody) SetRequestId(v string) *ListInstancesV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesV2ResponseBody) SetTotalCount(v int64) *ListInstancesV2ResponseBody {
	s.TotalCount = &v
	return s
}

type ListInstancesV2ResponseBodyInstances struct {
	// 实例计算类型
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	// 工作空间内是否他人可见
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// 累计运行时间（ms）
	AccumulatedRunningTimeInMs *int64 `json:"AccumulatedRunningTimeInMs,omitempty" xml:"AccumulatedRunningTimeInMs,omitempty"`
	// 数据集集合
	Datasets []*ListInstancesV2ResponseBodyInstancesDatasets `json:"Datasets,omitempty" xml:"Datasets,omitempty" type:"Repeated"`
	// 实例对应的Ecs规格
	EcsSpec *string `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	// 环境变量
	EnvironmentVariables map[string]*string `json:"EnvironmentVariables,omitempty" xml:"EnvironmentVariables,omitempty"`
	// 实例创建时间
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 实例修改时间
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// 镜像Id
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// 镜像地址
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// 定时关机任务
	InstanceShutdownTimer *ListInstancesV2ResponseBodyInstancesInstanceShutdownTimer `json:"InstanceShutdownTimer,omitempty" xml:"InstanceShutdownTimer,omitempty" type:"Struct"`
	// 实例Url
	InstanceUrl *string `json:"InstanceUrl,omitempty" xml:"InstanceUrl,omitempty"`
	// Jupyterlab Url
	JupyterlabUrl *string `json:"JupyterlabUrl,omitempty" xml:"JupyterlabUrl,omitempty"`
	// 最新保存的用户镜像
	LatestSnapshot *ListInstancesV2ResponseBodyInstancesLatestSnapshot `json:"LatestSnapshot,omitempty" xml:"LatestSnapshot,omitempty" type:"Struct"`
	// 支付类型
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// 实例错误代码
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// 实例错误原因
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// 实例状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// terminal url
	TerminalUrl *string `json:"TerminalUrl,omitempty" xml:"TerminalUrl,omitempty"`
	// 用户Id
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// user vpc配置
	UserVpc *ListInstancesV2ResponseBodyInstancesUserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	// web ide url
	WebIDEUrl *string `json:"WebIDEUrl,omitempty" xml:"WebIDEUrl,omitempty"`
	// 工作空间Id
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// 工作空间名称
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s ListInstancesV2ResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesV2ResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListInstancesV2ResponseBodyInstances) SetAcceleratorType(v string) *ListInstancesV2ResponseBodyInstances {
	s.AcceleratorType = &v
	return s
}

func (s *ListInstancesV2ResponseBodyInstances) SetAccessibility(v string) *ListInstancesV2ResponseBodyInstances {
	s.Accessibility = &v
	return s
}

func (s *ListInstancesV2ResponseBodyInstances) SetAccumulatedRunningTimeInMs(v int64) *ListInstancesV2ResponseBodyInstances {
	s.AccumulatedRunningTimeInMs = &v
	return s
}

func (s *ListInstancesV2ResponseBodyInstances) SetDatasets(v []*ListInstancesV2ResponseBodyInstancesDatasets) *ListInstancesV2ResponseBodyInstances {
	s.Datasets = v
	return s
}

func (s *ListInstancesV2ResponseBodyInstances) SetEcsSpec(v string) *ListInstancesV2ResponseBodyInstances {
	s.EcsSpec = &v
	return s
}

func (s *ListInstancesV2ResponseBodyInstances) SetEnvironmentVariables(v map[string]*string) *ListInstancesV2ResponseBodyInstances {
	s.EnvironmentVariables = v
	return s
}

func (s *ListInstancesV2ResponseBodyInstances) SetGmtCreateTime(v string) *ListInstancesV2ResponseBodyInstances {
	s.GmtCreateTime = &v
	return s
}

func (s *ListInstancesV2ResponseBodyInstances) SetGmtModifiedTime(v string) *ListInstancesV2ResponseBodyInstances {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListInstancesV2ResponseBodyInstances) SetImageId(v string) *ListInstancesV2ResponseBodyInstances {
	s.ImageId = &v
	return s
}

func (s *ListInstancesV2ResponseBodyInstances) SetImageName(v string) *ListInstancesV2ResponseBodyInstances {
	s.ImageName = &v
	return s
}

func (s *ListInstancesV2ResponseBodyInstances) SetImageUrl(v string) *ListInstancesV2ResponseBodyInstances {
	s.ImageUrl = &v
	return s
}

func (s *ListInstancesV2ResponseBodyInstances) SetInstanceId(v string) *ListInstancesV2ResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesV2ResponseBodyInstances) SetInstanceName(v string) *ListInstancesV2ResponseBodyInstances {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesV2ResponseBodyInstances) SetInstanceShutdownTimer(v *ListInstancesV2ResponseBodyInstancesInstanceShutdownTimer) *ListInstancesV2ResponseBodyInstances {
	s.InstanceShutdownTimer = v
	return s
}

func (s *ListInstancesV2ResponseBodyInstances) SetInstanceUrl(v string) *ListInstancesV2ResponseBodyInstances {
	s.InstanceUrl = &v
	return s
}

func (s *ListInstancesV2ResponseBodyInstances) SetJupyterlabUrl(v string) *ListInstancesV2ResponseBodyInstances {
	s.JupyterlabUrl = &v
	return s
}

func (s *ListInstancesV2ResponseBodyInstances) SetLatestSnapshot(v *ListInstancesV2ResponseBodyInstancesLatestSnapshot) *ListInstancesV2ResponseBodyInstances {
	s.LatestSnapshot = v
	return s
}

func (s *ListInstancesV2ResponseBodyInstances) SetPaymentType(v string) *ListInstancesV2ResponseBodyInstances {
	s.PaymentType = &v
	return s
}

func (s *ListInstancesV2ResponseBodyInstances) SetReasonCode(v string) *ListInstancesV2ResponseBodyInstances {
	s.ReasonCode = &v
	return s
}

func (s *ListInstancesV2ResponseBodyInstances) SetReasonMessage(v string) *ListInstancesV2ResponseBodyInstances {
	s.ReasonMessage = &v
	return s
}

func (s *ListInstancesV2ResponseBodyInstances) SetStatus(v string) *ListInstancesV2ResponseBodyInstances {
	s.Status = &v
	return s
}

func (s *ListInstancesV2ResponseBodyInstances) SetTerminalUrl(v string) *ListInstancesV2ResponseBodyInstances {
	s.TerminalUrl = &v
	return s
}

func (s *ListInstancesV2ResponseBodyInstances) SetUserId(v string) *ListInstancesV2ResponseBodyInstances {
	s.UserId = &v
	return s
}

func (s *ListInstancesV2ResponseBodyInstances) SetUserVpc(v *ListInstancesV2ResponseBodyInstancesUserVpc) *ListInstancesV2ResponseBodyInstances {
	s.UserVpc = v
	return s
}

func (s *ListInstancesV2ResponseBodyInstances) SetWebIDEUrl(v string) *ListInstancesV2ResponseBodyInstances {
	s.WebIDEUrl = &v
	return s
}

func (s *ListInstancesV2ResponseBodyInstances) SetWorkspaceId(v string) *ListInstancesV2ResponseBodyInstances {
	s.WorkspaceId = &v
	return s
}

func (s *ListInstancesV2ResponseBodyInstances) SetWorkspaceName(v string) *ListInstancesV2ResponseBodyInstances {
	s.WorkspaceName = &v
	return s
}

type ListInstancesV2ResponseBodyInstancesDatasets struct {
	// 数据集Id
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// 容器内挂载路径
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
}

func (s ListInstancesV2ResponseBodyInstancesDatasets) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesV2ResponseBodyInstancesDatasets) GoString() string {
	return s.String()
}

func (s *ListInstancesV2ResponseBodyInstancesDatasets) SetDatasetId(v string) *ListInstancesV2ResponseBodyInstancesDatasets {
	s.DatasetId = &v
	return s
}

func (s *ListInstancesV2ResponseBodyInstancesDatasets) SetMountPath(v string) *ListInstancesV2ResponseBodyInstancesDatasets {
	s.MountPath = &v
	return s
}

type ListInstancesV2ResponseBodyInstancesInstanceShutdownTimer struct {
	// 设定关机时间
	DueTime *string `json:"DueTime,omitempty" xml:"DueTime,omitempty"`
	// 创建时间
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 修改时间
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 剩余关机时间（ms）
	RemainingTimeInMs *int64 `json:"RemainingTimeInMs,omitempty" xml:"RemainingTimeInMs,omitempty"`
}

func (s ListInstancesV2ResponseBodyInstancesInstanceShutdownTimer) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesV2ResponseBodyInstancesInstanceShutdownTimer) GoString() string {
	return s.String()
}

func (s *ListInstancesV2ResponseBodyInstancesInstanceShutdownTimer) SetDueTime(v string) *ListInstancesV2ResponseBodyInstancesInstanceShutdownTimer {
	s.DueTime = &v
	return s
}

func (s *ListInstancesV2ResponseBodyInstancesInstanceShutdownTimer) SetGmtCreateTime(v string) *ListInstancesV2ResponseBodyInstancesInstanceShutdownTimer {
	s.GmtCreateTime = &v
	return s
}

func (s *ListInstancesV2ResponseBodyInstancesInstanceShutdownTimer) SetGmtModifiedTime(v string) *ListInstancesV2ResponseBodyInstancesInstanceShutdownTimer {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListInstancesV2ResponseBodyInstancesInstanceShutdownTimer) SetInstanceId(v string) *ListInstancesV2ResponseBodyInstancesInstanceShutdownTimer {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesV2ResponseBodyInstancesInstanceShutdownTimer) SetRemainingTimeInMs(v int64) *ListInstancesV2ResponseBodyInstancesInstanceShutdownTimer {
	s.RemainingTimeInMs = &v
	return s
}

type ListInstancesV2ResponseBodyInstancesLatestSnapshot struct {
	// 快照创建时间
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 快照修改时间
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// 镜像Id
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// 镜像Url
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// 镜像仓库Url
	RepositoryUrl *string `json:"RepositoryUrl,omitempty" xml:"RepositoryUrl,omitempty"`
}

func (s ListInstancesV2ResponseBodyInstancesLatestSnapshot) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesV2ResponseBodyInstancesLatestSnapshot) GoString() string {
	return s.String()
}

func (s *ListInstancesV2ResponseBodyInstancesLatestSnapshot) SetGmtCreateTime(v string) *ListInstancesV2ResponseBodyInstancesLatestSnapshot {
	s.GmtCreateTime = &v
	return s
}

func (s *ListInstancesV2ResponseBodyInstancesLatestSnapshot) SetGmtModifiedTime(v string) *ListInstancesV2ResponseBodyInstancesLatestSnapshot {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListInstancesV2ResponseBodyInstancesLatestSnapshot) SetImageId(v string) *ListInstancesV2ResponseBodyInstancesLatestSnapshot {
	s.ImageId = &v
	return s
}

func (s *ListInstancesV2ResponseBodyInstancesLatestSnapshot) SetImageName(v string) *ListInstancesV2ResponseBodyInstancesLatestSnapshot {
	s.ImageName = &v
	return s
}

func (s *ListInstancesV2ResponseBodyInstancesLatestSnapshot) SetImageUrl(v string) *ListInstancesV2ResponseBodyInstancesLatestSnapshot {
	s.ImageUrl = &v
	return s
}

func (s *ListInstancesV2ResponseBodyInstancesLatestSnapshot) SetRepositoryUrl(v string) *ListInstancesV2ResponseBodyInstancesLatestSnapshot {
	s.RepositoryUrl = &v
	return s
}

type ListInstancesV2ResponseBodyInstancesUserVpc struct {
	// Security Group Id
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// VSwitch Id
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// Vpc Id
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListInstancesV2ResponseBodyInstancesUserVpc) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesV2ResponseBodyInstancesUserVpc) GoString() string {
	return s.String()
}

func (s *ListInstancesV2ResponseBodyInstancesUserVpc) SetSecurityGroupId(v string) *ListInstancesV2ResponseBodyInstancesUserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *ListInstancesV2ResponseBodyInstancesUserVpc) SetVSwitchId(v string) *ListInstancesV2ResponseBodyInstancesUserVpc {
	s.VSwitchId = &v
	return s
}

func (s *ListInstancesV2ResponseBodyInstancesUserVpc) SetVpcId(v string) *ListInstancesV2ResponseBodyInstancesUserVpc {
	s.VpcId = &v
	return s
}

type ListInstancesV2Response struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInstancesV2ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInstancesV2Response) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesV2Response) GoString() string {
	return s.String()
}

func (s *ListInstancesV2Response) SetHeaders(v map[string]*string) *ListInstancesV2Response {
	s.Headers = v
	return s
}

func (s *ListInstancesV2Response) SetBody(v *ListInstancesV2ResponseBody) *ListInstancesV2Response {
	s.Body = v
	return s
}

type StartInstanceResponseBody struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartInstanceResponseBody) SetInstanceId(v string) *StartInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *StartInstanceResponseBody) SetRequestId(v string) *StartInstanceResponseBody {
	s.RequestId = &v
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

type StartInstanceV2ResponseBody struct {
	// 工作空间Id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartInstanceV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceV2ResponseBody) GoString() string {
	return s.String()
}

func (s *StartInstanceV2ResponseBody) SetInstanceId(v string) *StartInstanceV2ResponseBody {
	s.InstanceId = &v
	return s
}

func (s *StartInstanceV2ResponseBody) SetRequestId(v string) *StartInstanceV2ResponseBody {
	s.RequestId = &v
	return s
}

type StartInstanceV2Response struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartInstanceV2ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartInstanceV2Response) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceV2Response) GoString() string {
	return s.String()
}

func (s *StartInstanceV2Response) SetHeaders(v map[string]*string) *StartInstanceV2Response {
	s.Headers = v
	return s
}

func (s *StartInstanceV2Response) SetBody(v *StartInstanceV2ResponseBody) *StartInstanceV2Response {
	s.Body = v
	return s
}

type StopInstanceRequest struct {
	// 是否保存镜像后停止
	SaveImage *bool `json:"SaveImage,omitempty" xml:"SaveImage,omitempty"`
}

func (s StopInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StopInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopInstanceRequest) SetSaveImage(v bool) *StopInstanceRequest {
	s.SaveImage = &v
	return s
}

type StopInstanceResponseBody struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopInstanceResponseBody) SetInstanceId(v string) *StopInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *StopInstanceResponseBody) SetRequestId(v string) *StopInstanceResponseBody {
	s.RequestId = &v
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

type StopInstanceV2Request struct {
	SaveImage *bool `json:"SaveImage,omitempty" xml:"SaveImage,omitempty"`
}

func (s StopInstanceV2Request) String() string {
	return tea.Prettify(s)
}

func (s StopInstanceV2Request) GoString() string {
	return s.String()
}

func (s *StopInstanceV2Request) SetSaveImage(v bool) *StopInstanceV2Request {
	s.SaveImage = &v
	return s
}

type StopInstanceV2ResponseBody struct {
	// 工作空间Id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopInstanceV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopInstanceV2ResponseBody) GoString() string {
	return s.String()
}

func (s *StopInstanceV2ResponseBody) SetInstanceId(v string) *StopInstanceV2ResponseBody {
	s.InstanceId = &v
	return s
}

func (s *StopInstanceV2ResponseBody) SetRequestId(v string) *StopInstanceV2ResponseBody {
	s.RequestId = &v
	return s
}

type StopInstanceV2Response struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopInstanceV2ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopInstanceV2Response) String() string {
	return tea.Prettify(s)
}

func (s StopInstanceV2Response) GoString() string {
	return s.String()
}

func (s *StopInstanceV2Response) SetHeaders(v map[string]*string) *StopInstanceV2Response {
	s.Headers = v
	return s
}

func (s *StopInstanceV2Response) SetBody(v *StopInstanceV2ResponseBody) *StopInstanceV2Response {
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
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例镜像ID
	InstanceSnapshotId *string `json:"InstanceSnapshotId,omitempty" xml:"InstanceSnapshotId,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateInstanceSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceSnapshotResponseBody) SetInstanceId(v string) *UpdateInstanceSnapshotResponseBody {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceSnapshotResponseBody) SetInstanceSnapshotId(v string) *UpdateInstanceSnapshotResponseBody {
	s.InstanceSnapshotId = &v
	return s
}

func (s *UpdateInstanceSnapshotResponseBody) SetRequestId(v string) *UpdateInstanceSnapshotResponseBody {
	s.RequestId = &v
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

type UpdateInstanceV2Request struct {
	// 实例计算类型
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	// 工作空间内是否他人可见
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// 累计运行时间（ms）
	AccumulatedRunningTimeInMs *int64 `json:"AccumulatedRunningTimeInMs,omitempty" xml:"AccumulatedRunningTimeInMs,omitempty"`
	// 数据集集合
	Datasets []*UpdateInstanceV2RequestDatasets `json:"Datasets,omitempty" xml:"Datasets,omitempty" type:"Repeated"`
	// 实例对应的Ecs规格
	EcsSpec *string `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	// 环境变量
	EnvironmentVariables map[string]*string `json:"EnvironmentVariables,omitempty" xml:"EnvironmentVariables,omitempty"`
	// 实例创建时间
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 实例修改时间
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// 镜像Id
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// 镜像地址
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// 定时关机任务
	InstanceShutdownTimer *UpdateInstanceV2RequestInstanceShutdownTimer `json:"InstanceShutdownTimer,omitempty" xml:"InstanceShutdownTimer,omitempty" type:"Struct"`
	// 实例Url
	InstanceUrl *string `json:"InstanceUrl,omitempty" xml:"InstanceUrl,omitempty"`
	// Jupyterlab Url
	JupyterlabUrl *string `json:"JupyterlabUrl,omitempty" xml:"JupyterlabUrl,omitempty"`
	// 最新保存的用户镜像
	LatestSnapshot *UpdateInstanceV2RequestLatestSnapshot `json:"LatestSnapshot,omitempty" xml:"LatestSnapshot,omitempty" type:"Struct"`
	// 支付类型
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// 实例错误代码
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// 实例错误原因
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// 实例状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// terminal url
	TerminalUrl *string `json:"TerminalUrl,omitempty" xml:"TerminalUrl,omitempty"`
	// 用户Id
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// user vpc配置
	UserVpc *UpdateInstanceV2RequestUserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	// web ide url
	WebIDEUrl *string `json:"WebIDEUrl,omitempty" xml:"WebIDEUrl,omitempty"`
	// 工作空间Id
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// 工作空间名称
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s UpdateInstanceV2Request) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceV2Request) GoString() string {
	return s.String()
}

func (s *UpdateInstanceV2Request) SetAcceleratorType(v string) *UpdateInstanceV2Request {
	s.AcceleratorType = &v
	return s
}

func (s *UpdateInstanceV2Request) SetAccessibility(v string) *UpdateInstanceV2Request {
	s.Accessibility = &v
	return s
}

func (s *UpdateInstanceV2Request) SetAccumulatedRunningTimeInMs(v int64) *UpdateInstanceV2Request {
	s.AccumulatedRunningTimeInMs = &v
	return s
}

func (s *UpdateInstanceV2Request) SetDatasets(v []*UpdateInstanceV2RequestDatasets) *UpdateInstanceV2Request {
	s.Datasets = v
	return s
}

func (s *UpdateInstanceV2Request) SetEcsSpec(v string) *UpdateInstanceV2Request {
	s.EcsSpec = &v
	return s
}

func (s *UpdateInstanceV2Request) SetEnvironmentVariables(v map[string]*string) *UpdateInstanceV2Request {
	s.EnvironmentVariables = v
	return s
}

func (s *UpdateInstanceV2Request) SetGmtCreateTime(v string) *UpdateInstanceV2Request {
	s.GmtCreateTime = &v
	return s
}

func (s *UpdateInstanceV2Request) SetGmtModifiedTime(v string) *UpdateInstanceV2Request {
	s.GmtModifiedTime = &v
	return s
}

func (s *UpdateInstanceV2Request) SetImageId(v string) *UpdateInstanceV2Request {
	s.ImageId = &v
	return s
}

func (s *UpdateInstanceV2Request) SetImageName(v string) *UpdateInstanceV2Request {
	s.ImageName = &v
	return s
}

func (s *UpdateInstanceV2Request) SetImageUrl(v string) *UpdateInstanceV2Request {
	s.ImageUrl = &v
	return s
}

func (s *UpdateInstanceV2Request) SetInstanceId(v string) *UpdateInstanceV2Request {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceV2Request) SetInstanceName(v string) *UpdateInstanceV2Request {
	s.InstanceName = &v
	return s
}

func (s *UpdateInstanceV2Request) SetInstanceShutdownTimer(v *UpdateInstanceV2RequestInstanceShutdownTimer) *UpdateInstanceV2Request {
	s.InstanceShutdownTimer = v
	return s
}

func (s *UpdateInstanceV2Request) SetInstanceUrl(v string) *UpdateInstanceV2Request {
	s.InstanceUrl = &v
	return s
}

func (s *UpdateInstanceV2Request) SetJupyterlabUrl(v string) *UpdateInstanceV2Request {
	s.JupyterlabUrl = &v
	return s
}

func (s *UpdateInstanceV2Request) SetLatestSnapshot(v *UpdateInstanceV2RequestLatestSnapshot) *UpdateInstanceV2Request {
	s.LatestSnapshot = v
	return s
}

func (s *UpdateInstanceV2Request) SetPaymentType(v string) *UpdateInstanceV2Request {
	s.PaymentType = &v
	return s
}

func (s *UpdateInstanceV2Request) SetReasonCode(v string) *UpdateInstanceV2Request {
	s.ReasonCode = &v
	return s
}

func (s *UpdateInstanceV2Request) SetReasonMessage(v string) *UpdateInstanceV2Request {
	s.ReasonMessage = &v
	return s
}

func (s *UpdateInstanceV2Request) SetStatus(v string) *UpdateInstanceV2Request {
	s.Status = &v
	return s
}

func (s *UpdateInstanceV2Request) SetTerminalUrl(v string) *UpdateInstanceV2Request {
	s.TerminalUrl = &v
	return s
}

func (s *UpdateInstanceV2Request) SetUserId(v string) *UpdateInstanceV2Request {
	s.UserId = &v
	return s
}

func (s *UpdateInstanceV2Request) SetUserVpc(v *UpdateInstanceV2RequestUserVpc) *UpdateInstanceV2Request {
	s.UserVpc = v
	return s
}

func (s *UpdateInstanceV2Request) SetWebIDEUrl(v string) *UpdateInstanceV2Request {
	s.WebIDEUrl = &v
	return s
}

func (s *UpdateInstanceV2Request) SetWorkspaceId(v string) *UpdateInstanceV2Request {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateInstanceV2Request) SetWorkspaceName(v string) *UpdateInstanceV2Request {
	s.WorkspaceName = &v
	return s
}

type UpdateInstanceV2RequestDatasets struct {
	// 数据集Id
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// 容器内挂载路径
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
}

func (s UpdateInstanceV2RequestDatasets) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceV2RequestDatasets) GoString() string {
	return s.String()
}

func (s *UpdateInstanceV2RequestDatasets) SetDatasetId(v string) *UpdateInstanceV2RequestDatasets {
	s.DatasetId = &v
	return s
}

func (s *UpdateInstanceV2RequestDatasets) SetMountPath(v string) *UpdateInstanceV2RequestDatasets {
	s.MountPath = &v
	return s
}

type UpdateInstanceV2RequestInstanceShutdownTimer struct {
	// 设定关机时间
	DueTime *string `json:"DueTime,omitempty" xml:"DueTime,omitempty"`
	// 创建时间
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 修改时间
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 剩余关机时间（ms）
	RemainingTimeInMs *int64 `json:"RemainingTimeInMs,omitempty" xml:"RemainingTimeInMs,omitempty"`
}

func (s UpdateInstanceV2RequestInstanceShutdownTimer) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceV2RequestInstanceShutdownTimer) GoString() string {
	return s.String()
}

func (s *UpdateInstanceV2RequestInstanceShutdownTimer) SetDueTime(v string) *UpdateInstanceV2RequestInstanceShutdownTimer {
	s.DueTime = &v
	return s
}

func (s *UpdateInstanceV2RequestInstanceShutdownTimer) SetGmtCreateTime(v string) *UpdateInstanceV2RequestInstanceShutdownTimer {
	s.GmtCreateTime = &v
	return s
}

func (s *UpdateInstanceV2RequestInstanceShutdownTimer) SetGmtModifiedTime(v string) *UpdateInstanceV2RequestInstanceShutdownTimer {
	s.GmtModifiedTime = &v
	return s
}

func (s *UpdateInstanceV2RequestInstanceShutdownTimer) SetInstanceId(v string) *UpdateInstanceV2RequestInstanceShutdownTimer {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceV2RequestInstanceShutdownTimer) SetRemainingTimeInMs(v int64) *UpdateInstanceV2RequestInstanceShutdownTimer {
	s.RemainingTimeInMs = &v
	return s
}

type UpdateInstanceV2RequestLatestSnapshot struct {
	// 快照创建时间
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 快照修改时间
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// 镜像Id
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// 镜像Url
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// 镜像仓库Url
	RepositoryUrl *string `json:"RepositoryUrl,omitempty" xml:"RepositoryUrl,omitempty"`
}

func (s UpdateInstanceV2RequestLatestSnapshot) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceV2RequestLatestSnapshot) GoString() string {
	return s.String()
}

func (s *UpdateInstanceV2RequestLatestSnapshot) SetGmtCreateTime(v string) *UpdateInstanceV2RequestLatestSnapshot {
	s.GmtCreateTime = &v
	return s
}

func (s *UpdateInstanceV2RequestLatestSnapshot) SetGmtModifiedTime(v string) *UpdateInstanceV2RequestLatestSnapshot {
	s.GmtModifiedTime = &v
	return s
}

func (s *UpdateInstanceV2RequestLatestSnapshot) SetImageId(v string) *UpdateInstanceV2RequestLatestSnapshot {
	s.ImageId = &v
	return s
}

func (s *UpdateInstanceV2RequestLatestSnapshot) SetImageName(v string) *UpdateInstanceV2RequestLatestSnapshot {
	s.ImageName = &v
	return s
}

func (s *UpdateInstanceV2RequestLatestSnapshot) SetImageUrl(v string) *UpdateInstanceV2RequestLatestSnapshot {
	s.ImageUrl = &v
	return s
}

func (s *UpdateInstanceV2RequestLatestSnapshot) SetRepositoryUrl(v string) *UpdateInstanceV2RequestLatestSnapshot {
	s.RepositoryUrl = &v
	return s
}

type UpdateInstanceV2RequestUserVpc struct {
	// Security Group Id
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// VSwitch Id
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// Vpc Id
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s UpdateInstanceV2RequestUserVpc) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceV2RequestUserVpc) GoString() string {
	return s.String()
}

func (s *UpdateInstanceV2RequestUserVpc) SetSecurityGroupId(v string) *UpdateInstanceV2RequestUserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *UpdateInstanceV2RequestUserVpc) SetVSwitchId(v string) *UpdateInstanceV2RequestUserVpc {
	s.VSwitchId = &v
	return s
}

func (s *UpdateInstanceV2RequestUserVpc) SetVpcId(v string) *UpdateInstanceV2RequestUserVpc {
	s.VpcId = &v
	return s
}

type UpdateInstanceV2ResponseBody struct {
	// 工作空间Id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateInstanceV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceV2ResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceV2ResponseBody) SetInstanceId(v string) *UpdateInstanceV2ResponseBody {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceV2ResponseBody) SetRequestId(v string) *UpdateInstanceV2ResponseBody {
	s.RequestId = &v
	return s
}

type UpdateInstanceV2Response struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateInstanceV2ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateInstanceV2Response) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceV2Response) GoString() string {
	return s.String()
}

func (s *UpdateInstanceV2Response) SetHeaders(v map[string]*string) *UpdateInstanceV2Response {
	s.Headers = v
	return s
}

func (s *UpdateInstanceV2Response) SetBody(v *UpdateInstanceV2ResponseBody) *UpdateInstanceV2Response {
	s.Body = v
	return s
}

type UpdateV3InstanceByUserRequest struct {
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateV3InstanceByUserRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateV3InstanceByUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateV3InstanceByUserRequest) SetUserId(v string) *UpdateV3InstanceByUserRequest {
	s.UserId = &v
	return s
}

func (s *UpdateV3InstanceByUserRequest) SetWorkspaceId(v string) *UpdateV3InstanceByUserRequest {
	s.WorkspaceId = &v
	return s
}

type UpdateV3InstanceByUserResponseBody struct {
	// 状态码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// http状态码
	HttpStatusCode *int32   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	IdList         []*int64 `json:"IdList,omitempty" xml:"IdList,omitempty" type:"Repeated"`
	// 说明
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 成功标志
	Success     *bool     `json:"Success,omitempty" xml:"Success,omitempty"`
	UserSet     []*string `json:"UserSet,omitempty" xml:"UserSet,omitempty" type:"Repeated"`
	WorkspaceId *string   `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateV3InstanceByUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateV3InstanceByUserResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateV3InstanceByUserResponseBody) SetCode(v string) *UpdateV3InstanceByUserResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateV3InstanceByUserResponseBody) SetHttpStatusCode(v int32) *UpdateV3InstanceByUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateV3InstanceByUserResponseBody) SetIdList(v []*int64) *UpdateV3InstanceByUserResponseBody {
	s.IdList = v
	return s
}

func (s *UpdateV3InstanceByUserResponseBody) SetMessage(v string) *UpdateV3InstanceByUserResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateV3InstanceByUserResponseBody) SetRequestId(v string) *UpdateV3InstanceByUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateV3InstanceByUserResponseBody) SetSuccess(v bool) *UpdateV3InstanceByUserResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateV3InstanceByUserResponseBody) SetUserSet(v []*string) *UpdateV3InstanceByUserResponseBody {
	s.UserSet = v
	return s
}

func (s *UpdateV3InstanceByUserResponseBody) SetWorkspaceId(v string) *UpdateV3InstanceByUserResponseBody {
	s.WorkspaceId = &v
	return s
}

type UpdateV3InstanceByUserResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateV3InstanceByUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateV3InstanceByUserResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateV3InstanceByUserResponse) GoString() string {
	return s.String()
}

func (s *UpdateV3InstanceByUserResponse) SetHeaders(v map[string]*string) *UpdateV3InstanceByUserResponse {
	s.Headers = v
	return s
}

func (s *UpdateV3InstanceByUserResponse) SetBody(v *UpdateV3InstanceByUserResponseBody) *UpdateV3InstanceByUserResponse {
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
	if !tea.BoolValue(util.IsUnset(request.DatasetList)) {
		body["DatasetList"] = request.DatasetList
	}

	if !tea.BoolValue(util.IsUnset(request.EcsSpec)) {
		body["EcsSpec"] = request.EcsSpec
	}

	if !tea.BoolValue(util.IsUnset(request.Environments)) {
		body["Environments"] = request.Environments
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		body["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["ImageUrl"] = request.ImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.IsPublic)) {
		body["IsPublic"] = request.IsPublic
	}

	if !tea.BoolValue(util.IsUnset(request.NasFileSystemId)) {
		body["NasFileSystemId"] = request.NasFileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["UserName"] = request.UserName
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.UserVpc))) {
		body["UserVpc"] = request.UserVpc
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstance"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ScheduleTime)) {
		body["ScheduleTime"] = request.ScheduleTime
	}

	if !tea.BoolValue(util.IsUnset(request.TtlInMillis)) {
		body["TtlInMillis"] = request.TtlInMillis
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstanceShutdownTimer"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(InstanceId) + "/shutdownTimer"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInstanceShutdownTimerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateInstanceShutdownTimerV2(InstanceId *string, request *CreateInstanceShutdownTimerV2Request) (_result *CreateInstanceShutdownTimerV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateInstanceShutdownTimerV2Response{}
	_body, _err := client.CreateInstanceShutdownTimerV2WithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateInstanceShutdownTimerV2WithOptions(InstanceId *string, request *CreateInstanceShutdownTimerV2Request, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateInstanceShutdownTimerV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DueTime)) {
		body["DueTime"] = request.DueTime
	}

	if !tea.BoolValue(util.IsUnset(request.RemainingTimeInMs)) {
		body["RemainingTimeInMs"] = request.RemainingTimeInMs
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstanceShutdownTimerV2"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instances/" + tea.StringValue(InstanceId) + "/shutdowntimer"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInstanceShutdownTimerV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceSnapshotDescription)) {
		body["InstanceSnapshotDescription"] = request.InstanceSnapshotDescription
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceSnapshotName)) {
		body["InstanceSnapshotName"] = request.InstanceSnapshotName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceSnapshotRepoUrl)) {
		body["InstanceSnapshotRepoUrl"] = request.InstanceSnapshotRepoUrl
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstanceSnapshot"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(InstanceId) + "/snapshots"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInstanceSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateInstanceSnapshotV2(InstanceId *string, request *CreateInstanceSnapshotV2Request) (_result *CreateInstanceSnapshotV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateInstanceSnapshotV2Response{}
	_body, _err := client.CreateInstanceSnapshotV2WithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateInstanceSnapshotV2WithOptions(InstanceId *string, request *CreateInstanceSnapshotV2Request, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateInstanceSnapshotV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["ImageUrl"] = request.ImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotDescription)) {
		body["SnapshotDescription"] = request.SnapshotDescription
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotName)) {
		body["SnapshotName"] = request.SnapshotName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstanceSnapshotV2"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instances/" + tea.StringValue(InstanceId) + "/snapshots"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInstanceSnapshotV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateInstanceV2(request *CreateInstanceV2Request) (_result *CreateInstanceV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateInstanceV2Response{}
	_body, _err := client.CreateInstanceV2WithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateInstanceV2WithOptions(request *CreateInstanceV2Request, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateInstanceV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Accessibility)) {
		body["Accessibility"] = request.Accessibility
	}

	if !tea.BoolValue(util.IsUnset(request.Datasets)) {
		body["Datasets"] = request.Datasets
	}

	if !tea.BoolValue(util.IsUnset(request.EcsSpec)) {
		body["EcsSpec"] = request.EcsSpec
	}

	if !tea.BoolValue(util.IsUnset(request.EnvironmentVariables)) {
		body["EnvironmentVariables"] = request.EnvironmentVariables
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		body["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["ImageUrl"] = request.ImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.UserVpc))) {
		body["UserVpc"] = request.UserVpc
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstanceV2"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instances"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInstanceV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
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
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstance"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(InstanceId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstanceShutdownTimer"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(InstanceId) + "/shutdownTimer"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteInstanceShutdownTimerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteInstanceShutdownTimerV2(InstanceId *string) (_result *DeleteInstanceShutdownTimerV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteInstanceShutdownTimerV2Response{}
	_body, _err := client.DeleteInstanceShutdownTimerV2WithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteInstanceShutdownTimerV2WithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteInstanceShutdownTimerV2Response, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstanceShutdownTimerV2"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instances/" + tea.StringValue(InstanceId) + "/shutdowntimer"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteInstanceShutdownTimerV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
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
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	InstanceSnapshotId = openapiutil.GetEncodeParam(InstanceSnapshotId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstanceSnapshot"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(InstanceId) + "/snapshots/" + tea.StringValue(InstanceSnapshotId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteInstanceSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteInstanceSnapshotV2(InstanceId *string, SnapshotId *string) (_result *DeleteInstanceSnapshotV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteInstanceSnapshotV2Response{}
	_body, _err := client.DeleteInstanceSnapshotV2WithOptions(InstanceId, SnapshotId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteInstanceSnapshotV2WithOptions(InstanceId *string, SnapshotId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteInstanceSnapshotV2Response, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	SnapshotId = openapiutil.GetEncodeParam(SnapshotId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstanceSnapshotV2"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instances/" + tea.StringValue(InstanceId) + "/snapshots/" + tea.StringValue(SnapshotId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteInstanceSnapshotV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteInstanceV2(InstanceId *string) (_result *DeleteInstanceV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteInstanceV2Response{}
	_body, _err := client.DeleteInstanceV2WithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteInstanceV2WithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteInstanceV2Response, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstanceV2"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instances/" + tea.StringValue(InstanceId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteInstanceV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("Foobar"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/foobar"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &FoobarResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Foobar1() (_result *Foobar1Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &Foobar1Response{}
	_body, _err := client.Foobar1WithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) Foobar1WithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *Foobar1Response, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("Foobar1"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/foobar1"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("any"),
	}
	_result = &Foobar1Response{}
	_body, _err := client.CallApi(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("GetAuthorization"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/authorization"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAuthorizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDashboardStatistics(request *GetDashboardStatisticsRequest) (_result *GetDashboardStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDashboardStatisticsResponse{}
	_body, _err := client.GetDashboardStatisticsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDashboardStatisticsWithOptions(request *GetDashboardStatisticsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDashboardStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDashboardStatistics"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/statistics/dashboard"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDashboardStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstance"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(InstanceId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstanceShutdownTimer"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(InstanceId) + "/shutdownTimer"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceShutdownTimerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstanceShutdownTimerV2(InstanceId *string) (_result *GetInstanceShutdownTimerV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceShutdownTimerV2Response{}
	_body, _err := client.GetInstanceShutdownTimerV2WithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstanceShutdownTimerV2WithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetInstanceShutdownTimerV2Response, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstanceShutdownTimerV2"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instances/" + tea.StringValue(InstanceId) + "/shutdowntimer"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceShutdownTimerV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
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
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	InstanceSnapshotId = openapiutil.GetEncodeParam(InstanceSnapshotId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstanceSnapshot"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(InstanceId) + "/snapshots/" + tea.StringValue(InstanceSnapshotId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstanceSnapshotV2(InstanceId *string, SnapshotId *string) (_result *GetInstanceSnapshotV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceSnapshotV2Response{}
	_body, _err := client.GetInstanceSnapshotV2WithOptions(InstanceId, SnapshotId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstanceSnapshotV2WithOptions(InstanceId *string, SnapshotId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetInstanceSnapshotV2Response, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	SnapshotId = openapiutil.GetEncodeParam(SnapshotId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstanceSnapshotV2"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instances/" + tea.StringValue(InstanceId) + "/snapshots/" + tea.StringValue(SnapshotId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceSnapshotV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstanceV2(InstanceId *string) (_result *GetInstanceV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceV2Response{}
	_body, _err := client.GetInstanceV2WithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstanceV2WithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetInstanceV2Response, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstanceV2"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instances/" + tea.StringValue(InstanceId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("GetInstancesStatistics"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/statistics/instances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstancesStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserConfigV2() (_result *GetUserConfigV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUserConfigV2Response{}
	_body, _err := client.GetUserConfigV2WithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserConfigV2WithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetUserConfigV2Response, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserConfigV2"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/userconfig"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserConfigV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("ListEcsSpecs"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/ecsSpecs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEcsSpecsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEcsSpecsV2(request *ListEcsSpecsV2Request) (_result *ListEcsSpecsV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEcsSpecsV2Response{}
	_body, _err := client.ListEcsSpecsV2WithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEcsSpecsV2WithOptions(request *ListEcsSpecsV2Request, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListEcsSpecsV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorType)) {
		query["AcceleratorType"] = request.AcceleratorType
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEcsSpecsV2"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/ecsspecs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEcsSpecsV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.ImageNameContains)) {
		query["ImageNameContains"] = request.ImageNameContains
	}

	if !tea.BoolValue(util.IsUnset(request.ImageTypeEquals)) {
		query["ImageTypeEquals"] = request.ImageTypeEquals
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		query["Product"] = request.Product
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListImages"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/images"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstanceSnapshotV2(InstanceId *string, request *ListInstanceSnapshotV2Request) (_result *ListInstanceSnapshotV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstanceSnapshotV2Response{}
	_body, _err := client.ListInstanceSnapshotV2WithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstanceSnapshotV2WithOptions(InstanceId *string, request *ListInstanceSnapshotV2Request, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstanceSnapshotV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstanceSnapshotV2"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instances/" + tea.StringValue(InstanceId) + "/snapshots"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstanceSnapshotV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
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
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstanceSnapshots"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(InstanceId) + "/snapshots"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstanceSnapshotsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstanceStatisticsV2(request *ListInstanceStatisticsV2Request) (_result *ListInstanceStatisticsV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstanceStatisticsV2Response{}
	_body, _err := client.ListInstanceStatisticsV2WithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstanceStatisticsV2WithOptions(request *ListInstanceStatisticsV2Request, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstanceStatisticsV2Response, _err error) {
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
	params := &openapi.Params{
		Action:      tea.String("ListInstanceStatisticsV2"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instancestatistics"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstanceStatisticsV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.InWorkspace)) {
		query["InWorkspace"] = request.InWorkspace
	}

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

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrder)) {
		query["SortOrder"] = request.SortOrder
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceIdEquals)) {
		query["WorkspaceIdEquals"] = request.WorkspaceIdEquals
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstances"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstancesStatus(request *ListInstancesStatusRequest) (_result *ListInstancesStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstancesStatusResponse{}
	_body, _err := client.ListInstancesStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstancesStatusWithOptions(request *ListInstancesStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstancesStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstancesStatus"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/statuses/instances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstancesStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstancesV2(request *ListInstancesV2Request) (_result *ListInstancesV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstancesV2Response{}
	_body, _err := client.ListInstancesV2WithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstancesV2WithOptions(request *ListInstancesV2Request, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstancesV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstancesV2"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstancesV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
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
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StartInstance"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(InstanceId) + "/start"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StartInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartInstanceV2(InstanceId *string) (_result *StartInstanceV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartInstanceV2Response{}
	_body, _err := client.StartInstanceV2WithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartInstanceV2WithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartInstanceV2Response, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StartInstanceV2"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instances/" + tea.StringValue(InstanceId) + "/start"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StartInstanceV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopInstance(InstanceId *string, request *StopInstanceRequest) (_result *StopInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopInstanceResponse{}
	_body, _err := client.StopInstanceWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopInstanceWithOptions(InstanceId *string, request *StopInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SaveImage)) {
		query["SaveImage"] = request.SaveImage
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopInstance"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(InstanceId) + "/stop"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopInstanceV2(InstanceId *string, request *StopInstanceV2Request) (_result *StopInstanceV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopInstanceV2Response{}
	_body, _err := client.StopInstanceV2WithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopInstanceV2WithOptions(InstanceId *string, request *StopInstanceV2Request, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopInstanceV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SaveImage)) {
		body["SaveImage"] = request.SaveImage
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StopInstanceV2"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instances/" + tea.StringValue(InstanceId) + "/stop"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopInstanceV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
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
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstance"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(InstanceId)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	InstanceSnapshotId = openapiutil.GetEncodeParam(InstanceSnapshotId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceSnapshotName)) {
		body["InstanceSnapshotName"] = request.InstanceSnapshotName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstanceSnapshot"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(InstanceId) + "/snapshots/" + tea.StringValue(InstanceSnapshotId)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInstanceSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateInstanceV2(InstanceId *string, request *UpdateInstanceV2Request) (_result *UpdateInstanceV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateInstanceV2Response{}
	_body, _err := client.UpdateInstanceV2WithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateInstanceV2WithOptions(InstanceId *string, request *UpdateInstanceV2Request, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateInstanceV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorType)) {
		body["AcceleratorType"] = request.AcceleratorType
	}

	if !tea.BoolValue(util.IsUnset(request.Accessibility)) {
		body["Accessibility"] = request.Accessibility
	}

	if !tea.BoolValue(util.IsUnset(request.AccumulatedRunningTimeInMs)) {
		body["AccumulatedRunningTimeInMs"] = request.AccumulatedRunningTimeInMs
	}

	if !tea.BoolValue(util.IsUnset(request.Datasets)) {
		body["Datasets"] = request.Datasets
	}

	if !tea.BoolValue(util.IsUnset(request.EcsSpec)) {
		body["EcsSpec"] = request.EcsSpec
	}

	if !tea.BoolValue(util.IsUnset(request.EnvironmentVariables)) {
		body["EnvironmentVariables"] = request.EnvironmentVariables
	}

	if !tea.BoolValue(util.IsUnset(request.GmtCreateTime)) {
		body["GmtCreateTime"] = request.GmtCreateTime
	}

	if !tea.BoolValue(util.IsUnset(request.GmtModifiedTime)) {
		body["GmtModifiedTime"] = request.GmtModifiedTime
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		body["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.ImageName)) {
		body["ImageName"] = request.ImageName
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["ImageUrl"] = request.ImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.InstanceShutdownTimer))) {
		body["InstanceShutdownTimer"] = request.InstanceShutdownTimer
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceUrl)) {
		body["InstanceUrl"] = request.InstanceUrl
	}

	if !tea.BoolValue(util.IsUnset(request.JupyterlabUrl)) {
		body["JupyterlabUrl"] = request.JupyterlabUrl
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.LatestSnapshot))) {
		body["LatestSnapshot"] = request.LatestSnapshot
	}

	if !tea.BoolValue(util.IsUnset(request.PaymentType)) {
		body["PaymentType"] = request.PaymentType
	}

	if !tea.BoolValue(util.IsUnset(request.ReasonCode)) {
		body["ReasonCode"] = request.ReasonCode
	}

	if !tea.BoolValue(util.IsUnset(request.ReasonMessage)) {
		body["ReasonMessage"] = request.ReasonMessage
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TerminalUrl)) {
		body["TerminalUrl"] = request.TerminalUrl
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.UserVpc))) {
		body["UserVpc"] = request.UserVpc
	}

	if !tea.BoolValue(util.IsUnset(request.WebIDEUrl)) {
		body["WebIDEUrl"] = request.WebIDEUrl
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceName)) {
		body["WorkspaceName"] = request.WorkspaceName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstanceV2"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instances/" + tea.StringValue(InstanceId)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInstanceV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateV3InstanceByUser(request *UpdateV3InstanceByUserRequest) (_result *UpdateV3InstanceByUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateV3InstanceByUserResponse{}
	_body, _err := client.UpdateV3InstanceByUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateV3InstanceByUserWithOptions(request *UpdateV3InstanceByUserRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateV3InstanceByUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateV3InstanceByUser"),
		Version:     tea.String("2021-02-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/migrate/user"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateV3InstanceByUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
