// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServcieScheduleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIndex(v int32) *DescribeServcieScheduleResponseBody
	GetIndex() *int32
	SetInstanceId(v string) *DescribeServcieScheduleResponseBody
	GetInstanceId() *string
	SetInstanceIp(v string) *DescribeServcieScheduleResponseBody
	GetInstanceIp() *string
	SetInstancePort(v int32) *DescribeServcieScheduleResponseBody
	GetInstancePort() *int32
	SetPodAbstractInfo(v *DescribeServcieScheduleResponseBodyPodAbstractInfo) *DescribeServcieScheduleResponseBody
	GetPodAbstractInfo() *DescribeServcieScheduleResponseBodyPodAbstractInfo
	SetRequestId(v string) *DescribeServcieScheduleResponseBody
	GetRequestId() *string
	SetRequestRepeated(v bool) *DescribeServcieScheduleResponseBody
	GetRequestRepeated() *bool
	SetTcpPorts(v string) *DescribeServcieScheduleResponseBody
	GetTcpPorts() *string
}

type DescribeServcieScheduleResponseBody struct {
	// The index number of the scheduled virtual device (pod).
	//
	// example:
	//
	// 2
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The ID of the scheduled instance.
	//
	// example:
	//
	// i-5myukg7hnpbto7m024002****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the scheduled instance.
	//
	// example:
	//
	// 120.26.XX.XX
	InstanceIp *string `json:"InstanceIp,omitempty" xml:"InstanceIp,omitempty"`
	// The start port of the scheduled instance.
	//
	// example:
	//
	// 3306
	InstancePort *int32 `json:"InstancePort,omitempty" xml:"InstancePort,omitempty"`
	// The summary information about the scheduled virtual device.
	PodAbstractInfo *DescribeServcieScheduleResponseBodyPodAbstractInfo `json:"PodAbstractInfo,omitempty" xml:"PodAbstractInfo,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is repeated.
	//
	// example:
	//
	// false
	RequestRepeated *bool `json:"RequestRepeated,omitempty" xml:"RequestRepeated,omitempty"`
	// The TCP port range of the scheduled instance or container. The value is in the ${from}-$-{to} format. Example: 80-88.
	//
	// example:
	//
	// 80-88
	TcpPorts *string `json:"TcpPorts,omitempty" xml:"TcpPorts,omitempty"`
}

func (s DescribeServcieScheduleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeServcieScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServcieScheduleResponseBody) GetIndex() *int32 {
	return s.Index
}

func (s *DescribeServcieScheduleResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeServcieScheduleResponseBody) GetInstanceIp() *string {
	return s.InstanceIp
}

func (s *DescribeServcieScheduleResponseBody) GetInstancePort() *int32 {
	return s.InstancePort
}

func (s *DescribeServcieScheduleResponseBody) GetPodAbstractInfo() *DescribeServcieScheduleResponseBodyPodAbstractInfo {
	return s.PodAbstractInfo
}

func (s *DescribeServcieScheduleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeServcieScheduleResponseBody) GetRequestRepeated() *bool {
	return s.RequestRepeated
}

func (s *DescribeServcieScheduleResponseBody) GetTcpPorts() *string {
	return s.TcpPorts
}

func (s *DescribeServcieScheduleResponseBody) SetIndex(v int32) *DescribeServcieScheduleResponseBody {
	s.Index = &v
	return s
}

func (s *DescribeServcieScheduleResponseBody) SetInstanceId(v string) *DescribeServcieScheduleResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeServcieScheduleResponseBody) SetInstanceIp(v string) *DescribeServcieScheduleResponseBody {
	s.InstanceIp = &v
	return s
}

func (s *DescribeServcieScheduleResponseBody) SetInstancePort(v int32) *DescribeServcieScheduleResponseBody {
	s.InstancePort = &v
	return s
}

func (s *DescribeServcieScheduleResponseBody) SetPodAbstractInfo(v *DescribeServcieScheduleResponseBodyPodAbstractInfo) *DescribeServcieScheduleResponseBody {
	s.PodAbstractInfo = v
	return s
}

func (s *DescribeServcieScheduleResponseBody) SetRequestId(v string) *DescribeServcieScheduleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServcieScheduleResponseBody) SetRequestRepeated(v bool) *DescribeServcieScheduleResponseBody {
	s.RequestRepeated = &v
	return s
}

func (s *DescribeServcieScheduleResponseBody) SetTcpPorts(v string) *DescribeServcieScheduleResponseBody {
	s.TcpPorts = &v
	return s
}

func (s *DescribeServcieScheduleResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeServcieScheduleResponseBodyPodAbstractInfo struct {
	// The name of the container service.
	//
	// example:
	//
	// android
	ContainerService *bool `json:"ContainerService,omitempty" xml:"ContainerService,omitempty"`
	// The information about the container.
	ContainerStatuses *DescribeServcieScheduleResponseBodyPodAbstractInfoContainerStatuses `json:"ContainerStatuses,omitempty" xml:"ContainerStatuses,omitempty" type:"Struct"`
	// The name of the pod.
	//
	// example:
	//
	// gcs-prod-websocket-eip-telecom
	Name *bool `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// default-aliyun
	Namespace *bool `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The pod scope.
	//
	// example:
	//
	// FDN
	ResourceScope *bool `json:"ResourceScope,omitempty" xml:"ResourceScope,omitempty"`
	// The status of the pod.
	//
	// example:
	//
	// RUNNING
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeServcieScheduleResponseBodyPodAbstractInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeServcieScheduleResponseBodyPodAbstractInfo) GoString() string {
	return s.String()
}

func (s *DescribeServcieScheduleResponseBodyPodAbstractInfo) GetContainerService() *bool {
	return s.ContainerService
}

func (s *DescribeServcieScheduleResponseBodyPodAbstractInfo) GetContainerStatuses() *DescribeServcieScheduleResponseBodyPodAbstractInfoContainerStatuses {
	return s.ContainerStatuses
}

func (s *DescribeServcieScheduleResponseBodyPodAbstractInfo) GetName() *bool {
	return s.Name
}

func (s *DescribeServcieScheduleResponseBodyPodAbstractInfo) GetNamespace() *bool {
	return s.Namespace
}

func (s *DescribeServcieScheduleResponseBodyPodAbstractInfo) GetResourceScope() *bool {
	return s.ResourceScope
}

func (s *DescribeServcieScheduleResponseBodyPodAbstractInfo) GetStatus() *bool {
	return s.Status
}

func (s *DescribeServcieScheduleResponseBodyPodAbstractInfo) SetContainerService(v bool) *DescribeServcieScheduleResponseBodyPodAbstractInfo {
	s.ContainerService = &v
	return s
}

func (s *DescribeServcieScheduleResponseBodyPodAbstractInfo) SetContainerStatuses(v *DescribeServcieScheduleResponseBodyPodAbstractInfoContainerStatuses) *DescribeServcieScheduleResponseBodyPodAbstractInfo {
	s.ContainerStatuses = v
	return s
}

func (s *DescribeServcieScheduleResponseBodyPodAbstractInfo) SetName(v bool) *DescribeServcieScheduleResponseBodyPodAbstractInfo {
	s.Name = &v
	return s
}

func (s *DescribeServcieScheduleResponseBodyPodAbstractInfo) SetNamespace(v bool) *DescribeServcieScheduleResponseBodyPodAbstractInfo {
	s.Namespace = &v
	return s
}

func (s *DescribeServcieScheduleResponseBodyPodAbstractInfo) SetResourceScope(v bool) *DescribeServcieScheduleResponseBodyPodAbstractInfo {
	s.ResourceScope = &v
	return s
}

func (s *DescribeServcieScheduleResponseBodyPodAbstractInfo) SetStatus(v bool) *DescribeServcieScheduleResponseBodyPodAbstractInfo {
	s.Status = &v
	return s
}

func (s *DescribeServcieScheduleResponseBodyPodAbstractInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeServcieScheduleResponseBodyPodAbstractInfoContainerStatuses struct {
	ContainerStatus []*DescribeServcieScheduleResponseBodyPodAbstractInfoContainerStatusesContainerStatus `json:"ContainerStatus,omitempty" xml:"ContainerStatus,omitempty" type:"Repeated"`
}

func (s DescribeServcieScheduleResponseBodyPodAbstractInfoContainerStatuses) String() string {
	return dara.Prettify(s)
}

func (s DescribeServcieScheduleResponseBodyPodAbstractInfoContainerStatuses) GoString() string {
	return s.String()
}

func (s *DescribeServcieScheduleResponseBodyPodAbstractInfoContainerStatuses) GetContainerStatus() []*DescribeServcieScheduleResponseBodyPodAbstractInfoContainerStatusesContainerStatus {
	return s.ContainerStatus
}

func (s *DescribeServcieScheduleResponseBodyPodAbstractInfoContainerStatuses) SetContainerStatus(v []*DescribeServcieScheduleResponseBodyPodAbstractInfoContainerStatusesContainerStatus) *DescribeServcieScheduleResponseBodyPodAbstractInfoContainerStatuses {
	s.ContainerStatus = v
	return s
}

func (s *DescribeServcieScheduleResponseBodyPodAbstractInfoContainerStatuses) Validate() error {
	return dara.Validate(s)
}

type DescribeServcieScheduleResponseBodyPodAbstractInfoContainerStatusesContainerStatus struct {
	// The ID of the container.
	//
	// example:
	//
	// container_e79_1638372147094_158091_02_000001
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	// The name of the container.
	//
	// example:
	//
	// nginx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeServcieScheduleResponseBodyPodAbstractInfoContainerStatusesContainerStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeServcieScheduleResponseBodyPodAbstractInfoContainerStatusesContainerStatus) GoString() string {
	return s.String()
}

func (s *DescribeServcieScheduleResponseBodyPodAbstractInfoContainerStatusesContainerStatus) GetContainerId() *string {
	return s.ContainerId
}

func (s *DescribeServcieScheduleResponseBodyPodAbstractInfoContainerStatusesContainerStatus) GetName() *string {
	return s.Name
}

func (s *DescribeServcieScheduleResponseBodyPodAbstractInfoContainerStatusesContainerStatus) SetContainerId(v string) *DescribeServcieScheduleResponseBodyPodAbstractInfoContainerStatusesContainerStatus {
	s.ContainerId = &v
	return s
}

func (s *DescribeServcieScheduleResponseBodyPodAbstractInfoContainerStatusesContainerStatus) SetName(v string) *DescribeServcieScheduleResponseBodyPodAbstractInfoContainerStatusesContainerStatus {
	s.Name = &v
	return s
}

func (s *DescribeServcieScheduleResponseBodyPodAbstractInfoContainerStatusesContainerStatus) Validate() error {
	return dara.Validate(s)
}
