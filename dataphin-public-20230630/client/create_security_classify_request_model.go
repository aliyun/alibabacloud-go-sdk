// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecurityClassifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommand(v *CreateSecurityClassifyRequestCreateCommand) *CreateSecurityClassifyRequest
	GetCreateCommand() *CreateSecurityClassifyRequestCreateCommand
	SetOpTenantId(v int64) *CreateSecurityClassifyRequest
	GetOpTenantId() *int64
}

type CreateSecurityClassifyRequest struct {
	// This parameter is required.
	CreateCommand *CreateSecurityClassifyRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateSecurityClassifyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityClassifyRequest) GoString() string {
	return s.String()
}

func (s *CreateSecurityClassifyRequest) GetCreateCommand() *CreateSecurityClassifyRequestCreateCommand {
	return s.CreateCommand
}

func (s *CreateSecurityClassifyRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateSecurityClassifyRequest) SetCreateCommand(v *CreateSecurityClassifyRequestCreateCommand) *CreateSecurityClassifyRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateSecurityClassifyRequest) SetOpTenantId(v int64) *CreateSecurityClassifyRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateSecurityClassifyRequest) Validate() error {
	if s.CreateCommand != nil {
		if err := s.CreateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSecurityClassifyRequestCreateCommand struct {
	// example:
	//
	// test
	Abbreviation          *string                                                            `json:"Abbreviation,omitempty" xml:"Abbreviation,omitempty"`
	AdvancedConditionList []*CreateSecurityClassifyRequestCreateCommandAdvancedConditionList `json:"AdvancedConditionList,omitempty" xml:"AdvancedConditionList,omitempty" type:"Repeated"`
	// example:
	//
	// test
	Description     *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	FeatureNameList []*string `json:"FeatureNameList,omitempty" xml:"FeatureNameList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// L2
	LevelName *string `json:"LevelName,omitempty" xml:"LevelName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// /
	ParentPath *string `json:"ParentPath,omitempty" xml:"ParentPath,omitempty"`
	// example:
	//
	// 5
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// ENABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateSecurityClassifyRequestCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityClassifyRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateSecurityClassifyRequestCreateCommand) GetAbbreviation() *string {
	return s.Abbreviation
}

func (s *CreateSecurityClassifyRequestCreateCommand) GetAdvancedConditionList() []*CreateSecurityClassifyRequestCreateCommandAdvancedConditionList {
	return s.AdvancedConditionList
}

func (s *CreateSecurityClassifyRequestCreateCommand) GetDescription() *string {
	return s.Description
}

func (s *CreateSecurityClassifyRequestCreateCommand) GetFeatureNameList() []*string {
	return s.FeatureNameList
}

func (s *CreateSecurityClassifyRequestCreateCommand) GetLevelName() *string {
	return s.LevelName
}

func (s *CreateSecurityClassifyRequestCreateCommand) GetName() *string {
	return s.Name
}

func (s *CreateSecurityClassifyRequestCreateCommand) GetParentPath() *string {
	return s.ParentPath
}

func (s *CreateSecurityClassifyRequestCreateCommand) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateSecurityClassifyRequestCreateCommand) GetStatus() *string {
	return s.Status
}

func (s *CreateSecurityClassifyRequestCreateCommand) SetAbbreviation(v string) *CreateSecurityClassifyRequestCreateCommand {
	s.Abbreviation = &v
	return s
}

func (s *CreateSecurityClassifyRequestCreateCommand) SetAdvancedConditionList(v []*CreateSecurityClassifyRequestCreateCommandAdvancedConditionList) *CreateSecurityClassifyRequestCreateCommand {
	s.AdvancedConditionList = v
	return s
}

func (s *CreateSecurityClassifyRequestCreateCommand) SetDescription(v string) *CreateSecurityClassifyRequestCreateCommand {
	s.Description = &v
	return s
}

func (s *CreateSecurityClassifyRequestCreateCommand) SetFeatureNameList(v []*string) *CreateSecurityClassifyRequestCreateCommand {
	s.FeatureNameList = v
	return s
}

func (s *CreateSecurityClassifyRequestCreateCommand) SetLevelName(v string) *CreateSecurityClassifyRequestCreateCommand {
	s.LevelName = &v
	return s
}

func (s *CreateSecurityClassifyRequestCreateCommand) SetName(v string) *CreateSecurityClassifyRequestCreateCommand {
	s.Name = &v
	return s
}

func (s *CreateSecurityClassifyRequestCreateCommand) SetParentPath(v string) *CreateSecurityClassifyRequestCreateCommand {
	s.ParentPath = &v
	return s
}

func (s *CreateSecurityClassifyRequestCreateCommand) SetPriority(v int32) *CreateSecurityClassifyRequestCreateCommand {
	s.Priority = &v
	return s
}

func (s *CreateSecurityClassifyRequestCreateCommand) SetStatus(v string) *CreateSecurityClassifyRequestCreateCommand {
	s.Status = &v
	return s
}

func (s *CreateSecurityClassifyRequestCreateCommand) Validate() error {
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

type CreateSecurityClassifyRequestCreateCommandAdvancedConditionList struct {
	// This parameter is required.
	//
	// example:
	//
	// 234
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AND
	Operate    *string                                                                      `json:"Operate,omitempty" xml:"Operate,omitempty"`
	OptionList []*CreateSecurityClassifyRequestCreateCommandAdvancedConditionListOptionList `json:"OptionList,omitempty" xml:"OptionList,omitempty" type:"Repeated"`
	// example:
	//
	// 123
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// example:
	//
	// TABLE_NAME
	Property *string `json:"Property,omitempty" xml:"Property,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// EXPRESSION
	Relation  *string   `json:"Relation,omitempty" xml:"Relation,omitempty"`
	ValueList []*string `json:"ValueList,omitempty" xml:"ValueList,omitempty" type:"Repeated"`
}

func (s CreateSecurityClassifyRequestCreateCommandAdvancedConditionList) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityClassifyRequestCreateCommandAdvancedConditionList) GoString() string {
	return s.String()
}

func (s *CreateSecurityClassifyRequestCreateCommandAdvancedConditionList) GetId() *string {
	return s.Id
}

func (s *CreateSecurityClassifyRequestCreateCommandAdvancedConditionList) GetOperate() *string {
	return s.Operate
}

func (s *CreateSecurityClassifyRequestCreateCommandAdvancedConditionList) GetOptionList() []*CreateSecurityClassifyRequestCreateCommandAdvancedConditionListOptionList {
	return s.OptionList
}

func (s *CreateSecurityClassifyRequestCreateCommandAdvancedConditionList) GetParentId() *string {
	return s.ParentId
}

func (s *CreateSecurityClassifyRequestCreateCommandAdvancedConditionList) GetProperty() *string {
	return s.Property
}

func (s *CreateSecurityClassifyRequestCreateCommandAdvancedConditionList) GetRelation() *string {
	return s.Relation
}

func (s *CreateSecurityClassifyRequestCreateCommandAdvancedConditionList) GetValueList() []*string {
	return s.ValueList
}

func (s *CreateSecurityClassifyRequestCreateCommandAdvancedConditionList) SetId(v string) *CreateSecurityClassifyRequestCreateCommandAdvancedConditionList {
	s.Id = &v
	return s
}

func (s *CreateSecurityClassifyRequestCreateCommandAdvancedConditionList) SetOperate(v string) *CreateSecurityClassifyRequestCreateCommandAdvancedConditionList {
	s.Operate = &v
	return s
}

func (s *CreateSecurityClassifyRequestCreateCommandAdvancedConditionList) SetOptionList(v []*CreateSecurityClassifyRequestCreateCommandAdvancedConditionListOptionList) *CreateSecurityClassifyRequestCreateCommandAdvancedConditionList {
	s.OptionList = v
	return s
}

func (s *CreateSecurityClassifyRequestCreateCommandAdvancedConditionList) SetParentId(v string) *CreateSecurityClassifyRequestCreateCommandAdvancedConditionList {
	s.ParentId = &v
	return s
}

func (s *CreateSecurityClassifyRequestCreateCommandAdvancedConditionList) SetProperty(v string) *CreateSecurityClassifyRequestCreateCommandAdvancedConditionList {
	s.Property = &v
	return s
}

func (s *CreateSecurityClassifyRequestCreateCommandAdvancedConditionList) SetRelation(v string) *CreateSecurityClassifyRequestCreateCommandAdvancedConditionList {
	s.Relation = &v
	return s
}

func (s *CreateSecurityClassifyRequestCreateCommandAdvancedConditionList) SetValueList(v []*string) *CreateSecurityClassifyRequestCreateCommandAdvancedConditionList {
	s.ValueList = v
	return s
}

func (s *CreateSecurityClassifyRequestCreateCommandAdvancedConditionList) Validate() error {
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

type CreateSecurityClassifyRequestCreateCommandAdvancedConditionListOptionList struct {
	// This parameter is required.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateSecurityClassifyRequestCreateCommandAdvancedConditionListOptionList) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityClassifyRequestCreateCommandAdvancedConditionListOptionList) GoString() string {
	return s.String()
}

func (s *CreateSecurityClassifyRequestCreateCommandAdvancedConditionListOptionList) GetKey() *string {
	return s.Key
}

func (s *CreateSecurityClassifyRequestCreateCommandAdvancedConditionListOptionList) GetValue() *string {
	return s.Value
}

func (s *CreateSecurityClassifyRequestCreateCommandAdvancedConditionListOptionList) SetKey(v string) *CreateSecurityClassifyRequestCreateCommandAdvancedConditionListOptionList {
	s.Key = &v
	return s
}

func (s *CreateSecurityClassifyRequestCreateCommandAdvancedConditionListOptionList) SetValue(v string) *CreateSecurityClassifyRequestCreateCommandAdvancedConditionListOptionList {
	s.Value = &v
	return s
}

func (s *CreateSecurityClassifyRequestCreateCommandAdvancedConditionListOptionList) Validate() error {
	return dara.Validate(s)
}
