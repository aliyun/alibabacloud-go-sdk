// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolicyDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *DescribePolicyDetailsResponseBody
	GetAction() *string
	SetCategory(v string) *DescribePolicyDetailsResponseBody
	GetCategory() *string
	SetDescription(v string) *DescribePolicyDetailsResponseBody
	GetDescription() *string
	SetIsDeleted(v int32) *DescribePolicyDetailsResponseBody
	GetIsDeleted() *int32
	SetName(v string) *DescribePolicyDetailsResponseBody
	GetName() *string
	SetNoConfig(v int32) *DescribePolicyDetailsResponseBody
	GetNoConfig() *int32
	SetSeverity(v string) *DescribePolicyDetailsResponseBody
	GetSeverity() *string
	SetTemplate(v string) *DescribePolicyDetailsResponseBody
	GetTemplate() *string
}

type DescribePolicyDetailsResponseBody struct {
	// The governance action of the rule. Valid values:
	//
	// - `enforce`: blocks non-compliant deployments.
	//
	// - `inform`: generates alerts.
	//
	// example:
	//
	// enforce
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// The rule templatetype.
	//
	// example:
	//
	// k8s-general
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// The description of the rule template.
	//
	// example:
	//
	// Requires container images to begin with a repo string from a specified list
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Indicates whether the rule is deleted. Valid values:
	//
	// - 0: The rule is not deleted.
	//
	// - 1: The rule is deleted.
	//
	// example:
	//
	// 0
	IsDeleted *int32 `json:"is_deleted,omitempty" xml:"is_deleted,omitempty"`
	// The name of the policy governance rule.
	//
	// example:
	//
	// ACKAllowedRepos
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Indicates whether policy configuration is required. Valid values:
	//
	// - 0: Parameter settings are required.
	//
	// - 1: No parameter settings are required.
	//
	// example:
	//
	// 0
	NoConfig *int32 `json:"no_config,omitempty" xml:"no_config,omitempty"`
	// The governance severity level of the rule. Valid values:
	//
	// 	- `high`: high risk.
	//
	// 	- `medium`: medium risk.
	//
	// 	- `low`: low risk.
	//
	// example:
	//
	// high
	Severity *string `json:"severity,omitempty" xml:"severity,omitempty"`
	// The details of the rule template.
	//
	// example:
	//
	// 详情请参见请求示例
	Template *string `json:"template,omitempty" xml:"template,omitempty"`
}

func (s DescribePolicyDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolicyDetailsResponseBody) GetAction() *string {
	return s.Action
}

func (s *DescribePolicyDetailsResponseBody) GetCategory() *string {
	return s.Category
}

func (s *DescribePolicyDetailsResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribePolicyDetailsResponseBody) GetIsDeleted() *int32 {
	return s.IsDeleted
}

func (s *DescribePolicyDetailsResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribePolicyDetailsResponseBody) GetNoConfig() *int32 {
	return s.NoConfig
}

func (s *DescribePolicyDetailsResponseBody) GetSeverity() *string {
	return s.Severity
}

func (s *DescribePolicyDetailsResponseBody) GetTemplate() *string {
	return s.Template
}

func (s *DescribePolicyDetailsResponseBody) SetAction(v string) *DescribePolicyDetailsResponseBody {
	s.Action = &v
	return s
}

func (s *DescribePolicyDetailsResponseBody) SetCategory(v string) *DescribePolicyDetailsResponseBody {
	s.Category = &v
	return s
}

func (s *DescribePolicyDetailsResponseBody) SetDescription(v string) *DescribePolicyDetailsResponseBody {
	s.Description = &v
	return s
}

func (s *DescribePolicyDetailsResponseBody) SetIsDeleted(v int32) *DescribePolicyDetailsResponseBody {
	s.IsDeleted = &v
	return s
}

func (s *DescribePolicyDetailsResponseBody) SetName(v string) *DescribePolicyDetailsResponseBody {
	s.Name = &v
	return s
}

func (s *DescribePolicyDetailsResponseBody) SetNoConfig(v int32) *DescribePolicyDetailsResponseBody {
	s.NoConfig = &v
	return s
}

func (s *DescribePolicyDetailsResponseBody) SetSeverity(v string) *DescribePolicyDetailsResponseBody {
	s.Severity = &v
	return s
}

func (s *DescribePolicyDetailsResponseBody) SetTemplate(v string) *DescribePolicyDetailsResponseBody {
	s.Template = &v
	return s
}

func (s *DescribePolicyDetailsResponseBody) Validate() error {
	return dara.Validate(s)
}
