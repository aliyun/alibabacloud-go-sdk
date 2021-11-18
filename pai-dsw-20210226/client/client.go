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

type CreateImageRequest struct {
	// 镜像描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 实例名称
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 镜像名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 镜像仓库
	Repository *string `json:"Repository,omitempty" xml:"Repository,omitempty"`
}

func (s CreateImageRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateImageRequest) GoString() string {
	return s.String()
}

func (s *CreateImageRequest) SetDescription(v string) *CreateImageRequest {
	s.Description = &v
	return s
}

func (s *CreateImageRequest) SetInstanceId(v string) *CreateImageRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateImageRequest) SetName(v string) *CreateImageRequest {
	s.Name = &v
	return s
}

func (s *CreateImageRequest) SetRepository(v string) *CreateImageRequest {
	s.Repository = &v
	return s
}

type CreateImageResponseBody struct {
	// 保存的镜像Id
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateImageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageResponseBody) SetImageId(v string) *CreateImageResponseBody {
	s.ImageId = &v
	return s
}

func (s *CreateImageResponseBody) SetRequestId(v string) *CreateImageResponseBody {
	s.RequestId = &v
	return s
}

type CreateImageResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateImageResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateImageResponse) GoString() string {
	return s.String()
}

func (s *CreateImageResponse) SetHeaders(v map[string]*string) *CreateImageResponse {
	s.Headers = v
	return s
}

func (s *CreateImageResponse) SetBody(v *CreateImageResponseBody) *CreateImageResponse {
	s.Body = v
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

type GetInstanceImageResponseBody struct {
	Image *Image `json:"Image,omitempty" xml:"Image,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInstanceImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceImageResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceImageResponseBody) SetImage(v *Image) *GetInstanceImageResponseBody {
	s.Image = v
	return s
}

func (s *GetInstanceImageResponseBody) SetRequestId(v string) *GetInstanceImageResponseBody {
	s.RequestId = &v
	return s
}

type GetInstanceImageResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetInstanceImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInstanceImageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceImageResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceImageResponse) SetHeaders(v map[string]*string) *GetInstanceImageResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceImageResponse) SetBody(v *GetInstanceImageResponseBody) *GetInstanceImageResponse {
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

type GetInstanceTypeResponseBody struct {
	// cpu核数
	CpuCoreCount *int64 `json:"CpuCoreCount,omitempty" xml:"CpuCoreCount,omitempty"`
	// 折扣
	Discount *float32 `json:"Discount,omitempty" xml:"Discount,omitempty"`
	// 国内价格
	DomesticPrice *float32 `json:"DomesticPrice,omitempty" xml:"DomesticPrice,omitempty"`
	// GPU卡数
	GPUAmount *int64 `json:"GPUAmount,omitempty" xml:"GPUAmount,omitempty"`
	// GPU卡型
	GPUSpec *string `json:"GPUSpec,omitempty" xml:"GPUSpec,omitempty"`
	// 实例接收带宽
	InstanceBandwidthRx *int64 `json:"InstanceBandwidthRx,omitempty" xml:"InstanceBandwidthRx,omitempty"`
	// 实例发送带宽
	InstanceBandwidthTx *int64 `json:"InstanceBandwidthTx,omitempty" xml:"InstanceBandwidthTx,omitempty"`
	// 实例每秒发包数
	InstancePpsRx *int64 `json:"InstancePpsRx,omitempty" xml:"InstancePpsRx,omitempty"`
	// 实例每秒收包数
	InstancePpsTx *int64 `json:"InstancePpsTx,omitempty" xml:"InstancePpsTx,omitempty"`
	// 规格族
	InstanceTypeFamily *string `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
	// 实例类型id
	InstanceTypeId *string `json:"InstanceTypeId,omitempty" xml:"InstanceTypeId,omitempty"`
	// 是否国际站用户
	International *bool `json:"International,omitempty" xml:"International,omitempty"`
	// 存储盘容量
	LocalStorageCapacity *int64 `json:"LocalStorageCapacity,omitempty" xml:"LocalStorageCapacity,omitempty"`
	// 内存容量
	MemorySize *float32 `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	// 价格
	Price *float32 `json:"Price,omitempty" xml:"Price,omitempty"`
	// 价格（人民币）
	PriceCNY *float32 `json:"PriceCNY,omitempty" xml:"PriceCNY,omitempty"`
	// 价格（美元）
	PriceUSD *float32 `json:"PriceUSD,omitempty" xml:"PriceUSD,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 资源类型
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// 系统盘存储类型
	SystemDiskCategory *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	// 系统盘容量
	SystemDiskSize *int64 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s GetInstanceTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceTypeResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceTypeResponseBody) SetCpuCoreCount(v int64) *GetInstanceTypeResponseBody {
	s.CpuCoreCount = &v
	return s
}

func (s *GetInstanceTypeResponseBody) SetDiscount(v float32) *GetInstanceTypeResponseBody {
	s.Discount = &v
	return s
}

func (s *GetInstanceTypeResponseBody) SetDomesticPrice(v float32) *GetInstanceTypeResponseBody {
	s.DomesticPrice = &v
	return s
}

func (s *GetInstanceTypeResponseBody) SetGPUAmount(v int64) *GetInstanceTypeResponseBody {
	s.GPUAmount = &v
	return s
}

func (s *GetInstanceTypeResponseBody) SetGPUSpec(v string) *GetInstanceTypeResponseBody {
	s.GPUSpec = &v
	return s
}

func (s *GetInstanceTypeResponseBody) SetInstanceBandwidthRx(v int64) *GetInstanceTypeResponseBody {
	s.InstanceBandwidthRx = &v
	return s
}

func (s *GetInstanceTypeResponseBody) SetInstanceBandwidthTx(v int64) *GetInstanceTypeResponseBody {
	s.InstanceBandwidthTx = &v
	return s
}

func (s *GetInstanceTypeResponseBody) SetInstancePpsRx(v int64) *GetInstanceTypeResponseBody {
	s.InstancePpsRx = &v
	return s
}

func (s *GetInstanceTypeResponseBody) SetInstancePpsTx(v int64) *GetInstanceTypeResponseBody {
	s.InstancePpsTx = &v
	return s
}

func (s *GetInstanceTypeResponseBody) SetInstanceTypeFamily(v string) *GetInstanceTypeResponseBody {
	s.InstanceTypeFamily = &v
	return s
}

func (s *GetInstanceTypeResponseBody) SetInstanceTypeId(v string) *GetInstanceTypeResponseBody {
	s.InstanceTypeId = &v
	return s
}

func (s *GetInstanceTypeResponseBody) SetInternational(v bool) *GetInstanceTypeResponseBody {
	s.International = &v
	return s
}

func (s *GetInstanceTypeResponseBody) SetLocalStorageCapacity(v int64) *GetInstanceTypeResponseBody {
	s.LocalStorageCapacity = &v
	return s
}

func (s *GetInstanceTypeResponseBody) SetMemorySize(v float32) *GetInstanceTypeResponseBody {
	s.MemorySize = &v
	return s
}

func (s *GetInstanceTypeResponseBody) SetPrice(v float32) *GetInstanceTypeResponseBody {
	s.Price = &v
	return s
}

func (s *GetInstanceTypeResponseBody) SetPriceCNY(v float32) *GetInstanceTypeResponseBody {
	s.PriceCNY = &v
	return s
}

func (s *GetInstanceTypeResponseBody) SetPriceUSD(v float32) *GetInstanceTypeResponseBody {
	s.PriceUSD = &v
	return s
}

func (s *GetInstanceTypeResponseBody) SetRequestId(v string) *GetInstanceTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceTypeResponseBody) SetResourceType(v string) *GetInstanceTypeResponseBody {
	s.ResourceType = &v
	return s
}

