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
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F61CDR30-E83C-4FDA-BF73-9A94CDD44229
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The versions.
	Versions []*ListEdgeContainerAppVersionsResponseBodyVersions `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
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
	return dara.Validate(s)
}

type ListEdgeContainerAppVersionsResponseBodyVersions struct {
	// The application ID.
	//
	// example:
	//
	// app-88068867578379****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The containers in the version.
	Containers []*ListEdgeContainerAppVersionsResponseBodyVersionsContainers `json:"Containers,omitempty" xml:"Containers,omitempty" type:"Repeated"`
	// The time when the version was created. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-11-10T02:53:16Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the version was last released. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-02-10T02:48:36Z
	LastPublishTime *string `json:"LastPublishTime,omitempty" xml:"LastPublishTime,omitempty"`
	// The version name.
	//
	// example:
	//
	// version01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The time when the version was released. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-02-10T02:48:36Z
	PublishTime *string `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	// The remarks.
	//
	// example:
	//
	// test version
	Remarks *string `json:"Remarks,omitempty" xml:"Remarks,omitempty"`
	// The status of the current version. Valid values:
	//
	// 	- **created**
	//
	// 	- **failed**
	//
	// 	- **creating**
	//
	// example:
	//
	// created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the version was last modified. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-04-16 10:51:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The version ID.
	//
	// example:
	//
	// ver-87962637161651****
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
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
	return dara.Validate(s)
}

type ListEdgeContainerAppVersionsResponseBodyVersionsContainers struct {
	// The arguments that are passed to the container startup command.
	//
	// example:
	//
	// -c a=1
	Args *string `json:"Args,omitempty" xml:"Args,omitempty"`
	// The command that is used to start the container.
	//
	// example:
	//
	// openresty -g \\"daemon off;\\"
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The environment variables of the container.
	//
	// example:
	//
	// ENV=prod
	EnvVariables *string `json:"EnvVariables,omitempty" xml:"EnvVariables,omitempty"`
	// The address of the container image.
	//
	// example:
	//
	// nginx
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The container name.
	//
	// example:
	//
	// container1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The command that is run before the container is started. Format: `{"exec":{"command":["cat","/etc/group"\\]}}`.
	//
	// If you want to cancel this configuration, set the parameter value to `""` or `{}`. If you do not specify this parameter, this configuration is ignored.
	//
	// example:
	//
	// {\\"exec\\":{\\"command\\":[\\"bash\\",\\"-c\\",\\"cd /home/admin/
	PostStart *string `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	// The command that is run before the container is stopped.
	//
	// example:
	//
	// sh stop.sh
	PreStop *string `json:"PreStop,omitempty" xml:"PreStop,omitempty"`
	// The container probe content.
	ProbeContent *ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent `json:"ProbeContent,omitempty" xml:"ProbeContent,omitempty" type:"Struct"`
	// The probe type of the container.
	//
	// example:
	//
	// httpGet
	ProbeType *string `json:"ProbeType,omitempty" xml:"ProbeType,omitempty"`
	// The compute specification of the container.
	//
	// example:
	//
	// 1C2G
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
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
	return dara.Validate(s)
}

type ListEdgeContainerAppVersionsResponseBodyVersionsContainersProbeContent struct {
	// The probe command.
	//
	// example:
	//
	// openresty -g  "daemon off;
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
	// test.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The request headers that are included in the container health check request.
	//
	// example:
	//
	// {\\"Content-Type\\":\\"application/json\\"}
	HttpHeaders *string `json:"HttpHeaders,omitempty" xml:"HttpHeaders,omitempty"`
	// The latency for container probe initialization.
	//
	// example:
	//
	// 10
	InitialDelaySeconds *int32 `json:"InitialDelaySeconds,omitempty" xml:"InitialDelaySeconds,omitempty"`
	// The path of the container health check.
	//
	// example:
	//
	// /health_check
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The interval between container health checks.
	//
	// example:
	//
	// 5
	PeriodSeconds *int32 `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	// The port of the container health check.
	//
	// example:
	//
	// 80
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
	// 30
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
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
