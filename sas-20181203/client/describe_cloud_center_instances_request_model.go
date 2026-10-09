// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudCenterInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCriteria(v string) *DescribeCloudCenterInstancesRequest
	GetCriteria() *string
	SetCurrentPage(v int32) *DescribeCloudCenterInstancesRequest
	GetCurrentPage() *int32
	SetFlags(v string) *DescribeCloudCenterInstancesRequest
	GetFlags() *string
	SetImportance(v int32) *DescribeCloudCenterInstancesRequest
	GetImportance() *int32
	SetLang(v string) *DescribeCloudCenterInstancesRequest
	GetLang() *string
	SetLogicalExp(v string) *DescribeCloudCenterInstancesRequest
	GetLogicalExp() *string
	SetMachineTypes(v string) *DescribeCloudCenterInstancesRequest
	GetMachineTypes() *string
	SetNextToken(v string) *DescribeCloudCenterInstancesRequest
	GetNextToken() *string
	SetNoGroupTrace(v bool) *DescribeCloudCenterInstancesRequest
	GetNoGroupTrace() *bool
	SetPageSize(v int32) *DescribeCloudCenterInstancesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeCloudCenterInstancesRequest
	GetRegionId() *string
	SetResourceDirectoryAccountId(v int64) *DescribeCloudCenterInstancesRequest
	GetResourceDirectoryAccountId() *int64
	SetUseNextToken(v bool) *DescribeCloudCenterInstancesRequest
	GetUseNextToken() *bool
}

type DescribeCloudCenterInstancesRequest struct {
	// The search criteria for assets. This parameter is in JSON format. Pay attention to case sensitivity when entering parameter values.
	//
	// > You can search for assets by instance ID, instance name, VPC ID, region, public IP address, and other criteria. Call [DescribeCriteria](~~DescribeCriteria~~) to query the supported search criteria.
	//
	// example:
	//
	// [{"name":"riskStatus","value":"YES"},{"name":"internetIp","value":"1.2.XX.XX"}]
	Criteria *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// The page number from which to start displaying query results. Default value: **1**. This means results are displayed starting from page 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The asset vendor. Separate multiple vendors with commas (,). Valid values:
	//
	// - **0**: Alibaba Cloud asset
	//
	// - **1**: off-cloud asset
	//
	// - **2**: IDC asset
	//
	// - **3**, **4**, **5**, **7**, **14**, **16**: assets from other cloud vendors
	//
	// - **8**: lightweight asset
	//
	// - **9**: SAE
	//
	// - **10**: PAI
	//
	// example:
	//
	// 1,2,3
	Flags *string `json:"Flags,omitempty" xml:"Flags,omitempty"`
	// The importance level of the asset. Valid values:
	//
	// - **2**: important asset
	//
	// - **1**: general asset
	//
	// - **0**: test asset
	//
	// example:
	//
	// 2
	Importance *int32 `json:"Importance,omitempty" xml:"Importance,omitempty"`
	// The language of the request and response messages. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The logical relationship between multiple search criteria. Default value: **OR**. Valid values:
	//
	// - **OR**: The multiple search criteria have an OR relationship.
	//
	// - **AND**: The multiple search criteria have an AND relationship.
	//
	// example:
	//
	// OR
	LogicalExp *string `json:"LogicalExp,omitempty" xml:"LogicalExp,omitempty"`
	// The type of assets to query. Valid values:
	//
	// - **ecs**: server
	//
	// - **cloud_product**: cloud product
	//
	// - **eci**: Elastic Container Instance
	//
	// - **rund**: RunD container instance
	//
	// - **runc**: RunC container instance
	//
	// example:
	//
	// ecs
	MachineTypes *string `json:"MachineTypes,omitempty" xml:"MachineTypes,omitempty"`
	// The NextToken value returned when using the NextToken method. Leave this parameter empty for the first request.
	//
	// example:
	//
	// E17B501887A2D3AA5E8360A6EFA3B***
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Specifies whether to apply internationalization to the default group **Ungrouped**. Default value: **false**. Valid values:
	//
	// - **true**: Internationalization is not applied. When the GroupTrace parameter returns the Security Center default group **Ungrouped**, it is still displayed as **Ungrouped**.
	//
	// - **false**: Internationalization is applied. When the GroupTrace parameter returns the Security Center default group **Ungrouped**, it is displayed as **default**.
	//
	// example:
	//
	// false
	NoGroupTrace *bool `json:"NoGroupTrace,omitempty" xml:"NoGroupTrace,omitempty"`
	// The number of assets to display per page in a paged query. Settings take effect per page. Default value: **20**. This means 20 assets are displayed per page.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Deprecated
	//
	// The ID of the region where the instance to query resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The primary account ID of the resource directory member accounts.
	//
	// > Call [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) to obtain this parameter.
	//
	// example:
	//
	// 1232428423234****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// Specifies whether to use the NextToken method to retrieve the asset list. If this parameter is set to true, TotalCount is no longer returned. Valid values:
	//
	// - **true**: Use the NextToken method.
	//
	// - **false**: Do not use the NextToken method.
	//
	// example:
	//
	// false
	UseNextToken *bool `json:"UseNextToken,omitempty" xml:"UseNextToken,omitempty"`
}

