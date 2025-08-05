// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListProjectsResponseBodyData) *ListProjectsResponseBody
	GetData() *ListProjectsResponseBodyData
	SetRequestId(v string) *ListProjectsResponseBody
	GetRequestId() *string
}

type ListProjectsResponseBody struct {
	// The data returned.
	Data *ListProjectsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0b16399216671970335563173e2340
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListProjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBody) GetData() *ListProjectsResponseBodyData {
	return s.Data
}

func (s *ListProjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProjectsResponseBody) SetData(v *ListProjectsResponseBodyData) *ListProjectsResponseBody {
	s.Data = v
	return s
}

func (s *ListProjectsResponseBody) SetRequestId(v string) *ListProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListProjectsResponseBodyData struct {
	// A pagination token. Only continuous page turning is supported. If NextToken is not empty, the next page exists. The value of NextToken can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kvikyUl3ChyRxN+qLPvtOb
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Indicates the marker after which the returned list begins.
	//
	// example:
	//
	// cHlvZHBzX3VkZl8xMDExNV8xNDU3NDI4NDkzKg==
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 10
	MaxItem *int32 `json:"maxItem,omitempty" xml:"maxItem,omitempty"`
	// The list of projects.
	Projects []*ListProjectsResponseBodyDataProjects `json:"projects,omitempty" xml:"projects,omitempty" type:"Repeated"`
}

func (s ListProjectsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProjectsResponseBodyData) GetMarker() *string {
	return s.Marker
}

func (s *ListProjectsResponseBodyData) GetMaxItem() *int32 {
	return s.MaxItem
}

func (s *ListProjectsResponseBodyData) GetProjects() []*ListProjectsResponseBodyDataProjects {
	return s.Projects
}

func (s *ListProjectsResponseBodyData) SetNextToken(v string) *ListProjectsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListProjectsResponseBodyData) SetMarker(v string) *ListProjectsResponseBodyData {
	s.Marker = &v
	return s
}

func (s *ListProjectsResponseBodyData) SetMaxItem(v int32) *ListProjectsResponseBodyData {
	s.MaxItem = &v
	return s
}

func (s *ListProjectsResponseBodyData) SetProjects(v []*ListProjectsResponseBodyDataProjects) *ListProjectsResponseBodyData {
	s.Projects = v
	return s
}

