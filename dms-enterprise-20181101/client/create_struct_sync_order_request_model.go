// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStructSyncOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachmentKey(v string) *CreateStructSyncOrderRequest
	GetAttachmentKey() *string
	SetComment(v string) *CreateStructSyncOrderRequest
	GetComment() *string
	SetParam(v *CreateStructSyncOrderRequestParam) *CreateStructSyncOrderRequest
	GetParam() *CreateStructSyncOrderRequestParam
	SetRelatedUserList(v []*int64) *CreateStructSyncOrderRequest
	GetRelatedUserList() []*int64
	SetTid(v int64) *CreateStructSyncOrderRequest
	GetTid() *int64
}

type CreateStructSyncOrderRequest struct {
	// The key of an attachment that is returned after the attachment is uploaded. You can call the [GetUserUploadFileJob](https://help.aliyun.com/document_detail/206069.html) operation to query the key of the attachment.
	//
	// example:
	//
	// upload_3c7edea3-e4c3-4403-857d-737043036f69_test.sql
	AttachmentKey *string `json:"AttachmentKey,omitempty" xml:"AttachmentKey,omitempty"`
	// The remarks of the ticket.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The parameters of the ticket.
	//
	// This parameter is required.
	Param *CreateStructSyncOrderRequestParam `json:"Param,omitempty" xml:"Param,omitempty" type:"Struct"`
	// The IDs of the stakeholders.
	RelatedUserList []*int64 `json:"RelatedUserList,omitempty" xml:"RelatedUserList,omitempty" type:"Repeated"`
	// The ID of the tenant.
	//
	// >  To view the tenant ID, move the pointer over the profile picture in the upper-right corner of the DMS console. For more information, see [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 1
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateStructSyncOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStructSyncOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateStructSyncOrderRequest) GetAttachmentKey() *string {
	return s.AttachmentKey
}

func (s *CreateStructSyncOrderRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateStructSyncOrderRequest) GetParam() *CreateStructSyncOrderRequestParam {
	return s.Param
}

func (s *CreateStructSyncOrderRequest) GetRelatedUserList() []*int64 {
	return s.RelatedUserList
}

func (s *CreateStructSyncOrderRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateStructSyncOrderRequest) SetAttachmentKey(v string) *CreateStructSyncOrderRequest {
	s.AttachmentKey = &v
	return s
}

func (s *CreateStructSyncOrderRequest) SetComment(v string) *CreateStructSyncOrderRequest {
	s.Comment = &v
	return s
}

func (s *CreateStructSyncOrderRequest) SetParam(v *CreateStructSyncOrderRequestParam) *CreateStructSyncOrderRequest {
	s.Param = v
	return s
}

func (s *CreateStructSyncOrderRequest) SetRelatedUserList(v []*int64) *CreateStructSyncOrderRequest {
	s.RelatedUserList = v
	return s
}

func (s *CreateStructSyncOrderRequest) SetTid(v int64) *CreateStructSyncOrderRequest {
	s.Tid = &v
	return s
}

func (s *CreateStructSyncOrderRequest) Validate() error {
	return dara.Validate(s)
}

type CreateStructSyncOrderRequestParam struct {
	// Specifies whether to skip an error that occurs in executing an SQL statement. Valid values:
	//
	// 	- **true**: continues to execute subsequent SQL statements if an error occurs in executing an SQL statement.
	//
	// 	- **false**: stops executing subsequent SQL statements if an error occurs in executing an SQL statement.
	//
	// example:
	//
	// false
	IgnoreError *bool `json:"IgnoreError,omitempty" xml:"IgnoreError,omitempty"`
	// The information about the base database.
	//
	// This parameter is required.
	Source *CreateStructSyncOrderRequestParamSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	// The information about the table of which you want to synchronize the schema.
	TableInfoList []*CreateStructSyncOrderRequestParamTableInfoList `json:"TableInfoList,omitempty" xml:"TableInfoList,omitempty" type:"Repeated"`
	// The information about the database to which you want to synchronize the schema of a table.
	//
	// This parameter is required.
	Target *CreateStructSyncOrderRequestParamTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
}

