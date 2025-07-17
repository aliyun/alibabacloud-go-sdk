// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserPermissionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListUserPermissionsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListUserPermissionsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListUserPermissionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListUserPermissionsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListUserPermissionsResponseBody
	GetTotalCount() *int64
	SetUserPermissions(v *ListUserPermissionsResponseBodyUserPermissions) *ListUserPermissionsResponseBody
	GetUserPermissions() *ListUserPermissionsResponseBodyUserPermissions
}

type ListUserPermissionsResponseBody struct {
	// The error code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C51420E3-144A-4A94-B473-8662FCF4AD10
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - true: The request is successful.
	//
	// - false: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries that meet the query conditions.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The details of the permissions that the user has.
	UserPermissions *ListUserPermissionsResponseBodyUserPermissions `json:"UserPermissions,omitempty" xml:"UserPermissions,omitempty" type:"Struct"`
}

func (s ListUserPermissionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserPermissionsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListUserPermissionsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListUserPermissionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserPermissionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListUserPermissionsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListUserPermissionsResponseBody) GetUserPermissions() *ListUserPermissionsResponseBodyUserPermissions {
	return s.UserPermissions
}

func (s *ListUserPermissionsResponseBody) SetErrorCode(v string) *ListUserPermissionsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListUserPermissionsResponseBody) SetErrorMessage(v string) *ListUserPermissionsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListUserPermissionsResponseBody) SetRequestId(v string) *ListUserPermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserPermissionsResponseBody) SetSuccess(v bool) *ListUserPermissionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListUserPermissionsResponseBody) SetTotalCount(v int64) *ListUserPermissionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUserPermissionsResponseBody) SetUserPermissions(v *ListUserPermissionsResponseBodyUserPermissions) *ListUserPermissionsResponseBody {
	s.UserPermissions = v
	return s
}

