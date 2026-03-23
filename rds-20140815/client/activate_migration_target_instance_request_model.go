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
	// This parameter is required.
	DBInstanceName  *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	ForceSwitch     *string `json:"ForceSwitch,omitempty" xml:"ForceSwitch,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SwitchTime      *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	SwitchTimeMode  *string `json:"SwitchTimeMode,omitempty" xml:"SwitchTimeMode,omitempty"`
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
