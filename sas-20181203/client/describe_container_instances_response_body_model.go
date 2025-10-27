// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContainerInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContainerInstanceList(v []*DescribeContainerInstancesResponseBodyContainerInstanceList) *DescribeContainerInstancesResponseBody
	GetContainerInstanceList() []*DescribeContainerInstancesResponseBodyContainerInstanceList
	SetPageInfo(v *DescribeContainerInstancesResponseBodyPageInfo) *DescribeContainerInstancesResponseBody
	GetPageInfo() *DescribeContainerInstancesResponseBodyPageInfo
	SetRequestId(v string) *DescribeContainerInstancesResponseBody
	GetRequestId() *string
}

type DescribeContainerInstancesResponseBody struct {
	// The details of the container asset.
	ContainerInstanceList []*DescribeContainerInstancesResponseBodyContainerInstanceList `json:"ContainerInstanceList,omitempty" xml:"ContainerInstanceList,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeContainerInstancesResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 79CFF74D-E967-5407-8A78-EE03B925FDAA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeContainerInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContainerInstancesResponseBody) GetContainerInstanceList() []*DescribeContainerInstancesResponseBodyContainerInstanceList {
	return s.ContainerInstanceList
}

func (s *DescribeContainerInstancesResponseBody) GetPageInfo() *DescribeContainerInstancesResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeContainerInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeContainerInstancesResponseBody) SetContainerInstanceList(v []*DescribeContainerInstancesResponseBodyContainerInstanceList) *DescribeContainerInstancesResponseBody {
	s.ContainerInstanceList = v
	return s
}

func (s *DescribeContainerInstancesResponseBody) SetPageInfo(v *DescribeContainerInstancesResponseBodyPageInfo) *DescribeContainerInstancesResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeContainerInstancesResponseBody) SetRequestId(v string) *DescribeContainerInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContainerInstancesResponseBody) Validate() error {
	if s.ContainerInstanceList != nil {
		for _, item := range s.ContainerInstanceList {
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

type DescribeContainerInstancesResponseBodyContainerInstanceList struct {
	// The number of alerts.
	//
	// example:
	//
	// 1
	AlarmCount *int32 `json:"AlarmCount,omitempty" xml:"AlarmCount,omitempty"`
	// Indicates whether alerts are generated for the container. Valid values:
	//
	// 	- **YES**
	//
	// 	- **NO**
	//
	// example:
	//
	// YES
	AlarmStatus *string `json:"AlarmStatus,omitempty" xml:"AlarmStatus,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// alibaba-log-controller
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// cfb41a869c71e4678a97021582dd8a****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// test
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The ID of the container.
	//
	// example:
	//
	// 48a6d9a92435a13ad573372c3f3c63b7e04d106458141df9f92155709d****
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	// The timestamp when the cluster was created. Unit: milliseconds.
	//
	// example:
	//
	// 1670368337000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// Indicates whether the asset is exposed to the Internet.
	//
	// 	- **1**: exposed
	//
	// 	- **0**: not exposed
	//
	// example:
	//
	// 1
	Exposed *int32 `json:"Exposed,omitempty" xml:"Exposed,omitempty"`
	// The exposure details. The value is a JSON string.
	//
	// example:
	//
	// [{}]
	ExposedDetail *string `json:"ExposedDetail,omitempty" xml:"ExposedDetail,omitempty"`
	// The number of baseline risks.
	//
	// example:
	//
	// 1
	HcCount *int32 `json:"HcCount,omitempty" xml:"HcCount,omitempty"`
	// Indicates whether baseline risks are detected. Valid values:
	//
	// 	- **NO**
	//
	// 	- **YES**
	//
	// example:
	//
	// YES
	HcStatus *string `json:"HcStatus,omitempty" xml:"HcStatus,omitempty"`
	// The IP address of the host.
	//
	// example:
	//
	// 172.24.XX.XX
	HostIp *string `json:"HostIp,omitempty" xml:"HostIp,omitempty"`
	// The image of the container.
	//
	// example:
	//
	// docker.io/library/nginx:latest
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The digest value of the image.
	//
	// example:
	//
	// 5b046e2de8c490819125193ee2eb71a66f2cc16c032dcd8b69ead4be1024****
	ImageDigest *string `json:"ImageDigest,omitempty" xml:"ImageDigest,omitempty"`
	// The image ID.
	//
	// example:
	//
	// registry-vpc.cn-beijing.aliyuncs.com/acs/log-controller@sha256:5b046e2de8c490819125193ee2eb71a66f2cc16c032dcd8b69ead4be1024****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the image repository.
	//
	// example:
	//
	// log-controller
	ImageRepoName *string `json:"ImageRepoName,omitempty" xml:"ImageRepoName,omitempty"`
	// The namespace of the image repository.
	//
	// example:
	//
	// acs
	ImageRepoNamespace *string `json:"ImageRepoNamespace,omitempty" xml:"ImageRepoNamespace,omitempty"`
	// The tag that is added to the image.
	//
	// example:
	//
	// 0.3.1.0-dfa2010-aliyun
	ImageRepoTag *string `json:"ImageRepoTag,omitempty" xml:"ImageRepoTag,omitempty"`
	// The UUID of the image.
	//
	// example:
	//
	// 5f23dfbceec289a49ac94e035e2****
	ImageUuid *string `json:"ImageUuid,omitempty" xml:"ImageUuid,omitempty"`
	// The instance ID of the asset.
	//
	// example:
	//
	// i-2zegzjyotydfkz9****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The namespace.
	//
	// example:
	//
	// kube-system
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The node information.
	//
	// example:
	//
	// test
	NodeInfo *string `json:"NodeInfo,omitempty" xml:"NodeInfo,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// test
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The pod.
	//
	// example:
	//
	// alibaba-log-controller-6f847f8786-mk2mg
	Pod *string `json:"Pod,omitempty" xml:"Pod,omitempty"`
	// The IP address of the pod.
	//
	// example:
	//
	// 172.24.XX.XX
	PodIp *string `json:"PodIp,omitempty" xml:"PodIp,omitempty"`
	// The region ID of the container.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of risks.
	//
	// example:
	//
	// 1
	RiskCount *string `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	// Indicates whether risks exist. Valid values:
	//
	// 	- **NO**
	//
	// 	- **YES**
	//
	// example:
	//
	// YES
	RiskStatus *string `json:"RiskStatus,omitempty" xml:"RiskStatus,omitempty"`
	// The update identifier of the container.
	//
	// example:
	//
	// 79cff74d-e967-5407-8a78-ee03b9****
	UpdateMark *string `json:"UpdateMark,omitempty" xml:"UpdateMark,omitempty"`
	// The number of vulnerabilities that are detected in the container cluster.
	//
	// example:
	//
	// 15
	VulCount *int32 `json:"VulCount,omitempty" xml:"VulCount,omitempty"`
	// Indicates whether vulnerabilities are detected in the container. Valid values:
	//
	// 	- **YES**
	//
	// 	- **NO**
	//
	// example:
	//
	// YES
	VulStatus *string `json:"VulStatus,omitempty" xml:"VulStatus,omitempty"`
}

