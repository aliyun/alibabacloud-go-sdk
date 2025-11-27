// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeDBInstanceEngineVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpgradeDBInstanceEngineVersionResponseBody
	GetRequestId() *string
	SetTaskId(v string) *UpgradeDBInstanceEngineVersionResponseBody
	GetTaskId() *string
}

type UpgradeDBInstanceEngineVersionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 65BDA532-28AF-4122-AA39-B382721EEE64
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 10254125
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpgradeDBInstanceEngineVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDBInstanceEngineVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceEngineVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeDBInstanceEngineVersionResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *UpgradeDBInstanceEngineVersionResponseBody) SetRequestId(v string) *UpgradeDBInstanceEngineVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeDBInstanceEngineVersionResponseBody) SetTaskId(v string) *UpgradeDBInstanceEngineVersionResponseBody {
	s.TaskId = &v
	return s
}

func (s *UpgradeDBInstanceEngineVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