func (s DescribeCloudCenterInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudCenterInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudCenterInstancesRequest) GetCriteria() *string {
	return s.Criteria
}

func (s *DescribeCloudCenterInstancesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeCloudCenterInstancesRequest) GetFlags() *string {
	return s.Flags
}

func (s *DescribeCloudCenterInstancesRequest) GetImportance() *int32 {
	return s.Importance
}

func (s *DescribeCloudCenterInstancesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeCloudCenterInstancesRequest) GetLogicalExp() *string {
	return s.LogicalExp
}

func (s *DescribeCloudCenterInstancesRequest) GetMachineTypes() *string {
	return s.MachineTypes
}

func (s *DescribeCloudCenterInstancesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeCloudCenterInstancesRequest) GetNoGroupTrace() *bool {
	return s.NoGroupTrace
}

func (s *DescribeCloudCenterInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCloudCenterInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCloudCenterInstancesRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *DescribeCloudCenterInstancesRequest) GetUseNextToken() *bool {
	return s.UseNextToken
}

func (s *DescribeCloudCenterInstancesRequest) SetCriteria(v string) *DescribeCloudCenterInstancesRequest {
	s.Criteria = &v
	return s
}

func (s *DescribeCloudCenterInstancesRequest) SetCurrentPage(v int32) *DescribeCloudCenterInstancesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCloudCenterInstancesRequest) SetFlags(v string) *DescribeCloudCenterInstancesRequest {
	s.Flags = &v
	return s
}

func (s *DescribeCloudCenterInstancesRequest) SetImportance(v int32) *DescribeCloudCenterInstancesRequest {
	s.Importance = &v
	return s
}

func (s *DescribeCloudCenterInstancesRequest) SetLang(v string) *DescribeCloudCenterInstancesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCloudCenterInstancesRequest) SetLogicalExp(v string) *DescribeCloudCenterInstancesRequest {
	s.LogicalExp = &v
	return s
}

func (s *DescribeCloudCenterInstancesRequest) SetMachineTypes(v string) *DescribeCloudCenterInstancesRequest {
	s.MachineTypes = &v
	return s
}

func (s *DescribeCloudCenterInstancesRequest) SetNextToken(v string) *DescribeCloudCenterInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeCloudCenterInstancesRequest) SetNoGroupTrace(v bool) *DescribeCloudCenterInstancesRequest {
	s.NoGroupTrace = &v
	return s
}

func (s *DescribeCloudCenterInstancesRequest) SetPageSize(v int32) *DescribeCloudCenterInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudCenterInstancesRequest) SetRegionId(v string) *DescribeCloudCenterInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCloudCenterInstancesRequest) SetResourceDirectoryAccountId(v int64) *DescribeCloudCenterInstancesRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DescribeCloudCenterInstancesRequest) SetUseNextToken(v bool) *DescribeCloudCenterInstancesRequest {
	s.UseNextToken = &v
	return s
}

func (s *DescribeCloudCenterInstancesRequest) Validate() error {
	return dara.Validate(s)
}
