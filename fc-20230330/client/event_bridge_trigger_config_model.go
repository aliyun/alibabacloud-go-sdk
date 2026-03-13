// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEventBridgeTriggerConfig interface {
  dara.Model
  String() string
  GoString() string
  SetAsyncInvocationType(v bool) *EventBridgeTriggerConfig
  GetAsyncInvocationType() *bool 
  SetEventRuleFilterPattern(v string) *EventBridgeTriggerConfig
  GetEventRuleFilterPattern() *string 
  SetEventSinkConfig(v *EventSinkConfig) *EventBridgeTriggerConfig
  GetEventSinkConfig() *EventSinkConfig 
  SetEventSourceConfig(v *EventSourceConfig) *EventBridgeTriggerConfig
  GetEventSourceConfig() *EventSourceConfig 
  SetRunOptions(v *RunOptions) *EventBridgeTriggerConfig
  GetRunOptions() *RunOptions 
  SetTriggerEnable(v bool) *EventBridgeTriggerConfig
  GetTriggerEnable() *bool 
}

type EventBridgeTriggerConfig struct {
  // Whether to invoke the function in asynchronous mode. Valid values:
  // 
  // 	- **true**
  // 
  // 	- **false**
  // 
  // >  The default value is **false**.
  // 
  // example:
  // 
  // true
  AsyncInvocationType *bool `json:"asyncInvocationType,omitempty" xml:"asyncInvocationType,omitempty"`
  // The event pattern. The value is in the JSON format. For more information, see [Event patterns](https://help.aliyun.com/document_detail/181432.html).
  // 
  // example:
  // 
  // {}
  EventRuleFilterPattern *string `json:"eventRuleFilterPattern,omitempty" xml:"eventRuleFilterPattern,omitempty"`
  // The event destination configurations.
  EventSinkConfig *EventSinkConfig `json:"eventSinkConfig,omitempty" xml:"eventSinkConfig,omitempty"`
  // The event source configurations.
  EventSourceConfig *EventSourceConfig `json:"eventSourceConfig,omitempty" xml:"eventSourceConfig,omitempty"`
  // The runtime configurations.
  RunOptions *RunOptions `json:"runOptions,omitempty" xml:"runOptions,omitempty"`
  // Whether to enable the trigger. Valid values:
  // 
  // 	- **true**
  // 
  // 	- **false**
  // 
  // >  The default value is **true**.
  // 
  // example:
  // 
  // true
  TriggerEnable *bool `json:"triggerEnable,omitempty" xml:"triggerEnable,omitempty"`
}

func (s EventBridgeTriggerConfig) String() string {
  return dara.Prettify(s)
}

func (s EventBridgeTriggerConfig) GoString() string {
  return s.String()
}

func (s *EventBridgeTriggerConfig) GetAsyncInvocationType() *bool  {
  return s.AsyncInvocationType
}

func (s *EventBridgeTriggerConfig) GetEventRuleFilterPattern() *string  {
  return s.EventRuleFilterPattern
}

func (s *EventBridgeTriggerConfig) GetEventSinkConfig() *EventSinkConfig  {
  return s.EventSinkConfig
}

func (s *EventBridgeTriggerConfig) GetEventSourceConfig() *EventSourceConfig  {
  return s.EventSourceConfig
}

func (s *EventBridgeTriggerConfig) GetRunOptions() *RunOptions  {
  return s.RunOptions
}

func (s *EventBridgeTriggerConfig) GetTriggerEnable() *bool  {
  return s.TriggerEnable
}

func (s *EventBridgeTriggerConfig) SetAsyncInvocationType(v bool) *EventBridgeTriggerConfig {
  s.AsyncInvocationType = &v
  return s
}

func (s *EventBridgeTriggerConfig) SetEventRuleFilterPattern(v string) *EventBridgeTriggerConfig {
  s.EventRuleFilterPattern = &v
  return s
}

func (s *EventBridgeTriggerConfig) SetEventSinkConfig(v *EventSinkConfig) *EventBridgeTriggerConfig {
  s.EventSinkConfig = v
  return s
}

func (s *EventBridgeTriggerConfig) SetEventSourceConfig(v *EventSourceConfig) *EventBridgeTriggerConfig {
  s.EventSourceConfig = v
  return s
}

func (s *EventBridgeTriggerConfig) SetRunOptions(v *RunOptions) *EventBridgeTriggerConfig {
  s.RunOptions = v
  return s
}

func (s *EventBridgeTriggerConfig) SetTriggerEnable(v bool) *EventBridgeTriggerConfig {
  s.TriggerEnable = &v
  return s
}

func (s *EventBridgeTriggerConfig) Validate() error {
  if s.EventSinkConfig != nil {
    if err := s.EventSinkConfig.Validate(); err != nil {
      return err
    }
  }
  if s.EventSourceConfig != nil {
    if err := s.EventSourceConfig.Validate(); err != nil {
      return err
    }
  }
  if s.RunOptions != nil {
    if err := s.RunOptions.Validate(); err != nil {
      return err
    }
  }
  return nil
}

