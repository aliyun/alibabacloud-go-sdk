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
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	Containers []*CreateEdgeContainerAppVersionRequestContainers `json:"Containers,omitempty" xml:"Containers,omitempty" type:"Repeated"`
	// This parameter is required.
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	ACRImageInfo *CreateEdgeContainerAppVersionRequestContainersACRImageInfo `json:"ACRImageInfo,omitempty" xml:"ACRImageInfo,omitempty" type:"Struct"`
	Args         *string                                                     `json:"Args,omitempty" xml:"Args,omitempty"`
	Command      *string                                                     `json:"Command,omitempty" xml:"Command,omitempty"`
	EnvVariables *string                                                     `json:"EnvVariables,omitempty" xml:"EnvVariables,omitempty"`
	// This parameter is required.
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// This parameter is required.
	IsACRImage *bool `json:"IsACRImage,omitempty" xml:"IsACRImage,omitempty"`
	// This parameter is required.
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PostStart *string `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	PreStop   *string `json:"PreStop,omitempty" xml:"PreStop,omitempty"`
	// This parameter is required.
	ProbeContent *CreateEdgeContainerAppVersionRequestContainersProbeContent `json:"ProbeContent,omitempty" xml:"ProbeContent,omitempty" type:"Struct"`
	// This parameter is required.
	ProbeType *string `json:"ProbeType,omitempty" xml:"ProbeType,omitempty"`
	// This parameter is required.
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// This parameter is required.
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
	Domain               *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IsEnterpriseRegistry *bool   `json:"IsEnterpriseRegistry,omitempty" xml:"IsEnterpriseRegistry,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RepoId               *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	RepoName             *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	RepoNamespace        *string `json:"RepoNamespace,omitempty" xml:"RepoNamespace,omitempty"`
	Tag                  *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	TagUrl               *string `json:"TagUrl,omitempty" xml:"TagUrl,omitempty"`
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
	Command             *string `json:"Command,omitempty" xml:"Command,omitempty"`
	FailureThreshold    *int32  `json:"FailureThreshold,omitempty" xml:"FailureThreshold,omitempty"`
	Host                *string `json:"Host,omitempty" xml:"Host,omitempty"`
	HttpHeaders         *string `json:"HttpHeaders,omitempty" xml:"HttpHeaders,omitempty"`
	InitialDelaySeconds *int32  `json:"InitialDelaySeconds,omitempty" xml:"InitialDelaySeconds,omitempty"`
	Path                *string `json:"Path,omitempty" xml:"Path,omitempty"`
	PeriodSeconds       *int32  `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	Port                *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Scheme              *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
	SuccessThreshold    *int32  `json:"SuccessThreshold,omitempty" xml:"SuccessThreshold,omitempty"`
	TimeoutSeconds      *int32  `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
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
