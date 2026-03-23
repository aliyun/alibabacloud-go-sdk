// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteReplicationLinkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DeleteReplicationLinkResponseBody
	GetDBInstanceId() *string
	SetRequestId(v string) *DeleteReplicationLinkResponseBody
	GetRequestId() *string
	SetTaskId(v int64) *DeleteReplicationLinkResponseBody
	GetTaskId() *int64
	SetTaskName(v string) *DeleteReplicationLinkResponseBody
	GetTaskName() *string
}

type DeleteReplicationLinkResponseBody struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId       *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName     *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s DeleteReplicationLinkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteReplicationLinkResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteReplicationLinkResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteReplicationLinkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteReplicationLinkResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *DeleteReplicationLinkResponseBody) GetTaskName() *string {
	return s.TaskName
}

func (s *DeleteReplicationLinkResponseBody) SetDBInstanceId(v string) *DeleteReplicationLinkResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteReplicationLinkResponseBody) SetRequestId(v string) *DeleteReplicationLinkResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteReplicationLinkResponseBody) SetTaskId(v int64) *DeleteReplicationLinkResponseBody {
	s.TaskId = &v
	return s
}

func (s *DeleteReplicationLinkResponseBody) SetTaskName(v string) *DeleteReplicationLinkResponseBody {
	s.TaskName = &v
	return s
}

func (s *DeleteReplicationLinkResponseBody) Validate() error {
	return dara.Validate(s)
}
