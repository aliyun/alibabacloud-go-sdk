// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListInstancesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListInstancesResponseBody
	GetErrorMessage() *string
	SetInstanceList(v *ListInstancesResponseBodyInstanceList) *ListInstancesResponseBody
	GetInstanceList() *ListInstancesResponseBodyInstanceList
	SetRequestId(v string) *ListInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListInstancesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListInstancesResponseBody
	GetTotalCount() *int64
}

type ListInstancesResponseBody struct {
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
	// The information about the database instances that are returned.
	InstanceList *ListInstancesResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// B4B07137-F6AE-4756-8474-7F92BB6C4E04
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of database instances that are returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListInstancesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListInstancesResponseBody) GetInstanceList() *ListInstancesResponseBodyInstanceList {
	return s.InstanceList
}

func (s *ListInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListInstancesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListInstancesResponseBody) SetErrorCode(v string) *ListInstancesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListInstancesResponseBody) SetErrorMessage(v string) *ListInstancesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListInstancesResponseBody) SetInstanceList(v *ListInstancesResponseBodyInstanceList) *ListInstancesResponseBody {
	s.InstanceList = v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetSuccess(v bool) *ListInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *ListInstancesResponseBody) SetTotalCount(v int64) *ListInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListInstancesResponseBody) Validate() error {
	if s.InstanceList != nil {
		if err := s.InstanceList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInstancesResponseBodyInstanceList struct {
	Instance []*ListInstancesResponseBodyInstanceListInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s ListInstancesResponseBodyInstanceList) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstanceList) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstanceList) GetInstance() []*ListInstancesResponseBodyInstanceListInstance {
	return s.Instance
}

