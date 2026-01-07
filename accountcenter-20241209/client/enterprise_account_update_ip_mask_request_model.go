// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountUpdateIpMaskRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppName(v string) *EnterpriseAccountUpdateIpMaskRequest
  GetAppName() *string 
  SetIpMasksJson(v string) *EnterpriseAccountUpdateIpMaskRequest
  GetIpMasksJson() *string 
  SetOrientedEcId(v string) *EnterpriseAccountUpdateIpMaskRequest
  GetOrientedEcId() *string 
  SetOrientedLeId(v string) *EnterpriseAccountUpdateIpMaskRequest
  GetOrientedLeId() *string 
  SetOrientedNbId(v string) *EnterpriseAccountUpdateIpMaskRequest
  GetOrientedNbId() *string 
  SetPk(v string) *EnterpriseAccountUpdateIpMaskRequest
  GetPk() *string 
  SetRequestId(v string) *EnterpriseAccountUpdateIpMaskRequest
  GetRequestId() *string 
  SetStatus(v string) *EnterpriseAccountUpdateIpMaskRequest
  GetStatus() *string 
}

type EnterpriseAccountUpdateIpMaskRequest struct {
  AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
  // This parameter is required.
  IpMasksJson *string `json:"IpMasksJson,omitempty" xml:"IpMasksJson,omitempty"`
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

func (s EnterpriseAccountUpdateIpMaskRequest) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountUpdateIpMaskRequest) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountUpdateIpMaskRequest) GetAppName() *string  {
  return s.AppName
}

func (s *EnterpriseAccountUpdateIpMaskRequest) GetIpMasksJson() *string  {
  return s.IpMasksJson
}

func (s *EnterpriseAccountUpdateIpMaskRequest) GetOrientedEcId() *string  {
  return s.OrientedEcId
}

func (s *EnterpriseAccountUpdateIpMaskRequest) GetOrientedLeId() *string  {
  return s.OrientedLeId
}

func (s *EnterpriseAccountUpdateIpMaskRequest) GetOrientedNbId() *string  {
  return s.OrientedNbId
}

func (s *EnterpriseAccountUpdateIpMaskRequest) GetPk() *string  {
  return s.Pk
}

func (s *EnterpriseAccountUpdateIpMaskRequest) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseAccountUpdateIpMaskRequest) GetStatus() *string  {
  return s.Status
}

func (s *EnterpriseAccountUpdateIpMaskRequest) SetAppName(v string) *EnterpriseAccountUpdateIpMaskRequest {
  s.AppName = &v
  return s
}

func (s *EnterpriseAccountUpdateIpMaskRequest) SetIpMasksJson(v string) *EnterpriseAccountUpdateIpMaskRequest {
  s.IpMasksJson = &v
  return s
}

func (s *EnterpriseAccountUpdateIpMaskRequest) SetOrientedEcId(v string) *EnterpriseAccountUpdateIpMaskRequest {
  s.OrientedEcId = &v
  return s
}

func (s *EnterpriseAccountUpdateIpMaskRequest) SetOrientedLeId(v string) *EnterpriseAccountUpdateIpMaskRequest {
  s.OrientedLeId = &v
  return s
}

func (s *EnterpriseAccountUpdateIpMaskRequest) SetOrientedNbId(v string) *EnterpriseAccountUpdateIpMaskRequest {
  s.OrientedNbId = &v
  return s
}

func (s *EnterpriseAccountUpdateIpMaskRequest) SetPk(v string) *EnterpriseAccountUpdateIpMaskRequest {
  s.Pk = &v
  return s
}

func (s *EnterpriseAccountUpdateIpMaskRequest) SetRequestId(v string) *EnterpriseAccountUpdateIpMaskRequest {
  s.RequestId = &v
  return s
}

func (s *EnterpriseAccountUpdateIpMaskRequest) SetStatus(v string) *EnterpriseAccountUpdateIpMaskRequest {
  s.Status = &v
  return s
}

func (s *EnterpriseAccountUpdateIpMaskRequest) Validate() error {
  return dara.Validate(s)
}

