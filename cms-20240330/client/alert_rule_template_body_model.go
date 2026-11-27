// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlertRuleTemplateBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *AlertRuleTemplateBody
	GetDescription() *string
	SetLabels(v string) *AlertRuleTemplateBody
	GetLabels() *string
	SetRuleConfigs(v string) *AlertRuleTemplateBody
	GetRuleConfigs() *string
	SetScenes(v string) *AlertRuleTemplateBody
	GetScenes() *string
	SetSchemaVersion(v string) *AlertRuleTemplateBody
	GetSchemaVersion() *string
	SetStatus(v int32) *AlertRuleTemplateBody
	GetStatus() *int32
	SetTemplateName(v string) *AlertRuleTemplateBody
	GetTemplateName() *string
}

type AlertRuleTemplateBody struct {
	// The template description.
	//
	// example:
	//
	// Alert template for sustained CPU usage above 80%
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The business-level key-value labels. This is a JSON object string, such as {"env":"prod"}. The default value is an empty object. An empty string is equivalent to not passing this parameter.
	//
	// example:
	//
	// {"env":"prod"}
	Labels *string `json:"labels,omitempty" xml:"labels,omitempty"`
	// The rule configurations. This is a v2 wire JSON object string, and the top level must include datasourceConfig.
	//
	// example:
	//
	// {"datasourceConfig":{"type":"PROMETHEUS","regionId":"cn-hangzhou"},"queryConfig":{"type":"PROMQL_QUERY","expr":"up == 0"}}
	RuleConfigs *string `json:"ruleConfigs,omitempty" xml:"ruleConfigs,omitempty"`
	// The scenario labels. This is a JSON array string, such as ["Runbook"]. The default value is an empty array. An empty string is equivalent to not passing this parameter.
	//
	// example:
	//
	// ["Runbook"]
	Scenes *string `json:"scenes,omitempty" xml:"scenes,omitempty"`
	// The schema version of the template. The default value is v2. v1 is rejected by the backend.
	//
	// example:
	//
	// v2
	SchemaVersion *string `json:"schemaVersion,omitempty" xml:"schemaVersion,omitempty"`
	// The enabling status of the template. Valid values: 1: Enabled. 0: Disabled. Default value: 1.
	//
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// The template name.
	//
	// example:
	//
	// prod-cpu-high-template
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
}

func (s AlertRuleTemplateBody) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleTemplateBody) GoString() string {
	return s.String()
}

func (s *AlertRuleTemplateBody) GetDescription() *string {
	return s.Description
}

func (s *AlertRuleTemplateBody) GetLabels() *string {
	return s.Labels
}

func (s *AlertRuleTemplateBody) GetRuleConfigs() *string {
	return s.RuleConfigs
}

func (s *AlertRuleTemplateBody) GetScenes() *string {
	return s.Scenes
}

func (s *AlertRuleTemplateBody) GetSchemaVersion() *string {
	return s.SchemaVersion
}

func (s *AlertRuleTemplateBody) GetStatus() *int32 {
	return s.Status
}

func (s *AlertRuleTemplateBody) GetTemplateName() *string {
	return s.TemplateName
}

func (s *AlertRuleTemplateBody) SetDescription(v string) *AlertRuleTemplateBody {
	s.Description = &v
	return s
}

func (s *AlertRuleTemplateBody) SetLabels(v string) *AlertRuleTemplateBody {
	s.Labels = &v
	return s
}

func (s *AlertRuleTemplateBody) SetRuleConfigs(v string) *AlertRuleTemplateBody {
	s.RuleConfigs = &v
	return s
}

func (s *AlertRuleTemplateBody) SetScenes(v string) *AlertRuleTemplateBody {
	s.Scenes = &v
	return s
}

func (s *AlertRuleTemplateBody) SetSchemaVersion(v string) *AlertRuleTemplateBody {
	s.SchemaVersion = &v
	return s
}

func (s *AlertRuleTemplateBody) SetStatus(v int32) *AlertRuleTemplateBody {
	s.Status = &v
	return s
}

func (s *AlertRuleTemplateBody) SetTemplateName(v string) *AlertRuleTemplateBody {
	s.TemplateName = &v
	return s
}

func (s *AlertRuleTemplateBody) Validate() error {
	return dara.Validate(s)
}
