// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableBackupLogRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBackupPlanId(v string) *EnableBackupLogRequest
  GetBackupPlanId() *string 
  SetClientToken(v string) *EnableBackupLogRequest
  GetClientToken() *string 
  SetEnableMysqlPhysicalBackupBinlog(v string) *EnableBackupLogRequest
  GetEnableMysqlPhysicalBackupBinlog() *string 
  SetOwnerId(v string) *EnableBackupLogRequest
  GetOwnerId() *string 
}

type EnableBackupLogRequest struct {
  // The ID of the backup plan. Call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to query the ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // dbstooi******
  BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
  // Any string value.
  // 
  // example:
  // 
  // dbs
  ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
  // example:
  // 
  // true
  EnableMysqlPhysicalBackupBinlog *string `json:"EnableMysqlPhysicalBackupBinlog,omitempty" xml:"EnableMysqlPhysicalBackupBinlog,omitempty"`
  OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s EnableBackupLogRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableBackupLogRequest) GoString() string {
  return s.String()
}

func (s *EnableBackupLogRequest) GetBackupPlanId() *string  {
  return s.BackupPlanId
}

func (s *EnableBackupLogRequest) GetClientToken() *string  {
  return s.ClientToken
}

func (s *EnableBackupLogRequest) GetEnableMysqlPhysicalBackupBinlog() *string  {
  return s.EnableMysqlPhysicalBackupBinlog
}

func (s *EnableBackupLogRequest) GetOwnerId() *string  {
  return s.OwnerId
}

func (s *EnableBackupLogRequest) SetBackupPlanId(v string) *EnableBackupLogRequest {
  s.BackupPlanId = &v
  return s
}

func (s *EnableBackupLogRequest) SetClientToken(v string) *EnableBackupLogRequest {
  s.ClientToken = &v
  return s
}

func (s *EnableBackupLogRequest) SetEnableMysqlPhysicalBackupBinlog(v string) *EnableBackupLogRequest {
  s.EnableMysqlPhysicalBackupBinlog = &v
  return s
}

func (s *EnableBackupLogRequest) SetOwnerId(v string) *EnableBackupLogRequest {
  s.OwnerId = &v
  return s
}

func (s *EnableBackupLogRequest) Validate() error {
  return dara.Validate(s)
}

