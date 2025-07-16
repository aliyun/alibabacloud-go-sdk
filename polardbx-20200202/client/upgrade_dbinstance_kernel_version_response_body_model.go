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
	// example:
	//
	// pxc-hzjasd****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// polarx-kernel_5.4.7-16001481_xcluster-20200910
	TargetMinorVersion *string `json:"TargetMinorVersion,omitempty" xml:"TargetMinorVersion,omitempty"`
	// example:
	//
	// 422922413
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
