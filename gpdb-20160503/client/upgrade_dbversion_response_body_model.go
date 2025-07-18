// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeDBVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *UpgradeDBVersionResponseBody
	GetDBInstanceId() *string
	SetDBInstanceName(v string) *UpgradeDBVersionResponseBody
	GetDBInstanceName() *string
	SetRequestId(v string) *UpgradeDBVersionResponseBody
	GetRequestId() *string
	SetTaskId(v string) *UpgradeDBVersionResponseBody
	GetTaskId() *string
}

type UpgradeDBVersionResponseBody struct {
	// This parameter is no longer returned.
	//
	// example:
	//
	// gp-wz9kmr708m155j***
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// gp-wz9kmr708m155j***
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 25C11EE5-B7E8-481A-A07C-BD619971A570
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 101450956
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpgradeDBVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDBVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeDBVersionResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *UpgradeDBVersionResponseBody) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *UpgradeDBVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeDBVersionResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *UpgradeDBVersionResponseBody) SetDBInstanceId(v string) *UpgradeDBVersionResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *UpgradeDBVersionResponseBody) SetDBInstanceName(v string) *UpgradeDBVersionResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *UpgradeDBVersionResponseBody) SetRequestId(v string) *UpgradeDBVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeDBVersionResponseBody) SetTaskId(v string) *UpgradeDBVersionResponseBody {
	s.TaskId = &v
	return s
}

func (s *UpgradeDBVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
