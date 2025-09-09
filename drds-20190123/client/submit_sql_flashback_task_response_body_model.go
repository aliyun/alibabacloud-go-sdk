// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSqlFlashbackTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SubmitSqlFlashbackTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitSqlFlashbackTaskResponseBody
	GetSuccess() *bool
	SetTaskId(v int64) *SubmitSqlFlashbackTaskResponseBody
	GetTaskId() *int64
}

type SubmitSqlFlashbackTaskResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// DSSDF-SEWE-*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the database creation failure records were removed from the DRDS instance.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the replication task.
	//
	// example:
	//
	// 1111
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SubmitSqlFlashbackTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitSqlFlashbackTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSqlFlashbackTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitSqlFlashbackTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitSqlFlashbackTaskResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *SubmitSqlFlashbackTaskResponseBody) SetRequestId(v string) *SubmitSqlFlashbackTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitSqlFlashbackTaskResponseBody) SetSuccess(v bool) *SubmitSqlFlashbackTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitSqlFlashbackTaskResponseBody) SetTaskId(v int64) *SubmitSqlFlashbackTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *SubmitSqlFlashbackTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
