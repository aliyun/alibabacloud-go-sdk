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

type AccelerationInfo struct {
	// 镜像加速状态，取值 Preparing 或 Ready
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s AccelerationInfo) String() string {
	return tea.Prettify(s)
}

func (s AccelerationInfo) GoString() string {
	return s.String()
}

func (s *AccelerationInfo) SetStatus(v string) *AccelerationInfo {
	s.Status = &v
	return s
}

type CertConfig struct {
	// 证书名称
	CertName *string `json:"certName,omitempty" xml:"certName,omitempty"`
	// 证书，如果是证书链则依次填写多个证书
	Certificate *string `json:"certificate,omitempty" xml:"certificate,omitempty"`
	// 私钥
	PrivateKey *string `json:"privateKey,omitempty" xml:"privateKey,omitempty"`
}

func (s CertConfig) String() string {
	return tea.Prettify(s)
}

func (s CertConfig) GoString() string {
	return s.String()
}

func (s *CertConfig) SetCertName(v string) *CertConfig {
	s.CertName = &v
	return s
}

func (s *CertConfig) SetCertificate(v string) *CertConfig {
	s.Certificate = &v
	return s
}

func (s *CertConfig) SetPrivateKey(v string) *CertConfig {
	s.PrivateKey = &v
	return s
}

type Code struct {
	// codeCheckSum
	CodeCheckSum *string `json:"codeCheckSum,omitempty" xml:"codeCheckSum,omitempty"`
	// error
	Err *string `json:"err,omitempty" xml:"err,omitempty"`
	// 函数代码包的OSS bucket name
	OssBucketName *string `json:"ossBucketName,omitempty" xml:"ossBucketName,omitempty"`
	// 函数代码包的OSS对象名
	OssObjectName *string `json:"ossObjectName,omitempty" xml:"ossObjectName,omitempty"`
	// 直接在request body中上传code zip包的base64编码
	ZipFile *string `json:"zipFile,omitempty" xml:"zipFile,omitempty"`
}

func (s Code) String() string {
	return tea.Prettify(s)
}

func (s Code) GoString() string {
	return s.String()
}

func (s *Code) SetCodeCheckSum(v string) *Code {
	s.CodeCheckSum = &v
	return s
}

func (s *Code) SetErr(v string) *Code {
	s.Err = &v
	return s
}

func (s *Code) SetOssBucketName(v string) *Code {
	s.OssBucketName = &v
	return s
}

func (s *Code) SetOssObjectName(v string) *Code {
	s.OssObjectName = &v
	return s
}

func (s *Code) SetZipFile(v string) *Code {
	s.ZipFile = &v
	return s
}

type CustomContainerConfig struct {
	// 镜像加速类型，取值Default为开启加速，None为关闭加速，默认关闭
	AccelerationType *string `json:"accelerationType,omitempty" xml:"accelerationType,omitempty"`
	// 容器启动参数
	Args *string `json:"args,omitempty" xml:"args,omitempty"`
	// 容器启动命令，等同于 Docker ENTRYPOINT
	Command *string `json:"command,omitempty" xml:"command,omitempty"`
	// 容器镜像地址，实例值： registry-vpc.cn-hangzhou.aliyuncs.com/fc-demo/helloworld:v1beta1
	Image *string `json:"image,omitempty" xml:"image,omitempty"`
}

func (s CustomContainerConfig) String() string {
	return tea.Prettify(s)
}

func (s CustomContainerConfig) GoString() string {
	return s.String()
}

func (s *CustomContainerConfig) SetAccelerationType(v string) *CustomContainerConfig {
	s.AccelerationType = &v
	return s
}

func (s *CustomContainerConfig) SetArgs(v string) *CustomContainerConfig {
	s.Args = &v
	return s
}

func (s *CustomContainerConfig) SetCommand(v string) *CustomContainerConfig {
	s.Command = &v
	return s
}

func (s *CustomContainerConfig) SetImage(v string) *CustomContainerConfig {
	s.Image = &v
	return s
}

type CustomContainerConfigInfo struct {
	AccelerationInfo *AccelerationInfo `json:"accelerationInfo,omitempty" xml:"accelerationInfo,omitempty"`
	// 镜像加速类型，取值Default为开启加速，None为关闭加速，默认关闭
	AccelerationType *string `json:"accelerationType,omitempty" xml:"accelerationType,omitempty"`
	// 容器启动参数
	Args *string `json:"args,omitempty" xml:"args,omitempty"`
	// 容器启动命令，等同于 Docker ENTRYPOINT
	Command *string `json:"command,omitempty" xml:"command,omitempty"`
	// 容器镜像地址，实例值： registry-vpc.cn-hangzhou.aliyuncs.com/fc-demo/helloworld:v1beta1
	Image *string `json:"image,omitempty" xml:"image,omitempty"`
}

func (s CustomContainerConfigInfo) String() string {
	return tea.Prettify(s)
}

func (s CustomContainerConfigInfo) GoString() string {
	return s.String()
}

func (s *CustomContainerConfigInfo) SetAccelerationInfo(v *AccelerationInfo) *CustomContainerConfigInfo {
	s.AccelerationInfo = v
	return s
}

func (s *CustomContainerConfigInfo) SetAccelerationType(v string) *CustomContainerConfigInfo {
	s.AccelerationType = &v
	return s
}

func (s *CustomContainerConfigInfo) SetArgs(v string) *CustomContainerConfigInfo {
	s.Args = &v
	return s
}

func (s *CustomContainerConfigInfo) SetCommand(v string) *CustomContainerConfigInfo {
	s.Command = &v
	return s
}

func (s *CustomContainerConfigInfo) SetImage(v string) *CustomContainerConfigInfo {
	s.Image = &v
	return s
}

type CustomDNS struct {
	// DNS resolver 配置参数列表
	DnsOptions []*DNSOption `json:"dnsOptions,omitempty" xml:"dnsOptions,omitempty" type:"Repeated"`
	// DNS 服务器的 IP 地址列表
	NameServers []*string `json:"nameServers,omitempty" xml:"nameServers,omitempty" type:"Repeated"`
	// DNS 搜索域的列表
	Searches []*string `json:"searches,omitempty" xml:"searches,omitempty" type:"Repeated"`
}

func (s CustomDNS) String() string {
	return tea.Prettify(s)
}

func (s CustomDNS) GoString() string {
	return s.String()
}

func (s *CustomDNS) SetDnsOptions(v []*DNSOption) *CustomDNS {
	s.DnsOptions = v
	return s
}

func (s *CustomDNS) SetNameServers(v []*string) *CustomDNS {
	s.NameServers = v
	return s
}

func (s *CustomDNS) SetSearches(v []*string) *CustomDNS {
	s.Searches = v
	return s
}

type CustomRuntimeConfig struct {
	// 启动入口命令接收的参数
	Args []*string `json:"args,omitempty" xml:"args,omitempty" type:"Repeated"`
	// 启动入口命令
	Command []*string `json:"command,omitempty" xml:"command,omitempty" type:"Repeated"`
}

func (s CustomRuntimeConfig) String() string {
	return tea.Prettify(s)
}

func (s CustomRuntimeConfig) GoString() string {
	return s.String()
}

func (s *CustomRuntimeConfig) SetArgs(v []*string) *CustomRuntimeConfig {
	s.Args = v
	return s
}

func (s *CustomRuntimeConfig) SetCommand(v []*string) *CustomRuntimeConfig {
	s.Command = v
	return s
}

type DNSOption struct {
	// DNS option 名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// DNS option 值
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DNSOption) String() string {
	return tea.Prettify(s)
}

func (s DNSOption) GoString() string {
	return s.String()
}

func (s *DNSOption) SetName(v string) *DNSOption {
	s.Name = &v
	return s
}

func (s *DNSOption) SetValue(v string) *DNSOption {
	s.Value = &v
	return s
}

type Destination struct {
	// destination
	Destination *string `json:"destination,omitempty" xml:"destination,omitempty"`
}

func (s Destination) String() string {
	return tea.Prettify(s)
}

func (s Destination) GoString() string {
	return s.String()
}

func (s *Destination) SetDestination(v string) *Destination {
	s.Destination = &v
	return s
}

type DestinationConfig struct {
	OnFailure *Destination `json:"onFailure,omitempty" xml:"onFailure,omitempty"`
	OnSuccess *Destination `json:"onSuccess,omitempty" xml:"onSuccess,omitempty"`
}

func (s DestinationConfig) String() string {
	return tea.Prettify(s)
}

func (s DestinationConfig) GoString() string {
	return s.String()
}

func (s *DestinationConfig) SetOnFailure(v *Destination) *DestinationConfig {
	s.OnFailure = v
	return s
}

func (s *DestinationConfig) SetOnSuccess(v *Destination) *DestinationConfig {
	s.OnSuccess = v
	return s
}

type InstanceLifecycleConfig struct {
	PreFreeze *LifecycleHook `json:"preFreeze,omitempty" xml:"preFreeze,omitempty"`
	PreStop   *LifecycleHook `json:"preStop,omitempty" xml:"preStop,omitempty"`
}

func (s InstanceLifecycleConfig) String() string {
	return tea.Prettify(s)
}

func (s InstanceLifecycleConfig) GoString() string {
	return s.String()
}

func (s *InstanceLifecycleConfig) SetPreFreeze(v *LifecycleHook) *InstanceLifecycleConfig {
	s.PreFreeze = v
	return s
}

func (s *InstanceLifecycleConfig) SetPreStop(v *LifecycleHook) *InstanceLifecycleConfig {
	s.PreStop = v
	return s
}

type JaegerConfig struct {
	// endpoint
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
}

func (s JaegerConfig) String() string {
	return tea.Prettify(s)
}

func (s JaegerConfig) GoString() string {
	return s.String()
}

func (s *JaegerConfig) SetEndpoint(v string) *JaegerConfig {
	s.Endpoint = &v
	return s
}

type Layer struct {
	// 层访问类型
	Acl *int32 `json:"acl,omitempty" xml:"acl,omitempty"`
	// arn
	Arn *string `json:"arn,omitempty" xml:"arn,omitempty"`
	// 层代码
	Code *LayerCode `json:"code,omitempty" xml:"code,omitempty"`
	// 层Checksum
	CodeChecksum *string `json:"codeChecksum,omitempty" xml:"codeChecksum,omitempty"`
	// 层代码大小
	CodeSize *int64 `json:"codeSize,omitempty" xml:"codeSize,omitempty"`
	// compatibleRuntime
	CompatibleRuntime []*string `json:"compatibleRuntime,omitempty" xml:"compatibleRuntime,omitempty" type:"Repeated"`
	// 层创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 层描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 层名称
	LayerName *string `json:"layerName,omitempty" xml:"layerName,omitempty"`
	// 层版本
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s Layer) String() string {
	return tea.Prettify(s)
}

func (s Layer) GoString() string {
	return s.String()
}

func (s *Layer) SetAcl(v int32) *Layer {
	s.Acl = &v
	return s
}

func (s *Layer) SetArn(v string) *Layer {
	s.Arn = &v
	return s
}

func (s *Layer) SetCode(v *LayerCode) *Layer {
	s.Code = v
	return s
}

func (s *Layer) SetCodeChecksum(v string) *Layer {
	s.CodeChecksum = &v
	return s
}

func (s *Layer) SetCodeSize(v int64) *Layer {
	s.CodeSize = &v
	return s
}

func (s *Layer) SetCompatibleRuntime(v []*string) *Layer {
	s.CompatibleRuntime = v
	return s
}

func (s *Layer) SetCreateTime(v string) *Layer {
	s.CreateTime = &v
	return s
}

func (s *Layer) SetDescription(v string) *Layer {
	s.Description = &v
	return s
}

func (s *Layer) SetLayerName(v string) *Layer {
	s.LayerName = &v
	return s
}

func (s *Layer) SetVersion(v int32) *Layer {
	s.Version = &v
	return s
}

type LayerCode struct {
	// 层代码位置
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// 层代码类型
	RepositoryType *string `json:"repositoryType,omitempty" xml:"repositoryType,omitempty"`
}

func (s LayerCode) String() string {
	return tea.Prettify(s)
}

func (s LayerCode) GoString() string {
	return s.String()
}

func (s *LayerCode) SetLocation(v string) *LayerCode {
	s.Location = &v
	return s
}

func (s *LayerCode) SetRepositoryType(v string) *LayerCode {
	s.RepositoryType = &v
	return s
}

type LifecycleHook struct {
	// handler name
	Handler *string `json:"handler,omitempty" xml:"handler,omitempty"`
	// timeout in second
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
}

func (s LifecycleHook) String() string {
	return tea.Prettify(s)
}

func (s LifecycleHook) GoString() string {
	return s.String()
}

func (s *LifecycleHook) SetHandler(v string) *LifecycleHook {
	s.Handler = &v
	return s
}

func (s *LifecycleHook) SetTimeout(v int32) *LifecycleHook {
	s.Timeout = &v
	return s
}

type LogConfig struct {
	// 开启实例级别指标
	EnableInstanceMetrics *bool `json:"enableInstanceMetrics,omitempty" xml:"enableInstanceMetrics,omitempty"`
	// 开启请求级别指标
	EnableRequestMetrics *bool `json:"enableRequestMetrics,omitempty" xml:"enableRequestMetrics,omitempty"`
	// 日志切分规则
	LogBeginRule *string `json:"logBeginRule,omitempty" xml:"logBeginRule,omitempty"`
	// 日志库
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// 日志项目
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
}

func (s LogConfig) String() string {
	return tea.Prettify(s)
}

func (s LogConfig) GoString() string {
	return s.String()
}

func (s *LogConfig) SetEnableInstanceMetrics(v bool) *LogConfig {
	s.EnableInstanceMetrics = &v
	return s
}

func (s *LogConfig) SetEnableRequestMetrics(v bool) *LogConfig {
	s.EnableRequestMetrics = &v
	return s
}

func (s *LogConfig) SetLogBeginRule(v string) *LogConfig {
	s.LogBeginRule = &v
	return s
}

func (s *LogConfig) SetLogstore(v string) *LogConfig {
	s.Logstore = &v
	return s
}

func (s *LogConfig) SetProject(v string) *LogConfig {
	s.Project = &v
	return s
}

type NASConfig struct {
	// groupID
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// 挂载点
	MountPoints []*NASConfigMountPoints `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
	// userID
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s NASConfig) String() string {
	return tea.Prettify(s)
}

func (s NASConfig) GoString() string {
	return s.String()
}

func (s *NASConfig) SetGroupId(v string) *NASConfig {
	s.GroupId = &v
	return s
}

func (s *NASConfig) SetMountPoints(v []*NASConfigMountPoints) *NASConfig {
	s.MountPoints = v
	return s
}

func (s *NASConfig) SetUserId(v string) *NASConfig {
	s.UserId = &v
	return s
}

type NASConfigMountPoints struct {
	// 本地挂载目录
	MountDir *string `json:"mountDir,omitempty" xml:"mountDir,omitempty"`
	// NAS服务器地址
	ServerAddr *string `json:"serverAddr,omitempty" xml:"serverAddr,omitempty"`
}

func (s NASConfigMountPoints) String() string {
	return tea.Prettify(s)
}

func (s NASConfigMountPoints) GoString() string {
	return s.String()
}

func (s *NASConfigMountPoints) SetMountDir(v string) *NASConfigMountPoints {
	s.MountDir = &v
	return s
}

func (s *NASConfigMountPoints) SetServerAddr(v string) *NASConfigMountPoints {
	s.ServerAddr = &v
	return s
}

type OnDemandConfig struct {
	// todo
	MaximumInstanceCount *int64 `json:"maximumInstanceCount,omitempty" xml:"maximumInstanceCount,omitempty"`
	// 函数详情
	Resource *string `json:"resource,omitempty" xml:"resource,omitempty"`
}

func (s OnDemandConfig) String() string {
	return tea.Prettify(s)
}

func (s OnDemandConfig) GoString() string {
	return s.String()
}

func (s *OnDemandConfig) SetMaximumInstanceCount(v int64) *OnDemandConfig {
	s.MaximumInstanceCount = &v
	return s
}

func (s *OnDemandConfig) SetResource(v string) *OnDemandConfig {
	s.Resource = &v
	return s
}

type OpenReservedCapacity struct {
	// createdTime
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// cu
	Cu *int64 `json:"cu,omitempty" xml:"cu,omitempty"`
	// deadline
	Deadline *string `json:"deadline,omitempty" xml:"deadline,omitempty"`
	// instanceId
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// isRefunded
	IsRefunded *string `json:"isRefunded,omitempty" xml:"isRefunded,omitempty"`
	// lastModifiedTime
	LastModifiedTime *string `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
}

func (s OpenReservedCapacity) String() string {
	return tea.Prettify(s)
}

func (s OpenReservedCapacity) GoString() string {
	return s.String()
}

func (s *OpenReservedCapacity) SetCreatedTime(v string) *OpenReservedCapacity {
	s.CreatedTime = &v
	return s
}

func (s *OpenReservedCapacity) SetCu(v int64) *OpenReservedCapacity {
	s.Cu = &v
	return s
}

func (s *OpenReservedCapacity) SetDeadline(v string) *OpenReservedCapacity {
	s.Deadline = &v
	return s
}

func (s *OpenReservedCapacity) SetInstanceId(v string) *OpenReservedCapacity {
	s.InstanceId = &v
	return s
}

func (s *OpenReservedCapacity) SetIsRefunded(v string) *OpenReservedCapacity {
	s.IsRefunded = &v
	return s
}

func (s *OpenReservedCapacity) SetLastModifiedTime(v string) *OpenReservedCapacity {
	s.LastModifiedTime = &v
	return s
}

type OutputCodeLocation struct {
	// location
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// repositoryType
	RepositoryType *string `json:"repositoryType,omitempty" xml:"repositoryType,omitempty"`
}

func (s OutputCodeLocation) String() string {
	return tea.Prettify(s)
}

func (s OutputCodeLocation) GoString() string {
	return s.String()
}

func (s *OutputCodeLocation) SetLocation(v string) *OutputCodeLocation {
	s.Location = &v
	return s
}

func (s *OutputCodeLocation) SetRepositoryType(v string) *OutputCodeLocation {
	s.RepositoryType = &v
	return s
}

type PathConfig struct {
	// 该路径/方法对应的函数名称
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	// 请求方法，不填表示当前路径的所有方法匹配同一函数
	Methods []*string `json:"methods,omitempty" xml:"methods,omitempty" type:"Repeated"`
	// 请求路径
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// 该路径/方法对应服务的版本/别名
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	// 该路径/方法对应的服务名称
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
}

func (s PathConfig) String() string {
	return tea.Prettify(s)
}

func (s PathConfig) GoString() string {
	return s.String()
}

func (s *PathConfig) SetFunctionName(v string) *PathConfig {
	s.FunctionName = &v
	return s
}

func (s *PathConfig) SetMethods(v []*string) *PathConfig {
	s.Methods = v
	return s
}

func (s *PathConfig) SetPath(v string) *PathConfig {
	s.Path = &v
	return s
}

func (s *PathConfig) SetQualifier(v string) *PathConfig {
	s.Qualifier = &v
	return s
}

func (s *PathConfig) SetServiceName(v string) *PathConfig {
	s.ServiceName = &v
	return s
}

type Resource struct {
	// resourceArn
	ResourceArn *string `json:"resourceArn,omitempty" xml:"resourceArn,omitempty"`
	// tags
	Tags map[string]*string `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s Resource) String() string {
	return tea.Prettify(s)
}

func (s Resource) GoString() string {
	return s.String()
}

func (s *Resource) SetResourceArn(v string) *Resource {
	s.ResourceArn = &v
	return s
}

func (s *Resource) SetTags(v map[string]*string) *Resource {
	s.Tags = v
	return s
}

type RouteConfig struct {
	// routes
	Routes []*PathConfig `json:"routes,omitempty" xml:"routes,omitempty" type:"Repeated"`
}

func (s RouteConfig) String() string {
	return tea.Prettify(s)
}

func (s RouteConfig) GoString() string {
	return s.String()
}

func (s *RouteConfig) SetRoutes(v []*PathConfig) *RouteConfig {
	s.Routes = v
	return s
}

type ScheduledActions struct {
	// endTime
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// scheduleExpression
	ScheduleExpression *string `json:"scheduleExpression,omitempty" xml:"scheduleExpression,omitempty"`
	// startTime
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// target
	Target *int64 `json:"target,omitempty" xml:"target,omitempty"`
}

func (s ScheduledActions) String() string {
	return tea.Prettify(s)
}

func (s ScheduledActions) GoString() string {
	return s.String()
}

func (s *ScheduledActions) SetEndTime(v string) *ScheduledActions {
	s.EndTime = &v
	return s
}

func (s *ScheduledActions) SetName(v string) *ScheduledActions {
	s.Name = &v
	return s
}

func (s *ScheduledActions) SetScheduleExpression(v string) *ScheduledActions {
	s.ScheduleExpression = &v
	return s
}

func (s *ScheduledActions) SetStartTime(v string) *ScheduledActions {
	s.StartTime = &v
	return s
}

func (s *ScheduledActions) SetTarget(v int64) *ScheduledActions {
	s.Target = &v
	return s
}

type StatefulAsyncInvocation struct {
	// alreadyRetriedTimes
	AlreadyRetriedTimes *int64 `json:"alreadyRetriedTimes,omitempty" xml:"alreadyRetriedTimes,omitempty"`
	// destinationStatus
	DestinationStatus *string `json:"destinationStatus,omitempty" xml:"destinationStatus,omitempty"`
	// endTime
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// functionName
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	// invocationErrorMessage
	InvocationErrorMessage *string `json:"invocationErrorMessage,omitempty" xml:"invocationErrorMessage,omitempty"`
	// invocationId
	InvocationId *string `json:"invocationId,omitempty" xml:"invocationId,omitempty"`
	// invocationPayload
	InvocationPayload *string `json:"invocationPayload,omitempty" xml:"invocationPayload,omitempty"`
	// qualifier
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// serviceName
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	// startedTime
	StartedTime *int64 `json:"startedTime,omitempty" xml:"startedTime,omitempty"`
	// status
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s StatefulAsyncInvocation) String() string {
	return tea.Prettify(s)
}

func (s StatefulAsyncInvocation) GoString() string {
	return s.String()
}

func (s *StatefulAsyncInvocation) SetAlreadyRetriedTimes(v int64) *StatefulAsyncInvocation {
	s.AlreadyRetriedTimes = &v
	return s
}

func (s *StatefulAsyncInvocation) SetDestinationStatus(v string) *StatefulAsyncInvocation {
	s.DestinationStatus = &v
	return s
}

func (s *StatefulAsyncInvocation) SetEndTime(v int64) *StatefulAsyncInvocation {
	s.EndTime = &v
	return s
}

func (s *StatefulAsyncInvocation) SetFunctionName(v string) *StatefulAsyncInvocation {
	s.FunctionName = &v
	return s
}

func (s *StatefulAsyncInvocation) SetInvocationErrorMessage(v string) *StatefulAsyncInvocation {
	s.InvocationErrorMessage = &v
	return s
}

func (s *StatefulAsyncInvocation) SetInvocationId(v string) *StatefulAsyncInvocation {
	s.InvocationId = &v
	return s
}

func (s *StatefulAsyncInvocation) SetInvocationPayload(v string) *StatefulAsyncInvocation {
	s.InvocationPayload = &v
	return s
}

func (s *StatefulAsyncInvocation) SetQualifier(v string) *StatefulAsyncInvocation {
	s.Qualifier = &v
	return s
}

func (s *StatefulAsyncInvocation) SetRequestId(v string) *StatefulAsyncInvocation {
	s.RequestId = &v
	return s
}

func (s *StatefulAsyncInvocation) SetServiceName(v string) *StatefulAsyncInvocation {
	s.ServiceName = &v
	return s
}

func (s *StatefulAsyncInvocation) SetStartedTime(v int64) *StatefulAsyncInvocation {
	s.StartedTime = &v
	return s
}

func (s *StatefulAsyncInvocation) SetStatus(v string) *StatefulAsyncInvocation {
	s.Status = &v
	return s
}

type TargetTrackingPolicies struct {
	// endTime
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// maxCapacity
	MaxCapacity *int64 `json:"maxCapacity,omitempty" xml:"maxCapacity,omitempty"`
	// metricTarget
	MetricTarget *float64 `json:"metricTarget,omitempty" xml:"metricTarget,omitempty"`
	// metricType
	MetricType *string `json:"metricType,omitempty" xml:"metricType,omitempty"`
	// minCapacity
	MinCapacity *int64 `json:"minCapacity,omitempty" xml:"minCapacity,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// startTime
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s TargetTrackingPolicies) String() string {
	return tea.Prettify(s)
}

func (s TargetTrackingPolicies) GoString() string {
	return s.String()
}

func (s *TargetTrackingPolicies) SetEndTime(v string) *TargetTrackingPolicies {
	s.EndTime = &v
	return s
}

func (s *TargetTrackingPolicies) SetMaxCapacity(v int64) *TargetTrackingPolicies {
	s.MaxCapacity = &v
	return s
}

func (s *TargetTrackingPolicies) SetMetricTarget(v float64) *TargetTrackingPolicies {
	s.MetricTarget = &v
	return s
}

func (s *TargetTrackingPolicies) SetMetricType(v string) *TargetTrackingPolicies {
	s.MetricType = &v
	return s
}

func (s *TargetTrackingPolicies) SetMinCapacity(v int64) *TargetTrackingPolicies {
	s.MinCapacity = &v
	return s
}

func (s *TargetTrackingPolicies) SetName(v string) *TargetTrackingPolicies {
	s.Name = &v
	return s
}

func (s *TargetTrackingPolicies) SetStartTime(v string) *TargetTrackingPolicies {
	s.StartTime = &v
	return s
}

type TracingConfig struct {
	JaegerConfig *JaegerConfig `json:"jaegerConfig,omitempty" xml:"jaegerConfig,omitempty"`
	// 链路追踪参数
	Params map[string]*string `json:"params,omitempty" xml:"params,omitempty"`
	// 链路追踪类型
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s TracingConfig) String() string {
	return tea.Prettify(s)
}

func (s TracingConfig) GoString() string {
	return s.String()
}

func (s *TracingConfig) SetJaegerConfig(v *JaegerConfig) *TracingConfig {
	s.JaegerConfig = v
	return s
}

func (s *TracingConfig) SetParams(v map[string]*string) *TracingConfig {
	s.Params = v
	return s
}

func (s *TracingConfig) SetType(v string) *TracingConfig {
	s.Type = &v
	return s
}

type VPCConfig struct {
	// Role
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// 安全组ID
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	// VSwitch ID列表
	VSwitchIds []*string `json:"vSwitchIds,omitempty" xml:"vSwitchIds,omitempty" type:"Repeated"`
	// VPC ID
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s VPCConfig) String() string {
	return tea.Prettify(s)
}

func (s VPCConfig) GoString() string {
	return s.String()
}

func (s *VPCConfig) SetRole(v string) *VPCConfig {
	s.Role = &v
	return s
}

func (s *VPCConfig) SetSecurityGroupId(v string) *VPCConfig {
	s.SecurityGroupId = &v
	return s
}

func (s *VPCConfig) SetVSwitchIds(v []*string) *VPCConfig {
	s.VSwitchIds = v
	return s
}

func (s *VPCConfig) SetVpcId(v string) *VPCConfig {
	s.VpcId = &v
	return s
}

type CreateAliasHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s CreateAliasHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateAliasHeaders) GoString() string {
	return s.String()
}

func (s *CreateAliasHeaders) SetCommonHeaders(v map[string]*string) *CreateAliasHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateAliasHeaders) SetXFcAccountId(v string) *CreateAliasHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *CreateAliasHeaders) SetXFcCodeChecksum(v string) *CreateAliasHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *CreateAliasHeaders) SetXFcDate(v string) *CreateAliasHeaders {
	s.XFcDate = &v
	return s
}

func (s *CreateAliasHeaders) SetXFcInvocationType(v string) *CreateAliasHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *CreateAliasHeaders) SetXFcLogType(v string) *CreateAliasHeaders {
	s.XFcLogType = &v
	return s
}

func (s *CreateAliasHeaders) SetXFcTraceId(v string) *CreateAliasHeaders {
	s.XFcTraceId = &v
	return s
}

type CreateAliasRequest struct {
	// 额外版本权重
	AdditionalVersionWeight map[string]*float32 `json:"additionalVersionWeight,omitempty" xml:"additionalVersionWeight,omitempty"`
	// 别名名称
	AliasName *string `json:"aliasName,omitempty" xml:"aliasName,omitempty"`
	// 别名描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 版本ID
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s CreateAliasRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAliasRequest) GoString() string {
	return s.String()
}

func (s *CreateAliasRequest) SetAdditionalVersionWeight(v map[string]*float32) *CreateAliasRequest {
	s.AdditionalVersionWeight = v
	return s
}

func (s *CreateAliasRequest) SetAliasName(v string) *CreateAliasRequest {
	s.AliasName = &v
	return s
}

func (s *CreateAliasRequest) SetDescription(v string) *CreateAliasRequest {
	s.Description = &v
	return s
}

func (s *CreateAliasRequest) SetVersionId(v string) *CreateAliasRequest {
	s.VersionId = &v
	return s
}

type CreateAliasResponseBody struct {
	// 额外版本权重
	AdditionalVersionWeight map[string]*float32 `json:"additionalVersionWeight,omitempty" xml:"additionalVersionWeight,omitempty"`
	// 别名名称
	AliasName *string `json:"aliasName,omitempty" xml:"aliasName,omitempty"`
	// 创建时间
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// 别名描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 上次更新时间
	LastModifiedTime *string `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	// 版本ID
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s CreateAliasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAliasResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAliasResponseBody) SetAdditionalVersionWeight(v map[string]*float32) *CreateAliasResponseBody {
	s.AdditionalVersionWeight = v
	return s
}

func (s *CreateAliasResponseBody) SetAliasName(v string) *CreateAliasResponseBody {
	s.AliasName = &v
	return s
}

func (s *CreateAliasResponseBody) SetCreatedTime(v string) *CreateAliasResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *CreateAliasResponseBody) SetDescription(v string) *CreateAliasResponseBody {
	s.Description = &v
	return s
}

func (s *CreateAliasResponseBody) SetLastModifiedTime(v string) *CreateAliasResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *CreateAliasResponseBody) SetVersionId(v string) *CreateAliasResponseBody {
	s.VersionId = &v
	return s
}

type CreateAliasResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAliasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAliasResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAliasResponse) GoString() string {
	return s.String()
}

func (s *CreateAliasResponse) SetHeaders(v map[string]*string) *CreateAliasResponse {
	s.Headers = v
	return s
}

func (s *CreateAliasResponse) SetBody(v *CreateAliasResponseBody) *CreateAliasResponse {
	s.Body = v
	return s
}

type CreateCustomDomainHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s CreateCustomDomainHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomDomainHeaders) GoString() string {
	return s.String()
}

func (s *CreateCustomDomainHeaders) SetCommonHeaders(v map[string]*string) *CreateCustomDomainHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateCustomDomainHeaders) SetXFcAccountId(v string) *CreateCustomDomainHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *CreateCustomDomainHeaders) SetXFcCodeChecksum(v string) *CreateCustomDomainHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *CreateCustomDomainHeaders) SetXFcDate(v string) *CreateCustomDomainHeaders {
	s.XFcDate = &v
	return s
}

func (s *CreateCustomDomainHeaders) SetXFcInvocationType(v string) *CreateCustomDomainHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *CreateCustomDomainHeaders) SetXFcLogType(v string) *CreateCustomDomainHeaders {
	s.XFcLogType = &v
	return s
}

func (s *CreateCustomDomainHeaders) SetXFcTraceId(v string) *CreateCustomDomainHeaders {
	s.XFcTraceId = &v
	return s
}

type CreateCustomDomainRequest struct {
	CertConfig  *CertConfig  `json:"certConfig,omitempty" xml:"certConfig,omitempty"`
	DomainName  *string      `json:"domainName,omitempty" xml:"domainName,omitempty"`
	Protocol    *string      `json:"protocol,omitempty" xml:"protocol,omitempty"`
	RouteConfig *RouteConfig `json:"routeConfig,omitempty" xml:"routeConfig,omitempty"`
}

func (s CreateCustomDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomDomainRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomDomainRequest) SetCertConfig(v *CertConfig) *CreateCustomDomainRequest {
	s.CertConfig = v
	return s
}

func (s *CreateCustomDomainRequest) SetDomainName(v string) *CreateCustomDomainRequest {
	s.DomainName = &v
	return s
}

func (s *CreateCustomDomainRequest) SetProtocol(v string) *CreateCustomDomainRequest {
	s.Protocol = &v
	return s
}

func (s *CreateCustomDomainRequest) SetRouteConfig(v *RouteConfig) *CreateCustomDomainRequest {
	s.RouteConfig = v
	return s
}

type CreateCustomDomainResponseBody struct {
	AccountId   *string     `json:"accountId,omitempty" xml:"accountId,omitempty"`
	ApiVersion  *string     `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
	CertConfig  *CertConfig `json:"certConfig,omitempty" xml:"certConfig,omitempty"`
	CreatedTime *string     `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// Id of the request
	DomainName       *string      `json:"domainName,omitempty" xml:"domainName,omitempty"`
	LastModifiedTime *string      `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	Protocol         *string      `json:"protocol,omitempty" xml:"protocol,omitempty"`
	RouteConfig      *RouteConfig `json:"routeConfig,omitempty" xml:"routeConfig,omitempty"`
}

func (s CreateCustomDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomDomainResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomDomainResponseBody) SetAccountId(v string) *CreateCustomDomainResponseBody {
	s.AccountId = &v
	return s
}

func (s *CreateCustomDomainResponseBody) SetApiVersion(v string) *CreateCustomDomainResponseBody {
	s.ApiVersion = &v
	return s
}

func (s *CreateCustomDomainResponseBody) SetCertConfig(v *CertConfig) *CreateCustomDomainResponseBody {
	s.CertConfig = v
	return s
}

func (s *CreateCustomDomainResponseBody) SetCreatedTime(v string) *CreateCustomDomainResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *CreateCustomDomainResponseBody) SetDomainName(v string) *CreateCustomDomainResponseBody {
	s.DomainName = &v
	return s
}

func (s *CreateCustomDomainResponseBody) SetLastModifiedTime(v string) *CreateCustomDomainResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *CreateCustomDomainResponseBody) SetProtocol(v string) *CreateCustomDomainResponseBody {
	s.Protocol = &v
	return s
}

func (s *CreateCustomDomainResponseBody) SetRouteConfig(v *RouteConfig) *CreateCustomDomainResponseBody {
	s.RouteConfig = v
	return s
}

type CreateCustomDomainResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCustomDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCustomDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomDomainResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomDomainResponse) SetHeaders(v map[string]*string) *CreateCustomDomainResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomDomainResponse) SetBody(v *CreateCustomDomainResponseBody) *CreateCustomDomainResponse {
	s.Body = v
	return s
}

type CreateFunctionHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s CreateFunctionHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionHeaders) GoString() string {
	return s.String()
}

func (s *CreateFunctionHeaders) SetCommonHeaders(v map[string]*string) *CreateFunctionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateFunctionHeaders) SetXFcAccountId(v string) *CreateFunctionHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *CreateFunctionHeaders) SetXFcCodeChecksum(v string) *CreateFunctionHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *CreateFunctionHeaders) SetXFcDate(v string) *CreateFunctionHeaders {
	s.XFcDate = &v
	return s
}

func (s *CreateFunctionHeaders) SetXFcInvocationType(v string) *CreateFunctionHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *CreateFunctionHeaders) SetXFcLogType(v string) *CreateFunctionHeaders {
	s.XFcLogType = &v
	return s
}

func (s *CreateFunctionHeaders) SetXFcTraceId(v string) *CreateFunctionHeaders {
	s.XFcTraceId = &v
	return s
}

type CreateFunctionRequest struct {
	// 自定义、自定义容器运行时 HTTP Server 的监听端口
	CaPort                *int32                 `json:"caPort,omitempty" xml:"caPort,omitempty"`
	Code                  *Code                  `json:"code,omitempty" xml:"code,omitempty"`
	CustomContainerConfig *CustomContainerConfig `json:"customContainerConfig,omitempty" xml:"customContainerConfig,omitempty"`
	// 函数自定义DNS配置
	CustomDNS *CustomDNS `json:"customDNS,omitempty" xml:"customDNS,omitempty"`
	// Custom Runtime函数详细配置
	CustomRuntimeConfig *CustomRuntimeConfig `json:"customRuntimeConfig,omitempty" xml:"customRuntimeConfig,omitempty"`
	// 函数描述
	Description          *string            `json:"description,omitempty" xml:"description,omitempty"`
	EnvironmentVariables map[string]*string `json:"environmentVariables,omitempty" xml:"environmentVariables,omitempty"`
	// 函数名称
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	// function执行的入口，具体格式和语言相关
	Handler *string `json:"handler,omitempty" xml:"handler,omitempty"`
	// 初始化function运行的超时时间，单位为秒，最小1秒，默认3秒。初始化function超过这个时间后会被终止执行
	InitializationTimeout *int32 `json:"initializationTimeout,omitempty" xml:"initializationTimeout,omitempty"`
	// 初始化 function 执行的入口，具体格式和语言相关
	Initializer             *string                  `json:"initializer,omitempty" xml:"initializer,omitempty"`
	InstanceConcurrency     *int32                   `json:"instanceConcurrency,omitempty" xml:"instanceConcurrency,omitempty"`
	InstanceLifecycleConfig *InstanceLifecycleConfig `json:"instanceLifecycleConfig,omitempty" xml:"instanceLifecycleConfig,omitempty"`
	InstanceType            *string                  `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	// 层列表
	Layers []*string `json:"layers,omitempty" xml:"layers,omitempty" type:"Repeated"`
	// function的内存规格，单位为MB，为64MB的倍数
	MemorySize *int32 `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	// function运行的语言环境，目前支持nodejs6, nodejs8, python2.7, python3, java8
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// function运行的超时时间，单位为秒，最小1秒，默认3秒。function超过这个时间后会被终止执行
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
}

func (s CreateFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionRequest) GoString() string {
	return s.String()
}

func (s *CreateFunctionRequest) SetCaPort(v int32) *CreateFunctionRequest {
	s.CaPort = &v
	return s
}

func (s *CreateFunctionRequest) SetCode(v *Code) *CreateFunctionRequest {
	s.Code = v
	return s
}

func (s *CreateFunctionRequest) SetCustomContainerConfig(v *CustomContainerConfig) *CreateFunctionRequest {
	s.CustomContainerConfig = v
	return s
}

func (s *CreateFunctionRequest) SetCustomDNS(v *CustomDNS) *CreateFunctionRequest {
	s.CustomDNS = v
	return s
}

func (s *CreateFunctionRequest) SetCustomRuntimeConfig(v *CustomRuntimeConfig) *CreateFunctionRequest {
	s.CustomRuntimeConfig = v
	return s
}

func (s *CreateFunctionRequest) SetDescription(v string) *CreateFunctionRequest {
	s.Description = &v
	return s
}

func (s *CreateFunctionRequest) SetEnvironmentVariables(v map[string]*string) *CreateFunctionRequest {
	s.EnvironmentVariables = v
	return s
}

func (s *CreateFunctionRequest) SetFunctionName(v string) *CreateFunctionRequest {
	s.FunctionName = &v
	return s
}

func (s *CreateFunctionRequest) SetHandler(v string) *CreateFunctionRequest {
	s.Handler = &v
	return s
}

