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
	ReportLanguage *string                                   `json:"ReportLanguage,omitempty" xml:"ReportLanguage,omitempty"`
	ReportScope    []*CreateReportTemplateRequestReportScope `json:"ReportScope,omitempty" xml:"ReportScope,omitempty" type:"Repeated"`
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
	// example:
	//
	// RuleId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// In
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
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
