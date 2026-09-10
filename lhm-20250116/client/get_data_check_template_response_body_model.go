// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCheckTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDataCheckTemplateResponseBodyData) *GetDataCheckTemplateResponseBody
	GetData() *GetDataCheckTemplateResponseBodyData
	SetErrCode(v string) *GetDataCheckTemplateResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *GetDataCheckTemplateResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *GetDataCheckTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataCheckTemplateResponseBody
	GetSuccess() *bool
}

type GetDataCheckTemplateResponseBody struct {
	// The data body returned by the operation. For the field structure, refer to the child field descriptions below.
	Data *GetDataCheckTemplateResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error code. An empty string is returned if the call is successful.
	//
	// example:
	//
	// Success
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message. An empty string is returned if the call is successful.
	//
	// example:
	//
	// success
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 4C467B38-3910-4477-9B0B-6963D83B4E72
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful. Valid values: true: The call is successful. false: The call failed. If the call failed, check errCode and errMessage for details.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetDataCheckTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataCheckTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataCheckTemplateResponseBody) GetData() *GetDataCheckTemplateResponseBodyData {
	return s.Data
}

func (s *GetDataCheckTemplateResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetDataCheckTemplateResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *GetDataCheckTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataCheckTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataCheckTemplateResponseBody) SetData(v *GetDataCheckTemplateResponseBodyData) *GetDataCheckTemplateResponseBody {
	s.Data = v
	return s
}

func (s *GetDataCheckTemplateResponseBody) SetErrCode(v string) *GetDataCheckTemplateResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetDataCheckTemplateResponseBody) SetErrMessage(v string) *GetDataCheckTemplateResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetDataCheckTemplateResponseBody) SetRequestId(v string) *GetDataCheckTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataCheckTemplateResponseBody) SetSuccess(v bool) *GetDataCheckTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataCheckTemplateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataCheckTemplateResponseBodyData struct {
	// The list of check rules for basic data type metrics. This field is required when checkType is set to 1 (metric comparison).
	BasicMetricRules []*GetDataCheckTemplateResponseBodyDataBasicMetricRules `json:"basicMetricRules,omitempty" xml:"basicMetricRules,omitempty" type:"Repeated"`
	// The check rule type. Valid values: 0: data volume comparison. 1: metric comparison. 2: weak content comparison. 3: custom comparison. 4: full-text comparison. 5: null rate comparison.
	//
	// example:
	//
	// 1
	CheckType *int32 `json:"checkType,omitempty" xml:"checkType,omitempty"`
	// The Chinese name of the check type (used in export report fields).
	//
	// example:
	//
	// 指标比对
	CheckTypeExport *string `json:"checkTypeExport,omitempty" xml:"checkTypeExport,omitempty"`
	// The check type name.
	//
	// example:
	//
	// 1
	CheckTypeName *int32 `json:"checkTypeName,omitempty" xml:"checkTypeName,omitempty"`
	// The list of check rules for composite data type metrics. This field is used when checkType is set to 1 (metric comparison).
	ComplexMetricRules []*GetDataCheckTemplateResponseBodyDataComplexMetricRules `json:"complexMetricRules,omitempty" xml:"complexMetricRules,omitempty" type:"Repeated"`
	// The list of data source engine relationships (data source engines associated with the template).
	DsEngineRels []*GetDataCheckTemplateResponseBodyDataDsEngineRels `json:"dsEngineRels,omitempty" xml:"dsEngineRels,omitempty" type:"Repeated"`
	// The full-text comparison rule. This field has a value when checkType is set to 4 (full-text comparison). For the field structure, refer to the child field descriptions below.
	FulltextRule *GetDataCheckTemplateResponseBodyDataFulltextRule `json:"fulltextRule,omitempty" xml:"fulltextRule,omitempty" type:"Struct"`
	// The list of metric check rules. This parameter has a value when checkType is set to 1 (metric comparison).
	MetricRules []*GetDataCheckTemplateResponseBodyDataMetricRules `json:"metricRules,omitempty" xml:"metricRules,omitempty" type:"Repeated"`
	// The list of null rate check rules. This parameter has a value when checkType is set to 5 (null rate comparison).
	NullRules []*GetDataCheckTemplateResponseBodyDataNullRules `json:"nullRules,omitempty" xml:"nullRules,omitempty" type:"Repeated"`
	// The template description.
	//
	// example:
	//
	// Description of the data volume check template
	TemplateDesc *string `json:"templateDesc,omitempty" xml:"templateDesc,omitempty"`
	// The check template ID (logical foreign key) that uniquely identifies a check template.
	//
	// example:
	//
	// 1001
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// The name of the check template.
	//
	// example:
	//
	// Data volume check template
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
	// The weak content check rule. This parameter has a value and is required when checkType is set to 2 (weak content comparison). For the field structure, see the child field descriptions.
	WeakContentRule *GetDataCheckTemplateResponseBodyDataWeakContentRule `json:"weakContentRule,omitempty" xml:"weakContentRule,omitempty" type:"Struct"`
}

func (s GetDataCheckTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDataCheckTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDataCheckTemplateResponseBodyData) GetBasicMetricRules() []*GetDataCheckTemplateResponseBodyDataBasicMetricRules {
	return s.BasicMetricRules
}

func (s *GetDataCheckTemplateResponseBodyData) GetCheckType() *int32 {
	return s.CheckType
}

func (s *GetDataCheckTemplateResponseBodyData) GetCheckTypeExport() *string {
	return s.CheckTypeExport
}

func (s *GetDataCheckTemplateResponseBodyData) GetCheckTypeName() *int32 {
	return s.CheckTypeName
}

func (s *GetDataCheckTemplateResponseBodyData) GetComplexMetricRules() []*GetDataCheckTemplateResponseBodyDataComplexMetricRules {
	return s.ComplexMetricRules
}

func (s *GetDataCheckTemplateResponseBodyData) GetDsEngineRels() []*GetDataCheckTemplateResponseBodyDataDsEngineRels {
	return s.DsEngineRels
}

func (s *GetDataCheckTemplateResponseBodyData) GetFulltextRule() *GetDataCheckTemplateResponseBodyDataFulltextRule {
	return s.FulltextRule
}

func (s *GetDataCheckTemplateResponseBodyData) GetMetricRules() []*GetDataCheckTemplateResponseBodyDataMetricRules {
	return s.MetricRules
}

func (s *GetDataCheckTemplateResponseBodyData) GetNullRules() []*GetDataCheckTemplateResponseBodyDataNullRules {
	return s.NullRules
}

func (s *GetDataCheckTemplateResponseBodyData) GetTemplateDesc() *string {
	return s.TemplateDesc
}

