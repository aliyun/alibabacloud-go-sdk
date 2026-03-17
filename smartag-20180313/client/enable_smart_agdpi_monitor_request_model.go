// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSmartAGDpiMonitorRequest interface {
  dara.Model
  String() string
  GoString() string
  SetClientToken(v string) *EnableSmartAGDpiMonitorRequest
  GetClientToken() *string 
  SetDryRun(v bool) *EnableSmartAGDpiMonitorRequest
  GetDryRun() *bool 
  SetOwnerAccount(v string) *EnableSmartAGDpiMonitorRequest
  GetOwnerAccount() *string 
  SetOwnerId(v int64) *EnableSmartAGDpiMonitorRequest
  GetOwnerId() *int64 
  SetRegionId(v string) *EnableSmartAGDpiMonitorRequest
  GetRegionId() *string 
  SetResourceOwnerAccount(v string) *EnableSmartAGDpiMonitorRequest
  GetResourceOwnerAccount() *string 
  SetResourceOwnerId(v int64) *EnableSmartAGDpiMonitorRequest
  GetResourceOwnerId() *int64 
  SetSlsLogStore(v string) *EnableSmartAGDpiMonitorRequest
  GetSlsLogStore() *string 
  SetSlsProjectName(v string) *EnableSmartAGDpiMonitorRequest
  GetSlsProjectName() *string 
  SetSmartAGId(v string) *EnableSmartAGDpiMonitorRequest
  GetSmartAGId() *string 
}

type EnableSmartAGDpiMonitorRequest struct {
  // The client token that is used to ensure the idempotence of the request.
  // 
  // You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
  // 
  // example:
  // 
  // 02fb3da4-130e-11e9****
  ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
  // Specifies whether to check the validity of the request without actually making the request. Valid values:
  // 
  // 	- **true**: checks the validity of the request but does not enable or disables the DPI feature for the SAG instance. Check items include the request format, instance status, and whether the required parameters are set. If the request fails the check, an error message is returned. If the request passes the request, the ID of the request is returned.
  // 
  // 	- **false**: makes the request and enables the DPI feature for the SAG instance after the request passes the check. This is the default value.
  // 
  // example:
  // 
  // true
  DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
  OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  // The ID of the region where the SAG instance is deployed.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-shanghai
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
  ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
  // The name of the Logstore.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // test2
  SlsLogStore *string `json:"SlsLogStore,omitempty" xml:"SlsLogStore,omitempty"`
  // The name of the Log Service project.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // test1
  SlsProjectName *string `json:"SlsProjectName,omitempty" xml:"SlsProjectName,omitempty"`
  // The ID of the SAG instance.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // sag-vwmylqc9521p5l****
  SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s EnableSmartAGDpiMonitorRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableSmartAGDpiMonitorRequest) GoString() string {
  return s.String()
}

func (s *EnableSmartAGDpiMonitorRequest) GetClientToken() *string  {
  return s.ClientToken
}

func (s *EnableSmartAGDpiMonitorRequest) GetDryRun() *bool  {
  return s.DryRun
}

func (s *EnableSmartAGDpiMonitorRequest) GetOwnerAccount() *string  {
  return s.OwnerAccount
}

func (s *EnableSmartAGDpiMonitorRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *EnableSmartAGDpiMonitorRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableSmartAGDpiMonitorRequest) GetResourceOwnerAccount() *string  {
  return s.ResourceOwnerAccount
}

func (s *EnableSmartAGDpiMonitorRequest) GetResourceOwnerId() *int64  {
  return s.ResourceOwnerId
}

func (s *EnableSmartAGDpiMonitorRequest) GetSlsLogStore() *string  {
  return s.SlsLogStore
}

func (s *EnableSmartAGDpiMonitorRequest) GetSlsProjectName() *string  {
  return s.SlsProjectName
}

func (s *EnableSmartAGDpiMonitorRequest) GetSmartAGId() *string  {
  return s.SmartAGId
}

func (s *EnableSmartAGDpiMonitorRequest) SetClientToken(v string) *EnableSmartAGDpiMonitorRequest {
  s.ClientToken = &v
  return s
}

func (s *EnableSmartAGDpiMonitorRequest) SetDryRun(v bool) *EnableSmartAGDpiMonitorRequest {
  s.DryRun = &v
  return s
}

func (s *EnableSmartAGDpiMonitorRequest) SetOwnerAccount(v string) *EnableSmartAGDpiMonitorRequest {
  s.OwnerAccount = &v
  return s
}

func (s *EnableSmartAGDpiMonitorRequest) SetOwnerId(v int64) *EnableSmartAGDpiMonitorRequest {
  s.OwnerId = &v
  return s
}

func (s *EnableSmartAGDpiMonitorRequest) SetRegionId(v string) *EnableSmartAGDpiMonitorRequest {
  s.RegionId = &v
  return s
}

func (s *EnableSmartAGDpiMonitorRequest) SetResourceOwnerAccount(v string) *EnableSmartAGDpiMonitorRequest {
  s.ResourceOwnerAccount = &v
  return s
}

func (s *EnableSmartAGDpiMonitorRequest) SetResourceOwnerId(v int64) *EnableSmartAGDpiMonitorRequest {
  s.ResourceOwnerId = &v
  return s
}

func (s *EnableSmartAGDpiMonitorRequest) SetSlsLogStore(v string) *EnableSmartAGDpiMonitorRequest {
  s.SlsLogStore = &v
  return s
}

func (s *EnableSmartAGDpiMonitorRequest) SetSlsProjectName(v string) *EnableSmartAGDpiMonitorRequest {
  s.SlsProjectName = &v
  return s
}

func (s *EnableSmartAGDpiMonitorRequest) SetSmartAGId(v string) *EnableSmartAGDpiMonitorRequest {
  s.SmartAGId = &v
  return s
}

func (s *EnableSmartAGDpiMonitorRequest) Validate() error {
  return dara.Validate(s)
}

