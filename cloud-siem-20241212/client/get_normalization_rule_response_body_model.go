// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNormalizationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNormalizationRule(v *GetNormalizationRuleResponseBodyNormalizationRule) *GetNormalizationRuleResponseBody
	GetNormalizationRule() *GetNormalizationRuleResponseBodyNormalizationRule
	SetRequestId(v string) *GetNormalizationRuleResponseBody
	GetRequestId() *string
}

type GetNormalizationRuleResponseBody struct {
	NormalizationRule *GetNormalizationRuleResponseBodyNormalizationRule `json:"NormalizationRule,omitempty" xml:"NormalizationRule,omitempty" type:"Struct"`
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNormalizationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNormalizationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetNormalizationRuleResponseBody) GetNormalizationRule() *GetNormalizationRuleResponseBodyNormalizationRule {
	return s.NormalizationRule
}

func (s *GetNormalizationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNormalizationRuleResponseBody) SetNormalizationRule(v *GetNormalizationRuleResponseBodyNormalizationRule) *GetNormalizationRuleResponseBody {
	s.NormalizationRule = v
	return s
}

func (s *GetNormalizationRuleResponseBody) SetRequestId(v string) *GetNormalizationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNormalizationRuleResponseBody) Validate() error {
	if s.NormalizationRule != nil {
		if err := s.NormalizationRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetNormalizationRuleResponseBodyNormalizationRule struct {
	// example:
	//
	// 1733269771123。
	CreateTime          *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExtendContentPacked *string `json:"ExtendContentPacked,omitempty" xml:"ExtendContentPacked,omitempty"`
	// example:
	//
	// NETWORK_CATEGORY。
	NormalizationCategoryId *string `json:"NormalizationCategoryId,omitempty" xml:"NormalizationCategoryId,omitempty"`
	// example:
	//
	// normalization_rule_Z57np。
	NormalizationRuleDescription *string `json:"NormalizationRuleDescription,omitempty" xml:"NormalizationRuleDescription,omitempty"`
	// example:
	//
	// 	- | pack-fields -include=\\"[\\s\\S]+\\" as extend_content。
	NormalizationRuleExpression *string `json:"NormalizationRuleExpression,omitempty" xml:"NormalizationRuleExpression,omitempty"`
	// example:
	//
	// SPL。
	NormalizationRuleFormat *string `json:"NormalizationRuleFormat,omitempty" xml:"NormalizationRuleFormat,omitempty"`
	// example:
	//
	// nr-z0b2ssjteut85uoh9nzp。
	NormalizationRuleId  *string   `json:"NormalizationRuleId,omitempty" xml:"NormalizationRuleId,omitempty"`
	NormalizationRuleIds []*string `json:"NormalizationRuleIds,omitempty" xml:"NormalizationRuleIds,omitempty" type:"Repeated"`
	// example:
	//
	// both。
	NormalizationRuleMode *string `json:"NormalizationRuleMode,omitempty" xml:"NormalizationRuleMode,omitempty"`
	// example:
	//
	// normalization_rule_Z57np。
	NormalizationRuleName *string `json:"NormalizationRuleName,omitempty" xml:"NormalizationRuleName,omitempty"`
	// example:
	//
	// started。
	NormalizationRuleStatus *string `json:"NormalizationRuleStatus,omitempty" xml:"NormalizationRuleStatus,omitempty"`
	// example:
	//
	// predefined。
	NormalizationRuleType *string `json:"NormalizationRuleType,omitempty" xml:"NormalizationRuleType,omitempty"`
	// example:
	//
	// V1。
	NormalizationRuleVersion *int32 `json:"NormalizationRuleVersion,omitempty" xml:"NormalizationRuleVersion,omitempty"`
	// example:
	//
	// HTTP_ACTIVITY。
	NormalizationSchemaId *string `json:"NormalizationSchemaId,omitempty" xml:"NormalizationSchemaId,omitempty"`
	OrderField            *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	// example:
	//
	// alibaba_cloud_sas。
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// 1733269771123。
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// alibaba_cloud。
	VendorId *string `json:"VendorId,omitempty" xml:"VendorId,omitempty"`
}

func (s GetNormalizationRuleResponseBodyNormalizationRule) String() string {
	return dara.Prettify(s)
}

func (s GetNormalizationRuleResponseBodyNormalizationRule) GoString() string {
	return s.String()
}

func (s *GetNormalizationRuleResponseBodyNormalizationRule) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetNormalizationRuleResponseBodyNormalizationRule) GetExtendContentPacked() *string {
	return s.ExtendContentPacked
}

func (s *GetNormalizationRuleResponseBodyNormalizationRule) GetNormalizationCategoryId() *string {
	return s.NormalizationCategoryId
}

func (s *GetNormalizationRuleResponseBodyNormalizationRule) GetNormalizationRuleDescription() *string {
	return s.NormalizationRuleDescription
}

func (s *GetNormalizationRuleResponseBodyNormalizationRule) GetNormalizationRuleExpression() *string {
	return s.NormalizationRuleExpression
}

func (s *GetNormalizationRuleResponseBodyNormalizationRule) GetNormalizationRuleFormat() *string {
	return s.NormalizationRuleFormat
}

func (s *GetNormalizationRuleResponseBodyNormalizationRule) GetNormalizationRuleId() *string {
	return s.NormalizationRuleId
}

func (s *GetNormalizationRuleResponseBodyNormalizationRule) GetNormalizationRuleIds() []*string {
	return s.NormalizationRuleIds
}

func (s *GetNormalizationRuleResponseBodyNormalizationRule) GetNormalizationRuleMode() *string {
	return s.NormalizationRuleMode
}

func (s *GetNormalizationRuleResponseBodyNormalizationRule) GetNormalizationRuleName() *string {
	return s.NormalizationRuleName
}

func (s *GetNormalizationRuleResponseBodyNormalizationRule) GetNormalizationRuleStatus() *string {
	return s.NormalizationRuleStatus
}

func (s *GetNormalizationRuleResponseBodyNormalizationRule) GetNormalizationRuleType() *string {
	return s.NormalizationRuleType
}

func (s *GetNormalizationRuleResponseBodyNormalizationRule) GetNormalizationRuleVersion() *int32 {
	return s.NormalizationRuleVersion
}

func (s *GetNormalizationRuleResponseBodyNormalizationRule) GetNormalizationSchemaId() *string {
	return s.NormalizationSchemaId
}

func (s *GetNormalizationRuleResponseBodyNormalizationRule) GetOrderField() *string {
	return s.OrderField
}

func (s *GetNormalizationRuleResponseBodyNormalizationRule) GetProductId() *string {
	return s.ProductId
}

func (s *GetNormalizationRuleResponseBodyNormalizationRule) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetNormalizationRuleResponseBodyNormalizationRule) GetVendorId() *string {
	return s.VendorId
}

