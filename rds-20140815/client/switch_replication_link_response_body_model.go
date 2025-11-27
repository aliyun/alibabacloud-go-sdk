// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchReplicationLinkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *SwitchReplicationLinkResponseBody
	GetDBInstanceId() *string
	SetRequestId(v string) *SwitchReplicationLinkResponseBody
	GetRequestId() *string
	SetTaskId(v int64) *SwitchReplicationLinkResponseBody
	GetTaskId() *int64
	SetTaskName(v string) *SwitchReplicationLinkResponseBody
	GetTaskName() *string
}

type SwitchReplicationLinkResponseBody struct {
	// The ID of the DR instance.
	//
	// example:
	//
	// 135****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2F2DD69B-90AF-1E72-923C-87575658A9D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 159****
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task name.
	//
	// example:
	//
	// zbtest
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s SwitchReplicationLinkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SwitchReplicationLinkResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchReplicationLinkResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *SwitchReplicationLinkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SwitchReplicationLinkResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *SwitchReplicationLinkResponseBody) GetTaskName() *string {
	return s.TaskName
}

func (s *SwitchReplicationLinkResponseBody) SetDBInstanceId(v string) *SwitchReplicationLinkResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *SwitchReplicationLinkResponseBody) SetRequestId(v string) *SwitchReplicationLinkResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchReplicationLinkResponseBody) SetTaskId(v int64) *SwitchReplicationLinkResponseBody {
	s.TaskId = &v
	return s
}

func (s *SwitchReplicationLinkResponseBody) SetTaskName(v string) *SwitchReplicationLinkResponseBody {
	s.TaskName = &v
	return s
}

func (s *SwitchReplicationLinkResponseBody) Validate() error {
	return dara.Validate(s)
}
