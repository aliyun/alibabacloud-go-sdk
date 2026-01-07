// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountUpdateSecurityMobileLoginStatusRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppName(v string) *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest
  GetAppName() *string 
  SetOrientedEcId(v string) *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest
  GetOrientedEcId() *string 
  SetOrientedLeId(v string) *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest
  GetOrientedLeId() *string 
  SetOrientedNbId(v string) *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest
  GetOrientedNbId() *string 
  SetPk(v string) *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest
  GetPk() *string 
  SetRequestId(v string) *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest
  GetRequestId() *string 
  SetStatus(v string) *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest
  GetStatus() *string 
}

type EnterpriseAccountUpdateSecurityMobileLoginStatusRequest struct {
  AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
  OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
  OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
  OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
  // This parameter is required.
  Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
  // This parameter is required.
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // This parameter is required.
  Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s EnterpriseAccountUpdateSecurityMobileLoginStatusRequest) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountUpdateSecurityMobileLoginStatusRequest) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest) GetAppName() *string  {
  return s.AppName
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest) GetOrientedEcId() *string  {
  return s.OrientedEcId
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest) GetOrientedLeId() *string  {
  return s.OrientedLeId
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest) GetOrientedNbId() *string  {
  return s.OrientedNbId
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest) GetPk() *string  {
  return s.Pk
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest) GetStatus() *string  {
  return s.Status
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest) SetAppName(v string) *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest {
  s.AppName = &v
  return s
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest) SetOrientedEcId(v string) *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest {
  s.OrientedEcId = &v
  return s
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest) SetOrientedLeId(v string) *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest {
  s.OrientedLeId = &v
  return s
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest) SetOrientedNbId(v string) *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest {
  s.OrientedNbId = &v
  return s
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest) SetPk(v string) *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest {
  s.Pk = &v
  return s
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest) SetRequestId(v string) *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest {
  s.RequestId = &v
  return s
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest) SetStatus(v string) *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest {
  s.Status = &v
  return s
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest) Validate() error {
  return dara.Validate(s)
}

