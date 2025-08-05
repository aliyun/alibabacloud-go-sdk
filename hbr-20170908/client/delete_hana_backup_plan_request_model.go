// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHanaBackupPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteHanaBackupPlanRequest
	GetClusterId() *string
	SetPlanId(v string) *DeleteHanaBackupPlanRequest
	GetPlanId() *string
	SetResourceGroupId(v string) *DeleteHanaBackupPlanRequest
	GetResourceGroupId() *string
	SetVaultId(v string) *DeleteHanaBackupPlanRequest
	GetVaultId() *string
}

type DeleteHanaBackupPlanRequest struct {
	// The ID of the SAP HANA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cl-000br3******0ooy2
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the backup plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// pl-00035lc8pwp1azdf3qku
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm4ebtpkzx7zy
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the backup vault.
	//
	// This parameter is required.
	//
	// example:
	//
	// v-0007o******1ssno
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DeleteHanaBackupPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHanaBackupPlanRequest) GoString() string {
	return s.String()
}

func (s *DeleteHanaBackupPlanRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteHanaBackupPlanRequest) GetPlanId() *string {
	return s.PlanId
}

func (s *DeleteHanaBackupPlanRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteHanaBackupPlanRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *DeleteHanaBackupPlanRequest) SetClusterId(v string) *DeleteHanaBackupPlanRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteHanaBackupPlanRequest) SetPlanId(v string) *DeleteHanaBackupPlanRequest {
	s.PlanId = &v
	return s
}

func (s *DeleteHanaBackupPlanRequest) SetResourceGroupId(v string) *DeleteHanaBackupPlanRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteHanaBackupPlanRequest) SetVaultId(v string) *DeleteHanaBackupPlanRequest {
	s.VaultId = &v
	return s
}

func (s *DeleteHanaBackupPlanRequest) Validate() error {
	return dara.Validate(s)
}
