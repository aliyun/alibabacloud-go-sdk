// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationTemplates(v []*ListApplicationTemplateResponseBodyApplicationTemplates) *ListApplicationTemplateResponseBody
	GetApplicationTemplates() []*ListApplicationTemplateResponseBodyApplicationTemplates
	SetPageNo(v int64) *ListApplicationTemplateResponseBody
	GetPageNo() *int64
	SetPageSize(v int64) *ListApplicationTemplateResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListApplicationTemplateResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListApplicationTemplateResponseBody
	GetTotalCount() *int64
}

type ListApplicationTemplateResponseBody struct {
	// The list of application templates.
	ApplicationTemplates []*ListApplicationTemplateResponseBodyApplicationTemplates `json:"ApplicationTemplates,omitempty" xml:"ApplicationTemplates,omitempty" type:"Repeated"`
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
	// The total number of application templates that meet the filter conditions.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApplicationTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationTemplateResponseBody) GetApplicationTemplates() []*ListApplicationTemplateResponseBodyApplicationTemplates {
	return s.ApplicationTemplates
}

func (s *ListApplicationTemplateResponseBody) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListApplicationTemplateResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListApplicationTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationTemplateResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListApplicationTemplateResponseBody) SetApplicationTemplates(v []*ListApplicationTemplateResponseBodyApplicationTemplates) *ListApplicationTemplateResponseBody {
	s.ApplicationTemplates = v
	return s
}

func (s *ListApplicationTemplateResponseBody) SetPageNo(v int64) *ListApplicationTemplateResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListApplicationTemplateResponseBody) SetPageSize(v int64) *ListApplicationTemplateResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListApplicationTemplateResponseBody) SetRequestId(v string) *ListApplicationTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationTemplateResponseBody) SetTotalCount(v int64) *ListApplicationTemplateResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListApplicationTemplateResponseBody) Validate() error {
	if s.ApplicationTemplates != nil {
		for _, item := range s.ApplicationTemplates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationTemplateResponseBodyApplicationTemplates struct {
	// The application template ID.
	//
	// example:
	//
	// ati-4f0691c56b8d4f06
	ApplicationTemplateId *string `json:"ApplicationTemplateId,omitempty" xml:"ApplicationTemplateId,omitempty"`
	// The name of the application template.
	//
	// example:
	//
	// Application template 01
	ApplicationTemplateName *string `json:"ApplicationTemplateName,omitempty" xml:"ApplicationTemplateName,omitempty"`
	// The statistics on the number of built-in applications in the template, categorized by the final effective action.
	BuiltInApplicationActionStatistics *ListApplicationTemplateResponseBodyApplicationTemplatesBuiltInApplicationActionStatistics `json:"BuiltInApplicationActionStatistics,omitempty" xml:"BuiltInApplicationActionStatistics,omitempty" type:"Struct"`
	// The switch status of built-in application control. Valid values:
	//
	// - **Enable**: enabled
	//
	// - **Disable**: disabled
	//
	// example:
	//
	// Enable
	BuiltInApplicationStatus *string `json:"BuiltInApplicationStatus,omitempty" xml:"BuiltInApplicationStatus,omitempty"`
	// The creation time of the application template. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1731465600
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The number of custom application rules under the template.
	//
	// example:
	//
	// 3
	CustomRuleCount *int64 `json:"CustomRuleCount,omitempty" xml:"CustomRuleCount,omitempty"`
	// The switch status of custom application rules. Valid values:
	//
	// - **Enable**: enabled
	//
	// - **Disable**: disabled
	//
	// example:
	//
	// Enable
	CustomRuleStatus *string `json:"CustomRuleStatus,omitempty" xml:"CustomRuleStatus,omitempty"`
	// The description of the application template.
	//
	// example:
	//
	// Office network application control
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The master switch status of the application template. Valid values:
	//
	// - **Enable**: enabled
	//
	// - **Disable**: disabled
	//
	// example:
	//
	// Enable
	TemplateStatus *string `json:"TemplateStatus,omitempty" xml:"TemplateStatus,omitempty"`
	// The action for unrecognized applications. Valid values:
	//
	// - **Allow**: allow
	//
	// - **Alert**: monitor
	//
	// - **Deny**: block
	//
	// example:
	//
	// Alert
	UnknownAction *string `json:"UnknownAction,omitempty" xml:"UnknownAction,omitempty"`
	// The last modification time of the application template. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1731469200
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListApplicationTemplateResponseBodyApplicationTemplates) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationTemplateResponseBodyApplicationTemplates) GoString() string {
	return s.String()
}

func (s *ListApplicationTemplateResponseBodyApplicationTemplates) GetApplicationTemplateId() *string {
	return s.ApplicationTemplateId
}

func (s *ListApplicationTemplateResponseBodyApplicationTemplates) GetApplicationTemplateName() *string {
	return s.ApplicationTemplateName
}

func (s *ListApplicationTemplateResponseBodyApplicationTemplates) GetBuiltInApplicationActionStatistics() *ListApplicationTemplateResponseBodyApplicationTemplatesBuiltInApplicationActionStatistics {
	return s.BuiltInApplicationActionStatistics
}

func (s *ListApplicationTemplateResponseBodyApplicationTemplates) GetBuiltInApplicationStatus() *string {
	return s.BuiltInApplicationStatus
}

func (s *ListApplicationTemplateResponseBodyApplicationTemplates) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListApplicationTemplateResponseBodyApplicationTemplates) GetCustomRuleCount() *int64 {
	return s.CustomRuleCount
}

