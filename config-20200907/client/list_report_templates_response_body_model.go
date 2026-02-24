// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListReportTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListReportTemplatesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListReportTemplatesResponseBody
	GetNextToken() *string
	SetReportTemplateList(v []*ListReportTemplatesResponseBodyReportTemplateList) *ListReportTemplatesResponseBody
	GetReportTemplateList() []*ListReportTemplatesResponseBodyReportTemplateList
	SetRequestId(v string) *ListReportTemplatesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListReportTemplatesResponseBody
	GetTotalCount() *int32
}

type ListReportTemplatesResponseBody struct {
	// The maximum number of entries to return per page. Valid values: 1 to 50. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// If the response is truncated, use NextToken to send another request and get results after the truncation point.
	//
	// example:
	//
	// aVCjqNaSy0Ps7zSMGu25****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The list of report templates.
	ReportTemplateList []*ListReportTemplatesResponseBodyReportTemplateList `json:"ReportTemplateList,omitempty" xml:"ReportTemplateList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 8195B664-9565-4685-89AC-8B5F04B44B92
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of templates.
	//
	// example:
	//
	// 7
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListReportTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListReportTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListReportTemplatesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListReportTemplatesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListReportTemplatesResponseBody) GetReportTemplateList() []*ListReportTemplatesResponseBodyReportTemplateList {
	return s.ReportTemplateList
}

func (s *ListReportTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListReportTemplatesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListReportTemplatesResponseBody) SetMaxResults(v int32) *ListReportTemplatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListReportTemplatesResponseBody) SetNextToken(v string) *ListReportTemplatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListReportTemplatesResponseBody) SetReportTemplateList(v []*ListReportTemplatesResponseBodyReportTemplateList) *ListReportTemplatesResponseBody {
	s.ReportTemplateList = v
	return s
}

func (s *ListReportTemplatesResponseBody) SetRequestId(v string) *ListReportTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListReportTemplatesResponseBody) SetTotalCount(v int32) *ListReportTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListReportTemplatesResponseBody) Validate() error {
	if s.ReportTemplateList != nil {
		for _, item := range s.ReportTemplateList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListReportTemplatesResponseBodyReportTemplateList struct {
	// The format of the report. Only Excel is supported.
	//
	// example:
	//
	// excel
	ReportFileFormats *string `json:"ReportFileFormats,omitempty" xml:"ReportFileFormats,omitempty"`
	// The aggregation granularity of the report.
	//
	// From the management account perspective, the following options are supported:
	//
	// - AllInOne: Aggregate all accounts in the template scope into one report.
	//
	// - GroupByAggregator: Generate reports by aggregator group. Output as one compressed file.
	//
	// - GroupByAccount: Generate separate reports for each account (default). Output as one compressed file.
	//
	// Member accounts support only GroupByAccount.
	//
	// example:
	//
	// AllInOne
	ReportGranularity *string `json:"ReportGranularity,omitempty" xml:"ReportGranularity,omitempty"`
	// The language of the report. Valid values: zh-CN and en-US. Default value: en-US.
	//
	// example:
	//
	// en-US
	ReportLanguage *string `json:"ReportLanguage,omitempty" xml:"ReportLanguage,omitempty"`
	// An array that defines which rules appear in the audit report. Each ReportScope object uses OR logic (additive logic).
	//
	// > For example, if the array has two items — the first specifies RuleId cr-1 and the second specifies RuleId cr-2 — then the report covers both cr-1 and cr-2.
	ReportScope []*ListReportTemplatesResponseBodyReportTemplateListReportScope `json:"ReportScope,omitempty" xml:"ReportScope,omitempty" type:"Repeated"`
	// The description of the report template.
	//
	// example:
	//
	// test-description
	ReportTemplateDescription *string `json:"ReportTemplateDescription,omitempty" xml:"ReportTemplateDescription,omitempty"`
	// The ID of the report template.
	//
	// example:
	//
	// crt-xxx
	ReportTemplateId *string `json:"ReportTemplateId,omitempty" xml:"ReportTemplateId,omitempty"`
	// The name of the report template.
	//
	// example:
	//
	// test-name
	ReportTemplateName *string `json:"ReportTemplateName,omitempty" xml:"ReportTemplateName,omitempty"`
	// The subscription frequency of the report. If this parameter is specified, it must be a Quartz-formatted cron expression.
	//
	// The format is: second minute hour day month weekday. Common examples:
	//
	// - Run daily at 00:00: 0 0 0 \\	- \\	- ?
	//
	// - Run every Monday at 15:30: 0 30 15 ? \\	- MON
	//
	// - Run on the first day of each month at 02:00: 0 0 2 1 \\	- ?
	//
	// Where:
	//
	// - "\\*" means any value.
	//
	// - "?" means no specific value for day or weekday.
	//
	// - "MON" means Monday.
	//
	// > Times are in UTC+8. Adjust your cron expression based on your local time zone.
	//
	// > The system tries to run reports at the scheduled time, but delays may occur due to report generation. Each template can trigger at most one notification per day.
	//
	// > In Quartz, weekdays are numbered starting from Sunday: 1 = Sunday, 2 = Monday, 3 = Tuesday, 4 = Wednesday, 5 = Thursday, 6 = Friday, 7 = Saturday.
	//
	// example:
	//
	// 0 0 13 	- 	- ?
	SubscriptionFrequency *string `json:"SubscriptionFrequency,omitempty" xml:"SubscriptionFrequency,omitempty"`
}

func (s ListReportTemplatesResponseBodyReportTemplateList) String() string {
	return dara.Prettify(s)
}

func (s ListReportTemplatesResponseBodyReportTemplateList) GoString() string {
	return s.String()
}

func (s *ListReportTemplatesResponseBodyReportTemplateList) GetReportFileFormats() *string {
	return s.ReportFileFormats
}

func (s *ListReportTemplatesResponseBodyReportTemplateList) GetReportGranularity() *string {
	return s.ReportGranularity
}

func (s *ListReportTemplatesResponseBodyReportTemplateList) GetReportLanguage() *string {
	return s.ReportLanguage
}

func (s *ListReportTemplatesResponseBodyReportTemplateList) GetReportScope() []*ListReportTemplatesResponseBodyReportTemplateListReportScope {
	return s.ReportScope
}

func (s *ListReportTemplatesResponseBodyReportTemplateList) GetReportTemplateDescription() *string {
	return s.ReportTemplateDescription
}

func (s *ListReportTemplatesResponseBodyReportTemplateList) GetReportTemplateId() *string {
	return s.ReportTemplateId
}

func (s *ListReportTemplatesResponseBodyReportTemplateList) GetReportTemplateName() *string {
	return s.ReportTemplateName
}

func (s *ListReportTemplatesResponseBodyReportTemplateList) GetSubscriptionFrequency() *string {
	return s.SubscriptionFrequency
}

func (s *ListReportTemplatesResponseBodyReportTemplateList) SetReportFileFormats(v string) *ListReportTemplatesResponseBodyReportTemplateList {
	s.ReportFileFormats = &v
	return s
}

func (s *ListReportTemplatesResponseBodyReportTemplateList) SetReportGranularity(v string) *ListReportTemplatesResponseBodyReportTemplateList {
	s.ReportGranularity = &v
	return s
}

func (s *ListReportTemplatesResponseBodyReportTemplateList) SetReportLanguage(v string) *ListReportTemplatesResponseBodyReportTemplateList {
	s.ReportLanguage = &v
	return s
}

func (s *ListReportTemplatesResponseBodyReportTemplateList) SetReportScope(v []*ListReportTemplatesResponseBodyReportTemplateListReportScope) *ListReportTemplatesResponseBodyReportTemplateList {
	s.ReportScope = v
	return s
}

func (s *ListReportTemplatesResponseBodyReportTemplateList) SetReportTemplateDescription(v string) *ListReportTemplatesResponseBodyReportTemplateList {
	s.ReportTemplateDescription = &v
	return s
}

func (s *ListReportTemplatesResponseBodyReportTemplateList) SetReportTemplateId(v string) *ListReportTemplatesResponseBodyReportTemplateList {
	s.ReportTemplateId = &v
	return s
}

func (s *ListReportTemplatesResponseBodyReportTemplateList) SetReportTemplateName(v string) *ListReportTemplatesResponseBodyReportTemplateList {
	s.ReportTemplateName = &v
	return s
}

func (s *ListReportTemplatesResponseBodyReportTemplateList) SetSubscriptionFrequency(v string) *ListReportTemplatesResponseBodyReportTemplateList {
	s.SubscriptionFrequency = &v
	return s
}

func (s *ListReportTemplatesResponseBodyReportTemplateList) Validate() error {
	if s.ReportScope != nil {
		for _, item := range s.ReportScope {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListReportTemplatesResponseBodyReportTemplateListReportScope struct {
	// The key for the report scope. Supported values:
	//
	// - AggregatorId
	//
	// - CompliancePackId
	//
	// - RuleId
	//
	// example:
	//
	// RuleId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The matching logic. Only In is supported.
	//
	// example:
	//
	// In
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The value for the report scope. For multiple values of the same type — such as multiple rule IDs — separate them with commas (,).
	//
	// example:
	//
	// cr-1,cr-2
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListReportTemplatesResponseBodyReportTemplateListReportScope) String() string {
	return dara.Prettify(s)
}

func (s ListReportTemplatesResponseBodyReportTemplateListReportScope) GoString() string {
	return s.String()
}

func (s *ListReportTemplatesResponseBodyReportTemplateListReportScope) GetKey() *string {
	return s.Key
}

func (s *ListReportTemplatesResponseBodyReportTemplateListReportScope) GetMatchType() *string {
	return s.MatchType
}

func (s *ListReportTemplatesResponseBodyReportTemplateListReportScope) GetValue() *string {
	return s.Value
}

func (s *ListReportTemplatesResponseBodyReportTemplateListReportScope) SetKey(v string) *ListReportTemplatesResponseBodyReportTemplateListReportScope {
	s.Key = &v
	return s
}

func (s *ListReportTemplatesResponseBodyReportTemplateListReportScope) SetMatchType(v string) *ListReportTemplatesResponseBodyReportTemplateListReportScope {
	s.MatchType = &v
	return s
}

func (s *ListReportTemplatesResponseBodyReportTemplateListReportScope) SetValue(v string) *ListReportTemplatesResponseBodyReportTemplateListReportScope {
	s.Value = &v
	return s
}

func (s *ListReportTemplatesResponseBodyReportTemplateListReportScope) Validate() error {
	return dara.Validate(s)
}