func (s *GetDataCheckTemplateResponseBodyData) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetDataCheckTemplateResponseBodyData) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetDataCheckTemplateResponseBodyData) GetWeakContentRule() *GetDataCheckTemplateResponseBodyDataWeakContentRule {
	return s.WeakContentRule
}

func (s *GetDataCheckTemplateResponseBodyData) SetBasicMetricRules(v []*GetDataCheckTemplateResponseBodyDataBasicMetricRules) *GetDataCheckTemplateResponseBodyData {
	s.BasicMetricRules = v
	return s
}

func (s *GetDataCheckTemplateResponseBodyData) SetCheckType(v int32) *GetDataCheckTemplateResponseBodyData {
	s.CheckType = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyData) SetCheckTypeExport(v string) *GetDataCheckTemplateResponseBodyData {
	s.CheckTypeExport = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyData) SetCheckTypeName(v int32) *GetDataCheckTemplateResponseBodyData {
	s.CheckTypeName = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyData) SetComplexMetricRules(v []*GetDataCheckTemplateResponseBodyDataComplexMetricRules) *GetDataCheckTemplateResponseBodyData {
	s.ComplexMetricRules = v
	return s
}

func (s *GetDataCheckTemplateResponseBodyData) SetDsEngineRels(v []*GetDataCheckTemplateResponseBodyDataDsEngineRels) *GetDataCheckTemplateResponseBodyData {
	s.DsEngineRels = v
	return s
}

func (s *GetDataCheckTemplateResponseBodyData) SetFulltextRule(v *GetDataCheckTemplateResponseBodyDataFulltextRule) *GetDataCheckTemplateResponseBodyData {
	s.FulltextRule = v
	return s
}

func (s *GetDataCheckTemplateResponseBodyData) SetMetricRules(v []*GetDataCheckTemplateResponseBodyDataMetricRules) *GetDataCheckTemplateResponseBodyData {
	s.MetricRules = v
	return s
}

func (s *GetDataCheckTemplateResponseBodyData) SetNullRules(v []*GetDataCheckTemplateResponseBodyDataNullRules) *GetDataCheckTemplateResponseBodyData {
	s.NullRules = v
	return s
}

func (s *GetDataCheckTemplateResponseBodyData) SetTemplateDesc(v string) *GetDataCheckTemplateResponseBodyData {
	s.TemplateDesc = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyData) SetTemplateId(v string) *GetDataCheckTemplateResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyData) SetTemplateName(v string) *GetDataCheckTemplateResponseBodyData {
	s.TemplateName = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyData) SetWeakContentRule(v *GetDataCheckTemplateResponseBodyDataWeakContentRule) *GetDataCheckTemplateResponseBodyData {
	s.WeakContentRule = v
	return s
}

