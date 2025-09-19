// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAffectedMaliciousFileImagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeAffectedMaliciousFileImagesRequest
	GetClusterId() *string
	SetClusterName(v string) *DescribeAffectedMaliciousFileImagesRequest
	GetClusterName() *string
	SetContainerId(v string) *DescribeAffectedMaliciousFileImagesRequest
	GetContainerId() *string
	SetCurrentPage(v int32) *DescribeAffectedMaliciousFileImagesRequest
	GetCurrentPage() *int32
	SetImage(v string) *DescribeAffectedMaliciousFileImagesRequest
	GetImage() *string
	SetImageDigest(v string) *DescribeAffectedMaliciousFileImagesRequest
	GetImageDigest() *string
	SetImageLayer(v string) *DescribeAffectedMaliciousFileImagesRequest
	GetImageLayer() *string
	SetImageTag(v string) *DescribeAffectedMaliciousFileImagesRequest
	GetImageTag() *string
	SetLang(v string) *DescribeAffectedMaliciousFileImagesRequest
	GetLang() *string
	SetLevels(v string) *DescribeAffectedMaliciousFileImagesRequest
	GetLevels() *string
	SetMaliciousMd5(v string) *DescribeAffectedMaliciousFileImagesRequest
	GetMaliciousMd5() *string
	SetNamespace(v string) *DescribeAffectedMaliciousFileImagesRequest
	GetNamespace() *string
	SetPageSize(v string) *DescribeAffectedMaliciousFileImagesRequest
	GetPageSize() *string
	SetPod(v string) *DescribeAffectedMaliciousFileImagesRequest
	GetPod() *string
	SetRepoId(v string) *DescribeAffectedMaliciousFileImagesRequest
	GetRepoId() *string
	SetRepoInstanceId(v string) *DescribeAffectedMaliciousFileImagesRequest
	GetRepoInstanceId() *string
	SetRepoName(v string) *DescribeAffectedMaliciousFileImagesRequest
	GetRepoName() *string
	SetRepoNamespace(v string) *DescribeAffectedMaliciousFileImagesRequest
	GetRepoNamespace() *string
	SetRepoRegionId(v string) *DescribeAffectedMaliciousFileImagesRequest
	GetRepoRegionId() *string
	SetScanRange(v []*string) *DescribeAffectedMaliciousFileImagesRequest
	GetScanRange() []*string
	SetStatus(v string) *DescribeAffectedMaliciousFileImagesRequest
	GetStatus() *string
}