func (s *CreateFunctionRequest) SetInitializationTimeout(v int32) *CreateFunctionRequest {
	s.InitializationTimeout = &v
	return s
}

func (s *CreateFunctionRequest) SetInitializer(v string) *CreateFunctionRequest {
	s.Initializer = &v
	return s
}

func (s *CreateFunctionRequest) SetInstanceConcurrency(v int32) *CreateFunctionRequest {
	s.InstanceConcurrency = &v
	return s
}

func (s *CreateFunctionRequest) SetInstanceLifecycleConfig(v *InstanceLifecycleConfig) *CreateFunctionRequest {
	s.InstanceLifecycleConfig = v
	return s
}

func (s *CreateFunctionRequest) SetInstanceType(v string) *CreateFunctionRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateFunctionRequest) SetLayers(v []*string) *CreateFunctionRequest {
	s.Layers = v
	return s
}

func (s *CreateFunctionRequest) SetMemorySize(v int32) *CreateFunctionRequest {
	s.MemorySize = &v
	return s
}

func (s *CreateFunctionRequest) SetRuntime(v string) *CreateFunctionRequest {
	s.Runtime = &v
	return s
}

func (s *CreateFunctionRequest) SetTimeout(v int32) *CreateFunctionRequest {
	s.Timeout = &v
	return s
}

type CreateFunctionResponseBody struct {
	// 自定义、自定义容器运行时 HTTP Server 的监听端口
	CaPort *int32 `json:"caPort,omitempty" xml:"caPort,omitempty"`
	// function code包的CRC64值
	CodeChecksum *string `json:"codeChecksum,omitempty" xml:"codeChecksum,omitempty"`
	// 系统返回的function的code包大小，单位为byte Example : 1024
	CodeSize *int64 `json:"codeSize,omitempty" xml:"codeSize,omitempty"`
	// function创建时间
	CreatedTime           *string                `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	CustomContainerConfig *CustomContainerConfig `json:"customContainerConfig,omitempty" xml:"customContainerConfig,omitempty"`
	// 函数自定义DNS配置
	CustomDNS *CustomDNS `json:"customDNS,omitempty" xml:"customDNS,omitempty"`
	// Custom Runtime函数详细配置
	CustomRuntimeConfig *CustomRuntimeConfig `json:"customRuntimeConfig,omitempty" xml:"customRuntimeConfig,omitempty"`
	// 函数描述
	Description          *string            `json:"description,omitempty" xml:"description,omitempty"`
	EnvironmentVariables map[string]*string `json:"environmentVariables,omitempty" xml:"environmentVariables,omitempty"`
	// 系统为每个function生成的唯一ID
	FunctionId *string `json:"functionId,omitempty" xml:"functionId,omitempty"`
	// 函数名称
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	// function的执行入口
	Handler *string `json:"handler,omitempty" xml:"handler,omitempty"`
	// 初始化function运行的超时时间，单位为秒，最小1秒，默认3秒。初始化function超过这个时间后会被终止执行
	InitializationTimeout *int32 `json:"initializationTimeout,omitempty" xml:"initializationTimeout,omitempty"`
	// 初始化 function 执行的入口，具体格式和语言相关
	Initializer             *string                  `json:"initializer,omitempty" xml:"initializer,omitempty"`
	InstanceConcurrency     *int32                   `json:"instanceConcurrency,omitempty" xml:"instanceConcurrency,omitempty"`
	InstanceLifecycleConfig *InstanceLifecycleConfig `json:"instanceLifecycleConfig,omitempty" xml:"instanceLifecycleConfig,omitempty"`
	InstanceType            *string                  `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	// function上次修改时间
	LastModifiedTime *string   `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	Layers           []*string `json:"layers,omitempty" xml:"layers,omitempty" type:"Repeated"`
	// function设置的内存大小，单位为MB
	MemorySize *int32 `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	// function运行的语言环境，目前支持nodejs6, nodejs8, python2.7, python3, java8
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// 运行的超时时间，单位为秒
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
}

func (s CreateFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFunctionResponseBody) SetCaPort(v int32) *CreateFunctionResponseBody {
	s.CaPort = &v
	return s
}

func (s *CreateFunctionResponseBody) SetCodeChecksum(v string) *CreateFunctionResponseBody {
	s.CodeChecksum = &v
	return s
}

func (s *CreateFunctionResponseBody) SetCodeSize(v int64) *CreateFunctionResponseBody {
	s.CodeSize = &v
	return s
}

func (s *CreateFunctionResponseBody) SetCreatedTime(v string) *CreateFunctionResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *CreateFunctionResponseBody) SetCustomContainerConfig(v *CustomContainerConfig) *CreateFunctionResponseBody {
	s.CustomContainerConfig = v
	return s
}

func (s *CreateFunctionResponseBody) SetCustomDNS(v *CustomDNS) *CreateFunctionResponseBody {
	s.CustomDNS = v
	return s
}

func (s *CreateFunctionResponseBody) SetCustomRuntimeConfig(v *CustomRuntimeConfig) *CreateFunctionResponseBody {
	s.CustomRuntimeConfig = v
	return s
}

func (s *CreateFunctionResponseBody) SetDescription(v string) *CreateFunctionResponseBody {
	s.Description = &v
	return s
}

func (s *CreateFunctionResponseBody) SetEnvironmentVariables(v map[string]*string) *CreateFunctionResponseBody {
	s.EnvironmentVariables = v
	return s
}

func (s *CreateFunctionResponseBody) SetFunctionId(v string) *CreateFunctionResponseBody {
	s.FunctionId = &v
	return s
}

func (s *CreateFunctionResponseBody) SetFunctionName(v string) *CreateFunctionResponseBody {
	s.FunctionName = &v
	return s
}

func (s *CreateFunctionResponseBody) SetHandler(v string) *CreateFunctionResponseBody {
	s.Handler = &v
	return s
}

func (s *CreateFunctionResponseBody) SetInitializationTimeout(v int32) *CreateFunctionResponseBody {
	s.InitializationTimeout = &v
	return s
}

func (s *CreateFunctionResponseBody) SetInitializer(v string) *CreateFunctionResponseBody {
	s.Initializer = &v
	return s
}

func (s *CreateFunctionResponseBody) SetInstanceConcurrency(v int32) *CreateFunctionResponseBody {
	s.InstanceConcurrency = &v
	return s
}

func (s *CreateFunctionResponseBody) SetInstanceLifecycleConfig(v *InstanceLifecycleConfig) *CreateFunctionResponseBody {
	s.InstanceLifecycleConfig = v
	return s
}

func (s *CreateFunctionResponseBody) SetInstanceType(v string) *CreateFunctionResponseBody {
	s.InstanceType = &v
	return s
}

func (s *CreateFunctionResponseBody) SetLastModifiedTime(v string) *CreateFunctionResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *CreateFunctionResponseBody) SetLayers(v []*string) *CreateFunctionResponseBody {
	s.Layers = v
	return s
}

func (s *CreateFunctionResponseBody) SetMemorySize(v int32) *CreateFunctionResponseBody {
	s.MemorySize = &v
	return s
}

func (s *CreateFunctionResponseBody) SetRuntime(v string) *CreateFunctionResponseBody {
	s.Runtime = &v
	return s
}

func (s *CreateFunctionResponseBody) SetTimeout(v int32) *CreateFunctionResponseBody {
	s.Timeout = &v
	return s
}

type CreateFunctionResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionResponse) GoString() string {
	return s.String()
}

func (s *CreateFunctionResponse) SetHeaders(v map[string]*string) *CreateFunctionResponse {
	s.Headers = v
	return s
}

func (s *CreateFunctionResponse) SetBody(v *CreateFunctionResponseBody) *CreateFunctionResponse {
	s.Body = v
	return s
}

type CreateLayerVersionHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s CreateLayerVersionHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateLayerVersionHeaders) GoString() string {
	return s.String()
}

func (s *CreateLayerVersionHeaders) SetCommonHeaders(v map[string]*string) *CreateLayerVersionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateLayerVersionHeaders) SetXFcAccountId(v string) *CreateLayerVersionHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *CreateLayerVersionHeaders) SetXFcCodeChecksum(v string) *CreateLayerVersionHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *CreateLayerVersionHeaders) SetXFcDate(v string) *CreateLayerVersionHeaders {
	s.XFcDate = &v
	return s
}

func (s *CreateLayerVersionHeaders) SetXFcInvocationType(v string) *CreateLayerVersionHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *CreateLayerVersionHeaders) SetXFcLogType(v string) *CreateLayerVersionHeaders {
	s.XFcLogType = &v
	return s
}

func (s *CreateLayerVersionHeaders) SetXFcTraceId(v string) *CreateLayerVersionHeaders {
	s.XFcTraceId = &v
	return s
}

type CreateLayerVersionRequest struct {
	Code              *Code     `json:"Code,omitempty" xml:"Code,omitempty"`
	CompatibleRuntime []*string `json:"compatibleRuntime,omitempty" xml:"compatibleRuntime,omitempty" type:"Repeated"`
	Description       *string   `json:"description,omitempty" xml:"description,omitempty"`
}

func (s CreateLayerVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLayerVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateLayerVersionRequest) SetCode(v *Code) *CreateLayerVersionRequest {
	s.Code = v
	return s
}

func (s *CreateLayerVersionRequest) SetCompatibleRuntime(v []*string) *CreateLayerVersionRequest {
	s.CompatibleRuntime = v
	return s
}

func (s *CreateLayerVersionRequest) SetDescription(v string) *CreateLayerVersionRequest {
	s.Description = &v
	return s
}

type CreateLayerVersionResponseBody struct {
	Acl               *int32              `json:"acl,omitempty" xml:"acl,omitempty"`
	Arn               *string             `json:"arn,omitempty" xml:"arn,omitempty"`
	Code              *OutputCodeLocation `json:"code,omitempty" xml:"code,omitempty"`
	CodeChecksum      *string             `json:"codeChecksum,omitempty" xml:"codeChecksum,omitempty"`
	Codesize          *int64              `json:"codesize,omitempty" xml:"codesize,omitempty"`
	CompatibleRuntime []*string           `json:"compatibleRuntime,omitempty" xml:"compatibleRuntime,omitempty" type:"Repeated"`
	CreateTime        *string             `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description       *string             `json:"description,omitempty" xml:"description,omitempty"`
	LayerName         *string             `json:"layerName,omitempty" xml:"layerName,omitempty"`
	Version           *int32              `json:"version,omitempty" xml:"version,omitempty"`
}

func (s CreateLayerVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLayerVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLayerVersionResponseBody) SetAcl(v int32) *CreateLayerVersionResponseBody {
	s.Acl = &v
	return s
}

func (s *CreateLayerVersionResponseBody) SetArn(v string) *CreateLayerVersionResponseBody {
	s.Arn = &v
	return s
}

func (s *CreateLayerVersionResponseBody) SetCode(v *OutputCodeLocation) *CreateLayerVersionResponseBody {
	s.Code = v
	return s
}

func (s *CreateLayerVersionResponseBody) SetCodeChecksum(v string) *CreateLayerVersionResponseBody {
	s.CodeChecksum = &v
	return s
}

func (s *CreateLayerVersionResponseBody) SetCodesize(v int64) *CreateLayerVersionResponseBody {
	s.Codesize = &v
	return s
}

func (s *CreateLayerVersionResponseBody) SetCompatibleRuntime(v []*string) *CreateLayerVersionResponseBody {
	s.CompatibleRuntime = v
	return s
}

func (s *CreateLayerVersionResponseBody) SetCreateTime(v string) *CreateLayerVersionResponseBody {
	s.CreateTime = &v
	return s
}

func (s *CreateLayerVersionResponseBody) SetDescription(v string) *CreateLayerVersionResponseBody {
	s.Description = &v
	return s
}

func (s *CreateLayerVersionResponseBody) SetLayerName(v string) *CreateLayerVersionResponseBody {
	s.LayerName = &v
	return s
}

func (s *CreateLayerVersionResponseBody) SetVersion(v int32) *CreateLayerVersionResponseBody {
	s.Version = &v
	return s
}

type CreateLayerVersionResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateLayerVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLayerVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLayerVersionResponse) GoString() string {
	return s.String()
}

func (s *CreateLayerVersionResponse) SetHeaders(v map[string]*string) *CreateLayerVersionResponse {
	s.Headers = v
	return s
}

func (s *CreateLayerVersionResponse) SetBody(v *CreateLayerVersionResponseBody) *CreateLayerVersionResponse {
	s.Body = v
	return s
}

type CreateServiceHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s CreateServiceHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceHeaders) GoString() string {
	return s.String()
}

func (s *CreateServiceHeaders) SetCommonHeaders(v map[string]*string) *CreateServiceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateServiceHeaders) SetXFcAccountId(v string) *CreateServiceHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *CreateServiceHeaders) SetXFcCodeChecksum(v string) *CreateServiceHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *CreateServiceHeaders) SetXFcDate(v string) *CreateServiceHeaders {
	s.XFcDate = &v
	return s
}

func (s *CreateServiceHeaders) SetXFcInvocationType(v string) *CreateServiceHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *CreateServiceHeaders) SetXFcLogType(v string) *CreateServiceHeaders {
	s.XFcLogType = &v
	return s
}

func (s *CreateServiceHeaders) SetXFcTraceId(v string) *CreateServiceHeaders {
	s.XFcTraceId = &v
	return s
}

type CreateServiceRequest struct {
	// 服务描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 公网访问设置
	InternetAccess *bool      `json:"internetAccess,omitempty" xml:"internetAccess,omitempty"`
	LogConfig      *LogConfig `json:"logConfig,omitempty" xml:"logConfig,omitempty"`
	NasConfig      *NASConfig `json:"nasConfig,omitempty" xml:"nasConfig,omitempty"`
	// 服务角色
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// 服务名称
	ServiceName   *string        `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	TracingConfig *TracingConfig `json:"tracingConfig,omitempty" xml:"tracingConfig,omitempty"`
	VpcConfig     *VPCConfig     `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty"`
}

func (s CreateServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceRequest) SetDescription(v string) *CreateServiceRequest {
	s.Description = &v
	return s
}

func (s *CreateServiceRequest) SetInternetAccess(v bool) *CreateServiceRequest {
	s.InternetAccess = &v
	return s
}

func (s *CreateServiceRequest) SetLogConfig(v *LogConfig) *CreateServiceRequest {
	s.LogConfig = v
	return s
}

func (s *CreateServiceRequest) SetNasConfig(v *NASConfig) *CreateServiceRequest {
	s.NasConfig = v
	return s
}

func (s *CreateServiceRequest) SetRole(v string) *CreateServiceRequest {
	s.Role = &v
	return s
}

func (s *CreateServiceRequest) SetServiceName(v string) *CreateServiceRequest {
	s.ServiceName = &v
	return s
}

func (s *CreateServiceRequest) SetTracingConfig(v *TracingConfig) *CreateServiceRequest {
	s.TracingConfig = v
	return s
}

func (s *CreateServiceRequest) SetVpcConfig(v *VPCConfig) *CreateServiceRequest {
	s.VpcConfig = v
	return s
}

type CreateServiceResponseBody struct {
	// 创建时间
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// 服务描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 公网访问设置
	InternetAccess *bool `json:"internetAccess,omitempty" xml:"internetAccess,omitempty"`
	// 上次更新时间
	LastModifiedTime *string    `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	LogConfig        *LogConfig `json:"logConfig,omitempty" xml:"logConfig,omitempty"`
	NasConfig        *NASConfig `json:"nasConfig,omitempty" xml:"nasConfig,omitempty"`
	// 服务角色
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// 服务ID
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// 服务名称
	ServiceName   *string        `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	TracingConfig *TracingConfig `json:"tracingConfig,omitempty" xml:"tracingConfig,omitempty"`
	VpcConfig     *VPCConfig     `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty"`
}

func (s CreateServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceResponseBody) SetCreatedTime(v string) *CreateServiceResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *CreateServiceResponseBody) SetDescription(v string) *CreateServiceResponseBody {
	s.Description = &v
	return s
}

func (s *CreateServiceResponseBody) SetInternetAccess(v bool) *CreateServiceResponseBody {
	s.InternetAccess = &v
	return s
}

func (s *CreateServiceResponseBody) SetLastModifiedTime(v string) *CreateServiceResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *CreateServiceResponseBody) SetLogConfig(v *LogConfig) *CreateServiceResponseBody {
	s.LogConfig = v
	return s
}

func (s *CreateServiceResponseBody) SetNasConfig(v *NASConfig) *CreateServiceResponseBody {
	s.NasConfig = v
	return s
}

func (s *CreateServiceResponseBody) SetRole(v string) *CreateServiceResponseBody {
	s.Role = &v
	return s
}

func (s *CreateServiceResponseBody) SetServiceId(v string) *CreateServiceResponseBody {
	s.ServiceId = &v
	return s
}

func (s *CreateServiceResponseBody) SetServiceName(v string) *CreateServiceResponseBody {
	s.ServiceName = &v
	return s
}

func (s *CreateServiceResponseBody) SetTracingConfig(v *TracingConfig) *CreateServiceResponseBody {
	s.TracingConfig = v
	return s
}

func (s *CreateServiceResponseBody) SetVpcConfig(v *VPCConfig) *CreateServiceResponseBody {
	s.VpcConfig = v
	return s
}

type CreateServiceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceResponse) SetHeaders(v map[string]*string) *CreateServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceResponse) SetBody(v *CreateServiceResponseBody) *CreateServiceResponse {
	s.Body = v
	return s
}

type CreateTriggerHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s CreateTriggerHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateTriggerHeaders) GoString() string {
	return s.String()
}

func (s *CreateTriggerHeaders) SetCommonHeaders(v map[string]*string) *CreateTriggerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateTriggerHeaders) SetXFcAccountId(v string) *CreateTriggerHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *CreateTriggerHeaders) SetXFcCodeChecksum(v string) *CreateTriggerHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *CreateTriggerHeaders) SetXFcDate(v string) *CreateTriggerHeaders {
	s.XFcDate = &v
	return s
}

func (s *CreateTriggerHeaders) SetXFcInvocationType(v string) *CreateTriggerHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *CreateTriggerHeaders) SetXFcLogType(v string) *CreateTriggerHeaders {
	s.XFcLogType = &v
	return s
}

func (s *CreateTriggerHeaders) SetXFcTraceId(v string) *CreateTriggerHeaders {
	s.XFcTraceId = &v
	return s
}

type CreateTriggerRequest struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// event source，如OSS，使用该role去invoke function
	InvocationRole *string `json:"invocationRole,omitempty" xml:"invocationRole,omitempty"`
	// service版本
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	// event source的Aliyun Resource Name（ARN
	SourceArn *string `json:"sourceArn,omitempty" xml:"sourceArn,omitempty"`
	// trigger配置，针对不同的trigger类型，trigger配置会有所不同
	TriggerConfig *string `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty"`
	// trigger名称
	TriggerName *string `json:"triggerName,omitempty" xml:"triggerName,omitempty"`
	// trigger类型，如 oss, log, tablestore, timer, http, cdn_events, mns_topic
	TriggerType *string `json:"triggerType,omitempty" xml:"triggerType,omitempty"`
}

func (s CreateTriggerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTriggerRequest) GoString() string {
	return s.String()
}

func (s *CreateTriggerRequest) SetDescription(v string) *CreateTriggerRequest {
	s.Description = &v
	return s
}

func (s *CreateTriggerRequest) SetInvocationRole(v string) *CreateTriggerRequest {
	s.InvocationRole = &v
	return s
}

func (s *CreateTriggerRequest) SetQualifier(v string) *CreateTriggerRequest {
	s.Qualifier = &v
	return s
}

func (s *CreateTriggerRequest) SetSourceArn(v string) *CreateTriggerRequest {
	s.SourceArn = &v
	return s
}

func (s *CreateTriggerRequest) SetTriggerConfig(v string) *CreateTriggerRequest {
	s.TriggerConfig = &v
	return s
}

func (s *CreateTriggerRequest) SetTriggerName(v string) *CreateTriggerRequest {
	s.TriggerName = &v
	return s
}

func (s *CreateTriggerRequest) SetTriggerType(v string) *CreateTriggerRequest {
	s.TriggerType = &v
	return s
}

type CreateTriggerResponseBody struct {
	// 创建时间
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 域名名称，使用域名名称拼接上函数计算域名，可以采用HTTP协议调用到触发器对应版本的函数。例如{domainName}.cn-shanghai.fc.aliyuncs.com
	DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
	// 调用函数使用的RAM角色的ARN
	InvocationRole *string `json:"invocationRole,omitempty" xml:"invocationRole,omitempty"`
	// 上次修改时间
	LastModifiedTime *string `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	// service版本
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	// event source的Aliyun Resource Name（ARN
	SourceArn *string `json:"sourceArn,omitempty" xml:"sourceArn,omitempty"`
	// trigger配置对象
	TriggerConfig *string `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty"`
	TriggerId     *string `json:"triggerId,omitempty" xml:"triggerId,omitempty"`
	// trigger名称
	TriggerName *string `json:"triggerName,omitempty" xml:"triggerName,omitempty"`
	// trigger类型，如 oss, log, tablestore, timer, http, cdn_events, mns_topic
	TriggerType *string `json:"triggerType,omitempty" xml:"triggerType,omitempty"`
}

func (s CreateTriggerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTriggerResponseBody) SetCreatedTime(v string) *CreateTriggerResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *CreateTriggerResponseBody) SetDescription(v string) *CreateTriggerResponseBody {
	s.Description = &v
	return s
}

func (s *CreateTriggerResponseBody) SetDomainName(v string) *CreateTriggerResponseBody {
	s.DomainName = &v
	return s
}

func (s *CreateTriggerResponseBody) SetInvocationRole(v string) *CreateTriggerResponseBody {
	s.InvocationRole = &v
	return s
}

func (s *CreateTriggerResponseBody) SetLastModifiedTime(v string) *CreateTriggerResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *CreateTriggerResponseBody) SetQualifier(v string) *CreateTriggerResponseBody {
	s.Qualifier = &v
	return s
}

func (s *CreateTriggerResponseBody) SetSourceArn(v string) *CreateTriggerResponseBody {
	s.SourceArn = &v
	return s
}

func (s *CreateTriggerResponseBody) SetTriggerConfig(v string) *CreateTriggerResponseBody {
	s.TriggerConfig = &v
	return s
}

func (s *CreateTriggerResponseBody) SetTriggerId(v string) *CreateTriggerResponseBody {
	s.TriggerId = &v
	return s
}

func (s *CreateTriggerResponseBody) SetTriggerName(v string) *CreateTriggerResponseBody {
	s.TriggerName = &v
	return s
}

func (s *CreateTriggerResponseBody) SetTriggerType(v string) *CreateTriggerResponseBody {
	s.TriggerType = &v
	return s
}

type CreateTriggerResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTriggerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTriggerResponse) GoString() string {
	return s.String()
}

func (s *CreateTriggerResponse) SetHeaders(v map[string]*string) *CreateTriggerResponse {
	s.Headers = v
	return s
}

func (s *CreateTriggerResponse) SetBody(v *CreateTriggerResponseBody) *CreateTriggerResponse {
	s.Body = v
	return s
}

type CreateVpcBindingHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s CreateVpcBindingHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcBindingHeaders) GoString() string {
	return s.String()
}

func (s *CreateVpcBindingHeaders) SetCommonHeaders(v map[string]*string) *CreateVpcBindingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateVpcBindingHeaders) SetXFcAccountId(v string) *CreateVpcBindingHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *CreateVpcBindingHeaders) SetXFcCodeChecksum(v string) *CreateVpcBindingHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *CreateVpcBindingHeaders) SetXFcDate(v string) *CreateVpcBindingHeaders {
	s.XFcDate = &v
	return s
}

func (s *CreateVpcBindingHeaders) SetXFcInvocationType(v string) *CreateVpcBindingHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *CreateVpcBindingHeaders) SetXFcLogType(v string) *CreateVpcBindingHeaders {
	s.XFcLogType = &v
	return s
}

func (s *CreateVpcBindingHeaders) SetXFcTraceId(v string) *CreateVpcBindingHeaders {
	s.XFcTraceId = &v
	return s
}

type CreateVpcBindingRequest struct {
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s CreateVpcBindingRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcBindingRequest) GoString() string {
	return s.String()
}

func (s *CreateVpcBindingRequest) SetVpcId(v string) *CreateVpcBindingRequest {
	s.VpcId = &v
	return s
}

type CreateVpcBindingResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s CreateVpcBindingResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcBindingResponse) GoString() string {
	return s.String()
}

func (s *CreateVpcBindingResponse) SetHeaders(v map[string]*string) *CreateVpcBindingResponse {
	s.Headers = v
	return s
}

type DeleteAliasHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	IfMatch           *string            `json:"If-Match,omitempty" xml:"If-Match,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s DeleteAliasHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteAliasHeaders) GoString() string {
	return s.String()
}

func (s *DeleteAliasHeaders) SetCommonHeaders(v map[string]*string) *DeleteAliasHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteAliasHeaders) SetIfMatch(v string) *DeleteAliasHeaders {
	s.IfMatch = &v
	return s
}

func (s *DeleteAliasHeaders) SetXFcAccountId(v string) *DeleteAliasHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *DeleteAliasHeaders) SetXFcCodeChecksum(v string) *DeleteAliasHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *DeleteAliasHeaders) SetXFcDate(v string) *DeleteAliasHeaders {
	s.XFcDate = &v
	return s
}

func (s *DeleteAliasHeaders) SetXFcInvocationType(v string) *DeleteAliasHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *DeleteAliasHeaders) SetXFcLogType(v string) *DeleteAliasHeaders {
	s.XFcLogType = &v
	return s
}

func (s *DeleteAliasHeaders) SetXFcTraceId(v string) *DeleteAliasHeaders {
	s.XFcTraceId = &v
	return s
}

type DeleteAliasResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteAliasResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAliasResponse) GoString() string {
	return s.String()
}

func (s *DeleteAliasResponse) SetHeaders(v map[string]*string) *DeleteAliasResponse {
	s.Headers = v
	return s
}

type DeleteCustomDomainHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s DeleteCustomDomainHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomDomainHeaders) GoString() string {
	return s.String()
}

func (s *DeleteCustomDomainHeaders) SetCommonHeaders(v map[string]*string) *DeleteCustomDomainHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteCustomDomainHeaders) SetXFcAccountId(v string) *DeleteCustomDomainHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *DeleteCustomDomainHeaders) SetXFcCodeChecksum(v string) *DeleteCustomDomainHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *DeleteCustomDomainHeaders) SetXFcDate(v string) *DeleteCustomDomainHeaders {
	s.XFcDate = &v
	return s
}

func (s *DeleteCustomDomainHeaders) SetXFcInvocationType(v string) *DeleteCustomDomainHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *DeleteCustomDomainHeaders) SetXFcLogType(v string) *DeleteCustomDomainHeaders {
	s.XFcLogType = &v
	return s
}

func (s *DeleteCustomDomainHeaders) SetXFcTraceId(v string) *DeleteCustomDomainHeaders {
	s.XFcTraceId = &v
	return s
}

type DeleteCustomDomainResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteCustomDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomDomainResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomDomainResponse) SetHeaders(v map[string]*string) *DeleteCustomDomainResponse {
	s.Headers = v
	return s
}

type DeleteFunctionHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// 用于确保实际更改的资源和期望更改的资源是一致的，该值来自Create，Get和Update API的响应
	IfMatch           *string `json:"If-Match,omitempty" xml:"If-Match,omitempty"`
	XFcAccountId      *string `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s DeleteFunctionHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteFunctionHeaders) GoString() string {
	return s.String()
}

func (s *DeleteFunctionHeaders) SetCommonHeaders(v map[string]*string) *DeleteFunctionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteFunctionHeaders) SetIfMatch(v string) *DeleteFunctionHeaders {
	s.IfMatch = &v
	return s
}

func (s *DeleteFunctionHeaders) SetXFcAccountId(v string) *DeleteFunctionHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *DeleteFunctionHeaders) SetXFcCodeChecksum(v string) *DeleteFunctionHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *DeleteFunctionHeaders) SetXFcDate(v string) *DeleteFunctionHeaders {
	s.XFcDate = &v
	return s
}

func (s *DeleteFunctionHeaders) SetXFcInvocationType(v string) *DeleteFunctionHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *DeleteFunctionHeaders) SetXFcLogType(v string) *DeleteFunctionHeaders {
	s.XFcLogType = &v
	return s
}

func (s *DeleteFunctionHeaders) SetXFcTraceId(v string) *DeleteFunctionHeaders {
	s.XFcTraceId = &v
	return s
}

type DeleteFunctionResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFunctionResponse) GoString() string {
	return s.String()
}

func (s *DeleteFunctionResponse) SetHeaders(v map[string]*string) *DeleteFunctionResponse {
	s.Headers = v
	return s
}

type DeleteFunctionAsyncInvokeConfigHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s DeleteFunctionAsyncInvokeConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteFunctionAsyncInvokeConfigHeaders) GoString() string {
	return s.String()
}

func (s *DeleteFunctionAsyncInvokeConfigHeaders) SetCommonHeaders(v map[string]*string) *DeleteFunctionAsyncInvokeConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteFunctionAsyncInvokeConfigHeaders) SetXFcAccountId(v string) *DeleteFunctionAsyncInvokeConfigHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *DeleteFunctionAsyncInvokeConfigHeaders) SetXFcCodeChecksum(v string) *DeleteFunctionAsyncInvokeConfigHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *DeleteFunctionAsyncInvokeConfigHeaders) SetXFcDate(v string) *DeleteFunctionAsyncInvokeConfigHeaders {
	s.XFcDate = &v
	return s
}

func (s *DeleteFunctionAsyncInvokeConfigHeaders) SetXFcInvocationType(v string) *DeleteFunctionAsyncInvokeConfigHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *DeleteFunctionAsyncInvokeConfigHeaders) SetXFcLogType(v string) *DeleteFunctionAsyncInvokeConfigHeaders {
	s.XFcLogType = &v
	return s
}

func (s *DeleteFunctionAsyncInvokeConfigHeaders) SetXFcTraceId(v string) *DeleteFunctionAsyncInvokeConfigHeaders {
	s.XFcTraceId = &v
	return s
}

type DeleteFunctionAsyncInvokeConfigRequest struct {
	// 限定符
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s DeleteFunctionAsyncInvokeConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFunctionAsyncInvokeConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteFunctionAsyncInvokeConfigRequest) SetQualifier(v string) *DeleteFunctionAsyncInvokeConfigRequest {
	s.Qualifier = &v
	return s
}

type DeleteFunctionAsyncInvokeConfigResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteFunctionAsyncInvokeConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFunctionAsyncInvokeConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteFunctionAsyncInvokeConfigResponse) SetHeaders(v map[string]*string) *DeleteFunctionAsyncInvokeConfigResponse {
	s.Headers = v
	return s
}

type DeleteFunctionOnDemandConfigHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	IfMatch           *string            `json:"If-Match,omitempty" xml:"If-Match,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s DeleteFunctionOnDemandConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteFunctionOnDemandConfigHeaders) GoString() string {
	return s.String()
}

func (s *DeleteFunctionOnDemandConfigHeaders) SetCommonHeaders(v map[string]*string) *DeleteFunctionOnDemandConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteFunctionOnDemandConfigHeaders) SetIfMatch(v string) *DeleteFunctionOnDemandConfigHeaders {
	s.IfMatch = &v
	return s
}

func (s *DeleteFunctionOnDemandConfigHeaders) SetXFcAccountId(v string) *DeleteFunctionOnDemandConfigHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *DeleteFunctionOnDemandConfigHeaders) SetXFcCodeChecksum(v string) *DeleteFunctionOnDemandConfigHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *DeleteFunctionOnDemandConfigHeaders) SetXFcDate(v string) *DeleteFunctionOnDemandConfigHeaders {
	s.XFcDate = &v
	return s
}

func (s *DeleteFunctionOnDemandConfigHeaders) SetXFcInvocationType(v string) *DeleteFunctionOnDemandConfigHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *DeleteFunctionOnDemandConfigHeaders) SetXFcLogType(v string) *DeleteFunctionOnDemandConfigHeaders {
	s.XFcLogType = &v
	return s
}

func (s *DeleteFunctionOnDemandConfigHeaders) SetXFcTraceId(v string) *DeleteFunctionOnDemandConfigHeaders {
	s.XFcTraceId = &v
	return s
}

type DeleteFunctionOnDemandConfigRequest struct {
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s DeleteFunctionOnDemandConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFunctionOnDemandConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteFunctionOnDemandConfigRequest) SetQualifier(v string) *DeleteFunctionOnDemandConfigRequest {
	s.Qualifier = &v
	return s
}

type DeleteFunctionOnDemandConfigResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteFunctionOnDemandConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFunctionOnDemandConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteFunctionOnDemandConfigResponse) SetHeaders(v map[string]*string) *DeleteFunctionOnDemandConfigResponse {
	s.Headers = v
	return s
}

type DeleteLayerVersionHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s DeleteLayerVersionHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteLayerVersionHeaders) GoString() string {
	return s.String()
}

func (s *DeleteLayerVersionHeaders) SetCommonHeaders(v map[string]*string) *DeleteLayerVersionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteLayerVersionHeaders) SetXFcAccountId(v string) *DeleteLayerVersionHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *DeleteLayerVersionHeaders) SetXFcCodeChecksum(v string) *DeleteLayerVersionHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *DeleteLayerVersionHeaders) SetXFcDate(v string) *DeleteLayerVersionHeaders {
	s.XFcDate = &v
	return s
}

func (s *DeleteLayerVersionHeaders) SetXFcInvocationType(v string) *DeleteLayerVersionHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *DeleteLayerVersionHeaders) SetXFcLogType(v string) *DeleteLayerVersionHeaders {
	s.XFcLogType = &v
	return s
}

func (s *DeleteLayerVersionHeaders) SetXFcTraceId(v string) *DeleteLayerVersionHeaders {
	s.XFcTraceId = &v
	return s
}

type DeleteLayerVersionResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteLayerVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLayerVersionResponse) GoString() string {
	return s.String()
}

func (s *DeleteLayerVersionResponse) SetHeaders(v map[string]*string) *DeleteLayerVersionResponse {
	s.Headers = v
	return s
}

type DeleteServiceHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// 用于确保实际更改的资源和期望更改的资源是一致的，该值来自Create，Get和Update API的响应
	IfMatch           *string `json:"If-Match,omitempty" xml:"If-Match,omitempty"`
	XFcAccountId      *string `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s DeleteServiceHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceHeaders) GoString() string {
	return s.String()
}

func (s *DeleteServiceHeaders) SetCommonHeaders(v map[string]*string) *DeleteServiceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteServiceHeaders) SetIfMatch(v string) *DeleteServiceHeaders {
	s.IfMatch = &v
	return s
}

func (s *DeleteServiceHeaders) SetXFcAccountId(v string) *DeleteServiceHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *DeleteServiceHeaders) SetXFcCodeChecksum(v string) *DeleteServiceHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *DeleteServiceHeaders) SetXFcDate(v string) *DeleteServiceHeaders {
	s.XFcDate = &v
	return s
}

func (s *DeleteServiceHeaders) SetXFcInvocationType(v string) *DeleteServiceHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *DeleteServiceHeaders) SetXFcLogType(v string) *DeleteServiceHeaders {
	s.XFcLogType = &v
	return s
}

func (s *DeleteServiceHeaders) SetXFcTraceId(v string) *DeleteServiceHeaders {
	s.XFcTraceId = &v
	return s
}

type DeleteServiceResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceResponse) SetHeaders(v map[string]*string) *DeleteServiceResponse {
	s.Headers = v
	return s
}

type DeleteServiceVersionHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s DeleteServiceVersionHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceVersionHeaders) GoString() string {
	return s.String()
}

func (s *DeleteServiceVersionHeaders) SetCommonHeaders(v map[string]*string) *DeleteServiceVersionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteServiceVersionHeaders) SetXFcAccountId(v string) *DeleteServiceVersionHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *DeleteServiceVersionHeaders) SetXFcCodeChecksum(v string) *DeleteServiceVersionHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *DeleteServiceVersionHeaders) SetXFcDate(v string) *DeleteServiceVersionHeaders {
	s.XFcDate = &v
	return s
}

func (s *DeleteServiceVersionHeaders) SetXFcInvocationType(v string) *DeleteServiceVersionHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *DeleteServiceVersionHeaders) SetXFcLogType(v string) *DeleteServiceVersionHeaders {
	s.XFcLogType = &v
	return s
}

func (s *DeleteServiceVersionHeaders) SetXFcTraceId(v string) *DeleteServiceVersionHeaders {
	s.XFcTraceId = &v
	return s
}

type DeleteServiceVersionResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteServiceVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceVersionResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceVersionResponse) SetHeaders(v map[string]*string) *DeleteServiceVersionResponse {
	s.Headers = v
	return s
}

type DeleteTriggerHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// 用于确保实际更改的资源和期望更改的资源是一致的，该值来自Create，Get和Update API的响应
	IfMatch           *string `json:"If-Match,omitempty" xml:"If-Match,omitempty"`
	XFcAccountId      *string `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s DeleteTriggerHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteTriggerHeaders) GoString() string {
	return s.String()
}

func (s *DeleteTriggerHeaders) SetCommonHeaders(v map[string]*string) *DeleteTriggerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteTriggerHeaders) SetIfMatch(v string) *DeleteTriggerHeaders {
	s.IfMatch = &v
	return s
}

func (s *DeleteTriggerHeaders) SetXFcAccountId(v string) *DeleteTriggerHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *DeleteTriggerHeaders) SetXFcCodeChecksum(v string) *DeleteTriggerHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *DeleteTriggerHeaders) SetXFcDate(v string) *DeleteTriggerHeaders {
	s.XFcDate = &v
	return s
}

func (s *DeleteTriggerHeaders) SetXFcInvocationType(v string) *DeleteTriggerHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *DeleteTriggerHeaders) SetXFcLogType(v string) *DeleteTriggerHeaders {
	s.XFcLogType = &v
	return s
}

func (s *DeleteTriggerHeaders) SetXFcTraceId(v string) *DeleteTriggerHeaders {
	s.XFcTraceId = &v
	return s
}

type DeleteTriggerResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteTriggerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTriggerResponse) GoString() string {
	return s.String()
}

func (s *DeleteTriggerResponse) SetHeaders(v map[string]*string) *DeleteTriggerResponse {
	s.Headers = v
	return s
}

type DeleteVpcBindingHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s DeleteVpcBindingHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpcBindingHeaders) GoString() string {
	return s.String()
}

func (s *DeleteVpcBindingHeaders) SetCommonHeaders(v map[string]*string) *DeleteVpcBindingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteVpcBindingHeaders) SetXFcAccountId(v string) *DeleteVpcBindingHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *DeleteVpcBindingHeaders) SetXFcCodeChecksum(v string) *DeleteVpcBindingHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *DeleteVpcBindingHeaders) SetXFcDate(v string) *DeleteVpcBindingHeaders {
	s.XFcDate = &v
	return s
}

func (s *DeleteVpcBindingHeaders) SetXFcInvocationType(v string) *DeleteVpcBindingHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *DeleteVpcBindingHeaders) SetXFcLogType(v string) *DeleteVpcBindingHeaders {
	s.XFcLogType = &v
	return s
}

func (s *DeleteVpcBindingHeaders) SetXFcTraceId(v string) *DeleteVpcBindingHeaders {
	s.XFcTraceId = &v
	return s
}

type DeleteVpcBindingResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteVpcBindingResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpcBindingResponse) GoString() string {
	return s.String()
}

func (s *DeleteVpcBindingResponse) SetHeaders(v map[string]*string) *DeleteVpcBindingResponse {
	s.Headers = v
	return s
}

type DeregisterEventSourceHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s DeregisterEventSourceHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeregisterEventSourceHeaders) GoString() string {
	return s.String()
}

func (s *DeregisterEventSourceHeaders) SetCommonHeaders(v map[string]*string) *DeregisterEventSourceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeregisterEventSourceHeaders) SetXFcAccountId(v string) *DeregisterEventSourceHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *DeregisterEventSourceHeaders) SetXFcCodeChecksum(v string) *DeregisterEventSourceHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *DeregisterEventSourceHeaders) SetXFcDate(v string) *DeregisterEventSourceHeaders {
	s.XFcDate = &v
	return s
}

func (s *DeregisterEventSourceHeaders) SetXFcInvocationType(v string) *DeregisterEventSourceHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *DeregisterEventSourceHeaders) SetXFcLogType(v string) *DeregisterEventSourceHeaders {
	s.XFcLogType = &v
	return s
}

func (s *DeregisterEventSourceHeaders) SetXFcTraceId(v string) *DeregisterEventSourceHeaders {
	s.XFcTraceId = &v
	return s
}

type DeregisterEventSourceRequest struct {
	// 别名或版本
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s DeregisterEventSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeregisterEventSourceRequest) GoString() string {
	return s.String()
}

func (s *DeregisterEventSourceRequest) SetQualifier(v string) *DeregisterEventSourceRequest {
	s.Qualifier = &v
	return s
}

type DeregisterEventSourceResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeregisterEventSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeregisterEventSourceResponse) GoString() string {
	return s.String()
}

func (s *DeregisterEventSourceResponse) SetHeaders(v map[string]*string) *DeregisterEventSourceResponse {
	s.Headers = v
	return s
}

type GetAccountSettingsHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s GetAccountSettingsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAccountSettingsHeaders) GoString() string {
	return s.String()
}

func (s *GetAccountSettingsHeaders) SetCommonHeaders(v map[string]*string) *GetAccountSettingsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAccountSettingsHeaders) SetXFcAccountId(v string) *GetAccountSettingsHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *GetAccountSettingsHeaders) SetXFcCodeChecksum(v string) *GetAccountSettingsHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *GetAccountSettingsHeaders) SetXFcDate(v string) *GetAccountSettingsHeaders {
	s.XFcDate = &v
	return s
}

func (s *GetAccountSettingsHeaders) SetXFcInvocationType(v string) *GetAccountSettingsHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *GetAccountSettingsHeaders) SetXFcLogType(v string) *GetAccountSettingsHeaders {
	s.XFcLogType = &v
	return s
}

func (s *GetAccountSettingsHeaders) SetXFcTraceId(v string) *GetAccountSettingsHeaders {
	s.XFcTraceId = &v
	return s
}

type GetAccountSettingsResponseBody struct {
	// 可用区列表
	AvailableAZs []*string `json:"availableAZs,omitempty" xml:"availableAZs,omitempty" type:"Repeated"`
}

func (s GetAccountSettingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccountSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountSettingsResponseBody) SetAvailableAZs(v []*string) *GetAccountSettingsResponseBody {
	s.AvailableAZs = v
	return s
}

type GetAccountSettingsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAccountSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAccountSettingsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccountSettingsResponse) GoString() string {
	return s.String()
}

func (s *GetAccountSettingsResponse) SetHeaders(v map[string]*string) *GetAccountSettingsResponse {
	s.Headers = v
	return s
}

func (s *GetAccountSettingsResponse) SetBody(v *GetAccountSettingsResponseBody) *GetAccountSettingsResponse {
	s.Body = v
	return s
}

type GetAliasHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s GetAliasHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAliasHeaders) GoString() string {
	return s.String()
}

func (s *GetAliasHeaders) SetCommonHeaders(v map[string]*string) *GetAliasHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAliasHeaders) SetXFcAccountId(v string) *GetAliasHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *GetAliasHeaders) SetXFcCodeChecksum(v string) *GetAliasHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *GetAliasHeaders) SetXFcDate(v string) *GetAliasHeaders {
	s.XFcDate = &v
	return s
}

func (s *GetAliasHeaders) SetXFcInvocationType(v string) *GetAliasHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *GetAliasHeaders) SetXFcLogType(v string) *GetAliasHeaders {
	s.XFcLogType = &v
	return s
}

func (s *GetAliasHeaders) SetXFcTraceId(v string) *GetAliasHeaders {
	s.XFcTraceId = &v
	return s
}

type GetAliasResponseBody struct {
	// 额外版本权重
	AdditionalVersionWeight map[string]*float32 `json:"additionalVersionWeight,omitempty" xml:"additionalVersionWeight,omitempty"`
	// 别名名称
	AliasName *string `json:"aliasName,omitempty" xml:"aliasName,omitempty"`
	// 创建时间
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// 别名描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 上次更新时间
	LastModifiedTime *string `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	// 版本ID
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s GetAliasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAliasResponseBody) GoString() string {
	return s.String()
}

func (s *GetAliasResponseBody) SetAdditionalVersionWeight(v map[string]*float32) *GetAliasResponseBody {
	s.AdditionalVersionWeight = v
	return s
}

func (s *GetAliasResponseBody) SetAliasName(v string) *GetAliasResponseBody {
	s.AliasName = &v
	return s
}

func (s *GetAliasResponseBody) SetCreatedTime(v string) *GetAliasResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *GetAliasResponseBody) SetDescription(v string) *GetAliasResponseBody {
	s.Description = &v
	return s
}

func (s *GetAliasResponseBody) SetLastModifiedTime(v string) *GetAliasResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *GetAliasResponseBody) SetVersionId(v string) *GetAliasResponseBody {
	s.VersionId = &v
	return s
}

type GetAliasResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAliasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAliasResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAliasResponse) GoString() string {
	return s.String()
}

func (s *GetAliasResponse) SetHeaders(v map[string]*string) *GetAliasResponse {
	s.Headers = v
	return s
}

func (s *GetAliasResponse) SetBody(v *GetAliasResponseBody) *GetAliasResponse {
	s.Body = v
	return s
}

type GetCustomDomainHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s GetCustomDomainHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCustomDomainHeaders) GoString() string {
	return s.String()
}

func (s *GetCustomDomainHeaders) SetCommonHeaders(v map[string]*string) *GetCustomDomainHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCustomDomainHeaders) SetXFcAccountId(v string) *GetCustomDomainHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *GetCustomDomainHeaders) SetXFcCodeChecksum(v string) *GetCustomDomainHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *GetCustomDomainHeaders) SetXFcDate(v string) *GetCustomDomainHeaders {
	s.XFcDate = &v
	return s
}

func (s *GetCustomDomainHeaders) SetXFcInvocationType(v string) *GetCustomDomainHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *GetCustomDomainHeaders) SetXFcLogType(v string) *GetCustomDomainHeaders {
	s.XFcLogType = &v
	return s
}

func (s *GetCustomDomainHeaders) SetXFcTraceId(v string) *GetCustomDomainHeaders {
	s.XFcTraceId = &v
	return s
}

type GetCustomDomainResponseBody struct {
	AccountId   *string     `json:"accountId,omitempty" xml:"accountId,omitempty"`
	ApiVersion  *string     `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
	CertConfig  *CertConfig `json:"certConfig,omitempty" xml:"certConfig,omitempty"`
	CreatedTime *string     `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// Id of the request
	DomainName       *string      `json:"domainName,omitempty" xml:"domainName,omitempty"`
	LastModifiedTime *string      `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	Protocol         *string      `json:"protocol,omitempty" xml:"protocol,omitempty"`
	RouteConfig      *RouteConfig `json:"routeConfig,omitempty" xml:"routeConfig,omitempty"`
}

func (s GetCustomDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCustomDomainResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomDomainResponseBody) SetAccountId(v string) *GetCustomDomainResponseBody {
	s.AccountId = &v
	return s
}

func (s *GetCustomDomainResponseBody) SetApiVersion(v string) *GetCustomDomainResponseBody {
	s.ApiVersion = &v
	return s
}

func (s *GetCustomDomainResponseBody) SetCertConfig(v *CertConfig) *GetCustomDomainResponseBody {
	s.CertConfig = v
	return s
}

func (s *GetCustomDomainResponseBody) SetCreatedTime(v string) *GetCustomDomainResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *GetCustomDomainResponseBody) SetDomainName(v string) *GetCustomDomainResponseBody {
	s.DomainName = &v
	return s
}

func (s *GetCustomDomainResponseBody) SetLastModifiedTime(v string) *GetCustomDomainResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *GetCustomDomainResponseBody) SetProtocol(v string) *GetCustomDomainResponseBody {
	s.Protocol = &v
	return s
}

func (s *GetCustomDomainResponseBody) SetRouteConfig(v *RouteConfig) *GetCustomDomainResponseBody {
	s.RouteConfig = v
	return s
}

type GetCustomDomainResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCustomDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCustomDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCustomDomainResponse) GoString() string {
	return s.String()
}

func (s *GetCustomDomainResponse) SetHeaders(v map[string]*string) *GetCustomDomainResponse {
	s.Headers = v
	return s
}

func (s *GetCustomDomainResponse) SetBody(v *GetCustomDomainResponseBody) *GetCustomDomainResponse {
	s.Body = v
	return s
}

type GetFunctionHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s GetFunctionHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionHeaders) GoString() string {
	return s.String()
}

func (s *GetFunctionHeaders) SetCommonHeaders(v map[string]*string) *GetFunctionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFunctionHeaders) SetXFcAccountId(v string) *GetFunctionHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *GetFunctionHeaders) SetXFcCodeChecksum(v string) *GetFunctionHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *GetFunctionHeaders) SetXFcDate(v string) *GetFunctionHeaders {
	s.XFcDate = &v
	return s
}

func (s *GetFunctionHeaders) SetXFcInvocationType(v string) *GetFunctionHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *GetFunctionHeaders) SetXFcLogType(v string) *GetFunctionHeaders {
	s.XFcLogType = &v
	return s
}

func (s *GetFunctionHeaders) SetXFcTraceId(v string) *GetFunctionHeaders {
	s.XFcTraceId = &v
	return s
}

type GetFunctionRequest struct {
	// service版本, 可以是versionId或者aliasName
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s GetFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionRequest) GoString() string {
	return s.String()
}

func (s *GetFunctionRequest) SetQualifier(v string) *GetFunctionRequest {
	s.Qualifier = &v
	return s
}

type GetFunctionResponseBody struct {
	// 自定义、自定义容器运行时 HTTP Server 的监听端口
	CaPort *int32 `json:"caPort,omitempty" xml:"caPort,omitempty"`
	// function code包的CRC64值
	CodeChecksum *string `json:"codeChecksum,omitempty" xml:"codeChecksum,omitempty"`
	// 系统返回的function的code包大小，单位为byte Example : 1024
	CodeSize *int64 `json:"codeSize,omitempty" xml:"codeSize,omitempty"`
	// function创建时间
	CreatedTime           *string                    `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	CustomContainerConfig *CustomContainerConfigInfo `json:"customContainerConfig,omitempty" xml:"customContainerConfig,omitempty"`
	// 函数自定义DNS配置
	CustomDNS *CustomDNS `json:"customDNS,omitempty" xml:"customDNS,omitempty"`
	// Custom Runtime函数详细配置
	CustomRuntimeConfig *CustomRuntimeConfig `json:"customRuntimeConfig,omitempty" xml:"customRuntimeConfig,omitempty"`
	// 函数描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 为函数设置的环境变量，可以在函数中获取环境变量的值
	EnvironmentVariables map[string]*string `json:"environmentVariables,omitempty" xml:"environmentVariables,omitempty"`
	// 系统为每个function生成的唯一ID
	FunctionId *string `json:"functionId,omitempty" xml:"functionId,omitempty"`
	// 函数名称
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	// function的执行入口
	Handler *string `json:"handler,omitempty" xml:"handler,omitempty"`
	// 初始化function运行的超时时间，单位为秒，最小1秒，默认3秒。初始化function超过这个时间后会被终止执行
	InitializationTimeout *int32 `json:"initializationTimeout,omitempty" xml:"initializationTimeout,omitempty"`
	// 初始化 function 执行的入口，具体格式和语言相关
	Initializer             *string                  `json:"initializer,omitempty" xml:"initializer,omitempty"`
	InstanceConcurrency     *int32                   `json:"instanceConcurrency,omitempty" xml:"instanceConcurrency,omitempty"`
	InstanceLifecycleConfig *InstanceLifecycleConfig `json:"instanceLifecycleConfig,omitempty" xml:"instanceLifecycleConfig,omitempty"`
	InstanceType            *string                  `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	// function上次修改时间
	LastModifiedTime *string   `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	Layers           []*string `json:"layers,omitempty" xml:"layers,omitempty" type:"Repeated"`
	// function设置的内存大小，单位为MB
	MemorySize *int32 `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	// function运行的语言环境，目前支持nodejs6, nodejs8, python2.7, python3, java8
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// 运行的超时时间，单位为秒
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
}

func (s GetFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *GetFunctionResponseBody) SetCaPort(v int32) *GetFunctionResponseBody {
	s.CaPort = &v
	return s
}

func (s *GetFunctionResponseBody) SetCodeChecksum(v string) *GetFunctionResponseBody {
	s.CodeChecksum = &v
	return s
}

func (s *GetFunctionResponseBody) SetCodeSize(v int64) *GetFunctionResponseBody {
	s.CodeSize = &v
	return s
}

func (s *GetFunctionResponseBody) SetCreatedTime(v string) *GetFunctionResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *GetFunctionResponseBody) SetCustomContainerConfig(v *CustomContainerConfigInfo) *GetFunctionResponseBody {
	s.CustomContainerConfig = v
	return s
}

func (s *GetFunctionResponseBody) SetCustomDNS(v *CustomDNS) *GetFunctionResponseBody {
	s.CustomDNS = v
	return s
}

func (s *GetFunctionResponseBody) SetCustomRuntimeConfig(v *CustomRuntimeConfig) *GetFunctionResponseBody {
	s.CustomRuntimeConfig = v
	return s
}

func (s *GetFunctionResponseBody) SetDescription(v string) *GetFunctionResponseBody {
	s.Description = &v
	return s
}

func (s *GetFunctionResponseBody) SetEnvironmentVariables(v map[string]*string) *GetFunctionResponseBody {
	s.EnvironmentVariables = v
	return s
}

func (s *GetFunctionResponseBody) SetFunctionId(v string) *GetFunctionResponseBody {
	s.FunctionId = &v
	return s
}

func (s *GetFunctionResponseBody) SetFunctionName(v string) *GetFunctionResponseBody {
	s.FunctionName = &v
	return s
}

func (s *GetFunctionResponseBody) SetHandler(v string) *GetFunctionResponseBody {
	s.Handler = &v
	return s
}

func (s *GetFunctionResponseBody) SetInitializationTimeout(v int32) *GetFunctionResponseBody {
	s.InitializationTimeout = &v
	return s
}

func (s *GetFunctionResponseBody) SetInitializer(v string) *GetFunctionResponseBody {
	s.Initializer = &v
	return s
}

func (s *GetFunctionResponseBody) SetInstanceConcurrency(v int32) *GetFunctionResponseBody {
	s.InstanceConcurrency = &v
	return s
}

func (s *GetFunctionResponseBody) SetInstanceLifecycleConfig(v *InstanceLifecycleConfig) *GetFunctionResponseBody {
	s.InstanceLifecycleConfig = v
	return s
}

func (s *GetFunctionResponseBody) SetInstanceType(v string) *GetFunctionResponseBody {
	s.InstanceType = &v
	return s
}

func (s *GetFunctionResponseBody) SetLastModifiedTime(v string) *GetFunctionResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *GetFunctionResponseBody) SetLayers(v []*string) *GetFunctionResponseBody {
	s.Layers = v
	return s
}

func (s *GetFunctionResponseBody) SetMemorySize(v int32) *GetFunctionResponseBody {
	s.MemorySize = &v
	return s
}

func (s *GetFunctionResponseBody) SetRuntime(v string) *GetFunctionResponseBody {
	s.Runtime = &v
	return s
}

func (s *GetFunctionResponseBody) SetTimeout(v int32) *GetFunctionResponseBody {
	s.Timeout = &v
	return s
}

type GetFunctionResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionResponse) GoString() string {
	return s.String()
}

func (s *GetFunctionResponse) SetHeaders(v map[string]*string) *GetFunctionResponse {
	s.Headers = v
	return s
}

func (s *GetFunctionResponse) SetBody(v *GetFunctionResponseBody) *GetFunctionResponse {
	s.Body = v
	return s
}

type GetFunctionAsyncInvokeConfigHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s GetFunctionAsyncInvokeConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionAsyncInvokeConfigHeaders) GoString() string {
	return s.String()
}

func (s *GetFunctionAsyncInvokeConfigHeaders) SetCommonHeaders(v map[string]*string) *GetFunctionAsyncInvokeConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFunctionAsyncInvokeConfigHeaders) SetXFcAccountId(v string) *GetFunctionAsyncInvokeConfigHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *GetFunctionAsyncInvokeConfigHeaders) SetXFcCodeChecksum(v string) *GetFunctionAsyncInvokeConfigHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *GetFunctionAsyncInvokeConfigHeaders) SetXFcDate(v string) *GetFunctionAsyncInvokeConfigHeaders {
	s.XFcDate = &v
	return s
}

func (s *GetFunctionAsyncInvokeConfigHeaders) SetXFcInvocationType(v string) *GetFunctionAsyncInvokeConfigHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *GetFunctionAsyncInvokeConfigHeaders) SetXFcLogType(v string) *GetFunctionAsyncInvokeConfigHeaders {
	s.XFcLogType = &v
	return s
}

func (s *GetFunctionAsyncInvokeConfigHeaders) SetXFcTraceId(v string) *GetFunctionAsyncInvokeConfigHeaders {
	s.XFcTraceId = &v
	return s
}

type GetFunctionAsyncInvokeConfigRequest struct {
	// 限定符
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s GetFunctionAsyncInvokeConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionAsyncInvokeConfigRequest) GoString() string {
	return s.String()
}

func (s *GetFunctionAsyncInvokeConfigRequest) SetQualifier(v string) *GetFunctionAsyncInvokeConfigRequest {
	s.Qualifier = &v
	return s
}

type GetFunctionAsyncInvokeConfigResponseBody struct {
	// 创建时间
	CreatedTime       *string            `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	DestinationConfig *DestinationConfig `json:"destinationConfig,omitempty" xml:"destinationConfig,omitempty"`
	// 函数名称
	Function *string `json:"function,omitempty" xml:"function,omitempty"`
	// 最后更改时间
	LastModifiedTime *string `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	// 消息最大存活时长
	MaxAsyncEventAgeInSeconds *int64 `json:"maxAsyncEventAgeInSeconds,omitempty" xml:"maxAsyncEventAgeInSeconds,omitempty"`
	// 异步调用失败后的最大重试次数
	MaxAsyncRetryAttempts *int64 `json:"maxAsyncRetryAttempts,omitempty" xml:"maxAsyncRetryAttempts,omitempty"`
	// 限定符
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	// 服务名称
	Service            *string `json:"service,omitempty" xml:"service,omitempty"`
	StatefulInvocation *bool   `json:"statefulInvocation,omitempty" xml:"statefulInvocation,omitempty"`
}

func (s GetFunctionAsyncInvokeConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionAsyncInvokeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetFunctionAsyncInvokeConfigResponseBody) SetCreatedTime(v string) *GetFunctionAsyncInvokeConfigResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *GetFunctionAsyncInvokeConfigResponseBody) SetDestinationConfig(v *DestinationConfig) *GetFunctionAsyncInvokeConfigResponseBody {
	s.DestinationConfig = v
	return s
}

func (s *GetFunctionAsyncInvokeConfigResponseBody) SetFunction(v string) *GetFunctionAsyncInvokeConfigResponseBody {
	s.Function = &v
	return s
}

func (s *GetFunctionAsyncInvokeConfigResponseBody) SetLastModifiedTime(v string) *GetFunctionAsyncInvokeConfigResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *GetFunctionAsyncInvokeConfigResponseBody) SetMaxAsyncEventAgeInSeconds(v int64) *GetFunctionAsyncInvokeConfigResponseBody {
	s.MaxAsyncEventAgeInSeconds = &v
	return s
}

func (s *GetFunctionAsyncInvokeConfigResponseBody) SetMaxAsyncRetryAttempts(v int64) *GetFunctionAsyncInvokeConfigResponseBody {
	s.MaxAsyncRetryAttempts = &v
	return s
}

func (s *GetFunctionAsyncInvokeConfigResponseBody) SetQualifier(v string) *GetFunctionAsyncInvokeConfigResponseBody {
	s.Qualifier = &v
	return s
}

func (s *GetFunctionAsyncInvokeConfigResponseBody) SetService(v string) *GetFunctionAsyncInvokeConfigResponseBody {
	s.Service = &v
	return s
}

func (s *GetFunctionAsyncInvokeConfigResponseBody) SetStatefulInvocation(v bool) *GetFunctionAsyncInvokeConfigResponseBody {
	s.StatefulInvocation = &v
	return s
}

type GetFunctionAsyncInvokeConfigResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFunctionAsyncInvokeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFunctionAsyncInvokeConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionAsyncInvokeConfigResponse) GoString() string {
	return s.String()
}

func (s *GetFunctionAsyncInvokeConfigResponse) SetHeaders(v map[string]*string) *GetFunctionAsyncInvokeConfigResponse {
	s.Headers = v
	return s
}

func (s *GetFunctionAsyncInvokeConfigResponse) SetBody(v *GetFunctionAsyncInvokeConfigResponseBody) *GetFunctionAsyncInvokeConfigResponse {
	s.Body = v
	return s
}

type GetFunctionCodeHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s GetFunctionCodeHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionCodeHeaders) GoString() string {
	return s.String()
}

func (s *GetFunctionCodeHeaders) SetCommonHeaders(v map[string]*string) *GetFunctionCodeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFunctionCodeHeaders) SetXFcAccountId(v string) *GetFunctionCodeHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *GetFunctionCodeHeaders) SetXFcCodeChecksum(v string) *GetFunctionCodeHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *GetFunctionCodeHeaders) SetXFcDate(v string) *GetFunctionCodeHeaders {
	s.XFcDate = &v
	return s
}

func (s *GetFunctionCodeHeaders) SetXFcInvocationType(v string) *GetFunctionCodeHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *GetFunctionCodeHeaders) SetXFcLogType(v string) *GetFunctionCodeHeaders {
	s.XFcLogType = &v
	return s
}

func (s *GetFunctionCodeHeaders) SetXFcTraceId(v string) *GetFunctionCodeHeaders {
	s.XFcTraceId = &v
	return s
}

type GetFunctionCodeRequest struct {
	// service版本, 可以是versionId或者aliasName
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s GetFunctionCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionCodeRequest) GoString() string {
	return s.String()
}

func (s *GetFunctionCodeRequest) SetQualifier(v string) *GetFunctionCodeRequest {
	s.Qualifier = &v
	return s
}

type GetFunctionCodeResponseBody struct {
	// function code包的CRC64值
	Checksum *string `json:"checksum,omitempty" xml:"checksum,omitempty"`
	// 获取function代码的URL
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetFunctionCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetFunctionCodeResponseBody) SetChecksum(v string) *GetFunctionCodeResponseBody {
	s.Checksum = &v
	return s
}

func (s *GetFunctionCodeResponseBody) SetUrl(v string) *GetFunctionCodeResponseBody {
	s.Url = &v
	return s
}

type GetFunctionCodeResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFunctionCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFunctionCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionCodeResponse) GoString() string {
	return s.String()
}

func (s *GetFunctionCodeResponse) SetHeaders(v map[string]*string) *GetFunctionCodeResponse {
	s.Headers = v
	return s
}

func (s *GetFunctionCodeResponse) SetBody(v *GetFunctionCodeResponseBody) *GetFunctionCodeResponse {
	s.Body = v
	return s
}

type GetFunctionOnDemandConfigHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s GetFunctionOnDemandConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionOnDemandConfigHeaders) GoString() string {
	return s.String()
}

func (s *GetFunctionOnDemandConfigHeaders) SetCommonHeaders(v map[string]*string) *GetFunctionOnDemandConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFunctionOnDemandConfigHeaders) SetXFcAccountId(v string) *GetFunctionOnDemandConfigHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *GetFunctionOnDemandConfigHeaders) SetXFcCodeChecksum(v string) *GetFunctionOnDemandConfigHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *GetFunctionOnDemandConfigHeaders) SetXFcDate(v string) *GetFunctionOnDemandConfigHeaders {
	s.XFcDate = &v
	return s
}

func (s *GetFunctionOnDemandConfigHeaders) SetXFcInvocationType(v string) *GetFunctionOnDemandConfigHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *GetFunctionOnDemandConfigHeaders) SetXFcLogType(v string) *GetFunctionOnDemandConfigHeaders {
	s.XFcLogType = &v
	return s
}

func (s *GetFunctionOnDemandConfigHeaders) SetXFcTraceId(v string) *GetFunctionOnDemandConfigHeaders {
	s.XFcTraceId = &v
	return s
}

type GetFunctionOnDemandConfigRequest struct {
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s GetFunctionOnDemandConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionOnDemandConfigRequest) GoString() string {
	return s.String()
}

func (s *GetFunctionOnDemandConfigRequest) SetQualifier(v string) *GetFunctionOnDemandConfigRequest {
	s.Qualifier = &v
	return s
}

type GetFunctionOnDemandConfigResponseBody struct {
	MaximumInstanceCount *int64 `json:"maximumInstanceCount,omitempty" xml:"maximumInstanceCount,omitempty"`
	// Id of the request
	Resource *string `json:"resource,omitempty" xml:"resource,omitempty"`
}

func (s GetFunctionOnDemandConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionOnDemandConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetFunctionOnDemandConfigResponseBody) SetMaximumInstanceCount(v int64) *GetFunctionOnDemandConfigResponseBody {
	s.MaximumInstanceCount = &v
	return s
}

func (s *GetFunctionOnDemandConfigResponseBody) SetResource(v string) *GetFunctionOnDemandConfigResponseBody {
	s.Resource = &v
	return s
}

type GetFunctionOnDemandConfigResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFunctionOnDemandConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFunctionOnDemandConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionOnDemandConfigResponse) GoString() string {
	return s.String()
}

func (s *GetFunctionOnDemandConfigResponse) SetHeaders(v map[string]*string) *GetFunctionOnDemandConfigResponse {
	s.Headers = v
	return s
}

func (s *GetFunctionOnDemandConfigResponse) SetBody(v *GetFunctionOnDemandConfigResponseBody) *GetFunctionOnDemandConfigResponse {
	s.Body = v
	return s
}

type GetLayerVersionHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s GetLayerVersionHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetLayerVersionHeaders) GoString() string {
	return s.String()
}

func (s *GetLayerVersionHeaders) SetCommonHeaders(v map[string]*string) *GetLayerVersionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetLayerVersionHeaders) SetXFcAccountId(v string) *GetLayerVersionHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *GetLayerVersionHeaders) SetXFcCodeChecksum(v string) *GetLayerVersionHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *GetLayerVersionHeaders) SetXFcDate(v string) *GetLayerVersionHeaders {
	s.XFcDate = &v
	return s
}

func (s *GetLayerVersionHeaders) SetXFcInvocationType(v string) *GetLayerVersionHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *GetLayerVersionHeaders) SetXFcLogType(v string) *GetLayerVersionHeaders {
	s.XFcLogType = &v
	return s
}

func (s *GetLayerVersionHeaders) SetXFcTraceId(v string) *GetLayerVersionHeaders {
	s.XFcTraceId = &v
	return s
}

type GetLayerVersionResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *Layer             `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLayerVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLayerVersionResponse) GoString() string {
	return s.String()
}

func (s *GetLayerVersionResponse) SetHeaders(v map[string]*string) *GetLayerVersionResponse {
	s.Headers = v
	return s
}

func (s *GetLayerVersionResponse) SetBody(v *Layer) *GetLayerVersionResponse {
	s.Body = v
	return s
}

type GetLayerVersionByArnHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s GetLayerVersionByArnHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetLayerVersionByArnHeaders) GoString() string {
	return s.String()
}

func (s *GetLayerVersionByArnHeaders) SetCommonHeaders(v map[string]*string) *GetLayerVersionByArnHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetLayerVersionByArnHeaders) SetXFcAccountId(v string) *GetLayerVersionByArnHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *GetLayerVersionByArnHeaders) SetXFcCodeChecksum(v string) *GetLayerVersionByArnHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *GetLayerVersionByArnHeaders) SetXFcDate(v string) *GetLayerVersionByArnHeaders {
	s.XFcDate = &v
	return s
}

func (s *GetLayerVersionByArnHeaders) SetXFcInvocationType(v string) *GetLayerVersionByArnHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *GetLayerVersionByArnHeaders) SetXFcLogType(v string) *GetLayerVersionByArnHeaders {
	s.XFcLogType = &v
	return s
}

func (s *GetLayerVersionByArnHeaders) SetXFcTraceId(v string) *GetLayerVersionByArnHeaders {
	s.XFcTraceId = &v
	return s
}

type GetLayerVersionByArnResponseBody struct {
	Acl               *int32              `json:"acl,omitempty" xml:"acl,omitempty"`
	Arn               *string             `json:"arn,omitempty" xml:"arn,omitempty"`
	Code              *OutputCodeLocation `json:"code,omitempty" xml:"code,omitempty"`
	CodeChecksum      *string             `json:"codeChecksum,omitempty" xml:"codeChecksum,omitempty"`
	Codesize          *int64              `json:"codesize,omitempty" xml:"codesize,omitempty"`
	CompatibleRuntime []*string           `json:"compatibleRuntime,omitempty" xml:"compatibleRuntime,omitempty" type:"Repeated"`
	CreateTime        *string             `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description       *string             `json:"description,omitempty" xml:"description,omitempty"`
	LayerName         *string             `json:"layerName,omitempty" xml:"layerName,omitempty"`
	Version           *int32              `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetLayerVersionByArnResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLayerVersionByArnResponseBody) GoString() string {
	return s.String()
}

func (s *GetLayerVersionByArnResponseBody) SetAcl(v int32) *GetLayerVersionByArnResponseBody {
	s.Acl = &v
	return s
}

func (s *GetLayerVersionByArnResponseBody) SetArn(v string) *GetLayerVersionByArnResponseBody {
	s.Arn = &v
	return s
}

func (s *GetLayerVersionByArnResponseBody) SetCode(v *OutputCodeLocation) *GetLayerVersionByArnResponseBody {
	s.Code = v
	return s
}

func (s *GetLayerVersionByArnResponseBody) SetCodeChecksum(v string) *GetLayerVersionByArnResponseBody {
	s.CodeChecksum = &v
	return s
}

func (s *GetLayerVersionByArnResponseBody) SetCodesize(v int64) *GetLayerVersionByArnResponseBody {
	s.Codesize = &v
	return s
}

func (s *GetLayerVersionByArnResponseBody) SetCompatibleRuntime(v []*string) *GetLayerVersionByArnResponseBody {
	s.CompatibleRuntime = v
	return s
}

func (s *GetLayerVersionByArnResponseBody) SetCreateTime(v string) *GetLayerVersionByArnResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetLayerVersionByArnResponseBody) SetDescription(v string) *GetLayerVersionByArnResponseBody {
	s.Description = &v
	return s
}

func (s *GetLayerVersionByArnResponseBody) SetLayerName(v string) *GetLayerVersionByArnResponseBody {
	s.LayerName = &v
	return s
}

func (s *GetLayerVersionByArnResponseBody) SetVersion(v int32) *GetLayerVersionByArnResponseBody {
	s.Version = &v
	return s
}

type GetLayerVersionByArnResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetLayerVersionByArnResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLayerVersionByArnResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLayerVersionByArnResponse) GoString() string {
	return s.String()
}

func (s *GetLayerVersionByArnResponse) SetHeaders(v map[string]*string) *GetLayerVersionByArnResponse {
	s.Headers = v
	return s
}

func (s *GetLayerVersionByArnResponse) SetBody(v *GetLayerVersionByArnResponseBody) *GetLayerVersionByArnResponse {
	s.Body = v
	return s
}

type GetProvisionConfigHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s GetProvisionConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetProvisionConfigHeaders) GoString() string {
	return s.String()
}

func (s *GetProvisionConfigHeaders) SetCommonHeaders(v map[string]*string) *GetProvisionConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetProvisionConfigHeaders) SetXFcAccountId(v string) *GetProvisionConfigHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *GetProvisionConfigHeaders) SetXFcCodeChecksum(v string) *GetProvisionConfigHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *GetProvisionConfigHeaders) SetXFcDate(v string) *GetProvisionConfigHeaders {
	s.XFcDate = &v
	return s
}

func (s *GetProvisionConfigHeaders) SetXFcInvocationType(v string) *GetProvisionConfigHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *GetProvisionConfigHeaders) SetXFcLogType(v string) *GetProvisionConfigHeaders {
	s.XFcLogType = &v
	return s
}

func (s *GetProvisionConfigHeaders) SetXFcTraceId(v string) *GetProvisionConfigHeaders {
	s.XFcTraceId = &v
	return s
}

type GetProvisionConfigRequest struct {
	// 别名名称
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s GetProvisionConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProvisionConfigRequest) GoString() string {
	return s.String()
}

func (s *GetProvisionConfigRequest) SetQualifier(v string) *GetProvisionConfigRequest {
	s.Qualifier = &v
	return s
}

type GetProvisionConfigResponseBody struct {
	// 实际资源个数
	Current *int64 `json:"current,omitempty" xml:"current,omitempty"`
	// 预留实例创建失败时的错误信息
	CurrentError *string `json:"currentError,omitempty" xml:"currentError,omitempty"`
	// 资源描述
	Resource *string `json:"resource,omitempty" xml:"resource,omitempty"`
	// 定时策略配置
	ScheduledActions []*ScheduledActions `json:"scheduledActions,omitempty" xml:"scheduledActions,omitempty" type:"Repeated"`
	// 目标资源个数
	Target *int64 `json:"target,omitempty" xml:"target,omitempty"`
	// 指标追踪伸缩策略配置
	TargetTrackingPolicies []*TargetTrackingPolicies `json:"targetTrackingPolicies,omitempty" xml:"targetTrackingPolicies,omitempty" type:"Repeated"`
}

func (s GetProvisionConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProvisionConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetProvisionConfigResponseBody) SetCurrent(v int64) *GetProvisionConfigResponseBody {
	s.Current = &v
	return s
}

func (s *GetProvisionConfigResponseBody) SetCurrentError(v string) *GetProvisionConfigResponseBody {
	s.CurrentError = &v
	return s
}

func (s *GetProvisionConfigResponseBody) SetResource(v string) *GetProvisionConfigResponseBody {
	s.Resource = &v
	return s
}

func (s *GetProvisionConfigResponseBody) SetScheduledActions(v []*ScheduledActions) *GetProvisionConfigResponseBody {
	s.ScheduledActions = v
	return s
}

func (s *GetProvisionConfigResponseBody) SetTarget(v int64) *GetProvisionConfigResponseBody {
	s.Target = &v
	return s
}

func (s *GetProvisionConfigResponseBody) SetTargetTrackingPolicies(v []*TargetTrackingPolicies) *GetProvisionConfigResponseBody {
	s.TargetTrackingPolicies = v
	return s
}

type GetProvisionConfigResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetProvisionConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProvisionConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProvisionConfigResponse) GoString() string {
	return s.String()
}

func (s *GetProvisionConfigResponse) SetHeaders(v map[string]*string) *GetProvisionConfigResponse {
	s.Headers = v
	return s
}

func (s *GetProvisionConfigResponse) SetBody(v *GetProvisionConfigResponseBody) *GetProvisionConfigResponse {
	s.Body = v
	return s
}

type GetResourceTagsHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s GetResourceTagsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetResourceTagsHeaders) GoString() string {
	return s.String()
}

func (s *GetResourceTagsHeaders) SetCommonHeaders(v map[string]*string) *GetResourceTagsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetResourceTagsHeaders) SetXFcAccountId(v string) *GetResourceTagsHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *GetResourceTagsHeaders) SetXFcCodeChecksum(v string) *GetResourceTagsHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *GetResourceTagsHeaders) SetXFcDate(v string) *GetResourceTagsHeaders {
	s.XFcDate = &v
	return s
}

func (s *GetResourceTagsHeaders) SetXFcInvocationType(v string) *GetResourceTagsHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *GetResourceTagsHeaders) SetXFcLogType(v string) *GetResourceTagsHeaders {
	s.XFcLogType = &v
	return s
}

func (s *GetResourceTagsHeaders) SetXFcTraceId(v string) *GetResourceTagsHeaders {
	s.XFcTraceId = &v
	return s
}

type GetResourceTagsRequest struct {
	// Resource ARN 全称或者简称
	ResourceArn *string `json:"resourceArn,omitempty" xml:"resourceArn,omitempty"`
}

func (s GetResourceTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceTagsRequest) GoString() string {
	return s.String()
}

func (s *GetResourceTagsRequest) SetResourceArn(v string) *GetResourceTagsRequest {
	s.ResourceArn = &v
	return s
}

type GetResourceTagsResponseBody struct {
	// Resource ARN 全称
	ResourceArn *string `json:"resourceArn,omitempty" xml:"resourceArn,omitempty"`
	// tag 列表
	Tags map[string]*string `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s GetResourceTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceTagsResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceTagsResponseBody) SetResourceArn(v string) *GetResourceTagsResponseBody {
	s.ResourceArn = &v
	return s
}

func (s *GetResourceTagsResponseBody) SetTags(v map[string]*string) *GetResourceTagsResponseBody {
	s.Tags = v
	return s
}

type GetResourceTagsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetResourceTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetResourceTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceTagsResponse) GoString() string {
	return s.String()
}

func (s *GetResourceTagsResponse) SetHeaders(v map[string]*string) *GetResourceTagsResponse {
	s.Headers = v
	return s
}

func (s *GetResourceTagsResponse) SetBody(v *GetResourceTagsResponseBody) *GetResourceTagsResponse {
	s.Body = v
	return s
}

type GetServiceHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s GetServiceHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetServiceHeaders) GoString() string {
	return s.String()
}

func (s *GetServiceHeaders) SetCommonHeaders(v map[string]*string) *GetServiceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetServiceHeaders) SetXFcAccountId(v string) *GetServiceHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *GetServiceHeaders) SetXFcCodeChecksum(v string) *GetServiceHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *GetServiceHeaders) SetXFcDate(v string) *GetServiceHeaders {
	s.XFcDate = &v
	return s
}

func (s *GetServiceHeaders) SetXFcInvocationType(v string) *GetServiceHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *GetServiceHeaders) SetXFcLogType(v string) *GetServiceHeaders {
	s.XFcLogType = &v
	return s
}

func (s *GetServiceHeaders) SetXFcTraceId(v string) *GetServiceHeaders {
	s.XFcTraceId = &v
	return s
}

type GetServiceRequest struct {
	// 限定符
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s GetServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceRequest) GoString() string {
	return s.String()
}

func (s *GetServiceRequest) SetQualifier(v string) *GetServiceRequest {
	s.Qualifier = &v
	return s
}

type GetServiceResponseBody struct {
	// 创建时间
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// 服务描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 公网访问设置
	InternetAccess *bool `json:"internetAccess,omitempty" xml:"internetAccess,omitempty"`
	// 上次更新时间
	LastModifiedTime *string    `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	LogConfig        *LogConfig `json:"logConfig,omitempty" xml:"logConfig,omitempty"`
	NasConfig        *NASConfig `json:"nasConfig,omitempty" xml:"nasConfig,omitempty"`
	// 服务角色
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// 服务ID
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// 服务名称
	ServiceName   *string        `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	TracingConfig *TracingConfig `json:"tracingConfig,omitempty" xml:"tracingConfig,omitempty"`
	VpcConfig     *VPCConfig     `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty"`
}

