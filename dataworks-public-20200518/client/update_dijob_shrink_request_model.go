// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDIJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDIJobId(v int64) *UpdateDIJobShrinkRequest
	GetDIJobId() *int64
	SetDescription(v string) *UpdateDIJobShrinkRequest
	GetDescription() *string
	SetJobSettingsShrink(v string) *UpdateDIJobShrinkRequest
	GetJobSettingsShrink() *string
	SetResourceSettingsShrink(v string) *UpdateDIJobShrinkRequest
	GetResourceSettingsShrink() *string
	SetTableMappingsShrink(v string) *UpdateDIJobShrinkRequest
	GetTableMappingsShrink() *string
	SetTransformationRulesShrink(v string) *UpdateDIJobShrinkRequest
	GetTransformationRulesShrink() *string
}

type UpdateDIJobShrinkRequest struct {
	// The ID of the synchronization task.
	//
	// example:
	//
	// 11588
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// The description of the synchronization task.
	//
	// example:
	//
	// Synchronize mysql to hologres
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The settings for the dimension of the synchronization task. The settings include processing policies for DDL messages, policies for data type mappings between source fields and destination fields, and runtime parameters of the synchronization task.
	JobSettingsShrink *string `json:"JobSettings,omitempty" xml:"JobSettings,omitempty"`
	// The resource settings.
	ResourceSettingsShrink *string `json:"ResourceSettings,omitempty" xml:"ResourceSettings,omitempty"`
	// The list of mappings between rules used to select synchronization objects in the source and transformation rules applied to the selected synchronization objects. Each entry in the list displays a mapping between a rule used to select synchronization objects and a transformation rule applied to the selected synchronization objects.
	TableMappingsShrink *string `json:"TableMappings,omitempty" xml:"TableMappings,omitempty"`
	// The list of transformation rules that you want to apply to the synchronization objects selected from the source. Each entry in the list defines a transformation rule.
	TransformationRulesShrink *string `json:"TransformationRules,omitempty" xml:"TransformationRules,omitempty"`
}

func (s UpdateDIJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDIJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDIJobShrinkRequest) GetDIJobId() *int64 {
	return s.DIJobId
}

func (s *UpdateDIJobShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDIJobShrinkRequest) GetJobSettingsShrink() *string {
	return s.JobSettingsShrink
}

func (s *UpdateDIJobShrinkRequest) GetResourceSettingsShrink() *string {
	return s.ResourceSettingsShrink
}

func (s *UpdateDIJobShrinkRequest) GetTableMappingsShrink() *string {
	return s.TableMappingsShrink
}

func (s *UpdateDIJobShrinkRequest) GetTransformationRulesShrink() *string {
	return s.TransformationRulesShrink
}

func (s *UpdateDIJobShrinkRequest) SetDIJobId(v int64) *UpdateDIJobShrinkRequest {
	s.DIJobId = &v
	return s
}

func (s *UpdateDIJobShrinkRequest) SetDescription(v string) *UpdateDIJobShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateDIJobShrinkRequest) SetJobSettingsShrink(v string) *UpdateDIJobShrinkRequest {
	s.JobSettingsShrink = &v
	return s
}

func (s *UpdateDIJobShrinkRequest) SetResourceSettingsShrink(v string) *UpdateDIJobShrinkRequest {
	s.ResourceSettingsShrink = &v
	return s
}

func (s *UpdateDIJobShrinkRequest) SetTableMappingsShrink(v string) *UpdateDIJobShrinkRequest {
	s.TableMappingsShrink = &v
	return s
}

func (s *UpdateDIJobShrinkRequest) SetTransformationRulesShrink(v string) *UpdateDIJobShrinkRequest {
	s.TransformationRulesShrink = &v
	return s
}

func (s *UpdateDIJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
