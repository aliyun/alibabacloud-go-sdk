// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeContainerAppVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetEdgeContainerAppVersionResponseBody
	GetRequestId() *string
	SetVersion(v *GetEdgeContainerAppVersionResponseBodyVersion) *GetEdgeContainerAppVersionResponseBody
	GetVersion() *GetEdgeContainerAppVersionResponseBodyVersion
}

type GetEdgeContainerAppVersionResponseBody struct {
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Version   *GetEdgeContainerAppVersionResponseBodyVersion `json:"Version,omitempty" xml:"Version,omitempty" type:"Struct"`
}

func (s GetEdgeContainerAppVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerAppVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerAppVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEdgeContainerAppVersionResponseBody) GetVersion() *GetEdgeContainerAppVersionResponseBodyVersion {
	return s.Version
}

func (s *GetEdgeContainerAppVersionResponseBody) SetRequestId(v string) *GetEdgeContainerAppVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBody) SetVersion(v *GetEdgeContainerAppVersionResponseBodyVersion) *GetEdgeContainerAppVersionResponseBody {
	s.Version = v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBody) Validate() error {
	if s.Version != nil {
		if err := s.Version.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetEdgeContainerAppVersionResponseBodyVersion struct {
	AppId           *string                                                    `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Containers      []*GetEdgeContainerAppVersionResponseBodyVersionContainers `json:"Containers,omitempty" xml:"Containers,omitempty" type:"Repeated"`
	CreateTime      *string                                                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	LastPublishTime *string                                                    `json:"LastPublishTime,omitempty" xml:"LastPublishTime,omitempty"`
	Name            *string                                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	PublishTime     *string                                                    `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	Remarks         *string                                                    `json:"Remarks,omitempty" xml:"Remarks,omitempty"`
	Status          *string                                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime      *string                                                    `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	VersionId       *string                                                    `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetEdgeContainerAppVersionResponseBodyVersion) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerAppVersionResponseBodyVersion) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerAppVersionResponseBodyVersion) GetAppId() *string {
	return s.AppId
}

func (s *GetEdgeContainerAppVersionResponseBodyVersion) GetContainers() []*GetEdgeContainerAppVersionResponseBodyVersionContainers {
	return s.Containers
}

func (s *GetEdgeContainerAppVersionResponseBodyVersion) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetEdgeContainerAppVersionResponseBodyVersion) GetLastPublishTime() *string {
	return s.LastPublishTime
}

func (s *GetEdgeContainerAppVersionResponseBodyVersion) GetName() *string {
	return s.Name
}

func (s *GetEdgeContainerAppVersionResponseBodyVersion) GetPublishTime() *string {
	return s.PublishTime
}

func (s *GetEdgeContainerAppVersionResponseBodyVersion) GetRemarks() *string {
	return s.Remarks
}

func (s *GetEdgeContainerAppVersionResponseBodyVersion) GetStatus() *string {
	return s.Status
}

func (s *GetEdgeContainerAppVersionResponseBodyVersion) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetEdgeContainerAppVersionResponseBodyVersion) GetVersionId() *string {
	return s.VersionId
}

