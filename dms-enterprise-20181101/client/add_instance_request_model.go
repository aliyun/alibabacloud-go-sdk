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
	// example:
	//
	// dblink_test
	DataLinkName *string `json:"DataLinkName,omitempty" xml:"DataLinkName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test***
	DatabasePassword *string `json:"DatabasePassword,omitempty" xml:"DatabasePassword,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testsdb
	DatabaseUser *string `json:"DatabaseUser,omitempty" xml:"DatabaseUser,omitempty"`
	// example:
	//
	// 27****
	DbaId *int64 `json:"DbaId,omitempty" xml:"DbaId,omitempty"`
	// example:
	//
	// 2
	DdlOnline *int32 `json:"DdlOnline,omitempty" xml:"DdlOnline,omitempty"`
	// example:
	//
	// i-2zei9gs1t7h8l7ac****
	EcsInstanceId *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	EcsRegion *string `json:"EcsRegion,omitempty" xml:"EcsRegion,omitempty"`
	// example:
	//
	// Y
	EnableSellCommon *string `json:"EnableSellCommon,omitempty" xml:"EnableSellCommon,omitempty"`
	// example:
	//
	// Y
	EnableSellSitd *string `json:"EnableSellSitd,omitempty" xml:"EnableSellSitd,omitempty"`
	// example:
	//
	// NULL
	EnableSellStable *string `json:"EnableSellStable,omitempty" xml:"EnableSellStable,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Y
	EnableSellTrust *string `json:"EnableSellTrust,omitempty" xml:"EnableSellTrust,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// product
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 86400
	ExportTimeout *int32 `json:"ExportTimeout,omitempty" xml:"ExportTimeout,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 172.XX.XXX.254
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// This parameter is required.
	InstanceAlias *string `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// RDS
	InstanceSource *string `json:"InstanceSource,omitempty" xml:"InstanceSource,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MySQL
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7200
	QueryTimeout *int32  `json:"QueryTimeout,omitempty" xml:"QueryTimeout,omitempty"`
	SafeRule     *string `json:"SafeRule,omitempty" xml:"SafeRule,omitempty"`
	// example:
	//
	// testSid
	Sid *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	// example:
	//
	// false
	SkipTest *bool `json:"SkipTest,omitempty" xml:"SkipTest,omitempty"`
	// example:
	//
	// 31***
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// INNER
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// example:
	//
	// 23****
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// example:
	//
	// 1
	UseDsql *int32 `json:"UseDsql,omitempty" xml:"UseDsql,omitempty"`
	UseSsl  *int32 `json:"UseSsl,omitempty" xml:"UseSsl,omitempty"`
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