func (s DescribeContainerInstancesResponseBodyContainerInstanceList) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerInstancesResponseBodyContainerInstanceList) GoString() string {
	return s.String()
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) GetAlarmCount() *int32 {
	return s.AlarmCount
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) GetAlarmStatus() *string {
	return s.AlarmStatus
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) GetAppName() *string {
	return s.AppName
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) GetContainerId() *string {
	return s.ContainerId
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) GetExposed() *int32 {
	return s.Exposed
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) GetExposedDetail() *string {
	return s.ExposedDetail
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) GetHcCount() *int32 {
	return s.HcCount
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) GetHcStatus() *string {
	return s.HcStatus
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) GetHostIp() *string {
	return s.HostIp
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) GetImage() *string {
	return s.Image
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) GetImageDigest() *string {
	return s.ImageDigest
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) GetImageRepoName() *string {
	return s.ImageRepoName
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) GetImageRepoNamespace() *string {
	return s.ImageRepoNamespace
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) GetImageRepoTag() *string {
	return s.ImageRepoTag
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) GetImageUuid() *string {
	return s.ImageUuid
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) GetNodeInfo() *string {
	return s.NodeInfo
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) GetNodeName() *string {
	return s.NodeName
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) GetPod() *string {
	return s.Pod
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) GetPodIp() *string {
	return s.PodIp
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) GetRiskCount() *string {
	return s.RiskCount
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) GetRiskStatus() *string {
	return s.RiskStatus
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) GetUpdateMark() *string {
	return s.UpdateMark
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) GetVulCount() *int32 {
	return s.VulCount
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) GetVulStatus() *string {
	return s.VulStatus
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) SetAlarmCount(v int32) *DescribeContainerInstancesResponseBodyContainerInstanceList {
	s.AlarmCount = &v
	return s
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) SetAlarmStatus(v string) *DescribeContainerInstancesResponseBodyContainerInstanceList {
	s.AlarmStatus = &v
	return s
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) SetAppName(v string) *DescribeContainerInstancesResponseBodyContainerInstanceList {
	s.AppName = &v
	return s
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) SetClusterId(v string) *DescribeContainerInstancesResponseBodyContainerInstanceList {
	s.ClusterId = &v
	return s
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) SetClusterName(v string) *DescribeContainerInstancesResponseBodyContainerInstanceList {
	s.ClusterName = &v
	return s
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) SetContainerId(v string) *DescribeContainerInstancesResponseBodyContainerInstanceList {
	s.ContainerId = &v
	return s
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) SetCreateTimestamp(v int64) *DescribeContainerInstancesResponseBodyContainerInstanceList {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) SetExposed(v int32) *DescribeContainerInstancesResponseBodyContainerInstanceList {
	s.Exposed = &v
	return s
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) SetExposedDetail(v string) *DescribeContainerInstancesResponseBodyContainerInstanceList {
	s.ExposedDetail = &v
	return s
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) SetHcCount(v int32) *DescribeContainerInstancesResponseBodyContainerInstanceList {
	s.HcCount = &v
	return s
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) SetHcStatus(v string) *DescribeContainerInstancesResponseBodyContainerInstanceList {
	s.HcStatus = &v
	return s
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) SetHostIp(v string) *DescribeContainerInstancesResponseBodyContainerInstanceList {
	s.HostIp = &v
	return s
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) SetImage(v string) *DescribeContainerInstancesResponseBodyContainerInstanceList {
	s.Image = &v
	return s
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) SetImageDigest(v string) *DescribeContainerInstancesResponseBodyContainerInstanceList {
	s.ImageDigest = &v
	return s
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) SetImageId(v string) *DescribeContainerInstancesResponseBodyContainerInstanceList {
	s.ImageId = &v
	return s
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) SetImageRepoName(v string) *DescribeContainerInstancesResponseBodyContainerInstanceList {
	s.ImageRepoName = &v
	return s
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) SetImageRepoNamespace(v string) *DescribeContainerInstancesResponseBodyContainerInstanceList {
	s.ImageRepoNamespace = &v
	return s
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) SetImageRepoTag(v string) *DescribeContainerInstancesResponseBodyContainerInstanceList {
	s.ImageRepoTag = &v
	return s
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) SetImageUuid(v string) *DescribeContainerInstancesResponseBodyContainerInstanceList {
	s.ImageUuid = &v
	return s
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) SetInstanceId(v string) *DescribeContainerInstancesResponseBodyContainerInstanceList {
	s.InstanceId = &v
	return s
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) SetNamespace(v string) *DescribeContainerInstancesResponseBodyContainerInstanceList {
	s.Namespace = &v
	return s
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) SetNodeInfo(v string) *DescribeContainerInstancesResponseBodyContainerInstanceList {
	s.NodeInfo = &v
	return s
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) SetNodeName(v string) *DescribeContainerInstancesResponseBodyContainerInstanceList {
	s.NodeName = &v
	return s
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) SetPod(v string) *DescribeContainerInstancesResponseBodyContainerInstanceList {
	s.Pod = &v
	return s
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) SetPodIp(v string) *DescribeContainerInstancesResponseBodyContainerInstanceList {
	s.PodIp = &v
	return s
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) SetRegionId(v string) *DescribeContainerInstancesResponseBodyContainerInstanceList {
	s.RegionId = &v
	return s
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) SetRiskCount(v string) *DescribeContainerInstancesResponseBodyContainerInstanceList {
	s.RiskCount = &v
	return s
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) SetRiskStatus(v string) *DescribeContainerInstancesResponseBodyContainerInstanceList {
	s.RiskStatus = &v
	return s
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) SetUpdateMark(v string) *DescribeContainerInstancesResponseBodyContainerInstanceList {
	s.UpdateMark = &v
	return s
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) SetVulCount(v int32) *DescribeContainerInstancesResponseBodyContainerInstanceList {
	s.VulCount = &v
	return s
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) SetVulStatus(v string) *DescribeContainerInstancesResponseBodyContainerInstanceList {
	s.VulStatus = &v
	return s
}

func (s *DescribeContainerInstancesResponseBodyContainerInstanceList) Validate() error {
	return dara.Validate(s)
}

type DescribeContainerInstancesResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 3
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 45
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeContainerInstancesResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerInstancesResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeContainerInstancesResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeContainerInstancesResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeContainerInstancesResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeContainerInstancesResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeContainerInstancesResponseBodyPageInfo) SetCount(v int32) *DescribeContainerInstancesResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeContainerInstancesResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeContainerInstancesResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeContainerInstancesResponseBodyPageInfo) SetPageSize(v int32) *DescribeContainerInstancesResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeContainerInstancesResponseBodyPageInfo) SetTotalCount(v int32) *DescribeContainerInstancesResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeContainerInstancesResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
