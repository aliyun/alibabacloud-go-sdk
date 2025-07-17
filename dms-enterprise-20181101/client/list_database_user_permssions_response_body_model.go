// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatabaseUserPermssionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListDatabaseUserPermssionsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDatabaseUserPermssionsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListDatabaseUserPermssionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDatabaseUserPermssionsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListDatabaseUserPermssionsResponseBody
	GetTotalCount() *int64
	SetUserPermissions(v *ListDatabaseUserPermssionsResponseBodyUserPermissions) *ListDatabaseUserPermssionsResponseBody
	GetUserPermissions() *ListDatabaseUserPermssionsResponseBodyUserPermissions
}

type ListDatabaseUserPermssionsResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 68075D06-7406-4887-83A7-F558A4D28C13
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The details of user permissions.
	UserPermissions *ListDatabaseUserPermssionsResponseBodyUserPermissions `json:"UserPermissions,omitempty" xml:"UserPermissions,omitempty" type:"Struct"`
}

func (s ListDatabaseUserPermssionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDatabaseUserPermssionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatabaseUserPermssionsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDatabaseUserPermssionsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDatabaseUserPermssionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDatabaseUserPermssionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDatabaseUserPermssionsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDatabaseUserPermssionsResponseBody) GetUserPermissions() *ListDatabaseUserPermssionsResponseBodyUserPermissions {
	return s.UserPermissions
}