func (s GetServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBody) SetCreatedTime(v string) *GetServiceResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *GetServiceResponseBody) SetDescription(v string) *GetServiceResponseBody {
	s.Description = &v
	return s
}

func (s *GetServiceResponseBody) SetInternetAccess(v bool) *GetServiceResponseBody {
	s.InternetAccess = &v
	return s
}

func (s *GetServiceResponseBody) SetLastModifiedTime(v string) *GetServiceResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *GetServiceResponseBody) SetLogConfig(v *LogConfig) *GetServiceResponseBody {
	s.LogConfig = v
	return s
}

func (s *GetServiceResponseBody) SetNasConfig(v *NASConfig) *GetServiceResponseBody {
	s.NasConfig = v
	return s
}

func (s *GetServiceResponseBody) SetRole(v string) *GetServiceResponseBody {
	s.Role = &v
	return s
}

func (s *GetServiceResponseBody) SetServiceId(v string) *GetServiceResponseBody {
	s.ServiceId = &v
	return s
}

func (s *GetServiceResponseBody) SetServiceName(v string) *GetServiceResponseBody {
	s.ServiceName = &v
	return s
}

func (s *GetServiceResponseBody) SetTracingConfig(v *TracingConfig) *GetServiceResponseBody {
	s.TracingConfig = v
	return s
}

func (s *GetServiceResponseBody) SetVpcConfig(v *VPCConfig) *GetServiceResponseBody {
	s.VpcConfig = v
	return s
}

type GetServiceResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponse) GoString() string {
	return s.String()
}

func (s *GetServiceResponse) SetHeaders(v map[string]*string) *GetServiceResponse {
	s.Headers = v
	return s
}

func (s *GetServiceResponse) SetBody(v *GetServiceResponseBody) *GetServiceResponse {
	s.Body = v
	return s
}

type GetStatefulAsyncInvocationHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s GetStatefulAsyncInvocationHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetStatefulAsyncInvocationHeaders) GoString() string {
	return s.String()
}

func (s *GetStatefulAsyncInvocationHeaders) SetCommonHeaders(v map[string]*string) *GetStatefulAsyncInvocationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetStatefulAsyncInvocationHeaders) SetXFcAccountId(v string) *GetStatefulAsyncInvocationHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *GetStatefulAsyncInvocationHeaders) SetXFcCodeChecksum(v string) *GetStatefulAsyncInvocationHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *GetStatefulAsyncInvocationHeaders) SetXFcDate(v string) *GetStatefulAsyncInvocationHeaders {
	s.XFcDate = &v
	return s
}

func (s *GetStatefulAsyncInvocationHeaders) SetXFcInvocationType(v string) *GetStatefulAsyncInvocationHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *GetStatefulAsyncInvocationHeaders) SetXFcLogType(v string) *GetStatefulAsyncInvocationHeaders {
	s.XFcLogType = &v
	return s
}

func (s *GetStatefulAsyncInvocationHeaders) SetXFcTraceId(v string) *GetStatefulAsyncInvocationHeaders {
	s.XFcTraceId = &v
	return s
}

type GetStatefulAsyncInvocationRequest struct {
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s GetStatefulAsyncInvocationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStatefulAsyncInvocationRequest) GoString() string {
	return s.String()
}

func (s *GetStatefulAsyncInvocationRequest) SetQualifier(v string) *GetStatefulAsyncInvocationRequest {
	s.Qualifier = &v
	return s
}

type GetStatefulAsyncInvocationResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StatefulAsyncInvocation `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetStatefulAsyncInvocationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStatefulAsyncInvocationResponse) GoString() string {
	return s.String()
}

func (s *GetStatefulAsyncInvocationResponse) SetHeaders(v map[string]*string) *GetStatefulAsyncInvocationResponse {
	s.Headers = v
	return s
}

func (s *GetStatefulAsyncInvocationResponse) SetBody(v *StatefulAsyncInvocation) *GetStatefulAsyncInvocationResponse {
	s.Body = v
	return s
}

type GetTriggerHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s GetTriggerHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetTriggerHeaders) GoString() string {
	return s.String()
}

func (s *GetTriggerHeaders) SetCommonHeaders(v map[string]*string) *GetTriggerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTriggerHeaders) SetXFcAccountId(v string) *GetTriggerHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *GetTriggerHeaders) SetXFcCodeChecksum(v string) *GetTriggerHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *GetTriggerHeaders) SetXFcDate(v string) *GetTriggerHeaders {
	s.XFcDate = &v
	return s
}

func (s *GetTriggerHeaders) SetXFcInvocationType(v string) *GetTriggerHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *GetTriggerHeaders) SetXFcLogType(v string) *GetTriggerHeaders {
	s.XFcLogType = &v
	return s
}

func (s *GetTriggerHeaders) SetXFcTraceId(v string) *GetTriggerHeaders {
	s.XFcTraceId = &v
	return s
}

type GetTriggerResponseBody struct {
	// 创建时间
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 域名名称，使用域名名称拼接上函数计算域名，可以采用HTTP协议调用到触发器对应版本的函数。例如{domainName}.cn-shanghai.fc.aliyuncs.com
	DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
	// 调用函数使用的RAM角色的ARN
	InvocationRole *string `json:"invocationRole,omitempty" xml:"invocationRole,omitempty"`
	// 上次修改时间
	LastModifiedTime *string `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	// service版本
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	// event source的Aliyun Resource Name（ARN
	SourceArn *string `json:"sourceArn,omitempty" xml:"sourceArn,omitempty"`
	// trigger配置对象
	TriggerConfig *string `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty"`
	TriggerId     *string `json:"triggerId,omitempty" xml:"triggerId,omitempty"`
	// trigger名称
	TriggerName *string `json:"triggerName,omitempty" xml:"triggerName,omitempty"`
	// trigger类型，如 oss, log, tablestore, timer, http, cdn_events, mns_topic
	TriggerType *string `json:"triggerType,omitempty" xml:"triggerType,omitempty"`
}

func (s GetTriggerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *GetTriggerResponseBody) SetCreatedTime(v string) *GetTriggerResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *GetTriggerResponseBody) SetDescription(v string) *GetTriggerResponseBody {
	s.Description = &v
	return s
}

func (s *GetTriggerResponseBody) SetDomainName(v string) *GetTriggerResponseBody {
	s.DomainName = &v
	return s
}

func (s *GetTriggerResponseBody) SetInvocationRole(v string) *GetTriggerResponseBody {
	s.InvocationRole = &v
	return s
}

func (s *GetTriggerResponseBody) SetLastModifiedTime(v string) *GetTriggerResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *GetTriggerResponseBody) SetQualifier(v string) *GetTriggerResponseBody {
	s.Qualifier = &v
	return s
}

func (s *GetTriggerResponseBody) SetSourceArn(v string) *GetTriggerResponseBody {
	s.SourceArn = &v
	return s
}

func (s *GetTriggerResponseBody) SetTriggerConfig(v string) *GetTriggerResponseBody {
	s.TriggerConfig = &v
	return s
}

func (s *GetTriggerResponseBody) SetTriggerId(v string) *GetTriggerResponseBody {
	s.TriggerId = &v
	return s
}

func (s *GetTriggerResponseBody) SetTriggerName(v string) *GetTriggerResponseBody {
	s.TriggerName = &v
	return s
}

func (s *GetTriggerResponseBody) SetTriggerType(v string) *GetTriggerResponseBody {
	s.TriggerType = &v
	return s
}

type GetTriggerResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTriggerResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTriggerResponse) GoString() string {
	return s.String()
}

func (s *GetTriggerResponse) SetHeaders(v map[string]*string) *GetTriggerResponse {
	s.Headers = v
	return s
}

func (s *GetTriggerResponse) SetBody(v *GetTriggerResponseBody) *GetTriggerResponse {
	s.Body = v
	return s
}

type InvokeFunctionHeaders struct {
	CommonHeaders   map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId    *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate         *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	// 调用方式:Sync或者Async，默认值：Sync
	XFcInvocationType *string `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	// 请求返回日志类型, Tail 为返回函数日志最后 4KB 数据，None 或空值则返回不带有日志，默认为 None
	XFcLogType                   *string `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcStatefulAsyncInvocationId *string `json:"X-Fc-Stateful-Async-Invocation-Id,omitempty" xml:"X-Fc-Stateful-Async-Invocation-Id,omitempty"`
	XFcTraceId                   *string `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s InvokeFunctionHeaders) String() string {
	return tea.Prettify(s)
}

func (s InvokeFunctionHeaders) GoString() string {
	return s.String()
}

func (s *InvokeFunctionHeaders) SetCommonHeaders(v map[string]*string) *InvokeFunctionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InvokeFunctionHeaders) SetXFcAccountId(v string) *InvokeFunctionHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *InvokeFunctionHeaders) SetXFcCodeChecksum(v string) *InvokeFunctionHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *InvokeFunctionHeaders) SetXFcDate(v string) *InvokeFunctionHeaders {
	s.XFcDate = &v
	return s
}

func (s *InvokeFunctionHeaders) SetXFcInvocationType(v string) *InvokeFunctionHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *InvokeFunctionHeaders) SetXFcLogType(v string) *InvokeFunctionHeaders {
	s.XFcLogType = &v
	return s
}

func (s *InvokeFunctionHeaders) SetXFcStatefulAsyncInvocationId(v string) *InvokeFunctionHeaders {
	s.XFcStatefulAsyncInvocationId = &v
	return s
}

func (s *InvokeFunctionHeaders) SetXFcTraceId(v string) *InvokeFunctionHeaders {
	s.XFcTraceId = &v
	return s
}

type InvokeFunctionRequest struct {
	// 事件（event），binary type。函数计算服务将event传递给用户function来处理
	Body []byte `json:"body,omitempty" xml:"body,omitempty"`
	// service版本, 可以是versionId或者aliasName
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s InvokeFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s InvokeFunctionRequest) GoString() string {
	return s.String()
}

func (s *InvokeFunctionRequest) SetBody(v []byte) *InvokeFunctionRequest {
	s.Body = v
	return s
}

func (s *InvokeFunctionRequest) SetQualifier(v string) *InvokeFunctionRequest {
	s.Qualifier = &v
	return s
}

type InvokeFunctionResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    []byte             `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InvokeFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s InvokeFunctionResponse) GoString() string {
	return s.String()
}

func (s *InvokeFunctionResponse) SetHeaders(v map[string]*string) *InvokeFunctionResponse {
	s.Headers = v
	return s
}

func (s *InvokeFunctionResponse) SetBody(v []byte) *InvokeFunctionResponse {
	s.Body = v
	return s
}

type ListAliasesHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s ListAliasesHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListAliasesHeaders) GoString() string {
	return s.String()
}

func (s *ListAliasesHeaders) SetCommonHeaders(v map[string]*string) *ListAliasesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListAliasesHeaders) SetXFcAccountId(v string) *ListAliasesHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *ListAliasesHeaders) SetXFcCodeChecksum(v string) *ListAliasesHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *ListAliasesHeaders) SetXFcDate(v string) *ListAliasesHeaders {
	s.XFcDate = &v
	return s
}

func (s *ListAliasesHeaders) SetXFcInvocationType(v string) *ListAliasesHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *ListAliasesHeaders) SetXFcLogType(v string) *ListAliasesHeaders {
	s.XFcLogType = &v
	return s
}

func (s *ListAliasesHeaders) SetXFcTraceId(v string) *ListAliasesHeaders {
	s.XFcTraceId = &v
	return s
}

type ListAliasesRequest struct {
	// 最多返回个数
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// 下次查询token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 前缀
	Prefix *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
	// 起始key
	StartKey *string `json:"startKey,omitempty" xml:"startKey,omitempty"`
}

func (s ListAliasesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAliasesRequest) GoString() string {
	return s.String()
}

func (s *ListAliasesRequest) SetLimit(v int32) *ListAliasesRequest {
	s.Limit = &v
	return s
}

func (s *ListAliasesRequest) SetNextToken(v string) *ListAliasesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAliasesRequest) SetPrefix(v string) *ListAliasesRequest {
	s.Prefix = &v
	return s
}

func (s *ListAliasesRequest) SetStartKey(v string) *ListAliasesRequest {
	s.StartKey = &v
	return s
}

type ListAliasesResponseBody struct {
	// 别名列表
	Aliases []*ListAliasesResponseBodyAliases `json:"aliases,omitempty" xml:"aliases,omitempty" type:"Repeated"`
	// 下次查询token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListAliasesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAliasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAliasesResponseBody) SetAliases(v []*ListAliasesResponseBodyAliases) *ListAliasesResponseBody {
	s.Aliases = v
	return s
}

func (s *ListAliasesResponseBody) SetNextToken(v string) *ListAliasesResponseBody {
	s.NextToken = &v
	return s
}

type ListAliasesResponseBodyAliases struct {
	// 额外版本权重
	AdditionalVersionWeight map[string]*float32 `json:"additionalVersionWeight,omitempty" xml:"additionalVersionWeight,omitempty"`
	// 别名名称
	AliasName *string `json:"aliasName,omitempty" xml:"aliasName,omitempty"`
	// 创建时间
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// 别名描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 上次更新时间
	LastModifiedTime *string `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	// 版本ID
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s ListAliasesResponseBodyAliases) String() string {
	return tea.Prettify(s)
}

func (s ListAliasesResponseBodyAliases) GoString() string {
	return s.String()
}

func (s *ListAliasesResponseBodyAliases) SetAdditionalVersionWeight(v map[string]*float32) *ListAliasesResponseBodyAliases {
	s.AdditionalVersionWeight = v
	return s
}

func (s *ListAliasesResponseBodyAliases) SetAliasName(v string) *ListAliasesResponseBodyAliases {
	s.AliasName = &v
	return s
}

func (s *ListAliasesResponseBodyAliases) SetCreatedTime(v string) *ListAliasesResponseBodyAliases {
	s.CreatedTime = &v
	return s
}

func (s *ListAliasesResponseBodyAliases) SetDescription(v string) *ListAliasesResponseBodyAliases {
	s.Description = &v
	return s
}

func (s *ListAliasesResponseBodyAliases) SetLastModifiedTime(v string) *ListAliasesResponseBodyAliases {
	s.LastModifiedTime = &v
	return s
}

func (s *ListAliasesResponseBodyAliases) SetVersionId(v string) *ListAliasesResponseBodyAliases {
	s.VersionId = &v
	return s
}

type ListAliasesResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAliasesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAliasesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAliasesResponse) GoString() string {
	return s.String()
}

func (s *ListAliasesResponse) SetHeaders(v map[string]*string) *ListAliasesResponse {
	s.Headers = v
	return s
}

func (s *ListAliasesResponse) SetBody(v *ListAliasesResponseBody) *ListAliasesResponse {
	s.Body = v
	return s
}

type ListCustomDomainsHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s ListCustomDomainsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListCustomDomainsHeaders) GoString() string {
	return s.String()
}

func (s *ListCustomDomainsHeaders) SetCommonHeaders(v map[string]*string) *ListCustomDomainsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListCustomDomainsHeaders) SetXFcAccountId(v string) *ListCustomDomainsHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *ListCustomDomainsHeaders) SetXFcCodeChecksum(v string) *ListCustomDomainsHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *ListCustomDomainsHeaders) SetXFcDate(v string) *ListCustomDomainsHeaders {
	s.XFcDate = &v
	return s
}

func (s *ListCustomDomainsHeaders) SetXFcInvocationType(v string) *ListCustomDomainsHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *ListCustomDomainsHeaders) SetXFcLogType(v string) *ListCustomDomainsHeaders {
	s.XFcLogType = &v
	return s
}

func (s *ListCustomDomainsHeaders) SetXFcTraceId(v string) *ListCustomDomainsHeaders {
	s.XFcTraceId = &v
	return s
}

type ListCustomDomainsRequest struct {
	Limit     *int32  `json:"limit,omitempty" xml:"limit,omitempty"`
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Prefix    *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
	StartKey  *string `json:"startKey,omitempty" xml:"startKey,omitempty"`
}

func (s ListCustomDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCustomDomainsRequest) GoString() string {
	return s.String()
}

func (s *ListCustomDomainsRequest) SetLimit(v int32) *ListCustomDomainsRequest {
	s.Limit = &v
	return s
}

func (s *ListCustomDomainsRequest) SetNextToken(v string) *ListCustomDomainsRequest {
	s.NextToken = &v
	return s
}

func (s *ListCustomDomainsRequest) SetPrefix(v string) *ListCustomDomainsRequest {
	s.Prefix = &v
	return s
}

func (s *ListCustomDomainsRequest) SetStartKey(v string) *ListCustomDomainsRequest {
	s.StartKey = &v
	return s
}

type ListCustomDomainsResponseBody struct {
	CustomDomains []*ListCustomDomainsResponseBodyCustomDomains `json:"customDomains,omitempty" xml:"customDomains,omitempty" type:"Repeated"`
	NextToken     *string                                       `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListCustomDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCustomDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomDomainsResponseBody) SetCustomDomains(v []*ListCustomDomainsResponseBodyCustomDomains) *ListCustomDomainsResponseBody {
	s.CustomDomains = v
	return s
}

func (s *ListCustomDomainsResponseBody) SetNextToken(v string) *ListCustomDomainsResponseBody {
	s.NextToken = &v
	return s
}

type ListCustomDomainsResponseBodyCustomDomains struct {
	AccountId        *string      `json:"accountId,omitempty" xml:"accountId,omitempty"`
	ApiVersion       *string      `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
	CertConfig       *CertConfig  `json:"certConfig,omitempty" xml:"certConfig,omitempty"`
	CreatedTime      *string      `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	DomainName       *string      `json:"domainName,omitempty" xml:"domainName,omitempty"`
	LastModifiedTime *string      `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	Protocol         *string      `json:"protocol,omitempty" xml:"protocol,omitempty"`
	RouteConfig      *RouteConfig `json:"routeConfig,omitempty" xml:"routeConfig,omitempty"`
}

func (s ListCustomDomainsResponseBodyCustomDomains) String() string {
	return tea.Prettify(s)
}

func (s ListCustomDomainsResponseBodyCustomDomains) GoString() string {
	return s.String()
}

func (s *ListCustomDomainsResponseBodyCustomDomains) SetAccountId(v string) *ListCustomDomainsResponseBodyCustomDomains {
	s.AccountId = &v
	return s
}

func (s *ListCustomDomainsResponseBodyCustomDomains) SetApiVersion(v string) *ListCustomDomainsResponseBodyCustomDomains {
	s.ApiVersion = &v
	return s
}

func (s *ListCustomDomainsResponseBodyCustomDomains) SetCertConfig(v *CertConfig) *ListCustomDomainsResponseBodyCustomDomains {
	s.CertConfig = v
	return s
}

func (s *ListCustomDomainsResponseBodyCustomDomains) SetCreatedTime(v string) *ListCustomDomainsResponseBodyCustomDomains {
	s.CreatedTime = &v
	return s
}

func (s *ListCustomDomainsResponseBodyCustomDomains) SetDomainName(v string) *ListCustomDomainsResponseBodyCustomDomains {
	s.DomainName = &v
	return s
}

func (s *ListCustomDomainsResponseBodyCustomDomains) SetLastModifiedTime(v string) *ListCustomDomainsResponseBodyCustomDomains {
	s.LastModifiedTime = &v
	return s
}

func (s *ListCustomDomainsResponseBodyCustomDomains) SetProtocol(v string) *ListCustomDomainsResponseBodyCustomDomains {
	s.Protocol = &v
	return s
}

func (s *ListCustomDomainsResponseBodyCustomDomains) SetRouteConfig(v *RouteConfig) *ListCustomDomainsResponseBodyCustomDomains {
	s.RouteConfig = v
	return s
}

type ListCustomDomainsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCustomDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCustomDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCustomDomainsResponse) GoString() string {
	return s.String()
}

func (s *ListCustomDomainsResponse) SetHeaders(v map[string]*string) *ListCustomDomainsResponse {
	s.Headers = v
	return s
}

func (s *ListCustomDomainsResponse) SetBody(v *ListCustomDomainsResponseBody) *ListCustomDomainsResponse {
	s.Body = v
	return s
}

type ListEventSourcesHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s ListEventSourcesHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListEventSourcesHeaders) GoString() string {
	return s.String()
}

func (s *ListEventSourcesHeaders) SetCommonHeaders(v map[string]*string) *ListEventSourcesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListEventSourcesHeaders) SetXFcAccountId(v string) *ListEventSourcesHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *ListEventSourcesHeaders) SetXFcCodeChecksum(v string) *ListEventSourcesHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *ListEventSourcesHeaders) SetXFcDate(v string) *ListEventSourcesHeaders {
	s.XFcDate = &v
	return s
}

func (s *ListEventSourcesHeaders) SetXFcInvocationType(v string) *ListEventSourcesHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *ListEventSourcesHeaders) SetXFcLogType(v string) *ListEventSourcesHeaders {
	s.XFcLogType = &v
	return s
}

func (s *ListEventSourcesHeaders) SetXFcTraceId(v string) *ListEventSourcesHeaders {
	s.XFcTraceId = &v
	return s
}

type ListEventSourcesRequest struct {
	// 别名或版本
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s ListEventSourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEventSourcesRequest) GoString() string {
	return s.String()
}

func (s *ListEventSourcesRequest) SetQualifier(v string) *ListEventSourcesRequest {
	s.Qualifier = &v
	return s
}

type ListEventSourcesResponseBody struct {
	// 事件源列表
	EventSources []*ListEventSourcesResponseBodyEventSources `json:"eventSources,omitempty" xml:"eventSources,omitempty" type:"Repeated"`
}

func (s ListEventSourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEventSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListEventSourcesResponseBody) SetEventSources(v []*ListEventSourcesResponseBodyEventSources) *ListEventSourcesResponseBody {
	s.EventSources = v
	return s
}

type ListEventSourcesResponseBodyEventSources struct {
	// 创建时间
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// 事件源资源标识符
	SourceArn *string `json:"sourceArn,omitempty" xml:"sourceArn,omitempty"`
}

func (s ListEventSourcesResponseBodyEventSources) String() string {
	return tea.Prettify(s)
}

func (s ListEventSourcesResponseBodyEventSources) GoString() string {
	return s.String()
}

func (s *ListEventSourcesResponseBodyEventSources) SetCreatedTime(v string) *ListEventSourcesResponseBodyEventSources {
	s.CreatedTime = &v
	return s
}

func (s *ListEventSourcesResponseBodyEventSources) SetSourceArn(v string) *ListEventSourcesResponseBodyEventSources {
	s.SourceArn = &v
	return s
}

type ListEventSourcesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListEventSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEventSourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEventSourcesResponse) GoString() string {
	return s.String()
}

func (s *ListEventSourcesResponse) SetHeaders(v map[string]*string) *ListEventSourcesResponse {
	s.Headers = v
	return s
}

func (s *ListEventSourcesResponse) SetBody(v *ListEventSourcesResponseBody) *ListEventSourcesResponse {
	s.Body = v
	return s
}

type ListFunctionAsyncInvokeConfigsHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s ListFunctionAsyncInvokeConfigsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionAsyncInvokeConfigsHeaders) GoString() string {
	return s.String()
}

func (s *ListFunctionAsyncInvokeConfigsHeaders) SetCommonHeaders(v map[string]*string) *ListFunctionAsyncInvokeConfigsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListFunctionAsyncInvokeConfigsHeaders) SetXFcAccountId(v string) *ListFunctionAsyncInvokeConfigsHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *ListFunctionAsyncInvokeConfigsHeaders) SetXFcCodeChecksum(v string) *ListFunctionAsyncInvokeConfigsHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *ListFunctionAsyncInvokeConfigsHeaders) SetXFcDate(v string) *ListFunctionAsyncInvokeConfigsHeaders {
	s.XFcDate = &v
	return s
}

func (s *ListFunctionAsyncInvokeConfigsHeaders) SetXFcInvocationType(v string) *ListFunctionAsyncInvokeConfigsHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *ListFunctionAsyncInvokeConfigsHeaders) SetXFcLogType(v string) *ListFunctionAsyncInvokeConfigsHeaders {
	s.XFcLogType = &v
	return s
}

func (s *ListFunctionAsyncInvokeConfigsHeaders) SetXFcTraceId(v string) *ListFunctionAsyncInvokeConfigsHeaders {
	s.XFcTraceId = &v
	return s
}

type ListFunctionAsyncInvokeConfigsRequest struct {
	// 最多返回个数
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// 下次查询token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListFunctionAsyncInvokeConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionAsyncInvokeConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListFunctionAsyncInvokeConfigsRequest) SetLimit(v int32) *ListFunctionAsyncInvokeConfigsRequest {
	s.Limit = &v
	return s
}

func (s *ListFunctionAsyncInvokeConfigsRequest) SetNextToken(v string) *ListFunctionAsyncInvokeConfigsRequest {
	s.NextToken = &v
	return s
}

type ListFunctionAsyncInvokeConfigsResponseBody struct {
	// 异步配置列表
	Configs []*ListFunctionAsyncInvokeConfigsResponseBodyConfigs `json:"configs,omitempty" xml:"configs,omitempty" type:"Repeated"`
	// 下次查询token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListFunctionAsyncInvokeConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionAsyncInvokeConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFunctionAsyncInvokeConfigsResponseBody) SetConfigs(v []*ListFunctionAsyncInvokeConfigsResponseBodyConfigs) *ListFunctionAsyncInvokeConfigsResponseBody {
	s.Configs = v
	return s
}

func (s *ListFunctionAsyncInvokeConfigsResponseBody) SetNextToken(v string) *ListFunctionAsyncInvokeConfigsResponseBody {
	s.NextToken = &v
	return s
}

type ListFunctionAsyncInvokeConfigsResponseBodyConfigs struct {
	// 创建时间
	CreatedTime       *string            `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	DestinationConfig *DestinationConfig `json:"destinationConfig,omitempty" xml:"destinationConfig,omitempty"`
	// 函数名称
	Function *string `json:"function,omitempty" xml:"function,omitempty"`
	// 最后更改时间
	LastModifiedTime *string `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	// 消息最大存活时长
	MaxAsyncEventAgeInSeconds *int64 `json:"maxAsyncEventAgeInSeconds,omitempty" xml:"maxAsyncEventAgeInSeconds,omitempty"`
	// 异步调用失败后的最大重试次数
	MaxAsyncRetryAttempts *int64 `json:"maxAsyncRetryAttempts,omitempty" xml:"maxAsyncRetryAttempts,omitempty"`
	// 限定符
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	// 服务名称
	Service            *string `json:"service,omitempty" xml:"service,omitempty"`
	StatefulInvocation *bool   `json:"statefulInvocation,omitempty" xml:"statefulInvocation,omitempty"`
}

func (s ListFunctionAsyncInvokeConfigsResponseBodyConfigs) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionAsyncInvokeConfigsResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *ListFunctionAsyncInvokeConfigsResponseBodyConfigs) SetCreatedTime(v string) *ListFunctionAsyncInvokeConfigsResponseBodyConfigs {
	s.CreatedTime = &v
	return s
}

func (s *ListFunctionAsyncInvokeConfigsResponseBodyConfigs) SetDestinationConfig(v *DestinationConfig) *ListFunctionAsyncInvokeConfigsResponseBodyConfigs {
	s.DestinationConfig = v
	return s
}

func (s *ListFunctionAsyncInvokeConfigsResponseBodyConfigs) SetFunction(v string) *ListFunctionAsyncInvokeConfigsResponseBodyConfigs {
	s.Function = &v
	return s
}

func (s *ListFunctionAsyncInvokeConfigsResponseBodyConfigs) SetLastModifiedTime(v string) *ListFunctionAsyncInvokeConfigsResponseBodyConfigs {
	s.LastModifiedTime = &v
	return s
}

func (s *ListFunctionAsyncInvokeConfigsResponseBodyConfigs) SetMaxAsyncEventAgeInSeconds(v int64) *ListFunctionAsyncInvokeConfigsResponseBodyConfigs {
	s.MaxAsyncEventAgeInSeconds = &v
	return s
}

func (s *ListFunctionAsyncInvokeConfigsResponseBodyConfigs) SetMaxAsyncRetryAttempts(v int64) *ListFunctionAsyncInvokeConfigsResponseBodyConfigs {
	s.MaxAsyncRetryAttempts = &v
	return s
}

func (s *ListFunctionAsyncInvokeConfigsResponseBodyConfigs) SetQualifier(v string) *ListFunctionAsyncInvokeConfigsResponseBodyConfigs {
	s.Qualifier = &v
	return s
}

func (s *ListFunctionAsyncInvokeConfigsResponseBodyConfigs) SetService(v string) *ListFunctionAsyncInvokeConfigsResponseBodyConfigs {
	s.Service = &v
	return s
}

func (s *ListFunctionAsyncInvokeConfigsResponseBodyConfigs) SetStatefulInvocation(v bool) *ListFunctionAsyncInvokeConfigsResponseBodyConfigs {
	s.StatefulInvocation = &v
	return s
}

type ListFunctionAsyncInvokeConfigsResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFunctionAsyncInvokeConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFunctionAsyncInvokeConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionAsyncInvokeConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListFunctionAsyncInvokeConfigsResponse) SetHeaders(v map[string]*string) *ListFunctionAsyncInvokeConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListFunctionAsyncInvokeConfigsResponse) SetBody(v *ListFunctionAsyncInvokeConfigsResponseBody) *ListFunctionAsyncInvokeConfigsResponse {
	s.Body = v
	return s
}

type ListFunctionsHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s ListFunctionsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsHeaders) GoString() string {
	return s.String()
}

func (s *ListFunctionsHeaders) SetCommonHeaders(v map[string]*string) *ListFunctionsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListFunctionsHeaders) SetXFcAccountId(v string) *ListFunctionsHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *ListFunctionsHeaders) SetXFcCodeChecksum(v string) *ListFunctionsHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *ListFunctionsHeaders) SetXFcDate(v string) *ListFunctionsHeaders {
	s.XFcDate = &v
	return s
}

func (s *ListFunctionsHeaders) SetXFcInvocationType(v string) *ListFunctionsHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *ListFunctionsHeaders) SetXFcLogType(v string) *ListFunctionsHeaders {
	s.XFcLogType = &v
	return s
}

func (s *ListFunctionsHeaders) SetXFcTraceId(v string) *ListFunctionsHeaders {
	s.XFcTraceId = &v
	return s
}

type ListFunctionsRequest struct {
	// 限定此次返回资源的数量。如果不设定，默认返回20，最大不能超过100。返回结果可能小于指定的数量，但不会多于指定的数量
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// 用来返回更多结果。第一次查询不需要提供这个参数，后续查询的token从返回结果中获取
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 限定返回的资源名称必须以prefix作为前缀
	Prefix *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
	// service版本, 可以是versionId或者aliasName
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	// 设定结果从startKey之后（包括startKey）按字母排序的第一个开始返回
	StartKey *string `json:"startKey,omitempty" xml:"startKey,omitempty"`
}

func (s ListFunctionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsRequest) GoString() string {
	return s.String()
}

func (s *ListFunctionsRequest) SetLimit(v int32) *ListFunctionsRequest {
	s.Limit = &v
	return s
}

func (s *ListFunctionsRequest) SetNextToken(v string) *ListFunctionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListFunctionsRequest) SetPrefix(v string) *ListFunctionsRequest {
	s.Prefix = &v
	return s
}

func (s *ListFunctionsRequest) SetQualifier(v string) *ListFunctionsRequest {
	s.Qualifier = &v
	return s
}

func (s *ListFunctionsRequest) SetStartKey(v string) *ListFunctionsRequest {
	s.StartKey = &v
	return s
}

type ListFunctionsResponseBody struct {
	// 函数列表
	Functions []*ListFunctionsResponseBodyFunctions `json:"functions,omitempty" xml:"functions,omitempty" type:"Repeated"`
	// 用来返回更多的查询结果。如果这个值没有返回，则说明没有更多结果
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListFunctionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBody) SetFunctions(v []*ListFunctionsResponseBodyFunctions) *ListFunctionsResponseBody {
	s.Functions = v
	return s
}

func (s *ListFunctionsResponseBody) SetNextToken(v string) *ListFunctionsResponseBody {
	s.NextToken = &v
	return s
}

type ListFunctionsResponseBodyFunctions struct {
	// 自定义、自定义容器运行时 HTTP Server 的监听端口
	CaPort *int32 `json:"caPort,omitempty" xml:"caPort,omitempty"`
	// function code包的CRC64值
	CodeChecksum *string `json:"codeChecksum,omitempty" xml:"codeChecksum,omitempty"`
	// 系统返回的function的code包大小，单位为byte Example : 1024
	CodeSize *int64 `json:"codeSize,omitempty" xml:"codeSize,omitempty"`
	// function创建时间
	CreatedTime           *string                `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	CustomContainerConfig *CustomContainerConfig `json:"customContainerConfig,omitempty" xml:"customContainerConfig,omitempty"`
	// 函数描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 为函数设置的环境变量，可以在函数中获取环境变量的值
	EnvironmentVariables map[string]*string `json:"environmentVariables,omitempty" xml:"environmentVariables,omitempty"`
	// 系统为每个function生成的唯一ID
	FunctionId *string `json:"functionId,omitempty" xml:"functionId,omitempty"`
	// 函数名称
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	// function的执行入口
	Handler *string `json:"handler,omitempty" xml:"handler,omitempty"`
	// 初始化function运行的超时时间，单位为秒，最小1秒，默认3秒。初始化function超过这个时间后会被终止执行
	InitializationTimeout *int32 `json:"initializationTimeout,omitempty" xml:"initializationTimeout,omitempty"`
	// 初始化 function 执行的入口，具体格式和语言相关
	Initializer             *string                  `json:"initializer,omitempty" xml:"initializer,omitempty"`
	InstanceConcurrency     *int32                   `json:"instanceConcurrency,omitempty" xml:"instanceConcurrency,omitempty"`
	InstanceLifecycleConfig *InstanceLifecycleConfig `json:"instanceLifecycleConfig,omitempty" xml:"instanceLifecycleConfig,omitempty"`
	InstanceType            *string                  `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	// function上次修改时间
	LastModifiedTime *string   `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	Layers           []*string `json:"layers,omitempty" xml:"layers,omitempty" type:"Repeated"`
	// function设置的内存大小，单位为MB
	MemorySize *int32 `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	// function运行的语言环境，目前支持nodejs6, nodejs8, python2.7, python3, java8
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// 运行的超时时间，单位为秒
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
}

func (s ListFunctionsResponseBodyFunctions) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsResponseBodyFunctions) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBodyFunctions) SetCaPort(v int32) *ListFunctionsResponseBodyFunctions {
	s.CaPort = &v
	return s
}

func (s *ListFunctionsResponseBodyFunctions) SetCodeChecksum(v string) *ListFunctionsResponseBodyFunctions {
	s.CodeChecksum = &v
	return s
}

func (s *ListFunctionsResponseBodyFunctions) SetCodeSize(v int64) *ListFunctionsResponseBodyFunctions {
	s.CodeSize = &v
	return s
}

func (s *ListFunctionsResponseBodyFunctions) SetCreatedTime(v string) *ListFunctionsResponseBodyFunctions {
	s.CreatedTime = &v
	return s
}

func (s *ListFunctionsResponseBodyFunctions) SetCustomContainerConfig(v *CustomContainerConfig) *ListFunctionsResponseBodyFunctions {
	s.CustomContainerConfig = v
	return s
}

func (s *ListFunctionsResponseBodyFunctions) SetDescription(v string) *ListFunctionsResponseBodyFunctions {
	s.Description = &v
	return s
}

func (s *ListFunctionsResponseBodyFunctions) SetEnvironmentVariables(v map[string]*string) *ListFunctionsResponseBodyFunctions {
	s.EnvironmentVariables = v
	return s
}

func (s *ListFunctionsResponseBodyFunctions) SetFunctionId(v string) *ListFunctionsResponseBodyFunctions {
	s.FunctionId = &v
	return s
}

func (s *ListFunctionsResponseBodyFunctions) SetFunctionName(v string) *ListFunctionsResponseBodyFunctions {
	s.FunctionName = &v
	return s
}

func (s *ListFunctionsResponseBodyFunctions) SetHandler(v string) *ListFunctionsResponseBodyFunctions {
	s.Handler = &v
	return s
}

func (s *ListFunctionsResponseBodyFunctions) SetInitializationTimeout(v int32) *ListFunctionsResponseBodyFunctions {
	s.InitializationTimeout = &v
	return s
}

func (s *ListFunctionsResponseBodyFunctions) SetInitializer(v string) *ListFunctionsResponseBodyFunctions {
	s.Initializer = &v
	return s
}

func (s *ListFunctionsResponseBodyFunctions) SetInstanceConcurrency(v int32) *ListFunctionsResponseBodyFunctions {
	s.InstanceConcurrency = &v
	return s
}

func (s *ListFunctionsResponseBodyFunctions) SetInstanceLifecycleConfig(v *InstanceLifecycleConfig) *ListFunctionsResponseBodyFunctions {
	s.InstanceLifecycleConfig = v
	return s
}

func (s *ListFunctionsResponseBodyFunctions) SetInstanceType(v string) *ListFunctionsResponseBodyFunctions {
	s.InstanceType = &v
	return s
}

func (s *ListFunctionsResponseBodyFunctions) SetLastModifiedTime(v string) *ListFunctionsResponseBodyFunctions {
	s.LastModifiedTime = &v
	return s
}

func (s *ListFunctionsResponseBodyFunctions) SetLayers(v []*string) *ListFunctionsResponseBodyFunctions {
	s.Layers = v
	return s
}

func (s *ListFunctionsResponseBodyFunctions) SetMemorySize(v int32) *ListFunctionsResponseBodyFunctions {
	s.MemorySize = &v
	return s
}

func (s *ListFunctionsResponseBodyFunctions) SetRuntime(v string) *ListFunctionsResponseBodyFunctions {
	s.Runtime = &v
	return s
}

func (s *ListFunctionsResponseBodyFunctions) SetTimeout(v int32) *ListFunctionsResponseBodyFunctions {
	s.Timeout = &v
	return s
}

type ListFunctionsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFunctionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFunctionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsResponse) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponse) SetHeaders(v map[string]*string) *ListFunctionsResponse {
	s.Headers = v
	return s
}

func (s *ListFunctionsResponse) SetBody(v *ListFunctionsResponseBody) *ListFunctionsResponse {
	s.Body = v
	return s
}

type ListInstancesHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId  *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
}

func (s ListInstancesHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesHeaders) GoString() string {
	return s.String()
}

func (s *ListInstancesHeaders) SetCommonHeaders(v map[string]*string) *ListInstancesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListInstancesHeaders) SetXFcAccountId(v string) *ListInstancesHeaders {
	s.XFcAccountId = &v
	return s
}

type ListInstancesRequest struct {
	// 限定此次返回资源的数量。如果不设定，默认返回20，最大不能超过100。返回结果可能小于指定的数量，但不会多于指定的数量
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// 用来返回更多结果。第一次查询不需要提供这个参数，后续查询的token从返回结果中获取
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// service版本, 可以是versionId或者aliasName
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) SetLimit(v int32) *ListInstancesRequest {
	s.Limit = &v
	return s
}

func (s *ListInstancesRequest) SetNextToken(v string) *ListInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *ListInstancesRequest) SetQualifier(v string) *ListInstancesRequest {
	s.Qualifier = &v
	return s
}

type ListInstancesResponseBody struct {
	Instances []*ListInstancesResponseBodyInstances `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
}

