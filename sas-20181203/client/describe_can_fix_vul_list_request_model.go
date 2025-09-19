// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCanFixVulListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliasName(v string) *DescribeCanFixVulListRequest
	GetAliasName() *string
	SetClusterId(v string) *DescribeCanFixVulListRequest
	GetClusterId() *string
	SetClusterName(v string) *DescribeCanFixVulListRequest
	GetClusterName() *string
	SetContainerId(v string) *DescribeCanFixVulListRequest
	GetContainerId() *string
	SetCurrentPage(v int32) *DescribeCanFixVulListRequest
	GetCurrentPage() *int32
	SetDealed(v string) *DescribeCanFixVulListRequest
	GetDealed() *string
	SetDigest(v string) *DescribeCanFixVulListRequest
	GetDigest() *string
	SetImage(v string) *DescribeCanFixVulListRequest
	GetImage() *string
	SetInstanceId(v string) *DescribeCanFixVulListRequest
	GetInstanceId() *string
	SetName(v string) *DescribeCanFixVulListRequest
	GetName() *string
	SetNamespace(v string) *DescribeCanFixVulListRequest
	GetNamespace() *string
	SetNecessity(v string) *DescribeCanFixVulListRequest
	GetNecessity() *string
	SetPageSize(v int32) *DescribeCanFixVulListRequest
	GetPageSize() *int32
	SetPod(v string) *DescribeCanFixVulListRequest
	GetPod() *string
	SetRegionId(v string) *DescribeCanFixVulListRequest
	GetRegionId() *string
	SetRepoId(v string) *DescribeCanFixVulListRequest
	GetRepoId() *string
	SetRepoInstanceId(v string) *DescribeCanFixVulListRequest
	GetRepoInstanceId() *string
	SetRepoName(v string) *DescribeCanFixVulListRequest
	GetRepoName() *string
	SetRepoNamespace(v string) *DescribeCanFixVulListRequest
	GetRepoNamespace() *string
	SetRepoRegionId(v string) *DescribeCanFixVulListRequest
	GetRepoRegionId() *string
	SetScanRange(v []*string) *DescribeCanFixVulListRequest
	GetScanRange() []*string
	SetStatusList(v string) *DescribeCanFixVulListRequest
	GetStatusList() *string
	SetTag(v string) *DescribeCanFixVulListRequest
	GetTag() *string
	SetType(v string) *DescribeCanFixVulListRequest
	GetType() *string
	SetUuids(v string) *DescribeCanFixVulListRequest
	GetUuids() *string
}

