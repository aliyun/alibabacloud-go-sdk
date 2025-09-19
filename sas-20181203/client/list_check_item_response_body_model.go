// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCheckItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCheckItems(v []*ListCheckItemResponseBodyCheckItems) *ListCheckItemResponseBody
	GetCheckItems() []*ListCheckItemResponseBodyCheckItems
	SetPageInfo(v *ListCheckItemResponseBodyPageInfo) *ListCheckItemResponseBody
	GetPageInfo() *ListCheckItemResponseBodyPageInfo
	SetRequestId(v string) *ListCheckItemResponseBody
	GetRequestId() *string
}

type ListCheckItemResponseBody struct {
	// The check items.
	CheckItems []*ListCheckItemResponseBodyCheckItems `json:"CheckItems,omitempty" xml:"CheckItems,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *ListCheckItemResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 9F4E6157-9600-5588-86B9-38F09067****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCheckItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCheckItemResponseBody) GoString() string {
	return s.String()
}

func (s *ListCheckItemResponseBody) GetCheckItems() []*ListCheckItemResponseBodyCheckItems {
	return s.CheckItems
}

func (s *ListCheckItemResponseBody) GetPageInfo() *ListCheckItemResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListCheckItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCheckItemResponseBody) SetCheckItems(v []*ListCheckItemResponseBodyCheckItems) *ListCheckItemResponseBody {
	s.CheckItems = v
	return s
}

func (s *ListCheckItemResponseBody) SetPageInfo(v *ListCheckItemResponseBodyPageInfo) *ListCheckItemResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListCheckItemResponseBody) SetRequestId(v string) *ListCheckItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCheckItemResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListCheckItemResponseBodyCheckItems struct {
	// The ID of the check item.
	//
	// example:
	//
	// 21
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The name of the check item.
	//
	// example:
	//
	// Enable deletion protection
	CheckShowName *string `json:"CheckShowName,omitempty" xml:"CheckShowName,omitempty"`
	// The source type of the Situation Awareness check item:
	//
	// - **CUSTOM**: User-defined
	//
	// - **SYSTEM**: Predefined by the Situation Awareness platform
	//
	// example:
	//
	// SYSTEM
	CheckType *string `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	// The check items.
	CustomConfigs []*ListCheckItemResponseBodyCheckItemsCustomConfigs `json:"CustomConfigs,omitempty" xml:"CustomConfigs,omitempty" type:"Repeated"`
	// The description of the check item.
	Description *ListCheckItemResponseBodyCheckItemsDescription `json:"Description,omitempty" xml:"Description,omitempty" type:"Struct"`
	// The estimated quota that will be consumed by this check item.
	//
	// example:
	//
	// 30
	EstimatedCount *int32 `json:"EstimatedCount,omitempty" xml:"EstimatedCount,omitempty"`
	// The asset subtype of the cloud service. Valid values:
	//
	// 	- If **InstanceType*	- is set to **ECS**, this parameter supports the following valid values:
	//
	//     	- **INSTANCE**
	//
	//     	- **DISK**
	//
	//     	- **SECURITY_GROUP**
	//
	// 	- If **InstanceType*	- is set to **ACR**, this parameter supports the following valid values:
	//
	//     	- **REPOSITORY_ENTERPRISE**
	//
	//     	- **REPOSITORY_PERSON**
	//
	// 	- If **InstanceType*	- is set to **RAM**, this parameter supports the following valid values:
	//
	//     	- **ALIAS**
	//
	//     	- **USER**
	//
	//     	- **POLICY**
	//
	//     	- **GROUP**
	//
	// 	- If **InstanceType*	- is set to **WAF**, this parameter supports the following valid value:
	//
	//     	- **DOMAIN**
	//
	// 	- If **InstanceType*	- is set to other values, this parameter supports the following valid values:
	//
	//     	- **INSTANCE**
	//
	// example:
	//
	// ECS
	InstanceSubType *string `json:"InstanceSubType,omitempty" xml:"InstanceSubType,omitempty"`
	// The asset type of the cloud service. Valid values:
	//
	// 	- **ECS**: Elastic Compute Service (ECS).
	//
	// 	- **SLB**: Server Load Balancer (SLB).
	//
	// 	- **RDS**: ApsaraDB RDS.
	//
	// 	- **MONGODB**: ApsaraDB for MongoDB (MongoDB).
	//
	// 	- **KVSTORE**: ApsaraDB for Redis (Redis).
	//
	// 	- **ACR**: Container Registry.
	//
	// 	- **CSK**: Container Service for Kubernetes (ACK).
	//
	// 	- **VPC**: Virtual Private Cloud (VPC).
	//
	// 	- **ACTIONTRAIL**: ActionTrail.
	//
	// 	- **CDN**: Alibaba Cloud CDN (CDN).
	//
	// 	- **CAS**: Certificate Management Service (formerly SSL Certificates Service).
	//
	// 	- **RDC**: Apsara Devops.
	//
	// 	- **RAM**: Resource Access Management (RAM).
	//
	// 	- **DDOS**: Anti-DDoS.
	//
	// 	- **WAF**: Web Application Firewall (WAF).
	//
	// 	- **OSS**: Object Storage Service (OSS).
	//
	// 	- **POLARDB**: PolarDB.
	//
	// 	- **POSTGRESQL**: ApsaraDB RDS for PostgreSQL.
	//
	// 	- **MSE**: Microservices Engine (MSE).
	//
	// 	- **NAS**: File Storage NAS (NAS).
	//
	// 	- **SDDP**: Sensitive Data Discovery and Protection (SDDP).
	//
	// 	- **EIP**: Elastic IP Address (EIP).
	//
	// example:
	//
	// OSS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The risk level of the check item. Valid values:
	//
	// 	- **HIGH**
	//
	// 	- **MEDIUM**
	//
	// 	- **LOW**
	//
	// example:
	//
	// HIGH
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The IDs of the sections associated with the check items.
	SectionIds []*int64 `json:"SectionIds,omitempty" xml:"SectionIds,omitempty" type:"Repeated"`
	// The type of the cloud asset. Valid values:
	//
	// 	- **0**: an asset provided by Alibaba Cloud.
	//
	// 	- **1**: an asset outside Alibaba Cloud.
	//
	// 	- **2**: an asset in a data center.
	//
	// 	- **3**, **4**, **5**, and **7**: other cloud asset.
	//
	// 	- **8**: a simple application server.
	//
	// example:
	//
	// 0
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s ListCheckItemResponseBodyCheckItems) String() string {
	return dara.Prettify(s)
}

