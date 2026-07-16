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
	// The application ID. You can call the [ListEdgeContainerApps](~~ListEdgeContainerApps~~) operation to obtain the application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// app-88068867578379****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The container group to deploy for this version, including specific image information. The image information consists of the image address, startup commands, parameters, environment variables, and probe rules. Multiple images are supported. This parameter is a JSON array.
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
	// The version name. The name must be **6 to 128*	- characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// verson1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The remarks.
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
	if s.Containers != nil {
		for _, item := range s.Containers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateEdgeContainerAppVersionRequestContainers struct {
	// The ACR image information.
	ACRImageInfo *CreateEdgeContainerAppVersionRequestContainersACRImageInfo `json:"ACRImageInfo,omitempty" xml:"ACRImageInfo,omitempty" type:"Struct"`
	// The startup parameters. Separate multiple parameters with spaces.
	//
	// example:
	//
	// -a
	Args *string `json:"Args,omitempty" xml:"Args,omitempty"`
	// The startup command. Separate multiple commands with spaces.
	//
	// example:
	//
	// nginx
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The environment variables, in the format of key1=val1,key2=val2.
	//
	// example:
	//
	// VITE_APP_TITLE=My App
	EnvVariables *string `json:"EnvVariables,omitempty" xml:"EnvVariables,omitempty"`
	// The image address.
	//
	// This parameter is required.
	//
	// example:
	//
	// registry-vpc.cn-shenzhen.aliyuncs.com/lihe****h/ea****ts_serv****am:3.**
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// Specifies whether the image is an Alibaba Cloud Container Registry (ACR) image.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	IsACRImage *bool `json:"IsACRImage,omitempty" xml:"IsACRImage,omitempty"`
	// The container name. The name must be unique within the same container group.
	//
	// This parameter is required.
	//
	// example:
	//
	// lxg-demo-er
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The command to run before the container starts. Separate multiple commands with spaces. This command runs before the service starts and is typically used for initialization operations.
	//
	// example:
	//
	// sh poststart.sh "echo hello world"
	PostStart *string `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	// The command to run before the container stops. Separate multiple commands with spaces. This command runs before the service exits and is typically used for cleanup operations.
	//
	// example:
	//
	// sh prestop.sh "echo hello world"
	PreStop *string `json:"PreStop,omitempty" xml:"PreStop,omitempty"`
	// The container health probe content.
	//
	// This parameter is required.
	//
	// example:
	//
	// For details, see the definition of readiness probes in Kubernetes.
	ProbeContent *CreateEdgeContainerAppVersionRequestContainersProbeContent `json:"ProbeContent,omitempty" xml:"ProbeContent,omitempty" type:"Struct"`
	// The probe type. Valid values:
	//
	// - **exec**: command-based.
	//
	// - **tcpSocket**: TCP-based.
	//
	// - **httpGet**: HTTP-based.
	//
	// This parameter is required.
	//
	// example:
	//
	// exec
	ProbeType *string `json:"ProbeType,omitempty" xml:"ProbeType,omitempty"`
	// The container specifications. This parameter specifies the computing specifications. Valid values: 1C2G, 2C4G, 2C8G, 4C8G, 4C16G, 8C16G, and 8C32G.
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
	if s.ACRImageInfo != nil {
		if err := s.ACRImageInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ProbeContent != nil {
		if err := s.ProbeContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateEdgeContainerAppVersionRequestContainersACRImageInfo struct {
	// The ACR image domain name.
	//
	// example:
	//
	// 1500.***.net
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ACR instance ID.
	//
	// example:
	//
	// xcdn-9axbo****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether the image is an enterprise-level image.
	//
	// example:
	//
	// false
	IsEnterpriseRegistry *bool `json:"IsEnterpriseRegistry,omitempty" xml:"IsEnterpriseRegistry,omitempty"`
	// The list of regions for the ACR instance.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The repository ID of the image.
	//
	// example:
	//
	// crr-h1ghghu60ct****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The image repository name.
	//
	// example:
	//
	// test_71
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The namespace of the image repository.
	//
	// example:
	//
	// safeline
	RepoNamespace *string `json:"RepoNamespace,omitempty" xml:"RepoNamespace,omitempty"`
	// The ACR image tag.
	//
	// example:
	//
	// 3.40.2
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The ACR image tag URL.
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
	// The probe command for the exec probe type.
	//
	// example:
	//
	// echo ok
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The number of consecutive failed health checks required.
	//
	// example:
	//
	// 3
	FailureThreshold *int32 `json:"FailureThreshold,omitempty" xml:"FailureThreshold,omitempty"`
	// The domain name for the health check.
	//
	// example:
	//
	// www.rewrite.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The HTTP request headers.
	//
	// example:
	//
	// [{\\"Content-Type\\":\\"application/json\\"}]
	HttpHeaders *string `json:"HttpHeaders,omitempty" xml:"HttpHeaders,omitempty"`
	// The initial delay before the container probe starts, in seconds.
	//
	// example:
	//
	// 1
	InitialDelaySeconds *int32 `json:"InitialDelaySeconds,omitempty" xml:"InitialDelaySeconds,omitempty"`
	// The path for the container health check.
	//
	// example:
	//
	// /
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The interval between container health checks, in seconds.
	//
	// example:
	//
	// 1
	PeriodSeconds *int32 `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	// The port for the container health check.
	//
	// example:
	//
	// 9991
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The request protocol for the health check.
	//
	// example:
	//
	// http
	Scheme *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
	// The number of consecutive successful health checks required.
	//
	// example:
	//
	// 1
	SuccessThreshold *int32 `json:"SuccessThreshold,omitempty" xml:"SuccessThreshold,omitempty"`
	// The timeout period for the container health check, in seconds.
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
