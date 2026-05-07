// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrder(v string) *ListJobTemplatesRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListJobTemplatesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListJobTemplatesRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListJobTemplatesRequest
	GetSortBy() *string
	SetTemplateId(v string) *ListJobTemplatesRequest
	GetTemplateId() *string
	SetTemplateName(v string) *ListJobTemplatesRequest
	GetTemplateName() *string
	SetUserId(v string) *ListJobTemplatesRequest
	GetUserId() *string
	SetWorkspaceId(v string) *ListJobTemplatesRequest
	GetWorkspaceId() *string
}

type ListJobTemplatesRequest struct {
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// 按模板 ID 精确筛选。
	//
	// example:
	//
	// tpl1****6jcq2q
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 按模板名称模糊筛选。
	//
	// example:
	//
	// job-template-example-1778047****
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// 按创建者用户 ID 筛选。
	//
	// example:
	//
	// 20**************02
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 工作空间 ID。如何获取工作空间 ID，请参见 ListWorkspaces。
	//
	// This parameter is required.
	//
	// example:
	//
	// 4***9
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListJobTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListJobTemplatesRequest) GetOrder() *string {
	return s.Order
}

func (s *ListJobTemplatesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListJobTemplatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListJobTemplatesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListJobTemplatesRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListJobTemplatesRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListJobTemplatesRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListJobTemplatesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListJobTemplatesRequest) SetOrder(v string) *ListJobTemplatesRequest {
	s.Order = &v
	return s
}

func (s *ListJobTemplatesRequest) SetPageNumber(v int32) *ListJobTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobTemplatesRequest) SetPageSize(v int32) *ListJobTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobTemplatesRequest) SetSortBy(v string) *ListJobTemplatesRequest {
	s.SortBy = &v
	return s
}

func (s *ListJobTemplatesRequest) SetTemplateId(v string) *ListJobTemplatesRequest {
	s.TemplateId = &v
	return s
}

func (s *ListJobTemplatesRequest) SetTemplateName(v string) *ListJobTemplatesRequest {
	s.TemplateName = &v
	return s
}

func (s *ListJobTemplatesRequest) SetUserId(v string) *ListJobTemplatesRequest {
	s.UserId = &v
	return s
}

func (s *ListJobTemplatesRequest) SetWorkspaceId(v string) *ListJobTemplatesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListJobTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
