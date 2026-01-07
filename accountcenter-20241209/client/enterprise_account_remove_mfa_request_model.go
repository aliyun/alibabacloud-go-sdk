// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountRemoveMfaRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppName(v string) *EnterpriseAccountRemoveMfaRequest
  GetAppName() *string 
  SetOrientedEcId(v string) *EnterpriseAccountRemoveMfaRequest
  GetOrientedEcId() *string 
  SetOrientedLeId(v string) *EnterpriseAccountRemoveMfaRequest
  GetOrientedLeId() *string 
  SetOrientedNbId(v string) *EnterpriseAccountRemoveMfaRequest
  GetOrientedNbId() *string 
  SetPk(v string) *EnterpriseAccountRemoveMfaRequest
  GetPk() *string 
  SetRequestId(v string) *EnterpriseAccountRemoveMfaRequest
  GetRequestId() *string 
}

type EnterpriseAccountRemoveMfaRequest struct {
  AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
  OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
  OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
  OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
  // This parameter is required.
  Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
  // This parameter is required.
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnterpriseAccountRemoveMfaRequest) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountRemoveMfaRequest) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountRemoveMfaRequest) GetAppName() *string  {
  return s.AppName
}

func (s *EnterpriseAccountRemoveMfaRequest) GetOrientedEcId() *string  {
  return s.OrientedEcId
}

func (s *EnterpriseAccountRemoveMfaRequest) GetOrientedLeId() *string  {
  return s.OrientedLeId
}

func (s *EnterpriseAccountRemoveMfaRequest) GetOrientedNbId() *string  {
  return s.OrientedNbId
}

func (s *EnterpriseAccountRemoveMfaRequest) GetPk() *string  {
  return s.Pk
}

func (s *EnterpriseAccountRemoveMfaRequest) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseAccountRemoveMfaRequest) SetAppName(v string) *EnterpriseAccountRemoveMfaRequest {
  s.AppName = &v
  return s
}

func (s *EnterpriseAccountRemoveMfaRequest) SetOrientedEcId(v string) *EnterpriseAccountRemoveMfaRequest {
  s.OrientedEcId = &v
  return s
}

func (s *EnterpriseAccountRemoveMfaRequest) SetOrientedLeId(v string) *EnterpriseAccountRemoveMfaRequest {
  s.OrientedLeId = &v
  return s
}

func (s *EnterpriseAccountRemoveMfaRequest) SetOrientedNbId(v string) *EnterpriseAccountRemoveMfaRequest {
  s.OrientedNbId = &v
  return s
}

func (s *EnterpriseAccountRemoveMfaRequest) SetPk(v string) *EnterpriseAccountRemoveMfaRequest {
  s.Pk = &v
  return s
}

func (s *EnterpriseAccountRemoveMfaRequest) SetRequestId(v string) *EnterpriseAccountRemoveMfaRequest {
  s.RequestId = &v
  return s
}

func (s *EnterpriseAccountRemoveMfaRequest) Validate() error {
  return dara.Validate(s)
}

