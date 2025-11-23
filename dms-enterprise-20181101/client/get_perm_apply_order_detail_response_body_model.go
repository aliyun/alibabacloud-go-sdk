// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPermApplyOrderDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetPermApplyOrderDetailResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetPermApplyOrderDetailResponseBody
	GetErrorMessage() *string
	SetPermApplyOrderDetail(v *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetail) *GetPermApplyOrderDetailResponseBody
	GetPermApplyOrderDetail() *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetail
	SetRequestId(v string) *GetPermApplyOrderDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPermApplyOrderDetailResponseBody
	GetSuccess() *bool
}

type GetPermApplyOrderDetailResponseBody struct {
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
	// The details of the permission application ticket.
	PermApplyOrderDetail *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetail `json:"PermApplyOrderDetail,omitempty" xml:"PermApplyOrderDetail,omitempty" type:"Struct"`
	// The request ID. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// A4C4499E-5AC2-4318-9FCF-03E426781A04
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPermApplyOrderDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPermApplyOrderDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetPermApplyOrderDetailResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetPermApplyOrderDetailResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetPermApplyOrderDetailResponseBody) GetPermApplyOrderDetail() *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetail {
	return s.PermApplyOrderDetail
}

func (s *GetPermApplyOrderDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPermApplyOrderDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPermApplyOrderDetailResponseBody) SetErrorCode(v string) *GetPermApplyOrderDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBody) SetErrorMessage(v string) *GetPermApplyOrderDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBody) SetPermApplyOrderDetail(v *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetail) *GetPermApplyOrderDetailResponseBody {
	s.PermApplyOrderDetail = v
	return s
}

