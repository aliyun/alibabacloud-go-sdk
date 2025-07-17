// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataLinkName(v string) *RegisterInstanceRequest
	GetDataLinkName() *string
	SetDatabasePassword(v string) *RegisterInstanceRequest
	GetDatabasePassword() *string
	SetDatabaseUser(v string) *RegisterInstanceRequest
	GetDatabaseUser() *string
	SetDbaUid(v int64) *RegisterInstanceRequest
	GetDbaUid() *int64
	SetDbaUidByString(v string) *RegisterInstanceRequest
	GetDbaUidByString() *string
	SetDdlOnline(v int32) *RegisterInstanceRequest
	GetDdlOnline() *int32
	SetEcsInstanceId(v string) *RegisterInstanceRequest
	GetEcsInstanceId() *string
	SetEcsRegion(v string) *RegisterInstanceRequest
	GetEcsRegion() *string
	SetEnableSellSitd(v string) *RegisterInstanceRequest
	GetEnableSellSitd() *string
	SetEnvType(v string) *RegisterInstanceRequest
	GetEnvType() *string
	SetExportTimeout(v int32) *RegisterInstanceRequest
	GetExportTimeout() *int32
	SetHost(v string) *RegisterInstanceRequest
	GetHost() *string
	SetInstanceAlias(v string) *RegisterInstanceRequest
	GetInstanceAlias() *string
	SetInstanceSource(v string) *RegisterInstanceRequest
	GetInstanceSource() *string
	SetInstanceType(v string) *RegisterInstanceRequest
	GetInstanceType() *string
	SetNetworkType(v string) *RegisterInstanceRequest
	GetNetworkType() *string
	SetPort(v int32) *RegisterInstanceRequest
	GetPort() *int32
	SetQueryTimeout(v int32) *RegisterInstanceRequest
	GetQueryTimeout() *int32
	SetResourceGroup(v string) *RegisterInstanceRequest
	GetResourceGroup() *string
	SetSafeRule(v string) *RegisterInstanceRequest
	GetSafeRule() *string
	SetSid(v string) *RegisterInstanceRequest
	GetSid() *string
	SetSkipTest(v bool) *RegisterInstanceRequest
	GetSkipTest() *bool
	SetTemplateId(v int64) *RegisterInstanceRequest
	GetTemplateId() *int64
	SetTemplateType(v string) *RegisterInstanceRequest
	GetTemplateType() *string
	SetTid(v int64) *RegisterInstanceRequest
	GetTid() *int64
	SetUseDsql(v int32) *RegisterInstanceRequest
	GetUseDsql() *int32
	SetVpcId(v string) *RegisterInstanceRequest
	GetVpcId() *string
}

