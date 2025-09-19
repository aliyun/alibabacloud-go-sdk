// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPodRiskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *ListPodRiskResponseBodyPageInfo) *ListPodRiskResponseBody
	GetPageInfo() *ListPodRiskResponseBodyPageInfo
	SetPodRiskList(v []*ListPodRiskResponseBodyPodRiskList) *ListPodRiskResponseBody
	GetPodRiskList() []*ListPodRiskResponseBodyPodRiskList
	SetRequestId(v string) *ListPodRiskResponseBody
	GetRequestId() *string
}

type ListPodRiskResponseBody struct {
	// The pagination information.
	PageInfo *ListPodRiskResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// An array that consists of the risks.
	PodRiskList []*ListPodRiskResponseBodyPodRiskList `json:"PodRiskList,omitempty" xml:"PodRiskList,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 94254AD7-5026-5847-833B-403C2326BD6E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPodRiskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPodRiskResponseBody) GoString() string {
	return s.String()
}

func (s *ListPodRiskResponseBody) GetPageInfo() *ListPodRiskResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListPodRiskResponseBody) GetPodRiskList() []*ListPodRiskResponseBodyPodRiskList {
	return s.PodRiskList
}

func (s *ListPodRiskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPodRiskResponseBody) SetPageInfo(v *ListPodRiskResponseBodyPageInfo) *ListPodRiskResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListPodRiskResponseBody) SetPodRiskList(v []*ListPodRiskResponseBodyPodRiskList) *ListPodRiskResponseBody {
	s.PodRiskList = v
	return s
}

func (s *ListPodRiskResponseBody) SetRequestId(v string) *ListPodRiskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPodRiskResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListPodRiskResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
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
	// The number of entries returned per page.
	//
	// example:
	//
	// 2-
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 29
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPodRiskResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListPodRiskResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListPodRiskResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *ListPodRiskResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListPodRiskResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPodRiskResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPodRiskResponseBodyPageInfo) SetCount(v int32) *ListPodRiskResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListPodRiskResponseBodyPageInfo) SetCurrentPage(v int32) *ListPodRiskResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListPodRiskResponseBodyPageInfo) SetPageSize(v int32) *ListPodRiskResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListPodRiskResponseBodyPageInfo) SetTotalCount(v int32) *ListPodRiskResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListPodRiskResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type ListPodRiskResponseBodyPodRiskList struct {
	// The number of alerts that are generated for the pod.
	//
	// example:
	//
	// 10
	AlarmCount *int32 `json:"AlarmCount,omitempty" xml:"AlarmCount,omitempty"`
	// The ID of the container cluster.
	//
	// example:
	//
	// c1d903a628af043659a043af59d89****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// test66
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The timestamp that indicates the time when the pod was created. Unit: milliseconds.
	//
	// example:
	//
	// 1644283112720
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The number of baseline risk items that are detected in the pod.
	//
	// example:
	//
	// 1
	HcCount *int32 `json:"HcCount,omitempty" xml:"HcCount,omitempty"`
	// The instance ID of the node.
	//
	// example:
	//
	// i-7yvdq597****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The namespace of the Kubernetes cluster.
	//
	// example:
	//
	// kube-system
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// test
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The name of the pod.
	//
	// example:
	//
	// expoit-xxx-b****
	Pod *string `json:"Pod,omitempty" xml:"Pod,omitempty"`
	// The IP address of the pod.
	//
	// example:
	//
	// 172.0.XXX.XXX
	PodIp *string `json:"PodIp,omitempty" xml:"PodIp,omitempty"`
	// The number of vulnerabilities that are detected in the pod.
	//
	// example:
	//
	// 1
	VulCount *int32 `json:"VulCount,omitempty" xml:"VulCount,omitempty"`
}

func (s ListPodRiskResponseBodyPodRiskList) String() string {
	return dara.Prettify(s)
}

func (s ListPodRiskResponseBodyPodRiskList) GoString() string {
	return s.String()
}

func (s *ListPodRiskResponseBodyPodRiskList) GetAlarmCount() *int32 {
	return s.AlarmCount
}

func (s *ListPodRiskResponseBodyPodRiskList) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListPodRiskResponseBodyPodRiskList) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListPodRiskResponseBodyPodRiskList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListPodRiskResponseBodyPodRiskList) GetHcCount() *int32 {
	return s.HcCount
}

func (s *ListPodRiskResponseBodyPodRiskList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListPodRiskResponseBodyPodRiskList) GetNamespace() *string {
	return s.Namespace
}

func (s *ListPodRiskResponseBodyPodRiskList) GetNodeName() *string {
	return s.NodeName
}

func (s *ListPodRiskResponseBodyPodRiskList) GetPod() *string {
	return s.Pod
}

func (s *ListPodRiskResponseBodyPodRiskList) GetPodIp() *string {
	return s.PodIp
}

func (s *ListPodRiskResponseBodyPodRiskList) GetVulCount() *int32 {
	return s.VulCount
}

func (s *ListPodRiskResponseBodyPodRiskList) SetAlarmCount(v int32) *ListPodRiskResponseBodyPodRiskList {
	s.AlarmCount = &v
	return s
}

func (s *ListPodRiskResponseBodyPodRiskList) SetClusterId(v string) *ListPodRiskResponseBodyPodRiskList {
	s.ClusterId = &v
	return s
}

func (s *ListPodRiskResponseBodyPodRiskList) SetClusterName(v string) *ListPodRiskResponseBodyPodRiskList {
	s.ClusterName = &v
	return s
}

func (s *ListPodRiskResponseBodyPodRiskList) SetCreateTime(v int64) *ListPodRiskResponseBodyPodRiskList {
	s.CreateTime = &v
	return s
}

func (s *ListPodRiskResponseBodyPodRiskList) SetHcCount(v int32) *ListPodRiskResponseBodyPodRiskList {
	s.HcCount = &v
	return s
}

func (s *ListPodRiskResponseBodyPodRiskList) SetInstanceId(v string) *ListPodRiskResponseBodyPodRiskList {
	s.InstanceId = &v
	return s
}

func (s *ListPodRiskResponseBodyPodRiskList) SetNamespace(v string) *ListPodRiskResponseBodyPodRiskList {
	s.Namespace = &v
	return s
}

func (s *ListPodRiskResponseBodyPodRiskList) SetNodeName(v string) *ListPodRiskResponseBodyPodRiskList {
	s.NodeName = &v
	return s
}

func (s *ListPodRiskResponseBodyPodRiskList) SetPod(v string) *ListPodRiskResponseBodyPodRiskList {
	s.Pod = &v
	return s
}

func (s *ListPodRiskResponseBodyPodRiskList) SetPodIp(v string) *ListPodRiskResponseBodyPodRiskList {
	s.PodIp = &v
	return s
}

func (s *ListPodRiskResponseBodyPodRiskList) SetVulCount(v int32) *ListPodRiskResponseBodyPodRiskList {
	s.VulCount = &v
	return s
}

func (s *ListPodRiskResponseBodyPodRiskList) Validate() error {
	return dara.Validate(s)
}
