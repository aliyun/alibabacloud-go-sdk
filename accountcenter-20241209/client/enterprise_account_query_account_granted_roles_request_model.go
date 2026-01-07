// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountQueryAccountGrantedRolesRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppName(v string) *EnterpriseAccountQueryAccountGrantedRolesRequest
  GetAppName() *string 
  SetIsOpenApi(v bool) *EnterpriseAccountQueryAccountGrantedRolesRequest
  GetIsOpenApi() *bool 
  SetOrientedEcId(v string) *EnterpriseAccountQueryAccountGrantedRolesRequest
  GetOrientedEcId() *string 
  SetOrientedLeId(v string) *EnterpriseAccountQueryAccountGrantedRolesRequest
  GetOrientedLeId() *string 
  SetOrientedNbId(v string) *EnterpriseAccountQueryAccountGrantedRolesRequest
  GetOrientedNbId() *string 
  SetPk(v string) *EnterpriseAccountQueryAccountGrantedRolesRequest
  GetPk() *string 
  SetShowCompleteInfo(v bool) *EnterpriseAccountQueryAccountGrantedRolesRequest
  GetShowCompleteInfo() *bool 
}

type EnterpriseAccountQueryAccountGrantedRolesRequest struct {
  AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
  IsOpenApi *bool `json:"IsOpenApi,omitempty" xml:"IsOpenApi,omitempty"`
  OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
  OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
  OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
  Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
  ShowCompleteInfo *bool `json:"ShowCompleteInfo,omitempty" xml:"ShowCompleteInfo,omitempty"`
}

func (s EnterpriseAccountQueryAccountGrantedRolesRequest) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountQueryAccountGrantedRolesRequest) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountQueryAccountGrantedRolesRequest) GetAppName() *string  {
  return s.AppName
}

func (s *EnterpriseAccountQueryAccountGrantedRolesRequest) GetIsOpenApi() *bool  {
  return s.IsOpenApi
}

func (s *EnterpriseAccountQueryAccountGrantedRolesRequest) GetOrientedEcId() *string  {
  return s.OrientedEcId
}

func (s *EnterpriseAccountQueryAccountGrantedRolesRequest) GetOrientedLeId() *string  {
  return s.OrientedLeId
}

func (s *EnterpriseAccountQueryAccountGrantedRolesRequest) GetOrientedNbId() *string  {
  return s.OrientedNbId
}

func (s *EnterpriseAccountQueryAccountGrantedRolesRequest) GetPk() *string  {
  return s.Pk
}

func (s *EnterpriseAccountQueryAccountGrantedRolesRequest) GetShowCompleteInfo() *bool  {
  return s.ShowCompleteInfo
}

func (s *EnterpriseAccountQueryAccountGrantedRolesRequest) SetAppName(v string) *EnterpriseAccountQueryAccountGrantedRolesRequest {
  s.AppName = &v
  return s
}

func (s *EnterpriseAccountQueryAccountGrantedRolesRequest) SetIsOpenApi(v bool) *EnterpriseAccountQueryAccountGrantedRolesRequest {
  s.IsOpenApi = &v
  return s
}

func (s *EnterpriseAccountQueryAccountGrantedRolesRequest) SetOrientedEcId(v string) *EnterpriseAccountQueryAccountGrantedRolesRequest {
  s.OrientedEcId = &v
  return s
}

func (s *EnterpriseAccountQueryAccountGrantedRolesRequest) SetOrientedLeId(v string) *EnterpriseAccountQueryAccountGrantedRolesRequest {
  s.OrientedLeId = &v
  return s
}

func (s *EnterpriseAccountQueryAccountGrantedRolesRequest) SetOrientedNbId(v string) *EnterpriseAccountQueryAccountGrantedRolesRequest {
  s.OrientedNbId = &v
  return s
}

func (s *EnterpriseAccountQueryAccountGrantedRolesRequest) SetPk(v string) *EnterpriseAccountQueryAccountGrantedRolesRequest {
  s.Pk = &v
  return s
}

func (s *EnterpriseAccountQueryAccountGrantedRolesRequest) SetShowCompleteInfo(v bool) *EnterpriseAccountQueryAccountGrantedRolesRequest {
  s.ShowCompleteInfo = &v
  return s
}

func (s *EnterpriseAccountQueryAccountGrantedRolesRequest) Validate() error {
  return dara.Validate(s)
}

