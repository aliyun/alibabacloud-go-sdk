// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectBasicMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *UpdateProjectBasicMetaRequest
	GetComment() *string
	SetProperties(v *UpdateProjectBasicMetaRequestProperties) *UpdateProjectBasicMetaRequest
	GetProperties() *UpdateProjectBasicMetaRequestProperties
}

type UpdateProjectBasicMetaRequest struct {
	// The project description.
	//
	// example:
	//
	// BI_Analysis
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// The basic properties of the project.
	Properties *UpdateProjectBasicMetaRequestProperties `json:"properties,omitempty" xml:"properties,omitempty" type:"Struct"`
}

func (s UpdateProjectBasicMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectBasicMetaRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectBasicMetaRequest) GetComment() *string {
	return s.Comment
}

func (s *UpdateProjectBasicMetaRequest) GetProperties() *UpdateProjectBasicMetaRequestProperties {
	return s.Properties
}

func (s *UpdateProjectBasicMetaRequest) SetComment(v string) *UpdateProjectBasicMetaRequest {
	s.Comment = &v
	return s
}

func (s *UpdateProjectBasicMetaRequest) SetProperties(v *UpdateProjectBasicMetaRequestProperties) *UpdateProjectBasicMetaRequest {
	s.Properties = v
	return s
}

func (s *UpdateProjectBasicMetaRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateProjectBasicMetaRequestProperties struct {
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
	EnableDr       *bool `json:"enableDr,omitempty" xml:"enableDr,omitempty"`
	// Indicates whether the routing of the Tunnel resource group is enabled.
	//
	// - true: The data transfer tasks that are submitted by the project by default use the Tunnel resource group that is bound to the project.
	//
	// - false: The data transfer tasks that are submitted by the project by default use the Tunnel shared resource group.
	//
	// example:
	//
	// true
	EnableTunnelQuotaRoute *bool `json:"enableTunnelQuotaRoute,omitempty" xml:"enableTunnelQuotaRoute,omitempty"`
	// The storage encryption properties.
	Encryption *UpdateProjectBasicMetaRequestPropertiesEncryption `json:"encryption,omitempty" xml:"encryption,omitempty" type:"Struct"`
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
	TableLifecycle *UpdateProjectBasicMetaRequestPropertiesTableLifecycle `json:"tableLifecycle,omitempty" xml:"tableLifecycle,omitempty" type:"Struct"`
	// The time zone that is used by your project. The time zone is the same as the time zone specified by `odps.sql.timezone` .
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// The <props="china">[Data Transmission Service](https://help.aliyun.com/zh/maxcompute/user-guide/overview-of-dts)
	//
	// <props="intl">[Data Transmission Service](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/overview-of-dts) resource group that is bound to the project.
	//
	// - Default resource group: The Tunnel shared resource group is used. You cannot use the subscription-based Tunnel resource group for the project. The default resource group is automatically used by the Tunnel service of your project, regardless of the parameter setting.
	//
	// - Subscription-based Tunnel resource group: You can use the subscription-based Tunnel resource group for the project.
	//
	// example:
	//
	// Default
	TunnelQuota *string `json:"tunnelQuota,omitempty" xml:"tunnelQuota,omitempty"`
	// The data type edition. Valid values:
	//
	// - *1*: MaxCompute V1.0 data type edition
	//
	// - *2*: MaxCompute V2.0 data type edition
	//
	// - *hive*: Hive-compatible data type edition
	//
	// For more information about the differences among the three data type editions, see <props="china">[Data Type Versions](https://help.aliyun.com/zh/maxcompute/user-guide/data-type-editions)
	//
	// <props="intl">[Data Type Versions](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/data-type-editions).
	//
	// example:
	//
	// 2.0
	TypeSystem *string `json:"typeSystem,omitempty" xml:"typeSystem,omitempty"`
}

func (s UpdateProjectBasicMetaRequestProperties) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectBasicMetaRequestProperties) GoString() string {
	return s.String()
}