func (s *GetNormalizationRuleResponseBodyNormalizationRule) SetCreateTime(v int64) *GetNormalizationRuleResponseBodyNormalizationRule {
	s.CreateTime = &v
	return s
}

func (s *GetNormalizationRuleResponseBodyNormalizationRule) SetExtendContentPacked(v string) *GetNormalizationRuleResponseBodyNormalizationRule {
	s.ExtendContentPacked = &v
	return s
}

func (s *GetNormalizationRuleResponseBodyNormalizationRule) SetNormalizationCategoryId(v string) *GetNormalizationRuleResponseBodyNormalizationRule {
	s.NormalizationCategoryId = &v
	return s
}

func (s *GetNormalizationRuleResponseBodyNormalizationRule) SetNormalizationRuleDescription(v string) *GetNormalizationRuleResponseBodyNormalizationRule {
	s.NormalizationRuleDescription = &v
	return s
}

func (s *GetNormalizationRuleResponseBodyNormalizationRule) SetNormalizationRuleExpression(v string) *GetNormalizationRuleResponseBodyNormalizationRule {
	s.NormalizationRuleExpression = &v
	return s
}

func (s *GetNormalizationRuleResponseBodyNormalizationRule) SetNormalizationRuleFormat(v string) *GetNormalizationRuleResponseBodyNormalizationRule {
	s.NormalizationRuleFormat = &v
	return s
}

func (s *GetNormalizationRuleResponseBodyNormalizationRule) SetNormalizationRuleId(v string) *GetNormalizationRuleResponseBodyNormalizationRule {
	s.NormalizationRuleId = &v
	return s
}

func (s *GetNormalizationRuleResponseBodyNormalizationRule) SetNormalizationRuleIds(v []*string) *GetNormalizationRuleResponseBodyNormalizationRule {
	s.NormalizationRuleIds = v
	return s
}

func (s *GetNormalizationRuleResponseBodyNormalizationRule) SetNormalizationRuleMode(v string) *GetNormalizationRuleResponseBodyNormalizationRule {
	s.NormalizationRuleMode = &v
	return s
}

func (s *GetNormalizationRuleResponseBodyNormalizationRule) SetNormalizationRuleName(v string) *GetNormalizationRuleResponseBodyNormalizationRule {
	s.NormalizationRuleName = &v
	return s
}

func (s *GetNormalizationRuleResponseBodyNormalizationRule) SetNormalizationRuleStatus(v string) *GetNormalizationRuleResponseBodyNormalizationRule {
	s.NormalizationRuleStatus = &v
	return s
}

func (s *GetNormalizationRuleResponseBodyNormalizationRule) SetNormalizationRuleType(v string) *GetNormalizationRuleResponseBodyNormalizationRule {
	s.NormalizationRuleType = &v
	return s
}

func (s *GetNormalizationRuleResponseBodyNormalizationRule) SetNormalizationRuleVersion(v int32) *GetNormalizationRuleResponseBodyNormalizationRule {
	s.NormalizationRuleVersion = &v
	return s
}

func (s *GetNormalizationRuleResponseBodyNormalizationRule) SetNormalizationSchemaId(v string) *GetNormalizationRuleResponseBodyNormalizationRule {
	s.NormalizationSchemaId = &v
	return s
}

func (s *GetNormalizationRuleResponseBodyNormalizationRule) SetOrderField(v string) *GetNormalizationRuleResponseBodyNormalizationRule {
	s.OrderField = &v
	return s
}

func (s *GetNormalizationRuleResponseBodyNormalizationRule) SetProductId(v string) *GetNormalizationRuleResponseBodyNormalizationRule {
	s.ProductId = &v
	return s
}

func (s *GetNormalizationRuleResponseBodyNormalizationRule) SetUpdateTime(v int64) *GetNormalizationRuleResponseBodyNormalizationRule {
	s.UpdateTime = &v
	return s
}

func (s *GetNormalizationRuleResponseBodyNormalizationRule) SetVendorId(v string) *GetNormalizationRuleResponseBodyNormalizationRule {
	s.VendorId = &v
	return s
}

func (s *GetNormalizationRuleResponseBodyNormalizationRule) Validate() error {
	return dara.Validate(s)
}
