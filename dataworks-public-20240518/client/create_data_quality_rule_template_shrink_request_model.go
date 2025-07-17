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
	// The check settings for sample data.
	CheckingConfigShrink *string `json:"CheckingConfig,omitempty" xml:"CheckingConfig,omitempty"`
	// The directory in which the template is stored. Slashes (/) are used to separate directory levels. The name of each directory level can be up to 1,024 characters in length. It cannot contain whitespace characters or slashes (/).
	//
	// example:
	//
	// /ods/order_data
	DirectoryPath *string `json:"DirectoryPath,omitempty" xml:"DirectoryPath,omitempty"`
	// The name of the template. The name can be up to 512 characters in length and can contain digits, letters, and punctuation marks.
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
	// The sampling settings.
	SamplingConfigShrink *string `json:"SamplingConfig,omitempty" xml:"SamplingConfig,omitempty"`
	// The applicable scope of the template. Valid values:
	//
	// 	- Tenant: The template is available in all workspaces in the current tenant.
	//
	// 	- Project: The template is available only in the current workspace.
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
