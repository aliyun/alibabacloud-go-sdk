// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageVulListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliasName(v string) *DescribeImageVulListRequest
	GetAliasName() *string
	SetClusterId(v string) *DescribeImageVulListRequest
	GetClusterId() *string
	SetClusterName(v string) *DescribeImageVulListRequest
	GetClusterName() *string
	SetContainerId(v string) *DescribeImageVulListRequest
	GetContainerId() *string
	SetCurrentPage(v int32) *DescribeImageVulListRequest
	GetCurrentPage() *int32
	SetDealed(v string) *DescribeImageVulListRequest
	GetDealed() *string
	SetDigest(v string) *DescribeImageVulListRequest
	GetDigest() *string
	SetImage(v string) *DescribeImageVulListRequest
	GetImage() *string
	SetInstanceId(v string) *DescribeImageVulListRequest
	GetInstanceId() *string
	SetLang(v string) *DescribeImageVulListRequest
	GetLang() *string
	SetName(v string) *DescribeImageVulListRequest
	GetName() *string
	SetNamespace(v string) *DescribeImageVulListRequest
	GetNamespace() *string
	SetNecessity(v string) *DescribeImageVulListRequest
	GetNecessity() *string
	SetPageSize(v int32) *DescribeImageVulListRequest
	GetPageSize() *int32
	SetPod(v string) *DescribeImageVulListRequest
	GetPod() *string
	SetRegionId(v string) *DescribeImageVulListRequest
	GetRegionId() *string
	SetRepoId(v string) *DescribeImageVulListRequest
	GetRepoId() *string
	SetRepoInstanceId(v string) *DescribeImageVulListRequest
	GetRepoInstanceId() *string
	SetRepoName(v string) *DescribeImageVulListRequest
	GetRepoName() *string
	SetRepoNamespace(v string) *DescribeImageVulListRequest
	GetRepoNamespace() *string
	SetRepoRegionId(v string) *DescribeImageVulListRequest
	GetRepoRegionId() *string
	SetRuleTag(v string) *DescribeImageVulListRequest
	GetRuleTag() *string
	SetScanRange(v []*string) *DescribeImageVulListRequest
	GetScanRange() []*string
	SetStatusList(v string) *DescribeImageVulListRequest
	GetStatusList() *string
	SetTag(v string) *DescribeImageVulListRequest
	GetTag() *string
	SetType(v string) *DescribeImageVulListRequest
	GetType() *string
	SetUuids(v string) *DescribeImageVulListRequest
	GetUuids() *string
}

