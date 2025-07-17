// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartDataCorrectSQLJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v int64) *RestartDataCorrectSQLJobRequest
	GetJobId() *int64
	SetOrderId(v int64) *RestartDataCorrectSQLJobRequest
	GetOrderId() *int64
	SetRealLoginUserUid(v string) *RestartDataCorrectSQLJobRequest
	GetRealLoginUserUid() *string
	SetTid(v int64) *RestartDataCorrectSQLJobRequest
	GetTid() *int64
	SetType(v string) *RestartDataCorrectSQLJobRequest
	GetType() *string
}

type RestartDataCorrectSQLJobRequest struct {
	// The ID of the SQL task. You can call the [GetDataCorrectTaskDetail](https://help.aliyun.com/document_detail/208481.html) and [ListDBTaskSQLJob](https://help.aliyun.com/document_detail/207049.html) operations to obtain the value of this parameter.
	//
	// If the Type parameter is set to SINGLE, you must pass the value of the JobId parameter to confirm the ID of the SQL task that you want to rerun.
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
	// 453****
	OrderId          *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RealLoginUserUid *string `json:"RealLoginUserUid,omitempty" xml:"RealLoginUserUid,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) operation to query the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The type of the rerun operation. Valid values:
	//
	// 	- **ALL**: reruns all SQL tasks.
	//
	// 	- **SINGLE**: reruns a single SQL task.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s RestartDataCorrectSQLJobRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartDataCorrectSQLJobRequest) GoString() string {
	return s.String()
}

func (s *RestartDataCorrectSQLJobRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *RestartDataCorrectSQLJobRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *RestartDataCorrectSQLJobRequest) GetRealLoginUserUid() *string {
	return s.RealLoginUserUid
}

func (s *RestartDataCorrectSQLJobRequest) GetTid() *int64 {
	return s.Tid
}

func (s *RestartDataCorrectSQLJobRequest) GetType() *string {
	return s.Type
}

func (s *RestartDataCorrectSQLJobRequest) SetJobId(v int64) *RestartDataCorrectSQLJobRequest {
	s.JobId = &v
	return s
}

func (s *RestartDataCorrectSQLJobRequest) SetOrderId(v int64) *RestartDataCorrectSQLJobRequest {
	s.OrderId = &v
	return s
}

func (s *RestartDataCorrectSQLJobRequest) SetRealLoginUserUid(v string) *RestartDataCorrectSQLJobRequest {
	s.RealLoginUserUid = &v
	return s
}

func (s *RestartDataCorrectSQLJobRequest) SetTid(v int64) *RestartDataCorrectSQLJobRequest {
	s.Tid = &v
	return s
}

func (s *RestartDataCorrectSQLJobRequest) SetType(v string) *RestartDataCorrectSQLJobRequest {
	s.Type = &v
	return s
}

func (s *RestartDataCorrectSQLJobRequest) Validate() error {
	return dara.Validate(s)
}
