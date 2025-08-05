// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetProjectResponseBodyData) *GetProjectResponseBody
	GetData() *GetProjectResponseBodyData
	SetErrorCode(v string) *GetProjectResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GetProjectResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *GetProjectResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetProjectResponseBody
	GetRequestId() *string
}

type GetProjectResponseBody struct {
	// The data returned.
	Data *GetProjectResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// 040002
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// error message.
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// The HTTP status code.
	//
	// 	- 1xx: informational response. The request is received and is being processed.
	//
	// 	- 2xx: success. The request is successfully received, understood, and accepted by the server.
	//
	// 	- 3xx: redirection. The request is redirected, and further actions are required to complete the request.
	//
	// 	- 4xx: client error. The request contains invalid request parameters and syntaxes, or specific request conditions cannot be met.
	//
	// 	- 5xx: server error. The server cannot meet requirements due to other reasons.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0b87b7b316643495896551555e855b
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProjectResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBody) GetData() *GetProjectResponseBodyData {
	return s.Data
}

func (s *GetProjectResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetProjectResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetProjectResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProjectResponseBody) SetData(v *GetProjectResponseBodyData) *GetProjectResponseBody {
	s.Data = v
	return s
}

func (s *GetProjectResponseBody) SetErrorCode(v string) *GetProjectResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetProjectResponseBody) SetErrorMsg(v string) *GetProjectResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetProjectResponseBody) SetHttpCode(v int32) *GetProjectResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetProjectResponseBody) SetRequestId(v string) *GetProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetProjectResponseBodyData struct {
	// The project description.
	//
	// example:
	//
	// maxcompute project
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
	// The default computing quota that is used to allocate computing resources. If you do not specify a computing quota for your project, the jobs that are initiated by your project consume the computing resources in the default quota. For more information about how to use computing resources, see [Use quota groups for computing resources](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/use-of-computing-resources).
	//
	// example:
	//
	// quota_a
	DefaultQuota *string `json:"defaultQuota,omitempty" xml:"defaultQuota,omitempty"`
	// The information about the IP address whitelist.
	IpWhiteList *GetProjectResponseBodyDataIpWhiteList `json:"ipWhiteList,omitempty" xml:"ipWhiteList,omitempty" type:"Struct"`
	// The project name.
	//
	// example:
	//
	// odps_project
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The account information of the project owner.
	//
	// example:
	//
	// 1565950907343451
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// The billing method of the default computing quota.
	//
	// example:
	//
	// PayAsYouGo
	ProductType *string `json:"productType,omitempty" xml:"productType,omitempty"`
	// The basic properties of the project.
	Properties *GetProjectResponseBodyDataProperties `json:"properties,omitempty" xml:"properties,omitempty" type:"Struct"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The instance ID and billing method of the default computing quota.
	SaleTag *GetProjectResponseBodyDataSaleTag `json:"saleTag,omitempty" xml:"saleTag,omitempty" type:"Struct"`
	// The permission properties.
	SecurityProperties *GetProjectResponseBodyDataSecurityProperties `json:"securityProperties,omitempty" xml:"securityProperties,omitempty" type:"Struct"`
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
	// The list of `Super_Administrator` role members of the project.
	SuperAdmins []*string `json:"superAdmins,omitempty" xml:"superAdmins,omitempty" type:"Repeated"`
	// Indicates whether data storage by schema is supported. MaxCompute supports the schema feature. This feature allows you to classify objects such as tables, resources, and user-defined functions (UDFs) in a project by schema. You can create multiple schemas in a project. For more information, see [Schema-related operations](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/schema-related-operations).
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

func (s GetProjectResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyData) GetComment() *string {
	return s.Comment
}

func (s *GetProjectResponseBodyData) GetCostStorage() *string {
	return s.CostStorage
}

func (s *GetProjectResponseBodyData) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *GetProjectResponseBodyData) GetDefaultQuota() *string {
	return s.DefaultQuota
}

func (s *GetProjectResponseBodyData) GetIpWhiteList() *GetProjectResponseBodyDataIpWhiteList {
	return s.IpWhiteList
}

func (s *GetProjectResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetProjectResponseBodyData) GetOwner() *string {
	return s.Owner
}

func (s *GetProjectResponseBodyData) GetProductType() *string {
	return s.ProductType
}

func (s *GetProjectResponseBodyData) GetProperties() *GetProjectResponseBodyDataProperties {
	return s.Properties
}

func (s *GetProjectResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetProjectResponseBodyData) GetSaleTag() *GetProjectResponseBodyDataSaleTag {
	return s.SaleTag
}

func (s *GetProjectResponseBodyData) GetSecurityProperties() *GetProjectResponseBodyDataSecurityProperties {
	return s.SecurityProperties
}

func (s *GetProjectResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetProjectResponseBodyData) GetSuperAdmins() []*string {
	return s.SuperAdmins
}

func (s *GetProjectResponseBodyData) GetThreeTierModel() *bool {
	return s.ThreeTierModel
}

func (s *GetProjectResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetProjectResponseBodyData) SetComment(v string) *GetProjectResponseBodyData {
	s.Comment = &v
	return s
}

func (s *GetProjectResponseBodyData) SetCostStorage(v string) *GetProjectResponseBodyData {
	s.CostStorage = &v
	return s
}

func (s *GetProjectResponseBodyData) SetCreatedTime(v int64) *GetProjectResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *GetProjectResponseBodyData) SetDefaultQuota(v string) *GetProjectResponseBodyData {
	s.DefaultQuota = &v
	return s
}

func (s *GetProjectResponseBodyData) SetIpWhiteList(v *GetProjectResponseBodyDataIpWhiteList) *GetProjectResponseBodyData {
	s.IpWhiteList = v
	return s
}

func (s *GetProjectResponseBodyData) SetName(v string) *GetProjectResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetProjectResponseBodyData) SetOwner(v string) *GetProjectResponseBodyData {
	s.Owner = &v
	return s
}

func (s *GetProjectResponseBodyData) SetProductType(v string) *GetProjectResponseBodyData {
	s.ProductType = &v
	return s
}

func (s *GetProjectResponseBodyData) SetProperties(v *GetProjectResponseBodyDataProperties) *GetProjectResponseBodyData {
	s.Properties = v
	return s
}

func (s *GetProjectResponseBodyData) SetRegionId(v string) *GetProjectResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetProjectResponseBodyData) SetSaleTag(v *GetProjectResponseBodyDataSaleTag) *GetProjectResponseBodyData {
	s.SaleTag = v
	return s
}

func (s *GetProjectResponseBodyData) SetSecurityProperties(v *GetProjectResponseBodyDataSecurityProperties) *GetProjectResponseBodyData {
	s.SecurityProperties = v
	return s
}

func (s *GetProjectResponseBodyData) SetStatus(v string) *GetProjectResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetProjectResponseBodyData) SetSuperAdmins(v []*string) *GetProjectResponseBodyData {
	s.SuperAdmins = v
	return s
}

func (s *GetProjectResponseBodyData) SetThreeTierModel(v bool) *GetProjectResponseBodyData {
	s.ThreeTierModel = &v
	return s
}

func (s *GetProjectResponseBodyData) SetType(v string) *GetProjectResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetProjectResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetProjectResponseBodyDataIpWhiteList struct {
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

func (s GetProjectResponseBodyDataIpWhiteList) String() string {
	return dara.Prettify(s)
}

func (s GetProjectResponseBodyDataIpWhiteList) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyDataIpWhiteList) GetIpList() *string {
	return s.IpList
}

func (s *GetProjectResponseBodyDataIpWhiteList) GetVpcIpList() *string {
	return s.VpcIpList
}

func (s *GetProjectResponseBodyDataIpWhiteList) SetIpList(v string) *GetProjectResponseBodyDataIpWhiteList {
	s.IpList = &v
	return s
}

func (s *GetProjectResponseBodyDataIpWhiteList) SetVpcIpList(v string) *GetProjectResponseBodyDataIpWhiteList {
	s.VpcIpList = &v
	return s
}

func (s *GetProjectResponseBodyDataIpWhiteList) Validate() error {
	return dara.Validate(s)
}

type GetProjectResponseBodyDataProperties struct {
	// Indicates whether a full table scan is allowed in the project. A full table scan occupies a large number of resources, which reduces data processing efficiency. By default, the full table scan feature is disabled.
	//
	// example:
	//
	// false
	AllowFullScan *bool  `json:"allowFullScan,omitempty" xml:"allowFullScan,omitempty"`
	AutoMvQuotaGb *int64 `json:"autoMvQuotaGb,omitempty" xml:"autoMvQuotaGb,omitempty"`
	// The Tunnel parent resource group that is bound to the project. You do not need to pay attention to this group.
	//
	// example:
	//
	// No value
	ElderTunnelQuota *string `json:"elderTunnelQuota,omitempty" xml:"elderTunnelQuota,omitempty"`
	EnableAutoMv     *bool   `json:"enableAutoMv,omitempty" xml:"enableAutoMv,omitempty"`
	// Indicates whether the DECIMAL type of the MaxCompute V2.0 data type edition is enabled.
	//
	// example:
	//
	// true
	EnableDecimal2 *bool `json:"enableDecimal2,omitempty" xml:"enableDecimal2,omitempty"`
	EnableDr       *bool `json:"enableDr,omitempty" xml:"enableDr,omitempty"`
	// Indicates whether external table caching is forcefully enabled.
	//
	// example:
	//
	// true
	EnableFdcCacheForce *bool `json:"enableFdcCacheForce,omitempty" xml:"enableFdcCacheForce,omitempty"`
	// Indicates whether [tiered storage](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/tiered-storage) is enabled.
	//
	// example:
	//
	// true
	EnableTieredStorage *bool `json:"enableTieredStorage,omitempty" xml:"enableTieredStorage,omitempty"`
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
	Encryption *GetProjectResponseBodyDataPropertiesEncryption `json:"encryption,omitempty" xml:"encryption,omitempty" type:"Struct"`
	// The properties of the external project.
	ExternalProjectProperties *GetProjectResponseBodyDataPropertiesExternalProjectProperties `json:"externalProjectProperties,omitempty" xml:"externalProjectProperties,omitempty" type:"Struct"`
	// The quota for external table caching.
	//
	// example:
	//
	// fdc_quota
	FdcQuota *string `json:"fdcQuota,omitempty" xml:"fdcQuota,omitempty"`
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
	// The [storage tier](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/tiered-storage) information.
	StorageTierInfo *GetProjectResponseBodyDataPropertiesStorageTierInfo `json:"storageTierInfo,omitempty" xml:"storageTierInfo,omitempty" type:"Struct"`
	// The table lifecycle properties.
	TableLifecycle *GetProjectResponseBodyDataPropertiesTableLifecycle `json:"tableLifecycle,omitempty" xml:"tableLifecycle,omitempty" type:"Struct"`
	// The [properties of tiered storage lifecycle rules](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/tiered-storage#f61fc9db76nna). After you configure the properties, the system triggers automatic switching of storage tiers based on the rules.
	TableLifecycleConfig *GetProjectResponseBodyDataPropertiesTableLifecycleConfig `json:"tableLifecycleConfig,omitempty" xml:"tableLifecycleConfig,omitempty" type:"Struct"`
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
	// Quota
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
	// 2.0
	TypeSystem *string `json:"typeSystem,omitempty" xml:"typeSystem,omitempty"`
}

func (s GetProjectResponseBodyDataProperties) String() string {
	return dara.Prettify(s)
}

func (s GetProjectResponseBodyDataProperties) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyDataProperties) GetAllowFullScan() *bool {
	return s.AllowFullScan
}

func (s *GetProjectResponseBodyDataProperties) GetAutoMvQuotaGb() *int64 {
	return s.AutoMvQuotaGb
}

func (s *GetProjectResponseBodyDataProperties) GetElderTunnelQuota() *string {
	return s.ElderTunnelQuota
}

func (s *GetProjectResponseBodyDataProperties) GetEnableAutoMv() *bool {
	return s.EnableAutoMv
}

func (s *GetProjectResponseBodyDataProperties) GetEnableDecimal2() *bool {
	return s.EnableDecimal2
}

func (s *GetProjectResponseBodyDataProperties) GetEnableDr() *bool {
	return s.EnableDr
}

func (s *GetProjectResponseBodyDataProperties) GetEnableFdcCacheForce() *bool {
	return s.EnableFdcCacheForce
}

func (s *GetProjectResponseBodyDataProperties) GetEnableTieredStorage() *bool {
	return s.EnableTieredStorage
}

func (s *GetProjectResponseBodyDataProperties) GetEnableTunnelQuotaRoute() *bool {
	return s.EnableTunnelQuotaRoute
}

func (s *GetProjectResponseBodyDataProperties) GetEncryption() *GetProjectResponseBodyDataPropertiesEncryption {
	return s.Encryption
}

func (s *GetProjectResponseBodyDataProperties) GetExternalProjectProperties() *GetProjectResponseBodyDataPropertiesExternalProjectProperties {
	return s.ExternalProjectProperties
}

func (s *GetProjectResponseBodyDataProperties) GetFdcQuota() *string {
	return s.FdcQuota
}

func (s *GetProjectResponseBodyDataProperties) GetRetentionDays() *int64 {
	return s.RetentionDays
}

func (s *GetProjectResponseBodyDataProperties) GetSqlMeteringMax() *string {
	return s.SqlMeteringMax
}

func (s *GetProjectResponseBodyDataProperties) GetStorageTierInfo() *GetProjectResponseBodyDataPropertiesStorageTierInfo {
	return s.StorageTierInfo
}

func (s *GetProjectResponseBodyDataProperties) GetTableLifecycle() *GetProjectResponseBodyDataPropertiesTableLifecycle {
	return s.TableLifecycle
}

func (s *GetProjectResponseBodyDataProperties) GetTableLifecycleConfig() *GetProjectResponseBodyDataPropertiesTableLifecycleConfig {
	return s.TableLifecycleConfig
}

func (s *GetProjectResponseBodyDataProperties) GetTimezone() *string {
	return s.Timezone
}

func (s *GetProjectResponseBodyDataProperties) GetTunnelQuota() *string {
	return s.TunnelQuota
}

func (s *GetProjectResponseBodyDataProperties) GetTypeSystem() *string {
	return s.TypeSystem
}

func (s *GetProjectResponseBodyDataProperties) SetAllowFullScan(v bool) *GetProjectResponseBodyDataProperties {
	s.AllowFullScan = &v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetAutoMvQuotaGb(v int64) *GetProjectResponseBodyDataProperties {
	s.AutoMvQuotaGb = &v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetElderTunnelQuota(v string) *GetProjectResponseBodyDataProperties {
	s.ElderTunnelQuota = &v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetEnableAutoMv(v bool) *GetProjectResponseBodyDataProperties {
	s.EnableAutoMv = &v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetEnableDecimal2(v bool) *GetProjectResponseBodyDataProperties {
	s.EnableDecimal2 = &v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetEnableDr(v bool) *GetProjectResponseBodyDataProperties {
	s.EnableDr = &v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetEnableFdcCacheForce(v bool) *GetProjectResponseBodyDataProperties {
	s.EnableFdcCacheForce = &v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetEnableTieredStorage(v bool) *GetProjectResponseBodyDataProperties {
	s.EnableTieredStorage = &v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetEnableTunnelQuotaRoute(v bool) *GetProjectResponseBodyDataProperties {
	s.EnableTunnelQuotaRoute = &v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetEncryption(v *GetProjectResponseBodyDataPropertiesEncryption) *GetProjectResponseBodyDataProperties {
	s.Encryption = v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetExternalProjectProperties(v *GetProjectResponseBodyDataPropertiesExternalProjectProperties) *GetProjectResponseBodyDataProperties {
	s.ExternalProjectProperties = v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetFdcQuota(v string) *GetProjectResponseBodyDataProperties {
	s.FdcQuota = &v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetRetentionDays(v int64) *GetProjectResponseBodyDataProperties {
	s.RetentionDays = &v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetSqlMeteringMax(v string) *GetProjectResponseBodyDataProperties {
	s.SqlMeteringMax = &v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetStorageTierInfo(v *GetProjectResponseBodyDataPropertiesStorageTierInfo) *GetProjectResponseBodyDataProperties {
	s.StorageTierInfo = v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetTableLifecycle(v *GetProjectResponseBodyDataPropertiesTableLifecycle) *GetProjectResponseBodyDataProperties {
	s.TableLifecycle = v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetTableLifecycleConfig(v *GetProjectResponseBodyDataPropertiesTableLifecycleConfig) *GetProjectResponseBodyDataProperties {
	s.TableLifecycleConfig = v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetTimezone(v string) *GetProjectResponseBodyDataProperties {
	s.Timezone = &v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetTunnelQuota(v string) *GetProjectResponseBodyDataProperties {
	s.TunnelQuota = &v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetTypeSystem(v string) *GetProjectResponseBodyDataProperties {
	s.TypeSystem = &v
	return s
}

func (s *GetProjectResponseBodyDataProperties) Validate() error {
	return dara.Validate(s)
}

type GetProjectResponseBodyDataPropertiesEncryption struct {
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

func (s GetProjectResponseBodyDataPropertiesEncryption) String() string {
	return dara.Prettify(s)
}

func (s GetProjectResponseBodyDataPropertiesEncryption) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyDataPropertiesEncryption) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *GetProjectResponseBodyDataPropertiesEncryption) GetEnable() *bool {
	return s.Enable
}

func (s *GetProjectResponseBodyDataPropertiesEncryption) GetKey() *string {
	return s.Key
}

func (s *GetProjectResponseBodyDataPropertiesEncryption) SetAlgorithm(v string) *GetProjectResponseBodyDataPropertiesEncryption {
	s.Algorithm = &v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesEncryption) SetEnable(v bool) *GetProjectResponseBodyDataPropertiesEncryption {
	s.Enable = &v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesEncryption) SetKey(v string) *GetProjectResponseBodyDataPropertiesEncryption {
	s.Key = &v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesEncryption) Validate() error {
	return dara.Validate(s)
}

type GetProjectResponseBodyDataPropertiesExternalProjectProperties struct {
	ExternalCatalogId *string `json:"externalCatalogId,omitempty" xml:"externalCatalogId,omitempty"`
	ForeignServerName *string `json:"foreignServerName,omitempty" xml:"foreignServerName,omitempty"`
	ForeignServerType *string `json:"foreignServerType,omitempty" xml:"foreignServerType,omitempty"`
	// Indicates whether the external project is an external project for [data lakehouse solution 2.0](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/lake-warehouse-integrated-2-0-use-guide).
	//
	// example:
	//
	// true
	IsExternalCatalogBound *string `json:"isExternalCatalogBound,omitempty" xml:"isExternalCatalogBound,omitempty"`
	TableFormat            *string `json:"tableFormat,omitempty" xml:"tableFormat,omitempty"`
	Warehouse              *string `json:"warehouse,omitempty" xml:"warehouse,omitempty"`
}

func (s GetProjectResponseBodyDataPropertiesExternalProjectProperties) String() string {
	return dara.Prettify(s)
}

func (s GetProjectResponseBodyDataPropertiesExternalProjectProperties) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyDataPropertiesExternalProjectProperties) GetExternalCatalogId() *string {
	return s.ExternalCatalogId
}

func (s *GetProjectResponseBodyDataPropertiesExternalProjectProperties) GetForeignServerName() *string {
	return s.ForeignServerName
}

func (s *GetProjectResponseBodyDataPropertiesExternalProjectProperties) GetForeignServerType() *string {
	return s.ForeignServerType
}

func (s *GetProjectResponseBodyDataPropertiesExternalProjectProperties) GetIsExternalCatalogBound() *string {
	return s.IsExternalCatalogBound
}

func (s *GetProjectResponseBodyDataPropertiesExternalProjectProperties) GetTableFormat() *string {
	return s.TableFormat
}

func (s *GetProjectResponseBodyDataPropertiesExternalProjectProperties) GetWarehouse() *string {
	return s.Warehouse
}

func (s *GetProjectResponseBodyDataPropertiesExternalProjectProperties) SetExternalCatalogId(v string) *GetProjectResponseBodyDataPropertiesExternalProjectProperties {
	s.ExternalCatalogId = &v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesExternalProjectProperties) SetForeignServerName(v string) *GetProjectResponseBodyDataPropertiesExternalProjectProperties {
	s.ForeignServerName = &v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesExternalProjectProperties) SetForeignServerType(v string) *GetProjectResponseBodyDataPropertiesExternalProjectProperties {
	s.ForeignServerType = &v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesExternalProjectProperties) SetIsExternalCatalogBound(v string) *GetProjectResponseBodyDataPropertiesExternalProjectProperties {
	s.IsExternalCatalogBound = &v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesExternalProjectProperties) SetTableFormat(v string) *GetProjectResponseBodyDataPropertiesExternalProjectProperties {
	s.TableFormat = &v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesExternalProjectProperties) SetWarehouse(v string) *GetProjectResponseBodyDataPropertiesExternalProjectProperties {
	s.Warehouse = &v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesExternalProjectProperties) Validate() error {
	return dara.Validate(s)
}

type GetProjectResponseBodyDataPropertiesStorageTierInfo struct {
	// The backup storage usage.
	//
	// example:
	//
	// 86672917
	ProjectBackupSize *int64 `json:"projectBackupSize,omitempty" xml:"projectBackupSize,omitempty"`
	// The total storage usage.
	//
	// example:
	//
	// 56066037
	ProjectTotalSize *int64 `json:"projectTotalSize,omitempty" xml:"projectTotalSize,omitempty"`
	// The [storage tier](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/tiered-storage) information.
	StorageTierSize *GetProjectResponseBodyDataPropertiesStorageTierInfoStorageTierSize `json:"storageTierSize,omitempty" xml:"storageTierSize,omitempty" type:"Struct"`
}

func (s GetProjectResponseBodyDataPropertiesStorageTierInfo) String() string {
	return dara.Prettify(s)
}

func (s GetProjectResponseBodyDataPropertiesStorageTierInfo) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyDataPropertiesStorageTierInfo) GetProjectBackupSize() *int64 {
	return s.ProjectBackupSize
}

func (s *GetProjectResponseBodyDataPropertiesStorageTierInfo) GetProjectTotalSize() *int64 {
	return s.ProjectTotalSize
}

func (s *GetProjectResponseBodyDataPropertiesStorageTierInfo) GetStorageTierSize() *GetProjectResponseBodyDataPropertiesStorageTierInfoStorageTierSize {
	return s.StorageTierSize
}

func (s *GetProjectResponseBodyDataPropertiesStorageTierInfo) SetProjectBackupSize(v int64) *GetProjectResponseBodyDataPropertiesStorageTierInfo {
	s.ProjectBackupSize = &v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesStorageTierInfo) SetProjectTotalSize(v int64) *GetProjectResponseBodyDataPropertiesStorageTierInfo {
	s.ProjectTotalSize = &v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesStorageTierInfo) SetStorageTierSize(v *GetProjectResponseBodyDataPropertiesStorageTierInfoStorageTierSize) *GetProjectResponseBodyDataPropertiesStorageTierInfo {
	s.StorageTierSize = v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesStorageTierInfo) Validate() error {
	return dara.Validate(s)
}

type GetProjectResponseBodyDataPropertiesStorageTierInfoStorageTierSize struct {
	// The storage usage at the long-term storage tier.
	//
	// example:
	//
	// 21764917
	LongTermSize *int64 `json:"longTermSize,omitempty" xml:"longTermSize,omitempty"`
	// The storage usage at the Infrequent Access (IA) layer.
	//
	// example:
	//
	// 767693
	LowFrequencySize *int64 `json:"lowFrequencySize,omitempty" xml:"lowFrequencySize,omitempty"`
	// The storage usage at the standard storage tier.
	//
	// example:
	//
	// 27649172
	StandardSize *int64 `json:"standardSize,omitempty" xml:"standardSize,omitempty"`
}

func (s GetProjectResponseBodyDataPropertiesStorageTierInfoStorageTierSize) String() string {
	return dara.Prettify(s)
}

func (s GetProjectResponseBodyDataPropertiesStorageTierInfoStorageTierSize) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyDataPropertiesStorageTierInfoStorageTierSize) GetLongTermSize() *int64 {
	return s.LongTermSize
}

func (s *GetProjectResponseBodyDataPropertiesStorageTierInfoStorageTierSize) GetLowFrequencySize() *int64 {
	return s.LowFrequencySize
}

func (s *GetProjectResponseBodyDataPropertiesStorageTierInfoStorageTierSize) GetStandardSize() *int64 {
	return s.StandardSize
}

func (s *GetProjectResponseBodyDataPropertiesStorageTierInfoStorageTierSize) SetLongTermSize(v int64) *GetProjectResponseBodyDataPropertiesStorageTierInfoStorageTierSize {
	s.LongTermSize = &v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesStorageTierInfoStorageTierSize) SetLowFrequencySize(v int64) *GetProjectResponseBodyDataPropertiesStorageTierInfoStorageTierSize {
	s.LowFrequencySize = &v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesStorageTierInfoStorageTierSize) SetStandardSize(v int64) *GetProjectResponseBodyDataPropertiesStorageTierInfoStorageTierSize {
	s.StandardSize = &v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesStorageTierInfoStorageTierSize) Validate() error {
	return dara.Validate(s)
}

type GetProjectResponseBodyDataPropertiesTableLifecycle struct {
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

func (s GetProjectResponseBodyDataPropertiesTableLifecycle) String() string {
	return dara.Prettify(s)
}

func (s GetProjectResponseBodyDataPropertiesTableLifecycle) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyDataPropertiesTableLifecycle) GetType() *string {
	return s.Type
}

func (s *GetProjectResponseBodyDataPropertiesTableLifecycle) GetValue() *string {
	return s.Value
}

func (s *GetProjectResponseBodyDataPropertiesTableLifecycle) SetType(v string) *GetProjectResponseBodyDataPropertiesTableLifecycle {
	s.Type = &v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesTableLifecycle) SetValue(v string) *GetProjectResponseBodyDataPropertiesTableLifecycle {
	s.Value = &v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesTableLifecycle) Validate() error {
	return dara.Validate(s)
}

type GetProjectResponseBodyDataPropertiesTableLifecycleConfig struct {
	// The information about the long-term storage tier.
	TierToLongterm *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLongterm `json:"TierToLongterm,omitempty" xml:"TierToLongterm,omitempty" type:"Struct"`
	// The information about the IA storage tier.
	TierToLowFrequency *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLowFrequency `json:"TierToLowFrequency,omitempty" xml:"TierToLowFrequency,omitempty" type:"Struct"`
}

func (s GetProjectResponseBodyDataPropertiesTableLifecycleConfig) String() string {
	return dara.Prettify(s)
}

func (s GetProjectResponseBodyDataPropertiesTableLifecycleConfig) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyDataPropertiesTableLifecycleConfig) GetTierToLongterm() *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLongterm {
	return s.TierToLongterm
}

func (s *GetProjectResponseBodyDataPropertiesTableLifecycleConfig) GetTierToLowFrequency() *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLowFrequency {
	return s.TierToLowFrequency
}

func (s *GetProjectResponseBodyDataPropertiesTableLifecycleConfig) SetTierToLongterm(v *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLongterm) *GetProjectResponseBodyDataPropertiesTableLifecycleConfig {
	s.TierToLongterm = v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesTableLifecycleConfig) SetTierToLowFrequency(v *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLowFrequency) *GetProjectResponseBodyDataPropertiesTableLifecycleConfig {
	s.TierToLowFrequency = v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesTableLifecycleConfig) Validate() error {
	return dara.Validate(s)
}

type GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLongterm struct {
	// The system triggers an automatic storage tier change N days after the last access time of data. N is specified by this parameter and corresponds to `LastAccessTime` that is configured for the table or partition.
	//
	// >  If LastAccessTime of a table or partition is left empty, the following rules are applied:
	//
	// 	- For tables or partitions that you created before October 1, 2023, 2023.10.01 00:00:00 in UTC+0 is considered as the last access time.
	//
	// 	- For tables or partitions that you created on or after October 1, 2023, if no data is accessed, the table or partition creation time is considered as the last access time.
	//
	// example:
	//
	// 180
	DaysAfterLastAccessGreaterThan *int64 `json:"DaysAfterLastAccessGreaterThan,omitempty" xml:"DaysAfterLastAccessGreaterThan,omitempty"`
	// The system triggers an automatic storage tier change N days after the last modification time of data. N is specified by this parameter and corresponds to `LastModifiedTime` that is configured for the table or partition.
	//
	// example:
	//
	// 180
	DaysAfterLastModificationGreaterThan *int64 `json:"DaysAfterLastModificationGreaterThan,omitempty" xml:"DaysAfterLastModificationGreaterThan,omitempty"`
	// The period after the previous storage tier change time.
	//
	// example:
	//
	// 1
	DaysAfterLastTierModificationGreaterThan *int64 `json:"DaysAfterLastTierModificationGreaterThan,omitempty" xml:"DaysAfterLastTierModificationGreaterThan,omitempty"`
}

func (s GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLongterm) String() string {
	return dara.Prettify(s)
}

func (s GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLongterm) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLongterm) GetDaysAfterLastAccessGreaterThan() *int64 {
	return s.DaysAfterLastAccessGreaterThan
}

func (s *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLongterm) GetDaysAfterLastModificationGreaterThan() *int64 {
	return s.DaysAfterLastModificationGreaterThan
}

func (s *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLongterm) GetDaysAfterLastTierModificationGreaterThan() *int64 {
	return s.DaysAfterLastTierModificationGreaterThan
}

func (s *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLongterm) SetDaysAfterLastAccessGreaterThan(v int64) *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLongterm {
	s.DaysAfterLastAccessGreaterThan = &v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLongterm) SetDaysAfterLastModificationGreaterThan(v int64) *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLongterm {
	s.DaysAfterLastModificationGreaterThan = &v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLongterm) SetDaysAfterLastTierModificationGreaterThan(v int64) *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLongterm {
	s.DaysAfterLastTierModificationGreaterThan = &v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLongterm) Validate() error {
	return dara.Validate(s)
}

type GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLowFrequency struct {
	// The system triggers an automatic storage tier change N days after the last access time of data. N is specified by this parameter and corresponds to `LastAccessTime` that is configured for the table or partition.
	//
	// >  If LastAccessTime of a table or partition is left empty, the following rules are applied:
	//
	// 	- For tables or partitions that you created before October 1, 2023, 2023.10.01 00:00:00 in UTC+0 is considered as the last access time.
	//
	// 	- For tables or partitions that you created on or after October 1, 2023, if no data is accessed, the table or partition creation time is considered as the last access time.
	//
	// example:
	//
	// 30
	DaysAfterLastAccessGreaterThan *int64 `json:"DaysAfterLastAccessGreaterThan,omitempty" xml:"DaysAfterLastAccessGreaterThan,omitempty"`
	// The system triggers an automatic storage tier change N days after the last modification time of data. N is specified by this parameter and corresponds to `LastModifiedTime` that is configured for the table or partition.
	//
	// example:
	//
	// 30
	DaysAfterLastModificationGreaterThan *int64 `json:"DaysAfterLastModificationGreaterThan,omitempty" xml:"DaysAfterLastModificationGreaterThan,omitempty"`
	// The period after the previous storage tier change time.
	//
	// example:
	//
	// 1
	DaysAfterLastTierModificationGreaterThan *int64 `json:"DaysAfterLastTierModificationGreaterThan,omitempty" xml:"DaysAfterLastTierModificationGreaterThan,omitempty"`
}

func (s GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLowFrequency) String() string {
	return dara.Prettify(s)
}

func (s GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLowFrequency) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLowFrequency) GetDaysAfterLastAccessGreaterThan() *int64 {
	return s.DaysAfterLastAccessGreaterThan
}

func (s *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLowFrequency) GetDaysAfterLastModificationGreaterThan() *int64 {
	return s.DaysAfterLastModificationGreaterThan
}

func (s *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLowFrequency) GetDaysAfterLastTierModificationGreaterThan() *int64 {
	return s.DaysAfterLastTierModificationGreaterThan
}

func (s *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLowFrequency) SetDaysAfterLastAccessGreaterThan(v int64) *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLowFrequency {
	s.DaysAfterLastAccessGreaterThan = &v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLowFrequency) SetDaysAfterLastModificationGreaterThan(v int64) *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLowFrequency {
	s.DaysAfterLastModificationGreaterThan = &v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLowFrequency) SetDaysAfterLastTierModificationGreaterThan(v int64) *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLowFrequency {
	s.DaysAfterLastTierModificationGreaterThan = &v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLowFrequency) Validate() error {
	return dara.Validate(s)
}

type GetProjectResponseBodyDataSaleTag struct {
	// The instance ID of the default computing quota.
	//
	// example:
	//
	// project_name
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The billing method of the default computing quota.
	//
	// example:
	//
	// project
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s GetProjectResponseBodyDataSaleTag) String() string {
	return dara.Prettify(s)
}

func (s GetProjectResponseBodyDataSaleTag) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyDataSaleTag) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetProjectResponseBodyDataSaleTag) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetProjectResponseBodyDataSaleTag) SetResourceId(v string) *GetProjectResponseBodyDataSaleTag {
	s.ResourceId = &v
	return s
}

func (s *GetProjectResponseBodyDataSaleTag) SetResourceType(v string) *GetProjectResponseBodyDataSaleTag {
	s.ResourceType = &v
	return s
}

func (s *GetProjectResponseBodyDataSaleTag) Validate() error {
	return dara.Validate(s)
}

type GetProjectResponseBodyDataSecurityProperties struct {
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
	ProjectProtection *GetProjectResponseBodyDataSecurityPropertiesProjectProtection `json:"projectProtection,omitempty" xml:"projectProtection,omitempty" type:"Struct"`
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

func (s GetProjectResponseBodyDataSecurityProperties) String() string {
	return dara.Prettify(s)
}

func (s GetProjectResponseBodyDataSecurityProperties) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyDataSecurityProperties) GetEnableDownloadPrivilege() *bool {
	return s.EnableDownloadPrivilege
}

func (s *GetProjectResponseBodyDataSecurityProperties) GetLabelSecurity() *bool {
	return s.LabelSecurity
}

func (s *GetProjectResponseBodyDataSecurityProperties) GetObjectCreatorHasAccessPermission() *bool {
	return s.ObjectCreatorHasAccessPermission
}

func (s *GetProjectResponseBodyDataSecurityProperties) GetObjectCreatorHasGrantPermission() *bool {
	return s.ObjectCreatorHasGrantPermission
}

func (s *GetProjectResponseBodyDataSecurityProperties) GetProjectProtection() *GetProjectResponseBodyDataSecurityPropertiesProjectProtection {
	return s.ProjectProtection
}

func (s *GetProjectResponseBodyDataSecurityProperties) GetUsingAcl() *bool {
	return s.UsingAcl
}

func (s *GetProjectResponseBodyDataSecurityProperties) GetUsingPolicy() *bool {
	return s.UsingPolicy
}

func (s *GetProjectResponseBodyDataSecurityProperties) SetEnableDownloadPrivilege(v bool) *GetProjectResponseBodyDataSecurityProperties {
	s.EnableDownloadPrivilege = &v
	return s
}

func (s *GetProjectResponseBodyDataSecurityProperties) SetLabelSecurity(v bool) *GetProjectResponseBodyDataSecurityProperties {
	s.LabelSecurity = &v
	return s
}

func (s *GetProjectResponseBodyDataSecurityProperties) SetObjectCreatorHasAccessPermission(v bool) *GetProjectResponseBodyDataSecurityProperties {
	s.ObjectCreatorHasAccessPermission = &v
	return s
}

func (s *GetProjectResponseBodyDataSecurityProperties) SetObjectCreatorHasGrantPermission(v bool) *GetProjectResponseBodyDataSecurityProperties {
	s.ObjectCreatorHasGrantPermission = &v
	return s
}

func (s *GetProjectResponseBodyDataSecurityProperties) SetProjectProtection(v *GetProjectResponseBodyDataSecurityPropertiesProjectProtection) *GetProjectResponseBodyDataSecurityProperties {
	s.ProjectProtection = v
	return s
}

func (s *GetProjectResponseBodyDataSecurityProperties) SetUsingAcl(v bool) *GetProjectResponseBodyDataSecurityProperties {
	s.UsingAcl = &v
	return s
}

func (s *GetProjectResponseBodyDataSecurityProperties) SetUsingPolicy(v bool) *GetProjectResponseBodyDataSecurityProperties {
	s.UsingPolicy = &v
	return s
}

func (s *GetProjectResponseBodyDataSecurityProperties) Validate() error {
	return dara.Validate(s)
}

type GetProjectResponseBodyDataSecurityPropertiesProjectProtection struct {
	// If you enable the project data protection mechanism, you can configure exception or trusted projects. This allows specified users to transfer data of a specified object to a specified project. The project data protection mechanism does not take effect in all the situations that are specified in the exception policy.
	//
	// example:
	//
	// {
	//
	//     "Version": "1",
	//
	//     "Statement":
	//
	//     [{
	//
	//         "Effect":"Allow",
	//
	//         "Principal":"<Principal>",
	//
	//         "Action":["odps:<Action1>[, <Action2>, ...]"],
	//
	//         "Resource":"acs:odps:*:<Resource>",
	//
	//         "Condition":{
	//
	//             "StringEquals": {
	//
	//                 "odps:TaskType":["<Tasktype>"]
	//
	//             }
	//
	//         }
	//
	//     }]
	//
	//     }
	ExceptionPolicy *string `json:"exceptionPolicy,omitempty" xml:"exceptionPolicy,omitempty"`
	// Indicates whether the [data protection mechanism](https://www.alibabacloud.com/help/zh/maxcompute/security-and-compliance/project-data-protection) is enabled for the project. This allows or denies data transfer across projects. By default, the data protection mechanism is disabled.
	//
	// example:
	//
	// true
	Protected *bool `json:"protected,omitempty" xml:"protected,omitempty"`
}

func (s GetProjectResponseBodyDataSecurityPropertiesProjectProtection) String() string {
	return dara.Prettify(s)
}

func (s GetProjectResponseBodyDataSecurityPropertiesProjectProtection) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyDataSecurityPropertiesProjectProtection) GetExceptionPolicy() *string {
	return s.ExceptionPolicy
}

func (s *GetProjectResponseBodyDataSecurityPropertiesProjectProtection) GetProtected() *bool {
	return s.Protected
}

func (s *GetProjectResponseBodyDataSecurityPropertiesProjectProtection) SetExceptionPolicy(v string) *GetProjectResponseBodyDataSecurityPropertiesProjectProtection {
	s.ExceptionPolicy = &v
	return s
}

func (s *GetProjectResponseBodyDataSecurityPropertiesProjectProtection) SetProtected(v bool) *GetProjectResponseBodyDataSecurityPropertiesProjectProtection {
	s.Protected = &v
	return s
}

func (s *GetProjectResponseBodyDataSecurityPropertiesProjectProtection) Validate() error {
	return dara.Validate(s)
}
