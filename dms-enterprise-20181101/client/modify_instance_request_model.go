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
	// example:
	//
	// dblink_test
	DataLinkName *string `json:"DataLinkName,omitempty" xml:"DataLinkName,omitempty"`
	// example:
	//
	// test***
	DatabasePassword *string `json:"DatabasePassword,omitempty" xml:"DatabasePassword,omitempty"`
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
	// example:
	//
	// Y
	EnableSellTrust *string `json:"EnableSellTrust,omitempty" xml:"EnableSellTrust,omitempty"`
	// example:
	//
	// dev
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// example:
	//
	// 86400
	ExportTimeout *int32 `json:"ExportTimeout,omitempty" xml:"ExportTimeout,omitempty"`
	// example:
	//
	// 192.XXX.0.56
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// instance_test
	InstanceAlias *string `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 183****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// RDS
	InstanceSource *string `json:"InstanceSource,omitempty" xml:"InstanceSource,omitempty"`
	// example:
	//
	// MySQL
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
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
