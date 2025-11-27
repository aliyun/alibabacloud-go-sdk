// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceDelayedReplicationTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyDBInstanceDelayedReplicationTimeResponseBody
	GetDBInstanceId() *string
	SetReadSQLReplicationTime(v string) *ModifyDBInstanceDelayedReplicationTimeResponseBody
	GetReadSQLReplicationTime() *string
	SetRequestId(v string) *ModifyDBInstanceDelayedReplicationTimeResponseBody
	GetRequestId() *string
	SetTaskId(v string) *ModifyDBInstanceDelayedReplicationTimeResponseBody
	GetTaskId() *string
}

type ModifyDBInstanceDelayedReplicationTimeResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// rm-uf6wjk5*****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The replication latency of the read-only instance. Unit: seconds.
	//
	// example:
	//
	// 100
	ReadSQLReplicationTime *string `json:"ReadSQLReplicationTime,omitempty" xml:"ReadSQLReplicationTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EFB6083A-7699-489B-8278-C0CB4793A96E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 1715482.0
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyDBInstanceDelayedReplicationTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceDelayedReplicationTimeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceDelayedReplicationTimeResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceDelayedReplicationTimeResponseBody) GetReadSQLReplicationTime() *string {
	return s.ReadSQLReplicationTime
}

func (s *ModifyDBInstanceDelayedReplicationTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBInstanceDelayedReplicationTimeResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *ModifyDBInstanceDelayedReplicationTimeResponseBody) SetDBInstanceId(v string) *ModifyDBInstanceDelayedReplicationTimeResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceDelayedReplicationTimeResponseBody) SetReadSQLReplicationTime(v string) *ModifyDBInstanceDelayedReplicationTimeResponseBody {
	s.ReadSQLReplicationTime = &v
	return s
}

func (s *ModifyDBInstanceDelayedReplicationTimeResponseBody) SetRequestId(v string) *ModifyDBInstanceDelayedReplicationTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstanceDelayedReplicationTimeResponseBody) SetTaskId(v string) *ModifyDBInstanceDelayedReplicationTimeResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModifyDBInstanceDelayedReplicationTimeResponseBody) Validate() error {
	return dara.Validate(s)
}