type RegisterInstanceRequest struct {
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
	// dblink_test
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
	// dmstest
	DatabaseUser *string `json:"DatabaseUser,omitempty" xml:"DatabaseUser,omitempty"`
	// The ID of the user who assumes the DBA role of the database instance. You can call the [ListUsers](https://help.aliyun.com/document_detail/141938.html) or [GetInstance](https://help.aliyun.com/document_detail/141567.html) operation to query the user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 22275482072787****
	DbaUid *int64 `json:"DbaUid,omitempty" xml:"DbaUid,omitempty"`
	// The ID of the user who assumes the DBA role of the database instance. If the user ID is a non-numeric value such as a role or an account, you can use this parameter to replace DbaUid.
	//
	// example:
	//
	// 22275482072787****
	DbaUidByString *string `json:"DbaUidByString,omitempty" xml:"DbaUidByString,omitempty"`
	// Specifies whether to enable the lock-free schema change feature for the database instance. Valid values:
	//
	// 	- **0**: disables the lock-free schema change feature.
	//
	// 	- **1**: uses the online DDL of MySQL first.
	//
	// 	- **2**: uses the lock-free schema change feature of DMS first.
	//
	// > Supported database types: ApsaraDB RDS for MySQL, PolarDB for MySQL, ApsaraDB MyBase for MySQL, and third-party MySQL databases.
	//
	// example:
	//
	// 2
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
	// > This parameter is required if the InstanceSource parameter is set to RDS, ECS_OWN, or VPC_IDC.
	//
	// example:
	//
	// cn-hangzhou
	EcsRegion *string `json:"EcsRegion,omitempty" xml:"EcsRegion,omitempty"`
	// 	- **Y:*	- enables the sensitive data protection feature
	//
	// 	- **NULL or other:*	- disables the sensitive data protection feature
	//
	// example:
	//
	// Y
	EnableSellSitd *string `json:"EnableSellSitd,omitempty" xml:"EnableSellSitd,omitempty"`
	// The type of the environment in which the database instance is deployed. Valid values:
	//
	// 	- product: production environment
	//
	// 	- dev: development environment
	//
	// 	- pre: pre-release environment
	//
	// 	- test: test environment
	//
	// 	- sit: system integration testing (SIT) environment
	//
	// 	- uat: user acceptance testing (UAT) environment
	//
	// 	- pet: stress testing environment
	//
	// 	- stag: staging environment
	//
	// This parameter is required.
	//
	// example:
	//
	// product
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The timeout period for exporting data from the database instance. Unit: seconds.
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
	// 172.XX.XXX.254
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The alias of the database instance. Specify an alias that can help you identify the database instance in DMS.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test instance
	InstanceAlias *string `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
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
	// RDS
	InstanceSource *string `json:"InstanceSource,omitempty" xml:"InstanceSource,omitempty"`
	// The type of the database. For more information about the valid values of this parameter, see [DbType parameter](https://help.aliyun.com/document_detail/198106.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// MySQL
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The network type of the database instance. Valid values:
	//
	// 	- **CLASSIC:*	- classic network
	//
	// 	- **VPC:*	- VPC
	//
	// This parameter is required.
	//
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The port that is used to connect to the database instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The timeout period for querying data in the database instance. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	QueryTimeout  *int32  `json:"QueryTimeout,omitempty" xml:"QueryTimeout,omitempty"`
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// The name of the security rule set (GroupName) for the database instance. You can call the [ListStandardGroups](https://help.aliyun.com/document_detail/417891.html) or [GetInstance](https://help.aliyun.com/document_detail/141567.html) operation to query the name of the security rule set.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test
	SafeRule *string `json:"SafeRule,omitempty" xml:"SafeRule,omitempty"`
	// The system ID (SID) of the database.
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
	// true
	SkipTest *bool `json:"SkipTest,omitempty" xml:"SkipTest,omitempty"`
	// The ID of the classification template. You can call the [ListClassificationTemplates](https://help.aliyun.com/document_detail/460613.html) operation to query the template ID.
	//
	// example:
	//
	// 31***
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
	// 1
	UseDsql *int32 `json:"UseDsql,omitempty" xml:"UseDsql,omitempty"`
	// The ID of the VPC to which the database instance belongs.
	//
	// > This parameter is required if the InstanceSource parameter is set to VPC_IDC.
	//
	// example:
	//
	// vpc-xxxxxxxxxxxxxxxxxxxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s RegisterInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RegisterInstanceRequest) GoString() string {
	return s.String()
}

func (s *RegisterInstanceRequest) GetDataLinkName() *string {
	return s.DataLinkName
}

func (s *RegisterInstanceRequest) GetDatabasePassword() *string {
	return s.DatabasePassword
}

func (s *RegisterInstanceRequest) GetDatabaseUser() *string {
	return s.DatabaseUser
}

func (s *RegisterInstanceRequest) GetDbaUid() *int64 {
	return s.DbaUid
}

func (s *RegisterInstanceRequest) GetDbaUidByString() *string {
	return s.DbaUidByString
}

func (s *RegisterInstanceRequest) GetDdlOnline() *int32 {
	return s.DdlOnline
}

func (s *RegisterInstanceRequest) GetEcsInstanceId() *string {
	return s.EcsInstanceId
}

func (s *RegisterInstanceRequest) GetEcsRegion() *string {
	return s.EcsRegion
}

func (s *RegisterInstanceRequest) GetEnableSellSitd() *string {
	return s.EnableSellSitd
}

func (s *RegisterInstanceRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *RegisterInstanceRequest) GetExportTimeout() *int32 {
	return s.ExportTimeout
}

func (s *RegisterInstanceRequest) GetHost() *string {
	return s.Host
}

func (s *RegisterInstanceRequest) GetInstanceAlias() *string {
	return s.InstanceAlias
}

func (s *RegisterInstanceRequest) GetInstanceSource() *string {
	return s.InstanceSource
}

func (s *RegisterInstanceRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *RegisterInstanceRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *RegisterInstanceRequest) GetPort() *int32 {
	return s.Port
}

func (s *RegisterInstanceRequest) GetQueryTimeout() *int32 {
	return s.QueryTimeout
}

func (s *RegisterInstanceRequest) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *RegisterInstanceRequest) GetSafeRule() *string {
	return s.SafeRule
}

func (s *RegisterInstanceRequest) GetSid() *string {
	return s.Sid
}

func (s *RegisterInstanceRequest) GetSkipTest() *bool {
	return s.SkipTest
}

func (s *RegisterInstanceRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *RegisterInstanceRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *RegisterInstanceRequest) GetTid() *int64 {
	return s.Tid
}

func (s *RegisterInstanceRequest) GetUseDsql() *int32 {
	return s.UseDsql
}

func (s *RegisterInstanceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *RegisterInstanceRequest) SetDataLinkName(v string) *RegisterInstanceRequest {
	s.DataLinkName = &v
	return s
}

func (s *RegisterInstanceRequest) SetDatabasePassword(v string) *RegisterInstanceRequest {
	s.DatabasePassword = &v
	return s
}

func (s *RegisterInstanceRequest) SetDatabaseUser(v string) *RegisterInstanceRequest {
	s.DatabaseUser = &v
	return s
}

func (s *RegisterInstanceRequest) SetDbaUid(v int64) *RegisterInstanceRequest {
	s.DbaUid = &v
	return s
}

func (s *RegisterInstanceRequest) SetDbaUidByString(v string) *RegisterInstanceRequest {
	s.DbaUidByString = &v
	return s
}

func (s *RegisterInstanceRequest) SetDdlOnline(v int32) *RegisterInstanceRequest {
	s.DdlOnline = &v
	return s
}

func (s *RegisterInstanceRequest) SetEcsInstanceId(v string) *RegisterInstanceRequest {
	s.EcsInstanceId = &v
	return s
}

func (s *RegisterInstanceRequest) SetEcsRegion(v string) *RegisterInstanceRequest {
	s.EcsRegion = &v
	return s
}

func (s *RegisterInstanceRequest) SetEnableSellSitd(v string) *RegisterInstanceRequest {
	s.EnableSellSitd = &v
	return s
}

func (s *RegisterInstanceRequest) SetEnvType(v string) *RegisterInstanceRequest {
	s.EnvType = &v
	return s
}

func (s *RegisterInstanceRequest) SetExportTimeout(v int32) *RegisterInstanceRequest {
	s.ExportTimeout = &v
	return s
}

func (s *RegisterInstanceRequest) SetHost(v string) *RegisterInstanceRequest {
	s.Host = &v
	return s
}

func (s *RegisterInstanceRequest) SetInstanceAlias(v string) *RegisterInstanceRequest {
	s.InstanceAlias = &v
	return s
}

func (s *RegisterInstanceRequest) SetInstanceSource(v string) *RegisterInstanceRequest {
	s.InstanceSource = &v
	return s
}

func (s *RegisterInstanceRequest) SetInstanceType(v string) *RegisterInstanceRequest {
	s.InstanceType = &v
	return s
}

func (s *RegisterInstanceRequest) SetNetworkType(v string) *RegisterInstanceRequest {
	s.NetworkType = &v
	return s
}

func (s *RegisterInstanceRequest) SetPort(v int32) *RegisterInstanceRequest {
	s.Port = &v
	return s
}

func (s *RegisterInstanceRequest) SetQueryTimeout(v int32) *RegisterInstanceRequest {
	s.QueryTimeout = &v
	return s
}

func (s *RegisterInstanceRequest) SetResourceGroup(v string) *RegisterInstanceRequest {
	s.ResourceGroup = &v
	return s
}

func (s *RegisterInstanceRequest) SetSafeRule(v string) *RegisterInstanceRequest {
	s.SafeRule = &v
	return s
}

func (s *RegisterInstanceRequest) SetSid(v string) *RegisterInstanceRequest {
	s.Sid = &v
	return s
}

func (s *RegisterInstanceRequest) SetSkipTest(v bool) *RegisterInstanceRequest {
	s.SkipTest = &v
	return s
}

func (s *RegisterInstanceRequest) SetTemplateId(v int64) *RegisterInstanceRequest {
	s.TemplateId = &v
	return s
}

func (s *RegisterInstanceRequest) SetTemplateType(v string) *RegisterInstanceRequest {
	s.TemplateType = &v
	return s
}

func (s *RegisterInstanceRequest) SetTid(v int64) *RegisterInstanceRequest {
	s.Tid = &v
	return s
}

func (s *RegisterInstanceRequest) SetUseDsql(v int32) *RegisterInstanceRequest {
	s.UseDsql = &v
	return s
}

func (s *RegisterInstanceRequest) SetVpcId(v string) *RegisterInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *RegisterInstanceRequest) Validate() error {
	return dara.Validate(s)
}
