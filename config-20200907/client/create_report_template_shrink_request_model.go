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
	// example:
	//
	// excel
	ReportFileFormats *string `json:"ReportFileFormats,omitempty" xml:"ReportFileFormats,omitempty"`
	// example:
	//
	// ReportGranularity
	ReportGranularity *string `json:"ReportGranularity,omitempty" xml:"ReportGranularity,omitempty"`
	// example:
	//
	// zh-CN
	ReportLanguage    *string `json:"ReportLanguage,omitempty" xml:"ReportLanguage,omitempty"`
	ReportScopeShrink *string `json:"ReportScope,omitempty" xml:"ReportScope,omitempty"`
	// example:
	//
	// test-report-description
	ReportTemplateDescription *string `json:"ReportTemplateDescription,omitempty" xml:"ReportTemplateDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-report-name
	ReportTemplateName *string `json:"ReportTemplateName,omitempty" xml:"ReportTemplateName,omitempty"`
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
