// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataCorrectOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachmentKey(v string) *CreateDataCorrectOrderRequest
	GetAttachmentKey() *string
	SetComment(v string) *CreateDataCorrectOrderRequest
	GetComment() *string
	SetParam(v *CreateDataCorrectOrderRequestParam) *CreateDataCorrectOrderRequest
	GetParam() *CreateDataCorrectOrderRequestParam
	SetRealLoginUserUid(v string) *CreateDataCorrectOrderRequest
	GetRealLoginUserUid() *string
	SetRelatedUserList(v []*int64) *CreateDataCorrectOrderRequest
	GetRelatedUserList() []*int64
	SetTid(v int64) *CreateDataCorrectOrderRequest
	GetTid() *int64
}

type CreateDataCorrectOrderRequest struct {
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
	Param *CreateDataCorrectOrderRequestParam `json:"Param,omitempty" xml:"Param,omitempty" type:"Struct"`
	// The ID of the Alibaba Cloud account that is used to call the API operation.
	//
	// example:
	//
	// 21400447956867****
	RealLoginUserUid *string `json:"RealLoginUserUid,omitempty" xml:"RealLoginUserUid,omitempty"`
	// The stakeholders of the data change. All stakeholders can view the ticket details and assist in the approval process. Irrelevant users other than DMS administrators and database administrators (DBAs) are not allowed to view the ticket details.
	RelatedUserList []*int64 `json:"RelatedUserList,omitempty" xml:"RelatedUserList,omitempty" type:"Repeated"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateDataCorrectOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataCorrectOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateDataCorrectOrderRequest) GetAttachmentKey() *string {
	return s.AttachmentKey
}

func (s *CreateDataCorrectOrderRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateDataCorrectOrderRequest) GetParam() *CreateDataCorrectOrderRequestParam {
	return s.Param
}

func (s *CreateDataCorrectOrderRequest) GetRealLoginUserUid() *string {
	return s.RealLoginUserUid
}

func (s *CreateDataCorrectOrderRequest) GetRelatedUserList() []*int64 {
	return s.RelatedUserList
}

func (s *CreateDataCorrectOrderRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateDataCorrectOrderRequest) SetAttachmentKey(v string) *CreateDataCorrectOrderRequest {
	s.AttachmentKey = &v
	return s
}

func (s *CreateDataCorrectOrderRequest) SetComment(v string) *CreateDataCorrectOrderRequest {
	s.Comment = &v
	return s
}

func (s *CreateDataCorrectOrderRequest) SetParam(v *CreateDataCorrectOrderRequestParam) *CreateDataCorrectOrderRequest {
	s.Param = v
	return s
}

func (s *CreateDataCorrectOrderRequest) SetRealLoginUserUid(v string) *CreateDataCorrectOrderRequest {
	s.RealLoginUserUid = &v
	return s
}

func (s *CreateDataCorrectOrderRequest) SetRelatedUserList(v []*int64) *CreateDataCorrectOrderRequest {
	s.RelatedUserList = v
	return s
}

func (s *CreateDataCorrectOrderRequest) SetTid(v int64) *CreateDataCorrectOrderRequest {
	s.Tid = &v
	return s
}

func (s *CreateDataCorrectOrderRequest) Validate() error {
	return dara.Validate(s)
}

type CreateDataCorrectOrderRequestParam struct {
	// The key of the attachment that contains the SQL statements used to change data. You can call the [GetUserUploadFileJob](https://help.aliyun.com/document_detail/206069.html) operation to obtain the attachment key from the value of AttachmentKey.
	//
	// >  This parameter is required if you set **SqlType*	- to **ATTACHMENT**.
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
	// The databases whose data you want to change.
	//
	// This parameter is required.
	DbItemList []*CreateDataCorrectOrderRequestParamDbItemList `json:"DbItemList,omitempty" xml:"DbItemList,omitempty" type:"Repeated"`
	// The estimated number of data rows that may be affected by the data change.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	EstimateAffectRows *int64 `json:"EstimateAffectRows,omitempty" xml:"EstimateAffectRows,omitempty"`
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
	// >
	//
	// 	- This parameter is required if you set **SqlType*	- to **TEXT**.
	//
	// 	- The size of the SQL statement cannot exceed 15 MB.
	//
	// example:
	//
	// update base_user set id = 1 where id  = 1;
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
	// >  This parameter is required if you set **RollbackSqlType*	- to **TEXT**.
	//
	// example:
	//
	// update base_user set id = 1 where id  = 1;
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
	// The format of the SQL statements used to change data. Valid values:
	//
	// 	- **TEXT**
	//
	// 	- **ATTACHMENT**
	//
	// This parameter is required.
	//
	// example:
	//
	// TEXT
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
}

func (s CreateDataCorrectOrderRequestParam) String() string {
	return dara.Prettify(s)
}

func (s CreateDataCorrectOrderRequestParam) GoString() string {
	return s.String()
}

func (s *CreateDataCorrectOrderRequestParam) GetAttachmentName() *string {
	return s.AttachmentName
}

func (s *CreateDataCorrectOrderRequestParam) GetClassify() *string {
	return s.Classify
}

func (s *CreateDataCorrectOrderRequestParam) GetDbItemList() []*CreateDataCorrectOrderRequestParamDbItemList {
	return s.DbItemList
}

func (s *CreateDataCorrectOrderRequestParam) GetEstimateAffectRows() *int64 {
	return s.EstimateAffectRows
}

func (s *CreateDataCorrectOrderRequestParam) GetExecMode() *string {
	return s.ExecMode
}

func (s *CreateDataCorrectOrderRequestParam) GetExecSQL() *string {
	return s.ExecSQL
}

func (s *CreateDataCorrectOrderRequestParam) GetRollbackAttachmentName() *string {
	return s.RollbackAttachmentName
}

func (s *CreateDataCorrectOrderRequestParam) GetRollbackSQL() *string {
	return s.RollbackSQL
}

func (s *CreateDataCorrectOrderRequestParam) GetRollbackSqlType() *string {
	return s.RollbackSqlType
}

func (s *CreateDataCorrectOrderRequestParam) GetSqlType() *string {
	return s.SqlType
}

func (s *CreateDataCorrectOrderRequestParam) SetAttachmentName(v string) *CreateDataCorrectOrderRequestParam {
	s.AttachmentName = &v
	return s
}

func (s *CreateDataCorrectOrderRequestParam) SetClassify(v string) *CreateDataCorrectOrderRequestParam {
	s.Classify = &v
	return s
}

func (s *CreateDataCorrectOrderRequestParam) SetDbItemList(v []*CreateDataCorrectOrderRequestParamDbItemList) *CreateDataCorrectOrderRequestParam {
	s.DbItemList = v
	return s
}

func (s *CreateDataCorrectOrderRequestParam) SetEstimateAffectRows(v int64) *CreateDataCorrectOrderRequestParam {
	s.EstimateAffectRows = &v
	return s
}

func (s *CreateDataCorrectOrderRequestParam) SetExecMode(v string) *CreateDataCorrectOrderRequestParam {
	s.ExecMode = &v
	return s
}

func (s *CreateDataCorrectOrderRequestParam) SetExecSQL(v string) *CreateDataCorrectOrderRequestParam {
	s.ExecSQL = &v
	return s
}

func (s *CreateDataCorrectOrderRequestParam) SetRollbackAttachmentName(v string) *CreateDataCorrectOrderRequestParam {
	s.RollbackAttachmentName = &v
	return s
}

func (s *CreateDataCorrectOrderRequestParam) SetRollbackSQL(v string) *CreateDataCorrectOrderRequestParam {
	s.RollbackSQL = &v
	return s
}

func (s *CreateDataCorrectOrderRequestParam) SetRollbackSqlType(v string) *CreateDataCorrectOrderRequestParam {
	s.RollbackSqlType = &v
	return s
}

func (s *CreateDataCorrectOrderRequestParam) SetSqlType(v string) *CreateDataCorrectOrderRequestParam {
	s.SqlType = &v
	return s
}

func (s *CreateDataCorrectOrderRequestParam) Validate() error {
	return dara.Validate(s)
}

type CreateDataCorrectOrderRequestParamDbItemList struct {
	// The database ID. The database can be a physical database or a logical database.
	//
	// 	- To query the ID of a physical database, call the [ListDatabases](https://help.aliyun.com/document_detail/141873.html) or [SearchDatabase](https://help.aliyun.com/document_detail/141876.html) operation.
	//
	// 	- To query the ID of a logical database, call the [ListLogicDatabases](https://help.aliyun.com/document_detail/141874.html) or [SearchDatabase](https://help.aliyun.com/document_detail/141876.html) operation.
	//
	// example:
	//
	// 1860****
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The instance ID. You can call the ListInstances or GetInstance operation to query the instance ID.
	//
	// >
	//
	// > The instance change feature is supported only by ApsaraDB RDS for MySQL instances, PolarDB for MySQL clusters, and AnalyticDB for MySQL clusters.
	//
	// example:
	//
	// 1860****
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether the database is a logical database. Valid values:
	//
	// 	- **true**: The database is a logical database.
	//
	// 	- **false**: The database is a physical database.
	//
	// example:
	//
	// false
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
}

func (s CreateDataCorrectOrderRequestParamDbItemList) String() string {
	return dara.Prettify(s)
}

func (s CreateDataCorrectOrderRequestParamDbItemList) GoString() string {
	return s.String()
}

func (s *CreateDataCorrectOrderRequestParamDbItemList) GetDbId() *int64 {
	return s.DbId
}

func (s *CreateDataCorrectOrderRequestParamDbItemList) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *CreateDataCorrectOrderRequestParamDbItemList) GetLogic() *bool {
	return s.Logic
}

func (s *CreateDataCorrectOrderRequestParamDbItemList) SetDbId(v int64) *CreateDataCorrectOrderRequestParamDbItemList {
	s.DbId = &v
	return s
}

func (s *CreateDataCorrectOrderRequestParamDbItemList) SetInstanceId(v int64) *CreateDataCorrectOrderRequestParamDbItemList {
	s.InstanceId = &v
	return s
}

func (s *CreateDataCorrectOrderRequestParamDbItemList) SetLogic(v bool) *CreateDataCorrectOrderRequestParamDbItemList {
	s.Logic = &v
	return s
}

func (s *CreateDataCorrectOrderRequestParamDbItemList) Validate() error {
	return dara.Validate(s)
}
