// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReportTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetReportTemplate(v *GetReportTemplateResponseBodyReportTemplate) *GetReportTemplateResponseBody
	GetReportTemplate() *GetReportTemplateResponseBodyReportTemplate
	SetRequestId(v string) *GetReportTemplateResponseBody
	GetRequestId() *string
}

type GetReportTemplateResponseBody struct {
	// Report template.
	ReportTemplate *GetReportTemplateResponseBodyReportTemplate `json:"ReportTemplate,omitempty" xml:"ReportTemplate,omitempty" type:"Struct"`
	// Request ID.
	//
	// example:
	//
	// A7A0FFF8-0B44-40C6-8BBF-3A185EFDF3F7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetReportTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetReportTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetReportTemplateResponseBody) GetReportTemplate() *GetReportTemplateResponseBodyReportTemplate {
	return s.ReportTemplate
}

func (s *GetReportTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetReportTemplateResponseBody) SetReportTemplate(v *GetReportTemplateResponseBodyReportTemplate) *GetReportTemplateResponseBody {
	s.ReportTemplate = v
	return s
}

func (s *GetReportTemplateResponseBody) SetRequestId(v string) *GetReportTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetReportTemplateResponseBody) Validate() error {
	if s.ReportTemplate != nil {
		if err := s.ReportTemplate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetReportTemplateResponseBodyReportTemplate struct {
	// Report file format.
	//
	// example:
	//
	// excel
	ReportFileFormats *string `json:"ReportFileFormats,omitempty" xml:"ReportFileFormats,omitempty"`
	// Aggregation granularity of the report.
	//
	// example:
	//
	// AllInOne
	ReportGranularity *string `json:"ReportGranularity,omitempty" xml:"ReportGranularity,omitempty"`
	// Report language. Valid values: zh-CN and en-US. Default is en-US if empty.
	//
	// example:
	//
	// en-US
	ReportLanguage *string `json:"ReportLanguage,omitempty" xml:"ReportLanguage,omitempty"`
	// Array of report scopes. Each scope defines a set of rules included in the audit report. Scopes use OR logic. That is, rules from all scopes are combined.
	//
	// > If the array has two items, and the first specifies RuleId cr-1 while the second specifies RuleId cr-2, then the report covers both cr-1 and cr-2.
	ReportScope []*GetReportTemplateResponseBodyReportTemplateReportScope `json:"ReportScope,omitempty" xml:"ReportScope,omitempty" type:"Repeated"`
	// Description of the report template.
	//
	// example:
	//
	// test-description
	ReportTemplateDescription *string `json:"ReportTemplateDescription,omitempty" xml:"ReportTemplateDescription,omitempty"`
	// ID of the report template.
	//
	// example:
	//
	// crt-xxx
	ReportTemplateId *string `json:"ReportTemplateId,omitempty" xml:"ReportTemplateId,omitempty"`
	// Name of the report template.
	//
	// example:
	//
	// test-name
	ReportTemplateName *string `json:"ReportTemplateName,omitempty" xml:"ReportTemplateName,omitempty"`
	// Subscription frequency for the report. If this field is not empty, it contains a Quartz-formatted cron expression that triggers notifications.
	//
	// The format is: seconds minutes hours day-of-month month day-of-week. Common examples include the following:
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
	// - "?" means no specific value for the day-of-month or day-of-week field.
	//
	// - MON means Monday.
	//
	// > Trigger times are in UTC+8. Adjust your cron expression based on your time zone.
	//
	// > The system tries to trigger notifications as close as possible to the scheduled time. Delays may occur due to report generation status. A single template can trigger at most one notification per day.
	//
	// > In Quartz, days of the week are numbered: 1 = Sunday, 2 = Monday, 3 = Tuesday, 4 = Wednesday, 5 = Thursday, 6 = Friday, 7 = Saturday.
	//
	// example:
	//
	// 0 0 0 	- 	- ?
	SubscriptionFrequency *string `json:"SubscriptionFrequency,omitempty" xml:"SubscriptionFrequency,omitempty"`
}

func (s GetReportTemplateResponseBodyReportTemplate) String() string {
	return dara.Prettify(s)
}

func (s GetReportTemplateResponseBodyReportTemplate) GoString() string {
	return s.String()
}

func (s *GetReportTemplateResponseBodyReportTemplate) GetReportFileFormats() *string {
	return s.ReportFileFormats
}

func (s *GetReportTemplateResponseBodyReportTemplate) GetReportGranularity() *string {
	return s.ReportGranularity
}

func (s *GetReportTemplateResponseBodyReportTemplate) GetReportLanguage() *string {
	return s.ReportLanguage
}

func (s *GetReportTemplateResponseBodyReportTemplate) GetReportScope() []*GetReportTemplateResponseBodyReportTemplateReportScope {
	return s.ReportScope
}

func (s *GetReportTemplateResponseBodyReportTemplate) GetReportTemplateDescription() *string {
	return s.ReportTemplateDescription
}

func (s *GetReportTemplateResponseBodyReportTemplate) GetReportTemplateId() *string {
	return s.ReportTemplateId
}

func (s *GetReportTemplateResponseBodyReportTemplate) GetReportTemplateName() *string {
	return s.ReportTemplateName
}

func (s *GetReportTemplateResponseBodyReportTemplate) GetSubscriptionFrequency() *string {
	return s.SubscriptionFrequency
}

func (s *GetReportTemplateResponseBodyReportTemplate) SetReportFileFormats(v string) *GetReportTemplateResponseBodyReportTemplate {
	s.ReportFileFormats = &v
	return s
}

func (s *GetReportTemplateResponseBodyReportTemplate) SetReportGranularity(v string) *GetReportTemplateResponseBodyReportTemplate {
	s.ReportGranularity = &v
	return s
}

func (s *GetReportTemplateResponseBodyReportTemplate) SetReportLanguage(v string) *GetReportTemplateResponseBodyReportTemplate {
	s.ReportLanguage = &v
	return s
}

func (s *GetReportTemplateResponseBodyReportTemplate) SetReportScope(v []*GetReportTemplateResponseBodyReportTemplateReportScope) *GetReportTemplateResponseBodyReportTemplate {
	s.ReportScope = v
	return s
}

func (s *GetReportTemplateResponseBodyReportTemplate) SetReportTemplateDescription(v string) *GetReportTemplateResponseBodyReportTemplate {
	s.ReportTemplateDescription = &v
	return s
}

func (s *GetReportTemplateResponseBodyReportTemplate) SetReportTemplateId(v string) *GetReportTemplateResponseBodyReportTemplate {
	s.ReportTemplateId = &v
	return s
}

func (s *GetReportTemplateResponseBodyReportTemplate) SetReportTemplateName(v string) *GetReportTemplateResponseBodyReportTemplate {
	s.ReportTemplateName = &v
	return s
}

func (s *GetReportTemplateResponseBodyReportTemplate) SetSubscriptionFrequency(v string) *GetReportTemplateResponseBodyReportTemplate {
	s.SubscriptionFrequency = &v
	return s
}

func (s *GetReportTemplateResponseBodyReportTemplate) Validate() error {
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

type GetReportTemplateResponseBodyReportTemplateReportScope struct {
	// Key for the report scope. Supported keys:
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
	// Matching logic. Only In is supported.
	//
	// example:
	//
	// In
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// Value for the report scope. For multiple values of the same type, such as multiple rule IDs, separate them with commas.
	//
	// example:
	//
	// cr-1,cr-2
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetReportTemplateResponseBodyReportTemplateReportScope) String() string {
	return dara.Prettify(s)
}

func (s GetReportTemplateResponseBodyReportTemplateReportScope) GoString() string {
	return s.String()
}

func (s *GetReportTemplateResponseBodyReportTemplateReportScope) GetKey() *string {
	return s.Key
}

func (s *GetReportTemplateResponseBodyReportTemplateReportScope) GetMatchType() *string {
	return s.MatchType
}

func (s *GetReportTemplateResponseBodyReportTemplateReportScope) GetValue() *string {
	return s.Value
}

func (s *GetReportTemplateResponseBodyReportTemplateReportScope) SetKey(v string) *GetReportTemplateResponseBodyReportTemplateReportScope {
	s.Key = &v
	return s
}

func (s *GetReportTemplateResponseBodyReportTemplateReportScope) SetMatchType(v string) *GetReportTemplateResponseBodyReportTemplateReportScope {
	s.MatchType = &v
	return s
}

func (s *GetReportTemplateResponseBodyReportTemplateReportScope) SetValue(v string) *GetReportTemplateResponseBodyReportTemplateReportScope {
	s.Value = &v
	return s
}

func (s *GetReportTemplateResponseBodyReportTemplateReportScope) Validate() error {
	return dara.Validate(s)
}
