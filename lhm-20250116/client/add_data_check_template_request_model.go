// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDataCheckTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBasicMetricRules(v []*AddDataCheckTemplateRequestBasicMetricRules) *AddDataCheckTemplateRequest
	GetBasicMetricRules() []*AddDataCheckTemplateRequestBasicMetricRules
	SetCheckType(v int32) *AddDataCheckTemplateRequest
	GetCheckType() *int32
	SetComplexMetricRules(v []*AddDataCheckTemplateRequestComplexMetricRules) *AddDataCheckTemplateRequest
	GetComplexMetricRules() []*AddDataCheckTemplateRequestComplexMetricRules
	SetDsEngineRels(v []*AddDataCheckTemplateRequestDsEngineRels) *AddDataCheckTemplateRequest
	GetDsEngineRels() []*AddDataCheckTemplateRequestDsEngineRels
	SetFulltextRule(v *AddDataCheckTemplateRequestFulltextRule) *AddDataCheckTemplateRequest
	GetFulltextRule() *AddDataCheckTemplateRequestFulltextRule
	SetMetricRules(v []*AddDataCheckTemplateRequestMetricRules) *AddDataCheckTemplateRequest
	GetMetricRules() []*AddDataCheckTemplateRequestMetricRules
	SetNullRules(v []*AddDataCheckTemplateRequestNullRules) *AddDataCheckTemplateRequest
	GetNullRules() []*AddDataCheckTemplateRequestNullRules
	SetRequestId(v string) *AddDataCheckTemplateRequest
	GetRequestId() *string
	SetTemplateDesc(v string) *AddDataCheckTemplateRequest
	GetTemplateDesc() *string
	SetTemplateName(v string) *AddDataCheckTemplateRequest
	GetTemplateName() *string
	SetTenantId(v string) *AddDataCheckTemplateRequest
	GetTenantId() *string
	SetWeakContentRule(v *AddDataCheckTemplateRequestWeakContentRule) *AddDataCheckTemplateRequest
	GetWeakContentRule() *AddDataCheckTemplateRequestWeakContentRule
}

type AddDataCheckTemplateRequest struct {
	// The list of metric verification rules for basic data types. This field is required when checkType is set to 1 (metric comparison).
	BasicMetricRules []*AddDataCheckTemplateRequestBasicMetricRules `json:"basicMetricRules,omitempty" xml:"basicMetricRules,omitempty" type:"Repeated"`
	// The verification rule type. Valid values:
	//
	// - 0: data volume comparison.
	//
	// - 1: metric comparison.
	//
	// - 2: weak content comparison.
	//
	// - 3: custom comparison.
	//
	// - 4: full-text comparison.
	//
	// - 5: null rate comparison.
	//
	// example:
	//
	// 1
	CheckType *int32 `json:"checkType,omitempty" xml:"checkType,omitempty"`
	// The list of check rules for complex data type metrics. Used when checkType is set to 1 (metric comparison).
	ComplexMetricRules []*AddDataCheckTemplateRequestComplexMetricRules `json:"complexMetricRules,omitempty" xml:"complexMetricRules,omitempty" type:"Repeated"`
	// The list of datasource engine relationships (datasource engines associated with the template).
	DsEngineRels []*AddDataCheckTemplateRequestDsEngineRels `json:"dsEngineRels,omitempty" xml:"dsEngineRels,omitempty" type:"Repeated"`
	// The full-text comparison rule. This parameter has a value when checkType is set to 4 (full-text comparison). For the field structure, see the child field descriptions.
	FulltextRule *AddDataCheckTemplateRequestFulltextRule `json:"fulltextRule,omitempty" xml:"fulltextRule,omitempty" type:"Struct"`
	// The list of metric check rules. This parameter has a value when checkType is set to 1 (metric comparison).
	MetricRules []*AddDataCheckTemplateRequestMetricRules `json:"metricRules,omitempty" xml:"metricRules,omitempty" type:"Repeated"`
	// The list of null value rate check rules. This parameter has a value when checkType is set to 5 (null value rate comparison).
	NullRules []*AddDataCheckTemplateRequestNullRules `json:"nullRules,omitempty" xml:"nullRules,omitempty" type:"Repeated"`
	// The request ID, which is used to locate and troubleshoot issues of this call.
	//
	// example:
	//
	// 4C467B38-3910-4477-9B0B-6963D83B4E72
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The template description.
	//
	// example:
	//
	// Description of the data volume check template
	TemplateDesc *string `json:"templateDesc,omitempty" xml:"templateDesc,omitempty"`
	// The name of the check template.
	//
	// example:
	//
	// Data volume check template
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 10001
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The weak content check rule. This parameter has a value and is required when checkType is set to 2 (weak content comparison). For the field structure, refer to the child field descriptions below.
	WeakContentRule *AddDataCheckTemplateRequestWeakContentRule `json:"weakContentRule,omitempty" xml:"weakContentRule,omitempty" type:"Struct"`
}

func (s AddDataCheckTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDataCheckTemplateRequest) GoString() string {
	return s.String()
}

func (s *AddDataCheckTemplateRequest) GetBasicMetricRules() []*AddDataCheckTemplateRequestBasicMetricRules {
	return s.BasicMetricRules
}

func (s *AddDataCheckTemplateRequest) GetCheckType() *int32 {
	return s.CheckType
}

func (s *AddDataCheckTemplateRequest) GetComplexMetricRules() []*AddDataCheckTemplateRequestComplexMetricRules {
	return s.ComplexMetricRules
}

func (s *AddDataCheckTemplateRequest) GetDsEngineRels() []*AddDataCheckTemplateRequestDsEngineRels {
	return s.DsEngineRels
}

func (s *AddDataCheckTemplateRequest) GetFulltextRule() *AddDataCheckTemplateRequestFulltextRule {
	return s.FulltextRule
}

func (s *AddDataCheckTemplateRequest) GetMetricRules() []*AddDataCheckTemplateRequestMetricRules {
	return s.MetricRules
}

func (s *AddDataCheckTemplateRequest) GetNullRules() []*AddDataCheckTemplateRequestNullRules {
	return s.NullRules
}

func (s *AddDataCheckTemplateRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDataCheckTemplateRequest) GetTemplateDesc() *string {
	return s.TemplateDesc
}

func (s *AddDataCheckTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *AddDataCheckTemplateRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *AddDataCheckTemplateRequest) GetWeakContentRule() *AddDataCheckTemplateRequestWeakContentRule {
	return s.WeakContentRule
}

func (s *AddDataCheckTemplateRequest) SetBasicMetricRules(v []*AddDataCheckTemplateRequestBasicMetricRules) *AddDataCheckTemplateRequest {
	s.BasicMetricRules = v
	return s
}

func (s *AddDataCheckTemplateRequest) SetCheckType(v int32) *AddDataCheckTemplateRequest {
	s.CheckType = &v
	return s
}

func (s *AddDataCheckTemplateRequest) SetComplexMetricRules(v []*AddDataCheckTemplateRequestComplexMetricRules) *AddDataCheckTemplateRequest {
	s.ComplexMetricRules = v
	return s
}

func (s *AddDataCheckTemplateRequest) SetDsEngineRels(v []*AddDataCheckTemplateRequestDsEngineRels) *AddDataCheckTemplateRequest {
	s.DsEngineRels = v
	return s
}

func (s *AddDataCheckTemplateRequest) SetFulltextRule(v *AddDataCheckTemplateRequestFulltextRule) *AddDataCheckTemplateRequest {
	s.FulltextRule = v
	return s
}

func (s *AddDataCheckTemplateRequest) SetMetricRules(v []*AddDataCheckTemplateRequestMetricRules) *AddDataCheckTemplateRequest {
	s.MetricRules = v
	return s
}

func (s *AddDataCheckTemplateRequest) SetNullRules(v []*AddDataCheckTemplateRequestNullRules) *AddDataCheckTemplateRequest {
	s.NullRules = v
	return s
}

