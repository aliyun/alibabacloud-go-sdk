// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataQualityRuleTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckingConfigShrink(v string) *CreateDataQualityRuleTemplateShrinkRequest
	GetCheckingConfigShrink() *string
	SetDirectoryPath(v string) *CreateDataQualityRuleTemplateShrinkRequest
	GetDirectoryPath() *string
	SetName(v string) *CreateDataQualityRuleTemplateShrinkRequest
	GetName() *string
	SetProjectId(v int64) *CreateDataQualityRuleTemplateShrinkRequest
	GetProjectId() *int64
	SetSamplingConfigShrink(v string) *CreateDataQualityRuleTemplateShrinkRequest
	GetSamplingConfigShrink() *string
	SetVisibleScope(v string) *CreateDataQualityRuleTemplateShrinkRequest
	GetVisibleScope() *string
}

type CreateDataQualityRuleTemplateShrinkRequest struct {
	// The sample verification settings.
	CheckingConfigShrink *string `json:"CheckingConfig,omitempty" xml:"CheckingConfig,omitempty"`
	// The directory path where the custom template is stored. Levels are separated by forward slashes (/). Each level name can be up to 1024 characters in length and cannot contain whitespace characters or forward slashes.
	//
	// example:
	//
	// /ods/order_data
	DirectoryPath *string `json:"DirectoryPath,omitempty" xml:"DirectoryPath,omitempty"`
	// The name of the rule template. The name can contain digits, letters, Chinese characters, and half-width or full-width punctuation marks. The name can be up to 512 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// Table row Count Verification
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The DataWorks workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The settings required for sample collection.
	SamplingConfigShrink *string `json:"SamplingConfig,omitempty" xml:"SamplingConfig,omitempty"`
	// The visibility scope of the template. Valid values:
	//
	// - Tenant: available to the entire tenant.
	//
	// - Project: available only in the current project.
	//
	// example:
	//
	// Project
	VisibleScope *string `json:"VisibleScope,omitempty" xml:"VisibleScope,omitempty"`
}

func (s CreateDataQualityRuleTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityRuleTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDataQualityRuleTemplateShrinkRequest) GetCheckingConfigShrink() *string {
	return s.CheckingConfigShrink
}

func (s *CreateDataQualityRuleTemplateShrinkRequest) GetDirectoryPath() *string {
	return s.DirectoryPath
}

func (s *CreateDataQualityRuleTemplateShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateDataQualityRuleTemplateShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateDataQualityRuleTemplateShrinkRequest) GetSamplingConfigShrink() *string {
	return s.SamplingConfigShrink
}

func (s *CreateDataQualityRuleTemplateShrinkRequest) GetVisibleScope() *string {
	return s.VisibleScope
}

func (s *CreateDataQualityRuleTemplateShrinkRequest) SetCheckingConfigShrink(v string) *CreateDataQualityRuleTemplateShrinkRequest {
	s.CheckingConfigShrink = &v
	return s
}

func (s *CreateDataQualityRuleTemplateShrinkRequest) SetDirectoryPath(v string) *CreateDataQualityRuleTemplateShrinkRequest {
	s.DirectoryPath = &v
	return s
}

func (s *CreateDataQualityRuleTemplateShrinkRequest) SetName(v string) *CreateDataQualityRuleTemplateShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateDataQualityRuleTemplateShrinkRequest) SetProjectId(v int64) *CreateDataQualityRuleTemplateShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDataQualityRuleTemplateShrinkRequest) SetSamplingConfigShrink(v string) *CreateDataQualityRuleTemplateShrinkRequest {
	s.SamplingConfigShrink = &v
	return s
}

func (s *CreateDataQualityRuleTemplateShrinkRequest) SetVisibleScope(v string) *CreateDataQualityRuleTemplateShrinkRequest {
	s.VisibleScope = &v
	return s
}

func (s *CreateDataQualityRuleTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
