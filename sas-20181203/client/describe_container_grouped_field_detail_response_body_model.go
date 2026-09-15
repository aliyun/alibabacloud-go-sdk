// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContainerGroupedFieldDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeContainerGroupedFieldDetailResponseBodyData) *DescribeContainerGroupedFieldDetailResponseBody
	GetData() *DescribeContainerGroupedFieldDetailResponseBodyData
	SetRequestId(v string) *DescribeContainerGroupedFieldDetailResponseBody
	GetRequestId() *string
}

type DescribeContainerGroupedFieldDetailResponseBody struct {
	// The returned data.
	Data *DescribeContainerGroupedFieldDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7532B7EE-7CE7-5F4D-BF04-B12447DD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeContainerGroupedFieldDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerGroupedFieldDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupedFieldDetailResponseBody) GetData() *DescribeContainerGroupedFieldDetailResponseBodyData {
	return s.Data
}

func (s *DescribeContainerGroupedFieldDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeContainerGroupedFieldDetailResponseBody) SetData(v *DescribeContainerGroupedFieldDetailResponseBodyData) *DescribeContainerGroupedFieldDetailResponseBody {
	s.Data = v
	return s
}

func (s *DescribeContainerGroupedFieldDetailResponseBody) SetRequestId(v string) *DescribeContainerGroupedFieldDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContainerGroupedFieldDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeContainerGroupedFieldDetailResponseBodyData struct {
	// The number of alerts.
	//
	// example:
	//
	// 1
	AlarmCount *int32 `json:"AlarmCount,omitempty" xml:"AlarmCount,omitempty"`
	// The application name.
	//
	// example:
	//
	// node-local-dns
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The version of the current online server in the cluster.
	//
	// example:
	//
	// 1.14.8-aliyun.1
	ClusterCurrentVersion *string `json:"ClusterCurrentVersion,omitempty" xml:"ClusterCurrentVersion,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// c8ca91e0907d94efaba7fb0827eb9****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// zhhtest
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The cluster status. Valid values:
	//
	// - STARTING: Starting.
	//
	// - START_FAILED: Failed to start.
	//
	// - BOOTSTRAPPING: Initializing bootstrap actions.
	//
	// - RUNNING: Running.
	//
	// - TERMINATING: Terminating.
	//
	// - TERMINATED: Terminated.
	//
	// - TERMINATED_WITH_ERRORS: Terminated due to an exception.
	//
	// - TERMINATE_FAILED: Failed to terminate.
	//
	// example:
	//
	// RUNNING
	ClusterState *string `json:"ClusterState,omitempty" xml:"ClusterState,omitempty"`
	// The cluster type. Valid values:
	//
	// - **Kubernetes**: dedicated Kubernetes cluster.
	//
	// - **ManagedKubernetes**: standard managed cluster (edge cluster).
	//
	// - **Ask**: standard serverless cluster.
	//
	// example:
	//
	// ManagedKubernetes
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The number of containers.
	//
	// example:
	//
	// 1
	ContainerCount *int32 `json:"ContainerCount,omitempty" xml:"ContainerCount,omitempty"`
	// The creation time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1702433618301
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The number of instances.
	//
	// example:
	//
	// 1
	InstanceCount *int32 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// The namespace.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The node name.
	//
	// example:
	//
	// shangliang-test
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The POD name.
	//
	// example:
	//
	// node-local-dns-zwsxl
	Pod *string `json:"Pod,omitempty" xml:"Pod,omitempty"`
	// The number of PODs.
	//
	// example:
	//
	// 1
	PodCount *int32 `json:"PodCount,omitempty" xml:"PodCount,omitempty"`
	// The IP address of the POD.
	//
	// example:
	//
	// 192.168.1.1
	PodIp *string `json:"PodIp,omitempty" xml:"PodIp,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of vulnerabilities.
	//
	// example:
	//
	// 1
	VulCount *int32 `json:"VulCount,omitempty" xml:"VulCount,omitempty"`
}

func (s DescribeContainerGroupedFieldDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerGroupedFieldDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupedFieldDetailResponseBodyData) GetAlarmCount() *int32 {
	return s.AlarmCount
}

func (s *DescribeContainerGroupedFieldDetailResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *DescribeContainerGroupedFieldDetailResponseBodyData) GetClusterCurrentVersion() *string {
	return s.ClusterCurrentVersion
}

func (s *DescribeContainerGroupedFieldDetailResponseBodyData) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeContainerGroupedFieldDetailResponseBodyData) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeContainerGroupedFieldDetailResponseBodyData) GetClusterState() *string {
	return s.ClusterState
}

