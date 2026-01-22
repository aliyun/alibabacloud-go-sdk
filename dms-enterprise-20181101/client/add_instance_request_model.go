// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataLinkName(v string) *AddInstanceRequest
	GetDataLinkName() *string
	SetDatabasePassword(v string) *AddInstanceRequest
	GetDatabasePassword() *string
	SetDatabaseUser(v string) *AddInstanceRequest
	GetDatabaseUser() *string
	SetDbaId(v int64) *AddInstanceRequest
	GetDbaId() *int64
	SetDdlOnline(v int32) *AddInstanceRequest
	GetDdlOnline() *int32
	SetEcsInstanceId(v string) *AddInstanceRequest
	GetEcsInstanceId() *string
	SetEcsRegion(v string) *AddInstanceRequest
	GetEcsRegion() *string
	SetEnableSellCommon(v string) *AddInstanceRequest
	GetEnableSellCommon() *string
	SetEnableSellSitd(v string) *AddInstanceRequest
	GetEnableSellSitd() *string
	SetEnableSellStable(v string) *AddInstanceRequest
	GetEnableSellStable() *string
	SetEnableSellTrust(v string) *AddInstanceRequest
	GetEnableSellTrust() *string
	SetEnvType(v string) *AddInstanceRequest
	GetEnvType() *string
	SetExportTimeout(v int32) *AddInstanceRequest
	GetExportTimeout() *int32
	SetHost(v string) *AddInstanceRequest
	GetHost() *string
	SetInstanceAlias(v string) *AddInstanceRequest
	GetInstanceAlias() *string
	SetInstanceSource(v string) *AddInstanceRequest
	GetInstanceSource() *string
	SetInstanceType(v string) *AddInstanceRequest
	GetInstanceType() *string
	SetNetworkType(v string) *AddInstanceRequest
	GetNetworkType() *string
	SetPort(v int32) *AddInstanceRequest
	GetPort() *int32
	SetQueryTimeout(v int32) *AddInstanceRequest
	GetQueryTimeout() *int32
	SetRoleArn(v string) *AddInstanceRequest
	GetRoleArn() *string
	SetSafeRule(v string) *AddInstanceRequest
	GetSafeRule() *string
	SetSid(v string) *AddInstanceRequest
	GetSid() *string
	SetSkipTest(v bool) *AddInstanceRequest
	GetSkipTest() *bool
	SetTemplateId(v int64) *AddInstanceRequest
	GetTemplateId() *int64
	SetTemplateType(v string) *AddInstanceRequest
	GetTemplateType() *string
	SetTid(v int64) *AddInstanceRequest
	GetTid() *int64
	SetUseDsql(v int32) *AddInstanceRequest
	GetUseDsql() *int32
	SetUseSsl(v int32) *AddInstanceRequest
	GetUseSsl() *int32
	SetVpcId(v string) *AddInstanceRequest
	GetVpcId() *string
}