func (s CreateStructSyncOrderRequestParam) String() string {
	return dara.Prettify(s)
}

func (s CreateStructSyncOrderRequestParam) GoString() string {
	return s.String()
}

func (s *CreateStructSyncOrderRequestParam) GetIgnoreError() *bool {
	return s.IgnoreError
}

func (s *CreateStructSyncOrderRequestParam) GetSource() *CreateStructSyncOrderRequestParamSource {
	return s.Source
}

func (s *CreateStructSyncOrderRequestParam) GetTableInfoList() []*CreateStructSyncOrderRequestParamTableInfoList {
	return s.TableInfoList
}

func (s *CreateStructSyncOrderRequestParam) GetTarget() *CreateStructSyncOrderRequestParamTarget {
	return s.Target
}

func (s *CreateStructSyncOrderRequestParam) SetIgnoreError(v bool) *CreateStructSyncOrderRequestParam {
	s.IgnoreError = &v
	return s
}

func (s *CreateStructSyncOrderRequestParam) SetSource(v *CreateStructSyncOrderRequestParamSource) *CreateStructSyncOrderRequestParam {
	s.Source = v
	return s
}

func (s *CreateStructSyncOrderRequestParam) SetTableInfoList(v []*CreateStructSyncOrderRequestParamTableInfoList) *CreateStructSyncOrderRequestParam {
	s.TableInfoList = v
	return s
}

func (s *CreateStructSyncOrderRequestParam) SetTarget(v *CreateStructSyncOrderRequestParamTarget) *CreateStructSyncOrderRequestParam {
	s.Target = v
	return s
}

func (s *CreateStructSyncOrderRequestParam) Validate() error {
	return dara.Validate(s)
}

