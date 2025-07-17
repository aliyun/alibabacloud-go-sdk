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
	// example:
	//
	// order_attachment.txt
	AttachmentKey *string `json:"AttachmentKey,omitempty" xml:"AttachmentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// This parameter is required.
	Param           *CreateProcCorrectOrderRequestParam `json:"Param,omitempty" xml:"Param,omitempty" type:"Struct"`
	RelatedUserList []*int64                            `json:"RelatedUserList,omitempty" xml:"RelatedUserList,omitempty" type:"Repeated"`
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
	return dara.Validate(s)
}

type CreateProcCorrectOrderRequestParam struct {
	// example:
	//
	// test
	Classify *string `json:"Classify,omitempty" xml:"Classify,omitempty"`
	// This parameter is required.
	DbItemList []*CreateProcCorrectOrderRequestParamDbItemList `json:"DbItemList,omitempty" xml:"DbItemList,omitempty" type:"Repeated"`
	ExecMode   *string                                         `json:"ExecMode,omitempty" xml:"ExecMode,omitempty"`
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
	// example:
	//
	// test_rollback.sql
	RollbackAttachmentName *string `json:"RollbackAttachmentName,omitempty" xml:"RollbackAttachmentName,omitempty"`
	// example:
	//
	// empty
	RollbackSQL *string `json:"RollbackSQL,omitempty" xml:"RollbackSQL,omitempty"`
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
	return dara.Validate(s)
}

type CreateProcCorrectOrderRequestParamDbItemList struct {
	// This parameter is required.
	//
	// example:
	//
	// 19721978
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
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
