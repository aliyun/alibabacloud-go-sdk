// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupedMaliciousFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeGroupedMaliciousFilesRequest
	GetClusterId() *string
	SetCurrentPage(v int32) *DescribeGroupedMaliciousFilesRequest
	GetCurrentPage() *int32
	SetFuzzyMaliciousName(v string) *DescribeGroupedMaliciousFilesRequest
	GetFuzzyMaliciousName() *string
	SetImageDigest(v string) *DescribeGroupedMaliciousFilesRequest
	GetImageDigest() *string
	SetImageLayer(v string) *DescribeGroupedMaliciousFilesRequest
	GetImageLayer() *string
	SetImageTag(v string) *DescribeGroupedMaliciousFilesRequest
	GetImageTag() *string
	SetLang(v string) *DescribeGroupedMaliciousFilesRequest
	GetLang() *string
	SetLevels(v string) *DescribeGroupedMaliciousFilesRequest
	GetLevels() *string
	SetMaliciousMd5(v string) *DescribeGroupedMaliciousFilesRequest
	GetMaliciousMd5() *string
	SetPageSize(v string) *DescribeGroupedMaliciousFilesRequest
	GetPageSize() *string
	SetRepoId(v string) *DescribeGroupedMaliciousFilesRequest
	GetRepoId() *string
	SetRepoInstanceId(v string) *DescribeGroupedMaliciousFilesRequest
	GetRepoInstanceId() *string
	SetRepoName(v string) *DescribeGroupedMaliciousFilesRequest
	GetRepoName() *string
	SetRepoNamespace(v string) *DescribeGroupedMaliciousFilesRequest
	GetRepoNamespace() *string
	SetRepoRegionId(v string) *DescribeGroupedMaliciousFilesRequest
	GetRepoRegionId() *string
	SetScanRange(v []*string) *DescribeGroupedMaliciousFilesRequest
	GetScanRange() []*string
}

type DescribeGroupedMaliciousFilesRequest struct {
	// The cluster ID of the container on which the malicious image sample is detected.
	//
	// example:
	//
	// c556c8133b5ad4378b7fc533ddbda****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The number of the page to return. Pages start from page **1**. Default value: **1**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The name of the malicious image sample that you want to query.
	//
	// >  Fuzzy match is supported.
	//
	// example:
	//
	// Mining
	FuzzyMaliciousName *string `json:"FuzzyMaliciousName,omitempty" xml:"FuzzyMaliciousName,omitempty"`
	// The image digest.
	//
	// example:
	//
	// 6a5e103187b31a94592a47a5858617f7****
	ImageDigest *string `json:"ImageDigest,omitempty" xml:"ImageDigest,omitempty"`
	// The image layer.
	//
	// example:
	//
	// 27213ad375b53628dd152a5ca****
	ImageLayer *string `json:"ImageLayer,omitempty" xml:"ImageLayer,omitempty"`
	// The image tag.
	//
	// example:
	//
	// 0.2
	ImageTag *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The severity of the malicious image sample that you want to query. You can enter multiple severities. Separate the severities with commas (,). Valid values:
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
	Levels *string `json:"Levels,omitempty" xml:"Levels,omitempty"`
	// The MD5 hash value of the malicious image sample.
	//
	// example:
	//
	// d836968041f7683b5459****
	MaliciousMd5 *string `json:"MaliciousMd5,omitempty" xml:"MaliciousMd5,omitempty"`
	// The number of entries to return on each page. Default value: **20**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the image repository.
	//
	// >  You can call the [ListRepository](https://help.aliyun.com/document_detail/145293.html) operation to query the IDs of image repositories from the value of the **RepoId*	- response parameter.
	//
	// example:
	//
	// crr-vridcl4****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The ID of the container image.
	//
	// >  You can call the [ListRepository](https://help.aliyun.com/document_detail/145293.html) operation to query the IDs of container images from the value of the **InstanceId*	- response parameter.
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
	// cn-shanghai
	RepoRegionId *string `json:"RepoRegionId,omitempty" xml:"RepoRegionId,omitempty"`
	// The types of the assets that you want to scan.
	ScanRange []*string `json:"ScanRange,omitempty" xml:"ScanRange,omitempty" type:"Repeated"`
}

func (s DescribeGroupedMaliciousFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupedMaliciousFilesRequest) GoString() string {
	return s.String()
}

func (s *DescribeGroupedMaliciousFilesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeGroupedMaliciousFilesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeGroupedMaliciousFilesRequest) GetFuzzyMaliciousName() *string {
	return s.FuzzyMaliciousName
}

func (s *DescribeGroupedMaliciousFilesRequest) GetImageDigest() *string {
	return s.ImageDigest
}

func (s *DescribeGroupedMaliciousFilesRequest) GetImageLayer() *string {
	return s.ImageLayer
}

func (s *DescribeGroupedMaliciousFilesRequest) GetImageTag() *string {
	return s.ImageTag
}

func (s *DescribeGroupedMaliciousFilesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeGroupedMaliciousFilesRequest) GetLevels() *string {
	return s.Levels
}

func (s *DescribeGroupedMaliciousFilesRequest) GetMaliciousMd5() *string {
	return s.MaliciousMd5
}

func (s *DescribeGroupedMaliciousFilesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeGroupedMaliciousFilesRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *DescribeGroupedMaliciousFilesRequest) GetRepoInstanceId() *string {
	return s.RepoInstanceId
}

func (s *DescribeGroupedMaliciousFilesRequest) GetRepoName() *string {
	return s.RepoName
}

func (s *DescribeGroupedMaliciousFilesRequest) GetRepoNamespace() *string {
	return s.RepoNamespace
}

func (s *DescribeGroupedMaliciousFilesRequest) GetRepoRegionId() *string {
	return s.RepoRegionId
}

func (s *DescribeGroupedMaliciousFilesRequest) GetScanRange() []*string {
	return s.ScanRange
}

func (s *DescribeGroupedMaliciousFilesRequest) SetClusterId(v string) *DescribeGroupedMaliciousFilesRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesRequest) SetCurrentPage(v int32) *DescribeGroupedMaliciousFilesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesRequest) SetFuzzyMaliciousName(v string) *DescribeGroupedMaliciousFilesRequest {
	s.FuzzyMaliciousName = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesRequest) SetImageDigest(v string) *DescribeGroupedMaliciousFilesRequest {
	s.ImageDigest = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesRequest) SetImageLayer(v string) *DescribeGroupedMaliciousFilesRequest {
	s.ImageLayer = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesRequest) SetImageTag(v string) *DescribeGroupedMaliciousFilesRequest {
	s.ImageTag = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesRequest) SetLang(v string) *DescribeGroupedMaliciousFilesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesRequest) SetLevels(v string) *DescribeGroupedMaliciousFilesRequest {
	s.Levels = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesRequest) SetMaliciousMd5(v string) *DescribeGroupedMaliciousFilesRequest {
	s.MaliciousMd5 = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesRequest) SetPageSize(v string) *DescribeGroupedMaliciousFilesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesRequest) SetRepoId(v string) *DescribeGroupedMaliciousFilesRequest {
	s.RepoId = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesRequest) SetRepoInstanceId(v string) *DescribeGroupedMaliciousFilesRequest {
	s.RepoInstanceId = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesRequest) SetRepoName(v string) *DescribeGroupedMaliciousFilesRequest {
	s.RepoName = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesRequest) SetRepoNamespace(v string) *DescribeGroupedMaliciousFilesRequest {
	s.RepoNamespace = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesRequest) SetRepoRegionId(v string) *DescribeGroupedMaliciousFilesRequest {
	s.RepoRegionId = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesRequest) SetScanRange(v []*string) *DescribeGroupedMaliciousFilesRequest {
	s.ScanRange = v
	return s
}

func (s *DescribeGroupedMaliciousFilesRequest) Validate() error {
	return dara.Validate(s)
}
