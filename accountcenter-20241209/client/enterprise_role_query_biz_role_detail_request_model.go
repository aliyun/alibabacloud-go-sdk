// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseRoleQueryBizRoleDetailRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBizRoleCode(v string) *EnterpriseRoleQueryBizRoleDetailRequest
  GetBizRoleCode() *string 
  SetEncryptTicket(v string) *EnterpriseRoleQueryBizRoleDetailRequest
  GetEncryptTicket() *string 
  SetOrientedEcId(v string) *EnterpriseRoleQueryBizRoleDetailRequest
  GetOrientedEcId() *string 
  SetOrientedLeId(v string) *EnterpriseRoleQueryBizRoleDetailRequest
  GetOrientedLeId() *string 
  SetOrientedNbId(v string) *EnterpriseRoleQueryBizRoleDetailRequest
  GetOrientedNbId() *string 
}

type EnterpriseRoleQueryBizRoleDetailRequest struct {
  BizRoleCode *string `json:"BizRoleCode,omitempty" xml:"BizRoleCode,omitempty"`
  EncryptTicket *string `json:"EncryptTicket,omitempty" xml:"EncryptTicket,omitempty"`
  OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
  OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
  OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
}

func (s EnterpriseRoleQueryBizRoleDetailRequest) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseRoleQueryBizRoleDetailRequest) GoString() string {
  return s.String()
}

func (s *EnterpriseRoleQueryBizRoleDetailRequest) GetBizRoleCode() *string  {
  return s.BizRoleCode
}

func (s *EnterpriseRoleQueryBizRoleDetailRequest) GetEncryptTicket() *string  {
  return s.EncryptTicket
}

func (s *EnterpriseRoleQueryBizRoleDetailRequest) GetOrientedEcId() *string  {
  return s.OrientedEcId
}

func (s *EnterpriseRoleQueryBizRoleDetailRequest) GetOrientedLeId() *string  {
  return s.OrientedLeId
}

func (s *EnterpriseRoleQueryBizRoleDetailRequest) GetOrientedNbId() *string  {
  return s.OrientedNbId
}

func (s *EnterpriseRoleQueryBizRoleDetailRequest) SetBizRoleCode(v string) *EnterpriseRoleQueryBizRoleDetailRequest {
  s.BizRoleCode = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleDetailRequest) SetEncryptTicket(v string) *EnterpriseRoleQueryBizRoleDetailRequest {
  s.EncryptTicket = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleDetailRequest) SetOrientedEcId(v string) *EnterpriseRoleQueryBizRoleDetailRequest {
  s.OrientedEcId = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleDetailRequest) SetOrientedLeId(v string) *EnterpriseRoleQueryBizRoleDetailRequest {
  s.OrientedLeId = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleDetailRequest) SetOrientedNbId(v string) *EnterpriseRoleQueryBizRoleDetailRequest {
  s.OrientedNbId = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleDetailRequest) Validate() error {
  return dara.Validate(s)
}

