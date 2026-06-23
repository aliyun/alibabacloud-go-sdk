// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataQualityRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckingConfigShrink(v string) *CreateDataQualityRuleShrinkRequest
	GetCheckingConfigShrink() *string
	SetDescription(v string) *CreateDataQualityRuleShrinkRequest
	GetDescription() *string
	SetEnabled(v bool) *CreateDataQualityRuleShrinkRequest
	GetEnabled() *bool
	SetErrorHandlersShrink(v string) *CreateDataQualityRuleShrinkRequest
	GetErrorHandlersShrink() *string
	SetName(v string) *CreateDataQualityRuleShrinkRequest
	GetName() *string
	SetProjectId(v int64) *CreateDataQualityRuleShrinkRequest
	GetProjectId() *int64
	SetSamplingConfigShrink(v string) *CreateDataQualityRuleShrinkRequest
	GetSamplingConfigShrink() *string
	SetSeverity(v string) *CreateDataQualityRuleShrinkRequest
	GetSeverity() *string
	SetTargetShrink(v string) *CreateDataQualityRuleShrinkRequest
	GetTargetShrink() *string
	SetTemplateCode(v string) *CreateDataQualityRuleShrinkRequest
	GetTemplateCode() *string
}

type CreateDataQualityRuleShrinkRequest struct {
	// The sample check settings.
	CheckingConfigShrink *string `json:"CheckingConfig,omitempty" xml:"CheckingConfig,omitempty"`
	// The description of the rule. The maximum length is 500 characters.
	//
	// example:
	//
	// this is a odps _sql task
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable the data quality rule.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The list of issue handlers for the data quality rule check.
	ErrorHandlersShrink *string `json:"ErrorHandlers,omitempty" xml:"ErrorHandlers,omitempty"`
	// The name of the rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// The table cannot be empty.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the DataWorks workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10726
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The settings required for sample collection.
	SamplingConfigShrink *string `json:"SamplingConfig,omitempty" xml:"SamplingConfig,omitempty"`
	// The severity of the rule for the business (corresponding to the strong/weak rule on the page). Valid values:
	//
	// - Normal
	//
	// - High
	//
	// example:
	//
	// Normal
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The object monitored by the rule.
	TargetShrink *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The unique identifier of the rule template that the rule references.
	//
	// example:
	//
	// SYSTEM:table:table_count:fixed
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s CreateDataQualityRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDataQualityRuleShrinkRequest) GetCheckingConfigShrink() *string {
	return s.CheckingConfigShrink
}

func (s *CreateDataQualityRuleShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDataQualityRuleShrinkRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateDataQualityRuleShrinkRequest) GetErrorHandlersShrink() *string {
	return s.ErrorHandlersShrink
}

func (s *CreateDataQualityRuleShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateDataQualityRuleShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateDataQualityRuleShrinkRequest) GetSamplingConfigShrink() *string {
	return s.SamplingConfigShrink
}

func (s *CreateDataQualityRuleShrinkRequest) GetSeverity() *string {
	return s.Severity
}

func (s *CreateDataQualityRuleShrinkRequest) GetTargetShrink() *string {
	return s.TargetShrink
}

func (s *CreateDataQualityRuleShrinkRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *CreateDataQualityRuleShrinkRequest) SetCheckingConfigShrink(v string) *CreateDataQualityRuleShrinkRequest {
	s.CheckingConfigShrink = &v
	return s
}

func (s *CreateDataQualityRuleShrinkRequest) SetDescription(v string) *CreateDataQualityRuleShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateDataQualityRuleShrinkRequest) SetEnabled(v bool) *CreateDataQualityRuleShrinkRequest {
	s.Enabled = &v
	return s
}

func (s *CreateDataQualityRuleShrinkRequest) SetErrorHandlersShrink(v string) *CreateDataQualityRuleShrinkRequest {
	s.ErrorHandlersShrink = &v
	return s
}

func (s *CreateDataQualityRuleShrinkRequest) SetName(v string) *CreateDataQualityRuleShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateDataQualityRuleShrinkRequest) SetProjectId(v int64) *CreateDataQualityRuleShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDataQualityRuleShrinkRequest) SetSamplingConfigShrink(v string) *CreateDataQualityRuleShrinkRequest {
	s.SamplingConfigShrink = &v
	return s
}

func (s *CreateDataQualityRuleShrinkRequest) SetSeverity(v string) *CreateDataQualityRuleShrinkRequest {
	s.Severity = &v
	return s
}

func (s *CreateDataQualityRuleShrinkRequest) SetTargetShrink(v string) *CreateDataQualityRuleShrinkRequest {
	s.TargetShrink = &v
	return s
}

func (s *CreateDataQualityRuleShrinkRequest) SetTemplateCode(v string) *CreateDataQualityRuleShrinkRequest {
	s.TemplateCode = &v
	return s
}

func (s *CreateDataQualityRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
