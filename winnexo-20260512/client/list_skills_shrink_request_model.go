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
	// 绑定状态：BOUND(已绑定) / UNBOUND(未绑定的全局技能)；必须与 operatingObjectName 同时传入
	//
	// example:
	//
	// BOUND
	BindStatus *string `json:"bindStatus,omitempty" xml:"bindStatus,omitempty"`
	// 技能筛选维度：ALL/BUILTIN/CUSTOM/DRAFT/ALL_WITH_DRAFTS
	//
	// example:
	//
	// ALL
	FilterType *string `json:"filterType,omitempty" xml:"filterType,omitempty"`
	// 按技能名称或描述模糊匹配
	//
	// example:
	//
	// 示例关键词
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 数字员工名称；必须与 bindStatus 同时传入
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// 页码，从 1 开始
	//
	// example:
	//
	// string_value
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页数量，范围 1-100
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 按标签过滤，数组任一命中即匹配
	//
	// example:
	//
	// string_value
	TagsShrink *string `json:"tags,omitempty" xml:"tags,omitempty"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
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
