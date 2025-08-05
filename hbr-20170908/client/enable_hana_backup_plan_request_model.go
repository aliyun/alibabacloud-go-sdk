// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableHanaBackupPlanRequest interface {
  dara.Model
  String() string
  GoString() string
  SetClusterId(v string) *EnableHanaBackupPlanRequest
  GetClusterId() *string 
  SetPlanId(v string) *EnableHanaBackupPlanRequest
  GetPlanId() *string 
  SetResourceGroupId(v string) *EnableHanaBackupPlanRequest
  GetResourceGroupId() *string 
  SetVaultId(v string) *EnableHanaBackupPlanRequest
  GetVaultId() *string 
}

type EnableHanaBackupPlanRequest struct {
  // The ID of the SAP HANA instance.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cl-0001zfcn******0pr3
  ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
  // The ID of the backup plan.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // plan-*********************
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
  // v-00030j3c******sn
  VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s EnableHanaBackupPlanRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableHanaBackupPlanRequest) GoString() string {
  return s.String()
}

func (s *EnableHanaBackupPlanRequest) GetClusterId() *string  {
  return s.ClusterId
}

func (s *EnableHanaBackupPlanRequest) GetPlanId() *string  {
  return s.PlanId
}

func (s *EnableHanaBackupPlanRequest) GetResourceGroupId() *string  {
  return s.ResourceGroupId
}

func (s *EnableHanaBackupPlanRequest) GetVaultId() *string  {
  return s.VaultId
}

func (s *EnableHanaBackupPlanRequest) SetClusterId(v string) *EnableHanaBackupPlanRequest {
  s.ClusterId = &v
  return s
}

func (s *EnableHanaBackupPlanRequest) SetPlanId(v string) *EnableHanaBackupPlanRequest {
  s.PlanId = &v
  return s
}

func (s *EnableHanaBackupPlanRequest) SetResourceGroupId(v string) *EnableHanaBackupPlanRequest {
  s.ResourceGroupId = &v
  return s
}

func (s *EnableHanaBackupPlanRequest) SetVaultId(v string) *EnableHanaBackupPlanRequest {
  s.VaultId = &v
  return s
}

func (s *EnableHanaBackupPlanRequest) Validate() error {
  return dara.Validate(s)
}

