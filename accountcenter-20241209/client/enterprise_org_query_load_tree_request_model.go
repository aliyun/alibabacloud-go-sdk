// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseOrgQueryLoadTreeRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEncryptTicket(v string) *EnterpriseOrgQueryLoadTreeRequest
  GetEncryptTicket() *string 
  SetLoadOrgOnly(v bool) *EnterpriseOrgQueryLoadTreeRequest
  GetLoadOrgOnly() *bool 
  SetOrientedEcId(v string) *EnterpriseOrgQueryLoadTreeRequest
  GetOrientedEcId() *string 
  SetOrientedLeId(v string) *EnterpriseOrgQueryLoadTreeRequest
  GetOrientedLeId() *string 
  SetOrientedNbId(v string) *EnterpriseOrgQueryLoadTreeRequest
  GetOrientedNbId() *string 
  SetRequestId(v string) *EnterpriseOrgQueryLoadTreeRequest
  GetRequestId() *string 
}

type EnterpriseOrgQueryLoadTreeRequest struct {
  EncryptTicket *string `json:"EncryptTicket,omitempty" xml:"EncryptTicket,omitempty"`
  // example:
  // 
  // true
  LoadOrgOnly *bool `json:"LoadOrgOnly,omitempty" xml:"LoadOrgOnly,omitempty"`
  OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
  OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
  OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // CF20ED94-D406-512F-9798-4E1F65720BF6
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnterpriseOrgQueryLoadTreeRequest) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseOrgQueryLoadTreeRequest) GoString() string {
  return s.String()
}

func (s *EnterpriseOrgQueryLoadTreeRequest) GetEncryptTicket() *string  {
  return s.EncryptTicket
}

func (s *EnterpriseOrgQueryLoadTreeRequest) GetLoadOrgOnly() *bool  {
  return s.LoadOrgOnly
}

func (s *EnterpriseOrgQueryLoadTreeRequest) GetOrientedEcId() *string  {
  return s.OrientedEcId
}

func (s *EnterpriseOrgQueryLoadTreeRequest) GetOrientedLeId() *string  {
  return s.OrientedLeId
}

func (s *EnterpriseOrgQueryLoadTreeRequest) GetOrientedNbId() *string  {
  return s.OrientedNbId
}

func (s *EnterpriseOrgQueryLoadTreeRequest) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseOrgQueryLoadTreeRequest) SetEncryptTicket(v string) *EnterpriseOrgQueryLoadTreeRequest {
  s.EncryptTicket = &v
  return s
}

func (s *EnterpriseOrgQueryLoadTreeRequest) SetLoadOrgOnly(v bool) *EnterpriseOrgQueryLoadTreeRequest {
  s.LoadOrgOnly = &v
  return s
}

func (s *EnterpriseOrgQueryLoadTreeRequest) SetOrientedEcId(v string) *EnterpriseOrgQueryLoadTreeRequest {
  s.OrientedEcId = &v
  return s
}

func (s *EnterpriseOrgQueryLoadTreeRequest) SetOrientedLeId(v string) *EnterpriseOrgQueryLoadTreeRequest {
  s.OrientedLeId = &v
  return s
}

func (s *EnterpriseOrgQueryLoadTreeRequest) SetOrientedNbId(v string) *EnterpriseOrgQueryLoadTreeRequest {
  s.OrientedNbId = &v
  return s
}

func (s *EnterpriseOrgQueryLoadTreeRequest) SetRequestId(v string) *EnterpriseOrgQueryLoadTreeRequest {
  s.RequestId = &v
  return s
}

func (s *EnterpriseOrgQueryLoadTreeRequest) Validate() error {
  return dara.Validate(s)
}