func (s *GetPermApplyOrderDetailResponseBody) SetRequestId(v string) *GetPermApplyOrderDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBody) SetSuccess(v bool) *GetPermApplyOrderDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBody) Validate() error {
	if s.PermApplyOrderDetail != nil {
		if err := s.PermApplyOrderDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPermApplyOrderDetailResponseBodyPermApplyOrderDetail struct {
	// The type of objects on which you apply for permissions. Valid values:
	//
	// 	- **DB**: database
	//
	// 	- **TAB**: table
	//
	// 	- **COL**: column
	//
	// 	- **INSTANT**: instance
	//
	// example:
	//
	// DB
	ApplyType *string `json:"ApplyType,omitempty" xml:"ApplyType,omitempty"`
	// The type of the permissions that you apply for. Valid values:
	//
	// 	- **1**: the permissions to query information.
	//
	// 	- **2**: the permissions to export information.
	//
	// 	- **3**: the permissions to query and export information.
	//
	// 	- **4**: the permissions to modify information.
	//
	// 	- **5**: the permissions to query and modify information.
	//
	// 	- **6**: the permissions to export and modify information.
	//
	// 	- **7**: the permissions to query, export, and modify information.
	//
	// 	- **8**: the permissions to log on to the database.
	//
	// example:
	//
	// 7
	PermType *int64 `json:"PermType,omitempty" xml:"PermType,omitempty"`
	// The list of resources.
	Resources []*GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// The validity duration of the permissions. Unit: seconds.
	//
	// example:
	//
	// 3600
	Seconds *int64 `json:"Seconds,omitempty" xml:"Seconds,omitempty"`
}

func (s GetPermApplyOrderDetailResponseBodyPermApplyOrderDetail) String() string {
	return dara.Prettify(s)
}

func (s GetPermApplyOrderDetailResponseBodyPermApplyOrderDetail) GoString() string {
	return s.String()
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetail) GetApplyType() *string {
	return s.ApplyType
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetail) GetPermType() *int64 {
	return s.PermType
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetail) GetResources() []*GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResources {
	return s.Resources
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetail) GetSeconds() *int64 {
	return s.Seconds
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetail) SetApplyType(v string) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetail {
	s.ApplyType = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetail) SetPermType(v int64) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetail {
	s.PermType = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetail) SetResources(v []*GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResources) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetail {
	s.Resources = v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetail) SetSeconds(v int64) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetail {
	s.Seconds = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetail) Validate() error {
	if s.Resources != nil {
		for _, item := range s.Resources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResources struct {
	// The information about the column.
	ColumnInfo *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesColumnInfo `json:"ColumnInfo,omitempty" xml:"ColumnInfo,omitempty" type:"Struct"`
	// The information about the database.
	DatabaseInfo *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo `json:"DatabaseInfo,omitempty" xml:"DatabaseInfo,omitempty" type:"Struct"`
	// The information about the instance.
	InstanceInfo *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo `json:"InstanceInfo,omitempty" xml:"InstanceInfo,omitempty" type:"Struct"`
	RowInfo      *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesRowInfo      `json:"RowInfo,omitempty" xml:"RowInfo,omitempty" type:"Struct"`
	RowValueInfo *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesRowValueInfo `json:"RowValueInfo,omitempty" xml:"RowValueInfo,omitempty" type:"Struct"`
	// The information about the table.
	TableInfo *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesTableInfo `json:"TableInfo,omitempty" xml:"TableInfo,omitempty" type:"Struct"`
}

func (s GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResources) String() string {
	return dara.Prettify(s)
}

func (s GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResources) GoString() string {
	return s.String()
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResources) GetColumnInfo() *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesColumnInfo {
	return s.ColumnInfo
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResources) GetDatabaseInfo() *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo {
	return s.DatabaseInfo
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResources) GetInstanceInfo() *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo {
	return s.InstanceInfo
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResources) GetRowInfo() *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesRowInfo {
	return s.RowInfo
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResources) GetRowValueInfo() *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesRowValueInfo {
	return s.RowValueInfo
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResources) GetTableInfo() *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesTableInfo {
	return s.TableInfo
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResources) SetColumnInfo(v *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesColumnInfo) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResources {
	s.ColumnInfo = v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResources) SetDatabaseInfo(v *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResources {
	s.DatabaseInfo = v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResources) SetInstanceInfo(v *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResources {
	s.InstanceInfo = v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResources) SetRowInfo(v *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesRowInfo) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResources {
	s.RowInfo = v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResources) SetRowValueInfo(v *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesRowValueInfo) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResources {
	s.RowValueInfo = v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResources) SetTableInfo(v *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesTableInfo) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResources {
	s.TableInfo = v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResources) Validate() error {
	if s.ColumnInfo != nil {
		if err := s.ColumnInfo.Validate(); err != nil {
			return err
		}
	}
	if s.DatabaseInfo != nil {
		if err := s.DatabaseInfo.Validate(); err != nil {
			return err
		}
	}
	if s.InstanceInfo != nil {
		if err := s.InstanceInfo.Validate(); err != nil {
			return err
		}
	}
	if s.RowInfo != nil {
		if err := s.RowInfo.Validate(); err != nil {
			return err
		}
	}
	if s.RowValueInfo != nil {
		if err := s.RowValueInfo.Validate(); err != nil {
			return err
		}
	}
	if s.TableInfo != nil {
		if err := s.TableInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesColumnInfo struct {
	// The name of the column.
	//
	// example:
	//
	// test_col
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// test_tb
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesColumnInfo) String() string {
	return dara.Prettify(s)
}

func (s GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesColumnInfo) GoString() string {
	return s.String()
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesColumnInfo) GetColumnName() *string {
	return s.ColumnName
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesColumnInfo) GetTableName() *string {
	return s.TableName
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesColumnInfo) SetColumnName(v string) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesColumnInfo {
	s.ColumnName = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesColumnInfo) SetTableName(v string) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesColumnInfo {
	s.TableName = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesColumnInfo) Validate() error {
	return dara.Validate(s)
}

type GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo struct {
	// The database ID.
	//
	// example:
	//
	// 12345
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The type of the database engine.
	//
	// example:
	//
	// MySQL
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The type of the environment to which the instance belongs. For more information, see [Change the environment type of an instance](https://help.aliyun.com/document_detail/163309.html).
	//
	// example:
	//
	// product
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// Indicates whether the database is a logical database. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// The IDs of the owners of the database.
	OwnerIds []*int64 `json:"OwnerIds,omitempty" xml:"OwnerIds,omitempty" type:"Repeated"`
	// The nicknames of the owners of the database.
	OwnerNickNames []*string `json:"OwnerNickNames,omitempty" xml:"OwnerNickNames,omitempty" type:"Repeated"`
	// The name that is used to search for the database.
	//
	// example:
	//
	// test@xxxx:3306【test】
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
}

func (s GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo) String() string {
	return dara.Prettify(s)
}

func (s GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo) GoString() string {
	return s.String()
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo) GetDbId() *int64 {
	return s.DbId
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo) GetDbType() *string {
	return s.DbType
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo) GetEnvType() *string {
	return s.EnvType
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo) GetLogic() *bool {
	return s.Logic
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo) GetOwnerIds() []*int64 {
	return s.OwnerIds
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo) GetOwnerNickNames() []*string {
	return s.OwnerNickNames
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo) GetSearchName() *string {
	return s.SearchName
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo) SetDbId(v int64) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo {
	s.DbId = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo) SetDbType(v string) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo {
	s.DbType = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo) SetEnvType(v string) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo {
	s.EnvType = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo) SetLogic(v bool) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo {
	s.Logic = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo) SetOwnerIds(v []*int64) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo {
	s.OwnerIds = v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo) SetOwnerNickNames(v []*string) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo {
	s.OwnerNickNames = v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo) SetSearchName(v string) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo {
	s.SearchName = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo) Validate() error {
	return dara.Validate(s)
}

type GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo struct {
	// The type of the database engine.
	//
	// example:
	//
	// MySQL
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The ID of the database administrator (DBA) of the instance.
	//
	// example:
	//
	// 12345
	DbaId *int64 `json:"DbaId,omitempty" xml:"DbaId,omitempty"`
	// The nickname of the DBA of the instance.
	//
	// example:
	//
	// test_dba
	DbaNickName *string `json:"DbaNickName,omitempty" xml:"DbaNickName,omitempty"`
	// The type of the environment to which the instance belongs. For more information, see [Change the environment type of an instance](https://help.aliyun.com/document_detail/163309.html).
	//
	// example:
	//
	// product
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The endpoint of the instance.
	//
	// example:
	//
	// xxxx
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// 12345
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IDs of the owners of the instance.
	OwnerIds []*int64 `json:"OwnerIds,omitempty" xml:"OwnerIds,omitempty" type:"Repeated"`
	// The nicknames of the owners of the instance.
	OwnerNickName []*string `json:"OwnerNickName,omitempty" xml:"OwnerNickName,omitempty" type:"Repeated"`
	// The port that is used to connect to the instance.
	//
	// example:
	//
	// 3306
	Port *int64 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The name that is used to search for the instance.
	//
	// example:
	//
	// xxxx:3306
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
}

func (s GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo) String() string {
	return dara.Prettify(s)
}

func (s GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo) GoString() string {
	return s.String()
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo) GetDbType() *string {
	return s.DbType
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo) GetDbaId() *int64 {
	return s.DbaId
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo) GetDbaNickName() *string {
	return s.DbaNickName
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo) GetEnvType() *string {
	return s.EnvType
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo) GetHost() *string {
	return s.Host
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo) GetOwnerIds() []*int64 {
	return s.OwnerIds
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo) GetOwnerNickName() []*string {
	return s.OwnerNickName
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo) GetPort() *int64 {
	return s.Port
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo) GetSearchName() *string {
	return s.SearchName
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo) SetDbType(v string) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo {
	s.DbType = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo) SetDbaId(v int64) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo {
	s.DbaId = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo) SetDbaNickName(v string) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo {
	s.DbaNickName = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo) SetEnvType(v string) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo {
	s.EnvType = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo) SetHost(v string) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo {
	s.Host = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo) SetInstanceId(v string) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo {
	s.InstanceId = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo) SetOwnerIds(v []*int64) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo {
	s.OwnerIds = v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo) SetOwnerNickName(v []*string) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo {
	s.OwnerNickName = v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo) SetPort(v int64) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo {
	s.Port = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo) SetSearchName(v string) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo {
	s.SearchName = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo) Validate() error {
	return dara.Validate(s)
}

type GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesRowInfo struct {
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	DbId       *int64  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	Logic      *bool   `json:"Logic,omitempty" xml:"Logic,omitempty"`
	MatchMode  *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	RowGroupId *int64  `json:"RowGroupId,omitempty" xml:"RowGroupId,omitempty"`
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName  *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesRowInfo) String() string {
	return dara.Prettify(s)
}

func (s GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesRowInfo) GoString() string {
	return s.String()
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesRowInfo) GetColumnName() *string {
	return s.ColumnName
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesRowInfo) GetDbId() *int64 {
	return s.DbId
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesRowInfo) GetLogic() *bool {
	return s.Logic
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesRowInfo) GetMatchMode() *string {
	return s.MatchMode
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesRowInfo) GetRowGroupId() *int64 {
	return s.RowGroupId
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesRowInfo) GetSchemaName() *string {
	return s.SchemaName
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesRowInfo) GetTableName() *string {
	return s.TableName
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesRowInfo) SetColumnName(v string) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesRowInfo {
	s.ColumnName = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesRowInfo) SetDbId(v int64) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesRowInfo {
	s.DbId = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesRowInfo) SetLogic(v bool) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesRowInfo {
	s.Logic = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesRowInfo) SetMatchMode(v string) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesRowInfo {
	s.MatchMode = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesRowInfo) SetRowGroupId(v int64) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesRowInfo {
	s.RowGroupId = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesRowInfo) SetSchemaName(v string) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesRowInfo {
	s.SchemaName = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesRowInfo) SetTableName(v string) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesRowInfo {
	s.TableName = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesRowInfo) Validate() error {
	return dara.Validate(s)
}

type GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesRowValueInfo struct {
	RowValue *string `json:"RowValue,omitempty" xml:"RowValue,omitempty"`
}

func (s GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesRowValueInfo) String() string {
	return dara.Prettify(s)
}

func (s GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesRowValueInfo) GoString() string {
	return s.String()
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesRowValueInfo) GetRowValue() *string {
	return s.RowValue
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesRowValueInfo) SetRowValue(v string) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesRowValueInfo {
	s.RowValue = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesRowValueInfo) Validate() error {
	return dara.Validate(s)
}

type GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesTableInfo struct {
	// The name of the table.
	//
	// example:
	//
	// test_tb
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesTableInfo) String() string {
	return dara.Prettify(s)
}

func (s GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesTableInfo) GoString() string {
	return s.String()
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesTableInfo) GetTableName() *string {
	return s.TableName
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesTableInfo) SetTableName(v string) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesTableInfo {
	s.TableName = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesTableInfo) Validate() error {
	return dara.Validate(s)
}
