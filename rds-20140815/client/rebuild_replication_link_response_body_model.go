// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebuildReplicationLinkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *RebuildReplicationLinkResponseBody
	GetDBInstanceId() *string
	SetRequestId(v string) *RebuildReplicationLinkResponseBody
	GetRequestId() *string
	SetTaskId(v int64) *RebuildReplicationLinkResponseBody
	GetTaskId() *int64
	SetTaskName(v string) *RebuildReplicationLinkResponseBody
	GetTaskName() *string
}

type RebuildReplicationLinkResponseBody struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId       *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName     *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s RebuildReplicationLinkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RebuildReplicationLinkResponseBody) GoString() string {
	return s.String()
}

func (s *RebuildReplicationLinkResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *RebuildReplicationLinkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RebuildReplicationLinkResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *RebuildReplicationLinkResponseBody) GetTaskName() *string {
	return s.TaskName
}

func (s *RebuildReplicationLinkResponseBody) SetDBInstanceId(v string) *RebuildReplicationLinkResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *RebuildReplicationLinkResponseBody) SetRequestId(v string) *RebuildReplicationLinkResponseBody {
	s.RequestId = &v
	return s
}

func (s *RebuildReplicationLinkResponseBody) SetTaskId(v int64) *RebuildReplicationLinkResponseBody {
	s.TaskId = &v
	return s
}

func (s *RebuildReplicationLinkResponseBody) SetTaskName(v string) *RebuildReplicationLinkResponseBody {
	s.TaskName = &v
	return s
}

func (s *RebuildReplicationLinkResponseBody) Validate() error {
	return dara.Validate(s)
}
