// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateReportTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReportFileFormats(v string) *UpdateReportTemplateRequest
	GetReportFileFormats() *string
	SetReportGranularity(v string) *UpdateReportTemplateRequest
	GetReportGranularity() *string
	SetReportLanguage(v string) *UpdateReportTemplateRequest
	GetReportLanguage() *string
	SetReportScope(v []*UpdateReportTemplateRequestReportScope) *UpdateReportTemplateRequest
	GetReportScope() []*UpdateReportTemplateRequestReportScope
	SetReportTemplateDescription(v string) *UpdateReportTemplateRequest
	GetReportTemplateDescription() *string
	SetReportTemplateId(v string) *UpdateReportTemplateRequest
	GetReportTemplateId() *string
	SetReportTemplateName(v string) *UpdateReportTemplateRequest
	GetReportTemplateName() *string
	SetSubscriptionFrequency(v string) *UpdateReportTemplateRequest
	GetSubscriptionFrequency() *string
}

type UpdateReportTemplateRequest struct {
	// The format of the report.
	//
	// example:
	//
	// excel
	ReportFileFormats *string `json:"ReportFileFormats,omitempty" xml:"ReportFileFormats,omitempty"`
	// The aggregation granularity of the report.
	//
	// For a management account, the following values are supported:
	//
	// - AllInOne: A single report is generated for all accounts within the template scope.
	//
	// - GroupByAggregator: Reports are aggregated by account group. A compressed file is generated.
	//
	// - GroupByAccount: A separate report is generated for each account. This is the default value. A compressed file is generated.
	//
	// For a member account, only GroupByAccount is supported.
	//
	// example:
	//
	// AllInOne
	ReportGranularity *string `json:"ReportGranularity,omitempty" xml:"ReportGranularity,omitempty"`
	// The language of the report. Valid values: zh-CN and en-US. If you leave this parameter empty, the default value en-US is used.
	//
	// example:
	//
	// en-US
	ReportLanguage *string `json:"ReportLanguage,omitempty" xml:"ReportLanguage,omitempty"`
	// An array that specifies the report scope. It is used to select the range of rules to include in the audit report. The logical relationship between multiple ReportScope objects in the array is OR. This means the scopes are added together.
	//
	// > For example, if the array contains two ReportScope objects, where the first specifies the rule In cr-1 and the second specifies the rule In cr-2, the report scope includes both cr-1 and cr-2.
	ReportScope []*UpdateReportTemplateRequestReportScope `json:"ReportScope,omitempty" xml:"ReportScope,omitempty" type:"Repeated"`
	// The description of the report template.
	//
	// example:
	//
	// test-description
	ReportTemplateDescription *string `json:"ReportTemplateDescription,omitempty" xml:"ReportTemplateDescription,omitempty"`
	// The ID of the report template.
	//
	// This parameter is required.
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
	// The frequency for subscribing to the report. If this parameter is not empty, it specifies a cron expression in the Quartz format that triggers a subscription notification.
	//
	// The format is: Second Minute Hour Day Month Week. The following list provides common examples of cron expressions:
	//
	// - To run at 00:00 every day: 0 0 0 \\	- \\	- ?
	//
	// - To run at 15:30 every Monday: 0 30 15 ? \\	- MON
	//
	// - To run at 02:00 on the first day of every month: 0 0 2 1 \\	- ?
	//
	// Where:
	//
	// - "\\*" indicates any value.
	//
	// - ? is used in the Day and Week fields and indicates that no specific value is specified.
	//
	// - MON indicates Monday.
	//
	// > The trigger time is in UTC+8. You can convert the cron expression based on your time zone.
	//
	// > The system attempts to trigger the notification at the time specified by the cron expression. However, a delay may occur due to the report generation status. The cron expression limits the notification for the same template to a maximum of once per day.
	//
	// > In addition to using MON for Monday, you can also use numbers. In the Quartz framework, 1 represents Sunday, 2 represents Monday, 3 represents Tuesday, 4 represents Wednesday, 5 represents Thursday, 6 represents Friday, and 7 represents Saturday.
	//
	// example:
	//
	// 0 0 0 	- 	- ?
	SubscriptionFrequency *string `json:"SubscriptionFrequency,omitempty" xml:"SubscriptionFrequency,omitempty"`
}

