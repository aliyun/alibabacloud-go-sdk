// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyDBInstanceConfigRequest
	GetClientToken() *string
	SetConfigName(v string) *ModifyDBInstanceConfigRequest
	GetConfigName() *string
	SetConfigValue(v string) *ModifyDBInstanceConfigRequest
	GetConfigValue() *string
	SetDBInstanceId(v string) *ModifyDBInstanceConfigRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *ModifyDBInstanceConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBInstanceConfigRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *ModifyDBInstanceConfigRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifyDBInstanceConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceConfigRequest
	GetResourceOwnerId() *int64
	SetSwitchTime(v string) *ModifyDBInstanceConfigRequest
	GetSwitchTime() *string
	SetSwitchTimeMode(v string) *ModifyDBInstanceConfigRequest
	GetSwitchTimeMode() *string
}

type ModifyDBInstanceConfigRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 6000170000591aed949d0f****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The name of the configuration item to modify. This parameter is used together with ConfigValue.
	//
	// <details>
	//
	// <summary>ApsaraDB RDS for PostgreSQL configuration items</summary>
	//
	// - **pgbouncer**: Modifies the PgBouncer feature.
	//
	// - **encryptionKey**: Modifies the cloud disk encryption feature.
	//
	// - **duckdb_create_databases**: Configures databases of the primary instance as DuckDB column store databases in batches.
	//
	// - **duckdb_prepare_dependency**: Configures the primary instance with one click so that its parameters and minor engine version meet the [prerequisites](https://help.aliyun.com/document_detail/2977241.html) for creating a DuckDB-based analytical instance. If the primary instance already has read-only instances, the read-only instances are also updated.
	//
	// - **enable_db_visible_by_connect_rls**: Enables CONNECT RLS on the instance to control database visibility.
	//
	// - **set_db_visible_by_connect_rls**: Enables CONNECT RLS on a database to control database visibility. This can be called only after CONNECT RLS is enabled on the instance.
	//
	// </details>
	//
	// <details>
	//
	// <summary>ApsaraDB RDS for SQL Server configuration items</summary>
	//
	// <props="intl">
	//
	// - **clear_errorlog**: Clears error logs.
	//
	// - **encryptionKey**: Modifies the cloud disk encryption feature. Serverless instances and shared instance types do not support this feature.
	//
	//
	//
	// <props="china">
	//
	//
	// - **backup_recovery_model**: Enables the simple recovery model feature. Only Basic Edition instances support this feature. **This feature cannot be disabled after it is enabled**.
	//
	// - **clear_errorlog**: Clears error logs.
	//
	// - **encryptionKey**: Modifies the cloud disk encryption feature. Serverless instances and shared instance types do not support this feature.
	//
	//
	// </details>
	//
	// This parameter is required.
	//
	// example:
	//
	// pgbouncer
	ConfigName *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	// The value of the configuration item to modify. This parameter is used together with ConfigName.
	//
	// <details>
	//
	// <summary>ApsaraDB RDS for PostgreSQL configuration item values</summary>
	//
	// - PgBouncer feature: **true*	- (enable) or **false*	- (disable).
	//
	// - Cloud disk encryption feature:
	//
	//   - **ServiceKey**: Uses an automatically generated key from Alibaba Cloud, which is the RDS-managed service key (Default Service CMK), to enable cloud disk encryption.
	//
	//   - **<Key>**: Uses a custom key to enable cloud disk encryption or replaces the current key. Example: `494c98ce-f2b5-48ab-96ab-36c986b6****`.
	//
	//   - **disabled**: Disables cloud disk encryption.
	//
	// - One-click fix for prerequisites to create a DuckDB-based analytical instance: **duckdb_prepare_dependency**
	//
	// - Configure databases of the primary instance as DuckDB column store databases in batches. The value is a JSON string. Example: `{"dbNames": "db1,db2,db3", "accountName": "yourSuperAccountName"}`, where:
	//
	//   - **dbNames**: The names of databases to convert to DuckDB column store databases. Separate multiple database names with commas (,).
	//
	//   - **accountName**: The privileged user. Specify only one privileged user.
	//
	// - Enable CONNECT RLS on the instance to control database visibility: **true*	- to enable.
	//
	// - Enable CONNECT RLS on a database to control database visibility: The **database names*	- managed by CONNECT RLS. Separate multiple database names with commas (,). Example: **testdb1,testdb2**. When a client connects to a database with CONNECT RLS enabled, the database list is displayed based on whether the client has CONNECT permissions on other databases.
	//
	// </details>
	//
	// <details>
	//
	// <summary>ApsaraDB RDS for SQL Server configuration item values</summary>
	//
	// <props="intl">
	//
	// - Error log cleanup feature: **1*	- (confirm cleanup).
	//
	// - Cloud disk encryption feature (**this feature cannot be disabled after it is enabled**):
	//
	//   - **serviceKey**: Uses an automatically generated key from Alibaba Cloud, which is the RDS-managed service key (Default Service CMK), to enable cloud disk encryption.
	//
	//   - **<Key>**: Uses a custom key to enable cloud disk encryption or replaces the current key. Example: `494c98ce-f2b5-48ab-96ab-36c986b6****`.
	//
	//
	//
	//
	// <props="china">
	//
	// - Simple recovery feature: **simple*	- (enable simple recovery).
	//
	// - Error log cleanup feature: **1*	- (confirm cleanup).
	//
	// - Cloud disk encryption feature (**this feature cannot be disabled after it is enabled**):
	//
	//   - **serviceKey**: Uses an automatically generated key from Alibaba Cloud, which is the RDS-managed service key (Default Service CMK), to enable cloud disk encryption.
	//
	//   - **<Key>**: Uses a custom key to enable cloud disk encryption or replaces the current key. Example: `494c98ce-f2b5-48ab-96ab-36c986b6****`.
	//
	//
	//
	// </details>
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	// The instance ID. You can call DescribeDBInstances to obtain the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pgm-2ze****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The resource group ID. You can call DescribeDBInstanceAttribute to obtain the resource group ID.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The time at which the modification takes effect. We recommend that you perform specification changes during off-peak hours. Format: <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC).
	//
	// example:
	//
	// 2025-05-06T09:24:00Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// The switchover time. Valid values:
	//
	// - **Immediate**: The modification takes effect immediately.
	//
	// - **MaintainTime**: The modification takes effect during the maintenance window. You can call ModifyDBInstanceMaintainTime to modify the maintenance window.
	//
	// example:
	//
	// Immediate
	SwitchTimeMode *string `json:"SwitchTimeMode,omitempty" xml:"SwitchTimeMode,omitempty"`
}

