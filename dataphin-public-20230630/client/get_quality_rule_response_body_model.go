// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetQualityRuleResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetQualityRuleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetQualityRuleResponseBody
	GetMessage() *string
	SetQualityRuleInfo(v *GetQualityRuleResponseBodyQualityRuleInfo) *GetQualityRuleResponseBody
	GetQualityRuleInfo() *GetQualityRuleResponseBodyQualityRuleInfo
	SetRequestId(v string) *GetQualityRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetQualityRuleResponseBody
	GetSuccess() *bool
}

type GetQualityRuleResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The details of the backend exception.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The quality rule details.
	QualityRuleInfo *GetQualityRuleResponseBodyQualityRuleInfo `json:"QualityRuleInfo,omitempty" xml:"QualityRuleInfo,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQualityRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQualityRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetQualityRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetQualityRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetQualityRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetQualityRuleResponseBody) GetQualityRuleInfo() *GetQualityRuleResponseBodyQualityRuleInfo {
	return s.QualityRuleInfo
}

func (s *GetQualityRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQualityRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetQualityRuleResponseBody) SetCode(v string) *GetQualityRuleResponseBody {
	s.Code = &v
	return s
}

func (s *GetQualityRuleResponseBody) SetHttpStatusCode(v int32) *GetQualityRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetQualityRuleResponseBody) SetMessage(v string) *GetQualityRuleResponseBody {
	s.Message = &v
	return s
}

func (s *GetQualityRuleResponseBody) SetQualityRuleInfo(v *GetQualityRuleResponseBodyQualityRuleInfo) *GetQualityRuleResponseBody {
	s.QualityRuleInfo = v
	return s
}

func (s *GetQualityRuleResponseBody) SetRequestId(v string) *GetQualityRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQualityRuleResponseBody) SetSuccess(v bool) *GetQualityRuleResponseBody {
	s.Success = &v
	return s
}

func (s *GetQualityRuleResponseBody) Validate() error {
	if s.QualityRuleInfo != nil {
		if err := s.QualityRuleInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQualityRuleResponseBodyQualityRuleInfo struct {
	// The rule business attribute configurations.
	AttributeWithValueList []*GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueList `json:"AttributeWithValueList,omitempty" xml:"AttributeWithValueList,omitempty" type:"Repeated"`
	// The rule categories. Valid values:
	//
	// - CONSISTENT: consistency
	//
	// - EFFECTIVE: validity
	//
	// - TIMELINESE: timeliness
	//
	// - ACCURATE: accuracy
	//
	// - UNIQUENESS: uniqueness
	//
	// - COMPLETENESS: completeness
	//
	// - STABILITY: stability
	//
	// - CUSTOM: custom.
	CatalogList []*string `json:"CatalogList,omitempty" xml:"CatalogList,omitempty" type:"Repeated"`
	// The creation time.
	//
	// example:
	//
	// 2025-06-30 00:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The creator.
	//
	// example:
	//
	// 30012011
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The creator name.
	//
	// example:
	//
	// test
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// The description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether exception archiving is enabled.
	EnableErrorArchive *bool `json:"EnableErrorArchive,omitempty" xml:"EnableErrorArchive,omitempty"`
	// The rule configuration key-value pairs. These vary by templatetype. Different templatetypes return different form key-value pair configurations.
	FormPropertyList []*GetQualityRuleResponseBodyQualityRuleInfoFormPropertyList `json:"FormPropertyList,omitempty" xml:"FormPropertyList,omitempty" type:"Repeated"`
	// The quality rule ID.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the user who last modified the rule.
	//
	// example:
	//
	// 30012011
	Modifier *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// The name of the user who last modified the rule.
	//
	// example:
	//
	// test
	ModifierName *string `json:"ModifierName,omitempty" xml:"ModifierName,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2025-06-30 00:00:00
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The quality rule name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of schedules bound to the rule.
	ScheduleBindList []*GetQualityRuleResponseBodyQualityRuleInfoScheduleBindList `json:"ScheduleBindList,omitempty" xml:"ScheduleBindList,omitempty" type:"Repeated"`
	// The quality rule status. Valid values:
	//
	// - ENABLE
	//
	// - DISABLE.
	//
	// example:
	//
	// ENABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The rule strength. Valid values:
	//
	// - STRONG
	//
	// - WEAK.
	//
	// example:
	//
	// STRONG
	Strength *string `json:"Strength,omitempty" xml:"Strength,omitempty"`
	// The template ID.
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The template name.
	//
	// example:
	//
	// 1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The templatetype group. Valid values:
	//
	// - SYSTEM: system preset
	//
	// - CUSTOM: custom template
	//
	// - TEMPLATE: union of SYSTEM and CUSTOM
	//
	// - CUSTOM_SQL: custom SQL template.
	//
	// example:
	//
	// CUSTOM
	TemplateScope *string `json:"TemplateScope,omitempty" xml:"TemplateScope,omitempty"`
	// The templatetype. Valid values:
	//
	//   - FIELD_NULL_VALUE_VALIDATE: field null value check
	//
	//   - FIELD_EMPTY_STRING_VALIDATE: field empty character string check
	//
	//   - FIELD_UNIQUE_VALIDATE: field uniqueness check
	//
	//   - FIELD_GROUP_COUNT_VALIDATE: field unique value count check
	//
	//   - FIELD_DUPLICATE_VALUE_COUNT_VALIDATE: field duplicate value count check
	//
	//   - FUNCTION_TIME_COMPARE: time function comparison
	//
	//   - SINGLE_TABLE_TIME_COMPARE: non-partitioned table time field comparison
	//
	//   - DOUBLE_TABLE_TIME_COMPARE: two-table time field comparison
	//
	//   - FIELD_FORMAT_VALIDATE: field format check
	//
	//   - FIELD_LENGTH_VALIDATE: field length check
	//
	//   - FIELD_VALUE_RANGE_VALIDATE: field value range check
	//
	//   - CODE_TABLE_COMPARE: lookup table reference comparison
	//
	//   - STANDARD_CODE_TABLE_COMPARE: data standard lookup table reference comparison
	//
	//   - SINGLE_TABLE_FIELD_VALUE_COMPARE: non-partitioned table field value consistency comparison
	//
	//   - SINGLE_TABLE_FIELD_STATISTICAL_COMPARE: non-partitioned table field statistical value consistency comparison
	//
	//   - SINGLE_TABLE_FIELD_EXP_COMPARE: non-partitioned table field business logic consistency comparison
	//
	//   - DOUBLE_TABLE_FIELD_VALUE_COMPARE: two-table field value consistency comparison
	//
	//   - DOUBLE_TABLE_FIELD_STATISTICAL_COMPARE: two-table field statistical value consistency comparison
	//
	//   - CROSS_DOUBLE_TABLE_FIELD_STATISTICAL_COMPARE: cross-source two-table field statistical value consistency comparison
	//
	//   - DOUBLE_TABLE_FIELD_EXP_COMPARE: two-table field business logic consistency comparison
	//
	//   - TABLE_STABILITY_VALIDATE: table stability check
	//
	//   - TABLE_FLUCTUATION_VALIDATE: table fluctuation check
	//
	//   - FIELD_STABILITY_VALIDATE: field stability check
	//
	//   - FIELD_FLUCTUATION_VALIDATE: field fluctuation check
	//
	//   - CUSTOM_STATISTICAL_VALIDATE: custom statistical metric check
	//
	//   - CUSTOM_DATA_DETAILS_VALIDATE: custom data details check
	//
	//   - DATASOURCE_AVAILABLE_CHECK: data source connectivity monitoring
	//
	//   - TABLE_SCHEMA_CHECK: table schema change monitoring
	//
	//   - REAL_TIME_OFFLINE_COMPARE: real-time and offline comparison
	//
	//   - REAL_TIME_STATISTICAL_VALIDATE: real-time statistical value monitoring
	//
	//   - REAL_TIME_MULTI_CHAIN_COMPARE: real-time multi-link comparison, and more.
	//
	// example:
	//
	// FIELD_NULL_VALUE_VALIDATE
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// The ID of the most recent test run task.
	//
	// example:
	//
	// 22
	TestRunRuleTaskId *int64 `json:"TestRunRuleTaskId,omitempty" xml:"TestRunRuleTaskId,omitempty"`
	// The status of the most recent test run task. Valid values: NOT_RUN, WAITING, RUNNING, SUCCESS, FAILED.
	//
	// example:
	//
	// SUCCESS
	TestRunRuleTaskStatus *string `json:"TestRunRuleTaskStatus,omitempty" xml:"TestRunRuleTaskStatus,omitempty"`
	// Indicates whether the test run validation passed.
	TestRunRuleValidateResult *bool `json:"TestRunRuleValidateResult,omitempty" xml:"TestRunRuleValidateResult,omitempty"`
	// The list of validation conditions.
	ValidateConditionList []*GetQualityRuleResponseBodyQualityRuleInfoValidateConditionList `json:"ValidateConditionList,omitempty" xml:"ValidateConditionList,omitempty" type:"Repeated"`
	// The validation object.
	ValidateObject *GetQualityRuleResponseBodyQualityRuleInfoValidateObject `json:"ValidateObject,omitempty" xml:"ValidateObject,omitempty" type:"Struct"`
	// The ID of the associated monitoring task.
	//
	// example:
	//
	// 1
	WatchId *int64 `json:"WatchId,omitempty" xml:"WatchId,omitempty"`
}

