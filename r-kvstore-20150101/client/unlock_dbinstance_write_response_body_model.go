// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnlockDBInstanceWriteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *UnlockDBInstanceWriteResponseBody
	GetDBInstanceName() *string
	SetRequestId(v string) *UnlockDBInstanceWriteResponseBody
	GetRequestId() *string
	SetTaskId(v int64) *UnlockDBInstanceWriteResponseBody
	GetTaskId() *int64
}

type UnlockDBInstanceWriteResponseBody struct {
	// The name of the instance.
	//
	// example:
	//
	// r-2ev03avw0r0552***
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5D622714-AEDD-4609-9167-F5DDD3D1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 10****
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UnlockDBInstanceWriteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnlockDBInstanceWriteResponseBody) GoString() string {
	return s.String()
}

func (s *UnlockDBInstanceWriteResponseBody) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *UnlockDBInstanceWriteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnlockDBInstanceWriteResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *UnlockDBInstanceWriteResponseBody) SetDBInstanceName(v string) *UnlockDBInstanceWriteResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *UnlockDBInstanceWriteResponseBody) SetRequestId(v string) *UnlockDBInstanceWriteResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnlockDBInstanceWriteResponseBody) SetTaskId(v int64) *UnlockDBInstanceWriteResponseBody {
	s.TaskId = &v
	return s
}

func (s *UnlockDBInstanceWriteResponseBody) Validate() error {
	return dara.Validate(s)
}
