// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAffectedMaliciousFileImagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAffectedMaliciousFileImagesResponse(v []*DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) *DescribeAffectedMaliciousFileImagesResponseBody
	GetAffectedMaliciousFileImagesResponse() []*DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse
	SetPageInfo(v *DescribeAffectedMaliciousFileImagesResponseBodyPageInfo) *DescribeAffectedMaliciousFileImagesResponseBody
	GetPageInfo() *DescribeAffectedMaliciousFileImagesResponseBodyPageInfo
	SetRequestId(v string) *DescribeAffectedMaliciousFileImagesResponseBody
	GetRequestId() *string
}

type DescribeAffectedMaliciousFileImagesResponseBody struct {
	// An array consisting of the images that have malicious image samples.
	AffectedMaliciousFileImagesResponse []*DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse `json:"AffectedMaliciousFileImagesResponse,omitempty" xml:"AffectedMaliciousFileImagesResponse,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeAffectedMaliciousFileImagesResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// ACF97412-FD09-4D1F-994F-34DF12BREF20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAffectedMaliciousFileImagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAffectedMaliciousFileImagesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAffectedMaliciousFileImagesResponseBody) GetAffectedMaliciousFileImagesResponse() []*DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	return s.AffectedMaliciousFileImagesResponse
}

func (s *DescribeAffectedMaliciousFileImagesResponseBody) GetPageInfo() *DescribeAffectedMaliciousFileImagesResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeAffectedMaliciousFileImagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAffectedMaliciousFileImagesResponseBody) SetAffectedMaliciousFileImagesResponse(v []*DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) *DescribeAffectedMaliciousFileImagesResponseBody {
	s.AffectedMaliciousFileImagesResponse = v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBody) SetPageInfo(v *DescribeAffectedMaliciousFileImagesResponseBodyPageInfo) *DescribeAffectedMaliciousFileImagesResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBody) SetRequestId(v string) *DescribeAffectedMaliciousFileImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse struct {
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
	// The image digest.
	//
	// example:
	//
	// 6a5e1031a5858617f7d8a179ead6****
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The URL to download the malicious image sample.
	//
	// example:
	//
	// https://aegis-metadata-file.oss-cn-shanghai.aliyuncs.com/
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The path to the image file.
	//
	// example:
	//
	// /d836968041f7683b5605a****
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// The timestamp of the first scan.
	//
	// example:
	//
	// 1594907349000
	FirstScanTimestamp *int64 `json:"FirstScanTimestamp,omitempty" xml:"FirstScanTimestamp,omitempty"`
	// The text that is highlighted.
	//
	// example:
	//
	// {"ruleVersion":"highlight_20210908","ruleId":600106,"events":[[2,54]]}
	HighLight *string `json:"HighLight,omitempty" xml:"HighLight,omitempty"`
	// The ID of alert event.
	//
	// example:
	//
	// 1000040
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the image.
	//
	// example:
	//
	// registry.cn-wulanchabu.aliyuncs.com/sas_test/huxin-test-001:nuxeo6-conta****
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The UUID of the image.
	//
	// example:
	//
	// e05c0de798217637868ef4****
	ImageUuid *string `json:"ImageUuid,omitempty" xml:"ImageUuid,omitempty"`
	// The name of the ECS instance.
	//
	// example:
	//
	// sql-test-001
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the server.
	//
	// example:
	//
	// 47.101.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the server.
	//
	// example:
	//
	// 172.22.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The timestamp of the last scan.
	//
	// example:
	//
	// 1596522785000
	LatestScanTimestamp *int64 `json:"LatestScanTimestamp,omitempty" xml:"LatestScanTimestamp,omitempty"`
	// The timestamp of the last verification.
	//
	// example:
	//
	// 1596522711000
	LatestVerifyTimestamp *int64 `json:"LatestVerifyTimestamp,omitempty" xml:"LatestVerifyTimestamp,omitempty"`
	// The image layer.
	//
	// example:
	//
	// 27213ad3447f0209dd152a5cadea****
	Layer *string `json:"Layer,omitempty" xml:"Layer,omitempty"`
	// The severity of the malicious image sample. Valid values:
	//
	// 	- **serious**
	//
	// 	- **suspicious**
	//
	// 	- **remind**
	//
	// example:
	//
	// serious
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The MD5 hash value of the malicious image sample.
	//
	// example:
	//
	// d836968041f768300d9605a****
	MaliciousMd5 *string `json:"MaliciousMd5,omitempty" xml:"MaliciousMd5,omitempty"`
	// The namespace to which the image repository belongs.
	//
	// example:
	//
	// hanghai-namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The pod.
	//
	// example:
	//
	// 22222-7xsqq
	Pod *string `json:"Pod,omitempty" xml:"Pod,omitempty"`
	// The ID of the image repository.
	//
	// example:
	//
	// crr-vridcl4****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The ID of the container image.
	//
	// example:
	//
	// cri-datvail3m****
	RepoInstanceId *string `json:"RepoInstanceId,omitempty" xml:"RepoInstanceId,omitempty"`
	// The name of the image repository.
	//
	// example:
	//
	// centos
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The region ID of the image repository.
	//
	// example:
	//
	// cn-shanghai
	RepoRegionId *string `json:"RepoRegionId,omitempty" xml:"RepoRegionId,omitempty"`
	// The handling status of the malicious image sample. Valid values:
	//
	// 	- **0**: unhandled
	//
	// 	- **1**: handled
	//
	// 	- **2**: verifying
	//
	// 	- **3**: added to the whitelist
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag that is added to the image.
	//
	// example:
	//
	// 0.2
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The ID of the task object.
	//
	// example:
	//
	// m-bp17m0pc0xprzbwo****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the task object.
	//
	// example:
	//
	// source-test-obj-9LaLJ
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The object type. Valid value:
	//
	// 	- **ECS_IMAGE**
	//
	// 	- **ECS_SNAPSHOT**
	//
	// example:
	//
	// ECS_IMAGE
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// 48a473c4-1650-4931-a822-7e6c2c28****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) GetContainerId() *string {
	return s.ContainerId
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) GetDigest() *string {
	return s.Digest
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) GetFilePath() *string {
	return s.FilePath
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) GetFirstScanTimestamp() *int64 {
	return s.FirstScanTimestamp
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) GetHighLight() *string {
	return s.HighLight
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) GetId() *int64 {
	return s.Id
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) GetImage() *string {
	return s.Image
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) GetImageUuid() *string {
	return s.ImageUuid
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) GetLatestScanTimestamp() *int64 {
	return s.LatestScanTimestamp
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) GetLatestVerifyTimestamp() *int64 {
	return s.LatestVerifyTimestamp
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) GetLayer() *string {
	return s.Layer
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) GetLevel() *string {
	return s.Level
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) GetMaliciousMd5() *string {
	return s.MaliciousMd5
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) GetPod() *string {
	return s.Pod
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) GetRepoId() *string {
	return s.RepoId
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) GetRepoInstanceId() *string {
	return s.RepoInstanceId
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) GetRepoName() *string {
	return s.RepoName
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) GetRepoRegionId() *string {
	return s.RepoRegionId
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) GetTag() *string {
	return s.Tag
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) GetTargetId() *string {
	return s.TargetId
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) GetTargetName() *string {
	return s.TargetName
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) GetTargetType() *string {
	return s.TargetType
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetClusterId(v string) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.ClusterId = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetClusterName(v string) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.ClusterName = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetContainerId(v string) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.ContainerId = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetDigest(v string) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.Digest = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetDownloadUrl(v string) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.DownloadUrl = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetFilePath(v string) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.FilePath = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetFirstScanTimestamp(v int64) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.FirstScanTimestamp = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetHighLight(v string) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.HighLight = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetId(v int64) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.Id = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetImage(v string) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.Image = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetImageUuid(v string) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.ImageUuid = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetInstanceName(v string) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.InstanceName = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetInternetIp(v string) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.InternetIp = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetIntranetIp(v string) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.IntranetIp = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetLatestScanTimestamp(v int64) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.LatestScanTimestamp = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetLatestVerifyTimestamp(v int64) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.LatestVerifyTimestamp = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetLayer(v string) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.Layer = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetLevel(v string) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.Level = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetMaliciousMd5(v string) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.MaliciousMd5 = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetNamespace(v string) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.Namespace = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetPod(v string) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.Pod = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetRepoId(v string) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.RepoId = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetRepoInstanceId(v string) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.RepoInstanceId = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetRepoName(v string) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.RepoName = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetRepoRegionId(v string) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.RepoRegionId = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetStatus(v int32) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.Status = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetTag(v string) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.Tag = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetTargetId(v string) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.TargetId = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetTargetName(v string) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.TargetName = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetTargetType(v string) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.TargetType = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetUuid(v string) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.Uuid = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) Validate() error {
	return dara.Validate(s)
}

type DescribeAffectedMaliciousFileImagesResponseBodyPageInfo struct {
	// The number of images that have malicious image samples returned on the current page.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page. Pages start from page **1**. Default value: **1**.
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
	// The total number of images that have malicious image samples.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAffectedMaliciousFileImagesResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeAffectedMaliciousFileImagesResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyPageInfo) SetCount(v int32) *DescribeAffectedMaliciousFileImagesResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeAffectedMaliciousFileImagesResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyPageInfo) SetPageSize(v int32) *DescribeAffectedMaliciousFileImagesResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyPageInfo) SetTotalCount(v int32) *DescribeAffectedMaliciousFileImagesResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