func (s *ListDatabaseUserPermssionsResponseBody) SetErrorCode(v string) *ListDatabaseUserPermssionsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBody) SetErrorMessage(v string) *ListDatabaseUserPermssionsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBody) SetRequestId(v string) *ListDatabaseUserPermssionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBody) SetSuccess(v bool) *ListDatabaseUserPermssionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBody) SetTotalCount(v int64) *ListDatabaseUserPermssionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBody) SetUserPermissions(v *ListDatabaseUserPermssionsResponseBodyUserPermissions) *ListDatabaseUserPermssionsResponseBody {
	s.UserPermissions = v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDatabaseUserPermssionsResponseBodyUserPermissions struct {
	UserPermission []*ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission `json:"UserPermission,omitempty" xml:"UserPermission,omitempty" type:"Repeated"`
}

func (s ListDatabaseUserPermssionsResponseBodyUserPermissions) String() string {
	return dara.Prettify(s)
}

func (s ListDatabaseUserPermssionsResponseBodyUserPermissions) GoString() string {
	return s.String()
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissions) GetUserPermission() []*ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission {
	return s.UserPermission
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissions) SetUserPermission(v []*ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) *ListDatabaseUserPermssionsResponseBodyUserPermissions {
	s.UserPermission = v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissions) Validate() error {
	return dara.Validate(s)
}

type ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission struct {
	// The alias of the database instance.
	//
	// example:
	//
	// instance_alias
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The name of a column.
	//
	// example:
	//
	// column_name
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The ID of the database.
	//
	// example:
	//
	// 1234
	DbId *string `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The type of the database engine.
	//
	// example:
	//
	// mysql
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The type of resources on which the user has permissions.
	//
	// example:
	//
	// DATABASE
	DsType *string `json:"DsType,omitempty" xml:"DsType,omitempty"`
	// The type of the environment to which the database belongs.
	//
	// example:
	//
	// product
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// 1443
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the database is a logical database.
	//
	// example:
	//
	// false
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// The details of user permissions.
	PermDetails *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetails `json:"PermDetails,omitempty" xml:"PermDetails,omitempty" type:"Struct"`
	// The name of the database.
	//
	// example:
	//
	// xxx@xxxx:3306
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name that is used to search for the database.
	//
	// example:
	//
	// xxx
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
	// The ID of the table.
	//
	// example:
	//
	// 42345
	TableId *string `json:"TableId,omitempty" xml:"TableId,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// table_name
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// 14324
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The nickname of the user.
	//
	// example:
	//
	// user_nick_name
	UserNickName *string `json:"UserNickName,omitempty" xml:"UserNickName,omitempty"`
}

func (s ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) String() string {
	return dara.Prettify(s)
}

func (s ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) GoString() string {
	return s.String()
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) GetAlias() *string {
	return s.Alias
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) GetColumnName() *string {
	return s.ColumnName
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) GetDbId() *string {
	return s.DbId
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) GetDbType() *string {
	return s.DbType
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) GetDsType() *string {
	return s.DsType
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) GetEnvType() *string {
	return s.EnvType
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) GetLogic() *bool {
	return s.Logic
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) GetPermDetails() *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetails {
	return s.PermDetails
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) GetSchemaName() *string {
	return s.SchemaName
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) GetSearchName() *string {
	return s.SearchName
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) GetTableId() *string {
	return s.TableId
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) GetTableName() *string {
	return s.TableName
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) GetUserId() *string {
	return s.UserId
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) GetUserNickName() *string {
	return s.UserNickName
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) SetAlias(v string) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission {
	s.Alias = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) SetColumnName(v string) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission {
	s.ColumnName = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) SetDbId(v string) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission {
	s.DbId = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) SetDbType(v string) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission {
	s.DbType = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) SetDsType(v string) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission {
	s.DsType = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) SetEnvType(v string) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission {
	s.EnvType = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) SetInstanceId(v string) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission {
	s.InstanceId = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) SetLogic(v bool) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission {
	s.Logic = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) SetPermDetails(v *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetails) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission {
	s.PermDetails = v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) SetSchemaName(v string) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission {
	s.SchemaName = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) SetSearchName(v string) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission {
	s.SearchName = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) SetTableId(v string) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission {
	s.TableId = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) SetTableName(v string) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission {
	s.TableName = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) SetUserId(v string) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission {
	s.UserId = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) SetUserNickName(v string) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission {
	s.UserNickName = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) Validate() error {
	return dara.Validate(s)
}

type ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetails struct {
	PermDetail []*ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail `json:"PermDetail,omitempty" xml:"PermDetail,omitempty" type:"Repeated"`
}

func (s ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetails) String() string {
	return dara.Prettify(s)
}

func (s ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetails) GoString() string {
	return s.String()
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetails) GetPermDetail() []*ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail {
	return s.PermDetail
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetails) SetPermDetail(v []*ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetails {
	s.PermDetail = v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetails) Validate() error {
	return dara.Validate(s)
}

type ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail struct {
	// The time when the permission was created.
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
	// The extra information. This parameter is reserved.
	//
	// example:
	//
	// DEFAULT
	ExtraData *string `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
	// The description of the entity that authorizes the permission.
	//
	// example:
	//
	// xxx grant
	OriginFrom *string `json:"OriginFrom,omitempty" xml:"OriginFrom,omitempty"`
	// The type of the permission. Valid values:
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
	// 13434
	UserAccessId *string `json:"UserAccessId,omitempty" xml:"UserAccessId,omitempty"`
}

func (s ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) String() string {
	return dara.Prettify(s)
}

func (s ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) GoString() string {
	return s.String()
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) GetCreateDate() *string {
	return s.CreateDate
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) GetExpireDate() *string {
	return s.ExpireDate
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) GetExtraData() *string {
	return s.ExtraData
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) GetOriginFrom() *string {
	return s.OriginFrom
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) GetPermType() *string {
	return s.PermType
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) GetUserAccessId() *string {
	return s.UserAccessId
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) SetCreateDate(v string) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail {
	s.CreateDate = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) SetExpireDate(v string) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail {
	s.ExpireDate = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) SetExtraData(v string) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail {
	s.ExtraData = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) SetOriginFrom(v string) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail {
	s.OriginFrom = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) SetPermType(v string) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail {
	s.PermType = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) SetUserAccessId(v string) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail {
	s.UserAccessId = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) Validate() error {
	return dara.Validate(s)
}
