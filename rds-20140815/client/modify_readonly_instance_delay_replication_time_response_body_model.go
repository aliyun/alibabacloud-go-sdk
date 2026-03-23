// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyReadonlyInstanceDelayReplicationTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyReadonlyInstanceDelayReplicationTimeResponseBody
	GetDBInstanceId() *string
	SetReadSQLReplicationTime(v string) *ModifyReadonlyInstanceDelayReplicationTimeResponseBody
	GetReadSQLReplicationTime() *string
	SetRequestId(v string) *ModifyReadonlyInstanceDelayReplicationTimeResponseBody
	GetRequestId() *string
	SetTaskId(v string) *ModifyReadonlyInstanceDelayReplicationTimeResponseBody
	GetTaskId() *string
}

type ModifyReadonlyInstanceDelayReplicationTimeResponseBody struct {
	DBInstanceId           *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ReadSQLReplicationTime *string `json:"ReadSQLReplicationTime,omitempty" xml:"ReadSQLReplicationTime,omitempty"`
	RequestId              *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId                 *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyReadonlyInstanceDelayReplicationTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyReadonlyInstanceDelayReplicationTimeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyReadonlyInstanceDelayReplicationTimeResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyReadonlyInstanceDelayReplicationTimeResponseBody) GetReadSQLReplicationTime() *string {
	return s.ReadSQLReplicationTime
}

func (s *ModifyReadonlyInstanceDelayReplicationTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyReadonlyInstanceDelayReplicationTimeResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *ModifyReadonlyInstanceDelayReplicationTimeResponseBody) SetDBInstanceId(v string) *ModifyReadonlyInstanceDelayReplicationTimeResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyReadonlyInstanceDelayReplicationTimeResponseBody) SetReadSQLReplicationTime(v string) *ModifyReadonlyInstanceDelayReplicationTimeResponseBody {
	s.ReadSQLReplicationTime = &v
	return s
}

func (s *ModifyReadonlyInstanceDelayReplicationTimeResponseBody) SetRequestId(v string) *ModifyReadonlyInstanceDelayReplicationTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyReadonlyInstanceDelayReplicationTimeResponseBody) SetTaskId(v string) *ModifyReadonlyInstanceDelayReplicationTimeResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModifyReadonlyInstanceDelayReplicationTimeResponseBody) Validate() error {
	return dara.Validate(s)
}
