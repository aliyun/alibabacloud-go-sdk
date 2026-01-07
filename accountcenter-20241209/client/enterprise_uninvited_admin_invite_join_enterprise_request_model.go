// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseUninvitedAdminInviteJoinEnterpriseRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEcId(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest
  GetEcId() *string 
  SetEncryptTicket(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest
  GetEncryptTicket() *string 
  SetInviteePk(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest
  GetInviteePk() *string 
  SetLeId(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest
  GetLeId() *string 
  SetNbId(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest
  GetNbId() *string 
  SetRemark(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest
  GetRemark() *string 
}

type EnterpriseUninvitedAdminInviteJoinEnterpriseRequest struct {
  EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
  EncryptTicket *string `json:"EncryptTicket,omitempty" xml:"EncryptTicket,omitempty"`
  InviteePk *string `json:"InviteePk,omitempty" xml:"InviteePk,omitempty"`
  LeId *string `json:"LeId,omitempty" xml:"LeId,omitempty"`
  NbId *string `json:"NbId,omitempty" xml:"NbId,omitempty"`
  Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s EnterpriseUninvitedAdminInviteJoinEnterpriseRequest) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseUninvitedAdminInviteJoinEnterpriseRequest) GoString() string {
  return s.String()
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest) GetEcId() *string  {
  return s.EcId
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest) GetEncryptTicket() *string  {
  return s.EncryptTicket
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest) GetInviteePk() *string  {
  return s.InviteePk
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest) GetLeId() *string  {
  return s.LeId
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest) GetNbId() *string  {
  return s.NbId
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest) GetRemark() *string  {
  return s.Remark
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest) SetEcId(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest {
  s.EcId = &v
  return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest) SetEncryptTicket(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest {
  s.EncryptTicket = &v
  return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest) SetInviteePk(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest {
  s.InviteePk = &v
  return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest) SetLeId(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest {
  s.LeId = &v
  return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest) SetNbId(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest {
  s.NbId = &v
  return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest) SetRemark(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest {
  s.Remark = &v
  return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest) Validate() error {
  return dara.Validate(s)
}