type DescribeAffectedMaliciousFileImagesRequest struct {
	// The ID of the container cluster.
	//
	// >  You can call the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) operation to query the IDs of container clusters.
	//
	// example:
	//
	// c60b77fe62093480db6164a3c2fa5****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// minikube
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The ID of the container.
	//
	// example:
	//
	// cc20a1024011c44b6a8710d6f8b****
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	// The number of the page to return. Pages start from page **1**. Default value: **1**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The name of the container image.
	//
	// example:
	//
	// registry.cn-wulanchabu.aliyuncs.com/sas_test/huxin-test-001:nuxeo6-****
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The image digest.
	//
	// example:
	//
	// 6a5e103187b31a94592a47a5858617f7a179ead61df7606****
	ImageDigest *string `json:"ImageDigest,omitempty" xml:"ImageDigest,omitempty"`
	// The image layer.
	//
	// example:
	//
	// 27213ad375b53628dd152a5ca****
	ImageLayer *string `json:"ImageLayer,omitempty" xml:"ImageLayer,omitempty"`
	// The tag that is added to the image.
	//
	// example:
	//
	// 0.2
	ImageTag *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	// The language of the content within the request and the response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The severity level of the malicious image sample. Separate multiple severity levels with commas (,). Valid values:
	//
	// 	- **serious**
	//
	// 	- **suspicious**
	//
	// 	- **remind**
	//
	// example:
	//
	// serious,suspicious
	Levels *string `json:"Levels,omitempty" xml:"Levels,omitempty"`
	// The MD5 hash value of the malicious image sample.
	//
	// >  You can call the [DescribeGroupedMaliciousFiles](~~DescribeGroupedMaliciousFiles~~) operation to query the MD5 hash values of malicious image samples.
	//
	// example:
	//
	// d836968041f7683b5459****
	MaliciousMd5 *string `json:"MaliciousMd5,omitempty" xml:"MaliciousMd5,omitempty"`
	// The namespace.
	//
	// example:
	//
	// test-002
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The number of entries to return on each page. Default value: **20**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The pod.
	//
	// example:
	//
	// 22222-7xsqq
	Pod *string `json:"Pod,omitempty" xml:"Pod,omitempty"`
	// The ID of the image repository.
	//
	// >  You can call the [ListRepository](https://help.aliyun.com/document_detail/451339.html) operation to query the IDs of image repositories from the value of the **RepoId*	- response parameter.
	//
	// example:
	//
	// crr-vridcl4****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The ID of the container image.
	//
	// >  You can call the [ListRepository](https://help.aliyun.com/document_detail/451339.html) operation to query the IDs of container images from the value of the **InstanceId*	- response parameter.
	//
	// example:
	//
	// cri-datvailb****
	RepoInstanceId *string `json:"RepoInstanceId,omitempty" xml:"RepoInstanceId,omitempty"`
	// The name of the image repository.
	//
	// >  Fuzzy match is supported.
	//
	// example:
	//
	// centos
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The namespace to which the image repository belongs.
	//
	// >  Fuzzy match is supported.
	//
	// example:
	//
	// hanghai-namespace
	RepoNamespace *string `json:"RepoNamespace,omitempty" xml:"RepoNamespace,omitempty"`
	// The region ID of the image repository. Valid values:
	//
	// 	- **cn-beijing**: China (Beijing)
	//
	// 	- **cn-zhangjiakou**: China (Zhangjiakou)
	//
	// 	- **cn-hangzhou**: China (Hangzhou)
	//
	// 	- **cn-shanghai**: China (Shanghai)
	//
	// 	- **cn-shenzhen**: China (Shenzhen)
	//
	// 	- **cn-hongkong**: China (Hong Kong)
	//
	// 	- **ap-southeast-1**: Singapore
	//
	// 	- **ap-southeast-5**: Indonesia (Jakarta)
	//
	// 	- **us-east-1**: US (Virginia)
	//
	// 	- **us-west-1**: US (Silicon Valley)
	//
	// 	- **eu-central-1**: Germany (Frankfurt)
	//
	// 	- **eu-west-1**: UK (London)
	//
	// example:
	//
	// cn-hangzhou
	RepoRegionId *string `json:"RepoRegionId,omitempty" xml:"RepoRegionId,omitempty"`
	// The types of the assets that you want to scan.
	ScanRange []*string `json:"ScanRange,omitempty" xml:"ScanRange,omitempty" type:"Repeated"`
	// The status of the malicious image sample. Valid values:
	//
	// 	- **0**: The malicious image sample is not handled.
	//
	// 	- **1**: The malicious image sample is handled.
	//
	// 	- **2**: The malicious image sample is being verified.
	//
	// 	- **3**: The malicious image sample is added to the whitelist.
	//
	// example:
	//
	// 0
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeAffectedMaliciousFileImagesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAffectedMaliciousFileImagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAffectedMaliciousFileImagesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeAffectedMaliciousFileImagesRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeAffectedMaliciousFileImagesRequest) GetContainerId() *string {
	return s.ContainerId
}

func (s *DescribeAffectedMaliciousFileImagesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeAffectedMaliciousFileImagesRequest) GetImage() *string {
	return s.Image
}

func (s *DescribeAffectedMaliciousFileImagesRequest) GetImageDigest() *string {
	return s.ImageDigest
}

func (s *DescribeAffectedMaliciousFileImagesRequest) GetImageLayer() *string {
	return s.ImageLayer
}

