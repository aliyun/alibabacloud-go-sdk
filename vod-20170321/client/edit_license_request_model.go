// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditLicenseRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppItemId(v string) *EditLicenseRequest
  GetAppItemId() *string 
  SetAppName(v string) *EditLicenseRequest
  GetAppName() *string 
  SetAppPlatforms(v string) *EditLicenseRequest
  GetAppPlatforms() *string 
  SetContractNo(v string) *EditLicenseRequest
  GetContractNo() *string 
  SetInstanceId(v string) *EditLicenseRequest
  GetInstanceId() *string 
  SetSdkModels(v string) *EditLicenseRequest
  GetSdkModels() *string 
}

type EditLicenseRequest struct {
  AppItemId *string `json:"AppItemId,omitempty" xml:"AppItemId,omitempty"`
  AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
  AppPlatforms *string `json:"AppPlatforms,omitempty" xml:"AppPlatforms,omitempty"`
  ContractNo *string `json:"ContractNo,omitempty" xml:"ContractNo,omitempty"`
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  SdkModels *string `json:"SdkModels,omitempty" xml:"SdkModels,omitempty"`
}

func (s EditLicenseRequest) String() string {
  return dara.Prettify(s)
}

func (s EditLicenseRequest) GoString() string {
  return s.String()
}

func (s *EditLicenseRequest) GetAppItemId() *string  {
  return s.AppItemId
}

func (s *EditLicenseRequest) GetAppName() *string  {
  return s.AppName
}

func (s *EditLicenseRequest) GetAppPlatforms() *string  {
  return s.AppPlatforms
}

func (s *EditLicenseRequest) GetContractNo() *string  {
  return s.ContractNo
}

func (s *EditLicenseRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EditLicenseRequest) GetSdkModels() *string  {
  return s.SdkModels
}

func (s *EditLicenseRequest) SetAppItemId(v string) *EditLicenseRequest {
  s.AppItemId = &v
  return s
}

func (s *EditLicenseRequest) SetAppName(v string) *EditLicenseRequest {
  s.AppName = &v
  return s
}

func (s *EditLicenseRequest) SetAppPlatforms(v string) *EditLicenseRequest {
  s.AppPlatforms = &v
  return s
}

func (s *EditLicenseRequest) SetContractNo(v string) *EditLicenseRequest {
  s.ContractNo = &v
  return s
}

func (s *EditLicenseRequest) SetInstanceId(v string) *EditLicenseRequest {
  s.InstanceId = &v
  return s
}

func (s *EditLicenseRequest) SetSdkModels(v string) *EditLicenseRequest {
  s.SdkModels = &v
  return s
}

func (s *EditLicenseRequest) Validate() error {
  return dara.Validate(s)
}