func (s *GetInstanceTypeResponseBody) SetSystemDiskCategory(v string) *GetInstanceTypeResponseBody {
	s.SystemDiskCategory = &v
	return s
}

func (s *GetInstanceTypeResponseBody) SetSystemDiskSize(v int64) *GetInstanceTypeResponseBody {
	s.SystemDiskSize = &v
	return s
}

type GetInstanceTypeResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetInstanceTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInstanceTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceTypeResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceTypeResponse) SetHeaders(v map[string]*string) *GetInstanceTypeResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceTypeResponse) SetBody(v *GetInstanceTypeResponseBody) *GetInstanceTypeResponse {
	s.Body = v
	return s
}

type GetInstanceUrlResponseBody struct {
	// webide的链接
	Ide *string `json:"Ide,omitempty" xml:"Ide,omitempty"`
	// jupyterlab的链接
	Lab *string `json:"Lab,omitempty" xml:"Lab,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// terminal终端的链接
	Terminal *string `json:"Terminal,omitempty" xml:"Terminal,omitempty"`
}

func (s GetInstanceUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceUrlResponseBody) SetIde(v string) *GetInstanceUrlResponseBody {
	s.Ide = &v
	return s
}

func (s *GetInstanceUrlResponseBody) SetLab(v string) *GetInstanceUrlResponseBody {
	s.Lab = &v
	return s
}

func (s *GetInstanceUrlResponseBody) SetRequestId(v string) *GetInstanceUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceUrlResponseBody) SetTerminal(v string) *GetInstanceUrlResponseBody {
	s.Terminal = &v
	return s
}

type GetInstanceUrlResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetInstanceUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInstanceUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceUrlResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceUrlResponse) SetHeaders(v map[string]*string) *GetInstanceUrlResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceUrlResponse) SetBody(v *GetInstanceUrlResponseBody) *GetInstanceUrlResponse {
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

type GetUserConfigResponseBody struct {
	// 当前版本
	CurrentFeatureVersion *string `json:"CurrentFeatureVersion,omitempty" xml:"CurrentFeatureVersion,omitempty"`
	// 是否启用v2功能
	EnableEmrCluster *bool `json:"EnableEmrCluster,omitempty" xml:"EnableEmrCluster,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 是否显示特价版功能
	UseOnSaleVersion *bool `json:"UseOnSaleVersion,omitempty" xml:"UseOnSaleVersion,omitempty"`
	// 是否使用团队版功能（v21）
	UseV21Feature *bool `json:"UseV21Feature,omitempty" xml:"UseV21Feature,omitempty"`
}

func (s GetUserConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserConfigResponseBody) SetCurrentFeatureVersion(v string) *GetUserConfigResponseBody {
	s.CurrentFeatureVersion = &v
	return s
}

func (s *GetUserConfigResponseBody) SetEnableEmrCluster(v bool) *GetUserConfigResponseBody {
	s.EnableEmrCluster = &v
	return s
}

func (s *GetUserConfigResponseBody) SetRequestId(v string) *GetUserConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserConfigResponseBody) SetUseOnSaleVersion(v bool) *GetUserConfigResponseBody {
	s.UseOnSaleVersion = &v
	return s
}

func (s *GetUserConfigResponseBody) SetUseV21Feature(v bool) *GetUserConfigResponseBody {
	s.UseV21Feature = &v
	return s
}

type GetUserConfigResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUserConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserConfigResponse) GoString() string {
	return s.String()
}

func (s *GetUserConfigResponse) SetHeaders(v map[string]*string) *GetUserConfigResponse {
	s.Headers = v
	return s
}

func (s *GetUserConfigResponse) SetBody(v *GetUserConfigResponseBody) *GetUserConfigResponse {
	s.Body = v
	return s
}

