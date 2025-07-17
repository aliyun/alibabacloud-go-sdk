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
	SetId(v int64) *UpdateDIJobShrinkRequest
	GetId() *int64
	SetJobSettingsShrink(v string) *UpdateDIJobShrinkRequest
	GetJobSettingsShrink() *string
	SetProjectId(v int64) *UpdateDIJobShrinkRequest
	GetProjectId() *int64
	SetResourceSettingsShrink(v string) *UpdateDIJobShrinkRequest
	GetResourceSettingsShrink() *string
	SetTableMappingsShrink(v string) *UpdateDIJobShrinkRequest
	GetTableMappingsShrink() *string
	SetTransformationRulesShrink(v string) *UpdateDIJobShrinkRequest
	GetTransformationRulesShrink() *string
}

type UpdateDIJobShrinkRequest struct {
	// Deprecated
	//
	// This parameter is deprecated. Use the Id parameter instead.
	//
	// example:
	//
	// 11588
	DIJobId     *int64  `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the synchronization task.
	//
	// example:
	//
	// 11588
	Id                *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	JobSettingsShrink *string `json:"JobSettings,omitempty" xml:"JobSettings,omitempty"`
	// The DataWorks workspace ID. You can call the [ListProjects](https://help.aliyun.com/document_detail/178393.html) operation to obtain the ID.
	//
	// example:
	//
	// 10000
	ProjectId                 *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ResourceSettingsShrink    *string `json:"ResourceSettings,omitempty" xml:"ResourceSettings,omitempty"`
	TableMappingsShrink       *string `json:"TableMappings,omitempty" xml:"TableMappings,omitempty"`
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

func (s *UpdateDIJobShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateDIJobShrinkRequest) GetJobSettingsShrink() *string {
	return s.JobSettingsShrink
}

func (s *UpdateDIJobShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
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

func (s *UpdateDIJobShrinkRequest) SetId(v int64) *UpdateDIJobShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateDIJobShrinkRequest) SetJobSettingsShrink(v string) *UpdateDIJobShrinkRequest {
	s.JobSettingsShrink = &v
	return s
}

func (s *UpdateDIJobShrinkRequest) SetProjectId(v int64) *UpdateDIJobShrinkRequest {
	s.ProjectId = &v
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
