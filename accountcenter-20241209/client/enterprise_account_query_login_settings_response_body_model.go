// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountQueryLoginSettingsResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnterpriseAccountQueryLoginSettingsResponseBody
  GetCode() *string 
  SetData(v *EnterpriseAccountQueryLoginSettingsResponseBodyData) *EnterpriseAccountQueryLoginSettingsResponseBody
  GetData() *EnterpriseAccountQueryLoginSettingsResponseBodyData 
  SetMessage(v string) *EnterpriseAccountQueryLoginSettingsResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnterpriseAccountQueryLoginSettingsResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnterpriseAccountQueryLoginSettingsResponseBody
  GetSuccess() *bool 
}

type EnterpriseAccountQueryLoginSettingsResponseBody struct {
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Data *EnterpriseAccountQueryLoginSettingsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseAccountQueryLoginSettingsResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountQueryLoginSettingsResponseBody) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBody) GetData() *EnterpriseAccountQueryLoginSettingsResponseBodyData  {
  return s.Data
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBody) SetCode(v string) *EnterpriseAccountQueryLoginSettingsResponseBody {
  s.Code = &v
  return s
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBody) SetData(v *EnterpriseAccountQueryLoginSettingsResponseBodyData) *EnterpriseAccountQueryLoginSettingsResponseBody {
  s.Data = v
  return s
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBody) SetMessage(v string) *EnterpriseAccountQueryLoginSettingsResponseBody {
  s.Message = &v
  return s
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBody) SetRequestId(v string) *EnterpriseAccountQueryLoginSettingsResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBody) SetSuccess(v bool) *EnterpriseAccountQueryLoginSettingsResponseBody {
  s.Success = &v
  return s
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EnterpriseAccountQueryLoginSettingsResponseBodyData struct {
  IpMaskDto *EnterpriseAccountQueryLoginSettingsResponseBodyDataIpMaskDto `json:"IpMaskDto,omitempty" xml:"IpMaskDto,omitempty" type:"Struct"`
  MfaBindStatus *string `json:"MfaBindStatus,omitempty" xml:"MfaBindStatus,omitempty"`
  RiskControlDto *EnterpriseAccountQueryLoginSettingsResponseBodyDataRiskControlDto `json:"RiskControlDto,omitempty" xml:"RiskControlDto,omitempty" type:"Struct"`
  SecurityMobileLoginStatus *string `json:"SecurityMobileLoginStatus,omitempty" xml:"SecurityMobileLoginStatus,omitempty"`
  SessionExpireTime *int32 `json:"SessionExpireTime,omitempty" xml:"SessionExpireTime,omitempty"`
}

func (s EnterpriseAccountQueryLoginSettingsResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountQueryLoginSettingsResponseBodyData) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBodyData) GetIpMaskDto() *EnterpriseAccountQueryLoginSettingsResponseBodyDataIpMaskDto  {
  return s.IpMaskDto
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBodyData) GetMfaBindStatus() *string  {
  return s.MfaBindStatus
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBodyData) GetRiskControlDto() *EnterpriseAccountQueryLoginSettingsResponseBodyDataRiskControlDto  {
  return s.RiskControlDto
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBodyData) GetSecurityMobileLoginStatus() *string  {
  return s.SecurityMobileLoginStatus
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBodyData) GetSessionExpireTime() *int32  {
  return s.SessionExpireTime
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBodyData) SetIpMaskDto(v *EnterpriseAccountQueryLoginSettingsResponseBodyDataIpMaskDto) *EnterpriseAccountQueryLoginSettingsResponseBodyData {
  s.IpMaskDto = v
  return s
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBodyData) SetMfaBindStatus(v string) *EnterpriseAccountQueryLoginSettingsResponseBodyData {
  s.MfaBindStatus = &v
  return s
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBodyData) SetRiskControlDto(v *EnterpriseAccountQueryLoginSettingsResponseBodyDataRiskControlDto) *EnterpriseAccountQueryLoginSettingsResponseBodyData {
  s.RiskControlDto = v
  return s
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBodyData) SetSecurityMobileLoginStatus(v string) *EnterpriseAccountQueryLoginSettingsResponseBodyData {
  s.SecurityMobileLoginStatus = &v
  return s
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBodyData) SetSessionExpireTime(v int32) *EnterpriseAccountQueryLoginSettingsResponseBodyData {
  s.SessionExpireTime = &v
  return s
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBodyData) Validate() error {
  if s.IpMaskDto != nil {
    if err := s.IpMaskDto.Validate(); err != nil {
      return err
    }
  }
  if s.RiskControlDto != nil {
    if err := s.RiskControlDto.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EnterpriseAccountQueryLoginSettingsResponseBodyDataIpMaskDto struct {
  IpMaskEnabledStatus *string `json:"IpMaskEnabledStatus,omitempty" xml:"IpMaskEnabledStatus,omitempty"`
  IpMasks []*string `json:"IpMasks,omitempty" xml:"IpMasks,omitempty" type:"Repeated"`
}

func (s EnterpriseAccountQueryLoginSettingsResponseBodyDataIpMaskDto) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountQueryLoginSettingsResponseBodyDataIpMaskDto) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBodyDataIpMaskDto) GetIpMaskEnabledStatus() *string  {
  return s.IpMaskEnabledStatus
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBodyDataIpMaskDto) GetIpMasks() []*string  {
  return s.IpMasks
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBodyDataIpMaskDto) SetIpMaskEnabledStatus(v string) *EnterpriseAccountQueryLoginSettingsResponseBodyDataIpMaskDto {
  s.IpMaskEnabledStatus = &v
  return s
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBodyDataIpMaskDto) SetIpMasks(v []*string) *EnterpriseAccountQueryLoginSettingsResponseBodyDataIpMaskDto {
  s.IpMasks = v
  return s
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBodyDataIpMaskDto) Validate() error {
  return dara.Validate(s)
}