func (s *ListApplicationTemplateResponseBodyApplicationTemplates) GetCustomRuleStatus() *string {
	return s.CustomRuleStatus
}

func (s *ListApplicationTemplateResponseBodyApplicationTemplates) GetDescription() *string {
	return s.Description
}

func (s *ListApplicationTemplateResponseBodyApplicationTemplates) GetTemplateStatus() *string {
	return s.TemplateStatus
}

func (s *ListApplicationTemplateResponseBodyApplicationTemplates) GetUnknownAction() *string {
	return s.UnknownAction
}

func (s *ListApplicationTemplateResponseBodyApplicationTemplates) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListApplicationTemplateResponseBodyApplicationTemplates) SetApplicationTemplateId(v string) *ListApplicationTemplateResponseBodyApplicationTemplates {
	s.ApplicationTemplateId = &v
	return s
}

func (s *ListApplicationTemplateResponseBodyApplicationTemplates) SetApplicationTemplateName(v string) *ListApplicationTemplateResponseBodyApplicationTemplates {
	s.ApplicationTemplateName = &v
	return s
}

func (s *ListApplicationTemplateResponseBodyApplicationTemplates) SetBuiltInApplicationActionStatistics(v *ListApplicationTemplateResponseBodyApplicationTemplatesBuiltInApplicationActionStatistics) *ListApplicationTemplateResponseBodyApplicationTemplates {
	s.BuiltInApplicationActionStatistics = v
	return s
}

func (s *ListApplicationTemplateResponseBodyApplicationTemplates) SetBuiltInApplicationStatus(v string) *ListApplicationTemplateResponseBodyApplicationTemplates {
	s.BuiltInApplicationStatus = &v
	return s
}

func (s *ListApplicationTemplateResponseBodyApplicationTemplates) SetCreateTime(v int64) *ListApplicationTemplateResponseBodyApplicationTemplates {
	s.CreateTime = &v
	return s
}

func (s *ListApplicationTemplateResponseBodyApplicationTemplates) SetCustomRuleCount(v int64) *ListApplicationTemplateResponseBodyApplicationTemplates {
	s.CustomRuleCount = &v
	return s
}

func (s *ListApplicationTemplateResponseBodyApplicationTemplates) SetCustomRuleStatus(v string) *ListApplicationTemplateResponseBodyApplicationTemplates {
	s.CustomRuleStatus = &v
	return s
}

