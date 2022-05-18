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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateInstanceResponse) SetStatusCode(v int32) *CreateInstanceResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateInstanceShutdownTimerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateInstanceShutdownTimerResponse) SetStatusCode(v int32) *CreateInstanceShutdownTimerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceShutdownTimerResponse) SetBody(v *CreateInstanceShutdownTimerResponseBody) *CreateInstanceShutdownTimerResponse {
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateInstanceSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateInstanceSnapshotResponse) SetStatusCode(v int32) *CreateInstanceSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceSnapshotResponse) SetBody(v *CreateInstanceSnapshotResponseBody) *CreateInstanceSnapshotResponse {
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteInstanceResponse) SetStatusCode(v int32) *DeleteInstanceResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteInstanceShutdownTimerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteInstanceShutdownTimerResponse) SetStatusCode(v int32) *DeleteInstanceShutdownTimerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstanceShutdownTimerResponse) SetBody(v *DeleteInstanceShutdownTimerResponseBody) *DeleteInstanceShutdownTimerResponse {
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteInstanceSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteInstanceSnapshotResponse) SetStatusCode(v int32) *DeleteInstanceSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstanceSnapshotResponse) SetBody(v *DeleteInstanceSnapshotResponseBody) *DeleteInstanceSnapshotResponse {
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetAuthorizationResponse) SetStatusCode(v int32) *GetAuthorizationResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDashboardStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetDashboardStatisticsResponse) SetStatusCode(v int32) *GetDashboardStatisticsResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetInstanceResponse) SetStatusCode(v int32) *GetInstanceResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetInstanceShutdownTimerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetInstanceShutdownTimerResponse) SetStatusCode(v int32) *GetInstanceShutdownTimerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceShutdownTimerResponse) SetBody(v *GetInstanceShutdownTimerResponseBody) *GetInstanceShutdownTimerResponse {
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetInstanceSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetInstanceSnapshotResponse) SetStatusCode(v int32) *GetInstanceSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceSnapshotResponse) SetBody(v *GetInstanceSnapshotResponseBody) *GetInstanceSnapshotResponse {
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetInstancesStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetInstancesStatisticsResponse) SetStatusCode(v int32) *GetInstancesStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstancesStatisticsResponse) SetBody(v *GetInstancesStatisticsResponseBody) *GetInstancesStatisticsResponse {
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListEcsSpecsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListEcsSpecsResponse) SetStatusCode(v int32) *ListEcsSpecsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEcsSpecsResponse) SetBody(v *ListEcsSpecsResponseBody) *ListEcsSpecsResponse {
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListImagesResponse) SetStatusCode(v int32) *ListImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListImagesResponse) SetBody(v *ListImagesResponseBody) *ListImagesResponse {
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListInstanceSnapshotsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListInstanceSnapshotsResponse) SetStatusCode(v int32) *ListInstanceSnapshotsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceSnapshotsResponse) SetBody(v *ListInstanceSnapshotsResponseBody) *ListInstanceSnapshotsResponse {
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListInstancesResponse) SetStatusCode(v int32) *ListInstancesResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListInstancesStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListInstancesStatusResponse) SetStatusCode(v int32) *ListInstancesStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstancesStatusResponse) SetBody(v *ListInstancesStatusResponseBody) *ListInstancesStatusResponse {
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *StartInstanceResponse) SetStatusCode(v int32) *StartInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartInstanceResponse) SetBody(v *StartInstanceResponseBody) *StartInstanceResponse {
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *StopInstanceResponse) SetStatusCode(v int32) *StopInstanceResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateInstanceResponse) SetStatusCode(v int32) *UpdateInstanceResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateInstanceSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateInstanceSnapshotResponse) SetStatusCode(v int32) *UpdateInstanceSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceSnapshotResponse) SetBody(v *UpdateInstanceSnapshotResponseBody) *UpdateInstanceSnapshotResponse {
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateV3InstanceByUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateV3InstanceByUserResponse) SetStatusCode(v int32) *UpdateV3InstanceByUserResponse {
	s.StatusCode = &v
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
