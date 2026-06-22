// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPreCheckDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceUuid(v string) *QueryPreCheckDatabaseRequest
	GetInstanceUuid() *string
	SetTaskId(v string) *QueryPreCheckDatabaseRequest
	GetTaskId() *string
	SetUniRegionId(v string) *QueryPreCheckDatabaseRequest
	GetUniRegionId() *string
}

type QueryPreCheckDatabaseRequest struct {
	// The unique identifier of the server database backup client.
	//
	// > You can call the [DescribeUniBackupDatabase](~~DescribeUniBackupDatabase~~) operation to obtain this parameter.
	//
	// example:
	//
	// ebc895506c6911ed800000163e0e****
	InstanceUuid *string `json:"InstanceUuid,omitempty" xml:"InstanceUuid,omitempty"`
	// The ID of the database pre-check task.
	//
	// > You can call the [StartPreCheckDatabase](~~StartPreCheckDatabase~~) operation to obtain this parameter.
	//
	// example:
	//
	// t-000bc9nqwxsbyvod****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The region ID of the database server.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	UniRegionId *string `json:"UniRegionId,omitempty" xml:"UniRegionId,omitempty"`
}

func (s QueryPreCheckDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPreCheckDatabaseRequest) GoString() string {
	return s.String()
}

func (s *QueryPreCheckDatabaseRequest) GetInstanceUuid() *string {
	return s.InstanceUuid
}

func (s *QueryPreCheckDatabaseRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryPreCheckDatabaseRequest) GetUniRegionId() *string {
	return s.UniRegionId
}

func (s *QueryPreCheckDatabaseRequest) SetInstanceUuid(v string) *QueryPreCheckDatabaseRequest {
	s.InstanceUuid = &v
	return s
}

func (s *QueryPreCheckDatabaseRequest) SetTaskId(v string) *QueryPreCheckDatabaseRequest {
	s.TaskId = &v
	return s
}

func (s *QueryPreCheckDatabaseRequest) SetUniRegionId(v string) *QueryPreCheckDatabaseRequest {
	s.UniRegionId = &v
	return s
}

func (s *QueryPreCheckDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
