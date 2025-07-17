// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDBTaskSQLJobLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v int64) *GetDBTaskSQLJobLogRequest
	GetJobId() *int64
	SetTid(v int64) *GetDBTaskSQLJobLogRequest
	GetTid() *int64
}

type GetDBTaskSQLJobLogRequest struct {
	// The ID of the SQL task. You can call the [ListDBTaskSQLJob](https://help.aliyun.com/document_detail/207049.html) operation to query the ID of the SQL task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1276****
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) operation to query the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetDBTaskSQLJobLogRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDBTaskSQLJobLogRequest) GoString() string {
	return s.String()
}

func (s *GetDBTaskSQLJobLogRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *GetDBTaskSQLJobLogRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetDBTaskSQLJobLogRequest) SetJobId(v int64) *GetDBTaskSQLJobLogRequest {
	s.JobId = &v
	return s
}

func (s *GetDBTaskSQLJobLogRequest) SetTid(v int64) *GetDBTaskSQLJobLogRequest {
	s.Tid = &v
	return s
}

func (s *GetDBTaskSQLJobLogRequest) Validate() error {
	return dara.Validate(s)
}
