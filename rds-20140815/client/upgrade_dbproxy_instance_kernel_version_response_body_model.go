// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeDBProxyInstanceKernelVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *UpgradeDBProxyInstanceKernelVersionResponseBody
	GetDBInstanceName() *string
	SetRequestId(v string) *UpgradeDBProxyInstanceKernelVersionResponseBody
	GetRequestId() *string
	SetTaskId(v string) *UpgradeDBProxyInstanceKernelVersionResponseBody
	GetTaskId() *string
}

type UpgradeDBProxyInstanceKernelVersionResponseBody struct {
	// The ID of the database proxy of the instance.
	//
	// example:
	//
	// bu9***
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 44537EC8-DFA2-4745-B579-E733FF2C5B9A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 33436****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpgradeDBProxyInstanceKernelVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDBProxyInstanceKernelVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeDBProxyInstanceKernelVersionResponseBody) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *UpgradeDBProxyInstanceKernelVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeDBProxyInstanceKernelVersionResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *UpgradeDBProxyInstanceKernelVersionResponseBody) SetDBInstanceName(v string) *UpgradeDBProxyInstanceKernelVersionResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *UpgradeDBProxyInstanceKernelVersionResponseBody) SetRequestId(v string) *UpgradeDBProxyInstanceKernelVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeDBProxyInstanceKernelVersionResponseBody) SetTaskId(v string) *UpgradeDBProxyInstanceKernelVersionResponseBody {
	s.TaskId = &v
	return s
}

func (s *UpgradeDBProxyInstanceKernelVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
