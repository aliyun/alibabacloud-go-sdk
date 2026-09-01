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
	// The conditions for searching assets. This parameter is in JSON format. Note that the parameter values are case-sensitive.
	//
	// > You can search for assets by instance ID, instance name, VPC ID, region, public IP address, and other conditions. Call the [DescribeCriteria](~~DescribeCriteria~~) operation to query the supported search conditions.
	//
	// example:
	//
	// [{"name":"riskStatus","value":"YES"},{"name":"internetIp","value":"1.2.XX.XX"}]
	Criteria *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// The page number of the first page to return. Default value: **1**, which indicates that the query results are returned starting from page 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The asset vendor. Separate multiple asset vendors with commas (,). Valid values:
	//
	// example:
	//
	// 1,2,3
	Flags *string `json:"Flags,omitempty" xml:"Flags,omitempty"`
	// The importance level of the asset. Valid values:
	//
	// - **2**: Important asset.
	//
	// - **1**: General asset.
	//
	// - **0**: Test asset.
	//
	// example:
	//
	// 2
	Importance *int32 `json:"Importance,omitempty" xml:"Importance,omitempty"`
	// The language of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The logical relationship between multiple search conditions. Default value: **OR**. Valid values:
	//
	// - **OR**: The search conditions have an **OR*	- relationship.
	//
	// - **AND**: The search conditions have an **AND*	- relationship.
	//
	// example:
	//
	// OR
	LogicalExp *string `json:"LogicalExp,omitempty" xml:"LogicalExp,omitempty"`
	// The type of asset to query. Valid values:
	//
	// - **ecs**: server.
	//
	// - **cloud_product**: cloud product.
	//
	// - **eci**: elastic container instance.
	//
	// - **rund**: RunD container instance.
	//
	// - **runc**: RunC container instance.
	//
	// example:
	//
	// ecs
	MachineTypes *string `json:"MachineTypes,omitempty" xml:"MachineTypes,omitempty"`
	// The NextToken value returned when the NextToken method is used. Leave this parameter empty for the first request.
	//
	// example:
	//
	// E17B501887A2D3AA5E8360A6EFA3B***
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Specifies whether to disable internationalization for the default group name **未分组**. Default value: **false**. Valid values:
	//
	// - **true**: Internationalization is disabled. If the value of the GroupTrace response parameter is the default Security Center group **未分组**, the value is still displayed as **未分组**.
	//
	// - **false**: Internationalization is enabled. If the value of the GroupTrace response parameter is the default Security Center group **未分组**, the value is displayed as **default**.
	//
	// example:
	//
	// false
	NoGroupTrace *bool `json:"NoGroupTrace,omitempty" xml:"NoGroupTrace,omitempty"`
	// The number of assets to display on each page in a paged conditional query. Default value: **20**, which indicates that 20 asset records are displayed on each page.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Deprecated
	//
	// The region ID of the instance to query.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud account that corresponds to the member account in the resource directory.
	//
	// >Call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain this parameter.
	//
	// example:
	//
	// 1232428423234****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// Specifies whether to use the NextToken method to retrieve asset list data. If this parameter is set to true, TotalCount is no longer returned. Valid values:
	//
	// - **true**: Uses the NextToken method.
	//
	// - **false**: Does not use the NextToken method.
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
