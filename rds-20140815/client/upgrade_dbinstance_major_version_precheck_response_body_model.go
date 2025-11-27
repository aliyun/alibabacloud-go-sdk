// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeDBInstanceMajorVersionPrecheckResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *UpgradeDBInstanceMajorVersionPrecheckResponseBody
	GetDBInstanceName() *string
	SetRequestId(v string) *UpgradeDBInstanceMajorVersionPrecheckResponseBody
	GetRequestId() *string
	SetTargetMajorVersion(v string) *UpgradeDBInstanceMajorVersionPrecheckResponseBody
	GetTargetMajorVersion() *string
	SetTaskId(v string) *UpgradeDBInstanceMajorVersionPrecheckResponseBody
	GetTaskId() *string
}

type UpgradeDBInstanceMajorVersionPrecheckResponseBody struct {
	// The instance name.
	//
	// example:
	//
	// pgm-bp1c808s731l****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 99C1FEEE-FB44-5342-8EBA-DC1E1A1557A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The new major engine version of the instance.
	//
	// example:
	//
	// 12.0
	TargetMajorVersion *string `json:"TargetMajorVersion,omitempty" xml:"TargetMajorVersion,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 41698****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpgradeDBInstanceMajorVersionPrecheckResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDBInstanceMajorVersionPrecheckResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceMajorVersionPrecheckResponseBody) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *UpgradeDBInstanceMajorVersionPrecheckResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeDBInstanceMajorVersionPrecheckResponseBody) GetTargetMajorVersion() *string {
	return s.TargetMajorVersion
}

func (s *UpgradeDBInstanceMajorVersionPrecheckResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *UpgradeDBInstanceMajorVersionPrecheckResponseBody) SetDBInstanceName(v string) *UpgradeDBInstanceMajorVersionPrecheckResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionPrecheckResponseBody) SetRequestId(v string) *UpgradeDBInstanceMajorVersionPrecheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionPrecheckResponseBody) SetTargetMajorVersion(v string) *UpgradeDBInstanceMajorVersionPrecheckResponseBody {
	s.TargetMajorVersion = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionPrecheckResponseBody) SetTaskId(v string) *UpgradeDBInstanceMajorVersionPrecheckResponseBody {
	s.TaskId = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionPrecheckResponseBody) Validate() error {
	return dara.Validate(s)
}
