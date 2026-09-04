// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBindStatus(v string) *ListSkillsRequest
	GetBindStatus() *string
	SetFilterType(v string) *ListSkillsRequest
	GetFilterType() *string
	SetKeyword(v string) *ListSkillsRequest
	GetKeyword() *string
	SetOperatingObjectName(v string) *ListSkillsRequest
	GetOperatingObjectName() *string
	SetPage(v int32) *ListSkillsRequest
	GetPage() *int32
	SetPageSize(v int32) *ListSkillsRequest
	GetPageSize() *int32
	SetTags(v []*string) *ListSkillsRequest
	GetTags() []*string
	SetTenantId(v string) *ListSkillsRequest
	GetTenantId() *string
}

type ListSkillsRequest struct {
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
	Tags []*string `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListSkillsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSkillsRequest) GoString() string {
	return s.String()
}

func (s *ListSkillsRequest) GetBindStatus() *string {
	return s.BindStatus
}

func (s *ListSkillsRequest) GetFilterType() *string {
	return s.FilterType
}

func (s *ListSkillsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListSkillsRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *ListSkillsRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListSkillsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSkillsRequest) GetTags() []*string {
	return s.Tags
}

func (s *ListSkillsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListSkillsRequest) SetBindStatus(v string) *ListSkillsRequest {
	s.BindStatus = &v
	return s
}

func (s *ListSkillsRequest) SetFilterType(v string) *ListSkillsRequest {
	s.FilterType = &v
	return s
}

func (s *ListSkillsRequest) SetKeyword(v string) *ListSkillsRequest {
	s.Keyword = &v
	return s
}

func (s *ListSkillsRequest) SetOperatingObjectName(v string) *ListSkillsRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *ListSkillsRequest) SetPage(v int32) *ListSkillsRequest {
	s.Page = &v
	return s
}

func (s *ListSkillsRequest) SetPageSize(v int32) *ListSkillsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSkillsRequest) SetTags(v []*string) *ListSkillsRequest {
	s.Tags = v
	return s
}

func (s *ListSkillsRequest) SetTenantId(v string) *ListSkillsRequest {
	s.TenantId = &v
	return s
}

func (s *ListSkillsRequest) Validate() error {
	return dara.Validate(s)
}
