// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReportTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReportFileFormats(v string) *CreateReportTemplateRequest
	GetReportFileFormats() *string
	SetReportGranularity(v string) *CreateReportTemplateRequest
	GetReportGranularity() *string
	SetReportLanguage(v string) *CreateReportTemplateRequest
	GetReportLanguage() *string
	SetReportScope(v []*CreateReportTemplateRequestReportScope) *CreateReportTemplateRequest
	GetReportScope() []*CreateReportTemplateRequestReportScope
	SetReportTemplateDescription(v string) *CreateReportTemplateRequest
	GetReportTemplateDescription() *string
	SetReportTemplateName(v string) *CreateReportTemplateRequest
	GetReportTemplateName() *string
	SetSubscriptionFrequency(v string) *CreateReportTemplateRequest
	GetSubscriptionFrequency() *string
}

type CreateReportTemplateRequest struct {
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
	ReportScope []*CreateReportTemplateRequestReportScope `json:"ReportScope,omitempty" xml:"ReportScope,omitempty" type:"Repeated"`
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

func (s CreateReportTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateReportTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateReportTemplateRequest) GetReportFileFormats() *string {
	return s.ReportFileFormats
}

func (s *CreateReportTemplateRequest) GetReportGranularity() *string {
	return s.ReportGranularity
}

func (s *CreateReportTemplateRequest) GetReportLanguage() *string {
	return s.ReportLanguage
}

func (s *CreateReportTemplateRequest) GetReportScope() []*CreateReportTemplateRequestReportScope {
	return s.ReportScope
}

func (s *CreateReportTemplateRequest) GetReportTemplateDescription() *string {
	return s.ReportTemplateDescription
}

func (s *CreateReportTemplateRequest) GetReportTemplateName() *string {
	return s.ReportTemplateName
}

func (s *CreateReportTemplateRequest) GetSubscriptionFrequency() *string {
	return s.SubscriptionFrequency
}

func (s *CreateReportTemplateRequest) SetReportFileFormats(v string) *CreateReportTemplateRequest {
	s.ReportFileFormats = &v
	return s
}

func (s *CreateReportTemplateRequest) SetReportGranularity(v string) *CreateReportTemplateRequest {
	s.ReportGranularity = &v
	return s
}

func (s *CreateReportTemplateRequest) SetReportLanguage(v string) *CreateReportTemplateRequest {
	s.ReportLanguage = &v
	return s
}

func (s *CreateReportTemplateRequest) SetReportScope(v []*CreateReportTemplateRequestReportScope) *CreateReportTemplateRequest {
	s.ReportScope = v
	return s
}

func (s *CreateReportTemplateRequest) SetReportTemplateDescription(v string) *CreateReportTemplateRequest {
	s.ReportTemplateDescription = &v
	return s
}

func (s *CreateReportTemplateRequest) SetReportTemplateName(v string) *CreateReportTemplateRequest {
	s.ReportTemplateName = &v
	return s
}

func (s *CreateReportTemplateRequest) SetSubscriptionFrequency(v string) *CreateReportTemplateRequest {
	s.SubscriptionFrequency = &v
	return s
}

func (s *CreateReportTemplateRequest) Validate() error {
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

type CreateReportTemplateRequestReportScope struct {
	// The key for the report scope. Currently supports:
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
	// Matching logic. Currently, only In is supported.
	//
	// example:
	//
	// In
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The value for the report scope. Multiple items of the same type, such as multiple Rule IDs, can be separated by English commas (,).
	//
	// example:
	//
	// cr-1,cr-2
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateReportTemplateRequestReportScope) String() string {
	return dara.Prettify(s)
}

func (s CreateReportTemplateRequestReportScope) GoString() string {
	return s.String()
}

func (s *CreateReportTemplateRequestReportScope) GetKey() *string {
	return s.Key
}

func (s *CreateReportTemplateRequestReportScope) GetMatchType() *string {
	return s.MatchType
}

func (s *CreateReportTemplateRequestReportScope) GetValue() *string {
	return s.Value
}

func (s *CreateReportTemplateRequestReportScope) SetKey(v string) *CreateReportTemplateRequestReportScope {
	s.Key = &v
	return s
}

func (s *CreateReportTemplateRequestReportScope) SetMatchType(v string) *CreateReportTemplateRequestReportScope {
	s.MatchType = &v
	return s
}

func (s *CreateReportTemplateRequestReportScope) SetValue(v string) *CreateReportTemplateRequestReportScope {
	s.Value = &v
	return s
}

func (s *CreateReportTemplateRequestReportScope) Validate() error {
	return dara.Validate(s)
}
