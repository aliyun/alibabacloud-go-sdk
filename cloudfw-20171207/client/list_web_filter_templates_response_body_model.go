// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWebFilterTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int64) *ListWebFilterTemplatesResponseBody
	GetPageNo() *int64
	SetPageSize(v int64) *ListWebFilterTemplatesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListWebFilterTemplatesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListWebFilterTemplatesResponseBody
	GetTotalCount() *int64
	SetWebFilterTemplates(v []*ListWebFilterTemplatesResponseBodyWebFilterTemplates) *ListWebFilterTemplatesResponseBody
	GetWebFilterTemplates() []*ListWebFilterTemplatesResponseBodyWebFilterTemplates
}

type ListWebFilterTemplatesResponseBody struct {
	// The current page number.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of web filtering templates that meet the filter conditions.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of web filtering templates.
	WebFilterTemplates []*ListWebFilterTemplatesResponseBodyWebFilterTemplates `json:"WebFilterTemplates,omitempty" xml:"WebFilterTemplates,omitempty" type:"Repeated"`
}

func (s ListWebFilterTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWebFilterTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWebFilterTemplatesResponseBody) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListWebFilterTemplatesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListWebFilterTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWebFilterTemplatesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListWebFilterTemplatesResponseBody) GetWebFilterTemplates() []*ListWebFilterTemplatesResponseBodyWebFilterTemplates {
	return s.WebFilterTemplates
}

func (s *ListWebFilterTemplatesResponseBody) SetPageNo(v int64) *ListWebFilterTemplatesResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListWebFilterTemplatesResponseBody) SetPageSize(v int64) *ListWebFilterTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListWebFilterTemplatesResponseBody) SetRequestId(v string) *ListWebFilterTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWebFilterTemplatesResponseBody) SetTotalCount(v int64) *ListWebFilterTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListWebFilterTemplatesResponseBody) SetWebFilterTemplates(v []*ListWebFilterTemplatesResponseBodyWebFilterTemplates) *ListWebFilterTemplatesResponseBody {
	s.WebFilterTemplates = v
	return s
}

