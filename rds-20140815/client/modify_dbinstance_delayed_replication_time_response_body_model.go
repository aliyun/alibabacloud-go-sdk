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
	DBInstanceId           *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ReadSQLReplicationTime *string `json:"ReadSQLReplicationTime,omitempty" xml:"ReadSQLReplicationTime,omitempty"`
	RequestId              *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId                 *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
