// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDataCorrectExecSQLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExecSQL(v string) *ModifyDataCorrectExecSQLRequest
	GetExecSQL() *string
	SetOrderId(v int64) *ModifyDataCorrectExecSQLRequest
	GetOrderId() *int64
	SetRealLoginUserUid(v string) *ModifyDataCorrectExecSQLRequest
	GetRealLoginUserUid() *string
	SetTid(v int64) *ModifyDataCorrectExecSQLRequest
	GetTid() *int64
}

type ModifyDataCorrectExecSQLRequest struct {
	// The new SQL script.
	//
	// This parameter is required.
	//
	// example:
	//
	// update tb set id = 1 where id = 1;
	ExecSQL *string `json:"ExecSQL,omitempty" xml:"ExecSQL,omitempty"`
	// The ID of the data change ticket.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4328****
	OrderId          *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RealLoginUserUid *string `json:"RealLoginUserUid,omitempty" xml:"RealLoginUserUid,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) operation to query the tenant ID.
	//
	// example:
	//
	// 4****
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ModifyDataCorrectExecSQLRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDataCorrectExecSQLRequest) GoString() string {
	return s.String()
}

func (s *ModifyDataCorrectExecSQLRequest) GetExecSQL() *string {
	return s.ExecSQL
}

func (s *ModifyDataCorrectExecSQLRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ModifyDataCorrectExecSQLRequest) GetRealLoginUserUid() *string {
	return s.RealLoginUserUid
}

func (s *ModifyDataCorrectExecSQLRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ModifyDataCorrectExecSQLRequest) SetExecSQL(v string) *ModifyDataCorrectExecSQLRequest {
	s.ExecSQL = &v
	return s
}

func (s *ModifyDataCorrectExecSQLRequest) SetOrderId(v int64) *ModifyDataCorrectExecSQLRequest {
	s.OrderId = &v
	return s
}

func (s *ModifyDataCorrectExecSQLRequest) SetRealLoginUserUid(v string) *ModifyDataCorrectExecSQLRequest {
	s.RealLoginUserUid = &v
	return s
}

func (s *ModifyDataCorrectExecSQLRequest) SetTid(v int64) *ModifyDataCorrectExecSQLRequest {
	s.Tid = &v
	return s
}

func (s *ModifyDataCorrectExecSQLRequest) Validate() error {
	return dara.Validate(s)
}
