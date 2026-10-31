// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertQualityTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpsertQualityTemplateRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *UpsertQualityTemplateRequest
	GetOpUserId() *string
	SetUpsertCommand(v *UpsertQualityTemplateRequestUpsertCommand) *UpsertQualityTemplateRequest
	GetUpsertCommand() *UpsertQualityTemplateRequestUpsertCommand
}

type UpsertQualityTemplateRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The upsert command.
	//
	// This parameter is required.
	UpsertCommand *UpsertQualityTemplateRequestUpsertCommand `json:"UpsertCommand,omitempty" xml:"UpsertCommand,omitempty" type:"Struct"`
}

func (s UpsertQualityTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpsertQualityTemplateRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpsertQualityTemplateRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *UpsertQualityTemplateRequest) GetUpsertCommand() *UpsertQualityTemplateRequestUpsertCommand {
	return s.UpsertCommand
}

func (s *UpsertQualityTemplateRequest) SetOpTenantId(v int64) *UpsertQualityTemplateRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpsertQualityTemplateRequest) SetOpUserId(v string) *UpsertQualityTemplateRequest {
	s.OpUserId = &v
	return s
}

func (s *UpsertQualityTemplateRequest) SetUpsertCommand(v *UpsertQualityTemplateRequestUpsertCommand) *UpsertQualityTemplateRequest {
	s.UpsertCommand = v
	return s
}

