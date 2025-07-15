// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableLiveRealtimeLogDeliveryRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDomainName(v string) *EnableLiveRealtimeLogDeliveryRequest
  GetDomainName() *string 
  SetOwnerId(v int64) *EnableLiveRealtimeLogDeliveryRequest
  GetOwnerId() *int64 
  SetRegionId(v string) *EnableLiveRealtimeLogDeliveryRequest
  GetRegionId() *string 
}

type EnableLiveRealtimeLogDeliveryRequest struct {
  // The streaming domain for which you want to enable real-time log delivery.
  // 
  // Separate multiple streaming domains with commas (,).
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // example.com
  DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s EnableLiveRealtimeLogDeliveryRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableLiveRealtimeLogDeliveryRequest) GoString() string {
  return s.String()
}

func (s *EnableLiveRealtimeLogDeliveryRequest) GetDomainName() *string  {
  return s.DomainName
}

func (s *EnableLiveRealtimeLogDeliveryRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *EnableLiveRealtimeLogDeliveryRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableLiveRealtimeLogDeliveryRequest) SetDomainName(v string) *EnableLiveRealtimeLogDeliveryRequest {
  s.DomainName = &v
  return s
}

func (s *EnableLiveRealtimeLogDeliveryRequest) SetOwnerId(v int64) *EnableLiveRealtimeLogDeliveryRequest {
  s.OwnerId = &v
  return s
}

func (s *EnableLiveRealtimeLogDeliveryRequest) SetRegionId(v string) *EnableLiveRealtimeLogDeliveryRequest {
  s.RegionId = &v
  return s
}

func (s *EnableLiveRealtimeLogDeliveryRequest) Validate() error {
  return dara.Validate(s)
}

