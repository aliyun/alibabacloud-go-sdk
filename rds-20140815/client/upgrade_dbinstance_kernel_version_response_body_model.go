// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeDBInstanceKernelVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *UpgradeDBInstanceKernelVersionResponseBody
	GetDBInstanceName() *string
	SetRequestId(v string) *UpgradeDBInstanceKernelVersionResponseBody
	GetRequestId() *string
	SetTargetMinorVersion(v string) *UpgradeDBInstanceKernelVersionResponseBody
	GetTargetMinorVersion() *string
	SetTaskId(v string) *UpgradeDBInstanceKernelVersionResponseBody
	GetTaskId() *string
}

type UpgradeDBInstanceKernelVersionResponseBody struct {
	// The ID of the instance.
	//
	// example:
	//
	// rm-bpxxxxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// DA2ECBA0-4745-4491-9166-799FF8984AC9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The new minor engine version of the instance.
	//
	// example:
	//
	// xcluster80_20210305
	TargetMinorVersion *string `json:"TargetMinorVersion,omitempty" xml:"TargetMinorVersion,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 226917****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpgradeDBInstanceKernelVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDBInstanceKernelVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceKernelVersionResponseBody) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *UpgradeDBInstanceKernelVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeDBInstanceKernelVersionResponseBody) GetTargetMinorVersion() *string {
	return s.TargetMinorVersion
}

func (s *UpgradeDBInstanceKernelVersionResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *UpgradeDBInstanceKernelVersionResponseBody) SetDBInstanceName(v string) *UpgradeDBInstanceKernelVersionResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionResponseBody) SetRequestId(v string) *UpgradeDBInstanceKernelVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionResponseBody) SetTargetMinorVersion(v string) *UpgradeDBInstanceKernelVersionResponseBody {
	s.TargetMinorVersion = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionResponseBody) SetTaskId(v string) *UpgradeDBInstanceKernelVersionResponseBody {
	s.TaskId = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
