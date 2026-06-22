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
	// The alias of the vulnerability to query.
	//
	// example:
	//
	// RHSA-2019:0230-Important: polkit security update
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The Asset Type where the vulnerability is detected. Separate multiple types with commas (,). Valid values:
	//
	// - **ECS**: host asset
	//
	// - **CONTAINER**: container asset.
	//
	// example:
	//
	// ECS,CONTAINER
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// The vulnerability type. This query condition is valid only for application vulnerabilities. Separate multiple values with commas (,). Valid values:
	//
	// - **sca**: software constituency parsing vulnerability
	//
	// - **app**: application vulnerability.
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
	// The container search field. Valid values:
	//
	// - **instanceId**: instance ID
	//
	// - **appName**: application name
	//
	// - **clusterId**: cluster ID
	//
	// - **regionId**: region
	//
	// - **nodeName**: node name
	//
	// - **namespace**: namespace
	//
	// - **clusterName**: cluster name
	//
	// - **image**: image name
	//
	// - **imageRepoName**: image repository name
	//
	// - **imageRepoNamespace**: image repository namespace
	//
	// - **imageRepoTag**: image tag
	//
	// - **imageDigest**: image digest.
	//
	// example:
	//
	// appName
	ContainerFieldName *string `json:"ContainerFieldName,omitempty" xml:"ContainerFieldName,omitempty"`
	// The value that corresponds to **ContainerFieldName**.
	//
	// example:
	//
	// cc914b0df156d40148412afe4a581****
	ContainerFieldValue *string `json:"ContainerFieldValue,omitempty" xml:"ContainerFieldValue,omitempty"`
	// The page number of the first page to display in the query results. Default value: **1**, which indicates that the results start from page 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The CVE ID.
	//
	// > Call the [DescribeVulListPage](~~DescribeVulListPage~~) operation to obtain this parameter.
	//
	// example:
	//
	// CVE-2017-15420
	CveId *string `json:"CveId,omitempty" xml:"CveId,omitempty"`
	// Specifies whether the vulnerability is handled. Valid values:
	//
	// - **y**: handled
	//
	// - **n**: not handled.
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
	// The language type of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The priority of the vulnerability fix to query. Separate multiple priorities with commas (,). Valid values:
	//
	// - **asap**: high
	//
	// - **later**: medium
	//
	// - **nntf**: low.
	//
	// example:
	//
	// asap,later,nntf
	Necessity *string `json:"Necessity,omitempty" xml:"Necessity,omitempty"`
	// The number of vulnerability entries per page in a paged query. Default value: 10, which indicates that 10 vulnerability entries are displayed per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether Runtime Application Self-Protection (RASP) supports real-time protection against the vulnerability. Valid values:
	//
	// - **0**: Not supported.
	//
	// - **1**: Supported.
	//
	// example:
	//
	// 1
	RaspDefend *int32 `json:"RaspDefend,omitempty" xml:"RaspDefend,omitempty"`
	// The ID of the Alibaba Cloud account that is added as one of the member accounts in a resource folder.
	//
	// > Invoke the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain this parameter.
	//
	// example:
	//
	// 127608589417****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The label used for filtering. Valid values:
	//
	// - **Restart required**
	//
	// - **Remote utilization**
	//
	// - **EXP exists**
	//
	// - **Available**
	//
	// - **Privilege escalation**
	//
	// - **Code execution**
	//
	// example:
	//
	// Code Execution
	SearchTags *string `json:"SearchTags,omitempty" xml:"SearchTags,omitempty"`
	// The container query type. Valid values:
	//
	// - **containerId**: container ID
	//
	// - **uuid**: asset ID.
	//
	// example:
	//
	// containerId
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The type of the vulnerability to query. Default value: cve. Valid values:
	//
	// - **cve**: Linux software vulnerability
	//
	// - **sys**: Windows system vulnerability
	//
	// - **cms**: Web-CMS vulnerability
	//
	// - **app**: application vulnerability (network scan)
	//
	// - **sca**: application vulnerability (software constituency parsing).
	//
	// example:
	//
	// cve
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The UUIDs of the servers to query. Separate multiple UUIDs with commas (,).
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
