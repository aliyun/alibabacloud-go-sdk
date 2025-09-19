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
	// The search conditions. The value of this parameter is in the JSON format and is case-sensitive.
	//
	// >  You can search for an asset by using the search conditions, such as the instance ID, instance name, VPC ID, region, or public IP address. You can call the [DescribeCriteria](https://help.aliyun.com/document_detail/149773.html) operation to query the supported search conditions.
	//
	// example:
	//
	// [{"name":"riskStatus","value":"YES"},{"name":"internetIp","value":"1.2.XX.XX"}]
	Criteria *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Asset vendor. Multiple asset vendors should be separated by a comma (,). Values:
	//
	// - **0**: an asset provided by Alibaba Cloud
	//
	// - **1**: an asset outside Alibaba Cloud
	//
	// - **2**: an asset in a data center
	//
	// - **3**, **4**, **5**, **7**, **14**, **16**: an asset from a third-party cloud service provider
	//
	// - **8**: a lightweight asset
	//
	// - **9**: a Serverless App Engine (SAE) instance
	//
	// - **10**: an instance in Platform for AI (PAI)
	//
	// example:
	//
	// 1,2,3
	Flags *string `json:"Flags,omitempty" xml:"Flags,omitempty"`
	// The importance of the asset. Valid values:
	//
	// 	- **2**: an important asset
	//
	// 	- **1**: a common asset
	//
	// 	- **0**: a test asset
	//
	// example:
	//
	// 2
	Importance *int32 `json:"Importance,omitempty" xml:"Importance,omitempty"`
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
	// The logical relationship among multiple search conditions. Valid values:
	//
	// 	- **OR**: The logical relationship among search conditions is **OR**.
	//
	// 	- **AND**: The logical relationship among search conditions is **AND**.
	//
	// example:
	//
	// OR
	LogicalExp *string `json:"LogicalExp,omitempty" xml:"LogicalExp,omitempty"`
	// The type of asset to be queried. Values:
	//
	// - **ecs**: Server
	//
	// - **cloud_product**: Cloud Product
	//
	// - **eci**: Elastic Container Instance
	//
	// - **rund**: RunD Container Instance
	//
	// - **runc**: RunC Container Instance
	//
	// example:
	//
	// ecs
	MachineTypes *string `json:"MachineTypes,omitempty" xml:"MachineTypes,omitempty"`
	// The value of NextToken that is returned when the NextToken method is used. You do not need to specify this parameter for the first request.
	//
	// example:
	//
	// E17B501887A2D3AA5E8360A6EFA3B***
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Specifies whether to internationalize the name of the default group. Valid values:
	//
	// 	- **true**: The system returns the Chinese name of the default group for the GroupTrace response parameter.
	//
	// 	- **false**: The system returns default for the GroupTrace response parameter.
	//
	// example:
	//
	// false
	NoGroupTrace *bool `json:"NoGroupTrace,omitempty" xml:"NoGroupTrace,omitempty"`
	// The number of entries to return on each page. Default value: **20**.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Deprecated
	//
	// The ID of the region in which the asset resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Alibaba Cloud account ID of the member in the resource directory.
	//
	// >  You can call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain the IDs.
	//
	// example:
	//
	// 1232428423234****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// Specifies whether to use the NextToken method to retrieve a new page of results. If you set UseNextToken to true, the value of TotalCount is not returned. Valid values:
	//
	// - **true**: The NextToken method is used.
	//
	// - **false**: The NextToken method is not used.
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
