// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVulListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliasName(v string) *DescribeVulListRequest
	GetAliasName() *string
	SetAttachTypes(v string) *DescribeVulListRequest
	GetAttachTypes() *string
	SetClusterId(v string) *DescribeVulListRequest
	GetClusterId() *string
	SetCurrentPage(v int32) *DescribeVulListRequest
	GetCurrentPage() *int32
	SetDealed(v string) *DescribeVulListRequest
	GetDealed() *string
	SetGroupId(v string) *DescribeVulListRequest
	GetGroupId() *string
	SetIds(v string) *DescribeVulListRequest
	GetIds() *string
	SetLang(v string) *DescribeVulListRequest
	GetLang() *string
	SetName(v string) *DescribeVulListRequest
	GetName() *string
	SetNecessity(v string) *DescribeVulListRequest
	GetNecessity() *string
	SetNextToken(v string) *DescribeVulListRequest
	GetNextToken() *string
	SetPageSize(v int32) *DescribeVulListRequest
	GetPageSize() *int32
	SetRaspDefend(v int32) *DescribeVulListRequest
	GetRaspDefend() *int32
	SetRemark(v string) *DescribeVulListRequest
	GetRemark() *string
	SetResourceDirectoryAccountId(v int64) *DescribeVulListRequest
	GetResourceDirectoryAccountId() *int64
	SetStatusList(v string) *DescribeVulListRequest
	GetStatusList() *string
	SetTargetType(v string) *DescribeVulListRequest
	GetTargetType() *string
	SetType(v string) *DescribeVulListRequest
	GetType() *string
	SetUseNextToken(v bool) *DescribeVulListRequest
	GetUseNextToken() *bool
	SetUuids(v string) *DescribeVulListRequest
	GetUuids() *string
	SetVpcInstanceIds(v string) *DescribeVulListRequest
	GetVpcInstanceIds() *string
}

type DescribeVulListRequest struct {
	// The name of the vulnerability.
	//
	// example:
	//
	// RHSA-2019:0230-Important: polkit security update
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The additional type of the vulnerabilities. You need to specify this parameter when you query application vulnerabilities. Set the value to **sca**. If you set **Type*	- to **app**, you must specify this parameter.
	//
	// > If you set this parameter to **sca**, application vulnerabilities and the vulnerabilities that are detected based on software component analysis are queried. If you do not specify this parameter, only application vulnerabilities are queried.
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
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Specifies whether the vulnerabilities are fixed. Valid values:
	//
	// 	- **y**: yes
	//
	// 	- **n**: no
	//
	// example:
	//
	// n
	Dealed *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	// The ID of the asset group.
	//
	// > You can call the [DescribeAllGroups](~~DescribeAllGroups~~) operation to query the IDs of asset groups.
	//
	// example:
	//
	// 9207613
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The IDs of vulnerabilities. You can specify up to 50 IDs. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// 282,281,283
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
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
	// The alias of the vulnerability.
	//
	// example:
	//
	// oval:com.redhat.rhsa:def:20172836
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// The pagination token that is used in the next request to retrieve a new page of results. You must specify the token that is obtained from the previous query as the value of NextToken. You do not need to specify this parameter for the first request.
	//
	// example:
	//
	// E17B501887A2D3AA5E8360A6EFA3B***
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries per page. Default value: **10**.
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
	// 0
	RaspDefend *int32 `json:"RaspDefend,omitempty" xml:"RaspDefend,omitempty"`
	// The remarks for the asset affected by the vulnerability. The value can be the private IP address, public IP address, or name of the asset.
	//
	// example:
	//
	// 1.2.XX.XX
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The Alibaba Cloud account ID of the member in the resource directory.
	//
	// >  You can call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain the IDs.
	//
	// example:
	//
	// 1232428423234****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The status of the vulnerability. Separate multiple statuses with commas (,). Valid values:
	//
	// - 1: unfixed
	//
	// - 2: fix failed
	//
	// - 3: rollback failed
	//
	// - 4: being fixed
	//
	// - 5: being rolled back
	//
	// - 6: being verified
	//
	// - 7: fixed
	//
	// - 8: fixed and to be restarted
	//
	// - 9: rolled back
	//
	// - 10: ignored
	//
	// - 11: rolled back and to be restarted
	//
	// - 12: not found
	//
	// - 20: expired
	//
	// example:
	//
	// 1,2,3
	StatusList *string `json:"StatusList,omitempty" xml:"StatusList,omitempty"`
	// The type of the asset on which the vulnerability is detected. Valid values:
	//
	// 	- **k8s**: Kubernetes component.
	//
	// 	- **uuid**: server.
	//
	// 	- **containerId**: container.
	//
	// example:
	//
	// k8s
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The type of the vulnerability. Valid values:
	//
	// 	- **cve**: Linux software vulnerability
	//
	// 	- **sys**: Windows system vulnerability
	//
	// 	- **cms**: Web-CMS vulnerability.
	//
	// 	- **app**: application vulnerability that is detected by using web scanner
	//
	// 	- **emg**: urgent vulnerability.
	//
	// 	- **sca**: application vulnerability that is detected by using software component analysis
	//
	// This parameter is required.
	//
	// example:
	//
	// cve
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Specifies whether to use NextToken to query the data of vulnerabilities. If you set UseNextToken to true, the value of TotalCount is not returned. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	UseNextToken *bool `json:"UseNextToken,omitempty" xml:"UseNextToken,omitempty"`
	// The UUIDs of the servers on which you want to query the vulnerabilities. Separate multiple UUIDs with commas (,).
	//
	// >  You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to obtain the UUIDs.
	//
	// example:
	//
	// 1587bedb-fdb4-48c4-9330-****
	Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
	// The ID of the virtual private cloud (VPC) in which the vulnerabilities are detected. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// ins-133****,ins-5414****
	VpcInstanceIds *string `json:"VpcInstanceIds,omitempty" xml:"VpcInstanceIds,omitempty"`
}