func (s GetQualityRuleResponseBodyQualityRuleInfo) String() string {
	return dara.Prettify(s)
}

func (s GetQualityRuleResponseBodyQualityRuleInfo) GoString() string {
	return s.String()
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) GetAttributeWithValueList() []*GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueList {
	return s.AttributeWithValueList
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) GetCatalogList() []*string {
	return s.CatalogList
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) GetCreator() *string {
	return s.Creator
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) GetCreatorName() *string {
	return s.CreatorName
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) GetDescription() *string {
	return s.Description
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) GetEnableErrorArchive() *bool {
	return s.EnableErrorArchive
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) GetFormPropertyList() []*GetQualityRuleResponseBodyQualityRuleInfoFormPropertyList {
	return s.FormPropertyList
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) GetId() *int64 {
	return s.Id
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) GetModifier() *string {
	return s.Modifier
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) GetModifierName() *string {
	return s.ModifierName
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) GetName() *string {
	return s.Name
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) GetScheduleBindList() []*GetQualityRuleResponseBodyQualityRuleInfoScheduleBindList {
	return s.ScheduleBindList
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) GetStatus() *string {
	return s.Status
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) GetStrength() *string {
	return s.Strength
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) GetTemplateScope() *string {
	return s.TemplateScope
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) GetTemplateType() *string {
	return s.TemplateType
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) GetTestRunRuleTaskId() *int64 {
	return s.TestRunRuleTaskId
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) GetTestRunRuleTaskStatus() *string {
	return s.TestRunRuleTaskStatus
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) GetTestRunRuleValidateResult() *bool {
	return s.TestRunRuleValidateResult
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) GetValidateConditionList() []*GetQualityRuleResponseBodyQualityRuleInfoValidateConditionList {
	return s.ValidateConditionList
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) GetValidateObject() *GetQualityRuleResponseBodyQualityRuleInfoValidateObject {
	return s.ValidateObject
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) GetWatchId() *int64 {
	return s.WatchId
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) SetAttributeWithValueList(v []*GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueList) *GetQualityRuleResponseBodyQualityRuleInfo {
	s.AttributeWithValueList = v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) SetCatalogList(v []*string) *GetQualityRuleResponseBodyQualityRuleInfo {
	s.CatalogList = v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) SetCreateTime(v string) *GetQualityRuleResponseBodyQualityRuleInfo {
	s.CreateTime = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) SetCreator(v string) *GetQualityRuleResponseBodyQualityRuleInfo {
	s.Creator = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) SetCreatorName(v string) *GetQualityRuleResponseBodyQualityRuleInfo {
	s.CreatorName = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) SetDescription(v string) *GetQualityRuleResponseBodyQualityRuleInfo {
	s.Description = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) SetEnableErrorArchive(v bool) *GetQualityRuleResponseBodyQualityRuleInfo {
	s.EnableErrorArchive = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) SetFormPropertyList(v []*GetQualityRuleResponseBodyQualityRuleInfoFormPropertyList) *GetQualityRuleResponseBodyQualityRuleInfo {
	s.FormPropertyList = v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) SetId(v int64) *GetQualityRuleResponseBodyQualityRuleInfo {
	s.Id = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) SetModifier(v string) *GetQualityRuleResponseBodyQualityRuleInfo {
	s.Modifier = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) SetModifierName(v string) *GetQualityRuleResponseBodyQualityRuleInfo {
	s.ModifierName = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) SetModifyTime(v string) *GetQualityRuleResponseBodyQualityRuleInfo {
	s.ModifyTime = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) SetName(v string) *GetQualityRuleResponseBodyQualityRuleInfo {
	s.Name = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) SetScheduleBindList(v []*GetQualityRuleResponseBodyQualityRuleInfoScheduleBindList) *GetQualityRuleResponseBodyQualityRuleInfo {
	s.ScheduleBindList = v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) SetStatus(v string) *GetQualityRuleResponseBodyQualityRuleInfo {
	s.Status = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) SetStrength(v string) *GetQualityRuleResponseBodyQualityRuleInfo {
	s.Strength = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) SetTemplateId(v int64) *GetQualityRuleResponseBodyQualityRuleInfo {
	s.TemplateId = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) SetTemplateName(v string) *GetQualityRuleResponseBodyQualityRuleInfo {
	s.TemplateName = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) SetTemplateScope(v string) *GetQualityRuleResponseBodyQualityRuleInfo {
	s.TemplateScope = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) SetTemplateType(v string) *GetQualityRuleResponseBodyQualityRuleInfo {
	s.TemplateType = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) SetTestRunRuleTaskId(v int64) *GetQualityRuleResponseBodyQualityRuleInfo {
	s.TestRunRuleTaskId = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) SetTestRunRuleTaskStatus(v string) *GetQualityRuleResponseBodyQualityRuleInfo {
	s.TestRunRuleTaskStatus = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) SetTestRunRuleValidateResult(v bool) *GetQualityRuleResponseBodyQualityRuleInfo {
	s.TestRunRuleValidateResult = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) SetValidateConditionList(v []*GetQualityRuleResponseBodyQualityRuleInfoValidateConditionList) *GetQualityRuleResponseBodyQualityRuleInfo {
	s.ValidateConditionList = v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) SetValidateObject(v *GetQualityRuleResponseBodyQualityRuleInfoValidateObject) *GetQualityRuleResponseBodyQualityRuleInfo {
	s.ValidateObject = v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) SetWatchId(v int64) *GetQualityRuleResponseBodyQualityRuleInfo {
	s.WatchId = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfo) Validate() error {
	if s.AttributeWithValueList != nil {
		for _, item := range s.AttributeWithValueList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FormPropertyList != nil {
		for _, item := range s.FormPropertyList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ScheduleBindList != nil {
		for _, item := range s.ScheduleBindList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ValidateConditionList != nil {
		for _, item := range s.ValidateConditionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ValidateObject != nil {
		if err := s.ValidateObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueList struct {
	// The attribute details.
	AttributeInfo *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfo `json:"AttributeInfo,omitempty" xml:"AttributeInfo,omitempty" type:"Struct"`
	// The property value.
	AttributeValue *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeValue `json:"AttributeValue,omitempty" xml:"AttributeValue,omitempty" type:"Struct"`
}

func (s GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueList) String() string {
	return dara.Prettify(s)
}

func (s GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueList) GoString() string {
	return s.String()
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueList) GetAttributeInfo() *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfo {
	return s.AttributeInfo
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueList) GetAttributeValue() *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeValue {
	return s.AttributeValue
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueList) SetAttributeInfo(v *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfo) *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueList {
	s.AttributeInfo = v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueList) SetAttributeValue(v *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeValue) *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueList {
	s.AttributeValue = v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueList) Validate() error {
	if s.AttributeInfo != nil {
		if err := s.AttributeInfo.Validate(); err != nil {
			return err
		}
	}
	if s.AttributeValue != nil {
		if err := s.AttributeValue.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfo struct {
	// The description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the attribute is enabled.
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The attribute ID.
	//
	// example:
	//
	// 612415804007
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The attribute name.
	//
	// example:
	//
	// attr01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the attribute is required.
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
	// Indicates whether the attribute is searchable.
	Searchable *bool `json:"Searchable,omitempty" xml:"Searchable,omitempty"`
	// The attribute value configuration details.
	ValueConfig *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfig `json:"ValueConfig,omitempty" xml:"ValueConfig,omitempty" type:"Struct"`
}

func (s GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfo) String() string {
	return dara.Prettify(s)
}

func (s GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfo) GoString() string {
	return s.String()
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfo) GetDescription() *string {
	return s.Description
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfo) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfo) GetId() *int64 {
	return s.Id
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfo) GetName() *string {
	return s.Name
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfo) GetRequired() *bool {
	return s.Required
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfo) GetSearchable() *bool {
	return s.Searchable
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfo) GetValueConfig() *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfig {
	return s.ValueConfig
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfo) SetDescription(v string) *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfo {
	s.Description = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfo) SetEnabled(v bool) *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfo {
	s.Enabled = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfo) SetId(v int64) *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfo {
	s.Id = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfo) SetName(v string) *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfo {
	s.Name = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfo) SetRequired(v bool) *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfo {
	s.Required = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfo) SetSearchable(v bool) *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfo {
	s.Searchable = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfo) SetValueConfig(v *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfig) *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfo {
	s.ValueConfig = v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfo) Validate() error {
	if s.ValueConfig != nil {
		if err := s.ValueConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfig struct {
	// The attribute field data type. Valid values:
	//
	// - STRING: text
	//
	// - BIGINT: integer
	//
	// - DOUBLE: floating-point
	//
	// - BOOLEAN: Boolean
	//
	// - DATE: date
	//
	// - DATETIME: datetime.
	//
	// example:
	//
	// STRING
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The attribute default value.
	DefaultValue *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfigDefaultValue `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty" type:"Struct"`
	// The attribute field length. This constrains the maximum length of text-type attribute values.
	//
	// example:
	//
	// 6921666
	Length *int32 `json:"Length,omitempty" xml:"Length,omitempty"`
	// The attribute value input method. Valid values:
	//
	// - CUSTOMIZED: custom input
	//
	// - SINGLE_ENUM: single-select dropdown
	//
	// - MULTIPLE_ENUMS: multi-select dropdown
	//
	// - RANGE: range interval.
	//
	// example:
	//
	// CUSTOMIZED
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The attribute option values. This applies only to attributes with a single-select dropdown or multi-select dropdown input method.
	ValueEnumList []*string `json:"ValueEnumList,omitempty" xml:"ValueEnumList,omitempty" type:"Repeated"`
}

func (s GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfig) String() string {
	return dara.Prettify(s)
}

func (s GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfig) GoString() string {
	return s.String()
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfig) GetDataType() *string {
	return s.DataType
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfig) GetDefaultValue() *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfigDefaultValue {
	return s.DefaultValue
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfig) GetLength() *int32 {
	return s.Length
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfig) GetType() *string {
	return s.Type
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfig) GetValueEnumList() []*string {
	return s.ValueEnumList
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfig) SetDataType(v string) *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfig {
	s.DataType = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfig) SetDefaultValue(v *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfigDefaultValue) *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfig {
	s.DefaultValue = v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfig) SetLength(v int32) *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfig {
	s.Length = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfig) SetType(v string) *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfig {
	s.Type = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfig) SetValueEnumList(v []*string) *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfig {
	s.ValueEnumList = v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfig) Validate() error {
	if s.DefaultValue != nil {
		if err := s.DefaultValue.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfigDefaultValue struct {
	// Indicates whether the maximum value is included.
	IncludeMaxValue *bool `json:"IncludeMaxValue,omitempty" xml:"IncludeMaxValue,omitempty"`
	// Indicates whether the minimum value is included.
	IncludeMinValue *bool `json:"IncludeMinValue,omitempty" xml:"IncludeMinValue,omitempty"`
	// The maximum value. This applies to range interval attributes.
	//
	// example:
	//
	// 100
	MaxValue *string `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	// The minimum value. This applies to range interval attributes.
	//
	// example:
	//
	// 1
	MinValue *string `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	// The list of attribute values. This applies to attributes with custom input, single-select dropdown, or multi-select dropdown input methods.
	ValueList []*string `json:"ValueList,omitempty" xml:"ValueList,omitempty" type:"Repeated"`
}

func (s GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfigDefaultValue) String() string {
	return dara.Prettify(s)
}

func (s GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfigDefaultValue) GoString() string {
	return s.String()
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfigDefaultValue) GetIncludeMaxValue() *bool {
	return s.IncludeMaxValue
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfigDefaultValue) GetIncludeMinValue() *bool {
	return s.IncludeMinValue
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfigDefaultValue) GetMaxValue() *string {
	return s.MaxValue
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfigDefaultValue) GetMinValue() *string {
	return s.MinValue
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfigDefaultValue) GetValueList() []*string {
	return s.ValueList
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfigDefaultValue) SetIncludeMaxValue(v bool) *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfigDefaultValue {
	s.IncludeMaxValue = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfigDefaultValue) SetIncludeMinValue(v bool) *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfigDefaultValue {
	s.IncludeMinValue = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfigDefaultValue) SetMaxValue(v string) *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfigDefaultValue {
	s.MaxValue = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfigDefaultValue) SetMinValue(v string) *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfigDefaultValue {
	s.MinValue = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfigDefaultValue) SetValueList(v []*string) *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfigDefaultValue {
	s.ValueList = v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfigDefaultValue) Validate() error {
	return dara.Validate(s)
}

type GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeValue struct {
	// Indicates whether the maximum value is included.
	IncludeMaxValue *bool `json:"IncludeMaxValue,omitempty" xml:"IncludeMaxValue,omitempty"`
	// Indicates whether the minimum value is included.
	IncludeMinValue *bool `json:"IncludeMinValue,omitempty" xml:"IncludeMinValue,omitempty"`
	// The maximum value. This applies to range interval attributes.
	//
	// example:
	//
	// 100
	MaxValue *string `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	// The minimum value. This applies to range interval attributes.
	//
	// example:
	//
	// 1
	MinValue *string `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	// The list of attribute values. This applies to attributes with custom input, single-select dropdown, or multi-select dropdown input methods.
	ValueList []*string `json:"ValueList,omitempty" xml:"ValueList,omitempty" type:"Repeated"`
}

func (s GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeValue) String() string {
	return dara.Prettify(s)
}

func (s GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeValue) GoString() string {
	return s.String()
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeValue) GetIncludeMaxValue() *bool {
	return s.IncludeMaxValue
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeValue) GetIncludeMinValue() *bool {
	return s.IncludeMinValue
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeValue) GetMaxValue() *string {
	return s.MaxValue
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeValue) GetMinValue() *string {
	return s.MinValue
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeValue) GetValueList() []*string {
	return s.ValueList
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeValue) SetIncludeMaxValue(v bool) *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeValue {
	s.IncludeMaxValue = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeValue) SetIncludeMinValue(v bool) *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeValue {
	s.IncludeMinValue = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeValue) SetMaxValue(v string) *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeValue {
	s.MaxValue = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeValue) SetMinValue(v string) *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeValue {
	s.MinValue = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeValue) SetValueList(v []*string) *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeValue {
	s.ValueList = v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeValue) Validate() error {
	return dara.Validate(s)
}

type GetQualityRuleResponseBodyQualityRuleInfoFormPropertyList struct {
	// The control type.
	//
	// example:
	//
	// expression
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	// The property name.
	//
	// example:
	//
	// col
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The property value.
	//
	// example:
	//
	// abc
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetQualityRuleResponseBodyQualityRuleInfoFormPropertyList) String() string {
	return dara.Prettify(s)
}

func (s GetQualityRuleResponseBodyQualityRuleInfoFormPropertyList) GoString() string {
	return s.String()
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoFormPropertyList) GetComponentType() *string {
	return s.ComponentType
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoFormPropertyList) GetName() *string {
	return s.Name
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoFormPropertyList) GetValue() *string {
	return s.Value
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoFormPropertyList) SetComponentType(v string) *GetQualityRuleResponseBodyQualityRuleInfoFormPropertyList {
	s.ComponentType = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoFormPropertyList) SetName(v string) *GetQualityRuleResponseBodyQualityRuleInfoFormPropertyList {
	s.Name = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoFormPropertyList) SetValue(v string) *GetQualityRuleResponseBodyQualityRuleInfoFormPropertyList {
	s.Value = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoFormPropertyList) Validate() error {
	return dara.Validate(s)
}

type GetQualityRuleResponseBodyQualityRuleInfoScheduleBindList struct {
	// The schedule ID.
	//
	// example:
	//
	// 1
	ScheduleId *int64 `json:"ScheduleId,omitempty" xml:"ScheduleId,omitempty"`
	// The schedule name.
	//
	// example:
	//
	// test
	ScheduleName *string `json:"ScheduleName,omitempty" xml:"ScheduleName,omitempty"`
}

func (s GetQualityRuleResponseBodyQualityRuleInfoScheduleBindList) String() string {
	return dara.Prettify(s)
}

func (s GetQualityRuleResponseBodyQualityRuleInfoScheduleBindList) GoString() string {
	return s.String()
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoScheduleBindList) GetScheduleId() *int64 {
	return s.ScheduleId
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoScheduleBindList) GetScheduleName() *string {
	return s.ScheduleName
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoScheduleBindList) SetScheduleId(v int64) *GetQualityRuleResponseBodyQualityRuleInfoScheduleBindList {
	s.ScheduleId = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoScheduleBindList) SetScheduleName(v string) *GetQualityRuleResponseBodyQualityRuleInfoScheduleBindList {
	s.ScheduleName = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoScheduleBindList) Validate() error {
	return dara.Validate(s)
}

type GetQualityRuleResponseBodyQualityRuleInfoValidateConditionList struct {
	// The condition node ID.
	//
	// example:
	//
	// 268
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The metric.
	//
	// example:
	//
	// test
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// The metric name.
	//
	// example:
	//
	// test
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The operator. Valid values: EQUAL, NOT_EQUAL, LARGER, SMALLER, LARGE_OR_EQUAL, SMALLER_OR_EQUAL, AND, OR.
	//
	// example:
	//
	// AND
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The operator name.
	//
	// example:
	//
	// 且
	OperatorName *string `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
	// The parent condition node ID.
	//
	// example:
	//
	// 123
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The condition type. Valid values:
	//
	// - RELATION: relationship
	//
	// - EXPRESSION: expression.
	//
	// example:
	//
	// RELATION
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value.
	//
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetQualityRuleResponseBodyQualityRuleInfoValidateConditionList) String() string {
	return dara.Prettify(s)
}

func (s GetQualityRuleResponseBodyQualityRuleInfoValidateConditionList) GoString() string {
	return s.String()
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoValidateConditionList) GetId() *string {
	return s.Id
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoValidateConditionList) GetMetric() *string {
	return s.Metric
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoValidateConditionList) GetMetricName() *string {
	return s.MetricName
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoValidateConditionList) GetOperator() *string {
	return s.Operator
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoValidateConditionList) GetOperatorName() *string {
	return s.OperatorName
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoValidateConditionList) GetParentId() *string {
	return s.ParentId
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoValidateConditionList) GetType() *string {
	return s.Type
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoValidateConditionList) GetValue() *string {
	return s.Value
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoValidateConditionList) SetId(v string) *GetQualityRuleResponseBodyQualityRuleInfoValidateConditionList {
	s.Id = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoValidateConditionList) SetMetric(v string) *GetQualityRuleResponseBodyQualityRuleInfoValidateConditionList {
	s.Metric = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoValidateConditionList) SetMetricName(v string) *GetQualityRuleResponseBodyQualityRuleInfoValidateConditionList {
	s.MetricName = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoValidateConditionList) SetOperator(v string) *GetQualityRuleResponseBodyQualityRuleInfoValidateConditionList {
	s.Operator = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoValidateConditionList) SetOperatorName(v string) *GetQualityRuleResponseBodyQualityRuleInfoValidateConditionList {
	s.OperatorName = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoValidateConditionList) SetParentId(v string) *GetQualityRuleResponseBodyQualityRuleInfoValidateConditionList {
	s.ParentId = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoValidateConditionList) SetType(v string) *GetQualityRuleResponseBodyQualityRuleInfoValidateConditionList {
	s.Type = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoValidateConditionList) SetValue(v string) *GetQualityRuleResponseBodyQualityRuleInfoValidateConditionList {
	s.Value = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoValidateConditionList) Validate() error {
	return dara.Validate(s)
}

type GetQualityRuleResponseBodyQualityRuleInfoValidateObject struct {
	// The validation object name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The validation object type. Valid values: UNKNOWN, TABLE, COLUMN, DATASOURCE, DATASOURCE_TABLE, REALTIME, INDEX, CHAIN.
	//
	// example:
	//
	// TABLE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetQualityRuleResponseBodyQualityRuleInfoValidateObject) String() string {
	return dara.Prettify(s)
}

func (s GetQualityRuleResponseBodyQualityRuleInfoValidateObject) GoString() string {
	return s.String()
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoValidateObject) GetName() *string {
	return s.Name
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoValidateObject) GetType() *string {
	return s.Type
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoValidateObject) SetName(v string) *GetQualityRuleResponseBodyQualityRuleInfoValidateObject {
	s.Name = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoValidateObject) SetType(v string) *GetQualityRuleResponseBodyQualityRuleInfoValidateObject {
	s.Type = &v
	return s
}

func (s *GetQualityRuleResponseBodyQualityRuleInfoValidateObject) Validate() error {
	return dara.Validate(s)
}
