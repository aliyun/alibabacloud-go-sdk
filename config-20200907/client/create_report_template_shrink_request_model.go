// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReportTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReportFileFormats(v string) *CreateReportTemplateShrinkRequest
	GetReportFileFormats() *string
	SetReportGranularity(v string) *CreateReportTemplateShrinkRequest
	GetReportGranularity() *string
	SetReportLanguage(v string) *CreateReportTemplateShrinkRequest
	GetReportLanguage() *string
	SetReportScopeShrink(v string) *CreateReportTemplateShrinkRequest
	GetReportScopeShrink() *string
	SetReportTemplateDescription(v string) *CreateReportTemplateShrinkRequest
	GetReportTemplateDescription() *string
	SetReportTemplateName(v string) *CreateReportTemplateShrinkRequest
	GetReportTemplateName() *string
	SetSubscriptionFrequency(v string) *CreateReportTemplateShrinkRequest
	GetSubscriptionFrequency() *string
}

type CreateReportTemplateShrinkRequest struct {
	// Report format. Currently, only Excel is supported.
	//
	// example:
	//
	// excel
	ReportFileFormats *string `json:"ReportFileFormats,omitempty" xml:"ReportFileFormats,omitempty"`
	// Report aggregation granularity.
	//
	// Supported for management accounts:
	//
	// - AllInOne (All accounts within the template scope are aggregated into one report)
	//
	// - GroupByAggregator (Reports are aggregated by account group, generating a compressed file)
	//
	// - GroupByAccount (Each account generates a separate report (default), generating a compressed file)
	//
	// Member accounts only support GroupByAccount.
	//
	// example:
	//
	// GroupByAccount
	ReportGranularity *string `json:"ReportGranularity,omitempty" xml:"ReportGranularity,omitempty"`
	// Report language. Supports zh-CN and en-US. If empty, the default is en-US.
	//
	// example:
	//
	// zh-CN
	ReportLanguage *string `json:"ReportLanguage,omitempty" xml:"ReportLanguage,omitempty"`
	// An array of report scopes, used to select the range of rules included in the audit report. The logic between each ReportScope in the array is OR, which means additive logic.
	//
	// > For example, if the array size is 2, the first ReportScope specifies rule In cr-1, and the second ReportScope specifies rule In cr-2, then the rule scope defined by this report is cr-1 and cr-2.
	ReportScopeShrink *string `json:"ReportScope,omitempty" xml:"ReportScope,omitempty"`
	// Report template description
	//
	// example:
	//
	// test-report-description
	ReportTemplateDescription *string `json:"ReportTemplateDescription,omitempty" xml:"ReportTemplateDescription,omitempty"`
	// Report template name
	//
	// This parameter is required.
	//
	// example:
	//
	// test-report-name
	ReportTemplateName *string `json:"ReportTemplateName,omitempty" xml:"ReportTemplateName,omitempty"`
	// Report subscription frequency. If this field is not empty, it is a Quartz-format Cron expression that triggers subscription notifications.
	//
	// Format: second minute hour day month week. The following are common Cron expression examples:
	//
	// - Execute at 0:00 every day: 0 0 0 \\	- \\	- ?
	//
	// - Execute at 15:30 every Monday: 0 30 15 ? \\	- MON
	//
	// - Execute at 2:00 on the 1st of every month: 0 0 2 1 \\	- ?
	//
	// Where:
	//
	// - "\\*" indicates any value
	//
	// - "?" is used for day and week fields, indicating no specific value
	//
	// - "MON" indicates Monday
	//
	// > The trigger time is UTC+8. Adjust the Cron expression settings based on your time zone.
	//
	// > We try to trigger notifications according to the Cron expression time, but there might be delays due to report generation status. A Cron expression limits the same template to trigger notifications at most once per day.
	//
	// > 1 represents Sunday; 2 represents Monday; 3 represents Tuesday; 4 represents Wednesday; 5 represents Thursday; 6 represents Friday; 7 represents Saturday
	//
	// example:
	//
	// 0 0 9 	- 	- ?
	SubscriptionFrequency *string `json:"SubscriptionFrequency,omitempty" xml:"SubscriptionFrequency,omitempty"`
}

func (s CreateReportTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateReportTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateReportTemplateShrinkRequest) GetReportFileFormats() *string {
	return s.ReportFileFormats
}

func (s *CreateReportTemplateShrinkRequest) GetReportGranularity() *string {
	return s.ReportGranularity
}

func (s *CreateReportTemplateShrinkRequest) GetReportLanguage() *string {
	return s.ReportLanguage
}

func (s *CreateReportTemplateShrinkRequest) GetReportScopeShrink() *string {
	return s.ReportScopeShrink
}

func (s *CreateReportTemplateShrinkRequest) GetReportTemplateDescription() *string {
	return s.ReportTemplateDescription
}

func (s *CreateReportTemplateShrinkRequest) GetReportTemplateName() *string {
	return s.ReportTemplateName
}

func (s *CreateReportTemplateShrinkRequest) GetSubscriptionFrequency() *string {
	return s.SubscriptionFrequency
}

func (s *CreateReportTemplateShrinkRequest) SetReportFileFormats(v string) *CreateReportTemplateShrinkRequest {
	s.ReportFileFormats = &v
	return s
}

func (s *CreateReportTemplateShrinkRequest) SetReportGranularity(v string) *CreateReportTemplateShrinkRequest {
	s.ReportGranularity = &v
	return s
}

func (s *CreateReportTemplateShrinkRequest) SetReportLanguage(v string) *CreateReportTemplateShrinkRequest {
	s.ReportLanguage = &v
	return s
}

func (s *CreateReportTemplateShrinkRequest) SetReportScopeShrink(v string) *CreateReportTemplateShrinkRequest {
	s.ReportScopeShrink = &v
	return s
}

func (s *CreateReportTemplateShrinkRequest) SetReportTemplateDescription(v string) *CreateReportTemplateShrinkRequest {
	s.ReportTemplateDescription = &v
	return s
}

func (s *CreateReportTemplateShrinkRequest) SetReportTemplateName(v string) *CreateReportTemplateShrinkRequest {
	s.ReportTemplateName = &v
	return s
}

func (s *CreateReportTemplateShrinkRequest) SetSubscriptionFrequency(v string) *CreateReportTemplateShrinkRequest {
	s.SubscriptionFrequency = &v
	return s
}

func (s *CreateReportTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
