// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobPlansShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHasTemplate(v bool) *ListJobPlansShrinkRequest
	GetHasTemplate() *bool
	SetJobPlanName(v string) *ListJobPlansShrinkRequest
	GetJobPlanName() *string
	SetJobPlanType(v string) *ListJobPlansShrinkRequest
	GetJobPlanType() *string
	SetOrder(v string) *ListJobPlansShrinkRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListJobPlansShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListJobPlansShrinkRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListJobPlansShrinkRequest
	GetSortBy() *string
	SetTagShrink(v string) *ListJobPlansShrinkRequest
	GetTagShrink() *string
	SetTemplateId(v string) *ListJobPlansShrinkRequest
	GetTemplateId() *string
	SetWorkspaceId(v string) *ListJobPlansShrinkRequest
	GetWorkspaceId() *string
}

type ListJobPlansShrinkRequest struct {
	// Specifies whether to filter by template association. Valid values:
	//
	// - true: Returns only scenario-specific job plans that have a template.
	//
	// - false: Returns only general-purpose job plans that do not have a template.
	//
	// If this parameter is not specified, no filtering is applied. If both this parameter and TemplateId are specified, the value of TemplateId takes precedence.
	//
	// example:
	//
	// true
	HasTemplate *bool `json:"HasTemplate,omitempty" xml:"HasTemplate,omitempty"`
	// The name of the job plan.
	//
	// example:
	//
	// JobPlan1
	JobPlanName *string `json:"JobPlanName,omitempty" xml:"JobPlanName,omitempty"`
	// The type of the job plan.
	//
	// example:
	//
	// Distillation
	JobPlanType *string `json:"JobPlanType,omitempty" xml:"JobPlanType,omitempty"`
	// The sort order.
	//
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The field by which to sort the results.
	//
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The list of tags.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The distillation template ID. Filters results to return only scenario-specific tasks that use the specified template.
	//
	// example:
	//
	// advanced_cot_distill
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 32495
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListJobPlansShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobPlansShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListJobPlansShrinkRequest) GetHasTemplate() *bool {
	return s.HasTemplate
}

func (s *ListJobPlansShrinkRequest) GetJobPlanName() *string {
	return s.JobPlanName
}

func (s *ListJobPlansShrinkRequest) GetJobPlanType() *string {
	return s.JobPlanType
}

func (s *ListJobPlansShrinkRequest) GetOrder() *string {
	return s.Order
}

func (s *ListJobPlansShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListJobPlansShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListJobPlansShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListJobPlansShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *ListJobPlansShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListJobPlansShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListJobPlansShrinkRequest) SetHasTemplate(v bool) *ListJobPlansShrinkRequest {
	s.HasTemplate = &v
	return s
}

func (s *ListJobPlansShrinkRequest) SetJobPlanName(v string) *ListJobPlansShrinkRequest {
	s.JobPlanName = &v
	return s
}

func (s *ListJobPlansShrinkRequest) SetJobPlanType(v string) *ListJobPlansShrinkRequest {
	s.JobPlanType = &v
	return s
}

func (s *ListJobPlansShrinkRequest) SetOrder(v string) *ListJobPlansShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListJobPlansShrinkRequest) SetPageNumber(v int32) *ListJobPlansShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobPlansShrinkRequest) SetPageSize(v int32) *ListJobPlansShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobPlansShrinkRequest) SetSortBy(v string) *ListJobPlansShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListJobPlansShrinkRequest) SetTagShrink(v string) *ListJobPlansShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *ListJobPlansShrinkRequest) SetTemplateId(v string) *ListJobPlansShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *ListJobPlansShrinkRequest) SetWorkspaceId(v string) *ListJobPlansShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListJobPlansShrinkRequest) Validate() error {
	return dara.Validate(s)
}
