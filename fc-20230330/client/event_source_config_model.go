// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEventSourceConfig interface {
  dara.Model
  String() string
  GoString() string
  SetEventSourceParameters(v *EventSourceParameters) *EventSourceConfig
  GetEventSourceParameters() *EventSourceParameters 
  SetEventSourceType(v string) *EventSourceConfig
  GetEventSourceType() *string 
}

type EventSourceConfig struct {
  // The event source. Custom event sources include Message Service (MNS), ApsaraMQ for RocketMQ, ApsaraMQ for RabbitMQ, ApsaraMQ for Kafka, ApsaraMQ for MQTT, and Data Transmission Service (DTS).
  EventSourceParameters *EventSourceParameters `json:"eventSourceParameters,omitempty" xml:"eventSourceParameters,omitempty"`
  // The event source type. Valid values:
  // 
  // 	- **Default**: Alibaba Cloud EventBridge sources
  // 
  // 	- **MNS**: Message Service (MNS)
  // 
  // 	- **RocketMQ**: ApsaraMQ for RocketMQ
  // 
  // 	- **RabbitMQ**: ApsaraMQ for RabbitMQ
  // 
  // 	- **Kafka**: ApsaraMQ for Kafka
  // 
  // 	- **MQTT**: ApsaraMQ for MQTT
  // 
  // 	- **DTS**: DTS
  // 
  // >  This parameter cannot be updated. If you specify this parameter when you update the configurations, it does not take effect.
  // 
  // example:
  // 
  // MNS
  EventSourceType *string `json:"eventSourceType,omitempty" xml:"eventSourceType,omitempty"`
}

func (s EventSourceConfig) String() string {
  return dara.Prettify(s)
}

func (s EventSourceConfig) GoString() string {
  return s.String()
}

func (s *EventSourceConfig) GetEventSourceParameters() *EventSourceParameters  {
  return s.EventSourceParameters
}

func (s *EventSourceConfig) GetEventSourceType() *string  {
  return s.EventSourceType
}

func (s *EventSourceConfig) SetEventSourceParameters(v *EventSourceParameters) *EventSourceConfig {
  s.EventSourceParameters = v
  return s
}

func (s *EventSourceConfig) SetEventSourceType(v string) *EventSourceConfig {
  s.EventSourceType = &v
  return s
}

func (s *EventSourceConfig) Validate() error {
  if s.EventSourceParameters != nil {
    if err := s.EventSourceParameters.Validate(); err != nil {
      return err
    }
  }
  return nil
}

