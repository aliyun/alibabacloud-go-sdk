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
	// example:
	//
	// excel
	ReportFileFormats *string `json:"ReportFileFormats,omitempty" xml:"ReportFileFormats,omitempty"`
	// example:
	//
	// AllInOne
	ReportGranularity *string `json:"ReportGranularity,omitempty" xml:"ReportGranularity,omitempty"`
	ReportLanguage    *string `json:"ReportLanguage,omitempty" xml:"ReportLanguage,omitempty"`
	ReportScopeShrink *string `json:"ReportScope,omitempty" xml:"ReportScope,omitempty"`
	// example:
	//
	// test-description
	ReportTemplateDescription *string `json:"ReportTemplateDescription,omitempty" xml:"ReportTemplateDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// crt-xxx
	ReportTemplateId *string `json:"ReportTemplateId,omitempty" xml:"ReportTemplateId,omitempty"`
	// example:
	//
	// test-name
	ReportTemplateName *string `json:"ReportTemplateName,omitempty" xml:"ReportTemplateName,omitempty"`
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
