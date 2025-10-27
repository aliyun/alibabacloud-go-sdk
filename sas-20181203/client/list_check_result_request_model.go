// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCheckResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckIds(v []*int64) *ListCheckResultRequest
	GetCheckIds() []*int64
	SetCheckKey(v string) *ListCheckResultRequest
	GetCheckKey() *string
	SetCheckTypes(v []*string) *ListCheckResultRequest
	GetCheckTypes() []*string
	SetCurrentPage(v int32) *ListCheckResultRequest
	GetCurrentPage() *int32
	SetCustomParam(v bool) *ListCheckResultRequest
	GetCustomParam() *bool
	SetInstanceIds(v []*string) *ListCheckResultRequest
	GetInstanceIds() []*string
	SetInstanceTypes(v []*string) *ListCheckResultRequest
	GetInstanceTypes() []*string
	SetLang(v string) *ListCheckResultRequest
	GetLang() *string
	SetOperationTypes(v []*string) *ListCheckResultRequest
	GetOperationTypes() []*string
	SetPageSize(v int32) *ListCheckResultRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListCheckResultRequest
	GetRegionId() *string
	SetRequirementIds(v []*int64) *ListCheckResultRequest
	GetRequirementIds() []*int64
	SetResourceDirectoryAccountId(v int64) *ListCheckResultRequest
	GetResourceDirectoryAccountId() *int64
	SetRiskLevels(v []*string) *ListCheckResultRequest
	GetRiskLevels() []*string
	SetSortTypes(v []*string) *ListCheckResultRequest
	GetSortTypes() []*string
	SetStandardIds(v []*int64) *ListCheckResultRequest
	GetStandardIds() []*int64
	SetStatuses(v []*string) *ListCheckResultRequest
	GetStatuses() []*string
	SetTaskSources(v []*string) *ListCheckResultRequest
	GetTaskSources() []*string
	SetTypes(v []*string) *ListCheckResultRequest
	GetTypes() []*string
	SetVendors(v []*string) *ListCheckResultRequest
	GetVendors() []*string
}

