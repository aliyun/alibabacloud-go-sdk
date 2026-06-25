// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExpandPhoneDataVolumeRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAutoPay(v bool) *ExpandPhoneDataVolumeRequest
  GetAutoPay() *bool 
  SetBizRegionId(v string) *ExpandPhoneDataVolumeRequest
  GetBizRegionId() *string 
  SetInstanceIds(v []*string) *ExpandPhoneDataVolumeRequest
  GetInstanceIds() []*string 
  SetPaidCallBackUrl(v string) *ExpandPhoneDataVolumeRequest
  GetPaidCallBackUrl() *string 
  SetPhoneDataVolume(v int32) *ExpandPhoneDataVolumeRequest
  GetPhoneDataVolume() *int32 
  SetPromotionId(v string) *ExpandPhoneDataVolumeRequest
  GetPromotionId() *string 
}

type ExpandPhoneDataVolumeRequest struct {
  // Whether to enable automatic payment. The default value is false.
  // 
  // example:
  // 
  // false
  AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
  // The region ID.
  // 
  // example:
  // 
  // cn-hangzhou
  BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
  // A list of cloud phone matrix instance IDs. You can specify 1 to 100 IDs.
  InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
  PaidCallBackUrl *string `json:"PaidCallBackUrl,omitempty" xml:"PaidCallBackUrl,omitempty"`
  // The target size of the phone storage, in GiB.	Notice: The new value must be greater than the current size of the phone storage.
  // 
  // example:
  // 
  // 10
  PhoneDataVolume *int32 `json:"PhoneDataVolume,omitempty" xml:"PhoneDataVolume,omitempty"`
  // The promotion ID.
  // 
  // example:
  // 
  // 50003308011****
  PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
}

func (s ExpandPhoneDataVolumeRequest) String() string {
  return dara.Prettify(s)
}

func (s ExpandPhoneDataVolumeRequest) GoString() string {
  return s.String()
}

func (s *ExpandPhoneDataVolumeRequest) GetAutoPay() *bool  {
  return s.AutoPay
}

func (s *ExpandPhoneDataVolumeRequest) GetBizRegionId() *string  {
  return s.BizRegionId
}

func (s *ExpandPhoneDataVolumeRequest) GetInstanceIds() []*string  {
  return s.InstanceIds
}

func (s *ExpandPhoneDataVolumeRequest) GetPaidCallBackUrl() *string  {
  return s.PaidCallBackUrl
}

func (s *ExpandPhoneDataVolumeRequest) GetPhoneDataVolume() *int32  {
  return s.PhoneDataVolume
}

func (s *ExpandPhoneDataVolumeRequest) GetPromotionId() *string  {
  return s.PromotionId
}

func (s *ExpandPhoneDataVolumeRequest) SetAutoPay(v bool) *ExpandPhoneDataVolumeRequest {
  s.AutoPay = &v
  return s
}

func (s *ExpandPhoneDataVolumeRequest) SetBizRegionId(v string) *ExpandPhoneDataVolumeRequest {
  s.BizRegionId = &v
  return s
}

func (s *ExpandPhoneDataVolumeRequest) SetInstanceIds(v []*string) *ExpandPhoneDataVolumeRequest {
  s.InstanceIds = v
  return s
}

func (s *ExpandPhoneDataVolumeRequest) SetPaidCallBackUrl(v string) *ExpandPhoneDataVolumeRequest {
  s.PaidCallBackUrl = &v
  return s
}

func (s *ExpandPhoneDataVolumeRequest) SetPhoneDataVolume(v int32) *ExpandPhoneDataVolumeRequest {
  s.PhoneDataVolume = &v
  return s
}

func (s *ExpandPhoneDataVolumeRequest) SetPromotionId(v string) *ExpandPhoneDataVolumeRequest {
  s.PromotionId = &v
  return s
}

func (s *ExpandPhoneDataVolumeRequest) Validate() error {
  return dara.Validate(s)
}

