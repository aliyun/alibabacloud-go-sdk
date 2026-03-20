// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableBackupEncryptionRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDBInstanceName(v string) *EnableBackupEncryptionRequest
  GetDBInstanceName() *string 
  SetEncryptionKey(v string) *EnableBackupEncryptionRequest
  GetEncryptionKey() *string 
  SetResourceOwnerId(v int64) *EnableBackupEncryptionRequest
  GetResourceOwnerId() *int64 
}

type EnableBackupEncryptionRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // rm-wz951f7f******
  DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
  // example:
  // 
  // 564cf6c4-d2ee-495b-b265-5724******
  EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
  ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s EnableBackupEncryptionRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableBackupEncryptionRequest) GoString() string {
  return s.String()
}

func (s *EnableBackupEncryptionRequest) GetDBInstanceName() *string  {
  return s.DBInstanceName
}

func (s *EnableBackupEncryptionRequest) GetEncryptionKey() *string  {
  return s.EncryptionKey
}

func (s *EnableBackupEncryptionRequest) GetResourceOwnerId() *int64  {
  return s.ResourceOwnerId
}

func (s *EnableBackupEncryptionRequest) SetDBInstanceName(v string) *EnableBackupEncryptionRequest {
  s.DBInstanceName = &v
  return s
}

func (s *EnableBackupEncryptionRequest) SetEncryptionKey(v string) *EnableBackupEncryptionRequest {
  s.EncryptionKey = &v
  return s
}

func (s *EnableBackupEncryptionRequest) SetResourceOwnerId(v int64) *EnableBackupEncryptionRequest {
  s.ResourceOwnerId = &v
  return s
}

func (s *EnableBackupEncryptionRequest) Validate() error {
  return dara.Validate(s)
}

