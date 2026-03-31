// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDryRunConfigRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigurationItem(v string) *DryRunConfigRuleRequest
	GetConfigurationItem() *string
	SetResourceType(v string) *DryRunConfigRuleRequest
	GetResourceType() *string
}

type DryRunConfigRuleRequest struct {
	ConfigurationItem *string `json:"ConfigurationItem,omitempty" xml:"ConfigurationItem,omitempty"`
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DryRunConfigRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DryRunConfigRuleRequest) GoString() string {
	return s.String()
}

func (s *DryRunConfigRuleRequest) GetConfigurationItem() *string {
	return s.ConfigurationItem
}

func (s *DryRunConfigRuleRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DryRunConfigRuleRequest) SetConfigurationItem(v string) *DryRunConfigRuleRequest {
	s.ConfigurationItem = &v
	return s
}

func (s *DryRunConfigRuleRequest) SetResourceType(v string) *DryRunConfigRuleRequest {
	s.ResourceType = &v
	return s
}

func (s *DryRunConfigRuleRequest) Validate() error {
	return dara.Validate(s)
}
