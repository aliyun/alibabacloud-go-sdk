// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageRiskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImageRiskList(v []*ListImageRiskResponseBodyImageRiskList) *ListImageRiskResponseBody
	GetImageRiskList() []*ListImageRiskResponseBodyImageRiskList
	SetPageInfo(v *ListImageRiskResponseBodyPageInfo) *ListImageRiskResponseBody
	GetPageInfo() *ListImageRiskResponseBodyPageInfo
	SetRequestId(v string) *ListImageRiskResponseBody
	GetRequestId() *string
}

type ListImageRiskResponseBody struct {
	// An array that consists of security information about the image.
	ImageRiskList []*ListImageRiskResponseBodyImageRiskList `json:"ImageRiskList,omitempty" xml:"ImageRiskList,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *ListImageRiskResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 050ACC7A-D4FD-55C6-B861-BA9569C1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListImageRiskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListImageRiskResponseBody) GoString() string {
	return s.String()
}

func (s *ListImageRiskResponseBody) GetImageRiskList() []*ListImageRiskResponseBodyImageRiskList {
	return s.ImageRiskList
}

func (s *ListImageRiskResponseBody) GetPageInfo() *ListImageRiskResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListImageRiskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListImageRiskResponseBody) SetImageRiskList(v []*ListImageRiskResponseBodyImageRiskList) *ListImageRiskResponseBody {
	s.ImageRiskList = v
	return s
}

func (s *ListImageRiskResponseBody) SetPageInfo(v *ListImageRiskResponseBodyPageInfo) *ListImageRiskResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListImageRiskResponseBody) SetRequestId(v string) *ListImageRiskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListImageRiskResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListImageRiskResponseBodyImageRiskList struct {
	// The digest value of the image.
	//
	// example:
	//
	// 3f8efc2184cf1d24936b49c27286a284714b77be34c80c9ee38ca6bf322445****
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// An array that consists of the details of the endpoint.
	EndPointList []*ListImageRiskResponseBodyImageRiskListEndPointList `json:"EndPointList,omitempty" xml:"EndPointList,omitempty" type:"Repeated"`
	// The endpoint of Container Registry.
	//
	// example:
	//
	// https://172.20.XXX.XXX/test
	Endpoints *string `json:"Endpoints,omitempty" xml:"Endpoints,omitempty"`
	// The image of the container.
	//
	// example:
	//
	// .aliyuncs.com/sas_test/baseline:exploit
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The registration status of the image repository. Valid values:
	//
	// 	- **IN_SAS**: The image repository is registered with Security Center.
	//
	// 	- **NOT_IN_SAS**: The image repository is not registered with Security Center.
	//
	// example:
	//
	// IN_SAS
	ImageAccessType *string `json:"ImageAccessType,omitempty" xml:"ImageAccessType,omitempty"`
	// The ID of the image.
	//
	// example:
	//
	// f922bfeb6960436fe3f0e7b62fc6b9a0b47980986669c367c22433269404****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The public endpoint of the image repository.
	//
	// example:
	//
	// ****registry-registry.cn-shenzhen-finance-1.cr.aliyuncs.com/xxxx/docker-****
	InternetURLs *string `json:"InternetURLs,omitempty" xml:"InternetURLs,omitempty"`
	// The region of the image repository.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the image repository. Valid values:
	//
	// 	- **acr**
	//
	// 	- **harbor**
	//
	// 	- **quay**
	//
	// 	- **CI/CD**: Jenkins
	//
	// example:
	//
	// harbor
	RegistryType *string `json:"RegistryType,omitempty" xml:"RegistryType,omitempty"`
	// The ID of the image repository.
	//
	// example:
	//
	// crr-bk2l746eyxca1****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The name of the image repository.
	//
	// example:
	//
	// scan_test
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace to which the repository belongs.
	//
	// example:
	//
	// vultar***
	RepoNamespace *string `json:"RepoNamespace,omitempty" xml:"RepoNamespace,omitempty"`
	// The type of the repository. Valid values:
	//
	// 	- `PUBLIC`
	//
	// 	- `PRIVATE`
	//
	// example:
	//
	// PRIVATE
	RepoType *string `json:"RepoType,omitempty" xml:"RepoType,omitempty"`
	// The statistics on a security event.
	//
	// example:
	//
	// {
	//
	// 	"baselineNum": 0,
	//
	// 	"newSuspicious": 0,
	//
	// 	"vul": 0
	//
	// }
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The tag that is added to the image.
	//
	// example:
	//
	// 0.1.0
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// Indicates whether the image version is immutable. If the image version is immutable, only the image of the latest version in the image repository can be overwritten. Valid values:
	//
	// 	- **0**: The image version is mutable.
	//
	// 	- **1**: The image version is immutable.
	//
	// example:
	//
	// 0
	TagImmutable *int32 `json:"TagImmutable,omitempty" xml:"TagImmutable,omitempty"`
	// The UUID of the image.
	//
	// example:
	//
	// 624778f3-5bf2-423c-ac0c-47a62c05****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// The endpoint of the image repository in the VPC.
	//
	// example:
	//
	// ****-registry-registry-vpc.cn-shenzhen-finance-1.cr.aliyuncs.com/xxxx/docker-****
	VpcURLs *string `json:"VpcURLs,omitempty" xml:"VpcURLs,omitempty"`
}

func (s ListImageRiskResponseBodyImageRiskList) String() string {
	return dara.Prettify(s)
}

func (s ListImageRiskResponseBodyImageRiskList) GoString() string {
	return s.String()
}

func (s *ListImageRiskResponseBodyImageRiskList) GetDigest() *string {
	return s.Digest
}

func (s *ListImageRiskResponseBodyImageRiskList) GetEndPointList() []*ListImageRiskResponseBodyImageRiskListEndPointList {
	return s.EndPointList
}

func (s *ListImageRiskResponseBodyImageRiskList) GetEndpoints() *string {
	return s.Endpoints
}

func (s *ListImageRiskResponseBodyImageRiskList) GetImage() *string {
	return s.Image
}

func (s *ListImageRiskResponseBodyImageRiskList) GetImageAccessType() *string {
	return s.ImageAccessType
}

func (s *ListImageRiskResponseBodyImageRiskList) GetImageId() *string {
	return s.ImageId
}

func (s *ListImageRiskResponseBodyImageRiskList) GetInternetURLs() *string {
	return s.InternetURLs
}

func (s *ListImageRiskResponseBodyImageRiskList) GetRegionId() *string {
	return s.RegionId
}

func (s *ListImageRiskResponseBodyImageRiskList) GetRegistryType() *string {
	return s.RegistryType
}

func (s *ListImageRiskResponseBodyImageRiskList) GetRepoId() *string {
	return s.RepoId
}

func (s *ListImageRiskResponseBodyImageRiskList) GetRepoName() *string {
	return s.RepoName
}

func (s *ListImageRiskResponseBodyImageRiskList) GetRepoNamespace() *string {
	return s.RepoNamespace
}

func (s *ListImageRiskResponseBodyImageRiskList) GetRepoType() *string {
	return s.RepoType
}

func (s *ListImageRiskResponseBodyImageRiskList) GetStatistics() *string {
	return s.Statistics
}

func (s *ListImageRiskResponseBodyImageRiskList) GetTag() *string {
	return s.Tag
}

func (s *ListImageRiskResponseBodyImageRiskList) GetTagImmutable() *int32 {
	return s.TagImmutable
}

func (s *ListImageRiskResponseBodyImageRiskList) GetUuid() *string {
	return s.Uuid
}

func (s *ListImageRiskResponseBodyImageRiskList) GetVpcURLs() *string {
	return s.VpcURLs
}

func (s *ListImageRiskResponseBodyImageRiskList) SetDigest(v string) *ListImageRiskResponseBodyImageRiskList {
	s.Digest = &v
	return s
}

func (s *ListImageRiskResponseBodyImageRiskList) SetEndPointList(v []*ListImageRiskResponseBodyImageRiskListEndPointList) *ListImageRiskResponseBodyImageRiskList {
	s.EndPointList = v
	return s
}

func (s *ListImageRiskResponseBodyImageRiskList) SetEndpoints(v string) *ListImageRiskResponseBodyImageRiskList {
	s.Endpoints = &v
	return s
}

func (s *ListImageRiskResponseBodyImageRiskList) SetImage(v string) *ListImageRiskResponseBodyImageRiskList {
	s.Image = &v
	return s
}

func (s *ListImageRiskResponseBodyImageRiskList) SetImageAccessType(v string) *ListImageRiskResponseBodyImageRiskList {
	s.ImageAccessType = &v
	return s
}

func (s *ListImageRiskResponseBodyImageRiskList) SetImageId(v string) *ListImageRiskResponseBodyImageRiskList {
	s.ImageId = &v
	return s
}

func (s *ListImageRiskResponseBodyImageRiskList) SetInternetURLs(v string) *ListImageRiskResponseBodyImageRiskList {
	s.InternetURLs = &v
	return s
}

func (s *ListImageRiskResponseBodyImageRiskList) SetRegionId(v string) *ListImageRiskResponseBodyImageRiskList {
	s.RegionId = &v
	return s
}

func (s *ListImageRiskResponseBodyImageRiskList) SetRegistryType(v string) *ListImageRiskResponseBodyImageRiskList {
	s.RegistryType = &v
	return s
}

func (s *ListImageRiskResponseBodyImageRiskList) SetRepoId(v string) *ListImageRiskResponseBodyImageRiskList {
	s.RepoId = &v
	return s
}

func (s *ListImageRiskResponseBodyImageRiskList) SetRepoName(v string) *ListImageRiskResponseBodyImageRiskList {
	s.RepoName = &v
	return s
}

func (s *ListImageRiskResponseBodyImageRiskList) SetRepoNamespace(v string) *ListImageRiskResponseBodyImageRiskList {
	s.RepoNamespace = &v
	return s
}

func (s *ListImageRiskResponseBodyImageRiskList) SetRepoType(v string) *ListImageRiskResponseBodyImageRiskList {
	s.RepoType = &v
	return s
}

func (s *ListImageRiskResponseBodyImageRiskList) SetStatistics(v string) *ListImageRiskResponseBodyImageRiskList {
	s.Statistics = &v
	return s
}

func (s *ListImageRiskResponseBodyImageRiskList) SetTag(v string) *ListImageRiskResponseBodyImageRiskList {
	s.Tag = &v
	return s
}

func (s *ListImageRiskResponseBodyImageRiskList) SetTagImmutable(v int32) *ListImageRiskResponseBodyImageRiskList {
	s.TagImmutable = &v
	return s
}

func (s *ListImageRiskResponseBodyImageRiskList) SetUuid(v string) *ListImageRiskResponseBodyImageRiskList {
	s.Uuid = &v
	return s
}

func (s *ListImageRiskResponseBodyImageRiskList) SetVpcURLs(v string) *ListImageRiskResponseBodyImageRiskList {
	s.VpcURLs = &v
	return s
}

func (s *ListImageRiskResponseBodyImageRiskList) Validate() error {
	return dara.Validate(s)
}

type ListImageRiskResponseBodyImageRiskListEndPointList struct {
	// An array that consists the details of the domain name in the endpoint.
	Domains []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// The type of the domain name in the endpoint. Valid values:
	//
	// 	- **internet**: Internet
	//
	// 	- **intranet**: internal network
	//
	// example:
	//
	// internet
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListImageRiskResponseBodyImageRiskListEndPointList) String() string {
	return dara.Prettify(s)
}

func (s ListImageRiskResponseBodyImageRiskListEndPointList) GoString() string {
	return s.String()
}

func (s *ListImageRiskResponseBodyImageRiskListEndPointList) GetDomains() []*string {
	return s.Domains
}

func (s *ListImageRiskResponseBodyImageRiskListEndPointList) GetType() *string {
	return s.Type
}

func (s *ListImageRiskResponseBodyImageRiskListEndPointList) SetDomains(v []*string) *ListImageRiskResponseBodyImageRiskListEndPointList {
	s.Domains = v
	return s
}

func (s *ListImageRiskResponseBodyImageRiskListEndPointList) SetType(v string) *ListImageRiskResponseBodyImageRiskListEndPointList {
	s.Type = &v
	return s
}

func (s *ListImageRiskResponseBodyImageRiskListEndPointList) Validate() error {
	return dara.Validate(s)
}

type ListImageRiskResponseBodyPageInfo struct {
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
	// 23
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListImageRiskResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListImageRiskResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListImageRiskResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *ListImageRiskResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListImageRiskResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListImageRiskResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListImageRiskResponseBodyPageInfo) SetCount(v int32) *ListImageRiskResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListImageRiskResponseBodyPageInfo) SetCurrentPage(v int32) *ListImageRiskResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListImageRiskResponseBodyPageInfo) SetPageSize(v int32) *ListImageRiskResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListImageRiskResponseBodyPageInfo) SetTotalCount(v int32) *ListImageRiskResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListImageRiskResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