func (s *UpsertQualityTemplateRequest) Validate() error {
	if s.UpsertCommand != nil {
		if err := s.UpsertCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpsertQualityTemplateRequestUpsertCommand struct {
	// The category of the template. Valid values: CONSISTENT (consistency), EFFECTIVE (effectiveness), TIMELINESE (timeliness), ACCURATE (accuracy), UNIQUENESS (uniqueness), COMPLETENESS (completeness), STABILITY (stability), and CUSTOM (custom).
	//
	// This parameter is required.
	//
	// example:
	//
	// CONSISTENT
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// The description of the template.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The key-value pairs for rule configuration.
	FormPropertyList []*UpsertQualityTemplateRequestUpsertCommandFormPropertyList `json:"FormPropertyList,omitempty" xml:"FormPropertyList,omitempty" type:"Repeated"`
	// The template ID. If this parameter is specified, the template is updated. If this parameter is not specified, a new template is created.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the template owner.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30012011
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The supported data source types, such as MySQL, Oracle, Microsoft SQL Server, MaxCompute, and Hive.
	SupportDataSourceTypeList []*string `json:"SupportDataSourceTypeList,omitempty" xml:"SupportDataSourceTypeList,omitempty" type:"Repeated"`
	// The templatetype. Valid values: FIELD_NULL_VALUE_VALIDATE (field null value check), FIELD_EMPTY_STRING_VALIDATE (field empty character check), FIELD_UNIQUE_VALIDATE (field uniqueness check), FIELD_GROUP_COUNT_VALIDATE (field unique value count check), FIELD_DUPLICATE_VALUE_COUNT_VALIDATE (field duplicate value count check), FUNCTION_TIME_COMPARE (time function comparison), SINGLE_TABLE_TIME_COMPARE (non-partitioned table time field comparison), DOUBLE_TABLE_TIME_COMPARE (two-table time field comparison), FIELD_FORMAT_VALIDATE (field format check), FIELD_LENGTH_VALIDATE (field length check), FIELD_VALUE_RANGE_VALIDATE (field value range check), CODE_TABLE_COMPARE (lookup table reference comparison), STANDARD_CODE_TABLE_COMPARE (data standard lookup table reference comparison), SINGLE_TABLE_FIELD_VALUE_COMPARE (non-partitioned table field value consistency comparison), SINGLE_TABLE_FIELD_STATISTICAL_COMPARE (non-partitioned table field statistical value consistency comparison), SINGLE_TABLE_FIELD_EXP_COMPARE (non-partitioned table field business logic consistency comparison), DOUBLE_TABLE_FIELD_VALUE_COMPARE (two-table field value consistency comparison), DOUBLE_TABLE_FIELD_STATISTICAL_COMPARE (two-table field statistical value consistency comparison), CROSS_DOUBLE_TABLE_FIELD_STATISTICAL_COMPARE (cross-source two-table field statistical value consistency comparison), DOUBLE_TABLE_FIELD_EXP_COMPARE (two-table field business logic consistency comparison), TABLE_STABILITY_VALIDATE (table stability check), TABLE_FLUCTUATION_VALIDATE (table fluctuation check), FIELD_STABILITY_VALIDATE (field stability check), FIELD_FLUCTUATION_VALIDATE (field fluctuation check), CUSTOM_STATISTICAL_VALIDATE (custom statistical metric check), CUSTOM_DATA_DETAILS_VALIDATE (custom data details check), DATASOURCE_AVAILABLE_CHECK (data source connectivity monitoring), TABLE_SCHEMA_CHECK (table schema change monitoring), REAL_TIME_OFFLINE_COMPARE (real-time and offline comparison), REAL_TIME_STATISTICAL_VALIDATE (real-time statistical value monitoring), and REAL_TIME_MULTI_CHAIN_COMPARE (real-time multi-link comparison).
	//
	// This parameter is required.
	//
	// example:
	//
	// FIELD_NULL_VALUE_VALIDATE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpsertQualityTemplateRequestUpsertCommand) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityTemplateRequestUpsertCommand) GoString() string {
	return s.String()
}

func (s *UpsertQualityTemplateRequestUpsertCommand) GetCatalog() *string {
	return s.Catalog
}

func (s *UpsertQualityTemplateRequestUpsertCommand) GetDescription() *string {
	return s.Description
}

func (s *UpsertQualityTemplateRequestUpsertCommand) GetFormPropertyList() []*UpsertQualityTemplateRequestUpsertCommandFormPropertyList {
	return s.FormPropertyList
}

func (s *UpsertQualityTemplateRequestUpsertCommand) GetId() *int64 {
	return s.Id
}

func (s *UpsertQualityTemplateRequestUpsertCommand) GetName() *string {
	return s.Name
}

func (s *UpsertQualityTemplateRequestUpsertCommand) GetOwner() *string {
	return s.Owner
}

func (s *UpsertQualityTemplateRequestUpsertCommand) GetSupportDataSourceTypeList() []*string {
	return s.SupportDataSourceTypeList
}

func (s *UpsertQualityTemplateRequestUpsertCommand) GetType() *string {
	return s.Type
}

func (s *UpsertQualityTemplateRequestUpsertCommand) SetCatalog(v string) *UpsertQualityTemplateRequestUpsertCommand {
	s.Catalog = &v
	return s
}

func (s *UpsertQualityTemplateRequestUpsertCommand) SetDescription(v string) *UpsertQualityTemplateRequestUpsertCommand {
	s.Description = &v
	return s
}

func (s *UpsertQualityTemplateRequestUpsertCommand) SetFormPropertyList(v []*UpsertQualityTemplateRequestUpsertCommandFormPropertyList) *UpsertQualityTemplateRequestUpsertCommand {
	s.FormPropertyList = v
	return s
}

func (s *UpsertQualityTemplateRequestUpsertCommand) SetId(v int64) *UpsertQualityTemplateRequestUpsertCommand {
	s.Id = &v
	return s
}

func (s *UpsertQualityTemplateRequestUpsertCommand) SetName(v string) *UpsertQualityTemplateRequestUpsertCommand {
	s.Name = &v
	return s
}

func (s *UpsertQualityTemplateRequestUpsertCommand) SetOwner(v string) *UpsertQualityTemplateRequestUpsertCommand {
	s.Owner = &v
	return s
}

func (s *UpsertQualityTemplateRequestUpsertCommand) SetSupportDataSourceTypeList(v []*string) *UpsertQualityTemplateRequestUpsertCommand {
	s.SupportDataSourceTypeList = v
	return s
}

func (s *UpsertQualityTemplateRequestUpsertCommand) SetType(v string) *UpsertQualityTemplateRequestUpsertCommand {
	s.Type = &v
	return s
}

func (s *UpsertQualityTemplateRequestUpsertCommand) Validate() error {
	if s.FormPropertyList != nil {
		for _, item := range s.FormPropertyList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpsertQualityTemplateRequestUpsertCommandFormPropertyList struct {
	// The type of the control.
	//
	// example:
	//
	// expression
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	// The name of the property.
	//
	// example:
	//
	// col
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the property.
	//
	// example:
	//
	// abc
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpsertQualityTemplateRequestUpsertCommandFormPropertyList) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityTemplateRequestUpsertCommandFormPropertyList) GoString() string {
	return s.String()
}

func (s *UpsertQualityTemplateRequestUpsertCommandFormPropertyList) GetComponentType() *string {
	return s.ComponentType
}

func (s *UpsertQualityTemplateRequestUpsertCommandFormPropertyList) GetName() *string {
	return s.Name
}

func (s *UpsertQualityTemplateRequestUpsertCommandFormPropertyList) GetValue() *string {
	return s.Value
}

func (s *UpsertQualityTemplateRequestUpsertCommandFormPropertyList) SetComponentType(v string) *UpsertQualityTemplateRequestUpsertCommandFormPropertyList {
	s.ComponentType = &v
	return s
}

func (s *UpsertQualityTemplateRequestUpsertCommandFormPropertyList) SetName(v string) *UpsertQualityTemplateRequestUpsertCommandFormPropertyList {
	s.Name = &v
	return s
}

func (s *UpsertQualityTemplateRequestUpsertCommandFormPropertyList) SetValue(v string) *UpsertQualityTemplateRequestUpsertCommandFormPropertyList {
	s.Value = &v
	return s
}

func (s *UpsertQualityTemplateRequestUpsertCommandFormPropertyList) Validate() error {
	return dara.Validate(s)
}
