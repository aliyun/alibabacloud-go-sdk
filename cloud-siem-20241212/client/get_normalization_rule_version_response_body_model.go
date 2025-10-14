// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNormalizationRuleVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNormalizationRuleVersion(v *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion) *GetNormalizationRuleVersionResponseBody
	GetNormalizationRuleVersion() *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion
	SetRequestId(v string) *GetNormalizationRuleVersionResponseBody
	GetRequestId() *string
}

type GetNormalizationRuleVersionResponseBody struct {
	NormalizationRuleVersion *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion `json:"NormalizationRuleVersion,omitempty" xml:"NormalizationRuleVersion,omitempty" type:"Struct"`
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNormalizationRuleVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNormalizationRuleVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetNormalizationRuleVersionResponseBody) GetNormalizationRuleVersion() *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion {
	return s.NormalizationRuleVersion
}

func (s *GetNormalizationRuleVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNormalizationRuleVersionResponseBody) SetNormalizationRuleVersion(v *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion) *GetNormalizationRuleVersionResponseBody {
	s.NormalizationRuleVersion = v
	return s
}

func (s *GetNormalizationRuleVersionResponseBody) SetRequestId(v string) *GetNormalizationRuleVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNormalizationRuleVersionResponseBody) Validate() error {
	if s.NormalizationRuleVersion != nil {
		if err := s.NormalizationRuleVersion.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion struct {
	// example:
	//
	// 1733269771123。
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
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
	NormalizationRuleId *string `json:"NormalizationRuleId,omitempty" xml:"NormalizationRuleId,omitempty"`
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
	NormalizationRuleVersion     *int32  `json:"NormalizationRuleVersion,omitempty" xml:"NormalizationRuleVersion,omitempty"`
	NormalizationRuleVersionName *string `json:"NormalizationRuleVersionName,omitempty" xml:"NormalizationRuleVersionName,omitempty"`
	// example:
	//
	// HTTP_ACTIVITY。
	NormalizationSchemaId *string `json:"NormalizationSchemaId,omitempty" xml:"NormalizationSchemaId,omitempty"`
	// example:
	//
	// alibaba_cloud_sas。
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 1733269771123。
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// alibaba_cloud。
	VendorId *string `json:"VendorId,omitempty" xml:"VendorId,omitempty"`
}

func (s GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion) String() string {
	return dara.Prettify(s)
}

func (s GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion) GoString() string {
	return s.String()
}

func (s *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion) GetNormalizationCategoryId() *string {
	return s.NormalizationCategoryId
}

func (s *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion) GetNormalizationRuleDescription() *string {
	return s.NormalizationRuleDescription
}

func (s *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion) GetNormalizationRuleExpression() *string {
	return s.NormalizationRuleExpression
}

func (s *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion) GetNormalizationRuleFormat() *string {
	return s.NormalizationRuleFormat
}

func (s *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion) GetNormalizationRuleId() *string {
	return s.NormalizationRuleId
}

func (s *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion) GetNormalizationRuleName() *string {
	return s.NormalizationRuleName
}

func (s *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion) GetNormalizationRuleStatus() *string {
	return s.NormalizationRuleStatus
}

func (s *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion) GetNormalizationRuleType() *string {
	return s.NormalizationRuleType
}

func (s *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion) GetNormalizationRuleVersion() *int32 {
	return s.NormalizationRuleVersion
}

func (s *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion) GetNormalizationRuleVersionName() *string {
	return s.NormalizationRuleVersionName
}

func (s *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion) GetNormalizationSchemaId() *string {
	return s.NormalizationSchemaId
}

func (s *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion) GetProductId() *string {
	return s.ProductId
}

func (s *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion) GetRegionId() *string {
	return s.RegionId
}

func (s *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion) GetVendorId() *string {
	return s.VendorId
}

func (s *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion) SetCreateTime(v int64) *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion {
	s.CreateTime = &v
	return s
}

func (s *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion) SetNormalizationCategoryId(v string) *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion {
	s.NormalizationCategoryId = &v
	return s
}

func (s *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion) SetNormalizationRuleDescription(v string) *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion {
	s.NormalizationRuleDescription = &v
	return s
}

func (s *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion) SetNormalizationRuleExpression(v string) *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion {
	s.NormalizationRuleExpression = &v
	return s
}

func (s *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion) SetNormalizationRuleFormat(v string) *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion {
	s.NormalizationRuleFormat = &v
	return s
}

func (s *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion) SetNormalizationRuleId(v string) *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion {
	s.NormalizationRuleId = &v
	return s
}

func (s *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion) SetNormalizationRuleName(v string) *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion {
	s.NormalizationRuleName = &v
	return s
}

func (s *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion) SetNormalizationRuleStatus(v string) *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion {
	s.NormalizationRuleStatus = &v
	return s
}

func (s *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion) SetNormalizationRuleType(v string) *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion {
	s.NormalizationRuleType = &v
	return s
}

func (s *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion) SetNormalizationRuleVersion(v int32) *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion {
	s.NormalizationRuleVersion = &v
	return s
}

func (s *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion) SetNormalizationRuleVersionName(v string) *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion {
	s.NormalizationRuleVersionName = &v
	return s
}

func (s *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion) SetNormalizationSchemaId(v string) *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion {
	s.NormalizationSchemaId = &v
	return s
}

func (s *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion) SetProductId(v string) *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion {
	s.ProductId = &v
	return s
}

func (s *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion) SetRegionId(v string) *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion {
	s.RegionId = &v
	return s
}

func (s *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion) SetUpdateTime(v int64) *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion {
	s.UpdateTime = &v
	return s
}

func (s *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion) SetVendorId(v string) *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion {
	s.VendorId = &v
	return s
}

func (s *GetNormalizationRuleVersionResponseBodyNormalizationRuleVersion) Validate() error {
	return dara.Validate(s)
}
