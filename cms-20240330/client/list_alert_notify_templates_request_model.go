// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertNotifyTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertNotifyTemplateIds(v []*string) *ListAlertNotifyTemplatesRequest
	GetAlertNotifyTemplateIds() []*string
	SetAlertNotifyTemplateName(v string) *ListAlertNotifyTemplatesRequest
	GetAlertNotifyTemplateName() *string
	SetPageNumber(v int64) *ListAlertNotifyTemplatesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListAlertNotifyTemplatesRequest
	GetPageSize() *int64
	SetProgramLang(v string) *ListAlertNotifyTemplatesRequest
	GetProgramLang() *string
	SetType(v string) *ListAlertNotifyTemplatesRequest
	GetType() *string
}

type ListAlertNotifyTemplatesRequest struct {
	AlertNotifyTemplateIds []*string `json:"alertNotifyTemplateIds,omitempty" xml:"alertNotifyTemplateIds,omitempty" type:"Repeated"`
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

func (s ListAlertNotifyTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAlertNotifyTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListAlertNotifyTemplatesRequest) GetAlertNotifyTemplateIds() []*string {
	return s.AlertNotifyTemplateIds
}

func (s *ListAlertNotifyTemplatesRequest) GetAlertNotifyTemplateName() *string {
	return s.AlertNotifyTemplateName
}

func (s *ListAlertNotifyTemplatesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListAlertNotifyTemplatesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAlertNotifyTemplatesRequest) GetProgramLang() *string {
	return s.ProgramLang
}

func (s *ListAlertNotifyTemplatesRequest) GetType() *string {
	return s.Type
}

func (s *ListAlertNotifyTemplatesRequest) SetAlertNotifyTemplateIds(v []*string) *ListAlertNotifyTemplatesRequest {
	s.AlertNotifyTemplateIds = v
	return s
}

func (s *ListAlertNotifyTemplatesRequest) SetAlertNotifyTemplateName(v string) *ListAlertNotifyTemplatesRequest {
	s.AlertNotifyTemplateName = &v
	return s
}

func (s *ListAlertNotifyTemplatesRequest) SetPageNumber(v int64) *ListAlertNotifyTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAlertNotifyTemplatesRequest) SetPageSize(v int64) *ListAlertNotifyTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAlertNotifyTemplatesRequest) SetProgramLang(v string) *ListAlertNotifyTemplatesRequest {
	s.ProgramLang = &v
	return s
}

func (s *ListAlertNotifyTemplatesRequest) SetType(v string) *ListAlertNotifyTemplatesRequest {
	s.Type = &v
	return s
}

func (s *ListAlertNotifyTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
