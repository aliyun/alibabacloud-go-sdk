// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableHanaBackupPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DisableHanaBackupPlanRequest
	GetClusterId() *string
	SetPlanId(v string) *DisableHanaBackupPlanRequest
	GetPlanId() *string
	SetResourceGroupId(v string) *DisableHanaBackupPlanRequest
	GetResourceGroupId() *string
	SetVaultId(v string) *DisableHanaBackupPlanRequest
	GetVaultId() *string
}

type DisableHanaBackupPlanRequest struct {
	// The ID of the SAP HANA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cl-0003tu******y5oc
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the backup plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// pl-0006o11ectqr650ceoct
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm3erpwweavki
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the backup vault.
	//
	// This parameter is required.
	//
	// example:
	//
	// v-000f9z******vilrr
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DisableHanaBackupPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableHanaBackupPlanRequest) GoString() string {
	return s.String()
}

func (s *DisableHanaBackupPlanRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DisableHanaBackupPlanRequest) GetPlanId() *string {
	return s.PlanId
}

func (s *DisableHanaBackupPlanRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DisableHanaBackupPlanRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *DisableHanaBackupPlanRequest) SetClusterId(v string) *DisableHanaBackupPlanRequest {
	s.ClusterId = &v
	return s
}

func (s *DisableHanaBackupPlanRequest) SetPlanId(v string) *DisableHanaBackupPlanRequest {
	s.PlanId = &v
	return s
}

func (s *DisableHanaBackupPlanRequest) SetResourceGroupId(v string) *DisableHanaBackupPlanRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DisableHanaBackupPlanRequest) SetVaultId(v string) *DisableHanaBackupPlanRequest {
	s.VaultId = &v
	return s
}

func (s *DisableHanaBackupPlanRequest) Validate() error {
	return dara.Validate(s)
}
