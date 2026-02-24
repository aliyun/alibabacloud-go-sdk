// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateReportTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReportFileFormats(v string) *UpdateReportTemplateShrinkRequest
	GetReportFileFormats() *string
	SetReportGranularity(v string) *UpdateReportTemplateShrinkRequest
	GetReportGranularity() *string
	SetReportLanguage(v string) *UpdateReportTemplateShrinkRequest
	GetReportLanguage() *string
	SetReportScopeShrink(v string) *UpdateReportTemplateShrinkRequest
	GetReportScopeShrink() *string
	SetReportTemplateDescription(v string) *UpdateReportTemplateShrinkRequest
	GetReportTemplateDescription() *string
	SetReportTemplateId(v string) *UpdateReportTemplateShrinkRequest
	GetReportTemplateId() *string
	SetReportTemplateName(v string) *UpdateReportTemplateShrinkRequest
	GetReportTemplateName() *string
	SetSubscriptionFrequency(v string) *UpdateReportTemplateShrinkRequest
	GetSubscriptionFrequency() *string
}

type UpdateReportTemplateShrinkRequest struct {
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
	ReportScopeShrink *string `json:"ReportScope,omitempty" xml:"ReportScope,omitempty"`
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

func (s UpdateReportTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateReportTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateReportTemplateShrinkRequest) GetReportFileFormats() *string {
	return s.ReportFileFormats
}

func (s *UpdateReportTemplateShrinkRequest) GetReportGranularity() *string {
	return s.ReportGranularity
}

func (s *UpdateReportTemplateShrinkRequest) GetReportLanguage() *string {
	return s.ReportLanguage
}

func (s *UpdateReportTemplateShrinkRequest) GetReportScopeShrink() *string {
	return s.ReportScopeShrink
}

func (s *UpdateReportTemplateShrinkRequest) GetReportTemplateDescription() *string {
	return s.ReportTemplateDescription
}

func (s *UpdateReportTemplateShrinkRequest) GetReportTemplateId() *string {
	return s.ReportTemplateId
}

func (s *UpdateReportTemplateShrinkRequest) GetReportTemplateName() *string {
	return s.ReportTemplateName
}

func (s *UpdateReportTemplateShrinkRequest) GetSubscriptionFrequency() *string {
	return s.SubscriptionFrequency
}

func (s *UpdateReportTemplateShrinkRequest) SetReportFileFormats(v string) *UpdateReportTemplateShrinkRequest {
	s.ReportFileFormats = &v
	return s
}

func (s *UpdateReportTemplateShrinkRequest) SetReportGranularity(v string) *UpdateReportTemplateShrinkRequest {
	s.ReportGranularity = &v
	return s
}

func (s *UpdateReportTemplateShrinkRequest) SetReportLanguage(v string) *UpdateReportTemplateShrinkRequest {
	s.ReportLanguage = &v
	return s
}

func (s *UpdateReportTemplateShrinkRequest) SetReportScopeShrink(v string) *UpdateReportTemplateShrinkRequest {
	s.ReportScopeShrink = &v
	return s
}

func (s *UpdateReportTemplateShrinkRequest) SetReportTemplateDescription(v string) *UpdateReportTemplateShrinkRequest {
	s.ReportTemplateDescription = &v
	return s
}

func (s *UpdateReportTemplateShrinkRequest) SetReportTemplateId(v string) *UpdateReportTemplateShrinkRequest {
	s.ReportTemplateId = &v
	return s
}

func (s *UpdateReportTemplateShrinkRequest) SetReportTemplateName(v string) *UpdateReportTemplateShrinkRequest {
	s.ReportTemplateName = &v
	return s
}

func (s *UpdateReportTemplateShrinkRequest) SetSubscriptionFrequency(v string) *UpdateReportTemplateShrinkRequest {
	s.SubscriptionFrequency = &v
	return s
}

func (s *UpdateReportTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