type DescribeImageVulListRequest struct {
	// The alias of the vulnerability.
	//
	// example:
	//
	// High severity vulnerability that affects org.eclipse.jetty:jetty-server
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The ID of the cluster to which the container belongs.
	//
	// example:
	//
	// cc20a1024011c44b6a8710d6f8b****
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
	// c08d5fc1a329a4b88950a253d082f****
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Specifies whether the vulnerability is handled. Valid values:
	//
	// 	- **y**: yes
	//
	// 	- **n**: no
	//
	// example:
	//
	// y
	Dealed *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	// The digest of the image.
	//
	// example:
	//
	// 8f0fbdb41d3d1ade4ffdf21558443f4c03342010563bb8c43ccc09594d507012
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The name of the image.
	//
	// example:
	//
	// registry.cn-wulanchabu.aliyuncs.com/sas_test/huxin-test-001:nuxeo6-****
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The instance ID of the asset.
	//
	// example:
	//
	// 1-qeqewqw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	// The name of the vulnerability.
	//
	// example:
	//
	// debian:10:CVE-2019-9893
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace.
	//
	// example:
	//
	// test-002
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The priority to fix the vulnerability. Valid values:
	//
	// 	- **asap**: high. You must fix the vulnerability at the earliest opportunity.
	//
	// 	- **later**: medium. You can fix the vulnerability based on your business requirements.
	//
	// 	- **nntf**: low. You can ignore the vulnerability.
	//
	// example:
	//
	// asap
	Necessity *string `json:"Necessity,omitempty" xml:"Necessity,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
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
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the image repository.
	//
	// example:
	//
	// qew****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The instance ID of the image repository.
	//
	// example:
	//
	// i-qewqrqcsadf****
	RepoInstanceId *string `json:"RepoInstanceId,omitempty" xml:"RepoInstanceId,omitempty"`
	// The name of the image.
	//
	// example:
	//
	// libssh2
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The namespace to which the image repository belongs.
	//
	// example:
	//
	// libssh2
	RepoNamespace *string `json:"RepoNamespace,omitempty" xml:"RepoNamespace,omitempty"`
	// The region ID of the image repository.
	//
	// example:
	//
	// cn-hangzhou
	RepoRegionId *string `json:"RepoRegionId,omitempty" xml:"RepoRegionId,omitempty"`
	// The tag of this vulnerability. Valid values:
	//
	// 	- **AI**: AI-related components.
	//
	// example:
	//
	// AI
	RuleTag *string `json:"RuleTag,omitempty" xml:"RuleTag,omitempty"`
	// The types of the assets that you want to scan.
	ScanRange []*string `json:"ScanRange,omitempty" xml:"ScanRange,omitempty" type:"Repeated"`
	// The status of the vulnerability. Valid values:
	//
	// 	- **1**: unfixed
	//
	// 	- **4**: being fixed
	//
	// 	- **7**: fixed
	//
	// example:
	//
	// 1
	StatusList *string `json:"StatusList,omitempty" xml:"StatusList,omitempty"`
	// The tag that is added to the image vulnerability.
	//
	// example:
	//
	// oval
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The type of the vulnerability. Set the value to **cve**, which indicates image vulnerabilities.
	//
	// This parameter is required.
	//
	// example:
	//
	// cve
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The UUIDs of the assets. Separate multiple UUIDs with commas (,).
	//
	// example:
	//
	// 0004a32a0305a7f6ab5ff9600d47****
	Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s DescribeImageVulListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageVulListRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageVulListRequest) GetAliasName() *string {
	return s.AliasName
}

func (s *DescribeImageVulListRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeImageVulListRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeImageVulListRequest) GetContainerId() *string {
	return s.ContainerId
}

func (s *DescribeImageVulListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeImageVulListRequest) GetDealed() *string {
	return s.Dealed
}

func (s *DescribeImageVulListRequest) GetDigest() *string {
	return s.Digest
}

func (s *DescribeImageVulListRequest) GetImage() *string {
	return s.Image
}

func (s *DescribeImageVulListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeImageVulListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeImageVulListRequest) GetName() *string {
	return s.Name
}

func (s *DescribeImageVulListRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeImageVulListRequest) GetNecessity() *string {
	return s.Necessity
}

func (s *DescribeImageVulListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageVulListRequest) GetPod() *string {
	return s.Pod
}

func (s *DescribeImageVulListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeImageVulListRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *DescribeImageVulListRequest) GetRepoInstanceId() *string {
	return s.RepoInstanceId
}

func (s *DescribeImageVulListRequest) GetRepoName() *string {
	return s.RepoName
}

func (s *DescribeImageVulListRequest) GetRepoNamespace() *string {
	return s.RepoNamespace
}

func (s *DescribeImageVulListRequest) GetRepoRegionId() *string {
	return s.RepoRegionId
}

func (s *DescribeImageVulListRequest) GetRuleTag() *string {
	return s.RuleTag
}

func (s *DescribeImageVulListRequest) GetScanRange() []*string {
	return s.ScanRange
}

func (s *DescribeImageVulListRequest) GetStatusList() *string {
	return s.StatusList
}

func (s *DescribeImageVulListRequest) GetTag() *string {
	return s.Tag
}

func (s *DescribeImageVulListRequest) GetType() *string {
	return s.Type
}

func (s *DescribeImageVulListRequest) GetUuids() *string {
	return s.Uuids
}

func (s *DescribeImageVulListRequest) SetAliasName(v string) *DescribeImageVulListRequest {
	s.AliasName = &v
	return s
}

func (s *DescribeImageVulListRequest) SetClusterId(v string) *DescribeImageVulListRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeImageVulListRequest) SetClusterName(v string) *DescribeImageVulListRequest {
	s.ClusterName = &v
	return s
}

func (s *DescribeImageVulListRequest) SetContainerId(v string) *DescribeImageVulListRequest {
	s.ContainerId = &v
	return s
}

func (s *DescribeImageVulListRequest) SetCurrentPage(v int32) *DescribeImageVulListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageVulListRequest) SetDealed(v string) *DescribeImageVulListRequest {
	s.Dealed = &v
	return s
}

func (s *DescribeImageVulListRequest) SetDigest(v string) *DescribeImageVulListRequest {
	s.Digest = &v
	return s
}

func (s *DescribeImageVulListRequest) SetImage(v string) *DescribeImageVulListRequest {
	s.Image = &v
	return s
}

func (s *DescribeImageVulListRequest) SetInstanceId(v string) *DescribeImageVulListRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeImageVulListRequest) SetLang(v string) *DescribeImageVulListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeImageVulListRequest) SetName(v string) *DescribeImageVulListRequest {
	s.Name = &v
	return s
}

func (s *DescribeImageVulListRequest) SetNamespace(v string) *DescribeImageVulListRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeImageVulListRequest) SetNecessity(v string) *DescribeImageVulListRequest {
	s.Necessity = &v
	return s
}

func (s *DescribeImageVulListRequest) SetPageSize(v int32) *DescribeImageVulListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeImageVulListRequest) SetPod(v string) *DescribeImageVulListRequest {
	s.Pod = &v
	return s
}

func (s *DescribeImageVulListRequest) SetRegionId(v string) *DescribeImageVulListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeImageVulListRequest) SetRepoId(v string) *DescribeImageVulListRequest {
	s.RepoId = &v
	return s
}

func (s *DescribeImageVulListRequest) SetRepoInstanceId(v string) *DescribeImageVulListRequest {
	s.RepoInstanceId = &v
	return s
}

func (s *DescribeImageVulListRequest) SetRepoName(v string) *DescribeImageVulListRequest {
	s.RepoName = &v
	return s
}

func (s *DescribeImageVulListRequest) SetRepoNamespace(v string) *DescribeImageVulListRequest {
	s.RepoNamespace = &v
	return s
}

func (s *DescribeImageVulListRequest) SetRepoRegionId(v string) *DescribeImageVulListRequest {
	s.RepoRegionId = &v
	return s
}

func (s *DescribeImageVulListRequest) SetRuleTag(v string) *DescribeImageVulListRequest {
	s.RuleTag = &v
	return s
}

func (s *DescribeImageVulListRequest) SetScanRange(v []*string) *DescribeImageVulListRequest {
	s.ScanRange = v
	return s
}

func (s *DescribeImageVulListRequest) SetStatusList(v string) *DescribeImageVulListRequest {
	s.StatusList = &v
	return s
}

func (s *DescribeImageVulListRequest) SetTag(v string) *DescribeImageVulListRequest {
	s.Tag = &v
	return s
}

func (s *DescribeImageVulListRequest) SetType(v string) *DescribeImageVulListRequest {
	s.Type = &v
	return s
}

func (s *DescribeImageVulListRequest) SetUuids(v string) *DescribeImageVulListRequest {
	s.Uuids = &v
	return s
}

func (s *DescribeImageVulListRequest) Validate() error {
	return dara.Validate(s)
}
