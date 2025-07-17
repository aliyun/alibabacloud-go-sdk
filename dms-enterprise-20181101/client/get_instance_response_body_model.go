// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetInstanceResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetInstanceResponseBody
	GetErrorMessage() *string
	SetInstance(v *GetInstanceResponseBodyInstance) *GetInstanceResponseBody
	GetInstance() *GetInstanceResponseBodyInstance
	SetRequestId(v string) *GetInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetInstanceResponseBody
	GetSuccess() *bool
}

type GetInstanceResponseBody struct {
	// The error code that is returned.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message that is returned.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The details of the database instance.
	Instance *GetInstanceResponseBodyInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// F6C80B69-3203-56AC-8021-18BA72A6F4E1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetInstanceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetInstanceResponseBody) GetInstance() *GetInstanceResponseBodyInstance {
	return s.Instance
}

func (s *GetInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetInstanceResponseBody) SetErrorCode(v string) *GetInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetInstanceResponseBody) SetErrorMessage(v string) *GetInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetInstanceResponseBody) SetInstance(v *GetInstanceResponseBodyInstance) *GetInstanceResponseBody {
	s.Instance = v
	return s
}

func (s *GetInstanceResponseBody) SetRequestId(v string) *GetInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceResponseBody) SetSuccess(v bool) *GetInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *GetInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyInstance struct {
	// The name of the database link for the database instance.
	//
	// example:
	//
	// test
	DataLinkName *string `json:"DataLinkName,omitempty" xml:"DataLinkName,omitempty"`
	// The password that is used to log on to the database.
	//
	// example:
	//
	// ******
	DatabasePassword *string `json:"DatabasePassword,omitempty" xml:"DatabasePassword,omitempty"`
	// The account that is used to log on to the database instance.
	//
	// example:
	//
	// dbuser
	DatabaseUser *string `json:"DatabaseUser,omitempty" xml:"DatabaseUser,omitempty"`
	// The ID of the database administrator (DBA) for the database instance.
	//
	// example:
	//
	// 29****
	DbaId *string `json:"DbaId,omitempty" xml:"DbaId,omitempty"`
	// The nickname of the DBA for the database instance.
	//
	// example:
	//
	// dbaname
	DbaNickName *string `json:"DbaNickName,omitempty" xml:"DbaNickName,omitempty"`
	// Indicates whether the lock-free schema change feature is enabled for the database instance.
	//
	// example:
	//
	// 0
	DdlOnline *int32 `json:"DdlOnline,omitempty" xml:"DdlOnline,omitempty"`
	// The ID of the Elastic Compute Service (ECS) instance on which the database instance is deployed.
	//
	// example:
	//
	// i-bp124ldpklqz59y3****
	EcsInstanceId *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	// The ID of the region in which the database instance resides.
	//
	// example:
	//
	// cn-beijing
	EcsRegion *string `json:"EcsRegion,omitempty" xml:"EcsRegion,omitempty"`
	// The type of the environment to which the database instance belongs. Valid values:
	//
	// 	- **product**: production environment
	//
	// 	- **dev**: development environment
	//
	// 	- **pre**: staging environment
	//
	// 	- **test**: test environment
	//
	// 	- **sit**: system integration testing (SIT) environment
	//
	// 	- **uat**: user acceptance testing (UAT) environment
	//
	// 	- **pet**: stress testing environment
	//
	// 	- **stag**: STAG environment
	//
	// example:
	//
	// test
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The timeout period for exporting data from the database instance.
	//
	// example:
	//
	// 86400
	ExportTimeout *int32 `json:"ExportTimeout,omitempty" xml:"ExportTimeout,omitempty"`
	// The host address that is used to connect to the database instance.
	//
	// example:
	//
	// 192.168.XXX.XXX
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The alias of the database instance.
	//
	// example:
	//
	// test
	InstanceAlias *string `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	// The ID of the database instance.
	//
	// example:
	//
	// 188****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The source of the database instance.
	//
	// example:
	//
	// ECS_OWN
	InstanceSource *string `json:"InstanceSource,omitempty" xml:"InstanceSource,omitempty"`
	// The type of the database instance.
	//
	// example:
	//
	// postgresql
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The IDs of the owners for the database instance.
	OwnerIdList *GetInstanceResponseBodyInstanceOwnerIdList `json:"OwnerIdList,omitempty" xml:"OwnerIdList,omitempty" type:"Struct"`
	// The nicknames of the owners for the database instance.
	OwnerNameList *GetInstanceResponseBodyInstanceOwnerNameList `json:"OwnerNameList,omitempty" xml:"OwnerNameList,omitempty" type:"Struct"`
	// The port number that is used to connect to the database instance.
	//
	// example:
	//
	// 5432
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The timeout period for querying data in the database instance.
	//
	// example:
	//
	// 7200
	QueryTimeout *int32 `json:"QueryTimeout,omitempty" xml:"QueryTimeout,omitempty"`
	// The ID of the security rule set for the database instance.
	//
	// example:
	//
	// 3****
	SafeRuleId *string `json:"SafeRuleId,omitempty" xml:"SafeRuleId,omitempty"`
	// Whether sensitive data protection is enabled.  Valid values:
	//
	// - **true**: Enable.
	//
	// - **false**: Close.
	//
	// example:
	//
	// false
	SellSitd  *string `json:"SellSitd,omitempty" xml:"SellSitd,omitempty"`
	SellTrust *string `json:"SellTrust,omitempty" xml:"SellTrust,omitempty"`
	// The SID of the database instance.
	//
	// example:
	//
	// test
	Sid *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	// The control mode of the database instance.
	StandardGroup *GetInstanceResponseBodyInstanceStandardGroup `json:"StandardGroup,omitempty" xml:"StandardGroup,omitempty" type:"Struct"`
	// The status of the database instance. Valid values:
	//
	// 	- **NORMAL**: normal
	//
	// 	- **DISABLE**: disabled
	//
	// example:
	//
	// NORMAL
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// Indicates whether the cross-database query feature is enabled for the database instance. Valid values:
	//
	// 	- **0**: disabled
	//
	// 	- **1**: enabled
	//
	// example:
	//
	// 0
	UseDsql *int32 `json:"UseDsql,omitempty" xml:"UseDsql,omitempty"`
	// The ID of the virtual private cloud (VPC) to which the database instance belongs.
	//
	// example:
	//
	// vpc-o6wrloqsdqc9io3mg****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetInstanceResponseBodyInstance) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyInstance) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyInstance) GetDataLinkName() *string {
	return s.DataLinkName
}

func (s *GetInstanceResponseBodyInstance) GetDatabasePassword() *string {
	return s.DatabasePassword
}

func (s *GetInstanceResponseBodyInstance) GetDatabaseUser() *string {
	return s.DatabaseUser
}

func (s *GetInstanceResponseBodyInstance) GetDbaId() *string {
	return s.DbaId
}

func (s *GetInstanceResponseBodyInstance) GetDbaNickName() *string {
	return s.DbaNickName
}

func (s *GetInstanceResponseBodyInstance) GetDdlOnline() *int32 {
	return s.DdlOnline
}

func (s *GetInstanceResponseBodyInstance) GetEcsInstanceId() *string {
	return s.EcsInstanceId
}

func (s *GetInstanceResponseBodyInstance) GetEcsRegion() *string {
	return s.EcsRegion
}

func (s *GetInstanceResponseBodyInstance) GetEnvType() *string {
	return s.EnvType
}

func (s *GetInstanceResponseBodyInstance) GetExportTimeout() *int32 {
	return s.ExportTimeout
}

func (s *GetInstanceResponseBodyInstance) GetHost() *string {
	return s.Host
}

func (s *GetInstanceResponseBodyInstance) GetInstanceAlias() *string {
	return s.InstanceAlias
}

func (s *GetInstanceResponseBodyInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceResponseBodyInstance) GetInstanceSource() *string {
	return s.InstanceSource
}

func (s *GetInstanceResponseBodyInstance) GetInstanceType() *string {
	return s.InstanceType
}

func (s *GetInstanceResponseBodyInstance) GetOwnerIdList() *GetInstanceResponseBodyInstanceOwnerIdList {
	return s.OwnerIdList
}

func (s *GetInstanceResponseBodyInstance) GetOwnerNameList() *GetInstanceResponseBodyInstanceOwnerNameList {
	return s.OwnerNameList
}

func (s *GetInstanceResponseBodyInstance) GetPort() *int32 {
	return s.Port
}

func (s *GetInstanceResponseBodyInstance) GetQueryTimeout() *int32 {
	return s.QueryTimeout
}

func (s *GetInstanceResponseBodyInstance) GetSafeRuleId() *string {
	return s.SafeRuleId
}

func (s *GetInstanceResponseBodyInstance) GetSellSitd() *string {
	return s.SellSitd
}

func (s *GetInstanceResponseBodyInstance) GetSellTrust() *string {
	return s.SellTrust
}

func (s *GetInstanceResponseBodyInstance) GetSid() *string {
	return s.Sid
}

func (s *GetInstanceResponseBodyInstance) GetStandardGroup() *GetInstanceResponseBodyInstanceStandardGroup {
	return s.StandardGroup
}

func (s *GetInstanceResponseBodyInstance) GetState() *string {
	return s.State
}

func (s *GetInstanceResponseBodyInstance) GetUseDsql() *int32 {
	return s.UseDsql
}

func (s *GetInstanceResponseBodyInstance) GetVpcId() *string {
	return s.VpcId
}

func (s *GetInstanceResponseBodyInstance) SetDataLinkName(v string) *GetInstanceResponseBodyInstance {
	s.DataLinkName = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetDatabasePassword(v string) *GetInstanceResponseBodyInstance {
	s.DatabasePassword = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetDatabaseUser(v string) *GetInstanceResponseBodyInstance {
	s.DatabaseUser = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetDbaId(v string) *GetInstanceResponseBodyInstance {
	s.DbaId = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetDbaNickName(v string) *GetInstanceResponseBodyInstance {
	s.DbaNickName = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetDdlOnline(v int32) *GetInstanceResponseBodyInstance {
	s.DdlOnline = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetEcsInstanceId(v string) *GetInstanceResponseBodyInstance {
	s.EcsInstanceId = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetEcsRegion(v string) *GetInstanceResponseBodyInstance {
	s.EcsRegion = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetEnvType(v string) *GetInstanceResponseBodyInstance {
	s.EnvType = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetExportTimeout(v int32) *GetInstanceResponseBodyInstance {
	s.ExportTimeout = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetHost(v string) *GetInstanceResponseBodyInstance {
	s.Host = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetInstanceAlias(v string) *GetInstanceResponseBodyInstance {
	s.InstanceAlias = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetInstanceId(v string) *GetInstanceResponseBodyInstance {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetInstanceSource(v string) *GetInstanceResponseBodyInstance {
	s.InstanceSource = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetInstanceType(v string) *GetInstanceResponseBodyInstance {
	s.InstanceType = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetOwnerIdList(v *GetInstanceResponseBodyInstanceOwnerIdList) *GetInstanceResponseBodyInstance {
	s.OwnerIdList = v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetOwnerNameList(v *GetInstanceResponseBodyInstanceOwnerNameList) *GetInstanceResponseBodyInstance {
	s.OwnerNameList = v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetPort(v int32) *GetInstanceResponseBodyInstance {
	s.Port = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetQueryTimeout(v int32) *GetInstanceResponseBodyInstance {
	s.QueryTimeout = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetSafeRuleId(v string) *GetInstanceResponseBodyInstance {
	s.SafeRuleId = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetSellSitd(v string) *GetInstanceResponseBodyInstance {
	s.SellSitd = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetSellTrust(v string) *GetInstanceResponseBodyInstance {
	s.SellTrust = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetSid(v string) *GetInstanceResponseBodyInstance {
	s.Sid = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetStandardGroup(v *GetInstanceResponseBodyInstanceStandardGroup) *GetInstanceResponseBodyInstance {
	s.StandardGroup = v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetState(v string) *GetInstanceResponseBodyInstance {
	s.State = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetUseDsql(v int32) *GetInstanceResponseBodyInstance {
	s.UseDsql = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetVpcId(v string) *GetInstanceResponseBodyInstance {
	s.VpcId = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyInstanceOwnerIdList struct {
	OwnerIds []*string `json:"OwnerIds,omitempty" xml:"OwnerIds,omitempty" type:"Repeated"`
}

func (s GetInstanceResponseBodyInstanceOwnerIdList) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyInstanceOwnerIdList) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyInstanceOwnerIdList) GetOwnerIds() []*string {
	return s.OwnerIds
}

func (s *GetInstanceResponseBodyInstanceOwnerIdList) SetOwnerIds(v []*string) *GetInstanceResponseBodyInstanceOwnerIdList {
	s.OwnerIds = v
	return s
}

func (s *GetInstanceResponseBodyInstanceOwnerIdList) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyInstanceOwnerNameList struct {
	OwnerNames []*string `json:"OwnerNames,omitempty" xml:"OwnerNames,omitempty" type:"Repeated"`
}

func (s GetInstanceResponseBodyInstanceOwnerNameList) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyInstanceOwnerNameList) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyInstanceOwnerNameList) GetOwnerNames() []*string {
	return s.OwnerNames
}

func (s *GetInstanceResponseBodyInstanceOwnerNameList) SetOwnerNames(v []*string) *GetInstanceResponseBodyInstanceOwnerNameList {
	s.OwnerNames = v
	return s
}

func (s *GetInstanceResponseBodyInstanceOwnerNameList) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyInstanceStandardGroup struct {
	// The type of the control mode. Valid values:
	//
	// 	- **COMMON**: Security Collaboration
	//
	// 	- **NONE_CONTROL**: Flexible Management
	//
	// 	- **STABLE**: Stable Change
	//
	// example:
	//
	// NONE_CONTROL
	GroupMode *string `json:"GroupMode,omitempty" xml:"GroupMode,omitempty"`
	// The name of the security rule set corresponding to the control mode.
	//
	// example:
	//
	// test group name
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s GetInstanceResponseBodyInstanceStandardGroup) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyInstanceStandardGroup) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyInstanceStandardGroup) GetGroupMode() *string {
	return s.GroupMode
}

func (s *GetInstanceResponseBodyInstanceStandardGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *GetInstanceResponseBodyInstanceStandardGroup) SetGroupMode(v string) *GetInstanceResponseBodyInstanceStandardGroup {
	s.GroupMode = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceStandardGroup) SetGroupName(v string) *GetInstanceResponseBodyInstanceStandardGroup {
	s.GroupName = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceStandardGroup) Validate() error {
	return dara.Validate(s)
}