func (s *ListWebFilterTemplatesResponseBody) Validate() error {
	if s.WebFilterTemplates != nil {
		for _, item := range s.WebFilterTemplates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWebFilterTemplatesResponseBodyWebFilterTemplates struct {
	// The creation time of the web filtering template. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1731465600
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the web filtering template.
	//
	// example:
	//
	// Web filtering template for URL access control
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the web filtering template.
	//
	// example:
	//
	// Web filtering template 01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The statistics on the number of web filtering rules in the template categorized by action.
	RuleActionStatistics *ListWebFilterTemplatesResponseBodyWebFilterTemplatesRuleActionStatistics `json:"RuleActionStatistics,omitempty" xml:"RuleActionStatistics,omitempty" type:"Struct"`
	// The status of the web filtering template. Valid values:
	//
	// - **Enable**: Enabled.
	//
	// - **Disable**: Disabled.
	//
	// example:
	//
	// Enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the web filtering template.
	//
	// example:
	//
	// wft-3f2a9c1d4b5e6f70
	WebFilterTemplateId *string `json:"WebFilterTemplateId,omitempty" xml:"WebFilterTemplateId,omitempty"`
}

func (s ListWebFilterTemplatesResponseBodyWebFilterTemplates) String() string {
	return dara.Prettify(s)
}

func (s ListWebFilterTemplatesResponseBodyWebFilterTemplates) GoString() string {
	return s.String()
}

func (s *ListWebFilterTemplatesResponseBodyWebFilterTemplates) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListWebFilterTemplatesResponseBodyWebFilterTemplates) GetDescription() *string {
	return s.Description
}

func (s *ListWebFilterTemplatesResponseBodyWebFilterTemplates) GetName() *string {
	return s.Name
}

func (s *ListWebFilterTemplatesResponseBodyWebFilterTemplates) GetRuleActionStatistics() *ListWebFilterTemplatesResponseBodyWebFilterTemplatesRuleActionStatistics {
	return s.RuleActionStatistics
}

func (s *ListWebFilterTemplatesResponseBodyWebFilterTemplates) GetStatus() *string {
	return s.Status
}

func (s *ListWebFilterTemplatesResponseBodyWebFilterTemplates) GetWebFilterTemplateId() *string {
	return s.WebFilterTemplateId
}

func (s *ListWebFilterTemplatesResponseBodyWebFilterTemplates) SetCreateTime(v int64) *ListWebFilterTemplatesResponseBodyWebFilterTemplates {
	s.CreateTime = &v
	return s
}

func (s *ListWebFilterTemplatesResponseBodyWebFilterTemplates) SetDescription(v string) *ListWebFilterTemplatesResponseBodyWebFilterTemplates {
	s.Description = &v
	return s
}

func (s *ListWebFilterTemplatesResponseBodyWebFilterTemplates) SetName(v string) *ListWebFilterTemplatesResponseBodyWebFilterTemplates {
	s.Name = &v
	return s
}

func (s *ListWebFilterTemplatesResponseBodyWebFilterTemplates) SetRuleActionStatistics(v *ListWebFilterTemplatesResponseBodyWebFilterTemplatesRuleActionStatistics) *ListWebFilterTemplatesResponseBodyWebFilterTemplates {
	s.RuleActionStatistics = v
	return s
}

func (s *ListWebFilterTemplatesResponseBodyWebFilterTemplates) SetStatus(v string) *ListWebFilterTemplatesResponseBodyWebFilterTemplates {
	s.Status = &v
	return s
}

func (s *ListWebFilterTemplatesResponseBodyWebFilterTemplates) SetWebFilterTemplateId(v string) *ListWebFilterTemplatesResponseBodyWebFilterTemplates {
	s.WebFilterTemplateId = &v
	return s
}

func (s *ListWebFilterTemplatesResponseBodyWebFilterTemplates) Validate() error {
	if s.RuleActionStatistics != nil {
		if err := s.RuleActionStatistics.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListWebFilterTemplatesResponseBodyWebFilterTemplatesRuleActionStatistics struct {
	// The number of web filtering rules whose action is Alert in the template.
	//
	// example:
	//
	// 2
	Alert *int64 `json:"Alert,omitempty" xml:"Alert,omitempty"`
	// The number of web filtering rules whose action is Allow in the template.
	//
	// example:
	//
	// 3
	Allow *int64 `json:"Allow,omitempty" xml:"Allow,omitempty"`
	// The number of web filtering rules whose action is Deny in the template.
	//
	// example:
	//
	// 1
	Deny *int64 `json:"Deny,omitempty" xml:"Deny,omitempty"`
}

func (s ListWebFilterTemplatesResponseBodyWebFilterTemplatesRuleActionStatistics) String() string {
	return dara.Prettify(s)
}

func (s ListWebFilterTemplatesResponseBodyWebFilterTemplatesRuleActionStatistics) GoString() string {
	return s.String()
}

func (s *ListWebFilterTemplatesResponseBodyWebFilterTemplatesRuleActionStatistics) GetAlert() *int64 {
	return s.Alert
}

func (s *ListWebFilterTemplatesResponseBodyWebFilterTemplatesRuleActionStatistics) GetAllow() *int64 {
	return s.Allow
}

func (s *ListWebFilterTemplatesResponseBodyWebFilterTemplatesRuleActionStatistics) GetDeny() *int64 {
	return s.Deny
}

func (s *ListWebFilterTemplatesResponseBodyWebFilterTemplatesRuleActionStatistics) SetAlert(v int64) *ListWebFilterTemplatesResponseBodyWebFilterTemplatesRuleActionStatistics {
	s.Alert = &v
	return s
}

func (s *ListWebFilterTemplatesResponseBodyWebFilterTemplatesRuleActionStatistics) SetAllow(v int64) *ListWebFilterTemplatesResponseBodyWebFilterTemplatesRuleActionStatistics {
	s.Allow = &v
	return s
}

func (s *ListWebFilterTemplatesResponseBodyWebFilterTemplatesRuleActionStatistics) SetDeny(v int64) *ListWebFilterTemplatesResponseBodyWebFilterTemplatesRuleActionStatistics {
	s.Deny = &v
	return s
}

func (s *ListWebFilterTemplatesResponseBodyWebFilterTemplatesRuleActionStatistics) Validate() error {
	return dara.Validate(s)
}
