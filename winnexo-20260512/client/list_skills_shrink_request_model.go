// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBindStatus(v string) *ListSkillsShrinkRequest
	GetBindStatus() *string
	SetFilterType(v string) *ListSkillsShrinkRequest
	GetFilterType() *string
	SetKeyword(v string) *ListSkillsShrinkRequest
	GetKeyword() *string
	SetOperatingObjectName(v string) *ListSkillsShrinkRequest
	GetOperatingObjectName() *string
	SetPage(v int32) *ListSkillsShrinkRequest
	GetPage() *int32
	SetPageSize(v int32) *ListSkillsShrinkRequest
	GetPageSize() *int32
	SetTagsShrink(v string) *ListSkillsShrinkRequest
	GetTagsShrink() *string
	SetTenantId(v string) *ListSkillsShrinkRequest
	GetTenantId() *string
}

type ListSkillsShrinkRequest struct {
	// The binding status. Valid values: BOUND (bound) and UNBOUND (unbound global skills). Must be specified together with operatingObjectName.
	//
	// example:
	//
	// BOUND
	BindStatus *string `json:"bindStatus,omitempty" xml:"bindStatus,omitempty"`
	// The filter expression type.
	//
	// - SQL: SQL-based filtering.
	//
	// - TAG: Tag-based filtering.
	//
	// example:
	//
	// ALL
	FilterType *string `json:"filterType,omitempty" xml:"filterType,omitempty"`
	// The search keyword. Supports fuzzy search by API name or exact search by API ID.
	//
	// example:
	//
	// SampleKeyword
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The digital employee name. Used to calculate the CodeAgent allowedSkills whitelist based on binding relationships.
	//
	// example:
	//
	// 11111
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The tag filtering parameter.
	//
	// example:
	//
	// string_value
	TagsShrink *string `json:"tags,omitempty" xml:"tags,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListSkillsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSkillsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListSkillsShrinkRequest) GetBindStatus() *string {
	return s.BindStatus
}

func (s *ListSkillsShrinkRequest) GetFilterType() *string {
	return s.FilterType
}

func (s *ListSkillsShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListSkillsShrinkRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *ListSkillsShrinkRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListSkillsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSkillsShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *ListSkillsShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListSkillsShrinkRequest) SetBindStatus(v string) *ListSkillsShrinkRequest {
	s.BindStatus = &v
	return s
}

func (s *ListSkillsShrinkRequest) SetFilterType(v string) *ListSkillsShrinkRequest {
	s.FilterType = &v
	return s
}

func (s *ListSkillsShrinkRequest) SetKeyword(v string) *ListSkillsShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListSkillsShrinkRequest) SetOperatingObjectName(v string) *ListSkillsShrinkRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *ListSkillsShrinkRequest) SetPage(v int32) *ListSkillsShrinkRequest {
	s.Page = &v
	return s
}

func (s *ListSkillsShrinkRequest) SetPageSize(v int32) *ListSkillsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListSkillsShrinkRequest) SetTagsShrink(v string) *ListSkillsShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListSkillsShrinkRequest) SetTenantId(v string) *ListSkillsShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *ListSkillsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
