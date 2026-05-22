// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEdgeContainerAppVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListEdgeContainerAppVersionsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListEdgeContainerAppVersionsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListEdgeContainerAppVersionsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListEdgeContainerAppVersionsResponseBody
	GetTotalCount() *int32
	SetVersions(v []*ListEdgeContainerAppVersionsResponseBodyVersions) *ListEdgeContainerAppVersionsResponseBody
	GetVersions() []*ListEdgeContainerAppVersionsResponseBodyVersions
}

type ListEdgeContainerAppVersionsResponseBody struct {
	PageNumber *int32                                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Versions   []*ListEdgeContainerAppVersionsResponseBodyVersions `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
}

func (s ListEdgeContainerAppVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeContainerAppVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEdgeContainerAppVersionsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListEdgeContainerAppVersionsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEdgeContainerAppVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEdgeContainerAppVersionsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListEdgeContainerAppVersionsResponseBody) GetVersions() []*ListEdgeContainerAppVersionsResponseBodyVersions {
	return s.Versions
}

func (s *ListEdgeContainerAppVersionsResponseBody) SetPageNumber(v int32) *ListEdgeContainerAppVersionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListEdgeContainerAppVersionsResponseBody) SetPageSize(v int32) *ListEdgeContainerAppVersionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListEdgeContainerAppVersionsResponseBody) SetRequestId(v string) *ListEdgeContainerAppVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEdgeContainerAppVersionsResponseBody) SetTotalCount(v int32) *ListEdgeContainerAppVersionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListEdgeContainerAppVersionsResponseBody) SetVersions(v []*ListEdgeContainerAppVersionsResponseBodyVersions) *ListEdgeContainerAppVersionsResponseBody {
	s.Versions = v
	return s
}

func (s *ListEdgeContainerAppVersionsResponseBody) Validate() error {
	if s.Versions != nil {
		for _, item := range s.Versions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEdgeContainerAppVersionsResponseBodyVersions struct {
	AppId           *string                                                       `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Containers      []*ListEdgeContainerAppVersionsResponseBodyVersionsContainers `json:"Containers,omitempty" xml:"Containers,omitempty" type:"Repeated"`
	CreateTime      *string                                                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	LastPublishTime *string                                                       `json:"LastPublishTime,omitempty" xml:"LastPublishTime,omitempty"`
	Name            *string                                                       `json:"Name,omitempty" xml:"Name,omitempty"`
	PublishTime     *string                                                       `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	Remarks         *string                                                       `json:"Remarks,omitempty" xml:"Remarks,omitempty"`
	Status          *string                                                       `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime      *string                                                       `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	VersionId       *string                                                       `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s ListEdgeContainerAppVersionsResponseBodyVersions) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeContainerAppVersionsResponseBodyVersions) GoString() string {
	return s.String()
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersions) GetAppId() *string {
	return s.AppId
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersions) GetContainers() []*ListEdgeContainerAppVersionsResponseBodyVersionsContainers {
	return s.Containers
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersions) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersions) GetLastPublishTime() *string {
	return s.LastPublishTime
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersions) GetName() *string {
	return s.Name
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersions) GetPublishTime() *string {
	return s.PublishTime
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersions) GetRemarks() *string {
	return s.Remarks
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersions) GetStatus() *string {
	return s.Status
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersions) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersions) GetVersionId() *string {
	return s.VersionId
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersions) SetAppId(v string) *ListEdgeContainerAppVersionsResponseBodyVersions {
	s.AppId = &v
	return s
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersions) SetContainers(v []*ListEdgeContainerAppVersionsResponseBodyVersionsContainers) *ListEdgeContainerAppVersionsResponseBodyVersions {
	s.Containers = v
	return s
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersions) SetCreateTime(v string) *ListEdgeContainerAppVersionsResponseBodyVersions {
	s.CreateTime = &v
	return s
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersions) SetLastPublishTime(v string) *ListEdgeContainerAppVersionsResponseBodyVersions {
	s.LastPublishTime = &v
	return s
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersions) SetName(v string) *ListEdgeContainerAppVersionsResponseBodyVersions {
	s.Name = &v
	return s
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersions) SetPublishTime(v string) *ListEdgeContainerAppVersionsResponseBodyVersions {
	s.PublishTime = &v
	return s
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersions) SetRemarks(v string) *ListEdgeContainerAppVersionsResponseBodyVersions {
	s.Remarks = &v
	return s
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersions) SetStatus(v string) *ListEdgeContainerAppVersionsResponseBodyVersions {
	s.Status = &v
	return s
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersions) SetUpdateTime(v string) *ListEdgeContainerAppVersionsResponseBodyVersions {
	s.UpdateTime = &v
	return s
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersions) SetVersionId(v string) *ListEdgeContainerAppVersionsResponseBodyVersions {
	s.VersionId = &v
	return s
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersions) Validate() error {
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

type ListEdgeContainerAppVersionsResponseBodyVersionsContainers struct {
	Args         *string                                                                 `json:"Args,omitempty" xml:"Args,omitempty"`
	Command      *string                                                                 `json:"Command,omitempty" xml:"Command,omitempty"`
	EnvVariables *string                                                                 `json:"EnvVariables,omitempty" xml:"EnvVariables,omitempty"`
	Image        *string                                                                 `json:"Image,omitempty" xml:"Image,omitempty"`
	Name         *string                                                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	PostStart    *string                                                                 `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	PreStop      *string                                                                 `json:"PreStop,omitempty" xml:"PreStop,omitempty"`
	ProbeContent *ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent `json:"ProbeContent,omitempty" xml:"ProbeContent,omitempty" type:"Struct"`
	ProbeType    *string                                                                 `json:"ProbeType,omitempty" xml:"ProbeType,omitempty"`
	Spec         *string                                                                 `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s ListEdgeContainerAppVersionsResponseBodyVersionsContainers) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeContainerAppVersionsResponseBodyVersionsContainers) GoString() string {
	return s.String()
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainers) GetArgs() *string {
	return s.Args
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainers) GetCommand() *string {
	return s.Command
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainers) GetEnvVariables() *string {
	return s.EnvVariables
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainers) GetImage() *string {
	return s.Image
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainers) GetName() *string {
	return s.Name
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainers) GetPostStart() *string {
	return s.PostStart
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainers) GetPreStop() *string {
	return s.PreStop
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainers) GetProbeContent() *ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent {
	return s.ProbeContent
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainers) GetProbeType() *string {
	return s.ProbeType
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainers) GetSpec() *string {
	return s.Spec
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainers) SetArgs(v string) *ListEdgeContainerAppVersionsResponseBodyVersionsContainers {
	s.Args = &v
	return s
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainers) SetCommand(v string) *ListEdgeContainerAppVersionsResponseBodyVersionsContainers {
	s.Command = &v
	return s
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainers) SetEnvVariables(v string) *ListEdgeContainerAppVersionsResponseBodyVersionsContainers {
	s.EnvVariables = &v
	return s
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainers) SetImage(v string) *ListEdgeContainerAppVersionsResponseBodyVersionsContainers {
	s.Image = &v
	return s
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainers) SetName(v string) *ListEdgeContainerAppVersionsResponseBodyVersionsContainers {
	s.Name = &v
	return s
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainers) SetPostStart(v string) *ListEdgeContainerAppVersionsResponseBodyVersionsContainers {
	s.PostStart = &v
	return s
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainers) SetPreStop(v string) *ListEdgeContainerAppVersionsResponseBodyVersionsContainers {
	s.PreStop = &v
	return s
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainers) SetProbeContent(v *ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent) *ListEdgeContainerAppVersionsResponseBodyVersionsContainers {
	s.ProbeContent = v
	return s
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainers) SetProbeType(v string) *ListEdgeContainerAppVersionsResponseBodyVersionsContainers {
	s.ProbeType = &v
	return s
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainers) SetSpec(v string) *ListEdgeContainerAppVersionsResponseBodyVersionsContainers {
	s.Spec = &v
	return s
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainers) Validate() error {
	if s.ProbeContent != nil {
		if err := s.ProbeContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent struct {
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

func (s ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent) GoString() string {
	return s.String()
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent) GetCommand() *string {
	return s.Command
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent) GetFailureThreshold() *int32 {
	return s.FailureThreshold
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent) GetHost() *string {
	return s.Host
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent) GetHttpHeaders() *string {
	return s.HttpHeaders
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent) GetInitialDelaySeconds() *int32 {
	return s.InitialDelaySeconds
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent) GetPath() *string {
	return s.Path
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent) GetPeriodSeconds() *int32 {
	return s.PeriodSeconds
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent) GetPort() *int32 {
	return s.Port
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent) GetScheme() *string {
	return s.Scheme
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent) GetSuccessThreshold() *int32 {
	return s.SuccessThreshold
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent) SetCommand(v string) *ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent {
	s.Command = &v
	return s
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent) SetFailureThreshold(v int32) *ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent {
	s.FailureThreshold = &v
	return s
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent) SetHost(v string) *ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent {
	s.Host = &v
	return s
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent) SetHttpHeaders(v string) *ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent {
	s.HttpHeaders = &v
	return s
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent) SetInitialDelaySeconds(v int32) *ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent {
	s.InitialDelaySeconds = &v
	return s
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent) SetPath(v string) *ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent {
	s.Path = &v
	return s
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent) SetPeriodSeconds(v int32) *ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent {
	s.PeriodSeconds = &v
	return s
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent) SetPort(v int32) *ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent {
	s.Port = &v
	return s
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent) SetScheme(v string) *ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent {
	s.Scheme = &v
	return s
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent) SetSuccessThreshold(v int32) *ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent {
	s.SuccessThreshold = &v
	return s
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent) SetTimeoutSeconds(v int32) *ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent {
	s.TimeoutSeconds = &v
	return s
}

func (s *ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent) Validate() error {
	return dara.Validate(s)
}