type ListCheckResultRequest struct {
	// The IDs of the check items.
	CheckIds []*int64 `json:"CheckIds,omitempty" xml:"CheckIds,omitempty" type:"Repeated"`
	// The key that you want to use to search for check items in fuzzy match mode.
	//
	// example:
	//
	// OSS
	CheckKey *string `json:"CheckKey,omitempty" xml:"CheckKey,omitempty"`
	// Source type of the situation awareness check item.
	CheckTypes []*string `json:"CheckTypes,omitempty" xml:"CheckTypes,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 2
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Specifies whether the check item supports custom parameters. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	CustomParam *bool `json:"CustomParam,omitempty" xml:"CustomParam,omitempty"`
	// The instance IDs of the cloud services that you want to query. Separate multiple IDs with commas (,).
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The asset type of the cloud services. Valid values:
	//
	// 	- **ECS**: Elastic Compute Service (ECS)
	//
	// 	- **SLB**: Server Load Balancer (SLB)
	//
	// 	- **RDS**: ApsaraDB RDS
	//
	// 	- **MONGODB**: ApsaraDB for MongoDB (MongoDB)
	//
	// 	- **KVSTORE**: ApsaraDB for Redis (Redis)
	//
	// 	- **ACR**: Container Registry
	//
	// 	- **CSK**: Container Service for Kubernetes (ACK)
	//
	// 	- **VPC**: Virtual Private Cloud (VPC)
	//
	// 	- **ACTIONTRAIL**: ActionTrail
	//
	// 	- **CDN**: Alibaba Cloud CDN (CDN)
	//
	// 	- **CAS**: Certificate Management Service (formerly SSL Certificates Service)
	//
	// 	- **RDC**: Apsara Devops
	//
	// 	- **RAM**: Resource Access Management (RAM)
	//
	// 	- **DDOS**: Anti-DDoS
	//
	// 	- **WAF**: Web Application Firewall (WAF)
	//
	// 	- **OSS**: Object Storage Service (OSS)
	//
	// 	- **POLARDB**: PolarDB
	//
	// 	- **POSTGRESQL**: ApsaraDB RDS for PostgreSQL
	//
	// 	- **MSE**: Microservices Engine (MSE)
	//
	// 	- **NAS**: File Storage NAS (NAS)
	//
	// 	- **SDDP**: Sensitive Data Discovery and Protection (SDDP)
	//
	// 	- **EIP**: Elastic IP Address (EIP)
	InstanceTypes []*string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
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
	// Specifies whether fixing is supported. Valid values:
	//
	// 	- **SUPPORT_REPAIR**
	//
	// 	- **NOT_SUPPORT_REPAIR**
	OperationTypes []*string `json:"OperationTypes,omitempty" xml:"OperationTypes,omitempty" type:"Repeated"`
	// The number of entries per page. Maximum value: 100.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the instance. Valid values:
	//
	// 	- **cn-hangzhou**: International
	//
	// 	- **ap-southeast-1**: Singapore
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the requirements.
	RequirementIds []*int64 `json:"RequirementIds,omitempty" xml:"RequirementIds,omitempty" type:"Repeated"`
	// The Alibaba Cloud account ID of the member in the resource directory.
	//
	// >  You can call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain the IDs.
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The risk levels of check items. Separate multiple risk levels with commas (,). Valid values:
	//
	// 	- **HIGH**
	//
	// 	- **MEDIUM**
	//
	// 	- **LOW**
	RiskLevels []*string `json:"RiskLevels,omitempty" xml:"RiskLevels,omitempty" type:"Repeated"`
	// The types of the conditions based on which check items are sorted. Valid values:
	//
	// 	- **RISK_LEVEL**: risk level
	//
	// 	- **STATUS**: status
	SortTypes []*string `json:"SortTypes,omitempty" xml:"SortTypes,omitempty" type:"Repeated"`
	// The standard IDs.
	StandardIds []*int64 `json:"StandardIds,omitempty" xml:"StandardIds,omitempty" type:"Repeated"`
	// The statuses of check items. Separate multiple statuses with commas (,). Valid values:
	//
	// 	- **PASS**
	//
	// 	- **NOT_PASS**
	//
	// 	- **CHECKING**
	//
	// 	- **NOT_CHECK**
	//
	// 	- **WHITELIST**
	Statuses []*string `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
	// Delete the custom category in a custom inspection item.
	TaskSources []*string `json:"TaskSources,omitempty" xml:"TaskSources,omitempty" type:"Repeated"`
	// The types of check standards.
	Types []*string `json:"Types,omitempty" xml:"Types,omitempty" type:"Repeated"`
	// The cloud service providers. Valid values:
	//
	// 	- **ALIYUN**: Alibaba Cloud
	//
	// 	- **TENCENT**: Tencent Cloud
	//
	// 	- **AWS**: Amazon Web Services (AWS)
	//
	// 	- **MICROSOFT**: Microsoft Azure
	Vendors []*string `json:"Vendors,omitempty" xml:"Vendors,omitempty" type:"Repeated"`
}

func (s ListCheckResultRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCheckResultRequest) GoString() string {
	return s.String()
}

func (s *ListCheckResultRequest) GetCheckIds() []*int64 {
	return s.CheckIds
}

func (s *ListCheckResultRequest) GetCheckKey() *string {
	return s.CheckKey
}

func (s *ListCheckResultRequest) GetCheckTypes() []*string {
	return s.CheckTypes
}

func (s *ListCheckResultRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListCheckResultRequest) GetCustomParam() *bool {
	return s.CustomParam
}

func (s *ListCheckResultRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ListCheckResultRequest) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *ListCheckResultRequest) GetLang() *string {
	return s.Lang
}