type GetUserResourceAuthorizationStatusResponseBody struct {
	// 现金账户余额
	AccountBalance *float32 `json:"AccountBalance,omitempty" xml:"AccountBalance,omitempty"`
	// 金额是否充足
	AccountSufficient *bool `json:"AccountSufficient,omitempty" xml:"AccountSufficient,omitempty"`
	// 充值页面
	AccountTopUpPage *string `json:"AccountTopUpPage,omitempty" xml:"AccountTopUpPage,omitempty"`
	// 授权开通页面
	AllAuthorizationPage *string `json:"AllAuthorizationPage,omitempty" xml:"AllAuthorizationPage,omitempty"`
	// 购买页
	BuyPage *string `json:"BuyPage,omitempty" xml:"BuyPage,omitempty"`
	// 代金券金额
	CouponBalance *float32 `json:"CouponBalance,omitempty" xml:"CouponBalance,omitempty"`
	// 当前版本
	CurrentFeatureVersion *string `json:"CurrentFeatureVersion,omitempty" xml:"CurrentFeatureVersion,omitempty"`
	// 是否禁止金额验证
	DisableBalanceCheck *bool `json:"DisableBalanceCheck,omitempty" xml:"DisableBalanceCheck,omitempty"`
	// dsw默认角色授权页面
	DswDefaultAuthorizationPage *string `json:"DswDefaultAuthorizationPage,omitempty" xml:"DswDefaultAuthorizationPage,omitempty"`
	// 环境
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// ess开通页面
	EssConsolePage *string `json:"EssConsolePage,omitempty" xml:"EssConsolePage,omitempty"`
	// ess是否开通
	EssServiceAvailable *bool `json:"EssServiceAvailable,omitempty" xml:"EssServiceAvailable,omitempty"`
	// 是否通过购买验证
	HasAllAuthorization *bool `json:"HasAllAuthorization,omitempty" xml:"HasAllAuthorization,omitempty"`
	// 是否通过授权验证
	HasDswDefaultAuthorization *bool `json:"HasDswDefaultAuthorization,omitempty" xml:"HasDswDefaultAuthorization,omitempty"`
	// 是否国际站账号
	International *bool `json:"International,omitempty" xml:"International,omitempty"`
	// 是否子账号登录
	IsSubUser *bool `json:"IsSubUser,omitempty" xml:"IsSubUser,omitempty"`
	// nas控制台
	NasConsolePage *string `json:"NasConsolePage,omitempty" xml:"NasConsolePage,omitempty"`
	// 是否实名认证
	RealNameVerified *bool `json:"RealNameVerified,omitempty" xml:"RealNameVerified,omitempty"`
	// 实名认证页面
	RealNameVerifiedPage *string `json:"RealNameVerifiedPage,omitempty" xml:"RealNameVerifiedPage,omitempty"`
	// 地区
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 子账号授权开通页面
	SubUserAuthorizationPage *string `json:"SubUserAuthorizationPage,omitempty" xml:"SubUserAuthorizationPage,omitempty"`
	// 子账号是否授权通过
	SubUserAuthorized *bool `json:"SubUserAuthorized,omitempty" xml:"SubUserAuthorized,omitempty"`
	// 总金额
	TotalBalance *float32 `json:"TotalBalance,omitempty" xml:"TotalBalance,omitempty"`
}

func (s GetUserResourceAuthorizationStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserResourceAuthorizationStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResourceAuthorizationStatusResponseBody) SetAccountBalance(v float32) *GetUserResourceAuthorizationStatusResponseBody {
	s.AccountBalance = &v
	return s
}

func (s *GetUserResourceAuthorizationStatusResponseBody) SetAccountSufficient(v bool) *GetUserResourceAuthorizationStatusResponseBody {
	s.AccountSufficient = &v
	return s
}

func (s *GetUserResourceAuthorizationStatusResponseBody) SetAccountTopUpPage(v string) *GetUserResourceAuthorizationStatusResponseBody {
	s.AccountTopUpPage = &v
	return s
}

func (s *GetUserResourceAuthorizationStatusResponseBody) SetAllAuthorizationPage(v string) *GetUserResourceAuthorizationStatusResponseBody {
	s.AllAuthorizationPage = &v
	return s
}

func (s *GetUserResourceAuthorizationStatusResponseBody) SetBuyPage(v string) *GetUserResourceAuthorizationStatusResponseBody {
	s.BuyPage = &v
	return s
}

func (s *GetUserResourceAuthorizationStatusResponseBody) SetCouponBalance(v float32) *GetUserResourceAuthorizationStatusResponseBody {
	s.CouponBalance = &v
	return s
}

func (s *GetUserResourceAuthorizationStatusResponseBody) SetCurrentFeatureVersion(v string) *GetUserResourceAuthorizationStatusResponseBody {
	s.CurrentFeatureVersion = &v
	return s
}

func (s *GetUserResourceAuthorizationStatusResponseBody) SetDisableBalanceCheck(v bool) *GetUserResourceAuthorizationStatusResponseBody {
	s.DisableBalanceCheck = &v
	return s
}

func (s *GetUserResourceAuthorizationStatusResponseBody) SetDswDefaultAuthorizationPage(v string) *GetUserResourceAuthorizationStatusResponseBody {
	s.DswDefaultAuthorizationPage = &v
	return s
}

func (s *GetUserResourceAuthorizationStatusResponseBody) SetEnv(v string) *GetUserResourceAuthorizationStatusResponseBody {
	s.Env = &v
	return s
}

func (s *GetUserResourceAuthorizationStatusResponseBody) SetEssConsolePage(v string) *GetUserResourceAuthorizationStatusResponseBody {
	s.EssConsolePage = &v
	return s
}

func (s *GetUserResourceAuthorizationStatusResponseBody) SetEssServiceAvailable(v bool) *GetUserResourceAuthorizationStatusResponseBody {
	s.EssServiceAvailable = &v
	return s
}

func (s *GetUserResourceAuthorizationStatusResponseBody) SetHasAllAuthorization(v bool) *GetUserResourceAuthorizationStatusResponseBody {
	s.HasAllAuthorization = &v
	return s
}

func (s *GetUserResourceAuthorizationStatusResponseBody) SetHasDswDefaultAuthorization(v bool) *GetUserResourceAuthorizationStatusResponseBody {
	s.HasDswDefaultAuthorization = &v
	return s
}

func (s *GetUserResourceAuthorizationStatusResponseBody) SetInternational(v bool) *GetUserResourceAuthorizationStatusResponseBody {
	s.International = &v
	return s
}

