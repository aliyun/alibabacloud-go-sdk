// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertNotifyTemplatesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertNotifyTemplateIdsShrink(v string) *ListAlertNotifyTemplatesShrinkRequest
	GetAlertNotifyTemplateIdsShrink() *string
	SetAlertNotifyTemplateName(v string) *ListAlertNotifyTemplatesShrinkRequest
	GetAlertNotifyTemplateName() *string
	SetPageNumber(v int64) *ListAlertNotifyTemplatesShrinkRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListAlertNotifyTemplatesShrinkRequest
	GetPageSize() *int64
	SetProgramLang(v string) *ListAlertNotifyTemplatesShrinkRequest
	GetProgramLang() *string
	SetType(v string) *ListAlertNotifyTemplatesShrinkRequest
	GetType() *string
}

type ListAlertNotifyTemplatesShrinkRequest struct {
	AlertNotifyTemplateIdsShrink *string `json:"alertNotifyTemplateIds,omitempty" xml:"alertNotifyTemplateIds,omitempty"`
	// example:
	//
	// 告警通知模板
	AlertNotifyTemplateName *string `json:"alertNotifyTemplateName,omitempty" xml:"alertNotifyTemplateName,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// GO
	ProgramLang *string `json:"programLang,omitempty" xml:"programLang,omitempty"`
	// example:
	//
	// DING
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListAlertNotifyTemplatesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAlertNotifyTemplatesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAlertNotifyTemplatesShrinkRequest) GetAlertNotifyTemplateIdsShrink() *string {
	return s.AlertNotifyTemplateIdsShrink
}

func (s *ListAlertNotifyTemplatesShrinkRequest) GetAlertNotifyTemplateName() *string {
	return s.AlertNotifyTemplateName
}

func (s *ListAlertNotifyTemplatesShrinkRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListAlertNotifyTemplatesShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAlertNotifyTemplatesShrinkRequest) GetProgramLang() *string {
	return s.ProgramLang
}

func (s *ListAlertNotifyTemplatesShrinkRequest) GetType() *string {
	return s.Type
}

func (s *ListAlertNotifyTemplatesShrinkRequest) SetAlertNotifyTemplateIdsShrink(v string) *ListAlertNotifyTemplatesShrinkRequest {
	s.AlertNotifyTemplateIdsShrink = &v
	return s
}

func (s *ListAlertNotifyTemplatesShrinkRequest) SetAlertNotifyTemplateName(v string) *ListAlertNotifyTemplatesShrinkRequest {
	s.AlertNotifyTemplateName = &v
	return s
}

func (s *ListAlertNotifyTemplatesShrinkRequest) SetPageNumber(v int64) *ListAlertNotifyTemplatesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAlertNotifyTemplatesShrinkRequest) SetPageSize(v int64) *ListAlertNotifyTemplatesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListAlertNotifyTemplatesShrinkRequest) SetProgramLang(v string) *ListAlertNotifyTemplatesShrinkRequest {
	s.ProgramLang = &v
	return s
}

func (s *ListAlertNotifyTemplatesShrinkRequest) SetType(v string) *ListAlertNotifyTemplatesShrinkRequest {
	s.Type = &v
	return s
}

func (s *ListAlertNotifyTemplatesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