func (s *DescribeContainerGroupedFieldDetailResponseBodyData) GetClusterType() *string {
	return s.ClusterType
}

func (s *DescribeContainerGroupedFieldDetailResponseBodyData) GetContainerCount() *int32 {
	return s.ContainerCount
}

func (s *DescribeContainerGroupedFieldDetailResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeContainerGroupedFieldDetailResponseBodyData) GetInstanceCount() *int32 {
	return s.InstanceCount
}

func (s *DescribeContainerGroupedFieldDetailResponseBodyData) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeContainerGroupedFieldDetailResponseBodyData) GetNodeName() *string {
	return s.NodeName
}

func (s *DescribeContainerGroupedFieldDetailResponseBodyData) GetPod() *string {
	return s.Pod
}

func (s *DescribeContainerGroupedFieldDetailResponseBodyData) GetPodCount() *int32 {
	return s.PodCount
}

func (s *DescribeContainerGroupedFieldDetailResponseBodyData) GetPodIp() *string {
	return s.PodIp
}

func (s *DescribeContainerGroupedFieldDetailResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeContainerGroupedFieldDetailResponseBodyData) GetVulCount() *int32 {
	return s.VulCount
}

func (s *DescribeContainerGroupedFieldDetailResponseBodyData) SetAlarmCount(v int32) *DescribeContainerGroupedFieldDetailResponseBodyData {
	s.AlarmCount = &v
	return s
}

func (s *DescribeContainerGroupedFieldDetailResponseBodyData) SetAppName(v string) *DescribeContainerGroupedFieldDetailResponseBodyData {
	s.AppName = &v
	return s
}

func (s *DescribeContainerGroupedFieldDetailResponseBodyData) SetClusterCurrentVersion(v string) *DescribeContainerGroupedFieldDetailResponseBodyData {
	s.ClusterCurrentVersion = &v
	return s
}

func (s *DescribeContainerGroupedFieldDetailResponseBodyData) SetClusterId(v string) *DescribeContainerGroupedFieldDetailResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *DescribeContainerGroupedFieldDetailResponseBodyData) SetClusterName(v string) *DescribeContainerGroupedFieldDetailResponseBodyData {
	s.ClusterName = &v
	return s
}

func (s *DescribeContainerGroupedFieldDetailResponseBodyData) SetClusterState(v string) *DescribeContainerGroupedFieldDetailResponseBodyData {
	s.ClusterState = &v
	return s
}

func (s *DescribeContainerGroupedFieldDetailResponseBodyData) SetClusterType(v string) *DescribeContainerGroupedFieldDetailResponseBodyData {
	s.ClusterType = &v
	return s
}

func (s *DescribeContainerGroupedFieldDetailResponseBodyData) SetContainerCount(v int32) *DescribeContainerGroupedFieldDetailResponseBodyData {
	s.ContainerCount = &v
	return s
}

func (s *DescribeContainerGroupedFieldDetailResponseBodyData) SetCreateTime(v int64) *DescribeContainerGroupedFieldDetailResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DescribeContainerGroupedFieldDetailResponseBodyData) SetInstanceCount(v int32) *DescribeContainerGroupedFieldDetailResponseBodyData {
	s.InstanceCount = &v
	return s
}

func (s *DescribeContainerGroupedFieldDetailResponseBodyData) SetNamespace(v string) *DescribeContainerGroupedFieldDetailResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *DescribeContainerGroupedFieldDetailResponseBodyData) SetNodeName(v string) *DescribeContainerGroupedFieldDetailResponseBodyData {
	s.NodeName = &v
	return s
}

func (s *DescribeContainerGroupedFieldDetailResponseBodyData) SetPod(v string) *DescribeContainerGroupedFieldDetailResponseBodyData {
	s.Pod = &v
	return s
}

func (s *DescribeContainerGroupedFieldDetailResponseBodyData) SetPodCount(v int32) *DescribeContainerGroupedFieldDetailResponseBodyData {
	s.PodCount = &v
	return s
}

func (s *DescribeContainerGroupedFieldDetailResponseBodyData) SetPodIp(v string) *DescribeContainerGroupedFieldDetailResponseBodyData {
	s.PodIp = &v
	return s
}

func (s *DescribeContainerGroupedFieldDetailResponseBodyData) SetRegionId(v string) *DescribeContainerGroupedFieldDetailResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *DescribeContainerGroupedFieldDetailResponseBodyData) SetVulCount(v int32) *DescribeContainerGroupedFieldDetailResponseBodyData {
	s.VulCount = &v
	return s
}

func (s *DescribeContainerGroupedFieldDetailResponseBodyData) Validate() error {
	return dara.Validate(s)
}
