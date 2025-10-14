// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNormalizationRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListNormalizationRulesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListNormalizationRulesResponseBody
	GetNextToken() *string
	SetNormalizationRules(v []*ListNormalizationRulesResponseBodyNormalizationRules) *ListNormalizationRulesResponseBody
	GetNormalizationRules() []*ListNormalizationRulesResponseBodyNormalizationRules
	SetPageNumber(v int32) *ListNormalizationRulesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListNormalizationRulesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListNormalizationRulesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListNormalizationRulesResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListNormalizationRulesResponseBody
	GetTotalPage() *int32
}

type ListNormalizationRulesResponseBody struct {
	// example:
	//
	// 50。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****。
	NextToken          *string                                                 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	NormalizationRules []*ListNormalizationRulesResponseBodyNormalizationRules `json:"NormalizationRules,omitempty" xml:"NormalizationRules,omitempty" type:"Repeated"`
	// example:
	//
	// 1。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 2。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 5。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 1。
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListNormalizationRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNormalizationRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNormalizationRulesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNormalizationRulesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNormalizationRulesResponseBody) GetNormalizationRules() []*ListNormalizationRulesResponseBodyNormalizationRules {
	return s.NormalizationRules
}

func (s *ListNormalizationRulesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListNormalizationRulesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNormalizationRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNormalizationRulesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListNormalizationRulesResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListNormalizationRulesResponseBody) SetMaxResults(v int32) *ListNormalizationRulesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListNormalizationRulesResponseBody) SetNextToken(v string) *ListNormalizationRulesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListNormalizationRulesResponseBody) SetNormalizationRules(v []*ListNormalizationRulesResponseBodyNormalizationRules) *ListNormalizationRulesResponseBody {
	s.NormalizationRules = v
	return s
}

