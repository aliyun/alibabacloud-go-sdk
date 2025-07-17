// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStructSyncOrderDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetStructSyncOrderDetailResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetStructSyncOrderDetailResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetStructSyncOrderDetailResponseBody
	GetRequestId() *string
	SetStructSyncOrderDetail(v *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail) *GetStructSyncOrderDetailResponseBody
	GetStructSyncOrderDetail() *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail
	SetSuccess(v bool) *GetStructSyncOrderDetailResponseBody
	GetSuccess() *bool
}

type GetStructSyncOrderDetailResponseBody struct {
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
	// 4E1D2B4D-3E53-4ABC-999D-1D2520B3471A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the schema synchronization ticket.
	StructSyncOrderDetail *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail `json:"StructSyncOrderDetail,omitempty" xml:"StructSyncOrderDetail,omitempty" type:"Struct"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetStructSyncOrderDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStructSyncOrderDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetStructSyncOrderDetailResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetStructSyncOrderDetailResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetStructSyncOrderDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStructSyncOrderDetailResponseBody) GetStructSyncOrderDetail() *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail {
	return s.StructSyncOrderDetail
}

func (s *GetStructSyncOrderDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetStructSyncOrderDetailResponseBody) SetErrorCode(v string) *GetStructSyncOrderDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetStructSyncOrderDetailResponseBody) SetErrorMessage(v string) *GetStructSyncOrderDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetStructSyncOrderDetailResponseBody) SetRequestId(v string) *GetStructSyncOrderDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStructSyncOrderDetailResponseBody) SetStructSyncOrderDetail(v *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail) *GetStructSyncOrderDetailResponseBody {
	s.StructSyncOrderDetail = v
	return s
}

func (s *GetStructSyncOrderDetailResponseBody) SetSuccess(v bool) *GetStructSyncOrderDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetStructSyncOrderDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail struct {
	// Indicates whether to skip errors. Valid values:
	//
	// 	- **true**: skips the error and continues to execute SQL statements.
	//
	// 	- **false**: stops executing SQL statements.
	//
	// example:
	//
	// false
	IgnoreError *bool `json:"IgnoreError,omitempty" xml:"IgnoreError,omitempty"`
	// The information about the source database.
	SourceDatabaseInfo *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceDatabaseInfo `json:"SourceDatabaseInfo,omitempty" xml:"SourceDatabaseInfo,omitempty" type:"Struct"`
	// The schema version of the source database. Valid values:
	//
	// 	- **DATASOURCE**: the default latest version of the system
	//
	// 	- **VERSION**: a previous schema version that you manually specify
	//
	// example:
	//
	// VERSION
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The version information about the source instance.
	//
	// > This parameter is displayed only when the value of the **SourceType*	- parameter is **VERSION**.
	SourceVersionInfo *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceVersionInfo `json:"SourceVersionInfo,omitempty" xml:"SourceVersionInfo,omitempty" type:"Struct"`
	// The information about the table whose schema you want to synchronize.
	TableInfoList []*GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTableInfoList `json:"TableInfoList,omitempty" xml:"TableInfoList,omitempty" type:"Repeated"`
	// The information about the destination database.
	TargetDatabaseInfo *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetDatabaseInfo `json:"TargetDatabaseInfo,omitempty" xml:"TargetDatabaseInfo,omitempty" type:"Struct"`
	// The schema version of the destination database. Valid values:
	//
	// 	- **DATASOURCE**: the default latest version of the system
	//
	// 	- **VERSION**: a previous schema version that you manually specify
	//
	// example:
	//
	// DATASOURCE
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The version information about the destination instance.
	//
	// > This parameter is displayed only when the value of the **SourceType*	- parameter is **VERSION**.
	TargetVersionInfo *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetVersionInfo `json:"TargetVersionInfo,omitempty" xml:"TargetVersionInfo,omitempty" type:"Struct"`
}

func (s GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail) String() string {
	return dara.Prettify(s)
}

func (s GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail) GoString() string {
	return s.String()
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail) GetIgnoreError() *bool {
	return s.IgnoreError
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail) GetSourceDatabaseInfo() *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceDatabaseInfo {
	return s.SourceDatabaseInfo
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail) GetSourceType() *string {
	return s.SourceType
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail) GetSourceVersionInfo() *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceVersionInfo {
	return s.SourceVersionInfo
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail) GetTableInfoList() []*GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTableInfoList {
	return s.TableInfoList
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail) GetTargetDatabaseInfo() *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetDatabaseInfo {
	return s.TargetDatabaseInfo
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail) GetTargetType() *string {
	return s.TargetType
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail) GetTargetVersionInfo() *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetVersionInfo {
	return s.TargetVersionInfo
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail) SetIgnoreError(v bool) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail {
	s.IgnoreError = &v
	return s
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail) SetSourceDatabaseInfo(v *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceDatabaseInfo) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail {
	s.SourceDatabaseInfo = v
	return s
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail) SetSourceType(v string) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail {
	s.SourceType = &v
	return s
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail) SetSourceVersionInfo(v *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceVersionInfo) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail {
	s.SourceVersionInfo = v
	return s
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail) SetTableInfoList(v []*GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTableInfoList) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail {
	s.TableInfoList = v
	return s
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail) SetTargetDatabaseInfo(v *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetDatabaseInfo) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail {
	s.TargetDatabaseInfo = v
	return s
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail) SetTargetType(v string) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail {
	s.TargetType = &v
	return s
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail) SetTargetVersionInfo(v *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetVersionInfo) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail {
	s.TargetVersionInfo = v
	return s
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail) Validate() error {
	return dara.Validate(s)
}

type GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceDatabaseInfo struct {
	// The ID of the source database.
	//
	// example:
	//
	// 432532
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The type of the database engine.
	//
	// example:
	//
	// MySQL
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The type of the environment to which the database instance belongs. For more information, see [Change the environment type of an instance](https://help.aliyun.com/document_detail/163309.html).
	//
	// example:
	//
	// dev
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// Indicates whether the database is a logical database. Valid values:
	//
	// 	- **true**: The database is a logical database.
	//
	// 	- **false**: The database is not a logical database
	//
	// example:
	//
	// false
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// The name that is used to search for the database.
	//
	// example:
	//
	// test
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
}

func (s GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceDatabaseInfo) String() string {
	return dara.Prettify(s)
}

func (s GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceDatabaseInfo) GoString() string {
	return s.String()
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceDatabaseInfo) GetDbId() *int64 {
	return s.DbId
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceDatabaseInfo) GetDbType() *string {
	return s.DbType
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceDatabaseInfo) GetEnvType() *string {
	return s.EnvType
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceDatabaseInfo) GetLogic() *bool {
	return s.Logic
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceDatabaseInfo) GetSearchName() *string {
	return s.SearchName
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceDatabaseInfo) SetDbId(v int64) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceDatabaseInfo {
	s.DbId = &v
	return s
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceDatabaseInfo) SetDbType(v string) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceDatabaseInfo {
	s.DbType = &v
	return s
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceDatabaseInfo) SetEnvType(v string) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceDatabaseInfo {
	s.EnvType = &v
	return s
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceDatabaseInfo) SetLogic(v bool) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceDatabaseInfo {
	s.Logic = &v
	return s
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceDatabaseInfo) SetSearchName(v string) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceDatabaseInfo {
	s.SearchName = &v
	return s
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceDatabaseInfo) Validate() error {
	return dara.Validate(s)
}

type GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceVersionInfo struct {
	// The version number.
	//
	// example:
	//
	// e179bbb8163dcdcfacda24858bedb4d8006ae2b8
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceVersionInfo) String() string {
	return dara.Prettify(s)
}

func (s GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceVersionInfo) GoString() string {
	return s.String()
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceVersionInfo) GetVersionId() *string {
	return s.VersionId
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceVersionInfo) SetVersionId(v string) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceVersionInfo {
	s.VersionId = &v
	return s
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceVersionInfo) Validate() error {
	return dara.Validate(s)
}

type GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTableInfoList struct {
	// The name of the table whose schema you want to synchronize.
	//
	// example:
	//
	// test_tbl
	SourceTableName *string `json:"SourceTableName,omitempty" xml:"SourceTableName,omitempty"`
	// The name of the table to which you want to synchronize the schema of a table.
	//
	// example:
	//
	// test_tbl
	TargetTableName *string `json:"TargetTableName,omitempty" xml:"TargetTableName,omitempty"`
}

func (s GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTableInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTableInfoList) GoString() string {
	return s.String()
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTableInfoList) GetSourceTableName() *string {
	return s.SourceTableName
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTableInfoList) GetTargetTableName() *string {
	return s.TargetTableName
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTableInfoList) SetSourceTableName(v string) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTableInfoList {
	s.SourceTableName = &v
	return s
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTableInfoList) SetTargetTableName(v string) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTableInfoList {
	s.TargetTableName = &v
	return s
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTableInfoList) Validate() error {
	return dara.Validate(s)
}

type GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetDatabaseInfo struct {
	// The ID of the destination database.
	//
	// example:
	//
	// 432543
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The type of the database engine.
	//
	// example:
	//
	// MySQL
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The type of the environment to which the database instance belongs. For more information, see [Change the environment type of an instance](https://help.aliyun.com/document_detail/163309.html).
	//
	// example:
	//
	// dev
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// Indicates whether the database is a logical database. Valid values:
	//
	// 	- **true**: The database is a logical database.
	//
	// 	- **false**: The database is not a logical database
	//
	// example:
	//
	// false
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// The name that is used to search for the database.
	//
	// example:
	//
	// test
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
}

func (s GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetDatabaseInfo) String() string {
	return dara.Prettify(s)
}

func (s GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetDatabaseInfo) GoString() string {
	return s.String()
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetDatabaseInfo) GetDbId() *int64 {
	return s.DbId
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetDatabaseInfo) GetDbType() *string {
	return s.DbType
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetDatabaseInfo) GetEnvType() *string {
	return s.EnvType
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetDatabaseInfo) GetLogic() *bool {
	return s.Logic
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetDatabaseInfo) GetSearchName() *string {
	return s.SearchName
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetDatabaseInfo) SetDbId(v int64) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetDatabaseInfo {
	s.DbId = &v
	return s
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetDatabaseInfo) SetDbType(v string) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetDatabaseInfo {
	s.DbType = &v
	return s
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetDatabaseInfo) SetEnvType(v string) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetDatabaseInfo {
	s.EnvType = &v
	return s
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetDatabaseInfo) SetLogic(v bool) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetDatabaseInfo {
	s.Logic = &v
	return s
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetDatabaseInfo) SetSearchName(v string) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetDatabaseInfo {
	s.SearchName = &v
	return s
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetDatabaseInfo) Validate() error {
	return dara.Validate(s)
}

type GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetVersionInfo struct {
	// The version number.
	//
	// example:
	//
	// e179bbb8163dcdcfacda24858bedb4d8006ae2b8
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetVersionInfo) String() string {
	return dara.Prettify(s)
}

func (s GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetVersionInfo) GoString() string {
	return s.String()
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetVersionInfo) GetVersionId() *string {
	return s.VersionId
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetVersionInfo) SetVersionId(v string) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetVersionInfo {
	s.VersionId = &v
	return s
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetVersionInfo) Validate() error {
	return dara.Validate(s)
}