type AddInstanceRequest struct {
	// The name of the database link for cross-database queries.
	//
	// >
	//
	// 	- This property must be specified when UseDsql is set to 1.
	//
	// 	- The name can contain only lowercase letters and underscores (_).
	//
	// 	- The name must be unique within a tenant.
	//
	// example:
	//
	// dblink_test
	DataLinkName *string `json:"DataLinkName,omitempty" xml:"DataLinkName,omitempty"`
	// The password of the database account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test***
	DatabasePassword *string `json:"DatabasePassword,omitempty" xml:"DatabasePassword,omitempty"`
	// The name of the database account.
	//
	// This parameter is required.
	//
	// example:
	//
	// testsdb
	DatabaseUser *string `json:"DatabaseUser,omitempty" xml:"DatabaseUser,omitempty"`
	// The ID of the user who assumes the database administrator (DBA) role. You can call the [ListUsers](https://help.aliyun.com/document_detail/141938.html) or [GetInstance](https://help.aliyun.com/document_detail/141567.html) operation to obtain the value of this parameter.
	//
	// example:
	//
	// 27****
	DbaId *int64 `json:"DbaId,omitempty" xml:"DbaId,omitempty"`
	// Specifies whether to enable Lock-free Schema Change. Valid values:
	//
	// 	- **0**: does not enable lock-free schema change.
	//
	// 	- **1**: uses the native online DDL operations of MySQL first.
	//
	// 	- **2:*	- uses lock-free schema change first.
	//
	// > Supported databases include ApsaraDB RDS for MySQL, PolarDB for MySQL, ApsaraDB MyBase for MySQL, and third-party MySQL databases.
	//
	// example:
	//
	// 2
	DdlOnline *int32 `json:"DdlOnline,omitempty" xml:"DdlOnline,omitempty"`
	// The ID of the instance. If your instance is a database instance connected by using a database gateway, specify the gateway ID for this parameter.
	//
	// > This parameter is required if InstanceSource is set to ECS_OWN or GATEWAY.
	//
	// example:
	//
	// i-2zei9gs1t7h8l7ac****
	EcsInstanceId *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	// The region in which the ECS instance resides.
	//
	// > This parameter is required if InstanceSource is set to RDS, ECS_OWN, GATEWAY, or VPC_IDC.
	//
	// example:
	//
	// cn-hangzhou
	EcsRegion *string `json:"EcsRegion,omitempty" xml:"EcsRegion,omitempty"`
	// Specifies whether to enable Security Collaboration for the database instance. Valid values:
	//
	// 	- Y: enables Security Collaboration.
	//
	// 	- N: disables Security Collaboration.
	//
	// example:
	//
	// Y
	EnableSellCommon *string `json:"EnableSellCommon,omitempty" xml:"EnableSellCommon,omitempty"`
	// Specifies whether to enable sensitive data protection. Valid values:
	//
	// 	- Y: enables the sensitive data protection feature for the database instance.
	//
	// 	- N: disables the sensitive data protection feature for the database instance.
	//
	// example:
	//
	// Y
	EnableSellSitd *string `json:"EnableSellSitd,omitempty" xml:"EnableSellSitd,omitempty"`
	// Specifies whether to enable Stable Change for the database instance. Valid values:
	//
	// 	- Y: Enables Stable Change.
	//
	// 	- N: Disables Stable Change.
	//
	// example:
	//
	// NULL
	EnableSellStable *string `json:"EnableSellStable,omitempty" xml:"EnableSellStable,omitempty"`
	// Specifies whether to enable the security hosting feature for the database instance. Valid values:
	//
	// 	- Y: enables the security hosting feature for the database instance.
	//
	// 	- N: disables the security hosting feature for the database instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// Y
	EnableSellTrust *string `json:"EnableSellTrust,omitempty" xml:"EnableSellTrust,omitempty"`
	// The type of the environment to which the database instance belongs. Valid values:
	//
	// 	- **product:*	- the production environment.
	//
	// 	- **dev**: development environment.
	//
	// 	- **pre**: pre-release environment.
	//
	// 	- **test**: test environment.
	//
	// 	- **sit**: system integration testing (SIT) environment.
	//
	// 	- **uat**: user acceptance testing (UAT) environment.
	//
	// 	- **pet**: stress testing environment.
	//
	// 	- **stag**: staging environment.
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
	// 86400
	ExportTimeout *int32 `json:"ExportTimeout,omitempty" xml:"ExportTimeout,omitempty"`
	// The endpoint that is used to connect to the database.
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
	InstanceAlias *string `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	// The source of the database instance. Valid values:
	//
	// 	- **PUBLIC_OWN**: a self-managed database instance that is deployed on the Internet.
	//
	// 	- **RDS**: an ApsaraDB RDS instance.
	//
	// 	- **ECS_OWN**: a self-managed database instance that is deployed on an Elastic Compute Service (ECS) instance.
	//
	// 	- **VPC_IDC**: a self-managed database instance that is deployed in a data center connected over a virtual private cloud (VPC).
	//
	// 	- **GATEWAY**: a database instance connected by using a database gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// RDS
	InstanceSource *string `json:"InstanceSource,omitempty" xml:"InstanceSource,omitempty"`
	// The type of the database instance. For more information about the valid values of this parameter, see [DbType parameter](https://help.aliyun.com/document_detail/198106.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// MySQL
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The network type. Valid values:
	//
	// 	- **CLASSIC:*	- the classic network.
	//
	// 	- **VPC**
	//
	// This parameter is required.
	//
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The port that is used to connect to the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The timeout period for querying data from the database instance. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7200
	QueryTimeout *int32  `json:"QueryTimeout,omitempty" xml:"QueryTimeout,omitempty"`
	RoleArn      *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// The name of the security rule set for the database instance. This parameter is required if Security Collaboration is enabled. You can call the[ListStandardGroups](https://help.aliyun.com/document_detail/465940.html) or [GetInstance](https://help.aliyun.com/document_detail/465826.html) operation to obtain the name of the security rule set from GroupName.
	SafeRule *string `json:"SafeRule,omitempty" xml:"SafeRule,omitempty"`
	// The system ID (SID) of the database instance.
	//
	// > This parameter is required if InstanceType is set to ORACLE.
	//
	// example:
	//
	// testSid
	Sid *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	// Specifies whether to skip the connectivity test. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	SkipTest *bool `json:"SkipTest,omitempty" xml:"SkipTest,omitempty"`
	// The ID of the classification and grading template. You can call the [ListClassificationTemplates](https://help.aliyun.com/document_detail/465947.html) operation to query the template ID.
	//
	// example:
	//
	// 31***
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The type of the classification and grading template. You can call the [ListClassificationTemplates](https://help.aliyun.com/document_detail/465947.html) operation to query the template type.
	//
	// example:
	//
	// INNER
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// The ID of the tenant.
	//
	// > You can move the pointer over the profile picture in the upper-right corner of the DMS console to obtain the tenant ID.
	//
	// example:
	//
	// 23****
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// Specifies whether to enable cross-database query for the database instance. Valid values:
	//
	// 	- **0: does not enable cross-database query.**
	//
	// 	- **1**: enables cross-database query.
	//
	// > Supported databases include MySQL, SQL Server, PostgreSQL, PolarDB for Oracle, and Redis.
	//
	// example:
	//
	// 1
	UseDsql *int32 `json:"UseDsql,omitempty" xml:"UseDsql,omitempty"`
	// Specifies whether to allow Data Management Service (DMS) to connect to the database instance by using SSL connections. Before you use SSL connections, make sure that the SSL encryption feature is enabled for the database instance. Valid values:
	//
	// 	- **0*	- (default): DMS automatically checks whether self-negotiation is enabled for the database instance. DMS automatically checks whether the SSL encryption feature is enabled for the database instance. If the SSL encryption feature is enabled, DMS connects to the database instance by using SSL connections. Otherwise, DMS connects to the database instance without encryption.
	//
	// 	- **1**: DMS connects to the database instance by using SSL connections. This value is invalid if the SSL encryption feature is disabled for the database instance.
	//
	// 	- **-1**: DMS does not connect to the database instance by using SSL connections.
	//
	// >
	//
	// 	- This parameter is available only for a MySQL or Redis database instance.
	//
	// 	- SSL encrypts network connections at the transport layer to improve the security and integrity of data in transmission. However, SSL increases the response time of network connections.
	//
	// example:
	//
	// 0
	UseSsl *int32 `json:"UseSsl,omitempty" xml:"UseSsl,omitempty"`
	// The ID of the instance connected over a VPC.
	//
	// > This parameter is required if InstanceSource is set to VPC_IDC.
	//
	// example:
	//
	// vpc-2zef4o1hu7ljd****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s AddInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s AddInstanceRequest) GoString() string {
	return s.String()
}

func (s *AddInstanceRequest) GetDataLinkName() *string {
	return s.DataLinkName
}

func (s *AddInstanceRequest) GetDatabasePassword() *string {
	return s.DatabasePassword
}

func (s *AddInstanceRequest) GetDatabaseUser() *string {
	return s.DatabaseUser
}

func (s *AddInstanceRequest) GetDbaId() *int64 {
	return s.DbaId
}

func (s *AddInstanceRequest) GetDdlOnline() *int32 {
	return s.DdlOnline
}

func (s *AddInstanceRequest) GetEcsInstanceId() *string {
	return s.EcsInstanceId
}

func (s *AddInstanceRequest) GetEcsRegion() *string {
	return s.EcsRegion
}

func (s *AddInstanceRequest) GetEnableSellCommon() *string {
	return s.EnableSellCommon
}

func (s *AddInstanceRequest) GetEnableSellSitd() *string {
	return s.EnableSellSitd
}

func (s *AddInstanceRequest) GetEnableSellStable() *string {
	return s.EnableSellStable
}

func (s *AddInstanceRequest) GetEnableSellTrust() *string {
	return s.EnableSellTrust
}

func (s *AddInstanceRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *AddInstanceRequest) GetExportTimeout() *int32 {
	return s.ExportTimeout
}

func (s *AddInstanceRequest) GetHost() *string {
	return s.Host
}

func (s *AddInstanceRequest) GetInstanceAlias() *string {
	return s.InstanceAlias
}

func (s *AddInstanceRequest) GetInstanceSource() *string {
	return s.InstanceSource
}

func (s *AddInstanceRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *AddInstanceRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *AddInstanceRequest) GetPort() *int32 {
	return s.Port
}

func (s *AddInstanceRequest) GetQueryTimeout() *int32 {
	return s.QueryTimeout
}

func (s *AddInstanceRequest) GetRoleArn() *string {
	return s.RoleArn
}

func (s *AddInstanceRequest) GetSafeRule() *string {
	return s.SafeRule
}

func (s *AddInstanceRequest) GetSid() *string {
	return s.Sid
}

func (s *AddInstanceRequest) GetSkipTest() *bool {
	return s.SkipTest
}

func (s *AddInstanceRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *AddInstanceRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *AddInstanceRequest) GetTid() *int64 {
	return s.Tid
}

func (s *AddInstanceRequest) GetUseDsql() *int32 {
	return s.UseDsql
}

func (s *AddInstanceRequest) GetUseSsl() *int32 {
	return s.UseSsl
}

func (s *AddInstanceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *AddInstanceRequest) SetDataLinkName(v string) *AddInstanceRequest {
	s.DataLinkName = &v
	return s
}

func (s *AddInstanceRequest) SetDatabasePassword(v string) *AddInstanceRequest {
	s.DatabasePassword = &v
	return s
}

func (s *AddInstanceRequest) SetDatabaseUser(v string) *AddInstanceRequest {
	s.DatabaseUser = &v
	return s
}

func (s *AddInstanceRequest) SetDbaId(v int64) *AddInstanceRequest {
	s.DbaId = &v
	return s
}

func (s *AddInstanceRequest) SetDdlOnline(v int32) *AddInstanceRequest {
	s.DdlOnline = &v
	return s
}

func (s *AddInstanceRequest) SetEcsInstanceId(v string) *AddInstanceRequest {
	s.EcsInstanceId = &v
	return s
}

func (s *AddInstanceRequest) SetEcsRegion(v string) *AddInstanceRequest {
	s.EcsRegion = &v
	return s
}

func (s *AddInstanceRequest) SetEnableSellCommon(v string) *AddInstanceRequest {
	s.EnableSellCommon = &v
	return s
}

func (s *AddInstanceRequest) SetEnableSellSitd(v string) *AddInstanceRequest {
	s.EnableSellSitd = &v
	return s
}

func (s *AddInstanceRequest) SetEnableSellStable(v string) *AddInstanceRequest {
	s.EnableSellStable = &v
	return s
}

func (s *AddInstanceRequest) SetEnableSellTrust(v string) *AddInstanceRequest {
	s.EnableSellTrust = &v
	return s
}

func (s *AddInstanceRequest) SetEnvType(v string) *AddInstanceRequest {
	s.EnvType = &v
	return s
}

func (s *AddInstanceRequest) SetExportTimeout(v int32) *AddInstanceRequest {
	s.ExportTimeout = &v
	return s
}

func (s *AddInstanceRequest) SetHost(v string) *AddInstanceRequest {
	s.Host = &v
	return s
}

func (s *AddInstanceRequest) SetInstanceAlias(v string) *AddInstanceRequest {
	s.InstanceAlias = &v
	return s
}

func (s *AddInstanceRequest) SetInstanceSource(v string) *AddInstanceRequest {
	s.InstanceSource = &v
	return s
}

func (s *AddInstanceRequest) SetInstanceType(v string) *AddInstanceRequest {
	s.InstanceType = &v
	return s
}

func (s *AddInstanceRequest) SetNetworkType(v string) *AddInstanceRequest {
	s.NetworkType = &v
	return s
}

func (s *AddInstanceRequest) SetPort(v int32) *AddInstanceRequest {
	s.Port = &v
	return s
}

func (s *AddInstanceRequest) SetQueryTimeout(v int32) *AddInstanceRequest {
	s.QueryTimeout = &v
	return s
}

func (s *AddInstanceRequest) SetRoleArn(v string) *AddInstanceRequest {
	s.RoleArn = &v
	return s
}

func (s *AddInstanceRequest) SetSafeRule(v string) *AddInstanceRequest {
	s.SafeRule = &v
	return s
}

func (s *AddInstanceRequest) SetSid(v string) *AddInstanceRequest {
	s.Sid = &v
	return s
}

func (s *AddInstanceRequest) SetSkipTest(v bool) *AddInstanceRequest {
	s.SkipTest = &v
	return s
}

func (s *AddInstanceRequest) SetTemplateId(v int64) *AddInstanceRequest {
	s.TemplateId = &v
	return s
}

func (s *AddInstanceRequest) SetTemplateType(v string) *AddInstanceRequest {
	s.TemplateType = &v
	return s
}

func (s *AddInstanceRequest) SetTid(v int64) *AddInstanceRequest {
	s.Tid = &v
	return s
}

func (s *AddInstanceRequest) SetUseDsql(v int32) *AddInstanceRequest {
	s.UseDsql = &v
	return s
}

func (s *AddInstanceRequest) SetUseSsl(v int32) *AddInstanceRequest {
	s.UseSsl = &v
	return s
}

func (s *AddInstanceRequest) SetVpcId(v string) *AddInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *AddInstanceRequest) Validate() error {
	return dara.Validate(s)
}
