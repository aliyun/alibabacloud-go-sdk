// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupedContainerInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupedContainerInstanceList(v []*DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) *DescribeGroupedContainerInstancesResponseBody
	GetGroupedContainerInstanceList() []*DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList
	SetPageInfo(v *DescribeGroupedContainerInstancesResponseBodyPageInfo) *DescribeGroupedContainerInstancesResponseBody
	GetPageInfo() *DescribeGroupedContainerInstancesResponseBodyPageInfo
	SetRequestId(v string) *DescribeGroupedContainerInstancesResponseBody
	GetRequestId() *string
}

type DescribeGroupedContainerInstancesResponseBody struct {
	// The information about the container.
	GroupedContainerInstanceList []*DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList `json:"GroupedContainerInstanceList,omitempty" xml:"GroupedContainerInstanceList,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeGroupedContainerInstancesResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 4E5BFDCF-B9DD-430D-9DA4-151BCB581C9D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeGroupedContainerInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupedContainerInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGroupedContainerInstancesResponseBody) GetGroupedContainerInstanceList() []*DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	return s.GroupedContainerInstanceList
}

func (s *DescribeGroupedContainerInstancesResponseBody) GetPageInfo() *DescribeGroupedContainerInstancesResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeGroupedContainerInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGroupedContainerInstancesResponseBody) SetGroupedContainerInstanceList(v []*DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) *DescribeGroupedContainerInstancesResponseBody {
	s.GroupedContainerInstanceList = v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBody) SetPageInfo(v *DescribeGroupedContainerInstancesResponseBodyPageInfo) *DescribeGroupedContainerInstancesResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBody) SetRequestId(v string) *DescribeGroupedContainerInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBody) Validate() error {
	if s.GroupedContainerInstanceList != nil {
		for _, item := range s.GroupedContainerInstanceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList struct {
	// The number of alerts that are detected for the current pod, application, namespace, or cluster.
	//
	// example:
	//
	// 1
	AlarmCount *int32 `json:"AlarmCount,omitempty" xml:"AlarmCount,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// oss-liveness-probe
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// cf3824769c85441b4bf3****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// test
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The type of the cluster. Valid values:
	//
	// 	- **Kubernetes**: dedicated Kubernetes cluster.
	//
	// 	- **ManagedKubernetes**: standard managed cluster (edge cluster).
	//
	// 	- **Ask**: serverless Kubernetes (ASK) cluster.
	//
	// example:
	//
	// ManagedKubernetes
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The timestamp when the cluster was created. Unit: milliseconds.
	//
	// example:
	//
	// 1600076893000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The status of the cluster. Valid values:
	//
	// 	- **running**: The cluster is running.
	//
	// 	- **stopped**: The cluster is stopped.
	//
	// 	- **deleted**: The cluster is deleted.
	//
	// 	- **delete_failed**: The cluster failed to be deleted.
	//
	// 	- **failed**: The cluster failed to be created.
	//
	// example:
	//
	// running
	CusterState *string `json:"CusterState,omitempty" xml:"CusterState,omitempty"`
	// The number of baseline risks that are detected for the current pod, application, namespace, or cluster.
	//
	// example:
	//
	// 20
	HcCount *int32 `json:"HcCount,omitempty" xml:"HcCount,omitempty"`
	// The IP address of the host in the container cluster.
	//
	// example:
	//
	// 172.114.XX.XX
	HostIp *string `json:"HostIp,omitempty" xml:"HostIp,omitempty"`
	// The container image.
	//
	// example:
	//
	// registry-vpc.ap-southeast-5.aliyuncs.com/log-service-release/sls-connector:1.1.77
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The digest value of the image.
	//
	// example:
	//
	// 58e58c013f70bbfde140c8a55c1078074b3483479428d4069aa946827fd566cf
	ImageDigest *string `json:"ImageDigest,omitempty" xml:"ImageDigest,omitempty"`
	// The name of the image repository.
	//
	// example:
	//
	// test-003
	ImageRepoName *string `json:"ImageRepoName,omitempty" xml:"ImageRepoName,omitempty"`
	// The namespace of the image repository.
	//
	// example:
	//
	// name001
	ImageRepoNamespace *string `json:"ImageRepoNamespace,omitempty" xml:"ImageRepoNamespace,omitempty"`
	// The tag that is added to the image repository.
	//
	// example:
	//
	// dev-20220512-2
	ImageRepoTag *string `json:"ImageRepoTag,omitempty" xml:"ImageRepoTag,omitempty"`
	// The UUID of the image.
	//
	// example:
	//
	// e4bdec1d9ba7e0967111a7ea467c****
	ImageUuid *string `json:"ImageUuid,omitempty" xml:"ImageUuid,omitempty"`
	// The number of pods, applications, clusters, or namespaces.
	//
	// example:
	//
	// 9
	InstanceCount *int32 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// The ID of the server.
	//
	// example:
	//
	// i-8vb9ul5xec4tua4q****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The namespace of the cluster.
	//
	// example:
	//
	// kube-system
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The name of the pod.
	//
	// example:
	//
	// csi-plugin-2n****
	Pod *string `json:"Pod,omitempty" xml:"Pod,omitempty"`
	// The IP address of the pod.
	//
	// example:
	//
	// 172.114.XX.XX
	PodIp *string `json:"PodIp,omitempty" xml:"PodIp,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of at-risk instances.
	//
	// example:
	//
	// 1
	RiskInstanceCount *int32 `json:"RiskInstanceCount,omitempty" xml:"RiskInstanceCount,omitempty"`
	// The risk level. Valid values:
	//
	// 	- **high**
	//
	// 	- **medium**
	//
	// 	- **low**
	//
	// example:
	//
	// low
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// Indicates whether risks were detected. Valid values:
	//
	// 	- **NO**
	//
	// 	- **YES**
	//
	// example:
	//
	// NO
	RiskStatus *string `json:"RiskStatus,omitempty" xml:"RiskStatus,omitempty"`
	// Indicates whether the synchronization of cluster audit logs is enabled. Valid values:
	//
	// 	- **0**: disabled.
	//
	// 	- **1**: enabled.
	//
	// example:
	//
	// 1
	SyncOpen *int32 `json:"SyncOpen,omitempty" xml:"SyncOpen,omitempty"`
	// The status of the synchronization of cluster audit logs. Valid values:
	//
	// 	- **0**: The synchronization failed.
	//
	// 	- **1**: The synchronization is successful.
	//
	// example:
	//
	// 1
	SyncStatus *int32 `json:"SyncStatus,omitempty" xml:"SyncStatus,omitempty"`
	// The number of vulnerabilities that are detected for the current pod, application, namespace, or cluster.
	//
	// example:
	//
	// 1
	VulCount *int32 `json:"VulCount,omitempty" xml:"VulCount,omitempty"`
}

func (s DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) GoString() string {
	return s.String()
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) GetAlarmCount() *int32 {
	return s.AlarmCount
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) GetAppName() *string {
	return s.AppName
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) GetClusterType() *string {
	return s.ClusterType
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) GetCusterState() *string {
	return s.CusterState
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) GetHcCount() *int32 {
	return s.HcCount
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) GetHostIp() *string {
	return s.HostIp
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) GetImage() *string {
	return s.Image
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) GetImageDigest() *string {
	return s.ImageDigest
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) GetImageRepoName() *string {
	return s.ImageRepoName
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) GetImageRepoNamespace() *string {
	return s.ImageRepoNamespace
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) GetImageRepoTag() *string {
	return s.ImageRepoTag
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) GetImageUuid() *string {
	return s.ImageUuid
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) GetInstanceCount() *int32 {
	return s.InstanceCount
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) GetPod() *string {
	return s.Pod
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) GetPodIp() *string {
	return s.PodIp
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) GetRiskInstanceCount() *int32 {
	return s.RiskInstanceCount
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) GetRiskStatus() *string {
	return s.RiskStatus
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) GetSyncOpen() *int32 {
	return s.SyncOpen
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) GetSyncStatus() *int32 {
	return s.SyncStatus
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) GetVulCount() *int32 {
	return s.VulCount
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetAlarmCount(v int32) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.AlarmCount = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetAppName(v string) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.AppName = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetClusterId(v string) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.ClusterId = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetClusterName(v string) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.ClusterName = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetClusterType(v string) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.ClusterType = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetCreateTime(v int64) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.CreateTime = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetCusterState(v string) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.CusterState = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetHcCount(v int32) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.HcCount = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetHostIp(v string) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.HostIp = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetImage(v string) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.Image = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetImageDigest(v string) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.ImageDigest = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetImageRepoName(v string) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.ImageRepoName = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetImageRepoNamespace(v string) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.ImageRepoNamespace = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetImageRepoTag(v string) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.ImageRepoTag = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetImageUuid(v string) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.ImageUuid = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetInstanceCount(v int32) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.InstanceCount = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetInstanceId(v string) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.InstanceId = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetNamespace(v string) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.Namespace = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetPod(v string) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.Pod = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetPodIp(v string) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.PodIp = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetRegionId(v string) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.RegionId = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetRiskInstanceCount(v int32) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.RiskInstanceCount = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetRiskLevel(v string) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.RiskLevel = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetRiskStatus(v string) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.RiskStatus = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetSyncOpen(v int32) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.SyncOpen = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetSyncStatus(v int32) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.SyncStatus = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetVulCount(v int32) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.VulCount = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) Validate() error {
	return dara.Validate(s)
}

type DescribeGroupedContainerInstancesResponseBodyPageInfo struct {
	// The number of container assets returned on the current page.
	//
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of container assets returned.
	//
	// example:
	//
	// 25
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeGroupedContainerInstancesResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupedContainerInstancesResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeGroupedContainerInstancesResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeGroupedContainerInstancesResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeGroupedContainerInstancesResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGroupedContainerInstancesResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeGroupedContainerInstancesResponseBodyPageInfo) SetCount(v int32) *DescribeGroupedContainerInstancesResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeGroupedContainerInstancesResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyPageInfo) SetPageSize(v int32) *DescribeGroupedContainerInstancesResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyPageInfo) SetTotalCount(v int32) *DescribeGroupedContainerInstancesResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