type EnterpriseAccountQueryLoginSettingsResponseBodyDataRiskControlDto struct {
  ProtectLevel *string `json:"ProtectLevel,omitempty" xml:"ProtectLevel,omitempty"`
  RiskControlEnabled *bool `json:"RiskControlEnabled,omitempty" xml:"RiskControlEnabled,omitempty"`
  VerifyDetail *string `json:"VerifyDetail,omitempty" xml:"VerifyDetail,omitempty"`
  VerifyType *string `json:"VerifyType,omitempty" xml:"VerifyType,omitempty"`
}

func (s EnterpriseAccountQueryLoginSettingsResponseBodyDataRiskControlDto) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountQueryLoginSettingsResponseBodyDataRiskControlDto) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBodyDataRiskControlDto) GetProtectLevel() *string  {
  return s.ProtectLevel
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBodyDataRiskControlDto) GetRiskControlEnabled() *bool  {
  return s.RiskControlEnabled
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBodyDataRiskControlDto) GetVerifyDetail() *string  {
  return s.VerifyDetail
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBodyDataRiskControlDto) GetVerifyType() *string  {
  return s.VerifyType
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBodyDataRiskControlDto) SetProtectLevel(v string) *EnterpriseAccountQueryLoginSettingsResponseBodyDataRiskControlDto {
  s.ProtectLevel = &v
  return s
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBodyDataRiskControlDto) SetRiskControlEnabled(v bool) *EnterpriseAccountQueryLoginSettingsResponseBodyDataRiskControlDto {
  s.RiskControlEnabled = &v
  return s
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBodyDataRiskControlDto) SetVerifyDetail(v string) *EnterpriseAccountQueryLoginSettingsResponseBodyDataRiskControlDto {
  s.VerifyDetail = &v
  return s
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBodyDataRiskControlDto) SetVerifyType(v string) *EnterpriseAccountQueryLoginSettingsResponseBodyDataRiskControlDto {
  s.VerifyType = &v
  return s
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBodyDataRiskControlDto) Validate() error {
  return dara.Validate(s)
}

