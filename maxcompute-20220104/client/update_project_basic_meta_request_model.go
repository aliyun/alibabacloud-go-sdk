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
	// The description of the project.
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
	if s.Properties != nil {
		if err := s.Properties.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateProjectBasicMetaRequestProperties struct {
	// Specifies whether to allow full table scans in the project. A full table scan consumes a large amount of resources. To improve processing efficiency, this feature is disabled by default.
	//
	// example:
	//
	// false
	AllowFullScan *bool `json:"allowFullScan,omitempty" xml:"allowFullScan,omitempty"`
	// Specifies whether to enable the Decimal data type of MaxCompute V2.0 for the project.
	//
	// example:
	//
	// true
	EnableDecimal2 *bool `json:"enableDecimal2,omitempty" xml:"enableDecimal2,omitempty"`
	EnableDr       *bool `json:"enableDr,omitempty" xml:"enableDr,omitempty"`
	// Specifies whether to enable resource group-based routing for Data Transmission Service.
	//
	// - true: Data transmission tasks submitted in the project use the attached Data Transmission Service resource group by default.
	//
	// - false: Data transmission tasks submitted in the project use the shared Data Transmission Service resource group by default.
	//
	// example:
	//
	// true
	EnableTunnelQuotaRoute *bool `json:"enableTunnelQuotaRoute,omitempty" xml:"enableTunnelQuotaRoute,omitempty"`
	// The storage encryption properties.
	Encryption *UpdateProjectBasicMetaRequestPropertiesEncryption `json:"encryption,omitempty" xml:"encryption,omitempty" type:"Struct"`
	// The number of days to retain backup data. During this period, you can restore the current version to any backup version.
	//
	// The value must be an integer from 0 to 30. The default value is 1. A value of 0 disables the backup feature.
	//
	// example:
	//
	// 1
	RetentionDays *int64 `json:"retentionDays,omitempty" xml:"retentionDays,omitempty"`
	// The maximum consumption threshold for a single SQL job.
	//
	// Unit: Scanned data (GB) × Complexity.
	//
	// example:
	//
	// 1500
	SqlMeteringMax *string `json:"sqlMeteringMax,omitempty" xml:"sqlMeteringMax,omitempty"`
	// The lifecycle properties of the table.
	TableLifecycle *UpdateProjectBasicMetaRequestPropertiesTableLifecycle `json:"tableLifecycle,omitempty" xml:"tableLifecycle,omitempty" type:"Struct"`
	// The time zone of the project. This is the `odps.sql.timezone` property.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// The <props="intl">[Data Transmission Service](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/overview-of-dts) resource group attached to the project.
	//
	// - Default (shared Data Transmission Service resource group): The project is not allowed to use a subscription Data Transmission Service resource group. Regardless of the value of the default Data Transmission Service resource group, data transmission tasks submitted in the project automatically use the Default resource group.
	//
	// - Subscription Data Transmission Service resource group: The project is allowed to use a subscription Data Transmission Service resource group.
	//
	// example:
	//
	// Default
	TunnelQuota *string `json:"tunnelQuota,omitempty" xml:"tunnelQuota,omitempty"`
	// The data type edition. Valid values:
	//
	// - **1**: Edition 1.0
	//
	// - **2**: Edition 2.0
	//
	// - **hive**: Hive-compatible edition
	//
	// For more information about the differences between the data type editions, see <props="intl">[Data type editions](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/data-type-editions).
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
	if s.Encryption != nil {
		if err := s.Encryption.Validate(); err != nil {
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

type UpdateProjectBasicMetaRequestPropertiesEncryption struct {
	// The encryption algorithm. The key supports algorithms such as AES256, AESCTR, and RC4.
	//
	// example:
	//
	// AES256
	Algorithm *string `json:"algorithm,omitempty" xml:"algorithm,omitempty"`
	// Specifies whether to enable data encryption for the project. For more information about data encryption, see
	//
	// <props="intl">[Storage encryption](https://www.alibabacloud.com/help/zh/maxcompute/security-and-compliance/storage-encryption).
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The type of key used for data encryption. This can be the default MaxCompute key or a Bring-Your-Own-Key (BYOK). The default MaxCompute key is created within MaxCompute.
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
	// - **mandatory**: The Lifecycle clause is required. You must set a lifecycle for the table.
	//
	// - **optional**: The Lifecycle clause is optional when you create a table. If you do not set a lifecycle for the table, the table never expires.
	//
	// - **inherit**: If you do not set a lifecycle for the table when you create it, the lifecycle of the table is the value of odps.table.lifecycle.value.
	//
	// example:
	//
	// optional
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The lifecycle of the table in days. The value must be an integer from 1 to 37231. The default value is 37231.
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