func (s ListCheckItemResponseBodyCheckItems) GoString() string {
	return s.String()
}

func (s *ListCheckItemResponseBodyCheckItems) GetCheckId() *int64 {
	return s.CheckId
}

func (s *ListCheckItemResponseBodyCheckItems) GetCheckShowName() *string {
	return s.CheckShowName
}

func (s *ListCheckItemResponseBodyCheckItems) GetCheckType() *string {
	return s.CheckType
}

func (s *ListCheckItemResponseBodyCheckItems) GetCustomConfigs() []*ListCheckItemResponseBodyCheckItemsCustomConfigs {
	return s.CustomConfigs
}

func (s *ListCheckItemResponseBodyCheckItems) GetDescription() *ListCheckItemResponseBodyCheckItemsDescription {
	return s.Description
}

func (s *ListCheckItemResponseBodyCheckItems) GetEstimatedCount() *int32 {
	return s.EstimatedCount
}

func (s *ListCheckItemResponseBodyCheckItems) GetInstanceSubType() *string {
	return s.InstanceSubType
}

func (s *ListCheckItemResponseBodyCheckItems) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListCheckItemResponseBodyCheckItems) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *ListCheckItemResponseBodyCheckItems) GetSectionIds() []*int64 {
	return s.SectionIds
}

func (s *ListCheckItemResponseBodyCheckItems) GetVendor() *string {
	return s.Vendor
}

func (s *ListCheckItemResponseBodyCheckItems) SetCheckId(v int64) *ListCheckItemResponseBodyCheckItems {
	s.CheckId = &v
	return s
}

func (s *ListCheckItemResponseBodyCheckItems) SetCheckShowName(v string) *ListCheckItemResponseBodyCheckItems {
	s.CheckShowName = &v
	return s
}

func (s *ListCheckItemResponseBodyCheckItems) SetCheckType(v string) *ListCheckItemResponseBodyCheckItems {
	s.CheckType = &v
	return s
}

func (s *ListCheckItemResponseBodyCheckItems) SetCustomConfigs(v []*ListCheckItemResponseBodyCheckItemsCustomConfigs) *ListCheckItemResponseBodyCheckItems {
	s.CustomConfigs = v
	return s
}

func (s *ListCheckItemResponseBodyCheckItems) SetDescription(v *ListCheckItemResponseBodyCheckItemsDescription) *ListCheckItemResponseBodyCheckItems {
	s.Description = v
	return s
}

func (s *ListCheckItemResponseBodyCheckItems) SetEstimatedCount(v int32) *ListCheckItemResponseBodyCheckItems {
	s.EstimatedCount = &v
	return s
}

func (s *ListCheckItemResponseBodyCheckItems) SetInstanceSubType(v string) *ListCheckItemResponseBodyCheckItems {
	s.InstanceSubType = &v
	return s
}

func (s *ListCheckItemResponseBodyCheckItems) SetInstanceType(v string) *ListCheckItemResponseBodyCheckItems {
	s.InstanceType = &v
	return s
}

func (s *ListCheckItemResponseBodyCheckItems) SetRiskLevel(v string) *ListCheckItemResponseBodyCheckItems {
	s.RiskLevel = &v
	return s
}

func (s *ListCheckItemResponseBodyCheckItems) SetSectionIds(v []*int64) *ListCheckItemResponseBodyCheckItems {
	s.SectionIds = v
	return s
}

