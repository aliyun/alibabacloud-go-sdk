// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEventSinkConfig interface {
  dara.Model
  String() string
  GoString() string
  SetDeliveryOption(v *DeliveryOption) *EventSinkConfig
  GetDeliveryOption() *DeliveryOption 
}

type EventSinkConfig struct {
  DeliveryOption *DeliveryOption `json:"deliveryOption,omitempty" xml:"deliveryOption,omitempty"`
}

func (s EventSinkConfig) String() string {
  return dara.Prettify(s)
}

func (s EventSinkConfig) GoString() string {
  return s.String()
}

func (s *EventSinkConfig) GetDeliveryOption() *DeliveryOption  {
  return s.DeliveryOption
}

func (s *EventSinkConfig) SetDeliveryOption(v *DeliveryOption) *EventSinkConfig {
  s.DeliveryOption = v
  return s
}

func (s *EventSinkConfig) Validate() error {
  if s.DeliveryOption != nil {
    if err := s.DeliveryOption.Validate(); err != nil {
      return err
    }
  }
  return nil
}