func (s ListInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) SetInstances(v []*ListInstancesResponseBodyInstances) *ListInstancesResponseBody {
	s.Instances = v
	return s
}

type ListInstancesResponseBodyInstances struct {
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	VersionId  *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s ListInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstances) SetInstanceId(v string) *ListInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetVersionId(v string) *ListInstancesResponseBodyInstances {
	s.VersionId = &v
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

type ListLayerVersionsHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s ListLayerVersionsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListLayerVersionsHeaders) GoString() string {
	return s.String()
}

func (s *ListLayerVersionsHeaders) SetCommonHeaders(v map[string]*string) *ListLayerVersionsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListLayerVersionsHeaders) SetXFcAccountId(v string) *ListLayerVersionsHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *ListLayerVersionsHeaders) SetXFcCodeChecksum(v string) *ListLayerVersionsHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *ListLayerVersionsHeaders) SetXFcDate(v string) *ListLayerVersionsHeaders {
	s.XFcDate = &v
	return s
}

func (s *ListLayerVersionsHeaders) SetXFcInvocationType(v string) *ListLayerVersionsHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *ListLayerVersionsHeaders) SetXFcLogType(v string) *ListLayerVersionsHeaders {
	s.XFcLogType = &v
	return s
}

func (s *ListLayerVersionsHeaders) SetXFcTraceId(v string) *ListLayerVersionsHeaders {
	s.XFcTraceId = &v
	return s
}

type ListLayerVersionsRequest struct {
	// 本次读取的最大数据记录数量
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// 起始版本
	StartVersion *int32 `json:"startVersion,omitempty" xml:"startVersion,omitempty"`
}

func (s ListLayerVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLayerVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListLayerVersionsRequest) SetLimit(v int32) *ListLayerVersionsRequest {
	s.Limit = &v
	return s
}

func (s *ListLayerVersionsRequest) SetStartVersion(v int32) *ListLayerVersionsRequest {
	s.StartVersion = &v
	return s
}

type ListLayerVersionsResponseBody struct {
	// 层版本列表
	Layers []*Layer `json:"layers,omitempty" xml:"layers,omitempty" type:"Repeated"`
	// 剩余列表起始版本号
	NextVersion *int32 `json:"nextVersion,omitempty" xml:"nextVersion,omitempty"`
}

func (s ListLayerVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLayerVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLayerVersionsResponseBody) SetLayers(v []*Layer) *ListLayerVersionsResponseBody {
	s.Layers = v
	return s
}

func (s *ListLayerVersionsResponseBody) SetNextVersion(v int32) *ListLayerVersionsResponseBody {
	s.NextVersion = &v
	return s
}

type ListLayerVersionsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListLayerVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListLayerVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLayerVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListLayerVersionsResponse) SetHeaders(v map[string]*string) *ListLayerVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListLayerVersionsResponse) SetBody(v *ListLayerVersionsResponseBody) *ListLayerVersionsResponse {
	s.Body = v
	return s
}

type ListLayersHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s ListLayersHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListLayersHeaders) GoString() string {
	return s.String()
}

func (s *ListLayersHeaders) SetCommonHeaders(v map[string]*string) *ListLayersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListLayersHeaders) SetXFcAccountId(v string) *ListLayersHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *ListLayersHeaders) SetXFcCodeChecksum(v string) *ListLayersHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *ListLayersHeaders) SetXFcDate(v string) *ListLayersHeaders {
	s.XFcDate = &v
	return s
}

func (s *ListLayersHeaders) SetXFcInvocationType(v string) *ListLayersHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *ListLayersHeaders) SetXFcLogType(v string) *ListLayersHeaders {
	s.XFcLogType = &v
	return s
}

func (s *ListLayersHeaders) SetXFcTraceId(v string) *ListLayersHeaders {
	s.XFcTraceId = &v
	return s
}

type ListLayersRequest struct {
	// 最大返回条目数
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// 下一个层名称
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 层名称前缀
	Prefix *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
	// 起始层名称
	StartKey *string `json:"startKey,omitempty" xml:"startKey,omitempty"`
}

func (s ListLayersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLayersRequest) GoString() string {
	return s.String()
}

func (s *ListLayersRequest) SetLimit(v int32) *ListLayersRequest {
	s.Limit = &v
	return s
}

func (s *ListLayersRequest) SetNextToken(v string) *ListLayersRequest {
	s.NextToken = &v
	return s
}

func (s *ListLayersRequest) SetPrefix(v string) *ListLayersRequest {
	s.Prefix = &v
	return s
}

func (s *ListLayersRequest) SetStartKey(v string) *ListLayersRequest {
	s.StartKey = &v
	return s
}

type ListLayersResponseBody struct {
	// 层列表
	Layers []*Layer `json:"layers,omitempty" xml:"layers,omitempty" type:"Repeated"`
	// 剩余列表起始层名
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListLayersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLayersResponseBody) GoString() string {
	return s.String()
}

func (s *ListLayersResponseBody) SetLayers(v []*Layer) *ListLayersResponseBody {
	s.Layers = v
	return s
}

func (s *ListLayersResponseBody) SetNextToken(v string) *ListLayersResponseBody {
	s.NextToken = &v
	return s
}

type ListLayersResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListLayersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListLayersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLayersResponse) GoString() string {
	return s.String()
}

func (s *ListLayersResponse) SetHeaders(v map[string]*string) *ListLayersResponse {
	s.Headers = v
	return s
}

func (s *ListLayersResponse) SetBody(v *ListLayersResponseBody) *ListLayersResponse {
	s.Body = v
	return s
}

type ListOnDemandConfigsHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s ListOnDemandConfigsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListOnDemandConfigsHeaders) GoString() string {
	return s.String()
}

func (s *ListOnDemandConfigsHeaders) SetCommonHeaders(v map[string]*string) *ListOnDemandConfigsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListOnDemandConfigsHeaders) SetXFcAccountId(v string) *ListOnDemandConfigsHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *ListOnDemandConfigsHeaders) SetXFcCodeChecksum(v string) *ListOnDemandConfigsHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *ListOnDemandConfigsHeaders) SetXFcDate(v string) *ListOnDemandConfigsHeaders {
	s.XFcDate = &v
	return s
}

func (s *ListOnDemandConfigsHeaders) SetXFcInvocationType(v string) *ListOnDemandConfigsHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *ListOnDemandConfigsHeaders) SetXFcLogType(v string) *ListOnDemandConfigsHeaders {
	s.XFcLogType = &v
	return s
}

func (s *ListOnDemandConfigsHeaders) SetXFcTraceId(v string) *ListOnDemandConfigsHeaders {
	s.XFcTraceId = &v
	return s
}

type ListOnDemandConfigsRequest struct {
	// 限定此次返回资源的数量。如果不设定，默认返回20，最大不能超过100。返回结果可以小于指定的数量，但不会多于指定的数量。
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// 用来返回更多结果。第一次查询不需要提供这个参数，后续查询的Token从返回结果中获取。
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 限定返回的资源名称，名称必须以Prefix作为前缀，例如Prefix是a，则返回的资源名均是以a开始的。
	Prefix *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
	// 设定结果从startKey之后（包括startKey）按字母排序的第一个开始返回。
	StartKey *string `json:"startKey,omitempty" xml:"startKey,omitempty"`
}

func (s ListOnDemandConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOnDemandConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListOnDemandConfigsRequest) SetLimit(v int32) *ListOnDemandConfigsRequest {
	s.Limit = &v
	return s
}

func (s *ListOnDemandConfigsRequest) SetNextToken(v string) *ListOnDemandConfigsRequest {
	s.NextToken = &v
	return s
}

func (s *ListOnDemandConfigsRequest) SetPrefix(v string) *ListOnDemandConfigsRequest {
	s.Prefix = &v
	return s
}

func (s *ListOnDemandConfigsRequest) SetStartKey(v string) *ListOnDemandConfigsRequest {
	s.StartKey = &v
	return s
}

type ListOnDemandConfigsResponseBody struct {
	// 预留实例配置
	Configs []*OnDemandConfig `json:"configs,omitempty" xml:"configs,omitempty" type:"Repeated"`
	// 用来返回更多的查询结果。如果这个值没有返回，则说明没有更多结果。
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListOnDemandConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOnDemandConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOnDemandConfigsResponseBody) SetConfigs(v []*OnDemandConfig) *ListOnDemandConfigsResponseBody {
	s.Configs = v
	return s
}

func (s *ListOnDemandConfigsResponseBody) SetNextToken(v string) *ListOnDemandConfigsResponseBody {
	s.NextToken = &v
	return s
}

type ListOnDemandConfigsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListOnDemandConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOnDemandConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOnDemandConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListOnDemandConfigsResponse) SetHeaders(v map[string]*string) *ListOnDemandConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListOnDemandConfigsResponse) SetBody(v *ListOnDemandConfigsResponseBody) *ListOnDemandConfigsResponse {
	s.Body = v
	return s
}

type ListProvisionConfigsHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s ListProvisionConfigsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListProvisionConfigsHeaders) GoString() string {
	return s.String()
}

func (s *ListProvisionConfigsHeaders) SetCommonHeaders(v map[string]*string) *ListProvisionConfigsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListProvisionConfigsHeaders) SetXFcAccountId(v string) *ListProvisionConfigsHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *ListProvisionConfigsHeaders) SetXFcCodeChecksum(v string) *ListProvisionConfigsHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *ListProvisionConfigsHeaders) SetXFcDate(v string) *ListProvisionConfigsHeaders {
	s.XFcDate = &v
	return s
}

func (s *ListProvisionConfigsHeaders) SetXFcInvocationType(v string) *ListProvisionConfigsHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *ListProvisionConfigsHeaders) SetXFcLogType(v string) *ListProvisionConfigsHeaders {
	s.XFcLogType = &v
	return s
}

func (s *ListProvisionConfigsHeaders) SetXFcTraceId(v string) *ListProvisionConfigsHeaders {
	s.XFcTraceId = &v
	return s
}

type ListProvisionConfigsRequest struct {
	// 限定此次返回资源的数量。如果不设定，默认返回20，最大不能超过100。返回结果可能小于指定的数量，但不会多于指定的数量
	Limit *int64 `json:"limit,omitempty" xml:"limit,omitempty"`
	// 用来返回更多结果。第一次查询不需要提供这个参数，后续查询的token从返回结果中获取
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 限定返回的资源名称必须属于该qualifier。qualifier只能是aliasName，且必须和serviceName共同使用
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	// 限定返回的资源名称必须属于该service
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
}

func (s ListProvisionConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProvisionConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListProvisionConfigsRequest) SetLimit(v int64) *ListProvisionConfigsRequest {
	s.Limit = &v
	return s
}

func (s *ListProvisionConfigsRequest) SetNextToken(v string) *ListProvisionConfigsRequest {
	s.NextToken = &v
	return s
}

func (s *ListProvisionConfigsRequest) SetQualifier(v string) *ListProvisionConfigsRequest {
	s.Qualifier = &v
	return s
}

func (s *ListProvisionConfigsRequest) SetServiceName(v string) *ListProvisionConfigsRequest {
	s.ServiceName = &v
	return s
}

type ListProvisionConfigsResponseBody struct {
	// 下次查询的起始token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 预留实例列表
	ProvisionConfigs []*ListProvisionConfigsResponseBodyProvisionConfigs `json:"provisionConfigs,omitempty" xml:"provisionConfigs,omitempty" type:"Repeated"`
}

func (s ListProvisionConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProvisionConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProvisionConfigsResponseBody) SetNextToken(v string) *ListProvisionConfigsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListProvisionConfigsResponseBody) SetProvisionConfigs(v []*ListProvisionConfigsResponseBodyProvisionConfigs) *ListProvisionConfigsResponseBody {
	s.ProvisionConfigs = v
	return s
}

type ListProvisionConfigsResponseBodyProvisionConfigs struct {
	// 实际资源个数
	Current *int64 `json:"current,omitempty" xml:"current,omitempty"`
	// 预留实例创建失败时的错误信息
	CurrentError *string `json:"currentError,omitempty" xml:"currentError,omitempty"`
	// 资源描述
	Resource *string `json:"resource,omitempty" xml:"resource,omitempty"`
	// 定时策略配置
	ScheduledActions []*ScheduledActions `json:"scheduledActions,omitempty" xml:"scheduledActions,omitempty" type:"Repeated"`
	// 目标资源个数
	Target *int64 `json:"target,omitempty" xml:"target,omitempty"`
	// 指标追踪伸缩策略配置
	TargetTrackingPolicies []*TargetTrackingPolicies `json:"targetTrackingPolicies,omitempty" xml:"targetTrackingPolicies,omitempty" type:"Repeated"`
}

func (s ListProvisionConfigsResponseBodyProvisionConfigs) String() string {
	return tea.Prettify(s)
}

func (s ListProvisionConfigsResponseBodyProvisionConfigs) GoString() string {
	return s.String()
}

func (s *ListProvisionConfigsResponseBodyProvisionConfigs) SetCurrent(v int64) *ListProvisionConfigsResponseBodyProvisionConfigs {
	s.Current = &v
	return s
}

func (s *ListProvisionConfigsResponseBodyProvisionConfigs) SetCurrentError(v string) *ListProvisionConfigsResponseBodyProvisionConfigs {
	s.CurrentError = &v
	return s
}

func (s *ListProvisionConfigsResponseBodyProvisionConfigs) SetResource(v string) *ListProvisionConfigsResponseBodyProvisionConfigs {
	s.Resource = &v
	return s
}

func (s *ListProvisionConfigsResponseBodyProvisionConfigs) SetScheduledActions(v []*ScheduledActions) *ListProvisionConfigsResponseBodyProvisionConfigs {
	s.ScheduledActions = v
	return s
}

func (s *ListProvisionConfigsResponseBodyProvisionConfigs) SetTarget(v int64) *ListProvisionConfigsResponseBodyProvisionConfigs {
	s.Target = &v
	return s
}

func (s *ListProvisionConfigsResponseBodyProvisionConfigs) SetTargetTrackingPolicies(v []*TargetTrackingPolicies) *ListProvisionConfigsResponseBodyProvisionConfigs {
	s.TargetTrackingPolicies = v
	return s
}

type ListProvisionConfigsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListProvisionConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProvisionConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProvisionConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListProvisionConfigsResponse) SetHeaders(v map[string]*string) *ListProvisionConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListProvisionConfigsResponse) SetBody(v *ListProvisionConfigsResponseBody) *ListProvisionConfigsResponse {
	s.Body = v
	return s
}

type ListReservedCapacitiesHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s ListReservedCapacitiesHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListReservedCapacitiesHeaders) GoString() string {
	return s.String()
}

func (s *ListReservedCapacitiesHeaders) SetCommonHeaders(v map[string]*string) *ListReservedCapacitiesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListReservedCapacitiesHeaders) SetXFcAccountId(v string) *ListReservedCapacitiesHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *ListReservedCapacitiesHeaders) SetXFcCodeChecksum(v string) *ListReservedCapacitiesHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *ListReservedCapacitiesHeaders) SetXFcDate(v string) *ListReservedCapacitiesHeaders {
	s.XFcDate = &v
	return s
}

func (s *ListReservedCapacitiesHeaders) SetXFcInvocationType(v string) *ListReservedCapacitiesHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *ListReservedCapacitiesHeaders) SetXFcLogType(v string) *ListReservedCapacitiesHeaders {
	s.XFcLogType = &v
	return s
}

func (s *ListReservedCapacitiesHeaders) SetXFcTraceId(v string) *ListReservedCapacitiesHeaders {
	s.XFcTraceId = &v
	return s
}

type ListReservedCapacitiesRequest struct {
	Limit     *string `json:"limit,omitempty" xml:"limit,omitempty"`
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListReservedCapacitiesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListReservedCapacitiesRequest) GoString() string {
	return s.String()
}

func (s *ListReservedCapacitiesRequest) SetLimit(v string) *ListReservedCapacitiesRequest {
	s.Limit = &v
	return s
}

func (s *ListReservedCapacitiesRequest) SetNextToken(v string) *ListReservedCapacitiesRequest {
	s.NextToken = &v
	return s
}

type ListReservedCapacitiesResponseBody struct {
	// nextToken
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// reservedCapacities
	ReservedCapacities []*OpenReservedCapacity `json:"reservedCapacities,omitempty" xml:"reservedCapacities,omitempty" type:"Repeated"`
}

func (s ListReservedCapacitiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListReservedCapacitiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListReservedCapacitiesResponseBody) SetNextToken(v string) *ListReservedCapacitiesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListReservedCapacitiesResponseBody) SetReservedCapacities(v []*OpenReservedCapacity) *ListReservedCapacitiesResponseBody {
	s.ReservedCapacities = v
	return s
}

type ListReservedCapacitiesResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListReservedCapacitiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListReservedCapacitiesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListReservedCapacitiesResponse) GoString() string {
	return s.String()
}

func (s *ListReservedCapacitiesResponse) SetHeaders(v map[string]*string) *ListReservedCapacitiesResponse {
	s.Headers = v
	return s
}

func (s *ListReservedCapacitiesResponse) SetBody(v *ListReservedCapacitiesResponseBody) *ListReservedCapacitiesResponse {
	s.Body = v
	return s
}

type ListServiceVersionsHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s ListServiceVersionsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListServiceVersionsHeaders) GoString() string {
	return s.String()
}

func (s *ListServiceVersionsHeaders) SetCommonHeaders(v map[string]*string) *ListServiceVersionsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListServiceVersionsHeaders) SetXFcAccountId(v string) *ListServiceVersionsHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *ListServiceVersionsHeaders) SetXFcCodeChecksum(v string) *ListServiceVersionsHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *ListServiceVersionsHeaders) SetXFcDate(v string) *ListServiceVersionsHeaders {
	s.XFcDate = &v
	return s
}

func (s *ListServiceVersionsHeaders) SetXFcInvocationType(v string) *ListServiceVersionsHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *ListServiceVersionsHeaders) SetXFcLogType(v string) *ListServiceVersionsHeaders {
	s.XFcLogType = &v
	return s
}

func (s *ListServiceVersionsHeaders) SetXFcTraceId(v string) *ListServiceVersionsHeaders {
	s.XFcTraceId = &v
	return s
}

type ListServiceVersionsRequest struct {
	// 排序方向
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	// 最多返回个数
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// 下次查询token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 起始key
	StartKey *string `json:"startKey,omitempty" xml:"startKey,omitempty"`
}

func (s ListServiceVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListServiceVersionsRequest) SetDirection(v string) *ListServiceVersionsRequest {
	s.Direction = &v
	return s
}

func (s *ListServiceVersionsRequest) SetLimit(v int32) *ListServiceVersionsRequest {
	s.Limit = &v
	return s
}

func (s *ListServiceVersionsRequest) SetNextToken(v string) *ListServiceVersionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceVersionsRequest) SetStartKey(v string) *ListServiceVersionsRequest {
	s.StartKey = &v
	return s
}

type ListServiceVersionsResponseBody struct {
	// 排序方向
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	// 下次查询token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 版本列表
	Versions []*ListServiceVersionsResponseBodyVersions `json:"versions,omitempty" xml:"versions,omitempty" type:"Repeated"`
}

func (s ListServiceVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceVersionsResponseBody) SetDirection(v string) *ListServiceVersionsResponseBody {
	s.Direction = &v
	return s
}

func (s *ListServiceVersionsResponseBody) SetNextToken(v string) *ListServiceVersionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceVersionsResponseBody) SetVersions(v []*ListServiceVersionsResponseBodyVersions) *ListServiceVersionsResponseBody {
	s.Versions = v
	return s
}

type ListServiceVersionsResponseBodyVersions struct {
	// 创建时间
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// 版本描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 上次更新时间
	LastModifiedTime *string `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	// 版本ID
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s ListServiceVersionsResponseBodyVersions) String() string {
	return tea.Prettify(s)
}

func (s ListServiceVersionsResponseBodyVersions) GoString() string {
	return s.String()
}

func (s *ListServiceVersionsResponseBodyVersions) SetCreatedTime(v string) *ListServiceVersionsResponseBodyVersions {
	s.CreatedTime = &v
	return s
}

func (s *ListServiceVersionsResponseBodyVersions) SetDescription(v string) *ListServiceVersionsResponseBodyVersions {
	s.Description = &v
	return s
}

func (s *ListServiceVersionsResponseBodyVersions) SetLastModifiedTime(v string) *ListServiceVersionsResponseBodyVersions {
	s.LastModifiedTime = &v
	return s
}

func (s *ListServiceVersionsResponseBodyVersions) SetVersionId(v string) *ListServiceVersionsResponseBodyVersions {
	s.VersionId = &v
	return s
}

type ListServiceVersionsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListServiceVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListServiceVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListServiceVersionsResponse) SetHeaders(v map[string]*string) *ListServiceVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListServiceVersionsResponse) SetBody(v *ListServiceVersionsResponseBody) *ListServiceVersionsResponse {
	s.Body = v
	return s
}

type ListServicesHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s ListServicesHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListServicesHeaders) GoString() string {
	return s.String()
}

func (s *ListServicesHeaders) SetCommonHeaders(v map[string]*string) *ListServicesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListServicesHeaders) SetXFcAccountId(v string) *ListServicesHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *ListServicesHeaders) SetXFcCodeChecksum(v string) *ListServicesHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *ListServicesHeaders) SetXFcDate(v string) *ListServicesHeaders {
	s.XFcDate = &v
	return s
}

func (s *ListServicesHeaders) SetXFcInvocationType(v string) *ListServicesHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *ListServicesHeaders) SetXFcLogType(v string) *ListServicesHeaders {
	s.XFcLogType = &v
	return s
}

func (s *ListServicesHeaders) SetXFcTraceId(v string) *ListServicesHeaders {
	s.XFcTraceId = &v
	return s
}

type ListServicesRequest struct {
	// 最多返回个数
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// 下次查询token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 前缀
	Prefix *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
	// 起始key
	StartKey *string `json:"startKey,omitempty" xml:"startKey,omitempty"`
}

func (s ListServicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServicesRequest) GoString() string {
	return s.String()
}

func (s *ListServicesRequest) SetLimit(v int32) *ListServicesRequest {
	s.Limit = &v
	return s
}

func (s *ListServicesRequest) SetNextToken(v string) *ListServicesRequest {
	s.NextToken = &v
	return s
}

func (s *ListServicesRequest) SetPrefix(v string) *ListServicesRequest {
	s.Prefix = &v
	return s
}

func (s *ListServicesRequest) SetStartKey(v string) *ListServicesRequest {
	s.StartKey = &v
	return s
}

type ListServicesResponseBody struct {
	// 下次查询token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 服务列表
	Services []*ListServicesResponseBodyServices `json:"services,omitempty" xml:"services,omitempty" type:"Repeated"`
}

func (s ListServicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBody) SetNextToken(v string) *ListServicesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServicesResponseBody) SetServices(v []*ListServicesResponseBodyServices) *ListServicesResponseBody {
	s.Services = v
	return s
}

type ListServicesResponseBodyServices struct {
	// 创建时间
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// 服务描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 公网访问设置
	InternetAccess *bool `json:"internetAccess,omitempty" xml:"internetAccess,omitempty"`
	// 上次更新时间
	LastModifiedTime *string    `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	LogConfig        *LogConfig `json:"logConfig,omitempty" xml:"logConfig,omitempty"`
	NasConfig        *NASConfig `json:"nasConfig,omitempty" xml:"nasConfig,omitempty"`
	// 服务角色
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// 服务ID
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// 服务信息
	ServiceName   *string        `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	TracingConfig *TracingConfig `json:"tracingConfig,omitempty" xml:"tracingConfig,omitempty"`
	VpcConfig     *VPCConfig     `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty"`
}

func (s ListServicesResponseBodyServices) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBodyServices) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBodyServices) SetCreatedTime(v string) *ListServicesResponseBodyServices {
	s.CreatedTime = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetDescription(v string) *ListServicesResponseBodyServices {
	s.Description = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetInternetAccess(v bool) *ListServicesResponseBodyServices {
	s.InternetAccess = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetLastModifiedTime(v string) *ListServicesResponseBodyServices {
	s.LastModifiedTime = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetLogConfig(v *LogConfig) *ListServicesResponseBodyServices {
	s.LogConfig = v
	return s
}

func (s *ListServicesResponseBodyServices) SetNasConfig(v *NASConfig) *ListServicesResponseBodyServices {
	s.NasConfig = v
	return s
}

func (s *ListServicesResponseBodyServices) SetRole(v string) *ListServicesResponseBodyServices {
	s.Role = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetServiceId(v string) *ListServicesResponseBodyServices {
	s.ServiceId = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetServiceName(v string) *ListServicesResponseBodyServices {
	s.ServiceName = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetTracingConfig(v *TracingConfig) *ListServicesResponseBodyServices {
	s.TracingConfig = v
	return s
}

func (s *ListServicesResponseBodyServices) SetVpcConfig(v *VPCConfig) *ListServicesResponseBodyServices {
	s.VpcConfig = v
	return s
}

type ListServicesResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListServicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListServicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponse) GoString() string {
	return s.String()
}

func (s *ListServicesResponse) SetHeaders(v map[string]*string) *ListServicesResponse {
	s.Headers = v
	return s
}

func (s *ListServicesResponse) SetBody(v *ListServicesResponseBody) *ListServicesResponse {
	s.Body = v
	return s
}

type ListStatefulAsyncInvocationsHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s ListStatefulAsyncInvocationsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListStatefulAsyncInvocationsHeaders) GoString() string {
	return s.String()
}

func (s *ListStatefulAsyncInvocationsHeaders) SetCommonHeaders(v map[string]*string) *ListStatefulAsyncInvocationsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListStatefulAsyncInvocationsHeaders) SetXFcAccountId(v string) *ListStatefulAsyncInvocationsHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *ListStatefulAsyncInvocationsHeaders) SetXFcCodeChecksum(v string) *ListStatefulAsyncInvocationsHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *ListStatefulAsyncInvocationsHeaders) SetXFcDate(v string) *ListStatefulAsyncInvocationsHeaders {
	s.XFcDate = &v
	return s
}

func (s *ListStatefulAsyncInvocationsHeaders) SetXFcInvocationType(v string) *ListStatefulAsyncInvocationsHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *ListStatefulAsyncInvocationsHeaders) SetXFcLogType(v string) *ListStatefulAsyncInvocationsHeaders {
	s.XFcLogType = &v
	return s
}

func (s *ListStatefulAsyncInvocationsHeaders) SetXFcTraceId(v string) *ListStatefulAsyncInvocationsHeaders {
	s.XFcTraceId = &v
	return s
}

type ListStatefulAsyncInvocationsRequest struct {
	IncludePayload     *bool   `json:"includePayload,omitempty" xml:"includePayload,omitempty"`
	InvocationIdPrefix *string `json:"invocationIdPrefix,omitempty" xml:"invocationIdPrefix,omitempty"`
	Limit              *int32  `json:"limit,omitempty" xml:"limit,omitempty"`
	NextToken          *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Qualifier          *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	SortOrderByTime    *string `json:"sortOrderByTime,omitempty" xml:"sortOrderByTime,omitempty"`
	StartedTimeBegin   *int64  `json:"startedTimeBegin,omitempty" xml:"startedTimeBegin,omitempty"`
	StartedTimeEnd     *int64  `json:"startedTimeEnd,omitempty" xml:"startedTimeEnd,omitempty"`
	Status             *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListStatefulAsyncInvocationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListStatefulAsyncInvocationsRequest) GoString() string {
	return s.String()
}

func (s *ListStatefulAsyncInvocationsRequest) SetIncludePayload(v bool) *ListStatefulAsyncInvocationsRequest {
	s.IncludePayload = &v
	return s
}

func (s *ListStatefulAsyncInvocationsRequest) SetInvocationIdPrefix(v string) *ListStatefulAsyncInvocationsRequest {
	s.InvocationIdPrefix = &v
	return s
}

func (s *ListStatefulAsyncInvocationsRequest) SetLimit(v int32) *ListStatefulAsyncInvocationsRequest {
	s.Limit = &v
	return s
}

func (s *ListStatefulAsyncInvocationsRequest) SetNextToken(v string) *ListStatefulAsyncInvocationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListStatefulAsyncInvocationsRequest) SetQualifier(v string) *ListStatefulAsyncInvocationsRequest {
	s.Qualifier = &v
	return s
}

func (s *ListStatefulAsyncInvocationsRequest) SetSortOrderByTime(v string) *ListStatefulAsyncInvocationsRequest {
	s.SortOrderByTime = &v
	return s
}

func (s *ListStatefulAsyncInvocationsRequest) SetStartedTimeBegin(v int64) *ListStatefulAsyncInvocationsRequest {
	s.StartedTimeBegin = &v
	return s
}

func (s *ListStatefulAsyncInvocationsRequest) SetStartedTimeEnd(v int64) *ListStatefulAsyncInvocationsRequest {
	s.StartedTimeEnd = &v
	return s
}

func (s *ListStatefulAsyncInvocationsRequest) SetStatus(v string) *ListStatefulAsyncInvocationsRequest {
	s.Status = &v
	return s
}

type ListStatefulAsyncInvocationsResponseBody struct {
	Invocations []*StatefulAsyncInvocation `json:"invocations,omitempty" xml:"invocations,omitempty" type:"Repeated"`
	NextToken   *string                    `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListStatefulAsyncInvocationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListStatefulAsyncInvocationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListStatefulAsyncInvocationsResponseBody) SetInvocations(v []*StatefulAsyncInvocation) *ListStatefulAsyncInvocationsResponseBody {
	s.Invocations = v
	return s
}

func (s *ListStatefulAsyncInvocationsResponseBody) SetNextToken(v string) *ListStatefulAsyncInvocationsResponseBody {
	s.NextToken = &v
	return s
}

type ListStatefulAsyncInvocationsResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListStatefulAsyncInvocationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListStatefulAsyncInvocationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListStatefulAsyncInvocationsResponse) GoString() string {
	return s.String()
}

func (s *ListStatefulAsyncInvocationsResponse) SetHeaders(v map[string]*string) *ListStatefulAsyncInvocationsResponse {
	s.Headers = v
	return s
}

func (s *ListStatefulAsyncInvocationsResponse) SetBody(v *ListStatefulAsyncInvocationsResponseBody) *ListStatefulAsyncInvocationsResponse {
	s.Body = v
	return s
}

type ListTaggedResourcesHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s ListTaggedResourcesHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListTaggedResourcesHeaders) GoString() string {
	return s.String()
}

func (s *ListTaggedResourcesHeaders) SetCommonHeaders(v map[string]*string) *ListTaggedResourcesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListTaggedResourcesHeaders) SetXFcAccountId(v string) *ListTaggedResourcesHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *ListTaggedResourcesHeaders) SetXFcCodeChecksum(v string) *ListTaggedResourcesHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *ListTaggedResourcesHeaders) SetXFcDate(v string) *ListTaggedResourcesHeaders {
	s.XFcDate = &v
	return s
}

func (s *ListTaggedResourcesHeaders) SetXFcInvocationType(v string) *ListTaggedResourcesHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *ListTaggedResourcesHeaders) SetXFcLogType(v string) *ListTaggedResourcesHeaders {
	s.XFcLogType = &v
	return s
}

func (s *ListTaggedResourcesHeaders) SetXFcTraceId(v string) *ListTaggedResourcesHeaders {
	s.XFcTraceId = &v
	return s
}

type ListTaggedResourcesRequest struct {
	Limit     *int32  `json:"limit,omitempty" xml:"limit,omitempty"`
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListTaggedResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTaggedResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTaggedResourcesRequest) SetLimit(v int32) *ListTaggedResourcesRequest {
	s.Limit = &v
	return s
}

func (s *ListTaggedResourcesRequest) SetNextToken(v string) *ListTaggedResourcesRequest {
	s.NextToken = &v
	return s
}

type ListTaggedResourcesResponseBody struct {
	NextToken *string     `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Resources []*Resource `json:"resources,omitempty" xml:"resources,omitempty" type:"Repeated"`
}

func (s ListTaggedResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTaggedResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTaggedResourcesResponseBody) SetNextToken(v string) *ListTaggedResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTaggedResourcesResponseBody) SetResources(v []*Resource) *ListTaggedResourcesResponseBody {
	s.Resources = v
	return s
}

type ListTaggedResourcesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTaggedResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTaggedResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTaggedResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTaggedResourcesResponse) SetHeaders(v map[string]*string) *ListTaggedResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTaggedResourcesResponse) SetBody(v *ListTaggedResourcesResponseBody) *ListTaggedResourcesResponse {
	s.Body = v
	return s
}

type ListTriggersHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s ListTriggersHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListTriggersHeaders) GoString() string {
	return s.String()
}

func (s *ListTriggersHeaders) SetCommonHeaders(v map[string]*string) *ListTriggersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListTriggersHeaders) SetXFcAccountId(v string) *ListTriggersHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *ListTriggersHeaders) SetXFcCodeChecksum(v string) *ListTriggersHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *ListTriggersHeaders) SetXFcDate(v string) *ListTriggersHeaders {
	s.XFcDate = &v
	return s
}

func (s *ListTriggersHeaders) SetXFcInvocationType(v string) *ListTriggersHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *ListTriggersHeaders) SetXFcLogType(v string) *ListTriggersHeaders {
	s.XFcLogType = &v
	return s
}

func (s *ListTriggersHeaders) SetXFcTraceId(v string) *ListTriggersHeaders {
	s.XFcTraceId = &v
	return s
}

type ListTriggersRequest struct {
	// 限定此次返回资源的数量。如果不设定，默认返回20，最大不能超过100。返回结果可能小于指定的数量，但不会多于指定的数量
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// 用来返回更多结果。第一次查询不需要提供这个参数，后续查询的token从返回结果中获取
	NextToken *int32 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 限定返回的资源名称必须以prefix作为前缀
	Prefix *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
	// 设定结果从startKey之后（包括startKey）按字母排序的第一个开始返回
	StartKey *string `json:"startKey,omitempty" xml:"startKey,omitempty"`
}

func (s ListTriggersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTriggersRequest) GoString() string {
	return s.String()
}

func (s *ListTriggersRequest) SetLimit(v int32) *ListTriggersRequest {
	s.Limit = &v
	return s
}

func (s *ListTriggersRequest) SetNextToken(v int32) *ListTriggersRequest {
	s.NextToken = &v
	return s
}

func (s *ListTriggersRequest) SetPrefix(v string) *ListTriggersRequest {
	s.Prefix = &v
	return s
}

func (s *ListTriggersRequest) SetStartKey(v string) *ListTriggersRequest {
	s.StartKey = &v
	return s
}

type ListTriggersResponseBody struct {
	// 用来返回更多的查询结果。如果这个值没有返回，则说明没有更多结果
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Trigger列表
	Triggers []*ListTriggersResponseBodyTriggers `json:"triggers,omitempty" xml:"triggers,omitempty" type:"Repeated"`
}

func (s ListTriggersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTriggersResponseBody) GoString() string {
	return s.String()
}

func (s *ListTriggersResponseBody) SetNextToken(v string) *ListTriggersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTriggersResponseBody) SetTriggers(v []*ListTriggersResponseBodyTriggers) *ListTriggersResponseBody {
	s.Triggers = v
	return s
}

type ListTriggersResponseBodyTriggers struct {
	// 创建时间
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 域名名称，使用域名名称拼接上函数计算域名，可以采用HTTP协议调用到触发器对应版本的函数。例如{domainName}.cn-shanghai.fc.aliyuncs.com
	DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
	// 调用函数使用的RAM角色的ARN
	InvocationRole *string `json:"invocationRole,omitempty" xml:"invocationRole,omitempty"`
	// 上次修改时间
	LastModifiedTime *string `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	// service版本
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	// event source的Aliyun Resource Name（ARN
	SourceArn *string `json:"sourceArn,omitempty" xml:"sourceArn,omitempty"`
	// trigger配置对象
	TriggerConfig *string `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty"`
	TriggerId     *string `json:"triggerId,omitempty" xml:"triggerId,omitempty"`
	// trigger名称
	TriggerName *string `json:"triggerName,omitempty" xml:"triggerName,omitempty"`
	// trigger类型，如 oss, log, tablestore, timer, http, cdn_events, mns_topic
	TriggerType *string `json:"triggerType,omitempty" xml:"triggerType,omitempty"`
}

func (s ListTriggersResponseBodyTriggers) String() string {
	return tea.Prettify(s)
}

func (s ListTriggersResponseBodyTriggers) GoString() string {
	return s.String()
}

func (s *ListTriggersResponseBodyTriggers) SetCreatedTime(v string) *ListTriggersResponseBodyTriggers {
	s.CreatedTime = &v
	return s
}

func (s *ListTriggersResponseBodyTriggers) SetDescription(v string) *ListTriggersResponseBodyTriggers {
	s.Description = &v
	return s
}

func (s *ListTriggersResponseBodyTriggers) SetDomainName(v string) *ListTriggersResponseBodyTriggers {
	s.DomainName = &v
	return s
}

func (s *ListTriggersResponseBodyTriggers) SetInvocationRole(v string) *ListTriggersResponseBodyTriggers {
	s.InvocationRole = &v
	return s
}

func (s *ListTriggersResponseBodyTriggers) SetLastModifiedTime(v string) *ListTriggersResponseBodyTriggers {
	s.LastModifiedTime = &v
	return s
}

func (s *ListTriggersResponseBodyTriggers) SetQualifier(v string) *ListTriggersResponseBodyTriggers {
	s.Qualifier = &v
	return s
}

func (s *ListTriggersResponseBodyTriggers) SetSourceArn(v string) *ListTriggersResponseBodyTriggers {
	s.SourceArn = &v
	return s
}

func (s *ListTriggersResponseBodyTriggers) SetTriggerConfig(v string) *ListTriggersResponseBodyTriggers {
	s.TriggerConfig = &v
	return s
}

func (s *ListTriggersResponseBodyTriggers) SetTriggerId(v string) *ListTriggersResponseBodyTriggers {
	s.TriggerId = &v
	return s
}

func (s *ListTriggersResponseBodyTriggers) SetTriggerName(v string) *ListTriggersResponseBodyTriggers {
	s.TriggerName = &v
	return s
}

func (s *ListTriggersResponseBodyTriggers) SetTriggerType(v string) *ListTriggersResponseBodyTriggers {
	s.TriggerType = &v
	return s
}

type ListTriggersResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTriggersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTriggersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTriggersResponse) GoString() string {
	return s.String()
}

func (s *ListTriggersResponse) SetHeaders(v map[string]*string) *ListTriggersResponse {
	s.Headers = v
	return s
}

func (s *ListTriggersResponse) SetBody(v *ListTriggersResponseBody) *ListTriggersResponse {
	s.Body = v
	return s
}

type ListVpcBindingsHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s ListVpcBindingsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListVpcBindingsHeaders) GoString() string {
	return s.String()
}

func (s *ListVpcBindingsHeaders) SetCommonHeaders(v map[string]*string) *ListVpcBindingsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListVpcBindingsHeaders) SetXFcAccountId(v string) *ListVpcBindingsHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *ListVpcBindingsHeaders) SetXFcCodeChecksum(v string) *ListVpcBindingsHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *ListVpcBindingsHeaders) SetXFcDate(v string) *ListVpcBindingsHeaders {
	s.XFcDate = &v
	return s
}