func (s DescribeVulListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVulListRequest) GetAliasName() *string {
	return s.AliasName
}

func (s *DescribeVulListRequest) GetAttachTypes() *string {
	return s.AttachTypes
}

func (s *DescribeVulListRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeVulListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeVulListRequest) GetDealed() *string {
	return s.Dealed
}

func (s *DescribeVulListRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeVulListRequest) GetIds() *string {
	return s.Ids
}

func (s *DescribeVulListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeVulListRequest) GetName() *string {
	return s.Name
}

func (s *DescribeVulListRequest) GetNecessity() *string {
	return s.Necessity
}

func (s *DescribeVulListRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeVulListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVulListRequest) GetRaspDefend() *int32 {
	return s.RaspDefend
}

func (s *DescribeVulListRequest) GetRemark() *string {
	return s.Remark
}

func (s *DescribeVulListRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *DescribeVulListRequest) GetStatusList() *string {
	return s.StatusList
}

func (s *DescribeVulListRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *DescribeVulListRequest) GetType() *string {
	return s.Type
}

func (s *DescribeVulListRequest) GetUseNextToken() *bool {
	return s.UseNextToken
}

func (s *DescribeVulListRequest) GetUuids() *string {
	return s.Uuids
}

func (s *DescribeVulListRequest) GetVpcInstanceIds() *string {
	return s.VpcInstanceIds
}

func (s *DescribeVulListRequest) SetAliasName(v string) *DescribeVulListRequest {
	s.AliasName = &v
	return s
}

func (s *DescribeVulListRequest) SetAttachTypes(v string) *DescribeVulListRequest {
	s.AttachTypes = &v
	return s
}

func (s *DescribeVulListRequest) SetClusterId(v string) *DescribeVulListRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeVulListRequest) SetCurrentPage(v int32) *DescribeVulListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVulListRequest) SetDealed(v string) *DescribeVulListRequest {
	s.Dealed = &v
	return s
}

func (s *DescribeVulListRequest) SetGroupId(v string) *DescribeVulListRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeVulListRequest) SetIds(v string) *DescribeVulListRequest {
	s.Ids = &v
	return s
}

func (s *DescribeVulListRequest) SetLang(v string) *DescribeVulListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVulListRequest) SetName(v string) *DescribeVulListRequest {
	s.Name = &v
	return s
}

func (s *DescribeVulListRequest) SetNecessity(v string) *DescribeVulListRequest {
	s.Necessity = &v
	return s
}

func (s *DescribeVulListRequest) SetNextToken(v string) *DescribeVulListRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeVulListRequest) SetPageSize(v int32) *DescribeVulListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVulListRequest) SetRaspDefend(v int32) *DescribeVulListRequest {
	s.RaspDefend = &v
	return s
}

func (s *DescribeVulListRequest) SetRemark(v string) *DescribeVulListRequest {
	s.Remark = &v
	return s
}

func (s *DescribeVulListRequest) SetResourceDirectoryAccountId(v int64) *DescribeVulListRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DescribeVulListRequest) SetStatusList(v string) *DescribeVulListRequest {
	s.StatusList = &v
	return s
}

func (s *DescribeVulListRequest) SetTargetType(v string) *DescribeVulListRequest {
	s.TargetType = &v
	return s
}

func (s *DescribeVulListRequest) SetType(v string) *DescribeVulListRequest {
	s.Type = &v
	return s
}

func (s *DescribeVulListRequest) SetUseNextToken(v bool) *DescribeVulListRequest {
	s.UseNextToken = &v
	return s
}

func (s *DescribeVulListRequest) SetUuids(v string) *DescribeVulListRequest {
	s.Uuids = &v
	return s
}

func (s *DescribeVulListRequest) SetVpcInstanceIds(v string) *DescribeVulListRequest {
	s.VpcInstanceIds = &v
	return s
}

func (s *DescribeVulListRequest) Validate() error {
	return dara.Validate(s)
}
