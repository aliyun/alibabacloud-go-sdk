// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupedVulRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliasName(v string) *DescribeGroupedVulRequest
	GetAliasName() *string
	SetAssetType(v string) *DescribeGroupedVulRequest
	GetAssetType() *string
	SetAttachTypes(v string) *DescribeGroupedVulRequest
	GetAttachTypes() *string
	SetClusterId(v string) *DescribeGroupedVulRequest
	GetClusterId() *string
	SetContainerFieldName(v string) *DescribeGroupedVulRequest
	GetContainerFieldName() *string
	SetContainerFieldValue(v string) *DescribeGroupedVulRequest
	GetContainerFieldValue() *string
	SetCurrentPage(v int32) *DescribeGroupedVulRequest
	GetCurrentPage() *int32
	SetCveId(v string) *DescribeGroupedVulRequest
	GetCveId() *string
	SetDealed(v string) *DescribeGroupedVulRequest
	GetDealed() *string
	SetGroupId(v string) *DescribeGroupedVulRequest
	GetGroupId() *string
	SetLang(v string) *DescribeGroupedVulRequest
	GetLang() *string
	SetNecessity(v string) *DescribeGroupedVulRequest
	GetNecessity() *string
	SetPageSize(v int32) *DescribeGroupedVulRequest
	GetPageSize() *int32
	SetRaspDefend(v int32) *DescribeGroupedVulRequest
	GetRaspDefend() *int32
	SetResourceDirectoryAccountId(v int64) *DescribeGroupedVulRequest
	GetResourceDirectoryAccountId() *int64
	SetSearchTags(v string) *DescribeGroupedVulRequest
	GetSearchTags() *string
	SetTargetType(v string) *DescribeGroupedVulRequest
	GetTargetType() *string
	SetType(v string) *DescribeGroupedVulRequest
	GetType() *string
	SetUuids(v string) *DescribeGroupedVulRequest
	GetUuids() *string
}

type DescribeGroupedVulRequest struct {
	// The alias of the vulnerability.
	//
	// example:
	//
	// RHSA-2019:0230-Important: polkit security update
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The type of the asset on which the vulnerability is detected. Separate multiple types with commas (,). Valid values:
	//
	// 	- **ECS**: Elastic Compute Service (ECS) instance
	//
	// 	- **CONTAINER**: container
	//
	// example:
	//
	// ECS,CONTAINER
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// The type of the vulnerability. This parameter is valid only for application vulnerabilities. Separate multiple values with commas (,). Valid values:
	//
	// 	- **sca**: vulnerability that is detected based on software component analysis
	//
	// 	- **app**: application vulnerability
	//
	// example:
	//
	// sca
	AttachTypes *string `json:"AttachTypes,omitempty" xml:"AttachTypes,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// c88fb10da1168494091db6aafc5dd****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The key of the condition that is used to query containers. Valid values:
	//
	// 	- **instanceId**: the ID of the asset
	//
	// 	- **appName**: the name of the application
	//
	// 	- **clusterId**: the ID of the cluster
	//
	// 	- **regionId**: the ID of the region
	//
	// 	- **nodeName**: the name of the node
	//
	// 	- **namespace**: the namespace
	//
	// 	- **clusterName**: the name of the cluster
	//
	// 	- **image**: the name of the image
	//
	// 	- **imageRepoName**: the name of the image repository
	//
	// 	- **imageRepoNamespace**: the namespace to which the image repository belongs
	//
	// 	- **imageRepoTag**: the tag that is added to the image
	//
	// 	- **imageDigest**: the digest of the image
	//
	// example:
	//
	// appName
	ContainerFieldName *string `json:"ContainerFieldName,omitempty" xml:"ContainerFieldName,omitempty"`
	// The value specified by **ContainerFieldName**.
	//
	// example:
	//
	// cc914b0df156d40148412afe4a581****
	ContainerFieldValue *string `json:"ContainerFieldValue,omitempty" xml:"ContainerFieldValue,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The CVE ID.
	//
	// >  You can call the [DescribeVulListPage](~~DescribeVulListPage~~) operation to query the CVE ID.
	//
	// example:
	//
	// CVE-2017-15420
	CveId *string `json:"CveId,omitempty" xml:"CveId,omitempty"`
	// Specifies whether the vulnerability is handled. Valid values:
	//
	// 	- **y**: handled
	//
	// 	- **n**: not handled
	//
	// example:
	//
	// n
	Dealed *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	// The ID of the asset group.
	//
	// example:
	//
	// 235454
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
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
	// The priorities to fix the vulnerabilities. Separate multiple priorities with commas (,). Valid values:
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
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Indicates whether the application protection feature is supported. Valid values:
	//
	// - **0**: no.
	//
	// - **1**: yes.
	//
	// example:
	//
	// 1
	RaspDefend *int32 `json:"RaspDefend,omitempty" xml:"RaspDefend,omitempty"`
	// The Alibaba Cloud account ID of the member in the resource directory.
	//
	// >  You can call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to query the account ID.
	//
	// example:
	//
	// 127608589417****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The tag that is used to search for the vulnerability. Valid values:
	//
	// 	- **Restart required**
	//
	// 	- **Remote utilization**
	//
	// 	- **EXP exists**
	//
	// 	- **Available**
	//
	// 	- **Elevation of Privilege**
	//
	// 	- **Code Execution**
	//
	// example:
	//
	// Code Execution
	SearchTags *string `json:"SearchTags,omitempty" xml:"SearchTags,omitempty"`
	// The query type for containers. Valid values:
	//
	// 	- **containerId**: the ID of the container
	//
	// 	- **uuid**: the ID of the asset
	//
	// example:
	//
	// containerId
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The type of the vulnerability that you want to query. Default value: cve. Valid values:
	//
	// 	- **cve**: Linux software vulnerability
	//
	// 	- **sys**: Windows system vulnerability
	//
	// 	- **cms**: Web-CMS vulnerability
	//
	// 	- **app**: application vulnerability that is detected by network scanning
	//
	// 	- **sca**: application vulnerability that is detected by software component analysis
	//
	// example:
	//
	// cve
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The UUID of the server. Separate multiple UUIDs with commas (,).
	//
	// example:
	//
	// d42f938c-d962-48a0-90f9-05e4ea****
	Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s DescribeGroupedVulRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupedVulRequest) GoString() string {
	return s.String()
}

