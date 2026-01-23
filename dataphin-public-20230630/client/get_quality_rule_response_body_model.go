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
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message         *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	QualityRuleInfo *GetQualityRuleResponseBodyQualityRuleInfo `json:"QualityRuleInfo,omitempty" xml:"QualityRuleInfo,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	AttributeWithValueList []*GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueList `json:"AttributeWithValueList,omitempty" xml:"AttributeWithValueList,omitempty" type:"Repeated"`
	CatalogList            []*string                                                          `json:"CatalogList,omitempty" xml:"CatalogList,omitempty" type:"Repeated"`
	// example:
	//
	// 2025-06-30 00:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 30012011
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// test
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// example:
	//
	// test
	Description        *string                                                      `json:"Description,omitempty" xml:"Description,omitempty"`
	EnableErrorArchive *bool                                                        `json:"EnableErrorArchive,omitempty" xml:"EnableErrorArchive,omitempty"`
	FormPropertyList   []*GetQualityRuleResponseBodyQualityRuleInfoFormPropertyList `json:"FormPropertyList,omitempty" xml:"FormPropertyList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 30012011
	Modifier *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// example:
	//
	// test
	ModifierName *string `json:"ModifierName,omitempty" xml:"ModifierName,omitempty"`
	// example:
	//
	// 2025-06-30 00:00:00
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// test
	Name             *string                                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	ScheduleBindList []*GetQualityRuleResponseBodyQualityRuleInfoScheduleBindList `json:"ScheduleBindList,omitempty" xml:"ScheduleBindList,omitempty" type:"Repeated"`
	// example:
	//
	// ENABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// STRONG
	Strength *string `json:"Strength,omitempty" xml:"Strength,omitempty"`
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// 1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// CUSTOM
	TemplateScope *string `json:"TemplateScope,omitempty" xml:"TemplateScope,omitempty"`
	// example:
	//
	// FIELD_NULL_VALUE_VALIDATE
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// example:
	//
	// 22
	TestRunRuleTaskId *int64 `json:"TestRunRuleTaskId,omitempty" xml:"TestRunRuleTaskId,omitempty"`
	// example:
	//
	// SUCCESS
	TestRunRuleTaskStatus     *string                                                           `json:"TestRunRuleTaskStatus,omitempty" xml:"TestRunRuleTaskStatus,omitempty"`
	TestRunRuleValidateResult *bool                                                             `json:"TestRunRuleValidateResult,omitempty" xml:"TestRunRuleValidateResult,omitempty"`
	ValidateConditionList     []*GetQualityRuleResponseBodyQualityRuleInfoValidateConditionList `json:"ValidateConditionList,omitempty" xml:"ValidateConditionList,omitempty" type:"Repeated"`
	ValidateObject            *GetQualityRuleResponseBodyQualityRuleInfoValidateObject          `json:"ValidateObject,omitempty" xml:"ValidateObject,omitempty" type:"Struct"`
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
	AttributeInfo  *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfo  `json:"AttributeInfo,omitempty" xml:"AttributeInfo,omitempty" type:"Struct"`
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
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Enabled     *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// 612415804007
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// attr01
	Name        *string                                                                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	Required    *bool                                                                                    `json:"Required,omitempty" xml:"Required,omitempty"`
	Searchable  *bool                                                                                    `json:"Searchable,omitempty" xml:"Searchable,omitempty"`
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
	// example:
	//
	// STRING
	DataType     *string                                                                                              `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DefaultValue *GetQualityRuleResponseBodyQualityRuleInfoAttributeWithValueListAttributeInfoValueConfigDefaultValue `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty" type:"Struct"`
	// example:
	//
	// 6921666
	Length *int32 `json:"Length,omitempty" xml:"Length,omitempty"`
	// example:
	//
	// CUSTOMIZED
	Type          *string   `json:"Type,omitempty" xml:"Type,omitempty"`
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
	IncludeMaxValue *bool `json:"IncludeMaxValue,omitempty" xml:"IncludeMaxValue,omitempty"`
	IncludeMinValue *bool `json:"IncludeMinValue,omitempty" xml:"IncludeMinValue,omitempty"`
	// example:
	//
	// 100
	MaxValue *string `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	// example:
	//
	// 1
	MinValue  *string   `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
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
	IncludeMaxValue *bool `json:"IncludeMaxValue,omitempty" xml:"IncludeMaxValue,omitempty"`
	IncludeMinValue *bool `json:"IncludeMinValue,omitempty" xml:"IncludeMinValue,omitempty"`
	// example:
	//
	// 100
	MaxValue *string `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	// example:
	//
	// 1
	MinValue  *string   `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
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
	// example:
	//
	// expression
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	// example:
	//
	// col
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// example:
	//
	// 1
	ScheduleId *int64 `json:"ScheduleId,omitempty" xml:"ScheduleId,omitempty"`
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
	// example:
	//
	// 268
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// test
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// example:
	//
	// test
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// example:
	//
	// AND
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// example:
	//
	// 且
	OperatorName *string `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
	// example:
	//
	// 123
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// example:
	//
	// RELATION
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
