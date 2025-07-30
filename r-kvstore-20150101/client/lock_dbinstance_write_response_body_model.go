// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLockDBInstanceWriteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *LockDBInstanceWriteResponseBody
	GetDBInstanceName() *string
	SetLockReason(v string) *LockDBInstanceWriteResponseBody
	GetLockReason() *string
	SetRequestId(v string) *LockDBInstanceWriteResponseBody
	GetRequestId() *string
	SetTaskId(v int64) *LockDBInstanceWriteResponseBody
	GetTaskId() *int64
}

type LockDBInstanceWriteResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// r-2ev03avw0r0552***
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The reason why write operations on the instance are locked.
	//
	// example:
	//
	// lock reason
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2BE6E619-A657-42E3-AD2D-18F8428A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 21986****
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s LockDBInstanceWriteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LockDBInstanceWriteResponseBody) GoString() string {
	return s.String()
}

func (s *LockDBInstanceWriteResponseBody) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *LockDBInstanceWriteResponseBody) GetLockReason() *string {
	return s.LockReason
}

func (s *LockDBInstanceWriteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LockDBInstanceWriteResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *LockDBInstanceWriteResponseBody) SetDBInstanceName(v string) *LockDBInstanceWriteResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *LockDBInstanceWriteResponseBody) SetLockReason(v string) *LockDBInstanceWriteResponseBody {
	s.LockReason = &v
	return s
}

func (s *LockDBInstanceWriteResponseBody) SetRequestId(v string) *LockDBInstanceWriteResponseBody {
	s.RequestId = &v
	return s
}

func (s *LockDBInstanceWriteResponseBody) SetTaskId(v int64) *LockDBInstanceWriteResponseBody {
	s.TaskId = &v
	return s
}

func (s *LockDBInstanceWriteResponseBody) Validate() error {
	return dara.Validate(s)
}
