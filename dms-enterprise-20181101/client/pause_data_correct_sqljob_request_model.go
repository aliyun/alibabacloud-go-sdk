// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseDataCorrectSQLJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v int64) *PauseDataCorrectSQLJobRequest
	GetJobId() *int64
	SetOrderId(v int64) *PauseDataCorrectSQLJobRequest
	GetOrderId() *int64
	SetTid(v int64) *PauseDataCorrectSQLJobRequest
	GetTid() *int64
	SetType(v string) *PauseDataCorrectSQLJobRequest
	GetType() *string
}

type PauseDataCorrectSQLJobRequest struct {
	// The ID of the SQL task. You can call the [GetDataCorrectTaskDetail](https://help.aliyun.com/document_detail/208481.html) or [ListDBTaskSQLJob](https://help.aliyun.com/document_detail/207049.html) operation to obtain the value of this parameter.
	//
	// >  If Type is set to SINGLE, you must pass in the value of JobId to confirm the ID of the SQL task that you want to pause.
	//
	// example:
	//
	// 43253
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the data change ticket. You can call the [ListOrders](https://help.aliyun.com/document_detail/144643.html) operation to query the ID of the data change ticket.
	//
	// This parameter is required.
	//
	// example:
	//
	// 43253
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The tenant ID. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) operation to query the tenant ID.
	//
	// example:
	//
	// 4325
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The type of the pause operation. Valid values:
	//
	// 	- ALL: pauses all SQL tasks.
	//
	// 	- SINGLE: pauses a single SQL task.
	//
	// This parameter is required.
	//
	// example:
	//
	// SINGLE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PauseDataCorrectSQLJobRequest) String() string {
	return dara.Prettify(s)
}

func (s PauseDataCorrectSQLJobRequest) GoString() string {
	return s.String()
}

func (s *PauseDataCorrectSQLJobRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *PauseDataCorrectSQLJobRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *PauseDataCorrectSQLJobRequest) GetTid() *int64 {
	return s.Tid
}

func (s *PauseDataCorrectSQLJobRequest) GetType() *string {
	return s.Type
}

func (s *PauseDataCorrectSQLJobRequest) SetJobId(v int64) *PauseDataCorrectSQLJobRequest {
	s.JobId = &v
	return s
}

func (s *PauseDataCorrectSQLJobRequest) SetOrderId(v int64) *PauseDataCorrectSQLJobRequest {
	s.OrderId = &v
	return s
}

func (s *PauseDataCorrectSQLJobRequest) SetTid(v int64) *PauseDataCorrectSQLJobRequest {
	s.Tid = &v
	return s
}

func (s *PauseDataCorrectSQLJobRequest) SetType(v string) *PauseDataCorrectSQLJobRequest {
	s.Type = &v
	return s
}

func (s *PauseDataCorrectSQLJobRequest) Validate() error {
	return dara.Validate(s)
}
