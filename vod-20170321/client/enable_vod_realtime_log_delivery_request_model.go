// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableVodRealtimeLogDeliveryRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDomainName(v string) *EnableVodRealtimeLogDeliveryRequest
  GetDomainName() *string 
  SetOwnerId(v int64) *EnableVodRealtimeLogDeliveryRequest
  GetOwnerId() *int64 
}

type EnableVodRealtimeLogDeliveryRequest struct {
  // This parameter is required.
  DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s EnableVodRealtimeLogDeliveryRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableVodRealtimeLogDeliveryRequest) GoString() string {
  return s.String()
}

func (s *EnableVodRealtimeLogDeliveryRequest) GetDomainName() *string  {
  return s.DomainName
}

func (s *EnableVodRealtimeLogDeliveryRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *EnableVodRealtimeLogDeliveryRequest) SetDomainName(v string) *EnableVodRealtimeLogDeliveryRequest {
  s.DomainName = &v
  return s
}

func (s *EnableVodRealtimeLogDeliveryRequest) SetOwnerId(v int64) *EnableVodRealtimeLogDeliveryRequest {
  s.OwnerId = &v
  return s
}

func (s *EnableVodRealtimeLogDeliveryRequest) Validate() error {
  return dara.Validate(s)
}