func (s *ListApplicationTemplateResponseBodyApplicationTemplates) SetDescription(v string) *ListApplicationTemplateResponseBodyApplicationTemplates {
	s.Description = &v
	return s
}

func (s *ListApplicationTemplateResponseBodyApplicationTemplates) SetTemplateStatus(v string) *ListApplicationTemplateResponseBodyApplicationTemplates {
	s.TemplateStatus = &v
	return s
}

func (s *ListApplicationTemplateResponseBodyApplicationTemplates) SetUnknownAction(v string) *ListApplicationTemplateResponseBodyApplicationTemplates {
	s.UnknownAction = &v
	return s
}

func (s *ListApplicationTemplateResponseBodyApplicationTemplates) SetUpdateTime(v int64) *ListApplicationTemplateResponseBodyApplicationTemplates {
	s.UpdateTime = &v
	return s
}

func (s *ListApplicationTemplateResponseBodyApplicationTemplates) Validate() error {
	if s.BuiltInApplicationActionStatistics != nil {
		if err := s.BuiltInApplicationActionStatistics.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListApplicationTemplateResponseBodyApplicationTemplatesBuiltInApplicationActionStatistics struct {
	// The number of built-in applications whose final effective action is Alert.
	//
	// example:
	//
	// 5
	AlertApplicationCount *int64 `json:"AlertApplicationCount,omitempty" xml:"AlertApplicationCount,omitempty"`
	// The number of built-in applications whose final effective action is Allow.
	//
	// example:
	//
	// 10
	AllowApplicationCount *int64 `json:"AllowApplicationCount,omitempty" xml:"AllowApplicationCount,omitempty"`
	// The number of built-in applications whose final effective action is Deny.
	//
	// example:
	//
	// 2
	DenyApplicationCount *int64 `json:"DenyApplicationCount,omitempty" xml:"DenyApplicationCount,omitempty"`
}

func (s ListApplicationTemplateResponseBodyApplicationTemplatesBuiltInApplicationActionStatistics) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationTemplateResponseBodyApplicationTemplatesBuiltInApplicationActionStatistics) GoString() string {
	return s.String()
}

func (s *ListApplicationTemplateResponseBodyApplicationTemplatesBuiltInApplicationActionStatistics) GetAlertApplicationCount() *int64 {
	return s.AlertApplicationCount
}

func (s *ListApplicationTemplateResponseBodyApplicationTemplatesBuiltInApplicationActionStatistics) GetAllowApplicationCount() *int64 {
	return s.AllowApplicationCount
}

func (s *ListApplicationTemplateResponseBodyApplicationTemplatesBuiltInApplicationActionStatistics) GetDenyApplicationCount() *int64 {
	return s.DenyApplicationCount
}

func (s *ListApplicationTemplateResponseBodyApplicationTemplatesBuiltInApplicationActionStatistics) SetAlertApplicationCount(v int64) *ListApplicationTemplateResponseBodyApplicationTemplatesBuiltInApplicationActionStatistics {
	s.AlertApplicationCount = &v
	return s
}

func (s *ListApplicationTemplateResponseBodyApplicationTemplatesBuiltInApplicationActionStatistics) SetAllowApplicationCount(v int64) *ListApplicationTemplateResponseBodyApplicationTemplatesBuiltInApplicationActionStatistics {
	s.AllowApplicationCount = &v
	return s
}

func (s *ListApplicationTemplateResponseBodyApplicationTemplatesBuiltInApplicationActionStatistics) SetDenyApplicationCount(v int64) *ListApplicationTemplateResponseBodyApplicationTemplatesBuiltInApplicationActionStatistics {
	s.DenyApplicationCount = &v
	return s
}

func (s *ListApplicationTemplateResponseBodyApplicationTemplatesBuiltInApplicationActionStatistics) Validate() error {
	return dara.Validate(s)
}