func (s *GetUserResourceAuthorizationStatusResponseBody) SetIsSubUser(v bool) *GetUserResourceAuthorizationStatusResponseBody {
	s.IsSubUser = &v
	return s
}

func (s *GetUserResourceAuthorizationStatusResponseBody) SetNasConsolePage(v string) *GetUserResourceAuthorizationStatusResponseBody {
	s.NasConsolePage = &v
	return s
}

func (s *GetUserResourceAuthorizationStatusResponseBody) SetRealNameVerified(v bool) *GetUserResourceAuthorizationStatusResponseBody {
	s.RealNameVerified = &v
	return s
}

func (s *GetUserResourceAuthorizationStatusResponseBody) SetRealNameVerifiedPage(v string) *GetUserResourceAuthorizationStatusResponseBody {
	s.RealNameVerifiedPage = &v
	return s
}

func (s *GetUserResourceAuthorizationStatusResponseBody) SetRegion(v string) *GetUserResourceAuthorizationStatusResponseBody {
	s.Region = &v
	return s
}

func (s *GetUserResourceAuthorizationStatusResponseBody) SetRequestId(v string) *GetUserResourceAuthorizationStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserResourceAuthorizationStatusResponseBody) SetSubUserAuthorizationPage(v string) *GetUserResourceAuthorizationStatusResponseBody {
	s.SubUserAuthorizationPage = &v
	return s
}

func (s *GetUserResourceAuthorizationStatusResponseBody) SetSubUserAuthorized(v bool) *GetUserResourceAuthorizationStatusResponseBody {
	s.SubUserAuthorized = &v
	return s
}

func (s *GetUserResourceAuthorizationStatusResponseBody) SetTotalBalance(v float32) *GetUserResourceAuthorizationStatusResponseBody {
	s.TotalBalance = &v
	return s
}

type GetUserResourceAuthorizationStatusResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUserResourceAuthorizationStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserResourceAuthorizationStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserResourceAuthorizationStatusResponse) GoString() string {
	return s.String()
}

func (s *GetUserResourceAuthorizationStatusResponse) SetHeaders(v map[string]*string) *GetUserResourceAuthorizationStatusResponse {
	s.Headers = v
	return s
}

func (s *GetUserResourceAuthorizationStatusResponse) SetBody(v *GetUserResourceAuthorizationStatusResponseBody) *GetUserResourceAuthorizationStatusResponse {
	s.Body = v
	return s
}

type GetUserResourceStatusResponseBody struct {
	// 现金账户余额
	AccountBalance *float32 `json:"AccountBalance,omitempty" xml:"AccountBalance,omitempty"`
	// 金额是否充足
	AccountSufficient *bool `json:"AccountSufficient,omitempty" xml:"AccountSufficient,omitempty"`
	// 充值页面
	AccountTopUpPage *string `json:"AccountTopUpPage,omitempty" xml:"AccountTopUpPage,omitempty"`
	// 授权页面
	AllAuthorizationPage *string `json:"AllAuthorizationPage,omitempty" xml:"AllAuthorizationPage,omitempty"`
	// 代金券余额
	CouponBalance *float32 `json:"CouponBalance,omitempty" xml:"CouponBalance,omitempty"`
	// 环境
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// 是否通过购买条件验证
	HasAllAuthorization *bool `json:"HasAllAuthorization,omitempty" xml:"HasAllAuthorization,omitempty"`
	// 是否国际站账号
	International *bool `json:"International,omitempty" xml:"International,omitempty"`
	// 是否实名验证
	RealNameVerified *bool `json:"RealNameVerified,omitempty" xml:"RealNameVerified,omitempty"`
	// 实名验证页面
	RealNameVerifiedPage *string `json:"RealNameVerifiedPage,omitempty" xml:"RealNameVerifiedPage,omitempty"`
	// 地区
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 总余额
	TotalBalance *float32 `json:"TotalBalance,omitempty" xml:"TotalBalance,omitempty"`
}

func (s GetUserResourceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserResourceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResourceStatusResponseBody) SetAccountBalance(v float32) *GetUserResourceStatusResponseBody {
	s.AccountBalance = &v
	return s
}

func (s *GetUserResourceStatusResponseBody) SetAccountSufficient(v bool) *GetUserResourceStatusResponseBody {
	s.AccountSufficient = &v
	return s
}

func (s *GetUserResourceStatusResponseBody) SetAccountTopUpPage(v string) *GetUserResourceStatusResponseBody {
	s.AccountTopUpPage = &v
	return s
}

func (s *GetUserResourceStatusResponseBody) SetAllAuthorizationPage(v string) *GetUserResourceStatusResponseBody {
	s.AllAuthorizationPage = &v
	return s
}

func (s *GetUserResourceStatusResponseBody) SetCouponBalance(v float32) *GetUserResourceStatusResponseBody {
	s.CouponBalance = &v
	return s
}

func (s *GetUserResourceStatusResponseBody) SetEnv(v string) *GetUserResourceStatusResponseBody {
	s.Env = &v
	return s
}

func (s *GetUserResourceStatusResponseBody) SetHasAllAuthorization(v bool) *GetUserResourceStatusResponseBody {
	s.HasAllAuthorization = &v
	return s
}

func (s *GetUserResourceStatusResponseBody) SetInternational(v bool) *GetUserResourceStatusResponseBody {
	s.International = &v
	return s
}

func (s *GetUserResourceStatusResponseBody) SetRealNameVerified(v bool) *GetUserResourceStatusResponseBody {
	s.RealNameVerified = &v
	return s
}

func (s *GetUserResourceStatusResponseBody) SetRealNameVerifiedPage(v string) *GetUserResourceStatusResponseBody {
	s.RealNameVerifiedPage = &v
	return s
}

func (s *GetUserResourceStatusResponseBody) SetRegion(v string) *GetUserResourceStatusResponseBody {
	s.Region = &v
	return s
}

