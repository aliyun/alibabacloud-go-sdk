// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableBackupPlanRequest interface {
  dara.Model
  String() string
  GoString() string
  SetPlanId(v string) *EnableBackupPlanRequest
  GetPlanId() *string 
  SetSourceType(v string) *EnableBackupPlanRequest
  GetSourceType() *string 
  SetVaultId(v string) *EnableBackupPlanRequest
  GetVaultId() *string 
}

type EnableBackupPlanRequest struct {
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
  // 	- **ECS_FILE**: ECS files
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

func (s EnableBackupPlanRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableBackupPlanRequest) GoString() string {
  return s.String()
}

func (s *EnableBackupPlanRequest) GetPlanId() *string  {
  return s.PlanId
}

func (s *EnableBackupPlanRequest) GetSourceType() *string  {
  return s.SourceType
}

func (s *EnableBackupPlanRequest) GetVaultId() *string  {
  return s.VaultId
}

func (s *EnableBackupPlanRequest) SetPlanId(v string) *EnableBackupPlanRequest {
  s.PlanId = &v
  return s
}

func (s *EnableBackupPlanRequest) SetSourceType(v string) *EnableBackupPlanRequest {
  s.SourceType = &v
  return s
}

func (s *EnableBackupPlanRequest) SetVaultId(v string) *EnableBackupPlanRequest {
  s.VaultId = &v
  return s
}

func (s *EnableBackupPlanRequest) Validate() error {
  return dara.Validate(s)
}

