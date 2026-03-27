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
	SetFileSpec(v string) *UpdateDIJobShrinkRequest
	GetFileSpec() *string
	SetId(v int64) *UpdateDIJobShrinkRequest
	GetId() *int64
	SetJobSettingsShrink(v string) *UpdateDIJobShrinkRequest
	GetJobSettingsShrink() *string
	SetOwner(v string) *UpdateDIJobShrinkRequest
	GetOwner() *string
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
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// The task description.
	//
	// example:
	//
	// The description of the synchronization task.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FileSpec    *string `json:"FileSpec,omitempty" xml:"FileSpec,omitempty"`
	// The ID of the synchronization task.
	//
	// example:
	//
	// 11588
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The task-level settings, including DDL handling policies, column data type mapping between source and destination, and runtime parameters.
	JobSettingsShrink *string `json:"JobSettings,omitempty" xml:"JobSettings,omitempty"`
	// The task owner.
	//
	// example:
	//
	// 95279527
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The DataWorks workspace ID. You can call the [ListProjects](https://help.aliyun.com/document_detail/178393.html) operation to obtain the ID.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The resource settings.
	ResourceSettingsShrink *string `json:"ResourceSettings,omitempty" xml:"ResourceSettings,omitempty"`
	// The list of synchronization object transformation mappings. Each element describes a set of source object selection rules and the transformation rules applied to those objects.
	//
	// >  [ { "SourceObjectSelectionRules":[ { "ObjectType":"Database", "Action":"Include", "ExpressionType":"Exact", "Expression":"biz_db" }, { "ObjectType":"Schema", "Action":"Include", "ExpressionType":"Exact", "Expression":"s1" }, { "ObjectType":"Table", "Action":"Include", "ExpressionType":"Exact", "Expression":"table1" } ], "TransformationRuleNames":[ { "RuleName":"my_database_rename_rule", "RuleActionType":"Rename", "RuleTargetType":"Schema" } ] } ]
	TableMappingsShrink *string `json:"TableMappings,omitempty" xml:"TableMappings,omitempty"`
	// The list of synchronization object transformation rule definitions.
	//
	// >  [ { "RuleName":"my_database_rename_rule", "RuleActionType":"Rename", "RuleTargetType":"Schema", "RuleExpression":"{"expression":"${srcDatasoureName}_${srcDatabaseName}"}" } ]
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

func (s *UpdateDIJobShrinkRequest) GetFileSpec() *string {
	return s.FileSpec
}

func (s *UpdateDIJobShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateDIJobShrinkRequest) GetJobSettingsShrink() *string {
	return s.JobSettingsShrink
}

func (s *UpdateDIJobShrinkRequest) GetOwner() *string {
	return s.Owner
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

func (s *UpdateDIJobShrinkRequest) SetFileSpec(v string) *UpdateDIJobShrinkRequest {
	s.FileSpec = &v
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

func (s *UpdateDIJobShrinkRequest) SetOwner(v string) *UpdateDIJobShrinkRequest {
	s.Owner = &v
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
