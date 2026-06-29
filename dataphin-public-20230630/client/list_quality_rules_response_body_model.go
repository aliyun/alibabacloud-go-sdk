// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQualityRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListQualityRulesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListQualityRulesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListQualityRulesResponseBody
	GetMessage() *string
	SetPageResult(v *ListQualityRulesResponseBodyPageResult) *ListQualityRulesResponseBody
	GetPageResult() *ListQualityRulesResponseBodyPageResult
	SetRequestId(v string) *ListQualityRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListQualityRulesResponseBody
	GetSuccess() *bool
}

type ListQualityRulesResponseBody struct {
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
	// The backend exception details.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The paged query result.
	PageResult *ListQualityRulesResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListQualityRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListQualityRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListQualityRulesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListQualityRulesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListQualityRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListQualityRulesResponseBody) GetPageResult() *ListQualityRulesResponseBodyPageResult {
	return s.PageResult
}

func (s *ListQualityRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListQualityRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListQualityRulesResponseBody) SetCode(v string) *ListQualityRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListQualityRulesResponseBody) SetHttpStatusCode(v int32) *ListQualityRulesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListQualityRulesResponseBody) SetMessage(v string) *ListQualityRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListQualityRulesResponseBody) SetPageResult(v *ListQualityRulesResponseBodyPageResult) *ListQualityRulesResponseBody {
	s.PageResult = v
	return s
}

func (s *ListQualityRulesResponseBody) SetRequestId(v string) *ListQualityRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQualityRulesResponseBody) SetSuccess(v bool) *ListQualityRulesResponseBody {
	s.Success = &v
	return s
}

func (s *ListQualityRulesResponseBody) Validate() error {
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListQualityRulesResponseBodyPageResult struct {
	// The list of quality rules returned by page.
	QualityRuleList []*ListQualityRulesResponseBodyPageResultQualityRuleList `json:"QualityRuleList,omitempty" xml:"QualityRuleList,omitempty" type:"Repeated"`
	// The total number of records.
	//
	// example:
	//
	// 68
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListQualityRulesResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListQualityRulesResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListQualityRulesResponseBodyPageResult) GetQualityRuleList() []*ListQualityRulesResponseBodyPageResultQualityRuleList {
	return s.QualityRuleList
}

func (s *ListQualityRulesResponseBodyPageResult) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListQualityRulesResponseBodyPageResult) SetQualityRuleList(v []*ListQualityRulesResponseBodyPageResultQualityRuleList) *ListQualityRulesResponseBodyPageResult {
	s.QualityRuleList = v
	return s
}

