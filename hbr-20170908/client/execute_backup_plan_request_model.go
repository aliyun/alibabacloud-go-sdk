// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteBackupPlanRequest interface {
  dara.Model
  String() string
  GoString() string
  SetPlanId(v string) *ExecuteBackupPlanRequest
  GetPlanId() *string 
  SetRuleId(v string) *ExecuteBackupPlanRequest
  GetRuleId() *string 
  SetSourceType(v string) *ExecuteBackupPlanRequest
  GetSourceType() *string 
  SetVaultId(v string) *ExecuteBackupPlanRequest
  GetVaultId() *string 
}

type ExecuteBackupPlanRequest struct {
  // The ID of the backup plan.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // plan-*********************
  PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
  // The ID of the backup rule.
  // 
  // example:
  // 
  // rule-0002*****ux8
  RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
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

func (s ExecuteBackupPlanRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteBackupPlanRequest) GoString() string {
  return s.String()
}

func (s *ExecuteBackupPlanRequest) GetPlanId() *string  {
  return s.PlanId
}

func (s *ExecuteBackupPlanRequest) GetRuleId() *string  {
  return s.RuleId
}

func (s *ExecuteBackupPlanRequest) GetSourceType() *string  {
  return s.SourceType
}

func (s *ExecuteBackupPlanRequest) GetVaultId() *string  {
  return s.VaultId
}

func (s *ExecuteBackupPlanRequest) SetPlanId(v string) *ExecuteBackupPlanRequest {
  s.PlanId = &v
  return s
}

func (s *ExecuteBackupPlanRequest) SetRuleId(v string) *ExecuteBackupPlanRequest {
  s.RuleId = &v
  return s
}

func (s *ExecuteBackupPlanRequest) SetSourceType(v string) *ExecuteBackupPlanRequest {
  s.SourceType = &v
  return s
}

func (s *ExecuteBackupPlanRequest) SetVaultId(v string) *ExecuteBackupPlanRequest {
  s.VaultId = &v
  return s
}

func (s *ExecuteBackupPlanRequest) Validate() error {
  return dara.Validate(s)
}

