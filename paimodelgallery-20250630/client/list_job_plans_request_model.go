// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobPlansRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHasTemplate(v bool) *ListJobPlansRequest
	GetHasTemplate() *bool
	SetJobPlanName(v string) *ListJobPlansRequest
	GetJobPlanName() *string
	SetJobPlanType(v string) *ListJobPlansRequest
	GetJobPlanType() *string
	SetOrder(v string) *ListJobPlansRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListJobPlansRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListJobPlansRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListJobPlansRequest
	GetSortBy() *string
	SetTag(v []*ListJobPlansRequestTag) *ListJobPlansRequest
	GetTag() []*ListJobPlansRequestTag
	SetTemplateId(v string) *ListJobPlansRequest
	GetTemplateId() *string
	SetWorkspaceId(v string) *ListJobPlansRequest
	GetWorkspaceId() *string
}

type ListJobPlansRequest struct {
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
	Tag []*ListJobPlansRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s ListJobPlansRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobPlansRequest) GoString() string {
	return s.String()
}

func (s *ListJobPlansRequest) GetHasTemplate() *bool {
	return s.HasTemplate
}

func (s *ListJobPlansRequest) GetJobPlanName() *string {
	return s.JobPlanName
}

func (s *ListJobPlansRequest) GetJobPlanType() *string {
	return s.JobPlanType
}

func (s *ListJobPlansRequest) GetOrder() *string {
	return s.Order
}

func (s *ListJobPlansRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListJobPlansRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListJobPlansRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListJobPlansRequest) GetTag() []*ListJobPlansRequestTag {
	return s.Tag
}

func (s *ListJobPlansRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListJobPlansRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListJobPlansRequest) SetHasTemplate(v bool) *ListJobPlansRequest {
	s.HasTemplate = &v
	return s
}

func (s *ListJobPlansRequest) SetJobPlanName(v string) *ListJobPlansRequest {
	s.JobPlanName = &v
	return s
}

func (s *ListJobPlansRequest) SetJobPlanType(v string) *ListJobPlansRequest {
	s.JobPlanType = &v
	return s
}

func (s *ListJobPlansRequest) SetOrder(v string) *ListJobPlansRequest {
	s.Order = &v
	return s
}

func (s *ListJobPlansRequest) SetPageNumber(v int32) *ListJobPlansRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobPlansRequest) SetPageSize(v int32) *ListJobPlansRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobPlansRequest) SetSortBy(v string) *ListJobPlansRequest {
	s.SortBy = &v
	return s
}

func (s *ListJobPlansRequest) SetTag(v []*ListJobPlansRequestTag) *ListJobPlansRequest {
	s.Tag = v
	return s
}

func (s *ListJobPlansRequest) SetTemplateId(v string) *ListJobPlansRequest {
	s.TemplateId = &v
	return s
}

func (s *ListJobPlansRequest) SetWorkspaceId(v string) *ListJobPlansRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListJobPlansRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListJobPlansRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListJobPlansRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListJobPlansRequestTag) GoString() string {
	return s.String()
}

func (s *ListJobPlansRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListJobPlansRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListJobPlansRequestTag) SetKey(v string) *ListJobPlansRequestTag {
	s.Key = &v
	return s
}

func (s *ListJobPlansRequestTag) SetValue(v string) *ListJobPlansRequestTag {
	s.Value = &v
	return s
}

func (s *ListJobPlansRequestTag) Validate() error {
	return dara.Validate(s)
}
