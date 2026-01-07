// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountUpdateSessionExpireTimeRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppName(v string) *EnterpriseAccountUpdateSessionExpireTimeRequest
  GetAppName() *string 
  SetOrientedEcId(v string) *EnterpriseAccountUpdateSessionExpireTimeRequest
  GetOrientedEcId() *string 
  SetOrientedLeId(v string) *EnterpriseAccountUpdateSessionExpireTimeRequest
  GetOrientedLeId() *string 
  SetOrientedNbId(v string) *EnterpriseAccountUpdateSessionExpireTimeRequest
  GetOrientedNbId() *string 
  SetPk(v string) *EnterpriseAccountUpdateSessionExpireTimeRequest
  GetPk() *string 
  SetRequestId(v string) *EnterpriseAccountUpdateSessionExpireTimeRequest
  GetRequestId() *string 
  SetSessionExpireTime(v int32) *EnterpriseAccountUpdateSessionExpireTimeRequest
  GetSessionExpireTime() *int32 
}

type EnterpriseAccountUpdateSessionExpireTimeRequest struct {
  AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
  OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
  OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
  OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
  // This parameter is required.
  Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
  // This parameter is required.
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // This parameter is required.
  SessionExpireTime *int32 `json:"SessionExpireTime,omitempty" xml:"SessionExpireTime,omitempty"`
}

func (s EnterpriseAccountUpdateSessionExpireTimeRequest) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountUpdateSessionExpireTimeRequest) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountUpdateSessionExpireTimeRequest) GetAppName() *string  {
  return s.AppName
}

func (s *EnterpriseAccountUpdateSessionExpireTimeRequest) GetOrientedEcId() *string  {
  return s.OrientedEcId
}

func (s *EnterpriseAccountUpdateSessionExpireTimeRequest) GetOrientedLeId() *string  {
  return s.OrientedLeId
}

func (s *EnterpriseAccountUpdateSessionExpireTimeRequest) GetOrientedNbId() *string  {
  return s.OrientedNbId
}

func (s *EnterpriseAccountUpdateSessionExpireTimeRequest) GetPk() *string  {
  return s.Pk
}

func (s *EnterpriseAccountUpdateSessionExpireTimeRequest) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseAccountUpdateSessionExpireTimeRequest) GetSessionExpireTime() *int32  {
  return s.SessionExpireTime
}

func (s *EnterpriseAccountUpdateSessionExpireTimeRequest) SetAppName(v string) *EnterpriseAccountUpdateSessionExpireTimeRequest {
  s.AppName = &v
  return s
}

func (s *EnterpriseAccountUpdateSessionExpireTimeRequest) SetOrientedEcId(v string) *EnterpriseAccountUpdateSessionExpireTimeRequest {
  s.OrientedEcId = &v
  return s
}

func (s *EnterpriseAccountUpdateSessionExpireTimeRequest) SetOrientedLeId(v string) *EnterpriseAccountUpdateSessionExpireTimeRequest {
  s.OrientedLeId = &v
  return s
}

func (s *EnterpriseAccountUpdateSessionExpireTimeRequest) SetOrientedNbId(v string) *EnterpriseAccountUpdateSessionExpireTimeRequest {
  s.OrientedNbId = &v
  return s
}

func (s *EnterpriseAccountUpdateSessionExpireTimeRequest) SetPk(v string) *EnterpriseAccountUpdateSessionExpireTimeRequest {
  s.Pk = &v
  return s
}

func (s *EnterpriseAccountUpdateSessionExpireTimeRequest) SetRequestId(v string) *EnterpriseAccountUpdateSessionExpireTimeRequest {
  s.RequestId = &v
  return s
}

func (s *EnterpriseAccountUpdateSessionExpireTimeRequest) SetSessionExpireTime(v int32) *EnterpriseAccountUpdateSessionExpireTimeRequest {
  s.SessionExpireTime = &v
  return s
}

func (s *EnterpriseAccountUpdateSessionExpireTimeRequest) Validate() error {
  return dara.Validate(s)
}