func (s *UpdateProjectBasicMetaRequestProperties) GetAllowFullScan() *bool {
	return s.AllowFullScan
}

func (s *UpdateProjectBasicMetaRequestProperties) GetEnableDecimal2() *bool {
	return s.EnableDecimal2
}

func (s *UpdateProjectBasicMetaRequestProperties) GetEnableDr() *bool {
	return s.EnableDr
}

func (s *UpdateProjectBasicMetaRequestProperties) GetEnableTunnelQuotaRoute() *bool {
	return s.EnableTunnelQuotaRoute
}

func (s *UpdateProjectBasicMetaRequestProperties) GetEncryption() *UpdateProjectBasicMetaRequestPropertiesEncryption {
	return s.Encryption
}

func (s *UpdateProjectBasicMetaRequestProperties) GetRetentionDays() *int64 {
	return s.RetentionDays
}

func (s *UpdateProjectBasicMetaRequestProperties) GetSqlMeteringMax() *string {
	return s.SqlMeteringMax
}

func (s *UpdateProjectBasicMetaRequestProperties) GetTableLifecycle() *UpdateProjectBasicMetaRequestPropertiesTableLifecycle {
	return s.TableLifecycle
}

func (s *UpdateProjectBasicMetaRequestProperties) GetTimezone() *string {
	return s.Timezone
}

func (s *UpdateProjectBasicMetaRequestProperties) GetTunnelQuota() *string {
	return s.TunnelQuota
}

func (s *UpdateProjectBasicMetaRequestProperties) GetTypeSystem() *string {
	return s.TypeSystem
}

func (s *UpdateProjectBasicMetaRequestProperties) SetAllowFullScan(v bool) *UpdateProjectBasicMetaRequestProperties {
	s.AllowFullScan = &v
	return s
}

func (s *UpdateProjectBasicMetaRequestProperties) SetEnableDecimal2(v bool) *UpdateProjectBasicMetaRequestProperties {
	s.EnableDecimal2 = &v
	return s
}

func (s *UpdateProjectBasicMetaRequestProperties) SetEnableDr(v bool) *UpdateProjectBasicMetaRequestProperties {
	s.EnableDr = &v
	return s
}

func (s *UpdateProjectBasicMetaRequestProperties) SetEnableTunnelQuotaRoute(v bool) *UpdateProjectBasicMetaRequestProperties {
	s.EnableTunnelQuotaRoute = &v
	return s
}

func (s *UpdateProjectBasicMetaRequestProperties) SetEncryption(v *UpdateProjectBasicMetaRequestPropertiesEncryption) *UpdateProjectBasicMetaRequestProperties {
	s.Encryption = v
	return s
}

func (s *UpdateProjectBasicMetaRequestProperties) SetRetentionDays(v int64) *UpdateProjectBasicMetaRequestProperties {
	s.RetentionDays = &v
	return s
}

func (s *UpdateProjectBasicMetaRequestProperties) SetSqlMeteringMax(v string) *UpdateProjectBasicMetaRequestProperties {
	s.SqlMeteringMax = &v
	return s
}

func (s *UpdateProjectBasicMetaRequestProperties) SetTableLifecycle(v *UpdateProjectBasicMetaRequestPropertiesTableLifecycle) *UpdateProjectBasicMetaRequestProperties {
	s.TableLifecycle = v
	return s
}

func (s *UpdateProjectBasicMetaRequestProperties) SetTimezone(v string) *UpdateProjectBasicMetaRequestProperties {
	s.Timezone = &v
	return s
}

func (s *UpdateProjectBasicMetaRequestProperties) SetTunnelQuota(v string) *UpdateProjectBasicMetaRequestProperties {
	s.TunnelQuota = &v
	return s
}

