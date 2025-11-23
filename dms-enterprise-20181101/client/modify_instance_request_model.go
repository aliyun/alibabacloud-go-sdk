// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataLinkName(v string) *ModifyInstanceRequest
	GetDataLinkName() *string
	SetDatabasePassword(v string) *ModifyInstanceRequest
	GetDatabasePassword() *string
	SetDatabaseUser(v string) *ModifyInstanceRequest
	GetDatabaseUser() *string
	SetDbaId(v int64) *ModifyInstanceRequest
	GetDbaId() *int64
	SetDdlOnline(v int32) *ModifyInstanceRequest
	GetDdlOnline() *int32
	SetEcsInstanceId(v string) *ModifyInstanceRequest
	GetEcsInstanceId() *string
	SetEcsRegion(v string) *ModifyInstanceRequest
	GetEcsRegion() *string
	SetEnableSellCommon(v string) *ModifyInstanceRequest
	GetEnableSellCommon() *string
	SetEnableSellSitd(v string) *ModifyInstanceRequest
	GetEnableSellSitd() *string
	SetEnableSellStable(v string) *ModifyInstanceRequest
	GetEnableSellStable() *string
	SetEnableSellTrust(v string) *ModifyInstanceRequest
	GetEnableSellTrust() *string
	SetEnvType(v string) *ModifyInstanceRequest
	GetEnvType() *string
	SetExportTimeout(v int32) *ModifyInstanceRequest
	GetExportTimeout() *int32
	SetHost(v string) *ModifyInstanceRequest
	GetHost() *string
	SetInstanceAlias(v string) *ModifyInstanceRequest
	GetInstanceAlias() *string
	SetInstanceId(v string) *ModifyInstanceRequest
	GetInstanceId() *string
	SetInstanceSource(v string) *ModifyInstanceRequest
	GetInstanceSource() *string
	SetInstanceType(v string) *ModifyInstanceRequest
	GetInstanceType() *string
	SetNetworkType(v string) *ModifyInstanceRequest
	GetNetworkType() *string
	SetPort(v int32) *ModifyInstanceRequest
	GetPort() *int32
	SetQueryTimeout(v int32) *ModifyInstanceRequest
	GetQueryTimeout() *int32
	SetSafeRule(v string) *ModifyInstanceRequest
	GetSafeRule() *string
	SetSid(v string) *ModifyInstanceRequest
	GetSid() *string
	SetSkipTest(v bool) *ModifyInstanceRequest
	GetSkipTest() *bool
	SetTemplateId(v int64) *ModifyInstanceRequest
	GetTemplateId() *int64
	SetTemplateType(v string) *ModifyInstanceRequest
	GetTemplateType() *string
	SetTid(v int64) *ModifyInstanceRequest
	GetTid() *int64
	SetUseDsql(v int32) *ModifyInstanceRequest
	GetUseDsql() *int32
	SetUseSsl(v int32) *ModifyInstanceRequest
	GetUseSsl() *int32
	SetVpcId(v string) *ModifyInstanceRequest
	GetVpcId() *string
}

