// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEdgeContainerAppVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateEdgeContainerAppVersionRequest
	GetAppId() *string
	SetContainers(v []*CreateEdgeContainerAppVersionRequestContainers) *CreateEdgeContainerAppVersionRequest
	GetContainers() []*CreateEdgeContainerAppVersionRequestContainers
	SetName(v string) *CreateEdgeContainerAppVersionRequest
	GetName() *string
	SetRemarks(v string) *CreateEdgeContainerAppVersionRequest
	GetRemarks() *string
}

type CreateEdgeContainerAppVersionRequest struct {
	// The application ID, which can be obtained by calling the [ListEdgeContainerApps](~~ListEdgeContainerApps~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// app-88068867578379****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The container group to be deployed for this version, which contains information about images.\\
	//
	// The image data contains the image address, startup command, parameters, environment variables, and probe rules. You can specify one or more images. The parameter value is a JSON string.
	//
	// This parameter is required.
	//
	// example:
	//
	// [
	//
	//       {
	//
	//             "Name": "container1",
	//
	//             "Image": "image1",
	//
	//             "Spec": "1C2G",
	//
	//             "Command": "/bin/sh",
	//
	//             "Args": "-c hello",
	//
	//             "ProbeType": "tcpSocket",
	//
	//             "ProbeContent": "{\\"Port\\":8080}"
	//
	//       },
	//
	//       {
	//
	//             "Name": "container2",
	//
	//             "Image": "image2",
	//
	//             "Spec": "2C4G",
	//
	//             "ProbeType": "httpGet",
	//
	//             "ProbeContent": "{\\"Path\\":\\"/\\",\\"Port\\":80,\\"InitialDelaySeconds\\":10}"
	//
	//       }
	//
	// ]
	Containers []*CreateEdgeContainerAppVersionRequestContainers `json:"Containers,omitempty" xml:"Containers,omitempty" type:"Repeated"`
	// The version name, which must be 6 to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// verson1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The description of the version.
	//
	// example:
	//
	// test app
	Remarks *string `json:"Remarks,omitempty" xml:"Remarks,omitempty"`
}

func (s CreateEdgeContainerAppVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEdgeContainerAppVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateEdgeContainerAppVersionRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateEdgeContainerAppVersionRequest) GetContainers() []*CreateEdgeContainerAppVersionRequestContainers {
	return s.Containers
}

func (s *CreateEdgeContainerAppVersionRequest) GetName() *string {
	return s.Name
}

func (s *CreateEdgeContainerAppVersionRequest) GetRemarks() *string {
	return s.Remarks
}

func (s *CreateEdgeContainerAppVersionRequest) SetAppId(v string) *CreateEdgeContainerAppVersionRequest {
	s.AppId = &v
	return s
}

func (s *CreateEdgeContainerAppVersionRequest) SetContainers(v []*CreateEdgeContainerAppVersionRequestContainers) *CreateEdgeContainerAppVersionRequest {
	s.Containers = v
	return s
}

func (s *CreateEdgeContainerAppVersionRequest) SetName(v string) *CreateEdgeContainerAppVersionRequest {
	s.Name = &v
	return s
}

func (s *CreateEdgeContainerAppVersionRequest) SetRemarks(v string) *CreateEdgeContainerAppVersionRequest {
	s.Remarks = &v
	return s
}

func (s *CreateEdgeContainerAppVersionRequest) Validate() error {
	return dara.Validate(s)
}

