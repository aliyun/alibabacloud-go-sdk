// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHanaBackupPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPrefix(v string) *UpdateHanaBackupPlanRequest
	GetBackupPrefix() *string
	SetClusterId(v string) *UpdateHanaBackupPlanRequest
	GetClusterId() *string
	SetPlanId(v string) *UpdateHanaBackupPlanRequest
	GetPlanId() *string
	SetPlanName(v string) *UpdateHanaBackupPlanRequest
	GetPlanName() *string
	SetResourceGroupId(v string) *UpdateHanaBackupPlanRequest
	GetResourceGroupId() *string
	SetSchedule(v string) *UpdateHanaBackupPlanRequest
	GetSchedule() *string
	SetVaultId(v string) *UpdateHanaBackupPlanRequest
	GetVaultId() *string
}

type UpdateHanaBackupPlanRequest struct {
	// The backup prefix.
	//
	// example:
	//
	// COMPLETE_DATA_BACKUP
	BackupPrefix *string `json:"BackupPrefix,omitempty" xml:"BackupPrefix,omitempty"`
	// The ID of the SAP HANA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cl-0005dhe******f38
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the backup plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// pl-000br3cm4dqvmtph7cul
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The name of the backup plan.
	//
	// example:
	//
	// plan-20211109-162411
	PlanName *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmze36euddwjq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The backup policy. Format: `I|{startTime}|{interval}`. The system runs the first backup job at a point in time that is specified in the {startTime} parameter and the subsequent backup jobs at an interval that is specified in the {interval} parameter. The system does not run a backup job before the specified point in time. Each backup job, except the first one, starts only after the previous backup job is completed. For example, `I|1631685600|P1D` indicates that the system runs the first backup job at 14:00:00 on September 15, 2021 and the subsequent backup jobs once a day.
	//
	// 	- startTime: the time at which the system starts to run a backup job. The time follows the UNIX time format. Unit: seconds.
	//
	// 	- interval: the interval at which the system runs a backup job. The interval follows the ISO 8601 standard. For example, PT1H indicates an interval of 1 hour. P1D indicates an interval of one day.
	//
	// example:
	//
	// I|1602673264|P1D
	Schedule *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	// The ID of the backup vault.
	//
	// This parameter is required.
	//
	// example:
	//
	// v-0000rcw******5c6
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s UpdateHanaBackupPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHanaBackupPlanRequest) GoString() string {
	return s.String()
}

func (s *UpdateHanaBackupPlanRequest) GetBackupPrefix() *string {
	return s.BackupPrefix
}

func (s *UpdateHanaBackupPlanRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateHanaBackupPlanRequest) GetPlanId() *string {
	return s.PlanId
}

func (s *UpdateHanaBackupPlanRequest) GetPlanName() *string {
	return s.PlanName
}

func (s *UpdateHanaBackupPlanRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateHanaBackupPlanRequest) GetSchedule() *string {
	return s.Schedule
}

func (s *UpdateHanaBackupPlanRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *UpdateHanaBackupPlanRequest) SetBackupPrefix(v string) *UpdateHanaBackupPlanRequest {
	s.BackupPrefix = &v
	return s
}

func (s *UpdateHanaBackupPlanRequest) SetClusterId(v string) *UpdateHanaBackupPlanRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateHanaBackupPlanRequest) SetPlanId(v string) *UpdateHanaBackupPlanRequest {
	s.PlanId = &v
	return s
}

func (s *UpdateHanaBackupPlanRequest) SetPlanName(v string) *UpdateHanaBackupPlanRequest {
	s.PlanName = &v
	return s
}

func (s *UpdateHanaBackupPlanRequest) SetResourceGroupId(v string) *UpdateHanaBackupPlanRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateHanaBackupPlanRequest) SetSchedule(v string) *UpdateHanaBackupPlanRequest {
	s.Schedule = &v
	return s
}

func (s *UpdateHanaBackupPlanRequest) SetVaultId(v string) *UpdateHanaBackupPlanRequest {
	s.VaultId = &v
	return s
}

func (s *UpdateHanaBackupPlanRequest) Validate() error {
	return dara.Validate(s)
}