func (s *ListInstancesResponseBodyInstanceList) SetInstance(v []*ListInstancesResponseBodyInstanceListInstance) *ListInstancesResponseBodyInstanceList {
	s.Instance = v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) Validate() error {
	if s.Instance != nil {
		for _, item := range s.Instance {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstancesResponseBodyInstanceListInstance struct {
	// The name of the database link for the database instance.
	//
	// example:
	//
	// dblink_test
	DataLinkName *string `json:"DataLinkName,omitempty" xml:"DataLinkName,omitempty"`
	// The password that is used to log on to the database instance.
	//
	// example:
	//
	// ******
	DatabasePassword *string `json:"DatabasePassword,omitempty" xml:"DatabasePassword,omitempty"`
	// The account that is used to log on to the database.
	//
	// example:
	//
	// dbUser
	DatabaseUser *string `json:"DatabaseUser,omitempty" xml:"DatabaseUser,omitempty"`
	// The ID of the database administrator (DBA) of the database instance.
	//
	// example:
	//
	// 31****
	DbaId *string `json:"DbaId,omitempty" xml:"DbaId,omitempty"`
	// The nickname of the DBA of the instance.
	//
	// example:
	//
	// dbaName
	DbaNickName *string `json:"DbaNickName,omitempty" xml:"DbaNickName,omitempty"`
	// Indicates whether the lock-free schema change feature is enabled for the database instance.
	//
	// example:
	//
	// 1
	DdlOnline *int32 `json:"DdlOnline,omitempty" xml:"DdlOnline,omitempty"`
	// The ID of the ECS instance on which the database instance is deployed.
	//
	// example:
	//
	// 150****
	EcsInstanceId *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	// The ID of the region in which the database instance resides.
	//
	// example:
	//
	// cn-hangzhou
	EcsRegion *string `json:"EcsRegion,omitempty" xml:"EcsRegion,omitempty"`
	// The type of the environment to which the database instance belongs. Valid values:
	//
	// 	- **product:*	- production environment
	//
	// 	- **dev**: development environment
	//
	// 	- **pre**: pre-release environment
	//
	// 	- **test**: test environment
	//
	// 	- **sit**: SIT environment
	//
	// 	- **uat**: UAT environment
	//
	// 	- **pet**: stress testing environment
	//
	// 	- **stag:*	- staging environment
	//
	// example:
	//
	// product
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
	// ****.mysql.rds.aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The alias of the database instance.
	//
	// example:
	//
	// test
	InstanceAlias *string `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// 150***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The source of the database instance.
	//
	// example:
	//
	// RDS
	InstanceSource *string `json:"InstanceSource,omitempty" xml:"InstanceSource,omitempty"`
	// The type of the database instance.
	//
	// example:
	//
	// mysql
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The IDs of the owners of the database instance.
	OwnerIdList *ListInstancesResponseBodyInstanceListInstanceOwnerIdList `json:"OwnerIdList,omitempty" xml:"OwnerIdList,omitempty" type:"Struct"`
	// The nicknames of the owners of the database instance.
	OwnerNameList *ListInstancesResponseBodyInstanceListInstanceOwnerNameList `json:"OwnerNameList,omitempty" xml:"OwnerNameList,omitempty" type:"Struct"`
	// The port number that is used to connect to the database instance.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The timeout period for querying data in the database instance.
	//
	// example:
	//
	// 60
	QueryTimeout *int32 `json:"QueryTimeout,omitempty" xml:"QueryTimeout,omitempty"`
	// The ID of the security rule set of the database instance.
	//
	// example:
	//
	// 1
	SafeRuleId *string `json:"SafeRuleId,omitempty" xml:"SafeRuleId,omitempty"`
	// Indicates whether the sensitive data protection feature is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	SellSitd  *bool   `json:"SellSitd,omitempty" xml:"SellSitd,omitempty"`
	SellTrust *string `json:"SellTrust,omitempty" xml:"SellTrust,omitempty"`
	// The system ID (SID) of the database instance.
	//
	// example:
	//
	// test
	Sid *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	// The control mode of the database instance.
	StandardGroup *ListInstancesResponseBodyInstanceListInstanceStandardGroup `json:"StandardGroup,omitempty" xml:"StandardGroup,omitempty" type:"Struct"`
	// The status of the database instance.
	//
	// example:
	//
	// NORMAL
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// Indicates whether the cross-database query feature is enabled for the database instance. Valid values:
	//
	// 	- **0**: disabled
	//
	// 	- **1:**: enabled
	//
	// example:
	//
	// 1
	UseDsql *int32 `json:"UseDsql,omitempty" xml:"UseDsql,omitempty"`
	// The ID of the VPC to which the database instance belongs.
	//
	// example:
	//
	// vpc-o6wrloqsdqc9io3mg****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListInstancesResponseBodyInstanceListInstance) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstanceListInstance) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstanceListInstance) GetDataLinkName() *string {
	return s.DataLinkName
}

func (s *ListInstancesResponseBodyInstanceListInstance) GetDatabasePassword() *string {
	return s.DatabasePassword
}

func (s *ListInstancesResponseBodyInstanceListInstance) GetDatabaseUser() *string {
	return s.DatabaseUser
}

func (s *ListInstancesResponseBodyInstanceListInstance) GetDbaId() *string {
	return s.DbaId
}

func (s *ListInstancesResponseBodyInstanceListInstance) GetDbaNickName() *string {
	return s.DbaNickName
}

func (s *ListInstancesResponseBodyInstanceListInstance) GetDdlOnline() *int32 {
	return s.DdlOnline
}

func (s *ListInstancesResponseBodyInstanceListInstance) GetEcsInstanceId() *string {
	return s.EcsInstanceId
}

func (s *ListInstancesResponseBodyInstanceListInstance) GetEcsRegion() *string {
	return s.EcsRegion
}

func (s *ListInstancesResponseBodyInstanceListInstance) GetEnvType() *string {
	return s.EnvType
}

func (s *ListInstancesResponseBodyInstanceListInstance) GetExportTimeout() *int32 {
	return s.ExportTimeout
}

func (s *ListInstancesResponseBodyInstanceListInstance) GetHost() *string {
	return s.Host
}

func (s *ListInstancesResponseBodyInstanceListInstance) GetInstanceAlias() *string {
	return s.InstanceAlias
}

func (s *ListInstancesResponseBodyInstanceListInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancesResponseBodyInstanceListInstance) GetInstanceSource() *string {
	return s.InstanceSource
}

func (s *ListInstancesResponseBodyInstanceListInstance) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListInstancesResponseBodyInstanceListInstance) GetOwnerIdList() *ListInstancesResponseBodyInstanceListInstanceOwnerIdList {
	return s.OwnerIdList
}

func (s *ListInstancesResponseBodyInstanceListInstance) GetOwnerNameList() *ListInstancesResponseBodyInstanceListInstanceOwnerNameList {
	return s.OwnerNameList
}

func (s *ListInstancesResponseBodyInstanceListInstance) GetPort() *int32 {
	return s.Port
}

func (s *ListInstancesResponseBodyInstanceListInstance) GetQueryTimeout() *int32 {
	return s.QueryTimeout
}

func (s *ListInstancesResponseBodyInstanceListInstance) GetSafeRuleId() *string {
	return s.SafeRuleId
}

func (s *ListInstancesResponseBodyInstanceListInstance) GetSellSitd() *bool {
	return s.SellSitd
}

func (s *ListInstancesResponseBodyInstanceListInstance) GetSellTrust() *string {
	return s.SellTrust
}

func (s *ListInstancesResponseBodyInstanceListInstance) GetSid() *string {
	return s.Sid
}

func (s *ListInstancesResponseBodyInstanceListInstance) GetStandardGroup() *ListInstancesResponseBodyInstanceListInstanceStandardGroup {
	return s.StandardGroup
}

func (s *ListInstancesResponseBodyInstanceListInstance) GetState() *string {
	return s.State
}

func (s *ListInstancesResponseBodyInstanceListInstance) GetUseDsql() *int32 {
	return s.UseDsql
}

func (s *ListInstancesResponseBodyInstanceListInstance) GetVpcId() *string {
	return s.VpcId
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetDataLinkName(v string) *ListInstancesResponseBodyInstanceListInstance {
	s.DataLinkName = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetDatabasePassword(v string) *ListInstancesResponseBodyInstanceListInstance {
	s.DatabasePassword = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetDatabaseUser(v string) *ListInstancesResponseBodyInstanceListInstance {
	s.DatabaseUser = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetDbaId(v string) *ListInstancesResponseBodyInstanceListInstance {
	s.DbaId = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetDbaNickName(v string) *ListInstancesResponseBodyInstanceListInstance {
	s.DbaNickName = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetDdlOnline(v int32) *ListInstancesResponseBodyInstanceListInstance {
	s.DdlOnline = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetEcsInstanceId(v string) *ListInstancesResponseBodyInstanceListInstance {
	s.EcsInstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetEcsRegion(v string) *ListInstancesResponseBodyInstanceListInstance {
	s.EcsRegion = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetEnvType(v string) *ListInstancesResponseBodyInstanceListInstance {
	s.EnvType = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetExportTimeout(v int32) *ListInstancesResponseBodyInstanceListInstance {
	s.ExportTimeout = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetHost(v string) *ListInstancesResponseBodyInstanceListInstance {
	s.Host = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetInstanceAlias(v string) *ListInstancesResponseBodyInstanceListInstance {
	s.InstanceAlias = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetInstanceId(v string) *ListInstancesResponseBodyInstanceListInstance {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetInstanceSource(v string) *ListInstancesResponseBodyInstanceListInstance {
	s.InstanceSource = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetInstanceType(v string) *ListInstancesResponseBodyInstanceListInstance {
	s.InstanceType = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetOwnerIdList(v *ListInstancesResponseBodyInstanceListInstanceOwnerIdList) *ListInstancesResponseBodyInstanceListInstance {
	s.OwnerIdList = v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetOwnerNameList(v *ListInstancesResponseBodyInstanceListInstanceOwnerNameList) *ListInstancesResponseBodyInstanceListInstance {
	s.OwnerNameList = v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetPort(v int32) *ListInstancesResponseBodyInstanceListInstance {
	s.Port = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetQueryTimeout(v int32) *ListInstancesResponseBodyInstanceListInstance {
	s.QueryTimeout = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetSafeRuleId(v string) *ListInstancesResponseBodyInstanceListInstance {
	s.SafeRuleId = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetSellSitd(v bool) *ListInstancesResponseBodyInstanceListInstance {
	s.SellSitd = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetSellTrust(v string) *ListInstancesResponseBodyInstanceListInstance {
	s.SellTrust = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetSid(v string) *ListInstancesResponseBodyInstanceListInstance {
	s.Sid = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetStandardGroup(v *ListInstancesResponseBodyInstanceListInstanceStandardGroup) *ListInstancesResponseBodyInstanceListInstance {
	s.StandardGroup = v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetState(v string) *ListInstancesResponseBodyInstanceListInstance {
	s.State = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetUseDsql(v int32) *ListInstancesResponseBodyInstanceListInstance {
	s.UseDsql = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetVpcId(v string) *ListInstancesResponseBodyInstanceListInstance {
	s.VpcId = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) Validate() error {
	if s.OwnerIdList != nil {
		if err := s.OwnerIdList.Validate(); err != nil {
			return err
		}
	}
	if s.OwnerNameList != nil {
		if err := s.OwnerNameList.Validate(); err != nil {
			return err
		}
	}
	if s.StandardGroup != nil {
		if err := s.StandardGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInstancesResponseBodyInstanceListInstanceOwnerIdList struct {
	OwnerIds []*string `json:"OwnerIds,omitempty" xml:"OwnerIds,omitempty" type:"Repeated"`
}

func (s ListInstancesResponseBodyInstanceListInstanceOwnerIdList) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstanceListInstanceOwnerIdList) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstanceListInstanceOwnerIdList) GetOwnerIds() []*string {
	return s.OwnerIds
}

func (s *ListInstancesResponseBodyInstanceListInstanceOwnerIdList) SetOwnerIds(v []*string) *ListInstancesResponseBodyInstanceListInstanceOwnerIdList {
	s.OwnerIds = v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstanceOwnerIdList) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyInstanceListInstanceOwnerNameList struct {
	OwnerNames []*string `json:"OwnerNames,omitempty" xml:"OwnerNames,omitempty" type:"Repeated"`
}

func (s ListInstancesResponseBodyInstanceListInstanceOwnerNameList) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstanceListInstanceOwnerNameList) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstanceListInstanceOwnerNameList) GetOwnerNames() []*string {
	return s.OwnerNames
}

func (s *ListInstancesResponseBodyInstanceListInstanceOwnerNameList) SetOwnerNames(v []*string) *ListInstancesResponseBodyInstanceListInstanceOwnerNameList {
	s.OwnerNames = v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstanceOwnerNameList) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyInstanceListInstanceStandardGroup struct {
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
	// COMMON
	GroupMode *string `json:"GroupMode,omitempty" xml:"GroupMode,omitempty"`
	// The name of the security rule corresponding to the control mode.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s ListInstancesResponseBodyInstanceListInstanceStandardGroup) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstanceListInstanceStandardGroup) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstanceListInstanceStandardGroup) GetGroupMode() *string {
	return s.GroupMode
}

func (s *ListInstancesResponseBodyInstanceListInstanceStandardGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *ListInstancesResponseBodyInstanceListInstanceStandardGroup) SetGroupMode(v string) *ListInstancesResponseBodyInstanceListInstanceStandardGroup {
	s.GroupMode = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstanceStandardGroup) SetGroupName(v string) *ListInstancesResponseBodyInstanceListInstanceStandardGroup {
	s.GroupName = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstanceStandardGroup) Validate() error {
	return dara.Validate(s)
}
