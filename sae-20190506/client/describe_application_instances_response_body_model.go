// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeApplicationInstancesResponseBody
	GetCode() *string
	SetData(v *DescribeApplicationInstancesResponseBodyData) *DescribeApplicationInstancesResponseBody
	GetData() *DescribeApplicationInstancesResponseBodyData
	SetErrorCode(v string) *DescribeApplicationInstancesResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DescribeApplicationInstancesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeApplicationInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeApplicationInstancesResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DescribeApplicationInstancesResponseBody
	GetTraceId() *string
}

type DescribeApplicationInstancesResponseBody struct {
	// The HTTP status code. Valid values:
	//
	// 	- **2xx**: indicates that the request was successful.
	//
	// 	- **3xx**: indicates that the request was redirected.
	//
	// 	- **4xx**: indicates that the request was invalid.
	//
	// 	- **5xx**: indicates that a server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the application instance.
	Data *DescribeApplicationInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// - The **ErrorCode*	- parameter is not returned when the request succeeds.
	//
	// - The **ErrorCode*	- parameter is returned when the request fails. For more information, see **Error codes*	- in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the list of application instances was obtained. Valid values:
	//
	// 	- **true**: indicates that the list was obtained.
	//
	// 	- **false**: indicates that the list could not be obtained.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the trace. It is used to query the details of a request.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeApplicationInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationInstancesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeApplicationInstancesResponseBody) GetData() *DescribeApplicationInstancesResponseBodyData {
	return s.Data
}

func (s *DescribeApplicationInstancesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeApplicationInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeApplicationInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApplicationInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeApplicationInstancesResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeApplicationInstancesResponseBody) SetCode(v string) *DescribeApplicationInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBody) SetData(v *DescribeApplicationInstancesResponseBodyData) *DescribeApplicationInstancesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeApplicationInstancesResponseBody) SetErrorCode(v string) *DescribeApplicationInstancesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBody) SetMessage(v string) *DescribeApplicationInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBody) SetRequestId(v string) *DescribeApplicationInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBody) SetSuccess(v bool) *DescribeApplicationInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBody) SetTraceId(v string) *DescribeApplicationInstancesResponseBody {
	s.TraceId = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationInstancesResponseBodyData struct {
	// The number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The application instances.
	Instances []*DescribeApplicationInstancesResponseBodyDataInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of returned instances.
	//
	// example:
	//
	// 10
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s DescribeApplicationInstancesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeApplicationInstancesResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeApplicationInstancesResponseBodyData) GetInstances() []*DescribeApplicationInstancesResponseBodyDataInstances {
	return s.Instances
}

func (s *DescribeApplicationInstancesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApplicationInstancesResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *DescribeApplicationInstancesResponseBodyData) SetCurrentPage(v int32) *DescribeApplicationInstancesResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBodyData) SetInstances(v []*DescribeApplicationInstancesResponseBodyDataInstances) *DescribeApplicationInstancesResponseBodyData {
	s.Instances = v
	return s
}