func (s *GetUserResourceStatusResponseBody) SetRequestId(v string) *GetUserResourceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserResourceStatusResponseBody) SetTotalBalance(v float32) *GetUserResourceStatusResponseBody {
	s.TotalBalance = &v
	return s
}

type GetUserResourceStatusResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUserResourceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserResourceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserResourceStatusResponse) GoString() string {
	return s.String()
}

func (s *GetUserResourceStatusResponse) SetHeaders(v map[string]*string) *GetUserResourceStatusResponse {
	s.Headers = v
	return s
}

func (s *GetUserResourceStatusResponse) SetBody(v *GetUserResourceStatusResponseBody) *GetUserResourceStatusResponse {
	s.Body = v
	return s
}

type GetUserSpecialVersionGpuResourceInfoRequest struct {
	// 付费类型
	PayType *int64 `json:"PayType,omitempty" xml:"PayType,omitempty"`
}

func (s GetUserSpecialVersionGpuResourceInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserSpecialVersionGpuResourceInfoRequest) GoString() string {
	return s.String()
}

func (s *GetUserSpecialVersionGpuResourceInfoRequest) SetPayType(v int64) *GetUserSpecialVersionGpuResourceInfoRequest {
	s.PayType = &v
	return s
}

type GetUserSpecialVersionGpuResourceInfoResponseBody struct {
	GpuAvailableQuota *int64 `json:"GpuAvailableQuota,omitempty" xml:"GpuAvailableQuota,omitempty"`
	GpuTotalQuota     *int64 `json:"GpuTotalQuota,omitempty" xml:"GpuTotalQuota,omitempty"`
	// Id of the request
	RequestId *string         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources []*ResourceInfo `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s GetUserSpecialVersionGpuResourceInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserSpecialVersionGpuResourceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserSpecialVersionGpuResourceInfoResponseBody) SetGpuAvailableQuota(v int64) *GetUserSpecialVersionGpuResourceInfoResponseBody {
	s.GpuAvailableQuota = &v
	return s
}

func (s *GetUserSpecialVersionGpuResourceInfoResponseBody) SetGpuTotalQuota(v int64) *GetUserSpecialVersionGpuResourceInfoResponseBody {
	s.GpuTotalQuota = &v
	return s
}

func (s *GetUserSpecialVersionGpuResourceInfoResponseBody) SetRequestId(v string) *GetUserSpecialVersionGpuResourceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserSpecialVersionGpuResourceInfoResponseBody) SetResources(v []*ResourceInfo) *GetUserSpecialVersionGpuResourceInfoResponseBody {
	s.Resources = v
	return s
}

type GetUserSpecialVersionGpuResourceInfoResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUserSpecialVersionGpuResourceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserSpecialVersionGpuResourceInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserSpecialVersionGpuResourceInfoResponse) GoString() string {
	return s.String()
}

func (s *GetUserSpecialVersionGpuResourceInfoResponse) SetHeaders(v map[string]*string) *GetUserSpecialVersionGpuResourceInfoResponse {
	s.Headers = v
	return s
}

func (s *GetUserSpecialVersionGpuResourceInfoResponse) SetBody(v *GetUserSpecialVersionGpuResourceInfoResponseBody) *GetUserSpecialVersionGpuResourceInfoResponse {
	s.Body = v
	return s
}

type ListConfigsResponseBody struct {
	Configs []*Config `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConfigsResponseBody) SetConfigs(v []*Config) *ListConfigsResponseBody {
	s.Configs = v
	return s
}

func (s *ListConfigsResponseBody) SetRequestId(v string) *ListConfigsResponseBody {
	s.RequestId = &v
	return s
}

type ListConfigsResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListConfigsResponse) SetHeaders(v map[string]*string) *ListConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListConfigsResponse) SetBody(v *ListConfigsResponseBody) *ListConfigsResponse {
	s.Body = v
	return s
}

type ListDatasetsRequest struct {
	// 工作空间Id
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListDatasetsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDatasetsRequest) GoString() string {
	return s.String()
}

func (s *ListDatasetsRequest) SetWorkspaceId(v string) *ListDatasetsRequest {
	s.WorkspaceId = &v
	return s
}

type ListDatasetsResponseBody struct {
	Datasets []*Dataset `json:"Datasets,omitempty" xml:"Datasets,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDatasetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDatasetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatasetsResponseBody) SetDatasets(v []*Dataset) *ListDatasetsResponseBody {
	s.Datasets = v
	return s
}

func (s *ListDatasetsResponseBody) SetRequestId(v string) *ListDatasetsResponseBody {
	s.RequestId = &v
	return s
}

type ListDatasetsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDatasetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDatasetsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDatasetsResponse) GoString() string {
	return s.String()
}

func (s *ListDatasetsResponse) SetHeaders(v map[string]*string) *ListDatasetsResponse {
	s.Headers = v
	return s
}

func (s *ListDatasetsResponse) SetBody(v *ListDatasetsResponseBody) *ListDatasetsResponse {
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

type ListInstanceTypesRequest struct {
	// AcceleratorType
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
}

func (s ListInstanceTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceTypesRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceTypesRequest) SetAcceleratorType(v string) *ListInstanceTypesRequest {
	s.AcceleratorType = &v
	return s
}

type ListInstanceTypesResponseBody struct {
	InstanceTypes []*InstanceType `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInstanceTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceTypesResponseBody) SetInstanceTypes(v []*InstanceType) *ListInstanceTypesResponseBody {
	s.InstanceTypes = v
	return s
}

func (s *ListInstanceTypesResponseBody) SetRequestId(v string) *ListInstanceTypesResponseBody {
	s.RequestId = &v
	return s
}

type ListInstanceTypesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInstanceTypesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInstanceTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceTypesResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceTypesResponse) SetHeaders(v map[string]*string) *ListInstanceTypesResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceTypesResponse) SetBody(v *ListInstanceTypesResponseBody) *ListInstanceTypesResponse {
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

type ListNamespacesResponseBody struct {
	// 命名空间列表
	Namespaces []*ImageNamespace `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListNamespacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNamespacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNamespacesResponseBody) SetNamespaces(v []*ImageNamespace) *ListNamespacesResponseBody {
	s.Namespaces = v
	return s
}

func (s *ListNamespacesResponseBody) SetRequestId(v string) *ListNamespacesResponseBody {
	s.RequestId = &v
	return s
}

type ListNamespacesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListNamespacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNamespacesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNamespacesResponse) GoString() string {
	return s.String()
}

func (s *ListNamespacesResponse) SetHeaders(v map[string]*string) *ListNamespacesResponse {
	s.Headers = v
	return s
}

func (s *ListNamespacesResponse) SetBody(v *ListNamespacesResponseBody) *ListNamespacesResponse {
	s.Body = v
	return s
}

type ListNasesResponseBody struct {
	// nas文件系统列表
	Nases []*Nas `json:"Nases,omitempty" xml:"Nases,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListNasesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNasesResponseBody) SetNases(v []*Nas) *ListNasesResponseBody {
	s.Nases = v
	return s
}

func (s *ListNasesResponseBody) SetRequestId(v string) *ListNasesResponseBody {
	s.RequestId = &v
	return s
}

type ListNasesResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListNasesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNasesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNasesResponse) GoString() string {
	return s.String()
}

func (s *ListNasesResponse) SetHeaders(v map[string]*string) *ListNasesResponse {
	s.Headers = v
	return s
}

func (s *ListNasesResponse) SetBody(v *ListNasesResponseBody) *ListNasesResponse {
	s.Body = v
	return s
}

type ListNetworkSecurityGroupsResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// security groups
	SecurityGroups []*SecurityGroup `json:"SecurityGroups,omitempty" xml:"SecurityGroups,omitempty" type:"Repeated"`
}

func (s ListNetworkSecurityGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNetworkSecurityGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNetworkSecurityGroupsResponseBody) SetRequestId(v string) *ListNetworkSecurityGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNetworkSecurityGroupsResponseBody) SetSecurityGroups(v []*SecurityGroup) *ListNetworkSecurityGroupsResponseBody {
	s.SecurityGroups = v
	return s
}

type ListNetworkSecurityGroupsResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListNetworkSecurityGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNetworkSecurityGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNetworkSecurityGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListNetworkSecurityGroupsResponse) SetHeaders(v map[string]*string) *ListNetworkSecurityGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListNetworkSecurityGroupsResponse) SetBody(v *ListNetworkSecurityGroupsResponseBody) *ListNetworkSecurityGroupsResponse {
	s.Body = v
	return s
}

type ListNetworkVSwitchesResponseBody struct {
	// Id of the request
	RequestId *string    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VSwitches []*VSwitch `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
}

func (s ListNetworkVSwitchesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNetworkVSwitchesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNetworkVSwitchesResponseBody) SetRequestId(v string) *ListNetworkVSwitchesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNetworkVSwitchesResponseBody) SetVSwitches(v []*VSwitch) *ListNetworkVSwitchesResponseBody {
	s.VSwitches = v
	return s
}

type ListNetworkVSwitchesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListNetworkVSwitchesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNetworkVSwitchesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNetworkVSwitchesResponse) GoString() string {
	return s.String()
}

func (s *ListNetworkVSwitchesResponse) SetHeaders(v map[string]*string) *ListNetworkVSwitchesResponse {
	s.Headers = v
	return s
}

func (s *ListNetworkVSwitchesResponse) SetBody(v *ListNetworkVSwitchesResponseBody) *ListNetworkVSwitchesResponse {
	s.Body = v
	return s
}

type ListNetworkVpcsResponseBody struct {
	// RequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// vpc列表
	Vpcs []*Vpc `json:"Vpcs,omitempty" xml:"Vpcs,omitempty" type:"Repeated"`
}

func (s ListNetworkVpcsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNetworkVpcsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNetworkVpcsResponseBody) SetRequestId(v string) *ListNetworkVpcsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNetworkVpcsResponseBody) SetVpcs(v []*Vpc) *ListNetworkVpcsResponseBody {
	s.Vpcs = v
	return s
}

type ListNetworkVpcsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListNetworkVpcsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNetworkVpcsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNetworkVpcsResponse) GoString() string {
	return s.String()
}

func (s *ListNetworkVpcsResponse) SetHeaders(v map[string]*string) *ListNetworkVpcsResponse {
	s.Headers = v
	return s
}

func (s *ListNetworkVpcsResponse) SetBody(v *ListNetworkVpcsResponseBody) *ListNetworkVpcsResponse {
	s.Body = v
	return s
}

type ListRegionsResponseBody struct {
	Regions []*Region `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBody) SetRegions(v []*Region) *ListRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *ListRegionsResponseBody) SetRequestId(v string) *ListRegionsResponseBody {
	s.RequestId = &v
	return s
}

type ListRegionsResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListRegionsResponse) SetHeaders(v map[string]*string) *ListRegionsResponse {
	s.Headers = v
	return s
}

func (s *ListRegionsResponse) SetBody(v *ListRegionsResponseBody) *ListRegionsResponse {
	s.Body = v
	return s
}

type ListRepositoriesResponseBody struct {
	Repositories []*ImageRepository `json:"Repositories,omitempty" xml:"Repositories,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRepositoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepositoriesResponseBody) SetRepositories(v []*ImageRepository) *ListRepositoriesResponseBody {
	s.Repositories = v
	return s
}

func (s *ListRepositoriesResponseBody) SetRequestId(v string) *ListRepositoriesResponseBody {
	s.RequestId = &v
	return s
}

type ListRepositoriesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRepositoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRepositoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoriesResponse) GoString() string {
	return s.String()
}

func (s *ListRepositoriesResponse) SetHeaders(v map[string]*string) *ListRepositoriesResponse {
	s.Headers = v
	return s
}

func (s *ListRepositoriesResponse) SetBody(v *ListRepositoriesResponseBody) *ListRepositoriesResponse {
	s.Body = v
	return s
}

type ListUserClustersResponseBody struct {
	Clusters []*Cluster `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListUserClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserClustersResponseBody) SetClusters(v []*Cluster) *ListUserClustersResponseBody {
	s.Clusters = v
	return s
}

