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
	// example:
	//
	// excel
	ReportFileFormats *string `json:"ReportFileFormats,omitempty" xml:"ReportFileFormats,omitempty"`
	// example:
	//
	// AllInOne
	ReportGranularity *string                                   `json:"ReportGranularity,omitempty" xml:"ReportGranularity,omitempty"`
	ReportLanguage    *string                                   `json:"ReportLanguage,omitempty" xml:"ReportLanguage,omitempty"`
	ReportScope       []*UpdateReportTemplateRequestReportScope `json:"ReportScope,omitempty" xml:"ReportScope,omitempty" type:"Repeated"`
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
