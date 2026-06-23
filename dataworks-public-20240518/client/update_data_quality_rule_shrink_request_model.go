// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataQualityRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckingConfigShrink(v string) *UpdateDataQualityRuleShrinkRequest
	GetCheckingConfigShrink() *string
	SetDescription(v string) *UpdateDataQualityRuleShrinkRequest
	GetDescription() *string
	SetEnabled(v bool) *UpdateDataQualityRuleShrinkRequest
	GetEnabled() *bool
	SetErrorHandlersShrink(v string) *UpdateDataQualityRuleShrinkRequest
	GetErrorHandlersShrink() *string
	SetId(v int64) *UpdateDataQualityRuleShrinkRequest
	GetId() *int64
	SetName(v string) *UpdateDataQualityRuleShrinkRequest
	GetName() *string
	SetProjectId(v int64) *UpdateDataQualityRuleShrinkRequest
	GetProjectId() *int64
	SetSamplingConfigShrink(v string) *UpdateDataQualityRuleShrinkRequest
	GetSamplingConfigShrink() *string
	SetSeverity(v string) *UpdateDataQualityRuleShrinkRequest
	GetSeverity() *string
	SetTemplateCode(v string) *UpdateDataQualityRuleShrinkRequest
	GetTemplateCode() *string
}

type UpdateDataQualityRuleShrinkRequest struct {
	// The sample verification settings.
	CheckingConfigShrink *string `json:"CheckingConfig,omitempty" xml:"CheckingConfig,omitempty"`
	// The rule description. The maximum length is 500 characters.
	//
	// example:
	//
	// this is a odps _sql task
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether the rule is enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The list of issue handlers for data quality rule verification.
	ErrorHandlersShrink *string `json:"ErrorHandlers,omitempty" xml:"ErrorHandlers,omitempty"`
	// The rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The rule name. The name can be a combination of digits, English letters, Chinese characters, and half-width or full-width punctuation. The maximum length is 255 characters.
	//
	// example:
	//
	// The table cannot be empty.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace Settings page to obtain the workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The settings required for sample collection.
	SamplingConfigShrink *string `json:"SamplingConfig,omitempty" xml:"SamplingConfig,omitempty"`
	// The severity level of the rule for the business (corresponding to strong/weak rules on the page). Valid values:
	//
	// - Normal
	//
	// - High
	//
	// example:
	//
	// High
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The unique identifier of the rule template referenced by the rule.
	//
	// example:
	//
	// SYSTEM:table:table_count:fixed
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s UpdateDataQualityRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityRuleShrinkRequest) GetCheckingConfigShrink() *string {
	return s.CheckingConfigShrink
}

func (s *UpdateDataQualityRuleShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDataQualityRuleShrinkRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateDataQualityRuleShrinkRequest) GetErrorHandlersShrink() *string {
	return s.ErrorHandlersShrink
}

func (s *UpdateDataQualityRuleShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateDataQualityRuleShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateDataQualityRuleShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateDataQualityRuleShrinkRequest) GetSamplingConfigShrink() *string {
	return s.SamplingConfigShrink
}

func (s *UpdateDataQualityRuleShrinkRequest) GetSeverity() *string {
	return s.Severity
}

func (s *UpdateDataQualityRuleShrinkRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *UpdateDataQualityRuleShrinkRequest) SetCheckingConfigShrink(v string) *UpdateDataQualityRuleShrinkRequest {
	s.CheckingConfigShrink = &v
	return s
}

func (s *UpdateDataQualityRuleShrinkRequest) SetDescription(v string) *UpdateDataQualityRuleShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateDataQualityRuleShrinkRequest) SetEnabled(v bool) *UpdateDataQualityRuleShrinkRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateDataQualityRuleShrinkRequest) SetErrorHandlersShrink(v string) *UpdateDataQualityRuleShrinkRequest {
	s.ErrorHandlersShrink = &v
	return s
}

func (s *UpdateDataQualityRuleShrinkRequest) SetId(v int64) *UpdateDataQualityRuleShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateDataQualityRuleShrinkRequest) SetName(v string) *UpdateDataQualityRuleShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateDataQualityRuleShrinkRequest) SetProjectId(v int64) *UpdateDataQualityRuleShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateDataQualityRuleShrinkRequest) SetSamplingConfigShrink(v string) *UpdateDataQualityRuleShrinkRequest {
	s.SamplingConfigShrink = &v
	return s
}

func (s *UpdateDataQualityRuleShrinkRequest) SetSeverity(v string) *UpdateDataQualityRuleShrinkRequest {
	s.Severity = &v
	return s
}

func (s *UpdateDataQualityRuleShrinkRequest) SetTemplateCode(v string) *UpdateDataQualityRuleShrinkRequest {
	s.TemplateCode = &v
	return s
}

func (s *UpdateDataQualityRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