func (s *ListUserClustersResponseBody) SetRequestId(v string) *ListUserClustersResponseBody {
	s.RequestId = &v
	return s
}

type ListUserClustersResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUserClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUserClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserClustersResponse) GoString() string {
	return s.String()
}

func (s *ListUserClustersResponse) SetHeaders(v map[string]*string) *ListUserClustersResponse {
	s.Headers = v
	return s
}

func (s *ListUserClustersResponse) SetBody(v *ListUserClustersResponseBody) *ListUserClustersResponse {
	s.Body = v
	return s
}

type ListUserWorkNodesRequest struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s ListUserWorkNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserWorkNodesRequest) GoString() string {
	return s.String()
}

func (s *ListUserWorkNodesRequest) SetClusterId(v string) *ListUserWorkNodesRequest {
	s.ClusterId = &v
	return s
}

type ListUserWorkNodesResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListUserWorkNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserWorkNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserWorkNodesResponseBody) SetRequestId(v string) *ListUserWorkNodesResponseBody {
	s.RequestId = &v
	return s
}

type ListUserWorkNodesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUserWorkNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUserWorkNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserWorkNodesResponse) GoString() string {
	return s.String()
}

func (s *ListUserWorkNodesResponse) SetHeaders(v map[string]*string) *ListUserWorkNodesResponse {
	s.Headers = v
	return s
}