type CreateEdgeContainerAppVersionRequestContainers struct {
	// The information about the Container Registry image.
	ACRImageInfo *CreateEdgeContainerAppVersionRequestContainersACRImageInfo `json:"ACRImageInfo,omitempty" xml:"ACRImageInfo,omitempty" type:"Struct"`
	// The arguments that are passed to the container startup command. Separate the parameters with spaces.
	//
	// example:
	//
	// -a
	Args *string `json:"Args,omitempty" xml:"Args,omitempty"`
	// The command that is used to start the container. Separate the arguments with spaces.
	//
	// example:
	//
	// nginx
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The environment variables. Separate the environment variables with commas (,).
	//
	// example:
	//
	// VITE_APP_TITLE=My App
	EnvVariables *string `json:"EnvVariables,omitempty" xml:"EnvVariables,omitempty"`
	// The address of the image.
	//
	// This parameter is required.
	//
	// example:
	//
	// registry-vpc.cn-shenzhen.aliyuncs.com/lihe****h/ea****ts_serv****am:3.**
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// Specifies whether the image is a Container Registry image.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	IsACRImage *bool `json:"IsACRImage,omitempty" xml:"IsACRImage,omitempty"`
	// The name of the container. The name must be unique in the same container group.
	//
	// This parameter is required.
	//
	// example:
	//
	// lxg-demo-er
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The command that is run before the container is started. Separate the arguments with spaces.
	//
	// example:
	//
	// sh poststart.sh "echo hello world"
	PostStart *string `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	// The command that is run before the container is stopped. Separate the arguments with spaces.
	//
	// example:
	//
	// sh prestop.sh "echo hello world"
	PreStop *string `json:"PreStop,omitempty" xml:"PreStop,omitempty"`
	// The content of the container health probe.
	//
	// This parameter is required.
	ProbeContent *CreateEdgeContainerAppVersionRequestContainersProbeContent `json:"ProbeContent,omitempty" xml:"ProbeContent,omitempty" type:"Struct"`
	// The type of the probe. Valid values:
	//
	// 	- exec: the command type.
	//
	// 	- tcpSocket: the TCP probe type.
	//
	// 	- httpGet: the HTTP access type.
	//
	// This parameter is required.
	//
	// example:
	//
	// exec
	ProbeType *string `json:"ProbeType,omitempty" xml:"ProbeType,omitempty"`
	// The compute specification of the container. Valid values: 1C2G, 2C4G, 2C8G, 4C8G, 4C16G, 8C16G, and 8C32G.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1C2G
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The storage capacity. Valid values: 0.5G, 10G, 20G, and 30G.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.5G
	Storage *string `json:"Storage,omitempty" xml:"Storage,omitempty"`
}

func (s CreateEdgeContainerAppVersionRequestContainers) String() string {
	return dara.Prettify(s)
}

func (s CreateEdgeContainerAppVersionRequestContainers) GoString() string {
	return s.String()
}

func (s *CreateEdgeContainerAppVersionRequestContainers) GetACRImageInfo() *CreateEdgeContainerAppVersionRequestContainersACRImageInfo {
	return s.ACRImageInfo
}

func (s *CreateEdgeContainerAppVersionRequestContainers) GetArgs() *string {
	return s.Args
}

func (s *CreateEdgeContainerAppVersionRequestContainers) GetCommand() *string {
	return s.Command
}

func (s *CreateEdgeContainerAppVersionRequestContainers) GetEnvVariables() *string {
	return s.EnvVariables
}

func (s *CreateEdgeContainerAppVersionRequestContainers) GetImage() *string {
	return s.Image
}

func (s *CreateEdgeContainerAppVersionRequestContainers) GetIsACRImage() *bool {
	return s.IsACRImage
}

func (s *CreateEdgeContainerAppVersionRequestContainers) GetName() *string {
	return s.Name
}

func (s *CreateEdgeContainerAppVersionRequestContainers) GetPostStart() *string {
	return s.PostStart
}

func (s *CreateEdgeContainerAppVersionRequestContainers) GetPreStop() *string {
	return s.PreStop
}

func (s *CreateEdgeContainerAppVersionRequestContainers) GetProbeContent() *CreateEdgeContainerAppVersionRequestContainersProbeContent {
	return s.ProbeContent
}

func (s *CreateEdgeContainerAppVersionRequestContainers) GetProbeType() *string {
	return s.ProbeType
}

func (s *CreateEdgeContainerAppVersionRequestContainers) GetSpec() *string {
	return s.Spec
}

func (s *CreateEdgeContainerAppVersionRequestContainers) GetStorage() *string {
	return s.Storage
}

func (s *CreateEdgeContainerAppVersionRequestContainers) SetACRImageInfo(v *CreateEdgeContainerAppVersionRequestContainersACRImageInfo) *CreateEdgeContainerAppVersionRequestContainers {
	s.ACRImageInfo = v
	return s
}

func (s *CreateEdgeContainerAppVersionRequestContainers) SetArgs(v string) *CreateEdgeContainerAppVersionRequestContainers {
	s.Args = &v
	return s
}

func (s *CreateEdgeContainerAppVersionRequestContainers) SetCommand(v string) *CreateEdgeContainerAppVersionRequestContainers {
	s.Command = &v
	return s
}

func (s *CreateEdgeContainerAppVersionRequestContainers) SetEnvVariables(v string) *CreateEdgeContainerAppVersionRequestContainers {
	s.EnvVariables = &v
	return s
}

func (s *CreateEdgeContainerAppVersionRequestContainers) SetImage(v string) *CreateEdgeContainerAppVersionRequestContainers {
	s.Image = &v
	return s
}

func (s *CreateEdgeContainerAppVersionRequestContainers) SetIsACRImage(v bool) *CreateEdgeContainerAppVersionRequestContainers {
	s.IsACRImage = &v
	return s
}

func (s *CreateEdgeContainerAppVersionRequestContainers) SetName(v string) *CreateEdgeContainerAppVersionRequestContainers {
	s.Name = &v
	return s
}

func (s *CreateEdgeContainerAppVersionRequestContainers) SetPostStart(v string) *CreateEdgeContainerAppVersionRequestContainers {
	s.PostStart = &v
	return s
}

func (s *CreateEdgeContainerAppVersionRequestContainers) SetPreStop(v string) *CreateEdgeContainerAppVersionRequestContainers {
	s.PreStop = &v
	return s
}

func (s *CreateEdgeContainerAppVersionRequestContainers) SetProbeContent(v *CreateEdgeContainerAppVersionRequestContainersProbeContent) *CreateEdgeContainerAppVersionRequestContainers {
	s.ProbeContent = v
	return s
}

func (s *CreateEdgeContainerAppVersionRequestContainers) SetProbeType(v string) *CreateEdgeContainerAppVersionRequestContainers {
	s.ProbeType = &v
	return s
}

func (s *CreateEdgeContainerAppVersionRequestContainers) SetSpec(v string) *CreateEdgeContainerAppVersionRequestContainers {
	s.Spec = &v
	return s
}

func (s *CreateEdgeContainerAppVersionRequestContainers) SetStorage(v string) *CreateEdgeContainerAppVersionRequestContainers {
	s.Storage = &v
	return s
}

func (s *CreateEdgeContainerAppVersionRequestContainers) Validate() error {
	return dara.Validate(s)
}

type CreateEdgeContainerAppVersionRequestContainersACRImageInfo struct {
	// The domain name of the Container Registry image.
	//
	// example:
	//
	// 1500.***.net
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the Container Registry instance.
	//
	// example:
	//
	// xcdn-9axbo****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether the image is an enterprise-level Container Registry image.
	//
	// example:
	//
	// false
	IsEnterpriseRegistry *bool `json:"IsEnterpriseRegistry,omitempty" xml:"IsEnterpriseRegistry,omitempty"`
	// The regions in which the Container Registry instance resides.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the image repository.
	//
	// example:
	//
	// crr-h1ghghu60ct****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The name of the image repository.
	//
	// example:
	//
	// test_71
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The namespace to which the image repository belongs.
	//
	// example:
	//
	// safeline
	RepoNamespace *string `json:"RepoNamespace,omitempty" xml:"RepoNamespace,omitempty"`
	// The tag of the Container Registry image.
	//
	// example:
	//
	// 3.40.2
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The URL of the Container Registry image tag.
	TagUrl *string `json:"TagUrl,omitempty" xml:"TagUrl,omitempty"`
}

func (s CreateEdgeContainerAppVersionRequestContainersACRImageInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateEdgeContainerAppVersionRequestContainersACRImageInfo) GoString() string {
	return s.String()
}

func (s *CreateEdgeContainerAppVersionRequestContainersACRImageInfo) GetDomain() *string {
	return s.Domain
}

func (s *CreateEdgeContainerAppVersionRequestContainersACRImageInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateEdgeContainerAppVersionRequestContainersACRImageInfo) GetIsEnterpriseRegistry() *bool {
	return s.IsEnterpriseRegistry
}

func (s *CreateEdgeContainerAppVersionRequestContainersACRImageInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateEdgeContainerAppVersionRequestContainersACRImageInfo) GetRepoId() *string {
	return s.RepoId
}

func (s *CreateEdgeContainerAppVersionRequestContainersACRImageInfo) GetRepoName() *string {
	return s.RepoName
}

func (s *CreateEdgeContainerAppVersionRequestContainersACRImageInfo) GetRepoNamespace() *string {
	return s.RepoNamespace
}

func (s *CreateEdgeContainerAppVersionRequestContainersACRImageInfo) GetTag() *string {
	return s.Tag
}

func (s *CreateEdgeContainerAppVersionRequestContainersACRImageInfo) GetTagUrl() *string {
	return s.TagUrl
}

func (s *CreateEdgeContainerAppVersionRequestContainersACRImageInfo) SetDomain(v string) *CreateEdgeContainerAppVersionRequestContainersACRImageInfo {
	s.Domain = &v
	return s
}

func (s *CreateEdgeContainerAppVersionRequestContainersACRImageInfo) SetInstanceId(v string) *CreateEdgeContainerAppVersionRequestContainersACRImageInfo {
	s.InstanceId = &v
	return s
}

func (s *CreateEdgeContainerAppVersionRequestContainersACRImageInfo) SetIsEnterpriseRegistry(v bool) *CreateEdgeContainerAppVersionRequestContainersACRImageInfo {
	s.IsEnterpriseRegistry = &v
	return s
}

func (s *CreateEdgeContainerAppVersionRequestContainersACRImageInfo) SetRegionId(v string) *CreateEdgeContainerAppVersionRequestContainersACRImageInfo {
	s.RegionId = &v
	return s
}

func (s *CreateEdgeContainerAppVersionRequestContainersACRImageInfo) SetRepoId(v string) *CreateEdgeContainerAppVersionRequestContainersACRImageInfo {
	s.RepoId = &v
	return s
}

func (s *CreateEdgeContainerAppVersionRequestContainersACRImageInfo) SetRepoName(v string) *CreateEdgeContainerAppVersionRequestContainersACRImageInfo {
	s.RepoName = &v
	return s
}

func (s *CreateEdgeContainerAppVersionRequestContainersACRImageInfo) SetRepoNamespace(v string) *CreateEdgeContainerAppVersionRequestContainersACRImageInfo {
	s.RepoNamespace = &v
	return s
}

func (s *CreateEdgeContainerAppVersionRequestContainersACRImageInfo) SetTag(v string) *CreateEdgeContainerAppVersionRequestContainersACRImageInfo {
	s.Tag = &v
	return s
}

func (s *CreateEdgeContainerAppVersionRequestContainersACRImageInfo) SetTagUrl(v string) *CreateEdgeContainerAppVersionRequestContainersACRImageInfo {
	s.TagUrl = &v
	return s
}

func (s *CreateEdgeContainerAppVersionRequestContainersACRImageInfo) Validate() error {
	return dara.Validate(s)
}

type CreateEdgeContainerAppVersionRequestContainersProbeContent struct {
	// The command of the exec type probe.
	//
	// example:
	//
	// echo ok
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The number of consecutive failed health checks required for a container to be considered as unhealthy.
	//
	// example:
	//
	// 3
	FailureThreshold *int32 `json:"FailureThreshold,omitempty" xml:"FailureThreshold,omitempty"`
	// The domain name that is used for health checks.
	//
	// example:
	//
	// www.rewrite.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The request headers that are included in the container health check request.
	//
	// example:
	//
	// [{\\"Content-Type\\":\\"application/json\\"}]
	HttpHeaders *string `json:"HttpHeaders,omitempty" xml:"HttpHeaders,omitempty"`
	// The latency for container probe initialization.
	//
	// example:
	//
	// 1
	InitialDelaySeconds *int32 `json:"InitialDelaySeconds,omitempty" xml:"InitialDelaySeconds,omitempty"`
	// The health check path.
	//
	// example:
	//
	// /
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The interval between container health checks.
	//
	// example:
	//
	// 1
	PeriodSeconds *int32 `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	// The health check port.
	//
	// example:
	//
	// 9991
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol that the container health check request uses.
	//
	// example:
	//
	// http
	Scheme *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
	// The number of consecutive successful health checks required for a container to be considered as healthy.
	//
	// example:
	//
	// 1
	SuccessThreshold *int32 `json:"SuccessThreshold,omitempty" xml:"SuccessThreshold,omitempty"`
	// The timeout period of the container health check.
	//
	// example:
	//
	// 1
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s CreateEdgeContainerAppVersionRequestContainersProbeContent) String() string {
	return dara.Prettify(s)
}

