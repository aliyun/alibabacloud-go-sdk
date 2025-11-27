// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReplicationLinkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *CreateReplicationLinkResponseBody
	GetDBInstanceId() *string
	SetRequestId(v string) *CreateReplicationLinkResponseBody
	GetRequestId() *string
	SetTaskId(v int64) *CreateReplicationLinkResponseBody
	GetTaskId() *int64
	SetTaskName(v string) *CreateReplicationLinkResponseBody
	GetTaskName() *string
}

type CreateReplicationLinkResponseBody struct {
	// The ID of the DR instance.
	//
	// example:
	//
	// pgm-****.pg.rds.aliyuncs.com
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9F8C06AD-3F37-57A0-ABBF-ABD7824F55CE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 564532302
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task name.
	//
	// example:
	//
	// test01
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s CreateReplicationLinkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateReplicationLinkResponseBody) GoString() string {
	return s.String()
}

func (s *CreateReplicationLinkResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateReplicationLinkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateReplicationLinkResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *CreateReplicationLinkResponseBody) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateReplicationLinkResponseBody) SetDBInstanceId(v string) *CreateReplicationLinkResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *CreateReplicationLinkResponseBody) SetRequestId(v string) *CreateReplicationLinkResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateReplicationLinkResponseBody) SetTaskId(v int64) *CreateReplicationLinkResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateReplicationLinkResponseBody) SetTaskName(v string) *CreateReplicationLinkResponseBody {
	s.TaskName = &v
	return s
}

func (s *CreateReplicationLinkResponseBody) Validate() error {
	return dara.Validate(s)
}