func (s *DescribeApplicationInstancesResponseBodyData) SetPageSize(v int32) *DescribeApplicationInstancesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBodyData) SetTotalSize(v int32) *DescribeApplicationInstancesResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationInstancesResponseBodyDataInstances struct {
	// The start time of the instance creation process. Unit: milliseconds.
	//
	// example:
	//
	// 1558442609000
	CreateTimeStamp *int64 `json:"CreateTimeStamp,omitempty" xml:"CreateTimeStamp,omitempty"`
	// example:
	//
	// False
	DebugStatus *bool `json:"DebugStatus,omitempty" xml:"DebugStatus,omitempty"`
	// The elastic IP address (EIP).
	//
	// example:
	//
	// 8.129.XX.XXX
	Eip *string `json:"Eip,omitempty" xml:"Eip,omitempty"`
	// The end time of the instance creation process. Unit: milliseconds.
	//
	// example:
	//
	// 1642757283000
	FinishTimeStamp *int64 `json:"FinishTimeStamp,omitempty" xml:"FinishTimeStamp,omitempty"`
	// The ID of the group to which the instance belongs.
	//
	// example:
	//
	// b2a8a925-477a-4ed7-b825-d5e22500****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The address of the repository.
	//
	// >  If you deploy the application by using a JAR or WAR package, the image generated by SAE is not available for download.
	//
	// example:
	//
	// registry-vpc.cn-beijing.aliyuncs.com/sae-demo-image/cartservice:1.0
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The internal IP address of the instance.
	//
	// example:
	//
	// 192.168.X.X
	InstanceContainerIp *string `json:"InstanceContainerIp,omitempty" xml:"InstanceContainerIp,omitempty"`
	// The number of times that the instance restarted.
	//
	// example:
	//
	// 0
	InstanceContainerRestarts *int64 `json:"InstanceContainerRestarts,omitempty" xml:"InstanceContainerRestarts,omitempty"`
	// The state of the instance. Valid values:
	//
	// 	- **Error**: An error occurred during the instance startup.
	//
	// 	- **CrashLoopBackOff**: The container failed to start. An error occurred during the startup process and persisted after the restart.
	//
	// 	- **ErrImagePull**: An error occurred while the container image was being pulled from the instance.
	//
	// 	- **ImagePullBackOff**: The container image could not be obtained.
	//
	// 	- **Pending**: The instance is waiting to be scheduled.
	//
	// 	- **Unknown**: An unknown exception occurred.
	//
	// 	- **Terminating**: The instance creation process is being terminated.
	//
	// 	- **NotFound**: The instance cannot be found.
	//
	// 	- **PodInitializing**: The instance is being initialized.
	//
	// 	- **Init:0/1**: The instance is initialized.
	//
	// 	- **Running**: The instance is running.
	//
	// example:
	//
	// Running
	InstanceContainerStatus *string `json:"InstanceContainerStatus,omitempty" xml:"InstanceContainerStatus,omitempty"`
	// The configurations of health checks. Valid values:
	//
	// 	- **WithoutHealthCheckConfig**: Liveness and readiness checks are not configured.
	//
	// 	- **WithoutLivenessConfig**: The liveness check is not configured.
	//
	// 	- **WithoutReadinessConfig**: The readiness check is not configured.
	//
	// 	- **NotCheckedYet**: The health checks are not performed or are in progress.
	//
	// 	- **LivenessUnhealthy**: The instance failed the liveness check.
	//
	// 	- **ReadinessUnhealthy**: The instance failed the readiness check.
	//
	// 	- **Unhealthy**: The instance failed both liveness and readiness checks.
	//
	// 	- **Healthy**: The instance passed both liveness and readiness checks.
	//
	// example:
	//
	// WithoutHealthCheckConfig
	InstanceHealthStatus *string `json:"InstanceHealthStatus,omitempty" xml:"InstanceHealthStatus,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// b2a8a925-477a-4ed7-b825-d5e22500****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of the main container.
	//
	// example:
	//
	// Running
	MainContainerStatus *string `json:"MainContainerStatus,omitempty" xml:"MainContainerStatus,omitempty"`
	// The version of the package.
	//
	// example:
	//
	// 1609939496200
	PackageVersion *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty"`
	// The status of the sidecar container.
	SidecarContainersStatus []*DescribeApplicationInstancesResponseBodyDataInstancesSidecarContainersStatus `json:"SidecarContainersStatus,omitempty" xml:"SidecarContainersStatus,omitempty" type:"Repeated"`
	// example:
	//
	// 1750061980000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// If the health check of an application instance fails, the detailed failure cause or error message is returned. If the health check of an application instance passes, no response is returned.
	//
	// example:
	//
	// Null
	UnhealthyMessage *string `json:"UnhealthyMessage,omitempty" xml:"UnhealthyMessage,omitempty"`
	// The ID of the zone where the instance is deployed.
	//
	// example:
	//
	// vsw-***
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeApplicationInstancesResponseBodyDataInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationInstancesResponseBodyDataInstances) GoString() string {
	return s.String()
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) GetCreateTimeStamp() *int64 {
	return s.CreateTimeStamp
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) GetDebugStatus() *bool {
	return s.DebugStatus
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) GetEip() *string {
	return s.Eip
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) GetFinishTimeStamp() *int64 {
	return s.FinishTimeStamp
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) GetInstanceContainerIp() *string {
	return s.InstanceContainerIp
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) GetInstanceContainerRestarts() *int64 {
	return s.InstanceContainerRestarts
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) GetInstanceContainerStatus() *string {
	return s.InstanceContainerStatus
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) GetInstanceHealthStatus() *string {
	return s.InstanceHealthStatus
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) GetMainContainerStatus() *string {
	return s.MainContainerStatus
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) GetPackageVersion() *string {
	return s.PackageVersion
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) GetSidecarContainersStatus() []*DescribeApplicationInstancesResponseBodyDataInstancesSidecarContainersStatus {
	return s.SidecarContainersStatus
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) GetUnhealthyMessage() *string {
	return s.UnhealthyMessage
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) SetCreateTimeStamp(v int64) *DescribeApplicationInstancesResponseBodyDataInstances {
	s.CreateTimeStamp = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) SetDebugStatus(v bool) *DescribeApplicationInstancesResponseBodyDataInstances {
	s.DebugStatus = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) SetEip(v string) *DescribeApplicationInstancesResponseBodyDataInstances {
	s.Eip = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) SetFinishTimeStamp(v int64) *DescribeApplicationInstancesResponseBodyDataInstances {
	s.FinishTimeStamp = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) SetGroupId(v string) *DescribeApplicationInstancesResponseBodyDataInstances {
	s.GroupId = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) SetImageUrl(v string) *DescribeApplicationInstancesResponseBodyDataInstances {
	s.ImageUrl = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) SetInstanceContainerIp(v string) *DescribeApplicationInstancesResponseBodyDataInstances {
	s.InstanceContainerIp = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) SetInstanceContainerRestarts(v int64) *DescribeApplicationInstancesResponseBodyDataInstances {
	s.InstanceContainerRestarts = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) SetInstanceContainerStatus(v string) *DescribeApplicationInstancesResponseBodyDataInstances {
	s.InstanceContainerStatus = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) SetInstanceHealthStatus(v string) *DescribeApplicationInstancesResponseBodyDataInstances {
	s.InstanceHealthStatus = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) SetInstanceId(v string) *DescribeApplicationInstancesResponseBodyDataInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) SetMainContainerStatus(v string) *DescribeApplicationInstancesResponseBodyDataInstances {
	s.MainContainerStatus = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) SetPackageVersion(v string) *DescribeApplicationInstancesResponseBodyDataInstances {
	s.PackageVersion = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) SetSidecarContainersStatus(v []*DescribeApplicationInstancesResponseBodyDataInstancesSidecarContainersStatus) *DescribeApplicationInstancesResponseBodyDataInstances {
	s.SidecarContainersStatus = v
	return s
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) SetTimestamp(v int64) *DescribeApplicationInstancesResponseBodyDataInstances {
	s.Timestamp = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) SetUnhealthyMessage(v string) *DescribeApplicationInstancesResponseBodyDataInstances {
	s.UnhealthyMessage = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) SetVSwitchId(v string) *DescribeApplicationInstancesResponseBodyDataInstances {
	s.VSwitchId = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationInstancesResponseBodyDataInstancesSidecarContainersStatus struct {
	// The ID of the sidecar container.
	//
	// example:
	//
	// sidecar-test-01
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	// The status of the container.
	//
	// example:
	//
	// Running
	ContainerStatus *string `json:"ContainerStatus,omitempty" xml:"ContainerStatus,omitempty"`
	// The URL of the image.
	//
	// example:
	//
	// registry.cn-beijing.aliyuncs.com/sae-dev-test/******
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s DescribeApplicationInstancesResponseBodyDataInstancesSidecarContainersStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationInstancesResponseBodyDataInstancesSidecarContainersStatus) GoString() string {
	return s.String()
}

func (s *DescribeApplicationInstancesResponseBodyDataInstancesSidecarContainersStatus) GetContainerId() *string {
	return s.ContainerId
}

func (s *DescribeApplicationInstancesResponseBodyDataInstancesSidecarContainersStatus) GetContainerStatus() *string {
	return s.ContainerStatus
}

func (s *DescribeApplicationInstancesResponseBodyDataInstancesSidecarContainersStatus) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *DescribeApplicationInstancesResponseBodyDataInstancesSidecarContainersStatus) SetContainerId(v string) *DescribeApplicationInstancesResponseBodyDataInstancesSidecarContainersStatus {
	s.ContainerId = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBodyDataInstancesSidecarContainersStatus) SetContainerStatus(v string) *DescribeApplicationInstancesResponseBodyDataInstancesSidecarContainersStatus {
	s.ContainerStatus = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBodyDataInstancesSidecarContainersStatus) SetImageUrl(v string) *DescribeApplicationInstancesResponseBodyDataInstancesSidecarContainersStatus {
	s.ImageUrl = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBodyDataInstancesSidecarContainersStatus) Validate() error {
	return dara.Validate(s)
}
