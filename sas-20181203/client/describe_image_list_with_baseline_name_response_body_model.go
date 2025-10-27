// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageListWithBaselineNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImageInfos(v []*DescribeImageListWithBaselineNameResponseBodyImageInfos) *DescribeImageListWithBaselineNameResponseBody
	GetImageInfos() []*DescribeImageListWithBaselineNameResponseBodyImageInfos
	SetPageInfo(v *DescribeImageListWithBaselineNameResponseBodyPageInfo) *DescribeImageListWithBaselineNameResponseBody
	GetPageInfo() *DescribeImageListWithBaselineNameResponseBodyPageInfo
	SetRequestId(v string) *DescribeImageListWithBaselineNameResponseBody
	GetRequestId() *string
}

type DescribeImageListWithBaselineNameResponseBody struct {
	// The information about the images.
	ImageInfos []*DescribeImageListWithBaselineNameResponseBodyImageInfos `json:"ImageInfos,omitempty" xml:"ImageInfos,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeImageListWithBaselineNameResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 5B8C2156-2DB9-5A42-99E7-F2ED5AE9EA1F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeImageListWithBaselineNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageListWithBaselineNameResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageListWithBaselineNameResponseBody) GetImageInfos() []*DescribeImageListWithBaselineNameResponseBodyImageInfos {
	return s.ImageInfos
}

func (s *DescribeImageListWithBaselineNameResponseBody) GetPageInfo() *DescribeImageListWithBaselineNameResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeImageListWithBaselineNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageListWithBaselineNameResponseBody) SetImageInfos(v []*DescribeImageListWithBaselineNameResponseBodyImageInfos) *DescribeImageListWithBaselineNameResponseBody {
	s.ImageInfos = v
	return s
}

func (s *DescribeImageListWithBaselineNameResponseBody) SetPageInfo(v *DescribeImageListWithBaselineNameResponseBodyPageInfo) *DescribeImageListWithBaselineNameResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeImageListWithBaselineNameResponseBody) SetRequestId(v string) *DescribeImageListWithBaselineNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageListWithBaselineNameResponseBody) Validate() error {
	if s.ImageInfos != nil {
		for _, item := range s.ImageInfos {
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

type DescribeImageListWithBaselineNameResponseBodyImageInfos struct {
	// The ID of the cluster.
	//
	// example:
	//
	// c08d5fc1a329a4b88950a253d082f1****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// docker-law
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The ID of the container.
	//
	// example:
	//
	// 04d20e98c8e2c93b7b864372084320a15a58c8671e53c972ce3a71d9c163****
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	// The SHA-256 value of the image digest.
	//
	// example:
	//
	// 2e6daffce524ffeae66cccaa90c8fc47de912346dcec295c27395b6d66db6423
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The timestamp generated when the first scan was performed. Unit: milliseconds.
	//
	// example:
	//
	// 1649814050000
	FirstScanTime *int64 `json:"FirstScanTime,omitempty" xml:"FirstScanTime,omitempty"`
	// The number of images on which **high*	- baseline risks are detected.
	//
	// example:
	//
	// 1
	HighRiskImage *int32 `json:"HighRiskImage,omitempty" xml:"HighRiskImage,omitempty"`
	// The name of the image.
	//
	// example:
	//
	// registry.cn-wulanchabu.aliyuncs.com/sas_test/huxin-test-001:nuxeo6-conta****
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The timestamp when the image was created. Unit: milliseconds.
	//
	// example:
	//
	// 1636962328000
	ImageCreate *int64 `json:"ImageCreate,omitempty" xml:"ImageCreate,omitempty"`
	// The ID of the image.
	//
	// example:
	//
	// cddb5fd33b34a1fabb358d0a19497cdfe362fe624821cb250947af0ea5cc****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The size of the image.
	//
	// example:
	//
	// 157408623
	ImageSize *int32 `json:"ImageSize,omitempty" xml:"ImageSize,omitempty"`
	// The timestamp when the image was updated. Unit: milliseconds.
	//
	// example:
	//
	// 1636974116000
	ImageUpdate *int64 `json:"ImageUpdate,omitempty" xml:"ImageUpdate,omitempty"`
	// The ID of the image instance.
	//
	// example:
	//
	// cri-a595qp31knh9****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name of the server.
	//
	// example:
	//
	// pre.mongo-196
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the server.
	//
	// example:
	//
	// 47.96.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the server.
	//
	// example:
	//
	// 172.16.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The timestamp when the last baseline check was performed. Unit: milliseconds.
	//
	// example:
	//
	// 1649814050000
	LastScanTime *int64 `json:"LastScanTime,omitempty" xml:"LastScanTime,omitempty"`
	// The number of images on which **low*	- baseline risks are detected.
	//
	// example:
	//
	// 0
	LowRiskImage *int32 `json:"LowRiskImage,omitempty" xml:"LowRiskImage,omitempty"`
	// The number of images on which **medium*	- baseline risks are detected.
	//
	// example:
	//
	// 0
	MiddleRiskImage *int32 `json:"MiddleRiskImage,omitempty" xml:"MiddleRiskImage,omitempty"`
	// The namespace.
	//
	// example:
	//
	// test-002
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The number of images that do not have baseline risks.
	//
	// example:
	//
	// 0
	NoRiskImage *int32 `json:"NoRiskImage,omitempty" xml:"NoRiskImage,omitempty"`
	// The pod.
	//
	// example:
	//
	// 22222-7xsqq
	Pod *string `json:"Pod,omitempty" xml:"Pod,omitempty"`
	// The region ID of the image instance.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the image repository.
	//
	// example:
	//
	// crr-1lt6q7167yh6****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The name of the image repository.
	//
	// example:
	//
	// scanner
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The namespace to which the image repository belongs.
	//
	// example:
	//
	// libssh2
	RepoNamespace *string `json:"RepoNamespace,omitempty" xml:"RepoNamespace,omitempty"`
	// The type of the image repository.
	//
	// example:
	//
	// PRIVATE
	RepoType *string `json:"RepoType,omitempty" xml:"RepoType,omitempty"`
	// Indicates whether the image is at risk. Valid values:
	//
	// 	- **YES**
	//
	// 	- **NO**
	//
	// example:
	//
	// YES
	RiskStatus *string `json:"RiskStatus,omitempty" xml:"RiskStatus,omitempty"`
	// The version of the image.
	//
	// example:
	//
	// v1
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The ID of the asset on which the baseline check is performed.
	//
	// example:
	//
	// m-bp17m0pc0xprzbwo****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the asset on which the baseline check is performed.
	//
	// example:
	//
	// spod
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The type of the asset on which the baseline check is performed. Valid values:
	//
	// 	- ECS_IMAGE
	//
	// 	- ECS_SNAPSHOT
	//
	// example:
	//
	// ECS_IMAGE
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The total number of risk items that are detected on the image by using the baseline.
	//
	// example:
	//
	// 3
	TotalItemCount *int32 `json:"TotalItemCount,omitempty" xml:"TotalItemCount,omitempty"`
	// The UUID of the image.
	//
	// example:
	//
	// f58681174f944623345379e23b7b****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeImageListWithBaselineNameResponseBodyImageInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageListWithBaselineNameResponseBodyImageInfos) GoString() string {
	return s.String()
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) GetContainerId() *string {
	return s.ContainerId
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) GetDigest() *string {
	return s.Digest
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) GetFirstScanTime() *int64 {
	return s.FirstScanTime
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) GetHighRiskImage() *int32 {
	return s.HighRiskImage
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) GetImage() *string {
	return s.Image
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) GetImageCreate() *int64 {
	return s.ImageCreate
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) GetImageSize() *int32 {
	return s.ImageSize
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) GetImageUpdate() *int64 {
	return s.ImageUpdate
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) GetLastScanTime() *int64 {
	return s.LastScanTime
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) GetLowRiskImage() *int32 {
	return s.LowRiskImage
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) GetMiddleRiskImage() *int32 {
	return s.MiddleRiskImage
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) GetNoRiskImage() *int32 {
	return s.NoRiskImage
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) GetPod() *string {
	return s.Pod
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) GetRepoId() *string {
	return s.RepoId
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) GetRepoName() *string {
	return s.RepoName
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) GetRepoNamespace() *string {
	return s.RepoNamespace
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) GetRepoType() *string {
	return s.RepoType
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) GetRiskStatus() *string {
	return s.RiskStatus
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) GetTag() *string {
	return s.Tag
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) GetTargetId() *string {
	return s.TargetId
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) GetTargetName() *string {
	return s.TargetName
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) GetTargetType() *string {
	return s.TargetType
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) GetTotalItemCount() *int32 {
	return s.TotalItemCount
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) SetClusterId(v string) *DescribeImageListWithBaselineNameResponseBodyImageInfos {
	s.ClusterId = &v
	return s
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) SetClusterName(v string) *DescribeImageListWithBaselineNameResponseBodyImageInfos {
	s.ClusterName = &v
	return s
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) SetContainerId(v string) *DescribeImageListWithBaselineNameResponseBodyImageInfos {
	s.ContainerId = &v
	return s
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) SetDigest(v string) *DescribeImageListWithBaselineNameResponseBodyImageInfos {
	s.Digest = &v
	return s
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) SetFirstScanTime(v int64) *DescribeImageListWithBaselineNameResponseBodyImageInfos {
	s.FirstScanTime = &v
	return s
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) SetHighRiskImage(v int32) *DescribeImageListWithBaselineNameResponseBodyImageInfos {
	s.HighRiskImage = &v
	return s
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) SetImage(v string) *DescribeImageListWithBaselineNameResponseBodyImageInfos {
	s.Image = &v
	return s
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) SetImageCreate(v int64) *DescribeImageListWithBaselineNameResponseBodyImageInfos {
	s.ImageCreate = &v
	return s
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) SetImageId(v string) *DescribeImageListWithBaselineNameResponseBodyImageInfos {
	s.ImageId = &v
	return s
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) SetImageSize(v int32) *DescribeImageListWithBaselineNameResponseBodyImageInfos {
	s.ImageSize = &v
	return s
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) SetImageUpdate(v int64) *DescribeImageListWithBaselineNameResponseBodyImageInfos {
	s.ImageUpdate = &v
	return s
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) SetInstanceId(v string) *DescribeImageListWithBaselineNameResponseBodyImageInfos {
	s.InstanceId = &v
	return s
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) SetInstanceName(v string) *DescribeImageListWithBaselineNameResponseBodyImageInfos {
	s.InstanceName = &v
	return s
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) SetInternetIp(v string) *DescribeImageListWithBaselineNameResponseBodyImageInfos {
	s.InternetIp = &v
	return s
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) SetIntranetIp(v string) *DescribeImageListWithBaselineNameResponseBodyImageInfos {
	s.IntranetIp = &v
	return s
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) SetLastScanTime(v int64) *DescribeImageListWithBaselineNameResponseBodyImageInfos {
	s.LastScanTime = &v
	return s
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) SetLowRiskImage(v int32) *DescribeImageListWithBaselineNameResponseBodyImageInfos {
	s.LowRiskImage = &v
	return s
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) SetMiddleRiskImage(v int32) *DescribeImageListWithBaselineNameResponseBodyImageInfos {
	s.MiddleRiskImage = &v
	return s
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) SetNamespace(v string) *DescribeImageListWithBaselineNameResponseBodyImageInfos {
	s.Namespace = &v
	return s
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) SetNoRiskImage(v int32) *DescribeImageListWithBaselineNameResponseBodyImageInfos {
	s.NoRiskImage = &v
	return s
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) SetPod(v string) *DescribeImageListWithBaselineNameResponseBodyImageInfos {
	s.Pod = &v
	return s
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) SetRegionId(v string) *DescribeImageListWithBaselineNameResponseBodyImageInfos {
	s.RegionId = &v
	return s
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) SetRepoId(v string) *DescribeImageListWithBaselineNameResponseBodyImageInfos {
	s.RepoId = &v
	return s
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) SetRepoName(v string) *DescribeImageListWithBaselineNameResponseBodyImageInfos {
	s.RepoName = &v
	return s
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) SetRepoNamespace(v string) *DescribeImageListWithBaselineNameResponseBodyImageInfos {
	s.RepoNamespace = &v
	return s
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) SetRepoType(v string) *DescribeImageListWithBaselineNameResponseBodyImageInfos {
	s.RepoType = &v
	return s
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) SetRiskStatus(v string) *DescribeImageListWithBaselineNameResponseBodyImageInfos {
	s.RiskStatus = &v
	return s
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) SetTag(v string) *DescribeImageListWithBaselineNameResponseBodyImageInfos {
	s.Tag = &v
	return s
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) SetTargetId(v string) *DescribeImageListWithBaselineNameResponseBodyImageInfos {
	s.TargetId = &v
	return s
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) SetTargetName(v string) *DescribeImageListWithBaselineNameResponseBodyImageInfos {
	s.TargetName = &v
	return s
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) SetTargetType(v string) *DescribeImageListWithBaselineNameResponseBodyImageInfos {
	s.TargetType = &v
	return s
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) SetTotalItemCount(v int32) *DescribeImageListWithBaselineNameResponseBodyImageInfos {
	s.TotalItemCount = &v
	return s
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) SetUuid(v string) *DescribeImageListWithBaselineNameResponseBodyImageInfos {
	s.Uuid = &v
	return s
}

func (s *DescribeImageListWithBaselineNameResponseBodyImageInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeImageListWithBaselineNameResponseBodyPageInfo struct {
	// The number of the images returned on the current page.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of images on which baseline risks are detected.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeImageListWithBaselineNameResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageListWithBaselineNameResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeImageListWithBaselineNameResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeImageListWithBaselineNameResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeImageListWithBaselineNameResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageListWithBaselineNameResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeImageListWithBaselineNameResponseBodyPageInfo) SetCount(v int32) *DescribeImageListWithBaselineNameResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeImageListWithBaselineNameResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeImageListWithBaselineNameResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageListWithBaselineNameResponseBodyPageInfo) SetPageSize(v int32) *DescribeImageListWithBaselineNameResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeImageListWithBaselineNameResponseBodyPageInfo) SetTotalCount(v int32) *DescribeImageListWithBaselineNameResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeImageListWithBaselineNameResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