func (s *DescribeAffectedMaliciousFileImagesRequest) GetImageTag() *string {
	return s.ImageTag
}

func (s *DescribeAffectedMaliciousFileImagesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAffectedMaliciousFileImagesRequest) GetLevels() *string {
	return s.Levels
}

func (s *DescribeAffectedMaliciousFileImagesRequest) GetMaliciousMd5() *string {
	return s.MaliciousMd5
}

func (s *DescribeAffectedMaliciousFileImagesRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeAffectedMaliciousFileImagesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeAffectedMaliciousFileImagesRequest) GetPod() *string {
	return s.Pod
}

func (s *DescribeAffectedMaliciousFileImagesRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *DescribeAffectedMaliciousFileImagesRequest) GetRepoInstanceId() *string {
	return s.RepoInstanceId
}

func (s *DescribeAffectedMaliciousFileImagesRequest) GetRepoName() *string {
	return s.RepoName
}

func (s *DescribeAffectedMaliciousFileImagesRequest) GetRepoNamespace() *string {
	return s.RepoNamespace
}

func (s *DescribeAffectedMaliciousFileImagesRequest) GetRepoRegionId() *string {
	return s.RepoRegionId
}

func (s *DescribeAffectedMaliciousFileImagesRequest) GetScanRange() []*string {
	return s.ScanRange
}

func (s *DescribeAffectedMaliciousFileImagesRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeAffectedMaliciousFileImagesRequest) SetClusterId(v string) *DescribeAffectedMaliciousFileImagesRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesRequest) SetClusterName(v string) *DescribeAffectedMaliciousFileImagesRequest {
	s.ClusterName = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesRequest) SetContainerId(v string) *DescribeAffectedMaliciousFileImagesRequest {
	s.ContainerId = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesRequest) SetCurrentPage(v int32) *DescribeAffectedMaliciousFileImagesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesRequest) SetImage(v string) *DescribeAffectedMaliciousFileImagesRequest {
	s.Image = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesRequest) SetImageDigest(v string) *DescribeAffectedMaliciousFileImagesRequest {
	s.ImageDigest = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesRequest) SetImageLayer(v string) *DescribeAffectedMaliciousFileImagesRequest {
	s.ImageLayer = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesRequest) SetImageTag(v string) *DescribeAffectedMaliciousFileImagesRequest {
	s.ImageTag = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesRequest) SetLang(v string) *DescribeAffectedMaliciousFileImagesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesRequest) SetLevels(v string) *DescribeAffectedMaliciousFileImagesRequest {
	s.Levels = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesRequest) SetMaliciousMd5(v string) *DescribeAffectedMaliciousFileImagesRequest {
	s.MaliciousMd5 = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesRequest) SetNamespace(v string) *DescribeAffectedMaliciousFileImagesRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesRequest) SetPageSize(v string) *DescribeAffectedMaliciousFileImagesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesRequest) SetPod(v string) *DescribeAffectedMaliciousFileImagesRequest {
	s.Pod = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesRequest) SetRepoId(v string) *DescribeAffectedMaliciousFileImagesRequest {
	s.RepoId = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesRequest) SetRepoInstanceId(v string) *DescribeAffectedMaliciousFileImagesRequest {
	s.RepoInstanceId = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesRequest) SetRepoName(v string) *DescribeAffectedMaliciousFileImagesRequest {
	s.RepoName = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesRequest) SetRepoNamespace(v string) *DescribeAffectedMaliciousFileImagesRequest {
	s.RepoNamespace = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesRequest) SetRepoRegionId(v string) *DescribeAffectedMaliciousFileImagesRequest {
	s.RepoRegionId = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesRequest) SetScanRange(v []*string) *DescribeAffectedMaliciousFileImagesRequest {
	s.ScanRange = v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesRequest) SetStatus(v string) *DescribeAffectedMaliciousFileImagesRequest {
	s.Status = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesRequest) Validate() error {
	return dara.Validate(s)
}
