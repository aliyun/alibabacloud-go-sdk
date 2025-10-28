// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iElasticConfigStatus interface {
  dara.Model
  String() string
  GoString() string
  SetCurrentError(v string) *ElasticConfigStatus
  GetCurrentError() *string 
  SetCurrentInstances(v int64) *ElasticConfigStatus
  GetCurrentInstances() *int64 
  SetFunctionArn(v string) *ElasticConfigStatus
  GetFunctionArn() *string 
  SetMinInstances(v int64) *ElasticConfigStatus
  GetMinInstances() *int64 
  SetResidentPoolId(v string) *ElasticConfigStatus
  GetResidentPoolId() *string 
  SetScalingPolicies(v []*ScalingPolicy) *ElasticConfigStatus
  GetScalingPolicies() []*ScalingPolicy 
  SetScheduledPolicies(v []*ScheduledPolicy) *ElasticConfigStatus
  GetScheduledPolicies() []*ScheduledPolicy 
  SetTargetInstances(v int64) *ElasticConfigStatus
  GetTargetInstances() *int64 
}

type ElasticConfigStatus struct {
  CurrentError *string `json:"currentError,omitempty" xml:"currentError,omitempty"`
  CurrentInstances *int64 `json:"currentInstances,omitempty" xml:"currentInstances,omitempty"`
  FunctionArn *string `json:"functionArn,omitempty" xml:"functionArn,omitempty"`
  MinInstances *int64 `json:"minInstances,omitempty" xml:"minInstances,omitempty"`
  ResidentPoolId *string `json:"residentPoolId,omitempty" xml:"residentPoolId,omitempty"`
  ScalingPolicies []*ScalingPolicy `json:"scalingPolicies" xml:"scalingPolicies" type:"Repeated"`
  ScheduledPolicies []*ScheduledPolicy `json:"scheduledPolicies" xml:"scheduledPolicies" type:"Repeated"`
  TargetInstances *int64 `json:"targetInstances,omitempty" xml:"targetInstances,omitempty"`
}

func (s ElasticConfigStatus) String() string {
  return dara.Prettify(s)
}

func (s ElasticConfigStatus) GoString() string {
  return s.String()
}

func (s *ElasticConfigStatus) GetCurrentError() *string  {
  return s.CurrentError
}

func (s *ElasticConfigStatus) GetCurrentInstances() *int64  {
  return s.CurrentInstances
}

func (s *ElasticConfigStatus) GetFunctionArn() *string  {
  return s.FunctionArn
}

func (s *ElasticConfigStatus) GetMinInstances() *int64  {
  return s.MinInstances
}

func (s *ElasticConfigStatus) GetResidentPoolId() *string  {
  return s.ResidentPoolId
}

func (s *ElasticConfigStatus) GetScalingPolicies() []*ScalingPolicy  {
  return s.ScalingPolicies
}

func (s *ElasticConfigStatus) GetScheduledPolicies() []*ScheduledPolicy  {
  return s.ScheduledPolicies
}

func (s *ElasticConfigStatus) GetTargetInstances() *int64  {
  return s.TargetInstances
}

func (s *ElasticConfigStatus) SetCurrentError(v string) *ElasticConfigStatus {
  s.CurrentError = &v
  return s
}

func (s *ElasticConfigStatus) SetCurrentInstances(v int64) *ElasticConfigStatus {
  s.CurrentInstances = &v
  return s
}

func (s *ElasticConfigStatus) SetFunctionArn(v string) *ElasticConfigStatus {
  s.FunctionArn = &v
  return s
}

func (s *ElasticConfigStatus) SetMinInstances(v int64) *ElasticConfigStatus {
  s.MinInstances = &v
  return s
}

func (s *ElasticConfigStatus) SetResidentPoolId(v string) *ElasticConfigStatus {
  s.ResidentPoolId = &v
  return s
}

func (s *ElasticConfigStatus) SetScalingPolicies(v []*ScalingPolicy) *ElasticConfigStatus {
  s.ScalingPolicies = v
  return s
}

func (s *ElasticConfigStatus) SetScheduledPolicies(v []*ScheduledPolicy) *ElasticConfigStatus {
  s.ScheduledPolicies = v
  return s
}

func (s *ElasticConfigStatus) SetTargetInstances(v int64) *ElasticConfigStatus {
  s.TargetInstances = &v
  return s
}

func (s *ElasticConfigStatus) Validate() error {
  if s.ScalingPolicies != nil {
    for _, item := range s.ScalingPolicies {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  if s.ScheduledPolicies != nil {
    for _, item := range s.ScheduledPolicies {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

