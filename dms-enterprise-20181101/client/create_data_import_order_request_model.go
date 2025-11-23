// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataImportOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachmentKey(v string) *CreateDataImportOrderRequest
	GetAttachmentKey() *string
	SetComment(v string) *CreateDataImportOrderRequest
	GetComment() *string
	SetParam(v *CreateDataImportOrderRequestParam) *CreateDataImportOrderRequest
	GetParam() *CreateDataImportOrderRequestParam
	SetRealLoginUserUid(v string) *CreateDataImportOrderRequest
	GetRealLoginUserUid() *string
	SetRelatedUserList(v []*int64) *CreateDataImportOrderRequest
	GetRelatedUserList() []*int64
	SetTid(v int64) *CreateDataImportOrderRequest
	GetTid() *int64
}

type CreateDataImportOrderRequest struct {
	// The key of the attachment that provides more instructions for the ticket. You can call the [GetUserUploadFileJob](https://help.aliyun.com/document_detail/206069.html) operation to obtain the attachment key from the value of the AttachmentKey parameter.
	//
	// example:
	//
	// order_attachment.txt
	AttachmentKey *string `json:"AttachmentKey,omitempty" xml:"AttachmentKey,omitempty"`
	// The purpose or objective of the data import. This parameter is used to help reduce unnecessary communication.
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
	Param            *CreateDataImportOrderRequestParam `json:"Param,omitempty" xml:"Param,omitempty" type:"Struct"`
	RealLoginUserUid *string                            `json:"RealLoginUserUid,omitempty" xml:"RealLoginUserUid,omitempty"`
	// The stakeholders of the data import. All stakeholders can view the ticket details and assist in the approval process. Irrelevant users other than DMS administrators and database administrators (DBAs) are not allowed to view the ticket details.
	RelatedUserList []*int64 `json:"RelatedUserList,omitempty" xml:"RelatedUserList,omitempty" type:"Repeated"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateDataImportOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataImportOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateDataImportOrderRequest) GetAttachmentKey() *string {
	return s.AttachmentKey
}

func (s *CreateDataImportOrderRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateDataImportOrderRequest) GetParam() *CreateDataImportOrderRequestParam {
	return s.Param
}

func (s *CreateDataImportOrderRequest) GetRealLoginUserUid() *string {
	return s.RealLoginUserUid
}

func (s *CreateDataImportOrderRequest) GetRelatedUserList() []*int64 {
	return s.RelatedUserList
}

func (s *CreateDataImportOrderRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateDataImportOrderRequest) SetAttachmentKey(v string) *CreateDataImportOrderRequest {
	s.AttachmentKey = &v
	return s
}

func (s *CreateDataImportOrderRequest) SetComment(v string) *CreateDataImportOrderRequest {
	s.Comment = &v
	return s
}

func (s *CreateDataImportOrderRequest) SetParam(v *CreateDataImportOrderRequestParam) *CreateDataImportOrderRequest {
	s.Param = v
	return s
}

func (s *CreateDataImportOrderRequest) SetRealLoginUserUid(v string) *CreateDataImportOrderRequest {
	s.RealLoginUserUid = &v
	return s
}

func (s *CreateDataImportOrderRequest) SetRelatedUserList(v []*int64) *CreateDataImportOrderRequest {
	s.RelatedUserList = v
	return s
}

func (s *CreateDataImportOrderRequest) SetTid(v int64) *CreateDataImportOrderRequest {
	s.Tid = &v
	return s
}

func (s *CreateDataImportOrderRequest) Validate() error {
	if s.Param != nil {
		if err := s.Param.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataImportOrderRequestParam struct {
	// The key of the attachment that contains the SQL statements used to import data. You can call the [GetUserUploadFileJob](https://help.aliyun.com/document_detail/206069.html) operation to the attachment key from the value of the AttachmentKey parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// order_text
	AttachmentName *string `json:"AttachmentName,omitempty" xml:"AttachmentName,omitempty"`
	// The reason for the data import.
	//
	// example:
	//
	// test
	Classify *string `json:"Classify,omitempty" xml:"Classify,omitempty"`
	// The type of the CSV file. Valid values:
	//
	// 	- **true**: The first row in the CSV file contains field names.
	//
	// 	- **false**: The first row in the CSV file contains data.
	//
	// >  This parameter is required if you set the **FileType*	- parameter to **CSV**.
	//
	// example:
	//
	// true
	CsvFirstRowIsColumnDef *bool `json:"CsvFirstRowIsColumnDef,omitempty" xml:"CsvFirstRowIsColumnDef,omitempty"`
	// The database to which you want to import data. You can specify only one database.
	//
	// This parameter is required.
	DbItemList []*CreateDataImportOrderRequestParamDbItemList `json:"DbItemList,omitempty" xml:"DbItemList,omitempty" type:"Repeated"`
	ExecMode   *string                                        `json:"ExecMode,omitempty" xml:"ExecMode,omitempty"`
	// The encoding algorithm to be used by the destination database. Valid values:
	//
	// 	- **AUTO**: automatic identification
	//
	// 	- **UTF-8**: UTF-8 encoding
	//
	// 	- **GBK**: GBK encoding
	//
	// 	- **ISO-8859-1**: ISO-8859-1 encoding
	//
	// example:
	//
	// AUTO
	FileEncoding *string `json:"FileEncoding,omitempty" xml:"FileEncoding,omitempty"`
	// The format of the file for the data import. Valid values:
	//
	// 	- **SQL**: an SQL file
	//
	// 	- **CSV**: a CSV file
	//
	// This parameter is required.
	//
	// example:
	//
	// CSV
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// Specifies whether to skip an error that occurs. Valid values:
	//
	// 	- **true**: skips the error and continues to execute SQL statements.
	//
	// 	- **false**: stops executing SQL statements.
	//
	// example:
	//
	// false
	IgnoreError *bool `json:"IgnoreError,omitempty" xml:"IgnoreError,omitempty"`
	// The import mode. Valid values:
	//
	// 	- **FAST_MODE**: In the Execute step, the uploaded file is read and SQL statements are executed to import data to the specified destination database. Compared with the security mode, this mode can be used to import data in a less secure but more efficient manner.
	//
	// 	- **SAFE_MODE**: In the Precheck step, the uploaded file is parsed, and SQL statements or CSV file data is cached. In the Execute step, the cached SQL statements are read and executed to import data, or the cached CSV file data is read and imported to the specified destination database. This mode can be used to import data in a more secure but less efficient manner.
	//
	// example:
	//
	// FAST_MODE
	ImportMode *string `json:"ImportMode,omitempty" xml:"ImportMode,omitempty"`
	// The mode in which the data in the CSV format is to be written to the destination table. Valid values:
	//
	// 	- **INSERT**: The database checks the primary key when data is written. If a duplicate primary key value exists, an error message is returned.
	//
	// 	- **INSERT_IGNORE**: If the imported data contains data records that are the same as those in the destination table, the new data records are ignored.
	//
	// 	- **REPLACE_INTO**: If the imported data contains a row that has the same value for the primary key or unique index as one row in the destination table, the database deletes the existing row and inserts the new row into the destination table.
	//
	// >  This parameter is required if you set the **FileType*	- parameter to **CSV**.
	//
	// example:
	//
	// INSERT
	InsertType *string `json:"InsertType,omitempty" xml:"InsertType,omitempty"`
	// The key of the attachment that contains the SQL statements used to roll back the data import. You can call the [GetUserUploadFileJob](https://help.aliyun.com/document_detail/206069.html) operation to obtain the attachment key from the value of the AttachmentKey parameter.
	//
	// >  This parameter is required if you set the **RollbackSqlType*	- parameter to **ATTACHMENT**.
	//
	// example:
	//
	// rollback.sql
	RollbackAttachmentName *string `json:"RollbackAttachmentName,omitempty" xml:"RollbackAttachmentName,omitempty"`
	// The SQL statements used to roll back the data import.
	//
	// >  This parameter is required if you set the **RollbackSqlType*	- parameter to **TEXT**.
	//
	// example:
	//
	// empty
	RollbackSQL *string `json:"RollbackSQL,omitempty" xml:"RollbackSQL,omitempty"`
	// The format of the SQL statements used to roll back the data import. Valid values:
	//
	// 	- **TEXT**: text
	//
	// 	- **ATTACHMENT**: attachment
	//
	// example:
	//
	// TEXT
	RollbackSqlType *string `json:"RollbackSqlType,omitempty" xml:"RollbackSqlType,omitempty"`
	// The destination table to which you want to import the data in the CSV format.
	//
	// >  This parameter is required if you set the **FileType*	- parameter to **CSV**.
	//
	// example:
	//
	// Table_text
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s CreateDataImportOrderRequestParam) String() string {
	return dara.Prettify(s)
}

func (s CreateDataImportOrderRequestParam) GoString() string {
	return s.String()
}

func (s *CreateDataImportOrderRequestParam) GetAttachmentName() *string {
	return s.AttachmentName
}

func (s *CreateDataImportOrderRequestParam) GetClassify() *string {
	return s.Classify
}

func (s *CreateDataImportOrderRequestParam) GetCsvFirstRowIsColumnDef() *bool {
	return s.CsvFirstRowIsColumnDef
}

func (s *CreateDataImportOrderRequestParam) GetDbItemList() []*CreateDataImportOrderRequestParamDbItemList {
	return s.DbItemList
}

func (s *CreateDataImportOrderRequestParam) GetExecMode() *string {
	return s.ExecMode
}

func (s *CreateDataImportOrderRequestParam) GetFileEncoding() *string {
	return s.FileEncoding
}

func (s *CreateDataImportOrderRequestParam) GetFileType() *string {
	return s.FileType
}

func (s *CreateDataImportOrderRequestParam) GetIgnoreError() *bool {
	return s.IgnoreError
}

func (s *CreateDataImportOrderRequestParam) GetImportMode() *string {
	return s.ImportMode
}

func (s *CreateDataImportOrderRequestParam) GetInsertType() *string {
	return s.InsertType
}

func (s *CreateDataImportOrderRequestParam) GetRollbackAttachmentName() *string {
	return s.RollbackAttachmentName
}

func (s *CreateDataImportOrderRequestParam) GetRollbackSQL() *string {
	return s.RollbackSQL
}

func (s *CreateDataImportOrderRequestParam) GetRollbackSqlType() *string {
	return s.RollbackSqlType
}

func (s *CreateDataImportOrderRequestParam) GetTableName() *string {
	return s.TableName
}

func (s *CreateDataImportOrderRequestParam) SetAttachmentName(v string) *CreateDataImportOrderRequestParam {
	s.AttachmentName = &v
	return s
}

func (s *CreateDataImportOrderRequestParam) SetClassify(v string) *CreateDataImportOrderRequestParam {
	s.Classify = &v
	return s
}

func (s *CreateDataImportOrderRequestParam) SetCsvFirstRowIsColumnDef(v bool) *CreateDataImportOrderRequestParam {
	s.CsvFirstRowIsColumnDef = &v
	return s
}

func (s *CreateDataImportOrderRequestParam) SetDbItemList(v []*CreateDataImportOrderRequestParamDbItemList) *CreateDataImportOrderRequestParam {
	s.DbItemList = v
	return s
}

func (s *CreateDataImportOrderRequestParam) SetExecMode(v string) *CreateDataImportOrderRequestParam {
	s.ExecMode = &v
	return s
}

func (s *CreateDataImportOrderRequestParam) SetFileEncoding(v string) *CreateDataImportOrderRequestParam {
	s.FileEncoding = &v
	return s
}

func (s *CreateDataImportOrderRequestParam) SetFileType(v string) *CreateDataImportOrderRequestParam {
	s.FileType = &v
	return s
}

func (s *CreateDataImportOrderRequestParam) SetIgnoreError(v bool) *CreateDataImportOrderRequestParam {
	s.IgnoreError = &v
	return s
}

func (s *CreateDataImportOrderRequestParam) SetImportMode(v string) *CreateDataImportOrderRequestParam {
	s.ImportMode = &v
	return s
}

func (s *CreateDataImportOrderRequestParam) SetInsertType(v string) *CreateDataImportOrderRequestParam {
	s.InsertType = &v
	return s
}

func (s *CreateDataImportOrderRequestParam) SetRollbackAttachmentName(v string) *CreateDataImportOrderRequestParam {
	s.RollbackAttachmentName = &v
	return s
}

func (s *CreateDataImportOrderRequestParam) SetRollbackSQL(v string) *CreateDataImportOrderRequestParam {
	s.RollbackSQL = &v
	return s
}

func (s *CreateDataImportOrderRequestParam) SetRollbackSqlType(v string) *CreateDataImportOrderRequestParam {
	s.RollbackSqlType = &v
	return s
}

func (s *CreateDataImportOrderRequestParam) SetTableName(v string) *CreateDataImportOrderRequestParam {
	s.TableName = &v
	return s
}

func (s *CreateDataImportOrderRequestParam) Validate() error {
	if s.DbItemList != nil {
		for _, item := range s.DbItemList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateDataImportOrderRequestParamDbItemList struct {
	// The ID of the database. The database can be a physical database or a logical database.
	//
	// 	- To obtain the ID of a physical database, call the [ListDatabases](https://help.aliyun.com/document_detail/141873.html) or [SearchDatabase](https://help.aliyun.com/document_detail/141876.html) operation.
	//
	// 	- To obtain the ID of a logical database, call the [ListLogicDatabases](https://help.aliyun.com/document_detail/141874.html) or [SearchDatabase](https://help.aliyun.com/document_detail/141876.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1860****
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// Specifies whether the database is a logical database. Valid values:
	//
	// 	- **true**: The database is a logical database.
	//
	// 	- **false**: The database is a physical database.
	//
	// >  If you set this parameter to **true**, the database that you specify must be a logical database.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
}

func (s CreateDataImportOrderRequestParamDbItemList) String() string {
	return dara.Prettify(s)
}

func (s CreateDataImportOrderRequestParamDbItemList) GoString() string {
	return s.String()
}

func (s *CreateDataImportOrderRequestParamDbItemList) GetDbId() *int64 {
	return s.DbId
}

func (s *CreateDataImportOrderRequestParamDbItemList) GetLogic() *bool {
	return s.Logic
}

func (s *CreateDataImportOrderRequestParamDbItemList) SetDbId(v int64) *CreateDataImportOrderRequestParamDbItemList {
	s.DbId = &v
	return s
}

func (s *CreateDataImportOrderRequestParamDbItemList) SetLogic(v bool) *CreateDataImportOrderRequestParamDbItemList {
	s.Logic = &v
	return s
}

func (s *CreateDataImportOrderRequestParamDbItemList) Validate() error {
	return dara.Validate(s)
}
