// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseRoleDeleteBizRoleRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBizRoleCode(v string) *EnterpriseRoleDeleteBizRoleRequest
  GetBizRoleCode() *string 
  SetEncryptTicket(v string) *EnterpriseRoleDeleteBizRoleRequest
  GetEncryptTicket() *string 
  SetOrientedEcId(v string) *EnterpriseRoleDeleteBizRoleRequest
  GetOrientedEcId() *string 
  SetOrientedLeId(v string) *EnterpriseRoleDeleteBizRoleRequest
  GetOrientedLeId() *string 
  SetOrientedNbId(v string) *EnterpriseRoleDeleteBizRoleRequest
  GetOrientedNbId() *string 
}

type EnterpriseRoleDeleteBizRoleRequest struct {
  BizRoleCode *string `json:"BizRoleCode,omitempty" xml:"BizRoleCode,omitempty"`
  EncryptTicket *string `json:"EncryptTicket,omitempty" xml:"EncryptTicket,omitempty"`
  OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
  OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
  OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
}

func (s EnterpriseRoleDeleteBizRoleRequest) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseRoleDeleteBizRoleRequest) GoString() string {
  return s.String()
}

func (s *EnterpriseRoleDeleteBizRoleRequest) GetBizRoleCode() *string  {
  return s.BizRoleCode
}

func (s *EnterpriseRoleDeleteBizRoleRequest) GetEncryptTicket() *string  {
  return s.EncryptTicket
}

func (s *EnterpriseRoleDeleteBizRoleRequest) GetOrientedEcId() *string  {
  return s.OrientedEcId
}

func (s *EnterpriseRoleDeleteBizRoleRequest) GetOrientedLeId() *string  {
  return s.OrientedLeId
}

func (s *EnterpriseRoleDeleteBizRoleRequest) GetOrientedNbId() *string  {
  return s.OrientedNbId
}

func (s *EnterpriseRoleDeleteBizRoleRequest) SetBizRoleCode(v string) *EnterpriseRoleDeleteBizRoleRequest {
  s.BizRoleCode = &v
  return s
}

func (s *EnterpriseRoleDeleteBizRoleRequest) SetEncryptTicket(v string) *EnterpriseRoleDeleteBizRoleRequest {
  s.EncryptTicket = &v
  return s
}

func (s *EnterpriseRoleDeleteBizRoleRequest) SetOrientedEcId(v string) *EnterpriseRoleDeleteBizRoleRequest {
  s.OrientedEcId = &v
  return s
}

func (s *EnterpriseRoleDeleteBizRoleRequest) SetOrientedLeId(v string) *EnterpriseRoleDeleteBizRoleRequest {
  s.OrientedLeId = &v
  return s
}

func (s *EnterpriseRoleDeleteBizRoleRequest) SetOrientedNbId(v string) *EnterpriseRoleDeleteBizRoleRequest {
  s.OrientedNbId = &v
  return s
}

func (s *EnterpriseRoleDeleteBizRoleRequest) Validate() error {
  return dara.Validate(s)
}

