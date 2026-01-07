// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountUpdateAccountAliasRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAlias(v string) *EnterpriseAccountUpdateAccountAliasRequest
  GetAlias() *string 
  SetAppName(v string) *EnterpriseAccountUpdateAccountAliasRequest
  GetAppName() *string 
  SetEncryptTicket(v string) *EnterpriseAccountUpdateAccountAliasRequest
  GetEncryptTicket() *string 
  SetOrientedEcId(v string) *EnterpriseAccountUpdateAccountAliasRequest
  GetOrientedEcId() *string 
  SetOrientedLeId(v string) *EnterpriseAccountUpdateAccountAliasRequest
  GetOrientedLeId() *string 
  SetOrientedNbId(v string) *EnterpriseAccountUpdateAccountAliasRequest
  GetOrientedNbId() *string 
  SetPk(v string) *EnterpriseAccountUpdateAccountAliasRequest
  GetPk() *string 
  SetRequestId(v string) *EnterpriseAccountUpdateAccountAliasRequest
  GetRequestId() *string 
}

type EnterpriseAccountUpdateAccountAliasRequest struct {
  Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
  AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
  EncryptTicket *string `json:"EncryptTicket,omitempty" xml:"EncryptTicket,omitempty"`
  OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
  OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
  OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
  Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
  // This parameter is required.
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnterpriseAccountUpdateAccountAliasRequest) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountUpdateAccountAliasRequest) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountUpdateAccountAliasRequest) GetAlias() *string  {
  return s.Alias
}

func (s *EnterpriseAccountUpdateAccountAliasRequest) GetAppName() *string  {
  return s.AppName
}

func (s *EnterpriseAccountUpdateAccountAliasRequest) GetEncryptTicket() *string  {
  return s.EncryptTicket
}

func (s *EnterpriseAccountUpdateAccountAliasRequest) GetOrientedEcId() *string  {
  return s.OrientedEcId
}

func (s *EnterpriseAccountUpdateAccountAliasRequest) GetOrientedLeId() *string  {
  return s.OrientedLeId
}

func (s *EnterpriseAccountUpdateAccountAliasRequest) GetOrientedNbId() *string  {
  return s.OrientedNbId
}

func (s *EnterpriseAccountUpdateAccountAliasRequest) GetPk() *string  {
  return s.Pk
}

func (s *EnterpriseAccountUpdateAccountAliasRequest) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseAccountUpdateAccountAliasRequest) SetAlias(v string) *EnterpriseAccountUpdateAccountAliasRequest {
  s.Alias = &v
  return s
}

func (s *EnterpriseAccountUpdateAccountAliasRequest) SetAppName(v string) *EnterpriseAccountUpdateAccountAliasRequest {
  s.AppName = &v
  return s
}

func (s *EnterpriseAccountUpdateAccountAliasRequest) SetEncryptTicket(v string) *EnterpriseAccountUpdateAccountAliasRequest {
  s.EncryptTicket = &v
  return s
}

func (s *EnterpriseAccountUpdateAccountAliasRequest) SetOrientedEcId(v string) *EnterpriseAccountUpdateAccountAliasRequest {
  s.OrientedEcId = &v
  return s
}

func (s *EnterpriseAccountUpdateAccountAliasRequest) SetOrientedLeId(v string) *EnterpriseAccountUpdateAccountAliasRequest {
  s.OrientedLeId = &v
  return s
}

func (s *EnterpriseAccountUpdateAccountAliasRequest) SetOrientedNbId(v string) *EnterpriseAccountUpdateAccountAliasRequest {
  s.OrientedNbId = &v
  return s
}

func (s *EnterpriseAccountUpdateAccountAliasRequest) SetPk(v string) *EnterpriseAccountUpdateAccountAliasRequest {
  s.Pk = &v
  return s
}

func (s *EnterpriseAccountUpdateAccountAliasRequest) SetRequestId(v string) *EnterpriseAccountUpdateAccountAliasRequest {
  s.RequestId = &v
  return s
}

func (s *EnterpriseAccountUpdateAccountAliasRequest) Validate() error {
  return dara.Validate(s)
}

