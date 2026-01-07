// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountChangeLoginPasswordRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppName(v string) *EnterpriseAccountChangeLoginPasswordRequest
  GetAppName() *string 
  SetEncryptPassword(v string) *EnterpriseAccountChangeLoginPasswordRequest
  GetEncryptPassword() *string 
  SetOrientedEcId(v string) *EnterpriseAccountChangeLoginPasswordRequest
  GetOrientedEcId() *string 
  SetOrientedLeId(v string) *EnterpriseAccountChangeLoginPasswordRequest
  GetOrientedLeId() *string 
  SetOrientedNbId(v string) *EnterpriseAccountChangeLoginPasswordRequest
  GetOrientedNbId() *string 
  SetPk(v string) *EnterpriseAccountChangeLoginPasswordRequest
  GetPk() *string 
  SetRequestId(v string) *EnterpriseAccountChangeLoginPasswordRequest
  GetRequestId() *string 
}

type EnterpriseAccountChangeLoginPasswordRequest struct {
  AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
  // This parameter is required.
  EncryptPassword *string `json:"EncryptPassword,omitempty" xml:"EncryptPassword,omitempty"`
  OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
  OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
  OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
  // This parameter is required.
  Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
  // This parameter is required.
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnterpriseAccountChangeLoginPasswordRequest) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountChangeLoginPasswordRequest) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountChangeLoginPasswordRequest) GetAppName() *string  {
  return s.AppName
}

func (s *EnterpriseAccountChangeLoginPasswordRequest) GetEncryptPassword() *string  {
  return s.EncryptPassword
}

func (s *EnterpriseAccountChangeLoginPasswordRequest) GetOrientedEcId() *string  {
  return s.OrientedEcId
}

func (s *EnterpriseAccountChangeLoginPasswordRequest) GetOrientedLeId() *string  {
  return s.OrientedLeId
}

func (s *EnterpriseAccountChangeLoginPasswordRequest) GetOrientedNbId() *string  {
  return s.OrientedNbId
}

func (s *EnterpriseAccountChangeLoginPasswordRequest) GetPk() *string  {
  return s.Pk
}

func (s *EnterpriseAccountChangeLoginPasswordRequest) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseAccountChangeLoginPasswordRequest) SetAppName(v string) *EnterpriseAccountChangeLoginPasswordRequest {
  s.AppName = &v
  return s
}

func (s *EnterpriseAccountChangeLoginPasswordRequest) SetEncryptPassword(v string) *EnterpriseAccountChangeLoginPasswordRequest {
  s.EncryptPassword = &v
  return s
}

func (s *EnterpriseAccountChangeLoginPasswordRequest) SetOrientedEcId(v string) *EnterpriseAccountChangeLoginPasswordRequest {
  s.OrientedEcId = &v
  return s
}

func (s *EnterpriseAccountChangeLoginPasswordRequest) SetOrientedLeId(v string) *EnterpriseAccountChangeLoginPasswordRequest {
  s.OrientedLeId = &v
  return s
}

func (s *EnterpriseAccountChangeLoginPasswordRequest) SetOrientedNbId(v string) *EnterpriseAccountChangeLoginPasswordRequest {
  s.OrientedNbId = &v
  return s
}

func (s *EnterpriseAccountChangeLoginPasswordRequest) SetPk(v string) *EnterpriseAccountChangeLoginPasswordRequest {
  s.Pk = &v
  return s
}

func (s *EnterpriseAccountChangeLoginPasswordRequest) SetRequestId(v string) *EnterpriseAccountChangeLoginPasswordRequest {
  s.RequestId = &v
  return s
}

func (s *EnterpriseAccountChangeLoginPasswordRequest) Validate() error {
  return dara.Validate(s)
}