func (s *ListQualityRulesResponseBodyPageResult) SetTotalCount(v int64) *ListQualityRulesResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResult) Validate() error {
	if s.QualityRuleList != nil {
		for _, item := range s.QualityRuleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListQualityRulesResponseBodyPageResultQualityRuleList struct {
	// The rule business attribute configurations.
	AttributeWithValueList []*ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueList `json:"AttributeWithValueList,omitempty" xml:"AttributeWithValueList,omitempty" type:"Repeated"`
	// The rule category. Valid values:
	//
	// - CONSISTENT: consistency.
	//
	// - EFFECTIVE: validity.
	//
	// - TIMELINESE: timeliness.
	//
	// - ACCURATE: accuracy.
	//
	// - UNIQUENESS: uniqueness.
	//
	// - COMPLETENESS: completeness.
	//
	// - STABILITY: stability.
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
	// The rule configuration key-value pairs. These vary by templatetype: different templatetypes return different form key-value pair configurations.
	FormPropertyList []*ListQualityRulesResponseBodyPageResultQualityRuleListFormPropertyList `json:"FormPropertyList,omitempty" xml:"FormPropertyList,omitempty" type:"Repeated"`
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
	ScheduleBindList []*ListQualityRulesResponseBodyPageResultQualityRuleListScheduleBindList `json:"ScheduleBindList,omitempty" xml:"ScheduleBindList,omitempty" type:"Repeated"`
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
	// - SYSTEM: system preset.
	//
	// - CUSTOM: custom template.
	//
	// - TEMPLATE: union of SYSTEM and CUSTOM.
	//
	// - CUSTOM_SQL: custom SQL template.
	//
	// example:
	//
	// CUSTOM
	TemplateScope *string `json:"TemplateScope,omitempty" xml:"TemplateScope,omitempty"`
	// The templatetype. Valid values:
	//
	//   - FIELD_NULL_VALUE_VALIDATE: field null value validation.
	//
	//   - FIELD_EMPTY_STRING_VALIDATE: field empty character string validation.
	//
	//   - FIELD_UNIQUE_VALIDATE: field uniqueness validation.
	//
	//   - FIELD_GROUP_COUNT_VALIDATE: field unique value count validation.
	//
	//   - FIELD_DUPLICATE_VALUE_COUNT_VALIDATE: field duplicate value count validation.
	//
	//   - FUNCTION_TIME_COMPARE: time function comparison.
	//
	//   - SINGLE_TABLE_TIME_COMPARE: non-partitioned table time field comparison.
	//
	//   - DOUBLE_TABLE_TIME_COMPARE: two-table time field comparison.
	//
	//   - FIELD_FORMAT_VALIDATE: field format validation.
	//
	//   - FIELD_LENGTH_VALIDATE: field length validation.
	//
	//   - FIELD_VALUE_RANGE_VALIDATE: field value range validation.
	//
	//   - CODE_TABLE_COMPARE: lookup table reference comparison.
	//
	//   - STANDARD_CODE_TABLE_COMPARE: data standard lookup table reference comparison.
	//
	//   - SINGLE_TABLE_FIELD_VALUE_COMPARE: non-partitioned table field value consistency comparison.
	//
	//   - SINGLE_TABLE_FIELD_STATISTICAL_COMPARE: non-partitioned table field statistical value consistency comparison.
	//
	//   - SINGLE_TABLE_FIELD_EXP_COMPARE: non-partitioned table field business logic consistency comparison.
	//
	//   - DOUBLE_TABLE_FIELD_VALUE_COMPARE: two-table field value consistency comparison.
	//
	//   - DOUBLE_TABLE_FIELD_STATISTICAL_COMPARE: two-table field statistical value consistency comparison.
	//
	//   - CROSS_DOUBLE_TABLE_FIELD_STATISTICAL_COMPARE: cross-source two-table field statistical value consistency comparison.
	//
	//   - DOUBLE_TABLE_FIELD_EXP_COMPARE: two-table field business logic consistency comparison.
	//
	//   - TABLE_STABILITY_VALIDATE: table stability validation.
	//
	//   - TABLE_FLUCTUATION_VALIDATE: table fluctuation validation.
	//
	//   - FIELD_STABILITY_VALIDATE: field stability validation.
	//
	//   - FIELD_FLUCTUATION_VALIDATE: field fluctuation validation.
	//
	//   - CUSTOM_STATISTICAL_VALIDATE: custom statistical metric validation.
	//
	//   - CUSTOM_DATA_DETAILS_VALIDATE: custom data details validation.
	//
	//   - DATASOURCE_AVAILABLE_CHECK: data source connectivity check.
	//
	//   - TABLE_SCHEMA_CHECK: table schema change detection.
	//
	//   - REAL_TIME_OFFLINE_COMPARE: real-time and offline comparison.
	//
	//   - REAL_TIME_STATISTICAL_VALIDATE: real-time statistical value monitoring.
	//
	//   - REAL_TIME_MULTI_CHAIN_COMPARE: real-time multi-link comparison, and more.
	//
	// example:
	//
	// FIELD_NULL_VALUE_VALIDATE
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// The ID of the most recent trial run task.
	//
	// example:
	//
	// 22
	TestRunRuleTaskId *int64 `json:"TestRunRuleTaskId,omitempty" xml:"TestRunRuleTaskId,omitempty"`
	// The status of the most recent trial run task. Valid values: NOT_RUN, WAITING, RUNNING, SUCCESS, FAILED.
	//
	// example:
	//
	// SUCCESS
	TestRunRuleTaskStatus *string `json:"TestRunRuleTaskStatus,omitempty" xml:"TestRunRuleTaskStatus,omitempty"`
	// Indicates whether the trial run validation passed.
	TestRunRuleValidateResult *bool `json:"TestRunRuleValidateResult,omitempty" xml:"TestRunRuleValidateResult,omitempty"`
	// The list of validation conditions.
	ValidateConditionList []*ListQualityRulesResponseBodyPageResultQualityRuleListValidateConditionList `json:"ValidateConditionList,omitempty" xml:"ValidateConditionList,omitempty" type:"Repeated"`
	// The validation object.
	ValidateObject *ListQualityRulesResponseBodyPageResultQualityRuleListValidateObject `json:"ValidateObject,omitempty" xml:"ValidateObject,omitempty" type:"Struct"`
	// The ID of the monitoring to which the rule belongs.
	//
	// example:
	//
	// 1
	WatchId *int64 `json:"WatchId,omitempty" xml:"WatchId,omitempty"`
}

func (s ListQualityRulesResponseBodyPageResultQualityRuleList) String() string {
	return dara.Prettify(s)
}

func (s ListQualityRulesResponseBodyPageResultQualityRuleList) GoString() string {
	return s.String()
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) GetAttributeWithValueList() []*ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueList {
	return s.AttributeWithValueList
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) GetCatalogList() []*string {
	return s.CatalogList
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) GetCreator() *string {
	return s.Creator
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) GetCreatorName() *string {
	return s.CreatorName
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) GetDescription() *string {
	return s.Description
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) GetEnableErrorArchive() *bool {
	return s.EnableErrorArchive
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) GetFormPropertyList() []*ListQualityRulesResponseBodyPageResultQualityRuleListFormPropertyList {
	return s.FormPropertyList
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) GetId() *int64 {
	return s.Id
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) GetModifier() *string {
	return s.Modifier
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) GetModifierName() *string {
	return s.ModifierName
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) GetName() *string {
	return s.Name
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) GetScheduleBindList() []*ListQualityRulesResponseBodyPageResultQualityRuleListScheduleBindList {
	return s.ScheduleBindList
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) GetStatus() *string {
	return s.Status
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) GetStrength() *string {
	return s.Strength
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) GetTemplateScope() *string {
	return s.TemplateScope
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) GetTemplateType() *string {
	return s.TemplateType
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) GetTestRunRuleTaskId() *int64 {
	return s.TestRunRuleTaskId
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) GetTestRunRuleTaskStatus() *string {
	return s.TestRunRuleTaskStatus
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) GetTestRunRuleValidateResult() *bool {
	return s.TestRunRuleValidateResult
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) GetValidateConditionList() []*ListQualityRulesResponseBodyPageResultQualityRuleListValidateConditionList {
	return s.ValidateConditionList
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) GetValidateObject() *ListQualityRulesResponseBodyPageResultQualityRuleListValidateObject {
	return s.ValidateObject
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) GetWatchId() *int64 {
	return s.WatchId
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) SetAttributeWithValueList(v []*ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueList) *ListQualityRulesResponseBodyPageResultQualityRuleList {
	s.AttributeWithValueList = v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) SetCatalogList(v []*string) *ListQualityRulesResponseBodyPageResultQualityRuleList {
	s.CatalogList = v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) SetCreateTime(v string) *ListQualityRulesResponseBodyPageResultQualityRuleList {
	s.CreateTime = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) SetCreator(v string) *ListQualityRulesResponseBodyPageResultQualityRuleList {
	s.Creator = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) SetCreatorName(v string) *ListQualityRulesResponseBodyPageResultQualityRuleList {
	s.CreatorName = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) SetDescription(v string) *ListQualityRulesResponseBodyPageResultQualityRuleList {
	s.Description = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) SetEnableErrorArchive(v bool) *ListQualityRulesResponseBodyPageResultQualityRuleList {
	s.EnableErrorArchive = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) SetFormPropertyList(v []*ListQualityRulesResponseBodyPageResultQualityRuleListFormPropertyList) *ListQualityRulesResponseBodyPageResultQualityRuleList {
	s.FormPropertyList = v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) SetId(v int64) *ListQualityRulesResponseBodyPageResultQualityRuleList {
	s.Id = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) SetModifier(v string) *ListQualityRulesResponseBodyPageResultQualityRuleList {
	s.Modifier = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) SetModifierName(v string) *ListQualityRulesResponseBodyPageResultQualityRuleList {
	s.ModifierName = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) SetModifyTime(v string) *ListQualityRulesResponseBodyPageResultQualityRuleList {
	s.ModifyTime = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) SetName(v string) *ListQualityRulesResponseBodyPageResultQualityRuleList {
	s.Name = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) SetScheduleBindList(v []*ListQualityRulesResponseBodyPageResultQualityRuleListScheduleBindList) *ListQualityRulesResponseBodyPageResultQualityRuleList {
	s.ScheduleBindList = v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) SetStatus(v string) *ListQualityRulesResponseBodyPageResultQualityRuleList {
	s.Status = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) SetStrength(v string) *ListQualityRulesResponseBodyPageResultQualityRuleList {
	s.Strength = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) SetTemplateId(v int64) *ListQualityRulesResponseBodyPageResultQualityRuleList {
	s.TemplateId = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) SetTemplateName(v string) *ListQualityRulesResponseBodyPageResultQualityRuleList {
	s.TemplateName = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) SetTemplateScope(v string) *ListQualityRulesResponseBodyPageResultQualityRuleList {
	s.TemplateScope = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) SetTemplateType(v string) *ListQualityRulesResponseBodyPageResultQualityRuleList {
	s.TemplateType = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) SetTestRunRuleTaskId(v int64) *ListQualityRulesResponseBodyPageResultQualityRuleList {
	s.TestRunRuleTaskId = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) SetTestRunRuleTaskStatus(v string) *ListQualityRulesResponseBodyPageResultQualityRuleList {
	s.TestRunRuleTaskStatus = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) SetTestRunRuleValidateResult(v bool) *ListQualityRulesResponseBodyPageResultQualityRuleList {
	s.TestRunRuleValidateResult = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) SetValidateConditionList(v []*ListQualityRulesResponseBodyPageResultQualityRuleListValidateConditionList) *ListQualityRulesResponseBodyPageResultQualityRuleList {
	s.ValidateConditionList = v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) SetValidateObject(v *ListQualityRulesResponseBodyPageResultQualityRuleListValidateObject) *ListQualityRulesResponseBodyPageResultQualityRuleList {
	s.ValidateObject = v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) SetWatchId(v int64) *ListQualityRulesResponseBodyPageResultQualityRuleList {
	s.WatchId = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleList) Validate() error {
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

type ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueList struct {
	// The attribute details.
	AttributeInfo *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfo `json:"AttributeInfo,omitempty" xml:"AttributeInfo,omitempty" type:"Struct"`
	// The property value.
	AttributeValue *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeValue `json:"AttributeValue,omitempty" xml:"AttributeValue,omitempty" type:"Struct"`
}

func (s ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueList) String() string {
	return dara.Prettify(s)
}

func (s ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueList) GoString() string {
	return s.String()
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueList) GetAttributeInfo() *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfo {
	return s.AttributeInfo
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueList) GetAttributeValue() *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeValue {
	return s.AttributeValue
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueList) SetAttributeInfo(v *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfo) *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueList {
	s.AttributeInfo = v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueList) SetAttributeValue(v *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeValue) *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueList {
	s.AttributeValue = v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueList) Validate() error {
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

type ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfo struct {
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
	// -168890138815
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
	ValueConfig *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfig `json:"ValueConfig,omitempty" xml:"ValueConfig,omitempty" type:"Struct"`
}

func (s ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfo) String() string {
	return dara.Prettify(s)
}

func (s ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfo) GoString() string {
	return s.String()
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfo) GetDescription() *string {
	return s.Description
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfo) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfo) GetId() *int64 {
	return s.Id
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfo) GetName() *string {
	return s.Name
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfo) GetRequired() *bool {
	return s.Required
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfo) GetSearchable() *bool {
	return s.Searchable
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfo) GetValueConfig() *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfig {
	return s.ValueConfig
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfo) SetDescription(v string) *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfo {
	s.Description = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfo) SetEnabled(v bool) *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfo {
	s.Enabled = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfo) SetId(v int64) *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfo {
	s.Id = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfo) SetName(v string) *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfo {
	s.Name = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfo) SetRequired(v bool) *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfo {
	s.Required = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfo) SetSearchable(v bool) *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfo {
	s.Searchable = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfo) SetValueConfig(v *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfig) *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfo {
	s.ValueConfig = v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfo) Validate() error {
	if s.ValueConfig != nil {
		if err := s.ValueConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfig struct {
	// The attribute field type. Valid values:
	//
	// - STRING: text.
	//
	// - BIGINT: integer.
	//
	// - DOUBLE: floating-point.
	//
	// - BOOLEAN: Boolean.
	//
	// - DATE: date.
	//
	// - DATETIME: datetime.
	//
	// example:
	//
	// STRING
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The attribute default value.
	DefaultValue *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfigDefaultValue `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty" type:"Struct"`
	// The attribute field length. This constrains the maximum length of text-type attribute values.
	//
	// example:
	//
	// 6997283
	Length *int32 `json:"Length,omitempty" xml:"Length,omitempty"`
	// The attribute value input method. Valid values:
	//
	// - CUSTOMIZED: custom input.
	//
	// - SINGLE_ENUM: single-select dropdown.
	//
	// - MULTIPLE_ENUMS: multi-select dropdown.
	//
	// - RANGE: range interval.
	//
	// example:
	//
	// CUSTOMIZED
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The attribute option values. This applies only to attributes with the input method set to single-select dropdown or multi-select dropdown.
	ValueEnumList []*string `json:"ValueEnumList,omitempty" xml:"ValueEnumList,omitempty" type:"Repeated"`
}

func (s ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfig) String() string {
	return dara.Prettify(s)
}

func (s ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfig) GoString() string {
	return s.String()
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfig) GetDataType() *string {
	return s.DataType
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfig) GetDefaultValue() *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfigDefaultValue {
	return s.DefaultValue
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfig) GetLength() *int32 {
	return s.Length
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfig) GetType() *string {
	return s.Type
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfig) GetValueEnumList() []*string {
	return s.ValueEnumList
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfig) SetDataType(v string) *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfig {
	s.DataType = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfig) SetDefaultValue(v *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfigDefaultValue) *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfig {
	s.DefaultValue = v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfig) SetLength(v int32) *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfig {
	s.Length = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfig) SetType(v string) *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfig {
	s.Type = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfig) SetValueEnumList(v []*string) *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfig {
	s.ValueEnumList = v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfig) Validate() error {
	if s.DefaultValue != nil {
		if err := s.DefaultValue.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfigDefaultValue struct {
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
	// The list of attribute values. This applies to attributes with the input method set to custom input, single-select dropdown, or multi-select dropdown.
	ValueList []*string `json:"ValueList,omitempty" xml:"ValueList,omitempty" type:"Repeated"`
}

func (s ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfigDefaultValue) String() string {
	return dara.Prettify(s)
}

func (s ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfigDefaultValue) GoString() string {
	return s.String()
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfigDefaultValue) GetIncludeMaxValue() *bool {
	return s.IncludeMaxValue
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfigDefaultValue) GetIncludeMinValue() *bool {
	return s.IncludeMinValue
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfigDefaultValue) GetMaxValue() *string {
	return s.MaxValue
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfigDefaultValue) GetMinValue() *string {
	return s.MinValue
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfigDefaultValue) GetValueList() []*string {
	return s.ValueList
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfigDefaultValue) SetIncludeMaxValue(v bool) *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfigDefaultValue {
	s.IncludeMaxValue = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfigDefaultValue) SetIncludeMinValue(v bool) *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfigDefaultValue {
	s.IncludeMinValue = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfigDefaultValue) SetMaxValue(v string) *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfigDefaultValue {
	s.MaxValue = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfigDefaultValue) SetMinValue(v string) *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfigDefaultValue {
	s.MinValue = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfigDefaultValue) SetValueList(v []*string) *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfigDefaultValue {
	s.ValueList = v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeInfoValueConfigDefaultValue) Validate() error {
	return dara.Validate(s)
}

type ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeValue struct {
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
	// The list of attribute values. This applies to attributes with the input method set to custom input, single-select dropdown, or multi-select dropdown.
	ValueList []*string `json:"ValueList,omitempty" xml:"ValueList,omitempty" type:"Repeated"`
}

func (s ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeValue) String() string {
	return dara.Prettify(s)
}

func (s ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeValue) GoString() string {
	return s.String()
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeValue) GetIncludeMaxValue() *bool {
	return s.IncludeMaxValue
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeValue) GetIncludeMinValue() *bool {
	return s.IncludeMinValue
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeValue) GetMaxValue() *string {
	return s.MaxValue
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeValue) GetMinValue() *string {
	return s.MinValue
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeValue) GetValueList() []*string {
	return s.ValueList
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeValue) SetIncludeMaxValue(v bool) *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeValue {
	s.IncludeMaxValue = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeValue) SetIncludeMinValue(v bool) *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeValue {
	s.IncludeMinValue = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeValue) SetMaxValue(v string) *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeValue {
	s.MaxValue = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeValue) SetMinValue(v string) *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeValue {
	s.MinValue = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeValue) SetValueList(v []*string) *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeValue {
	s.ValueList = v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListAttributeWithValueListAttributeValue) Validate() error {
	return dara.Validate(s)
}

type ListQualityRulesResponseBodyPageResultQualityRuleListFormPropertyList struct {
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

func (s ListQualityRulesResponseBodyPageResultQualityRuleListFormPropertyList) String() string {
	return dara.Prettify(s)
}

func (s ListQualityRulesResponseBodyPageResultQualityRuleListFormPropertyList) GoString() string {
	return s.String()
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListFormPropertyList) GetComponentType() *string {
	return s.ComponentType
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListFormPropertyList) GetName() *string {
	return s.Name
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListFormPropertyList) GetValue() *string {
	return s.Value
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListFormPropertyList) SetComponentType(v string) *ListQualityRulesResponseBodyPageResultQualityRuleListFormPropertyList {
	s.ComponentType = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListFormPropertyList) SetName(v string) *ListQualityRulesResponseBodyPageResultQualityRuleListFormPropertyList {
	s.Name = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListFormPropertyList) SetValue(v string) *ListQualityRulesResponseBodyPageResultQualityRuleListFormPropertyList {
	s.Value = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListFormPropertyList) Validate() error {
	return dara.Validate(s)
}

type ListQualityRulesResponseBodyPageResultQualityRuleListScheduleBindList struct {
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

func (s ListQualityRulesResponseBodyPageResultQualityRuleListScheduleBindList) String() string {
	return dara.Prettify(s)
}

func (s ListQualityRulesResponseBodyPageResultQualityRuleListScheduleBindList) GoString() string {
	return s.String()
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListScheduleBindList) GetScheduleId() *int64 {
	return s.ScheduleId
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListScheduleBindList) GetScheduleName() *string {
	return s.ScheduleName
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListScheduleBindList) SetScheduleId(v int64) *ListQualityRulesResponseBodyPageResultQualityRuleListScheduleBindList {
	s.ScheduleId = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListScheduleBindList) SetScheduleName(v string) *ListQualityRulesResponseBodyPageResultQualityRuleListScheduleBindList {
	s.ScheduleName = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListScheduleBindList) Validate() error {
	return dara.Validate(s)
}

type ListQualityRulesResponseBodyPageResultQualityRuleListValidateConditionList struct {
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
	// - RELATION: relation.
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

func (s ListQualityRulesResponseBodyPageResultQualityRuleListValidateConditionList) String() string {
	return dara.Prettify(s)
}

func (s ListQualityRulesResponseBodyPageResultQualityRuleListValidateConditionList) GoString() string {
	return s.String()
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListValidateConditionList) GetId() *string {
	return s.Id
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListValidateConditionList) GetMetric() *string {
	return s.Metric
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListValidateConditionList) GetMetricName() *string {
	return s.MetricName
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListValidateConditionList) GetOperator() *string {
	return s.Operator
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListValidateConditionList) GetOperatorName() *string {
	return s.OperatorName
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListValidateConditionList) GetParentId() *string {
	return s.ParentId
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListValidateConditionList) GetType() *string {
	return s.Type
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListValidateConditionList) GetValue() *string {
	return s.Value
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListValidateConditionList) SetId(v string) *ListQualityRulesResponseBodyPageResultQualityRuleListValidateConditionList {
	s.Id = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListValidateConditionList) SetMetric(v string) *ListQualityRulesResponseBodyPageResultQualityRuleListValidateConditionList {
	s.Metric = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListValidateConditionList) SetMetricName(v string) *ListQualityRulesResponseBodyPageResultQualityRuleListValidateConditionList {
	s.MetricName = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListValidateConditionList) SetOperator(v string) *ListQualityRulesResponseBodyPageResultQualityRuleListValidateConditionList {
	s.Operator = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListValidateConditionList) SetOperatorName(v string) *ListQualityRulesResponseBodyPageResultQualityRuleListValidateConditionList {
	s.OperatorName = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListValidateConditionList) SetParentId(v string) *ListQualityRulesResponseBodyPageResultQualityRuleListValidateConditionList {
	s.ParentId = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListValidateConditionList) SetType(v string) *ListQualityRulesResponseBodyPageResultQualityRuleListValidateConditionList {
	s.Type = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListValidateConditionList) SetValue(v string) *ListQualityRulesResponseBodyPageResultQualityRuleListValidateConditionList {
	s.Value = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListValidateConditionList) Validate() error {
	return dara.Validate(s)
}

type ListQualityRulesResponseBodyPageResultQualityRuleListValidateObject struct {
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

func (s ListQualityRulesResponseBodyPageResultQualityRuleListValidateObject) String() string {
	return dara.Prettify(s)
}

func (s ListQualityRulesResponseBodyPageResultQualityRuleListValidateObject) GoString() string {
	return s.String()
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListValidateObject) GetName() *string {
	return s.Name
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListValidateObject) GetType() *string {
	return s.Type
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListValidateObject) SetName(v string) *ListQualityRulesResponseBodyPageResultQualityRuleListValidateObject {
	s.Name = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListValidateObject) SetType(v string) *ListQualityRulesResponseBodyPageResultQualityRuleListValidateObject {
	s.Type = &v
	return s
}

func (s *ListQualityRulesResponseBodyPageResultQualityRuleListValidateObject) Validate() error {
	return dara.Validate(s)
}
