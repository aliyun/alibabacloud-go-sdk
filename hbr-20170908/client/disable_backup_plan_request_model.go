// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableBackupPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPlanId(v string) *DisableBackupPlanRequest
	GetPlanId() *string
	SetSourceType(v string) *DisableBackupPlanRequest
	GetSourceType() *string
	SetVaultId(v string) *DisableBackupPlanRequest
	GetVaultId() *string
}

type DisableBackupPlanRequest struct {
	// The ID of the backup plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// plan-*********************
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- **ECS_FILE**: Elastic Compute Service (ECS) files
	//
	// 	- **OSS**: Object Storage Service (OSS) buckets
	//
	// 	- **NAS**: Apsara File Storage NAS (NAS) file systems
	//
	// example:
	//
	// ECS_FILE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The ID of the backup vault.
	//
	// example:
	//
	// v-*********************
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DisableBackupPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableBackupPlanRequest) GoString() string {
	return s.String()
}

func (s *DisableBackupPlanRequest) GetPlanId() *string {
	return s.PlanId
}

func (s *DisableBackupPlanRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *DisableBackupPlanRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *DisableBackupPlanRequest) SetPlanId(v string) *DisableBackupPlanRequest {
	s.PlanId = &v
	return s
}

func (s *DisableBackupPlanRequest) SetSourceType(v string) *DisableBackupPlanRequest {
	s.SourceType = &v
	return s
}

func (s *DisableBackupPlanRequest) SetVaultId(v string) *DisableBackupPlanRequest {
	s.VaultId = &v
	return s
}

func (s *DisableBackupPlanRequest) Validate() error {
	return dara.Validate(s)
}