func (s *ListVpcBindingsHeaders) SetXFcInvocationType(v string) *ListVpcBindingsHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *ListVpcBindingsHeaders) SetXFcLogType(v string) *ListVpcBindingsHeaders {
	s.XFcLogType = &v
	return s
}

func (s *ListVpcBindingsHeaders) SetXFcTraceId(v string) *ListVpcBindingsHeaders {
	s.XFcTraceId = &v
	return s
}

type ListVpcBindingsResponseBody struct {
	VpcIds []*string `json:"vpcIds,omitempty" xml:"vpcIds,omitempty" type:"Repeated"`
}

func (s ListVpcBindingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVpcBindingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpcBindingsResponseBody) SetVpcIds(v []*string) *ListVpcBindingsResponseBody {
	s.VpcIds = v
	return s
}

type ListVpcBindingsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListVpcBindingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVpcBindingsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVpcBindingsResponse) GoString() string {
	return s.String()
}

func (s *ListVpcBindingsResponse) SetHeaders(v map[string]*string) *ListVpcBindingsResponse {
	s.Headers = v
	return s
}

func (s *ListVpcBindingsResponse) SetBody(v *ListVpcBindingsResponseBody) *ListVpcBindingsResponse {
	s.Body = v
	return s
}

type PermanentDeleteLayerVersionHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s PermanentDeleteLayerVersionHeaders) String() string {
	return tea.Prettify(s)
}

func (s PermanentDeleteLayerVersionHeaders) GoString() string {
	return s.String()
}

func (s *PermanentDeleteLayerVersionHeaders) SetCommonHeaders(v map[string]*string) *PermanentDeleteLayerVersionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PermanentDeleteLayerVersionHeaders) SetXFcAccountId(v string) *PermanentDeleteLayerVersionHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *PermanentDeleteLayerVersionHeaders) SetXFcCodeChecksum(v string) *PermanentDeleteLayerVersionHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *PermanentDeleteLayerVersionHeaders) SetXFcDate(v string) *PermanentDeleteLayerVersionHeaders {
	s.XFcDate = &v
	return s
}

func (s *PermanentDeleteLayerVersionHeaders) SetXFcInvocationType(v string) *PermanentDeleteLayerVersionHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *PermanentDeleteLayerVersionHeaders) SetXFcLogType(v string) *PermanentDeleteLayerVersionHeaders {
	s.XFcLogType = &v
	return s
}

func (s *PermanentDeleteLayerVersionHeaders) SetXFcTraceId(v string) *PermanentDeleteLayerVersionHeaders {
	s.XFcTraceId = &v
	return s
}

type PermanentDeleteLayerVersionResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s PermanentDeleteLayerVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s PermanentDeleteLayerVersionResponse) GoString() string {
	return s.String()
}

func (s *PermanentDeleteLayerVersionResponse) SetHeaders(v map[string]*string) *PermanentDeleteLayerVersionResponse {
	s.Headers = v
	return s
}

type PublishLayerAsPublicHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s PublishLayerAsPublicHeaders) String() string {
	return tea.Prettify(s)
}

func (s PublishLayerAsPublicHeaders) GoString() string {
	return s.String()
}

func (s *PublishLayerAsPublicHeaders) SetCommonHeaders(v map[string]*string) *PublishLayerAsPublicHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PublishLayerAsPublicHeaders) SetXFcAccountId(v string) *PublishLayerAsPublicHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *PublishLayerAsPublicHeaders) SetXFcCodeChecksum(v string) *PublishLayerAsPublicHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *PublishLayerAsPublicHeaders) SetXFcDate(v string) *PublishLayerAsPublicHeaders {
	s.XFcDate = &v
	return s
}

func (s *PublishLayerAsPublicHeaders) SetXFcInvocationType(v string) *PublishLayerAsPublicHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *PublishLayerAsPublicHeaders) SetXFcLogType(v string) *PublishLayerAsPublicHeaders {
	s.XFcLogType = &v
	return s
}

func (s *PublishLayerAsPublicHeaders) SetXFcTraceId(v string) *PublishLayerAsPublicHeaders {
	s.XFcTraceId = &v
	return s
}

type PublishLayerAsPublicResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s PublishLayerAsPublicResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishLayerAsPublicResponse) GoString() string {
	return s.String()
}

func (s *PublishLayerAsPublicResponse) SetHeaders(v map[string]*string) *PublishLayerAsPublicResponse {
	s.Headers = v
	return s
}

type PublishServiceVersionHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	IfMatch           *string            `json:"If-Match,omitempty" xml:"If-Match,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s PublishServiceVersionHeaders) String() string {
	return tea.Prettify(s)
}

func (s PublishServiceVersionHeaders) GoString() string {
	return s.String()
}

func (s *PublishServiceVersionHeaders) SetCommonHeaders(v map[string]*string) *PublishServiceVersionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PublishServiceVersionHeaders) SetIfMatch(v string) *PublishServiceVersionHeaders {
	s.IfMatch = &v
	return s
}

func (s *PublishServiceVersionHeaders) SetXFcAccountId(v string) *PublishServiceVersionHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *PublishServiceVersionHeaders) SetXFcCodeChecksum(v string) *PublishServiceVersionHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *PublishServiceVersionHeaders) SetXFcDate(v string) *PublishServiceVersionHeaders {
	s.XFcDate = &v
	return s
}

func (s *PublishServiceVersionHeaders) SetXFcInvocationType(v string) *PublishServiceVersionHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *PublishServiceVersionHeaders) SetXFcLogType(v string) *PublishServiceVersionHeaders {
	s.XFcLogType = &v
	return s
}

func (s *PublishServiceVersionHeaders) SetXFcTraceId(v string) *PublishServiceVersionHeaders {
	s.XFcTraceId = &v
	return s
}

type PublishServiceVersionRequest struct {
	// 版本描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s PublishServiceVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishServiceVersionRequest) GoString() string {
	return s.String()
}

func (s *PublishServiceVersionRequest) SetDescription(v string) *PublishServiceVersionRequest {
	s.Description = &v
	return s
}

type PublishServiceVersionResponseBody struct {
	// 创建时间
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// 版本描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 上次更新时间
	LastModifiedTime *string `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	// 版本ID
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s PublishServiceVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishServiceVersionResponseBody) GoString() string {
	return s.String()
}

func (s *PublishServiceVersionResponseBody) SetCreatedTime(v string) *PublishServiceVersionResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *PublishServiceVersionResponseBody) SetDescription(v string) *PublishServiceVersionResponseBody {
	s.Description = &v
	return s
}

func (s *PublishServiceVersionResponseBody) SetLastModifiedTime(v string) *PublishServiceVersionResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *PublishServiceVersionResponseBody) SetVersionId(v string) *PublishServiceVersionResponseBody {
	s.VersionId = &v
	return s
}

type PublishServiceVersionResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PublishServiceVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PublishServiceVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishServiceVersionResponse) GoString() string {
	return s.String()
}

func (s *PublishServiceVersionResponse) SetHeaders(v map[string]*string) *PublishServiceVersionResponse {
	s.Headers = v
	return s
}

func (s *PublishServiceVersionResponse) SetBody(v *PublishServiceVersionResponseBody) *PublishServiceVersionResponse {
	s.Body = v
	return s
}

type PutFunctionAsyncInvokeConfigHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s PutFunctionAsyncInvokeConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s PutFunctionAsyncInvokeConfigHeaders) GoString() string {
	return s.String()
}

func (s *PutFunctionAsyncInvokeConfigHeaders) SetCommonHeaders(v map[string]*string) *PutFunctionAsyncInvokeConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PutFunctionAsyncInvokeConfigHeaders) SetXFcAccountId(v string) *PutFunctionAsyncInvokeConfigHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *PutFunctionAsyncInvokeConfigHeaders) SetXFcCodeChecksum(v string) *PutFunctionAsyncInvokeConfigHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *PutFunctionAsyncInvokeConfigHeaders) SetXFcDate(v string) *PutFunctionAsyncInvokeConfigHeaders {
	s.XFcDate = &v
	return s
}

func (s *PutFunctionAsyncInvokeConfigHeaders) SetXFcInvocationType(v string) *PutFunctionAsyncInvokeConfigHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *PutFunctionAsyncInvokeConfigHeaders) SetXFcLogType(v string) *PutFunctionAsyncInvokeConfigHeaders {
	s.XFcLogType = &v
	return s
}

func (s *PutFunctionAsyncInvokeConfigHeaders) SetXFcTraceId(v string) *PutFunctionAsyncInvokeConfigHeaders {
	s.XFcTraceId = &v
	return s
}

type PutFunctionAsyncInvokeConfigRequest struct {
	DestinationConfig *DestinationConfig `json:"destinationConfig,omitempty" xml:"destinationConfig,omitempty"`
	// 消息最大存活时长
	MaxAsyncEventAgeInSeconds *int64 `json:"maxAsyncEventAgeInSeconds,omitempty" xml:"maxAsyncEventAgeInSeconds,omitempty"`
	// 异步调用失败后的最大重试次数
	MaxAsyncRetryAttempts *int64 `json:"maxAsyncRetryAttempts,omitempty" xml:"maxAsyncRetryAttempts,omitempty"`
	StatefulInvocation    *bool  `json:"statefulInvocation,omitempty" xml:"statefulInvocation,omitempty"`
	// 别名或版本
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s PutFunctionAsyncInvokeConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s PutFunctionAsyncInvokeConfigRequest) GoString() string {
	return s.String()
}

func (s *PutFunctionAsyncInvokeConfigRequest) SetDestinationConfig(v *DestinationConfig) *PutFunctionAsyncInvokeConfigRequest {
	s.DestinationConfig = v
	return s
}

func (s *PutFunctionAsyncInvokeConfigRequest) SetMaxAsyncEventAgeInSeconds(v int64) *PutFunctionAsyncInvokeConfigRequest {
	s.MaxAsyncEventAgeInSeconds = &v
	return s
}

func (s *PutFunctionAsyncInvokeConfigRequest) SetMaxAsyncRetryAttempts(v int64) *PutFunctionAsyncInvokeConfigRequest {
	s.MaxAsyncRetryAttempts = &v
	return s
}

func (s *PutFunctionAsyncInvokeConfigRequest) SetStatefulInvocation(v bool) *PutFunctionAsyncInvokeConfigRequest {
	s.StatefulInvocation = &v
	return s
}

func (s *PutFunctionAsyncInvokeConfigRequest) SetQualifier(v string) *PutFunctionAsyncInvokeConfigRequest {
	s.Qualifier = &v
	return s
}

type PutFunctionAsyncInvokeConfigResponseBody struct {
	// 创建时间
	CreatedTime       *string            `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	DestinationConfig *DestinationConfig `json:"destinationConfig,omitempty" xml:"destinationConfig,omitempty"`
	// 函数名称
	Function *string `json:"function,omitempty" xml:"function,omitempty"`
	// 最后更改时间
	LastModifiedTime *string `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	// 消息最大存活时长
	MaxAsyncEventAgeInSeconds *int64 `json:"maxAsyncEventAgeInSeconds,omitempty" xml:"maxAsyncEventAgeInSeconds,omitempty"`
	// 异步调用失败后的最大重试次数
	MaxAsyncRetryAttempts *int64 `json:"maxAsyncRetryAttempts,omitempty" xml:"maxAsyncRetryAttempts,omitempty"`
	// 限定符
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	// 服务名称
	Service            *string `json:"service,omitempty" xml:"service,omitempty"`
	StatefulInvocation *bool   `json:"statefulInvocation,omitempty" xml:"statefulInvocation,omitempty"`
}

func (s PutFunctionAsyncInvokeConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutFunctionAsyncInvokeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *PutFunctionAsyncInvokeConfigResponseBody) SetCreatedTime(v string) *PutFunctionAsyncInvokeConfigResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *PutFunctionAsyncInvokeConfigResponseBody) SetDestinationConfig(v *DestinationConfig) *PutFunctionAsyncInvokeConfigResponseBody {
	s.DestinationConfig = v
	return s
}

func (s *PutFunctionAsyncInvokeConfigResponseBody) SetFunction(v string) *PutFunctionAsyncInvokeConfigResponseBody {
	s.Function = &v
	return s
}

func (s *PutFunctionAsyncInvokeConfigResponseBody) SetLastModifiedTime(v string) *PutFunctionAsyncInvokeConfigResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *PutFunctionAsyncInvokeConfigResponseBody) SetMaxAsyncEventAgeInSeconds(v int64) *PutFunctionAsyncInvokeConfigResponseBody {
	s.MaxAsyncEventAgeInSeconds = &v
	return s
}

func (s *PutFunctionAsyncInvokeConfigResponseBody) SetMaxAsyncRetryAttempts(v int64) *PutFunctionAsyncInvokeConfigResponseBody {
	s.MaxAsyncRetryAttempts = &v
	return s
}

func (s *PutFunctionAsyncInvokeConfigResponseBody) SetQualifier(v string) *PutFunctionAsyncInvokeConfigResponseBody {
	s.Qualifier = &v
	return s
}

func (s *PutFunctionAsyncInvokeConfigResponseBody) SetService(v string) *PutFunctionAsyncInvokeConfigResponseBody {
	s.Service = &v
	return s
}

func (s *PutFunctionAsyncInvokeConfigResponseBody) SetStatefulInvocation(v bool) *PutFunctionAsyncInvokeConfigResponseBody {
	s.StatefulInvocation = &v
	return s
}

type PutFunctionAsyncInvokeConfigResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PutFunctionAsyncInvokeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutFunctionAsyncInvokeConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s PutFunctionAsyncInvokeConfigResponse) GoString() string {
	return s.String()
}

func (s *PutFunctionAsyncInvokeConfigResponse) SetHeaders(v map[string]*string) *PutFunctionAsyncInvokeConfigResponse {
	s.Headers = v
	return s
}

func (s *PutFunctionAsyncInvokeConfigResponse) SetBody(v *PutFunctionAsyncInvokeConfigResponseBody) *PutFunctionAsyncInvokeConfigResponse {
	s.Body = v
	return s
}

type PutFunctionOnDemandConfigHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	IfMatch           *string            `json:"If-Match,omitempty" xml:"If-Match,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s PutFunctionOnDemandConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s PutFunctionOnDemandConfigHeaders) GoString() string {
	return s.String()
}

func (s *PutFunctionOnDemandConfigHeaders) SetCommonHeaders(v map[string]*string) *PutFunctionOnDemandConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PutFunctionOnDemandConfigHeaders) SetIfMatch(v string) *PutFunctionOnDemandConfigHeaders {
	s.IfMatch = &v
	return s
}

func (s *PutFunctionOnDemandConfigHeaders) SetXFcAccountId(v string) *PutFunctionOnDemandConfigHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *PutFunctionOnDemandConfigHeaders) SetXFcCodeChecksum(v string) *PutFunctionOnDemandConfigHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *PutFunctionOnDemandConfigHeaders) SetXFcDate(v string) *PutFunctionOnDemandConfigHeaders {
	s.XFcDate = &v
	return s
}

func (s *PutFunctionOnDemandConfigHeaders) SetXFcInvocationType(v string) *PutFunctionOnDemandConfigHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *PutFunctionOnDemandConfigHeaders) SetXFcLogType(v string) *PutFunctionOnDemandConfigHeaders {
	s.XFcLogType = &v
	return s
}

func (s *PutFunctionOnDemandConfigHeaders) SetXFcTraceId(v string) *PutFunctionOnDemandConfigHeaders {
	s.XFcTraceId = &v
	return s
}

type PutFunctionOnDemandConfigRequest struct {
	MaximumInstanceCount *int64  `json:"maximumInstanceCount,omitempty" xml:"maximumInstanceCount,omitempty"`
	Qualifier            *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s PutFunctionOnDemandConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s PutFunctionOnDemandConfigRequest) GoString() string {
	return s.String()
}

func (s *PutFunctionOnDemandConfigRequest) SetMaximumInstanceCount(v int64) *PutFunctionOnDemandConfigRequest {
	s.MaximumInstanceCount = &v
	return s
}

func (s *PutFunctionOnDemandConfigRequest) SetQualifier(v string) *PutFunctionOnDemandConfigRequest {
	s.Qualifier = &v
	return s
}

type PutFunctionOnDemandConfigResponseBody struct {
	MaximumInstanceCount *int64  `json:"maximumInstanceCount,omitempty" xml:"maximumInstanceCount,omitempty"`
	Resource             *string `json:"resource,omitempty" xml:"resource,omitempty"`
}

func (s PutFunctionOnDemandConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutFunctionOnDemandConfigResponseBody) GoString() string {
	return s.String()
}

func (s *PutFunctionOnDemandConfigResponseBody) SetMaximumInstanceCount(v int64) *PutFunctionOnDemandConfigResponseBody {
	s.MaximumInstanceCount = &v
	return s
}

func (s *PutFunctionOnDemandConfigResponseBody) SetResource(v string) *PutFunctionOnDemandConfigResponseBody {
	s.Resource = &v
	return s
}

type PutFunctionOnDemandConfigResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PutFunctionOnDemandConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutFunctionOnDemandConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s PutFunctionOnDemandConfigResponse) GoString() string {
	return s.String()
}

func (s *PutFunctionOnDemandConfigResponse) SetHeaders(v map[string]*string) *PutFunctionOnDemandConfigResponse {
	s.Headers = v
	return s
}

func (s *PutFunctionOnDemandConfigResponse) SetBody(v *PutFunctionOnDemandConfigResponseBody) *PutFunctionOnDemandConfigResponse {
	s.Body = v
	return s
}

type PutProvisionConfigHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s PutProvisionConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s PutProvisionConfigHeaders) GoString() string {
	return s.String()
}

func (s *PutProvisionConfigHeaders) SetCommonHeaders(v map[string]*string) *PutProvisionConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PutProvisionConfigHeaders) SetXFcAccountId(v string) *PutProvisionConfigHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *PutProvisionConfigHeaders) SetXFcCodeChecksum(v string) *PutProvisionConfigHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *PutProvisionConfigHeaders) SetXFcDate(v string) *PutProvisionConfigHeaders {
	s.XFcDate = &v
	return s
}

func (s *PutProvisionConfigHeaders) SetXFcInvocationType(v string) *PutProvisionConfigHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *PutProvisionConfigHeaders) SetXFcLogType(v string) *PutProvisionConfigHeaders {
	s.XFcLogType = &v
	return s
}

func (s *PutProvisionConfigHeaders) SetXFcTraceId(v string) *PutProvisionConfigHeaders {
	s.XFcTraceId = &v
	return s
}

type PutProvisionConfigRequest struct {
	// 定时策略配置
	ScheduledActions []*ScheduledActions `json:"scheduledActions,omitempty" xml:"scheduledActions,omitempty" type:"Repeated"`
	// 预留的目标资源个数
	Target *int64 `json:"target,omitempty" xml:"target,omitempty"`
	// 指标追踪伸缩策略配置
	TargetTrackingPolicies []*TargetTrackingPolicies `json:"targetTrackingPolicies,omitempty" xml:"targetTrackingPolicies,omitempty" type:"Repeated"`
	// 别名名称
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s PutProvisionConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s PutProvisionConfigRequest) GoString() string {
	return s.String()
}

func (s *PutProvisionConfigRequest) SetScheduledActions(v []*ScheduledActions) *PutProvisionConfigRequest {
	s.ScheduledActions = v
	return s
}

func (s *PutProvisionConfigRequest) SetTarget(v int64) *PutProvisionConfigRequest {
	s.Target = &v
	return s
}

func (s *PutProvisionConfigRequest) SetTargetTrackingPolicies(v []*TargetTrackingPolicies) *PutProvisionConfigRequest {
	s.TargetTrackingPolicies = v
	return s
}

func (s *PutProvisionConfigRequest) SetQualifier(v string) *PutProvisionConfigRequest {
	s.Qualifier = &v
	return s
}

type PutProvisionConfigResponseBody struct {
	// 实际资源个数
	Current *int64 `json:"current,omitempty" xml:"current,omitempty"`
	// 资源描述
	Resource *string `json:"resource,omitempty" xml:"resource,omitempty"`
	// 定时策略配置
	ScheduledActions []*ScheduledActions `json:"scheduledActions,omitempty" xml:"scheduledActions,omitempty" type:"Repeated"`
	// 目标资源个数
	Target *int64 `json:"target,omitempty" xml:"target,omitempty"`
	// 指标追踪伸缩策略配置
	TargetTrackingPolicies []*TargetTrackingPolicies `json:"targetTrackingPolicies,omitempty" xml:"targetTrackingPolicies,omitempty" type:"Repeated"`
}

func (s PutProvisionConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutProvisionConfigResponseBody) GoString() string {
	return s.String()
}

func (s *PutProvisionConfigResponseBody) SetCurrent(v int64) *PutProvisionConfigResponseBody {
	s.Current = &v
	return s
}

func (s *PutProvisionConfigResponseBody) SetResource(v string) *PutProvisionConfigResponseBody {
	s.Resource = &v
	return s
}

func (s *PutProvisionConfigResponseBody) SetScheduledActions(v []*ScheduledActions) *PutProvisionConfigResponseBody {
	s.ScheduledActions = v
	return s
}

func (s *PutProvisionConfigResponseBody) SetTarget(v int64) *PutProvisionConfigResponseBody {
	s.Target = &v
	return s
}

func (s *PutProvisionConfigResponseBody) SetTargetTrackingPolicies(v []*TargetTrackingPolicies) *PutProvisionConfigResponseBody {
	s.TargetTrackingPolicies = v
	return s
}

type PutProvisionConfigResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PutProvisionConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutProvisionConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s PutProvisionConfigResponse) GoString() string {
	return s.String()
}

func (s *PutProvisionConfigResponse) SetHeaders(v map[string]*string) *PutProvisionConfigResponse {
	s.Headers = v
	return s
}

func (s *PutProvisionConfigResponse) SetBody(v *PutProvisionConfigResponseBody) *PutProvisionConfigResponse {
	s.Body = v
	return s
}

type RegisterEventSourceHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s RegisterEventSourceHeaders) String() string {
	return tea.Prettify(s)
}

func (s RegisterEventSourceHeaders) GoString() string {
	return s.String()
}

func (s *RegisterEventSourceHeaders) SetCommonHeaders(v map[string]*string) *RegisterEventSourceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RegisterEventSourceHeaders) SetXFcAccountId(v string) *RegisterEventSourceHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *RegisterEventSourceHeaders) SetXFcCodeChecksum(v string) *RegisterEventSourceHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *RegisterEventSourceHeaders) SetXFcDate(v string) *RegisterEventSourceHeaders {
	s.XFcDate = &v
	return s
}

func (s *RegisterEventSourceHeaders) SetXFcInvocationType(v string) *RegisterEventSourceHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *RegisterEventSourceHeaders) SetXFcLogType(v string) *RegisterEventSourceHeaders {
	s.XFcLogType = &v
	return s
}

func (s *RegisterEventSourceHeaders) SetXFcTraceId(v string) *RegisterEventSourceHeaders {
	s.XFcTraceId = &v
	return s
}

type RegisterEventSourceRequest struct {
	// 事件源资源标识符
	SourceArn *string `json:"sourceArn,omitempty" xml:"sourceArn,omitempty"`
	// 别名或版本
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s RegisterEventSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterEventSourceRequest) GoString() string {
	return s.String()
}

func (s *RegisterEventSourceRequest) SetSourceArn(v string) *RegisterEventSourceRequest {
	s.SourceArn = &v
	return s
}

func (s *RegisterEventSourceRequest) SetQualifier(v string) *RegisterEventSourceRequest {
	s.Qualifier = &v
	return s
}

type RegisterEventSourceResponseBody struct {
	// 创建时间
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// 事件源资源标识符
	SourceArn *string `json:"sourceArn,omitempty" xml:"sourceArn,omitempty"`
}

func (s RegisterEventSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterEventSourceResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterEventSourceResponseBody) SetCreatedTime(v string) *RegisterEventSourceResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *RegisterEventSourceResponseBody) SetSourceArn(v string) *RegisterEventSourceResponseBody {
	s.SourceArn = &v
	return s
}

type RegisterEventSourceResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RegisterEventSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RegisterEventSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterEventSourceResponse) GoString() string {
	return s.String()
}

func (s *RegisterEventSourceResponse) SetHeaders(v map[string]*string) *RegisterEventSourceResponse {
	s.Headers = v
	return s
}

func (s *RegisterEventSourceResponse) SetBody(v *RegisterEventSourceResponseBody) *RegisterEventSourceResponse {
	s.Body = v
	return s
}

type StopStatefulAsyncInvocationHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s StopStatefulAsyncInvocationHeaders) String() string {
	return tea.Prettify(s)
}

func (s StopStatefulAsyncInvocationHeaders) GoString() string {
	return s.String()
}

func (s *StopStatefulAsyncInvocationHeaders) SetCommonHeaders(v map[string]*string) *StopStatefulAsyncInvocationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StopStatefulAsyncInvocationHeaders) SetXFcAccountId(v string) *StopStatefulAsyncInvocationHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *StopStatefulAsyncInvocationHeaders) SetXFcCodeChecksum(v string) *StopStatefulAsyncInvocationHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *StopStatefulAsyncInvocationHeaders) SetXFcDate(v string) *StopStatefulAsyncInvocationHeaders {
	s.XFcDate = &v
	return s
}

func (s *StopStatefulAsyncInvocationHeaders) SetXFcInvocationType(v string) *StopStatefulAsyncInvocationHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *StopStatefulAsyncInvocationHeaders) SetXFcLogType(v string) *StopStatefulAsyncInvocationHeaders {
	s.XFcLogType = &v
	return s
}

func (s *StopStatefulAsyncInvocationHeaders) SetXFcTraceId(v string) *StopStatefulAsyncInvocationHeaders {
	s.XFcTraceId = &v
	return s
}

type StopStatefulAsyncInvocationRequest struct {
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s StopStatefulAsyncInvocationRequest) String() string {
	return tea.Prettify(s)
}

func (s StopStatefulAsyncInvocationRequest) GoString() string {
	return s.String()
}

func (s *StopStatefulAsyncInvocationRequest) SetQualifier(v string) *StopStatefulAsyncInvocationRequest {
	s.Qualifier = &v
	return s
}

type StopStatefulAsyncInvocationResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s StopStatefulAsyncInvocationResponse) String() string {
	return tea.Prettify(s)
}

func (s StopStatefulAsyncInvocationResponse) GoString() string {
	return s.String()
}

func (s *StopStatefulAsyncInvocationResponse) SetHeaders(v map[string]*string) *StopStatefulAsyncInvocationResponse {
	s.Headers = v
	return s
}

type TagResourceHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s TagResourceHeaders) String() string {
	return tea.Prettify(s)
}

func (s TagResourceHeaders) GoString() string {
	return s.String()
}

func (s *TagResourceHeaders) SetCommonHeaders(v map[string]*string) *TagResourceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TagResourceHeaders) SetXFcAccountId(v string) *TagResourceHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *TagResourceHeaders) SetXFcCodeChecksum(v string) *TagResourceHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *TagResourceHeaders) SetXFcDate(v string) *TagResourceHeaders {
	s.XFcDate = &v
	return s
}

func (s *TagResourceHeaders) SetXFcInvocationType(v string) *TagResourceHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *TagResourceHeaders) SetXFcLogType(v string) *TagResourceHeaders {
	s.XFcLogType = &v
	return s
}

func (s *TagResourceHeaders) SetXFcTraceId(v string) *TagResourceHeaders {
	s.XFcTraceId = &v
	return s
}

type TagResourceRequest struct {
	ResourceArn *string            `json:"resourceArn,omitempty" xml:"resourceArn,omitempty"`
	Tags        map[string]*string `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s TagResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourceRequest) GoString() string {
	return s.String()
}

func (s *TagResourceRequest) SetResourceArn(v string) *TagResourceRequest {
	s.ResourceArn = &v
	return s
}

func (s *TagResourceRequest) SetTags(v map[string]*string) *TagResourceRequest {
	s.Tags = v
	return s
}

type TagResourceResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s TagResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourceResponse) GoString() string {
	return s.String()
}

func (s *TagResourceResponse) SetHeaders(v map[string]*string) *TagResourceResponse {
	s.Headers = v
	return s
}

type UntagResourceHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s UntagResourceHeaders) String() string {
	return tea.Prettify(s)
}

func (s UntagResourceHeaders) GoString() string {
	return s.String()
}

func (s *UntagResourceHeaders) SetCommonHeaders(v map[string]*string) *UntagResourceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UntagResourceHeaders) SetXFcAccountId(v string) *UntagResourceHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *UntagResourceHeaders) SetXFcCodeChecksum(v string) *UntagResourceHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *UntagResourceHeaders) SetXFcDate(v string) *UntagResourceHeaders {
	s.XFcDate = &v
	return s
}

func (s *UntagResourceHeaders) SetXFcInvocationType(v string) *UntagResourceHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *UntagResourceHeaders) SetXFcLogType(v string) *UntagResourceHeaders {
	s.XFcLogType = &v
	return s
}

func (s *UntagResourceHeaders) SetXFcTraceId(v string) *UntagResourceHeaders {
	s.XFcTraceId = &v
	return s
}

type UntagResourceRequest struct {
	// 删除所有 tag，默认值为 false
	All *bool `json:"all,omitempty" xml:"all,omitempty"`
	// 目前只能给 top level 资源 service 进行标签的相关操作, ARN 可以是类似 services/foo 或者 acs:fc:cn-shanghai:123456789:services/foo
	ResourceArn *string `json:"resourceArn,omitempty" xml:"resourceArn,omitempty"`
	// tag key 值列表， 最大为 20，当 all=false， length 至少为 1. 当 length 大于 1 时， 可以忽略 all 值
	TagKeys []*string `json:"tagKeys	,omitempty" xml:"tagKeys	,omitempty" type:"Repeated"`
}

func (s UntagResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourceRequest) GoString() string {
	return s.String()
}

func (s *UntagResourceRequest) SetAll(v bool) *UntagResourceRequest {
	s.All = &v
	return s
}

func (s *UntagResourceRequest) SetResourceArn(v string) *UntagResourceRequest {
	s.ResourceArn = &v
	return s
}

func (s *UntagResourceRequest) SetTagKeys(v []*string) *UntagResourceRequest {
	s.TagKeys = v
	return s
}

type UntagResourceResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s UntagResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourceResponse) GoString() string {
	return s.String()
}

func (s *UntagResourceResponse) SetHeaders(v map[string]*string) *UntagResourceResponse {
	s.Headers = v
	return s
}

type UpdateAliasHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	IfMatch           *string            `json:"If-Match,omitempty" xml:"If-Match,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s UpdateAliasHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateAliasHeaders) GoString() string {
	return s.String()
}

func (s *UpdateAliasHeaders) SetCommonHeaders(v map[string]*string) *UpdateAliasHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateAliasHeaders) SetIfMatch(v string) *UpdateAliasHeaders {
	s.IfMatch = &v
	return s
}

func (s *UpdateAliasHeaders) SetXFcAccountId(v string) *UpdateAliasHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *UpdateAliasHeaders) SetXFcCodeChecksum(v string) *UpdateAliasHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *UpdateAliasHeaders) SetXFcDate(v string) *UpdateAliasHeaders {
	s.XFcDate = &v
	return s
}

func (s *UpdateAliasHeaders) SetXFcInvocationType(v string) *UpdateAliasHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *UpdateAliasHeaders) SetXFcLogType(v string) *UpdateAliasHeaders {
	s.XFcLogType = &v
	return s
}

func (s *UpdateAliasHeaders) SetXFcTraceId(v string) *UpdateAliasHeaders {
	s.XFcTraceId = &v
	return s
}

type UpdateAliasRequest struct {
	// 额外版本权重
	AdditionalVersionWeight map[string]*float32 `json:"additionalVersionWeight,omitempty" xml:"additionalVersionWeight,omitempty"`
	// 别名描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 版本ID
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s UpdateAliasRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAliasRequest) GoString() string {
	return s.String()
}

func (s *UpdateAliasRequest) SetAdditionalVersionWeight(v map[string]*float32) *UpdateAliasRequest {
	s.AdditionalVersionWeight = v
	return s
}

func (s *UpdateAliasRequest) SetDescription(v string) *UpdateAliasRequest {
	s.Description = &v
	return s
}

func (s *UpdateAliasRequest) SetVersionId(v string) *UpdateAliasRequest {
	s.VersionId = &v
	return s
}

type UpdateAliasResponseBody struct {
	// 额外版本权重
	AdditionalVersionWeight map[string]*float32 `json:"additionalVersionWeight,omitempty" xml:"additionalVersionWeight,omitempty"`
	// 别名名称
	AliasName *string `json:"aliasName,omitempty" xml:"aliasName,omitempty"`
	// 创建时间
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// 别名描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 上次更新时间
	LastModifiedTime *string `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	// 版本ID
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s UpdateAliasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAliasResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAliasResponseBody) SetAdditionalVersionWeight(v map[string]*float32) *UpdateAliasResponseBody {
	s.AdditionalVersionWeight = v
	return s
}

func (s *UpdateAliasResponseBody) SetAliasName(v string) *UpdateAliasResponseBody {
	s.AliasName = &v
	return s
}

func (s *UpdateAliasResponseBody) SetCreatedTime(v string) *UpdateAliasResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *UpdateAliasResponseBody) SetDescription(v string) *UpdateAliasResponseBody {
	s.Description = &v
	return s
}

func (s *UpdateAliasResponseBody) SetLastModifiedTime(v string) *UpdateAliasResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *UpdateAliasResponseBody) SetVersionId(v string) *UpdateAliasResponseBody {
	s.VersionId = &v
	return s
}

type UpdateAliasResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAliasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAliasResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAliasResponse) GoString() string {
	return s.String()
}

func (s *UpdateAliasResponse) SetHeaders(v map[string]*string) *UpdateAliasResponse {
	s.Headers = v
	return s
}

func (s *UpdateAliasResponse) SetBody(v *UpdateAliasResponseBody) *UpdateAliasResponse {
	s.Body = v
	return s
}

type UpdateCustomDomainHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XFcAccountId      *string            `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string            `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string            `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string            `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string            `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string            `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s UpdateCustomDomainHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomDomainHeaders) GoString() string {
	return s.String()
}

func (s *UpdateCustomDomainHeaders) SetCommonHeaders(v map[string]*string) *UpdateCustomDomainHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateCustomDomainHeaders) SetXFcAccountId(v string) *UpdateCustomDomainHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *UpdateCustomDomainHeaders) SetXFcCodeChecksum(v string) *UpdateCustomDomainHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *UpdateCustomDomainHeaders) SetXFcDate(v string) *UpdateCustomDomainHeaders {
	s.XFcDate = &v
	return s
}

func (s *UpdateCustomDomainHeaders) SetXFcInvocationType(v string) *UpdateCustomDomainHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *UpdateCustomDomainHeaders) SetXFcLogType(v string) *UpdateCustomDomainHeaders {
	s.XFcLogType = &v
	return s
}

func (s *UpdateCustomDomainHeaders) SetXFcTraceId(v string) *UpdateCustomDomainHeaders {
	s.XFcTraceId = &v
	return s
}

type UpdateCustomDomainRequest struct {
	CertConfig  *CertConfig  `json:"certConfig,omitempty" xml:"certConfig,omitempty"`
	Protocol    *string      `json:"protocol,omitempty" xml:"protocol,omitempty"`
	RouteConfig *RouteConfig `json:"routeConfig,omitempty" xml:"routeConfig,omitempty"`
}

func (s UpdateCustomDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomDomainRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomDomainRequest) SetCertConfig(v *CertConfig) *UpdateCustomDomainRequest {
	s.CertConfig = v
	return s
}

func (s *UpdateCustomDomainRequest) SetProtocol(v string) *UpdateCustomDomainRequest {
	s.Protocol = &v
	return s
}

func (s *UpdateCustomDomainRequest) SetRouteConfig(v *RouteConfig) *UpdateCustomDomainRequest {
	s.RouteConfig = v
	return s
}

type UpdateCustomDomainResponseBody struct {
	AccountId        *string      `json:"accountId,omitempty" xml:"accountId,omitempty"`
	ApiVersion       *string      `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
	CertConfig       *CertConfig  `json:"certConfig,omitempty" xml:"certConfig,omitempty"`
	CreatedTime      *string      `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	DomainName       *string      `json:"domainName,omitempty" xml:"domainName,omitempty"`
	LastModifiedTime *string      `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	Protocol         *string      `json:"protocol,omitempty" xml:"protocol,omitempty"`
	RouteConfig      *RouteConfig `json:"routeConfig,omitempty" xml:"routeConfig,omitempty"`
}

func (s UpdateCustomDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomDomainResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomDomainResponseBody) SetAccountId(v string) *UpdateCustomDomainResponseBody {
	s.AccountId = &v
	return s
}

func (s *UpdateCustomDomainResponseBody) SetApiVersion(v string) *UpdateCustomDomainResponseBody {
	s.ApiVersion = &v
	return s
}

func (s *UpdateCustomDomainResponseBody) SetCertConfig(v *CertConfig) *UpdateCustomDomainResponseBody {
	s.CertConfig = v
	return s
}

func (s *UpdateCustomDomainResponseBody) SetCreatedTime(v string) *UpdateCustomDomainResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *UpdateCustomDomainResponseBody) SetDomainName(v string) *UpdateCustomDomainResponseBody {
	s.DomainName = &v
	return s
}

func (s *UpdateCustomDomainResponseBody) SetLastModifiedTime(v string) *UpdateCustomDomainResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *UpdateCustomDomainResponseBody) SetProtocol(v string) *UpdateCustomDomainResponseBody {
	s.Protocol = &v
	return s
}

func (s *UpdateCustomDomainResponseBody) SetRouteConfig(v *RouteConfig) *UpdateCustomDomainResponseBody {
	s.RouteConfig = v
	return s
}

type UpdateCustomDomainResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateCustomDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCustomDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomDomainResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomDomainResponse) SetHeaders(v map[string]*string) *UpdateCustomDomainResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomDomainResponse) SetBody(v *UpdateCustomDomainResponseBody) *UpdateCustomDomainResponse {
	s.Body = v
	return s
}

type UpdateFunctionHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// 用于确保实际更改的资源和期望更改的资源是一致的，该值来自Create，Get和Update API的响应
	IfMatch           *string `json:"If-Match,omitempty" xml:"If-Match,omitempty"`
	XFcAccountId      *string `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s UpdateFunctionHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionHeaders) GoString() string {
	return s.String()
}

func (s *UpdateFunctionHeaders) SetCommonHeaders(v map[string]*string) *UpdateFunctionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateFunctionHeaders) SetIfMatch(v string) *UpdateFunctionHeaders {
	s.IfMatch = &v
	return s
}

func (s *UpdateFunctionHeaders) SetXFcAccountId(v string) *UpdateFunctionHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *UpdateFunctionHeaders) SetXFcCodeChecksum(v string) *UpdateFunctionHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *UpdateFunctionHeaders) SetXFcDate(v string) *UpdateFunctionHeaders {
	s.XFcDate = &v
	return s
}

func (s *UpdateFunctionHeaders) SetXFcInvocationType(v string) *UpdateFunctionHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *UpdateFunctionHeaders) SetXFcLogType(v string) *UpdateFunctionHeaders {
	s.XFcLogType = &v
	return s
}