func (s ModifyDBInstanceConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConfigRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyDBInstanceConfigRequest) GetConfigName() *string {
	return s.ConfigName
}

func (s *ModifyDBInstanceConfigRequest) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *ModifyDBInstanceConfigRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBInstanceConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBInstanceConfigRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyDBInstanceConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBInstanceConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBInstanceConfigRequest) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *ModifyDBInstanceConfigRequest) GetSwitchTimeMode() *string {
	return s.SwitchTimeMode
}

func (s *ModifyDBInstanceConfigRequest) SetClientToken(v string) *ModifyDBInstanceConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDBInstanceConfigRequest) SetConfigName(v string) *ModifyDBInstanceConfigRequest {
	s.ConfigName = &v
	return s
}

func (s *ModifyDBInstanceConfigRequest) SetConfigValue(v string) *ModifyDBInstanceConfigRequest {
	s.ConfigValue = &v
	return s
}

func (s *ModifyDBInstanceConfigRequest) SetDBInstanceId(v string) *ModifyDBInstanceConfigRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceConfigRequest) SetOwnerAccount(v string) *ModifyDBInstanceConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceConfigRequest) SetOwnerId(v int64) *ModifyDBInstanceConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceConfigRequest) SetResourceGroupId(v string) *ModifyDBInstanceConfigRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyDBInstanceConfigRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceConfigRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceConfigRequest) SetSwitchTime(v string) *ModifyDBInstanceConfigRequest {
	s.SwitchTime = &v
	return s
}

func (s *ModifyDBInstanceConfigRequest) SetSwitchTimeMode(v string) *ModifyDBInstanceConfigRequest {
	s.SwitchTimeMode = &v
	return s
}

func (s *ModifyDBInstanceConfigRequest) Validate() error {
	return dara.Validate(s)
}