func (s *ListNormalizationRulesResponseBody) SetPageNumber(v int32) *ListNormalizationRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListNormalizationRulesResponseBody) SetPageSize(v int32) *ListNormalizationRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListNormalizationRulesResponseBody) SetRequestId(v string) *ListNormalizationRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNormalizationRulesResponseBody) SetTotalCount(v int32) *ListNormalizationRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListNormalizationRulesResponseBody) SetTotalPage(v int32) *ListNormalizationRulesResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListNormalizationRulesResponseBody) Validate() error {
	if s.NormalizationRules != nil {
		for _, item := range s.NormalizationRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNormalizationRulesResponseBodyNormalizationRules struct {
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
	NormalizationRuleId *string `json:"NormalizationRuleId,omitempty" xml:"NormalizationRuleId,omitempty"`
	// example:
	//
	// both。
	NormalizationRuleMode *string `json:"NormalizationRuleMode,omitempty" xml:"NormalizationRuleMode,omitempty"`
	// example:
	//
	// normalization_rule_Z57np。
	NormalizationRuleName       *string                                                                            `json:"NormalizationRuleName,omitempty" xml:"NormalizationRuleName,omitempty"`
	NormalizationRuleReferences []*ListNormalizationRulesResponseBodyNormalizationRulesNormalizationRuleReferences `json:"NormalizationRuleReferences,omitempty" xml:"NormalizationRuleReferences,omitempty" type:"Repeated"`
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
	NormalizationRuleVersion *string `json:"NormalizationRuleVersion,omitempty" xml:"NormalizationRuleVersion,omitempty"`
	// example:
	//
	// HTTP_ACTIVITY。
	NormalizationSchemaId *string `json:"NormalizationSchemaId,omitempty" xml:"NormalizationSchemaId,omitempty"`
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

func (s ListNormalizationRulesResponseBodyNormalizationRules) String() string {
	return dara.Prettify(s)
}

func (s ListNormalizationRulesResponseBodyNormalizationRules) GoString() string {
	return s.String()
}

func (s *ListNormalizationRulesResponseBodyNormalizationRules) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListNormalizationRulesResponseBodyNormalizationRules) GetExtendContentPacked() *string {
	return s.ExtendContentPacked
}

func (s *ListNormalizationRulesResponseBodyNormalizationRules) GetNormalizationCategoryId() *string {
	return s.NormalizationCategoryId
}

func (s *ListNormalizationRulesResponseBodyNormalizationRules) GetNormalizationRuleDescription() *string {
	return s.NormalizationRuleDescription
}

func (s *ListNormalizationRulesResponseBodyNormalizationRules) GetNormalizationRuleExpression() *string {
	return s.NormalizationRuleExpression
}

func (s *ListNormalizationRulesResponseBodyNormalizationRules) GetNormalizationRuleFormat() *string {
	return s.NormalizationRuleFormat
}

func (s *ListNormalizationRulesResponseBodyNormalizationRules) GetNormalizationRuleId() *string {
	return s.NormalizationRuleId
}

func (s *ListNormalizationRulesResponseBodyNormalizationRules) GetNormalizationRuleMode() *string {
	return s.NormalizationRuleMode
}

func (s *ListNormalizationRulesResponseBodyNormalizationRules) GetNormalizationRuleName() *string {
	return s.NormalizationRuleName
}

func (s *ListNormalizationRulesResponseBodyNormalizationRules) GetNormalizationRuleReferences() []*ListNormalizationRulesResponseBodyNormalizationRulesNormalizationRuleReferences {
	return s.NormalizationRuleReferences
}

func (s *ListNormalizationRulesResponseBodyNormalizationRules) GetNormalizationRuleStatus() *string {
	return s.NormalizationRuleStatus
}

func (s *ListNormalizationRulesResponseBodyNormalizationRules) GetNormalizationRuleType() *string {
	return s.NormalizationRuleType
}

func (s *ListNormalizationRulesResponseBodyNormalizationRules) GetNormalizationRuleVersion() *string {
	return s.NormalizationRuleVersion
}

func (s *ListNormalizationRulesResponseBodyNormalizationRules) GetNormalizationSchemaId() *string {
	return s.NormalizationSchemaId
}

func (s *ListNormalizationRulesResponseBodyNormalizationRules) GetProductId() *string {
	return s.ProductId
}

func (s *ListNormalizationRulesResponseBodyNormalizationRules) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListNormalizationRulesResponseBodyNormalizationRules) GetVendorId() *string {
	return s.VendorId
}

func (s *ListNormalizationRulesResponseBodyNormalizationRules) SetCreateTime(v int64) *ListNormalizationRulesResponseBodyNormalizationRules {
	s.CreateTime = &v
	return s
}

func (s *ListNormalizationRulesResponseBodyNormalizationRules) SetExtendContentPacked(v string) *ListNormalizationRulesResponseBodyNormalizationRules {
	s.ExtendContentPacked = &v
	return s
}

func (s *ListNormalizationRulesResponseBodyNormalizationRules) SetNormalizationCategoryId(v string) *ListNormalizationRulesResponseBodyNormalizationRules {
	s.NormalizationCategoryId = &v
	return s
}

func (s *ListNormalizationRulesResponseBodyNormalizationRules) SetNormalizationRuleDescription(v string) *ListNormalizationRulesResponseBodyNormalizationRules {
	s.NormalizationRuleDescription = &v
	return s
}

func (s *ListNormalizationRulesResponseBodyNormalizationRules) SetNormalizationRuleExpression(v string) *ListNormalizationRulesResponseBodyNormalizationRules {
	s.NormalizationRuleExpression = &v
	return s
}

func (s *ListNormalizationRulesResponseBodyNormalizationRules) SetNormalizationRuleFormat(v string) *ListNormalizationRulesResponseBodyNormalizationRules {
	s.NormalizationRuleFormat = &v
	return s
}

func (s *ListNormalizationRulesResponseBodyNormalizationRules) SetNormalizationRuleId(v string) *ListNormalizationRulesResponseBodyNormalizationRules {
	s.NormalizationRuleId = &v
	return s
}

func (s *ListNormalizationRulesResponseBodyNormalizationRules) SetNormalizationRuleMode(v string) *ListNormalizationRulesResponseBodyNormalizationRules {
	s.NormalizationRuleMode = &v
	return s
}

func (s *ListNormalizationRulesResponseBodyNormalizationRules) SetNormalizationRuleName(v string) *ListNormalizationRulesResponseBodyNormalizationRules {
	s.NormalizationRuleName = &v
	return s
}

func (s *ListNormalizationRulesResponseBodyNormalizationRules) SetNormalizationRuleReferences(v []*ListNormalizationRulesResponseBodyNormalizationRulesNormalizationRuleReferences) *ListNormalizationRulesResponseBodyNormalizationRules {
	s.NormalizationRuleReferences = v
	return s
}

func (s *ListNormalizationRulesResponseBodyNormalizationRules) SetNormalizationRuleStatus(v string) *ListNormalizationRulesResponseBodyNormalizationRules {
	s.NormalizationRuleStatus = &v
	return s
}

func (s *ListNormalizationRulesResponseBodyNormalizationRules) SetNormalizationRuleType(v string) *ListNormalizationRulesResponseBodyNormalizationRules {
	s.NormalizationRuleType = &v
	return s
}

func (s *ListNormalizationRulesResponseBodyNormalizationRules) SetNormalizationRuleVersion(v string) *ListNormalizationRulesResponseBodyNormalizationRules {
	s.NormalizationRuleVersion = &v
	return s
}

func (s *ListNormalizationRulesResponseBodyNormalizationRules) SetNormalizationSchemaId(v string) *ListNormalizationRulesResponseBodyNormalizationRules {
	s.NormalizationSchemaId = &v
	return s
}

func (s *ListNormalizationRulesResponseBodyNormalizationRules) SetProductId(v string) *ListNormalizationRulesResponseBodyNormalizationRules {
	s.ProductId = &v
	return s
}

func (s *ListNormalizationRulesResponseBodyNormalizationRules) SetUpdateTime(v int64) *ListNormalizationRulesResponseBodyNormalizationRules {
	s.UpdateTime = &v
	return s
}

func (s *ListNormalizationRulesResponseBodyNormalizationRules) SetVendorId(v string) *ListNormalizationRulesResponseBodyNormalizationRules {
	s.VendorId = &v
	return s
}

func (s *ListNormalizationRulesResponseBodyNormalizationRules) Validate() error {
	if s.NormalizationRuleReferences != nil {
		for _, item := range s.NormalizationRuleReferences {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNormalizationRulesResponseBodyNormalizationRulesNormalizationRuleReferences struct {
	// example:
	//
	// alibaba_cloud_bot_flow_ingestion_173326*******。
	DataIngestionId *string `json:"DataIngestionId,omitempty" xml:"DataIngestionId,omitempty"`
}

func (s ListNormalizationRulesResponseBodyNormalizationRulesNormalizationRuleReferences) String() string {
	return dara.Prettify(s)
}

func (s ListNormalizationRulesResponseBodyNormalizationRulesNormalizationRuleReferences) GoString() string {
	return s.String()
}

func (s *ListNormalizationRulesResponseBodyNormalizationRulesNormalizationRuleReferences) GetDataIngestionId() *string {
	return s.DataIngestionId
}

func (s *ListNormalizationRulesResponseBodyNormalizationRulesNormalizationRuleReferences) SetDataIngestionId(v string) *ListNormalizationRulesResponseBodyNormalizationRulesNormalizationRuleReferences {
	s.DataIngestionId = &v
	return s
}

func (s *ListNormalizationRulesResponseBodyNormalizationRulesNormalizationRuleReferences) Validate() error {
	return dara.Validate(s)
}
