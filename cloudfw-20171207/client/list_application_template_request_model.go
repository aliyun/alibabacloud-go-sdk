// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationTemplateIds(v []*string) *ListApplicationTemplateRequest
	GetApplicationTemplateIds() []*string
	SetApplicationTemplateName(v string) *ListApplicationTemplateRequest
	GetApplicationTemplateName() *string
	SetBuiltInApplicationStatus(v string) *ListApplicationTemplateRequest
	GetBuiltInApplicationStatus() *string
	SetCustomRuleStatus(v string) *ListApplicationTemplateRequest
	GetCustomRuleStatus() *string
	SetDescription(v string) *ListApplicationTemplateRequest
	GetDescription() *string
	SetPageNo(v int64) *ListApplicationTemplateRequest
	GetPageNo() *int64
	SetPageSize(v int64) *ListApplicationTemplateRequest
	GetPageSize() *int64
	SetTemplateStatus(v string) *ListApplicationTemplateRequest
	GetTemplateStatus() *string
	SetUnknownAction(v string) *ListApplicationTemplateRequest
	GetUnknownAction() *string
}

type ListApplicationTemplateRequest struct {
	// The list of application template IDs for batch exact matching.
	ApplicationTemplateIds []*string `json:"ApplicationTemplateIds,omitempty" xml:"ApplicationTemplateIds,omitempty" type:"Repeated"`
	// The application template name for fuzzy matching.
	//
	// example:
	//
	// Application template
	ApplicationTemplateName *string `json:"ApplicationTemplateName,omitempty" xml:"ApplicationTemplateName,omitempty"`
	// The switch status of built-in application control for exact matching. Valid values:
	//
	// - **Enable**: enabled
	//
	// - **Disable**: disabled
	//
	// example:
	//
	// Enable
	BuiltInApplicationStatus *string `json:"BuiltInApplicationStatus,omitempty" xml:"BuiltInApplicationStatus,omitempty"`
	// The switch status of custom application rules for exact matching. Valid values:
	//
	// - **Enable**: enabled
	//
	// - **Disable**: disabled
	//
	// example:
	//
	// Enable
	CustomRuleStatus *string `json:"CustomRuleStatus,omitempty" xml:"CustomRuleStatus,omitempty"`
	// The description of the application template for fuzzy matching.
	//
	// example:
	//
	// Office network application control
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The page number. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. The value must be greater than or equal to 1. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The master switch status of the application template for exact matching. Valid values:
	//
	// - **Enable**: enabled
	//
	// - **Disable**: disabled
	//
	// example:
	//
	// Enable
	TemplateStatus *string `json:"TemplateStatus,omitempty" xml:"TemplateStatus,omitempty"`
	// The action for unrecognized applications for exact matching. Valid values:
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
}

func (s ListApplicationTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationTemplateRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationTemplateRequest) GetApplicationTemplateIds() []*string {
	return s.ApplicationTemplateIds
}

func (s *ListApplicationTemplateRequest) GetApplicationTemplateName() *string {
	return s.ApplicationTemplateName
}

func (s *ListApplicationTemplateRequest) GetBuiltInApplicationStatus() *string {
	return s.BuiltInApplicationStatus
}

func (s *ListApplicationTemplateRequest) GetCustomRuleStatus() *string {
	return s.CustomRuleStatus
}

func (s *ListApplicationTemplateRequest) GetDescription() *string {
	return s.Description
}

func (s *ListApplicationTemplateRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListApplicationTemplateRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListApplicationTemplateRequest) GetTemplateStatus() *string {
	return s.TemplateStatus
}

func (s *ListApplicationTemplateRequest) GetUnknownAction() *string {
	return s.UnknownAction
}

func (s *ListApplicationTemplateRequest) SetApplicationTemplateIds(v []*string) *ListApplicationTemplateRequest {
	s.ApplicationTemplateIds = v
	return s
}

func (s *ListApplicationTemplateRequest) SetApplicationTemplateName(v string) *ListApplicationTemplateRequest {
	s.ApplicationTemplateName = &v
	return s
}

func (s *ListApplicationTemplateRequest) SetBuiltInApplicationStatus(v string) *ListApplicationTemplateRequest {
	s.BuiltInApplicationStatus = &v
	return s
}

func (s *ListApplicationTemplateRequest) SetCustomRuleStatus(v string) *ListApplicationTemplateRequest {
	s.CustomRuleStatus = &v
	return s
}

func (s *ListApplicationTemplateRequest) SetDescription(v string) *ListApplicationTemplateRequest {
	s.Description = &v
	return s
}

func (s *ListApplicationTemplateRequest) SetPageNo(v int64) *ListApplicationTemplateRequest {
	s.PageNo = &v
	return s
}

func (s *ListApplicationTemplateRequest) SetPageSize(v int64) *ListApplicationTemplateRequest {
	s.PageSize = &v
	return s
}

func (s *ListApplicationTemplateRequest) SetTemplateStatus(v string) *ListApplicationTemplateRequest {
	s.TemplateStatus = &v
	return s
}

func (s *ListApplicationTemplateRequest) SetUnknownAction(v string) *ListApplicationTemplateRequest {
	s.UnknownAction = &v
	return s
}

func (s *ListApplicationTemplateRequest) Validate() error {
	return dara.Validate(s)
}
