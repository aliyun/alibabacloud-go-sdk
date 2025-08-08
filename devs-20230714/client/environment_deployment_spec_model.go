// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnvironmentDeploymentSpec interface {
  dara.Model
  String() string
  GoString() string
  SetBaseline(v *EnvironmentSnapshot) *EnvironmentDeploymentSpec
  GetBaseline() *EnvironmentSnapshot 
  SetChanges(v *EnvironmentChanges) *EnvironmentDeploymentSpec
  GetChanges() *EnvironmentChanges 
  SetSkipRemoveResources(v bool) *EnvironmentDeploymentSpec
  GetSkipRemoveResources() *bool 
  SetTarget(v *EnvironmentStagedConfigs) *EnvironmentDeploymentSpec
  GetTarget() *EnvironmentStagedConfigs 
  SetWebhookCodeContext(v *WebhookCodeContext) *EnvironmentDeploymentSpec
  GetWebhookCodeContext() *WebhookCodeContext 
}

type EnvironmentDeploymentSpec struct {
  Baseline *EnvironmentSnapshot `json:"baseline,omitempty" xml:"baseline,omitempty"`
  Changes *EnvironmentChanges `json:"changes,omitempty" xml:"changes,omitempty"`
  // example:
  // 
  // false
  SkipRemoveResources *bool `json:"skipRemoveResources,omitempty" xml:"skipRemoveResources,omitempty"`
  Target *EnvironmentStagedConfigs `json:"target,omitempty" xml:"target,omitempty"`
  WebhookCodeContext *WebhookCodeContext `json:"webhookCodeContext,omitempty" xml:"webhookCodeContext,omitempty"`
}

func (s EnvironmentDeploymentSpec) String() string {
  return dara.Prettify(s)
}

func (s EnvironmentDeploymentSpec) GoString() string {
  return s.String()
}

func (s *EnvironmentDeploymentSpec) GetBaseline() *EnvironmentSnapshot  {
  return s.Baseline
}

func (s *EnvironmentDeploymentSpec) GetChanges() *EnvironmentChanges  {
  return s.Changes
}

func (s *EnvironmentDeploymentSpec) GetSkipRemoveResources() *bool  {
  return s.SkipRemoveResources
}

func (s *EnvironmentDeploymentSpec) GetTarget() *EnvironmentStagedConfigs  {
  return s.Target
}

func (s *EnvironmentDeploymentSpec) GetWebhookCodeContext() *WebhookCodeContext  {
  return s.WebhookCodeContext
}

func (s *EnvironmentDeploymentSpec) SetBaseline(v *EnvironmentSnapshot) *EnvironmentDeploymentSpec {
  s.Baseline = v
  return s
}

func (s *EnvironmentDeploymentSpec) SetChanges(v *EnvironmentChanges) *EnvironmentDeploymentSpec {
  s.Changes = v
  return s
}

func (s *EnvironmentDeploymentSpec) SetSkipRemoveResources(v bool) *EnvironmentDeploymentSpec {
  s.SkipRemoveResources = &v
  return s
}

func (s *EnvironmentDeploymentSpec) SetTarget(v *EnvironmentStagedConfigs) *EnvironmentDeploymentSpec {
  s.Target = v
  return s
}

func (s *EnvironmentDeploymentSpec) SetWebhookCodeContext(v *WebhookCodeContext) *EnvironmentDeploymentSpec {
  s.WebhookCodeContext = v
  return s
}

func (s *EnvironmentDeploymentSpec) Validate() error {
  return dara.Validate(s)
}

