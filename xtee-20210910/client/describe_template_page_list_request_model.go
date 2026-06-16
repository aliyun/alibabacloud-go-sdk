// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTemplatePageListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeTemplatePageListRequest
	GetLang() *string
	SetCurrentPage(v string) *DescribeTemplatePageListRequest
	GetCurrentPage() *string
	SetEventCodes(v string) *DescribeTemplatePageListRequest
	GetEventCodes() *string
	SetPageSize(v string) *DescribeTemplatePageListRequest
	GetPageSize() *string
	SetRegId(v string) *DescribeTemplatePageListRequest
	GetRegId() *string
	SetTemplateName(v string) *DescribeTemplatePageListRequest
	GetTemplateName() *string
	SetTemplateSearchItem(v string) *DescribeTemplatePageListRequest
	GetTemplateSearchItem() *string
	SetTemplateStatus(v string) *DescribeTemplatePageListRequest
	GetTemplateStatus() *string
	SetTemplateType(v string) *DescribeTemplatePageListRequest
	GetTemplateType() *string
}

type DescribeTemplatePageListRequest struct {
	// The language type for the request and response messages. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// The event code.
	//
	// example:
	//
	// d6_h1fe4cfa1g
	EventCodes *string `json:"eventCodes,omitempty" xml:"eventCodes,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// The event name.
	//
	// example:
	//
	// 注册事件
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
	// The search field for event templates.
	//
	// example:
	//
	// age
	TemplateSearchItem *string `json:"templateSearchItem,omitempty" xml:"templateSearchItem,omitempty"`
	// The event status.
	//
	// example:
	//
	// ONLINE
	TemplateStatus *string `json:"templateStatus,omitempty" xml:"templateStatus,omitempty"`
	// The templatetype.
	//
	// example:
	//
	// PUB_SERVICE
	TemplateType *string `json:"templateType,omitempty" xml:"templateType,omitempty"`
}

func (s DescribeTemplatePageListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplatePageListRequest) GoString() string {
	return s.String()
}

func (s *DescribeTemplatePageListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeTemplatePageListRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeTemplatePageListRequest) GetEventCodes() *string {
	return s.EventCodes
}

func (s *DescribeTemplatePageListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeTemplatePageListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeTemplatePageListRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DescribeTemplatePageListRequest) GetTemplateSearchItem() *string {
	return s.TemplateSearchItem
}

func (s *DescribeTemplatePageListRequest) GetTemplateStatus() *string {
	return s.TemplateStatus
}

func (s *DescribeTemplatePageListRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *DescribeTemplatePageListRequest) SetLang(v string) *DescribeTemplatePageListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeTemplatePageListRequest) SetCurrentPage(v string) *DescribeTemplatePageListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeTemplatePageListRequest) SetEventCodes(v string) *DescribeTemplatePageListRequest {
	s.EventCodes = &v
	return s
}

func (s *DescribeTemplatePageListRequest) SetPageSize(v string) *DescribeTemplatePageListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTemplatePageListRequest) SetRegId(v string) *DescribeTemplatePageListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeTemplatePageListRequest) SetTemplateName(v string) *DescribeTemplatePageListRequest {
	s.TemplateName = &v
	return s
}

func (s *DescribeTemplatePageListRequest) SetTemplateSearchItem(v string) *DescribeTemplatePageListRequest {
	s.TemplateSearchItem = &v
	return s
}

func (s *DescribeTemplatePageListRequest) SetTemplateStatus(v string) *DescribeTemplatePageListRequest {
	s.TemplateStatus = &v
	return s
}

func (s *DescribeTemplatePageListRequest) SetTemplateType(v string) *DescribeTemplatePageListRequest {
	s.TemplateType = &v
	return s
}

func (s *DescribeTemplatePageListRequest) Validate() error {
	return dara.Validate(s)
}
