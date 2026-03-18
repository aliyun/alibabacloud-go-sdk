// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableMultiAzRequest interface {
  dara.Model
  String() string
  GoString() string
  SetInstanceId(v string) *EnableMultiAzRequest
  GetInstanceId() *string 
  SetObservers(v []*EnableMultiAzRequestObservers) *EnableMultiAzRequest
  GetObservers() []*EnableMultiAzRequestObservers 
  SetPromotionOptionNo(v string) *EnableMultiAzRequest
  GetPromotionOptionNo() *string 
}

type EnableMultiAzRequest struct {
  // example:
  // 
  // c-238sjh237s12***
  InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
  Observers []*EnableMultiAzRequestObservers `json:"observers,omitempty" xml:"observers,omitempty" type:"Repeated"`
  // example:
  // 
  // youhuiquan_12378dfj6
  PromotionOptionNo *string `json:"promotionOptionNo,omitempty" xml:"promotionOptionNo,omitempty"`
}

func (s EnableMultiAzRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableMultiAzRequest) GoString() string {
  return s.String()
}

func (s *EnableMultiAzRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableMultiAzRequest) GetObservers() []*EnableMultiAzRequestObservers  {
  return s.Observers
}

func (s *EnableMultiAzRequest) GetPromotionOptionNo() *string  {
  return s.PromotionOptionNo
}

func (s *EnableMultiAzRequest) SetInstanceId(v string) *EnableMultiAzRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableMultiAzRequest) SetObservers(v []*EnableMultiAzRequestObservers) *EnableMultiAzRequest {
  s.Observers = v
  return s
}

func (s *EnableMultiAzRequest) SetPromotionOptionNo(v string) *EnableMultiAzRequest {
  s.PromotionOptionNo = &v
  return s
}

func (s *EnableMultiAzRequest) Validate() error {
  if s.Observers != nil {
    for _, item := range s.Observers {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type EnableMultiAzRequestObservers struct {
  // example:
  // 
  // vsw-x1232js012
  VswId *string `json:"vswId,omitempty" xml:"vswId,omitempty"`
  // example:
  // 
  // cn-hangzhou-h
  ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s EnableMultiAzRequestObservers) String() string {
  return dara.Prettify(s)
}

func (s EnableMultiAzRequestObservers) GoString() string {
  return s.String()
}

func (s *EnableMultiAzRequestObservers) GetVswId() *string  {
  return s.VswId
}

func (s *EnableMultiAzRequestObservers) GetZoneId() *string  {
  return s.ZoneId
}

func (s *EnableMultiAzRequestObservers) SetVswId(v string) *EnableMultiAzRequestObservers {
  s.VswId = &v
  return s
}

func (s *EnableMultiAzRequestObservers) SetZoneId(v string) *EnableMultiAzRequestObservers {
  s.ZoneId = &v
  return s
}

func (s *EnableMultiAzRequestObservers) Validate() error {
  return dara.Validate(s)
}