func (s CreateEdgeContainerAppVersionRequestContainersProbeContent) GoString() string {
	return s.String()
}

func (s *CreateEdgeContainerAppVersionRequestContainersProbeContent) GetCommand() *string {
	return s.Command
}

func (s *CreateEdgeContainerAppVersionRequestContainersProbeContent) GetFailureThreshold() *int32 {
	return s.FailureThreshold
}

func (s *CreateEdgeContainerAppVersionRequestContainersProbeContent) GetHost() *string {
	return s.Host
}

func (s *CreateEdgeContainerAppVersionRequestContainersProbeContent) GetHttpHeaders() *string {
	return s.HttpHeaders
}

func (s *CreateEdgeContainerAppVersionRequestContainersProbeContent) GetInitialDelaySeconds() *int32 {
	return s.InitialDelaySeconds
}

func (s *CreateEdgeContainerAppVersionRequestContainersProbeContent) GetPath() *string {
	return s.Path
}

func (s *CreateEdgeContainerAppVersionRequestContainersProbeContent) GetPeriodSeconds() *int32 {
	return s.PeriodSeconds
}

func (s *CreateEdgeContainerAppVersionRequestContainersProbeContent) GetPort() *int32 {
	return s.Port
}

func (s *CreateEdgeContainerAppVersionRequestContainersProbeContent) GetScheme() *string {
	return s.Scheme
}

