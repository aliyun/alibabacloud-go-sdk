// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountSeparateEaRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEncryptTicket(v string) *EnterpriseAccountSeparateEaRequest
  GetEncryptTicket() *string 
  SetOrientedEcId(v string) *EnterpriseAccountSeparateEaRequest
  GetOrientedEcId() *string 
  SetOrientedLeId(v string) *EnterpriseAccountSeparateEaRequest
  GetOrientedLeId() *string 
  SetOrientedNbId(v string) *EnterpriseAccountSeparateEaRequest
  GetOrientedNbId() *string 
  SetPk(v string) *EnterpriseAccountSeparateEaRequest
  GetPk() *string 
}

type EnterpriseAccountSeparateEaRequest struct {
  EncryptTicket *string `json:"EncryptTicket,omitempty" xml:"EncryptTicket,omitempty"`
  OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
  OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
  OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
  Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s EnterpriseAccountSeparateEaRequest) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountSeparateEaRequest) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountSeparateEaRequest) GetEncryptTicket() *string  {
  return s.EncryptTicket
}

func (s *EnterpriseAccountSeparateEaRequest) GetOrientedEcId() *string  {
  return s.OrientedEcId
}

func (s *EnterpriseAccountSeparateEaRequest) GetOrientedLeId() *string  {
  return s.OrientedLeId
}

func (s *EnterpriseAccountSeparateEaRequest) GetOrientedNbId() *string  {
  return s.OrientedNbId
}

func (s *EnterpriseAccountSeparateEaRequest) GetPk() *string  {
  return s.Pk
}

func (s *EnterpriseAccountSeparateEaRequest) SetEncryptTicket(v string) *EnterpriseAccountSeparateEaRequest {
  s.EncryptTicket = &v
  return s
}

func (s *EnterpriseAccountSeparateEaRequest) SetOrientedEcId(v string) *EnterpriseAccountSeparateEaRequest {
  s.OrientedEcId = &v
  return s
}

func (s *EnterpriseAccountSeparateEaRequest) SetOrientedLeId(v string) *EnterpriseAccountSeparateEaRequest {
  s.OrientedLeId = &v
  return s
}

func (s *EnterpriseAccountSeparateEaRequest) SetOrientedNbId(v string) *EnterpriseAccountSeparateEaRequest {
  s.OrientedNbId = &v
  return s
}

func (s *EnterpriseAccountSeparateEaRequest) SetPk(v string) *EnterpriseAccountSeparateEaRequest {
  s.Pk = &v
  return s
}

func (s *EnterpriseAccountSeparateEaRequest) Validate() error {
  return dara.Validate(s)
}

