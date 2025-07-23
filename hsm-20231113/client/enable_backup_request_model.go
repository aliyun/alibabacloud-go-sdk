// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableBackupRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBackupId(v string) *EnableBackupRequest
  GetBackupId() *string 
  SetInstanceId(v string) *EnableBackupRequest
  GetInstanceId() *string 
}

type EnableBackupRequest struct {
  // The ID of the backup.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // backup-1736207****
  BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
  // The ID of the HSM.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // hsm-cn-mp90fxef****
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s EnableBackupRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableBackupRequest) GoString() string {
  return s.String()
}

func (s *EnableBackupRequest) GetBackupId() *string  {
  return s.BackupId
}

func (s *EnableBackupRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableBackupRequest) SetBackupId(v string) *EnableBackupRequest {
  s.BackupId = &v
  return s
}

func (s *EnableBackupRequest) SetInstanceId(v string) *EnableBackupRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableBackupRequest) Validate() error {
  return dara.Validate(s)
}

