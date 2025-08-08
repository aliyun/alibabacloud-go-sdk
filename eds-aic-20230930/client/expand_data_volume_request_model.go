// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExpandDataVolumeRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAutoPay(v bool) *ExpandDataVolumeRequest
  GetAutoPay() *bool 
  SetBizRegionId(v string) *ExpandDataVolumeRequest
  GetBizRegionId() *string 
  SetNodeIds(v []*string) *ExpandDataVolumeRequest
  GetNodeIds() []*string 
  SetPhoneDataVolume(v int32) *ExpandDataVolumeRequest
  GetPhoneDataVolume() *int32 
  SetShareDataVolume(v int32) *ExpandDataVolumeRequest
  GetShareDataVolume() *int32 
}

type ExpandDataVolumeRequest struct {
  // example:
  // 
  // true
  AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
  // example:
  // 
  // cn-hangzhou
  BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
  NodeIds []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
  PhoneDataVolume *int32 `json:"PhoneDataVolume,omitempty" xml:"PhoneDataVolume,omitempty"`
  // example:
  // 
  // 100
  ShareDataVolume *int32 `json:"ShareDataVolume,omitempty" xml:"ShareDataVolume,omitempty"`
}

func (s ExpandDataVolumeRequest) String() string {
  return dara.Prettify(s)
}

func (s ExpandDataVolumeRequest) GoString() string {
  return s.String()
}

func (s *ExpandDataVolumeRequest) GetAutoPay() *bool  {
  return s.AutoPay
}

func (s *ExpandDataVolumeRequest) GetBizRegionId() *string  {
  return s.BizRegionId
}

func (s *ExpandDataVolumeRequest) GetNodeIds() []*string  {
  return s.NodeIds
}

func (s *ExpandDataVolumeRequest) GetPhoneDataVolume() *int32  {
  return s.PhoneDataVolume
}

func (s *ExpandDataVolumeRequest) GetShareDataVolume() *int32  {
  return s.ShareDataVolume
}

func (s *ExpandDataVolumeRequest) SetAutoPay(v bool) *ExpandDataVolumeRequest {
  s.AutoPay = &v
  return s
}

func (s *ExpandDataVolumeRequest) SetBizRegionId(v string) *ExpandDataVolumeRequest {
  s.BizRegionId = &v
  return s
}

func (s *ExpandDataVolumeRequest) SetNodeIds(v []*string) *ExpandDataVolumeRequest {
  s.NodeIds = v
  return s
}

func (s *ExpandDataVolumeRequest) SetPhoneDataVolume(v int32) *ExpandDataVolumeRequest {
  s.PhoneDataVolume = &v
  return s
}

func (s *ExpandDataVolumeRequest) SetShareDataVolume(v int32) *ExpandDataVolumeRequest {
  s.ShareDataVolume = &v
  return s
}

func (s *ExpandDataVolumeRequest) Validate() error {
  return dara.Validate(s)
}

