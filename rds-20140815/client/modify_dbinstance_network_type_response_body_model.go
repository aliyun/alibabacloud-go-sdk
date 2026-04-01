// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceNetworkTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionString(v string) *ModifyDBInstanceNetworkTypeResponseBody
	GetConnectionString() *string
	SetRequestId(v string) *ModifyDBInstanceNetworkTypeResponseBody
	GetRequestId() *string
	SetTaskId(v string) *ModifyDBInstanceNetworkTypeResponseBody
	GetTaskId() *string
}

type ModifyDBInstanceNetworkTypeResponseBody struct {
	// The endpoint that is used to connect to the instance.
	//
	// example:
	//
	// rm-bp1*****************.mysql.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 1025486523574
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyDBInstanceNetworkTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceNetworkTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceNetworkTypeResponseBody) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *ModifyDBInstanceNetworkTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBInstanceNetworkTypeResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *ModifyDBInstanceNetworkTypeResponseBody) SetConnectionString(v string) *ModifyDBInstanceNetworkTypeResponseBody {
	s.ConnectionString = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeResponseBody) SetRequestId(v string) *ModifyDBInstanceNetworkTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeResponseBody) SetTaskId(v string) *ModifyDBInstanceNetworkTypeResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
