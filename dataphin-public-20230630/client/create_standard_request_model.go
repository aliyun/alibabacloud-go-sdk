// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStandardRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommand(v *CreateStandardRequestCreateCommand) *CreateStandardRequest
	GetCreateCommand() *CreateStandardRequestCreateCommand
	SetOpTenantId(v int64) *CreateStandardRequest
	GetOpTenantId() *int64
}

type CreateStandardRequest struct {
	// This parameter is required.
	CreateCommand *CreateStandardRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateStandardRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardRequest) GoString() string {
	return s.String()
}

func (s *CreateStandardRequest) GetCreateCommand() *CreateStandardRequestCreateCommand {
	return s.CreateCommand
}

func (s *CreateStandardRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateStandardRequest) SetCreateCommand(v *CreateStandardRequestCreateCommand) *CreateStandardRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateStandardRequest) SetOpTenantId(v int64) *CreateStandardRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateStandardRequest) Validate() error {
	if s.CreateCommand != nil {
		if err := s.CreateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateStandardRequestCreateCommand struct {
	// example:
	//
	// test
	Description              *string                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	EffectiveTimeConfig      *CreateStandardRequestCreateCommandEffectiveTimeConfig `json:"EffectiveTimeConfig,omitempty" xml:"EffectiveTimeConfig,omitempty" type:"Struct"`
	NeedGenerateStandardCode *bool                                                  `json:"NeedGenerateStandardCode,omitempty" xml:"NeedGenerateStandardCode,omitempty"`
	// example:
	//
	// 300000913
	Owner                        *string                                                         `json:"Owner,omitempty" xml:"Owner,omitempty"`
	StandardGeneralMonitorConfig *CreateStandardRequestCreateCommandStandardGeneralMonitorConfig `json:"StandardGeneralMonitorConfig,omitempty" xml:"StandardGeneralMonitorConfig,omitempty" type:"Struct"`
	// This parameter is required.
	StandardSetReference *CreateStandardRequestCreateCommandStandardSetReference `json:"StandardSetReference,omitempty" xml:"StandardSetReference,omitempty" type:"Struct"`
	// This parameter is required.
	StandardTemplateReference *CreateStandardRequestCreateCommandStandardTemplateReference `json:"StandardTemplateReference,omitempty" xml:"StandardTemplateReference,omitempty" type:"Struct"`
}

func (s CreateStandardRequestCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateStandardRequestCreateCommand) GetDescription() *string {
	return s.Description
}

func (s *CreateStandardRequestCreateCommand) GetEffectiveTimeConfig() *CreateStandardRequestCreateCommandEffectiveTimeConfig {
	return s.EffectiveTimeConfig
}

func (s *CreateStandardRequestCreateCommand) GetNeedGenerateStandardCode() *bool {
	return s.NeedGenerateStandardCode
}

func (s *CreateStandardRequestCreateCommand) GetOwner() *string {
	return s.Owner
}

func (s *CreateStandardRequestCreateCommand) GetStandardGeneralMonitorConfig() *CreateStandardRequestCreateCommandStandardGeneralMonitorConfig {
	return s.StandardGeneralMonitorConfig
}

func (s *CreateStandardRequestCreateCommand) GetStandardSetReference() *CreateStandardRequestCreateCommandStandardSetReference {
	return s.StandardSetReference
}

func (s *CreateStandardRequestCreateCommand) GetStandardTemplateReference() *CreateStandardRequestCreateCommandStandardTemplateReference {
	return s.StandardTemplateReference
}

func (s *CreateStandardRequestCreateCommand) SetDescription(v string) *CreateStandardRequestCreateCommand {
	s.Description = &v
	return s
}

func (s *CreateStandardRequestCreateCommand) SetEffectiveTimeConfig(v *CreateStandardRequestCreateCommandEffectiveTimeConfig) *CreateStandardRequestCreateCommand {
	s.EffectiveTimeConfig = v
	return s
}

func (s *CreateStandardRequestCreateCommand) SetNeedGenerateStandardCode(v bool) *CreateStandardRequestCreateCommand {
	s.NeedGenerateStandardCode = &v
	return s
}

func (s *CreateStandardRequestCreateCommand) SetOwner(v string) *CreateStandardRequestCreateCommand {
	s.Owner = &v
	return s
}

func (s *CreateStandardRequestCreateCommand) SetStandardGeneralMonitorConfig(v *CreateStandardRequestCreateCommandStandardGeneralMonitorConfig) *CreateStandardRequestCreateCommand {
	s.StandardGeneralMonitorConfig = v
	return s
}

func (s *CreateStandardRequestCreateCommand) SetStandardSetReference(v *CreateStandardRequestCreateCommandStandardSetReference) *CreateStandardRequestCreateCommand {
	s.StandardSetReference = v
	return s
}

func (s *CreateStandardRequestCreateCommand) SetStandardTemplateReference(v *CreateStandardRequestCreateCommandStandardTemplateReference) *CreateStandardRequestCreateCommand {
	s.StandardTemplateReference = v
	return s
}

func (s *CreateStandardRequestCreateCommand) Validate() error {
	if s.EffectiveTimeConfig != nil {
		if err := s.EffectiveTimeConfig.Validate(); err != nil {
			return err
		}
	}
	if s.StandardGeneralMonitorConfig != nil {
		if err := s.StandardGeneralMonitorConfig.Validate(); err != nil {
			return err
		}
	}
	if s.StandardSetReference != nil {
		if err := s.StandardSetReference.Validate(); err != nil {
			return err
		}
	}
	if s.StandardTemplateReference != nil {
		if err := s.StandardTemplateReference.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateStandardRequestCreateCommandEffectiveTimeConfig struct {
	// example:
	//
	// 2025-12-30 00:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 2025-06-30 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TIME_PERIOD
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateStandardRequestCreateCommandEffectiveTimeConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardRequestCreateCommandEffectiveTimeConfig) GoString() string {
	return s.String()
}

func (s *CreateStandardRequestCreateCommandEffectiveTimeConfig) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateStandardRequestCreateCommandEffectiveTimeConfig) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateStandardRequestCreateCommandEffectiveTimeConfig) GetType() *string {
	return s.Type
}

func (s *CreateStandardRequestCreateCommandEffectiveTimeConfig) SetEndTime(v string) *CreateStandardRequestCreateCommandEffectiveTimeConfig {
	s.EndTime = &v
	return s
}

func (s *CreateStandardRequestCreateCommandEffectiveTimeConfig) SetStartTime(v string) *CreateStandardRequestCreateCommandEffectiveTimeConfig {
	s.StartTime = &v
	return s
}

func (s *CreateStandardRequestCreateCommandEffectiveTimeConfig) SetType(v string) *CreateStandardRequestCreateCommandEffectiveTimeConfig {
	s.Type = &v
	return s
}

func (s *CreateStandardRequestCreateCommandEffectiveTimeConfig) Validate() error {
	return dara.Validate(s)
}

type CreateStandardRequestCreateCommandStandardGeneralMonitorConfig struct {
	// This parameter is required.
	StandardMonitorConfigList []*CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList `json:"StandardMonitorConfigList,omitempty" xml:"StandardMonitorConfigList,omitempty" type:"Repeated"`
}

func (s CreateStandardRequestCreateCommandStandardGeneralMonitorConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardRequestCreateCommandStandardGeneralMonitorConfig) GoString() string {
	return s.String()
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfig) GetStandardMonitorConfigList() []*CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList {
	return s.StandardMonitorConfigList
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfig) SetStandardMonitorConfigList(v []*CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) *CreateStandardRequestCreateCommandStandardGeneralMonitorConfig {
	s.StandardMonitorConfigList = v
	return s
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfig) Validate() error {
	if s.StandardMonitorConfigList != nil {
		for _, item := range s.StandardMonitorConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList struct {
	// example:
	//
	// 112
	AttributeId            *int64                                                                                                         `json:"AttributeId,omitempty" xml:"AttributeId,omitempty"`
	AttributeMonitorConfig *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig `json:"AttributeMonitorConfig,omitempty" xml:"AttributeMonitorConfig,omitempty" type:"Struct"`
	// example:
	//
	// teset
	AttributeName *string `json:"AttributeName,omitempty" xml:"AttributeName,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// BY_SYSTEM_ATTRIBUTE
	MonitorFrom         *string                                                                                                     `json:"MonitorFrom,omitempty" xml:"MonitorFrom,omitempty"`
	QualityRuleTemplate *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate `json:"QualityRuleTemplate,omitempty" xml:"QualityRuleTemplate,omitempty" type:"Struct"`
	RuleConfigList      []*CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleConfigList    `json:"RuleConfigList,omitempty" xml:"RuleConfigList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// CUSTOMIZED
	RuleSubType            *string                                                                                                          `json:"RuleSubType,omitempty" xml:"RuleSubType,omitempty"`
	RuleValidateConfigList []*CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList `json:"RuleValidateConfigList,omitempty" xml:"RuleValidateConfigList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// METADATA
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) GoString() string {
	return s.String()
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) GetAttributeId() *int64 {
	return s.AttributeId
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) GetAttributeMonitorConfig() *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig {
	return s.AttributeMonitorConfig
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) GetAttributeName() *string {
	return s.AttributeName
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) GetDescription() *string {
	return s.Description
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) GetId() *int64 {
	return s.Id
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) GetMonitorFrom() *string {
	return s.MonitorFrom
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) GetQualityRuleTemplate() *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate {
	return s.QualityRuleTemplate
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) GetRuleConfigList() []*CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleConfigList {
	return s.RuleConfigList
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) GetRuleSubType() *string {
	return s.RuleSubType
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) GetRuleValidateConfigList() []*CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList {
	return s.RuleValidateConfigList
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) GetType() *string {
	return s.Type
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) SetAttributeId(v int64) *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList {
	s.AttributeId = &v
	return s
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) SetAttributeMonitorConfig(v *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig) *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList {
	s.AttributeMonitorConfig = v
	return s
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) SetAttributeName(v string) *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList {
	s.AttributeName = &v
	return s
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) SetDescription(v string) *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList {
	s.Description = &v
	return s
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) SetId(v int64) *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList {
	s.Id = &v
	return s
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) SetMonitorFrom(v string) *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList {
	s.MonitorFrom = &v
	return s
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) SetQualityRuleTemplate(v *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate) *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList {
	s.QualityRuleTemplate = v
	return s
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) SetRuleConfigList(v []*CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleConfigList) *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList {
	s.RuleConfigList = v
	return s
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) SetRuleName(v string) *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList {
	s.RuleName = &v
	return s
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) SetRuleSubType(v string) *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList {
	s.RuleSubType = &v
	return s
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) SetRuleValidateConfigList(v []*CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList {
	s.RuleValidateConfigList = v
	return s
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) SetType(v string) *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList {
	s.Type = &v
	return s
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) Validate() error {
	if s.AttributeMonitorConfig != nil {
		if err := s.AttributeMonitorConfig.Validate(); err != nil {
			return err
		}
	}
	if s.QualityRuleTemplate != nil {
		if err := s.QualityRuleTemplate.Validate(); err != nil {
			return err
		}
	}
	if s.RuleConfigList != nil {
		for _, item := range s.RuleConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RuleValidateConfigList != nil {
		for _, item := range s.RuleValidateConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig struct {
	// example:
	//
	// column1
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// This parameter is required.
	IsCaseSensitive *bool `json:"IsCaseSensitive,omitempty" xml:"IsCaseSensitive,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// METADATA
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig) GoString() string {
	return s.String()
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig) GetColumnName() *string {
	return s.ColumnName
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig) GetIsCaseSensitive() *bool {
	return s.IsCaseSensitive
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig) GetType() *string {
	return s.Type
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig) SetColumnName(v string) *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig {
	s.ColumnName = &v
	return s
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig) SetIsCaseSensitive(v bool) *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig {
	s.IsCaseSensitive = &v
	return s
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig) SetType(v string) *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig {
	s.Type = &v
	return s
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig) Validate() error {
	return dara.Validate(s)
}

type CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate struct {
	// This parameter is required.
	//
	// example:
	//
	// 22
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CUSTOMIZED
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CUSTOMIZED
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate) GoString() string {
	return s.String()
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate) GetId() *int64 {
	return s.Id
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate) GetName() *string {
	return s.Name
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate) GetType() *string {
	return s.Type
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate) SetId(v int64) *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate {
	s.Id = &v
	return s
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate) SetName(v string) *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate {
	s.Name = &v
	return s
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate) SetType(v string) *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate {
	s.Type = &v
	return s
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate) Validate() error {
	return dara.Validate(s)
}

type CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleConfigList struct {
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

func (s CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleConfigList) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleConfigList) GoString() string {
	return s.String()
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleConfigList) GetKey() *string {
	return s.Key
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleConfigList) GetValue() *string {
	return s.Value
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleConfigList) SetKey(v string) *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleConfigList {
	s.Key = &v
	return s
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleConfigList) SetValue(v string) *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleConfigList {
	s.Value = &v
	return s
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleConfigList) Validate() error {
	return dara.Validate(s)
}

type CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList struct {
	// This parameter is required.
	//
	// example:
	//
	// abc
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// a
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// example:
	//
	// test
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AND
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// example:
	//
	// a
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// RELATION
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) GoString() string {
	return s.String()
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) GetId() *string {
	return s.Id
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) GetMetric() *string {
	return s.Metric
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) GetMetricName() *string {
	return s.MetricName
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) GetOperator() *string {
	return s.Operator
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) GetParentId() *string {
	return s.ParentId
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) GetType() *string {
	return s.Type
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) GetValue() *string {
	return s.Value
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) SetId(v string) *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList {
	s.Id = &v
	return s
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) SetMetric(v string) *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList {
	s.Metric = &v
	return s
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) SetMetricName(v string) *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList {
	s.MetricName = &v
	return s
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) SetOperator(v string) *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList {
	s.Operator = &v
	return s
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) SetParentId(v string) *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList {
	s.ParentId = &v
	return s
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) SetType(v string) *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList {
	s.Type = &v
	return s
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) SetValue(v string) *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList {
	s.Value = &v
	return s
}

func (s *CreateStandardRequestCreateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) Validate() error {
	return dara.Validate(s)
}

type CreateStandardRequestCreateCommandStandardSetReference struct {
	// This parameter is required.
	//
	// example:
	//
	// 22
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateStandardRequestCreateCommandStandardSetReference) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardRequestCreateCommandStandardSetReference) GoString() string {
	return s.String()
}

func (s *CreateStandardRequestCreateCommandStandardSetReference) GetId() *int64 {
	return s.Id
}

func (s *CreateStandardRequestCreateCommandStandardSetReference) SetId(v int64) *CreateStandardRequestCreateCommandStandardSetReference {
	s.Id = &v
	return s
}

func (s *CreateStandardRequestCreateCommandStandardSetReference) Validate() error {
	return dara.Validate(s)
}

type CreateStandardRequestCreateCommandStandardTemplateReference struct {
	AttributeValueList []*CreateStandardRequestCreateCommandStandardTemplateReferenceAttributeValueList `json:"AttributeValueList,omitempty" xml:"AttributeValueList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 11
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateStandardRequestCreateCommandStandardTemplateReference) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardRequestCreateCommandStandardTemplateReference) GoString() string {
	return s.String()
}

func (s *CreateStandardRequestCreateCommandStandardTemplateReference) GetAttributeValueList() []*CreateStandardRequestCreateCommandStandardTemplateReferenceAttributeValueList {
	return s.AttributeValueList
}

func (s *CreateStandardRequestCreateCommandStandardTemplateReference) GetId() *int64 {
	return s.Id
}

func (s *CreateStandardRequestCreateCommandStandardTemplateReference) GetVersion() *int32 {
	return s.Version
}

func (s *CreateStandardRequestCreateCommandStandardTemplateReference) SetAttributeValueList(v []*CreateStandardRequestCreateCommandStandardTemplateReferenceAttributeValueList) *CreateStandardRequestCreateCommandStandardTemplateReference {
	s.AttributeValueList = v
	return s
}

func (s *CreateStandardRequestCreateCommandStandardTemplateReference) SetId(v int64) *CreateStandardRequestCreateCommandStandardTemplateReference {
	s.Id = &v
	return s
}

func (s *CreateStandardRequestCreateCommandStandardTemplateReference) SetVersion(v int32) *CreateStandardRequestCreateCommandStandardTemplateReference {
	s.Version = &v
	return s
}

func (s *CreateStandardRequestCreateCommandStandardTemplateReference) Validate() error {
	if s.AttributeValueList != nil {
		for _, item := range s.AttributeValueList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateStandardRequestCreateCommandStandardTemplateReferenceAttributeValueList struct {
	// example:
	//
	// 1011
	AttributeId *int64 `json:"AttributeId,omitempty" xml:"AttributeId,omitempty"`
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateStandardRequestCreateCommandStandardTemplateReferenceAttributeValueList) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardRequestCreateCommandStandardTemplateReferenceAttributeValueList) GoString() string {
	return s.String()
}

func (s *CreateStandardRequestCreateCommandStandardTemplateReferenceAttributeValueList) GetAttributeId() *int64 {
	return s.AttributeId
}

func (s *CreateStandardRequestCreateCommandStandardTemplateReferenceAttributeValueList) GetValue() *string {
	return s.Value
}

func (s *CreateStandardRequestCreateCommandStandardTemplateReferenceAttributeValueList) SetAttributeId(v int64) *CreateStandardRequestCreateCommandStandardTemplateReferenceAttributeValueList {
	s.AttributeId = &v
	return s
}

func (s *CreateStandardRequestCreateCommandStandardTemplateReferenceAttributeValueList) SetValue(v string) *CreateStandardRequestCreateCommandStandardTemplateReferenceAttributeValueList {
	s.Value = &v
	return s
}

func (s *CreateStandardRequestCreateCommandStandardTemplateReferenceAttributeValueList) Validate() error {
	return dara.Validate(s)
}
