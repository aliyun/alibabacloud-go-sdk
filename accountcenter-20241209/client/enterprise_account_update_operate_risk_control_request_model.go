// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountUpdateOperateRiskControlRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppName(v string) *EnterpriseAccountUpdateOperateRiskControlRequest
  GetAppName() *string 
  SetOrientedEcId(v string) *EnterpriseAccountUpdateOperateRiskControlRequest
  GetOrientedEcId() *string 
  SetOrientedLeId(v string) *EnterpriseAccountUpdateOperateRiskControlRequest
  GetOrientedLeId() *string 
  SetOrientedNbId(v string) *EnterpriseAccountUpdateOperateRiskControlRequest
  GetOrientedNbId() *string 
  SetPk(v string) *EnterpriseAccountUpdateOperateRiskControlRequest
  GetPk() *string 
  SetProductLevel(v string) *EnterpriseAccountUpdateOperateRiskControlRequest
  GetProductLevel() *string 
  SetRequestId(v string) *EnterpriseAccountUpdateOperateRiskControlRequest
  GetRequestId() *string 
  SetValidateType(v string) *EnterpriseAccountUpdateOperateRiskControlRequest
  GetValidateType() *string 
}

type EnterpriseAccountUpdateOperateRiskControlRequest struct {
  AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
  OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
  OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
  OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
  // This parameter is required.
  Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
  // This parameter is required.
  ProductLevel *string `json:"ProductLevel,omitempty" xml:"ProductLevel,omitempty"`
  // This parameter is required.
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // This parameter is required.
  ValidateType *string `json:"ValidateType,omitempty" xml:"ValidateType,omitempty"`
}

func (s EnterpriseAccountUpdateOperateRiskControlRequest) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountUpdateOperateRiskControlRequest) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountUpdateOperateRiskControlRequest) GetAppName() *string  {
  return s.AppName
}

func (s *EnterpriseAccountUpdateOperateRiskControlRequest) GetOrientedEcId() *string  {
  return s.OrientedEcId
}

func (s *EnterpriseAccountUpdateOperateRiskControlRequest) GetOrientedLeId() *string  {
  return s.OrientedLeId
}

func (s *EnterpriseAccountUpdateOperateRiskControlRequest) GetOrientedNbId() *string  {
  return s.OrientedNbId
}

func (s *EnterpriseAccountUpdateOperateRiskControlRequest) GetPk() *string  {
  return s.Pk
}

func (s *EnterpriseAccountUpdateOperateRiskControlRequest) GetProductLevel() *string  {
  return s.ProductLevel
}

func (s *EnterpriseAccountUpdateOperateRiskControlRequest) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseAccountUpdateOperateRiskControlRequest) GetValidateType() *string  {
  return s.ValidateType
}

func (s *EnterpriseAccountUpdateOperateRiskControlRequest) SetAppName(v string) *EnterpriseAccountUpdateOperateRiskControlRequest {
  s.AppName = &v
  return s
}

func (s *EnterpriseAccountUpdateOperateRiskControlRequest) SetOrientedEcId(v string) *EnterpriseAccountUpdateOperateRiskControlRequest {
  s.OrientedEcId = &v
  return s
}

func (s *EnterpriseAccountUpdateOperateRiskControlRequest) SetOrientedLeId(v string) *EnterpriseAccountUpdateOperateRiskControlRequest {
  s.OrientedLeId = &v
  return s
}

func (s *EnterpriseAccountUpdateOperateRiskControlRequest) SetOrientedNbId(v string) *EnterpriseAccountUpdateOperateRiskControlRequest {
  s.OrientedNbId = &v
  return s
}

func (s *EnterpriseAccountUpdateOperateRiskControlRequest) SetPk(v string) *EnterpriseAccountUpdateOperateRiskControlRequest {
  s.Pk = &v
  return s
}

func (s *EnterpriseAccountUpdateOperateRiskControlRequest) SetProductLevel(v string) *EnterpriseAccountUpdateOperateRiskControlRequest {
  s.ProductLevel = &v
  return s
}

func (s *EnterpriseAccountUpdateOperateRiskControlRequest) SetRequestId(v string) *EnterpriseAccountUpdateOperateRiskControlRequest {
  s.RequestId = &v
  return s
}

func (s *EnterpriseAccountUpdateOperateRiskControlRequest) SetValidateType(v string) *EnterpriseAccountUpdateOperateRiskControlRequest {
  s.ValidateType = &v
  return s
}

func (s *EnterpriseAccountUpdateOperateRiskControlRequest) Validate() error {
  return dara.Validate(s)
}