func (s *ListCheckResultRequest) GetOperationTypes() []*string {
	return s.OperationTypes
}

func (s *ListCheckResultRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCheckResultRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListCheckResultRequest) GetRequirementIds() []*int64 {
	return s.RequirementIds
}

func (s *ListCheckResultRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *ListCheckResultRequest) GetRiskLevels() []*string {
	return s.RiskLevels
}

func (s *ListCheckResultRequest) GetSortTypes() []*string {
	return s.SortTypes
}

func (s *ListCheckResultRequest) GetStandardIds() []*int64 {
	return s.StandardIds
}

func (s *ListCheckResultRequest) GetStatuses() []*string {
	return s.Statuses
}

func (s *ListCheckResultRequest) GetTaskSources() []*string {
	return s.TaskSources
}

func (s *ListCheckResultRequest) GetTypes() []*string {
	return s.Types
}

func (s *ListCheckResultRequest) GetVendors() []*string {
	return s.Vendors
}

func (s *ListCheckResultRequest) SetCheckIds(v []*int64) *ListCheckResultRequest {
	s.CheckIds = v
	return s
}

func (s *ListCheckResultRequest) SetCheckKey(v string) *ListCheckResultRequest {
	s.CheckKey = &v
	return s
}

func (s *ListCheckResultRequest) SetCheckTypes(v []*string) *ListCheckResultRequest {
	s.CheckTypes = v
	return s
}

func (s *ListCheckResultRequest) SetCurrentPage(v int32) *ListCheckResultRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListCheckResultRequest) SetCustomParam(v bool) *ListCheckResultRequest {
	s.CustomParam = &v
	return s
}

func (s *ListCheckResultRequest) SetInstanceIds(v []*string) *ListCheckResultRequest {
	s.InstanceIds = v
	return s
}

func (s *ListCheckResultRequest) SetInstanceTypes(v []*string) *ListCheckResultRequest {
	s.InstanceTypes = v
	return s
}

func (s *ListCheckResultRequest) SetLang(v string) *ListCheckResultRequest {
	s.Lang = &v
	return s
}

func (s *ListCheckResultRequest) SetOperationTypes(v []*string) *ListCheckResultRequest {
	s.OperationTypes = v
	return s
}

func (s *ListCheckResultRequest) SetPageSize(v int32) *ListCheckResultRequest {
	s.PageSize = &v
	return s
}

func (s *ListCheckResultRequest) SetRegionId(v string) *ListCheckResultRequest {
	s.RegionId = &v
	return s
}

func (s *ListCheckResultRequest) SetRequirementIds(v []*int64) *ListCheckResultRequest {
	s.RequirementIds = v
	return s
}

func (s *ListCheckResultRequest) SetResourceDirectoryAccountId(v int64) *ListCheckResultRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *ListCheckResultRequest) SetRiskLevels(v []*string) *ListCheckResultRequest {
	s.RiskLevels = v
	return s
}

func (s *ListCheckResultRequest) SetSortTypes(v []*string) *ListCheckResultRequest {
	s.SortTypes = v
	return s
}

func (s *ListCheckResultRequest) SetStandardIds(v []*int64) *ListCheckResultRequest {
	s.StandardIds = v
	return s
}

func (s *ListCheckResultRequest) SetStatuses(v []*string) *ListCheckResultRequest {
	s.Statuses = v
	return s
}

func (s *ListCheckResultRequest) SetTaskSources(v []*string) *ListCheckResultRequest {
	s.TaskSources = v
	return s
}

func (s *ListCheckResultRequest) SetTypes(v []*string) *ListCheckResultRequest {
	s.Types = v
	return s
}

func (s *ListCheckResultRequest) SetVendors(v []*string) *ListCheckResultRequest {
	s.Vendors = v
	return s
}

func (s *ListCheckResultRequest) Validate() error {
	return dara.Validate(s)
}
