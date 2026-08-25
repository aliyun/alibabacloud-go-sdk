// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEdition(v string) *DeleteBackupPlanRequest
	GetEdition() *string
	SetPlanId(v string) *DeleteBackupPlanRequest
	GetPlanId() *string
	SetRequireNoRunningJobs(v bool) *DeleteBackupPlanRequest
	GetRequireNoRunningJobs() *bool
	SetSourceType(v string) *DeleteBackupPlanRequest
	GetSourceType() *string
	SetVaultId(v string) *DeleteBackupPlanRequest
	GetVaultId() *string
}

type DeleteBackupPlanRequest struct {
	// The edition type. Valid values are BASIC and STANDARD. The default value is STANDARD.
	//
	// example:
	//
	// STANDARD
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// The ID of the backup plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// plan-*********************
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// Specifies whether to prevent the deletion of the backup plan if backup jobs are running.
	//
	// example:
	//
	// false
	RequireNoRunningJobs *bool `json:"RequireNoRunningJobs,omitempty" xml:"RequireNoRunningJobs,omitempty"`
	// The type of the data source. Valid values:
	//
	// - **ECS_FILE**: Backs up Elastic Compute Service (ECS) files.
	//
	// - **OSS**: Backs up Object Storage Service (OSS) buckets.
	//
	// - **NAS**: Backs up Network Attached Storage (NAS) file systems.
	//
	// - **UDM_ECS**: Backs up entire ECS instances.
	//
	// - **OTS**: Backs up Tablestore instances.
	//
	// example:
	//
	// ECS_FILE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The ID of the backup vault. This parameter is required if the SourceType parameter is not set to UDM_ECS.
	//
	// example:
	//
	// v-*********************
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DeleteBackupPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupPlanRequest) GoString() string {
	return s.String()
}

func (s *DeleteBackupPlanRequest) GetEdition() *string {
	return s.Edition
}

func (s *DeleteBackupPlanRequest) GetPlanId() *string {
	return s.PlanId
}

func (s *DeleteBackupPlanRequest) GetRequireNoRunningJobs() *bool {
	return s.RequireNoRunningJobs
}

func (s *DeleteBackupPlanRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *DeleteBackupPlanRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *DeleteBackupPlanRequest) SetEdition(v string) *DeleteBackupPlanRequest {
	s.Edition = &v
	return s
}

func (s *DeleteBackupPlanRequest) SetPlanId(v string) *DeleteBackupPlanRequest {
	s.PlanId = &v
	return s
}

func (s *DeleteBackupPlanRequest) SetRequireNoRunningJobs(v bool) *DeleteBackupPlanRequest {
	s.RequireNoRunningJobs = &v
	return s
}

func (s *DeleteBackupPlanRequest) SetSourceType(v string) *DeleteBackupPlanRequest {
	s.SourceType = &v
	return s
}

func (s *DeleteBackupPlanRequest) SetVaultId(v string) *DeleteBackupPlanRequest {
	s.VaultId = &v
	return s
}

func (s *DeleteBackupPlanRequest) Validate() error {
	return dara.Validate(s)
}