type DescribeCanFixVulListRequest struct {
	// The alias of the vulnerability that is specified in Common Vulnerabilities and Exposures (CVE).
	//
	// example:
	//
	// RHSA-2017:0184-Important: mysql security update
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The cluster ID.
	//
	// >  You can call the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) operation to query the IDs of clusters.
	//
	// example:
	//
	// c80f79959fd724a888e1187779b13****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// sas-test-cnnf
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The container ID.
	//
	// example:
	//
	// 48a6d9a92435a13ad573372c3f3c63b7e04d106458141df9f92155709d5a****
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Specifies whether the vulnerability is handled. Valid values:
	//
	// **y**: The vulnerability is handled. **n**: The vulnerability is not handled.
	//
	// example:
	//
	// n
	Dealed *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	// The unique identifier of the image.
	//
	// example:
	//
	// 8f0fbdb41d3d1ade4ffdf21558443f4c03342010563bb8c43ccc09594d50****
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The name of the image.
	//
	// example:
	//
	// registry.cn-wulanchabu.aliyuncs.com/sas_test/huxin-test-001:nuxeo6-****
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The ID of the container image.
	//
	// >  You can call the [ListRepository](https://help.aliyun.com/document_detail/451339.html) operation of Container Registry and obtain the ID of the container image from **InstanceId*	- in the response.
	//
	// example:
	//
	// cri-rv4nvbv8iju4****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the vulnerability.
	//
	// example:
	//
	// scan:AVD-2022-953356
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace of the cluster.
	//
	// >  You can call the [GetOpaClusterNamespaceList](~~GetOpaClusterNamespaceList~~) operation to query the namespaces of clusters.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The priority to fix the vulnerability. Separate multiple priorities with commas (,). Valid values:
	//
	// 	- **asap**: high
	//
	// 	- **later**: medium
	//
	// 	- **nntf**: low
	//
	// example:
	//
	// asap,later,nntf
	Necessity *string `json:"Necessity,omitempty" xml:"Necessity,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the container group.
	//
	// example:
	//
	// 22222-7xsqq
	Pod *string `json:"Pod,omitempty" xml:"Pod,omitempty"`
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
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the image repository.
	//
	// >  You can call the [ListRepository](https://help.aliyun.com/document_detail/145293.html) operation of Container Registry and obtain the ID of the image repository from **RepoId*	- in the response.
	//
	// example:
	//
	// crr-avo7qp02simz2njo
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The ID of the container image.
	//
	// >  You can call the [ListRepository](https://help.aliyun.com/document_detail/451339.html) operation of Container Registry and obtain the ID of the container image from **InstanceId*	- in the response.
	//
	// example:
	//
	// cri-rv4nvbv8iju4****
	RepoInstanceId *string `json:"RepoInstanceId,omitempty" xml:"RepoInstanceId,omitempty"`
	// The name of the image repository.
	//
	// >  Fuzzy match is supported.
	//
	// example:
	//
	// digital-account
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The namespace to which the image repository belongs.
	//
	// >  Fuzzy match is supported.
	//
	// example:
	//
	// ns-digital-dev
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
	// The type of the asset that you want to scan. Valid values:
	//
	// 	- **image**
	//
	// 	- **container**
	ScanRange []*string `json:"ScanRange,omitempty" xml:"ScanRange,omitempty" type:"Repeated"`
	// The status of the vulnerability. Valid values:
	//
	// 	- **1**: The vulnerability is unfixed.
	//
	// 	- **4**: The vulnerability is being fixed.
	//
	// 	- **7**:The vulnerability is fixed.
	//
	// example:
	//
	// 1
	StatusList *string `json:"StatusList,omitempty" xml:"StatusList,omitempty"`
	// The tag to add to the image.
	//
	// example:
	//
	// 0.1.0
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The type of the vulnerability. Valid values:
	//
	// 	- **cve**: system vulnerability
	//
	// 	- **sca**: application vulnerability
	//
	// This parameter is required.
	//
	// example:
	//
	// cve
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The UUID of the image. Separate multiple UUIDs with commas (,).
	//
	// example:
	//
	// d15df12472809c1c3b158606c0f1****
	Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s DescribeCanFixVulListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCanFixVulListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCanFixVulListRequest) GetAliasName() *string {
	return s.AliasName
}

func (s *DescribeCanFixVulListRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeCanFixVulListRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeCanFixVulListRequest) GetContainerId() *string {
	return s.ContainerId
}

func (s *DescribeCanFixVulListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeCanFixVulListRequest) GetDealed() *string {
	return s.Dealed
}

func (s *DescribeCanFixVulListRequest) GetDigest() *string {
	return s.Digest
}

func (s *DescribeCanFixVulListRequest) GetImage() *string {
	return s.Image
}

func (s *DescribeCanFixVulListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCanFixVulListRequest) GetName() *string {
	return s.Name
}

func (s *DescribeCanFixVulListRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeCanFixVulListRequest) GetNecessity() *string {
	return s.Necessity
}

func (s *DescribeCanFixVulListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCanFixVulListRequest) GetPod() *string {
	return s.Pod
}

func (s *DescribeCanFixVulListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCanFixVulListRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *DescribeCanFixVulListRequest) GetRepoInstanceId() *string {
	return s.RepoInstanceId
}

func (s *DescribeCanFixVulListRequest) GetRepoName() *string {
	return s.RepoName
}

func (s *DescribeCanFixVulListRequest) GetRepoNamespace() *string {
	return s.RepoNamespace
}

func (s *DescribeCanFixVulListRequest) GetRepoRegionId() *string {
	return s.RepoRegionId
}

func (s *DescribeCanFixVulListRequest) GetScanRange() []*string {
	return s.ScanRange
}

func (s *DescribeCanFixVulListRequest) GetStatusList() *string {
	return s.StatusList
}

func (s *DescribeCanFixVulListRequest) GetTag() *string {
	return s.Tag
}

func (s *DescribeCanFixVulListRequest) GetType() *string {
	return s.Type
}

func (s *DescribeCanFixVulListRequest) GetUuids() *string {
	return s.Uuids
}

func (s *DescribeCanFixVulListRequest) SetAliasName(v string) *DescribeCanFixVulListRequest {
	s.AliasName = &v
	return s
}

func (s *DescribeCanFixVulListRequest) SetClusterId(v string) *DescribeCanFixVulListRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeCanFixVulListRequest) SetClusterName(v string) *DescribeCanFixVulListRequest {
	s.ClusterName = &v
	return s
}

func (s *DescribeCanFixVulListRequest) SetContainerId(v string) *DescribeCanFixVulListRequest {
	s.ContainerId = &v
	return s
}

func (s *DescribeCanFixVulListRequest) SetCurrentPage(v int32) *DescribeCanFixVulListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCanFixVulListRequest) SetDealed(v string) *DescribeCanFixVulListRequest {
	s.Dealed = &v
	return s
}

func (s *DescribeCanFixVulListRequest) SetDigest(v string) *DescribeCanFixVulListRequest {
	s.Digest = &v
	return s
}

func (s *DescribeCanFixVulListRequest) SetImage(v string) *DescribeCanFixVulListRequest {
	s.Image = &v
	return s
}

func (s *DescribeCanFixVulListRequest) SetInstanceId(v string) *DescribeCanFixVulListRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeCanFixVulListRequest) SetName(v string) *DescribeCanFixVulListRequest {
	s.Name = &v
	return s
}

func (s *DescribeCanFixVulListRequest) SetNamespace(v string) *DescribeCanFixVulListRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeCanFixVulListRequest) SetNecessity(v string) *DescribeCanFixVulListRequest {
	s.Necessity = &v
	return s
}

func (s *DescribeCanFixVulListRequest) SetPageSize(v int32) *DescribeCanFixVulListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCanFixVulListRequest) SetPod(v string) *DescribeCanFixVulListRequest {
	s.Pod = &v
	return s
}

func (s *DescribeCanFixVulListRequest) SetRegionId(v string) *DescribeCanFixVulListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCanFixVulListRequest) SetRepoId(v string) *DescribeCanFixVulListRequest {
	s.RepoId = &v
	return s
}

func (s *DescribeCanFixVulListRequest) SetRepoInstanceId(v string) *DescribeCanFixVulListRequest {
	s.RepoInstanceId = &v
	return s
}

func (s *DescribeCanFixVulListRequest) SetRepoName(v string) *DescribeCanFixVulListRequest {
	s.RepoName = &v
	return s
}

func (s *DescribeCanFixVulListRequest) SetRepoNamespace(v string) *DescribeCanFixVulListRequest {
	s.RepoNamespace = &v
	return s
}

func (s *DescribeCanFixVulListRequest) SetRepoRegionId(v string) *DescribeCanFixVulListRequest {
	s.RepoRegionId = &v
	return s
}

func (s *DescribeCanFixVulListRequest) SetScanRange(v []*string) *DescribeCanFixVulListRequest {
	s.ScanRange = v
	return s
}

func (s *DescribeCanFixVulListRequest) SetStatusList(v string) *DescribeCanFixVulListRequest {
	s.StatusList = &v
	return s
}

func (s *DescribeCanFixVulListRequest) SetTag(v string) *DescribeCanFixVulListRequest {
	s.Tag = &v
	return s
}

func (s *DescribeCanFixVulListRequest) SetType(v string) *DescribeCanFixVulListRequest {
	s.Type = &v
	return s
}

func (s *DescribeCanFixVulListRequest) SetUuids(v string) *DescribeCanFixVulListRequest {
	s.Uuids = &v
	return s
}

func (s *DescribeCanFixVulListRequest) Validate() error {
	return dara.Validate(s)
}