func (s *CreateEdgeContainerAppVersionRequestContainersProbeContent) GetSuccessThreshold() *int32 {
	return s.SuccessThreshold
}

func (s *CreateEdgeContainerAppVersionRequestContainersProbeContent) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *CreateEdgeContainerAppVersionRequestContainersProbeContent) SetCommand(v string) *CreateEdgeContainerAppVersionRequestContainersProbeContent {
	s.Command = &v
	return s
}

func (s *CreateEdgeContainerAppVersionRequestContainersProbeContent) SetFailureThreshold(v int32) *CreateEdgeContainerAppVersionRequestContainersProbeContent {
	s.FailureThreshold = &v
	return s
}

func (s *CreateEdgeContainerAppVersionRequestContainersProbeContent) SetHost(v string) *CreateEdgeContainerAppVersionRequestContainersProbeContent {
	s.Host = &v
	return s
}

func (s *CreateEdgeContainerAppVersionRequestContainersProbeContent) SetHttpHeaders(v string) *CreateEdgeContainerAppVersionRequestContainersProbeContent {
	s.HttpHeaders = &v
	return s
}

func (s *CreateEdgeContainerAppVersionRequestContainersProbeContent) SetInitialDelaySeconds(v int32) *CreateEdgeContainerAppVersionRequestContainersProbeContent {
	s.InitialDelaySeconds = &v
	return s
}