func (s *ListUserPermissionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListUserPermissionsResponseBodyUserPermissions struct {
	UserPermission []*ListUserPermissionsResponseBodyUserPermissionsUserPermission `json:"UserPermission,omitempty" xml:"UserPermission,omitempty" type:"Repeated"`
}

func (s ListUserPermissionsResponseBodyUserPermissions) String() string {
	return dara.Prettify(s)
}

func (s ListUserPermissionsResponseBodyUserPermissions) GoString() string {
	return s.String()
}

func (s *ListUserPermissionsResponseBodyUserPermissions) GetUserPermission() []*ListUserPermissionsResponseBodyUserPermissionsUserPermission {
	return s.UserPermission
}

func (s *ListUserPermissionsResponseBodyUserPermissions) SetUserPermission(v []*ListUserPermissionsResponseBodyUserPermissionsUserPermission) *ListUserPermissionsResponseBodyUserPermissions {
	s.UserPermission = v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissions) Validate() error {
	return dara.Validate(s)
}

type ListUserPermissionsResponseBodyUserPermissionsUserPermission struct {
	// The alias of the instance.
	//
	// example:
	//
	// instance_alias
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The name of the field.
	//
	// example:
	//
	// column_name
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The ID of the database.
	//
	// example:
	//
	// 1860****
	DbId *string `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The type of the database. For more information about the valid values of this parameter, see [DbType parameter](https://www.alibabacloud.com/help/en/data-management-service/latest/dbtype-parameter).
	//
	// example:
	//
	// polardb
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The permissions on a specific type of objects that are granted to the user. Valid values:
	//
	// - DATABASE: permissions on physical databases
	//
	// - LOGIC_DATABASE: permissions on logical databases
	//
	// - TABLE: permissions on physical tables
	//
	// - LOGIC_TABLE: permissions on logical tables
	//
	// example:
	//
	// DATABASE
	DsType *string `json:"DsType,omitempty" xml:"DsType,omitempty"`
	// The type of the environment to which the database belongs. Valid values:
	//
	// - product: production environment
	//
	// - dev: development environment
	//
	// - pre: staging environment
	//
	// - test: test environment
	//
	// - sit: SIT environment
	//
	// - uat: UAT environment
	//
	// - pet: stress testing environment
	//
	// - stag: STAG environment
	//
	// example:
	//
	// dev
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The endpoint that is used to connect the database.
	//
	// example:
	//
	// rm-bp144d5ky4l4r****
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// 174****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the database is a logical database. Valid values:
	//
	// 	- true: The database is a logical database.
	//
	// 	- false: The database is a physical database.
	//
	// example:
	//
	// false
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// The details of permissions.
	PermDetails *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetails `json:"PermDetails,omitempty" xml:"PermDetails,omitempty" type:"Struct"`
	// The port that is used to connect to the instance.
	//
	// example:
	//
	// 3306
	Port *int64 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// test_db
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name that is used to search for the database.
	//
	// example:
	//
	// test_db@xxx:3306
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
	// The ID of the table.
	//
	// example:
	//
	// 13434
	TableId *string `json:"TableId,omitempty" xml:"TableId,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// test_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// 51****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The nickname of the user.
	//
	// example:
	//
	// nick_name
	UserNickName *string `json:"UserNickName,omitempty" xml:"UserNickName,omitempty"`
}

func (s ListUserPermissionsResponseBodyUserPermissionsUserPermission) String() string {
	return dara.Prettify(s)
}

func (s ListUserPermissionsResponseBodyUserPermissionsUserPermission) GoString() string {
	return s.String()
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) GetAlias() *string {
	return s.Alias
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) GetColumnName() *string {
	return s.ColumnName
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) GetDbId() *string {
	return s.DbId
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) GetDbType() *string {
	return s.DbType
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) GetDsType() *string {
	return s.DsType
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) GetEnvType() *string {
	return s.EnvType
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) GetHost() *string {
	return s.Host
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) GetLogic() *bool {
	return s.Logic
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) GetPermDetails() *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetails {
	return s.PermDetails
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) GetPort() *int64 {
	return s.Port
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) GetSchemaName() *string {
	return s.SchemaName
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) GetSearchName() *string {
	return s.SearchName
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) GetTableId() *string {
	return s.TableId
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) GetTableName() *string {
	return s.TableName
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) GetUserId() *string {
	return s.UserId
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) GetUserNickName() *string {
	return s.UserNickName
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) SetAlias(v string) *ListUserPermissionsResponseBodyUserPermissionsUserPermission {
	s.Alias = &v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) SetColumnName(v string) *ListUserPermissionsResponseBodyUserPermissionsUserPermission {
	s.ColumnName = &v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) SetDbId(v string) *ListUserPermissionsResponseBodyUserPermissionsUserPermission {
	s.DbId = &v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) SetDbType(v string) *ListUserPermissionsResponseBodyUserPermissionsUserPermission {
	s.DbType = &v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) SetDsType(v string) *ListUserPermissionsResponseBodyUserPermissionsUserPermission {
	s.DsType = &v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) SetEnvType(v string) *ListUserPermissionsResponseBodyUserPermissionsUserPermission {
	s.EnvType = &v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) SetHost(v string) *ListUserPermissionsResponseBodyUserPermissionsUserPermission {
	s.Host = &v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) SetInstanceId(v string) *ListUserPermissionsResponseBodyUserPermissionsUserPermission {
	s.InstanceId = &v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) SetLogic(v bool) *ListUserPermissionsResponseBodyUserPermissionsUserPermission {
	s.Logic = &v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) SetPermDetails(v *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetails) *ListUserPermissionsResponseBodyUserPermissionsUserPermission {
	s.PermDetails = v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) SetPort(v int64) *ListUserPermissionsResponseBodyUserPermissionsUserPermission {
	s.Port = &v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) SetSchemaName(v string) *ListUserPermissionsResponseBodyUserPermissionsUserPermission {
	s.SchemaName = &v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) SetSearchName(v string) *ListUserPermissionsResponseBodyUserPermissionsUserPermission {
	s.SearchName = &v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) SetTableId(v string) *ListUserPermissionsResponseBodyUserPermissionsUserPermission {
	s.TableId = &v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) SetTableName(v string) *ListUserPermissionsResponseBodyUserPermissionsUserPermission {
	s.TableName = &v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) SetUserId(v string) *ListUserPermissionsResponseBodyUserPermissionsUserPermission {
	s.UserId = &v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) SetUserNickName(v string) *ListUserPermissionsResponseBodyUserPermissionsUserPermission {
	s.UserNickName = &v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) Validate() error {
	return dara.Validate(s)
}

type ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetails struct {
	PermDetail []*ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail `json:"PermDetail,omitempty" xml:"PermDetail,omitempty" type:"Repeated"`
}

func (s ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetails) String() string {
	return dara.Prettify(s)
}

func (s ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetails) GoString() string {
	return s.String()
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetails) GetPermDetail() []*ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail {
	return s.PermDetail
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetails) SetPermDetail(v []*ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetails {
	s.PermDetail = v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetails) Validate() error {
	return dara.Validate(s)
}

type ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail struct {
	// The time when the permissions were granted.
	//
	// example:
	//
	// 2019-12-12 00:00:00
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The time when the permissions expire.
	//
	// example:
	//
	// 2020-12-12 00:00:00
	ExpireDate *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// This parameter is reserved.
	//
	// example:
	//
	// xxx
	ExtraData *string `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
	// The user who grants the permissions.
	//
	// example:
	//
	// xxx authorization
	OriginFrom *string `json:"OriginFrom,omitempty" xml:"OriginFrom,omitempty"`
	// The type of the permissions. Valid values:
	//
	// 	- QUERY: the query permissions
	//
	// 	- EXPORT: the export permissions
	//
	// 	- CORRECT: the change permissions
	//
	// example:
	//
	// QUERY
	PermType *string `json:"PermType,omitempty" xml:"PermType,omitempty"`
	// The ID of the authorization record.
	//
	// example:
	//
	// 758****
	UserAccessId *string `json:"UserAccessId,omitempty" xml:"UserAccessId,omitempty"`
}

func (s ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) String() string {
	return dara.Prettify(s)
}

func (s ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) GoString() string {
	return s.String()
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) GetCreateDate() *string {
	return s.CreateDate
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) GetExpireDate() *string {
	return s.ExpireDate
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) GetExtraData() *string {
	return s.ExtraData
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) GetOriginFrom() *string {
	return s.OriginFrom
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) GetPermType() *string {
	return s.PermType
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) GetUserAccessId() *string {
	return s.UserAccessId
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) SetCreateDate(v string) *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail {
	s.CreateDate = &v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) SetExpireDate(v string) *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail {
	s.ExpireDate = &v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) SetExtraData(v string) *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail {
	s.ExtraData = &v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) SetOriginFrom(v string) *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail {
	s.OriginFrom = &v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) SetPermType(v string) *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail {
	s.PermType = &v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) SetUserAccessId(v string) *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail {
	s.UserAccessId = &v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) Validate() error {
	return dara.Validate(s)
}
