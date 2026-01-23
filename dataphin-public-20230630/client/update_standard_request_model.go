// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStandardRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateStandardRequest
	GetOpTenantId() *int64
	SetUpdateCommand(v *UpdateStandardRequestUpdateCommand) *UpdateStandardRequest
	GetUpdateCommand() *UpdateStandardRequestUpdateCommand
}

type UpdateStandardRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommand *UpdateStandardRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateStandardRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardRequest) GoString() string {
	return s.String()
}

func (s *UpdateStandardRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateStandardRequest) GetUpdateCommand() *UpdateStandardRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateStandardRequest) SetOpTenantId(v int64) *UpdateStandardRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateStandardRequest) SetUpdateCommand(v *UpdateStandardRequestUpdateCommand) *UpdateStandardRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdateStandardRequest) Validate() error {
	if s.UpdateCommand != nil {
		if err := s.UpdateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateStandardRequestUpdateCommand struct {
	// example:
	//
	// test
	Description              *string                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	EffectiveTimeConfig      *UpdateStandardRequestUpdateCommandEffectiveTimeConfig `json:"EffectiveTimeConfig,omitempty" xml:"EffectiveTimeConfig,omitempty" type:"Struct"`
	NeedGenerateStandardCode *bool                                                  `json:"NeedGenerateStandardCode,omitempty" xml:"NeedGenerateStandardCode,omitempty"`
	// example:
	//
	// 300000913
	Owner                        *string                                                         `json:"Owner,omitempty" xml:"Owner,omitempty"`
	StandardGeneralMonitorConfig *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfig `json:"StandardGeneralMonitorConfig,omitempty" xml:"StandardGeneralMonitorConfig,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 1234
	StandardId *int64 `json:"StandardId,omitempty" xml:"StandardId,omitempty"`
	// This parameter is required.
	StandardSetReference *UpdateStandardRequestUpdateCommandStandardSetReference `json:"StandardSetReference,omitempty" xml:"StandardSetReference,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// draft
	StandardStatus *string `json:"StandardStatus,omitempty" xml:"StandardStatus,omitempty"`
	// This parameter is required.
	StandardTemplateReference *UpdateStandardRequestUpdateCommandStandardTemplateReference `json:"StandardTemplateReference,omitempty" xml:"StandardTemplateReference,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UpdateStandardRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateStandardRequestUpdateCommand) GetDescription() *string {
	return s.Description
}

func (s *UpdateStandardRequestUpdateCommand) GetEffectiveTimeConfig() *UpdateStandardRequestUpdateCommandEffectiveTimeConfig {
	return s.EffectiveTimeConfig
}

func (s *UpdateStandardRequestUpdateCommand) GetNeedGenerateStandardCode() *bool {
	return s.NeedGenerateStandardCode
}

func (s *UpdateStandardRequestUpdateCommand) GetOwner() *string {
	return s.Owner
}

func (s *UpdateStandardRequestUpdateCommand) GetStandardGeneralMonitorConfig() *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfig {
	return s.StandardGeneralMonitorConfig
}

func (s *UpdateStandardRequestUpdateCommand) GetStandardId() *int64 {
	return s.StandardId
}

func (s *UpdateStandardRequestUpdateCommand) GetStandardSetReference() *UpdateStandardRequestUpdateCommandStandardSetReference {
	return s.StandardSetReference
}

func (s *UpdateStandardRequestUpdateCommand) GetStandardStatus() *string {
	return s.StandardStatus
}

func (s *UpdateStandardRequestUpdateCommand) GetStandardTemplateReference() *UpdateStandardRequestUpdateCommandStandardTemplateReference {
	return s.StandardTemplateReference
}

func (s *UpdateStandardRequestUpdateCommand) GetVersion() *int32 {
	return s.Version
}

func (s *UpdateStandardRequestUpdateCommand) SetDescription(v string) *UpdateStandardRequestUpdateCommand {
	s.Description = &v
	return s
}

func (s *UpdateStandardRequestUpdateCommand) SetEffectiveTimeConfig(v *UpdateStandardRequestUpdateCommandEffectiveTimeConfig) *UpdateStandardRequestUpdateCommand {
	s.EffectiveTimeConfig = v
	return s
}

func (s *UpdateStandardRequestUpdateCommand) SetNeedGenerateStandardCode(v bool) *UpdateStandardRequestUpdateCommand {
	s.NeedGenerateStandardCode = &v
	return s
}

func (s *UpdateStandardRequestUpdateCommand) SetOwner(v string) *UpdateStandardRequestUpdateCommand {
	s.Owner = &v
	return s
}

func (s *UpdateStandardRequestUpdateCommand) SetStandardGeneralMonitorConfig(v *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfig) *UpdateStandardRequestUpdateCommand {
	s.StandardGeneralMonitorConfig = v
	return s
}

func (s *UpdateStandardRequestUpdateCommand) SetStandardId(v int64) *UpdateStandardRequestUpdateCommand {
	s.StandardId = &v
	return s
}

func (s *UpdateStandardRequestUpdateCommand) SetStandardSetReference(v *UpdateStandardRequestUpdateCommandStandardSetReference) *UpdateStandardRequestUpdateCommand {
	s.StandardSetReference = v
	return s
}

func (s *UpdateStandardRequestUpdateCommand) SetStandardStatus(v string) *UpdateStandardRequestUpdateCommand {
	s.StandardStatus = &v
	return s
}

func (s *UpdateStandardRequestUpdateCommand) SetStandardTemplateReference(v *UpdateStandardRequestUpdateCommandStandardTemplateReference) *UpdateStandardRequestUpdateCommand {
	s.StandardTemplateReference = v
	return s
}

func (s *UpdateStandardRequestUpdateCommand) SetVersion(v int32) *UpdateStandardRequestUpdateCommand {
	s.Version = &v
	return s
}

func (s *UpdateStandardRequestUpdateCommand) Validate() error {
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

type UpdateStandardRequestUpdateCommandEffectiveTimeConfig struct {
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

func (s UpdateStandardRequestUpdateCommandEffectiveTimeConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardRequestUpdateCommandEffectiveTimeConfig) GoString() string {
	return s.String()
}

func (s *UpdateStandardRequestUpdateCommandEffectiveTimeConfig) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateStandardRequestUpdateCommandEffectiveTimeConfig) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateStandardRequestUpdateCommandEffectiveTimeConfig) GetType() *string {
	return s.Type
}

func (s *UpdateStandardRequestUpdateCommandEffectiveTimeConfig) SetEndTime(v string) *UpdateStandardRequestUpdateCommandEffectiveTimeConfig {
	s.EndTime = &v
	return s
}

func (s *UpdateStandardRequestUpdateCommandEffectiveTimeConfig) SetStartTime(v string) *UpdateStandardRequestUpdateCommandEffectiveTimeConfig {
	s.StartTime = &v
	return s
}

func (s *UpdateStandardRequestUpdateCommandEffectiveTimeConfig) SetType(v string) *UpdateStandardRequestUpdateCommandEffectiveTimeConfig {
	s.Type = &v
	return s
}

func (s *UpdateStandardRequestUpdateCommandEffectiveTimeConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfig struct {
	// This parameter is required.
	StandardMonitorConfigList []*UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList `json:"StandardMonitorConfigList,omitempty" xml:"StandardMonitorConfigList,omitempty" type:"Repeated"`
}

func (s UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfig) GoString() string {
	return s.String()
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfig) GetStandardMonitorConfigList() []*UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList {
	return s.StandardMonitorConfigList
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfig) SetStandardMonitorConfigList(v []*UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfig {
	s.StandardMonitorConfigList = v
	return s
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfig) Validate() error {
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

type UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList struct {
	// example:
	//
	// 112
	AttributeId            *int64                                                                                                         `json:"AttributeId,omitempty" xml:"AttributeId,omitempty"`
	AttributeMonitorConfig *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig `json:"AttributeMonitorConfig,omitempty" xml:"AttributeMonitorConfig,omitempty" type:"Struct"`
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
	QualityRuleTemplate *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate `json:"QualityRuleTemplate,omitempty" xml:"QualityRuleTemplate,omitempty" type:"Struct"`
	RuleConfigList      []*UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleConfigList    `json:"RuleConfigList,omitempty" xml:"RuleConfigList,omitempty" type:"Repeated"`
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
	RuleValidateConfigList []*UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList `json:"RuleValidateConfigList,omitempty" xml:"RuleValidateConfigList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// METADATA
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) GoString() string {
	return s.String()
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) GetAttributeId() *int64 {
	return s.AttributeId
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) GetAttributeMonitorConfig() *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig {
	return s.AttributeMonitorConfig
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) GetAttributeName() *string {
	return s.AttributeName
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) GetDescription() *string {
	return s.Description
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) GetId() *int64 {
	return s.Id
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) GetMonitorFrom() *string {
	return s.MonitorFrom
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) GetQualityRuleTemplate() *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate {
	return s.QualityRuleTemplate
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) GetRuleConfigList() []*UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleConfigList {
	return s.RuleConfigList
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) GetRuleSubType() *string {
	return s.RuleSubType
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) GetRuleValidateConfigList() []*UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList {
	return s.RuleValidateConfigList
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) GetType() *string {
	return s.Type
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) SetAttributeId(v int64) *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList {
	s.AttributeId = &v
	return s
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) SetAttributeMonitorConfig(v *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig) *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList {
	s.AttributeMonitorConfig = v
	return s
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) SetAttributeName(v string) *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList {
	s.AttributeName = &v
	return s
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) SetDescription(v string) *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList {
	s.Description = &v
	return s
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) SetId(v int64) *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList {
	s.Id = &v
	return s
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) SetMonitorFrom(v string) *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList {
	s.MonitorFrom = &v
	return s
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) SetQualityRuleTemplate(v *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate) *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList {
	s.QualityRuleTemplate = v
	return s
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) SetRuleConfigList(v []*UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleConfigList) *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList {
	s.RuleConfigList = v
	return s
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) SetRuleName(v string) *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList {
	s.RuleName = &v
	return s
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) SetRuleSubType(v string) *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList {
	s.RuleSubType = &v
	return s
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) SetRuleValidateConfigList(v []*UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList {
	s.RuleValidateConfigList = v
	return s
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) SetType(v string) *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList {
	s.Type = &v
	return s
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigList) Validate() error {
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

type UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig struct {
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

func (s UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig) GoString() string {
	return s.String()
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig) GetColumnName() *string {
	return s.ColumnName
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig) GetIsCaseSensitive() *bool {
	return s.IsCaseSensitive
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig) GetType() *string {
	return s.Type
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig) SetColumnName(v string) *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig {
	s.ColumnName = &v
	return s
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig) SetIsCaseSensitive(v bool) *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig {
	s.IsCaseSensitive = &v
	return s
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig) SetType(v string) *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig {
	s.Type = &v
	return s
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate struct {
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

func (s UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate) GoString() string {
	return s.String()
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate) GetId() *int64 {
	return s.Id
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate) GetName() *string {
	return s.Name
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate) GetType() *string {
	return s.Type
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate) SetId(v int64) *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate {
	s.Id = &v
	return s
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate) SetName(v string) *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate {
	s.Name = &v
	return s
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate) SetType(v string) *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate {
	s.Type = &v
	return s
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate) Validate() error {
	return dara.Validate(s)
}

type UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleConfigList struct {
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

func (s UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleConfigList) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleConfigList) GoString() string {
	return s.String()
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleConfigList) GetKey() *string {
	return s.Key
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleConfigList) GetValue() *string {
	return s.Value
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleConfigList) SetKey(v string) *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleConfigList {
	s.Key = &v
	return s
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleConfigList) SetValue(v string) *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleConfigList {
	s.Value = &v
	return s
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleConfigList) Validate() error {
	return dara.Validate(s)
}

type UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList struct {
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

func (s UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) GoString() string {
	return s.String()
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) GetId() *string {
	return s.Id
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) GetMetric() *string {
	return s.Metric
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) GetMetricName() *string {
	return s.MetricName
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) GetOperator() *string {
	return s.Operator
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) GetParentId() *string {
	return s.ParentId
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) GetType() *string {
	return s.Type
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) GetValue() *string {
	return s.Value
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) SetId(v string) *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList {
	s.Id = &v
	return s
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) SetMetric(v string) *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList {
	s.Metric = &v
	return s
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) SetMetricName(v string) *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList {
	s.MetricName = &v
	return s
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) SetOperator(v string) *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList {
	s.Operator = &v
	return s
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) SetParentId(v string) *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList {
	s.ParentId = &v
	return s
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) SetType(v string) *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList {
	s.Type = &v
	return s
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) SetValue(v string) *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList {
	s.Value = &v
	return s
}

func (s *UpdateStandardRequestUpdateCommandStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) Validate() error {
	return dara.Validate(s)
}

type UpdateStandardRequestUpdateCommandStandardSetReference struct {
	// This parameter is required.
	//
	// example:
	//
	// 22
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateStandardRequestUpdateCommandStandardSetReference) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardRequestUpdateCommandStandardSetReference) GoString() string {
	return s.String()
}

func (s *UpdateStandardRequestUpdateCommandStandardSetReference) GetId() *int64 {
	return s.Id
}

func (s *UpdateStandardRequestUpdateCommandStandardSetReference) SetId(v int64) *UpdateStandardRequestUpdateCommandStandardSetReference {
	s.Id = &v
	return s
}

func (s *UpdateStandardRequestUpdateCommandStandardSetReference) Validate() error {
	return dara.Validate(s)
}

type UpdateStandardRequestUpdateCommandStandardTemplateReference struct {
	AttributeValueList []*UpdateStandardRequestUpdateCommandStandardTemplateReferenceAttributeValueList `json:"AttributeValueList,omitempty" xml:"AttributeValueList,omitempty" type:"Repeated"`
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

func (s UpdateStandardRequestUpdateCommandStandardTemplateReference) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardRequestUpdateCommandStandardTemplateReference) GoString() string {
	return s.String()
}

func (s *UpdateStandardRequestUpdateCommandStandardTemplateReference) GetAttributeValueList() []*UpdateStandardRequestUpdateCommandStandardTemplateReferenceAttributeValueList {
	return s.AttributeValueList
}

func (s *UpdateStandardRequestUpdateCommandStandardTemplateReference) GetId() *int64 {
	return s.Id
}

func (s *UpdateStandardRequestUpdateCommandStandardTemplateReference) GetVersion() *int32 {
	return s.Version
}

func (s *UpdateStandardRequestUpdateCommandStandardTemplateReference) SetAttributeValueList(v []*UpdateStandardRequestUpdateCommandStandardTemplateReferenceAttributeValueList) *UpdateStandardRequestUpdateCommandStandardTemplateReference {
	s.AttributeValueList = v
	return s
}

func (s *UpdateStandardRequestUpdateCommandStandardTemplateReference) SetId(v int64) *UpdateStandardRequestUpdateCommandStandardTemplateReference {
	s.Id = &v
	return s
}

func (s *UpdateStandardRequestUpdateCommandStandardTemplateReference) SetVersion(v int32) *UpdateStandardRequestUpdateCommandStandardTemplateReference {
	s.Version = &v
	return s
}

func (s *UpdateStandardRequestUpdateCommandStandardTemplateReference) Validate() error {
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

type UpdateStandardRequestUpdateCommandStandardTemplateReferenceAttributeValueList struct {
	// example:
	//
	// 1011
	AttributeId *int64 `json:"AttributeId,omitempty" xml:"AttributeId,omitempty"`
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateStandardRequestUpdateCommandStandardTemplateReferenceAttributeValueList) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardRequestUpdateCommandStandardTemplateReferenceAttributeValueList) GoString() string {
	return s.String()
}

func (s *UpdateStandardRequestUpdateCommandStandardTemplateReferenceAttributeValueList) GetAttributeId() *int64 {
	return s.AttributeId
}

func (s *UpdateStandardRequestUpdateCommandStandardTemplateReferenceAttributeValueList) GetValue() *string {
	return s.Value
}

func (s *UpdateStandardRequestUpdateCommandStandardTemplateReferenceAttributeValueList) SetAttributeId(v int64) *UpdateStandardRequestUpdateCommandStandardTemplateReferenceAttributeValueList {
	s.AttributeId = &v
	return s
}

func (s *UpdateStandardRequestUpdateCommandStandardTemplateReferenceAttributeValueList) SetValue(v string) *UpdateStandardRequestUpdateCommandStandardTemplateReferenceAttributeValueList {
	s.Value = &v
	return s
}

func (s *UpdateStandardRequestUpdateCommandStandardTemplateReferenceAttributeValueList) Validate() error {
	return dara.Validate(s)
}