func (s *ListCheckItemResponseBodyCheckItems) SetVendor(v string) *ListCheckItemResponseBodyCheckItems {
	s.Vendor = &v
	return s
}

func (s *ListCheckItemResponseBodyCheckItems) Validate() error {
	return dara.Validate(s)
}

type ListCheckItemResponseBodyCheckItemsCustomConfigs struct {
	// The default value of the check item. The value is a string.
	//
	// example:
	//
	// 0
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// The name of the check item.
	//
	// example:
	//
	// IPList
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The display name of the check item.
	//
	// example:
	//
	// Ensure RAM password policy prevents password reuse
	ShowName *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
	// The type of the check item. The value is a JSON string.
	//
	// example:
	//
	// {\\"type\\":\\"LIST\\",\\"range\\":[1,512],\\"listType\\":{\\"type\\":\\"STRING\\",\\"range\\":[0,22]}}
	TypeDefine *string `json:"TypeDefine,omitempty" xml:"TypeDefine,omitempty"`
	// The specified value of the check item. The value is a string.
	//
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListCheckItemResponseBodyCheckItemsCustomConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListCheckItemResponseBodyCheckItemsCustomConfigs) GoString() string {
	return s.String()
}

func (s *ListCheckItemResponseBodyCheckItemsCustomConfigs) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *ListCheckItemResponseBodyCheckItemsCustomConfigs) GetName() *string {
	return s.Name
}

func (s *ListCheckItemResponseBodyCheckItemsCustomConfigs) GetShowName() *string {
	return s.ShowName
}

func (s *ListCheckItemResponseBodyCheckItemsCustomConfigs) GetTypeDefine() *string {
	return s.TypeDefine
}

func (s *ListCheckItemResponseBodyCheckItemsCustomConfigs) GetValue() *string {
	return s.Value
}

func (s *ListCheckItemResponseBodyCheckItemsCustomConfigs) SetDefaultValue(v string) *ListCheckItemResponseBodyCheckItemsCustomConfigs {
	s.DefaultValue = &v
	return s
}

func (s *ListCheckItemResponseBodyCheckItemsCustomConfigs) SetName(v string) *ListCheckItemResponseBodyCheckItemsCustomConfigs {
	s.Name = &v
	return s
}

func (s *ListCheckItemResponseBodyCheckItemsCustomConfigs) SetShowName(v string) *ListCheckItemResponseBodyCheckItemsCustomConfigs {
	s.ShowName = &v
	return s
}

func (s *ListCheckItemResponseBodyCheckItemsCustomConfigs) SetTypeDefine(v string) *ListCheckItemResponseBodyCheckItemsCustomConfigs {
	s.TypeDefine = &v
	return s
}

func (s *ListCheckItemResponseBodyCheckItemsCustomConfigs) SetValue(v string) *ListCheckItemResponseBodyCheckItemsCustomConfigs {
	s.Value = &v
	return s
}

func (s *ListCheckItemResponseBodyCheckItemsCustomConfigs) Validate() error {
	return dara.Validate(s)
}

type ListCheckItemResponseBodyCheckItemsDescription struct {
	// The type of the description of the check item. Valid value:
	//
	// 	- **text**
	//
	// example:
	//
	// text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The content of the description for the check item when the Type parameter is text.
	//
	// example:
	//
	// The download of query results that are returned by SELECT statements in DataStudio must be prohibited at the MaxCompute level.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListCheckItemResponseBodyCheckItemsDescription) String() string {
	return dara.Prettify(s)
}

func (s ListCheckItemResponseBodyCheckItemsDescription) GoString() string {
	return s.String()
}

func (s *ListCheckItemResponseBodyCheckItemsDescription) GetType() *string {
	return s.Type
}

func (s *ListCheckItemResponseBodyCheckItemsDescription) GetValue() *string {
	return s.Value
}

func (s *ListCheckItemResponseBodyCheckItemsDescription) SetType(v string) *ListCheckItemResponseBodyCheckItemsDescription {
	s.Type = &v
	return s
}

func (s *ListCheckItemResponseBodyCheckItemsDescription) SetValue(v string) *ListCheckItemResponseBodyCheckItemsDescription {
	s.Value = &v
	return s
}

func (s *ListCheckItemResponseBodyCheckItemsDescription) Validate() error {
	return dara.Validate(s)
}

type ListCheckItemResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 149
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCheckItemResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListCheckItemResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListCheckItemResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *ListCheckItemResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListCheckItemResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCheckItemResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCheckItemResponseBodyPageInfo) SetCount(v int32) *ListCheckItemResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListCheckItemResponseBodyPageInfo) SetCurrentPage(v int32) *ListCheckItemResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListCheckItemResponseBodyPageInfo) SetPageSize(v int32) *ListCheckItemResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListCheckItemResponseBodyPageInfo) SetTotalCount(v int32) *ListCheckItemResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListCheckItemResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
