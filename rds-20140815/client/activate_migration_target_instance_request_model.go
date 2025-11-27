// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateMigrationTargetInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *ActivateMigrationTargetInstanceRequest
	GetDBInstanceName() *string
	SetForceSwitch(v string) *ActivateMigrationTargetInstanceRequest
	GetForceSwitch() *string
	SetResourceOwnerId(v int64) *ActivateMigrationTargetInstanceRequest
	GetResourceOwnerId() *int64
	SetSwitchTime(v string) *ActivateMigrationTargetInstanceRequest
	GetSwitchTime() *string
	SetSwitchTimeMode(v string) *ActivateMigrationTargetInstanceRequest
	GetSwitchTimeMode() *string
}

type ActivateMigrationTargetInstanceRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pgm-bp102g323jd4****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// Specifies whether to forcefully perform a switchover. Set the value to 1. The value 1 specifies a forceful switchover.
	//
	// example:
	//
	// 1
	ForceSwitch     *string `json:"ForceSwitch,omitempty" xml:"ForceSwitch,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// A reserved parameter. This parameter does not take effect.
	//
	// example:
	//
	// 2022-02-25T06:57:41Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// The time when you want to perform the switchover.
	//
	// Set the value to 0. The value 0 specifies an immediate switchover.
	//
	// example:
	//
	// 0
	SwitchTimeMode *string `json:"SwitchTimeMode,omitempty" xml:"SwitchTimeMode,omitempty"`
}

func (s ActivateMigrationTargetInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ActivateMigrationTargetInstanceRequest) GoString() string {
	return s.String()
}

func (s *ActivateMigrationTargetInstanceRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ActivateMigrationTargetInstanceRequest) GetForceSwitch() *string {
	return s.ForceSwitch
}

func (s *ActivateMigrationTargetInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ActivateMigrationTargetInstanceRequest) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *ActivateMigrationTargetInstanceRequest) GetSwitchTimeMode() *string {
	return s.SwitchTimeMode
}

func (s *ActivateMigrationTargetInstanceRequest) SetDBInstanceName(v string) *ActivateMigrationTargetInstanceRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ActivateMigrationTargetInstanceRequest) SetForceSwitch(v string) *ActivateMigrationTargetInstanceRequest {
	s.ForceSwitch = &v
	return s
}

func (s *ActivateMigrationTargetInstanceRequest) SetResourceOwnerId(v int64) *ActivateMigrationTargetInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ActivateMigrationTargetInstanceRequest) SetSwitchTime(v string) *ActivateMigrationTargetInstanceRequest {
	s.SwitchTime = &v
	return s
}

func (s *ActivateMigrationTargetInstanceRequest) SetSwitchTimeMode(v string) *ActivateMigrationTargetInstanceRequest {
	s.SwitchTimeMode = &v
	return s
}

func (s *ActivateMigrationTargetInstanceRequest) Validate() error {
	return dara.Validate(s)
}