func (s *CreateEdgeContainerAppVersionRequestContainersProbeContent) SetPath(v string) *CreateEdgeContainerAppVersionRequestContainersProbeContent {
	s.Path = &v
	return s
}

func (s *CreateEdgeContainerAppVersionRequestContainersProbeContent) SetPeriodSeconds(v int32) *CreateEdgeContainerAppVersionRequestContainersProbeContent {
	s.PeriodSeconds = &v
	return s
}

func (s *CreateEdgeContainerAppVersionRequestContainersProbeContent) SetPort(v int32) *CreateEdgeContainerAppVersionRequestContainersProbeContent {
	s.Port = &v
	return s
}

func (s *CreateEdgeContainerAppVersionRequestContainersProbeContent) SetScheme(v string) *CreateEdgeContainerAppVersionRequestContainersProbeContent {
	s.Scheme = &v
	return s
}

func (s *CreateEdgeContainerAppVersionRequestContainersProbeContent) SetSuccessThreshold(v int32) *CreateEdgeContainerAppVersionRequestContainersProbeContent {
	s.SuccessThreshold = &v
	return s
}

func (s *CreateEdgeContainerAppVersionRequestContainersProbeContent) SetTimeoutSeconds(v int32) *CreateEdgeContainerAppVersionRequestContainersProbeContent {
	s.TimeoutSeconds = &v
	return s
}

func (s *CreateEdgeContainerAppVersionRequestContainersProbeContent) Validate() error {
	return dara.Validate(s)
}
