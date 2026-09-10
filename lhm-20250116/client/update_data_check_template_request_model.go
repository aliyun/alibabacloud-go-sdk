// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataCheckTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBasicMetricRules(v []*UpdateDataCheckTemplateRequestBasicMetricRules) *UpdateDataCheckTemplateRequest
	GetBasicMetricRules() []*UpdateDataCheckTemplateRequestBasicMetricRules
	SetCheckType(v int32) *UpdateDataCheckTemplateRequest
	GetCheckType() *int32
	SetComplexMetricRules(v []*UpdateDataCheckTemplateRequestComplexMetricRules) *UpdateDataCheckTemplateRequest
	GetComplexMetricRules() []*UpdateDataCheckTemplateRequestComplexMetricRules
	SetDsEngineRels(v []*UpdateDataCheckTemplateRequestDsEngineRels) *UpdateDataCheckTemplateRequest
	GetDsEngineRels() []*UpdateDataCheckTemplateRequestDsEngineRels
	SetFulltextRule(v *UpdateDataCheckTemplateRequestFulltextRule) *UpdateDataCheckTemplateRequest
	GetFulltextRule() *UpdateDataCheckTemplateRequestFulltextRule
	SetMetricRules(v []*UpdateDataCheckTemplateRequestMetricRules) *UpdateDataCheckTemplateRequest
	GetMetricRules() []*UpdateDataCheckTemplateRequestMetricRules
	SetNullRules(v []*UpdateDataCheckTemplateRequestNullRules) *UpdateDataCheckTemplateRequest
	GetNullRules() []*UpdateDataCheckTemplateRequestNullRules
	SetRequestId(v string) *UpdateDataCheckTemplateRequest
	GetRequestId() *string
	SetTemplateDesc(v string) *UpdateDataCheckTemplateRequest
	GetTemplateDesc() *string
	SetTemplateId(v string) *UpdateDataCheckTemplateRequest
	GetTemplateId() *string
	SetTemplateName(v string) *UpdateDataCheckTemplateRequest
	GetTemplateName() *string
	SetWeakContentRule(v *UpdateDataCheckTemplateRequestWeakContentRule) *UpdateDataCheckTemplateRequest
	GetWeakContentRule() *UpdateDataCheckTemplateRequestWeakContentRule
}

