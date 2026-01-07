// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountQueryLoginSettingsRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppName(v string) *EnterpriseAccountQueryLoginSettingsRequest
  GetAppName() *string 
  SetIsOpenApi(v bool) *EnterpriseAccountQueryLoginSettingsRequest
  GetIsOpenApi() *bool 
  SetOrientedEcId(v string) *EnterpriseAccountQueryLoginSettingsRequest
  GetOrientedEcId() *string 
  SetOrientedLeId(v string) *EnterpriseAccountQueryLoginSettingsRequest
  GetOrientedLeId() *string 
  SetOrientedNbId(v string) *EnterpriseAccountQueryLoginSettingsRequest
  GetOrientedNbId() *string 
  SetPk(v string) *EnterpriseAccountQueryLoginSettingsRequest
  GetPk() *string 
  SetShowCompleteInfo(v bool) *EnterpriseAccountQueryLoginSettingsRequest
  GetShowCompleteInfo() *bool 
}

type EnterpriseAccountQueryLoginSettingsRequest struct {
  AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
  IsOpenApi *bool `json:"IsOpenApi,omitempty" xml:"IsOpenApi,omitempty"`
  OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
  OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
  OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
  Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
  ShowCompleteInfo *bool `json:"ShowCompleteInfo,omitempty" xml:"ShowCompleteInfo,omitempty"`
}

func (s EnterpriseAccountQueryLoginSettingsRequest) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountQueryLoginSettingsRequest) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountQueryLoginSettingsRequest) GetAppName() *string  {
  return s.AppName
}

func (s *EnterpriseAccountQueryLoginSettingsRequest) GetIsOpenApi() *bool  {
  return s.IsOpenApi
}

func (s *EnterpriseAccountQueryLoginSettingsRequest) GetOrientedEcId() *string  {
  return s.OrientedEcId
}

func (s *EnterpriseAccountQueryLoginSettingsRequest) GetOrientedLeId() *string  {
  return s.OrientedLeId
}

func (s *EnterpriseAccountQueryLoginSettingsRequest) GetOrientedNbId() *string  {
  return s.OrientedNbId
}

func (s *EnterpriseAccountQueryLoginSettingsRequest) GetPk() *string  {
  return s.Pk
}

func (s *EnterpriseAccountQueryLoginSettingsRequest) GetShowCompleteInfo() *bool  {
  return s.ShowCompleteInfo
}

func (s *EnterpriseAccountQueryLoginSettingsRequest) SetAppName(v string) *EnterpriseAccountQueryLoginSettingsRequest {
  s.AppName = &v
  return s
}

func (s *EnterpriseAccountQueryLoginSettingsRequest) SetIsOpenApi(v bool) *EnterpriseAccountQueryLoginSettingsRequest {
  s.IsOpenApi = &v
  return s
}

func (s *EnterpriseAccountQueryLoginSettingsRequest) SetOrientedEcId(v string) *EnterpriseAccountQueryLoginSettingsRequest {
  s.OrientedEcId = &v
  return s
}

func (s *EnterpriseAccountQueryLoginSettingsRequest) SetOrientedLeId(v string) *EnterpriseAccountQueryLoginSettingsRequest {
  s.OrientedLeId = &v
  return s
}

func (s *EnterpriseAccountQueryLoginSettingsRequest) SetOrientedNbId(v string) *EnterpriseAccountQueryLoginSettingsRequest {
  s.OrientedNbId = &v
  return s
}

func (s *EnterpriseAccountQueryLoginSettingsRequest) SetPk(v string) *EnterpriseAccountQueryLoginSettingsRequest {
  s.Pk = &v
  return s
}

func (s *EnterpriseAccountQueryLoginSettingsRequest) SetShowCompleteInfo(v bool) *EnterpriseAccountQueryLoginSettingsRequest {
  s.ShowCompleteInfo = &v
  return s
}

func (s *EnterpriseAccountQueryLoginSettingsRequest) Validate() error {
  return dara.Validate(s)
}