func (s *AddDataCheckTemplateRequest) SetRequestId(v string) *AddDataCheckTemplateRequest {
	s.RequestId = &v
	return s
}

func (s *AddDataCheckTemplateRequest) SetTemplateDesc(v string) *AddDataCheckTemplateRequest {
	s.TemplateDesc = &v
	return s
}

func (s *AddDataCheckTemplateRequest) SetTemplateName(v string) *AddDataCheckTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *AddDataCheckTemplateRequest) SetTenantId(v string) *AddDataCheckTemplateRequest {
	s.TenantId = &v
	return s
}

func (s *AddDataCheckTemplateRequest) SetWeakContentRule(v *AddDataCheckTemplateRequestWeakContentRule) *AddDataCheckTemplateRequest {
	s.WeakContentRule = v
	return s
}

func (s *AddDataCheckTemplateRequest) Validate() error {
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

type AddDataCheckTemplateRequestBasicMetricRules struct {
	// The verification methods (metric calculation methods). Separate multiple values with commas (,), such as SUM,AVG,MIN,MAX. The values must be within the range allowed by the templatetype.
	//
	// example:
	//
	// SUM,AVG
	CheckMethods *string `json:"checkMethods,omitempty" xml:"checkMethods,omitempty"`
	// Specifies whether to control floating-point precision. Valid values: 0 (no) and 1 (yes).
	//
	// example:
	//
	// 0
	ControlFloatPrecision *int32 `json:"controlFloatPrecision,omitempty" xml:"controlFloatPrecision,omitempty"`
	// The data type category. Valid values: 0 (native data type) and 1 (composite data type).
	//
	// example:
	//
	// 0
	DataTypeClassify *int32 `json:"dataTypeClassify,omitempty" xml:"dataTypeClassify,omitempty"`
	// The data type group that identifies the data type category to which the verification rule applies. Valid values: integers from 0 to 7. For the description of each value, see the enumeration values.
	//
	// example:
	//
	// 0
	DataTypeGroup *int32 `json:"dataTypeGroup,omitempty" xml:"dataTypeGroup,omitempty"`
	// The list of data types to which the verification rule applies. Configure this field based on your requirements.
	DataTypeList []*string `json:"dataTypeList,omitempty" xml:"dataTypeList,omitempty" type:"Repeated"`
	// The data types. Configure this field based on your requirements.
	//
	// example:
	//
	// BIGINT
	DataTypes *string `json:"dataTypes,omitempty" xml:"dataTypes,omitempty"`
	// The difference tolerance rate type. Valid values: 0 (unified) and 1 (custom). Default value: 0.
	//
	// example:
	//
	// 0
	DiffTolerateType *int32 `json:"diffTolerateType,omitempty" xml:"diffTolerateType,omitempty"`
	// The difference tolerance rate values. For the unified type, specify one value, such as {"SAME": 0}. For the custom type, specify a value for each tolerance type, such as {"SUM": 0.01, "AVG": 0.001}.
	DiffTolerateValues map[string]interface{} `json:"diffTolerateValues,omitempty" xml:"diffTolerateValues,omitempty"`
	// Specifies whether to enable decimal scale control for DECIMAL type comparison. Valid values: 0 (no) and 1 (yes).
	//
	// example:
	//
	// 0
	EnableDecimalScale *int32 `json:"enableDecimalScale,omitempty" xml:"enableDecimalScale,omitempty"`
	// The filter field names, separated by commas (,).
	//
	// example:
	//
	// col_a,col_b
	FilterColumnName *string `json:"filterColumnName,omitempty" xml:"filterColumnName,omitempty"`
	// **[Deprecated]*	- Use the filterColumnName field instead. This field is retained for backward compatibility.
	//
	// example:
	//
	// col_a,col_b
	FilterColumns *string `json:"filterColumns,omitempty" xml:"filterColumns,omitempty"`
	// The number of decimal places for floating-point values.
	//
	// example:
	//
	// 2
	FloatPrecision *int32 `json:"floatPrecision,omitempty" xml:"floatPrecision,omitempty"`
	// Specifies whether to ignore trailing zero differences in decimal parts. Valid values: 0 (no) and 1 (yes).
	//
	// example:
	//
	// 0
	IgnoreDecimalDiff *int32 `json:"ignoreDecimalDiff,omitempty" xml:"ignoreDecimalDiff,omitempty"`
	// Specifies whether to ignore trailing zeros in the decimal scale for DECIMAL type comparison. Valid values: 0 (no) and 1 (yes).
	//
	// example:
	//
	// 0
	IgnoreDecimalScaleSuffixZero *int32 `json:"ignoreDecimalScaleSuffixZero,omitempty" xml:"ignoreDecimalScaleSuffixZero,omitempty"`
	// Specifies whether to ignore differences between null values and empty strings. Valid values: 0 (no) and 1 (yes).
	//
	// example:
	//
	// 0
	IgnoreEmptyDiff *int32 `json:"ignoreEmptyDiff,omitempty" xml:"ignoreEmptyDiff,omitempty"`
	// Specifies whether to ignore zero values for numeric types. Valid values: 0 (no) and 1 (yes).
	//
	// example:
	//
	// 0
	IgnoreNumericZero *int32 `json:"ignoreNumericZero,omitempty" xml:"ignoreNumericZero,omitempty"`
	// Specifies whether to ignore empty strings and null values for string types. Valid values: 0 (no) and 1 (yes).
	//
	// example:
	//
	// 0
	IgnoreStringEmpty *int32 `json:"ignoreStringEmpty,omitempty" xml:"ignoreStringEmpty,omitempty"`
	// Specifies whether to ignore differences between null values and zero values. Valid values: 0 (no) and 1 (yes).
	//
	// example:
	//
	// 0
	IgnoreZeroDiff *int32 `json:"ignoreZeroDiff,omitempty" xml:"ignoreZeroDiff,omitempty"`
	// Specifies whether to enable count (data volume) verification. Valid values: 0 (no) and 1 (yes). Default value: 1.
	//
	// example:
	//
	// 1
	IsCountCheck *int32 `json:"isCountCheck,omitempty" xml:"isCountCheck,omitempty"`
	// The rule ID that uniquely identifies a verification rule.
	//
	// example:
	//
	// 1001
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// The specific decimal scale value for DECIMAL type comparison.
	//
	// example:
	//
	// 2
	SetDecimalScale *int32 `json:"setDecimalScale,omitempty" xml:"setDecimalScale,omitempty"`
}

func (s AddDataCheckTemplateRequestBasicMetricRules) String() string {
	return dara.Prettify(s)
}

func (s AddDataCheckTemplateRequestBasicMetricRules) GoString() string {
	return s.String()
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) GetCheckMethods() *string {
	return s.CheckMethods
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) GetControlFloatPrecision() *int32 {
	return s.ControlFloatPrecision
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) GetDataTypeClassify() *int32 {
	return s.DataTypeClassify
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) GetDataTypeGroup() *int32 {
	return s.DataTypeGroup
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) GetDataTypeList() []*string {
	return s.DataTypeList
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) GetDataTypes() *string {
	return s.DataTypes
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) GetDiffTolerateType() *int32 {
	return s.DiffTolerateType
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) GetDiffTolerateValues() map[string]interface{} {
	return s.DiffTolerateValues
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) GetEnableDecimalScale() *int32 {
	return s.EnableDecimalScale
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) GetFilterColumnName() *string {
	return s.FilterColumnName
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) GetFilterColumns() *string {
	return s.FilterColumns
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) GetFloatPrecision() *int32 {
	return s.FloatPrecision
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) GetIgnoreDecimalDiff() *int32 {
	return s.IgnoreDecimalDiff
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) GetIgnoreDecimalScaleSuffixZero() *int32 {
	return s.IgnoreDecimalScaleSuffixZero
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) GetIgnoreEmptyDiff() *int32 {
	return s.IgnoreEmptyDiff
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) GetIgnoreNumericZero() *int32 {
	return s.IgnoreNumericZero
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) GetIgnoreStringEmpty() *int32 {
	return s.IgnoreStringEmpty
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) GetIgnoreZeroDiff() *int32 {
	return s.IgnoreZeroDiff
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) GetIsCountCheck() *int32 {
	return s.IsCountCheck
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) GetRuleId() *string {
	return s.RuleId
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) GetSetDecimalScale() *int32 {
	return s.SetDecimalScale
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) SetCheckMethods(v string) *AddDataCheckTemplateRequestBasicMetricRules {
	s.CheckMethods = &v
	return s
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) SetControlFloatPrecision(v int32) *AddDataCheckTemplateRequestBasicMetricRules {
	s.ControlFloatPrecision = &v
	return s
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) SetDataTypeClassify(v int32) *AddDataCheckTemplateRequestBasicMetricRules {
	s.DataTypeClassify = &v
	return s
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) SetDataTypeGroup(v int32) *AddDataCheckTemplateRequestBasicMetricRules {
	s.DataTypeGroup = &v
	return s
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) SetDataTypeList(v []*string) *AddDataCheckTemplateRequestBasicMetricRules {
	s.DataTypeList = v
	return s
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) SetDataTypes(v string) *AddDataCheckTemplateRequestBasicMetricRules {
	s.DataTypes = &v
	return s
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) SetDiffTolerateType(v int32) *AddDataCheckTemplateRequestBasicMetricRules {
	s.DiffTolerateType = &v
	return s
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) SetDiffTolerateValues(v map[string]interface{}) *AddDataCheckTemplateRequestBasicMetricRules {
	s.DiffTolerateValues = v
	return s
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) SetEnableDecimalScale(v int32) *AddDataCheckTemplateRequestBasicMetricRules {
	s.EnableDecimalScale = &v
	return s
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) SetFilterColumnName(v string) *AddDataCheckTemplateRequestBasicMetricRules {
	s.FilterColumnName = &v
	return s
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) SetFilterColumns(v string) *AddDataCheckTemplateRequestBasicMetricRules {
	s.FilterColumns = &v
	return s
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) SetFloatPrecision(v int32) *AddDataCheckTemplateRequestBasicMetricRules {
	s.FloatPrecision = &v
	return s
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) SetIgnoreDecimalDiff(v int32) *AddDataCheckTemplateRequestBasicMetricRules {
	s.IgnoreDecimalDiff = &v
	return s
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) SetIgnoreDecimalScaleSuffixZero(v int32) *AddDataCheckTemplateRequestBasicMetricRules {
	s.IgnoreDecimalScaleSuffixZero = &v
	return s
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) SetIgnoreEmptyDiff(v int32) *AddDataCheckTemplateRequestBasicMetricRules {
	s.IgnoreEmptyDiff = &v
	return s
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) SetIgnoreNumericZero(v int32) *AddDataCheckTemplateRequestBasicMetricRules {
	s.IgnoreNumericZero = &v
	return s
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) SetIgnoreStringEmpty(v int32) *AddDataCheckTemplateRequestBasicMetricRules {
	s.IgnoreStringEmpty = &v
	return s
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) SetIgnoreZeroDiff(v int32) *AddDataCheckTemplateRequestBasicMetricRules {
	s.IgnoreZeroDiff = &v
	return s
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) SetIsCountCheck(v int32) *AddDataCheckTemplateRequestBasicMetricRules {
	s.IsCountCheck = &v
	return s
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) SetRuleId(v string) *AddDataCheckTemplateRequestBasicMetricRules {
	s.RuleId = &v
	return s
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) SetSetDecimalScale(v int32) *AddDataCheckTemplateRequestBasicMetricRules {
	s.SetDecimalScale = &v
	return s
}

func (s *AddDataCheckTemplateRequestBasicMetricRules) Validate() error {
	return dara.Validate(s)
}

type AddDataCheckTemplateRequestComplexMetricRules struct {
	// The check methods (metric calculation methods). Separate multiple values with commas, such as SUM,AVG,MIN,MAX. The values must be within the range allowed by the templatetype.
	//
	// example:
	//
	// SUM,AVG
	CheckMethods *string `json:"checkMethods,omitempty" xml:"checkMethods,omitempty"`
	// Specifies whether to control floating-point precision. Valid values: 0 (no) and 1 (yes).
	//
	// example:
	//
	// 0
	ControlFloatPrecision *int32 `json:"controlFloatPrecision,omitempty" xml:"controlFloatPrecision,omitempty"`
	// The data type category. Valid values: 0 (native data type) and 1 (composite data type).
	//
	// example:
	//
	// 0
	DataTypeClassify *int32 `json:"dataTypeClassify,omitempty" xml:"dataTypeClassify,omitempty"`
	// The data type group that identifies the data type category to which the verification rule applies. Valid values: integers from 0 to 7. For the description of each value, see the enumeration values.
	//
	// example:
	//
	// 0
	DataTypeGroup *int32 `json:"dataTypeGroup,omitempty" xml:"dataTypeGroup,omitempty"`
	// The list of data types to which the verification rule applies. Configure this field based on your requirements.
	DataTypeList []*string `json:"dataTypeList,omitempty" xml:"dataTypeList,omitempty" type:"Repeated"`
	// The data types. Configure this field based on your requirements.
	//
	// example:
	//
	// BIGINT
	DataTypes *string `json:"dataTypes,omitempty" xml:"dataTypes,omitempty"`
	// The difference tolerance rate type. Valid values: 0 (unified) and 1 (custom). Default value: 0.
	//
	// example:
	//
	// 0
	DiffTolerateType *int32 `json:"diffTolerateType,omitempty" xml:"diffTolerateType,omitempty"`
	// The difference tolerance rate values. For the unified type, specify one value, such as {"SAME": 0}. For the custom type, specify a value for each tolerance type, such as {"SUM": 0.01, "AVG": 0.001}.
	DiffTolerateValues map[string]interface{} `json:"diffTolerateValues,omitempty" xml:"diffTolerateValues,omitempty"`
	// Specifies whether to enable decimal scale control for DECIMAL type comparison. Valid values: 0 (no) and 1 (yes).
	//
	// example:
	//
	// 0
	EnableDecimalScale *int32 `json:"enableDecimalScale,omitempty" xml:"enableDecimalScale,omitempty"`
	// The filter field names, separated by commas (,).
	//
	// example:
	//
	// col_a,col_b
	FilterColumnName *string `json:"filterColumnName,omitempty" xml:"filterColumnName,omitempty"`
	// **[Deprecated]*	- Use the filterColumnName field instead. This field is retained for backward compatibility.
	//
	// example:
	//
	// col_a,col_b
	FilterColumns *string `json:"filterColumns,omitempty" xml:"filterColumns,omitempty"`
	// The number of decimal places for floating-point values.
	//
	// example:
	//
	// 2
	FloatPrecision *int32 `json:"floatPrecision,omitempty" xml:"floatPrecision,omitempty"`
	// Specifies whether to ignore trailing zero differences in decimal parts. Valid values: 0 (no) and 1 (yes).
	//
	// example:
	//
	// 0
	IgnoreDecimalDiff *int32 `json:"ignoreDecimalDiff,omitempty" xml:"ignoreDecimalDiff,omitempty"`
	// Specifies whether to ignore trailing zeros in the decimal scale for DECIMAL type comparison. Valid values: 0 (no) and 1 (yes).
	//
	// example:
	//
	// 0
	IgnoreDecimalScaleSuffixZero *int32 `json:"ignoreDecimalScaleSuffixZero,omitempty" xml:"ignoreDecimalScaleSuffixZero,omitempty"`
	// Specifies whether to ignore differences between null values and empty strings. Valid values: 0 (no) and 1 (yes).
	//
	// example:
	//
	// 0
	IgnoreEmptyDiff *int32 `json:"ignoreEmptyDiff,omitempty" xml:"ignoreEmptyDiff,omitempty"`
	// Specifies whether to ignore zero values for numeric types. Valid values: 0 (no) and 1 (yes).
	//
	// example:
	//
	// 0
	IgnoreNumericZero *int32 `json:"ignoreNumericZero,omitempty" xml:"ignoreNumericZero,omitempty"`
	// Specifies whether to ignore empty strings and null values for string types. Valid values: 0 (no) and 1 (yes).
	//
	// example:
	//
	// 0
	IgnoreStringEmpty *int32 `json:"ignoreStringEmpty,omitempty" xml:"ignoreStringEmpty,omitempty"`
	// Specifies whether to ignore differences between null values and zero values. Valid values: 0 (no) and 1 (yes).
	//
	// example:
	//
	// 0
	IgnoreZeroDiff *int32 `json:"ignoreZeroDiff,omitempty" xml:"ignoreZeroDiff,omitempty"`
	// Specifies whether to enable count (data volume) verification. Valid values: 0 (no) and 1 (yes). Default value: 1.
	//
	// example:
	//
	// 1
	IsCountCheck *int32 `json:"isCountCheck,omitempty" xml:"isCountCheck,omitempty"`
	// The rule ID that uniquely identifies a verification rule.
	//
	// example:
	//
	// 1001
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// The specific decimal scale value for DECIMAL type comparison.
	//
	// example:
	//
	// 2
	SetDecimalScale *int32 `json:"setDecimalScale,omitempty" xml:"setDecimalScale,omitempty"`
}

func (s AddDataCheckTemplateRequestComplexMetricRules) String() string {
	return dara.Prettify(s)
}

func (s AddDataCheckTemplateRequestComplexMetricRules) GoString() string {
	return s.String()
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) GetCheckMethods() *string {
	return s.CheckMethods
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) GetControlFloatPrecision() *int32 {
	return s.ControlFloatPrecision
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) GetDataTypeClassify() *int32 {
	return s.DataTypeClassify
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) GetDataTypeGroup() *int32 {
	return s.DataTypeGroup
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) GetDataTypeList() []*string {
	return s.DataTypeList
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) GetDataTypes() *string {
	return s.DataTypes
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) GetDiffTolerateType() *int32 {
	return s.DiffTolerateType
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) GetDiffTolerateValues() map[string]interface{} {
	return s.DiffTolerateValues
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) GetEnableDecimalScale() *int32 {
	return s.EnableDecimalScale
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) GetFilterColumnName() *string {
	return s.FilterColumnName
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) GetFilterColumns() *string {
	return s.FilterColumns
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) GetFloatPrecision() *int32 {
	return s.FloatPrecision
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) GetIgnoreDecimalDiff() *int32 {
	return s.IgnoreDecimalDiff
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) GetIgnoreDecimalScaleSuffixZero() *int32 {
	return s.IgnoreDecimalScaleSuffixZero
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) GetIgnoreEmptyDiff() *int32 {
	return s.IgnoreEmptyDiff
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) GetIgnoreNumericZero() *int32 {
	return s.IgnoreNumericZero
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) GetIgnoreStringEmpty() *int32 {
	return s.IgnoreStringEmpty
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) GetIgnoreZeroDiff() *int32 {
	return s.IgnoreZeroDiff
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) GetIsCountCheck() *int32 {
	return s.IsCountCheck
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) GetRuleId() *string {
	return s.RuleId
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) GetSetDecimalScale() *int32 {
	return s.SetDecimalScale
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) SetCheckMethods(v string) *AddDataCheckTemplateRequestComplexMetricRules {
	s.CheckMethods = &v
	return s
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) SetControlFloatPrecision(v int32) *AddDataCheckTemplateRequestComplexMetricRules {
	s.ControlFloatPrecision = &v
	return s
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) SetDataTypeClassify(v int32) *AddDataCheckTemplateRequestComplexMetricRules {
	s.DataTypeClassify = &v
	return s
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) SetDataTypeGroup(v int32) *AddDataCheckTemplateRequestComplexMetricRules {
	s.DataTypeGroup = &v
	return s
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) SetDataTypeList(v []*string) *AddDataCheckTemplateRequestComplexMetricRules {
	s.DataTypeList = v
	return s
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) SetDataTypes(v string) *AddDataCheckTemplateRequestComplexMetricRules {
	s.DataTypes = &v
	return s
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) SetDiffTolerateType(v int32) *AddDataCheckTemplateRequestComplexMetricRules {
	s.DiffTolerateType = &v
	return s
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) SetDiffTolerateValues(v map[string]interface{}) *AddDataCheckTemplateRequestComplexMetricRules {
	s.DiffTolerateValues = v
	return s
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) SetEnableDecimalScale(v int32) *AddDataCheckTemplateRequestComplexMetricRules {
	s.EnableDecimalScale = &v
	return s
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) SetFilterColumnName(v string) *AddDataCheckTemplateRequestComplexMetricRules {
	s.FilterColumnName = &v
	return s
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) SetFilterColumns(v string) *AddDataCheckTemplateRequestComplexMetricRules {
	s.FilterColumns = &v
	return s
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) SetFloatPrecision(v int32) *AddDataCheckTemplateRequestComplexMetricRules {
	s.FloatPrecision = &v
	return s
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) SetIgnoreDecimalDiff(v int32) *AddDataCheckTemplateRequestComplexMetricRules {
	s.IgnoreDecimalDiff = &v
	return s
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) SetIgnoreDecimalScaleSuffixZero(v int32) *AddDataCheckTemplateRequestComplexMetricRules {
	s.IgnoreDecimalScaleSuffixZero = &v
	return s
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) SetIgnoreEmptyDiff(v int32) *AddDataCheckTemplateRequestComplexMetricRules {
	s.IgnoreEmptyDiff = &v
	return s
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) SetIgnoreNumericZero(v int32) *AddDataCheckTemplateRequestComplexMetricRules {
	s.IgnoreNumericZero = &v
	return s
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) SetIgnoreStringEmpty(v int32) *AddDataCheckTemplateRequestComplexMetricRules {
	s.IgnoreStringEmpty = &v
	return s
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) SetIgnoreZeroDiff(v int32) *AddDataCheckTemplateRequestComplexMetricRules {
	s.IgnoreZeroDiff = &v
	return s
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) SetIsCountCheck(v int32) *AddDataCheckTemplateRequestComplexMetricRules {
	s.IsCountCheck = &v
	return s
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) SetRuleId(v string) *AddDataCheckTemplateRequestComplexMetricRules {
	s.RuleId = &v
	return s
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) SetSetDecimalScale(v int32) *AddDataCheckTemplateRequestComplexMetricRules {
	s.SetDecimalScale = &v
	return s
}

func (s *AddDataCheckTemplateRequestComplexMetricRules) Validate() error {
	return dara.Validate(s)
}

type AddDataCheckTemplateRequestDsEngineRels struct {
	// The ID of the datasource engine configuration.
	//
	// example:
	//
	// 1001
	DsEngineId *string `json:"dsEngineId,omitempty" xml:"dsEngineId,omitempty"`
	// The datasource type, such as Hive or MaxCompute.
	//
	// example:
	//
	// Hive
	DsType *string `json:"dsType,omitempty" xml:"dsType,omitempty"`
	// The list of covered check engine types, such as Tez or MapReduce. When specified as a string, separate multiple values with commas.
	EngineTypes []*string `json:"engineTypes,omitempty" xml:"engineTypes,omitempty" type:"Repeated"`
}

func (s AddDataCheckTemplateRequestDsEngineRels) String() string {
	return dara.Prettify(s)
}

func (s AddDataCheckTemplateRequestDsEngineRels) GoString() string {
	return s.String()
}

func (s *AddDataCheckTemplateRequestDsEngineRels) GetDsEngineId() *string {
	return s.DsEngineId
}

func (s *AddDataCheckTemplateRequestDsEngineRels) GetDsType() *string {
	return s.DsType
}

func (s *AddDataCheckTemplateRequestDsEngineRels) GetEngineTypes() []*string {
	return s.EngineTypes
}

func (s *AddDataCheckTemplateRequestDsEngineRels) SetDsEngineId(v string) *AddDataCheckTemplateRequestDsEngineRels {
	s.DsEngineId = &v
	return s
}

func (s *AddDataCheckTemplateRequestDsEngineRels) SetDsType(v string) *AddDataCheckTemplateRequestDsEngineRels {
	s.DsType = &v
	return s
}

func (s *AddDataCheckTemplateRequestDsEngineRels) SetEngineTypes(v []*string) *AddDataCheckTemplateRequestDsEngineRels {
	s.EngineTypes = v
	return s
}

func (s *AddDataCheckTemplateRequestDsEngineRels) Validate() error {
	return dara.Validate(s)
}

type AddDataCheckTemplateRequestFulltextRule struct {
	// The check mode. Valid values:
	//
	// - 0: row-by-row overall comparison.
	//
	// - 1: row-by-row column-by-column comparison.
	//
	// - 2: both row-by-row overall comparison and row-by-row column-by-column comparison.
	//
	// example:
	//
	// 0
	CheckMode *int32 `json:"checkMode,omitempty" xml:"checkMode,omitempty"`
	// The equality comparison type for row-by-row column-by-column comparison. Valid values:
	//
	// - 0: all field types.
	//
	// - 1: native primitive data types.
	//
	// - 2: complex data types.
	//
	// - 3: custom.
	//
	// example:
	//
	// 0
	ColumnEqualCmpType *int32 `json:"columnEqualCmpType,omitempty" xml:"columnEqualCmpType,omitempty"`
	// The custom type list for equality comparison during row-by-row column-by-column comparison. Separate multiple values with commas.
	//
	// example:
	//
	// ARRAY,MAP
	ColumnEqualCmpValues *string `json:"columnEqualCmpValues,omitempty" xml:"columnEqualCmpValues,omitempty"`
	// Specifies whether to enable cosine similarity during row-by-row column-by-column comparison. Valid values:
	//
	// - 0: Disabled.
	//
	// - 1: Enabled.
	//
	// example:
	//
	// 0
	ColumnIsCosine *int32 `json:"columnIsCosine,omitempty" xml:"columnIsCosine,omitempty"`
	// Specifies whether to ignore differences between null values and empty strings during row-by-row column-by-column comparison. Valid values:
	//
	// - 0: Not ignored.
	//
	// - 1: Ignored.
	//
	// example:
	//
	// 0
	ColumnIsIgnoreNull *int32 `json:"columnIsIgnoreNull,omitempty" xml:"columnIsIgnoreNull,omitempty"`
	// Specifies whether to ignore differences between null values and 0 values during row-by-row column-by-column comparison. Valid values:
	//
	// - 0: Not ignored.
	//
	// - 1: Ignored.
	//
	// example:
	//
	// 0
	ColumnIsIgnoreZero *int32 `json:"columnIsIgnoreZero,omitempty" xml:"columnIsIgnoreZero,omitempty"`
	// Specifies whether to enable sampling during row-by-row column-by-column comparison. Valid values:
	//
	// - 0: Disabled.
	//
	// - 1: Enabled.
	//
	// example:
	//
	// 0
	ColumnIsSamples *int32 `json:"columnIsSamples,omitempty" xml:"columnIsSamples,omitempty"`
	// The sampling method for row-by-row column-by-column comparison. Valid values:
	//
	// - 0: by row.
	//
	// - 1: by percentage.
	//
	// example:
	//
	// 0
	ColumnSamplesType *int32 `json:"columnSamplesType,omitempty" xml:"columnSamplesType,omitempty"`
	// The sampling value for row-by-row column-by-column comparison. The meaning depends on the sampling method: the number of rows when sampling by row, or the percentage value when sampling by percentage.
	//
	// example:
	//
	// 100
	ColumnSamplesValue *int32 `json:"columnSamplesValue,omitempty" xml:"columnSamplesValue,omitempty"`
	// The size comparison type for row-by-row column-by-column comparison. Valid values:
	//
	// - 0: all complex data types.
	//
	// - 1: custom.
	//
	// example:
	//
	// 0
	ColumnSizeCmpType *int32 `json:"columnSizeCmpType,omitempty" xml:"columnSizeCmpType,omitempty"`
	// The custom type list for size comparison during row-by-row column-by-column comparison. Separate multiple values with commas.
	//
	// example:
	//
	// ARRAY,MAP
	ColumnSizeCmpValues *string `json:"columnSizeCmpValues,omitempty" xml:"columnSizeCmpValues,omitempty"`
	// Specifies whether to enable primary key or composite primary key existence check. Valid values:
	//
	// - 0: Disabled.
	//
	// - 1: Enabled.
	//
	// example:
	//
	// 1
	IsPrimaryKeyCheck *int32 `json:"isPrimaryKeyCheck,omitempty" xml:"isPrimaryKeyCheck,omitempty"`
	// The row-by-row comparison method. Valid values:
	//
	// - 0: md5.
	//
	// - 1: crc32.
	//
	// example:
	//
	// 0
	LineCheckType *int32 `json:"lineCheckType,omitempty" xml:"lineCheckType,omitempty"`
	// Specifies whether to print all columns in the difference details during row-by-row comparison. Valid values:
	//
	// - 0: Not printed.
	//
	// - 1: Printed.
	//
	// example:
	//
	// 0
	LineIsPrintAll *int32 `json:"lineIsPrintAll,omitempty" xml:"lineIsPrintAll,omitempty"`
	// Specifies whether to enable sampling during row-by-row comparison. Valid values:
	//
	// - 0: Disabled.
	//
	// - 1: Enabled.
	//
	// example:
	//
	// 0
	LineIsSamples *int32 `json:"lineIsSamples,omitempty" xml:"lineIsSamples,omitempty"`
	// The sampling method for row-by-row comparison. Valid values:
	//
	// - 0: by row.
	//
	// - 1: by percentage.
	//
	// example:
	//
	// 0
	LineSamplesType *int32 `json:"lineSamplesType,omitempty" xml:"lineSamplesType,omitempty"`
	// The sampling value for row-by-row comparison. The meaning depends on the sampling method: the number of rows when sampling by row, or the percentage value when sampling by percentage.
	//
	// example:
	//
	// 100
	LineSamplesValue *int32 `json:"lineSamplesValue,omitempty" xml:"lineSamplesValue,omitempty"`
	// The rule ID that uniquely identifies a verification rule.
	//
	// example:
	//
	// 1001
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
}

func (s AddDataCheckTemplateRequestFulltextRule) String() string {
	return dara.Prettify(s)
}

func (s AddDataCheckTemplateRequestFulltextRule) GoString() string {
	return s.String()
}

func (s *AddDataCheckTemplateRequestFulltextRule) GetCheckMode() *int32 {
	return s.CheckMode
}

func (s *AddDataCheckTemplateRequestFulltextRule) GetColumnEqualCmpType() *int32 {
	return s.ColumnEqualCmpType
}

func (s *AddDataCheckTemplateRequestFulltextRule) GetColumnEqualCmpValues() *string {
	return s.ColumnEqualCmpValues
}

func (s *AddDataCheckTemplateRequestFulltextRule) GetColumnIsCosine() *int32 {
	return s.ColumnIsCosine
}

func (s *AddDataCheckTemplateRequestFulltextRule) GetColumnIsIgnoreNull() *int32 {
	return s.ColumnIsIgnoreNull
}

func (s *AddDataCheckTemplateRequestFulltextRule) GetColumnIsIgnoreZero() *int32 {
	return s.ColumnIsIgnoreZero
}

func (s *AddDataCheckTemplateRequestFulltextRule) GetColumnIsSamples() *int32 {
	return s.ColumnIsSamples
}

func (s *AddDataCheckTemplateRequestFulltextRule) GetColumnSamplesType() *int32 {
	return s.ColumnSamplesType
}

func (s *AddDataCheckTemplateRequestFulltextRule) GetColumnSamplesValue() *int32 {
	return s.ColumnSamplesValue
}

func (s *AddDataCheckTemplateRequestFulltextRule) GetColumnSizeCmpType() *int32 {
	return s.ColumnSizeCmpType
}

func (s *AddDataCheckTemplateRequestFulltextRule) GetColumnSizeCmpValues() *string {
	return s.ColumnSizeCmpValues
}

func (s *AddDataCheckTemplateRequestFulltextRule) GetIsPrimaryKeyCheck() *int32 {
	return s.IsPrimaryKeyCheck
}

func (s *AddDataCheckTemplateRequestFulltextRule) GetLineCheckType() *int32 {
	return s.LineCheckType
}

func (s *AddDataCheckTemplateRequestFulltextRule) GetLineIsPrintAll() *int32 {
	return s.LineIsPrintAll
}

func (s *AddDataCheckTemplateRequestFulltextRule) GetLineIsSamples() *int32 {
	return s.LineIsSamples
}

func (s *AddDataCheckTemplateRequestFulltextRule) GetLineSamplesType() *int32 {
	return s.LineSamplesType
}

func (s *AddDataCheckTemplateRequestFulltextRule) GetLineSamplesValue() *int32 {
	return s.LineSamplesValue
}

func (s *AddDataCheckTemplateRequestFulltextRule) GetRuleId() *string {
	return s.RuleId
}

func (s *AddDataCheckTemplateRequestFulltextRule) SetCheckMode(v int32) *AddDataCheckTemplateRequestFulltextRule {
	s.CheckMode = &v
	return s
}

func (s *AddDataCheckTemplateRequestFulltextRule) SetColumnEqualCmpType(v int32) *AddDataCheckTemplateRequestFulltextRule {
	s.ColumnEqualCmpType = &v
	return s
}

func (s *AddDataCheckTemplateRequestFulltextRule) SetColumnEqualCmpValues(v string) *AddDataCheckTemplateRequestFulltextRule {
	s.ColumnEqualCmpValues = &v
	return s
}

func (s *AddDataCheckTemplateRequestFulltextRule) SetColumnIsCosine(v int32) *AddDataCheckTemplateRequestFulltextRule {
	s.ColumnIsCosine = &v
	return s
}

func (s *AddDataCheckTemplateRequestFulltextRule) SetColumnIsIgnoreNull(v int32) *AddDataCheckTemplateRequestFulltextRule {
	s.ColumnIsIgnoreNull = &v
	return s
}

func (s *AddDataCheckTemplateRequestFulltextRule) SetColumnIsIgnoreZero(v int32) *AddDataCheckTemplateRequestFulltextRule {
	s.ColumnIsIgnoreZero = &v
	return s
}

func (s *AddDataCheckTemplateRequestFulltextRule) SetColumnIsSamples(v int32) *AddDataCheckTemplateRequestFulltextRule {
	s.ColumnIsSamples = &v
	return s
}

func (s *AddDataCheckTemplateRequestFulltextRule) SetColumnSamplesType(v int32) *AddDataCheckTemplateRequestFulltextRule {
	s.ColumnSamplesType = &v
	return s
}

func (s *AddDataCheckTemplateRequestFulltextRule) SetColumnSamplesValue(v int32) *AddDataCheckTemplateRequestFulltextRule {
	s.ColumnSamplesValue = &v
	return s
}

func (s *AddDataCheckTemplateRequestFulltextRule) SetColumnSizeCmpType(v int32) *AddDataCheckTemplateRequestFulltextRule {
	s.ColumnSizeCmpType = &v
	return s
}

func (s *AddDataCheckTemplateRequestFulltextRule) SetColumnSizeCmpValues(v string) *AddDataCheckTemplateRequestFulltextRule {
	s.ColumnSizeCmpValues = &v
	return s
}

func (s *AddDataCheckTemplateRequestFulltextRule) SetIsPrimaryKeyCheck(v int32) *AddDataCheckTemplateRequestFulltextRule {
	s.IsPrimaryKeyCheck = &v
	return s
}

func (s *AddDataCheckTemplateRequestFulltextRule) SetLineCheckType(v int32) *AddDataCheckTemplateRequestFulltextRule {
	s.LineCheckType = &v
	return s
}

func (s *AddDataCheckTemplateRequestFulltextRule) SetLineIsPrintAll(v int32) *AddDataCheckTemplateRequestFulltextRule {
	s.LineIsPrintAll = &v
	return s
}

func (s *AddDataCheckTemplateRequestFulltextRule) SetLineIsSamples(v int32) *AddDataCheckTemplateRequestFulltextRule {
	s.LineIsSamples = &v
	return s
}

func (s *AddDataCheckTemplateRequestFulltextRule) SetLineSamplesType(v int32) *AddDataCheckTemplateRequestFulltextRule {
	s.LineSamplesType = &v
	return s
}

func (s *AddDataCheckTemplateRequestFulltextRule) SetLineSamplesValue(v int32) *AddDataCheckTemplateRequestFulltextRule {
	s.LineSamplesValue = &v
	return s
}

func (s *AddDataCheckTemplateRequestFulltextRule) SetRuleId(v string) *AddDataCheckTemplateRequestFulltextRule {
	s.RuleId = &v
	return s
}

func (s *AddDataCheckTemplateRequestFulltextRule) Validate() error {
	return dara.Validate(s)
}

type AddDataCheckTemplateRequestMetricRules struct {
	// The check methods (metric calculation methods). Separate multiple values with commas, such as SUM,AVG,MIN,MAX. The values must be within the range allowed by the templatetype.
	//
	// example:
	//
	// SUM,AVG
	CheckMethods *string `json:"checkMethods,omitempty" xml:"checkMethods,omitempty"`
	// Specifies whether to control floating-point precision. Valid values: 0 (no) and 1 (yes).
	//
	// example:
	//
	// 0
	ControlFloatPrecision *int32 `json:"controlFloatPrecision,omitempty" xml:"controlFloatPrecision,omitempty"`
	// The data type category. Valid values: 0 (native data type) and 1 (composite data type).
	//
	// example:
	//
	// 0
	DataTypeClassify *int32 `json:"dataTypeClassify,omitempty" xml:"dataTypeClassify,omitempty"`
	// The data type group that identifies the data type category to which the verification rule applies. Valid values: integers from 0 to 7. For the description of each value, see the enumeration values.
	//
	// example:
	//
	// 0
	DataTypeGroup *int32 `json:"dataTypeGroup,omitempty" xml:"dataTypeGroup,omitempty"`
	// The list of data types to which the verification rule applies. Configure this field based on your requirements.
	DataTypeList []*string `json:"dataTypeList,omitempty" xml:"dataTypeList,omitempty" type:"Repeated"`
	// The data types. Configure this field based on your requirements.
	//
	// example:
	//
	// BIGINT
	DataTypes *string `json:"dataTypes,omitempty" xml:"dataTypes,omitempty"`
	// The difference tolerance rate type. Valid values: 0 (unified) and 1 (custom). Default value: 0.
	//
	// example:
	//
	// 0
	DiffTolerateType *int32 `json:"diffTolerateType,omitempty" xml:"diffTolerateType,omitempty"`
	// The difference tolerance rate values. For the unified type, specify one value, such as {"SAME": 0}. For the custom type, specify a value for each tolerance type, such as {"SUM": 0.01, "AVG": 0.001}.
	DiffTolerateValues map[string]interface{} `json:"diffTolerateValues,omitempty" xml:"diffTolerateValues,omitempty"`
	// Specifies whether to enable decimal scale control for DECIMAL type comparison. Valid values: 0 (no) and 1 (yes).
	//
	// example:
	//
	// 0
	EnableDecimalScale *int32 `json:"enableDecimalScale,omitempty" xml:"enableDecimalScale,omitempty"`
	// The filter field names, separated by commas (,).
	//
	// example:
	//
	// col_a,col_b
	FilterColumnName *string `json:"filterColumnName,omitempty" xml:"filterColumnName,omitempty"`
	// **[Deprecated]*	- Use the filterColumnName field instead. This field is retained for backward compatibility.
	//
	// example:
	//
	// col_a,col_b
	FilterColumns *string `json:"filterColumns,omitempty" xml:"filterColumns,omitempty"`
	// The number of decimal places for floating-point values.
	//
	// example:
	//
	// 2
	FloatPrecision *int32 `json:"floatPrecision,omitempty" xml:"floatPrecision,omitempty"`
	// Specifies whether to ignore trailing zero differences in decimal parts. Valid values: 0 (no) and 1 (yes).
	//
	// example:
	//
	// 0
	IgnoreDecimalDiff *int32 `json:"ignoreDecimalDiff,omitempty" xml:"ignoreDecimalDiff,omitempty"`
	// Specifies whether to ignore trailing zeros in the decimal scale for DECIMAL type comparison. Valid values: 0 (no) and 1 (yes).
	//
	// example:
	//
	// 0
	IgnoreDecimalScaleSuffixZero *int32 `json:"ignoreDecimalScaleSuffixZero,omitempty" xml:"ignoreDecimalScaleSuffixZero,omitempty"`
	// Specifies whether to ignore differences between null values and empty strings. Valid values: 0 (no) and 1 (yes).
	//
	// example:
	//
	// 0
	IgnoreEmptyDiff *int32 `json:"ignoreEmptyDiff,omitempty" xml:"ignoreEmptyDiff,omitempty"`
	// Specifies whether to ignore zero values for numeric types. Valid values: 0 (no) and 1 (yes).
	//
	// example:
	//
	// 0
	IgnoreNumericZero *int32 `json:"ignoreNumericZero,omitempty" xml:"ignoreNumericZero,omitempty"`
	// Specifies whether to ignore empty strings and null values for string types. Valid values: 0 (no) and 1 (yes).
	//
	// example:
	//
	// 0
	IgnoreStringEmpty *int32 `json:"ignoreStringEmpty,omitempty" xml:"ignoreStringEmpty,omitempty"`
	// Specifies whether to ignore differences between null values and zero values. Valid values: 0 (no) and 1 (yes).
	//
	// example:
	//
	// 0
	IgnoreZeroDiff *int32 `json:"ignoreZeroDiff,omitempty" xml:"ignoreZeroDiff,omitempty"`
	// Specifies whether to enable count (data volume) verification. Valid values: 0 (no) and 1 (yes). Default value: 1.
	//
	// example:
	//
	// 1
	IsCountCheck *int32 `json:"isCountCheck,omitempty" xml:"isCountCheck,omitempty"`
	// The rule ID that uniquely identifies a verification rule.
	//
	// example:
	//
	// 1001
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// The specific decimal scale value for DECIMAL type comparison.
	//
	// example:
	//
	// 2
	SetDecimalScale *int32 `json:"setDecimalScale,omitempty" xml:"setDecimalScale,omitempty"`
}

func (s AddDataCheckTemplateRequestMetricRules) String() string {
	return dara.Prettify(s)
}

func (s AddDataCheckTemplateRequestMetricRules) GoString() string {
	return s.String()
}

func (s *AddDataCheckTemplateRequestMetricRules) GetCheckMethods() *string {
	return s.CheckMethods
}

func (s *AddDataCheckTemplateRequestMetricRules) GetControlFloatPrecision() *int32 {
	return s.ControlFloatPrecision
}

func (s *AddDataCheckTemplateRequestMetricRules) GetDataTypeClassify() *int32 {
	return s.DataTypeClassify
}

func (s *AddDataCheckTemplateRequestMetricRules) GetDataTypeGroup() *int32 {
	return s.DataTypeGroup
}

func (s *AddDataCheckTemplateRequestMetricRules) GetDataTypeList() []*string {
	return s.DataTypeList
}

func (s *AddDataCheckTemplateRequestMetricRules) GetDataTypes() *string {
	return s.DataTypes
}

func (s *AddDataCheckTemplateRequestMetricRules) GetDiffTolerateType() *int32 {
	return s.DiffTolerateType
}

func (s *AddDataCheckTemplateRequestMetricRules) GetDiffTolerateValues() map[string]interface{} {
	return s.DiffTolerateValues
}

func (s *AddDataCheckTemplateRequestMetricRules) GetEnableDecimalScale() *int32 {
	return s.EnableDecimalScale
}

func (s *AddDataCheckTemplateRequestMetricRules) GetFilterColumnName() *string {
	return s.FilterColumnName
}

func (s *AddDataCheckTemplateRequestMetricRules) GetFilterColumns() *string {
	return s.FilterColumns
}

func (s *AddDataCheckTemplateRequestMetricRules) GetFloatPrecision() *int32 {
	return s.FloatPrecision
}

func (s *AddDataCheckTemplateRequestMetricRules) GetIgnoreDecimalDiff() *int32 {
	return s.IgnoreDecimalDiff
}

func (s *AddDataCheckTemplateRequestMetricRules) GetIgnoreDecimalScaleSuffixZero() *int32 {
	return s.IgnoreDecimalScaleSuffixZero
}

func (s *AddDataCheckTemplateRequestMetricRules) GetIgnoreEmptyDiff() *int32 {
	return s.IgnoreEmptyDiff
}

func (s *AddDataCheckTemplateRequestMetricRules) GetIgnoreNumericZero() *int32 {
	return s.IgnoreNumericZero
}

func (s *AddDataCheckTemplateRequestMetricRules) GetIgnoreStringEmpty() *int32 {
	return s.IgnoreStringEmpty
}

func (s *AddDataCheckTemplateRequestMetricRules) GetIgnoreZeroDiff() *int32 {
	return s.IgnoreZeroDiff
}

func (s *AddDataCheckTemplateRequestMetricRules) GetIsCountCheck() *int32 {
	return s.IsCountCheck
}

func (s *AddDataCheckTemplateRequestMetricRules) GetRuleId() *string {
	return s.RuleId
}

func (s *AddDataCheckTemplateRequestMetricRules) GetSetDecimalScale() *int32 {
	return s.SetDecimalScale
}

func (s *AddDataCheckTemplateRequestMetricRules) SetCheckMethods(v string) *AddDataCheckTemplateRequestMetricRules {
	s.CheckMethods = &v
	return s
}

func (s *AddDataCheckTemplateRequestMetricRules) SetControlFloatPrecision(v int32) *AddDataCheckTemplateRequestMetricRules {
	s.ControlFloatPrecision = &v
	return s
}

func (s *AddDataCheckTemplateRequestMetricRules) SetDataTypeClassify(v int32) *AddDataCheckTemplateRequestMetricRules {
	s.DataTypeClassify = &v
	return s
}

func (s *AddDataCheckTemplateRequestMetricRules) SetDataTypeGroup(v int32) *AddDataCheckTemplateRequestMetricRules {
	s.DataTypeGroup = &v
	return s
}

func (s *AddDataCheckTemplateRequestMetricRules) SetDataTypeList(v []*string) *AddDataCheckTemplateRequestMetricRules {
	s.DataTypeList = v
	return s
}

func (s *AddDataCheckTemplateRequestMetricRules) SetDataTypes(v string) *AddDataCheckTemplateRequestMetricRules {
	s.DataTypes = &v
	return s
}

func (s *AddDataCheckTemplateRequestMetricRules) SetDiffTolerateType(v int32) *AddDataCheckTemplateRequestMetricRules {
	s.DiffTolerateType = &v
	return s
}

func (s *AddDataCheckTemplateRequestMetricRules) SetDiffTolerateValues(v map[string]interface{}) *AddDataCheckTemplateRequestMetricRules {
	s.DiffTolerateValues = v
	return s
}

func (s *AddDataCheckTemplateRequestMetricRules) SetEnableDecimalScale(v int32) *AddDataCheckTemplateRequestMetricRules {
	s.EnableDecimalScale = &v
	return s
}

func (s *AddDataCheckTemplateRequestMetricRules) SetFilterColumnName(v string) *AddDataCheckTemplateRequestMetricRules {
	s.FilterColumnName = &v
	return s
}

func (s *AddDataCheckTemplateRequestMetricRules) SetFilterColumns(v string) *AddDataCheckTemplateRequestMetricRules {
	s.FilterColumns = &v
	return s
}

func (s *AddDataCheckTemplateRequestMetricRules) SetFloatPrecision(v int32) *AddDataCheckTemplateRequestMetricRules {
	s.FloatPrecision = &v
	return s
}

func (s *AddDataCheckTemplateRequestMetricRules) SetIgnoreDecimalDiff(v int32) *AddDataCheckTemplateRequestMetricRules {
	s.IgnoreDecimalDiff = &v
	return s
}

func (s *AddDataCheckTemplateRequestMetricRules) SetIgnoreDecimalScaleSuffixZero(v int32) *AddDataCheckTemplateRequestMetricRules {
	s.IgnoreDecimalScaleSuffixZero = &v
	return s
}

func (s *AddDataCheckTemplateRequestMetricRules) SetIgnoreEmptyDiff(v int32) *AddDataCheckTemplateRequestMetricRules {
	s.IgnoreEmptyDiff = &v
	return s
}

func (s *AddDataCheckTemplateRequestMetricRules) SetIgnoreNumericZero(v int32) *AddDataCheckTemplateRequestMetricRules {
	s.IgnoreNumericZero = &v
	return s
}

func (s *AddDataCheckTemplateRequestMetricRules) SetIgnoreStringEmpty(v int32) *AddDataCheckTemplateRequestMetricRules {
	s.IgnoreStringEmpty = &v
	return s
}

func (s *AddDataCheckTemplateRequestMetricRules) SetIgnoreZeroDiff(v int32) *AddDataCheckTemplateRequestMetricRules {
	s.IgnoreZeroDiff = &v
	return s
}

func (s *AddDataCheckTemplateRequestMetricRules) SetIsCountCheck(v int32) *AddDataCheckTemplateRequestMetricRules {
	s.IsCountCheck = &v
	return s
}

func (s *AddDataCheckTemplateRequestMetricRules) SetRuleId(v string) *AddDataCheckTemplateRequestMetricRules {
	s.RuleId = &v
	return s
}

func (s *AddDataCheckTemplateRequestMetricRules) SetSetDecimalScale(v int32) *AddDataCheckTemplateRequestMetricRules {
	s.SetDecimalScale = &v
	return s
}

func (s *AddDataCheckTemplateRequestMetricRules) Validate() error {
	return dara.Validate(s)
}

type AddDataCheckTemplateRequestNullRules struct {
	// The data type group that identifies the data type category to which the verification rule applies. Valid values: integers from 0 to 7. For the description of each value, see the enumeration values.
	//
	// example:
	//
	// 0
	DataTypeGroup *int32 `json:"dataTypeGroup,omitempty" xml:"dataTypeGroup,omitempty"`
	// The null value definitions, stored in JSON format.
	//
	// example:
	//
	// {}
	NullValues *string `json:"nullValues,omitempty" xml:"nullValues,omitempty"`
	// The rule ID that uniquely identifies a verification rule.
	//
	// example:
	//
	// 1001
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
}

func (s AddDataCheckTemplateRequestNullRules) String() string {
	return dara.Prettify(s)
}

func (s AddDataCheckTemplateRequestNullRules) GoString() string {
	return s.String()
}

func (s *AddDataCheckTemplateRequestNullRules) GetDataTypeGroup() *int32 {
	return s.DataTypeGroup
}

func (s *AddDataCheckTemplateRequestNullRules) GetNullValues() *string {
	return s.NullValues
}

func (s *AddDataCheckTemplateRequestNullRules) GetRuleId() *string {
	return s.RuleId
}

func (s *AddDataCheckTemplateRequestNullRules) SetDataTypeGroup(v int32) *AddDataCheckTemplateRequestNullRules {
	s.DataTypeGroup = &v
	return s
}

func (s *AddDataCheckTemplateRequestNullRules) SetNullValues(v string) *AddDataCheckTemplateRequestNullRules {
	s.NullValues = &v
	return s
}

func (s *AddDataCheckTemplateRequestNullRules) SetRuleId(v string) *AddDataCheckTemplateRequestNullRules {
	s.RuleId = &v
	return s
}

func (s *AddDataCheckTemplateRequestNullRules) Validate() error {
	return dara.Validate(s)
}

type AddDataCheckTemplateRequestWeakContentRule struct {
	// The filter column name expression.
	//
	// example:
	//
	// ^col_.*$
	FilterColumnExpression *string `json:"filterColumnExpression,omitempty" xml:"filterColumnExpression,omitempty"`
	// The filter column types, separated by vertical bars (|).
	FilterColumnTypes []*string `json:"filterColumnTypes,omitempty" xml:"filterColumnTypes,omitempty" type:"Repeated"`
	// The rule ID that uniquely identifies a verification rule.
	//
	// example:
	//
	// 1001
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// The weak content algorithm name: md5 or crc32.
	//
	// example:
	//
	// md5
	WeakContentAlgorithm *string `json:"weakContentAlgorithm,omitempty" xml:"weakContentAlgorithm,omitempty"`
}

func (s AddDataCheckTemplateRequestWeakContentRule) String() string {
	return dara.Prettify(s)
}

func (s AddDataCheckTemplateRequestWeakContentRule) GoString() string {
	return s.String()
}

func (s *AddDataCheckTemplateRequestWeakContentRule) GetFilterColumnExpression() *string {
	return s.FilterColumnExpression
}

func (s *AddDataCheckTemplateRequestWeakContentRule) GetFilterColumnTypes() []*string {
	return s.FilterColumnTypes
}

func (s *AddDataCheckTemplateRequestWeakContentRule) GetRuleId() *string {
	return s.RuleId
}

func (s *AddDataCheckTemplateRequestWeakContentRule) GetWeakContentAlgorithm() *string {
	return s.WeakContentAlgorithm
}

func (s *AddDataCheckTemplateRequestWeakContentRule) SetFilterColumnExpression(v string) *AddDataCheckTemplateRequestWeakContentRule {
	s.FilterColumnExpression = &v
	return s
}

func (s *AddDataCheckTemplateRequestWeakContentRule) SetFilterColumnTypes(v []*string) *AddDataCheckTemplateRequestWeakContentRule {
	s.FilterColumnTypes = v
	return s
}

func (s *AddDataCheckTemplateRequestWeakContentRule) SetRuleId(v string) *AddDataCheckTemplateRequestWeakContentRule {
	s.RuleId = &v
	return s
}

func (s *AddDataCheckTemplateRequestWeakContentRule) SetWeakContentAlgorithm(v string) *AddDataCheckTemplateRequestWeakContentRule {
	s.WeakContentAlgorithm = &v
	return s
}

func (s *AddDataCheckTemplateRequestWeakContentRule) Validate() error {
	return dara.Validate(s)
}