func (s *ListUserWorkNodesResponse) SetBody(v *ListUserWorkNodesResponseBody) *ListUserWorkNodesResponse {
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

func (client *Client) CreateImage(request *CreateImageRequest) (_result *CreateImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateImageResponse{}
	_body, _err := client.CreateImageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateImageWithOptions(request *CreateImageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Repository)) {
		body["Repository"] = request.Repository
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateImageResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateImage"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v1/images"), tea.String("json"), req, runtime)
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
	_result = &CreateInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateInstance"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v1/instances/"), tea.String("json"), req, runtime)
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
	_result = &CreateInstanceShutdownTimerResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateInstanceShutdownTimer"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v1/instances/"+tea.StringValue(InstanceId)+"/shutdownTimer"), tea.String("json"), req, runtime)
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
	_result = &CreateInstanceSnapshotResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateInstanceSnapshot"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v1/instances/"+tea.StringValue(InstanceId)+"/snapshots"), tea.String("json"), req, runtime)
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
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteInstance"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/api/v1/instances/"+tea.StringValue(InstanceId)), tea.String("json"), req, runtime)
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
	_result = &DeleteInstanceShutdownTimerResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteInstanceShutdownTimer"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/api/v1/instances/"+tea.StringValue(InstanceId)+"/shutdownTimer"), tea.String("json"), req, runtime)
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
	_result = &DeleteInstanceSnapshotResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteInstanceSnapshot"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/api/v1/instances/"+tea.StringValue(InstanceId)+"/snapshots/"+tea.StringValue(InstanceSnapshotId)), tea.String("json"), req, runtime)
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
	_result = &GetDashboardStatisticsResponse{}
	_body, _err := client.DoROARequest(tea.String("GetDashboardStatistics"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/statistics/dashboard"), tea.String("json"), req, runtime)
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
	_result = &GetInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("GetInstance"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/instances/"+tea.StringValue(InstanceId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstanceImage(InstanceId *string) (_result *GetInstanceImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceImageResponse{}
	_body, _err := client.GetInstanceImageWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstanceImageWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetInstanceImageResponse, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetInstanceImageResponse{}
	_body, _err := client.DoROARequest(tea.String("GetInstanceImage"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/instances/"+tea.StringValue(InstanceId)+"/image"), tea.String("json"), req, runtime)
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
	_result = &GetInstanceShutdownTimerResponse{}
	_body, _err := client.DoROARequest(tea.String("GetInstanceShutdownTimer"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/instances/"+tea.StringValue(InstanceId)+"/shutdownTimer"), tea.String("json"), req, runtime)
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
	_result = &GetInstanceSnapshotResponse{}
	_body, _err := client.DoROARequest(tea.String("GetInstanceSnapshot"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/instances/"+tea.StringValue(InstanceId)+"/snapshots/"+tea.StringValue(InstanceSnapshotId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstanceType(InstanceTypeId *string) (_result *GetInstanceTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceTypeResponse{}
	_body, _err := client.GetInstanceTypeWithOptions(InstanceTypeId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstanceTypeWithOptions(InstanceTypeId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetInstanceTypeResponse, _err error) {
	InstanceTypeId = openapiutil.GetEncodeParam(InstanceTypeId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetInstanceTypeResponse{}
	_body, _err := client.DoROARequest(tea.String("GetInstanceType"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/instanceTypes/"+tea.StringValue(InstanceTypeId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstanceUrl(InstanceId *string) (_result *GetInstanceUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceUrlResponse{}
	_body, _err := client.GetInstanceUrlWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstanceUrlWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetInstanceUrlResponse, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetInstanceUrlResponse{}
	_body, _err := client.DoROARequest(tea.String("GetInstanceUrl"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/instances/"+tea.StringValue(InstanceId)+"/url"), tea.String("json"), req, runtime)
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

func (client *Client) GetUserConfig() (_result *GetUserConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUserConfigResponse{}
	_body, _err := client.GetUserConfigWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserConfigWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetUserConfigResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetUserConfigResponse{}
	_body, _err := client.DoROARequest(tea.String("GetUserConfig"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/users/config"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserResourceAuthorizationStatus() (_result *GetUserResourceAuthorizationStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUserResourceAuthorizationStatusResponse{}
	_body, _err := client.GetUserResourceAuthorizationStatusWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserResourceAuthorizationStatusWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetUserResourceAuthorizationStatusResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetUserResourceAuthorizationStatusResponse{}
	_body, _err := client.DoROARequest(tea.String("GetUserResourceAuthorizationStatus"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/users/resourceAuthorizationStatus"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserResourceStatus() (_result *GetUserResourceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUserResourceStatusResponse{}
	_body, _err := client.GetUserResourceStatusWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserResourceStatusWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetUserResourceStatusResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetUserResourceStatusResponse{}
	_body, _err := client.DoROARequest(tea.String("GetUserResourceStatus"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/users/resourceStatus"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserSpecialVersionGpuResourceInfo(request *GetUserSpecialVersionGpuResourceInfoRequest) (_result *GetUserSpecialVersionGpuResourceInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUserSpecialVersionGpuResourceInfoResponse{}
	_body, _err := client.GetUserSpecialVersionGpuResourceInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserSpecialVersionGpuResourceInfoWithOptions(request *GetUserSpecialVersionGpuResourceInfoRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetUserSpecialVersionGpuResourceInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		query["PayType"] = request.PayType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetUserSpecialVersionGpuResourceInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetUserSpecialVersionGpuResourceInfo"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/users/specialVersionGpuResourceInfo"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListConfigs() (_result *ListConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListConfigsResponse{}
	_body, _err := client.ListConfigsWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListConfigsWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListConfigsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ListConfigsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListConfigs"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/configs"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDatasets(request *ListDatasetsRequest) (_result *ListDatasetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDatasetsResponse{}
	_body, _err := client.ListDatasetsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDatasetsWithOptions(request *ListDatasetsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDatasetsResponse, _err error) {
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
	_result = &ListDatasetsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListDatasets"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/datasets"), tea.String("json"), req, runtime)
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
	_result = &ListImagesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListImages"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/images"), tea.String("json"), req, runtime)
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
	_result = &ListInstanceSnapshotsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListInstanceSnapshots"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/instances/"+tea.StringValue(InstanceId)+"/snapshots"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstanceTypes(request *ListInstanceTypesRequest) (_result *ListInstanceTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstanceTypesResponse{}
	_body, _err := client.ListInstanceTypesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstanceTypesWithOptions(request *ListInstanceTypesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstanceTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorType)) {
		query["AcceleratorType"] = request.AcceleratorType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListInstanceTypesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListInstanceTypes"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/instanceTypes"), tea.String("json"), req, runtime)
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
	_result = &ListInstancesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListInstances"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/instances"), tea.String("json"), req, runtime)
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
	_result = &ListInstancesStatusResponse{}
	_body, _err := client.DoROARequest(tea.String("ListInstancesStatus"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/statuses/instances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNamespaces() (_result *ListNamespacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListNamespacesResponse{}
	_body, _err := client.ListNamespacesWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNamespacesWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListNamespacesResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ListNamespacesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListNamespaces"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/imageRegistry/namespaces"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNases() (_result *ListNasesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListNasesResponse{}
	_body, _err := client.ListNasesWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNasesWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListNasesResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ListNasesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListNases"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/nases"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNetworkSecurityGroups(VpcId *string) (_result *ListNetworkSecurityGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListNetworkSecurityGroupsResponse{}
	_body, _err := client.ListNetworkSecurityGroupsWithOptions(VpcId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNetworkSecurityGroupsWithOptions(VpcId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListNetworkSecurityGroupsResponse, _err error) {
	VpcId = openapiutil.GetEncodeParam(VpcId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ListNetworkSecurityGroupsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListNetworkSecurityGroups"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/network/vpcs/"+tea.StringValue(VpcId)+"/securityGroups"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNetworkVSwitches(VpcId *string) (_result *ListNetworkVSwitchesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListNetworkVSwitchesResponse{}
	_body, _err := client.ListNetworkVSwitchesWithOptions(VpcId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNetworkVSwitchesWithOptions(VpcId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListNetworkVSwitchesResponse, _err error) {
	VpcId = openapiutil.GetEncodeParam(VpcId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ListNetworkVSwitchesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListNetworkVSwitches"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/network/vpcs/"+tea.StringValue(VpcId)+"/vSwitches"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNetworkVpcs() (_result *ListNetworkVpcsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListNetworkVpcsResponse{}
	_body, _err := client.ListNetworkVpcsWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNetworkVpcsWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListNetworkVpcsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ListNetworkVpcsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListNetworkVpcs"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/network/vpcs"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRegions() (_result *ListRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRegionsResponse{}
	_body, _err := client.ListRegionsWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRegionsWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ListRegionsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListRegions"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/configs/regions"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRepositories(Namespace *string) (_result *ListRepositoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRepositoriesResponse{}
	_body, _err := client.ListRepositoriesWithOptions(Namespace, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRepositoriesWithOptions(Namespace *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRepositoriesResponse, _err error) {
	Namespace = openapiutil.GetEncodeParam(Namespace)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ListRepositoriesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListRepositories"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/imageRegistry/namespaces/"+tea.StringValue(Namespace)+"/repositories"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserClusters() (_result *ListUserClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListUserClustersResponse{}
	_body, _err := client.ListUserClustersWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUserClustersWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListUserClustersResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ListUserClustersResponse{}
	_body, _err := client.DoROARequest(tea.String("ListUserClusters"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/users/clusters"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserWorkNodes(request *ListUserWorkNodesRequest) (_result *ListUserWorkNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListUserWorkNodesResponse{}
	_body, _err := client.ListUserWorkNodesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUserWorkNodesWithOptions(request *ListUserWorkNodesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListUserWorkNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListUserWorkNodesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListUserWorkNodes"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/users/workerNodes"), tea.String("json"), req, runtime)
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
	_result = &StartInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("StartInstance"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/api/v1/instances/"+tea.StringValue(InstanceId)+"/start"), tea.String("json"), req, runtime)
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
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
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
	_result = &UpdateInstanceSnapshotResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateInstanceSnapshot"), tea.String("2021-02-26"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/api/v1/instances/"+tea.StringValue(InstanceId)+"/snapshots/"+tea.StringValue(InstanceSnapshotId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