type ModifyInstanceRequest struct {
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
	// The password of the account that is used to log on to the database instance.
	//
	// example:
	//
	// test***
	DatabasePassword *string `json:"DatabasePassword,omitempty" xml:"DatabasePassword,omitempty"`
	// The account that is used to log on to the database instance.
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
	// Specifies whether to enable lock-free schema change. Valid values:
	//
	// 	- **0**: Disable Lock-free Schema Change.
	//
	// 	- **1**: MySQL native online DDL first.
	//
	// 	- **2**: DMS native online DDL first.
	//
	// > Supported databases include ApsaraDB RDS for MySQL, PolarDB for MySQL, ApsaraDB MyBase for MySQL, and third-party MySQL databases.
	//
	// example:
	//
	// 2
	DdlOnline *int32 `json:"DdlOnline,omitempty" xml:"DdlOnline,omitempty"`
	// The ID of the ECS instance.
	//
	// >  This parameter is required if InstanceSource is set to ECS_OWN.
	//
	// example:
	//
	// i-2zei9gs1t7h8l7ac****
	EcsInstanceId *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	// The region in which the ECS instance resides.
	//
	// >  This parameter is required if InstanceSource is set to RDS, ECS_OWN, or VPC_IDC.
	//
	// example:
	//
	// cn-hangzhou
	EcsRegion *string `json:"EcsRegion,omitempty" xml:"EcsRegion,omitempty"`
	// Specifies whether to enable Security Collaboration for the database instance. Valid values:
	//
	// 	- Y: Enable.
	//
	// 	- N: Disable.
	//
	// example:
	//
	// Y
	EnableSellCommon *string `json:"EnableSellCommon,omitempty" xml:"EnableSellCommon,omitempty"`
	// Specifies whether to enable sensitive data protection. Valid values:
	//
	// 	- Y: Enable.
	//
	// 	- N: Disable.
	//
	// example:
	//
	// Y
	EnableSellSitd *string `json:"EnableSellSitd,omitempty" xml:"EnableSellSitd,omitempty"`
	// Specifies whether to enable Stable Change for the database instance. Valid values:
	//
	// 	- Y: Enable.
	//
	// 	- N: Disable.
	//
	// example:
	//
	// NULL
	EnableSellStable *string `json:"EnableSellStable,omitempty" xml:"EnableSellStable,omitempty"`
	// Specifies whether to enable the security hosting feature for the database instance. Valid values:
	//
	// 	- Y: Enable.
	//
	// 	- N: Disable.
	//
	// example:
	//
	// Y
	EnableSellTrust *string `json:"EnableSellTrust,omitempty" xml:"EnableSellTrust,omitempty"`
	// The type of the environment in which the database instance is deployed. Valid values:
	//
	// 	- **product**: production environment.
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
	// example:
	//
	// dev
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The timeout period for exporting data from the database instance. Unit: seconds.
	//
	// example:
	//
	// 86400
	ExportTimeout *int32 `json:"ExportTimeout,omitempty" xml:"ExportTimeout,omitempty"`
	// The endpoint that is used to connect to the database instance.
	//
	// example:
	//
	// 192.XXX.0.56
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The alias of the database instance. Specify an alias that can help you quickly identify the database instance in Data Management (DMS).
	//
	// example:
	//
	// instance_test
	InstanceAlias *string `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	// The ID of the instance. You can call the [ListInstances](https://help.aliyun.com/document_detail/141936.html) or [GetInstance](https://help.aliyun.com/document_detail/141567.html) operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 183****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	// example:
	//
	// RDS
	InstanceSource *string `json:"InstanceSource,omitempty" xml:"InstanceSource,omitempty"`
	// The type of the database instance. For more information about the valid values of this parameter, see [DbType parameter](https://help.aliyun.com/document_detail/198106.html).
	//
	// example:
	//
	// MySQL
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The network type of the database instance. Valid values:
	//
	// 	- **CLASSIC**: the classic network.
	//
	// 	- **VPC**: VPC.
	//
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The port that is used to connect to the database instance.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The timeout period for querying data from the database instance. Unit: seconds.
	//
	// example:
	//
	// 7200
	QueryTimeout *int32 `json:"QueryTimeout,omitempty" xml:"QueryTimeout,omitempty"`
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
	// Specifies whether to enable cross-instance query for the database instance. Valid values:
	//
	// 	- **0**: Disables cross-database query.
	//
	// 	- **1**: Enables cross-database query.
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
	// The VPC ID.
	//
	// >  This parameter is required if InstanceSource is set to VPC_IDC.
	//
	// example:
	//
	// vpc-bp10wnlcmor****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ModifyInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceRequest) GetDataLinkName() *string {
	return s.DataLinkName
}

func (s *ModifyInstanceRequest) GetDatabasePassword() *string {
	return s.DatabasePassword
}

func (s *ModifyInstanceRequest) GetDatabaseUser() *string {
	return s.DatabaseUser
}

func (s *ModifyInstanceRequest) GetDbaId() *int64 {
	return s.DbaId
}

func (s *ModifyInstanceRequest) GetDdlOnline() *int32 {
	return s.DdlOnline
}

func (s *ModifyInstanceRequest) GetEcsInstanceId() *string {
	return s.EcsInstanceId
}

func (s *ModifyInstanceRequest) GetEcsRegion() *string {
	return s.EcsRegion
}

func (s *ModifyInstanceRequest) GetEnableSellCommon() *string {
	return s.EnableSellCommon
}

func (s *ModifyInstanceRequest) GetEnableSellSitd() *string {
	return s.EnableSellSitd
}

func (s *ModifyInstanceRequest) GetEnableSellStable() *string {
	return s.EnableSellStable
}

func (s *ModifyInstanceRequest) GetEnableSellTrust() *string {
	return s.EnableSellTrust
}

func (s *ModifyInstanceRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *ModifyInstanceRequest) GetExportTimeout() *int32 {
	return s.ExportTimeout
}

func (s *ModifyInstanceRequest) GetHost() *string {
	return s.Host
}

func (s *ModifyInstanceRequest) GetInstanceAlias() *string {
	return s.InstanceAlias
}

func (s *ModifyInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceRequest) GetInstanceSource() *string {
	return s.InstanceSource
}

func (s *ModifyInstanceRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ModifyInstanceRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *ModifyInstanceRequest) GetPort() *int32 {
	return s.Port
}

func (s *ModifyInstanceRequest) GetQueryTimeout() *int32 {
	return s.QueryTimeout
}

func (s *ModifyInstanceRequest) GetSafeRule() *string {
	return s.SafeRule
}

func (s *ModifyInstanceRequest) GetSid() *string {
	return s.Sid
}

func (s *ModifyInstanceRequest) GetSkipTest() *bool {
	return s.SkipTest
}

func (s *ModifyInstanceRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *ModifyInstanceRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *ModifyInstanceRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ModifyInstanceRequest) GetUseDsql() *int32 {
	return s.UseDsql
}

func (s *ModifyInstanceRequest) GetUseSsl() *int32 {
	return s.UseSsl
}

func (s *ModifyInstanceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ModifyInstanceRequest) SetDataLinkName(v string) *ModifyInstanceRequest {
	s.DataLinkName = &v
	return s
}

func (s *ModifyInstanceRequest) SetDatabasePassword(v string) *ModifyInstanceRequest {
	s.DatabasePassword = &v
	return s
}

func (s *ModifyInstanceRequest) SetDatabaseUser(v string) *ModifyInstanceRequest {
	s.DatabaseUser = &v
	return s
}

func (s *ModifyInstanceRequest) SetDbaId(v int64) *ModifyInstanceRequest {
	s.DbaId = &v
	return s
}

func (s *ModifyInstanceRequest) SetDdlOnline(v int32) *ModifyInstanceRequest {
	s.DdlOnline = &v
	return s
}

func (s *ModifyInstanceRequest) SetEcsInstanceId(v string) *ModifyInstanceRequest {
	s.EcsInstanceId = &v
	return s
}

func (s *ModifyInstanceRequest) SetEcsRegion(v string) *ModifyInstanceRequest {
	s.EcsRegion = &v
	return s
}

func (s *ModifyInstanceRequest) SetEnableSellCommon(v string) *ModifyInstanceRequest {
	s.EnableSellCommon = &v
	return s
}

func (s *ModifyInstanceRequest) SetEnableSellSitd(v string) *ModifyInstanceRequest {
	s.EnableSellSitd = &v
	return s
}

func (s *ModifyInstanceRequest) SetEnableSellStable(v string) *ModifyInstanceRequest {
	s.EnableSellStable = &v
	return s
}

func (s *ModifyInstanceRequest) SetEnableSellTrust(v string) *ModifyInstanceRequest {
	s.EnableSellTrust = &v
	return s
}

func (s *ModifyInstanceRequest) SetEnvType(v string) *ModifyInstanceRequest {
	s.EnvType = &v
	return s
}

func (s *ModifyInstanceRequest) SetExportTimeout(v int32) *ModifyInstanceRequest {
	s.ExportTimeout = &v
	return s
}

func (s *ModifyInstanceRequest) SetHost(v string) *ModifyInstanceRequest {
	s.Host = &v
	return s
}

func (s *ModifyInstanceRequest) SetInstanceAlias(v string) *ModifyInstanceRequest {
	s.InstanceAlias = &v
	return s
}

func (s *ModifyInstanceRequest) SetInstanceId(v string) *ModifyInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceRequest) SetInstanceSource(v string) *ModifyInstanceRequest {
	s.InstanceSource = &v
	return s
}

func (s *ModifyInstanceRequest) SetInstanceType(v string) *ModifyInstanceRequest {
	s.InstanceType = &v
	return s
}

func (s *ModifyInstanceRequest) SetNetworkType(v string) *ModifyInstanceRequest {
	s.NetworkType = &v
	return s
}

func (s *ModifyInstanceRequest) SetPort(v int32) *ModifyInstanceRequest {
	s.Port = &v
	return s
}

func (s *ModifyInstanceRequest) SetQueryTimeout(v int32) *ModifyInstanceRequest {
	s.QueryTimeout = &v
	return s
}

func (s *ModifyInstanceRequest) SetSafeRule(v string) *ModifyInstanceRequest {
	s.SafeRule = &v
	return s
}

func (s *ModifyInstanceRequest) SetSid(v string) *ModifyInstanceRequest {
	s.Sid = &v
	return s
}

func (s *ModifyInstanceRequest) SetSkipTest(v bool) *ModifyInstanceRequest {
	s.SkipTest = &v
	return s
}

func (s *ModifyInstanceRequest) SetTemplateId(v int64) *ModifyInstanceRequest {
	s.TemplateId = &v
	return s
}

func (s *ModifyInstanceRequest) SetTemplateType(v string) *ModifyInstanceRequest {
	s.TemplateType = &v
	return s
}

func (s *ModifyInstanceRequest) SetTid(v int64) *ModifyInstanceRequest {
	s.Tid = &v
	return s
}

func (s *ModifyInstanceRequest) SetUseDsql(v int32) *ModifyInstanceRequest {
	s.UseDsql = &v
	return s
}

func (s *ModifyInstanceRequest) SetUseSsl(v int32) *ModifyInstanceRequest {
	s.UseSsl = &v
	return s
}

func (s *ModifyInstanceRequest) SetVpcId(v string) *ModifyInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *ModifyInstanceRequest) Validate() error {
	return dara.Validate(s)
}
