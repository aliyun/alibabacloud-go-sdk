// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSqlAuditRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAuditAccountName(v string) *EnableSqlAuditRequest
  GetAuditAccountName() *string 
  SetAuditAccountPassword(v string) *EnableSqlAuditRequest
  GetAuditAccountPassword() *string 
  SetDBInstanceId(v string) *EnableSqlAuditRequest
  GetDBInstanceId() *string 
  SetExpireAfterDays(v int32) *EnableSqlAuditRequest
  GetExpireAfterDays() *int32 
  SetRegionId(v string) *EnableSqlAuditRequest
  GetRegionId() *string 
}

type EnableSqlAuditRequest struct {
  // example:
  // 
  // test_daa
  AuditAccountName *string `json:"AuditAccountName,omitempty" xml:"AuditAccountName,omitempty"`
  // example:
  // 
  // Pw@11111
  AuditAccountPassword *string `json:"AuditAccountPassword,omitempty" xml:"AuditAccountPassword,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // pxc-****************
  DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
  // example:
  // 
  // 45
  ExpireAfterDays *int32 `json:"ExpireAfterDays,omitempty" xml:"ExpireAfterDays,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s EnableSqlAuditRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableSqlAuditRequest) GoString() string {
  return s.String()
}

func (s *EnableSqlAuditRequest) GetAuditAccountName() *string  {
  return s.AuditAccountName
}

func (s *EnableSqlAuditRequest) GetAuditAccountPassword() *string  {
  return s.AuditAccountPassword
}

func (s *EnableSqlAuditRequest) GetDBInstanceId() *string  {
  return s.DBInstanceId
}

func (s *EnableSqlAuditRequest) GetExpireAfterDays() *int32  {
  return s.ExpireAfterDays
}

func (s *EnableSqlAuditRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableSqlAuditRequest) SetAuditAccountName(v string) *EnableSqlAuditRequest {
  s.AuditAccountName = &v
  return s
}

func (s *EnableSqlAuditRequest) SetAuditAccountPassword(v string) *EnableSqlAuditRequest {
  s.AuditAccountPassword = &v
  return s
}

func (s *EnableSqlAuditRequest) SetDBInstanceId(v string) *EnableSqlAuditRequest {
  s.DBInstanceId = &v
  return s
}

func (s *EnableSqlAuditRequest) SetExpireAfterDays(v int32) *EnableSqlAuditRequest {
  s.ExpireAfterDays = &v
  return s
}

func (s *EnableSqlAuditRequest) SetRegionId(v string) *EnableSqlAuditRequest {
  s.RegionId = &v
  return s
}

func (s *EnableSqlAuditRequest) Validate() error {
  return dara.Validate(s)
}