type CreateStructSyncOrderRequestParamSource struct {
	// The ID of the database. You can call the [SearchDatabases](https://help.aliyun.com/document_detail/141876.html) operation to query the ID of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4324532
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The name that is used to search for the database. You can call the [SearchDatabases](https://help.aliyun.com/document_detail/141876.html) operation to query the name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	DbSearchName *string `json:"DbSearchName,omitempty" xml:"DbSearchName,omitempty"`
	// Specifies whether the database is a logical database. Valid values:
	//
	// 	- **true**: The database is a logical database.
	//
	// 	- **false**: The database is not a logical database.
	//
	// example:
	//
	// false
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// The version number of the schema. The default value is the latest schema version number. For more information, see [Manage schema versions](https://help.aliyun.com/document_detail/202275.html).
	//
	// example:
	//
	// e179bbb8163dcdcfacda24858bedb4d8006ae2b8
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s CreateStructSyncOrderRequestParamSource) String() string {
	return dara.Prettify(s)
}

func (s CreateStructSyncOrderRequestParamSource) GoString() string {
	return s.String()
}

func (s *CreateStructSyncOrderRequestParamSource) GetDbId() *int64 {
	return s.DbId
}

func (s *CreateStructSyncOrderRequestParamSource) GetDbSearchName() *string {
	return s.DbSearchName
}

func (s *CreateStructSyncOrderRequestParamSource) GetLogic() *bool {
	return s.Logic
}

func (s *CreateStructSyncOrderRequestParamSource) GetVersionId() *string {
	return s.VersionId
}

func (s *CreateStructSyncOrderRequestParamSource) SetDbId(v int64) *CreateStructSyncOrderRequestParamSource {
	s.DbId = &v
	return s
}

func (s *CreateStructSyncOrderRequestParamSource) SetDbSearchName(v string) *CreateStructSyncOrderRequestParamSource {
	s.DbSearchName = &v
	return s
}

func (s *CreateStructSyncOrderRequestParamSource) SetLogic(v bool) *CreateStructSyncOrderRequestParamSource {
	s.Logic = &v
	return s
}

func (s *CreateStructSyncOrderRequestParamSource) SetVersionId(v string) *CreateStructSyncOrderRequestParamSource {
	s.VersionId = &v
	return s
}

func (s *CreateStructSyncOrderRequestParamSource) Validate() error {
	return dara.Validate(s)
}

type CreateStructSyncOrderRequestParamTableInfoList struct {
	// The name of the source table.
	//
	// example:
	//
	// test_tbl
	SourceTableName *string `json:"SourceTableName,omitempty" xml:"SourceTableName,omitempty"`
	// The name of the destination table.
	//
	// example:
	//
	// test_tbl
	TargetTableName *string `json:"TargetTableName,omitempty" xml:"TargetTableName,omitempty"`
}

func (s CreateStructSyncOrderRequestParamTableInfoList) String() string {
	return dara.Prettify(s)
}

func (s CreateStructSyncOrderRequestParamTableInfoList) GoString() string {
	return s.String()
}

func (s *CreateStructSyncOrderRequestParamTableInfoList) GetSourceTableName() *string {
	return s.SourceTableName
}

func (s *CreateStructSyncOrderRequestParamTableInfoList) GetTargetTableName() *string {
	return s.TargetTableName
}

func (s *CreateStructSyncOrderRequestParamTableInfoList) SetSourceTableName(v string) *CreateStructSyncOrderRequestParamTableInfoList {
	s.SourceTableName = &v
	return s
}

func (s *CreateStructSyncOrderRequestParamTableInfoList) SetTargetTableName(v string) *CreateStructSyncOrderRequestParamTableInfoList {
	s.TargetTableName = &v
	return s
}

func (s *CreateStructSyncOrderRequestParamTableInfoList) Validate() error {
	return dara.Validate(s)
}

type CreateStructSyncOrderRequestParamTarget struct {
	// The ID of the database. You can call the [SearchDatabases](https://help.aliyun.com/document_detail/141876.html) operation to query the ID of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// 432432
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The name that is used to search for the database. You can call the [SearchDatabases](https://help.aliyun.com/document_detail/141876.html) operation to query the name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	DbSearchName *string `json:"DbSearchName,omitempty" xml:"DbSearchName,omitempty"`
	// Specifies whether the database is a logical database. Valid values:
	//
	// 	- **true**: The database is a logical database.
	//
	// 	- **false**: The database is not a logical database.
	//
	// example:
	//
	// false
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// The version number. By default, this parameter is left empty.
	//
	// >  If you specify the schema version number of the destination database, Data Management (DMS) only compares the schemas of the two databases.
	//
	// example:
	//
	// e179bbb8163dcdcfacda24858bedb4d8006ae2b8
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s CreateStructSyncOrderRequestParamTarget) String() string {
	return dara.Prettify(s)
}

func (s CreateStructSyncOrderRequestParamTarget) GoString() string {
	return s.String()
}

func (s *CreateStructSyncOrderRequestParamTarget) GetDbId() *int64 {
	return s.DbId
}

func (s *CreateStructSyncOrderRequestParamTarget) GetDbSearchName() *string {
	return s.DbSearchName
}

func (s *CreateStructSyncOrderRequestParamTarget) GetLogic() *bool {
	return s.Logic
}

func (s *CreateStructSyncOrderRequestParamTarget) GetVersionId() *string {
	return s.VersionId
}

func (s *CreateStructSyncOrderRequestParamTarget) SetDbId(v int64) *CreateStructSyncOrderRequestParamTarget {
	s.DbId = &v
	return s
}

func (s *CreateStructSyncOrderRequestParamTarget) SetDbSearchName(v string) *CreateStructSyncOrderRequestParamTarget {
	s.DbSearchName = &v
	return s
}

func (s *CreateStructSyncOrderRequestParamTarget) SetLogic(v bool) *CreateStructSyncOrderRequestParamTarget {
	s.Logic = &v
	return s
}

func (s *CreateStructSyncOrderRequestParamTarget) SetVersionId(v string) *CreateStructSyncOrderRequestParamTarget {
	s.VersionId = &v
	return s
}

func (s *CreateStructSyncOrderRequestParamTarget) Validate() error {
	return dara.Validate(s)
}
