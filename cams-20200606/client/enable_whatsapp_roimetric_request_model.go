// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableWhatsappROIMetricRequest interface {
  dara.Model
  String() string
  GoString() string
  SetCustSpaceId(v string) *EnableWhatsappROIMetricRequest
  GetCustSpaceId() *string 
  SetIsvCode(v string) *EnableWhatsappROIMetricRequest
  GetIsvCode() *string 
  SetOwnerId(v int64) *EnableWhatsappROIMetricRequest
  GetOwnerId() *int64 
  SetResourceOwnerAccount(v string) *EnableWhatsappROIMetricRequest
  GetResourceOwnerAccount() *string 
  SetResourceOwnerId(v int64) *EnableWhatsappROIMetricRequest
  GetResourceOwnerId() *int64 
}

type EnableWhatsappROIMetricRequest struct {
  // The space ID or instance ID of the ISV sub-customer. This is the channel ID, which can be viewed on the <props="china">[Channel Management](https://chatapp.console.aliyun.com/ChannelsManagement)<props="intl">[Channel Management](https://chatapp.console.alibabacloud.com/CustomerList) page.
  // 
  // example:
  // 
  // cams-************
  CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
  // The ISV verification code, which is used to verify that the RAM user is authorized by the ISV.
  // 
  // example:
  // 
  // skdi3kksloslikd****
  IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
  ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s EnableWhatsappROIMetricRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableWhatsappROIMetricRequest) GoString() string {
  return s.String()
}

func (s *EnableWhatsappROIMetricRequest) GetCustSpaceId() *string  {
  return s.CustSpaceId
}

func (s *EnableWhatsappROIMetricRequest) GetIsvCode() *string  {
  return s.IsvCode
}

func (s *EnableWhatsappROIMetricRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *EnableWhatsappROIMetricRequest) GetResourceOwnerAccount() *string  {
  return s.ResourceOwnerAccount
}

func (s *EnableWhatsappROIMetricRequest) GetResourceOwnerId() *int64  {
  return s.ResourceOwnerId
}

func (s *EnableWhatsappROIMetricRequest) SetCustSpaceId(v string) *EnableWhatsappROIMetricRequest {
  s.CustSpaceId = &v
  return s
}

func (s *EnableWhatsappROIMetricRequest) SetIsvCode(v string) *EnableWhatsappROIMetricRequest {
  s.IsvCode = &v
  return s
}

func (s *EnableWhatsappROIMetricRequest) SetOwnerId(v int64) *EnableWhatsappROIMetricRequest {
  s.OwnerId = &v
  return s
}

func (s *EnableWhatsappROIMetricRequest) SetResourceOwnerAccount(v string) *EnableWhatsappROIMetricRequest {
  s.ResourceOwnerAccount = &v
  return s
}

func (s *EnableWhatsappROIMetricRequest) SetResourceOwnerId(v int64) *EnableWhatsappROIMetricRequest {
  s.ResourceOwnerId = &v
  return s
}

func (s *EnableWhatsappROIMetricRequest) Validate() error {
  return dara.Validate(s)
}