func (s *UpdateProjectBasicMetaRequestProperties) SetTypeSystem(v string) *UpdateProjectBasicMetaRequestProperties {
	s.TypeSystem = &v
	return s
}

func (s *UpdateProjectBasicMetaRequestProperties) Validate() error {
	return dara.Validate(s)
}

type UpdateProjectBasicMetaRequestPropertiesEncryption struct {
	// The data encryption algorithm that is supported by the key. Valid values: AES256, AESCTR, and RC4.
	//
	// example:
	//
	// AES256
	Algorithm *string `json:"algorithm,omitempty" xml:"algorithm,omitempty"`
	// Indicates whether the data encryption feature needs to be enabled for the project. For more information about data encryption, see
	//
	// <props="china">[Storage Encryption](https://help.aliyun.com/zh/maxcompute/security-and-compliance/storage-encryption)
	//
	// <props="intl">[Storage Encryption](https://www.alibabacloud.com/help/zh/maxcompute/security-and-compliance/storage-encryption).
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The type of key that is used for data encryption. You can select MaxCompute Default Key or Bring Your Own Key (BYOK) as the key type. If you select MaxCompute Default Key, the default key that is created by MaxCompute is used.
	//
	// example:
	//
	// default
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s UpdateProjectBasicMetaRequestPropertiesEncryption) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectBasicMetaRequestPropertiesEncryption) GoString() string {
	return s.String()
}

func (s *UpdateProjectBasicMetaRequestPropertiesEncryption) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *UpdateProjectBasicMetaRequestPropertiesEncryption) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateProjectBasicMetaRequestPropertiesEncryption) GetKey() *string {
	return s.Key
}

func (s *UpdateProjectBasicMetaRequestPropertiesEncryption) SetAlgorithm(v string) *UpdateProjectBasicMetaRequestPropertiesEncryption {
	s.Algorithm = &v
	return s
}

func (s *UpdateProjectBasicMetaRequestPropertiesEncryption) SetEnable(v bool) *UpdateProjectBasicMetaRequestPropertiesEncryption {
	s.Enable = &v
	return s
}

func (s *UpdateProjectBasicMetaRequestPropertiesEncryption) SetKey(v string) *UpdateProjectBasicMetaRequestPropertiesEncryption {
	s.Key = &v
	return s
}

func (s *UpdateProjectBasicMetaRequestPropertiesEncryption) Validate() error {
	return dara.Validate(s)
}

type UpdateProjectBasicMetaRequestPropertiesTableLifecycle struct {
	// The lifecycle type. Valid values:
	//
	// - *mandatory*: The lifecycle clause is required in a table creation statement.
	//
	// - *optional*: The lifecycle clause is optional in a table creation statement. If you do not configure a lifecycle for a table, the table does not expire.
	//
	// - *inherit*: If you do not configure a lifecycle for a table when you create the table, the value of the odps.table.lifecycle.value parameter is used as the table lifecycle by default.
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

func (s UpdateProjectBasicMetaRequestPropertiesTableLifecycle) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectBasicMetaRequestPropertiesTableLifecycle) GoString() string {
	return s.String()
}

func (s *UpdateProjectBasicMetaRequestPropertiesTableLifecycle) GetType() *string {
	return s.Type
}

func (s *UpdateProjectBasicMetaRequestPropertiesTableLifecycle) GetValue() *string {
	return s.Value
}

func (s *UpdateProjectBasicMetaRequestPropertiesTableLifecycle) SetType(v string) *UpdateProjectBasicMetaRequestPropertiesTableLifecycle {
	s.Type = &v
	return s
}

func (s *UpdateProjectBasicMetaRequestPropertiesTableLifecycle) SetValue(v string) *UpdateProjectBasicMetaRequestPropertiesTableLifecycle {
	s.Value = &v
	return s
}

func (s *UpdateProjectBasicMetaRequestPropertiesTableLifecycle) Validate() error {
	return dara.Validate(s)
}
