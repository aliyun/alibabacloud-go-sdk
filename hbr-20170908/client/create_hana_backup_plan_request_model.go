// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHanaBackupPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPrefix(v string) *CreateHanaBackupPlanRequest
	GetBackupPrefix() *string
	SetBackupType(v string) *CreateHanaBackupPlanRequest
	GetBackupType() *string
	SetClusterId(v string) *CreateHanaBackupPlanRequest
	GetClusterId() *string
	SetDatabaseName(v string) *CreateHanaBackupPlanRequest
	GetDatabaseName() *string
	SetPlanName(v string) *CreateHanaBackupPlanRequest
	GetPlanName() *string
	SetResourceGroupId(v string) *CreateHanaBackupPlanRequest
	GetResourceGroupId() *string
	SetSchedule(v string) *CreateHanaBackupPlanRequest
	GetSchedule() *string
	SetVaultId(v string) *CreateHanaBackupPlanRequest
	GetVaultId() *string
}

type CreateHanaBackupPlanRequest struct {
	// The backup prefix.
	//
	// example:
	//
	// DIFF_DATA_BACKUP
	BackupPrefix *string `json:"BackupPrefix,omitempty" xml:"BackupPrefix,omitempty"`
	// The backup type. Valid values:
	//
	// 	- COMPLETE: full backup
	//
	// 	- INCREMENTAL: incremental backup
	//
	// 	- DIFFERENTIAL: differential backup
	//
	// This parameter is required.
	//
	// example:
	//
	// COMPLETE
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// The ID of the SAP HANA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cl-00024vyjj9******v
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// SYSTEMDB
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The name of the backup plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// plan-20220110-113108
	PlanName *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmvnf22m7itha
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The backup policy. Format: `I|{startTime}|{interval}`. The system runs the first backup job at a point in time that is specified in the {startTime} parameter and the subsequent backup jobs at an interval that is specified in the {interval} parameter. The system does not run a backup job before the specified point in time. Each backup job, except the first one, starts only after the previous backup job is completed. For example, `I|1631685600|P1D` specifies that the system runs the first backup job at 14:00:00 on September 15, 2021 and the subsequent backup jobs once a day.
	//
	// 	- startTime: the time at which the system starts to run a backup job. The time must follow the UNIX time format. Unit: seconds.
	//
	// 	- interval: the interval at which the system runs a backup job. The interval must follow the ISO 8601 standard. For example, PT1H specifies an interval of one hour. P1D specifies an interval of one day.
	//
	// This parameter is required.
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
	// v-0002pcwhdn******wmi
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s CreateHanaBackupPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHanaBackupPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateHanaBackupPlanRequest) GetBackupPrefix() *string {
	return s.BackupPrefix
}

func (s *CreateHanaBackupPlanRequest) GetBackupType() *string {
	return s.BackupType
}

func (s *CreateHanaBackupPlanRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateHanaBackupPlanRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *CreateHanaBackupPlanRequest) GetPlanName() *string {
	return s.PlanName
}

func (s *CreateHanaBackupPlanRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateHanaBackupPlanRequest) GetSchedule() *string {
	return s.Schedule
}

func (s *CreateHanaBackupPlanRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *CreateHanaBackupPlanRequest) SetBackupPrefix(v string) *CreateHanaBackupPlanRequest {
	s.BackupPrefix = &v
	return s
}

func (s *CreateHanaBackupPlanRequest) SetBackupType(v string) *CreateHanaBackupPlanRequest {
	s.BackupType = &v
	return s
}

func (s *CreateHanaBackupPlanRequest) SetClusterId(v string) *CreateHanaBackupPlanRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateHanaBackupPlanRequest) SetDatabaseName(v string) *CreateHanaBackupPlanRequest {
	s.DatabaseName = &v
	return s
}

func (s *CreateHanaBackupPlanRequest) SetPlanName(v string) *CreateHanaBackupPlanRequest {
	s.PlanName = &v
	return s
}

func (s *CreateHanaBackupPlanRequest) SetResourceGroupId(v string) *CreateHanaBackupPlanRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateHanaBackupPlanRequest) SetSchedule(v string) *CreateHanaBackupPlanRequest {
	s.Schedule = &v
	return s
}

func (s *CreateHanaBackupPlanRequest) SetVaultId(v string) *CreateHanaBackupPlanRequest {
	s.VaultId = &v
	return s
}

func (s *CreateHanaBackupPlanRequest) Validate() error {
	return dara.Validate(s)
}
