// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataLinkName(v string) *UpdateInstanceRequest
	GetDataLinkName() *string
	SetDatabasePassword(v string) *UpdateInstanceRequest
	GetDatabasePassword() *string
	SetDatabaseUser(v string) *UpdateInstanceRequest
	GetDatabaseUser() *string
	SetDbaId(v string) *UpdateInstanceRequest
	GetDbaId() *string
	SetDdlOnline(v int32) *UpdateInstanceRequest
	GetDdlOnline() *int32
	SetEcsInstanceId(v string) *UpdateInstanceRequest
	GetEcsInstanceId() *string
	SetEcsRegion(v string) *UpdateInstanceRequest
	GetEcsRegion() *string
	SetEnableSellSitd(v string) *UpdateInstanceRequest
	GetEnableSellSitd() *string
	SetEnvType(v string) *UpdateInstanceRequest
	GetEnvType() *string
	SetExportTimeout(v int32) *UpdateInstanceRequest
	GetExportTimeout() *int32
	SetHost(v string) *UpdateInstanceRequest
	GetHost() *string
	SetInstanceAlias(v string) *UpdateInstanceRequest
	GetInstanceAlias() *string
	SetInstanceId(v string) *UpdateInstanceRequest
	GetInstanceId() *string
	SetInstanceSource(v string) *UpdateInstanceRequest
	GetInstanceSource() *string
	SetInstanceType(v string) *UpdateInstanceRequest
	GetInstanceType() *string
	SetPort(v int32) *UpdateInstanceRequest
	GetPort() *int32
	SetQueryTimeout(v int32) *UpdateInstanceRequest
	GetQueryTimeout() *int32
	SetSafeRuleId(v string) *UpdateInstanceRequest
	GetSafeRuleId() *string
	SetSid(v string) *UpdateInstanceRequest
	GetSid() *string
	SetSkipTest(v bool) *UpdateInstanceRequest
	GetSkipTest() *bool
	SetTemplateId(v int64) *UpdateInstanceRequest
	GetTemplateId() *int64
	SetTemplateType(v string) *UpdateInstanceRequest
	GetTemplateType() *string
	SetTid(v int64) *UpdateInstanceRequest
	GetTid() *int64
	SetUseDsql(v int32) *UpdateInstanceRequest
	GetUseDsql() *int32
	SetVpcId(v string) *UpdateInstanceRequest
	GetVpcId() *string
}