type UpdateDataCheckTemplateRequest struct {
	BasicMetricRules []*UpdateDataCheckTemplateRequestBasicMetricRules `json:"basicMetricRules,omitempty" xml:"basicMetricRules,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	CheckType          *int32                                              `json:"checkType,omitempty" xml:"checkType,omitempty"`
	ComplexMetricRules []*UpdateDataCheckTemplateRequestComplexMetricRules `json:"complexMetricRules,omitempty" xml:"complexMetricRules,omitempty" type:"Repeated"`
	DsEngineRels       []*UpdateDataCheckTemplateRequestDsEngineRels       `json:"dsEngineRels,omitempty" xml:"dsEngineRels,omitempty" type:"Repeated"`
	FulltextRule       *UpdateDataCheckTemplateRequestFulltextRule         `json:"fulltextRule,omitempty" xml:"fulltextRule,omitempty" type:"Struct"`
	MetricRules        []*UpdateDataCheckTemplateRequestMetricRules        `json:"metricRules,omitempty" xml:"metricRules,omitempty" type:"Repeated"`
	NullRules          []*UpdateDataCheckTemplateRequestNullRules          `json:"nullRules,omitempty" xml:"nullRules,omitempty" type:"Repeated"`
	// example:
	//
	// 4C467B38-3910-4477-9B0B-6963D83B4E72
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 数据量校验模板描述
	TemplateDesc *string `json:"templateDesc,omitempty" xml:"templateDesc,omitempty"`
	// example:
	//
	// 1001
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// example:
	//
	// 数据量校验模板
	TemplateName    *string                                        `json:"templateName,omitempty" xml:"templateName,omitempty"`
	WeakContentRule *UpdateDataCheckTemplateRequestWeakContentRule `json:"weakContentRule,omitempty" xml:"weakContentRule,omitempty" type:"Struct"`
}

func (s UpdateDataCheckTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataCheckTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataCheckTemplateRequest) GetBasicMetricRules() []*UpdateDataCheckTemplateRequestBasicMetricRules {
	return s.BasicMetricRules
}

func (s *UpdateDataCheckTemplateRequest) GetCheckType() *int32 {
	return s.CheckType
}

func (s *UpdateDataCheckTemplateRequest) GetComplexMetricRules() []*UpdateDataCheckTemplateRequestComplexMetricRules {
	return s.ComplexMetricRules
}

func (s *UpdateDataCheckTemplateRequest) GetDsEngineRels() []*UpdateDataCheckTemplateRequestDsEngineRels {
	return s.DsEngineRels
}

func (s *UpdateDataCheckTemplateRequest) GetFulltextRule() *UpdateDataCheckTemplateRequestFulltextRule {
	return s.FulltextRule
}

func (s *UpdateDataCheckTemplateRequest) GetMetricRules() []*UpdateDataCheckTemplateRequestMetricRules {
	return s.MetricRules
}

func (s *UpdateDataCheckTemplateRequest) GetNullRules() []*UpdateDataCheckTemplateRequestNullRules {
	return s.NullRules
}

func (s *UpdateDataCheckTemplateRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataCheckTemplateRequest) GetTemplateDesc() *string {
	return s.TemplateDesc
}

func (s *UpdateDataCheckTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateDataCheckTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *UpdateDataCheckTemplateRequest) GetWeakContentRule() *UpdateDataCheckTemplateRequestWeakContentRule {
	return s.WeakContentRule
}

func (s *UpdateDataCheckTemplateRequest) SetBasicMetricRules(v []*UpdateDataCheckTemplateRequestBasicMetricRules) *UpdateDataCheckTemplateRequest {
	s.BasicMetricRules = v
	return s
}

func (s *UpdateDataCheckTemplateRequest) SetCheckType(v int32) *UpdateDataCheckTemplateRequest {
	s.CheckType = &v
	return s
}

func (s *UpdateDataCheckTemplateRequest) SetComplexMetricRules(v []*UpdateDataCheckTemplateRequestComplexMetricRules) *UpdateDataCheckTemplateRequest {
	s.ComplexMetricRules = v
	return s
}

func (s *UpdateDataCheckTemplateRequest) SetDsEngineRels(v []*UpdateDataCheckTemplateRequestDsEngineRels) *UpdateDataCheckTemplateRequest {
	s.DsEngineRels = v
	return s
}

func (s *UpdateDataCheckTemplateRequest) SetFulltextRule(v *UpdateDataCheckTemplateRequestFulltextRule) *UpdateDataCheckTemplateRequest {
	s.FulltextRule = v
	return s
}

func (s *UpdateDataCheckTemplateRequest) SetMetricRules(v []*UpdateDataCheckTemplateRequestMetricRules) *UpdateDataCheckTemplateRequest {
	s.MetricRules = v
	return s
}

func (s *UpdateDataCheckTemplateRequest) SetNullRules(v []*UpdateDataCheckTemplateRequestNullRules) *UpdateDataCheckTemplateRequest {
	s.NullRules = v
	return s
}

func (s *UpdateDataCheckTemplateRequest) SetRequestId(v string) *UpdateDataCheckTemplateRequest {
	s.RequestId = &v
	return s
}

func (s *UpdateDataCheckTemplateRequest) SetTemplateDesc(v string) *UpdateDataCheckTemplateRequest {
	s.TemplateDesc = &v
	return s
}

func (s *UpdateDataCheckTemplateRequest) SetTemplateId(v string) *UpdateDataCheckTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateDataCheckTemplateRequest) SetTemplateName(v string) *UpdateDataCheckTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *UpdateDataCheckTemplateRequest) SetWeakContentRule(v *UpdateDataCheckTemplateRequestWeakContentRule) *UpdateDataCheckTemplateRequest {
	s.WeakContentRule = v
	return s
}

func (s *UpdateDataCheckTemplateRequest) Validate() error {
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

type UpdateDataCheckTemplateRequestBasicMetricRules struct {
	// example:
	//
	// SUM,AVG
	CheckMethods *string `json:"checkMethods,omitempty" xml:"checkMethods,omitempty"`
	// example:
	//
	// 0
	ControlFloatPrecision *int32 `json:"controlFloatPrecision,omitempty" xml:"controlFloatPrecision,omitempty"`
	// example:
	//
	// 0
	DataTypeClassify *int32 `json:"dataTypeClassify,omitempty" xml:"dataTypeClassify,omitempty"`
	// example:
	//
	// 0
	DataTypeGroup *int32    `json:"dataTypeGroup,omitempty" xml:"dataTypeGroup,omitempty"`
	DataTypeList  []*string `json:"dataTypeList,omitempty" xml:"dataTypeList,omitempty" type:"Repeated"`
	// example:
	//
	// BIGINT
	DataTypes *string `json:"dataTypes,omitempty" xml:"dataTypes,omitempty"`
	// example:
	//
	// 0
	DiffTolerateType   *int32                 `json:"diffTolerateType,omitempty" xml:"diffTolerateType,omitempty"`
	DiffTolerateValues map[string]interface{} `json:"diffTolerateValues,omitempty" xml:"diffTolerateValues,omitempty"`
	// example:
	//
	// 0
	EnableDecimalScale *int32 `json:"enableDecimalScale,omitempty" xml:"enableDecimalScale,omitempty"`
	// example:
	//
	// col_a,col_b
	FilterColumnName *string `json:"filterColumnName,omitempty" xml:"filterColumnName,omitempty"`
	// example:
	//
	// col_a,col_b
	FilterColumns *string `json:"filterColumns,omitempty" xml:"filterColumns,omitempty"`
	// example:
	//
	// 2
	FloatPrecision *int32 `json:"floatPrecision,omitempty" xml:"floatPrecision,omitempty"`
	// example:
	//
	// 0
	IgnoreDecimalDiff *int32 `json:"ignoreDecimalDiff,omitempty" xml:"ignoreDecimalDiff,omitempty"`
	// example:
	//
	// 0
	IgnoreDecimalScaleSuffixZero *int32 `json:"ignoreDecimalScaleSuffixZero,omitempty" xml:"ignoreDecimalScaleSuffixZero,omitempty"`
	// example:
	//
	// 0
	IgnoreEmptyDiff *int32 `json:"ignoreEmptyDiff,omitempty" xml:"ignoreEmptyDiff,omitempty"`
	// example:
	//
	// 0
	IgnoreNumericZero *int32 `json:"ignoreNumericZero,omitempty" xml:"ignoreNumericZero,omitempty"`
	// example:
	//
	// 0
	IgnoreStringEmpty *int32 `json:"ignoreStringEmpty,omitempty" xml:"ignoreStringEmpty,omitempty"`
	// example:
	//
	// 0
	IgnoreZeroDiff *int32 `json:"ignoreZeroDiff,omitempty" xml:"ignoreZeroDiff,omitempty"`
	// example:
	//
	// 1
	IsCountCheck *int32 `json:"isCountCheck,omitempty" xml:"isCountCheck,omitempty"`
	// example:
	//
	// 1001
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// example:
	//
	// 2
	SetDecimalScale *int32 `json:"setDecimalScale,omitempty" xml:"setDecimalScale,omitempty"`
}

func (s UpdateDataCheckTemplateRequestBasicMetricRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataCheckTemplateRequestBasicMetricRules) GoString() string {
	return s.String()
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) GetCheckMethods() *string {
	return s.CheckMethods
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) GetControlFloatPrecision() *int32 {
	return s.ControlFloatPrecision
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) GetDataTypeClassify() *int32 {
	return s.DataTypeClassify
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) GetDataTypeGroup() *int32 {
	return s.DataTypeGroup
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) GetDataTypeList() []*string {
	return s.DataTypeList
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) GetDataTypes() *string {
	return s.DataTypes
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) GetDiffTolerateType() *int32 {
	return s.DiffTolerateType
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) GetDiffTolerateValues() map[string]interface{} {
	return s.DiffTolerateValues
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) GetEnableDecimalScale() *int32 {
	return s.EnableDecimalScale
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) GetFilterColumnName() *string {
	return s.FilterColumnName
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) GetFilterColumns() *string {
	return s.FilterColumns
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) GetFloatPrecision() *int32 {
	return s.FloatPrecision
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) GetIgnoreDecimalDiff() *int32 {
	return s.IgnoreDecimalDiff
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) GetIgnoreDecimalScaleSuffixZero() *int32 {
	return s.IgnoreDecimalScaleSuffixZero
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) GetIgnoreEmptyDiff() *int32 {
	return s.IgnoreEmptyDiff
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) GetIgnoreNumericZero() *int32 {
	return s.IgnoreNumericZero
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) GetIgnoreStringEmpty() *int32 {
	return s.IgnoreStringEmpty
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) GetIgnoreZeroDiff() *int32 {
	return s.IgnoreZeroDiff
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) GetIsCountCheck() *int32 {
	return s.IsCountCheck
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) GetRuleId() *string {
	return s.RuleId
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) GetSetDecimalScale() *int32 {
	return s.SetDecimalScale
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) SetCheckMethods(v string) *UpdateDataCheckTemplateRequestBasicMetricRules {
	s.CheckMethods = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) SetControlFloatPrecision(v int32) *UpdateDataCheckTemplateRequestBasicMetricRules {
	s.ControlFloatPrecision = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) SetDataTypeClassify(v int32) *UpdateDataCheckTemplateRequestBasicMetricRules {
	s.DataTypeClassify = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) SetDataTypeGroup(v int32) *UpdateDataCheckTemplateRequestBasicMetricRules {
	s.DataTypeGroup = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) SetDataTypeList(v []*string) *UpdateDataCheckTemplateRequestBasicMetricRules {
	s.DataTypeList = v
	return s
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) SetDataTypes(v string) *UpdateDataCheckTemplateRequestBasicMetricRules {
	s.DataTypes = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) SetDiffTolerateType(v int32) *UpdateDataCheckTemplateRequestBasicMetricRules {
	s.DiffTolerateType = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) SetDiffTolerateValues(v map[string]interface{}) *UpdateDataCheckTemplateRequestBasicMetricRules {
	s.DiffTolerateValues = v
	return s
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) SetEnableDecimalScale(v int32) *UpdateDataCheckTemplateRequestBasicMetricRules {
	s.EnableDecimalScale = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) SetFilterColumnName(v string) *UpdateDataCheckTemplateRequestBasicMetricRules {
	s.FilterColumnName = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) SetFilterColumns(v string) *UpdateDataCheckTemplateRequestBasicMetricRules {
	s.FilterColumns = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) SetFloatPrecision(v int32) *UpdateDataCheckTemplateRequestBasicMetricRules {
	s.FloatPrecision = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) SetIgnoreDecimalDiff(v int32) *UpdateDataCheckTemplateRequestBasicMetricRules {
	s.IgnoreDecimalDiff = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) SetIgnoreDecimalScaleSuffixZero(v int32) *UpdateDataCheckTemplateRequestBasicMetricRules {
	s.IgnoreDecimalScaleSuffixZero = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) SetIgnoreEmptyDiff(v int32) *UpdateDataCheckTemplateRequestBasicMetricRules {
	s.IgnoreEmptyDiff = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) SetIgnoreNumericZero(v int32) *UpdateDataCheckTemplateRequestBasicMetricRules {
	s.IgnoreNumericZero = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) SetIgnoreStringEmpty(v int32) *UpdateDataCheckTemplateRequestBasicMetricRules {
	s.IgnoreStringEmpty = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) SetIgnoreZeroDiff(v int32) *UpdateDataCheckTemplateRequestBasicMetricRules {
	s.IgnoreZeroDiff = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) SetIsCountCheck(v int32) *UpdateDataCheckTemplateRequestBasicMetricRules {
	s.IsCountCheck = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) SetRuleId(v string) *UpdateDataCheckTemplateRequestBasicMetricRules {
	s.RuleId = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) SetSetDecimalScale(v int32) *UpdateDataCheckTemplateRequestBasicMetricRules {
	s.SetDecimalScale = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestBasicMetricRules) Validate() error {
	return dara.Validate(s)
}

type UpdateDataCheckTemplateRequestComplexMetricRules struct {
	// example:
	//
	// SUM,AVG
	CheckMethods *string `json:"checkMethods,omitempty" xml:"checkMethods,omitempty"`
	// example:
	//
	// 0
	ControlFloatPrecision *int32 `json:"controlFloatPrecision,omitempty" xml:"controlFloatPrecision,omitempty"`
	// example:
	//
	// 0
	DataTypeClassify *int32 `json:"dataTypeClassify,omitempty" xml:"dataTypeClassify,omitempty"`
	// example:
	//
	// 0
	DataTypeGroup *int32    `json:"dataTypeGroup,omitempty" xml:"dataTypeGroup,omitempty"`
	DataTypeList  []*string `json:"dataTypeList,omitempty" xml:"dataTypeList,omitempty" type:"Repeated"`
	// example:
	//
	// BIGINT
	DataTypes *string `json:"dataTypes,omitempty" xml:"dataTypes,omitempty"`
	// example:
	//
	// 0
	DiffTolerateType   *int32                 `json:"diffTolerateType,omitempty" xml:"diffTolerateType,omitempty"`
	DiffTolerateValues map[string]interface{} `json:"diffTolerateValues,omitempty" xml:"diffTolerateValues,omitempty"`
	// example:
	//
	// 0
	EnableDecimalScale *int32 `json:"enableDecimalScale,omitempty" xml:"enableDecimalScale,omitempty"`
	// example:
	//
	// col_a,col_b
	FilterColumnName *string `json:"filterColumnName,omitempty" xml:"filterColumnName,omitempty"`
	// example:
	//
	// col_a,col_b
	FilterColumns *string `json:"filterColumns,omitempty" xml:"filterColumns,omitempty"`
	// example:
	//
	// 2
	FloatPrecision *int32 `json:"floatPrecision,omitempty" xml:"floatPrecision,omitempty"`
	// example:
	//
	// 0
	IgnoreDecimalDiff *int32 `json:"ignoreDecimalDiff,omitempty" xml:"ignoreDecimalDiff,omitempty"`
	// example:
	//
	// 0
	IgnoreDecimalScaleSuffixZero *int32 `json:"ignoreDecimalScaleSuffixZero,omitempty" xml:"ignoreDecimalScaleSuffixZero,omitempty"`
	// example:
	//
	// 0
	IgnoreEmptyDiff *int32 `json:"ignoreEmptyDiff,omitempty" xml:"ignoreEmptyDiff,omitempty"`
	// example:
	//
	// 0
	IgnoreNumericZero *int32 `json:"ignoreNumericZero,omitempty" xml:"ignoreNumericZero,omitempty"`
	// example:
	//
	// 0
	IgnoreStringEmpty *int32 `json:"ignoreStringEmpty,omitempty" xml:"ignoreStringEmpty,omitempty"`
	// example:
	//
	// 0
	IgnoreZeroDiff *int32 `json:"ignoreZeroDiff,omitempty" xml:"ignoreZeroDiff,omitempty"`
	// example:
	//
	// 1
	IsCountCheck *int32 `json:"isCountCheck,omitempty" xml:"isCountCheck,omitempty"`
	// example:
	//
	// 1001
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// example:
	//
	// 2
	SetDecimalScale *int32 `json:"setDecimalScale,omitempty" xml:"setDecimalScale,omitempty"`
}

func (s UpdateDataCheckTemplateRequestComplexMetricRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataCheckTemplateRequestComplexMetricRules) GoString() string {
	return s.String()
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) GetCheckMethods() *string {
	return s.CheckMethods
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) GetControlFloatPrecision() *int32 {
	return s.ControlFloatPrecision
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) GetDataTypeClassify() *int32 {
	return s.DataTypeClassify
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) GetDataTypeGroup() *int32 {
	return s.DataTypeGroup
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) GetDataTypeList() []*string {
	return s.DataTypeList
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) GetDataTypes() *string {
	return s.DataTypes
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) GetDiffTolerateType() *int32 {
	return s.DiffTolerateType
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) GetDiffTolerateValues() map[string]interface{} {
	return s.DiffTolerateValues
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) GetEnableDecimalScale() *int32 {
	return s.EnableDecimalScale
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) GetFilterColumnName() *string {
	return s.FilterColumnName
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) GetFilterColumns() *string {
	return s.FilterColumns
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) GetFloatPrecision() *int32 {
	return s.FloatPrecision
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) GetIgnoreDecimalDiff() *int32 {
	return s.IgnoreDecimalDiff
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) GetIgnoreDecimalScaleSuffixZero() *int32 {
	return s.IgnoreDecimalScaleSuffixZero
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) GetIgnoreEmptyDiff() *int32 {
	return s.IgnoreEmptyDiff
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) GetIgnoreNumericZero() *int32 {
	return s.IgnoreNumericZero
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) GetIgnoreStringEmpty() *int32 {
	return s.IgnoreStringEmpty
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) GetIgnoreZeroDiff() *int32 {
	return s.IgnoreZeroDiff
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) GetIsCountCheck() *int32 {
	return s.IsCountCheck
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) GetRuleId() *string {
	return s.RuleId
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) GetSetDecimalScale() *int32 {
	return s.SetDecimalScale
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) SetCheckMethods(v string) *UpdateDataCheckTemplateRequestComplexMetricRules {
	s.CheckMethods = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) SetControlFloatPrecision(v int32) *UpdateDataCheckTemplateRequestComplexMetricRules {
	s.ControlFloatPrecision = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) SetDataTypeClassify(v int32) *UpdateDataCheckTemplateRequestComplexMetricRules {
	s.DataTypeClassify = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) SetDataTypeGroup(v int32) *UpdateDataCheckTemplateRequestComplexMetricRules {
	s.DataTypeGroup = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) SetDataTypeList(v []*string) *UpdateDataCheckTemplateRequestComplexMetricRules {
	s.DataTypeList = v
	return s
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) SetDataTypes(v string) *UpdateDataCheckTemplateRequestComplexMetricRules {
	s.DataTypes = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) SetDiffTolerateType(v int32) *UpdateDataCheckTemplateRequestComplexMetricRules {
	s.DiffTolerateType = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) SetDiffTolerateValues(v map[string]interface{}) *UpdateDataCheckTemplateRequestComplexMetricRules {
	s.DiffTolerateValues = v
	return s
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) SetEnableDecimalScale(v int32) *UpdateDataCheckTemplateRequestComplexMetricRules {
	s.EnableDecimalScale = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) SetFilterColumnName(v string) *UpdateDataCheckTemplateRequestComplexMetricRules {
	s.FilterColumnName = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) SetFilterColumns(v string) *UpdateDataCheckTemplateRequestComplexMetricRules {
	s.FilterColumns = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) SetFloatPrecision(v int32) *UpdateDataCheckTemplateRequestComplexMetricRules {
	s.FloatPrecision = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) SetIgnoreDecimalDiff(v int32) *UpdateDataCheckTemplateRequestComplexMetricRules {
	s.IgnoreDecimalDiff = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) SetIgnoreDecimalScaleSuffixZero(v int32) *UpdateDataCheckTemplateRequestComplexMetricRules {
	s.IgnoreDecimalScaleSuffixZero = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) SetIgnoreEmptyDiff(v int32) *UpdateDataCheckTemplateRequestComplexMetricRules {
	s.IgnoreEmptyDiff = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) SetIgnoreNumericZero(v int32) *UpdateDataCheckTemplateRequestComplexMetricRules {
	s.IgnoreNumericZero = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) SetIgnoreStringEmpty(v int32) *UpdateDataCheckTemplateRequestComplexMetricRules {
	s.IgnoreStringEmpty = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) SetIgnoreZeroDiff(v int32) *UpdateDataCheckTemplateRequestComplexMetricRules {
	s.IgnoreZeroDiff = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) SetIsCountCheck(v int32) *UpdateDataCheckTemplateRequestComplexMetricRules {
	s.IsCountCheck = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) SetRuleId(v string) *UpdateDataCheckTemplateRequestComplexMetricRules {
	s.RuleId = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) SetSetDecimalScale(v int32) *UpdateDataCheckTemplateRequestComplexMetricRules {
	s.SetDecimalScale = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestComplexMetricRules) Validate() error {
	return dara.Validate(s)
}

type UpdateDataCheckTemplateRequestDsEngineRels struct {
	// example:
	//
	// 1001
	DsEngineId *string `json:"dsEngineId,omitempty" xml:"dsEngineId,omitempty"`
	// example:
	//
	// Hive
	DsType      *string   `json:"dsType,omitempty" xml:"dsType,omitempty"`
	EngineTypes []*string `json:"engineTypes,omitempty" xml:"engineTypes,omitempty" type:"Repeated"`
}

func (s UpdateDataCheckTemplateRequestDsEngineRels) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataCheckTemplateRequestDsEngineRels) GoString() string {
	return s.String()
}

func (s *UpdateDataCheckTemplateRequestDsEngineRels) GetDsEngineId() *string {
	return s.DsEngineId
}

func (s *UpdateDataCheckTemplateRequestDsEngineRels) GetDsType() *string {
	return s.DsType
}

func (s *UpdateDataCheckTemplateRequestDsEngineRels) GetEngineTypes() []*string {
	return s.EngineTypes
}

func (s *UpdateDataCheckTemplateRequestDsEngineRels) SetDsEngineId(v string) *UpdateDataCheckTemplateRequestDsEngineRels {
	s.DsEngineId = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestDsEngineRels) SetDsType(v string) *UpdateDataCheckTemplateRequestDsEngineRels {
	s.DsType = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestDsEngineRels) SetEngineTypes(v []*string) *UpdateDataCheckTemplateRequestDsEngineRels {
	s.EngineTypes = v
	return s
}

func (s *UpdateDataCheckTemplateRequestDsEngineRels) Validate() error {
	return dara.Validate(s)
}

type UpdateDataCheckTemplateRequestFulltextRule struct {
	// example:
	//
	// 0
	CheckMode *int32 `json:"checkMode,omitempty" xml:"checkMode,omitempty"`
	// example:
	//
	// 0
	ColumnEqualCmpType *int32 `json:"columnEqualCmpType,omitempty" xml:"columnEqualCmpType,omitempty"`
	// example:
	//
	// ARRAY,MAP
	ColumnEqualCmpValues *string `json:"columnEqualCmpValues,omitempty" xml:"columnEqualCmpValues,omitempty"`
	// example:
	//
	// 0
	ColumnIsCosine *int32 `json:"columnIsCosine,omitempty" xml:"columnIsCosine,omitempty"`
	// example:
	//
	// 0
	ColumnIsIgnoreNull *int32 `json:"columnIsIgnoreNull,omitempty" xml:"columnIsIgnoreNull,omitempty"`
	// example:
	//
	// 0
	ColumnIsIgnoreZero *int32 `json:"columnIsIgnoreZero,omitempty" xml:"columnIsIgnoreZero,omitempty"`
	// example:
	//
	// 0
	ColumnIsSamples *int32 `json:"columnIsSamples,omitempty" xml:"columnIsSamples,omitempty"`
	// example:
	//
	// 0
	ColumnSamplesType *int32 `json:"columnSamplesType,omitempty" xml:"columnSamplesType,omitempty"`
	// example:
	//
	// 100
	ColumnSamplesValue *int32 `json:"columnSamplesValue,omitempty" xml:"columnSamplesValue,omitempty"`
	// example:
	//
	// 0
	ColumnSizeCmpType *int32 `json:"columnSizeCmpType,omitempty" xml:"columnSizeCmpType,omitempty"`
	// example:
	//
	// ARRAY,MAP
	ColumnSizeCmpValues *string `json:"columnSizeCmpValues,omitempty" xml:"columnSizeCmpValues,omitempty"`
	// example:
	//
	// 1
	IsPrimaryKeyCheck *int32 `json:"isPrimaryKeyCheck,omitempty" xml:"isPrimaryKeyCheck,omitempty"`
	// example:
	//
	// 0
	LineCheckType *int32 `json:"lineCheckType,omitempty" xml:"lineCheckType,omitempty"`
	// example:
	//
	// 0
	LineIsPrintAll *int32 `json:"lineIsPrintAll,omitempty" xml:"lineIsPrintAll,omitempty"`
	// example:
	//
	// 0
	LineIsSamples *int32 `json:"lineIsSamples,omitempty" xml:"lineIsSamples,omitempty"`
	// example:
	//
	// 0
	LineSamplesType *int32 `json:"lineSamplesType,omitempty" xml:"lineSamplesType,omitempty"`
	// example:
	//
	// 100
	LineSamplesValue *int32 `json:"lineSamplesValue,omitempty" xml:"lineSamplesValue,omitempty"`
	// example:
	//
	// 1001
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
}

func (s UpdateDataCheckTemplateRequestFulltextRule) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataCheckTemplateRequestFulltextRule) GoString() string {
	return s.String()
}

func (s *UpdateDataCheckTemplateRequestFulltextRule) GetCheckMode() *int32 {
	return s.CheckMode
}

func (s *UpdateDataCheckTemplateRequestFulltextRule) GetColumnEqualCmpType() *int32 {
	return s.ColumnEqualCmpType
}

func (s *UpdateDataCheckTemplateRequestFulltextRule) GetColumnEqualCmpValues() *string {
	return s.ColumnEqualCmpValues
}

func (s *UpdateDataCheckTemplateRequestFulltextRule) GetColumnIsCosine() *int32 {
	return s.ColumnIsCosine
}

func (s *UpdateDataCheckTemplateRequestFulltextRule) GetColumnIsIgnoreNull() *int32 {
	return s.ColumnIsIgnoreNull
}

func (s *UpdateDataCheckTemplateRequestFulltextRule) GetColumnIsIgnoreZero() *int32 {
	return s.ColumnIsIgnoreZero
}

func (s *UpdateDataCheckTemplateRequestFulltextRule) GetColumnIsSamples() *int32 {
	return s.ColumnIsSamples
}

func (s *UpdateDataCheckTemplateRequestFulltextRule) GetColumnSamplesType() *int32 {
	return s.ColumnSamplesType
}

func (s *UpdateDataCheckTemplateRequestFulltextRule) GetColumnSamplesValue() *int32 {
	return s.ColumnSamplesValue
}

func (s *UpdateDataCheckTemplateRequestFulltextRule) GetColumnSizeCmpType() *int32 {
	return s.ColumnSizeCmpType
}

func (s *UpdateDataCheckTemplateRequestFulltextRule) GetColumnSizeCmpValues() *string {
	return s.ColumnSizeCmpValues
}

func (s *UpdateDataCheckTemplateRequestFulltextRule) GetIsPrimaryKeyCheck() *int32 {
	return s.IsPrimaryKeyCheck
}

func (s *UpdateDataCheckTemplateRequestFulltextRule) GetLineCheckType() *int32 {
	return s.LineCheckType
}

func (s *UpdateDataCheckTemplateRequestFulltextRule) GetLineIsPrintAll() *int32 {
	return s.LineIsPrintAll
}

func (s *UpdateDataCheckTemplateRequestFulltextRule) GetLineIsSamples() *int32 {
	return s.LineIsSamples
}

func (s *UpdateDataCheckTemplateRequestFulltextRule) GetLineSamplesType() *int32 {
	return s.LineSamplesType
}

func (s *UpdateDataCheckTemplateRequestFulltextRule) GetLineSamplesValue() *int32 {
	return s.LineSamplesValue
}

func (s *UpdateDataCheckTemplateRequestFulltextRule) GetRuleId() *string {
	return s.RuleId
}

func (s *UpdateDataCheckTemplateRequestFulltextRule) SetCheckMode(v int32) *UpdateDataCheckTemplateRequestFulltextRule {
	s.CheckMode = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestFulltextRule) SetColumnEqualCmpType(v int32) *UpdateDataCheckTemplateRequestFulltextRule {
	s.ColumnEqualCmpType = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestFulltextRule) SetColumnEqualCmpValues(v string) *UpdateDataCheckTemplateRequestFulltextRule {
	s.ColumnEqualCmpValues = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestFulltextRule) SetColumnIsCosine(v int32) *UpdateDataCheckTemplateRequestFulltextRule {
	s.ColumnIsCosine = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestFulltextRule) SetColumnIsIgnoreNull(v int32) *UpdateDataCheckTemplateRequestFulltextRule {
	s.ColumnIsIgnoreNull = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestFulltextRule) SetColumnIsIgnoreZero(v int32) *UpdateDataCheckTemplateRequestFulltextRule {
	s.ColumnIsIgnoreZero = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestFulltextRule) SetColumnIsSamples(v int32) *UpdateDataCheckTemplateRequestFulltextRule {
	s.ColumnIsSamples = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestFulltextRule) SetColumnSamplesType(v int32) *UpdateDataCheckTemplateRequestFulltextRule {
	s.ColumnSamplesType = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestFulltextRule) SetColumnSamplesValue(v int32) *UpdateDataCheckTemplateRequestFulltextRule {
	s.ColumnSamplesValue = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestFulltextRule) SetColumnSizeCmpType(v int32) *UpdateDataCheckTemplateRequestFulltextRule {
	s.ColumnSizeCmpType = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestFulltextRule) SetColumnSizeCmpValues(v string) *UpdateDataCheckTemplateRequestFulltextRule {
	s.ColumnSizeCmpValues = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestFulltextRule) SetIsPrimaryKeyCheck(v int32) *UpdateDataCheckTemplateRequestFulltextRule {
	s.IsPrimaryKeyCheck = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestFulltextRule) SetLineCheckType(v int32) *UpdateDataCheckTemplateRequestFulltextRule {
	s.LineCheckType = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestFulltextRule) SetLineIsPrintAll(v int32) *UpdateDataCheckTemplateRequestFulltextRule {
	s.LineIsPrintAll = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestFulltextRule) SetLineIsSamples(v int32) *UpdateDataCheckTemplateRequestFulltextRule {
	s.LineIsSamples = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestFulltextRule) SetLineSamplesType(v int32) *UpdateDataCheckTemplateRequestFulltextRule {
	s.LineSamplesType = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestFulltextRule) SetLineSamplesValue(v int32) *UpdateDataCheckTemplateRequestFulltextRule {
	s.LineSamplesValue = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestFulltextRule) SetRuleId(v string) *UpdateDataCheckTemplateRequestFulltextRule {
	s.RuleId = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestFulltextRule) Validate() error {
	return dara.Validate(s)
}

type UpdateDataCheckTemplateRequestMetricRules struct {
	// example:
	//
	// SUM,AVG
	CheckMethods *string `json:"checkMethods,omitempty" xml:"checkMethods,omitempty"`
	// example:
	//
	// 0
	ControlFloatPrecision *int32 `json:"controlFloatPrecision,omitempty" xml:"controlFloatPrecision,omitempty"`
	// example:
	//
	// 0
	DataTypeClassify *int32 `json:"dataTypeClassify,omitempty" xml:"dataTypeClassify,omitempty"`
	// example:
	//
	// 0
	DataTypeGroup *int32    `json:"dataTypeGroup,omitempty" xml:"dataTypeGroup,omitempty"`
	DataTypeList  []*string `json:"dataTypeList,omitempty" xml:"dataTypeList,omitempty" type:"Repeated"`
	// example:
	//
	// BIGINT
	DataTypes *string `json:"dataTypes,omitempty" xml:"dataTypes,omitempty"`
	// example:
	//
	// 0
	DiffTolerateType   *int32                 `json:"diffTolerateType,omitempty" xml:"diffTolerateType,omitempty"`
	DiffTolerateValues map[string]interface{} `json:"diffTolerateValues,omitempty" xml:"diffTolerateValues,omitempty"`
	// example:
	//
	// 0
	EnableDecimalScale *int32 `json:"enableDecimalScale,omitempty" xml:"enableDecimalScale,omitempty"`
	// example:
	//
	// col_a,col_b
	FilterColumnName *string `json:"filterColumnName,omitempty" xml:"filterColumnName,omitempty"`
	// example:
	//
	// col_a,col_b
	FilterColumns *string `json:"filterColumns,omitempty" xml:"filterColumns,omitempty"`
	// example:
	//
	// 2
	FloatPrecision *int32 `json:"floatPrecision,omitempty" xml:"floatPrecision,omitempty"`
	// example:
	//
	// 0
	IgnoreDecimalDiff *int32 `json:"ignoreDecimalDiff,omitempty" xml:"ignoreDecimalDiff,omitempty"`
	// example:
	//
	// 0
	IgnoreDecimalScaleSuffixZero *int32 `json:"ignoreDecimalScaleSuffixZero,omitempty" xml:"ignoreDecimalScaleSuffixZero,omitempty"`
	// example:
	//
	// 0
	IgnoreEmptyDiff *int32 `json:"ignoreEmptyDiff,omitempty" xml:"ignoreEmptyDiff,omitempty"`
	// example:
	//
	// 0
	IgnoreNumericZero *int32 `json:"ignoreNumericZero,omitempty" xml:"ignoreNumericZero,omitempty"`
	// example:
	//
	// 0
	IgnoreStringEmpty *int32 `json:"ignoreStringEmpty,omitempty" xml:"ignoreStringEmpty,omitempty"`
	// example:
	//
	// 0
	IgnoreZeroDiff *int32 `json:"ignoreZeroDiff,omitempty" xml:"ignoreZeroDiff,omitempty"`
	// example:
	//
	// 1
	IsCountCheck *int32 `json:"isCountCheck,omitempty" xml:"isCountCheck,omitempty"`
	// example:
	//
	// 1001
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// example:
	//
	// 2
	SetDecimalScale *int32 `json:"setDecimalScale,omitempty" xml:"setDecimalScale,omitempty"`
}

func (s UpdateDataCheckTemplateRequestMetricRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataCheckTemplateRequestMetricRules) GoString() string {
	return s.String()
}

func (s *UpdateDataCheckTemplateRequestMetricRules) GetCheckMethods() *string {
	return s.CheckMethods
}

func (s *UpdateDataCheckTemplateRequestMetricRules) GetControlFloatPrecision() *int32 {
	return s.ControlFloatPrecision
}

func (s *UpdateDataCheckTemplateRequestMetricRules) GetDataTypeClassify() *int32 {
	return s.DataTypeClassify
}

func (s *UpdateDataCheckTemplateRequestMetricRules) GetDataTypeGroup() *int32 {
	return s.DataTypeGroup
}

func (s *UpdateDataCheckTemplateRequestMetricRules) GetDataTypeList() []*string {
	return s.DataTypeList
}

func (s *UpdateDataCheckTemplateRequestMetricRules) GetDataTypes() *string {
	return s.DataTypes
}

func (s *UpdateDataCheckTemplateRequestMetricRules) GetDiffTolerateType() *int32 {
	return s.DiffTolerateType
}

func (s *UpdateDataCheckTemplateRequestMetricRules) GetDiffTolerateValues() map[string]interface{} {
	return s.DiffTolerateValues
}

func (s *UpdateDataCheckTemplateRequestMetricRules) GetEnableDecimalScale() *int32 {
	return s.EnableDecimalScale
}

func (s *UpdateDataCheckTemplateRequestMetricRules) GetFilterColumnName() *string {
	return s.FilterColumnName
}

func (s *UpdateDataCheckTemplateRequestMetricRules) GetFilterColumns() *string {
	return s.FilterColumns
}

func (s *UpdateDataCheckTemplateRequestMetricRules) GetFloatPrecision() *int32 {
	return s.FloatPrecision
}

func (s *UpdateDataCheckTemplateRequestMetricRules) GetIgnoreDecimalDiff() *int32 {
	return s.IgnoreDecimalDiff
}

func (s *UpdateDataCheckTemplateRequestMetricRules) GetIgnoreDecimalScaleSuffixZero() *int32 {
	return s.IgnoreDecimalScaleSuffixZero
}

func (s *UpdateDataCheckTemplateRequestMetricRules) GetIgnoreEmptyDiff() *int32 {
	return s.IgnoreEmptyDiff
}

func (s *UpdateDataCheckTemplateRequestMetricRules) GetIgnoreNumericZero() *int32 {
	return s.IgnoreNumericZero
}

func (s *UpdateDataCheckTemplateRequestMetricRules) GetIgnoreStringEmpty() *int32 {
	return s.IgnoreStringEmpty
}

func (s *UpdateDataCheckTemplateRequestMetricRules) GetIgnoreZeroDiff() *int32 {
	return s.IgnoreZeroDiff
}

func (s *UpdateDataCheckTemplateRequestMetricRules) GetIsCountCheck() *int32 {
	return s.IsCountCheck
}

func (s *UpdateDataCheckTemplateRequestMetricRules) GetRuleId() *string {
	return s.RuleId
}

func (s *UpdateDataCheckTemplateRequestMetricRules) GetSetDecimalScale() *int32 {
	return s.SetDecimalScale
}

func (s *UpdateDataCheckTemplateRequestMetricRules) SetCheckMethods(v string) *UpdateDataCheckTemplateRequestMetricRules {
	s.CheckMethods = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestMetricRules) SetControlFloatPrecision(v int32) *UpdateDataCheckTemplateRequestMetricRules {
	s.ControlFloatPrecision = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestMetricRules) SetDataTypeClassify(v int32) *UpdateDataCheckTemplateRequestMetricRules {
	s.DataTypeClassify = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestMetricRules) SetDataTypeGroup(v int32) *UpdateDataCheckTemplateRequestMetricRules {
	s.DataTypeGroup = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestMetricRules) SetDataTypeList(v []*string) *UpdateDataCheckTemplateRequestMetricRules {
	s.DataTypeList = v
	return s
}

func (s *UpdateDataCheckTemplateRequestMetricRules) SetDataTypes(v string) *UpdateDataCheckTemplateRequestMetricRules {
	s.DataTypes = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestMetricRules) SetDiffTolerateType(v int32) *UpdateDataCheckTemplateRequestMetricRules {
	s.DiffTolerateType = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestMetricRules) SetDiffTolerateValues(v map[string]interface{}) *UpdateDataCheckTemplateRequestMetricRules {
	s.DiffTolerateValues = v
	return s
}

func (s *UpdateDataCheckTemplateRequestMetricRules) SetEnableDecimalScale(v int32) *UpdateDataCheckTemplateRequestMetricRules {
	s.EnableDecimalScale = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestMetricRules) SetFilterColumnName(v string) *UpdateDataCheckTemplateRequestMetricRules {
	s.FilterColumnName = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestMetricRules) SetFilterColumns(v string) *UpdateDataCheckTemplateRequestMetricRules {
	s.FilterColumns = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestMetricRules) SetFloatPrecision(v int32) *UpdateDataCheckTemplateRequestMetricRules {
	s.FloatPrecision = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestMetricRules) SetIgnoreDecimalDiff(v int32) *UpdateDataCheckTemplateRequestMetricRules {
	s.IgnoreDecimalDiff = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestMetricRules) SetIgnoreDecimalScaleSuffixZero(v int32) *UpdateDataCheckTemplateRequestMetricRules {
	s.IgnoreDecimalScaleSuffixZero = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestMetricRules) SetIgnoreEmptyDiff(v int32) *UpdateDataCheckTemplateRequestMetricRules {
	s.IgnoreEmptyDiff = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestMetricRules) SetIgnoreNumericZero(v int32) *UpdateDataCheckTemplateRequestMetricRules {
	s.IgnoreNumericZero = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestMetricRules) SetIgnoreStringEmpty(v int32) *UpdateDataCheckTemplateRequestMetricRules {
	s.IgnoreStringEmpty = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestMetricRules) SetIgnoreZeroDiff(v int32) *UpdateDataCheckTemplateRequestMetricRules {
	s.IgnoreZeroDiff = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestMetricRules) SetIsCountCheck(v int32) *UpdateDataCheckTemplateRequestMetricRules {
	s.IsCountCheck = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestMetricRules) SetRuleId(v string) *UpdateDataCheckTemplateRequestMetricRules {
	s.RuleId = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestMetricRules) SetSetDecimalScale(v int32) *UpdateDataCheckTemplateRequestMetricRules {
	s.SetDecimalScale = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestMetricRules) Validate() error {
	return dara.Validate(s)
}

type UpdateDataCheckTemplateRequestNullRules struct {
	// example:
	//
	// 0
	DataTypeGroup *int32 `json:"dataTypeGroup,omitempty" xml:"dataTypeGroup,omitempty"`
	// example:
	//
	// {}
	NullValues *string `json:"nullValues,omitempty" xml:"nullValues,omitempty"`
	// example:
	//
	// 1001
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
}

func (s UpdateDataCheckTemplateRequestNullRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataCheckTemplateRequestNullRules) GoString() string {
	return s.String()
}

func (s *UpdateDataCheckTemplateRequestNullRules) GetDataTypeGroup() *int32 {
	return s.DataTypeGroup
}

func (s *UpdateDataCheckTemplateRequestNullRules) GetNullValues() *string {
	return s.NullValues
}

func (s *UpdateDataCheckTemplateRequestNullRules) GetRuleId() *string {
	return s.RuleId
}

func (s *UpdateDataCheckTemplateRequestNullRules) SetDataTypeGroup(v int32) *UpdateDataCheckTemplateRequestNullRules {
	s.DataTypeGroup = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestNullRules) SetNullValues(v string) *UpdateDataCheckTemplateRequestNullRules {
	s.NullValues = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestNullRules) SetRuleId(v string) *UpdateDataCheckTemplateRequestNullRules {
	s.RuleId = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestNullRules) Validate() error {
	return dara.Validate(s)
}

type UpdateDataCheckTemplateRequestWeakContentRule struct {
	// example:
	//
	// ^col_.*$
	FilterColumnExpression *string   `json:"filterColumnExpression,omitempty" xml:"filterColumnExpression,omitempty"`
	FilterColumnTypes      []*string `json:"filterColumnTypes,omitempty" xml:"filterColumnTypes,omitempty" type:"Repeated"`
	// example:
	//
	// 1001
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// example:
	//
	// md5
	WeakContentAlgorithm *string `json:"weakContentAlgorithm,omitempty" xml:"weakContentAlgorithm,omitempty"`
}

func (s UpdateDataCheckTemplateRequestWeakContentRule) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataCheckTemplateRequestWeakContentRule) GoString() string {
	return s.String()
}

func (s *UpdateDataCheckTemplateRequestWeakContentRule) GetFilterColumnExpression() *string {
	return s.FilterColumnExpression
}

func (s *UpdateDataCheckTemplateRequestWeakContentRule) GetFilterColumnTypes() []*string {
	return s.FilterColumnTypes
}

func (s *UpdateDataCheckTemplateRequestWeakContentRule) GetRuleId() *string {
	return s.RuleId
}

func (s *UpdateDataCheckTemplateRequestWeakContentRule) GetWeakContentAlgorithm() *string {
	return s.WeakContentAlgorithm
}

func (s *UpdateDataCheckTemplateRequestWeakContentRule) SetFilterColumnExpression(v string) *UpdateDataCheckTemplateRequestWeakContentRule {
	s.FilterColumnExpression = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestWeakContentRule) SetFilterColumnTypes(v []*string) *UpdateDataCheckTemplateRequestWeakContentRule {
	s.FilterColumnTypes = v
	return s
}

func (s *UpdateDataCheckTemplateRequestWeakContentRule) SetRuleId(v string) *UpdateDataCheckTemplateRequestWeakContentRule {
	s.RuleId = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestWeakContentRule) SetWeakContentAlgorithm(v string) *UpdateDataCheckTemplateRequestWeakContentRule {
	s.WeakContentAlgorithm = &v
	return s
}

func (s *UpdateDataCheckTemplateRequestWeakContentRule) Validate() error {
	return dara.Validate(s)
}
