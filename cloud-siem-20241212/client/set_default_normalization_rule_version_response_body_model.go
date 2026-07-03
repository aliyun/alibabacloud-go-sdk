// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultNormalizationRuleVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNormalizationRuleVersion(v *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion) *SetDefaultNormalizationRuleVersionResponseBody
	GetNormalizationRuleVersion() *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion
	SetRequestId(v string) *SetDefaultNormalizationRuleVersionResponseBody
	GetRequestId() *string
}

type SetDefaultNormalizationRuleVersionResponseBody struct {
	// The normalization rule version.
	NormalizationRuleVersion *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion `json:"NormalizationRuleVersion,omitempty" xml:"NormalizationRuleVersion,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDefaultNormalizationRuleVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultNormalizationRuleVersionResponseBody) GoString() string {
	return s.String()
}

func (s *SetDefaultNormalizationRuleVersionResponseBody) GetNormalizationRuleVersion() *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion {
	return s.NormalizationRuleVersion
}

func (s *SetDefaultNormalizationRuleVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDefaultNormalizationRuleVersionResponseBody) SetNormalizationRuleVersion(v *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion) *SetDefaultNormalizationRuleVersionResponseBody {
	s.NormalizationRuleVersion = v
	return s
}

func (s *SetDefaultNormalizationRuleVersionResponseBody) SetRequestId(v string) *SetDefaultNormalizationRuleVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDefaultNormalizationRuleVersionResponseBody) Validate() error {
	if s.NormalizationRuleVersion != nil {
		if err := s.NormalizationRuleVersion.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion struct {
	// The creation time.
	//
	// example:
	//
	// 1733269771123
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The normalization rule category ID.
	//
	// example:
	//
	// NETWORK_CATEGORY
	NormalizationCategoryId *string `json:"NormalizationCategoryId,omitempty" xml:"NormalizationCategoryId,omitempty"`
	// The normalization rule description.
	//
	// example:
	//
	// normalization_rule_Z57np
	NormalizationRuleDescription *string `json:"NormalizationRuleDescription,omitempty" xml:"NormalizationRuleDescription,omitempty"`
	// The normalization rule expression.
	//
	// example:
	//
	// 	- | pack-fields -include=\\"[\\s\\S]+\\" as extend_content
	NormalizationRuleExpression *string `json:"NormalizationRuleExpression,omitempty" xml:"NormalizationRuleExpression,omitempty"`
	// The normalization rule format.
	//
	// example:
	//
	// SPL
	NormalizationRuleFormat *string `json:"NormalizationRuleFormat,omitempty" xml:"NormalizationRuleFormat,omitempty"`
	// The normalization rule ID.
	//
	// example:
	//
	// nr-z0b2ssjteut85uoh9nzp
	NormalizationRuleId *string `json:"NormalizationRuleId,omitempty" xml:"NormalizationRuleId,omitempty"`
	// The normalization rule name.
	//
	// example:
	//
	// normalization_rule_Z57np
	NormalizationRuleName *string `json:"NormalizationRuleName,omitempty" xml:"NormalizationRuleName,omitempty"`
	// The normalization rule status.
	//
	// example:
	//
	// started
	NormalizationRuleStatus *string `json:"NormalizationRuleStatus,omitempty" xml:"NormalizationRuleStatus,omitempty"`
	// The normalization rule type. Valid values:
	//
	// - predefined: predefined normalization rule.
	//
	// - custom: custom normalization rule.
	//
	// example:
	//
	// predefined
	NormalizationRuleType *string `json:"NormalizationRuleType,omitempty" xml:"NormalizationRuleType,omitempty"`
	// The current version of the normalization rule.
	//
	// example:
	//
	// V1
	NormalizationRuleVersion *int32 `json:"NormalizationRuleVersion,omitempty" xml:"NormalizationRuleVersion,omitempty"`
	// The normalization schema ID.
	//
	// example:
	//
	// HTTP_ACTIVITY
	NormalizationSchemaId *string `json:"NormalizationSchemaId,omitempty" xml:"NormalizationSchemaId,omitempty"`
	// example:
	//
	// NETWORK_AND_WEB_SECURITY
	NormalizationSecurityDomainId *string `json:"NormalizationSecurityDomainId,omitempty" xml:"NormalizationSecurityDomainId,omitempty"`
	// The product ID.
	//
	// example:
	//
	// alibaba_cloud_sas
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The update time.
	//
	// example:
	//
	// 1733269771123
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The vendor ID associated with the normalization rule.
	//
	// example:
	//
	// alibaba_cloud
	VendorId *string `json:"VendorId,omitempty" xml:"VendorId,omitempty"`
}

func (s SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion) GoString() string {
	return s.String()
}

func (s *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion) GetNormalizationCategoryId() *string {
	return s.NormalizationCategoryId
}

func (s *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion) GetNormalizationRuleDescription() *string {
	return s.NormalizationRuleDescription
}

func (s *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion) GetNormalizationRuleExpression() *string {
	return s.NormalizationRuleExpression
}

func (s *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion) GetNormalizationRuleFormat() *string {
	return s.NormalizationRuleFormat
}

func (s *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion) GetNormalizationRuleId() *string {
	return s.NormalizationRuleId
}

func (s *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion) GetNormalizationRuleName() *string {
	return s.NormalizationRuleName
}

func (s *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion) GetNormalizationRuleStatus() *string {
	return s.NormalizationRuleStatus
}

func (s *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion) GetNormalizationRuleType() *string {
	return s.NormalizationRuleType
}

func (s *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion) GetNormalizationRuleVersion() *int32 {
	return s.NormalizationRuleVersion
}

func (s *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion) GetNormalizationSchemaId() *string {
	return s.NormalizationSchemaId
}

func (s *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion) GetNormalizationSecurityDomainId() *string {
	return s.NormalizationSecurityDomainId
}

func (s *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion) GetProductId() *string {
	return s.ProductId
}

func (s *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion) GetVendorId() *string {
	return s.VendorId
}

func (s *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion) SetCreateTime(v int64) *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion {
	s.CreateTime = &v
	return s
}

func (s *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion) SetNormalizationCategoryId(v string) *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion {
	s.NormalizationCategoryId = &v
	return s
}

func (s *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion) SetNormalizationRuleDescription(v string) *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion {
	s.NormalizationRuleDescription = &v
	return s
}

func (s *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion) SetNormalizationRuleExpression(v string) *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion {
	s.NormalizationRuleExpression = &v
	return s
}

func (s *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion) SetNormalizationRuleFormat(v string) *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion {
	s.NormalizationRuleFormat = &v
	return s
}

func (s *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion) SetNormalizationRuleId(v string) *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion {
	s.NormalizationRuleId = &v
	return s
}

func (s *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion) SetNormalizationRuleName(v string) *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion {
	s.NormalizationRuleName = &v
	return s
}

func (s *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion) SetNormalizationRuleStatus(v string) *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion {
	s.NormalizationRuleStatus = &v
	return s
}

func (s *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion) SetNormalizationRuleType(v string) *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion {
	s.NormalizationRuleType = &v
	return s
}

func (s *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion) SetNormalizationRuleVersion(v int32) *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion {
	s.NormalizationRuleVersion = &v
	return s
}

func (s *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion) SetNormalizationSchemaId(v string) *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion {
	s.NormalizationSchemaId = &v
	return s
}

func (s *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion) SetNormalizationSecurityDomainId(v string) *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion {
	s.NormalizationSecurityDomainId = &v
	return s
}

func (s *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion) SetProductId(v string) *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion {
	s.ProductId = &v
	return s
}

func (s *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion) SetUpdateTime(v int64) *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion {
	s.UpdateTime = &v
	return s
}

func (s *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion) SetVendorId(v string) *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion {
	s.VendorId = &v
	return s
}

func (s *SetDefaultNormalizationRuleVersionResponseBodyNormalizationRuleVersion) Validate() error {
	return dara.Validate(s)
}