type UpdateInstanceRequest struct {
	// The name of the database link for cross-database queries.
	//
	// >
	//
	// 	- This parameter is required if UseDsql is set to 1.
	//
	// 	- The name can contain only lowercase letters and underscores (_).
	//
	// 	- The name must be unique within a tenant.
	//
	// example:
	//
	// datalink_test
	DataLinkName *string `json:"DataLinkName,omitempty" xml:"DataLinkName,omitempty"`
	// The password that is used to log on to the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// ******
	DatabasePassword *string `json:"DatabasePassword,omitempty" xml:"DatabasePassword,omitempty"`
	// The account that is used to log on to the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbuser
	DatabaseUser *string `json:"DatabaseUser,omitempty" xml:"DatabaseUser,omitempty"`
	// The ID of the user who assumes the database administrator (DBA) role of the database instance. You can call the [ListUsers](https://help.aliyun.com/document_detail/141938.html) or [GetInstance](https://help.aliyun.com/document_detail/141567.html) operation to query the user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 27****
	DbaId *string `json:"DbaId,omitempty" xml:"DbaId,omitempty"`
	// Specifies whether to enable the lock-free schema change feature for the database instance. Valid values:
	//
	// 	- **0:*	- disables the lock-free schema change feature.
	//
	// 	- **1**: uses the online DDL of MySQL first.
	//
	// 	- **2**: uses the lock-free schema change feature of DMS first.
	//
	// example:
	//
	// 0
	DdlOnline *int32 `json:"DdlOnline,omitempty" xml:"DdlOnline,omitempty"`
	// The ID of the ECS instance on which the database instance is deployed.
	//
	// > This parameter is required if the InstanceSource parameter is set to ECS_OWN.
	//
	// example:
	//
	// i-2zei9gs1t7h8l7ac****
	EcsInstanceId *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	// The ID of the region in which the database instance resides.
	//
	// > This parameter is required if InstanceSource is set to RDS, ECS_OWN, and VPC_IDC.
	//
	// example:
	//
	// cn-hangzhou
	EcsRegion *string `json:"EcsRegion,omitempty" xml:"EcsRegion,omitempty"`
	// 	- **Y:*	- enables the sensitive data protection feature
	//
	// 	- **N:*	- disables the sensitive data protection feature
	//
	// 	- **NULL or other:*	- does not update the status of the sensitive data protection feature
	//
	// example:
	//
	// Y
	EnableSellSitd *string `json:"EnableSellSitd,omitempty" xml:"EnableSellSitd,omitempty"`
	// The type of the environment in which the database instance is deployed. Valid values:
	//
	// 	- **product:*	- production environment
	//
	// 	- **dev:*	- development environment
	//
	// 	- **pre:*	- pre-release environment
	//
	// 	- **test:*	- test environment
	//
	// 	- **sit:*	- system integration testing (SIT) environment
	//
	// 	- **uat:*	- user acceptance testing (UAT) environment
	//
	// 	- **pet:*	- stress testing environment
	//
	// 	- **stag:*	- staging environment
	//
	// This parameter is required.
	//
	// example:
	//
	// product
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The timeout period for exporting data from the database instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 600
	ExportTimeout *int32 `json:"ExportTimeout,omitempty" xml:"ExportTimeout,omitempty"`
	// The host address that is used to connect to the database instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.XXX.0.56
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The alias of the database instance. Specify an alias that can help you identify the database instance in DMS.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance_test
	InstanceAlias *string `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	// The ID of the database instance. You can call the [GetInstance](https://help.aliyun.com/document_detail/141567.html) operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 126****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The source of the database instance. Valid values:
	//
	// 	- **PUBLIC_OWN:*	- a self-managed database instance that is deployed on the Internet
	//
	// 	- **RDS:*	- an ApsaraDB RDS instance
	//
	// 	- **ECS_OWN:*	- a self-managed database that is deployed on an Elastic Compute Service (ECS) instance
	//
	// 	- **VPC_IDC:*	- a self-managed database instance that is deployed in a data center connected over a virtual private cloud (VPC)
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS_OWN
	InstanceSource *string `json:"InstanceSource,omitempty" xml:"InstanceSource,omitempty"`
	// The type of the database. For more information about the valid values of this parameter, see [DbType parameter](https://help.aliyun.com/document_detail/198106.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// mysql
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The port that is used to connect to the database instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The timeout period for querying data in the database instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	QueryTimeout *int32 `json:"QueryTimeout,omitempty" xml:"QueryTimeout,omitempty"`
	// The name of the security rule set (GroupName) for the instance. You can call the [ListStandardGroups](https://help.aliyun.com/document_detail/417891.html) or [GetInstance](https://help.aliyun.com/document_detail/141567.html) operation to query the name of the security rule set.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3****
	SafeRuleId *string `json:"SafeRuleId,omitempty" xml:"SafeRuleId,omitempty"`
	// The system ID (SID) of the database instance.
	//
	// > This parameter is required if the InstanceType parameter is set to ORACLE.
	//
	// example:
	//
	// XXX
	Sid *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	// Specifies whether to skip the connectivity test. Valid values:
	//
	// 	- **true:*	- skips the connectivity test
	//
	// 	- **false:*	- does not skip the connectivity test
	//
	// example:
	//
	// false
	SkipTest *bool `json:"SkipTest,omitempty" xml:"SkipTest,omitempty"`
	// The ID of the classification template. You can call the [ListClassificationTemplates](https://help.aliyun.com/document_detail/460613.html) operation to query the template ID.
	//
	// example:
	//
	// 3***
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The type of the classification template. You can call the [ListClassificationTemplates](https://help.aliyun.com/document_detail/460613.html) operation to query the template type.
	//
	// example:
	//
	// INNER
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) operation to query the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// Specifies whether to enable the cross-database query feature for the database instance. Valid values:
	//
	// 	- **0**: disables the cross-database query feature.
	//
	// 	- **1**: enables the cross-database query feature.
	//
	// > Supported database types: MySQL, SQL Server, PostgreSQL, PolarDB for PostgreSQL (compatible with Oracle), and ApsaraDB for Redis.
	//
	// example:
	//
	// 0
	UseDsql *int32 `json:"UseDsql,omitempty" xml:"UseDsql,omitempty"`
	// The ID of the VPC to which the database instance belongs.
	//
	// > This parameter is required if the InstanceSource parameter is set to VPC_IDC.
	//
	// example:
	//
	// vpc-xxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s UpdateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequest) GetDataLinkName() *string {
	return s.DataLinkName
}

func (s *UpdateInstanceRequest) GetDatabasePassword() *string {
	return s.DatabasePassword
}

func (s *UpdateInstanceRequest) GetDatabaseUser() *string {
	return s.DatabaseUser
}

func (s *UpdateInstanceRequest) GetDbaId() *string {
	return s.DbaId
}

func (s *UpdateInstanceRequest) GetDdlOnline() *int32 {
	return s.DdlOnline
}

func (s *UpdateInstanceRequest) GetEcsInstanceId() *string {
	return s.EcsInstanceId
}

func (s *UpdateInstanceRequest) GetEcsRegion() *string {
	return s.EcsRegion
}

func (s *UpdateInstanceRequest) GetEnableSellSitd() *string {
	return s.EnableSellSitd
}

func (s *UpdateInstanceRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *UpdateInstanceRequest) GetExportTimeout() *int32 {
	return s.ExportTimeout
}

func (s *UpdateInstanceRequest) GetHost() *string {
	return s.Host
}

func (s *UpdateInstanceRequest) GetInstanceAlias() *string {
	return s.InstanceAlias
}

func (s *UpdateInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateInstanceRequest) GetInstanceSource() *string {
	return s.InstanceSource
}

func (s *UpdateInstanceRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *UpdateInstanceRequest) GetPort() *int32 {
	return s.Port
}

func (s *UpdateInstanceRequest) GetQueryTimeout() *int32 {
	return s.QueryTimeout
}

func (s *UpdateInstanceRequest) GetSafeRuleId() *string {
	return s.SafeRuleId
}

func (s *UpdateInstanceRequest) GetSid() *string {
	return s.Sid
}

func (s *UpdateInstanceRequest) GetSkipTest() *bool {
	return s.SkipTest
}

func (s *UpdateInstanceRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *UpdateInstanceRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *UpdateInstanceRequest) GetTid() *int64 {
	return s.Tid
}

func (s *UpdateInstanceRequest) GetUseDsql() *int32 {
	return s.UseDsql
}

func (s *UpdateInstanceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateInstanceRequest) SetDataLinkName(v string) *UpdateInstanceRequest {
	s.DataLinkName = &v
	return s
}

func (s *UpdateInstanceRequest) SetDatabasePassword(v string) *UpdateInstanceRequest {
	s.DatabasePassword = &v
	return s
}

func (s *UpdateInstanceRequest) SetDatabaseUser(v string) *UpdateInstanceRequest {
	s.DatabaseUser = &v
	return s
}

func (s *UpdateInstanceRequest) SetDbaId(v string) *UpdateInstanceRequest {
	s.DbaId = &v
	return s
}

func (s *UpdateInstanceRequest) SetDdlOnline(v int32) *UpdateInstanceRequest {
	s.DdlOnline = &v
	return s
}

func (s *UpdateInstanceRequest) SetEcsInstanceId(v string) *UpdateInstanceRequest {
	s.EcsInstanceId = &v
	return s
}

func (s *UpdateInstanceRequest) SetEcsRegion(v string) *UpdateInstanceRequest {
	s.EcsRegion = &v
	return s
}

func (s *UpdateInstanceRequest) SetEnableSellSitd(v string) *UpdateInstanceRequest {
	s.EnableSellSitd = &v
	return s
}

func (s *UpdateInstanceRequest) SetEnvType(v string) *UpdateInstanceRequest {
	s.EnvType = &v
	return s
}

func (s *UpdateInstanceRequest) SetExportTimeout(v int32) *UpdateInstanceRequest {
	s.ExportTimeout = &v
	return s
}

func (s *UpdateInstanceRequest) SetHost(v string) *UpdateInstanceRequest {
	s.Host = &v
	return s
}

func (s *UpdateInstanceRequest) SetInstanceAlias(v string) *UpdateInstanceRequest {
	s.InstanceAlias = &v
	return s
}

func (s *UpdateInstanceRequest) SetInstanceId(v string) *UpdateInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceRequest) SetInstanceSource(v string) *UpdateInstanceRequest {
	s.InstanceSource = &v
	return s
}

func (s *UpdateInstanceRequest) SetInstanceType(v string) *UpdateInstanceRequest {
	s.InstanceType = &v
	return s
}

func (s *UpdateInstanceRequest) SetPort(v int32) *UpdateInstanceRequest {
	s.Port = &v
	return s
}

func (s *UpdateInstanceRequest) SetQueryTimeout(v int32) *UpdateInstanceRequest {
	s.QueryTimeout = &v
	return s
}

func (s *UpdateInstanceRequest) SetSafeRuleId(v string) *UpdateInstanceRequest {
	s.SafeRuleId = &v
	return s
}

func (s *UpdateInstanceRequest) SetSid(v string) *UpdateInstanceRequest {
	s.Sid = &v
	return s
}

func (s *UpdateInstanceRequest) SetSkipTest(v bool) *UpdateInstanceRequest {
	s.SkipTest = &v
	return s
}

func (s *UpdateInstanceRequest) SetTemplateId(v int64) *UpdateInstanceRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateInstanceRequest) SetTemplateType(v string) *UpdateInstanceRequest {
	s.TemplateType = &v
	return s
}

func (s *UpdateInstanceRequest) SetTid(v int64) *UpdateInstanceRequest {
	s.Tid = &v
	return s
}

func (s *UpdateInstanceRequest) SetUseDsql(v int32) *UpdateInstanceRequest {
	s.UseDsql = &v
	return s
}

func (s *UpdateInstanceRequest) SetVpcId(v string) *UpdateInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *UpdateInstanceRequest) Validate() error {
	return dara.Validate(s)
}
