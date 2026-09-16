// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListSkillResponseBodyData) *ListSkillResponseBody
	GetData() []*ListSkillResponseBodyData
	SetPageNumber(v int64) *ListSkillResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListSkillResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListSkillResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListSkillResponseBody
	GetTotalCount() *int32
}

type ListSkillResponseBody struct {
	// The skill list.
	Data []*ListSkillResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The unique request identifier.
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSkillResponseBody) GoString() string {
	return s.String()
}

func (s *ListSkillResponseBody) GetData() []*ListSkillResponseBodyData {
	return s.Data
}

func (s *ListSkillResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListSkillResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSkillResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSkillResponseBody) SetData(v []*ListSkillResponseBodyData) *ListSkillResponseBody {
	s.Data = v
	return s
}

func (s *ListSkillResponseBody) SetPageNumber(v int64) *ListSkillResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSkillResponseBody) SetPageSize(v int64) *ListSkillResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSkillResponseBody) SetRequestId(v string) *ListSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSkillResponseBody) SetTotalCount(v int32) *ListSkillResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSkillResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSkillResponseBodyData struct {
	// The ID of the currently active version.
	//
	// example:
	//
	// version-example
	ActiveVersionId *string `json:"ActiveVersionId,omitempty" xml:"ActiveVersionId,omitempty"`
	// The skill category.
	//
	// example:
	//
	// productivity
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The data content.
	//
	// example:
	//
	// {"MySQL": "MySQL optimization guide...","PostgreSQL": "PostgreSQL optimization guide..."}
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-02-04T21:14:45Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The list of database types.
	Dbtypes []*string `json:"Dbtypes,omitempty" xml:"Dbtypes,omitempty" type:"Repeated"`
	// The description.
	//
	// example:
	//
	// SQL Review Expert: Comprehensively reviews SQL for security, performance, and compliance, identifies risks, and provides optimization suggestions. Activated immediately when a user submits SQL or asks about "SQL review", "SQL Review", "any risks", or "how to optimize"
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the skill.
	//
	// example:
	//
	// Example Skill
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The public HTTPS URL of the current icon. Empty if not configured.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// https://example.com/skill-icon.png
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// The unique identifier of the skill.
	//
	// example:
	//
	// 9a2ba261-7bb2-41a7-9c6e-1799fb5b****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the skill is deleted.
	IsDeleted *bool `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	// The skill name.
	//
	// example:
	//
	// sql-review
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The visibility scope of the skill.
	//
	// example:
	//
	// PRIVATE
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The skill type.
	//
	// example:
	//
	// system
	SkillType *string `json:"SkillType,omitempty" xml:"SkillType,omitempty"`
	// The stable identifier of the skill.
	//
	// example:
	//
	// example-skill
	Slug *string `json:"Slug,omitempty" xml:"Slug,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2026-02-04T21:14:45Z
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s ListSkillResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSkillResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSkillResponseBodyData) GetActiveVersionId() *string {
	return s.ActiveVersionId
}

func (s *ListSkillResponseBodyData) GetCategory() *string {
	return s.Category
}

func (s *ListSkillResponseBodyData) GetContent() map[string]interface{} {
	return s.Content
}

func (s *ListSkillResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListSkillResponseBodyData) GetDbtypes() []*string {
	return s.Dbtypes
}

func (s *ListSkillResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListSkillResponseBodyData) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListSkillResponseBodyData) GetIcon() *string {
	return s.Icon
}

func (s *ListSkillResponseBodyData) GetId() *string {
	return s.Id
}

func (s *ListSkillResponseBodyData) GetIsDeleted() *bool {
	return s.IsDeleted
}

func (s *ListSkillResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListSkillResponseBodyData) GetScope() *string {
	return s.Scope
}

func (s *ListSkillResponseBodyData) GetSkillType() *string {
	return s.SkillType
}

func (s *ListSkillResponseBodyData) GetSlug() *string {
	return s.Slug
}

func (s *ListSkillResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListSkillResponseBodyData) SetActiveVersionId(v string) *ListSkillResponseBodyData {
	s.ActiveVersionId = &v
	return s
}

func (s *ListSkillResponseBodyData) SetCategory(v string) *ListSkillResponseBodyData {
	s.Category = &v
	return s
}

func (s *ListSkillResponseBodyData) SetContent(v map[string]interface{}) *ListSkillResponseBodyData {
	s.Content = v
	return s
}

func (s *ListSkillResponseBodyData) SetCreatedAt(v string) *ListSkillResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *ListSkillResponseBodyData) SetDbtypes(v []*string) *ListSkillResponseBodyData {
	s.Dbtypes = v
	return s
}

func (s *ListSkillResponseBodyData) SetDescription(v string) *ListSkillResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListSkillResponseBodyData) SetDisplayName(v string) *ListSkillResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *ListSkillResponseBodyData) SetIcon(v string) *ListSkillResponseBodyData {
	s.Icon = &v
	return s
}

func (s *ListSkillResponseBodyData) SetId(v string) *ListSkillResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListSkillResponseBodyData) SetIsDeleted(v bool) *ListSkillResponseBodyData {
	s.IsDeleted = &v
	return s
}

func (s *ListSkillResponseBodyData) SetName(v string) *ListSkillResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListSkillResponseBodyData) SetScope(v string) *ListSkillResponseBodyData {
	s.Scope = &v
	return s
}

func (s *ListSkillResponseBodyData) SetSkillType(v string) *ListSkillResponseBodyData {
	s.SkillType = &v
	return s
}

func (s *ListSkillResponseBodyData) SetSlug(v string) *ListSkillResponseBodyData {
	s.Slug = &v
	return s
}

func (s *ListSkillResponseBodyData) SetUpdatedAt(v string) *ListSkillResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *ListSkillResponseBodyData) Validate() error {
	return dara.Validate(s)
}