func (s *GetEdgeContainerAppVersionResponseBodyVersion) SetAppId(v string) *GetEdgeContainerAppVersionResponseBodyVersion {
	s.AppId = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersion) SetContainers(v []*GetEdgeContainerAppVersionResponseBodyVersionContainers) *GetEdgeContainerAppVersionResponseBodyVersion {
	s.Containers = v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersion) SetCreateTime(v string) *GetEdgeContainerAppVersionResponseBodyVersion {
	s.CreateTime = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersion) SetLastPublishTime(v string) *GetEdgeContainerAppVersionResponseBodyVersion {
	s.LastPublishTime = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersion) SetName(v string) *GetEdgeContainerAppVersionResponseBodyVersion {
	s.Name = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersion) SetPublishTime(v string) *GetEdgeContainerAppVersionResponseBodyVersion {
	s.PublishTime = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersion) SetRemarks(v string) *GetEdgeContainerAppVersionResponseBodyVersion {
	s.Remarks = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersion) SetStatus(v string) *GetEdgeContainerAppVersionResponseBodyVersion {
	s.Status = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersion) SetUpdateTime(v string) *GetEdgeContainerAppVersionResponseBodyVersion {
	s.UpdateTime = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersion) SetVersionId(v string) *GetEdgeContainerAppVersionResponseBodyVersion {
	s.VersionId = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersion) Validate() error {
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

type GetEdgeContainerAppVersionResponseBodyVersionContainers struct {
	ACRImageInfo *GetEdgeContainerAppVersionResponseBodyVersionContainersACRImageInfo `json:"ACRImageInfo,omitempty" xml:"ACRImageInfo,omitempty" type:"Struct"`
	Args         *string                                                              `json:"Args,omitempty" xml:"Args,omitempty"`
	Command      *string                                                              `json:"Command,omitempty" xml:"Command,omitempty"`
	EnvVariables *string                                                              `json:"EnvVariables,omitempty" xml:"EnvVariables,omitempty"`
	Image        *string                                                              `json:"Image,omitempty" xml:"Image,omitempty"`
	IsACRImage   *bool                                                                `json:"IsACRImage,omitempty" xml:"IsACRImage,omitempty"`
	Name         *string                                                              `json:"Name,omitempty" xml:"Name,omitempty"`
	PostStart    *string                                                              `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	PreStop      *string                                                              `json:"PreStop,omitempty" xml:"PreStop,omitempty"`
	ProbeContent *GetEdgeContainerAppVersionResponseBodyVersionContainersProbeContent `json:"ProbeContent,omitempty" xml:"ProbeContent,omitempty" type:"Struct"`
	ProbeType    *string                                                              `json:"ProbeType,omitempty" xml:"ProbeType,omitempty"`
	Spec         *string                                                              `json:"Spec,omitempty" xml:"Spec,omitempty"`
	Storage      *string                                                              `json:"Storage,omitempty" xml:"Storage,omitempty"`
}

func (s GetEdgeContainerAppVersionResponseBodyVersionContainers) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerAppVersionResponseBodyVersionContainers) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainers) GetACRImageInfo() *GetEdgeContainerAppVersionResponseBodyVersionContainersACRImageInfo {
	return s.ACRImageInfo
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainers) GetArgs() *string {
	return s.Args
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainers) GetCommand() *string {
	return s.Command
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainers) GetEnvVariables() *string {
	return s.EnvVariables
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainers) GetImage() *string {
	return s.Image
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainers) GetIsACRImage() *bool {
	return s.IsACRImage
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainers) GetName() *string {
	return s.Name
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainers) GetPostStart() *string {
	return s.PostStart
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainers) GetPreStop() *string {
	return s.PreStop
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainers) GetProbeContent() *GetEdgeContainerAppVersionResponseBodyVersionContainersProbeContent {
	return s.ProbeContent
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainers) GetProbeType() *string {
	return s.ProbeType
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainers) GetSpec() *string {
	return s.Spec
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainers) GetStorage() *string {
	return s.Storage
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainers) SetACRImageInfo(v *GetEdgeContainerAppVersionResponseBodyVersionContainersACRImageInfo) *GetEdgeContainerAppVersionResponseBodyVersionContainers {
	s.ACRImageInfo = v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainers) SetArgs(v string) *GetEdgeContainerAppVersionResponseBodyVersionContainers {
	s.Args = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainers) SetCommand(v string) *GetEdgeContainerAppVersionResponseBodyVersionContainers {
	s.Command = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainers) SetEnvVariables(v string) *GetEdgeContainerAppVersionResponseBodyVersionContainers {
	s.EnvVariables = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainers) SetImage(v string) *GetEdgeContainerAppVersionResponseBodyVersionContainers {
	s.Image = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainers) SetIsACRImage(v bool) *GetEdgeContainerAppVersionResponseBodyVersionContainers {
	s.IsACRImage = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainers) SetName(v string) *GetEdgeContainerAppVersionResponseBodyVersionContainers {
	s.Name = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainers) SetPostStart(v string) *GetEdgeContainerAppVersionResponseBodyVersionContainers {
	s.PostStart = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainers) SetPreStop(v string) *GetEdgeContainerAppVersionResponseBodyVersionContainers {
	s.PreStop = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainers) SetProbeContent(v *GetEdgeContainerAppVersionResponseBodyVersionContainersProbeContent) *GetEdgeContainerAppVersionResponseBodyVersionContainers {
	s.ProbeContent = v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainers) SetProbeType(v string) *GetEdgeContainerAppVersionResponseBodyVersionContainers {
	s.ProbeType = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainers) SetSpec(v string) *GetEdgeContainerAppVersionResponseBodyVersionContainers {
	s.Spec = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainers) SetStorage(v string) *GetEdgeContainerAppVersionResponseBodyVersionContainers {
	s.Storage = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainers) Validate() error {
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

type GetEdgeContainerAppVersionResponseBodyVersionContainersACRImageInfo struct {
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

func (s GetEdgeContainerAppVersionResponseBodyVersionContainersACRImageInfo) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerAppVersionResponseBodyVersionContainersACRImageInfo) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersACRImageInfo) GetDomain() *string {
	return s.Domain
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersACRImageInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersACRImageInfo) GetIsEnterpriseRegistry() *bool {
	return s.IsEnterpriseRegistry
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersACRImageInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersACRImageInfo) GetRepoId() *string {
	return s.RepoId
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersACRImageInfo) GetRepoName() *string {
	return s.RepoName
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersACRImageInfo) GetRepoNamespace() *string {
	return s.RepoNamespace
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersACRImageInfo) GetTag() *string {
	return s.Tag
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersACRImageInfo) GetTagUrl() *string {
	return s.TagUrl
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersACRImageInfo) SetDomain(v string) *GetEdgeContainerAppVersionResponseBodyVersionContainersACRImageInfo {
	s.Domain = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersACRImageInfo) SetInstanceId(v string) *GetEdgeContainerAppVersionResponseBodyVersionContainersACRImageInfo {
	s.InstanceId = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersACRImageInfo) SetIsEnterpriseRegistry(v bool) *GetEdgeContainerAppVersionResponseBodyVersionContainersACRImageInfo {
	s.IsEnterpriseRegistry = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersACRImageInfo) SetRegionId(v string) *GetEdgeContainerAppVersionResponseBodyVersionContainersACRImageInfo {
	s.RegionId = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersACRImageInfo) SetRepoId(v string) *GetEdgeContainerAppVersionResponseBodyVersionContainersACRImageInfo {
	s.RepoId = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersACRImageInfo) SetRepoName(v string) *GetEdgeContainerAppVersionResponseBodyVersionContainersACRImageInfo {
	s.RepoName = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersACRImageInfo) SetRepoNamespace(v string) *GetEdgeContainerAppVersionResponseBodyVersionContainersACRImageInfo {
	s.RepoNamespace = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersACRImageInfo) SetTag(v string) *GetEdgeContainerAppVersionResponseBodyVersionContainersACRImageInfo {
	s.Tag = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersACRImageInfo) SetTagUrl(v string) *GetEdgeContainerAppVersionResponseBodyVersionContainersACRImageInfo {
	s.TagUrl = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersACRImageInfo) Validate() error {
	return dara.Validate(s)
}

type GetEdgeContainerAppVersionResponseBodyVersionContainersProbeContent struct {
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

func (s GetEdgeContainerAppVersionResponseBodyVersionContainersProbeContent) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerAppVersionResponseBodyVersionContainersProbeContent) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersProbeContent) GetCommand() *string {
	return s.Command
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersProbeContent) GetFailureThreshold() *int32 {
	return s.FailureThreshold
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersProbeContent) GetHost() *string {
	return s.Host
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersProbeContent) GetHttpHeaders() *string {
	return s.HttpHeaders
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersProbeContent) GetInitialDelaySeconds() *int32 {
	return s.InitialDelaySeconds
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersProbeContent) GetPath() *string {
	return s.Path
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersProbeContent) GetPeriodSeconds() *int32 {
	return s.PeriodSeconds
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersProbeContent) GetPort() *int32 {
	return s.Port
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersProbeContent) GetScheme() *string {
	return s.Scheme
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersProbeContent) GetSuccessThreshold() *int32 {
	return s.SuccessThreshold
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersProbeContent) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersProbeContent) SetCommand(v string) *GetEdgeContainerAppVersionResponseBodyVersionContainersProbeContent {
	s.Command = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersProbeContent) SetFailureThreshold(v int32) *GetEdgeContainerAppVersionResponseBodyVersionContainersProbeContent {
	s.FailureThreshold = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersProbeContent) SetHost(v string) *GetEdgeContainerAppVersionResponseBodyVersionContainersProbeContent {
	s.Host = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersProbeContent) SetHttpHeaders(v string) *GetEdgeContainerAppVersionResponseBodyVersionContainersProbeContent {
	s.HttpHeaders = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersProbeContent) SetInitialDelaySeconds(v int32) *GetEdgeContainerAppVersionResponseBodyVersionContainersProbeContent {
	s.InitialDelaySeconds = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersProbeContent) SetPath(v string) *GetEdgeContainerAppVersionResponseBodyVersionContainersProbeContent {
	s.Path = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersProbeContent) SetPeriodSeconds(v int32) *GetEdgeContainerAppVersionResponseBodyVersionContainersProbeContent {
	s.PeriodSeconds = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersProbeContent) SetPort(v int32) *GetEdgeContainerAppVersionResponseBodyVersionContainersProbeContent {
	s.Port = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersProbeContent) SetScheme(v string) *GetEdgeContainerAppVersionResponseBodyVersionContainersProbeContent {
	s.Scheme = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersProbeContent) SetSuccessThreshold(v int32) *GetEdgeContainerAppVersionResponseBodyVersionContainersProbeContent {
	s.SuccessThreshold = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersProbeContent) SetTimeoutSeconds(v int32) *GetEdgeContainerAppVersionResponseBodyVersionContainersProbeContent {
	s.TimeoutSeconds = &v
	return s
}

func (s *GetEdgeContainerAppVersionResponseBodyVersionContainersProbeContent) Validate() error {
	return dara.Validate(s)
}
