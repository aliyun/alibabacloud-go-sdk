// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFreeLockCorrectOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachmentKey(v string) *CreateFreeLockCorrectOrderRequest
	GetAttachmentKey() *string
	SetComment(v string) *CreateFreeLockCorrectOrderRequest
	GetComment() *string
	SetParam(v *CreateFreeLockCorrectOrderRequestParam) *CreateFreeLockCorrectOrderRequest
	GetParam() *CreateFreeLockCorrectOrderRequestParam
	SetRealLoginUserUid(v string) *CreateFreeLockCorrectOrderRequest
	GetRealLoginUserUid() *string
	SetRelatedUserList(v []*int64) *CreateFreeLockCorrectOrderRequest
	GetRelatedUserList() []*int64
	SetTid(v int64) *CreateFreeLockCorrectOrderRequest
	GetTid() *int64
}

type CreateFreeLockCorrectOrderRequest struct {
	// The key of the attachment that provides more instructions for the ticket. You can call the [GetUserUploadFileJob](https://help.aliyun.com/document_detail/206069.html) operation to obtain the attachment key from the value of the AttachmentKey parameter.
	//
	// example:
	//
	// order_attachment.txt
	AttachmentKey *string `json:"AttachmentKey,omitempty" xml:"AttachmentKey,omitempty"`
	// The purpose or objective of the data change. This parameter is used to help reduce unnecessary communication.
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
	Param            *CreateFreeLockCorrectOrderRequestParam `json:"Param,omitempty" xml:"Param,omitempty" type:"Struct"`
	RealLoginUserUid *string                                 `json:"RealLoginUserUid,omitempty" xml:"RealLoginUserUid,omitempty"`
	// The stakeholders of the data change. All stakeholders can view the ticket details and assist in the approval process. Irrelevant users other than DMS administrators and database administrators (DBAs) are not allowed to view the ticket details.
	RelatedUserList []*int64 `json:"RelatedUserList,omitempty" xml:"RelatedUserList,omitempty" type:"Repeated"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateFreeLockCorrectOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFreeLockCorrectOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateFreeLockCorrectOrderRequest) GetAttachmentKey() *string {
	return s.AttachmentKey
}

func (s *CreateFreeLockCorrectOrderRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateFreeLockCorrectOrderRequest) GetParam() *CreateFreeLockCorrectOrderRequestParam {
	return s.Param
}

func (s *CreateFreeLockCorrectOrderRequest) GetRealLoginUserUid() *string {
	return s.RealLoginUserUid
}

func (s *CreateFreeLockCorrectOrderRequest) GetRelatedUserList() []*int64 {
	return s.RelatedUserList
}

func (s *CreateFreeLockCorrectOrderRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateFreeLockCorrectOrderRequest) SetAttachmentKey(v string) *CreateFreeLockCorrectOrderRequest {
	s.AttachmentKey = &v
	return s
}

func (s *CreateFreeLockCorrectOrderRequest) SetComment(v string) *CreateFreeLockCorrectOrderRequest {
	s.Comment = &v
	return s
}

func (s *CreateFreeLockCorrectOrderRequest) SetParam(v *CreateFreeLockCorrectOrderRequestParam) *CreateFreeLockCorrectOrderRequest {
	s.Param = v
	return s
}

func (s *CreateFreeLockCorrectOrderRequest) SetRealLoginUserUid(v string) *CreateFreeLockCorrectOrderRequest {
	s.RealLoginUserUid = &v
	return s
}

func (s *CreateFreeLockCorrectOrderRequest) SetRelatedUserList(v []*int64) *CreateFreeLockCorrectOrderRequest {
	s.RelatedUserList = v
	return s
}

func (s *CreateFreeLockCorrectOrderRequest) SetTid(v int64) *CreateFreeLockCorrectOrderRequest {
	s.Tid = &v
	return s
}

func (s *CreateFreeLockCorrectOrderRequest) Validate() error {
	return dara.Validate(s)
}

type CreateFreeLockCorrectOrderRequestParam struct {
	// The key of the attachment that contains the SQL statements used to change data. This parameter is not supported.
	//
	// example:
	//
	// test.sql
	AttachmentName *string `json:"AttachmentName,omitempty" xml:"AttachmentName,omitempty"`
	// The reason for the data change.
	//
	// example:
	//
	// test
	Classify *string `json:"Classify,omitempty" xml:"Classify,omitempty"`
	// The databases in which you want to change data.
	//
	// This parameter is required.
	DbItemList []*CreateFreeLockCorrectOrderRequestParamDbItemList `json:"DbItemList,omitempty" xml:"DbItemList,omitempty" type:"Repeated"`
	// The execution mode of the ticket after the ticket is approved. Valid values:
	//
	// 	- **COMMITOR**: The data change is performed by the user who submits the ticket.
	//
	// 	- **AUTO**: The data change is automatically performed after the ticket is approved.
	//
	// 	- **LAST_AUDITOR**: The data change is performed by the last approver of the ticket.
	//
	// example:
	//
	// COMMITOR
	ExecMode *string `json:"ExecMode,omitempty" xml:"ExecMode,omitempty"`
	// The SQL statements that you want to execute to change data.
	//
	// This parameter is required.
	//
	// example:
	//
	// delete from base_user where 1 = 1;
	ExecSQL *string `json:"ExecSQL,omitempty" xml:"ExecSQL,omitempty"`
	// The key of the attachment that contains the SQL statements used to roll back the data change.
	//
	// example:
	//
	// test_rollback.sql
	RollbackAttachmentName *string `json:"RollbackAttachmentName,omitempty" xml:"RollbackAttachmentName,omitempty"`
	// The SQL statements used to roll back the data change.
	//
	// example:
	//
	// empty
	RollbackSQL *string `json:"RollbackSQL,omitempty" xml:"RollbackSQL,omitempty"`
	// The format of the SQL statements used to roll back the data change. Valid values:
	//
	// 	- **TEXT**: text
	//
	// 	- **ATTACHMENT**: attachment. This value is not supported.
	//
	// example:
	//
	// TEXT
	RollbackSqlType *string `json:"RollbackSqlType,omitempty" xml:"RollbackSqlType,omitempty"`
	// The format of the SQL statements used to change data. Valid values:
	//
	// 	- **TEXT**: text
	//
	// 	- **ATTACHMENT**: attachment. This value is not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// TEXT
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
}

func (s CreateFreeLockCorrectOrderRequestParam) String() string {
	return dara.Prettify(s)
}

func (s CreateFreeLockCorrectOrderRequestParam) GoString() string {
	return s.String()
}

func (s *CreateFreeLockCorrectOrderRequestParam) GetAttachmentName() *string {
	return s.AttachmentName
}

func (s *CreateFreeLockCorrectOrderRequestParam) GetClassify() *string {
	return s.Classify
}

func (s *CreateFreeLockCorrectOrderRequestParam) GetDbItemList() []*CreateFreeLockCorrectOrderRequestParamDbItemList {
	return s.DbItemList
}

func (s *CreateFreeLockCorrectOrderRequestParam) GetExecMode() *string {
	return s.ExecMode
}

func (s *CreateFreeLockCorrectOrderRequestParam) GetExecSQL() *string {
	return s.ExecSQL
}

func (s *CreateFreeLockCorrectOrderRequestParam) GetRollbackAttachmentName() *string {
	return s.RollbackAttachmentName
}

func (s *CreateFreeLockCorrectOrderRequestParam) GetRollbackSQL() *string {
	return s.RollbackSQL
}

func (s *CreateFreeLockCorrectOrderRequestParam) GetRollbackSqlType() *string {
	return s.RollbackSqlType
}

func (s *CreateFreeLockCorrectOrderRequestParam) GetSqlType() *string {
	return s.SqlType
}

func (s *CreateFreeLockCorrectOrderRequestParam) SetAttachmentName(v string) *CreateFreeLockCorrectOrderRequestParam {
	s.AttachmentName = &v
	return s
}

func (s *CreateFreeLockCorrectOrderRequestParam) SetClassify(v string) *CreateFreeLockCorrectOrderRequestParam {
	s.Classify = &v
	return s
}

func (s *CreateFreeLockCorrectOrderRequestParam) SetDbItemList(v []*CreateFreeLockCorrectOrderRequestParamDbItemList) *CreateFreeLockCorrectOrderRequestParam {
	s.DbItemList = v
	return s
}

func (s *CreateFreeLockCorrectOrderRequestParam) SetExecMode(v string) *CreateFreeLockCorrectOrderRequestParam {
	s.ExecMode = &v
	return s
}

func (s *CreateFreeLockCorrectOrderRequestParam) SetExecSQL(v string) *CreateFreeLockCorrectOrderRequestParam {
	s.ExecSQL = &v
	return s
}

func (s *CreateFreeLockCorrectOrderRequestParam) SetRollbackAttachmentName(v string) *CreateFreeLockCorrectOrderRequestParam {
	s.RollbackAttachmentName = &v
	return s
}

func (s *CreateFreeLockCorrectOrderRequestParam) SetRollbackSQL(v string) *CreateFreeLockCorrectOrderRequestParam {
	s.RollbackSQL = &v
	return s
}

func (s *CreateFreeLockCorrectOrderRequestParam) SetRollbackSqlType(v string) *CreateFreeLockCorrectOrderRequestParam {
	s.RollbackSqlType = &v
	return s
}

func (s *CreateFreeLockCorrectOrderRequestParam) SetSqlType(v string) *CreateFreeLockCorrectOrderRequestParam {
	s.SqlType = &v
	return s
}

func (s *CreateFreeLockCorrectOrderRequestParam) Validate() error {
	return dara.Validate(s)
}

type CreateFreeLockCorrectOrderRequestParamDbItemList struct {
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
	// This parameter is required.
	//
	// example:
	//
	// false
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
}

func (s CreateFreeLockCorrectOrderRequestParamDbItemList) String() string {
	return dara.Prettify(s)
}

func (s CreateFreeLockCorrectOrderRequestParamDbItemList) GoString() string {
	return s.String()
}

func (s *CreateFreeLockCorrectOrderRequestParamDbItemList) GetDbId() *int64 {
	return s.DbId
}

func (s *CreateFreeLockCorrectOrderRequestParamDbItemList) GetLogic() *bool {
	return s.Logic
}

func (s *CreateFreeLockCorrectOrderRequestParamDbItemList) SetDbId(v int64) *CreateFreeLockCorrectOrderRequestParamDbItemList {
	s.DbId = &v
	return s
}

func (s *CreateFreeLockCorrectOrderRequestParamDbItemList) SetLogic(v bool) *CreateFreeLockCorrectOrderRequestParamDbItemList {
	s.Logic = &v
	return s
}

func (s *CreateFreeLockCorrectOrderRequestParamDbItemList) Validate() error {
	return dara.Validate(s)
}
