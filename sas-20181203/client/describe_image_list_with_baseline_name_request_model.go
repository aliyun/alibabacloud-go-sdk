// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageListWithBaselineNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaselineNameKey(v string) *DescribeImageListWithBaselineNameRequest
	GetBaselineNameKey() *string
	SetClusterId(v string) *DescribeImageListWithBaselineNameRequest
	GetClusterId() *string
	SetClusterName(v string) *DescribeImageListWithBaselineNameRequest
	GetClusterName() *string
	SetContainerId(v string) *DescribeImageListWithBaselineNameRequest
	GetContainerId() *string
	SetCriteria(v string) *DescribeImageListWithBaselineNameRequest
	GetCriteria() *string
	SetCriteriaType(v string) *DescribeImageListWithBaselineNameRequest
	GetCriteriaType() *string
	SetCurrentPage(v int32) *DescribeImageListWithBaselineNameRequest
	GetCurrentPage() *int32
	SetImage(v string) *DescribeImageListWithBaselineNameRequest
	GetImage() *string
	SetImageDigest(v string) *DescribeImageListWithBaselineNameRequest
	GetImageDigest() *string
	SetLang(v string) *DescribeImageListWithBaselineNameRequest
	GetLang() *string
	SetNamespace(v string) *DescribeImageListWithBaselineNameRequest
	GetNamespace() *string
	SetPageSize(v int32) *DescribeImageListWithBaselineNameRequest
	GetPageSize() *int32
	SetPod(v string) *DescribeImageListWithBaselineNameRequest
	GetPod() *string
	SetRepoInstanceId(v string) *DescribeImageListWithBaselineNameRequest
	GetRepoInstanceId() *string
	SetRepoName(v string) *DescribeImageListWithBaselineNameRequest
	GetRepoName() *string
	SetRepoNamespace(v string) *DescribeImageListWithBaselineNameRequest
	GetRepoNamespace() *string
	SetScanRange(v []*string) *DescribeImageListWithBaselineNameRequest
	GetScanRange() []*string
}

type DescribeImageListWithBaselineNameRequest struct {
	// The name of the image baseline check result.
	//
	// This parameter is required.
	//
	// example:
	//
	// ak_leak
	BaselineNameKey *string `json:"BaselineNameKey,omitempty" xml:"BaselineNameKey,omitempty"`
	// The ID of the container cluster to query.
	//
	// > You can call the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) operation to obtain this parameter.
	//
	// example:
	//
	// cc20a1024011c44b6a8710d6f8b****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// docker-law
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The container ID.
	//
	// example:
	//
	// c08d5fc1a329a4b88950a253d082f****
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	// The query condition for the baseline.
	//
	// example:
	//
	// Unauthorized access
	Criteria *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// The query type for the baseline. Valid values:
	//
	// - **BaselineNameAlias**: baseline name
	//
	// - **BaselineClassAlias**: baseline category.
	//
	// example:
	//
	// BaselineNameAlias
	CriteriaType *string `json:"CriteriaType,omitempty" xml:"CriteriaType,omitempty"`
	// The page number of the page to return. Default value: **1**, which indicates that the first page is returned.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The container image name.
	//
	// example:
	//
	// registry.cn-wulanchabu.aliyuncs.com/sas_test/huxin-test-001:nuxeo6-****
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The SHA256 value of the image digest.
	//
	// example:
	//
	// 2e6daffce524ffeae66cccaa90c8fc47de912346dcec295c27395b6d66db6423
	ImageDigest *string `json:"ImageDigest,omitempty" xml:"ImageDigest,omitempty"`
	// The language type for the request and response messages. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The namespace.
	//
	// example:
	//
	// test-002
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// Settings for paged query paging. The number of image baseline check result details to display per page. Default value: **10**, which indicates that 10 image baseline check result details are displayed per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The pod.
	//
	// example:
	//
	// 22222-7xsqq
	Pod *string `json:"Pod,omitempty" xml:"Pod,omitempty"`
	// The instance ID of the image repository.
	//
	// example:
	//
	// i-qewqrqcsadf****
	RepoInstanceId *string `json:"RepoInstanceId,omitempty" xml:"RepoInstanceId,omitempty"`
	// The name of the image repository.
	//
	// example:
	//
	// libssh2
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The namespace of the image repository.
	//
	// example:
	//
	// libssh2
	RepoNamespace *string `json:"RepoNamespace,omitempty" xml:"RepoNamespace,omitempty"`
	// The collection of scan ranges.
	ScanRange []*string `json:"ScanRange,omitempty" xml:"ScanRange,omitempty" type:"Repeated"`
}

func (s DescribeImageListWithBaselineNameRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageListWithBaselineNameRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageListWithBaselineNameRequest) GetBaselineNameKey() *string {
	return s.BaselineNameKey
}

func (s *DescribeImageListWithBaselineNameRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeImageListWithBaselineNameRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeImageListWithBaselineNameRequest) GetContainerId() *string {
	return s.ContainerId
}

func (s *DescribeImageListWithBaselineNameRequest) GetCriteria() *string {
	return s.Criteria
}

func (s *DescribeImageListWithBaselineNameRequest) GetCriteriaType() *string {
	return s.CriteriaType
}

func (s *DescribeImageListWithBaselineNameRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeImageListWithBaselineNameRequest) GetImage() *string {
	return s.Image
}

func (s *DescribeImageListWithBaselineNameRequest) GetImageDigest() *string {
	return s.ImageDigest
}

func (s *DescribeImageListWithBaselineNameRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeImageListWithBaselineNameRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeImageListWithBaselineNameRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageListWithBaselineNameRequest) GetPod() *string {
	return s.Pod
}

func (s *DescribeImageListWithBaselineNameRequest) GetRepoInstanceId() *string {
	return s.RepoInstanceId
}

func (s *DescribeImageListWithBaselineNameRequest) GetRepoName() *string {
	return s.RepoName
}

func (s *DescribeImageListWithBaselineNameRequest) GetRepoNamespace() *string {
	return s.RepoNamespace
}

func (s *DescribeImageListWithBaselineNameRequest) GetScanRange() []*string {
	return s.ScanRange
}

func (s *DescribeImageListWithBaselineNameRequest) SetBaselineNameKey(v string) *DescribeImageListWithBaselineNameRequest {
	s.BaselineNameKey = &v
	return s
}

func (s *DescribeImageListWithBaselineNameRequest) SetClusterId(v string) *DescribeImageListWithBaselineNameRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeImageListWithBaselineNameRequest) SetClusterName(v string) *DescribeImageListWithBaselineNameRequest {
	s.ClusterName = &v
	return s
}

func (s *DescribeImageListWithBaselineNameRequest) SetContainerId(v string) *DescribeImageListWithBaselineNameRequest {
	s.ContainerId = &v
	return s
}

func (s *DescribeImageListWithBaselineNameRequest) SetCriteria(v string) *DescribeImageListWithBaselineNameRequest {
	s.Criteria = &v
	return s
}

func (s *DescribeImageListWithBaselineNameRequest) SetCriteriaType(v string) *DescribeImageListWithBaselineNameRequest {
	s.CriteriaType = &v
	return s
}

func (s *DescribeImageListWithBaselineNameRequest) SetCurrentPage(v int32) *DescribeImageListWithBaselineNameRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageListWithBaselineNameRequest) SetImage(v string) *DescribeImageListWithBaselineNameRequest {
	s.Image = &v
	return s
}

func (s *DescribeImageListWithBaselineNameRequest) SetImageDigest(v string) *DescribeImageListWithBaselineNameRequest {
	s.ImageDigest = &v
	return s
}

func (s *DescribeImageListWithBaselineNameRequest) SetLang(v string) *DescribeImageListWithBaselineNameRequest {
	s.Lang = &v
	return s
}

func (s *DescribeImageListWithBaselineNameRequest) SetNamespace(v string) *DescribeImageListWithBaselineNameRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeImageListWithBaselineNameRequest) SetPageSize(v int32) *DescribeImageListWithBaselineNameRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeImageListWithBaselineNameRequest) SetPod(v string) *DescribeImageListWithBaselineNameRequest {
	s.Pod = &v
	return s
}

func (s *DescribeImageListWithBaselineNameRequest) SetRepoInstanceId(v string) *DescribeImageListWithBaselineNameRequest {
	s.RepoInstanceId = &v
	return s
}

func (s *DescribeImageListWithBaselineNameRequest) SetRepoName(v string) *DescribeImageListWithBaselineNameRequest {
	s.RepoName = &v
	return s
}

func (s *DescribeImageListWithBaselineNameRequest) SetRepoNamespace(v string) *DescribeImageListWithBaselineNameRequest {
	s.RepoNamespace = &v
	return s
}

func (s *DescribeImageListWithBaselineNameRequest) SetScanRange(v []*string) *DescribeImageListWithBaselineNameRequest {
	s.ScanRange = v
	return s
}

func (s *DescribeImageListWithBaselineNameRequest) Validate() error {
	return dara.Validate(s)
}
