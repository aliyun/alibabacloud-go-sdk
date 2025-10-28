// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEventSourceParameters interface {
  dara.Model
  String() string
  GoString() string
  SetSourceDTSParameters(v *SourceDTSParameters) *EventSourceParameters
  GetSourceDTSParameters() *SourceDTSParameters 
  SetSourceKafkaParameters(v *SourceKafkaParameters) *EventSourceParameters
  GetSourceKafkaParameters() *SourceKafkaParameters 
  SetSourceMNSParameters(v *SourceMNSParameters) *EventSourceParameters
  GetSourceMNSParameters() *SourceMNSParameters 
  SetSourceMQTTParameters(v *SourceMQTTParameters) *EventSourceParameters
  GetSourceMQTTParameters() *SourceMQTTParameters 
  SetSourceRabbitMQParameters(v *SourceRabbitMQParameters) *EventSourceParameters
  GetSourceRabbitMQParameters() *SourceRabbitMQParameters 
  SetSourceRocketMQParameters(v *SourceRocketMQParameters) *EventSourceParameters
  GetSourceRocketMQParameters() *SourceRocketMQParameters 
}

type EventSourceParameters struct {
  SourceDTSParameters *SourceDTSParameters `json:"sourceDTSParameters,omitempty" xml:"sourceDTSParameters,omitempty"`
  SourceKafkaParameters *SourceKafkaParameters `json:"sourceKafkaParameters,omitempty" xml:"sourceKafkaParameters,omitempty"`
  SourceMNSParameters *SourceMNSParameters `json:"sourceMNSParameters,omitempty" xml:"sourceMNSParameters,omitempty"`
  SourceMQTTParameters *SourceMQTTParameters `json:"sourceMQTTParameters,omitempty" xml:"sourceMQTTParameters,omitempty"`
  SourceRabbitMQParameters *SourceRabbitMQParameters `json:"sourceRabbitMQParameters,omitempty" xml:"sourceRabbitMQParameters,omitempty"`
  SourceRocketMQParameters *SourceRocketMQParameters `json:"sourceRocketMQParameters,omitempty" xml:"sourceRocketMQParameters,omitempty"`
}

func (s EventSourceParameters) String() string {
  return dara.Prettify(s)
}

func (s EventSourceParameters) GoString() string {
  return s.String()
}

func (s *EventSourceParameters) GetSourceDTSParameters() *SourceDTSParameters  {
  return s.SourceDTSParameters
}

func (s *EventSourceParameters) GetSourceKafkaParameters() *SourceKafkaParameters  {
  return s.SourceKafkaParameters
}

func (s *EventSourceParameters) GetSourceMNSParameters() *SourceMNSParameters  {
  return s.SourceMNSParameters
}

func (s *EventSourceParameters) GetSourceMQTTParameters() *SourceMQTTParameters  {
  return s.SourceMQTTParameters
}

func (s *EventSourceParameters) GetSourceRabbitMQParameters() *SourceRabbitMQParameters  {
  return s.SourceRabbitMQParameters
}

func (s *EventSourceParameters) GetSourceRocketMQParameters() *SourceRocketMQParameters  {
  return s.SourceRocketMQParameters
}

func (s *EventSourceParameters) SetSourceDTSParameters(v *SourceDTSParameters) *EventSourceParameters {
  s.SourceDTSParameters = v
  return s
}

func (s *EventSourceParameters) SetSourceKafkaParameters(v *SourceKafkaParameters) *EventSourceParameters {
  s.SourceKafkaParameters = v
  return s
}

func (s *EventSourceParameters) SetSourceMNSParameters(v *SourceMNSParameters) *EventSourceParameters {
  s.SourceMNSParameters = v
  return s
}

func (s *EventSourceParameters) SetSourceMQTTParameters(v *SourceMQTTParameters) *EventSourceParameters {
  s.SourceMQTTParameters = v
  return s
}

func (s *EventSourceParameters) SetSourceRabbitMQParameters(v *SourceRabbitMQParameters) *EventSourceParameters {
  s.SourceRabbitMQParameters = v
  return s
}

func (s *EventSourceParameters) SetSourceRocketMQParameters(v *SourceRocketMQParameters) *EventSourceParameters {
  s.SourceRocketMQParameters = v
  return s
}

func (s *EventSourceParameters) Validate() error {
  if s.SourceDTSParameters != nil {
    if err := s.SourceDTSParameters.Validate(); err != nil {
      return err
    }
  }
  if s.SourceKafkaParameters != nil {
    if err := s.SourceKafkaParameters.Validate(); err != nil {
      return err
    }
  }
  if s.SourceMNSParameters != nil {
    if err := s.SourceMNSParameters.Validate(); err != nil {
      return err
    }
  }
  if s.SourceMQTTParameters != nil {
    if err := s.SourceMQTTParameters.Validate(); err != nil {
      return err
    }
  }
  if s.SourceRabbitMQParameters != nil {
    if err := s.SourceRabbitMQParameters.Validate(); err != nil {
      return err
    }
  }
  if s.SourceRocketMQParameters != nil {
    if err := s.SourceRocketMQParameters.Validate(); err != nil {
      return err
    }
  }
  return nil
}

