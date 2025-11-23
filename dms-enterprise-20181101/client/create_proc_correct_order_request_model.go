// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProcCorrectOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachmentKey(v string) *CreateProcCorrectOrderRequest
	GetAttachmentKey() *string
	SetComment(v string) *CreateProcCorrectOrderRequest
	GetComment() *string
	SetParam(v *CreateProcCorrectOrderRequestParam) *CreateProcCorrectOrderRequest
	GetParam() *CreateProcCorrectOrderRequestParam
	SetRelatedUserList(v []*int64) *CreateProcCorrectOrderRequest
	GetRelatedUserList() []*int64
	SetTid(v int64) *CreateProcCorrectOrderRequest
	GetTid() *int64
}

type CreateProcCorrectOrderRequest struct {
	// The key of the attachment for the ticket. The attachment provides more instructions for this operation.
	//
	// You can call the [GetUserUploadFileJob](https://help.aliyun.com/document_detail/206069.html) operation to query the key of the attachment.
	//
	// example:
	//
	// order_attachment.txt
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
	Param *CreateProcCorrectOrderRequestParam `json:"Param,omitempty" xml:"Param,omitempty" type:"Struct"`
	// The operators that are related to the ticket.
	RelatedUserList []*int64 `json:"RelatedUserList,omitempty" xml:"RelatedUserList,omitempty" type:"Repeated"`
	// The ID of the tenant.
	//
	// >  To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the "Manage DMS tenants" topic.
	//
	// example:
	//
	// 4***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateProcCorrectOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProcCorrectOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateProcCorrectOrderRequest) GetAttachmentKey() *string {
	return s.AttachmentKey
}

func (s *CreateProcCorrectOrderRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateProcCorrectOrderRequest) GetParam() *CreateProcCorrectOrderRequestParam {
	return s.Param
}

func (s *CreateProcCorrectOrderRequest) GetRelatedUserList() []*int64 {
	return s.RelatedUserList
}

func (s *CreateProcCorrectOrderRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateProcCorrectOrderRequest) SetAttachmentKey(v string) *CreateProcCorrectOrderRequest {
	s.AttachmentKey = &v
	return s
}

func (s *CreateProcCorrectOrderRequest) SetComment(v string) *CreateProcCorrectOrderRequest {
	s.Comment = &v
	return s
}

func (s *CreateProcCorrectOrderRequest) SetParam(v *CreateProcCorrectOrderRequestParam) *CreateProcCorrectOrderRequest {
	s.Param = v
	return s
}

func (s *CreateProcCorrectOrderRequest) SetRelatedUserList(v []*int64) *CreateProcCorrectOrderRequest {
	s.RelatedUserList = v
	return s
}

func (s *CreateProcCorrectOrderRequest) SetTid(v int64) *CreateProcCorrectOrderRequest {
	s.Tid = &v
	return s
}

