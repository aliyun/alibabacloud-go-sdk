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
	// The returned data.
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListProjectsResponseBodyData struct {
	// The token for retrieving the next page of results. If this parameter is empty, all results have been returned.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kvikyUl3ChyRxN+qLPvtOb
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// A pagination marker used to retrieve the next page of results. This parameter is returned when the response is truncated.
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
	if s.Projects != nil {
		for _, item := range s.Projects {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProjectsResponseBodyDataProjects struct {
	// The description of the project.
	//
	// example:
	//
	// BI_Analysis
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// The total storage usage of the project, which represents the compressed, logical data size used for metering.
	//
	// example:
	//
	// 16489027
	CostStorage *string `json:"costStorage,omitempty" xml:"costStorage,omitempty"`
	// The time when the project was created, as a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1704380838000
	CreatedTime *int64 `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// The default compute quota. If you do not specify a quota for a job, the job consumes computing resources from this default quota. For more information about how to use computing resources, see <props="intl">[Use of computing resources](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/use-of-computing-resources).
	//
	// example:
	//
	// os_PayAsYouGoQuota
	DefaultQuota *string `json:"defaultQuota,omitempty" xml:"defaultQuota,omitempty"`
	// The IP whitelist.
	IpWhiteList *ListProjectsResponseBodyDataProjectsIpWhiteList `json:"ipWhiteList,omitempty" xml:"ipWhiteList,omitempty" type:"Struct"`
	// The name of the project.
	//
	// example:
	//
	// odps_project
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The owner of the project.
	//
	// example:
	//
	// ALIYUN$odps****@aliyunid.com
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// The basic properties of the project.
	Properties *ListProjectsResponseBodyDataProjectsProperties `json:"properties,omitempty" xml:"properties,omitempty" type:"Struct"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The instance ID and billing method of the default compute quota.
	SaleTag *ListProjectsResponseBodyDataProjectsSaleTag `json:"saleTag,omitempty" xml:"saleTag,omitempty" type:"Struct"`
	// The security-related properties.
	SecurityProperties *ListProjectsResponseBodyDataProjectsSecurityProperties `json:"securityProperties,omitempty" xml:"securityProperties,omitempty" type:"Struct"`
	// The status of the project. Valid values:
	//
	// - **AVAILABLE**: The project is running as expected.
	//
	// - **READONLY**: The project is read-only.
	//
	// - **FROZEN**: The project is frozen.
	//
	// - **DELETING**: The project is being deleted.
	//
	// example:
	//
	// AVAILABLE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Specifies whether the project uses a three-tier model (project > schema > object). In this model, schemas are used within a project to organize objects such as tables, resources, and user-defined functions (UDFs). For more information, see <props="intl">[Schema operations](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/schema-related-operations).
	//
	// example:
	//
	// true
	ThreeTierModel *bool `json:"threeTierModel,omitempty" xml:"threeTierModel,omitempty"`
	// The type of the project. Valid values:
	//
	// - **managed**: An internal project.
	//
	// - **external**: An external project.
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
	if s.IpWhiteList != nil {
		if err := s.IpWhiteList.Validate(); err != nil {
			return err
		}
	}
	if s.Properties != nil {
		if err := s.Properties.Validate(); err != nil {
			return err
		}
	}
	if s.SaleTag != nil {
		if err := s.SaleTag.Validate(); err != nil {
			return err
		}
	}
	if s.SecurityProperties != nil {
		if err := s.SecurityProperties.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListProjectsResponseBodyDataProjectsIpWhiteList struct {
	// The IP whitelist for access over the public network and from other Alibaba Cloud services.
	//
	// > If you configure only this IP whitelist, access over the public network and from other Alibaba Cloud services is restricted based on the whitelist, and all access from VPCs is denied.
	//
	// example:
	//
	// 10.88.111.3
	IpList *string `json:"ipList,omitempty" xml:"ipList,omitempty"`
	// The IP whitelist for access from VPCs.
	//
	// > If you configure only the VPC IP whitelist, access from VPCs is restricted based on the whitelist, and all access over the public network and from other Alibaba Cloud services is denied.
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
	// Specifies whether to allow a full table scan in the project. This feature is disabled by default because a full table scan can consume a large amount of computing resources.
	//
	// example:
	//
	// false
	AllowFullScan *bool `json:"allowFullScan,omitempty" xml:"allowFullScan,omitempty"`
	// Specifies whether to enable the MaxCompute 2.0 Decimal data type in the project.
	//
	// example:
	//
	// true
	EnableDecimal2 *bool `json:"enableDecimal2,omitempty" xml:"enableDecimal2,omitempty"`
	// Specifies whether to enable routing for the Data Transmission Service resource group.
	//
	// - true: Data transmission jobs submitted by default from the project use the bound Data Transmission Service resource group.
	//
	// - false: Data transmission jobs submitted by default from the project use the shared Data Transmission Service resource group.
	//
	// example:
	//
	// true
	EnableTunnelQuotaRoute *bool `json:"enableTunnelQuotaRoute,omitempty" xml:"enableTunnelQuotaRoute,omitempty"`
	// The storage encryption properties.
	Encryption *ListProjectsResponseBodyDataProjectsPropertiesEncryption `json:"encryption,omitempty" xml:"encryption,omitempty" type:"Struct"`
	// The properties of the external project.
	ExternalProjectProperties *ListProjectsResponseBodyDataProjectsPropertiesExternalProjectProperties `json:"externalProjectProperties,omitempty" xml:"externalProjectProperties,omitempty" type:"Struct"`
	// The number of retention days for backup data. You can restore data to any backup version that is created within the retention period. Valid values: `0` to `30`. Default value: `1`. A value of `0` indicates that the backup feature is disabled.
	//
	// example:
	//
	// 1
	RetentionDays *int64 `json:"retentionDays,omitempty" xml:"retentionDays,omitempty"`
	// The maximum metered cost for a single SQL statement. The cost is calculated by using the formula: (scanned data in GB) × (complexity).
	//
	// example:
	//
	// 1500
	SqlMeteringMax *string `json:"sqlMeteringMax,omitempty" xml:"sqlMeteringMax,omitempty"`
	// The table lifecycle properties.
	TableLifecycle *ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle `json:"tableLifecycle,omitempty" xml:"tableLifecycle,omitempty" type:"Struct"`
	// The time zone of the project. This parameter corresponds to the `odps.sql.timezone` property.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// The <props="intl">[Data Transmission Service](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/overview-of-dts) resource group that is bound to the project.
	//
	// - Default (shared Data Transmission Service resource group): The project cannot use subscription Data Transmission Service resource groups. Data Transmission Service jobs submitted from this project automatically use the Default resource group, regardless of the default setting for the Data Transmission Service resource group.
	//
	// - Subscription Data Transmission Service resource group: The project can use a subscription Data Transmission Service resource group.
	//
	// example:
	//
	// Default
	TunnelQuota *string `json:"tunnelQuota,omitempty" xml:"tunnelQuota,omitempty"`
	// The data type version. Valid values:
	//
	// - **1**: Version 1.0.
	//
	// - **2**: Version 2.0.
	//
	// - **hive**: A Hive-compatible type.
	//
	// For more information about the differences between the data type versions, see <props="intl">[Data type versions](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/data-type-editions).
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
	if s.Encryption != nil {
		if err := s.Encryption.Validate(); err != nil {
			return err
		}
	}
	if s.ExternalProjectProperties != nil {
		if err := s.ExternalProjectProperties.Validate(); err != nil {
			return err
		}
	}
	if s.TableLifecycle != nil {
		if err := s.TableLifecycle.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListProjectsResponseBodyDataProjectsPropertiesEncryption struct {
	// The data encryption algorithm. Supported algorithms include AES256, AESCTR, and RC4.
	//
	// example:
	//
	// AES256
	Algorithm *string `json:"algorithm,omitempty" xml:"algorithm,omitempty"`
	// Specifies whether to enable storage encryption for the project. For more information, see
	//
	// <props="intl">[Storage encryption](https://www.alibabacloud.com/help/zh/maxcompute/security-and-compliance/storage-encryption).
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The key for data encryption. You can use the default MaxCompute-managed key or a custom key with the Bring Your Own Key (BYOK) feature.
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
	// Specifies whether the project is an external project for <props="intl">[Integrated Lakehouse](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/lake-warehouse-integrated-2-0-use-guide).
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
	// The type of the lifecycle. Valid values:
	//
	// - **mandatory**: A lifecycle must be configured for each table.
	//
	// - **optional**: The lifecycle is optional. If unspecified for a table, the table does not expire.
	//
	// - **inherit**: If no lifecycle is specified for a table, the table inherits its lifecycle from the `odps.table.lifecycle.value` property.
	//
	// example:
	//
	// optional
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The lifecycle of the table, in days. Valid values: `1` to `37231`. Default value: `37231`.
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
	// The instance ID of the default compute quota.
	//
	// example:
	//
	// b7afb7d1-****-****-****-c393669c307b
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The billing method of the default compute quota.
	//
	// example:
	//
	// PayAsYouGo
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
	// Specifies whether to enable <props="intl">[download control](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/label-based-access-control). Default value: `false`.
	//
	// example:
	//
	// false
	EnableDownloadPrivilege *bool `json:"enableDownloadPrivilege,omitempty" xml:"enableDownloadPrivilege,omitempty"`
	// Specifies whether to enable <props="intl">[label-based access control](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/label-based-access-control). Default value: `false`.
	//
	// example:
	//
	// false
	LabelSecurity *bool `json:"labelSecurity,omitempty" xml:"labelSecurity,omitempty"`
	// Specifies whether the creator of an object can access it. Default value: `true`.
	//
	// example:
	//
	// true
	ObjectCreatorHasAccessPermission *bool `json:"objectCreatorHasAccessPermission,omitempty" xml:"objectCreatorHasAccessPermission,omitempty"`
	// Specifies whether the creator of an object can grant other users permissions on it. Default value: `true`.
	//
	// example:
	//
	// true
	ObjectCreatorHasGrantPermission *bool `json:"objectCreatorHasGrantPermission,omitempty" xml:"objectCreatorHasGrantPermission,omitempty"`
	// The <props="intl">[project data protection](https://www.alibabacloud.com/help/zh/maxcompute/security-and-compliance/project-data-protection) properties.
	ProjectProtection *ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection `json:"projectProtection,omitempty" xml:"projectProtection,omitempty" type:"Struct"`
	// Specifies whether to enable <props="intl">[ACL-based access control](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/acl-based-access-control). Default value: `true`.
	//
	// example:
	//
	// true
	UsingAcl *bool `json:"usingAcl,omitempty" xml:"usingAcl,omitempty"`
	// Specifies whether to enable <props="intl">[policy-based access control](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/policy-based-access-control-1). Default value: `true`.
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
	if s.ProjectProtection != nil {
		if err := s.ProjectProtection.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection struct {
	// If project data protection is enabled, you can configure an exception policy. This policy allows specified users to export data from specified objects to trusted projects, bypassing the data protection mechanism.
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
	// Specifies whether to enable <props="intl">[project data protection](https://www.alibabacloud.com/help/zh/maxcompute/security-and-compliance/project-data-protection) to prevent data from being exported from the project. Default value: `false`.
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
