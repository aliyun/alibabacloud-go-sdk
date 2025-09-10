// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotaApplicationTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListQuotaApplicationTemplatesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListQuotaApplicationTemplatesResponseBody
	GetNextToken() *string
	SetQuotaApplicationTemplates(v []*ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) *ListQuotaApplicationTemplatesResponseBody
	GetQuotaApplicationTemplates() []*ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates
	SetRequestId(v string) *ListQuotaApplicationTemplatesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListQuotaApplicationTemplatesResponseBody
	GetTotalCount() *int32
}

type ListQuotaApplicationTemplatesResponseBody struct {
	// The maximum number of records returned for the query.
	//
	// example:
	//
	// 30
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position at which the query ends.
	//
	// > An empty value indicates that all data is returned.
	//
	// example:
	//
	// 1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The information about quota templates.
	QuotaApplicationTemplates []*ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates `json:"QuotaApplicationTemplates,omitempty" xml:"QuotaApplicationTemplates,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D47B3A10-CDAC-5412-B2EE-EC9A3DBE9053
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records returned for the query.
	//
	// example:
	//
	// 40
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListQuotaApplicationTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaApplicationTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationTemplatesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListQuotaApplicationTemplatesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListQuotaApplicationTemplatesResponseBody) GetQuotaApplicationTemplates() []*ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates {
	return s.QuotaApplicationTemplates
}

func (s *ListQuotaApplicationTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListQuotaApplicationTemplatesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListQuotaApplicationTemplatesResponseBody) SetMaxResults(v int32) *ListQuotaApplicationTemplatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBody) SetNextToken(v string) *ListQuotaApplicationTemplatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBody) SetQuotaApplicationTemplates(v []*ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) *ListQuotaApplicationTemplatesResponseBody {
	s.QuotaApplicationTemplates = v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBody) SetRequestId(v string) *ListQuotaApplicationTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBody) SetTotalCount(v int32) *ListQuotaApplicationTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates struct {
	// None.
	ApplicableRange []*float32 `json:"ApplicableRange,omitempty" xml:"ApplicableRange,omitempty" type:"Repeated"`
	// The type of the adjustable value. Valid values:
	//
	// 	- continuous
	//
	// 	- discontinuous
	//
	// example:
	//
	// continuous
	ApplicableType *string `json:"ApplicableType,omitempty" xml:"ApplicableType,omitempty"`
	// The requested value of the quota.
	//
	// example:
	//
	// 802
	DesireValue *float32 `json:"DesireValue,omitempty" xml:"DesireValue,omitempty"`
	// The quota dimensions.
	//
	// Format: {"regionId":"Region"}.
	//
	// example:
	//
	// {"regionId":"cn-hangzhou"}
	Dimensions map[string]interface{} `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// The start time of the validity period of the quota. The value is displayed in UTC.
	//
	// example:
	//
	// 2022-09-28T06:07:00Z
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The language of the quota alert notification. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	EnvLanguage *string `json:"EnvLanguage,omitempty" xml:"EnvLanguage,omitempty"`
	// The end time of the validity period of the quota. The value is displayed in UTC.
	//
	// example:
	//
	// 2022-09-29T06:07:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the quota template.
	//
	// example:
	//
	// 1****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether Quota Center sends a notification about the application result. Valid values:
	//
	// 	- 0: no
	//
	// 	- 3: yes
	//
	// example:
	//
	// 0
	NoticeType *int32 `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// The calculation cycle of the quota.
	Period *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplatesPeriod `json:"Period,omitempty" xml:"Period,omitempty" type:"Struct"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the quota.
	//
	// example:
	//
	// q_security-groups
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The type of the quota. Valid values:
	//
	// 	- CommonQuota: general quota
	//
	// 	- FlowControl: API rate limit
	//
	// 	- WhiteListLabel: privilege
	//
	// example:
	//
	// CommonQuota
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
	// The description of the quota.
	//
	// example:
	//
	// The maximum number of security groups that can be created by the current account.
	QuotaDescription *string `json:"QuotaDescription,omitempty" xml:"QuotaDescription,omitempty"`
	// The name of the quota.
	//
	// example:
	//
	// Maximum Number of Security Groups
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
}

func (s ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) GetApplicableRange() []*float32 {
	return s.ApplicableRange
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) GetApplicableType() *string {
	return s.ApplicableType
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) GetDesireValue() *float32 {
	return s.DesireValue
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) GetDimensions() map[string]interface{} {
	return s.Dimensions
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) GetEnvLanguage() *string {
	return s.EnvLanguage
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) GetId() *string {
	return s.Id
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) GetNoticeType() *int32 {
	return s.NoticeType
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) GetPeriod() *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplatesPeriod {
	return s.Period
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) GetQuotaActionCode() *string {
	return s.QuotaActionCode
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) GetQuotaCategory() *string {
	return s.QuotaCategory
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) GetQuotaDescription() *string {
	return s.QuotaDescription
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) GetQuotaName() *string {
	return s.QuotaName
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) SetApplicableRange(v []*float32) *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates {
	s.ApplicableRange = v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) SetApplicableType(v string) *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates {
	s.ApplicableType = &v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) SetDesireValue(v float32) *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates {
	s.DesireValue = &v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) SetDimensions(v map[string]interface{}) *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates {
	s.Dimensions = v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) SetEffectiveTime(v string) *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates {
	s.EffectiveTime = &v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) SetEnvLanguage(v string) *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates {
	s.EnvLanguage = &v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) SetExpireTime(v string) *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates {
	s.ExpireTime = &v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) SetId(v string) *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates {
	s.Id = &v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) SetNoticeType(v int32) *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates {
	s.NoticeType = &v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) SetPeriod(v *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplatesPeriod) *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates {
	s.Period = v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) SetProductCode(v string) *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates {
	s.ProductCode = &v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) SetQuotaActionCode(v string) *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates {
	s.QuotaActionCode = &v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) SetQuotaCategory(v string) *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates {
	s.QuotaCategory = &v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) SetQuotaDescription(v string) *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates {
	s.QuotaDescription = &v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) SetQuotaName(v string) *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates {
	s.QuotaName = &v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplates) Validate() error {
	return dara.Validate(s)
}

type ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplatesPeriod struct {
	// The unit of the calculation cycle. Valid values:
	//
	// 	- second
	//
	// 	- minute
	//
	// 	- hour
	//
	// 	- day
	//
	// 	- week
	//
	// example:
	//
	// day
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The value of the calculation cycle.
	//
	// example:
	//
	// 1
	PeriodValue *int32 `json:"PeriodValue,omitempty" xml:"PeriodValue,omitempty"`
}

func (s ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplatesPeriod) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplatesPeriod) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplatesPeriod) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplatesPeriod) GetPeriodValue() *int32 {
	return s.PeriodValue
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplatesPeriod) SetPeriodUnit(v string) *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplatesPeriod {
	s.PeriodUnit = &v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplatesPeriod) SetPeriodValue(v int32) *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplatesPeriod {
	s.PeriodValue = &v
	return s
}

func (s *ListQuotaApplicationTemplatesResponseBodyQuotaApplicationTemplatesPeriod) Validate() error {
	return dara.Validate(s)
}
