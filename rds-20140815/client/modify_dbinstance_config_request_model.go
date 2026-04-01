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
	// The name of the configuration item that you want to modify. Valid values:
	//
	// 	- **pgbouncer**. This configuration item is supported for ApsaraDB RDS for PostgreSQL instances.
	//
	// 	- **clear_errorlog**. This configuration item is supported for ApsaraDB RDS for SQL Server instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// pgbouncer
	ConfigName *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	// The value of the configuration item that you want to modify. Valid values:
	//
	// 	- If you set ConfigName to pgbouncer, the valid values are **true*	- and **false**.
	//
	// 	- If you set ConfigName to clear_errorlog, set the value to **1**. The value 1 indicates that error logs are cleaned up.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The resource group ID. You can call the DescribeDBInstanceAttribute to obtain the resource group ID.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The update time. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2022-05-06T09:24:00Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// The time at which the modification takes effect. Valid values:
	//
	// - **Immediate**: immediately modifies the parameter. This is the default value.
	//
	// - **MaintainTime**: modifies the parameter during the maintenance window of the instance. You can call the ModifyDBInstanceMaintainTime operation to change the maintenance window.
	//
	// - **ScheduleTime**: modifies the parameter at the point in time that you specify. If you specify this value, you must also specify **SwitchTime**.
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
