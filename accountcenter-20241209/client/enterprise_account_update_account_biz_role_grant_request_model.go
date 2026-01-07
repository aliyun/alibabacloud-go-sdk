// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountUpdateAccountBizRoleGrantRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBizRoleCodeListJson(v string) *EnterpriseAccountUpdateAccountBizRoleGrantRequest
  GetBizRoleCodeListJson() *string 
  SetEncryptTicket(v string) *EnterpriseAccountUpdateAccountBizRoleGrantRequest
  GetEncryptTicket() *string 
  SetOrientedEcId(v string) *EnterpriseAccountUpdateAccountBizRoleGrantRequest
  GetOrientedEcId() *string 
  SetOrientedLeId(v string) *EnterpriseAccountUpdateAccountBizRoleGrantRequest
  GetOrientedLeId() *string 
  SetOrientedNbId(v string) *EnterpriseAccountUpdateAccountBizRoleGrantRequest
  GetOrientedNbId() *string 
  SetPk(v string) *EnterpriseAccountUpdateAccountBizRoleGrantRequest
  GetPk() *string 
}

type EnterpriseAccountUpdateAccountBizRoleGrantRequest struct {
  BizRoleCodeListJson *string `json:"BizRoleCodeListJson,omitempty" xml:"BizRoleCodeListJson,omitempty"`
  EncryptTicket *string `json:"EncryptTicket,omitempty" xml:"EncryptTicket,omitempty"`
  OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
  OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
  OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
  Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s EnterpriseAccountUpdateAccountBizRoleGrantRequest) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountUpdateAccountBizRoleGrantRequest) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantRequest) GetBizRoleCodeListJson() *string  {
  return s.BizRoleCodeListJson
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantRequest) GetEncryptTicket() *string  {
  return s.EncryptTicket
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantRequest) GetOrientedEcId() *string  {
  return s.OrientedEcId
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantRequest) GetOrientedLeId() *string  {
  return s.OrientedLeId
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantRequest) GetOrientedNbId() *string  {
  return s.OrientedNbId
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantRequest) GetPk() *string  {
  return s.Pk
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantRequest) SetBizRoleCodeListJson(v string) *EnterpriseAccountUpdateAccountBizRoleGrantRequest {
  s.BizRoleCodeListJson = &v
  return s
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantRequest) SetEncryptTicket(v string) *EnterpriseAccountUpdateAccountBizRoleGrantRequest {
  s.EncryptTicket = &v
  return s
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantRequest) SetOrientedEcId(v string) *EnterpriseAccountUpdateAccountBizRoleGrantRequest {
  s.OrientedEcId = &v
  return s
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantRequest) SetOrientedLeId(v string) *EnterpriseAccountUpdateAccountBizRoleGrantRequest {
  s.OrientedLeId = &v
  return s
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantRequest) SetOrientedNbId(v string) *EnterpriseAccountUpdateAccountBizRoleGrantRequest {
  s.OrientedNbId = &v
  return s
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantRequest) SetPk(v string) *EnterpriseAccountUpdateAccountBizRoleGrantRequest {
  s.Pk = &v
  return s
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantRequest) Validate() error {
  return dara.Validate(s)
}

