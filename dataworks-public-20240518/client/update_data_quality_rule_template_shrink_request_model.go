// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataQualityRuleTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckingConfigShrink(v string) *UpdateDataQualityRuleTemplateShrinkRequest
	GetCheckingConfigShrink() *string
	SetCode(v string) *UpdateDataQualityRuleTemplateShrinkRequest
	GetCode() *string
	SetDirectoryPath(v string) *UpdateDataQualityRuleTemplateShrinkRequest
	GetDirectoryPath() *string
	SetName(v string) *UpdateDataQualityRuleTemplateShrinkRequest
	GetName() *string
	SetProjectId(v int64) *UpdateDataQualityRuleTemplateShrinkRequest
	GetProjectId() *int64
	SetSamplingConfigShrink(v string) *UpdateDataQualityRuleTemplateShrinkRequest
	GetSamplingConfigShrink() *string
}

type UpdateDataQualityRuleTemplateShrinkRequest struct {
	// The check settings for sample data.
	CheckingConfigShrink *string `json:"CheckingConfig,omitempty" xml:"CheckingConfig,omitempty"`
	// The code for the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// USER_DEFINED:123
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The directory in which the template is stored. Slashes (/) are used to separate directory levels. The name of each directory level can be up to 1,024 characters in length. It cannot contain whitespace characters or slashes (/).
	//
	// example:
	//
	// /ods/order_data
	DirectoryPath *string `json:"DirectoryPath,omitempty" xml:"DirectoryPath,omitempty"`
	// The name of the template. The name can be up to 512 characters in length and can contain digits, letters, and punctuation marks.
	//
	// example:
	//
	// Table row Count Verification
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The sampling settings.
	SamplingConfigShrink *string `json:"SamplingConfig,omitempty" xml:"SamplingConfig,omitempty"`
}

func (s UpdateDataQualityRuleTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityRuleTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityRuleTemplateShrinkRequest) GetCheckingConfigShrink() *string {
	return s.CheckingConfigShrink
}

func (s *UpdateDataQualityRuleTemplateShrinkRequest) GetCode() *string {
	return s.Code
}

func (s *UpdateDataQualityRuleTemplateShrinkRequest) GetDirectoryPath() *string {
	return s.DirectoryPath
}

func (s *UpdateDataQualityRuleTemplateShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateDataQualityRuleTemplateShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateDataQualityRuleTemplateShrinkRequest) GetSamplingConfigShrink() *string {
	return s.SamplingConfigShrink
}

func (s *UpdateDataQualityRuleTemplateShrinkRequest) SetCheckingConfigShrink(v string) *UpdateDataQualityRuleTemplateShrinkRequest {
	s.CheckingConfigShrink = &v
	return s
}

func (s *UpdateDataQualityRuleTemplateShrinkRequest) SetCode(v string) *UpdateDataQualityRuleTemplateShrinkRequest {
	s.Code = &v
	return s
}

func (s *UpdateDataQualityRuleTemplateShrinkRequest) SetDirectoryPath(v string) *UpdateDataQualityRuleTemplateShrinkRequest {
	s.DirectoryPath = &v
	return s
}

func (s *UpdateDataQualityRuleTemplateShrinkRequest) SetName(v string) *UpdateDataQualityRuleTemplateShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateDataQualityRuleTemplateShrinkRequest) SetProjectId(v int64) *UpdateDataQualityRuleTemplateShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateDataQualityRuleTemplateShrinkRequest) SetSamplingConfigShrink(v string) *UpdateDataQualityRuleTemplateShrinkRequest {
	s.SamplingConfigShrink = &v
	return s
}

func (s *UpdateDataQualityRuleTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