func (s *UpdateFunctionHeaders) SetXFcTraceId(v string) *UpdateFunctionHeaders {
	s.XFcTraceId = &v
	return s
}

type UpdateFunctionRequest struct {
	InstanceConcurrency *int32 `json:"InstanceConcurrency,omitempty" xml:"InstanceConcurrency,omitempty"`
	// 自定义、自定义容器运行时 HTTP Server 的监听端口
	CaPort                *int32                 `json:"caPort,omitempty" xml:"caPort,omitempty"`
	Code                  *Code                  `json:"code,omitempty" xml:"code,omitempty"`
	CustomContainerConfig *CustomContainerConfig `json:"customContainerConfig,omitempty" xml:"customContainerConfig,omitempty"`
	// 函数自定义DNS配置
	CustomDNS *CustomDNS `json:"customDNS,omitempty" xml:"customDNS,omitempty"`
	// Custom Runtime函数详细配置
	CustomRuntimeConfig *CustomRuntimeConfig `json:"customRuntimeConfig,omitempty" xml:"customRuntimeConfig,omitempty"`
	// 函数描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 为函数设置的环境变量，可以在函数中获取环境变量的值
	EnvironmentVariables map[string]*string `json:"environmentVariables,omitempty" xml:"environmentVariables,omitempty"`
	// function执行的入口，具体格式和语言相关
	Handler *string `json:"handler,omitempty" xml:"handler,omitempty"`
	// 初始化function运行的超时时间，单位为秒，最小1秒，默认3秒。初始化function超过这个时间后会被终止执行
	InitializationTimeout *int32 `json:"initializationTimeout,omitempty" xml:"initializationTimeout,omitempty"`
	// 初始化 function 执行的入口，具体格式和语言相关
	Initializer             *string                  `json:"initializer,omitempty" xml:"initializer,omitempty"`
	InstanceLifecycleConfig *InstanceLifecycleConfig `json:"instanceLifecycleConfig,omitempty" xml:"instanceLifecycleConfig,omitempty"`
	InstanceType            *string                  `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	Layers                  []*string                `json:"layers,omitempty" xml:"layers,omitempty" type:"Repeated"`
	// function的内存规格，单位为MB，为64MB的倍数
	MemorySize *int32 `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	// runtime
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// function运行的超时时间，单位为秒，最小1秒，默认3秒。function超过这个时间后会被终止执行
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
}

func (s UpdateFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionRequest) GoString() string {
	return s.String()
}

func (s *UpdateFunctionRequest) SetInstanceConcurrency(v int32) *UpdateFunctionRequest {
	s.InstanceConcurrency = &v
	return s
}

func (s *UpdateFunctionRequest) SetCaPort(v int32) *UpdateFunctionRequest {
	s.CaPort = &v
	return s
}

func (s *UpdateFunctionRequest) SetCode(v *Code) *UpdateFunctionRequest {
	s.Code = v
	return s
}

func (s *UpdateFunctionRequest) SetCustomContainerConfig(v *CustomContainerConfig) *UpdateFunctionRequest {
	s.CustomContainerConfig = v
	return s
}

func (s *UpdateFunctionRequest) SetCustomDNS(v *CustomDNS) *UpdateFunctionRequest {
	s.CustomDNS = v
	return s
}

func (s *UpdateFunctionRequest) SetCustomRuntimeConfig(v *CustomRuntimeConfig) *UpdateFunctionRequest {
	s.CustomRuntimeConfig = v
	return s
}

func (s *UpdateFunctionRequest) SetDescription(v string) *UpdateFunctionRequest {
	s.Description = &v
	return s
}

func (s *UpdateFunctionRequest) SetEnvironmentVariables(v map[string]*string) *UpdateFunctionRequest {
	s.EnvironmentVariables = v
	return s
}

func (s *UpdateFunctionRequest) SetHandler(v string) *UpdateFunctionRequest {
	s.Handler = &v
	return s
}

func (s *UpdateFunctionRequest) SetInitializationTimeout(v int32) *UpdateFunctionRequest {
	s.InitializationTimeout = &v
	return s
}

func (s *UpdateFunctionRequest) SetInitializer(v string) *UpdateFunctionRequest {
	s.Initializer = &v
	return s
}

func (s *UpdateFunctionRequest) SetInstanceLifecycleConfig(v *InstanceLifecycleConfig) *UpdateFunctionRequest {
	s.InstanceLifecycleConfig = v
	return s
}

func (s *UpdateFunctionRequest) SetInstanceType(v string) *UpdateFunctionRequest {
	s.InstanceType = &v
	return s
}

func (s *UpdateFunctionRequest) SetLayers(v []*string) *UpdateFunctionRequest {
	s.Layers = v
	return s
}

func (s *UpdateFunctionRequest) SetMemorySize(v int32) *UpdateFunctionRequest {
	s.MemorySize = &v
	return s
}

func (s *UpdateFunctionRequest) SetRuntime(v string) *UpdateFunctionRequest {
	s.Runtime = &v
	return s
}

func (s *UpdateFunctionRequest) SetTimeout(v int32) *UpdateFunctionRequest {
	s.Timeout = &v
	return s
}

type UpdateFunctionResponseBody struct {
	// 自定义、自定义容器运行时 HTTP Server 的监听端口
	CaPort *int32 `json:"caPort,omitempty" xml:"caPort,omitempty"`
	// function code包的CRC64值
	CodeChecksum *string `json:"codeChecksum,omitempty" xml:"codeChecksum,omitempty"`
	// 系统返回的function的code包大小，单位为byte Example : 1024
	CodeSize *int64 `json:"codeSize,omitempty" xml:"codeSize,omitempty"`
	// function创建时间
	CreatedTime           *string                `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	CustomContainerConfig *CustomContainerConfig `json:"customContainerConfig,omitempty" xml:"customContainerConfig,omitempty"`
	// 函数自定义DNS配置
	CustomDNS *CustomDNS `json:"customDNS,omitempty" xml:"customDNS,omitempty"`
	// Custom Runtime函数详细配置
	CustomRuntimeConfig *CustomRuntimeConfig `json:"customRuntimeConfig,omitempty" xml:"customRuntimeConfig,omitempty"`
	// 函数描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 为函数设置的环境变量，可以在函数中获取环境变量的值
	EnvironmentVariables map[string]*string `json:"environmentVariables,omitempty" xml:"environmentVariables,omitempty"`
	// 系统为每个function生成的唯一ID
	FunctionId *string `json:"functionId,omitempty" xml:"functionId,omitempty"`
	// 函数名称
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	// function的执行入口
	Handler *string `json:"handler,omitempty" xml:"handler,omitempty"`
	// 初始化function运行的超时时间，单位为秒，最小1秒，默认3秒。初始化function超过这个时间后会被终止执行
	InitializationTimeout *int32 `json:"initializationTimeout,omitempty" xml:"initializationTimeout,omitempty"`
	// 初始化 function 执行的入口，具体格式和语言相关
	Initializer             *string                  `json:"initializer,omitempty" xml:"initializer,omitempty"`
	InstanceLifecycleConfig *InstanceLifecycleConfig `json:"instanceLifecycleConfig,omitempty" xml:"instanceLifecycleConfig,omitempty"`
	InstanceType            *string                  `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	// function上次修改时间
	LastModifiedTime *string   `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	Layers           []*string `json:"layers,omitempty" xml:"layers,omitempty" type:"Repeated"`
	// function设置的内存大小，单位为MB
	MemorySize *int32 `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	// function运行的语言环境，目前支持nodejs6, nodejs8, python2.7, python3, java8
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// 运行的超时时间，单位为秒
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
}

func (s UpdateFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFunctionResponseBody) SetCaPort(v int32) *UpdateFunctionResponseBody {
	s.CaPort = &v
	return s
}

func (s *UpdateFunctionResponseBody) SetCodeChecksum(v string) *UpdateFunctionResponseBody {
	s.CodeChecksum = &v
	return s
}

func (s *UpdateFunctionResponseBody) SetCodeSize(v int64) *UpdateFunctionResponseBody {
	s.CodeSize = &v
	return s
}

func (s *UpdateFunctionResponseBody) SetCreatedTime(v string) *UpdateFunctionResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *UpdateFunctionResponseBody) SetCustomContainerConfig(v *CustomContainerConfig) *UpdateFunctionResponseBody {
	s.CustomContainerConfig = v
	return s
}

func (s *UpdateFunctionResponseBody) SetCustomDNS(v *CustomDNS) *UpdateFunctionResponseBody {
	s.CustomDNS = v
	return s
}

func (s *UpdateFunctionResponseBody) SetCustomRuntimeConfig(v *CustomRuntimeConfig) *UpdateFunctionResponseBody {
	s.CustomRuntimeConfig = v
	return s
}

func (s *UpdateFunctionResponseBody) SetDescription(v string) *UpdateFunctionResponseBody {
	s.Description = &v
	return s
}

func (s *UpdateFunctionResponseBody) SetEnvironmentVariables(v map[string]*string) *UpdateFunctionResponseBody {
	s.EnvironmentVariables = v
	return s
}

func (s *UpdateFunctionResponseBody) SetFunctionId(v string) *UpdateFunctionResponseBody {
	s.FunctionId = &v
	return s
}

func (s *UpdateFunctionResponseBody) SetFunctionName(v string) *UpdateFunctionResponseBody {
	s.FunctionName = &v
	return s
}

func (s *UpdateFunctionResponseBody) SetHandler(v string) *UpdateFunctionResponseBody {
	s.Handler = &v
	return s
}

func (s *UpdateFunctionResponseBody) SetInitializationTimeout(v int32) *UpdateFunctionResponseBody {
	s.InitializationTimeout = &v
	return s
}

func (s *UpdateFunctionResponseBody) SetInitializer(v string) *UpdateFunctionResponseBody {
	s.Initializer = &v
	return s
}

func (s *UpdateFunctionResponseBody) SetInstanceLifecycleConfig(v *InstanceLifecycleConfig) *UpdateFunctionResponseBody {
	s.InstanceLifecycleConfig = v
	return s
}

func (s *UpdateFunctionResponseBody) SetInstanceType(v string) *UpdateFunctionResponseBody {
	s.InstanceType = &v
	return s
}

func (s *UpdateFunctionResponseBody) SetLastModifiedTime(v string) *UpdateFunctionResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *UpdateFunctionResponseBody) SetLayers(v []*string) *UpdateFunctionResponseBody {
	s.Layers = v
	return s
}

func (s *UpdateFunctionResponseBody) SetMemorySize(v int32) *UpdateFunctionResponseBody {
	s.MemorySize = &v
	return s
}

func (s *UpdateFunctionResponseBody) SetRuntime(v string) *UpdateFunctionResponseBody {
	s.Runtime = &v
	return s
}

func (s *UpdateFunctionResponseBody) SetTimeout(v int32) *UpdateFunctionResponseBody {
	s.Timeout = &v
	return s
}

type UpdateFunctionResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionResponse) GoString() string {
	return s.String()
}

func (s *UpdateFunctionResponse) SetHeaders(v map[string]*string) *UpdateFunctionResponse {
	s.Headers = v
	return s
}

func (s *UpdateFunctionResponse) SetBody(v *UpdateFunctionResponseBody) *UpdateFunctionResponse {
	s.Body = v
	return s
}

type UpdateServiceHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// 用于确保实际更改的资源和期望更改的资源是一致的，该值来自Create，Get和Update API的响应
	IfMatch           *string `json:"If-Match,omitempty" xml:"If-Match,omitempty"`
	XFcAccountId      *string `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s UpdateServiceHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceHeaders) GoString() string {
	return s.String()
}

func (s *UpdateServiceHeaders) SetCommonHeaders(v map[string]*string) *UpdateServiceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateServiceHeaders) SetIfMatch(v string) *UpdateServiceHeaders {
	s.IfMatch = &v
	return s
}

func (s *UpdateServiceHeaders) SetXFcAccountId(v string) *UpdateServiceHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *UpdateServiceHeaders) SetXFcCodeChecksum(v string) *UpdateServiceHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *UpdateServiceHeaders) SetXFcDate(v string) *UpdateServiceHeaders {
	s.XFcDate = &v
	return s
}

func (s *UpdateServiceHeaders) SetXFcInvocationType(v string) *UpdateServiceHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *UpdateServiceHeaders) SetXFcLogType(v string) *UpdateServiceHeaders {
	s.XFcLogType = &v
	return s
}

func (s *UpdateServiceHeaders) SetXFcTraceId(v string) *UpdateServiceHeaders {
	s.XFcTraceId = &v
	return s
}

type UpdateServiceRequest struct {
	// 服务描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 公网访问设置
	InternetAccess *bool      `json:"internetAccess,omitempty" xml:"internetAccess,omitempty"`
	LogConfig      *LogConfig `json:"logConfig,omitempty" xml:"logConfig,omitempty"`
	NasConfig      *NASConfig `json:"nasConfig,omitempty" xml:"nasConfig,omitempty"`
	// 服务角色
	Role          *string        `json:"role,omitempty" xml:"role,omitempty"`
	TracingConfig *TracingConfig `json:"tracingConfig,omitempty" xml:"tracingConfig,omitempty"`
	VpcConfig     *VPCConfig     `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty"`
}

func (s UpdateServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequest) SetDescription(v string) *UpdateServiceRequest {
	s.Description = &v
	return s
}

func (s *UpdateServiceRequest) SetInternetAccess(v bool) *UpdateServiceRequest {
	s.InternetAccess = &v
	return s
}

func (s *UpdateServiceRequest) SetLogConfig(v *LogConfig) *UpdateServiceRequest {
	s.LogConfig = v
	return s
}

func (s *UpdateServiceRequest) SetNasConfig(v *NASConfig) *UpdateServiceRequest {
	s.NasConfig = v
	return s
}

func (s *UpdateServiceRequest) SetRole(v string) *UpdateServiceRequest {
	s.Role = &v
	return s
}

func (s *UpdateServiceRequest) SetTracingConfig(v *TracingConfig) *UpdateServiceRequest {
	s.TracingConfig = v
	return s
}

func (s *UpdateServiceRequest) SetVpcConfig(v *VPCConfig) *UpdateServiceRequest {
	s.VpcConfig = v
	return s
}

type UpdateServiceResponseBody struct {
	// 创建时间
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// 服务描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 公网访问设置
	InternetAccess *bool `json:"internetAccess,omitempty" xml:"internetAccess,omitempty"`
	// 上次更新时间
	LastModifiedTime *string    `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	LogConfig        *LogConfig `json:"logConfig,omitempty" xml:"logConfig,omitempty"`
	NasConfig        *NASConfig `json:"nasConfig,omitempty" xml:"nasConfig,omitempty"`
	// 服务角色
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// 服务ID
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// 服务名称
	ServiceName   *string        `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	TracingConfig *TracingConfig `json:"tracingConfig,omitempty" xml:"tracingConfig,omitempty"`
	VpcConfig     *VPCConfig     `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty"`
}

func (s UpdateServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceResponseBody) SetCreatedTime(v string) *UpdateServiceResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *UpdateServiceResponseBody) SetDescription(v string) *UpdateServiceResponseBody {
	s.Description = &v
	return s
}

func (s *UpdateServiceResponseBody) SetInternetAccess(v bool) *UpdateServiceResponseBody {
	s.InternetAccess = &v
	return s
}

func (s *UpdateServiceResponseBody) SetLastModifiedTime(v string) *UpdateServiceResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *UpdateServiceResponseBody) SetLogConfig(v *LogConfig) *UpdateServiceResponseBody {
	s.LogConfig = v
	return s
}

func (s *UpdateServiceResponseBody) SetNasConfig(v *NASConfig) *UpdateServiceResponseBody {
	s.NasConfig = v
	return s
}

func (s *UpdateServiceResponseBody) SetRole(v string) *UpdateServiceResponseBody {
	s.Role = &v
	return s
}

func (s *UpdateServiceResponseBody) SetServiceId(v string) *UpdateServiceResponseBody {
	s.ServiceId = &v
	return s
}

func (s *UpdateServiceResponseBody) SetServiceName(v string) *UpdateServiceResponseBody {
	s.ServiceName = &v
	return s
}

func (s *UpdateServiceResponseBody) SetTracingConfig(v *TracingConfig) *UpdateServiceResponseBody {
	s.TracingConfig = v
	return s
}

func (s *UpdateServiceResponseBody) SetVpcConfig(v *VPCConfig) *UpdateServiceResponseBody {
	s.VpcConfig = v
	return s
}

type UpdateServiceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceResponse) SetHeaders(v map[string]*string) *UpdateServiceResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceResponse) SetBody(v *UpdateServiceResponseBody) *UpdateServiceResponse {
	s.Body = v
	return s
}

type UpdateTriggerHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// 用于确保实际更改的资源和期望更改的资源是一致的，该值来自Create，Get和Update API的响应
	IfMatch           *string `json:"If-Match,omitempty" xml:"If-Match,omitempty"`
	XFcAccountId      *string `json:"X-Fc-Account-Id,omitempty" xml:"X-Fc-Account-Id,omitempty"`
	XFcCodeChecksum   *string `json:"X-Fc-Code-Checksum,omitempty" xml:"X-Fc-Code-Checksum,omitempty"`
	XFcDate           *string `json:"X-Fc-Date,omitempty" xml:"X-Fc-Date,omitempty"`
	XFcInvocationType *string `json:"X-Fc-Invocation-Type,omitempty" xml:"X-Fc-Invocation-Type,omitempty"`
	XFcLogType        *string `json:"X-Fc-Log-Type,omitempty" xml:"X-Fc-Log-Type,omitempty"`
	XFcTraceId        *string `json:"X-Fc-Trace-Id,omitempty" xml:"X-Fc-Trace-Id,omitempty"`
}

func (s UpdateTriggerHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateTriggerHeaders) GoString() string {
	return s.String()
}

func (s *UpdateTriggerHeaders) SetCommonHeaders(v map[string]*string) *UpdateTriggerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateTriggerHeaders) SetIfMatch(v string) *UpdateTriggerHeaders {
	s.IfMatch = &v
	return s
}

func (s *UpdateTriggerHeaders) SetXFcAccountId(v string) *UpdateTriggerHeaders {
	s.XFcAccountId = &v
	return s
}

func (s *UpdateTriggerHeaders) SetXFcCodeChecksum(v string) *UpdateTriggerHeaders {
	s.XFcCodeChecksum = &v
	return s
}

func (s *UpdateTriggerHeaders) SetXFcDate(v string) *UpdateTriggerHeaders {
	s.XFcDate = &v
	return s
}

func (s *UpdateTriggerHeaders) SetXFcInvocationType(v string) *UpdateTriggerHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *UpdateTriggerHeaders) SetXFcLogType(v string) *UpdateTriggerHeaders {
	s.XFcLogType = &v
	return s
}

func (s *UpdateTriggerHeaders) SetXFcTraceId(v string) *UpdateTriggerHeaders {
	s.XFcTraceId = &v
	return s
}

type UpdateTriggerRequest struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// event source，如OSS，使用该role去invoke function
	InvocationRole *string `json:"invocationRole,omitempty" xml:"invocationRole,omitempty"`
	// service版本
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	// trigger配置，针对不同的trigger类型，trigger配置会有所不同
	TriggerConfig *string `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty"`
}

func (s UpdateTriggerRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTriggerRequest) GoString() string {
	return s.String()
}

func (s *UpdateTriggerRequest) SetDescription(v string) *UpdateTriggerRequest {
	s.Description = &v
	return s
}

func (s *UpdateTriggerRequest) SetInvocationRole(v string) *UpdateTriggerRequest {
	s.InvocationRole = &v
	return s
}

func (s *UpdateTriggerRequest) SetQualifier(v string) *UpdateTriggerRequest {
	s.Qualifier = &v
	return s
}

func (s *UpdateTriggerRequest) SetTriggerConfig(v string) *UpdateTriggerRequest {
	s.TriggerConfig = &v
	return s
}

type UpdateTriggerResponseBody struct {
	// 创建时间
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 域名名称，使用域名名称拼接上函数计算域名，可以采用HTTP协议调用到触发器对应版本的函数。例如{domainName}.cn-shanghai.fc.aliyuncs.com
	DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
	// 调用函数使用的RAM角色的ARN
	InvocationRole *string `json:"invocationRole,omitempty" xml:"invocationRole,omitempty"`
	// 上次修改时间
	LastModifiedTime *string `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	// service版本
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	// event source的Aliyun Resource Name（ARN
	SourceArn *string `json:"sourceArn,omitempty" xml:"sourceArn,omitempty"`
	// trigger配置对象
	TriggerConfig *string `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty"`
	TriggerId     *string `json:"triggerId,omitempty" xml:"triggerId,omitempty"`
	// trigger名称
	TriggerName *string `json:"triggerName,omitempty" xml:"triggerName,omitempty"`
	// trigger类型，如 oss, log, tablestore, timer, http, cdn_events, mns_topic
	TriggerType *string `json:"triggerType,omitempty" xml:"triggerType,omitempty"`
}

func (s UpdateTriggerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTriggerResponseBody) SetCreatedTime(v string) *UpdateTriggerResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *UpdateTriggerResponseBody) SetDescription(v string) *UpdateTriggerResponseBody {
	s.Description = &v
	return s
}

func (s *UpdateTriggerResponseBody) SetDomainName(v string) *UpdateTriggerResponseBody {
	s.DomainName = &v
	return s
}

func (s *UpdateTriggerResponseBody) SetInvocationRole(v string) *UpdateTriggerResponseBody {
	s.InvocationRole = &v
	return s
}

func (s *UpdateTriggerResponseBody) SetLastModifiedTime(v string) *UpdateTriggerResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *UpdateTriggerResponseBody) SetQualifier(v string) *UpdateTriggerResponseBody {
	s.Qualifier = &v
	return s
}

func (s *UpdateTriggerResponseBody) SetSourceArn(v string) *UpdateTriggerResponseBody {
	s.SourceArn = &v
	return s
}

func (s *UpdateTriggerResponseBody) SetTriggerConfig(v string) *UpdateTriggerResponseBody {
	s.TriggerConfig = &v
	return s
}

func (s *UpdateTriggerResponseBody) SetTriggerId(v string) *UpdateTriggerResponseBody {
	s.TriggerId = &v
	return s
}

func (s *UpdateTriggerResponseBody) SetTriggerName(v string) *UpdateTriggerResponseBody {
	s.TriggerName = &v
	return s
}

func (s *UpdateTriggerResponseBody) SetTriggerType(v string) *UpdateTriggerResponseBody {
	s.TriggerType = &v
	return s
}

type UpdateTriggerResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTriggerResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTriggerResponse) GoString() string {
	return s.String()
}

func (s *UpdateTriggerResponse) SetHeaders(v map[string]*string) *UpdateTriggerResponse {
	s.Headers = v
	return s
}