func (s *GetDataCheckTemplateResponseBodyData) Validate() error {
	if s.BasicMetricRules != nil {
		for _, item := range s.BasicMetricRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ComplexMetricRules != nil {
		for _, item := range s.ComplexMetricRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DsEngineRels != nil {
		for _, item := range s.DsEngineRels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FulltextRule != nil {
		if err := s.FulltextRule.Validate(); err != nil {
			return err
		}
	}
	if s.MetricRules != nil {
		for _, item := range s.MetricRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NullRules != nil {
		for _, item := range s.NullRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.WeakContentRule != nil {
		if err := s.WeakContentRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataCheckTemplateResponseBodyDataBasicMetricRules struct {
	// The check methods (metric calculation methods). Multiple values are separated by commas, such as SUM,AVG,MIN,MAX. The values must be within the range allowed by the templatetype.
	//
	// example:
	//
	// SUM,AVG
	CheckMethods *string `json:"checkMethods,omitempty" xml:"checkMethods,omitempty"`
	// The data type category. Valid values: 0: primitive data type. 1: composite data type.
	//
	// example:
	//
	// 0
	DataTypeClassify *int32 `json:"dataTypeClassify,omitempty" xml:"dataTypeClassify,omitempty"`
	// The data type group that identifies the data type category to which the check rule applies. The value is an integer from 0 to 7. For the meaning of each value, refer to the valid values.
	//
	// example:
	//
	// 0
	DataTypeGroup *int32 `json:"dataTypeGroup,omitempty" xml:"dataTypeGroup,omitempty"`
	// The list of data types to which the check rule applies. Configure this field as needed.
	DataTypeList []*string `json:"dataTypeList,omitempty" xml:"dataTypeList,omitempty" type:"Repeated"`
	// The data types. Configure this field as needed.
	//
	// example:
	//
	// BIGINT
	DataTypes *string `json:"dataTypes,omitempty" xml:"dataTypes,omitempty"`
	// The difference tolerance rate type. Valid values: 0: unified. 1: custom. Default value: 0.
	//
	// example:
	//
	// 0
	DiffTolerateType *int32 `json:"diffTolerateType,omitempty" xml:"diffTolerateType,omitempty"`
	// The difference tolerance rate values. When the type is unified, one value is used. When the type is custom, values are set by the configured tolerance type, such as sum:33,avg:99.
	DiffTolerateValues map[string]interface{} `json:"diffTolerateValues,omitempty" xml:"diffTolerateValues,omitempty"`
	// Specifies whether to enable decimal scale control for DECIMAL type comparison. Valid values: 0: no. 1: yes.
	//
	// example:
	//
	// 0
	EnableDecimalScale *int32 `json:"enableDecimalScale,omitempty" xml:"enableDecimalScale,omitempty"`
	// The filter column names, separated by commas.
	//
	// example:
	//
	// col_a,col_b
	FilterColumnName *string `json:"filterColumnName,omitempty" xml:"filterColumnName,omitempty"`
	// Specifies whether to ignore trailing zeros in decimal places for DECIMAL type comparison. Valid values: 0: no. 1: yes.
	//
	// example:
	//
	// 0
	IgnoreDecimalScaleSuffixZero *int32 `json:"ignoreDecimalScaleSuffixZero,omitempty" xml:"ignoreDecimalScaleSuffixZero,omitempty"`
	// Specifies whether to ignore zero values for numeric types. Valid values: 0: no. 1: yes.
	//
	// example:
	//
	// 0
	IgnoreNumericZero *int32 `json:"ignoreNumericZero,omitempty" xml:"ignoreNumericZero,omitempty"`
	// Specifies whether to ignore empty strings and null for string types. Valid values: 0: no. 1: yes.
	//
	// example:
	//
	// 0
	IgnoreStringEmpty *int32 `json:"ignoreStringEmpty,omitempty" xml:"ignoreStringEmpty,omitempty"`
	// Specifies whether to enable count (data volume) check. Valid values: 0: no. 1: yes. Default value: 1.
	//
	// example:
	//
	// 1
	IsCountCheck *int32 `json:"isCountCheck,omitempty" xml:"isCountCheck,omitempty"`
	// The rule ID that uniquely identifies a check rule.
	//
	// example:
	//
	// 1001
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// The specific number of decimal places for DECIMAL type comparison.
	//
	// example:
	//
	// 2
	SetDecimalScale *int32 `json:"setDecimalScale,omitempty" xml:"setDecimalScale,omitempty"`
}

func (s GetDataCheckTemplateResponseBodyDataBasicMetricRules) String() string {
	return dara.Prettify(s)
}

func (s GetDataCheckTemplateResponseBodyDataBasicMetricRules) GoString() string {
	return s.String()
}

func (s *GetDataCheckTemplateResponseBodyDataBasicMetricRules) GetCheckMethods() *string {
	return s.CheckMethods
}

func (s *GetDataCheckTemplateResponseBodyDataBasicMetricRules) GetDataTypeClassify() *int32 {
	return s.DataTypeClassify
}

func (s *GetDataCheckTemplateResponseBodyDataBasicMetricRules) GetDataTypeGroup() *int32 {
	return s.DataTypeGroup
}

func (s *GetDataCheckTemplateResponseBodyDataBasicMetricRules) GetDataTypeList() []*string {
	return s.DataTypeList
}

func (s *GetDataCheckTemplateResponseBodyDataBasicMetricRules) GetDataTypes() *string {
	return s.DataTypes
}

func (s *GetDataCheckTemplateResponseBodyDataBasicMetricRules) GetDiffTolerateType() *int32 {
	return s.DiffTolerateType
}

func (s *GetDataCheckTemplateResponseBodyDataBasicMetricRules) GetDiffTolerateValues() map[string]interface{} {
	return s.DiffTolerateValues
}

func (s *GetDataCheckTemplateResponseBodyDataBasicMetricRules) GetEnableDecimalScale() *int32 {
	return s.EnableDecimalScale
}

func (s *GetDataCheckTemplateResponseBodyDataBasicMetricRules) GetFilterColumnName() *string {
	return s.FilterColumnName
}

func (s *GetDataCheckTemplateResponseBodyDataBasicMetricRules) GetIgnoreDecimalScaleSuffixZero() *int32 {
	return s.IgnoreDecimalScaleSuffixZero
}

func (s *GetDataCheckTemplateResponseBodyDataBasicMetricRules) GetIgnoreNumericZero() *int32 {
	return s.IgnoreNumericZero
}

func (s *GetDataCheckTemplateResponseBodyDataBasicMetricRules) GetIgnoreStringEmpty() *int32 {
	return s.IgnoreStringEmpty
}

func (s *GetDataCheckTemplateResponseBodyDataBasicMetricRules) GetIsCountCheck() *int32 {
	return s.IsCountCheck
}

func (s *GetDataCheckTemplateResponseBodyDataBasicMetricRules) GetRuleId() *string {
	return s.RuleId
}

func (s *GetDataCheckTemplateResponseBodyDataBasicMetricRules) GetSetDecimalScale() *int32 {
	return s.SetDecimalScale
}

func (s *GetDataCheckTemplateResponseBodyDataBasicMetricRules) SetCheckMethods(v string) *GetDataCheckTemplateResponseBodyDataBasicMetricRules {
	s.CheckMethods = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataBasicMetricRules) SetDataTypeClassify(v int32) *GetDataCheckTemplateResponseBodyDataBasicMetricRules {
	s.DataTypeClassify = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataBasicMetricRules) SetDataTypeGroup(v int32) *GetDataCheckTemplateResponseBodyDataBasicMetricRules {
	s.DataTypeGroup = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataBasicMetricRules) SetDataTypeList(v []*string) *GetDataCheckTemplateResponseBodyDataBasicMetricRules {
	s.DataTypeList = v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataBasicMetricRules) SetDataTypes(v string) *GetDataCheckTemplateResponseBodyDataBasicMetricRules {
	s.DataTypes = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataBasicMetricRules) SetDiffTolerateType(v int32) *GetDataCheckTemplateResponseBodyDataBasicMetricRules {
	s.DiffTolerateType = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataBasicMetricRules) SetDiffTolerateValues(v map[string]interface{}) *GetDataCheckTemplateResponseBodyDataBasicMetricRules {
	s.DiffTolerateValues = v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataBasicMetricRules) SetEnableDecimalScale(v int32) *GetDataCheckTemplateResponseBodyDataBasicMetricRules {
	s.EnableDecimalScale = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataBasicMetricRules) SetFilterColumnName(v string) *GetDataCheckTemplateResponseBodyDataBasicMetricRules {
	s.FilterColumnName = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataBasicMetricRules) SetIgnoreDecimalScaleSuffixZero(v int32) *GetDataCheckTemplateResponseBodyDataBasicMetricRules {
	s.IgnoreDecimalScaleSuffixZero = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataBasicMetricRules) SetIgnoreNumericZero(v int32) *GetDataCheckTemplateResponseBodyDataBasicMetricRules {
	s.IgnoreNumericZero = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataBasicMetricRules) SetIgnoreStringEmpty(v int32) *GetDataCheckTemplateResponseBodyDataBasicMetricRules {
	s.IgnoreStringEmpty = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataBasicMetricRules) SetIsCountCheck(v int32) *GetDataCheckTemplateResponseBodyDataBasicMetricRules {
	s.IsCountCheck = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataBasicMetricRules) SetRuleId(v string) *GetDataCheckTemplateResponseBodyDataBasicMetricRules {
	s.RuleId = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataBasicMetricRules) SetSetDecimalScale(v int32) *GetDataCheckTemplateResponseBodyDataBasicMetricRules {
	s.SetDecimalScale = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataBasicMetricRules) Validate() error {
	return dara.Validate(s)
}

type GetDataCheckTemplateResponseBodyDataComplexMetricRules struct {
	// The check methods (metric calculation methods). Multiple values are separated by commas, such as SUM,AVG,MIN,MAX. The values must be within the range allowed by the templatetype.
	//
	// example:
	//
	// SUM,AVG
	CheckMethods *string `json:"checkMethods,omitempty" xml:"checkMethods,omitempty"`
	// The data type category. Valid values: 0: primitive data type. 1: composite data type.
	//
	// example:
	//
	// 0
	DataTypeClassify *int32 `json:"dataTypeClassify,omitempty" xml:"dataTypeClassify,omitempty"`
	// The data type group that identifies the data type category to which the check rule applies. The value is an integer from 0 to 7. For the meaning of each value, refer to the valid values.
	//
	// example:
	//
	// 0
	DataTypeGroup *int32 `json:"dataTypeGroup,omitempty" xml:"dataTypeGroup,omitempty"`
	// The list of data types to which the check rule applies. Configure this field as needed.
	DataTypeList []*string `json:"dataTypeList,omitempty" xml:"dataTypeList,omitempty" type:"Repeated"`
	// The data types. Configure this field as needed.
	//
	// example:
	//
	// BIGINT
	DataTypes *string `json:"dataTypes,omitempty" xml:"dataTypes,omitempty"`
	// The difference tolerance rate type. Valid values: 0: unified. 1: custom. Default value: 0.
	//
	// example:
	//
	// 0
	DiffTolerateType *int32 `json:"diffTolerateType,omitempty" xml:"diffTolerateType,omitempty"`
	// The difference tolerance rate values. When the type is unified, one value is used. When the type is custom, values are set by the configured tolerance type, such as sum:33,avg:99.
	DiffTolerateValues map[string]interface{} `json:"diffTolerateValues,omitempty" xml:"diffTolerateValues,omitempty"`
	// Specifies whether to enable decimal scale control for DECIMAL type comparison. Valid values: 0: no. 1: yes.
	//
	// example:
	//
	// 0
	EnableDecimalScale *int32 `json:"enableDecimalScale,omitempty" xml:"enableDecimalScale,omitempty"`
	// The filter column names, separated by commas.
	//
	// example:
	//
	// col_a,col_b
	FilterColumnName *string `json:"filterColumnName,omitempty" xml:"filterColumnName,omitempty"`
	// Specifies whether to ignore trailing zeros in decimal places for DECIMAL type comparison. Valid values: 0: no. 1: yes.
	//
	// example:
	//
	// 0
	IgnoreDecimalScaleSuffixZero *int32 `json:"ignoreDecimalScaleSuffixZero,omitempty" xml:"ignoreDecimalScaleSuffixZero,omitempty"`
	// Specifies whether to ignore zero values for numeric types. Valid values: 0: no. 1: yes.
	//
	// example:
	//
	// 0
	IgnoreNumericZero *int32 `json:"ignoreNumericZero,omitempty" xml:"ignoreNumericZero,omitempty"`
	// Specifies whether to ignore empty strings and null for string types. Valid values: 0: no. 1: yes.
	//
	// example:
	//
	// 0
	IgnoreStringEmpty *int32 `json:"ignoreStringEmpty,omitempty" xml:"ignoreStringEmpty,omitempty"`
	// Specifies whether to enable count (data volume) check. Valid values: 0: no. 1: yes. Default value: 1.
	//
	// example:
	//
	// 1
	IsCountCheck *int32 `json:"isCountCheck,omitempty" xml:"isCountCheck,omitempty"`
	// The rule ID that uniquely identifies a check rule.
	//
	// example:
	//
	// 1001
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// The specific number of decimal places for DECIMAL type comparison.
	//
	// example:
	//
	// 2
	SetDecimalScale *int32 `json:"setDecimalScale,omitempty" xml:"setDecimalScale,omitempty"`
}

func (s GetDataCheckTemplateResponseBodyDataComplexMetricRules) String() string {
	return dara.Prettify(s)
}

func (s GetDataCheckTemplateResponseBodyDataComplexMetricRules) GoString() string {
	return s.String()
}

func (s *GetDataCheckTemplateResponseBodyDataComplexMetricRules) GetCheckMethods() *string {
	return s.CheckMethods
}

func (s *GetDataCheckTemplateResponseBodyDataComplexMetricRules) GetDataTypeClassify() *int32 {
	return s.DataTypeClassify
}

func (s *GetDataCheckTemplateResponseBodyDataComplexMetricRules) GetDataTypeGroup() *int32 {
	return s.DataTypeGroup
}

func (s *GetDataCheckTemplateResponseBodyDataComplexMetricRules) GetDataTypeList() []*string {
	return s.DataTypeList
}

func (s *GetDataCheckTemplateResponseBodyDataComplexMetricRules) GetDataTypes() *string {
	return s.DataTypes
}

func (s *GetDataCheckTemplateResponseBodyDataComplexMetricRules) GetDiffTolerateType() *int32 {
	return s.DiffTolerateType
}

func (s *GetDataCheckTemplateResponseBodyDataComplexMetricRules) GetDiffTolerateValues() map[string]interface{} {
	return s.DiffTolerateValues
}

func (s *GetDataCheckTemplateResponseBodyDataComplexMetricRules) GetEnableDecimalScale() *int32 {
	return s.EnableDecimalScale
}

func (s *GetDataCheckTemplateResponseBodyDataComplexMetricRules) GetFilterColumnName() *string {
	return s.FilterColumnName
}

func (s *GetDataCheckTemplateResponseBodyDataComplexMetricRules) GetIgnoreDecimalScaleSuffixZero() *int32 {
	return s.IgnoreDecimalScaleSuffixZero
}

func (s *GetDataCheckTemplateResponseBodyDataComplexMetricRules) GetIgnoreNumericZero() *int32 {
	return s.IgnoreNumericZero
}

func (s *GetDataCheckTemplateResponseBodyDataComplexMetricRules) GetIgnoreStringEmpty() *int32 {
	return s.IgnoreStringEmpty
}

func (s *GetDataCheckTemplateResponseBodyDataComplexMetricRules) GetIsCountCheck() *int32 {
	return s.IsCountCheck
}

func (s *GetDataCheckTemplateResponseBodyDataComplexMetricRules) GetRuleId() *string {
	return s.RuleId
}

func (s *GetDataCheckTemplateResponseBodyDataComplexMetricRules) GetSetDecimalScale() *int32 {
	return s.SetDecimalScale
}

func (s *GetDataCheckTemplateResponseBodyDataComplexMetricRules) SetCheckMethods(v string) *GetDataCheckTemplateResponseBodyDataComplexMetricRules {
	s.CheckMethods = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataComplexMetricRules) SetDataTypeClassify(v int32) *GetDataCheckTemplateResponseBodyDataComplexMetricRules {
	s.DataTypeClassify = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataComplexMetricRules) SetDataTypeGroup(v int32) *GetDataCheckTemplateResponseBodyDataComplexMetricRules {
	s.DataTypeGroup = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataComplexMetricRules) SetDataTypeList(v []*string) *GetDataCheckTemplateResponseBodyDataComplexMetricRules {
	s.DataTypeList = v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataComplexMetricRules) SetDataTypes(v string) *GetDataCheckTemplateResponseBodyDataComplexMetricRules {
	s.DataTypes = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataComplexMetricRules) SetDiffTolerateType(v int32) *GetDataCheckTemplateResponseBodyDataComplexMetricRules {
	s.DiffTolerateType = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataComplexMetricRules) SetDiffTolerateValues(v map[string]interface{}) *GetDataCheckTemplateResponseBodyDataComplexMetricRules {
	s.DiffTolerateValues = v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataComplexMetricRules) SetEnableDecimalScale(v int32) *GetDataCheckTemplateResponseBodyDataComplexMetricRules {
	s.EnableDecimalScale = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataComplexMetricRules) SetFilterColumnName(v string) *GetDataCheckTemplateResponseBodyDataComplexMetricRules {
	s.FilterColumnName = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataComplexMetricRules) SetIgnoreDecimalScaleSuffixZero(v int32) *GetDataCheckTemplateResponseBodyDataComplexMetricRules {
	s.IgnoreDecimalScaleSuffixZero = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataComplexMetricRules) SetIgnoreNumericZero(v int32) *GetDataCheckTemplateResponseBodyDataComplexMetricRules {
	s.IgnoreNumericZero = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataComplexMetricRules) SetIgnoreStringEmpty(v int32) *GetDataCheckTemplateResponseBodyDataComplexMetricRules {
	s.IgnoreStringEmpty = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataComplexMetricRules) SetIsCountCheck(v int32) *GetDataCheckTemplateResponseBodyDataComplexMetricRules {
	s.IsCountCheck = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataComplexMetricRules) SetRuleId(v string) *GetDataCheckTemplateResponseBodyDataComplexMetricRules {
	s.RuleId = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataComplexMetricRules) SetSetDecimalScale(v int32) *GetDataCheckTemplateResponseBodyDataComplexMetricRules {
	s.SetDecimalScale = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataComplexMetricRules) Validate() error {
	return dara.Validate(s)
}

type GetDataCheckTemplateResponseBodyDataDsEngineRels struct {
	// The data source type, such as Hive or MaxCompute.
	//
	// example:
	//
	// Hive
	DsType *string `json:"dsType,omitempty" xml:"dsType,omitempty"`
	// The list of covered check engine types, such as Tez or MapReduce. When in string format, multiple values are separated by commas.
	EngineTypes []*string `json:"engineTypes,omitempty" xml:"engineTypes,omitempty" type:"Repeated"`
}

func (s GetDataCheckTemplateResponseBodyDataDsEngineRels) String() string {
	return dara.Prettify(s)
}

func (s GetDataCheckTemplateResponseBodyDataDsEngineRels) GoString() string {
	return s.String()
}

func (s *GetDataCheckTemplateResponseBodyDataDsEngineRels) GetDsType() *string {
	return s.DsType
}

func (s *GetDataCheckTemplateResponseBodyDataDsEngineRels) GetEngineTypes() []*string {
	return s.EngineTypes
}

func (s *GetDataCheckTemplateResponseBodyDataDsEngineRels) SetDsType(v string) *GetDataCheckTemplateResponseBodyDataDsEngineRels {
	s.DsType = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataDsEngineRels) SetEngineTypes(v []*string) *GetDataCheckTemplateResponseBodyDataDsEngineRels {
	s.EngineTypes = v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataDsEngineRels) Validate() error {
	return dara.Validate(s)
}

type GetDataCheckTemplateResponseBodyDataFulltextRule struct {
	// The check mode. Valid values: 0: row-level overall comparison. 1: row-level column-by-column comparison. 2: both row-level overall comparison and row-level column-by-column comparison.
	//
	// example:
	//
	// 0
	CheckMode *int32 `json:"checkMode,omitempty" xml:"checkMode,omitempty"`
	// The equality comparison type for row-level column-by-column comparison. Valid values: 0: all field types. 1: primitive basic data types. 2: composite data types. 3: custom.
	//
	// example:
	//
	// 0
	ColumnEqualCmpType *int32 `json:"columnEqualCmpType,omitempty" xml:"columnEqualCmpType,omitempty"`
	// The custom type list for equality comparison during row-level column-by-column comparison. Multiple values are separated by commas.
	//
	// example:
	//
	// ARRAY,MAP
	ColumnEqualCmpValues *string `json:"columnEqualCmpValues,omitempty" xml:"columnEqualCmpValues,omitempty"`
	// Specifies whether to enable cosine similarity for row-level column-by-column comparison. Valid values: 0: no. 1: yes.
	//
	// example:
	//
	// 0
	ColumnIsCosine *int32 `json:"columnIsCosine,omitempty" xml:"columnIsCosine,omitempty"`
	// Specifies whether to ignore differences between null values and empty strings during row-by-row and column-by-column comparison. Valid values: 0: No. 1: Yes.
	//
	// example:
	//
	// 0
	ColumnIsIgnoreNull *int32 `json:"columnIsIgnoreNull,omitempty" xml:"columnIsIgnoreNull,omitempty"`
	// Specifies whether to ignore differences between null values and 0 values during row-by-row and column-by-column comparison. Valid values: 0: No. 1: Yes.
	//
	// example:
	//
	// 0
	ColumnIsIgnoreZero *int32 `json:"columnIsIgnoreZero,omitempty" xml:"columnIsIgnoreZero,omitempty"`
	// Specifies whether to enable sampling during row-by-row and column-by-column comparison. Valid values: 0: No. 1: Yes.
	//
	// example:
	//
	// 0
	ColumnIsSamples *int32 `json:"columnIsSamples,omitempty" xml:"columnIsSamples,omitempty"`
	// The sampling method during row-by-row and column-by-column comparison. Valid values: 0: by row. 1: by percentage.
	//
	// example:
	//
	// 0
	ColumnSamplesType *int32 `json:"columnSamplesType,omitempty" xml:"columnSamplesType,omitempty"`
	// The sampling value during row-by-row and column-by-column comparison. The meaning depends on the sampling method. When sampling by row, this value represents the number of rows. When sampling by percentage, this value represents the percentage.
	//
	// example:
	//
	// 100
	ColumnSamplesValue *int32 `json:"columnSamplesValue,omitempty" xml:"columnSamplesValue,omitempty"`
	// The size comparison type during row-by-row and column-by-column comparison. Valid values: 0: all composite data types. 1: custom.
	//
	// example:
	//
	// 0
	ColumnSizeCmpType *int32 `json:"columnSizeCmpType,omitempty" xml:"columnSizeCmpType,omitempty"`
	// The custom type list for size comparison during row-by-row and column-by-column comparison. Multiple values are separated by commas (,).
	//
	// example:
	//
	// ARRAY,MAP
	ColumnSizeCmpValues *string `json:"columnSizeCmpValues,omitempty" xml:"columnSizeCmpValues,omitempty"`
	// Specifies whether to enable the existence check for primary keys or composite primary keys. Valid values: 0: No. 1: Yes.
	//
	// example:
	//
	// 1
	IsPrimaryKeyCheck *int32 `json:"isPrimaryKeyCheck,omitempty" xml:"isPrimaryKeyCheck,omitempty"`
	// The row-by-row comparison method. Valid values: 0: md5. 1: crc32.
	//
	// example:
	//
	// 0
	LineCheckType *int32 `json:"lineCheckType,omitempty" xml:"lineCheckType,omitempty"`
	// Specifies whether to print all columns in the difference details during row-by-row comparison. Valid values: 0: No. 1: Yes.
	//
	// example:
	//
	// 0
	LineIsPrintAll *int32 `json:"lineIsPrintAll,omitempty" xml:"lineIsPrintAll,omitempty"`
	// Specifies whether to enable sampling during row-by-row comparison. Valid values: 0: No. 1: Yes.
	//
	// example:
	//
	// 0
	LineIsSamples *int32 `json:"lineIsSamples,omitempty" xml:"lineIsSamples,omitempty"`
	// The sampling method during row-by-row comparison. Valid values: 0: by row. 1: by percentage.
	//
	// example:
	//
	// 0
	LineSamplesType *int32 `json:"lineSamplesType,omitempty" xml:"lineSamplesType,omitempty"`
	// The sampling value during row-by-row comparison. The meaning depends on the sampling method. When sampling by row, this value represents the number of rows. When sampling by percentage, this value represents the percentage.
	//
	// example:
	//
	// 100
	LineSamplesValue *int32 `json:"lineSamplesValue,omitempty" xml:"lineSamplesValue,omitempty"`
	// The rule ID that uniquely identifies a check rule.
	//
	// example:
	//
	// 1001
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
}

func (s GetDataCheckTemplateResponseBodyDataFulltextRule) String() string {
	return dara.Prettify(s)
}

func (s GetDataCheckTemplateResponseBodyDataFulltextRule) GoString() string {
	return s.String()
}

func (s *GetDataCheckTemplateResponseBodyDataFulltextRule) GetCheckMode() *int32 {
	return s.CheckMode
}

func (s *GetDataCheckTemplateResponseBodyDataFulltextRule) GetColumnEqualCmpType() *int32 {
	return s.ColumnEqualCmpType
}

func (s *GetDataCheckTemplateResponseBodyDataFulltextRule) GetColumnEqualCmpValues() *string {
	return s.ColumnEqualCmpValues
}

func (s *GetDataCheckTemplateResponseBodyDataFulltextRule) GetColumnIsCosine() *int32 {
	return s.ColumnIsCosine
}

func (s *GetDataCheckTemplateResponseBodyDataFulltextRule) GetColumnIsIgnoreNull() *int32 {
	return s.ColumnIsIgnoreNull
}

func (s *GetDataCheckTemplateResponseBodyDataFulltextRule) GetColumnIsIgnoreZero() *int32 {
	return s.ColumnIsIgnoreZero
}

func (s *GetDataCheckTemplateResponseBodyDataFulltextRule) GetColumnIsSamples() *int32 {
	return s.ColumnIsSamples
}

func (s *GetDataCheckTemplateResponseBodyDataFulltextRule) GetColumnSamplesType() *int32 {
	return s.ColumnSamplesType
}

func (s *GetDataCheckTemplateResponseBodyDataFulltextRule) GetColumnSamplesValue() *int32 {
	return s.ColumnSamplesValue
}

func (s *GetDataCheckTemplateResponseBodyDataFulltextRule) GetColumnSizeCmpType() *int32 {
	return s.ColumnSizeCmpType
}

func (s *GetDataCheckTemplateResponseBodyDataFulltextRule) GetColumnSizeCmpValues() *string {
	return s.ColumnSizeCmpValues
}

func (s *GetDataCheckTemplateResponseBodyDataFulltextRule) GetIsPrimaryKeyCheck() *int32 {
	return s.IsPrimaryKeyCheck
}

func (s *GetDataCheckTemplateResponseBodyDataFulltextRule) GetLineCheckType() *int32 {
	return s.LineCheckType
}

func (s *GetDataCheckTemplateResponseBodyDataFulltextRule) GetLineIsPrintAll() *int32 {
	return s.LineIsPrintAll
}

func (s *GetDataCheckTemplateResponseBodyDataFulltextRule) GetLineIsSamples() *int32 {
	return s.LineIsSamples
}

func (s *GetDataCheckTemplateResponseBodyDataFulltextRule) GetLineSamplesType() *int32 {
	return s.LineSamplesType
}

func (s *GetDataCheckTemplateResponseBodyDataFulltextRule) GetLineSamplesValue() *int32 {
	return s.LineSamplesValue
}

func (s *GetDataCheckTemplateResponseBodyDataFulltextRule) GetRuleId() *string {
	return s.RuleId
}

func (s *GetDataCheckTemplateResponseBodyDataFulltextRule) SetCheckMode(v int32) *GetDataCheckTemplateResponseBodyDataFulltextRule {
	s.CheckMode = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataFulltextRule) SetColumnEqualCmpType(v int32) *GetDataCheckTemplateResponseBodyDataFulltextRule {
	s.ColumnEqualCmpType = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataFulltextRule) SetColumnEqualCmpValues(v string) *GetDataCheckTemplateResponseBodyDataFulltextRule {
	s.ColumnEqualCmpValues = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataFulltextRule) SetColumnIsCosine(v int32) *GetDataCheckTemplateResponseBodyDataFulltextRule {
	s.ColumnIsCosine = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataFulltextRule) SetColumnIsIgnoreNull(v int32) *GetDataCheckTemplateResponseBodyDataFulltextRule {
	s.ColumnIsIgnoreNull = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataFulltextRule) SetColumnIsIgnoreZero(v int32) *GetDataCheckTemplateResponseBodyDataFulltextRule {
	s.ColumnIsIgnoreZero = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataFulltextRule) SetColumnIsSamples(v int32) *GetDataCheckTemplateResponseBodyDataFulltextRule {
	s.ColumnIsSamples = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataFulltextRule) SetColumnSamplesType(v int32) *GetDataCheckTemplateResponseBodyDataFulltextRule {
	s.ColumnSamplesType = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataFulltextRule) SetColumnSamplesValue(v int32) *GetDataCheckTemplateResponseBodyDataFulltextRule {
	s.ColumnSamplesValue = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataFulltextRule) SetColumnSizeCmpType(v int32) *GetDataCheckTemplateResponseBodyDataFulltextRule {
	s.ColumnSizeCmpType = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataFulltextRule) SetColumnSizeCmpValues(v string) *GetDataCheckTemplateResponseBodyDataFulltextRule {
	s.ColumnSizeCmpValues = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataFulltextRule) SetIsPrimaryKeyCheck(v int32) *GetDataCheckTemplateResponseBodyDataFulltextRule {
	s.IsPrimaryKeyCheck = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataFulltextRule) SetLineCheckType(v int32) *GetDataCheckTemplateResponseBodyDataFulltextRule {
	s.LineCheckType = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataFulltextRule) SetLineIsPrintAll(v int32) *GetDataCheckTemplateResponseBodyDataFulltextRule {
	s.LineIsPrintAll = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataFulltextRule) SetLineIsSamples(v int32) *GetDataCheckTemplateResponseBodyDataFulltextRule {
	s.LineIsSamples = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataFulltextRule) SetLineSamplesType(v int32) *GetDataCheckTemplateResponseBodyDataFulltextRule {
	s.LineSamplesType = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataFulltextRule) SetLineSamplesValue(v int32) *GetDataCheckTemplateResponseBodyDataFulltextRule {
	s.LineSamplesValue = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataFulltextRule) SetRuleId(v string) *GetDataCheckTemplateResponseBodyDataFulltextRule {
	s.RuleId = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataFulltextRule) Validate() error {
	return dara.Validate(s)
}

type GetDataCheckTemplateResponseBodyDataMetricRules struct {
	// The check methods (metric calculation methods). Multiple values are separated by commas (,), such as SUM,AVG,MIN,MAX. The values must be within the range allowed by the templatetype.
	//
	// example:
	//
	// SUM,AVG
	CheckMethods *string `json:"checkMethods,omitempty" xml:"checkMethods,omitempty"`
	// The data type category. Valid values: 0: primitive data type. 1: composite data type.
	//
	// example:
	//
	// 0
	DataTypeClassify *int32 `json:"dataTypeClassify,omitempty" xml:"dataTypeClassify,omitempty"`
	// The data type group that identifies the data type category to which the check rule applies. The value is an integer from 0 to 7. For the meaning of each value, refer to the valid values.
	//
	// example:
	//
	// 0
	DataTypeGroup *int32 `json:"dataTypeGroup,omitempty" xml:"dataTypeGroup,omitempty"`
	// The list of data types to which the check rule applies. Configure this field as needed.
	DataTypeList []*string `json:"dataTypeList,omitempty" xml:"dataTypeList,omitempty" type:"Repeated"`
	// The data types. Configure this field as needed.
	//
	// example:
	//
	// BIGINT
	DataTypes *string `json:"dataTypes,omitempty" xml:"dataTypes,omitempty"`
	// The difference tolerance rate type. Valid values: 0: unified. 1: custom. Default value: 0.
	//
	// example:
	//
	// 0
	DiffTolerateType *int32 `json:"diffTolerateType,omitempty" xml:"diffTolerateType,omitempty"`
	// The difference tolerance rate values. When the type is unified, one value is used. When the type is custom, values are set by the configured tolerance type, such as sum:33,avg:99.
	DiffTolerateValues map[string]interface{} `json:"diffTolerateValues,omitempty" xml:"diffTolerateValues,omitempty"`
	// Specifies whether to enable decimal scale control for DECIMAL type comparison. Valid values: 0: no. 1: yes.
	//
	// example:
	//
	// 0
	EnableDecimalScale *int32 `json:"enableDecimalScale,omitempty" xml:"enableDecimalScale,omitempty"`
	// The filter column names, separated by commas.
	//
	// example:
	//
	// col_a,col_b
	FilterColumnName *string `json:"filterColumnName,omitempty" xml:"filterColumnName,omitempty"`
	// Specifies whether to ignore trailing zeros in decimal places for DECIMAL type comparison. Valid values: 0: no. 1: yes.
	//
	// example:
	//
	// 0
	IgnoreDecimalScaleSuffixZero *int32 `json:"ignoreDecimalScaleSuffixZero,omitempty" xml:"ignoreDecimalScaleSuffixZero,omitempty"`
	// Specifies whether to ignore zero values for numeric types. Valid values: 0: no. 1: yes.
	//
	// example:
	//
	// 0
	IgnoreNumericZero *int32 `json:"ignoreNumericZero,omitempty" xml:"ignoreNumericZero,omitempty"`
	// Specifies whether to ignore empty strings and null for string types. Valid values: 0: no. 1: yes.
	//
	// example:
	//
	// 0
	IgnoreStringEmpty *int32 `json:"ignoreStringEmpty,omitempty" xml:"ignoreStringEmpty,omitempty"`
	// Specifies whether to enable count (data volume) check. Valid values: 0: no. 1: yes. Default value: 1.
	//
	// example:
	//
	// 1
	IsCountCheck *int32 `json:"isCountCheck,omitempty" xml:"isCountCheck,omitempty"`
	// The rule ID that uniquely identifies a check rule.
	//
	// example:
	//
	// 1001
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// The specific number of decimal places for DECIMAL type comparison.
	//
	// example:
	//
	// 2
	SetDecimalScale *int32 `json:"setDecimalScale,omitempty" xml:"setDecimalScale,omitempty"`
}

func (s GetDataCheckTemplateResponseBodyDataMetricRules) String() string {
	return dara.Prettify(s)
}

func (s GetDataCheckTemplateResponseBodyDataMetricRules) GoString() string {
	return s.String()
}

func (s *GetDataCheckTemplateResponseBodyDataMetricRules) GetCheckMethods() *string {
	return s.CheckMethods
}

func (s *GetDataCheckTemplateResponseBodyDataMetricRules) GetDataTypeClassify() *int32 {
	return s.DataTypeClassify
}

func (s *GetDataCheckTemplateResponseBodyDataMetricRules) GetDataTypeGroup() *int32 {
	return s.DataTypeGroup
}

func (s *GetDataCheckTemplateResponseBodyDataMetricRules) GetDataTypeList() []*string {
	return s.DataTypeList
}

func (s *GetDataCheckTemplateResponseBodyDataMetricRules) GetDataTypes() *string {
	return s.DataTypes
}

func (s *GetDataCheckTemplateResponseBodyDataMetricRules) GetDiffTolerateType() *int32 {
	return s.DiffTolerateType
}

func (s *GetDataCheckTemplateResponseBodyDataMetricRules) GetDiffTolerateValues() map[string]interface{} {
	return s.DiffTolerateValues
}

func (s *GetDataCheckTemplateResponseBodyDataMetricRules) GetEnableDecimalScale() *int32 {
	return s.EnableDecimalScale
}

func (s *GetDataCheckTemplateResponseBodyDataMetricRules) GetFilterColumnName() *string {
	return s.FilterColumnName
}

func (s *GetDataCheckTemplateResponseBodyDataMetricRules) GetIgnoreDecimalScaleSuffixZero() *int32 {
	return s.IgnoreDecimalScaleSuffixZero
}

func (s *GetDataCheckTemplateResponseBodyDataMetricRules) GetIgnoreNumericZero() *int32 {
	return s.IgnoreNumericZero
}

func (s *GetDataCheckTemplateResponseBodyDataMetricRules) GetIgnoreStringEmpty() *int32 {
	return s.IgnoreStringEmpty
}

func (s *GetDataCheckTemplateResponseBodyDataMetricRules) GetIsCountCheck() *int32 {
	return s.IsCountCheck
}

func (s *GetDataCheckTemplateResponseBodyDataMetricRules) GetRuleId() *string {
	return s.RuleId
}

func (s *GetDataCheckTemplateResponseBodyDataMetricRules) GetSetDecimalScale() *int32 {
	return s.SetDecimalScale
}

func (s *GetDataCheckTemplateResponseBodyDataMetricRules) SetCheckMethods(v string) *GetDataCheckTemplateResponseBodyDataMetricRules {
	s.CheckMethods = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataMetricRules) SetDataTypeClassify(v int32) *GetDataCheckTemplateResponseBodyDataMetricRules {
	s.DataTypeClassify = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataMetricRules) SetDataTypeGroup(v int32) *GetDataCheckTemplateResponseBodyDataMetricRules {
	s.DataTypeGroup = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataMetricRules) SetDataTypeList(v []*string) *GetDataCheckTemplateResponseBodyDataMetricRules {
	s.DataTypeList = v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataMetricRules) SetDataTypes(v string) *GetDataCheckTemplateResponseBodyDataMetricRules {
	s.DataTypes = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataMetricRules) SetDiffTolerateType(v int32) *GetDataCheckTemplateResponseBodyDataMetricRules {
	s.DiffTolerateType = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataMetricRules) SetDiffTolerateValues(v map[string]interface{}) *GetDataCheckTemplateResponseBodyDataMetricRules {
	s.DiffTolerateValues = v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataMetricRules) SetEnableDecimalScale(v int32) *GetDataCheckTemplateResponseBodyDataMetricRules {
	s.EnableDecimalScale = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataMetricRules) SetFilterColumnName(v string) *GetDataCheckTemplateResponseBodyDataMetricRules {
	s.FilterColumnName = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataMetricRules) SetIgnoreDecimalScaleSuffixZero(v int32) *GetDataCheckTemplateResponseBodyDataMetricRules {
	s.IgnoreDecimalScaleSuffixZero = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataMetricRules) SetIgnoreNumericZero(v int32) *GetDataCheckTemplateResponseBodyDataMetricRules {
	s.IgnoreNumericZero = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataMetricRules) SetIgnoreStringEmpty(v int32) *GetDataCheckTemplateResponseBodyDataMetricRules {
	s.IgnoreStringEmpty = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataMetricRules) SetIsCountCheck(v int32) *GetDataCheckTemplateResponseBodyDataMetricRules {
	s.IsCountCheck = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataMetricRules) SetRuleId(v string) *GetDataCheckTemplateResponseBodyDataMetricRules {
	s.RuleId = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataMetricRules) SetSetDecimalScale(v int32) *GetDataCheckTemplateResponseBodyDataMetricRules {
	s.SetDecimalScale = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataMetricRules) Validate() error {
	return dara.Validate(s)
}

type GetDataCheckTemplateResponseBodyDataNullRules struct {
	// The data type group that identifies the data type category to which the check rule applies. The value is an integer from 0 to 7. For the meaning of each value, refer to the valid values.
	//
	// example:
	//
	// 0
	DataTypeGroup *int32 `json:"dataTypeGroup,omitempty" xml:"dataTypeGroup,omitempty"`
	// The null values, stored in JSON format.
	//
	// example:
	//
	// {}
	NullValues *string `json:"nullValues,omitempty" xml:"nullValues,omitempty"`
	// The rule ID that uniquely identifies a check rule.
	//
	// example:
	//
	// 1001
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
}

func (s GetDataCheckTemplateResponseBodyDataNullRules) String() string {
	return dara.Prettify(s)
}

func (s GetDataCheckTemplateResponseBodyDataNullRules) GoString() string {
	return s.String()
}

func (s *GetDataCheckTemplateResponseBodyDataNullRules) GetDataTypeGroup() *int32 {
	return s.DataTypeGroup
}

func (s *GetDataCheckTemplateResponseBodyDataNullRules) GetNullValues() *string {
	return s.NullValues
}

func (s *GetDataCheckTemplateResponseBodyDataNullRules) GetRuleId() *string {
	return s.RuleId
}

func (s *GetDataCheckTemplateResponseBodyDataNullRules) SetDataTypeGroup(v int32) *GetDataCheckTemplateResponseBodyDataNullRules {
	s.DataTypeGroup = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataNullRules) SetNullValues(v string) *GetDataCheckTemplateResponseBodyDataNullRules {
	s.NullValues = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataNullRules) SetRuleId(v string) *GetDataCheckTemplateResponseBodyDataNullRules {
	s.RuleId = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataNullRules) Validate() error {
	return dara.Validate(s)
}

type GetDataCheckTemplateResponseBodyDataWeakContentRule struct {
	// The filter column name expression.
	//
	// example:
	//
	// ^col_.*$
	FilterColumnExpression *string `json:"filterColumnExpression,omitempty" xml:"filterColumnExpression,omitempty"`
	// The filter column types, separated by vertical bars (|).
	FilterColumnTypes []*string `json:"filterColumnTypes,omitempty" xml:"filterColumnTypes,omitempty" type:"Repeated"`
	// The rule ID that uniquely identifies a check rule.
	//
	// example:
	//
	// 1001
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// The weak content algorithm name. Valid values: md5 and crc32.
	//
	// example:
	//
	// md5
	WeakContentAlgorithm *string `json:"weakContentAlgorithm,omitempty" xml:"weakContentAlgorithm,omitempty"`
}

func (s GetDataCheckTemplateResponseBodyDataWeakContentRule) String() string {
	return dara.Prettify(s)
}

func (s GetDataCheckTemplateResponseBodyDataWeakContentRule) GoString() string {
	return s.String()
}

func (s *GetDataCheckTemplateResponseBodyDataWeakContentRule) GetFilterColumnExpression() *string {
	return s.FilterColumnExpression
}

func (s *GetDataCheckTemplateResponseBodyDataWeakContentRule) GetFilterColumnTypes() []*string {
	return s.FilterColumnTypes
}

func (s *GetDataCheckTemplateResponseBodyDataWeakContentRule) GetRuleId() *string {
	return s.RuleId
}

func (s *GetDataCheckTemplateResponseBodyDataWeakContentRule) GetWeakContentAlgorithm() *string {
	return s.WeakContentAlgorithm
}

func (s *GetDataCheckTemplateResponseBodyDataWeakContentRule) SetFilterColumnExpression(v string) *GetDataCheckTemplateResponseBodyDataWeakContentRule {
	s.FilterColumnExpression = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataWeakContentRule) SetFilterColumnTypes(v []*string) *GetDataCheckTemplateResponseBodyDataWeakContentRule {
	s.FilterColumnTypes = v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataWeakContentRule) SetRuleId(v string) *GetDataCheckTemplateResponseBodyDataWeakContentRule {
	s.RuleId = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataWeakContentRule) SetWeakContentAlgorithm(v string) *GetDataCheckTemplateResponseBodyDataWeakContentRule {
	s.WeakContentAlgorithm = &v
	return s
}

func (s *GetDataCheckTemplateResponseBodyDataWeakContentRule) Validate() error {
	return dara.Validate(s)
}