func (s *ListProjectsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListProjectsResponseBodyDataProjects struct {
	// The project description.
	//
	// example:
	//
	// maxcompute projects
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// The total storage usage. The storage space that is occupied by your project, which is the logical storage space after your project data is collected and compressed.
	//
	// example:
	//
	// 16489027
	CostStorage *string `json:"costStorage,omitempty" xml:"costStorage,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1704380838000
	CreatedTime *int64 `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// The default computing quota that is used to allocate computing resources. If you do not specify a computing quota for your project, the jobs that are initiated by your project consume the computing resources in the default quota. For more information about how to use computing resources, see [Use quota groups for computing resources](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/use-of-computing-resources)
	//
	// example:
	//
	// quotaA
	DefaultQuota *string `json:"defaultQuota,omitempty" xml:"defaultQuota,omitempty"`
	// The information about the IP address whitelist.
	IpWhiteList *ListProjectsResponseBodyDataProjectsIpWhiteList `json:"ipWhiteList,omitempty" xml:"ipWhiteList,omitempty" type:"Struct"`
	// The name of the project.
	//
	// example:
	//
	// odps_project
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The account information of the project owner.
	//
	// example:
	//
	// 1139815775606813
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// The basic properties of the project.
	Properties *ListProjectsResponseBodyDataProjectsProperties `json:"properties,omitempty" xml:"properties,omitempty" type:"Struct"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The instance ID and billing method of the default computing quota.
	SaleTag *ListProjectsResponseBodyDataProjectsSaleTag `json:"saleTag,omitempty" xml:"saleTag,omitempty" type:"Struct"`
	// The permission properties.
	SecurityProperties *ListProjectsResponseBodyDataProjectsSecurityProperties `json:"securityProperties,omitempty" xml:"securityProperties,omitempty" type:"Struct"`
	// The project status. Valid values:
	//
	// 	- **AVAILABLE**
	//
	// 	- **READONLY**
	//
	// 	- **FROZEN**
	//
	// 	- **DELETING**
	//
	// example:
	//
	// AVAILABLE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Indicates whether data storage by schema is supported. MaxCompute supports the schema feature. This feature allows you to classify objects such as tables, resources, and user-defined functions (UDFs) in a project by schema. You can create multiple schemas in a project. For more information, see [Schema-related operations](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/schema-related-operations).
	//
	// Valid values:
	//
	// 	- true: supported
	//
	// 	- false: not supported
	//
	// example:
	//
	// true
	ThreeTierModel *bool `json:"threeTierModel,omitempty" xml:"threeTierModel,omitempty"`
	// The project type. Valid values:
	//
	// 	- **managed**: internal project
	//
	// 	- **external**: external project
	//
	// example:
	//
	// managed
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListProjectsResponseBodyDataProjects) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBodyDataProjects) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyDataProjects) GetComment() *string {
	return s.Comment
}

func (s *ListProjectsResponseBodyDataProjects) GetCostStorage() *string {
	return s.CostStorage
}

func (s *ListProjectsResponseBodyDataProjects) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *ListProjectsResponseBodyDataProjects) GetDefaultQuota() *string {
	return s.DefaultQuota
}

func (s *ListProjectsResponseBodyDataProjects) GetIpWhiteList() *ListProjectsResponseBodyDataProjectsIpWhiteList {
	return s.IpWhiteList
}

func (s *ListProjectsResponseBodyDataProjects) GetName() *string {
	return s.Name
}

func (s *ListProjectsResponseBodyDataProjects) GetOwner() *string {
	return s.Owner
}

func (s *ListProjectsResponseBodyDataProjects) GetProperties() *ListProjectsResponseBodyDataProjectsProperties {
	return s.Properties
}

func (s *ListProjectsResponseBodyDataProjects) GetRegionId() *string {
	return s.RegionId
}

func (s *ListProjectsResponseBodyDataProjects) GetSaleTag() *ListProjectsResponseBodyDataProjectsSaleTag {
	return s.SaleTag
}

func (s *ListProjectsResponseBodyDataProjects) GetSecurityProperties() *ListProjectsResponseBodyDataProjectsSecurityProperties {
	return s.SecurityProperties
}

func (s *ListProjectsResponseBodyDataProjects) GetStatus() *string {
	return s.Status
}

func (s *ListProjectsResponseBodyDataProjects) GetThreeTierModel() *bool {
	return s.ThreeTierModel
}

func (s *ListProjectsResponseBodyDataProjects) GetType() *string {
	return s.Type
}

func (s *ListProjectsResponseBodyDataProjects) SetComment(v string) *ListProjectsResponseBodyDataProjects {
	s.Comment = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetCostStorage(v string) *ListProjectsResponseBodyDataProjects {
	s.CostStorage = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetCreatedTime(v int64) *ListProjectsResponseBodyDataProjects {
	s.CreatedTime = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetDefaultQuota(v string) *ListProjectsResponseBodyDataProjects {
	s.DefaultQuota = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetIpWhiteList(v *ListProjectsResponseBodyDataProjectsIpWhiteList) *ListProjectsResponseBodyDataProjects {
	s.IpWhiteList = v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetName(v string) *ListProjectsResponseBodyDataProjects {
	s.Name = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetOwner(v string) *ListProjectsResponseBodyDataProjects {
	s.Owner = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetProperties(v *ListProjectsResponseBodyDataProjectsProperties) *ListProjectsResponseBodyDataProjects {
	s.Properties = v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetRegionId(v string) *ListProjectsResponseBodyDataProjects {
	s.RegionId = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetSaleTag(v *ListProjectsResponseBodyDataProjectsSaleTag) *ListProjectsResponseBodyDataProjects {
	s.SaleTag = v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetSecurityProperties(v *ListProjectsResponseBodyDataProjectsSecurityProperties) *ListProjectsResponseBodyDataProjects {
	s.SecurityProperties = v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetStatus(v string) *ListProjectsResponseBodyDataProjects {
	s.Status = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetThreeTierModel(v bool) *ListProjectsResponseBodyDataProjects {
	s.ThreeTierModel = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetType(v string) *ListProjectsResponseBodyDataProjects {
	s.Type = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) Validate() error {
	return dara.Validate(s)
}

type ListProjectsResponseBodyDataProjectsIpWhiteList struct {
	// The IP address whitelist for access over the Internet or the network for interconnecting with other Alibaba Cloud services.
	//
	// >  If you configure only the IP address whitelist for access over the Internet or the network for interconnecting with other Alibaba Cloud services, the access over the Internet or the network for interconnecting with other Alibaba Cloud services is subject to configurations, and access over a virtual private cloud (VPC) is not allowed.
	//
	// example:
	//
	// 10.88.111.3
	IpList *string `json:"ipList,omitempty" xml:"ipList,omitempty"`
	// The IP address whitelist for access over a VPC.
	//
	// >  If you configure only the IP address whitelist for access over a VPC, the access over a VPC is subject to configurations, and the access over the Internet or the network for interconnecting with other Alibaba Cloud services is not allowed.
	//
	// example:
	//
	// 10.88.111.3
	VpcIpList *string `json:"vpcIpList,omitempty" xml:"vpcIpList,omitempty"`
}

func (s ListProjectsResponseBodyDataProjectsIpWhiteList) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBodyDataProjectsIpWhiteList) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyDataProjectsIpWhiteList) GetIpList() *string {
	return s.IpList
}

func (s *ListProjectsResponseBodyDataProjectsIpWhiteList) GetVpcIpList() *string {
	return s.VpcIpList
}

func (s *ListProjectsResponseBodyDataProjectsIpWhiteList) SetIpList(v string) *ListProjectsResponseBodyDataProjectsIpWhiteList {
	s.IpList = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsIpWhiteList) SetVpcIpList(v string) *ListProjectsResponseBodyDataProjectsIpWhiteList {
	s.VpcIpList = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsIpWhiteList) Validate() error {
	return dara.Validate(s)
}

type ListProjectsResponseBodyDataProjectsProperties struct {
	// Indicates whether a full table scan is allowed in the project. A full table scan occupies a large number of resources, which reduces data processing efficiency. By default, the full table scan feature is disabled.
	//
	// example:
	//
	// false
	AllowFullScan *bool `json:"allowFullScan,omitempty" xml:"allowFullScan,omitempty"`
	// Indicates whether the DECIMAL type of the MaxCompute V2.0 data type edition is enabled.
	//
	// example:
	//
	// true
	EnableDecimal2 *bool `json:"enableDecimal2,omitempty" xml:"enableDecimal2,omitempty"`
	// Indicates whether the routing of the Tunnel resource group is enabled.
	//
	// 	- true: The data transfer tasks that are submitted by the project by default use the Tunnel resource group that is bound to the project.
	//
	// 	- false: The data transfer tasks that are submitted by the project by default use the Tunnel shared resource group.
	//
	// example:
	//
	// true
	EnableTunnelQuotaRoute *bool `json:"enableTunnelQuotaRoute,omitempty" xml:"enableTunnelQuotaRoute,omitempty"`
	// The storage encryption properties.
	Encryption *ListProjectsResponseBodyDataProjectsPropertiesEncryption `json:"encryption,omitempty" xml:"encryption,omitempty" type:"Struct"`
	// The properties of the external project.
	ExternalProjectProperties *ListProjectsResponseBodyDataProjectsPropertiesExternalProjectProperties `json:"externalProjectProperties,omitempty" xml:"externalProjectProperties,omitempty" type:"Struct"`
	// The retention period for backup data. Unit: days. During the retention period, you can restore data of the version in use to the backup data of any version. Valid values: [0,30]. Default value: 1. The value 0 indicates that the backup feature is disabled.
	//
	// example:
	//
	// 1
	RetentionDays *int64 `json:"retentionDays,omitempty" xml:"retentionDays,omitempty"`
	// The maximum consumption threshold of a single SQL statement. Formula: Amount of scanned data (GB) × Complexity.
	//
	// example:
	//
	// 1500
	SqlMeteringMax *string `json:"sqlMeteringMax,omitempty" xml:"sqlMeteringMax,omitempty"`
	// The table lifecycle properties.
	TableLifecycle *ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle `json:"tableLifecycle,omitempty" xml:"tableLifecycle,omitempty" type:"Struct"`
	// The time zone that is used by your project. The time zone is the same as the time zone specified by `odps.sql.timezone`.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// The [Tunnel](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/overview-of-dts) resource group that is bound to the project.
	//
	// 	- Default resource group: The Tunnel shared resource group is used. You cannot use the subscription-based Tunnel resource group for the project. The default resource group is automatically used by the Tunnel service of your project, regardless of the parameter setting.
	//
	// 	- Subscription-based Tunnel resource group: You can use the subscription-based Tunnel resource group for the project.
	//
	// example:
	//
	// quota_tunnel
	TunnelQuota *string `json:"tunnelQuota,omitempty" xml:"tunnelQuota,omitempty"`
	// The data type edition. Valid values:
	//
	// 	- **1**: MaxCompute V1.0 data type edition
	//
	// 	- **2**: MaxCompute V2.0 data type edition
	//
	// 	- **hive**: Hive-compatible data type edition
	//
	// For more information about the differences among the three data type editions, see [Data type editions](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/data-type-editions).
	//
	// example:
	//
	// 2
	TypeSystem *string `json:"typeSystem,omitempty" xml:"typeSystem,omitempty"`
}

func (s ListProjectsResponseBodyDataProjectsProperties) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBodyDataProjectsProperties) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyDataProjectsProperties) GetAllowFullScan() *bool {
	return s.AllowFullScan
}

func (s *ListProjectsResponseBodyDataProjectsProperties) GetEnableDecimal2() *bool {
	return s.EnableDecimal2
}

func (s *ListProjectsResponseBodyDataProjectsProperties) GetEnableTunnelQuotaRoute() *bool {
	return s.EnableTunnelQuotaRoute
}

func (s *ListProjectsResponseBodyDataProjectsProperties) GetEncryption() *ListProjectsResponseBodyDataProjectsPropertiesEncryption {
	return s.Encryption
}

func (s *ListProjectsResponseBodyDataProjectsProperties) GetExternalProjectProperties() *ListProjectsResponseBodyDataProjectsPropertiesExternalProjectProperties {
	return s.ExternalProjectProperties
}

func (s *ListProjectsResponseBodyDataProjectsProperties) GetRetentionDays() *int64 {
	return s.RetentionDays
}

func (s *ListProjectsResponseBodyDataProjectsProperties) GetSqlMeteringMax() *string {
	return s.SqlMeteringMax
}

func (s *ListProjectsResponseBodyDataProjectsProperties) GetTableLifecycle() *ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle {
	return s.TableLifecycle
}

func (s *ListProjectsResponseBodyDataProjectsProperties) GetTimezone() *string {
	return s.Timezone
}

func (s *ListProjectsResponseBodyDataProjectsProperties) GetTunnelQuota() *string {
	return s.TunnelQuota
}

func (s *ListProjectsResponseBodyDataProjectsProperties) GetTypeSystem() *string {
	return s.TypeSystem
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetAllowFullScan(v bool) *ListProjectsResponseBodyDataProjectsProperties {
	s.AllowFullScan = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetEnableDecimal2(v bool) *ListProjectsResponseBodyDataProjectsProperties {
	s.EnableDecimal2 = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetEnableTunnelQuotaRoute(v bool) *ListProjectsResponseBodyDataProjectsProperties {
	s.EnableTunnelQuotaRoute = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetEncryption(v *ListProjectsResponseBodyDataProjectsPropertiesEncryption) *ListProjectsResponseBodyDataProjectsProperties {
	s.Encryption = v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetExternalProjectProperties(v *ListProjectsResponseBodyDataProjectsPropertiesExternalProjectProperties) *ListProjectsResponseBodyDataProjectsProperties {
	s.ExternalProjectProperties = v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetRetentionDays(v int64) *ListProjectsResponseBodyDataProjectsProperties {
	s.RetentionDays = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetSqlMeteringMax(v string) *ListProjectsResponseBodyDataProjectsProperties {
	s.SqlMeteringMax = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetTableLifecycle(v *ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle) *ListProjectsResponseBodyDataProjectsProperties {
	s.TableLifecycle = v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetTimezone(v string) *ListProjectsResponseBodyDataProjectsProperties {
	s.Timezone = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetTunnelQuota(v string) *ListProjectsResponseBodyDataProjectsProperties {
	s.TunnelQuota = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetTypeSystem(v string) *ListProjectsResponseBodyDataProjectsProperties {
	s.TypeSystem = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) Validate() error {
	return dara.Validate(s)
}

type ListProjectsResponseBodyDataProjectsPropertiesEncryption struct {
	// The data encryption algorithm that is supported by the key. Valid values: AES256, AESCTR, and RC4.
	//
	// example:
	//
	// SHA1
	Algorithm *string `json:"algorithm,omitempty" xml:"algorithm,omitempty"`
	// Indicates whether the data encryption feature needs to be enabled for the project. For more information about data encryption, see
	//
	// [Storage encryption](https://www.alibabacloud.com/help/zh/maxcompute/security-and-compliance/storage-encryption).
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The type of key that is used for data encryption. You can select MaxCompute Default Key or Bring Your Own Key (BYOK) as the key type. If you select MaxCompute Default Key, the default key that is created by MaxCompute is used.
	//
	// example:
	//
	// dafault
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s ListProjectsResponseBodyDataProjectsPropertiesEncryption) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBodyDataProjectsPropertiesEncryption) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesEncryption) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesEncryption) GetEnable() *bool {
	return s.Enable
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesEncryption) GetKey() *string {
	return s.Key
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesEncryption) SetAlgorithm(v string) *ListProjectsResponseBodyDataProjectsPropertiesEncryption {
	s.Algorithm = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesEncryption) SetEnable(v bool) *ListProjectsResponseBodyDataProjectsPropertiesEncryption {
	s.Enable = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesEncryption) SetKey(v string) *ListProjectsResponseBodyDataProjectsPropertiesEncryption {
	s.Key = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesEncryption) Validate() error {
	return dara.Validate(s)
}

type ListProjectsResponseBodyDataProjectsPropertiesExternalProjectProperties struct {
	// Indicates whether the external project is an external project for [data lakehouse solution 2.0](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/lake-warehouse-integrated-2-0-use-guide).
	//
	// example:
	//
	// true
	IsExternalCatalogBound *string `json:"isExternalCatalogBound,omitempty" xml:"isExternalCatalogBound,omitempty"`
}

func (s ListProjectsResponseBodyDataProjectsPropertiesExternalProjectProperties) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBodyDataProjectsPropertiesExternalProjectProperties) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesExternalProjectProperties) GetIsExternalCatalogBound() *string {
	return s.IsExternalCatalogBound
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesExternalProjectProperties) SetIsExternalCatalogBound(v string) *ListProjectsResponseBodyDataProjectsPropertiesExternalProjectProperties {
	s.IsExternalCatalogBound = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesExternalProjectProperties) Validate() error {
	return dara.Validate(s)
}

type ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle struct {
	// The lifecycle type. Valid values:
	//
	// 	- **mandatory**: The lifecycle clause is required in a table creation statement.
	//
	// 	- **optional**: The lifecycle clause is optional in a table creation statement. If you do not configure a lifecycle for a table, the table does not expire.
	//
	// 	- **inherit**: If you do not configure a lifecycle for a table when you create the table, the value of the odps.table.lifecycle.value parameter is used as the table lifecycle by default.
	//
	// example:
	//
	// optional
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The table lifecycle. Unit: days. Valid values: 1 to 37231. Default value: 37231.
	//
	// example:
	//
	// 37231
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle) GetType() *string {
	return s.Type
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle) GetValue() *string {
	return s.Value
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle) SetType(v string) *ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle {
	s.Type = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle) SetValue(v string) *ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle {
	s.Value = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle) Validate() error {
	return dara.Validate(s)
}

type ListProjectsResponseBodyDataProjectsSaleTag struct {
	// The instance ID of the default computing quota.
	//
	// example:
	//
	// "aaaa-bbbb"
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The billing method of the default computing quota.
	//
	// example:
	//
	// "project"
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ListProjectsResponseBodyDataProjectsSaleTag) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBodyDataProjectsSaleTag) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyDataProjectsSaleTag) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListProjectsResponseBodyDataProjectsSaleTag) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListProjectsResponseBodyDataProjectsSaleTag) SetResourceId(v string) *ListProjectsResponseBodyDataProjectsSaleTag {
	s.ResourceId = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsSaleTag) SetResourceType(v string) *ListProjectsResponseBodyDataProjectsSaleTag {
	s.ResourceType = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsSaleTag) Validate() error {
	return dara.Validate(s)
}

type ListProjectsResponseBodyDataProjectsSecurityProperties struct {
	// Indicates whether the [download control](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/label-based-access-control) feature is enabled. By default, this feature is disabled.
	//
	// example:
	//
	// false
	EnableDownloadPrivilege *bool `json:"enableDownloadPrivilege,omitempty" xml:"enableDownloadPrivilege,omitempty"`
	// Indicates whether the [label-based access control](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/label-based-access-control) feature is enabled. By default, this feature is disabled.
	//
	// example:
	//
	// false
	LabelSecurity *bool `json:"labelSecurity,omitempty" xml:"labelSecurity,omitempty"`
	// Indicates whether to allow the object creator to have the access permissions on the object. The default value is true, which indicates that the object creator has the access permissions on the object.
	//
	// example:
	//
	// true
	ObjectCreatorHasAccessPermission *bool `json:"objectCreatorHasAccessPermission,omitempty" xml:"objectCreatorHasAccessPermission,omitempty"`
	// Indicates whether the object creator has the authorization permissions on the object. The default value is true, which indicates that the object creator has the authorization permissions on the object.
	//
	// example:
	//
	// true
	ObjectCreatorHasGrantPermission *bool `json:"objectCreatorHasGrantPermission,omitempty" xml:"objectCreatorHasGrantPermission,omitempty"`
	// The properties of the [data protection mechanism](https://www.alibabacloud.com/help/zh/maxcompute/security-and-compliance/project-data-protection).
	ProjectProtection *ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection `json:"projectProtection,omitempty" xml:"projectProtection,omitempty" type:"Struct"`
	// Indicates whether the [ACL-based access control](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/acl-based-access-control) feature is enabled. By default, this feature is enabled.
	//
	// example:
	//
	// true
	UsingAcl *bool `json:"usingAcl,omitempty" xml:"usingAcl,omitempty"`
	// Indicates whether the [policy-based access control](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/policy-based-access-control-1) feature is enabled. By default, this feature is enabled.
	//
	// example:
	//
	// true
	UsingPolicy *bool `json:"usingPolicy,omitempty" xml:"usingPolicy,omitempty"`
}

func (s ListProjectsResponseBodyDataProjectsSecurityProperties) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBodyDataProjectsSecurityProperties) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) GetEnableDownloadPrivilege() *bool {
	return s.EnableDownloadPrivilege
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) GetLabelSecurity() *bool {
	return s.LabelSecurity
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) GetObjectCreatorHasAccessPermission() *bool {
	return s.ObjectCreatorHasAccessPermission
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) GetObjectCreatorHasGrantPermission() *bool {
	return s.ObjectCreatorHasGrantPermission
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) GetProjectProtection() *ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection {
	return s.ProjectProtection
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) GetUsingAcl() *bool {
	return s.UsingAcl
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) GetUsingPolicy() *bool {
	return s.UsingPolicy
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) SetEnableDownloadPrivilege(v bool) *ListProjectsResponseBodyDataProjectsSecurityProperties {
	s.EnableDownloadPrivilege = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) SetLabelSecurity(v bool) *ListProjectsResponseBodyDataProjectsSecurityProperties {
	s.LabelSecurity = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) SetObjectCreatorHasAccessPermission(v bool) *ListProjectsResponseBodyDataProjectsSecurityProperties {
	s.ObjectCreatorHasAccessPermission = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) SetObjectCreatorHasGrantPermission(v bool) *ListProjectsResponseBodyDataProjectsSecurityProperties {
	s.ObjectCreatorHasGrantPermission = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) SetProjectProtection(v *ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection) *ListProjectsResponseBodyDataProjectsSecurityProperties {
	s.ProjectProtection = v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) SetUsingAcl(v bool) *ListProjectsResponseBodyDataProjectsSecurityProperties {
	s.UsingAcl = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) SetUsingPolicy(v bool) *ListProjectsResponseBodyDataProjectsSecurityProperties {
	s.UsingPolicy = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) Validate() error {
	return dara.Validate(s)
}

type ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection struct {
	// If you enable the project data protection mechanism, you can configure exception or trusted projects. This allows specified users to transfer data of a specified object to a specified project. The project data protection mechanism does not take effect in all the situations that are specified in the exception policy.
	//
	// example:
	//
	// {
	//
	//       "Version": "1",
	//
	//       "Statement": [
	//
	//             {
	//
	//                   "Effect": "Allow",
	//
	//                   "Principal": "",
	//
	//                   "Action": [
	//
	//                         "odps:[, , ...]"
	//
	//                   ],
	//
	//                   "Resource": "acs:odps:*:",
	//
	//                   "Condition": {
	//
	//                         "StringEquals": {
	//
	//                               "odps:TaskType": [
	//
	//                                     ""
	//
	//                               ]
	//
	//                         }
	//
	//                   }
	//
	//             }
	//
	//       ]
	//
	// }
	ExceptionPolicy *string `json:"exceptionPolicy,omitempty" xml:"exceptionPolicy,omitempty"`
	// Indicates whether the [data protection mechanism](https://www.alibabacloud.com/help/zh/maxcompute/security-and-compliance/project-data-protection) is enabled for the project. This allows or denies data transfer across projects. By default, the data protection mechanism is disabled.
	//
	// example:
	//
	// true
	Protected *bool `json:"protected,omitempty" xml:"protected,omitempty"`
}

func (s ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection) GetExceptionPolicy() *string {
	return s.ExceptionPolicy
}

func (s *ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection) GetProtected() *bool {
	return s.Protected
}

func (s *ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection) SetExceptionPolicy(v string) *ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection {
	s.ExceptionPolicy = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection) SetProtected(v bool) *ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection {
	s.Protected = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection) Validate() error {
	return dara.Validate(s)
}