func (s UpdateReportTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateReportTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateReportTemplateRequest) GetReportFileFormats() *string {
	return s.ReportFileFormats
}

func (s *UpdateReportTemplateRequest) GetReportGranularity() *string {
	return s.ReportGranularity
}

func (s *UpdateReportTemplateRequest) GetReportLanguage() *string {
	return s.ReportLanguage
}

func (s *UpdateReportTemplateRequest) GetReportScope() []*UpdateReportTemplateRequestReportScope {
	return s.ReportScope
}

func (s *UpdateReportTemplateRequest) GetReportTemplateDescription() *string {
	return s.ReportTemplateDescription
}

func (s *UpdateReportTemplateRequest) GetReportTemplateId() *string {
	return s.ReportTemplateId
}

func (s *UpdateReportTemplateRequest) GetReportTemplateName() *string {
	return s.ReportTemplateName
}

func (s *UpdateReportTemplateRequest) GetSubscriptionFrequency() *string {
	return s.SubscriptionFrequency
}

func (s *UpdateReportTemplateRequest) SetReportFileFormats(v string) *UpdateReportTemplateRequest {
	s.ReportFileFormats = &v
	return s
}

func (s *UpdateReportTemplateRequest) SetReportGranularity(v string) *UpdateReportTemplateRequest {
	s.ReportGranularity = &v
	return s
}

func (s *UpdateReportTemplateRequest) SetReportLanguage(v string) *UpdateReportTemplateRequest {
	s.ReportLanguage = &v
	return s
}

func (s *UpdateReportTemplateRequest) SetReportScope(v []*UpdateReportTemplateRequestReportScope) *UpdateReportTemplateRequest {
	s.ReportScope = v
	return s
}

func (s *UpdateReportTemplateRequest) SetReportTemplateDescription(v string) *UpdateReportTemplateRequest {
	s.ReportTemplateDescription = &v
	return s
}

func (s *UpdateReportTemplateRequest) SetReportTemplateId(v string) *UpdateReportTemplateRequest {
	s.ReportTemplateId = &v
	return s
}

func (s *UpdateReportTemplateRequest) SetReportTemplateName(v string) *UpdateReportTemplateRequest {
	s.ReportTemplateName = &v
	return s
}

func (s *UpdateReportTemplateRequest) SetSubscriptionFrequency(v string) *UpdateReportTemplateRequest {
	s.SubscriptionFrequency = &v
	return s
}

func (s *UpdateReportTemplateRequest) Validate() error {
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

type UpdateReportTemplateRequestReportScope struct {
	// The key of the report scope. Supported values:
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
	// The value of the report scope. To specify multiple items of the same type, such as multiple rule IDs, separate them with commas (,).
	//
	// example:
	//
	// cr-1,cr-2
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateReportTemplateRequestReportScope) String() string {
	return dara.Prettify(s)
}

func (s UpdateReportTemplateRequestReportScope) GoString() string {
	return s.String()
}

func (s *UpdateReportTemplateRequestReportScope) GetKey() *string {
	return s.Key
}

func (s *UpdateReportTemplateRequestReportScope) GetMatchType() *string {
	return s.MatchType
}

func (s *UpdateReportTemplateRequestReportScope) GetValue() *string {
	return s.Value
}

func (s *UpdateReportTemplateRequestReportScope) SetKey(v string) *UpdateReportTemplateRequestReportScope {
	s.Key = &v
	return s
}

func (s *UpdateReportTemplateRequestReportScope) SetMatchType(v string) *UpdateReportTemplateRequestReportScope {
	s.MatchType = &v
	return s
}

func (s *UpdateReportTemplateRequestReportScope) SetValue(v string) *UpdateReportTemplateRequestReportScope {
	s.Value = &v
	return s
}

func (s *UpdateReportTemplateRequestReportScope) Validate() error {
	return dara.Validate(s)
}