func (s *UpdateTriggerResponse) SetBody(v *UpdateTriggerResponseBody) *UpdateTriggerResponse {
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
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-1":      tea.String("account-id.ap-northeast-1.fc.aliyuncs.com"),
		"ap-south-1":          tea.String("account-id.ap-south-1.fc.aliyuncs.com"),
		"ap-southeast-1":      tea.String("account-id.ap-southeast-1.fc.aliyuncs.com"),
		"ap-southeast-2":      tea.String("account-id.ap-southeast-2.fc.aliyuncs.com"),
		"ap-southeast-3":      tea.String("account-id.ap-southeast-3.fc.aliyuncs.com"),
		"ap-southeast-5":      tea.String("account-id.ap-southeast-5.fc.aliyuncs.com"),
		"cn-beijing":          tea.String("account-id.cn-beijing.fc.aliyuncs.com"),
		"cn-chengdu":          tea.String("account-id.cn-chengdu.fc.aliyuncs.com"),
		"cn-hangzhou":         tea.String("account-id.cn-hangzhou.fc.aliyuncs.com"),
		"cn-hangzhou-finance": tea.String("account-id.cn-hangzhou-finance.fc.aliyuncs.com"),
		"cn-hongkong":         tea.String("account-id.cn-hongkong.fc.aliyuncs.com"),
		"cn-huhehaote":        tea.String("account-id.cn-huhehaote.fc.aliyuncs.com"),
		"cn-north-2-gov-1":    tea.String("account-id.cn-north-2-gov-1.fc.aliyuncs.com"),
		"cn-qingdao":          tea.String("account-id.cn-qingdao.fc.aliyuncs.com"),
		"cn-shanghai":         tea.String("account-id.cn-shanghai.fc.aliyuncs.com"),
		"cn-shenzhen":         tea.String("account-id.cn-shenzhen.fc.aliyuncs.com"),
		"cn-zhangjiakou":      tea.String("account-id.cn-zhangjiakou.fc.aliyuncs.com"),
		"eu-central-1":        tea.String("account-id.eu-central-1.fc.aliyuncs.com"),
		"eu-west-1":           tea.String("account-id.eu-west-1.fc.aliyuncs.com"),
		"us-east-1":           tea.String("account-id.us-east-1.fc.aliyuncs.com"),
		"us-west-1":           tea.String("account-id.us-west-1.fc.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("fc-open"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateAlias(serviceName *string, request *CreateAliasRequest) (_result *CreateAliasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateAliasHeaders{}
	_result = &CreateAliasResponse{}
	_body, _err := client.CreateAliasWithOptions(serviceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAliasWithOptions(serviceName *string, request *CreateAliasRequest, headers *CreateAliasHeaders, runtime *util.RuntimeOptions) (_result *CreateAliasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	serviceName = openapiutil.GetEncodeParam(serviceName)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdditionalVersionWeight)) {
		body["additionalVersionWeight"] = request.AdditionalVersionWeight
	}

	if !tea.BoolValue(util.IsUnset(request.AliasName)) {
		body["aliasName"] = request.AliasName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		body["versionId"] = request.VersionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAlias"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName) + "/aliases"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAliasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCustomDomain(request *CreateCustomDomainRequest) (_result *CreateCustomDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateCustomDomainHeaders{}
	_result = &CreateCustomDomainResponse{}
	_body, _err := client.CreateCustomDomainWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCustomDomainWithOptions(request *CreateCustomDomainRequest, headers *CreateCustomDomainHeaders, runtime *util.RuntimeOptions) (_result *CreateCustomDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.CertConfig))) {
		body["certConfig"] = request.CertConfig
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		body["domainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Protocol)) {
		body["protocol"] = request.Protocol
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.RouteConfig))) {
		body["routeConfig"] = request.RouteConfig
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCustomDomain"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/custom-domains"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCustomDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFunction(serviceName *string, request *CreateFunctionRequest) (_result *CreateFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateFunctionHeaders{}
	_result = &CreateFunctionResponse{}
	_body, _err := client.CreateFunctionWithOptions(serviceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFunctionWithOptions(serviceName *string, request *CreateFunctionRequest, headers *CreateFunctionHeaders, runtime *util.RuntimeOptions) (_result *CreateFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	serviceName = openapiutil.GetEncodeParam(serviceName)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CaPort)) {
		body["caPort"] = request.CaPort
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Code))) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.CustomContainerConfig))) {
		body["customContainerConfig"] = request.CustomContainerConfig
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.CustomDNS))) {
		body["customDNS"] = request.CustomDNS
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.CustomRuntimeConfig))) {
		body["customRuntimeConfig"] = request.CustomRuntimeConfig
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EnvironmentVariables)) {
		body["environmentVariables"] = request.EnvironmentVariables
	}

	if !tea.BoolValue(util.IsUnset(request.FunctionName)) {
		body["functionName"] = request.FunctionName
	}

	if !tea.BoolValue(util.IsUnset(request.Handler)) {
		body["handler"] = request.Handler
	}

	if !tea.BoolValue(util.IsUnset(request.InitializationTimeout)) {
		body["initializationTimeout"] = request.InitializationTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.Initializer)) {
		body["initializer"] = request.Initializer
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceConcurrency)) {
		body["instanceConcurrency"] = request.InstanceConcurrency
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.InstanceLifecycleConfig))) {
		body["instanceLifecycleConfig"] = request.InstanceLifecycleConfig
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		body["instanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.Layers)) {
		body["layers"] = request.Layers
	}

	if !tea.BoolValue(util.IsUnset(request.MemorySize)) {
		body["memorySize"] = request.MemorySize
	}

	if !tea.BoolValue(util.IsUnset(request.Runtime)) {
		body["runtime"] = request.Runtime
	}

	if !tea.BoolValue(util.IsUnset(request.Timeout)) {
		body["timeout"] = request.Timeout
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFunction"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName) + "/functions"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLayerVersion(layerName *string, request *CreateLayerVersionRequest) (_result *CreateLayerVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateLayerVersionHeaders{}
	_result = &CreateLayerVersionResponse{}
	_body, _err := client.CreateLayerVersionWithOptions(layerName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLayerVersionWithOptions(layerName *string, request *CreateLayerVersionRequest, headers *CreateLayerVersionHeaders, runtime *util.RuntimeOptions) (_result *CreateLayerVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	layerName = openapiutil.GetEncodeParam(layerName)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Code))) {
		body["Code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.CompatibleRuntime)) {
		body["compatibleRuntime"] = request.CompatibleRuntime
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLayerVersion"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/layers/" + tea.StringValue(layerName) + "/versions"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLayerVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateService(request *CreateServiceRequest) (_result *CreateServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateServiceHeaders{}
	_result = &CreateServiceResponse{}
	_body, _err := client.CreateServiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateServiceWithOptions(request *CreateServiceRequest, headers *CreateServiceHeaders, runtime *util.RuntimeOptions) (_result *CreateServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InternetAccess)) {
		body["internetAccess"] = request.InternetAccess
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.LogConfig))) {
		body["logConfig"] = request.LogConfig
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.NasConfig))) {
		body["nasConfig"] = request.NasConfig
	}

	if !tea.BoolValue(util.IsUnset(request.Role)) {
		body["role"] = request.Role
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		body["serviceName"] = request.ServiceName
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.TracingConfig))) {
		body["tracingConfig"] = request.TracingConfig
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.VpcConfig))) {
		body["vpcConfig"] = request.VpcConfig
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateService"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTrigger(serviceName *string, functionName *string, request *CreateTriggerRequest) (_result *CreateTriggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateTriggerHeaders{}
	_result = &CreateTriggerResponse{}
	_body, _err := client.CreateTriggerWithOptions(serviceName, functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTriggerWithOptions(serviceName *string, functionName *string, request *CreateTriggerRequest, headers *CreateTriggerHeaders, runtime *util.RuntimeOptions) (_result *CreateTriggerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	serviceName = openapiutil.GetEncodeParam(serviceName)
	functionName = openapiutil.GetEncodeParam(functionName)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InvocationRole)) {
		body["invocationRole"] = request.InvocationRole
	}

	if !tea.BoolValue(util.IsUnset(request.Qualifier)) {
		body["qualifier"] = request.Qualifier
	}

	if !tea.BoolValue(util.IsUnset(request.SourceArn)) {
		body["sourceArn"] = request.SourceArn
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerConfig)) {
		body["triggerConfig"] = request.TriggerConfig
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerName)) {
		body["triggerName"] = request.TriggerName
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerType)) {
		body["triggerType"] = request.TriggerType
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTrigger"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName) + "/functions/" + tea.StringValue(functionName) + "/triggers"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTriggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVpcBinding(serviceName *string, request *CreateVpcBindingRequest) (_result *CreateVpcBindingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateVpcBindingHeaders{}
	_result = &CreateVpcBindingResponse{}
	_body, _err := client.CreateVpcBindingWithOptions(serviceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVpcBindingWithOptions(serviceName *string, request *CreateVpcBindingRequest, headers *CreateVpcBindingHeaders, runtime *util.RuntimeOptions) (_result *CreateVpcBindingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	serviceName = openapiutil.GetEncodeParam(serviceName)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["vpcId"] = request.VpcId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVpcBinding"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName) + "/bindings"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &CreateVpcBindingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAlias(serviceName *string, aliasName *string) (_result *DeleteAliasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteAliasHeaders{}
	_result = &DeleteAliasResponse{}
	_body, _err := client.DeleteAliasWithOptions(serviceName, aliasName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAliasWithOptions(serviceName *string, aliasName *string, headers *DeleteAliasHeaders, runtime *util.RuntimeOptions) (_result *DeleteAliasResponse, _err error) {
	serviceName = openapiutil.GetEncodeParam(serviceName)
	aliasName = openapiutil.GetEncodeParam(aliasName)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.IfMatch)) {
		realHeaders["If-Match"] = util.ToJSONString(headers.IfMatch)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAlias"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName) + "/aliases/" + tea.StringValue(aliasName)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteAliasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCustomDomain(domainName *string) (_result *DeleteCustomDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteCustomDomainHeaders{}
	_result = &DeleteCustomDomainResponse{}
	_body, _err := client.DeleteCustomDomainWithOptions(domainName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCustomDomainWithOptions(domainName *string, headers *DeleteCustomDomainHeaders, runtime *util.RuntimeOptions) (_result *DeleteCustomDomainResponse, _err error) {
	domainName = openapiutil.GetEncodeParam(domainName)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCustomDomain"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/custom-domains/" + tea.StringValue(domainName)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteCustomDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFunction(serviceName *string, functionName *string) (_result *DeleteFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteFunctionHeaders{}
	_result = &DeleteFunctionResponse{}
	_body, _err := client.DeleteFunctionWithOptions(serviceName, functionName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFunctionWithOptions(serviceName *string, functionName *string, headers *DeleteFunctionHeaders, runtime *util.RuntimeOptions) (_result *DeleteFunctionResponse, _err error) {
	serviceName = openapiutil.GetEncodeParam(serviceName)
	functionName = openapiutil.GetEncodeParam(functionName)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.IfMatch)) {
		realHeaders["If-Match"] = util.ToJSONString(headers.IfMatch)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFunction"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName) + "/functions/" + tea.StringValue(functionName)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFunctionAsyncInvokeConfig(serviceName *string, functionName *string, request *DeleteFunctionAsyncInvokeConfigRequest) (_result *DeleteFunctionAsyncInvokeConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteFunctionAsyncInvokeConfigHeaders{}
	_result = &DeleteFunctionAsyncInvokeConfigResponse{}
	_body, _err := client.DeleteFunctionAsyncInvokeConfigWithOptions(serviceName, functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFunctionAsyncInvokeConfigWithOptions(serviceName *string, functionName *string, request *DeleteFunctionAsyncInvokeConfigRequest, headers *DeleteFunctionAsyncInvokeConfigHeaders, runtime *util.RuntimeOptions) (_result *DeleteFunctionAsyncInvokeConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	serviceName = openapiutil.GetEncodeParam(serviceName)
	functionName = openapiutil.GetEncodeParam(functionName)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Qualifier)) {
		query["qualifier"] = request.Qualifier
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFunctionAsyncInvokeConfig"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName) + "/functions/" + tea.StringValue(functionName) + "/async-invoke-config"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteFunctionAsyncInvokeConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFunctionOnDemandConfig(serviceName *string, functionName *string, request *DeleteFunctionOnDemandConfigRequest) (_result *DeleteFunctionOnDemandConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteFunctionOnDemandConfigHeaders{}
	_result = &DeleteFunctionOnDemandConfigResponse{}
	_body, _err := client.DeleteFunctionOnDemandConfigWithOptions(serviceName, functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFunctionOnDemandConfigWithOptions(serviceName *string, functionName *string, request *DeleteFunctionOnDemandConfigRequest, headers *DeleteFunctionOnDemandConfigHeaders, runtime *util.RuntimeOptions) (_result *DeleteFunctionOnDemandConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	serviceName = openapiutil.GetEncodeParam(serviceName)
	functionName = openapiutil.GetEncodeParam(functionName)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Qualifier)) {
		query["qualifier"] = request.Qualifier
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.IfMatch)) {
		realHeaders["If-Match"] = util.ToJSONString(headers.IfMatch)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFunctionOnDemandConfig"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName) + "/functions/" + tea.StringValue(functionName) + "/on-demand-config"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteFunctionOnDemandConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLayerVersion(layerName *string, version *string) (_result *DeleteLayerVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteLayerVersionHeaders{}
	_result = &DeleteLayerVersionResponse{}
	_body, _err := client.DeleteLayerVersionWithOptions(layerName, version, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLayerVersionWithOptions(layerName *string, version *string, headers *DeleteLayerVersionHeaders, runtime *util.RuntimeOptions) (_result *DeleteLayerVersionResponse, _err error) {
	layerName = openapiutil.GetEncodeParam(layerName)
	version = openapiutil.GetEncodeParam(version)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLayerVersion"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/layers/" + tea.StringValue(layerName) + "/versions/" + tea.StringValue(version)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteLayerVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteService(serviceName *string) (_result *DeleteServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteServiceHeaders{}
	_result = &DeleteServiceResponse{}
	_body, _err := client.DeleteServiceWithOptions(serviceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteServiceWithOptions(serviceName *string, headers *DeleteServiceHeaders, runtime *util.RuntimeOptions) (_result *DeleteServiceResponse, _err error) {
	serviceName = openapiutil.GetEncodeParam(serviceName)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.IfMatch)) {
		realHeaders["If-Match"] = util.ToJSONString(headers.IfMatch)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteService"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteServiceVersion(serviceName *string, versionId *string) (_result *DeleteServiceVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteServiceVersionHeaders{}
	_result = &DeleteServiceVersionResponse{}
	_body, _err := client.DeleteServiceVersionWithOptions(serviceName, versionId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteServiceVersionWithOptions(serviceName *string, versionId *string, headers *DeleteServiceVersionHeaders, runtime *util.RuntimeOptions) (_result *DeleteServiceVersionResponse, _err error) {
	serviceName = openapiutil.GetEncodeParam(serviceName)
	versionId = openapiutil.GetEncodeParam(versionId)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteServiceVersion"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName) + "/versions/" + tea.StringValue(versionId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteServiceVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTrigger(serviceName *string, functionName *string, triggerName *string) (_result *DeleteTriggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteTriggerHeaders{}
	_result = &DeleteTriggerResponse{}
	_body, _err := client.DeleteTriggerWithOptions(serviceName, functionName, triggerName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTriggerWithOptions(serviceName *string, functionName *string, triggerName *string, headers *DeleteTriggerHeaders, runtime *util.RuntimeOptions) (_result *DeleteTriggerResponse, _err error) {
	serviceName = openapiutil.GetEncodeParam(serviceName)
	functionName = openapiutil.GetEncodeParam(functionName)
	triggerName = openapiutil.GetEncodeParam(triggerName)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.IfMatch)) {
		realHeaders["If-Match"] = util.ToJSONString(headers.IfMatch)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTrigger"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName) + "/functions/" + tea.StringValue(functionName) + "/triggers/" + tea.StringValue(triggerName)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteTriggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVpcBinding(serviceName *string, vpcId *string) (_result *DeleteVpcBindingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteVpcBindingHeaders{}
	_result = &DeleteVpcBindingResponse{}
	_body, _err := client.DeleteVpcBindingWithOptions(serviceName, vpcId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVpcBindingWithOptions(serviceName *string, vpcId *string, headers *DeleteVpcBindingHeaders, runtime *util.RuntimeOptions) (_result *DeleteVpcBindingResponse, _err error) {
	serviceName = openapiutil.GetEncodeParam(serviceName)
	vpcId = openapiutil.GetEncodeParam(vpcId)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVpcBinding"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName) + "/bindings/" + tea.StringValue(vpcId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteVpcBindingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeregisterEventSource(serviceName *string, functionName *string, sourceArn *string, request *DeregisterEventSourceRequest) (_result *DeregisterEventSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeregisterEventSourceHeaders{}
	_result = &DeregisterEventSourceResponse{}
	_body, _err := client.DeregisterEventSourceWithOptions(serviceName, functionName, sourceArn, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeregisterEventSourceWithOptions(serviceName *string, functionName *string, sourceArn *string, request *DeregisterEventSourceRequest, headers *DeregisterEventSourceHeaders, runtime *util.RuntimeOptions) (_result *DeregisterEventSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	serviceName = openapiutil.GetEncodeParam(serviceName)
	functionName = openapiutil.GetEncodeParam(functionName)
	sourceArn = openapiutil.GetEncodeParam(sourceArn)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Qualifier)) {
		query["qualifier"] = request.Qualifier
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeregisterEventSource"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName) + "/functions/" + tea.StringValue(functionName) + "/event-sources/" + tea.StringValue(sourceArn)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeregisterEventSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAccountSettings() (_result *GetAccountSettingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAccountSettingsHeaders{}
	_result = &GetAccountSettingsResponse{}
	_body, _err := client.GetAccountSettingsWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAccountSettingsWithOptions(headers *GetAccountSettingsHeaders, runtime *util.RuntimeOptions) (_result *GetAccountSettingsResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("GetAccountSettings"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/account-settings"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAccountSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAlias(serviceName *string, aliasName *string) (_result *GetAliasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAliasHeaders{}
	_result = &GetAliasResponse{}
	_body, _err := client.GetAliasWithOptions(serviceName, aliasName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAliasWithOptions(serviceName *string, aliasName *string, headers *GetAliasHeaders, runtime *util.RuntimeOptions) (_result *GetAliasResponse, _err error) {
	serviceName = openapiutil.GetEncodeParam(serviceName)
	aliasName = openapiutil.GetEncodeParam(aliasName)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("GetAlias"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName) + "/aliases/" + tea.StringValue(aliasName)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAliasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCustomDomain(domainName *string) (_result *GetCustomDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCustomDomainHeaders{}
	_result = &GetCustomDomainResponse{}
	_body, _err := client.GetCustomDomainWithOptions(domainName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCustomDomainWithOptions(domainName *string, headers *GetCustomDomainHeaders, runtime *util.RuntimeOptions) (_result *GetCustomDomainResponse, _err error) {
	domainName = openapiutil.GetEncodeParam(domainName)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("GetCustomDomain"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/custom-domains/" + tea.StringValue(domainName)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCustomDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFunction(serviceName *string, functionName *string, request *GetFunctionRequest) (_result *GetFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFunctionHeaders{}
	_result = &GetFunctionResponse{}
	_body, _err := client.GetFunctionWithOptions(serviceName, functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFunctionWithOptions(serviceName *string, functionName *string, request *GetFunctionRequest, headers *GetFunctionHeaders, runtime *util.RuntimeOptions) (_result *GetFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	serviceName = openapiutil.GetEncodeParam(serviceName)
	functionName = openapiutil.GetEncodeParam(functionName)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Qualifier)) {
		query["qualifier"] = request.Qualifier
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFunction"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName) + "/functions/" + tea.StringValue(functionName)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFunctionAsyncInvokeConfig(serviceName *string, functionName *string, request *GetFunctionAsyncInvokeConfigRequest) (_result *GetFunctionAsyncInvokeConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFunctionAsyncInvokeConfigHeaders{}
	_result = &GetFunctionAsyncInvokeConfigResponse{}
	_body, _err := client.GetFunctionAsyncInvokeConfigWithOptions(serviceName, functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFunctionAsyncInvokeConfigWithOptions(serviceName *string, functionName *string, request *GetFunctionAsyncInvokeConfigRequest, headers *GetFunctionAsyncInvokeConfigHeaders, runtime *util.RuntimeOptions) (_result *GetFunctionAsyncInvokeConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	serviceName = openapiutil.GetEncodeParam(serviceName)
	functionName = openapiutil.GetEncodeParam(functionName)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Qualifier)) {
		query["qualifier"] = request.Qualifier
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFunctionAsyncInvokeConfig"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName) + "/functions/" + tea.StringValue(functionName) + "/async-invoke-config"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFunctionAsyncInvokeConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFunctionCode(serviceName *string, functionName *string, request *GetFunctionCodeRequest) (_result *GetFunctionCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFunctionCodeHeaders{}
	_result = &GetFunctionCodeResponse{}
	_body, _err := client.GetFunctionCodeWithOptions(serviceName, functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFunctionCodeWithOptions(serviceName *string, functionName *string, request *GetFunctionCodeRequest, headers *GetFunctionCodeHeaders, runtime *util.RuntimeOptions) (_result *GetFunctionCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	serviceName = openapiutil.GetEncodeParam(serviceName)
	functionName = openapiutil.GetEncodeParam(functionName)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Qualifier)) {
		query["qualifier"] = request.Qualifier
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFunctionCode"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName) + "/functions/" + tea.StringValue(functionName) + "/code"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFunctionCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFunctionOnDemandConfig(serviceName *string, functionName *string, request *GetFunctionOnDemandConfigRequest) (_result *GetFunctionOnDemandConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFunctionOnDemandConfigHeaders{}
	_result = &GetFunctionOnDemandConfigResponse{}
	_body, _err := client.GetFunctionOnDemandConfigWithOptions(serviceName, functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFunctionOnDemandConfigWithOptions(serviceName *string, functionName *string, request *GetFunctionOnDemandConfigRequest, headers *GetFunctionOnDemandConfigHeaders, runtime *util.RuntimeOptions) (_result *GetFunctionOnDemandConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	serviceName = openapiutil.GetEncodeParam(serviceName)
	functionName = openapiutil.GetEncodeParam(functionName)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Qualifier)) {
		query["qualifier"] = request.Qualifier
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFunctionOnDemandConfig"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName) + "/functions/" + tea.StringValue(functionName) + "/on-demand-config"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFunctionOnDemandConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLayerVersion(layerName *string, version *string) (_result *GetLayerVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetLayerVersionHeaders{}
	_result = &GetLayerVersionResponse{}
	_body, _err := client.GetLayerVersionWithOptions(layerName, version, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLayerVersionWithOptions(layerName *string, version *string, headers *GetLayerVersionHeaders, runtime *util.RuntimeOptions) (_result *GetLayerVersionResponse, _err error) {
	layerName = openapiutil.GetEncodeParam(layerName)
	version = openapiutil.GetEncodeParam(version)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("GetLayerVersion"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/layers/" + tea.StringValue(layerName) + "/versions/" + tea.StringValue(version)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLayerVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLayerVersionByArn(arn *string) (_result *GetLayerVersionByArnResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetLayerVersionByArnHeaders{}
	_result = &GetLayerVersionByArnResponse{}
	_body, _err := client.GetLayerVersionByArnWithOptions(arn, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLayerVersionByArnWithOptions(arn *string, headers *GetLayerVersionByArnHeaders, runtime *util.RuntimeOptions) (_result *GetLayerVersionByArnResponse, _err error) {
	arn = openapiutil.GetEncodeParam(arn)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("GetLayerVersionByArn"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/layerarn/" + tea.StringValue(arn)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLayerVersionByArnResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProvisionConfig(serviceName *string, functionName *string, request *GetProvisionConfigRequest) (_result *GetProvisionConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetProvisionConfigHeaders{}
	_result = &GetProvisionConfigResponse{}
	_body, _err := client.GetProvisionConfigWithOptions(serviceName, functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProvisionConfigWithOptions(serviceName *string, functionName *string, request *GetProvisionConfigRequest, headers *GetProvisionConfigHeaders, runtime *util.RuntimeOptions) (_result *GetProvisionConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	serviceName = openapiutil.GetEncodeParam(serviceName)
	functionName = openapiutil.GetEncodeParam(functionName)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Qualifier)) {
		query["qualifier"] = request.Qualifier
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetProvisionConfig"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName) + "/functions/" + tea.StringValue(functionName) + "/provision-config"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProvisionConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetResourceTags(request *GetResourceTagsRequest) (_result *GetResourceTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetResourceTagsHeaders{}
	_result = &GetResourceTagsResponse{}
	_body, _err := client.GetResourceTagsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetResourceTagsWithOptions(request *GetResourceTagsRequest, headers *GetResourceTagsHeaders, runtime *util.RuntimeOptions) (_result *GetResourceTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceArn)) {
		query["resourceArn"] = request.ResourceArn
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResourceTags"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/tag"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourceTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetService(serviceName *string, request *GetServiceRequest) (_result *GetServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetServiceHeaders{}
	_result = &GetServiceResponse{}
	_body, _err := client.GetServiceWithOptions(serviceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetServiceWithOptions(serviceName *string, request *GetServiceRequest, headers *GetServiceHeaders, runtime *util.RuntimeOptions) (_result *GetServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	serviceName = openapiutil.GetEncodeParam(serviceName)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Qualifier)) {
		query["qualifier"] = request.Qualifier
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetService"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetStatefulAsyncInvocation(serviceName *string, functionName *string, invocationId *string, request *GetStatefulAsyncInvocationRequest) (_result *GetStatefulAsyncInvocationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetStatefulAsyncInvocationHeaders{}
	_result = &GetStatefulAsyncInvocationResponse{}
	_body, _err := client.GetStatefulAsyncInvocationWithOptions(serviceName, functionName, invocationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetStatefulAsyncInvocationWithOptions(serviceName *string, functionName *string, invocationId *string, request *GetStatefulAsyncInvocationRequest, headers *GetStatefulAsyncInvocationHeaders, runtime *util.RuntimeOptions) (_result *GetStatefulAsyncInvocationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	serviceName = openapiutil.GetEncodeParam(serviceName)
	functionName = openapiutil.GetEncodeParam(functionName)
	invocationId = openapiutil.GetEncodeParam(invocationId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Qualifier)) {
		query["qualifier"] = request.Qualifier
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetStatefulAsyncInvocation"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName) + "/functions/" + tea.StringValue(functionName) + "/stateful-async-invocations/" + tea.StringValue(invocationId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetStatefulAsyncInvocationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTrigger(serviceName *string, functionName *string, triggerName *string) (_result *GetTriggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTriggerHeaders{}
	_result = &GetTriggerResponse{}
	_body, _err := client.GetTriggerWithOptions(serviceName, functionName, triggerName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTriggerWithOptions(serviceName *string, functionName *string, triggerName *string, headers *GetTriggerHeaders, runtime *util.RuntimeOptions) (_result *GetTriggerResponse, _err error) {
	serviceName = openapiutil.GetEncodeParam(serviceName)
	functionName = openapiutil.GetEncodeParam(functionName)
	triggerName = openapiutil.GetEncodeParam(triggerName)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("GetTrigger"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName) + "/functions/" + tea.StringValue(functionName) + "/triggers/" + tea.StringValue(triggerName)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTriggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InvokeFunction(serviceName *string, functionName *string, request *InvokeFunctionRequest) (_result *InvokeFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InvokeFunctionHeaders{}
	_result = &InvokeFunctionResponse{}
	_body, _err := client.InvokeFunctionWithOptions(serviceName, functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InvokeFunctionWithOptions(serviceName *string, functionName *string, request *InvokeFunctionRequest, headers *InvokeFunctionHeaders, runtime *util.RuntimeOptions) (_result *InvokeFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	serviceName = openapiutil.GetEncodeParam(serviceName)
	functionName = openapiutil.GetEncodeParam(functionName)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Qualifier)) {
		query["qualifier"] = request.Qualifier
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcStatefulAsyncInvocationId)) {
		realHeaders["X-Fc-Stateful-Async-Invocation-Id"] = util.ToJSONString(headers.XFcStatefulAsyncInvocationId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("InvokeFunction"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName) + "/functions/" + tea.StringValue(functionName) + "/invocations"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("byte"),
	}
	_result = &InvokeFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAliases(serviceName *string, request *ListAliasesRequest) (_result *ListAliasesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListAliasesHeaders{}
	_result = &ListAliasesResponse{}
	_body, _err := client.ListAliasesWithOptions(serviceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAliasesWithOptions(serviceName *string, request *ListAliasesRequest, headers *ListAliasesHeaders, runtime *util.RuntimeOptions) (_result *ListAliasesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	serviceName = openapiutil.GetEncodeParam(serviceName)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Prefix)) {
		query["prefix"] = request.Prefix
	}

	if !tea.BoolValue(util.IsUnset(request.StartKey)) {
		query["startKey"] = request.StartKey
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAliases"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName) + "/aliases"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAliasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCustomDomains(request *ListCustomDomainsRequest) (_result *ListCustomDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListCustomDomainsHeaders{}
	_result = &ListCustomDomainsResponse{}
	_body, _err := client.ListCustomDomainsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCustomDomainsWithOptions(request *ListCustomDomainsRequest, headers *ListCustomDomainsHeaders, runtime *util.RuntimeOptions) (_result *ListCustomDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Prefix)) {
		query["prefix"] = request.Prefix
	}

	if !tea.BoolValue(util.IsUnset(request.StartKey)) {
		query["startKey"] = request.StartKey
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCustomDomains"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/custom-domains"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCustomDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEventSources(serviceName *string, functionName *string, request *ListEventSourcesRequest) (_result *ListEventSourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListEventSourcesHeaders{}
	_result = &ListEventSourcesResponse{}
	_body, _err := client.ListEventSourcesWithOptions(serviceName, functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEventSourcesWithOptions(serviceName *string, functionName *string, request *ListEventSourcesRequest, headers *ListEventSourcesHeaders, runtime *util.RuntimeOptions) (_result *ListEventSourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	serviceName = openapiutil.GetEncodeParam(serviceName)
	functionName = openapiutil.GetEncodeParam(functionName)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Qualifier)) {
		query["qualifier"] = request.Qualifier
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEventSources"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName) + "/functions/" + tea.StringValue(functionName) + "/event-sources"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEventSourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFunctionAsyncInvokeConfigs(serviceName *string, functionName *string, request *ListFunctionAsyncInvokeConfigsRequest) (_result *ListFunctionAsyncInvokeConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListFunctionAsyncInvokeConfigsHeaders{}
	_result = &ListFunctionAsyncInvokeConfigsResponse{}
	_body, _err := client.ListFunctionAsyncInvokeConfigsWithOptions(serviceName, functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFunctionAsyncInvokeConfigsWithOptions(serviceName *string, functionName *string, request *ListFunctionAsyncInvokeConfigsRequest, headers *ListFunctionAsyncInvokeConfigsHeaders, runtime *util.RuntimeOptions) (_result *ListFunctionAsyncInvokeConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	serviceName = openapiutil.GetEncodeParam(serviceName)
	functionName = openapiutil.GetEncodeParam(functionName)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFunctionAsyncInvokeConfigs"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName) + "/functions/" + tea.StringValue(functionName) + "/async-invoke-configs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFunctionAsyncInvokeConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFunctions(serviceName *string, request *ListFunctionsRequest) (_result *ListFunctionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListFunctionsHeaders{}
	_result = &ListFunctionsResponse{}
	_body, _err := client.ListFunctionsWithOptions(serviceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFunctionsWithOptions(serviceName *string, request *ListFunctionsRequest, headers *ListFunctionsHeaders, runtime *util.RuntimeOptions) (_result *ListFunctionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	serviceName = openapiutil.GetEncodeParam(serviceName)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Prefix)) {
		query["prefix"] = request.Prefix
	}

	if !tea.BoolValue(util.IsUnset(request.Qualifier)) {
		query["qualifier"] = request.Qualifier
	}

	if !tea.BoolValue(util.IsUnset(request.StartKey)) {
		query["startKey"] = request.StartKey
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFunctions"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName) + "/functions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFunctionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstances(serviceName *string, functionName *string, request *ListInstancesRequest) (_result *ListInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListInstancesHeaders{}
	_result = &ListInstancesResponse{}
	_body, _err := client.ListInstancesWithOptions(serviceName, functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstancesWithOptions(serviceName *string, functionName *string, request *ListInstancesRequest, headers *ListInstancesHeaders, runtime *util.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	serviceName = openapiutil.GetEncodeParam(serviceName)
	functionName = openapiutil.GetEncodeParam(functionName)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Qualifier)) {
		query["qualifier"] = request.Qualifier
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstances"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName) + "/functions/" + tea.StringValue(functionName) + "/instances"),
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

func (client *Client) ListLayerVersions(layerName *string, request *ListLayerVersionsRequest) (_result *ListLayerVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListLayerVersionsHeaders{}
	_result = &ListLayerVersionsResponse{}
	_body, _err := client.ListLayerVersionsWithOptions(layerName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListLayerVersionsWithOptions(layerName *string, request *ListLayerVersionsRequest, headers *ListLayerVersionsHeaders, runtime *util.RuntimeOptions) (_result *ListLayerVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	layerName = openapiutil.GetEncodeParam(layerName)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.StartVersion)) {
		query["startVersion"] = request.StartVersion
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLayerVersions"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/layers/" + tea.StringValue(layerName) + "/versions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListLayerVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListLayers(request *ListLayersRequest) (_result *ListLayersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListLayersHeaders{}
	_result = &ListLayersResponse{}
	_body, _err := client.ListLayersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListLayersWithOptions(request *ListLayersRequest, headers *ListLayersHeaders, runtime *util.RuntimeOptions) (_result *ListLayersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Prefix)) {
		query["prefix"] = request.Prefix
	}

	if !tea.BoolValue(util.IsUnset(request.StartKey)) {
		query["startKey"] = request.StartKey
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLayers"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/layers"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListLayersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOnDemandConfigs(request *ListOnDemandConfigsRequest) (_result *ListOnDemandConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListOnDemandConfigsHeaders{}
	_result = &ListOnDemandConfigsResponse{}
	_body, _err := client.ListOnDemandConfigsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOnDemandConfigsWithOptions(request *ListOnDemandConfigsRequest, headers *ListOnDemandConfigsHeaders, runtime *util.RuntimeOptions) (_result *ListOnDemandConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Prefix)) {
		query["prefix"] = request.Prefix
	}

	if !tea.BoolValue(util.IsUnset(request.StartKey)) {
		query["startKey"] = request.StartKey
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOnDemandConfigs"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/on-demand-configs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOnDemandConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProvisionConfigs(request *ListProvisionConfigsRequest) (_result *ListProvisionConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListProvisionConfigsHeaders{}
	_result = &ListProvisionConfigsResponse{}
	_body, _err := client.ListProvisionConfigsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProvisionConfigsWithOptions(request *ListProvisionConfigsRequest, headers *ListProvisionConfigsHeaders, runtime *util.RuntimeOptions) (_result *ListProvisionConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Qualifier)) {
		query["qualifier"] = request.Qualifier
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["serviceName"] = request.ServiceName
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProvisionConfigs"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/provision-configs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProvisionConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListReservedCapacities(request *ListReservedCapacitiesRequest) (_result *ListReservedCapacitiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListReservedCapacitiesHeaders{}
	_result = &ListReservedCapacitiesResponse{}
	_body, _err := client.ListReservedCapacitiesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListReservedCapacitiesWithOptions(request *ListReservedCapacitiesRequest, headers *ListReservedCapacitiesHeaders, runtime *util.RuntimeOptions) (_result *ListReservedCapacitiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListReservedCapacities"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/reserved-capacities"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListReservedCapacitiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListServiceVersions(serviceName *string, request *ListServiceVersionsRequest) (_result *ListServiceVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListServiceVersionsHeaders{}
	_result = &ListServiceVersionsResponse{}
	_body, _err := client.ListServiceVersionsWithOptions(serviceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListServiceVersionsWithOptions(serviceName *string, request *ListServiceVersionsRequest, headers *ListServiceVersionsHeaders, runtime *util.RuntimeOptions) (_result *ListServiceVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	serviceName = openapiutil.GetEncodeParam(serviceName)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		query["direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.StartKey)) {
		query["startKey"] = request.StartKey
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceVersions"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName) + "/versions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListServices(request *ListServicesRequest) (_result *ListServicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListServicesHeaders{}
	_result = &ListServicesResponse{}
	_body, _err := client.ListServicesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListServicesWithOptions(request *ListServicesRequest, headers *ListServicesHeaders, runtime *util.RuntimeOptions) (_result *ListServicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Prefix)) {
		query["prefix"] = request.Prefix
	}

	if !tea.BoolValue(util.IsUnset(request.StartKey)) {
		query["startKey"] = request.StartKey
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServices"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListStatefulAsyncInvocations(serviceName *string, functionName *string, request *ListStatefulAsyncInvocationsRequest) (_result *ListStatefulAsyncInvocationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListStatefulAsyncInvocationsHeaders{}
	_result = &ListStatefulAsyncInvocationsResponse{}
	_body, _err := client.ListStatefulAsyncInvocationsWithOptions(serviceName, functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListStatefulAsyncInvocationsWithOptions(serviceName *string, functionName *string, request *ListStatefulAsyncInvocationsRequest, headers *ListStatefulAsyncInvocationsHeaders, runtime *util.RuntimeOptions) (_result *ListStatefulAsyncInvocationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	serviceName = openapiutil.GetEncodeParam(serviceName)
	functionName = openapiutil.GetEncodeParam(functionName)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IncludePayload)) {
		query["includePayload"] = request.IncludePayload
	}

	if !tea.BoolValue(util.IsUnset(request.InvocationIdPrefix)) {
		query["invocationIdPrefix"] = request.InvocationIdPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Qualifier)) {
		query["qualifier"] = request.Qualifier
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrderByTime)) {
		query["sortOrderByTime"] = request.SortOrderByTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartedTimeBegin)) {
		query["startedTimeBegin"] = request.StartedTimeBegin
	}

	if !tea.BoolValue(util.IsUnset(request.StartedTimeEnd)) {
		query["startedTimeEnd"] = request.StartedTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListStatefulAsyncInvocations"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName) + "/functions/" + tea.StringValue(functionName) + "/stateful-async-invocations"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListStatefulAsyncInvocationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTaggedResources(request *ListTaggedResourcesRequest) (_result *ListTaggedResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListTaggedResourcesHeaders{}
	_result = &ListTaggedResourcesResponse{}
	_body, _err := client.ListTaggedResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTaggedResourcesWithOptions(request *ListTaggedResourcesRequest, headers *ListTaggedResourcesHeaders, runtime *util.RuntimeOptions) (_result *ListTaggedResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTaggedResources"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/tags"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTaggedResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTriggers(serviceName *string, functionName *string, request *ListTriggersRequest) (_result *ListTriggersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListTriggersHeaders{}
	_result = &ListTriggersResponse{}
	_body, _err := client.ListTriggersWithOptions(serviceName, functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTriggersWithOptions(serviceName *string, functionName *string, request *ListTriggersRequest, headers *ListTriggersHeaders, runtime *util.RuntimeOptions) (_result *ListTriggersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	serviceName = openapiutil.GetEncodeParam(serviceName)
	functionName = openapiutil.GetEncodeParam(functionName)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Prefix)) {
		query["prefix"] = request.Prefix
	}

	if !tea.BoolValue(util.IsUnset(request.StartKey)) {
		query["startKey"] = request.StartKey
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTriggers"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName) + "/functions/" + tea.StringValue(functionName) + "/triggers"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTriggersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVpcBindings(serviceName *string) (_result *ListVpcBindingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListVpcBindingsHeaders{}
	_result = &ListVpcBindingsResponse{}
	_body, _err := client.ListVpcBindingsWithOptions(serviceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVpcBindingsWithOptions(serviceName *string, headers *ListVpcBindingsHeaders, runtime *util.RuntimeOptions) (_result *ListVpcBindingsResponse, _err error) {
	serviceName = openapiutil.GetEncodeParam(serviceName)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("ListVpcBindings"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName) + "/bindings"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVpcBindingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PermanentDeleteLayerVersion(userID *string, layerName *string, version *string) (_result *PermanentDeleteLayerVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PermanentDeleteLayerVersionHeaders{}
	_result = &PermanentDeleteLayerVersionResponse{}
	_body, _err := client.PermanentDeleteLayerVersionWithOptions(userID, layerName, version, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PermanentDeleteLayerVersionWithOptions(userID *string, layerName *string, version *string, headers *PermanentDeleteLayerVersionHeaders, runtime *util.RuntimeOptions) (_result *PermanentDeleteLayerVersionResponse, _err error) {
	userID = openapiutil.GetEncodeParam(userID)
	layerName = openapiutil.GetEncodeParam(layerName)
	version = openapiutil.GetEncodeParam(version)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("PermanentDeleteLayerVersion"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/adminlayers/" + tea.StringValue(userID) + "/" + tea.StringValue(layerName) + "/versions/" + tea.StringValue(version)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &PermanentDeleteLayerVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PublishLayerAsPublic(layerName *string, version *string) (_result *PublishLayerAsPublicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PublishLayerAsPublicHeaders{}
	_result = &PublishLayerAsPublicResponse{}
	_body, _err := client.PublishLayerAsPublicWithOptions(layerName, version, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PublishLayerAsPublicWithOptions(layerName *string, version *string, headers *PublishLayerAsPublicHeaders, runtime *util.RuntimeOptions) (_result *PublishLayerAsPublicResponse, _err error) {
	layerName = openapiutil.GetEncodeParam(layerName)
	version = openapiutil.GetEncodeParam(version)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("PublishLayerAsPublic"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/layers/" + tea.StringValue(layerName) + "/versions/" + tea.StringValue(version)),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &PublishLayerAsPublicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PublishServiceVersion(serviceName *string, request *PublishServiceVersionRequest) (_result *PublishServiceVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PublishServiceVersionHeaders{}
	_result = &PublishServiceVersionResponse{}
	_body, _err := client.PublishServiceVersionWithOptions(serviceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PublishServiceVersionWithOptions(serviceName *string, request *PublishServiceVersionRequest, headers *PublishServiceVersionHeaders, runtime *util.RuntimeOptions) (_result *PublishServiceVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	serviceName = openapiutil.GetEncodeParam(serviceName)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.IfMatch)) {
		realHeaders["If-Match"] = util.ToJSONString(headers.IfMatch)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PublishServiceVersion"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName) + "/versions"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PublishServiceVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutFunctionAsyncInvokeConfig(serviceName *string, functionName *string, request *PutFunctionAsyncInvokeConfigRequest) (_result *PutFunctionAsyncInvokeConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PutFunctionAsyncInvokeConfigHeaders{}
	_result = &PutFunctionAsyncInvokeConfigResponse{}
	_body, _err := client.PutFunctionAsyncInvokeConfigWithOptions(serviceName, functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutFunctionAsyncInvokeConfigWithOptions(serviceName *string, functionName *string, request *PutFunctionAsyncInvokeConfigRequest, headers *PutFunctionAsyncInvokeConfigHeaders, runtime *util.RuntimeOptions) (_result *PutFunctionAsyncInvokeConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	serviceName = openapiutil.GetEncodeParam(serviceName)
	functionName = openapiutil.GetEncodeParam(functionName)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Qualifier)) {
		query["qualifier"] = request.Qualifier
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.DestinationConfig))) {
		body["destinationConfig"] = request.DestinationConfig
	}

	if !tea.BoolValue(util.IsUnset(request.MaxAsyncEventAgeInSeconds)) {
		body["maxAsyncEventAgeInSeconds"] = request.MaxAsyncEventAgeInSeconds
	}

	if !tea.BoolValue(util.IsUnset(request.MaxAsyncRetryAttempts)) {
		body["maxAsyncRetryAttempts"] = request.MaxAsyncRetryAttempts
	}

	if !tea.BoolValue(util.IsUnset(request.StatefulInvocation)) {
		body["statefulInvocation"] = request.StatefulInvocation
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PutFunctionAsyncInvokeConfig"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName) + "/functions/" + tea.StringValue(functionName) + "/async-invoke-config"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PutFunctionAsyncInvokeConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutFunctionOnDemandConfig(serviceName *string, functionName *string, request *PutFunctionOnDemandConfigRequest) (_result *PutFunctionOnDemandConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PutFunctionOnDemandConfigHeaders{}
	_result = &PutFunctionOnDemandConfigResponse{}
	_body, _err := client.PutFunctionOnDemandConfigWithOptions(serviceName, functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutFunctionOnDemandConfigWithOptions(serviceName *string, functionName *string, request *PutFunctionOnDemandConfigRequest, headers *PutFunctionOnDemandConfigHeaders, runtime *util.RuntimeOptions) (_result *PutFunctionOnDemandConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	serviceName = openapiutil.GetEncodeParam(serviceName)
	functionName = openapiutil.GetEncodeParam(functionName)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Qualifier)) {
		query["qualifier"] = request.Qualifier
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaximumInstanceCount)) {
		body["maximumInstanceCount"] = request.MaximumInstanceCount
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.IfMatch)) {
		realHeaders["If-Match"] = util.ToJSONString(headers.IfMatch)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PutFunctionOnDemandConfig"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName) + "/functions/" + tea.StringValue(functionName) + "/on-demand-config"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PutFunctionOnDemandConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutProvisionConfig(serviceName *string, functionName *string, request *PutProvisionConfigRequest) (_result *PutProvisionConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PutProvisionConfigHeaders{}
	_result = &PutProvisionConfigResponse{}
	_body, _err := client.PutProvisionConfigWithOptions(serviceName, functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutProvisionConfigWithOptions(serviceName *string, functionName *string, request *PutProvisionConfigRequest, headers *PutProvisionConfigHeaders, runtime *util.RuntimeOptions) (_result *PutProvisionConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	serviceName = openapiutil.GetEncodeParam(serviceName)
	functionName = openapiutil.GetEncodeParam(functionName)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Qualifier)) {
		query["qualifier"] = request.Qualifier
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ScheduledActions)) {
		body["scheduledActions"] = request.ScheduledActions
	}

	if !tea.BoolValue(util.IsUnset(request.Target)) {
		body["target"] = request.Target
	}

	if !tea.BoolValue(util.IsUnset(request.TargetTrackingPolicies)) {
		body["targetTrackingPolicies"] = request.TargetTrackingPolicies
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PutProvisionConfig"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName) + "/functions/" + tea.StringValue(functionName) + "/provision-config"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PutProvisionConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RegisterEventSource(serviceName *string, functionName *string, request *RegisterEventSourceRequest) (_result *RegisterEventSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RegisterEventSourceHeaders{}
	_result = &RegisterEventSourceResponse{}
	_body, _err := client.RegisterEventSourceWithOptions(serviceName, functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RegisterEventSourceWithOptions(serviceName *string, functionName *string, request *RegisterEventSourceRequest, headers *RegisterEventSourceHeaders, runtime *util.RuntimeOptions) (_result *RegisterEventSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	serviceName = openapiutil.GetEncodeParam(serviceName)
	functionName = openapiutil.GetEncodeParam(functionName)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Qualifier)) {
		query["qualifier"] = request.Qualifier
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SourceArn)) {
		body["sourceArn"] = request.SourceArn
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RegisterEventSource"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName) + "/functions/" + tea.StringValue(functionName) + "/event-sources"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RegisterEventSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopStatefulAsyncInvocation(serviceName *string, functionName *string, invocationId *string, request *StopStatefulAsyncInvocationRequest) (_result *StopStatefulAsyncInvocationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &StopStatefulAsyncInvocationHeaders{}
	_result = &StopStatefulAsyncInvocationResponse{}
	_body, _err := client.StopStatefulAsyncInvocationWithOptions(serviceName, functionName, invocationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopStatefulAsyncInvocationWithOptions(serviceName *string, functionName *string, invocationId *string, request *StopStatefulAsyncInvocationRequest, headers *StopStatefulAsyncInvocationHeaders, runtime *util.RuntimeOptions) (_result *StopStatefulAsyncInvocationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	serviceName = openapiutil.GetEncodeParam(serviceName)
	functionName = openapiutil.GetEncodeParam(functionName)
	invocationId = openapiutil.GetEncodeParam(invocationId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Qualifier)) {
		query["qualifier"] = request.Qualifier
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopStatefulAsyncInvocation"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName) + "/functions/" + tea.StringValue(functionName) + "/stateful-async-invocations/" + tea.StringValue(invocationId)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &StopStatefulAsyncInvocationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResource(request *TagResourceRequest) (_result *TagResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TagResourceHeaders{}
	_result = &TagResourceResponse{}
	_body, _err := client.TagResourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourceWithOptions(request *TagResourceRequest, headers *TagResourceHeaders, runtime *util.RuntimeOptions) (_result *TagResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceArn)) {
		body["resourceArn"] = request.ResourceArn
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		body["tags"] = request.Tags
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResource"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/tag"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &TagResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResource(request *UntagResourceRequest) (_result *UntagResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UntagResourceHeaders{}
	_result = &UntagResourceResponse{}
	_body, _err := client.UntagResourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourceWithOptions(request *UntagResourceRequest, headers *UntagResourceHeaders, runtime *util.RuntimeOptions) (_result *UntagResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		body["all"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceArn)) {
		body["resourceArn"] = request.ResourceArn
	}

	if !tea.BoolValue(util.IsUnset(request.TagKeys)) {
		body["tagKeys	"] = request.TagKeys
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResource"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/tag"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UntagResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAlias(serviceName *string, aliasName *string, request *UpdateAliasRequest) (_result *UpdateAliasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateAliasHeaders{}
	_result = &UpdateAliasResponse{}
	_body, _err := client.UpdateAliasWithOptions(serviceName, aliasName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAliasWithOptions(serviceName *string, aliasName *string, request *UpdateAliasRequest, headers *UpdateAliasHeaders, runtime *util.RuntimeOptions) (_result *UpdateAliasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	serviceName = openapiutil.GetEncodeParam(serviceName)
	aliasName = openapiutil.GetEncodeParam(aliasName)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdditionalVersionWeight)) {
		body["additionalVersionWeight"] = request.AdditionalVersionWeight
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		body["versionId"] = request.VersionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.IfMatch)) {
		realHeaders["If-Match"] = util.ToJSONString(headers.IfMatch)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAlias"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName) + "/aliases/" + tea.StringValue(aliasName)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAliasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCustomDomain(domainName *string, request *UpdateCustomDomainRequest) (_result *UpdateCustomDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateCustomDomainHeaders{}
	_result = &UpdateCustomDomainResponse{}
	_body, _err := client.UpdateCustomDomainWithOptions(domainName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCustomDomainWithOptions(domainName *string, request *UpdateCustomDomainRequest, headers *UpdateCustomDomainHeaders, runtime *util.RuntimeOptions) (_result *UpdateCustomDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	domainName = openapiutil.GetEncodeParam(domainName)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.CertConfig))) {
		body["certConfig"] = request.CertConfig
	}

	if !tea.BoolValue(util.IsUnset(request.Protocol)) {
		body["protocol"] = request.Protocol
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.RouteConfig))) {
		body["routeConfig"] = request.RouteConfig
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCustomDomain"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/custom-domains/" + tea.StringValue(domainName)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCustomDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFunction(serviceName *string, functionName *string, request *UpdateFunctionRequest) (_result *UpdateFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateFunctionHeaders{}
	_result = &UpdateFunctionResponse{}
	_body, _err := client.UpdateFunctionWithOptions(serviceName, functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFunctionWithOptions(serviceName *string, functionName *string, request *UpdateFunctionRequest, headers *UpdateFunctionHeaders, runtime *util.RuntimeOptions) (_result *UpdateFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	serviceName = openapiutil.GetEncodeParam(serviceName)
	functionName = openapiutil.GetEncodeParam(functionName)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceConcurrency)) {
		body["InstanceConcurrency"] = request.InstanceConcurrency
	}

	if !tea.BoolValue(util.IsUnset(request.CaPort)) {
		body["caPort"] = request.CaPort
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Code))) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.CustomContainerConfig))) {
		body["customContainerConfig"] = request.CustomContainerConfig
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.CustomDNS))) {
		body["customDNS"] = request.CustomDNS
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.CustomRuntimeConfig))) {
		body["customRuntimeConfig"] = request.CustomRuntimeConfig
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EnvironmentVariables)) {
		body["environmentVariables"] = request.EnvironmentVariables
	}

	if !tea.BoolValue(util.IsUnset(request.Handler)) {
		body["handler"] = request.Handler
	}

	if !tea.BoolValue(util.IsUnset(request.InitializationTimeout)) {
		body["initializationTimeout"] = request.InitializationTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.Initializer)) {
		body["initializer"] = request.Initializer
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.InstanceLifecycleConfig))) {
		body["instanceLifecycleConfig"] = request.InstanceLifecycleConfig
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		body["instanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.Layers)) {
		body["layers"] = request.Layers
	}

	if !tea.BoolValue(util.IsUnset(request.MemorySize)) {
		body["memorySize"] = request.MemorySize
	}

	if !tea.BoolValue(util.IsUnset(request.Runtime)) {
		body["runtime"] = request.Runtime
	}

	if !tea.BoolValue(util.IsUnset(request.Timeout)) {
		body["timeout"] = request.Timeout
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.IfMatch)) {
		realHeaders["If-Match"] = util.ToJSONString(headers.IfMatch)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateFunction"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName) + "/functions/" + tea.StringValue(functionName)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateService(serviceName *string, request *UpdateServiceRequest) (_result *UpdateServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateServiceHeaders{}
	_result = &UpdateServiceResponse{}
	_body, _err := client.UpdateServiceWithOptions(serviceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateServiceWithOptions(serviceName *string, request *UpdateServiceRequest, headers *UpdateServiceHeaders, runtime *util.RuntimeOptions) (_result *UpdateServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	serviceName = openapiutil.GetEncodeParam(serviceName)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InternetAccess)) {
		body["internetAccess"] = request.InternetAccess
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.LogConfig))) {
		body["logConfig"] = request.LogConfig
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.NasConfig))) {
		body["nasConfig"] = request.NasConfig
	}

	if !tea.BoolValue(util.IsUnset(request.Role)) {
		body["role"] = request.Role
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.TracingConfig))) {
		body["tracingConfig"] = request.TracingConfig
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.VpcConfig))) {
		body["vpcConfig"] = request.VpcConfig
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.IfMatch)) {
		realHeaders["If-Match"] = util.ToJSONString(headers.IfMatch)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateService"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTrigger(serviceName *string, functionName *string, triggerName *string, request *UpdateTriggerRequest) (_result *UpdateTriggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateTriggerHeaders{}
	_result = &UpdateTriggerResponse{}
	_body, _err := client.UpdateTriggerWithOptions(serviceName, functionName, triggerName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTriggerWithOptions(serviceName *string, functionName *string, triggerName *string, request *UpdateTriggerRequest, headers *UpdateTriggerHeaders, runtime *util.RuntimeOptions) (_result *UpdateTriggerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	serviceName = openapiutil.GetEncodeParam(serviceName)
	functionName = openapiutil.GetEncodeParam(functionName)
	triggerName = openapiutil.GetEncodeParam(triggerName)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InvocationRole)) {
		body["invocationRole"] = request.InvocationRole
	}

	if !tea.BoolValue(util.IsUnset(request.Qualifier)) {
		body["qualifier"] = request.Qualifier
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerConfig)) {
		body["triggerConfig"] = request.TriggerConfig
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.IfMatch)) {
		realHeaders["If-Match"] = util.ToJSONString(headers.IfMatch)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcCodeChecksum)) {
		realHeaders["X-Fc-Code-Checksum"] = util.ToJSONString(headers.XFcCodeChecksum)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["X-Fc-Invocation-Type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["X-Fc-Log-Type"] = util.ToJSONString(headers.XFcLogType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTrigger"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/services/" + tea.StringValue(serviceName) + "/functions/" + tea.StringValue(functionName) + "/triggers/" + tea.StringValue(triggerName)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTriggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
