// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecurityClassifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateSecurityClassifyRequest
	GetOpTenantId() *int64
	SetUpdateCommand(v *UpdateSecurityClassifyRequestUpdateCommand) *UpdateSecurityClassifyRequest
	GetUpdateCommand() *UpdateSecurityClassifyRequestUpdateCommand
}

type UpdateSecurityClassifyRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The update command.
	//
	// This parameter is required.
	UpdateCommand *UpdateSecurityClassifyRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateSecurityClassifyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecurityClassifyRequest) GoString() string {
	return s.String()
}

func (s *UpdateSecurityClassifyRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateSecurityClassifyRequest) GetUpdateCommand() *UpdateSecurityClassifyRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateSecurityClassifyRequest) SetOpTenantId(v int64) *UpdateSecurityClassifyRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateSecurityClassifyRequest) SetUpdateCommand(v *UpdateSecurityClassifyRequestUpdateCommand) *UpdateSecurityClassifyRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdateSecurityClassifyRequest) Validate() error {
	if s.UpdateCommand != nil {
		if err := s.UpdateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateSecurityClassifyRequestUpdateCommand struct {
	// The abbreviation of the classification.
	//
	// example:
	//
	// test
	Abbreviation *string `json:"Abbreviation,omitempty" xml:"Abbreviation,omitempty"`
	// The advanced configurations.
	AdvancedConditionList []*UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionList `json:"AdvancedConditionList,omitempty" xml:"AdvancedConditionList,omitempty" type:"Repeated"`
	// The description of the classification.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The feature names.
	FeatureNameList []*string `json:"FeatureNameList,omitempty" xml:"FeatureNameList,omitempty" type:"Repeated"`
	// The classification ID. If this parameter is not specified, the classification is matched and updated based on the classification name and parent path.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The data level name.
	//
	// This parameter is required.
	//
	// example:
	//
	// L2
	LevelName *string `json:"LevelName,omitempty" xml:"LevelName,omitempty"`
	// The classification name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parent path of the classification.
	//
	// example:
	//
	// /a/b/
	ParentPath *string `json:"ParentPath,omitempty" xml:"ParentPath,omitempty"`
	// The priority. Default value: 5.
	//
	// example:
	//
	// 5
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The classification status. Valid values:
	//
	// - ENABLE: enabled.
	//
	// - DISABLE: disabled.
	//
	// example:
	//
	// ENABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateSecurityClassifyRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecurityClassifyRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateSecurityClassifyRequestUpdateCommand) GetAbbreviation() *string {
	return s.Abbreviation
}

func (s *UpdateSecurityClassifyRequestUpdateCommand) GetAdvancedConditionList() []*UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionList {
	return s.AdvancedConditionList
}

func (s *UpdateSecurityClassifyRequestUpdateCommand) GetDescription() *string {
	return s.Description
}

func (s *UpdateSecurityClassifyRequestUpdateCommand) GetFeatureNameList() []*string {
	return s.FeatureNameList
}

func (s *UpdateSecurityClassifyRequestUpdateCommand) GetId() *int64 {
	return s.Id
}

func (s *UpdateSecurityClassifyRequestUpdateCommand) GetLevelName() *string {
	return s.LevelName
}

func (s *UpdateSecurityClassifyRequestUpdateCommand) GetName() *string {
	return s.Name
}

func (s *UpdateSecurityClassifyRequestUpdateCommand) GetParentPath() *string {
	return s.ParentPath
}

func (s *UpdateSecurityClassifyRequestUpdateCommand) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateSecurityClassifyRequestUpdateCommand) GetStatus() *string {
	return s.Status
}

func (s *UpdateSecurityClassifyRequestUpdateCommand) SetAbbreviation(v string) *UpdateSecurityClassifyRequestUpdateCommand {
	s.Abbreviation = &v
	return s
}

func (s *UpdateSecurityClassifyRequestUpdateCommand) SetAdvancedConditionList(v []*UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionList) *UpdateSecurityClassifyRequestUpdateCommand {
	s.AdvancedConditionList = v
	return s
}

func (s *UpdateSecurityClassifyRequestUpdateCommand) SetDescription(v string) *UpdateSecurityClassifyRequestUpdateCommand {
	s.Description = &v
	return s
}

func (s *UpdateSecurityClassifyRequestUpdateCommand) SetFeatureNameList(v []*string) *UpdateSecurityClassifyRequestUpdateCommand {
	s.FeatureNameList = v
	return s
}

func (s *UpdateSecurityClassifyRequestUpdateCommand) SetId(v int64) *UpdateSecurityClassifyRequestUpdateCommand {
	s.Id = &v
	return s
}

func (s *UpdateSecurityClassifyRequestUpdateCommand) SetLevelName(v string) *UpdateSecurityClassifyRequestUpdateCommand {
	s.LevelName = &v
	return s
}

func (s *UpdateSecurityClassifyRequestUpdateCommand) SetName(v string) *UpdateSecurityClassifyRequestUpdateCommand {
	s.Name = &v
	return s
}

func (s *UpdateSecurityClassifyRequestUpdateCommand) SetParentPath(v string) *UpdateSecurityClassifyRequestUpdateCommand {
	s.ParentPath = &v
	return s
}

func (s *UpdateSecurityClassifyRequestUpdateCommand) SetPriority(v int32) *UpdateSecurityClassifyRequestUpdateCommand {
	s.Priority = &v
	return s
}

func (s *UpdateSecurityClassifyRequestUpdateCommand) SetStatus(v string) *UpdateSecurityClassifyRequestUpdateCommand {
	s.Status = &v
	return s
}

func (s *UpdateSecurityClassifyRequestUpdateCommand) Validate() error {
	if s.AdvancedConditionList != nil {
		for _, item := range s.AdvancedConditionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionList struct {
	// The ID of the feature condition node. The ID is randomly generated by the application and must be unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// 234
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The condition operator. Valid values:
	//
	// - AND
	//
	// - OR
	//
	// - BUILT_IN_EXPRESSION: built-in expression.
	//
	// - IGNORE_CASE_EXPRESSION: regular expression with case-insensitive matching.
	//
	// - EXPRESSION: regular expression.
	//
	// - NOT_BELONG: does not belong to.
	//
	// - CONTAINS: contains.
	//
	// - NOT_CONTAINS: does not contain.
	//
	// - EQUAL
	//
	// - NOT_EQUAL.
	//
	// This parameter is required.
	//
	// example:
	//
	// AND
	Operate *string `json:"Operate,omitempty" xml:"Operate,omitempty"`
	// The additional properties, such as field content.
	OptionList []*UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionListOptionList `json:"OptionList,omitempty" xml:"OptionList,omitempty" type:"Repeated"`
	// The parent node of the feature condition.
	//
	// example:
	//
	// 123
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The feature property. Valid values:
	//
	// - TABLE_NAME: table name.
	//
	// - TABLE_DESC: table description.
	//
	// - FIELD_CONTENT: field content.
	//
	// - FIELD_NAME: field name.
	//
	// - FIELD_DESC: field description.
	//
	// - FIELD_TYPE: field type.
	//
	// example:
	//
	// TABLE_NAME
	Property *string `json:"Property,omitempty" xml:"Property,omitempty"`
	// The condition relation. Valid values:
	//
	// - EXPRESSION: expression.
	//
	// - RELATION: relation.
	//
	// This parameter is required.
	//
	// example:
	//
	// EXPRESSION
	Relation *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
	// The condition values.
	ValueList []*string `json:"ValueList,omitempty" xml:"ValueList,omitempty" type:"Repeated"`
}

func (s UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionList) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionList) GoString() string {
	return s.String()
}

func (s *UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionList) GetId() *string {
	return s.Id
}

func (s *UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionList) GetOperate() *string {
	return s.Operate
}

func (s *UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionList) GetOptionList() []*UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionListOptionList {
	return s.OptionList
}

func (s *UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionList) GetParentId() *string {
	return s.ParentId
}

func (s *UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionList) GetProperty() *string {
	return s.Property
}

func (s *UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionList) GetRelation() *string {
	return s.Relation
}

func (s *UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionList) GetValueList() []*string {
	return s.ValueList
}

func (s *UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionList) SetId(v string) *UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionList {
	s.Id = &v
	return s
}

func (s *UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionList) SetOperate(v string) *UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionList {
	s.Operate = &v
	return s
}

func (s *UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionList) SetOptionList(v []*UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionListOptionList) *UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionList {
	s.OptionList = v
	return s
}

func (s *UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionList) SetParentId(v string) *UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionList {
	s.ParentId = &v
	return s
}

func (s *UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionList) SetProperty(v string) *UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionList {
	s.Property = &v
	return s
}

func (s *UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionList) SetRelation(v string) *UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionList {
	s.Relation = &v
	return s
}

func (s *UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionList) SetValueList(v []*string) *UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionList {
	s.ValueList = v
	return s
}

func (s *UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionList) Validate() error {
	if s.OptionList != nil {
		for _, item := range s.OptionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionListOptionList struct {
	// The configuration item.
	//
	// This parameter is required.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The configuration item value.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionListOptionList) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionListOptionList) GoString() string {
	return s.String()
}

func (s *UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionListOptionList) GetKey() *string {
	return s.Key
}

func (s *UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionListOptionList) GetValue() *string {
	return s.Value
}

func (s *UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionListOptionList) SetKey(v string) *UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionListOptionList {
	s.Key = &v
	return s
}

func (s *UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionListOptionList) SetValue(v string) *UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionListOptionList {
	s.Value = &v
	return s
}

func (s *UpdateSecurityClassifyRequestUpdateCommandAdvancedConditionListOptionList) Validate() error {
	return dara.Validate(s)
}