func (s *CreateProcCorrectOrderRequest) Validate() error {
	if s.Param != nil {
		if err := s.Param.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateProcCorrectOrderRequestParam struct {
	// The reason for the programmable object change.
	//
	// example:
	//
	// test
	Classify *string `json:"Classify,omitempty" xml:"Classify,omitempty"`
	// The information about the database.
	//
	// This parameter is required.
	DbItemList []*CreateProcCorrectOrderRequestParamDbItemList `json:"DbItemList,omitempty" xml:"DbItemList,omitempty" type:"Repeated"`
	// The mode in which the data change ticket is executed after the ticket is approved. Valid values:
	//
	// 	- **COMMITOR**: The ticket is executed by the user who submits the ticket.
	//
	// 	- **AUTO**: The ticket is automatically executed after the ticket is approved.
	//
	// 	- **LAST_AUDITOR**: The ticket is executed by the last approver of the ticket.
	//
	// example:
	//
	// COMMITOR
	ExecMode *string `json:"ExecMode,omitempty" xml:"ExecMode,omitempty"`
	// The SQL statements for data change.
	//
	// This parameter is required.
	//
	// example:
	//
	// DELIMITER ///
	//
	// CREATE PROCEDURE GetAllProducts()
	//
	// BEGIN
	//
	// SELECT 	- FROM base_user;
	//
	// END ///
	ExecSQL *string `json:"ExecSQL,omitempty" xml:"ExecSQL,omitempty"`
	// The key of the attachment that contains the SQL statements used to roll back the data change. You can call the [GetUserUploadFileJob](https://help.aliyun.com/document_detail/206069.html) operation to obtain the attachment key from the value of AttachmentKey.
	//
	// >  This parameter is required if you set **RollbackSqlType*	- to **ATTACHMENT**.
	//
	// example:
	//
	// test_rollback.sql
	RollbackAttachmentName *string `json:"RollbackAttachmentName,omitempty" xml:"RollbackAttachmentName,omitempty"`
	// The SQL statements for rolling back the data change.
	//
	// >  This parameter is required if you set the **RollbackSqlType*	- parameter to **TEXT**.
	//
	// example:
	//
	// empty
	RollbackSQL *string `json:"RollbackSQL,omitempty" xml:"RollbackSQL,omitempty"`
	// The format of the SQL statements used to roll back the data change. Valid values:
	//
	// 	- **TEXT**
	//
	// 	- **ATTACHMENT**
	//
	// example:
	//
	// TEXT
	RollbackSqlType *string `json:"RollbackSqlType,omitempty" xml:"RollbackSqlType,omitempty"`
}

func (s CreateProcCorrectOrderRequestParam) String() string {
	return dara.Prettify(s)
}

func (s CreateProcCorrectOrderRequestParam) GoString() string {
	return s.String()
}

func (s *CreateProcCorrectOrderRequestParam) GetClassify() *string {
	return s.Classify
}

func (s *CreateProcCorrectOrderRequestParam) GetDbItemList() []*CreateProcCorrectOrderRequestParamDbItemList {
	return s.DbItemList
}

func (s *CreateProcCorrectOrderRequestParam) GetExecMode() *string {
	return s.ExecMode
}

func (s *CreateProcCorrectOrderRequestParam) GetExecSQL() *string {
	return s.ExecSQL
}

func (s *CreateProcCorrectOrderRequestParam) GetRollbackAttachmentName() *string {
	return s.RollbackAttachmentName
}

func (s *CreateProcCorrectOrderRequestParam) GetRollbackSQL() *string {
	return s.RollbackSQL
}

func (s *CreateProcCorrectOrderRequestParam) GetRollbackSqlType() *string {
	return s.RollbackSqlType
}

func (s *CreateProcCorrectOrderRequestParam) SetClassify(v string) *CreateProcCorrectOrderRequestParam {
	s.Classify = &v
	return s
}

func (s *CreateProcCorrectOrderRequestParam) SetDbItemList(v []*CreateProcCorrectOrderRequestParamDbItemList) *CreateProcCorrectOrderRequestParam {
	s.DbItemList = v
	return s
}

func (s *CreateProcCorrectOrderRequestParam) SetExecMode(v string) *CreateProcCorrectOrderRequestParam {
	s.ExecMode = &v
	return s
}

func (s *CreateProcCorrectOrderRequestParam) SetExecSQL(v string) *CreateProcCorrectOrderRequestParam {
	s.ExecSQL = &v
	return s
}

func (s *CreateProcCorrectOrderRequestParam) SetRollbackAttachmentName(v string) *CreateProcCorrectOrderRequestParam {
	s.RollbackAttachmentName = &v
	return s
}

func (s *CreateProcCorrectOrderRequestParam) SetRollbackSQL(v string) *CreateProcCorrectOrderRequestParam {
	s.RollbackSQL = &v
	return s
}

func (s *CreateProcCorrectOrderRequestParam) SetRollbackSqlType(v string) *CreateProcCorrectOrderRequestParam {
	s.RollbackSqlType = &v
	return s
}

func (s *CreateProcCorrectOrderRequestParam) Validate() error {
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

type CreateProcCorrectOrderRequestParamDbItemList struct {
	// The database ID. Databases are divided into physical databases and logical databases.
	//
	// 	- To query the ID of a physical database, call the [ListDatabases](https://help.aliyun.com/document_detail/141873.html) or [SearchDatabase](https://help.aliyun.com/document_detail/141876.html) operation.
	//
	// 	- To query the ID of a logical database, call the [ListLogicDatabases](https://help.aliyun.com/document_detail/141874.html) or [SearchDatabase](https://help.aliyun.com/document_detail/141876.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 19721978
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// Specifies whether the database is a logical database. Valid values:
	//
	// 	- **true**: The database is a logical database.
	//
	// 	- **false**: The database is a physical database.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
}

func (s CreateProcCorrectOrderRequestParamDbItemList) String() string {
	return dara.Prettify(s)
}

func (s CreateProcCorrectOrderRequestParamDbItemList) GoString() string {
	return s.String()
}

func (s *CreateProcCorrectOrderRequestParamDbItemList) GetDbId() *int64 {
	return s.DbId
}

func (s *CreateProcCorrectOrderRequestParamDbItemList) GetLogic() *bool {
	return s.Logic
}

func (s *CreateProcCorrectOrderRequestParamDbItemList) SetDbId(v int64) *CreateProcCorrectOrderRequestParamDbItemList {
	s.DbId = &v
	return s
}

func (s *CreateProcCorrectOrderRequestParamDbItemList) SetLogic(v bool) *CreateProcCorrectOrderRequestParamDbItemList {
	s.Logic = &v
	return s
}

func (s *CreateProcCorrectOrderRequestParamDbItemList) Validate() error {
	return dara.Validate(s)
}
