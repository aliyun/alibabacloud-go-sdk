// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableRightsSeparationRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAuditAccountDescription(v string) *EnableRightsSeparationRequest
  GetAuditAccountDescription() *string 
  SetAuditAccountName(v string) *EnableRightsSeparationRequest
  GetAuditAccountName() *string 
  SetAuditAccountPassword(v string) *EnableRightsSeparationRequest
  GetAuditAccountPassword() *string 
  SetDBInstanceName(v string) *EnableRightsSeparationRequest
  GetDBInstanceName() *string 
  SetRegionId(v string) *EnableRightsSeparationRequest
  GetRegionId() *string 
  SetSecurityAccountDescription(v string) *EnableRightsSeparationRequest
  GetSecurityAccountDescription() *string 
  SetSecurityAccountName(v string) *EnableRightsSeparationRequest
  GetSecurityAccountName() *string 
  SetSecurityAccountPassword(v string) *EnableRightsSeparationRequest
  GetSecurityAccountPassword() *string 
}

type EnableRightsSeparationRequest struct {
  // example:
  // 
  // test123
  AuditAccountDescription *string `json:"AuditAccountDescription,omitempty" xml:"AuditAccountDescription,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // account_audit
  AuditAccountName *string `json:"AuditAccountName,omitempty" xml:"AuditAccountName,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // ******
  AuditAccountPassword *string `json:"AuditAccountPassword,omitempty" xml:"AuditAccountPassword,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // pxc-htri0ori2r4k9p
  DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // example:
  // 
  // test123
  SecurityAccountDescription *string `json:"SecurityAccountDescription,omitempty" xml:"SecurityAccountDescription,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // account_sec
  SecurityAccountName *string `json:"SecurityAccountName,omitempty" xml:"SecurityAccountName,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // *****
  SecurityAccountPassword *string `json:"SecurityAccountPassword,omitempty" xml:"SecurityAccountPassword,omitempty"`
}

func (s EnableRightsSeparationRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableRightsSeparationRequest) GoString() string {
  return s.String()
}

func (s *EnableRightsSeparationRequest) GetAuditAccountDescription() *string  {
  return s.AuditAccountDescription
}

func (s *EnableRightsSeparationRequest) GetAuditAccountName() *string  {
  return s.AuditAccountName
}

func (s *EnableRightsSeparationRequest) GetAuditAccountPassword() *string  {
  return s.AuditAccountPassword
}

func (s *EnableRightsSeparationRequest) GetDBInstanceName() *string  {
  return s.DBInstanceName
}

func (s *EnableRightsSeparationRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableRightsSeparationRequest) GetSecurityAccountDescription() *string  {
  return s.SecurityAccountDescription
}

func (s *EnableRightsSeparationRequest) GetSecurityAccountName() *string  {
  return s.SecurityAccountName
}

func (s *EnableRightsSeparationRequest) GetSecurityAccountPassword() *string  {
  return s.SecurityAccountPassword
}

func (s *EnableRightsSeparationRequest) SetAuditAccountDescription(v string) *EnableRightsSeparationRequest {
  s.AuditAccountDescription = &v
  return s
}

func (s *EnableRightsSeparationRequest) SetAuditAccountName(v string) *EnableRightsSeparationRequest {
  s.AuditAccountName = &v
  return s
}

func (s *EnableRightsSeparationRequest) SetAuditAccountPassword(v string) *EnableRightsSeparationRequest {
  s.AuditAccountPassword = &v
  return s
}

func (s *EnableRightsSeparationRequest) SetDBInstanceName(v string) *EnableRightsSeparationRequest {
  s.DBInstanceName = &v
  return s
}

func (s *EnableRightsSeparationRequest) SetRegionId(v string) *EnableRightsSeparationRequest {
  s.RegionId = &v
  return s
}

func (s *EnableRightsSeparationRequest) SetSecurityAccountDescription(v string) *EnableRightsSeparationRequest {
  s.SecurityAccountDescription = &v
  return s
}

func (s *EnableRightsSeparationRequest) SetSecurityAccountName(v string) *EnableRightsSeparationRequest {
  s.SecurityAccountName = &v
  return s
}

func (s *EnableRightsSeparationRequest) SetSecurityAccountPassword(v string) *EnableRightsSeparationRequest {
  s.SecurityAccountPassword = &v
  return s
}

func (s *EnableRightsSeparationRequest) Validate() error {
  return dara.Validate(s)
}