func (s *DescribeGroupedVulRequest) GetAliasName() *string {
	return s.AliasName
}

func (s *DescribeGroupedVulRequest) GetAssetType() *string {
	return s.AssetType
}

func (s *DescribeGroupedVulRequest) GetAttachTypes() *string {
	return s.AttachTypes
}

func (s *DescribeGroupedVulRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeGroupedVulRequest) GetContainerFieldName() *string {
	return s.ContainerFieldName
}

func (s *DescribeGroupedVulRequest) GetContainerFieldValue() *string {
	return s.ContainerFieldValue
}

func (s *DescribeGroupedVulRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeGroupedVulRequest) GetCveId() *string {
	return s.CveId
}

func (s *DescribeGroupedVulRequest) GetDealed() *string {
	return s.Dealed
}

func (s *DescribeGroupedVulRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeGroupedVulRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeGroupedVulRequest) GetNecessity() *string {
	return s.Necessity
}

func (s *DescribeGroupedVulRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGroupedVulRequest) GetRaspDefend() *int32 {
	return s.RaspDefend
}

func (s *DescribeGroupedVulRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *DescribeGroupedVulRequest) GetSearchTags() *string {
	return s.SearchTags
}

func (s *DescribeGroupedVulRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *DescribeGroupedVulRequest) GetType() *string {
	return s.Type
}

func (s *DescribeGroupedVulRequest) GetUuids() *string {
	return s.Uuids
}

func (s *DescribeGroupedVulRequest) SetAliasName(v string) *DescribeGroupedVulRequest {
	s.AliasName = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetAssetType(v string) *DescribeGroupedVulRequest {
	s.AssetType = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetAttachTypes(v string) *DescribeGroupedVulRequest {
	s.AttachTypes = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetClusterId(v string) *DescribeGroupedVulRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetContainerFieldName(v string) *DescribeGroupedVulRequest {
	s.ContainerFieldName = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetContainerFieldValue(v string) *DescribeGroupedVulRequest {
	s.ContainerFieldValue = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetCurrentPage(v int32) *DescribeGroupedVulRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetCveId(v string) *DescribeGroupedVulRequest {
	s.CveId = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetDealed(v string) *DescribeGroupedVulRequest {
	s.Dealed = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetGroupId(v string) *DescribeGroupedVulRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetLang(v string) *DescribeGroupedVulRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetNecessity(v string) *DescribeGroupedVulRequest {
	s.Necessity = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetPageSize(v int32) *DescribeGroupedVulRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetRaspDefend(v int32) *DescribeGroupedVulRequest {
	s.RaspDefend = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetResourceDirectoryAccountId(v int64) *DescribeGroupedVulRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetSearchTags(v string) *DescribeGroupedVulRequest {
	s.SearchTags = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetTargetType(v string) *DescribeGroupedVulRequest {
	s.TargetType = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetType(v string) *DescribeGroupedVulRequest {
	s.Type = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetUuids(v string) *DescribeGroupedVulRequest {
	s.Uuids = &v
	return s
}

func (s *DescribeGroupedVulRequest) Validate() error {
	return dara.Validate(s)
}
